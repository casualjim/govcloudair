package vcloud

import (
	"fmt"

	"github.com/vmware/govcloudair/api"
)

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
}

// VAppTemplate gets the vApp template for this catalog item
func (ci *CatalogItem) VAppTemplate(client api.XMLClient) (*VAppTemplate, error) {
	if ci.Entity == nil {
		return nil, fmt.Errorf("no entity present in catalog item [%s]", ci.Name)
	}

	var template VAppTemplate
	if err := client.XMLRequest(HTTPGet, ci.Entity.HREF, ci.Entity.Type, nil, &template); err != nil {
		return nil, err
	}
	return &template, nil
}
