package vcloud

// XMLClient a client capable of making XML requests to a vcloud API
type XMLClient interface {
	XMLRequest(string, string, string, interface{}, interface{}) error
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

// ResourceReference represents a reference to a resource. Contains an href attribute, a resource status attribute, and optional name and type attributes.
// Type: ResourceReferenceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a reference to a resource. Contains an href attribute, a resource status attribute, and optional name and type attributes.
// Since: 0.9
type ResourceReference struct {
	HREF   string `xml:"href,attr"`
	ID     string `xml:"id,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"`
	Name   string `xml:"name,attr,omitempty"`
	Status string `xml:"status,attr,omitempty"`
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

// Property contains a key/value pair.
// Type: PropertyType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Since: 0.9
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
