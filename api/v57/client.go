package v57

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/vmware/govcloudair/schemas/vcloud"
)

const (
	// BaseURL for vCloud Air
	BaseURL = "https://vca.vmware.com/api"
	// LoginPath for vCloud Air
	LoginPath = "/iam/login"
	// InstancesPath for vCloud Air
	InstancesPath = "/sc/instances"
	// HeaderAccept the HTTP accept header key
	HeaderAccept = "Accept"
)

// Config is the client config for the vCloud Air API
type Config struct {
	// Override the default http client
	HTTP *http.Client
	// Username the username to use when authenticating
	Username string
	// Password the username to use when authenticating
	Password string
	// Debug, when true this will dump requests and responses with ALL parameters to the std logger.
	// All parameters also includes things like passwords etc, so be careful when you turn this on for live systems
	// because it's a security hole.
	Debug bool

	// BaseURL is the base url to use when talking to vCloud Air api's. Normal usage would not need to customize this URL.
	// This should be hugely useful in unit tests and stuff.
	BaseURL string
}

// NewAuthenticatedSession create a new vCloud Air authenticated client
func NewAuthenticatedSession(config Config) (interface{}, error) {
	if config.HTTP == nil {
		config.HTTP = http.DefaultClient
	}

	r, _ := http.NewRequest(vcloud.HTTPPost, config.BaseURL+LoginPath, nil)
	r.Header.Set("Accept", vcloud.JSONMimeV57)
	r.SetBasicAuth(config.Username, config.Password)

	if config.Debug {
		dr, _ := httputil.DumpRequestOut(r, true)
		log.Println(string(dr))
	}

	resp, err := config.HTTP.Do(r)
	if err != nil {
		return nil, err
	}
	if config.Debug {
		dr, _ := httputil.DumpResponse(resp, true)
		log.Println(string(dr))
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("Could not complete request with vca, because (status %d) %s\n", resp.StatusCode, resp.Status)
	}

	var result oAuthClient
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&result); err != nil {
		return nil, err
	}
	result.AuthToken = resp.Header.Get("vchs-authorization")
	result.Config = &config

	return result.instances()
}

type oAuthClient struct {
	AuthToken       string   `json:"-"`
	Config          *Config  `json:"-"`
	ServiceGroupIDs []string `json:"serviceGroupIds"`
	Info            struct {
		Instances []accountInstance `json:"instances"`
	}
}

func (a *oAuthClient) instances() ([]accountInstance, error) {
	if err := a.JSONRequest(vcloud.HTTPGet, InstancesPath, &a.Info); err != nil {
		return nil, err
	}
	return a.Info.Instances, nil
}

