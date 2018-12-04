// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 Neutron subnet resource within OpenStack.
type Subnet struct {
	s *pulumi.ResourceState
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOpt) (*Subnet, error) {
	if args == nil || args.NetworkId == nil {
		return nil, errors.New("missing required argument 'NetworkId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allocationPools"] = nil
		inputs["cidr"] = nil
		inputs["dnsNameservers"] = nil
		inputs["enableDhcp"] = nil
		inputs["gatewayIp"] = nil
		inputs["hostRoutes"] = nil
		inputs["ipVersion"] = nil
		inputs["ipv6AddressMode"] = nil
		inputs["ipv6RaMode"] = nil
		inputs["name"] = nil
		inputs["networkId"] = nil
		inputs["noGateway"] = nil
		inputs["region"] = nil
		inputs["subnetpoolId"] = nil
		inputs["tags"] = nil
		inputs["tenantId"] = nil
		inputs["valueSpecs"] = nil
	} else {
		inputs["allocationPools"] = args.AllocationPools
		inputs["cidr"] = args.Cidr
		inputs["dnsNameservers"] = args.DnsNameservers
		inputs["enableDhcp"] = args.EnableDhcp
		inputs["gatewayIp"] = args.GatewayIp
		inputs["hostRoutes"] = args.HostRoutes
		inputs["ipVersion"] = args.IpVersion
		inputs["ipv6AddressMode"] = args.Ipv6AddressMode
		inputs["ipv6RaMode"] = args.Ipv6RaMode
		inputs["name"] = args.Name
		inputs["networkId"] = args.NetworkId
		inputs["noGateway"] = args.NoGateway
		inputs["region"] = args.Region
		inputs["subnetpoolId"] = args.SubnetpoolId
		inputs["tags"] = args.Tags
		inputs["tenantId"] = args.TenantId
		inputs["valueSpecs"] = args.ValueSpecs
	}
	s, err := ctx.RegisterResource("openstack:networking/subnet:Subnet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subnet{s: s}, nil
}

// GetSubnet gets an existing Subnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubnetState, opts ...pulumi.ResourceOpt) (*Subnet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allocationPools"] = state.AllocationPools
		inputs["cidr"] = state.Cidr
		inputs["dnsNameservers"] = state.DnsNameservers
		inputs["enableDhcp"] = state.EnableDhcp
		inputs["gatewayIp"] = state.GatewayIp
		inputs["hostRoutes"] = state.HostRoutes
		inputs["ipVersion"] = state.IpVersion
		inputs["ipv6AddressMode"] = state.Ipv6AddressMode
		inputs["ipv6RaMode"] = state.Ipv6RaMode
		inputs["name"] = state.Name
		inputs["networkId"] = state.NetworkId
		inputs["noGateway"] = state.NoGateway
		inputs["region"] = state.Region
		inputs["subnetpoolId"] = state.SubnetpoolId
		inputs["tags"] = state.Tags
		inputs["tenantId"] = state.TenantId
		inputs["valueSpecs"] = state.ValueSpecs
	}
	s, err := ctx.ReadResource("openstack:networking/subnet:Subnet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subnet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Subnet) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Subnet) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// An array of sub-ranges of CIDR available for
// dynamic allocation to ports. The allocation_pool object structure is
// documented below. Changing this creates a new subnet.
func (r *Subnet) AllocationPools() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allocationPools"])
}

// CIDR representing IP range for this subnet, based on IP
// version. You can omit this option if you are creating a subnet from a
// subnet pool.
func (r *Subnet) Cidr() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cidr"])
}

// An array of DNS name server names used by hosts
// in this subnet. Changing this updates the DNS name servers for the existing
// subnet.
func (r *Subnet) DnsNameservers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dnsNameservers"])
}

// The administrative state of the network.
// Acceptable values are "true" and "false". Changing this value enables or
// disables the DHCP capabilities of the existing subnet. Defaults to true.
func (r *Subnet) EnableDhcp() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableDhcp"])
}

// Default gateway used by devices in this subnet.
// Leaving this blank and not setting `no_gateway` will cause a default
// gateway of `.1` to be used. Changing this updates the gateway IP of the
// existing subnet.
func (r *Subnet) GatewayIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["gatewayIp"])
}

// An array of routes that should be used by devices
// with IPs from this subnet (not including local subnet route). The host_route
// object structure is documented below. Changing this updates the host routes
// for the existing subnet.
func (r *Subnet) HostRoutes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["hostRoutes"])
}

// IP version, either 4 (default) or 6. Changing this creates a
// new subnet.
func (r *Subnet) IpVersion() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ipVersion"])
}

// The IPv6 address mode. Valid values are
// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
func (r *Subnet) Ipv6AddressMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6AddressMode"])
}

// The IPv6 Router Advertisement mode. Valid values
// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
func (r *Subnet) Ipv6RaMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6RaMode"])
}

// The name of the subnet. Changing this updates the name of
// the existing subnet.
func (r *Subnet) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The UUID of the parent network. Changing this
// creates a new subnet.
func (r *Subnet) NetworkId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkId"])
}

// Do not set a gateway IP on this subnet. Changing
// this removes or adds a default gateway IP of the existing subnet.
func (r *Subnet) NoGateway() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["noGateway"])
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a Neutron subnet. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// subnet.
func (r *Subnet) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The ID of the subnetpool associated with the subnet.
func (r *Subnet) SubnetpoolId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetpoolId"])
}

// A set of string tags for the subnet.
func (r *Subnet) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// The owner of the subnet. Required if admin wants to
// create a subnet for another tenant. Changing this creates a new subnet.
func (r *Subnet) TenantId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenantId"])
}

// Map of additional options.
func (r *Subnet) ValueSpecs() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["valueSpecs"])
}

// Input properties used for looking up and filtering Subnet resources.
type SubnetState struct {
	// An array of sub-ranges of CIDR available for
	// dynamic allocation to ports. The allocation_pool object structure is
	// documented below. Changing this creates a new subnet.
	AllocationPools interface{}
	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr interface{}
	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DnsNameservers interface{}
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDhcp interface{}
	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting `no_gateway` will cause a default
	// gateway of `.1` to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIp interface{}
	// An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	HostRoutes interface{}
	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IpVersion interface{}
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode interface{}
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode interface{}
	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name interface{}
	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkId interface{}
	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnet.
	Region interface{}
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId interface{}
	// A set of string tags for the subnet.
	Tags interface{}
	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantId interface{}
	// Map of additional options.
	ValueSpecs interface{}
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// An array of sub-ranges of CIDR available for
	// dynamic allocation to ports. The allocation_pool object structure is
	// documented below. Changing this creates a new subnet.
	AllocationPools interface{}
	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr interface{}
	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DnsNameservers interface{}
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDhcp interface{}
	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting `no_gateway` will cause a default
	// gateway of `.1` to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIp interface{}
	// An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	HostRoutes interface{}
	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IpVersion interface{}
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode interface{}
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode interface{}
	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name interface{}
	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkId interface{}
	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnet.
	Region interface{}
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId interface{}
	// A set of string tags for the subnet.
	Tags interface{}
	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantId interface{}
	// Map of additional options.
	ValueSpecs interface{}
}
