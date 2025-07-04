// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an existing snapshot.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/blockstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blockstorage.GetSnapshotV3(ctx, &blockstorage.GetSnapshotV3Args{
//				Name:       pulumi.StringRef("snapshot_1"),
//				MostRecent: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSnapshotV3(ctx *pulumi.Context, args *GetSnapshotV3Args, opts ...pulumi.InvokeOption) (*GetSnapshotV3Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSnapshotV3Result
	err := ctx.Invoke("openstack:blockstorage/getSnapshotV3:getSnapshotV3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshotV3.
type GetSnapshotV3Args struct {
	// Pick the most recently created snapshot if there
	// are multiple results.
	MostRecent *bool `pulumi:"mostRecent"`
	// The name of the snapshot.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Block Storage
	// client. If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The status of the snapshot.
	Status *string `pulumi:"status"`
	// The ID of the snapshot's volume.
	VolumeId *string `pulumi:"volumeId"`
}

// A collection of values returned by getSnapshotV3.
type GetSnapshotV3Result struct {
	// The snapshot's description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The snapshot's metadata.
	Metadata   map[string]string `pulumi:"metadata"`
	MostRecent *bool             `pulumi:"mostRecent"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The size of the snapshot.
	Size int `pulumi:"size"`
	// See Argument Reference above.
	Status string `pulumi:"status"`
	// See Argument Reference above.
	VolumeId string `pulumi:"volumeId"`
}

func GetSnapshotV3Output(ctx *pulumi.Context, args GetSnapshotV3OutputArgs, opts ...pulumi.InvokeOption) GetSnapshotV3ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetSnapshotV3ResultOutput, error) {
			args := v.(GetSnapshotV3Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:blockstorage/getSnapshotV3:getSnapshotV3", args, GetSnapshotV3ResultOutput{}, options).(GetSnapshotV3ResultOutput), nil
		}).(GetSnapshotV3ResultOutput)
}

// A collection of arguments for invoking getSnapshotV3.
type GetSnapshotV3OutputArgs struct {
	// Pick the most recently created snapshot if there
	// are multiple results.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// The name of the snapshot.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V3 Block Storage
	// client. If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The status of the snapshot.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The ID of the snapshot's volume.
	VolumeId pulumi.StringPtrInput `pulumi:"volumeId"`
}

func (GetSnapshotV3OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotV3Args)(nil)).Elem()
}

// A collection of values returned by getSnapshotV3.
type GetSnapshotV3ResultOutput struct{ *pulumi.OutputState }

func (GetSnapshotV3ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotV3Result)(nil)).Elem()
}

func (o GetSnapshotV3ResultOutput) ToGetSnapshotV3ResultOutput() GetSnapshotV3ResultOutput {
	return o
}

func (o GetSnapshotV3ResultOutput) ToGetSnapshotV3ResultOutputWithContext(ctx context.Context) GetSnapshotV3ResultOutput {
	return o
}

// The snapshot's description.
func (o GetSnapshotV3ResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSnapshotV3ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) string { return v.Id }).(pulumi.StringOutput)
}

// The snapshot's metadata.
func (o GetSnapshotV3ResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o GetSnapshotV3ResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// See Argument Reference above.
func (o GetSnapshotV3ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) string { return v.Name }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetSnapshotV3ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) string { return v.Region }).(pulumi.StringOutput)
}

// The size of the snapshot.
func (o GetSnapshotV3ResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) int { return v.Size }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o GetSnapshotV3ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) string { return v.Status }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetSnapshotV3ResultOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV3Result) string { return v.VolumeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSnapshotV3ResultOutput{})
}
