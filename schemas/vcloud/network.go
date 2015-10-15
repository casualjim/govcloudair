package vcloud

// LeaseSettingsSection represents vApp lease settings.
// Type: LeaseSettingsSectionType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents vApp lease settings.
// Since: 0.9
type LeaseSettingsSection struct {
	HREF                      string `xml:"href,attr,omitempty"`
	Type                      string `xml:"type,attr,omitempty"`
	DeploymentLeaseExpiration string `xml:"DeploymentLeaseExpiration,omitempty"`
	DeploymentLeaseInSeconds  int    `xml:"DeploymentLeaseInSeconds,omitempty"`
	Link                      *Link  `xml:"Link,omitempty"`
	StorageLeaseExpiration    string `xml:"StorageLeaseExpiration,omitempty"`
	StorageLeaseInSeconds     int    `xml:"StorageLeaseInSeconds,omitempty"`
}

// IPRange represents a range of IP addresses, start and end inclusive.
// Type: IpRangeType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a range of IP addresses, start and end inclusive.
// Since: 0.9
type IPRange struct {
	EndAddress   string `xml:"EndAddress"`   // End address of the IP range.
	StartAddress string `xml:"StartAddress"` // Start address of the IP range.
}

// DhcpService represents a DHCP network service.
// Type: DhcpServiceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a DHCP network service.
// Since:
type DhcpService struct {
	DefaultLeaseTime    int      `xml:"DefaultLeaseTime,omitempty"`    // Default lease in seconds for DHCP addresses.
	DomainName          string   `xml:"DomainName,omitempty"`          //	The domain name.
	IPRange             *IPRange `xml:"IpRange"`                       //	IP range for DHCP addresses.
	IsEnabled           bool     `xml:"IsEnabled"`                     // Enable or disable the service using this flag
	MaxLeaseTime        int      `xml:"MaxLeaseTime"`                  //	Max lease in seconds for DHCP addresses.
	PrimaryNameServer   string   `xml:"PrimaryNameServer,omitempty"`   // The primary name server.
	RouterIP            string   `xml:"RouterIp,omitempty"`            // Router IP.
	SecondaryNameServer string   `xml:"SecondaryNameServer,omitempty"` // The secondary name server.
	SubMask             string   `xml:"SubMask,omitempty"`             // The subnet mask.
}

// NetworkFeatures represents features of a network.
// Type: NetworkFeaturesType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents features of a network.
// Since:
type NetworkFeatures struct {
	DhcpService          *DhcpService          `xml:"DhcpService,omitempty"`          // Substitute for NetworkService. DHCP service settings
	FirewallService      *FirewallService      `xml:"FirewallService,omitempty"`      // Substitute for NetworkService. Firewall service settings
	NatService           *NatService           `xml:"NatService,omitempty"`           // Substitute for NetworkService. NAT service settings
	LoadBalancerService  *LoadBalancerService  `xml:"LoadBalancerService,omitempty"`  // Substitute for NetworkService. Load Balancer service settings
	StaticRoutingService *StaticRoutingService `xml:"StaticRoutingService,omitempty"` // Substitute for NetworkService. Static Routing service settings
	// TODO: Not Implemented
	// IpsecVpnService      IpsecVpnService      `xml:"IpsecVpnService,omitempty"`      // Substitute for NetworkService. Ipsec Vpn service settings
}

// IPAddresses a list of IP addresses
// Type: IpAddressesType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: A list of IP addresses.
// Since: 0.9
type IPAddresses struct {
	IPAddress string `xml:"IpAddress,omitempty"` // An IP address.
}

// IPRanges representsa list of IP ranges.
// Type: IpRangesType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a list of IP ranges.
// Since: 0.9
type IPRanges struct {
	IPRange []*IPRange `xml:"IpRange,omitempty"` // IP range.
}

