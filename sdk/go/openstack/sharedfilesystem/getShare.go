// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to get the ID of an available Shared File System share.
func LookupShare(ctx *pulumi.Context, args *LookupShareArgs, opts ...pulumi.InvokeOption) (*LookupShareResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupShareResult
	err := ctx.Invoke("openstack:sharedfilesystem/getShare:getShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getShare.
type LookupShareArgs struct {
	// The human-readable description for the share.
	Description *string `pulumi:"description"`
	// The export location path of the share. Available
	// since Manila API version 2.35.
	ExportLocationPath *string `pulumi:"exportLocationPath"`
	// The level of visibility for the share.
	// length.
	IsPublic *bool `pulumi:"isPublic"`
	// One or more metadata key and value pairs as a dictionary of
	// strings.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The name of the share.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Shared File System client.
	Region *string `pulumi:"region"`
	// The UUID of the share's share network.
	ShareNetworkId *string `pulumi:"shareNetworkId"`
	// The UUID of the share's base snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// A share status filter. A valid value is `creating`,
	// `error`, `available`, `deleting`, `errorDeleting`, `manageStarting`,
	// `manageError`, `unmanageStarting`, `unmanageError`, `unmanaged`,
	// `extending`, `extendingError`, `shrinking`, `shrinkingError`, or
	// `shrinkingPossibleDataLossError`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getShare.
type LookupShareResult struct {
	// The share availability zone.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// See Argument Reference above.
	Description string `pulumi:"description"`
	// See Argument Reference above.
	ExportLocationPath *string `pulumi:"exportLocationPath"`
	// A list of export locations. For example, when a share
	// server has more than one network interface, it can have multiple export
	// locations.
	ExportLocations []GetShareExportLocation `pulumi:"exportLocations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IsPublic bool `pulumi:"isPublic"`
	// See Argument Reference above.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	ProjectId string `pulumi:"projectId"`
	// The region in which to obtain the V2 Shared File System client.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	ShareNetworkId string `pulumi:"shareNetworkId"`
	// The share protocol.
	ShareProto string `pulumi:"shareProto"`
	// The share size, in GBs.
	Size int `pulumi:"size"`
	// See Argument Reference above.
	SnapshotId string `pulumi:"snapshotId"`
	// See Argument Reference above.
	Status string `pulumi:"status"`
}

func LookupShareOutput(ctx *pulumi.Context, args LookupShareOutputArgs, opts ...pulumi.InvokeOption) LookupShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupShareResult, error) {
			args := v.(LookupShareArgs)
			r, err := LookupShare(ctx, &args, opts...)
			var s LookupShareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupShareResultOutput)
}

// A collection of arguments for invoking getShare.
type LookupShareOutputArgs struct {
	// The human-readable description for the share.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The export location path of the share. Available
	// since Manila API version 2.35.
	ExportLocationPath pulumi.StringPtrInput `pulumi:"exportLocationPath"`
	// The level of visibility for the share.
	// length.
	IsPublic pulumi.BoolPtrInput `pulumi:"isPublic"`
	// One or more metadata key and value pairs as a dictionary of
	// strings.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The name of the share.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V2 Shared File System client.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The UUID of the share's share network.
	ShareNetworkId pulumi.StringPtrInput `pulumi:"shareNetworkId"`
	// The UUID of the share's base snapshot.
	SnapshotId pulumi.StringPtrInput `pulumi:"snapshotId"`
	// A share status filter. A valid value is `creating`,
	// `error`, `available`, `deleting`, `errorDeleting`, `manageStarting`,
	// `manageError`, `unmanageStarting`, `unmanageError`, `unmanaged`,
	// `extending`, `extendingError`, `shrinking`, `shrinkingError`, or
	// `shrinkingPossibleDataLossError`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (LookupShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareArgs)(nil)).Elem()
}

// A collection of values returned by getShare.
type LookupShareResultOutput struct{ *pulumi.OutputState }

func (LookupShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareResult)(nil)).Elem()
}

func (o LookupShareResultOutput) ToLookupShareResultOutput() LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) ToLookupShareResultOutputWithContext(ctx context.Context) LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupShareResult] {
	return pulumix.Output[LookupShareResult]{
		OutputState: o.OutputState,
	}
}

// The share availability zone.
func (o LookupShareResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Description }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) ExportLocationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareResult) *string { return v.ExportLocationPath }).(pulumi.StringPtrOutput)
}

// A list of export locations. For example, when a share
// server has more than one network interface, it can have multiple export
// locations.
func (o LookupShareResultOutput) ExportLocations() GetShareExportLocationArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []GetShareExportLocation { return v.ExportLocations }).(GetShareExportLocationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) IsPublic() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupShareResult) bool { return v.IsPublic }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupShareResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Name }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Shared File System client.
func (o LookupShareResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) ShareNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.ShareNetworkId }).(pulumi.StringOutput)
}

// The share protocol.
func (o LookupShareResultOutput) ShareProto() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.ShareProto }).(pulumi.StringOutput)
}

// The share size, in GBs.
func (o LookupShareResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupShareResult) int { return v.Size }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.SnapshotId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupShareResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupShareResultOutput{})
}
