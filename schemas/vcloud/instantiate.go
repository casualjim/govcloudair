package vcloud

import "encoding/xml"

// NewInstantiateVAppTemplateParams creates an instance of InstantiateVAppTemplateParams with xml namespaces filled out.
func NewInstantiateVAppTemplateParams() *InstantiateVAppTemplateParams {
	return &InstantiateVAppTemplateParams{
		Ovf:   NsOvf,
		Xsi:   NsXMLSchema,
		Xmlns: NsVCloud,
	}
}

// InstantiateVAppTemplateParams represents vApp template instantiation parameters.
// Type: InstantiateVAppTemplateParamsType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents vApp template instantiation parameters.
// Since: 0.9
type InstantiateVAppTemplateParams struct {
	XMLName xml.Name `xml:"InstantiateVAppTemplateParams"`
	Ovf     string   `xml:"xmlns:ovf,attr"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	// Attributes
	Name        string `xml:"name,attr,omitempty"`        // Typically used to name or identify the subject of the request. For example, the name of the object being created or modified.
	Deploy      bool   `xml:"deploy,attr"`                // True if the vApp should be deployed at instantiation. Defaults to true.
	PowerOn     bool   `xml:"powerOn,attr"`               // True if the vApp should be powered-on at instantiation. Defaults to true.
	LinkedClone bool   `xml:"linkedClone,attr,omitempty"` // Reserved. Unimplemented.
	// Elements
	Description         string                       `xml:"Description,omitempty"`         // Optional description.
	VAppParent          *Reference                   `xml:"VAppParent,omitempty"`          // Reserved. Unimplemented.
	InstantiationParams *InstantiationParams         `xml:"InstantiationParams,omitempty"` // Instantiation parameters for the composed vApp.
	Source              *Reference                   `xml:"Source"`                        // A reference to a source object such as a vApp or vApp template.
	IsSourceDelete      bool                         `xml:"IsSourceDelete,omitempty"`      // Set to true to delete the source object after the operation completes.
	SourcedItem         *SourcedCompositionItemParam `xml:"SourcedItem,omitempty"`         // Composition item. One of: vApp vAppTemplate Vm.
	AllEULAsAccepted    bool                         `xml:"AllEULAsAccepted,omitempty"`    // True confirms acceptance of all EULAs in a vApp template. Instantiation fails if this element is missing, empty, or set to false and one or more EulaSection elements are present.
}

// InstantiationParams is a container for ovf:Section_Type elements that specify vApp configuration on instantiate, compose, or recompose.
// Type: InstantiationParamsType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Container for ovf:Section_Type elements that specify vApp configuration on instantiate, compose, or recompose.
// Since: 0.9
type InstantiationParams struct {
	CustomizationSection         *CustomizationSection         `xml:"CustomizationSection,omitempty"`
	DefaultStorageProfileSection *DefaultStorageProfileSection `xml:"DefaultStorageProfileSection,omitempty"`
	GuestCustomizationSection    *GuestCustomizationSection    `xml:"GuestCustomizationSection,omitempty"`
	LeaseSettingsSection         *LeaseSettingsSection         `xml:"LeaseSettingsSection,omitempty"`
	NetworkConfigSection         *NetworkConfigSection         `xml:"NetworkConfigSection,omitempty"`
	NetworkConnectionSection     *NetworkConnectionSection     `xml:"NetworkConnectionSection,omitempty"`
	// TODO: Not Implemented
	// SnapshotSection              SnapshotSection              `xml:"SnapshotSection,omitempty"`
}

// SourcedCompositionItemParam represents a vApp, vApp template or Vm to include in a composed vApp.
// Type: SourcedCompositionItemParamType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a vApp, vApp template or Vm to include in a composed vApp.
// Since: 0.9
type SourcedCompositionItemParam struct {
	// Attributes
	SourceDelete bool `xml:"sourceDelete,attr,omitempty"` // True if the source item should be deleted after composition is complete.
	// Elements
	Source              *Reference           `xml:"Source"`                        // Reference to a vApp, vApp template or virtual machine to include in the composition. Changing the name of the newly created VM by specifying name attribute is deprecated. Include VmGeneralParams element instead.
	VMGeneralParams     *VMGeneralParams     `xml:"VmGeneralParams,omitempty"`     // Specify name, description, and other properties of a VM during instantiation.
	VAppScopedLocalID   string               `xml:"VAppScopedLocalId,omitempty"`   // If Source references a Vm, this value provides a unique identifier for the Vm in the scope of the composed vApp.
	InstantiationParams *InstantiationParams `xml:"InstantiationParams,omitempty"` // If Source references a Vm this can include any of the following OVF sections: VirtualHardwareSection OperatingSystemSection NetworkConnectionSection GuestCustomizationSection.
	NetworkAssignment   *NetworkAssignment   `xml:"NetworkAssignment,omitempty"`   // If Source references a Vm, this element maps a network name specified in the Vm to the network name of a vApp network defined in the composed vApp.
	StorageProfile      *Reference           `xml:"StorageProfile,omitempty"`      // If Source references a Vm, this element contains a reference to a storage profile to be used for the Vm. The specified storage profile must exist in the organization vDC that contains the composed vApp. If not specified, the default storage profile for the vDC is used.
	LocalityParams      *LocalityParams      `xml:"LocalityParams,omitempty"`      // Represents locality parameters. Locality parameters provide a hint that may help the placement engine optimize placement of a VM and an independent a Disk so that the VM can make efficient use of the disk.
}

// VMGeneralParams a set of overrides to source VM properties to apply to target VM during copying.
// Type: VmGeneralParamsType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: A set of overrides to source VM properties to apply to target VM during copying.
// Since: 5.6
type VMGeneralParams struct {
	// Elements
	Name               string `xml:"Name,omitempty"`               // Name of VM
	Description        string `xml:"Description,omitempty"`        // VM description
	NeedsCustomization bool   `xml:"NeedsCustomization,omitempty"` // True if this VM needs guest customization
}

// CustomizationSection represents a vApp template customization settings.
// Type: CustomizationSectionType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a vApp template customization settings.
// Since: 1.0
type CustomizationSection struct {
	// FIXME: OVF Section needs to be laid down correctly
	Info string `xml:"ovf:Info"`
	//
	GoldMaster             bool    `xml:"goldMaster,attr,omitempty"`
	HREF                   string  `xml:"href,attr,omitempty"`
	Type                   string  `xml:"type,attr,omitempty"`
	CustomizeOnInstantiate bool    `xml:"CustomizeOnInstantiate"`
	Link                   []*Link `xml:"Link,omitempty"`
}

// DefaultStorageProfileSection is the name of the storage profile that will be specified for this virtual machine. The named storage profile must exist in the organization vDC that contains the virtual machine. If not specified, the default storage profile for the vDC is used.
// Type: DefaultStorageProfileSection_Type
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Name of the storage profile that will be specified for this virtual machine. The named storage profile must exist in the organization vDC that contains the virtual machine. If not specified, the default storage profile for the vDC is used.
// Since: 5.1
type DefaultStorageProfileSection struct {
	StorageProfile string `xml:"StorageProfile,omitempty"`
}

// GuestCustomizationSection represents guest customization settings
// Type: GuestCustomizationSectionType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a guest customization settings.
// Since: 1.0
type GuestCustomizationSection struct {
	// Extends OVF Section_Type
	// Attributes
	Ovf   string `xml:"xmlns:ovf,attr,omitempty"`
	Xsi   string `xml:"xmlns:xsi,attr,omitempty"`
	Xmlns string `xml:"xmlns,attr,omitempty"`

	HREF string `xml:"href,attr,omitempty"` // A reference to the section in URL format.
	Type string `xml:"type,attr,omitempty"` // The MIME type of the section.
	// FIXME: Fix the OVF section
	Info string `xml:"ovf:Info"`
	// Elements
	Enabled               bool    `xml:"Enabled,omitempty"`               // True if guest customization is enabled.
	ChangeSid             bool    `xml:"ChangeSid,omitempty"`             // True if customization can change the Windows SID of this virtual machine.
	VirtualMachineID      string  `xml:"VirtualMachineId,omitempty"`      // Virtual machine ID to apply.
	JoinDomainEnabled     bool    `xml:"JoinDomainEnabled,omitempty"`     // True if this virtual machine can join a Windows Domain.
	UseOrgSettings        bool    `xml:"UseOrgSettings,omitempty"`        // True if customization should use organization settings (OrgGuestPersonalizationSettings) when joining a Windows Domain.
	DomainName            string  `xml:"DomainName,omitempty"`            // The name of the Windows Domain to join.
	DomainUserName        string  `xml:"DomainUserName,omitempty"`        // User name to specify when joining a Windows Domain.
	DomainUserPassword    string  `xml:"DomainUserPassword,omitempty"`    // Password to use with DomainUserName.
	MachineObjectOU       string  `xml:"MachineObjectOU,omitempty"`       // The name of the Windows Domain Organizational Unit (OU) in which the computer account for this virtual machine will be created.
	AdminPasswordEnabled  bool    `xml:"AdminPasswordEnabled,omitempty"`  // True if guest customization can modify administrator password settings for this virtual machine.
	AdminPasswordAuto     bool    `xml:"AdminPasswordAuto,omitempty"`     // True if the administrator password for this virtual machine should be automatically generated.
	AdminPassword         string  `xml:"AdminPassword,omitempty"`         // True if the administrator password for this virtual machine should be set to this string. (AdminPasswordAuto must be false.)
	AdminAutoLogonEnabled bool    `xml:"AdminAutoLogonEnabled,omitempty"` // True if guest administrator should automatically log into this virtual machine.
	AdminAutoLogonCount   int     `xml:"AdminAutoLogonCount,omitempty"`   // Number of times administrator can automatically log into this virtual machine. In case AdminAutoLogon is set to True, this value should be between 1 and 100. Otherwise, it should be 0.
	ResetPasswordRequired bool    `xml:"ResetPasswordRequired,omitempty"` // True if the administrator password for this virtual machine must be reset after first use.
	CustomizationScript   string  `xml:"CustomizationScript,omitempty"`   // Script to run on guest customization. The entire script must appear in this element. Use the XML entity &#13; to represent a newline. Unicode characters can be represented in the form &#xxxx; where xxxx is the character number.
	ComputerName          string  `xml:"ComputerName,omitempty"`          // Computer name to assign to this virtual machine.
	Link                  []*Link `xml:"Link,omitempty"`                  // A link to an operation on this section.
}

// LocalityParams represents locality parameters. Locality parameters provide a hint that may help the placement engine optimize placement of a VM with respect to another VM or an independent disk.  // Type: LocalityParamsType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents locality parameters. Locality parameters provide a hint that may help the placement engine optimize placement of a VM with respect to another VM or an independent disk.
// Since: 5.1
type LocalityParams struct {
	// Elements
	ResourceEntity *Reference `xml:"ResourceEntity,omitempty"` // Reference to a Disk, or a VM.
}

// NetworkAssignment maps a network name specified in a Vm to the network name of a vApp network defined in the VApp that contains the Vm
// Type: NetworkAssignmentType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Maps a network name specified in a Vm to the network name of a vApp network defined in the VApp that contains the Vm
// Since: 0.9
type NetworkAssignment struct {
	// Attributes
	InnerNetwork     string `xml:"innerNetwork,attr"`     // Name of the network as specified in the Vm.
	ContainerNetwork string `xml:"containerNetwork,attr"` // Name of the vApp network to map to.
}
