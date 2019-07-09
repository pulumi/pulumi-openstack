// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
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
// Ports. Please see the `openstack_compute_instance_v2` documentation for further
// documentation.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_port_v2.html.markdown.
type Port struct {
	s *pulumi.ResourceState
}

// NewPort registers a new resource with the given unique name, arguments, and options.
func NewPort(ctx *pulumi.Context,
	name string, args *PortArgs, opts ...pulumi.ResourceOpt) (*Port, error) {
	if args == nil || args.NetworkId == nil {
		return nil, errors.New("missing required argument 'NetworkId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adminStateUp"] = nil
		inputs["allowedAddressPairs"] = nil
		inputs["binding"] = nil
		inputs["description"] = nil
		inputs["deviceId"] = nil
		inputs["deviceOwner"] = nil
		inputs["dnsName"] = nil
		inputs["extraDhcpOptions"] = nil
		inputs["fixedIps"] = nil
		inputs["macAddress"] = nil
		inputs["name"] = nil
		inputs["networkId"] = nil
		inputs["noFixedIp"] = nil
		inputs["noSecurityGroups"] = nil
		inputs["portSecurityEnabled"] = nil
		inputs["region"] = nil
		inputs["securityGroupIds"] = nil
		inputs["tags"] = nil
		inputs["tenantId"] = nil
		inputs["valueSpecs"] = nil
	} else {
		inputs["adminStateUp"] = args.AdminStateUp
		inputs["allowedAddressPairs"] = args.AllowedAddressPairs
		inputs["binding"] = args.Binding
		inputs["description"] = args.Description
		inputs["deviceId"] = args.DeviceId
		inputs["deviceOwner"] = args.DeviceOwner
		inputs["dnsName"] = args.DnsName
		inputs["extraDhcpOptions"] = args.ExtraDhcpOptions
		inputs["fixedIps"] = args.FixedIps
		inputs["macAddress"] = args.MacAddress
		inputs["name"] = args.Name
		inputs["networkId"] = args.NetworkId
		inputs["noFixedIp"] = args.NoFixedIp
		inputs["noSecurityGroups"] = args.NoSecurityGroups
		inputs["portSecurityEnabled"] = args.PortSecurityEnabled
		inputs["region"] = args.Region
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["tags"] = args.Tags
		inputs["tenantId"] = args.TenantId
		inputs["valueSpecs"] = args.ValueSpecs
	}
	inputs["allFixedIps"] = nil
	inputs["allSecurityGroupIds"] = nil
	inputs["allTags"] = nil
	inputs["dnsAssignments"] = nil
	s, err := ctx.RegisterResource("openstack:networking/port:Port", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Port{s: s}, nil
}

// GetPort gets an existing Port resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPort(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PortState, opts ...pulumi.ResourceOpt) (*Port, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adminStateUp"] = state.AdminStateUp
		inputs["allFixedIps"] = state.AllFixedIps
		inputs["allSecurityGroupIds"] = state.AllSecurityGroupIds
		inputs["allTags"] = state.AllTags
		inputs["allowedAddressPairs"] = state.AllowedAddressPairs
		inputs["binding"] = state.Binding
		inputs["description"] = state.Description
		inputs["deviceId"] = state.DeviceId
		inputs["deviceOwner"] = state.DeviceOwner
		inputs["dnsAssignments"] = state.DnsAssignments
		inputs["dnsName"] = state.DnsName
		inputs["extraDhcpOptions"] = state.ExtraDhcpOptions
		inputs["fixedIps"] = state.FixedIps
		inputs["macAddress"] = state.MacAddress
		inputs["name"] = state.Name
		inputs["networkId"] = state.NetworkId
		inputs["noFixedIp"] = state.NoFixedIp
		inputs["noSecurityGroups"] = state.NoSecurityGroups
		inputs["portSecurityEnabled"] = state.PortSecurityEnabled
		inputs["region"] = state.Region
		inputs["securityGroupIds"] = state.SecurityGroupIds
		inputs["tags"] = state.Tags
		inputs["tenantId"] = state.TenantId
		inputs["valueSpecs"] = state.ValueSpecs
	}
	s, err := ctx.ReadResource("openstack:networking/port:Port", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Port{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Port) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Port) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Administrative up/down status for the port
// (must be "true" or "false" if provided). Changing this updates the
// `admin_state_up` of an existing port.
func (r *Port) AdminStateUp() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["adminStateUp"])
}

// The collection of Fixed IP addresses on the port in the
// order returned by the Network v2 API.
func (r *Port) AllFixedIps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allFixedIps"])
}

// The collection of Security Group IDs on the port
// which have been explicitly and implicitly added.
func (r *Port) AllSecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allSecurityGroupIds"])
}

// The collection of tags assigned on the port, which have been
// explicitly and implicitly added.
func (r *Port) AllTags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allTags"])
}

