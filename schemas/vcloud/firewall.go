package vcloud

// VMSelection represents details of an vm+nic+iptype selection.
// Type: VmSelectionType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents details of an vm+nic+iptype selection.
// Since: 5.1
type VMSelection struct {
	VAppScopedVMID string `xml:"VAppScopedVmId"` // VAppScopedVmId of VM to which this rule applies.
	VMNicID        int    `xml:"VmNicId"`        // VM NIC ID to which this rule applies.
	IPType         string `xml:"IpType"`         // The value can be one of:- assigned: assigned internal IP be automatically choosen. NAT: NATed external IP will be automatically choosen.
}

// FirewallRuleProtocols flags for a network protocol in a firewall rule
// Type: FirewallRuleType/Protocols
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description:
// Since:
type FirewallRuleProtocols struct {
	ICMP bool `xml:"Icmp,omitempty"` // True if the rule applies to the ICMP protocol.
	Any  bool `xml:"Any,omitempty"`  // True if the rule applies to any protocol.
	TCP  bool `xml:"Tcp,omitempty"`  // True if the rule applies to the TCP protocol.
	UDP  bool `xml:"Udp,omitempty"`  // True if the rule applies to the UDP protocol.
	// FIXME: this is supposed to extend protocol support to all the VSM supported protocols
	// Other string `xml:"Other,omitempty"` //	Any other protocol supported by vShield Manager
}

// FirewallRule represents a firewall rule
// Type: FirewallRuleType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a firewall rule.
// Since: 0.9
type FirewallRule struct {
	ID                   string                 `xml:"Id,omitempty"`                   // Firewall rule identifier.
	IsEnabled            bool                   `xml:"IsEnabled"`                      // Used to enable or disable the firewall rule. Default value is true.
	MatchOnTranslate     bool                   `xml:"MatchOnTranslate"`               // For DNATed traffic, match the firewall rules only after the destination IP is translated.
	Description          string                 `xml:"Description,omitempty"`          // A description of the rule.
	Policy               string                 `xml:"Policy,omitempty"`               // One of: drop (drop packets that match the rule), allow (allow packets that match the rule to pass through the firewall)
	Protocols            *FirewallRuleProtocols `xml:"Protocols,omitempty"`            // Specify the protocols to which the rule should be applied.
	IcmpSubType          string                 `xml:"IcmpSubType,omitempty"`          // ICMP subtype. One of: address-mask-request, address-mask-reply, destination-unreachable, echo-request, echo-reply, parameter-problem, redirect, router-advertisement, router-solicitation, source-quench, time-exceeded, timestamp-request, timestamp-reply, any.
	Port                 int                    `xml:"Port,omitempty"`                 // The port to which this rule applies. A value of -1 matches any port.
	DestinationPortRange string                 `xml:"DestinationPortRange,omitempty"` // Destination port range to which this rule applies.
	DestinationIP        string                 `xml:"DestinationIp,omitempty"`        // Destination IP address to which the rule applies. A value of Any matches any IP address.
	DestinationVM        *VMSelection           `xml:"DestinationVm,omitempty"`        // Details of the destination VM
	SourcePort           int                    `xml:"SourcePort,omitempty"`           // Destination port to which this rule applies. A value of -1 matches any port.
	SourcePortRange      string                 `xml:"SourcePortRange,omitempty"`      // Source port range to which this rule applies.
	SourceIP             string                 `xml:"SourceIp,omitempty"`             // Source IP address to which the rule applies. A value of Any matches any IP address.
	SourceVM             *VMSelection           `xml:"SourceVm,omitempty"`             // Details of the source Vm
	Direction            string                 `xml:"Direction,omitempty"`            // Direction of traffic to which rule applies. One of: in (rule applies to incoming traffic. This is the default value), out (rule applies to outgoing traffic).
	EnableLogging        bool                   `xml:"EnableLogging"`                  // Used to enable or disable firewall rule logging. Default value is false.
}

// FirewallService represent a network firewall service.
// Type: FirewallServiceType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Represents a network firewall service.
// Since:
type FirewallService struct {
	IsEnabled        bool            `xml:"IsEnabled"`               // Enable or disable the service using this flag
	DefaultAction    string          `xml:"DefaultAction,omitempty"` // Default action of the firewall. One of: drop (Default. Drop packets that match the rule.), allow (Allow packets that match the rule to pass through the firewall)
	LogDefaultAction bool            `xml:"LogDefaultAction"`        // Flag to enable logging for default action. Default value is false.
	FirewallRule     []*FirewallRule `xml:"FirewallRule,omitempty"`  //	A firewall rule.
}
