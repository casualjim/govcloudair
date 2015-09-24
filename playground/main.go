package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
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
	JSONMimeV57 = "application/json;version=5.7"
	// AnyXMLMime511 the wildcard xml mime for version 5.11 of the API
	AnyXMLMime511 = "application/*+xml;version=5.11"
)

const (
	MimeOrgList          = "application/vnd.vmware.vcloud.orgList+xml"
	MimeOrg              = "application/vnd.vmware.vcloud.org+xml"
	MimeCatalog          = "application/vnd.vmware.vcloud.catalog+xml"
	MimeVDC              = "application/vnd.vmware.vcloud.vdc+xml"
	MimeQueryRecords     = "application/vnd.vmware.vchs.query.records+xml"
	MimeAPIExtensibility = "application/vnd.vmware.vcloud.apiextensibility+xml"
	MimeEntity           = "application/vnd.vmware.vcloud.entity+xml"
	MimeQueryList        = "application/vnd.vmware.vcloud.query.queryList+xml"
	MimeSession          = "application/vnd.vmware.vcloud.session+xml"
	MimeTask             = "application/vnd.vmware.vcloud.task+xml"
	MimeError            = "application/vnd.vmware.vcloud.error+xml"
)

const (
	RelDown           = "down"
	RelRemove         = "remove"
	RelEntityResolver = "entityResolver"
	RelExtensibility  = "down:extensibility"
)