// An IP/MAC Address pair of additional IP
// addresses that can be active on this port. The structure is described
// below.
func (r *Port) AllowedAddressPairs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allowedAddressPairs"])
}

// The port binding allows to specify binding information
// for the port. The structure is described below.
func (r *Port) Binding() *pulumi.Output {
	return r.s.State["binding"]
}

// Human-readable description of the floating IP. Changing
// this updates the `description` of an existing port.
func (r *Port) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The ID of the device attached to the port. Changing this
// creates a new port.
func (r *Port) DeviceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deviceId"])
}

// The device owner of the Port. Changing this creates
// a new port.
func (r *Port) DeviceOwner() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deviceOwner"])
}

// The list of maps representing port DNS assignments.
func (r *Port) DnsAssignments() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dnsAssignments"])
}

// The port DNS name. Available, when Neutron DNS extension
// is enabled.
func (r *Port) DnsName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsName"])
}

// An extra DHCP option that needs to be configured
// on the port. The structure is described below. Can be specified multiple
// times.
func (r *Port) ExtraDhcpOptions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["extraDhcpOptions"])
}

// An array of desired IPs for
// this port. The structure is described below.
func (r *Port) FixedIps() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["fixedIps"])
}

// Specify a specific MAC address for the port. Changing
// this creates a new port.
func (r *Port) MacAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["macAddress"])
}

// A unique name for the port. Changing this
// updates the `name` of an existing port.
func (r *Port) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the network to attach the port to. Changing
// this creates a new port.
func (r *Port) NetworkId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkId"])
}

// Create a port with no fixed
// IP address. This will also remove any fixed IPs previously set on a port. `true`
// is the only valid value for this argument.
func (r *Port) NoFixedIp() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["noFixedIp"])
}

// If set to
// `true`, then no security groups are applied to the port. If set to `false` and
// no `security_group_ids` are specified, then the Port will yield to the default
// behavior of the Networking service, which is to usually apply the "default"
// security group.
func (r *Port) NoSecurityGroups() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["noSecurityGroups"])
}

// Whether to explicitly enable or disable
// port security on the port. Port Security is usually enabled by default, so
// omitting argument will usually result in a value of "true". Setting this
// explicitly to `false` will disable port security. In order to disable port
// security, the port must not have any security groups. Valid values are `true`
// and `false`.
func (r *Port) PortSecurityEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["portSecurityEnabled"])
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to create a port. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// port.
func (r *Port) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// A list
// of security group IDs to apply to the port. The security groups must be
// specified by ID and not name (as opposed to how they are configured with
// the Compute Instance).
func (r *Port) SecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["securityGroupIds"])
}

// A set of string tags for the port.
func (r *Port) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// The owner of the Port. Required if admin wants
// to create a port for another tenant. Changing this creates a new port.
func (r *Port) TenantId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenantId"])
}

// Map of additional options.
func (r *Port) ValueSpecs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["valueSpecs"])
}

// Input properties used for looking up and filtering Port resources.
type PortState struct {
	// Administrative up/down status for the port
	// (must be "true" or "false" if provided). Changing this updates the
	// `admin_state_up` of an existing port.
	AdminStateUp interface{}
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps interface{}
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds interface{}
	// The collection of tags assigned on the port, which have been
	// explicitly and implicitly added.
	AllTags interface{}
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs interface{}
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding interface{}
	// Human-readable description of the floating IP. Changing
	// this updates the `description` of an existing port.
	Description interface{}
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId interface{}
	// The device owner of the Port. Changing this creates
	// a new port.
	DeviceOwner interface{}
	// The list of maps representing port DNS assignments.
	DnsAssignments interface{}
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName interface{}
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions interface{}
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps interface{}
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress interface{}
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name interface{}
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId interface{}
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp interface{}
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `security_group_ids` are specified, then the Port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups interface{}
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled interface{}
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region interface{}
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds interface{}
	// A set of string tags for the port.
	Tags interface{}
	// The owner of the Port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId interface{}
	// Map of additional options.
	ValueSpecs interface{}
}

// The set of arguments for constructing a Port resource.
type PortArgs struct {
	// Administrative up/down status for the port
	// (must be "true" or "false" if provided). Changing this updates the
	// `admin_state_up` of an existing port.
	AdminStateUp interface{}
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs interface{}
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding interface{}
	// Human-readable description of the floating IP. Changing
	// this updates the `description` of an existing port.
	Description interface{}
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId interface{}
	// The device owner of the Port. Changing this creates
	// a new port.
	DeviceOwner interface{}
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName interface{}
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions interface{}
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps interface{}
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress interface{}
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name interface{}
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId interface{}
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp interface{}
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `security_group_ids` are specified, then the Port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups interface{}
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled interface{}
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region interface{}
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds interface{}
	// A set of string tags for the port.
	Tags interface{}
	// The owner of the Port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId interface{}
	// Map of additional options.
	ValueSpecs interface{}
}
