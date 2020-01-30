// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package networking

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 port resource within OpenStack.
//
// ## Notes
//
// ### Ports and Instances
//
// There are some notes to consider when connecting Instances to networks using
// Ports. Please see the `compute.Instance` documentation for further
// documentation.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_port_v2.html.markdown.
type Port struct {
	pulumi.CustomResourceState

	// Administrative up/down status for the port
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp pulumi.BoolOutput `pulumi:"adminStateUp"`
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps pulumi.StringArrayOutput `pulumi:"allFixedIps"`
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds pulumi.StringArrayOutput `pulumi:"allSecurityGroupIds"`
	// The collection of tags assigned on the port, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs PortAllowedAddressPairArrayOutput `pulumi:"allowedAddressPairs"`
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding PortBindingOutput `pulumi:"binding"`
	// Human-readable description of the floating IP. Changing
	// this updates the `description` of an existing port.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// The device owner of the Port. Changing this creates
	// a new port.
	DeviceOwner pulumi.StringOutput `pulumi:"deviceOwner"`
	// The list of maps representing port DNS assignments.
	DnsAssignments pulumi.MapArrayOutput `pulumi:"dnsAssignments"`
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions PortExtraDhcpOptionArrayOutput `pulumi:"extraDhcpOptions"`
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps PortFixedIpArrayOutput `pulumi:"fixedIps"`
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp pulumi.BoolPtrOutput `pulumi:"noFixedIp"`
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the Port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups pulumi.BoolPtrOutput `pulumi:"noSecurityGroups"`
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled pulumi.BoolOutput `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringOutput `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A set of string tags for the port.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The owner of the Port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewPort registers a new resource with the given unique name, arguments, and options.
func NewPort(ctx *pulumi.Context,
	name string, args *PortArgs, opts ...pulumi.ResourceOption) (*Port, error) {
	if args == nil || args.NetworkId == nil {
		return nil, errors.New("missing required argument 'NetworkId'")
	}
	if args == nil {
		args = &PortArgs{}
	}
	var resource Port
	err := ctx.RegisterResource("openstack:networking/port:Port", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPort gets an existing Port resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortState, opts ...pulumi.ResourceOption) (*Port, error) {
	var resource Port
	err := ctx.ReadResource("openstack:networking/port:Port", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Port resources.
type portState struct {
	// Administrative up/down status for the port
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps []string `pulumi:"allFixedIps"`
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds []string `pulumi:"allSecurityGroupIds"`
	// The collection of tags assigned on the port, which have been
	// explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs []PortAllowedAddressPair `pulumi:"allowedAddressPairs"`
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding *PortBinding `pulumi:"binding"`
	// Human-readable description of the floating IP. Changing
	// this updates the `description` of an existing port.
	Description *string `pulumi:"description"`
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId *string `pulumi:"deviceId"`
	// The device owner of the Port. Changing this creates
	// a new port.
	DeviceOwner *string `pulumi:"deviceOwner"`
	// The list of maps representing port DNS assignments.
	DnsAssignments []map[string]interface{} `pulumi:"dnsAssignments"`
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName *string `pulumi:"dnsName"`
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions []PortExtraDhcpOption `pulumi:"extraDhcpOptions"`
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps []PortFixedIp `pulumi:"fixedIps"`
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress *string `pulumi:"macAddress"`
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name *string `pulumi:"name"`
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId *string `pulumi:"networkId"`
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp *bool `pulumi:"noFixedIp"`
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the Port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups *bool `pulumi:"noSecurityGroups"`
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled *bool `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId *string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region *string `pulumi:"region"`
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A set of string tags for the port.
	Tags []string `pulumi:"tags"`
	// The owner of the Port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type PortState struct {
	// Administrative up/down status for the port
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp pulumi.BoolPtrInput
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps pulumi.StringArrayInput
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds pulumi.StringArrayInput
	// The collection of tags assigned on the port, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs PortAllowedAddressPairArrayInput
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding PortBindingPtrInput
	// Human-readable description of the floating IP. Changing
	// this updates the `description` of an existing port.
	Description pulumi.StringPtrInput
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId pulumi.StringPtrInput
	// The device owner of the Port. Changing this creates
	// a new port.
	DeviceOwner pulumi.StringPtrInput
	// The list of maps representing port DNS assignments.
	DnsAssignments pulumi.MapArrayInput
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName pulumi.StringPtrInput
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions PortExtraDhcpOptionArrayInput
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps PortFixedIpArrayInput
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress pulumi.StringPtrInput
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name pulumi.StringPtrInput
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId pulumi.StringPtrInput
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp pulumi.BoolPtrInput
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the Port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups pulumi.BoolPtrInput
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled pulumi.BoolPtrInput
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region pulumi.StringPtrInput
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds pulumi.StringArrayInput
	// A set of string tags for the port.
	Tags pulumi.StringArrayInput
	// The owner of the Port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (PortState) ElementType() reflect.Type {
	return reflect.TypeOf((*portState)(nil)).Elem()
}

type portArgs struct {
	// Administrative up/down status for the port
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs []PortAllowedAddressPair `pulumi:"allowedAddressPairs"`
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding *PortBinding `pulumi:"binding"`
	// Human-readable description of the floating IP. Changing
	// this updates the `description` of an existing port.
	Description *string `pulumi:"description"`
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId *string `pulumi:"deviceId"`
	// The device owner of the Port. Changing this creates
	// a new port.
	DeviceOwner *string `pulumi:"deviceOwner"`
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName *string `pulumi:"dnsName"`
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions []PortExtraDhcpOption `pulumi:"extraDhcpOptions"`
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps []PortFixedIp `pulumi:"fixedIps"`
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress *string `pulumi:"macAddress"`
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name *string `pulumi:"name"`
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId string `pulumi:"networkId"`
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp *bool `pulumi:"noFixedIp"`
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the Port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups *bool `pulumi:"noSecurityGroups"`
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled *bool `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId *string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region *string `pulumi:"region"`
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A set of string tags for the port.
	Tags []string `pulumi:"tags"`
	// The owner of the Port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Port resource.
type PortArgs struct {
	// Administrative up/down status for the port
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp pulumi.BoolPtrInput
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs PortAllowedAddressPairArrayInput
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding PortBindingPtrInput
	// Human-readable description of the floating IP. Changing
	// this updates the `description` of an existing port.
	Description pulumi.StringPtrInput
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId pulumi.StringPtrInput
	// The device owner of the Port. Changing this creates
	// a new port.
	DeviceOwner pulumi.StringPtrInput
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName pulumi.StringPtrInput
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions PortExtraDhcpOptionArrayInput
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps PortFixedIpArrayInput
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress pulumi.StringPtrInput
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name pulumi.StringPtrInput
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId pulumi.StringInput
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp pulumi.BoolPtrInput
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the Port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups pulumi.BoolPtrInput
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled pulumi.BoolPtrInput
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region pulumi.StringPtrInput
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds pulumi.StringArrayInput
	// A set of string tags for the port.
	Tags pulumi.StringArrayInput
	// The owner of the Port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (PortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portArgs)(nil)).Elem()
}
