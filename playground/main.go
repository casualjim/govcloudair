package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
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

	// JSONMimeV57 the json mime for version 5.7 of the API
	JSONMimeV57   = "application/json;version=5.7"
	AnyXMLMime511 = "application/*+xml;version=5.11"
)

var (
	vcaUsername string = os.Getenv("VCLOUDAIR_USERNAME")
	vcaPassword string = os.Getenv("VCLOUDAIR_PASSWORD")
)

func main() {
	client, err := newAuthenticatedClient(vcaUsername, vcaPassword)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.JSONRequest("GET", InstancesPath, &client.Info); err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(ses, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}

var (
	apiLoginResponse = `{"serviceGroupIds":["e7434542-25c7-4120-88ec-6038a5770828"]}`
)

func newAuthenticatedClient(user, password string) (*authenticatedClient, error) {
	r, _ := http.NewRequest("POST", BaseURL+LoginPath, nil)
	r.Header.Set(HeaderAccept, JSONMimeV57)
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
}

func (a *authenticatedClient) JSONRequest(method, path string, result interface{}) error {
	r, _ := http.NewRequest(method, BaseURL+path, nil)
	r.Header.Set(HeaderAccept, JSONMimeV57)

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

func (a *accountInstanceAttrs) Authenticate(username, password string) (*session, error) {
	r, _ := http.NewRequest("POST", a.SessionURI, nil)
	r.Header.Set(HeaderAccept, AnyXMLMime511)
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

	var ses session
	dec := xml.NewDecoder(resp.Body)
	if err := dec.Decode(&ses); err != nil {
		return nil, err
	}
	ses.Token = resp.Header.Get("x-vcloud-authorization")

	return &ses, nil
}

type session struct {
	Links []*Link `xml:"Link,omitempty"`
	Token string  `xml:"-"`
}

// Link extends reference type by adding relation attribute. Defines a hyper-link with a relationship, hyper-link reference, and an optional MIME type.
// Type: LinkType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Extends reference type by adding relation attribute. Defines a hyper-link with a relationship, hyper-link reference, and an optional MIME type.
// Since: 0.9
type Link struct {
	HREF string `xml:"href,attr"`
	ID   string `xml:"id,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	Rel  string `xml:"rel,attr"`
}
