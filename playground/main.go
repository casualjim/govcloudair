package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/vmware/govcloudair/schemas/vcloud"
)

var (
	// DumpInOut dumps the input and output of HTTP calls to the console
	DumpInOut = true
	// BaseURL for vCloud Air
	BaseURL = "https://vca.vmware.com/api"
	// LoginPath for vCloud Air
	LoginPath = "/iam/login"
	// InstancesPath for vCloud Air
	InstancesPath = "/sc/instances"

	// HeaderAccept the HTTP accept header key
	HeaderAccept = "Accept"
)

var (
	vcaUsername string = os.Getenv("VCLOUDAIR_USERNAME")
	vcaPassword string = os.Getenv("VCLOUDAIR_PASSWORD")
)

func main() {
	client, err := newAuthenticatedClient(vcaUsername, vcaPassword)
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.JSONRequest(vcloud.HTTPGet, InstancesPath, &client.Info); err != nil {
		log.Fatalln(err)
	}

	var attrs *accountInstanceAttrs
	for _, inst := range client.Info.Instances {
		attrs = inst.Attrs()
		if attrs != nil {
			break
		}
	}

	ses, err := attrs.Authenticate(vcaUsername, vcaPassword)
	if err != nil {
		log.Fatalln(err)
	}

	orgList, err := ses.OrgList()
	if err != nil {
		log.Fatalln(err)
	}

	org, err := orgList.FirstOrg(ses)
	if err != nil {
		log.Fatalln(err)
	}

	catalog, err := org.RetrievePublicCatalog(ses)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = catalog.ItemForName("VMware Photon OS - Tech Preview 2", ses)
	if err != nil {
		log.Fatalln(err)
	}

	vdc, err := org.FindVDC("VDC1", ses)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.MarshalIndent(vdc, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

}

var (
	apiLoginResponse = `{"serviceGroupIds":["e7434542-25c7-4120-88ec-6038a5770828"]}`
)

func newAuthenticatedClient(user, password string) (*authenticatedClient, error) {
	r, _ := http.NewRequest(vcloud.HTTPPost, BaseURL+LoginPath, nil)
	r.Header.Set(HeaderAccept, vcloud.JSONMimeV57)
	r.SetBasicAuth(user, password)

	if DumpInOut {
		dr, _ := httputil.DumpRequestOut(r, false)
		fmt.Println(string(dr))
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if DumpInOut {
		dr, _ := httputil.DumpResponse(resp, false)
		fmt.Println(string(dr))
	}

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("Could not complete request with vca, because (status %d) %s\n", resp.StatusCode, resp.Status)
	}

	var result authenticatedClient
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&result); err != nil {
		return nil, err
	}
	result.AuthToken = resp.Header.Get("vchs-authorization")

	return &result, nil
}

type authenticatedClient struct {
	AuthToken       string   `json:"-"`
	ServiceGroupIDs []string `json:"serviceGroupIds"`
	Info            struct {
		Instances []accountInstance `json:"instances"`
	}
	http *http.Client
}

func (a *authenticatedClient) JSONRequest(method, path string, result interface{}) error {
	r, _ := http.NewRequest(method, BaseURL+path, nil)
	r.Header.Set(HeaderAccept, vcloud.JSONMimeV57)

	if a.AuthToken != "" {
		r.Header.Set("Authorization", "Bearer "+a.AuthToken)
	}

	if DumpInOut {
		dr, _ := httputil.DumpRequestOut(r, false)
		fmt.Println(string(dr))
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if DumpInOut {
		dr, _ := httputil.DumpResponse(resp, false)
		fmt.Println(string(dr))
	}

	if resp.StatusCode/100 != 2 {
		log.Fatalf("Could not complete request with vca, because (status %d) %s\n", resp.StatusCode, resp.Status)
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
}

func (a *accountInstanceAttrs) Authenticate(username, password string) (*Session, error) {
	r, _ := http.NewRequest(vcloud.HTTPPost, a.SessionURI, nil)
	r.Header.Set(HeaderAccept, vcloud.AnyXMLMime511)
	r.SetBasicAuth(username+"@"+a.OrgName, password)

	if DumpInOut {
		dr, _ := httputil.DumpRequestOut(r, false)
		fmt.Println(string(dr))
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if DumpInOut {
		dr, _ := httputil.DumpResponse(resp, false)
		fmt.Println(string(dr))
	}

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

	Token string `xml:"-" json:"-"`
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

	if DumpInOut {
		dr, _ := httputil.DumpRequestOut(r, false)
		fmt.Println(string(dr))
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if DumpInOut {
		dr, _ := httputil.DumpResponse(resp, false)
		fmt.Println(string(dr))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	if resp.StatusCode/100 != 2 {
		log.Fatalf("Could not complete request with vca, because (status %d) %s\n", resp.StatusCode, resp.Status)
	}

	dec := xml.NewDecoder(bytes.NewReader(b))
	if err := dec.Decode(result); err != nil {
		return err
	}

	return nil
}
