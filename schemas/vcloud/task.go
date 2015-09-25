package vcloud

// Task represents an asynchronous operation in vCloud Director.
// Type: TaskType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents an asynchronous operation in vCloud Director.
// Since: 0.9
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
	Owner        *Reference  `xml:"Owner,omitempty"`
	Error        *Error      `xml:"Error,omitempty"`
	User         *Reference  `xml:"User,omitempty"`
	Organization *Reference  `xml:"Organization,omitempty"`
	Progress     int         `xml:"Progress,omitempty"`
	Params       interface{} `xml:"Params,omitempty"`
	Details      string      `xml:"Details,omitempty"`
}
