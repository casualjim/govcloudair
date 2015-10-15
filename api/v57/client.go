package v57

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/vmware/govcloudair/api"
	"github.com/vmware/govcloudair/schemas/vcloud"
)

const (
	// LoginPath for vCloud Air
	LoginPath = "/iam/login"
	// InstancesPath for vCloud Air
	InstancesPath = "/sc/instances"
	// HeaderAccept the HTTP accept header key
	HeaderAccept = "Accept"
)

// NewAuthenticatedSession create a new vCloud Air authenticated client
func NewAuthenticatedSession(config *api.Config) (*Session, error) {
	if config == nil {
		config = api.DefaultConfig()
	}

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
	result.Config = config

	instances, err := result.instances()
	if err != nil {
		return nil, err
	}

	var attrs *accountInstanceAttrs
	for _, inst := range instances {
		attrs = inst.Attrs()
		if attrs != nil {
			break
		}
	}

	if attrs == nil {
		return nil, fmt.Errorf("unable to determine session url")
	}

	attrs.config = config
	return attrs.Authenticate()
}

type oAuthClient struct {
	AuthToken       string      `json:"-"`
	Config          *api.Config `json:"-"`
	ServiceGroupIDs []string    `json:"serviceGroupIds"`
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
	config        *api.Config
}

func (a *accountInstanceAttrs) Authenticate() (*Session, error) {
	r, _ := http.NewRequest(vcloud.HTTPPost, a.SessionURI, nil)
	r.Header.Set(HeaderAccept, vcloud.AnyXMLMime511)
	r.SetBasicAuth(a.config.Username+"@"+a.OrgName, a.config.Password)

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
	ses.context = a.config
	a.config.Token = resp.Header.Get("x-vcloud-authorization")

	return &ses, nil
}

// Session represents an authenticated session for the vCloud Air API
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

	// token   string
	context *api.Config
}

// OrgList get the org list from the API
func (s *Session) OrgList() (*vcloud.OrgList, error) {
	return vcloud.FetchOrgList(s.Links, s)
}

// XMLRequest makes HTTP request that have XML bodies and get XML results
func (s *Session) XMLRequest(method, url, tpe string, body api.RequestBody, result interface{}) error {
	return api.XMLRequest(s.context, method, url, tpe, body, result)
}