// IPScope specifies network settings like gateway, network mask, DNS servers IP ranges etc
// Type: IpScopeType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Specify network settings like gateway, network mask, DNS servers, IP ranges, etc.
// Since: 0.9
type IPScope struct {
	IsInherited          bool            `xml:"IsInherited"`                    // True if the IP scope is inherit from parent network.
	Gateway              string          `xml:"Gateway,omitempty"`              // Gateway of the network.
	Netmask              string          `xml:"Netmask,omitempty"`              // Network mask.
	DNS1                 string          `xml:"Dns1,omitempty"`                 // Primary DNS server.
	DNS2                 string          `xml:"Dns2,omitempty"`                 // Secondary DNS server.
	DNSSuffix            string          `xml:"DnsSuffix,omitempty"`            // DNS suffix.
	IsEnabled            bool            `xml:"IsEnabled"`                      // Indicates if subnet is enabled or not. Default value is True.
	IPRanges             *IPRanges       `xml:"IpRanges,omitempty"`             // IP ranges used for static pool allocation in the network.
	AllocatedIPAddresses *IPAddresses    `xml:"AllocatedIpAddresses,omitempty"` // Read-only list of allocated IP addresses in the network.
	SubAllocations       *SubAllocations `xml:"SubAllocations,omitempty"`       // Read-only list of IP addresses that are sub allocated to edge gateways.
}

// SubAllocations a list of IP addresses that are sub allocated to edge gateways.
// Type: SubAllocationsType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: A list of IP addresses that are sub allocated to edge gateways.
// Since: 5.1
type SubAllocations struct {
	// Attributes
	HREF string `xml:"href,attr,omitempty"` // The URI of the entity.
	Type string `xml:"type,attr,omitempty"` // The MIME type of the entity.
	// Elements
	Link          []*Link        `xml:"Link,omitempty"`          // A reference to an entity or operation associated with this object.
	SubAllocation *SubAllocation `xml:"SubAllocation,omitempty"` // IP Range sub allocated to a edge gateway.
}

// SubAllocation IP range sub allocated to an edge gateway.
// Type: SubAllocationType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: IP range sub allocated to an edge gateway.
// Since: 5.1
type SubAllocation struct {
	EdgeGateway *Reference `xml:"EdgeGateway,omitempty"` // Edge gateway that uses this sub allocation.
	IPRanges    *IPRanges  `xml:"IpRanges,omitempty"`    // IP range sub allocated to the edge gateway.
}

// NetworkConfiguration the configuration applied to a network. This is an abstract base type. The concrete types include thos for vApp and Organization wide networks.
// Type: NetworkConfigurationType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: The configurations applied to a network. This is an abstract base type. The concrete types include those for vApp and Organization wide networks.
// Since: 0.9
type NetworkConfiguration struct {
	BackwardCompatibilityMode      bool             `xml:"BackwardCompatibilityMode"`
	Features                       *NetworkFeatures `xml:"Features,omitempty"`
	ParentNetwork                  *Reference       `xml:"ParentNetwork,omitempty"`
	FenceMode                      string           `xml:"FenceMode"`
	IPScopes                       *IPScopes        `xml:"IpScopes,omitempty"`
	RetainNetInfoAcrossDeployments bool             `xml:"RetainNetInfoAcrossDeployments"`
	// TODO: Not Implemented
	// RouterInfo                     RouterInfo           `xml:"RouterInfo,omitempty"`
	// SyslogServerSettings           SyslogServerSettings `xml:"SyslogServerSettings,omitempty"`
}

// IPScopes represents a list of IP scopes.
// Type: IpScopesType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a list of IP scopes.
// Since: 5.1
type IPScopes struct {
	IPScope []IPScope `xml:"IpScope"` // IP scope.
}

// VAppNetworkConfiguration representa a vApp network configuration
// Type: VAppNetworkConfigurationType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a vApp network configuration.
// Since: 0.9
type VAppNetworkConfiguration struct {
	HREF        string `xml:"href,attr,omitempty"`
	Type        string `xml:"type,attr,omitempty"`
	NetworkName string `xml:"networkName,attr"`

	Configuration *NetworkConfiguration `xml:"Configuration"`
	Description   string                `xml:"Description,omitempty"`
	IsDeployed    bool                  `xml:"IsDeployed"`
	Link          *Link                 `xml:"Link,omitempty"`
}

