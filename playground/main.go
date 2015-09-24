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
	// Version511 the 5.11 version
	Version511 = "5.11"
	// Version is the default version number
	Version = Version511

	// Public Catalog Name
	PublicCatalog = "Public Catalog"

	// Default Catalog Name
	DefaultCatalog = "Default Catalog"
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
	RelAdd            = "add"
	RelCopy           = "copy"
	RelMove           = "move"
	RelUp             = "up"
	RelDown           = "down"
	RelRemove         = "remove"
	RelEntityResolver = "entityResolver"
	RelExtensibility  = "down:extensibility"
	RelControlAccess  = "controlAccess"
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

	org, err := orgList.FirstOrg()
	if err != nil {
		log.Fatalln(err)
	}

	catalog, err := org.RetrievePublicCatalog()
	if err != nil {
		log.Fatalln(err)
	}

	catalogItem, err := catalog.ItemForName("VMware Photon OS - Tech Preview 2")
	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.MarshalIndent(catalogItem, "", "  ")
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
	http *http.Client
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
	orgList.session = s

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

	r.Header.Set(HeaderAccept, tpe+";version="+Version)
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

func (l LinkList) ForName(name, tpe, rel string) *Link {
	if rel == "" {
		rel = RelDown
	}

	for _, lnk := range l {
		if lnk != nil && lnk.Name == name && lnk.Type == tpe && lnk.Rel == rel {
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
	Links LinkList `xml:"Link,omitempty"`

	// OrgListType
	Orgs []Reference `xml:"Org,omitempty"`

	session *Session `xml:"-" json:"-"`
}

func (o *OrgList) FirstOrg() (*Org, error) {
	if len(o.Orgs) == 0 {
		return nil, errors.New("orgList has no orgs, can't get the first")
	}

	var org Org
	if err := o.session.XMLRequest(HTTPGet, o.Orgs[0].HREF, o.Orgs[0].Type, nil, &org); err != nil {
		return nil, err
	}

	org.session = o.session
	return &org, nil
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
	Links LinkList `xml:"Link,omitempty"`

	// entitytype
	Description string `xml:"Description,omitempty"`
	Tasks       []Task `xml:"Tasks>Task,omitempty"`

	// orgtype
	FullName  string `xml:"FullName"`
	IsEnabled bool   `xml:"IsEnabled,omitempty"`

	session *Session
}

func (o *Org) RetrievePublicCatalog() (*Catalog, error) {
	return o.RetrieveCatalog(PublicCatalog)
}

func (o *Org) RetrieveDefaultCatalog() (*Catalog, error) {
	return o.RetrieveCatalog(DefaultCatalog)
}

func (o *Org) RetrieveCatalog(name string) (*Catalog, error) {
	lnk := o.Links.ForName(name, MimeCatalog, RelDown)
	if lnk == nil {
		return nil, fmt.Errorf("no catalog link found for %q", o.ID)
	}

	var catalog Catalog
	if err := o.session.XMLRequest(HTTPGet, lnk.HREF, lnk.Type, nil, &catalog); err != nil {
		return nil, err
	}

	catalog.session = o.session
	return &catalog, nil
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

	// ResourceType
	Links LinkList `xml:"Link,omitempty"`

	// EntityType
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

// Owner represents the owner of this entity.
// Type: OwnerType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the owner of this entity.
// Since: 1.5
type Owner struct {
	// ResourceType
	HREF string `xml:"href,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`

	// ResourceType
	Links LinkList `xml:"Link,omitempty"`

	// OwnerType
	User *Reference `xml:"User"`
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

// Catalog represents the user view of a Catalog object.
// Type: CatalogType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the user view of a Catalog object.
// Since: 0.9
type Catalog struct {
	// ResourceType
	HREF string `xml:"href,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`

	// IdentifiableResourceType
	ID           string `xml:"id,attr,omitempty"`
	OperationKey string `xml:"operationKey,attr,omitempty"`

	// EntityType
	Name string `xml:"name,attr"`

	// ResourceType
	Links LinkList `xml:"Link,omitempty"`

	// EntityType
	Description string `xml:"Description,omitempty"`
	Tasks       []Task `xml:"Tasks>Task,omitempty"`

	// CatalogType
	Owner         *Owner      `xml:"Owner,omitempty"`
	CatalogItems  []Reference `xml:"CatalogItems>CatalogItem"`
	IsPublished   bool        `xml:"IsPublished"`
	DateCreated   string      `xml:"DateCreated"`
	VersionNumber int64       `xml:"VersionNumber"`

	session *Session `xml:"-" json:"-"`
}

func (c *Catalog) ItemForName(name string) (*CatalogItem, error) {
	for _, p := range c.CatalogItems {
		if p.Name == name {
			var ci CatalogItem
			if err := c.session.XMLRequest(HTTPGet, p.HREF, p.Type, nil, &ci); err != nil {
				return nil, err
			}

			ci.session = c.session
			return &ci, nil
		}
	}

	return nil, fmt.Errorf("no item found in catalog for %q", name)
}

// CatalogItem contains a reference to a VappTemplate or Media object and related metadata.
// Type: CatalogItemType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Contains a reference to a VappTemplate or Media object and related metadata.
// Since: 0.9
type CatalogItem struct {
	// ResourceType
	HREF string `xml:"href,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`

	// IdentifiableResourceType
	ID           string `xml:"id,attr,omitempty"`
	OperationKey string `xml:"operationKey,attr,omitempty"`

	// EntityType
	Name string `xml:"name,attr"`

	// CatalogItemType
	Size int64 `xml:"size,attr,omitempty"`

	// ResourceType
	Links LinkList `xml:"Link,omitempty"`

	// EntityType
	Description string `xml:"Description,omitempty"`
	Tasks       []Task `xml:"Tasks>Task,omitempty"`

	// CatalogItemType
	Entity        *Entity    `xml:"Entity"`
	Properties    []Property `xml:"Property,omitempty"`
	DateCreated   string     `xml:"DateCreated,omitempty"`
	VersionNumber int64      `xml:"VersionNumber,omitempty"`

	session *Session `xml:"-" json:"-"`
}

// Property
type Property struct {
	Key   string `xml:"key,attr"`
	Value string `xml:",innerxml"`
}

// Entity is a basic entity type in the vCloud object model. Includes a name, an optional description, and an optional list of links.
// Type: EntityType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Basic entity type in the vCloud object model. Includes a name, an optional description, and an optional list of links.
// Since: 0.9
type Entity struct {
	// ResourceType
	HREF string `xml:"href,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`

	// IdentifiableResourceType
	ID           string `xml:"id,attr,omitempty"`
	OperationKey string `xml:"operationKey,attr,omitempty"`

	// EntityType
	Name string `xml:"name,attr"`

	// ResourceType
	Links LinkList `xml:"Link,omitempty"`

	// EntityType
	Description string `xml:"Description,omitempty"`
	Tasks       []Task `xml:"Tasks>Task,omitempty"`
}
