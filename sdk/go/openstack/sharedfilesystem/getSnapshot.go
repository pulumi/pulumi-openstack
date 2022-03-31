// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available Shared File System snapshot.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/sharedfilesystem"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sharedfilesystem.GetSnapshot(ctx, &sharedfilesystem.GetSnapshotArgs{
// 			Name: pulumi.StringRef("snapshot_1"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSnapshot(ctx *pulumi.Context, args *GetSnapshotArgs, opts ...pulumi.InvokeOption) (*GetSnapshotResult, error) {
	var rv GetSnapshotResult
	err := ctx.Invoke("openstack:sharedfilesystem/getSnapshot:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshot.
type GetSnapshotArgs struct {
	// The human-readable description of the snapshot.
	Description *string `pulumi:"description"`
	// The name of the snapshot.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Shared File System client.
	Region *string `pulumi:"region"`
	// The UUID of the source share that was used to create the snapshot.
	ShareId *string `pulumi:"shareId"`
	// A snapshot status filter. A valid value is `available`, `error`,
	// `creating`, `deleting`, `manageStarting`, `manageError`, `unmanageStarting`,
	// `unmanageError` or `errorDeleting`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getSnapshot.
type GetSnapshotResult struct {
	// See Argument Reference above.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	ProjectId string `pulumi:"projectId"`
	Region    string `pulumi:"region"`
	// The UUID of the source share that was used to create the snapshot.
	ShareId string `pulumi:"shareId"`
	// The file system protocol of a share snapshot.
	ShareProto string `pulumi:"shareProto"`
	// The share snapshot size, in GBs.
	ShareSize int `pulumi:"shareSize"`
	// The snapshot size, in GBs.
	Size int `pulumi:"size"`
	// See Argument Reference above.
	Status string `pulumi:"status"`
}

func GetSnapshotOutput(ctx *pulumi.Context, args GetSnapshotOutputArgs, opts ...pulumi.InvokeOption) GetSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSnapshotResult, error) {
			args := v.(GetSnapshotArgs)
			r, err := GetSnapshot(ctx, &args, opts...)
			return *r, err
		}).(GetSnapshotResultOutput)
}

// A collection of arguments for invoking getSnapshot.
type GetSnapshotOutputArgs struct {
	// The human-readable description of the snapshot.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The name of the snapshot.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V2 Shared File System client.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The UUID of the source share that was used to create the snapshot.
	ShareId pulumi.StringPtrInput `pulumi:"shareId"`
	// A snapshot status filter. A valid value is `available`, `error`,
	// `creating`, `deleting`, `manageStarting`, `manageError`, `unmanageStarting`,
	// `unmanageError` or `errorDeleting`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getSnapshot.
type GetSnapshotResultOutput struct{ *pulumi.OutputState }

func (GetSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotResult)(nil)).Elem()
}

func (o GetSnapshotResultOutput) ToGetSnapshotResultOutput() GetSnapshotResultOutput {
	return o
}

func (o GetSnapshotResultOutput) ToGetSnapshotResultOutputWithContext(ctx context.Context) GetSnapshotResultOutput {
	return o
}

// See Argument Reference above.
func (o GetSnapshotResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetSnapshotResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetSnapshotResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotResult) string { return v.Region }).(pulumi.StringOutput)
}

// The UUID of the source share that was used to create the snapshot.
func (o GetSnapshotResultOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotResult) string { return v.ShareId }).(pulumi.StringOutput)
}

// The file system protocol of a share snapshot.
func (o GetSnapshotResultOutput) ShareProto() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotResult) string { return v.ShareProto }).(pulumi.StringOutput)
}

// The share snapshot size, in GBs.
func (o GetSnapshotResultOutput) ShareSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetSnapshotResult) int { return v.ShareSize }).(pulumi.IntOutput)
}

// The snapshot size, in GBs.
func (o GetSnapshotResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetSnapshotResult) int { return v.Size }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o GetSnapshotResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSnapshotResultOutput{})
}