// NetworkConfigSection is container for vApp networks.
// Type: NetworkConfigSectionType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Container for vApp networks.
// Since: 0.9
type NetworkConfigSection struct {
	// Extends OVF Section_Type
	// FIXME: Fix the OVF section
	Info string `xml:"ovf:Info"`
	//
	HREF          string                    `xml:"href,attr,omitempty"`
	Type          string                    `xml:"type,attr,omitempty"`
	Link          *Link                     `xml:"Link,omitempty"`
	NetworkConfig *VAppNetworkConfiguration `xml:"NetworkConfig,omitempty"`
}

// NetworkConnection represents a network connection in the virtual machine.
// Type: NetworkConnectionType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a network connection in the virtual machine.
// Since: 0.9
type NetworkConnection struct {
	Network                 string `xml:"network,attr"`                      // Name of the network to which this NIC is connected.
	NetworkConnectionIndex  int    `xml:"NetworkConnectionIndex"`            // Virtual slot number associated with this NIC. First slot number is 0.
	IsConnected             bool   `xml:"IsConnected"`                       // If the virtual machine is undeployed, this value specifies whether the NIC should be connected upon deployment. If the virtual machine is deployed, this value reports the current status of this NIC's connection, and can be updated to change that connection status.
	NeedsCustomization      bool   `xml:"needsCustomization,attr,omitempty"` // True if this NIC needs customization.
	ExternalIPAddress       string `xml:"ExternalIpAddress,omitempty"`       // If the network to which this NIC connects provides NAT services, the external address assigned to this NIC appears here.
	IPAddress               string `xml:"IpAddress,omitempty"`               // IP address assigned to this NIC.
	MACAddress              string `xml:"MACAddress,omitempty"`              // MAC address associated with the NIC.
	IPAddressAllocationMode string `xml:"IpAddressAllocationMode,omitempty"` // IP address allocation mode for this connection. One of: POOL (A static IP address is allocated automatically from a pool of addresses.) DHCP (The IP address is obtained from a DHCP service.) MANUAL (The IP address is assigned manually in the IpAddress element.) NONE (No IP addressing mode specified.)
	NetworkAdapterType      string `xml:"NetworkAdapterType,omitempty"`
}

// NetworkConnectionSection the container for the network connections of this virtual machine.
// Type: NetworkConnectionSectionType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Container for the network connections of this virtual machine.
// Since: 0.9
type NetworkConnectionSection struct {
	// Extends OVF Section_Type
	// FIXME: Fix the OVF section
	Info string `xml:"ovf:Info"`
	//
	HREF                          string             `xml:"href,attr,omitempty"`
	Type                          string             `xml:"type,attr,omitempty"`
	Link                          *Link              `xml:"Link,omitempty"`
	PrimaryNetworkConnectionIndex int                `xml:"PrimaryNetworkConnectionIndex"`
	NetworkConnection             *NetworkConnection `xml:"NetworkConnection,omitempty"`
}

// NatService represents a NAT network service.
// Type: NatServiceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a NAT network service.
// Since:
type NatService struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	// Elements
	IsEnabled  bool       `xml:"IsEnabled"`            // Enable or disable the service using this flag
	NatType    string     `xml:"NatType,omitempty"`    // One of: ipTranslation (use IP translation), portForwarding (use port forwarding)
	Policy     string     `xml:"Policy,omitempty"`     // One of: allowTraffic (Allow all traffic), allowTrafficIn (Allow inbound traffic only)
	NatRule    []*NatRule `xml:"NatRule,omitempty"`    // A NAT rule.
	ExternalIP string     `xml:"ExternalIp,omitempty"` // External IP address for rule.
}