const (
	HTTPGet    = "GET"
	HTTPPost   = "POST"
	HTTPPut    = "PUT"
	HTTPPatch  = "PATCH"
	HTTPDelete = "DELETE"
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

	if err := client.JSONRequest(HTTPGet, InstancesPath, &client.Info); err != nil {
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

	b, err := json.MarshalIndent(orgList, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

}

var (
	apiLoginResponse = `{"serviceGroupIds":["e7434542-25c7-4120-88ec-6038a5770828"]}`
)

func newAuthenticatedClient(user, password string) (*authenticatedClient, error) {
	r, _ := http.NewRequest(HTTPPost, BaseURL+LoginPath, nil)
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

func (a *accountInstanceAttrs) Authenticate(username, password string) (*Session, error) {
	r, _ := http.NewRequest(HTTPPost, a.SessionURI, nil)
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
	HREF  string   `xml:"href,attr,omitempty"`
	Type  string   `xml:"type,attr,omitempty"`
	Links LinkList `xml:"Link,omitempty"`

	// SessionType
	Org    string `xml:"org,attr,omitempty"`
	Roles  string `xml:"roles,attr,omitempty"`
	User   string `xml:"user,attr,omitempty"`
	UserID string `xml:"userId,attr,omitempty"`

	Token string `xml:"-"`
}

func (s *Session) OrgList() (*OrgList, error) {
	lnk := s.Links.ForType(MimeOrgList, RelDown)
	if lnk == nil {
		return nil, errors.New("Couldn't find the link for orgList")
	}

	var orgList OrgList
	if err := s.XMLRequest(HTTPGet, lnk.HREF, MimeOrgList, nil, &orgList); err != nil {
		return nil, err
	}

	return &orgList, nil
}

func (s *Session) XMLRequest(method, url, tpe string, body, result interface{}) error {
	var buf *bytes.Buffer
	if body != nil {
		buf = bytes.NewBuffer(nil)
		enc := xml.NewEncoder(buf)
		if err := enc.Encode(body); err != nil {
			return err
		}
	}

	r, _ := http.NewRequest(method, url, buf)
	r.Header.Set(HeaderAccept, tpe)
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

	if resp.StatusCode/100 != 2 {
		log.Fatalf("Could not complete request with vca, because (status %d) %s\n", resp.StatusCode, resp.Status)
	}

	dec := xml.NewDecoder(resp.Body)
	if err := dec.Decode(result); err != nil {
		return err
	}

	return nil
}

type LinkList []*Link

func (l LinkList) ForType(tpe, rel string) *Link {
	if rel == "" {
		rel = RelDown
	}

	for _, lnk := range l {
		if lnk != nil && lnk.Type == tpe && lnk.Rel == rel {
			return lnk
		}
	}
	return nil
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

type OrgList struct {
	HREF string `xml:"href,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`

	// ResourceType
	Link LinkList `xml:"Link,omitempty"`

	// OrgListType
	Orgs []Org `xml:"Org,omitempty"`
}

// Reference is a reference to a resource. Contains an href attribute and optional name and type attributes.
// Type: ReferenceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: A reference to a resource. Contains an href attribute and optional name and type attributes.
// Since: 0.9
type Reference struct {
	HREF string `xml:"href,attr"`
	ID   string `xml:"id,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

// Org represents the user view of a vCloud Director organization.
// Type: OrgType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the user view of a vCloud Director organization.
// Since: 0.9
type Org struct {
	HREF         string `xml:"href,attr,omitempty"`
	Type         string `xml:"type,attr,omitempty"`
	ID           string `xml:"id,attr,omitempty"`
	OperationKey string `xml:"operationKey,attr,omitempty"`
	Name         string `xml:"name,attr"`

	// resourcetype
	Link LinkList `xml:"Link,omitempty"`

	// entitytype
	Description sring  `xml:"Description,omitempty"`
	Tasks       []Task `xml:"Tasks>Task,omitempty"`

	// orgtype
	FullName  string `xml:"FullName"`
	IsEnabled bool   `xml:"IsEnabled,omitempty"`
}

type Task struct {
	// ResourceType
	HREF string `xml:"href,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`

	// IdentifiableResourceType
	ID           string `xml:"id,attr,omitempty"`
	OperationKey string `xml:"operationKey,attr,omitempty"`

	// EntityType
	Name string `xml:"name,attr"`

	// TaskType
	CancelRequested  bool   `xml:"cancelRequested,attr,omitempty"`
	EndTime          string `xml:"endTime,attr,omitempty"`
	ExpiryTime       string `xml:"expiryTime,attr,omitempty"`
	Operation        string `xml:"operation,attr,omitempty"`
	OperationName    string `xml:"operationName,attr,omitempty"`
	ServiceNamespace string `xml:"serviceNamespace,attr,omitempty"`
	StartTime        string `xml:"startTime,attr,omitempty"`
	Status           string `xml:"status,attr,omitempty"`

	// resourcetype
	Link LinkList `xml:"Link,omitempty"`

	// entitytype
	Description string `xml:"Description,omitempty"`
	Tasks       []Task `xml:"Tasks>Task,omitempty"`

	// TaskType
	Owner        *Reference  `xml:Owner,omitempty`
	Error        *Error      `xml:Error,omitempty`
	User         *Reference  `xml:User,omitempty`
	Organization *Reference  `xml:Organization,omitempty`
	Progress     int         `xml:Progress,omitempty`
	Params       interface{} `xml:Params,omitempty`
	Details      string      `xml:Details,omitempty`
}

// Error is the standard error message type used in the vCloud REST API.
// Type: ErrorType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: The standard error message type used in the vCloud REST API.
// Since: 0.9
type Error struct {
	MajorErrorCode          int    `xml:"majorErrorCode,attr"`
	Message                 string `xml:"message,attr"`
	MinorErrorCode          string `xml:"minorErrorCode,attr"`
	VendorSpecificErrorCode string `xml:"vendorSpecificErrorCode,attr,omitempty"`
	StackTrace              string `xml:"stackTrace,attr,omitempty"`

	TenantError *TenantError `xml:"TentantError,omitempty"`
}

// TenantError is the standard error message type used in the vCloud REST API.
// Type: ErrorType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: The standard error message type used in the vCloud REST API.
// Since: 0.9
type TenantError struct {
	MajorErrorCode          int    `xml:"majorErrorCode,attr"`
	Message                 string `xml:"message,attr"`
	MinorErrorCode          string `xml:"minorErrorCode,attr"`
	VendorSpecificErrorCode string `xml:"vendorSpecificErrorCode,attr,omitempty"`
	StackTrace              string `xml:"stackTrace,attr,omitempty"`
}
