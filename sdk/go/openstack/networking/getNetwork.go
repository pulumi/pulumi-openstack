// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack network.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.LookupNetwork(ctx, &networking.LookupNetworkArgs{
//				Name: pulumi.StringRef("tf_test_network"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkResult
	err := ctx.Invoke("openstack:networking/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	// Human-readable description of the network.
	Description *string `pulumi:"description"`
	// The external routing facility of the network.
	External *bool `pulumi:"external"`
	// The CIDR of a subnet within the network.
	MatchingSubnetCidr *string `pulumi:"matchingSubnetCidr"`
	// The network MTU to filter. Available, when Neutron `net-mtu`
	// extension is enabled.
	Mtu *int `pulumi:"mtu"`
	// The name of the network.
	Name *string `pulumi:"name"`
	// The ID of the network.
	NetworkId *string `pulumi:"networkId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve networks ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The status of the network.
	Status *string `pulumi:"status"`
	// The list of network tags to filter.
	Tags []string `pulumi:"tags"`
	// The owner of the network.
	TenantId *string `pulumi:"tenantId"`
	// The VLAN transparent attribute for the
	// network.
	TransparentVlan *bool `pulumi:"transparentVlan"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// The administrative state of the network.
	AdminStateUp string `pulumi:"adminStateUp"`
	// The set of string tags applied on the network.
	AllTags []string `pulumi:"allTags"`
	// The availability zone candidates for the network.
	AvailabilityZoneHints []string `pulumi:"availabilityZoneHints"`
	// See Argument Reference above.
	Description *string `pulumi:"description"`
	// The network DNS domain. Available, when Neutron DNS extension
	// is enabled
	DnsDomain string `pulumi:"dnsDomain"`
	// See Argument Reference above.
	External *bool `pulumi:"external"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	MatchingSubnetCidr *string `pulumi:"matchingSubnetCidr"`
	// See Argument Reference above.
	Mtu *int `pulumi:"mtu"`
	// See Argument Reference above.
	Name      *string `pulumi:"name"`
	NetworkId *string `pulumi:"networkId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// An array of one or more provider segment objects.
	Segments []GetNetworkSegment `pulumi:"segments"`
	// Specifies whether the network resource can be accessed by any
	// tenant or not.
	Shared string  `pulumi:"shared"`
	Status *string `pulumi:"status"`
	// A list of subnet IDs belonging to the network.
	Subnets  []string `pulumi:"subnets"`
	Tags     []string `pulumi:"tags"`
	TenantId *string  `pulumi:"tenantId"`
	// See Argument Reference above.
	TransparentVlan *bool `pulumi:"transparentVlan"`
}

func LookupNetworkOutput(ctx *pulumi.Context, args LookupNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkResultOutput, error) {
			args := v.(LookupNetworkArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupNetworkResult
			secret, err := ctx.InvokePackageRaw("openstack:networking/getNetwork:getNetwork", args, &rv, "", opts...)
			if err != nil {
				return LookupNetworkResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupNetworkResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupNetworkResultOutput), nil
			}
			return output, nil
		}).(LookupNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkOutputArgs struct {
	// Human-readable description of the network.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The external routing facility of the network.
	External pulumi.BoolPtrInput `pulumi:"external"`
	// The CIDR of a subnet within the network.
	MatchingSubnetCidr pulumi.StringPtrInput `pulumi:"matchingSubnetCidr"`
	// The network MTU to filter. Available, when Neutron `net-mtu`
	// extension is enabled.
	Mtu pulumi.IntPtrInput `pulumi:"mtu"`
	// The name of the network.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the network.
	NetworkId pulumi.StringPtrInput `pulumi:"networkId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve networks ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The status of the network.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The list of network tags to filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The owner of the network.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
	// The VLAN transparent attribute for the
	// network.
	TransparentVlan pulumi.BoolPtrInput `pulumi:"transparentVlan"`
}

func (LookupNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type LookupNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkResult)(nil)).Elem()
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutput() LookupNetworkResultOutput {
	return o
}

func (o LookupNetworkResultOutput) ToLookupNetworkResultOutputWithContext(ctx context.Context) LookupNetworkResultOutput {
	return o
}

// The administrative state of the network.
func (o LookupNetworkResultOutput) AdminStateUp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.AdminStateUp }).(pulumi.StringOutput)
}

// The set of string tags applied on the network.
func (o LookupNetworkResultOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []string { return v.AllTags }).(pulumi.StringArrayOutput)
}

// The availability zone candidates for the network.
func (o LookupNetworkResultOutput) AvailabilityZoneHints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []string { return v.AvailabilityZoneHints }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o LookupNetworkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The network DNS domain. Available, when Neutron DNS extension
// is enabled
func (o LookupNetworkResultOutput) DnsDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.DnsDomain }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupNetworkResultOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *bool { return v.External }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) MatchingSubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.MatchingSubnetCidr }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupNetworkResultOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *int { return v.Mtu }).(pulumi.IntPtrOutput)
}

// See Argument Reference above.
func (o LookupNetworkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkResultOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.NetworkId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupNetworkResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Region }).(pulumi.StringOutput)
}

// An array of one or more provider segment objects.
func (o LookupNetworkResultOutput) Segments() GetNetworkSegmentArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []GetNetworkSegment { return v.Segments }).(GetNetworkSegmentArrayOutput)
}

// Specifies whether the network resource can be accessed by any
// tenant or not.
func (o LookupNetworkResultOutput) Shared() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkResult) string { return v.Shared }).(pulumi.StringOutput)
}

func (o LookupNetworkResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A list of subnet IDs belonging to the network.
func (o LookupNetworkResultOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []string { return v.Subnets }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupNetworkResultOutput) TransparentVlan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkResult) *bool { return v.TransparentVlan }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkResultOutput{})
}