// NatRule represents a NAT rule.
// Type: NatRuleType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a NAT rule.
// Since: 0.9
type NatRule struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	// Elements
	Description        string                 `xml:"Description,omitempty"`        // A description of the rule.
	RuleType           string                 `xml:"RuleType,omitempty"`           // Type of NAT rule. One of: SNAT (source NAT), DNAT (destination NAT)
	IsEnabled          bool                   `xml:"IsEnabled"`                    // Used to enable or disable the firewall rule. Default value is true.
	ID                 string                 `xml:"Id,omitempty"`                 // Firewall rule identifier.
	GatewayNatRule     *GatewayNatRule        `xml:"GatewayNatRule,omitempty"`     // Defines SNAT and DNAT types.
	OneToOneBasicRule  *NatOneToOneBasicRule  `xml:"OneToOneBasicRule,omitempty"`  // Maps one internal IP address to one external IP address.
	OneToOneVMRule     *NatOneToOneVMRule     `xml:"OneToOneVmRule,omitempty"`     // Maps one VM NIC to one external IP addresses.
	PortForwardingRule *NatPortForwardingRule `xml:"PortForwardingRule,omitempty"` // Port forwarding internal to external IP addresses.
	VMRule             *NatVMRule             `xml:"VmRule,omitempty"`             // Port forwarding VM NIC to external IP addresses.
}

// GatewayNatRule represents the SNAT and DNAT rules.
// Type: GatewayNatRuleType represents the SNAT and DNAT rules.
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the SNAT and DNAT rules.
// Since: 5.1
type GatewayNatRule struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	// Elements
	Interface      *Reference `xml:"Interface,omitempty"`      // Interface to which rule is applied.
	OriginalIP     string     `xml:"OriginalIp"`               // Original IP for rule.
	OriginalPort   string     `xml:"OriginalPort,omitempty"`   // Original port for rule.
	TranslatedIP   string     `xml:"TranslatedIp"`             // Translated IP for rule.
	TranslatedPort string     `xml:"TranslatedPort,omitempty"` // Translated port for rule.
	Protocol       string     `xml:"Protocol,omitempty"`       // Protocol for rule.
	IcmpSubType    string     `xml:"IcmpSubType,omitempty"`    // ICMP subtype. One of: address-mask-request, address-mask-reply, destination-unreachable, echo-request, echo-reply, parameter-problem, redirect, router-advertisement, router-solicitation, source-quench, time-exceeded, timestamp-request, timestamp-reply, any.
}

// NatOneToOneBasicRule represents the NAT basic rule for one to one mapping of internal and external IP addresses from a network.
// Type: NatOneToOneBasicRuleType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the NAT basic rule for one to one mapping of internal and external IP addresses from a network.
// Since: 0.9
type NatOneToOneBasicRule struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	// Elements
	MappingMode       string `xml:"MappingMode"`       // One of: automatic (map IP addresses automatically), manual (map IP addresses manually using ExternalIpAddress and InternalIpAddress)
	ExternalIPAddress string `xml:"ExternalIpAddress"` // External IP address to map.
	InternalIPAddress string `xml:"InternalIpAddress"` // Internal IP address to map.
}

// NatOneToOneVMRule represents the NAT rule for one to one mapping of VM NIC and external IP addresses from a network.
// Type: NatOneToOneVmRuleType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the NAT rule for one to one mapping of VM NIC and external IP addresses from a network.
// Since: 0.9
type NatOneToOneVMRule struct {
	Xmlns string `xml:"xmlns,attr,omitempty"`
	// Elements
	MappingMode       string `xml:"MappingMode"`       // Mapping mode.
	ExternalIPAddress string `xml:"ExternalIpAddress"` // External IP address to map.
	VAppScopedVMID    string `xml:"VAppScopedVmId"`    // VAppScopedVmId of VM to which this rule applies.
	VMNicID           int    `xml:"VmNicId"`           // VM NIC ID to which this rule applies.
}

// NatPortForwardingRule represents the NAT rule for port forwarding between internal IP/port and external IP/port.
// Type: NatPortForwardingRuleType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the NAT rule for port forwarding between internal IP/port and external IP/port.
// Since: 0.9
type NatPortForwardingRule struct {
	ExternalIPAddress string `xml:"ExternalIpAddress"`  // External IP address to map.
	ExternalPort      int    `xml:"ExternalPort"`       // External port to forward to.
	InternalIPAddress string `xml:"InternalIpAddress"`  // Internal IP address to map.
	InternalPort      int    `xml:"InternalPort"`       // Internal port to forward to.
	Protocol          string `xml:"Protocol,omitempty"` // Protocol to forward. One of: TCP (forward TCP packets), UDP (forward UDP packets), TCP_UDP (forward TCP and UDP packets).
}

