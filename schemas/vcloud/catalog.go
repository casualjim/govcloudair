package vcloud

import (
	"fmt"

	"github.com/vmware/govcloudair/api"
)

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
}

func (c *Catalog) ItemForName(name string, client api.XMLClient) (*CatalogItem, error) {
	for _, p := range c.CatalogItems {
		if p.Name == name {
			var ci CatalogItem
			if err := client.XMLRequest(HTTPGet, p.HREF, p.Type, nil, &ci); err != nil {
				return nil, err
			}

			return &ci, nil
		}
	}

	return nil, fmt.Errorf("no item found in catalog for %q", name)
}
