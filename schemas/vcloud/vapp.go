package vcloud

import "encoding/xml"

type VAppTemplateFields struct {
	// ResourceType
	HREF string `xml:"href,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`

	// IdentifiableResourceType
	ID           string `xml:"id,attr,omitempty"`
	OperationKey string `xml:"operationKey,attr,omitempty"`

	// EntityType
	Name string `xml:"name,attr"`

	// ResourceEntityType
	Status int32 `xml:"status,attr,omitempty"`

	// ResourceType
	Links LinkList `xml:"Link,omitempty"`

	// EntityType
	Description string `xml:"Description,omitempty"`
	Tasks       []Task `xml:"Tasks>Task,omitempty"`

	// ResourceEntityType
	Files []File `xml:"Files>File"`

	OvfDescriptorUploaded string `xml:"ovfDescriptorUploaded,attr,omitempty"` // True if the OVF descriptor for this template has been uploaded.
	GoldMaster            bool   `xml:"goldMaster,attr,omitempty"`            // True if this template is a gold master.
	// Elements
	Owner    *Owner `xml:"Owner,omitempty"`       // vAppTemplate owner.
	Children []VM   `xml:"Children>Vm,omitempty"` // Container for virtual machines included in this vApp template.

	VAppScopedLocalID     string `xml:"VAppScopedLocalId"`               // A unique identifier for the Vm in the scope of the vApp template.
	DefaultStorageProfile string `xml:"DefaultStorageProfile,omitempty"` // The name of the storage profile to be used for this object. The named storage profile must exist in the organization vDC that contains the object. If not specified, the default storage profile for the vDC is used.
	DateCreated           string `xml:"DateCreated,omitempty"`           // Creation date/time of the template.

	// FIXME: Upstream bug? Missing NetworkConfigSection, LeaseSettingSection and
	// CustomizationSection at least, NetworkConnectionSection is required when
	// using ComposeVApp action in the context of a Children VM (still
	// referenced by VAppTemplateType).
	NetworkConfigSection *NetworkConfigSection `xml:"NetworkConfigSection,omitempty"`
	//NetworkConnectionSection *NetworkConnectionSection `xml:"NetworkConnectionSection,omitempty"`
	//LeaseSettingsSection     *LeaseSettingsSection     `xml:"LeaseSettingsSection,omitempty"`
	//CustomizationSection     *CustomizationSection     `xml:"CustomizationSection,omitempty"`
	//Section ovf.Section `xml:"Section,omitempty"`
}

// VAppTemplate represents a vApp template.
// Type: VAppTemplateType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a vApp template.
// Since: 0.9
type VAppTemplate struct {
	XMLName xml.Name `xml:"VAppTemplate"`
	VAppTemplateFields
}

type VM struct {
	XMLName xml.Name `xml:"Vm"`
	VAppTemplateFields
}

// File represents a file to be transferred (uploaded or downloaded).
// Type: FileType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a file to be transferred (uploaded or downloaded).
// Since: 0.9
type File struct {

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

	// FileType
	Size             int64  `xml:"size,attr,omitempty"`
	BytesTransferred int64  `xml:"bytesTransferred,attr,omitempty"`
	Checksum         string `xml:"checksum,attr,omitempty"`
}