// NatVMRule represents the NAT rule for port forwarding between VM NIC/port and external IP/port.
// Type: NatVmRuleType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents the NAT rule for port forwarding between VM NIC/port and external IP/port.
// Since: 0.9
type NatVMRule struct {
	ExternalIPAddress string `xml:"ExternalIpAddress,omitempty"` // External IP address to map.
	ExternalPort      int    `xml:"ExternalPort"`                // External port to forward to.
	VAppScopedVMID    string `xml:"VAppScopedVmId"`              // VAppScopedVmId of VM to which this rule applies.
	VMNicID           int    `xml:"VmNicId"`                     // VM NIC ID to which this rule applies.
	InternalPort      int    `xml:"InternalPort"`                // Internal port to forward to.
	Protocol          string `xml:"Protocol,omitempty"`          // Protocol to forward. One of: TCP (forward TCP packets), UDP (forward UDP packets), TCP_UDP (forward TCP and UDP packets).
}

// StaticRoutingService represents Static Routing network service.
// Type: StaticRoutingServiceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents Static Routing network service.
// Since: 1.5
type StaticRoutingService struct {
	IsEnabled   bool         `xml:"IsEnabled"`             // Enable or disable the service using this flag
	StaticRoute *StaticRoute `xml:"StaticRoute,omitempty"` // Details of each Static Route.
}

// StaticRoute represents a static route entry
// Type: StaticRouteType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description:
// Since:
type StaticRoute struct {
	Name             string     `xml:"Name"`                       // Name for the static route.
	Network          string     `xml:"Network"`                    // Network specification in CIDR.
	NextHopIP        string     `xml:"NextHopIp"`                  // IP Address of Next Hop router/gateway.
	Interface        string     `xml:"Interface,omitempty"`        // Interface to use for static routing. Internal and External are the supported values.
	GatewayInterface *Reference `xml:"GatewayInterface,omitempty"` // Gateway interface to which static route is bound.
}

// LoadBalancerService represents gateway load balancer service.
// Type: LoadBalancerServiceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents gateway load balancer service.
// Since: 5.1
type LoadBalancerService struct {
	IsEnabled     bool                       `xml:"IsEnabled"`               // Enable or disable the service using this flag
	Pool          *LoadBalancerPool          `xml:"Pool,omitempty"`          // List of load balancer pools.
	VirtualServer *LoadBalancerVirtualServer `xml:"VirtualServer,omitempty"` // List of load balancer virtual servers.
}

// LoadBalancerPool represents a load balancer pool.
// Type: LoadBalancerPoolType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a load balancer pool.
// Since: 5.1
type LoadBalancerPool struct {
	ID           string             `xml:"Id,omitempty"`           // Load balancer pool id.
	Name         string             `xml:"Name"`                   // Load balancer pool name.
	Description  string             `xml:"Description,omitempty"`  // Load balancer pool description.
	ServicePort  *LBPoolServicePort `xml:"ServicePort"`            // Load balancer pool service port.
	Member       *LBPoolMember      `xml:"Member"`                 // Load balancer pool member.
	Operational  bool               `xml:"Operational,omitempty"`  // True if the load balancer pool is operational.
	ErrorDetails string             `xml:"ErrorDetails,omitempty"` // Error details for this pool.
}

