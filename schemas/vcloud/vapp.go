package vcloud

import (
	"encoding/xml"

	"github.com/vmware/govcloudair/api"
)

// VAppTemplateFields contains the shared fields for a vapp template
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
	Owner    *Owner       `xml:"Owner,omitempty"`       // vAppTemplate owner.
	Children []VMTemplate `xml:"Children>Vm,omitempty"` // Container for virtual machines included in this vApp template.

	VAppScopedLocalID     string `xml:"VAppScopedLocalId"`               // A unique identifier for the Vm in the scope of the vApp template.
	DefaultStorageProfile string `xml:"DefaultStorageProfile,omitempty"` // The name of the storage profile to be used for this object. The named storage profile must exist in the organization vDC that contains the object. If not specified, the default storage profile for the vDC is used.
	DateCreated           string `xml:"DateCreated,omitempty"`           // Creation date/time of the template.

	// FIXME: Upstream bug? Missing NetworkConfigSection, LeaseSettingSection and
	// GuestCustomizationSection at least, NetworkConnectionSection is required when
	// using ComposeVApp action in the context of a Children VM (still
	// referenced by VAppTemplateType).
	NetworkConfigSection      *NetworkConfigSection      `xml:"NetworkConfigSection,omitempty"`
	NetworkConnectionSection  *NetworkConnectionSection  `xml:"NetworkConnectionSection,omitempty"`
	LeaseSettingsSection      *LeaseSettingsSection      `xml:"LeaseSettingsSection,omitempty"`
	GuestCustomizationSection *GuestCustomizationSection `xml:"GuestCustomizationSection,omitempty"`
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

// Ref returns the reference object to this vapp template fields
func (v *VAppTemplate) Ref() *Reference {
	var ref Reference
	ref.HREF = v.HREF
	ref.Type = v.Type
	ref.Name = v.Name
	ref.ID = v.ID
	return &ref
}

