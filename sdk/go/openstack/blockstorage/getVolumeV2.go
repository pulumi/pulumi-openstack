// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about an existing volume.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := blockstorage.LookupVolumeV2(ctx, &blockstorage.LookupVolumeV2Args{
// 			Name: pulumi.StringRef("volume_1"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVolumeV2(ctx *pulumi.Context, args *LookupVolumeV2Args, opts ...pulumi.InvokeOption) (*LookupVolumeV2Result, error) {
	var rv LookupVolumeV2Result
	err := ctx.Invoke("openstack:blockstorage/getVolumeV2:getVolumeV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolumeV2.
type LookupVolumeV2Args struct {
	// Indicates if the volume is bootable.
	Bootable *string `pulumi:"bootable"`
	// Metadata key/value pairs associated with the volume.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The name of the volume.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Block Storage
	// client. If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The status of the volume.
	Status *string `pulumi:"status"`
	// The type of the volume.
	VolumeType *string `pulumi:"volumeType"`
}

// A collection of values returned by getVolumeV2.
type LookupVolumeV2Result struct {
	// Indicates if the volume is bootable.
	Bootable string `pulumi:"bootable"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The size of the volume in GBs.
	Size int `pulumi:"size"`
	// The ID of the volume from which the current volume was created.
	SourceVolumeId string `pulumi:"sourceVolumeId"`
	// See Argument Reference above.
	Status string `pulumi:"status"`
	// The type of the volume.
	VolumeType string `pulumi:"volumeType"`
}

func LookupVolumeV2Output(ctx *pulumi.Context, args LookupVolumeV2OutputArgs, opts ...pulumi.InvokeOption) LookupVolumeV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeV2Result, error) {
			args := v.(LookupVolumeV2Args)
			r, err := LookupVolumeV2(ctx, &args, opts...)
			return *r, err
		}).(LookupVolumeV2ResultOutput)
}

// A collection of arguments for invoking getVolumeV2.
type LookupVolumeV2OutputArgs struct {
	// Indicates if the volume is bootable.
	Bootable pulumi.StringPtrInput `pulumi:"bootable"`
	// Metadata key/value pairs associated with the volume.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The name of the volume.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V2 Block Storage
	// client. If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The status of the volume.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The type of the volume.
	VolumeType pulumi.StringPtrInput `pulumi:"volumeType"`
}

func (LookupVolumeV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeV2Args)(nil)).Elem()
}

// A collection of values returned by getVolumeV2.
type LookupVolumeV2ResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeV2Result)(nil)).Elem()
}

func (o LookupVolumeV2ResultOutput) ToLookupVolumeV2ResultOutput() LookupVolumeV2ResultOutput {
	return o
}

func (o LookupVolumeV2ResultOutput) ToLookupVolumeV2ResultOutputWithContext(ctx context.Context) LookupVolumeV2ResultOutput {
	return o
}

// Indicates if the volume is bootable.
func (o LookupVolumeV2ResultOutput) Bootable() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) string { return v.Bootable }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVolumeV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupVolumeV2ResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// See Argument Reference above.
func (o LookupVolumeV2ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) string { return v.Name }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupVolumeV2ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) string { return v.Region }).(pulumi.StringOutput)
}

// The size of the volume in GBs.
func (o LookupVolumeV2ResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) int { return v.Size }).(pulumi.IntOutput)
}

// The ID of the volume from which the current volume was created.
func (o LookupVolumeV2ResultOutput) SourceVolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) string { return v.SourceVolumeId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupVolumeV2ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) string { return v.Status }).(pulumi.StringOutput)
}

// The type of the volume.
func (o LookupVolumeV2ResultOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeV2Result) string { return v.VolumeType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeV2ResultOutput{})
}