// LBPoolServicePort represents a service port in a load balancer pool.
// Type: LBPoolServicePortType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a service port in a load balancer pool.
// Since: 5.1
type LBPoolServicePort struct {
	IsEnabled       bool               `xml:"IsEnabled,omitempty"`       // True if this service port is enabled.
	Protocol        string             `xml:"Protocol"`                  // Load balancer protocol type. One of: HTTP, HTTPS, TCP.
	Algorithm       string             `xml:"Algorithm"`                 // Load Balancer algorithm type. One of: IP_HASH, ROUND_ROBIN, URI, LEAST_CONN.
	Port            string             `xml:"Port"`                      // Port for this service profile.
	HealthCheckPort string             `xml:"HealthCheckPort,omitempty"` // Health check port for this profile.
	HealthCheck     *LBPoolHealthCheck `xml:"HealthCheck,omitempty"`     // Health check list.
}

// LBPoolHealthCheck represents a service port health check list.
// Type: LBPoolHealthCheckType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a service port health check list.
// Since: 5.1
type LBPoolHealthCheck struct {
	Mode              string `xml:"Mode"`                        // Load balancer service port health check mode. One of: TCP, HTTP, SSL.
	URI               string `xml:"Uri,omitempty"`               // Load balancer service port health check URI.
	HealthThreshold   string `xml:"HealthThreshold,omitempty"`   // Health threshold for this service port.
	UnhealthThreshold string `xml:"UnhealthThreshold,omitempty"` // Unhealth check port for this profile.
	Interval          string `xml:"Interval,omitempty"`          // Interval between health checks.
	Timeout           string `xml:"Timeout,omitempty"`           // Health check timeout.
}

// LBPoolMember represents a member in a load balancer pool.
// Type: LBPoolMemberType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a member in a load balancer pool.
// Since: 5.1
type LBPoolMember struct {
	IPAddress   string             `xml:"IpAddress"`             // Ip Address for load balancer member.
	Weight      string             `xml:"Weight"`                // Weight of this member.
	ServicePort *LBPoolServicePort `xml:"ServicePort,omitempty"` // Load balancer member service port.
}

// LoadBalancerVirtualServer represents a load balancer virtual server.
// Type: LoadBalancerVirtualServerType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a load balancer virtual server.
// Since: 5.1
type LoadBalancerVirtualServer struct {
	IsEnabled             bool                           `xml:"IsEnabled,omitempty"`             // True if this virtual server is enabled.
	Name                  string                         `xml:"Name"`                            // Load balancer virtual server name.
	Description           string                         `xml:"Description,omitempty"`           // Load balancer virtual server description.
	Interface             *Reference                     `xml:"Interface"`                       // Gateway Interface to which Load Balancer Virtual Server is bound.
	IPAddress             string                         `xml:"IpAddress"`                       // Load balancer virtual server Ip Address.
	ServiceProfile        *LBVirtualServerServiceProfile `xml:"ServiceProfile"`                  // Load balancer virtual server service profiles.
	Logging               bool                           `xml:"Logging,omitempty"`               // Enable logging for this virtual server.
	Pool                  string                         `xml:"Pool"`                            // Name of Load balancer pool associated with this virtual server.
	LoadBalancerTemplates *VendorTemplate                `xml:"LoadBalancerTemplates,omitempty"` // Service template related attributes.
}

// LBVirtualServerServiceProfile represents service profile for a load balancing virtual server.
// Type: LBVirtualServerServiceProfileType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents service profile for a load balancing virtual server.
// Since: 5.1
type LBVirtualServerServiceProfile struct {
	IsEnabled   bool           `xml:"IsEnabled,omitempty"`   // True if this service profile is enabled.
	Protocol    string         `xml:"Protocol"`              // Load balancer Protocol type. One of: HTTP, HTTPS, TCP.
	Port        string         `xml:"Port"`                  // Port for this service profile.
	Persistence *LBPersistence `xml:"Persistence,omitempty"` // Persistence type for service profile.
}

// LBPersistence represents persistence type for a load balancer service profile.
// Type: LBPersistenceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents persistence type for a load balancer service profile.
// Since: 5.1
type LBPersistence struct {
	Method     string `xml:"Method"`               // Persistence method. One of: COOKIE, SSL_SESSION_ID.
	CookieName string `xml:"CookieName,omitempty"` // Cookie name when persistence method is COOKIE.
	CookieMode string `xml:"CookieMode,omitempty"` // Cookie Mode. One of: INSERT, PREFIX, APP.
}

