// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available Shared File System share network.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/sharedfilesystem"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sharedfilesystem.LookupShareNetwork(ctx, &sharedfilesystem.LookupShareNetworkArgs{
//				Name: pulumi.StringRef("sharenetwork_1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupShareNetwork(ctx *pulumi.Context, args *LookupShareNetworkArgs, opts ...pulumi.InvokeOption) (*LookupShareNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupShareNetworkResult
	err := ctx.Invoke("openstack:sharedfilesystem/getShareNetwork:getShareNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getShareNetwork.
type LookupShareNetworkArgs struct {
	// The human-readable description of the share network.
	Description *string `pulumi:"description"`
	// The IP version of the share network. Can either be 4 or 6.
	IpVersion *int `pulumi:"ipVersion"`
	// The name of the share network.
	Name *string `pulumi:"name"`
	// The share network type. Can either be VLAN, VXLAN,
	// GRE, or flat.
	NetworkType *string `pulumi:"networkType"`
	// The neutron network UUID of the share network.
	NeutronNetId *string `pulumi:"neutronNetId"`
	// The neutron subnet UUID of the share network.
	NeutronSubnetId *string `pulumi:"neutronSubnetId"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to read a share network. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The security service IDs associated with
	// the share network.
	SecurityServiceId *string `pulumi:"securityServiceId"`
	// The share network segmentation ID.
	SegmentationId *int `pulumi:"segmentationId"`
}

// A collection of values returned by getShareNetwork.
type LookupShareNetworkResult struct {
	// See Argument Reference above.
	Cidr string `pulumi:"cidr"`
	// See Argument Reference above.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IpVersion int `pulumi:"ipVersion"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	NetworkType string `pulumi:"networkType"`
	// See Argument Reference above.
	NeutronNetId string `pulumi:"neutronNetId"`
	// See Argument Reference above.
	NeutronSubnetId string `pulumi:"neutronSubnetId"`
	// The owner of the Share Network.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	SecurityServiceId *string `pulumi:"securityServiceId"`
	// The list of security service IDs associated with
	// the share network.
	SecurityServiceIds []string `pulumi:"securityServiceIds"`
	// See Argument Reference above.
	SegmentationId int `pulumi:"segmentationId"`
}

func LookupShareNetworkOutput(ctx *pulumi.Context, args LookupShareNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupShareNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupShareNetworkResult, error) {
			args := v.(LookupShareNetworkArgs)
			r, err := LookupShareNetwork(ctx, &args, opts...)
			var s LookupShareNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupShareNetworkResultOutput)
}

// A collection of arguments for invoking getShareNetwork.
type LookupShareNetworkOutputArgs struct {
	// The human-readable description of the share network.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The IP version of the share network. Can either be 4 or 6.
	IpVersion pulumi.IntPtrInput `pulumi:"ipVersion"`
	// The name of the share network.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The share network type. Can either be VLAN, VXLAN,
	// GRE, or flat.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	// The neutron network UUID of the share network.
	NeutronNetId pulumi.StringPtrInput `pulumi:"neutronNetId"`
	// The neutron subnet UUID of the share network.
	NeutronSubnetId pulumi.StringPtrInput `pulumi:"neutronSubnetId"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to read a share network. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The security service IDs associated with
	// the share network.
	SecurityServiceId pulumi.StringPtrInput `pulumi:"securityServiceId"`
	// The share network segmentation ID.
	SegmentationId pulumi.IntPtrInput `pulumi:"segmentationId"`
}

func (LookupShareNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getShareNetwork.
type LookupShareNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupShareNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareNetworkResult)(nil)).Elem()
}

func (o LookupShareNetworkResultOutput) ToLookupShareNetworkResultOutput() LookupShareNetworkResultOutput {
	return o
}

func (o LookupShareNetworkResultOutput) ToLookupShareNetworkResultOutputWithContext(ctx context.Context) LookupShareNetworkResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupShareNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) IpVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) int { return v.IpVersion }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.NetworkType }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) NeutronNetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.NeutronNetId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) NeutronSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.NeutronSubnetId }).(pulumi.StringOutput)
}

// The owner of the Share Network.
func (o LookupShareNetworkResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) SecurityServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) *string { return v.SecurityServiceId }).(pulumi.StringPtrOutput)
}

// The list of security service IDs associated with
// the share network.
func (o LookupShareNetworkResultOutput) SecurityServiceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) []string { return v.SecurityServiceIds }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o LookupShareNetworkResultOutput) SegmentationId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupShareNetworkResult) int { return v.SegmentationId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupShareNetworkResultOutput{})
}
