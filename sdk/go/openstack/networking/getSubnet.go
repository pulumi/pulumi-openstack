// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to get the ID of an available OpenStack subnet.
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
//			_, err := networking.LookupSubnet(ctx, &networking.LookupSubnetArgs{
//				Name: pulumi.StringRef("subnet_1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSubnetResult
	err := ctx.Invoke("openstack:networking/getSubnet:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetArgs struct {
	// The CIDR of the subnet.
	Cidr *string `pulumi:"cidr"`
	// Human-readable description of the subnet.
	Description *string `pulumi:"description"`
	// Deprecated: use dhcp_enabled instead
	DhcpDisabled *bool `pulumi:"dhcpDisabled"`
	// If the subnet has DHCP enabled.
	DhcpEnabled *bool `pulumi:"dhcpEnabled"`
	// The IP of the subnet's gateway.
	GatewayIp *string `pulumi:"gatewayIp"`
	// The IP version of the subnet (either 4 or 6).
	IpVersion *int `pulumi:"ipVersion"`
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode *string `pulumi:"ipv6AddressMode"`
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode *string `pulumi:"ipv6RaMode"`
	// The name of the subnet.
	Name *string `pulumi:"name"`
	// The ID of the network the subnet belongs to.
	NetworkId *string `pulumi:"networkId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve subnet ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The ID of the subnet.
	SubnetId *string `pulumi:"subnetId"`
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId *string `pulumi:"subnetpoolId"`
	// The list of subnet tags to filter.
	Tags []string `pulumi:"tags"`
	// The owner of the subnet.
	TenantId *string `pulumi:"tenantId"`
}

// A collection of values returned by getSubnet.
type LookupSubnetResult struct {
	// A set of string tags applied on the subnet.
	AllTags []string `pulumi:"allTags"`
	// Allocation pools of the subnet.
	AllocationPools []GetSubnetAllocationPool `pulumi:"allocationPools"`
	Cidr            string                    `pulumi:"cidr"`
	Description     string                    `pulumi:"description"`
	// Deprecated: use dhcp_enabled instead
	DhcpDisabled *bool `pulumi:"dhcpDisabled"`
	DhcpEnabled  *bool `pulumi:"dhcpEnabled"`
	// DNS Nameservers of the subnet.
	DnsNameservers []string `pulumi:"dnsNameservers"`
	// Whether the subnet has DHCP enabled or not.
	EnableDhcp bool   `pulumi:"enableDhcp"`
	GatewayIp  string `pulumi:"gatewayIp"`
	// Host Routes of the subnet.
	HostRoutes []GetSubnetHostRoute `pulumi:"hostRoutes"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	IpVersion       int    `pulumi:"ipVersion"`
	Ipv6AddressMode string `pulumi:"ipv6AddressMode"`
	Ipv6RaMode      string `pulumi:"ipv6RaMode"`
	Name            string `pulumi:"name"`
	NetworkId       string `pulumi:"networkId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// Service types of the subnet.
	ServiceTypes []string `pulumi:"serviceTypes"`
	SubnetId     string   `pulumi:"subnetId"`
	SubnetpoolId string   `pulumi:"subnetpoolId"`
	Tags         []string `pulumi:"tags"`
	TenantId     string   `pulumi:"tenantId"`
}

func LookupSubnetOutput(ctx *pulumi.Context, args LookupSubnetOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetResult, error) {
			args := v.(LookupSubnetArgs)
			r, err := LookupSubnet(ctx, &args, opts...)
			var s LookupSubnetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetResultOutput)
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetOutputArgs struct {
	// The CIDR of the subnet.
	Cidr pulumi.StringPtrInput `pulumi:"cidr"`
	// Human-readable description of the subnet.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Deprecated: use dhcp_enabled instead
	DhcpDisabled pulumi.BoolPtrInput `pulumi:"dhcpDisabled"`
	// If the subnet has DHCP enabled.
	DhcpEnabled pulumi.BoolPtrInput `pulumi:"dhcpEnabled"`
	// The IP of the subnet's gateway.
	GatewayIp pulumi.StringPtrInput `pulumi:"gatewayIp"`
	// The IP version of the subnet (either 4 or 6).
	IpVersion pulumi.IntPtrInput `pulumi:"ipVersion"`
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode pulumi.StringPtrInput `pulumi:"ipv6AddressMode"`
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode pulumi.StringPtrInput `pulumi:"ipv6RaMode"`
	// The name of the subnet.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the network the subnet belongs to.
	NetworkId pulumi.StringPtrInput `pulumi:"networkId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve subnet ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The ID of the subnet.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId pulumi.StringPtrInput `pulumi:"subnetpoolId"`
	// The list of subnet tags to filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The owner of the subnet.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupSubnetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetArgs)(nil)).Elem()
}

// A collection of values returned by getSubnet.
type LookupSubnetResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetResult)(nil)).Elem()
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutput() LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutputWithContext(ctx context.Context) LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSubnetResult] {
	return pulumix.Output[LookupSubnetResult]{
		OutputState: o.OutputState,
	}
}

// A set of string tags applied on the subnet.
func (o LookupSubnetResultOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []string { return v.AllTags }).(pulumi.StringArrayOutput)
}

// Allocation pools of the subnet.
func (o LookupSubnetResultOutput) AllocationPools() GetSubnetAllocationPoolArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []GetSubnetAllocationPool { return v.AllocationPools }).(GetSubnetAllocationPoolArrayOutput)
}

func (o LookupSubnetResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Cidr }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Description }).(pulumi.StringOutput)
}

// Deprecated: use dhcp_enabled instead
func (o LookupSubnetResultOutput) DhcpDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *bool { return v.DhcpDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSubnetResultOutput) DhcpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *bool { return v.DhcpEnabled }).(pulumi.BoolPtrOutput)
}

// DNS Nameservers of the subnet.
func (o LookupSubnetResultOutput) DnsNameservers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []string { return v.DnsNameservers }).(pulumi.StringArrayOutput)
}

// Whether the subnet has DHCP enabled or not.
func (o LookupSubnetResultOutput) EnableDhcp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetResult) bool { return v.EnableDhcp }).(pulumi.BoolOutput)
}

func (o LookupSubnetResultOutput) GatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.GatewayIp }).(pulumi.StringOutput)
}

// Host Routes of the subnet.
func (o LookupSubnetResultOutput) HostRoutes() GetSubnetHostRouteArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []GetSubnetHostRoute { return v.HostRoutes }).(GetSubnetHostRouteArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSubnetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) IpVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubnetResult) int { return v.IpVersion }).(pulumi.IntOutput)
}

func (o LookupSubnetResultOutput) Ipv6AddressMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Ipv6AddressMode }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Ipv6RaMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Ipv6RaMode }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupSubnetResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.Region }).(pulumi.StringOutput)
}

// Service types of the subnet.
func (o LookupSubnetResultOutput) ServiceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []string { return v.ServiceTypes }).(pulumi.StringArrayOutput)
}

func (o LookupSubnetResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) SubnetpoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.SubnetpoolId }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupSubnetResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetResultOutput{})
}
