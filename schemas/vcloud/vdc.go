package vcloud

// Vdc represents the user view of an organization vDC.
// Type: VdcType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the user view of an organization vDC.
// Since: 0.9
type Vdc struct {
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

	// VdcType
	Status string `xml:"status,attr,omitempty"`

	// EntityType
	Description string `xml:"Description,omitempty"`
	Tasks       []Task `xml:"Tasks>Task,omitempty"`

	// VdcType
	AllocationModel    string               `xml:"AllocationModel"`
	ComputeCapacity    *ComputeCapacity     `xml:"ComputeCapacity"`
	ResourceEntities   []*ResourceReference `xml:"ResourceEntities>ResourceEntity,omitempty"`
	AvailableNetworks  []*Reference         `xml:"AvailableNetworks>Network,omitempty"`
	Capabilities       *Capabilities        `xml:"Capabilities,omitempty"`
	NICQuota           int32                `xml:"NicQuota"`
	NetworkQuota       int32                `xml:"NetworkQuota"`
	UsedNetworkCount   int32                `xml:"UsedNetworkCount,omitempty"`
	VMQuota            int32                `xml:"VmQuota"`
	IsEnabled          bool                 `xml:"IsEnabled"`
	VdcStorageProfiles []*Reference         `xml:"VdcStorageProfiles>VdcStorageProfile"`
	VCPUInMHz          int64                `xml:"VCpuInMhz2,omitempty"`
}

// SupportedHardwareVersions contains a list of VMware virtual hardware versions supported in this vDC.
// Type: SupportedHardwareVersionsType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Contains a list of VMware virtual hardware versions supported in this vDC.
// Since: 1.5
type SupportedHardwareVersions struct {
	SupportedHardwareVersion []string `xml:"SupportedHardwareVersion,omitempty"` // A virtual hardware version supported in this vDC.
}

// Capabilities collection of supported hardware capabilities.
// Type: CapabilitiesType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Collection of supported hardware capabilities.
// Since: 1.5
type Capabilities struct {
	Required                  bool                       `xml:"required,attr,omitempty"`
	SupportedHardwareVersions *SupportedHardwareVersions `xml:"SupportedHardwareVersions,omitempty"` // Read-only list of virtual hardware versions supported by this vDC.
}

// ComputeCapacity represents vDC compute capacity.
// Type: ComputeCapacityType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents vDC compute capacity.
// Since: 0.9
type ComputeCapacity struct {
	CPU    *CapacityWithUsage `xml:"Cpu"`
	Memory *CapacityWithUsage `xml:"Memory"`
}

// CapacityWithUsage represents a capacity and usage of a given resource.
// Type: CapacityWithUsageType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a capacity and usage of a given resource.
// Since: 0.9
type CapacityWithUsage struct {
	Required bool `xml:"required,attr,omitempty"`

	// CapacityType
	Units     string `xml:"Units"`
	Allocated int64  `xml:"Allocated,omitempty"`
	Limit     int64  `xml:"Limit,omitempty"`

	// CapacityWithUsageType
	Overhead int64 `xml:"Overhead,omitempty"`
	Reserved int64 `xml:"Reserved,omitempty"`
	Used     int64 `xml:"Used,omitempty"`
}
