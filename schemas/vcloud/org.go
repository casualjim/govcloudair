package vcloud

import (
	"errors"
	"fmt"
)

// OrgList represents a list of organizations.
// Type: OrgListType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Since: 0.9
type OrgList struct {
	HREF string `xml:"href,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`

	// ResourceType
	Links LinkList `xml:"Link,omitempty"`

	// OrgListType
	Orgs []Reference `xml:"Org,omitempty"`
}

// FirstOrg retrieves the first organization from the org list
func (o *OrgList) FirstOrg(client XMLClient) (*Org, error) {
	if len(o.Orgs) == 0 {
		return nil, errors.New("orgList has no orgs, can't get the first")
	}

	var org Org
	if err := client.XMLRequest(HTTPGet, o.Orgs[0].HREF, o.Orgs[0].Type, nil, &org); err != nil {
		return nil, err
	}

	return &org, nil
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
}

// RetrievePublicCatalog convenience function to retrieve the public catalog
func (o *Org) RetrievePublicCatalog(client XMLClient) (*Catalog, error) {
	return o.RetrieveCatalog(PublicCatalog, client)
}

// RetrieveDefaultCatalog convenience function to retrieve the default catalog
func (o *Org) RetrieveDefaultCatalog(client XMLClient) (*Catalog, error) {
	return o.RetrieveCatalog(DefaultCatalog, client)
}

// RetrieveCatalog retrieve a named catalog
func (o *Org) RetrieveCatalog(name string, client XMLClient) (*Catalog, error) {
	lnk := o.Links.ForName(name, MimeCatalog, RelDown)
	if lnk == nil {
		return nil, fmt.Errorf("no catalog link found for %q", o.ID)
	}

	var catalog Catalog
	if err := client.XMLRequest(HTTPGet, lnk.HREF, lnk.Type, nil, &catalog); err != nil {
		return nil, err
	}
	return &catalog, nil
}

// FindVDC finds the named VDC for this org
func (o *Org) FindVDC(name string, client XMLClient) (*VDC, error) {
	lnk := o.Links.ForName(name, MimeVDC, RelDown)
	if lnk == nil {
		return nil, fmt.Errorf("no VDC link found for %q", o.ID)
	}

	var vdc VDC
	if err := client.XMLRequest(HTTPGet, lnk.HREF, lnk.Type, nil, &vdc); err != nil {
		return nil, err
	}
	return &vdc, nil
}