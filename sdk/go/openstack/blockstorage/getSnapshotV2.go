// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an existing snapshot.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blockstorage.GetSnapshotV2(ctx, &blockstorage.GetSnapshotV2Args{
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
// <!--End PulumiCodeChooser -->
func GetSnapshotV2(ctx *pulumi.Context, args *GetSnapshotV2Args, opts ...pulumi.InvokeOption) (*GetSnapshotV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSnapshotV2Result
	err := ctx.Invoke("openstack:blockstorage/getSnapshotV2:getSnapshotV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshotV2.
type GetSnapshotV2Args struct {
	// Pick the most recently created snapshot if there
	// are multiple results.
	MostRecent *bool `pulumi:"mostRecent"`
	// The name of the snapshot.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Block Storage
	// client. If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The status of the snapshot.
	Status *string `pulumi:"status"`
	// The ID of the snapshot's volume.
	VolumeId *string `pulumi:"volumeId"`
}

// A collection of values returned by getSnapshotV2.
type GetSnapshotV2Result struct {
	// The snapshot's description.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The snapshot's metadata.
	Metadata   map[string]interface{} `pulumi:"metadata"`
	MostRecent *bool                  `pulumi:"mostRecent"`
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

func GetSnapshotV2Output(ctx *pulumi.Context, args GetSnapshotV2OutputArgs, opts ...pulumi.InvokeOption) GetSnapshotV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSnapshotV2Result, error) {
			args := v.(GetSnapshotV2Args)
			r, err := GetSnapshotV2(ctx, &args, opts...)
			var s GetSnapshotV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSnapshotV2ResultOutput)
}

// A collection of arguments for invoking getSnapshotV2.
type GetSnapshotV2OutputArgs struct {
	// Pick the most recently created snapshot if there
	// are multiple results.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// The name of the snapshot.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V2 Block Storage
	// client. If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The status of the snapshot.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The ID of the snapshot's volume.
	VolumeId pulumi.StringPtrInput `pulumi:"volumeId"`
}

func (GetSnapshotV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotV2Args)(nil)).Elem()
}

// A collection of values returned by getSnapshotV2.
type GetSnapshotV2ResultOutput struct{ *pulumi.OutputState }

func (GetSnapshotV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotV2Result)(nil)).Elem()
}

func (o GetSnapshotV2ResultOutput) ToGetSnapshotV2ResultOutput() GetSnapshotV2ResultOutput {
	return o
}

func (o GetSnapshotV2ResultOutput) ToGetSnapshotV2ResultOutputWithContext(ctx context.Context) GetSnapshotV2ResultOutput {
	return o
}

// The snapshot's description.
func (o GetSnapshotV2ResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSnapshotV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// The snapshot's metadata.
func (o GetSnapshotV2ResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o GetSnapshotV2ResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// See Argument Reference above.
func (o GetSnapshotV2ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) string { return v.Name }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetSnapshotV2ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) string { return v.Region }).(pulumi.StringOutput)
}

// The size of the snapshot.
func (o GetSnapshotV2ResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) int { return v.Size }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o GetSnapshotV2ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) string { return v.Status }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetSnapshotV2ResultOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotV2Result) string { return v.VolumeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSnapshotV2ResultOutput{})
}
