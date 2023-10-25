// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a V2 Neutron subnet resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network1, err := networking.NewNetwork(ctx, "network1", &networking.NetworkArgs{
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewSubnet(ctx, "subnet1", &networking.SubnetArgs{
//				NetworkId: network1.ID(),
//				Cidr:      pulumi.String("192.168.199.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Subnets can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:networking/subnet:Subnet subnet_1 da4faf16-5546-41e4-8330-4d0002b74048
//
// ```
type Subnet struct {
	pulumi.CustomResourceState

	// The collection of ags assigned on the subnet, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// A block declaring the start and end range of
	// the IP addresses available for use with DHCP in this subnet. Multiple
	// `allocationPool` blocks can be declared, providing the subnet with more
	// than one range of IP addresses to use with DHCP. However, each IP range
	// must be from the same CIDR that the subnet is part of.
	// The `allocationPool` block is documented below.
	AllocationPools SubnetAllocationPoolArrayOutput `pulumi:"allocationPools"`
	// A block declaring the start and end range of the IP addresses available for
	// use with DHCP in this subnet.
	// The `allocationPools` block is documented below.
	//
	// Deprecated: use allocation_pool instead
	AllocationPoolsCollection SubnetAllocationPoolsCollectionArrayOutput `pulumi:"allocationPoolsCollection"`
	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// Human-readable description of the subnet. Changing this
	// updates the name of the existing subnet.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DnsNameservers pulumi.StringArrayOutput `pulumi:"dnsNameservers"`
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDhcp pulumi.BoolPtrOutput `pulumi:"enableDhcp"`
	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting `noGateway` will cause a default
	// gateway of `.1` to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIp pulumi.StringOutput `pulumi:"gatewayIp"`
	// (**Deprecated** - use `networking.SubnetRoute`
	// instead) An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The hostRoute
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	//
	// Deprecated: Use openstack_networking_subnet_route_v2 instead
	HostRoutes SubnetHostRouteArrayOutput `pulumi:"hostRoutes"`
	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IpVersion pulumi.IntPtrOutput `pulumi:"ipVersion"`
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode pulumi.StringOutput `pulumi:"ipv6AddressMode"`
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode pulumi.StringOutput `pulumi:"ipv6RaMode"`
	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway pulumi.BoolPtrOutput `pulumi:"noGateway"`
	// The prefix length to use when creating a subnet
	// from a subnet pool. The default subnet pool prefix length that was defined
	// when creating the subnet pool will be used if not provided. Changing this
	// creates a new subnet.
	PrefixLength pulumi.IntPtrOutput `pulumi:"prefixLength"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnet.
	Region pulumi.StringOutput `pulumi:"region"`
	// An array of service types used by the subnet.
	// Changing this updates the service types for the existing subnet.
	ServiceTypes pulumi.StringArrayOutput `pulumi:"serviceTypes"`
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId pulumi.StringPtrOutput `pulumi:"subnetpoolId"`
	// A set of string tags for the subnet.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOption) (*Subnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Subnet
	err := ctx.RegisterResource("openstack:networking/subnet:Subnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnet gets an existing Subnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetState, opts ...pulumi.ResourceOption) (*Subnet, error) {
	var resource Subnet
	err := ctx.ReadResource("openstack:networking/subnet:Subnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnet resources.
type subnetState struct {
	// The collection of ags assigned on the subnet, which have been
	// explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// A block declaring the start and end range of
	// the IP addresses available for use with DHCP in this subnet. Multiple
	// `allocationPool` blocks can be declared, providing the subnet with more
	// than one range of IP addresses to use with DHCP. However, each IP range
	// must be from the same CIDR that the subnet is part of.
	// The `allocationPool` block is documented below.
	AllocationPools []SubnetAllocationPool `pulumi:"allocationPools"`
	// A block declaring the start and end range of the IP addresses available for
	// use with DHCP in this subnet.
	// The `allocationPools` block is documented below.
	//
	// Deprecated: use allocation_pool instead
	AllocationPoolsCollection []SubnetAllocationPoolsCollection `pulumi:"allocationPoolsCollection"`
	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr *string `pulumi:"cidr"`
	// Human-readable description of the subnet. Changing this
	// updates the name of the existing subnet.
	Description *string `pulumi:"description"`
	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DnsNameservers []string `pulumi:"dnsNameservers"`
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDhcp *bool `pulumi:"enableDhcp"`
	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting `noGateway` will cause a default
	// gateway of `.1` to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIp *string `pulumi:"gatewayIp"`
	// (**Deprecated** - use `networking.SubnetRoute`
	// instead) An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The hostRoute
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	//
	// Deprecated: Use openstack_networking_subnet_route_v2 instead
	HostRoutes []SubnetHostRoute `pulumi:"hostRoutes"`
	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IpVersion *int `pulumi:"ipVersion"`
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode *string `pulumi:"ipv6AddressMode"`
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode *string `pulumi:"ipv6RaMode"`
	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name *string `pulumi:"name"`
	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkId *string `pulumi:"networkId"`
	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway *bool `pulumi:"noGateway"`
	// The prefix length to use when creating a subnet
	// from a subnet pool. The default subnet pool prefix length that was defined
	// when creating the subnet pool will be used if not provided. Changing this
	// creates a new subnet.
	PrefixLength *int `pulumi:"prefixLength"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnet.
	Region *string `pulumi:"region"`
	// An array of service types used by the subnet.
	// Changing this updates the service types for the existing subnet.
	ServiceTypes []string `pulumi:"serviceTypes"`
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId *string `pulumi:"subnetpoolId"`
	// A set of string tags for the subnet.
	Tags []string `pulumi:"tags"`
	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type SubnetState struct {
	// The collection of ags assigned on the subnet, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// A block declaring the start and end range of
	// the IP addresses available for use with DHCP in this subnet. Multiple
	// `allocationPool` blocks can be declared, providing the subnet with more
	// than one range of IP addresses to use with DHCP. However, each IP range
	// must be from the same CIDR that the subnet is part of.
	// The `allocationPool` block is documented below.
	AllocationPools SubnetAllocationPoolArrayInput
	// A block declaring the start and end range of the IP addresses available for
	// use with DHCP in this subnet.
	// The `allocationPools` block is documented below.
	//
	// Deprecated: use allocation_pool instead
	AllocationPoolsCollection SubnetAllocationPoolsCollectionArrayInput
	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr pulumi.StringPtrInput
	// Human-readable description of the subnet. Changing this
	// updates the name of the existing subnet.
	Description pulumi.StringPtrInput
	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DnsNameservers pulumi.StringArrayInput
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDhcp pulumi.BoolPtrInput
	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting `noGateway` will cause a default
	// gateway of `.1` to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIp pulumi.StringPtrInput
	// (**Deprecated** - use `networking.SubnetRoute`
	// instead) An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The hostRoute
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	//
	// Deprecated: Use openstack_networking_subnet_route_v2 instead
	HostRoutes SubnetHostRouteArrayInput
	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IpVersion pulumi.IntPtrInput
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode pulumi.StringPtrInput
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode pulumi.StringPtrInput
	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name pulumi.StringPtrInput
	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkId pulumi.StringPtrInput
	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway pulumi.BoolPtrInput
	// The prefix length to use when creating a subnet
	// from a subnet pool. The default subnet pool prefix length that was defined
	// when creating the subnet pool will be used if not provided. Changing this
	// creates a new subnet.
	PrefixLength pulumi.IntPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnet.
	Region pulumi.StringPtrInput
	// An array of service types used by the subnet.
	// Changing this updates the service types for the existing subnet.
	ServiceTypes pulumi.StringArrayInput
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId pulumi.StringPtrInput
	// A set of string tags for the subnet.
	Tags pulumi.StringArrayInput
	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (SubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetState)(nil)).Elem()
}

type subnetArgs struct {
	// A block declaring the start and end range of
	// the IP addresses available for use with DHCP in this subnet. Multiple
	// `allocationPool` blocks can be declared, providing the subnet with more
	// than one range of IP addresses to use with DHCP. However, each IP range
	// must be from the same CIDR that the subnet is part of.
	// The `allocationPool` block is documented below.
	AllocationPools []SubnetAllocationPool `pulumi:"allocationPools"`
	// A block declaring the start and end range of the IP addresses available for
	// use with DHCP in this subnet.
	// The `allocationPools` block is documented below.
	//
	// Deprecated: use allocation_pool instead
	AllocationPoolsCollection []SubnetAllocationPoolsCollection `pulumi:"allocationPoolsCollection"`
	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr *string `pulumi:"cidr"`
	// Human-readable description of the subnet. Changing this
	// updates the name of the existing subnet.
	Description *string `pulumi:"description"`
	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DnsNameservers []string `pulumi:"dnsNameservers"`
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDhcp *bool `pulumi:"enableDhcp"`
	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting `noGateway` will cause a default
	// gateway of `.1` to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIp *string `pulumi:"gatewayIp"`
	// (**Deprecated** - use `networking.SubnetRoute`
	// instead) An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The hostRoute
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	//
	// Deprecated: Use openstack_networking_subnet_route_v2 instead
	HostRoutes []SubnetHostRoute `pulumi:"hostRoutes"`
	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IpVersion *int `pulumi:"ipVersion"`
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode *string `pulumi:"ipv6AddressMode"`
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode *string `pulumi:"ipv6RaMode"`
	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name *string `pulumi:"name"`
	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkId string `pulumi:"networkId"`
	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway *bool `pulumi:"noGateway"`
	// The prefix length to use when creating a subnet
	// from a subnet pool. The default subnet pool prefix length that was defined
	// when creating the subnet pool will be used if not provided. Changing this
	// creates a new subnet.
	PrefixLength *int `pulumi:"prefixLength"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnet.
	Region *string `pulumi:"region"`
	// An array of service types used by the subnet.
	// Changing this updates the service types for the existing subnet.
	ServiceTypes []string `pulumi:"serviceTypes"`
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId *string `pulumi:"subnetpoolId"`
	// A set of string tags for the subnet.
	Tags []string `pulumi:"tags"`
	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// A block declaring the start and end range of
	// the IP addresses available for use with DHCP in this subnet. Multiple
	// `allocationPool` blocks can be declared, providing the subnet with more
	// than one range of IP addresses to use with DHCP. However, each IP range
	// must be from the same CIDR that the subnet is part of.
	// The `allocationPool` block is documented below.
	AllocationPools SubnetAllocationPoolArrayInput
	// A block declaring the start and end range of the IP addresses available for
	// use with DHCP in this subnet.
	// The `allocationPools` block is documented below.
	//
	// Deprecated: use allocation_pool instead
	AllocationPoolsCollection SubnetAllocationPoolsCollectionArrayInput
	// CIDR representing IP range for this subnet, based on IP
	// version. You can omit this option if you are creating a subnet from a
	// subnet pool.
	Cidr pulumi.StringPtrInput
	// Human-readable description of the subnet. Changing this
	// updates the name of the existing subnet.
	Description pulumi.StringPtrInput
	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	DnsNameservers pulumi.StringArrayInput
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	EnableDhcp pulumi.BoolPtrInput
	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting `noGateway` will cause a default
	// gateway of `.1` to be used. Changing this updates the gateway IP of the
	// existing subnet.
	GatewayIp pulumi.StringPtrInput
	// (**Deprecated** - use `networking.SubnetRoute`
	// instead) An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The hostRoute
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	//
	// Deprecated: Use openstack_networking_subnet_route_v2 instead
	HostRoutes SubnetHostRouteArrayInput
	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	IpVersion pulumi.IntPtrInput
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode pulumi.StringPtrInput
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode pulumi.StringPtrInput
	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	Name pulumi.StringPtrInput
	// The UUID of the parent network. Changing this
	// creates a new subnet.
	NetworkId pulumi.StringInput
	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	NoGateway pulumi.BoolPtrInput
	// The prefix length to use when creating a subnet
	// from a subnet pool. The default subnet pool prefix length that was defined
	// when creating the subnet pool will be used if not provided. Changing this
	// creates a new subnet.
	PrefixLength pulumi.IntPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnet.
	Region pulumi.StringPtrInput
	// An array of service types used by the subnet.
	// Changing this updates the service types for the existing subnet.
	ServiceTypes pulumi.StringArrayInput
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId pulumi.StringPtrInput
	// A set of string tags for the subnet.
	Tags pulumi.StringArrayInput
	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetArgs)(nil)).Elem()
}

type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(ctx context.Context) SubnetOutput
}

func (*Subnet) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (i *Subnet) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i *Subnet) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

func (i *Subnet) ToOutput(ctx context.Context) pulumix.Output[*Subnet] {
	return pulumix.Output[*Subnet]{
		OutputState: i.ToSubnetOutputWithContext(ctx).OutputState,
	}
}

// SubnetArrayInput is an input type that accepts SubnetArray and SubnetArrayOutput values.
// You can construct a concrete instance of `SubnetArrayInput` via:
//
//	SubnetArray{ SubnetArgs{...} }
type SubnetArrayInput interface {
	pulumi.Input

	ToSubnetArrayOutput() SubnetArrayOutput
	ToSubnetArrayOutputWithContext(context.Context) SubnetArrayOutput
}

type SubnetArray []SubnetInput

func (SubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subnet)(nil)).Elem()
}

func (i SubnetArray) ToSubnetArrayOutput() SubnetArrayOutput {
	return i.ToSubnetArrayOutputWithContext(context.Background())
}

func (i SubnetArray) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetArrayOutput)
}

func (i SubnetArray) ToOutput(ctx context.Context) pulumix.Output[[]*Subnet] {
	return pulumix.Output[[]*Subnet]{
		OutputState: i.ToSubnetArrayOutputWithContext(ctx).OutputState,
	}
}

// SubnetMapInput is an input type that accepts SubnetMap and SubnetMapOutput values.
// You can construct a concrete instance of `SubnetMapInput` via:
//
//	SubnetMap{ "key": SubnetArgs{...} }
type SubnetMapInput interface {
	pulumi.Input

	ToSubnetMapOutput() SubnetMapOutput
	ToSubnetMapOutputWithContext(context.Context) SubnetMapOutput
}

type SubnetMap map[string]SubnetInput

func (SubnetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subnet)(nil)).Elem()
}

func (i SubnetMap) ToSubnetMapOutput() SubnetMapOutput {
	return i.ToSubnetMapOutputWithContext(context.Background())
}

func (i SubnetMap) ToSubnetMapOutputWithContext(ctx context.Context) SubnetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetMapOutput)
}

func (i SubnetMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Subnet] {
	return pulumix.Output[map[string]*Subnet]{
		OutputState: i.ToSubnetMapOutputWithContext(ctx).OutputState,
	}
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) ToOutput(ctx context.Context) pulumix.Output[*Subnet] {
	return pulumix.Output[*Subnet]{
		OutputState: o.OutputState,
	}
}

// The collection of ags assigned on the subnet, which have been
// explicitly and implicitly added.
func (o SubnetOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringArrayOutput { return v.AllTags }).(pulumi.StringArrayOutput)
}

// A block declaring the start and end range of
// the IP addresses available for use with DHCP in this subnet. Multiple
// `allocationPool` blocks can be declared, providing the subnet with more
// than one range of IP addresses to use with DHCP. However, each IP range
// must be from the same CIDR that the subnet is part of.
// The `allocationPool` block is documented below.
func (o SubnetOutput) AllocationPools() SubnetAllocationPoolArrayOutput {
	return o.ApplyT(func(v *Subnet) SubnetAllocationPoolArrayOutput { return v.AllocationPools }).(SubnetAllocationPoolArrayOutput)
}

// A block declaring the start and end range of the IP addresses available for
// use with DHCP in this subnet.
// The `allocationPools` block is documented below.
//
// Deprecated: use allocation_pool instead
func (o SubnetOutput) AllocationPoolsCollection() SubnetAllocationPoolsCollectionArrayOutput {
	return o.ApplyT(func(v *Subnet) SubnetAllocationPoolsCollectionArrayOutput { return v.AllocationPoolsCollection }).(SubnetAllocationPoolsCollectionArrayOutput)
}

// CIDR representing IP range for this subnet, based on IP
// version. You can omit this option if you are creating a subnet from a
// subnet pool.
func (o SubnetOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.Cidr }).(pulumi.StringOutput)
}

// Human-readable description of the subnet. Changing this
// updates the name of the existing subnet.
func (o SubnetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An array of DNS name server names used by hosts
// in this subnet. Changing this updates the DNS name servers for the existing
// subnet.
func (o SubnetOutput) DnsNameservers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringArrayOutput { return v.DnsNameservers }).(pulumi.StringArrayOutput)
}

// The administrative state of the network.
// Acceptable values are "true" and "false". Changing this value enables or
// disables the DHCP capabilities of the existing subnet. Defaults to true.
func (o SubnetOutput) EnableDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.BoolPtrOutput { return v.EnableDhcp }).(pulumi.BoolPtrOutput)
}

// Default gateway used by devices in this subnet.
// Leaving this blank and not setting `noGateway` will cause a default
// gateway of `.1` to be used. Changing this updates the gateway IP of the
// existing subnet.
func (o SubnetOutput) GatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.GatewayIp }).(pulumi.StringOutput)
}

// (**Deprecated** - use `networking.SubnetRoute`
// instead) An array of routes that should be used by devices
// with IPs from this subnet (not including local subnet route). The hostRoute
// object structure is documented below. Changing this updates the host routes
// for the existing subnet.
//
// Deprecated: Use openstack_networking_subnet_route_v2 instead
func (o SubnetOutput) HostRoutes() SubnetHostRouteArrayOutput {
	return o.ApplyT(func(v *Subnet) SubnetHostRouteArrayOutput { return v.HostRoutes }).(SubnetHostRouteArrayOutput)
}

// IP version, either 4 (default) or 6. Changing this creates a
// new subnet.
func (o SubnetOutput) IpVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.IntPtrOutput { return v.IpVersion }).(pulumi.IntPtrOutput)
}

// The IPv6 address mode. Valid values are
// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
func (o SubnetOutput) Ipv6AddressMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.Ipv6AddressMode }).(pulumi.StringOutput)
}

// The IPv6 Router Advertisement mode. Valid values
// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
func (o SubnetOutput) Ipv6RaMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.Ipv6RaMode }).(pulumi.StringOutput)
}

// The name of the subnet. Changing this updates the name of
// the existing subnet.
func (o SubnetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The UUID of the parent network. Changing this
// creates a new subnet.
func (o SubnetOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// Do not set a gateway IP on this subnet. Changing
// this removes or adds a default gateway IP of the existing subnet.
func (o SubnetOutput) NoGateway() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.BoolPtrOutput { return v.NoGateway }).(pulumi.BoolPtrOutput)
}

// The prefix length to use when creating a subnet
// from a subnet pool. The default subnet pool prefix length that was defined
// when creating the subnet pool will be used if not provided. Changing this
// creates a new subnet.
func (o SubnetOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.IntPtrOutput { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a Neutron subnet. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// subnet.
func (o SubnetOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An array of service types used by the subnet.
// Changing this updates the service types for the existing subnet.
func (o SubnetOutput) ServiceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringArrayOutput { return v.ServiceTypes }).(pulumi.StringArrayOutput)
}

// The ID of the subnetpool associated with the subnet.
func (o SubnetOutput) SubnetpoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.SubnetpoolId }).(pulumi.StringPtrOutput)
}

// A set of string tags for the subnet.
func (o SubnetOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The owner of the subnet. Required if admin wants to
// create a subnet for another tenant. Changing this creates a new subnet.
func (o SubnetOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Map of additional options.
func (o SubnetOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *Subnet) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type SubnetArrayOutput struct{ *pulumi.OutputState }

func (SubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Subnet)(nil)).Elem()
}

func (o SubnetArrayOutput) ToSubnetArrayOutput() SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Subnet] {
	return pulumix.Output[[]*Subnet]{
		OutputState: o.OutputState,
	}
}

func (o SubnetArrayOutput) Index(i pulumi.IntInput) SubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Subnet {
		return vs[0].([]*Subnet)[vs[1].(int)]
	}).(SubnetOutput)
}

type SubnetMapOutput struct{ *pulumi.OutputState }

func (SubnetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Subnet)(nil)).Elem()
}

func (o SubnetMapOutput) ToSubnetMapOutput() SubnetMapOutput {
	return o
}

func (o SubnetMapOutput) ToSubnetMapOutputWithContext(ctx context.Context) SubnetMapOutput {
	return o
}

func (o SubnetMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Subnet] {
	return pulumix.Output[map[string]*Subnet]{
		OutputState: o.OutputState,
	}
}

func (o SubnetMapOutput) MapIndex(k pulumi.StringInput) SubnetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Subnet {
		return vs[0].(map[string]*Subnet)[vs[1].(string)]
	}).(SubnetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetInput)(nil)).Elem(), &Subnet{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetArrayInput)(nil)).Elem(), SubnetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetMapInput)(nil)).Elem(), SubnetMap{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetArrayOutput{})
	pulumi.RegisterOutputType(SubnetMapOutput{})
}