func (a *oAuthClient) JSONRequest(method, path string, result interface{}) error {
	r, _ := http.NewRequest(method, a.Config.BaseURL+path, nil)
	r.Header.Set(HeaderAccept, vcloud.JSONMimeV57)

	if a.AuthToken != "" {
		r.Header.Set("Authorization", "Bearer "+a.AuthToken)
	}

	if a.Config.Debug {
		dr, _ := httputil.DumpRequestOut(r, false)
		log.Println(string(dr))
	}

	resp, err := a.Config.HTTP.Do(r)
	if err != nil {
		return err
	}
	if a.Config.Debug {
		dr, _ := httputil.DumpResponse(resp, true)
		log.Println(string(dr))
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("Could not complete request with vca, because (status %d) %s\n", resp.StatusCode, resp.Status)
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(result); err != nil {
		return err
	}

	return nil
}

type accountInstance struct {
	APIURL             string   `json:"apiUrl"`
	DashboardURL       string   `json:"dashboardUrl"`
	Description        string   `json:"description"`
	ID                 string   `json:"id"`
	InstanceAttributes string   `json:"instanceAttributes"`
	InstanceVersion    string   `json:"instanceVersion"`
	Link               []string `json:"link"`
	Name               string   `json:"name"`
	PlanID             string   `json:"planId"`
	Region             string   `json:"region"`
	ServiceGroupID     string   `json:"serviceGroupId"`
}

func (a *accountInstance) Attrs() *accountInstanceAttrs {
	if !strings.HasPrefix(a.InstanceAttributes, "{") {
		return nil
	}

	var res accountInstanceAttrs
	if err := json.Unmarshal([]byte(a.InstanceAttributes), &res); err != nil {
		return nil
	}
	return &res
}

type accountInstanceAttrs struct {
	OrgName       string `json:"orgName"`
	SessionURI    string `json:"sessionUri"`
	APIVersionURI string `json:"apiVersionUri"`
	config        Config `json:"-"`
}

func (a *accountInstanceAttrs) Authenticate(username, password string) (*Session, error) {
	r, _ := http.NewRequest(vcloud.HTTPPost, a.SessionURI, nil)
	r.Header.Set(HeaderAccept, vcloud.AnyXMLMime511)
	r.SetBasicAuth(username+"@"+a.OrgName, password)

	if a.config.Debug {
		dr, _ := httputil.DumpRequestOut(r, false)
		log.Println(string(dr))
	}

	resp, err := a.config.HTTP.Do(r)
	if err != nil {
		return nil, err
	}
	if a.config.Debug {
		dr, _ := httputil.DumpResponse(resp, true)
		fmt.Println(string(dr))
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("Could not complete authenticating with vCloud, because (status %d) %s\n", resp.StatusCode, resp.Status)
	}

	var ses Session
	dec := xml.NewDecoder(resp.Body)
	if err := dec.Decode(&ses); err != nil {
		return nil, err
	}
	ses.Token = resp.Header.Get("x-vcloud-authorization")

	return &ses, nil
}

type Session struct {
	// ResourceType
	HREF  string          `xml:"href,attr,omitempty"`
	Type  string          `xml:"type,attr,omitempty"`
	Links vcloud.LinkList `xml:"Link,omitempty"`

	// SessionType
	Org    string `xml:"org,attr,omitempty"`
	Roles  string `xml:"roles,attr,omitempty"`
	User   string `xml:"user,attr,omitempty"`
	UserID string `xml:"userId,attr,omitempty"`

	Token   string `xml:"-" json:"-"`
	context Config `xml:"-" json:"-"`
}

func (s *Session) OrgList() (*vcloud.OrgList, error) {
	lnk := s.Links.ForType(vcloud.MimeOrgList, vcloud.RelDown)
	if lnk == nil {
		return nil, errors.New("Couldn't find the link for orgList")
	}

	var orgList vcloud.OrgList
	if err := s.XMLRequest(vcloud.HTTPGet, lnk.HREF, lnk.Type, nil, &orgList); err != nil {
		return nil, err
	}

	return &orgList, nil
}

func (s *Session) XMLRequest(method, url, tpe string, body, result interface{}) error {

	r, _ := http.NewRequest(method, url, nil)
	if body != nil {
		buf := bytes.NewBuffer(nil)
		enc := xml.NewEncoder(buf)
		if err := enc.Encode(body); err != nil {
			return err
		}
		r, _ = http.NewRequest(method, url, buf)
	}

	r.Header.Set(HeaderAccept, tpe+";version="+vcloud.Version)
	if body != nil {
		r.Header.Set("Content-Type", "application/xml")
	}

	if s.Token != "" {
		r.Header.Set("X-Vcloud-Authorization", s.Token)
	}

	if s.context.Debug {
		dr, _ := httputil.DumpRequestOut(r, true)
		fmt.Println(string(dr))
	}

	resp, err := s.context.HTTP.Do(r)
	if err != nil {
		return err
	}
	if s.context.Debug {
		dr, _ := httputil.DumpResponse(resp, true)
		fmt.Println(string(dr))
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		log.Fatalf("Could not complete request with vca, because (status %d) %s\n", resp.StatusCode, resp.Status)
	}

	dec := xml.NewDecoder(resp.Body)
	if err := dec.Decode(result); err != nil {
		return err
	}

	return nil
}