// VMTemplate represents a vApp child template.
// Type: VAppTemplateType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a vApp template.
// Since: 0.9
type VMTemplate struct {
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

// VApp representa a vApp
// Type: VAppType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a vApp.
// Since: 0.9
type VApp struct {
	// Attributes
	HREF                  string `xml:"href,attr,omitempty"`                  // The URI of the entity.
	Type                  string `xml:"type,attr,omitempty"`                  // The MIME type of the entity.
	ID                    string `xml:"id,attr,omitempty"`                    // The entity identifier, expressed in URN format. The value of this attribute uniquely identifies the entity, persists for the life of the entity, and is never reused.
	OperationKey          string `xml:"operationKey,attr,omitempty"`          // Optional unique identifier to support idempotent semantics for create and delete operations.
	Name                  string `xml:"name,attr"`                            // The name of the entity.
	Status                int    `xml:"status,attr,omitempty"`                // Creation status of the resource entity.
	Deployed              bool   `xml:"deployed,attr,omitempty"`              // True if the virtual machine is deployed.
	OvfDescriptorUploaded bool   `xml:"ovfDescriptorUploaded,attr,omitempty"` // Read-only indicator that the OVF descriptor for this vApp has been uploaded.
	// Elements
	Link        []*Link    `xml:"Link,omitempty"`        // A reference to an entity or operation associated with this object.
	Description string     `xml:"Description,omitempty"` // Optional description.
	Tasks       []*Task    `xml:"Tasks>Task,omitempty"`  // A list of queued, running, or recently completed tasks associated with this entity.
	Files       []*File    `xml:"Files>File,omitempty"`  // Represents a list of files to be transferred (uploaded or downloaded). Each File in the list is part of the ResourceEntity.
	VAppParent  *Reference `xml:"VAppParent,omitempty"`  // Reserved. Unimplemented.
	// TODO: OVF Sections to be implemented
	// Section OVF_Section `xml:"Section"`
	DateCreated       string `xml:"DateCreated,omitempty"`       // Creation date/time of the vApp.
	Owner             *Owner `xml:"Owner,omitempty"`             // vApp owner.
	InMaintenanceMode bool   `xml:"InMaintenanceMode,omitempty"` // True if this vApp is in maintenance mode. Prevents users from changing vApp metadata.
	Children          []VM   `xml:"Children>Vm,omitempty"`       // Container for virtual machines included in this vApp.
}

// Refresh refreshes this vApp
func (v *VApp) Refresh(client api.XMLClient) error {
	var nw VApp
	if err := client.XMLRequest(HTTPGet, v.HREF, v.Type, nil, &nw); err != nil {
		return err
	}
	*v = nw
	return nil
}

// VM represents a virtual machine
// Type: VmType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a virtual machine.
// Since: 0.9
type VM struct {
	// Attributes
	Ovf   string `xml:"xmlns:ovf,attr,omitempty"`
	Xsi   string `xml:"xmlns:xsi,attr,omitempty"`
	Xmlns string `xml:"xmlns,attr,omitempty"`

	HREF                    string `xml:"href,attr,omitempty"`                    // The URI of the entity.
	Type                    string `xml:"type,attr,omitempty"`                    // The MIME type of the entity.
	ID                      string `xml:"id,attr,omitempty"`                      // The entity identifier, expressed in URN format. The value of this attribute uniquely identifies the entity, persists for the life of the entity, and is never reused
	OperationKey            string `xml:"operationKey,attr,omitempty"`            // Optional unique identifier to support idempotent semantics for create and delete operations.
	Name                    string `xml:"name,attr"`                              // The name of the entity.
	Status                  int    `xml:"status,attr,omitempty"`                  // Creation status of the resource entity.
	Deployed                bool   `xml:"deployed,attr,omitempty"`                // True if the virtual machine is deployed.
	NeedsCustomization      bool   `xml:"needsCustomization,attr,omitempty"`      // True if this virtual machine needs customization.
	NestedHypervisorEnabled bool   `xml:"nestedHypervisorEnabled,attr,omitempty"` // True if hardware-assisted CPU virtualization capabilities in the host should be exposed to the guest operating system.
	// Elements
	Link        []*Link    `xml:"Link,omitempty"`        // A reference to an entity or operation associated with this object.
	Description string     `xml:"Description,omitempty"` // Optional description.
	Tasks       []*Task    `xml:"Tasks>Task,omitempty"`  // A list of queued, running, or recently completed tasks associated with this entity.
	Files       []*File    `xml:"Files>File,omitempty"`  // Represents a list of files to be transferred (uploaded or downloaded). Each File in the list is part of the ResourceEntity.
	VAppParent  *Reference `xml:"VAppParent,omitempty"`  // Reserved. Unimplemented.
	// TODO: OVF Sections to be implemented
	// Section OVF_Section `xml:"Section,omitempty"
	DateCreated string `xml:"DateCreated"` // Creation date/time of the vApp.

	// FIXME: Upstream bug? Missing NetworkConnectionSection
	NetworkConnectionSection *NetworkConnectionSection `xml:"NetworkConnectionSection,omitempty"`

	VAppScopedLocalID string `xml:"VAppScopedLocalId,omitempty"` // A unique identifier for the virtual machine in the scope of the vApp.

	// TODO: OVF Sections to be implemented
	// Environment OVF_Environment `xml:"Environment,omitempty"

	VMCapabilities *VMCapabilities `xml:"VmCapabilities,omitempty"` // Allows you to specify certain capabilities of this virtual machine.
	StorageProfile *Reference      `xml:"StorageProfile,omitempty"` // A reference to a storage profile to be used for this object. The specified storage profile must exist in the organization vDC that contains the object. If not specified, the default storage profile for the vDC is used.
}

// VMCapabilities allows you to specify certain capabilities of this virtual machine.
// Type: VmCapabilitiesType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Allows you to specify certain capabilities of this virtual machine.
// Since: 5.1
type VMCapabilities struct {
	HREF                string  `xml:"href,attr,omitempty"`
	Type                string  `xml:"type,attr,omitempty"`
	CPUHotAddEnabled    bool    `xml:"CpuHotAddEnabled,omitempty"`
	Link                []*Link `xml:"Link,omitempty"`
	MemoryHotAddEnabled bool    `xml:"MemoryHotAddEnabled,omitempty"`
}