// VendorTemplate is information about a vendor service template. This is optional.
// Type: VendorTemplateType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Information about a vendor service template. This is optional.
// Since: 5.1
type VendorTemplate struct {
	Name string `xml:"Name"` // Name of the vendor template. This is required.
	ID   string `xml:"Id"`   // ID of the vendor template. This is required.
}

// GatewayIpsecVpnService represents gateway IPsec VPN service.
// Type: GatewayIpsecVpnServiceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents gateway IPsec VPN service.
// Since: 5.1
type GatewayIpsecVpnService struct {
	IsEnabled bool                     `xml:"IsEnabled"`          // Enable or disable the service using this flag
	Endpoint  *GatewayIpsecVpnEndpoint `xml:"Endpoint,omitempty"` // List of IPSec VPN Service Endpoints.
	Tunnel    []*GatewayIpsecVpnTunnel `xml:"Tunnel"`             // List of IPSec VPN tunnels.
}

// GatewayIpsecVpnEndpoint represents an IPSec VPN endpoint.
// Type: GatewayIpsecVpnEndpointType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents an IPSec VPN endpoint.
// Since: 5.1
type GatewayIpsecVpnEndpoint struct {
	Network  *Reference `xml:"Network"`            // External network reference.
	PublicIP string     `xml:"PublicIp,omitempty"` // Public IP for IPSec endpoint.
}

// GatewayIpsecVpnTunnel represents an IPSec VPN tunnel.
// Type: GatewayIpsecVpnTunnelType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents an IPSec VPN tunnel.
// Since: 5.1
type GatewayIpsecVpnTunnel struct {
	Name        string `xml:"Name"`                  // The name of the tunnel.
	Description string `xml:"Description,omitempty"` // A description of the tunnel.
	// TODO: Fix this in a better way
	IpsecVpnThirdPartyPeer *IpsecVpnThirdPartyPeer `xml:"IpsecVpnThirdPartyPeer,omitempty"` // Details about the peer network.
	PeerIPAddress          string                  `xml:"PeerIpAddress"`                    // IP address of the peer endpoint.
	PeerID                 string                  `xml:"PeerId"`                           // Id for the peer end point
	LocalIPAddress         string                  `xml:"LocalIpAddress"`                   // Address of the local network.
	LocalID                string                  `xml:"LocalId"`                          // Id for local end point
	LocalSubnet            *IpsecVpnSubnet         `xml:"LocalSubnet"`                      // List of local subnets in the tunnel.
	PeerSubnet             *IpsecVpnSubnet         `xml:"PeerSubnet"`                       // List of peer subnets in the tunnel.
	SharedSecret           string                  `xml:"SharedSecret"`                     // Shared secret used for authentication.
	SharedSecretEncrypted  bool                    `xml:"SharedSecretEncrypted,omitempty"`  // True if shared secret is encrypted.
	EncryptionProtocol     string                  `xml:"EncryptionProtocol"`               // Encryption protocol to be used. One of: AES, AES256, TRIPLEDES
	Mtu                    int                     `xml:"Mtu"`                              // MTU for the tunnel.
	IsEnabled              bool                    `xml:"IsEnabled,omitempty"`              // True if the tunnel is enabled.
	IsOperational          bool                    `xml:"IsOperational,omitempty"`          // True if the tunnel is operational.
	ErrorDetails           string                  `xml:"ErrorDetails,omitempty"`           // Error details of the tunnel.
}

// IpsecVpnThirdPartyPeer represents details about a peer network
type IpsecVpnThirdPartyPeer struct {
	PeerID string `xml:"PeerId,omitempty"` // Id for the peer end point
}

// IpsecVpnSubnet represents subnet details.
// Type: IpsecVpnSubnetType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents subnet details.
// Since: 5.1
type IpsecVpnSubnet struct {
	Name    string `xml:"Name"`    // Gateway Name.
	Gateway string `xml:"Gateway"` // Subnet Gateway.
	Netmask string `xml:"Netmask"` // Subnet Netmask.
}
