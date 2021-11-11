// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack flavor.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := 512
// 		opt1 := 1
// 		_, err := compute.LookupFlavor(ctx, &compute.LookupFlavorArgs{
// 			Ram:   &opt0,
// 			Vcpus: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupFlavor(ctx *pulumi.Context, args *LookupFlavorArgs, opts ...pulumi.InvokeOption) (*LookupFlavorResult, error) {
	var rv LookupFlavorResult
	err := ctx.Invoke("openstack:compute/getFlavor:getFlavor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlavor.
type LookupFlavorArgs struct {
	// The exact amount of disk (in gigabytes).
	Disk *int `pulumi:"disk"`
	// The ID of the flavor. Conflicts with the `name`,
	// `minRam` and `minDisk`
	FlavorId *string `pulumi:"flavorId"`
	// The flavor visibility.
	IsPublic *bool `pulumi:"isPublic"`
	// The minimum amount of disk (in gigabytes). Conflicts
	// with the `flavorId`.
	MinDisk *int `pulumi:"minDisk"`
	// The minimum amount of RAM (in megabytes). Conflicts
	// with the `flavorId`.
	MinRam *int `pulumi:"minRam"`
	// The name of the flavor. Conflicts with the `flavorId`.
	Name *string `pulumi:"name"`
	// The exact amount of RAM (in megabytes).
	Ram *int `pulumi:"ram"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The `rxTxFactor` of the flavor.
	RxTxFactor *float64 `pulumi:"rxTxFactor"`
	// The amount of swap (in gigabytes).
	Swap *int `pulumi:"swap"`
	// The amount of VCPUs.
	Vcpus *int `pulumi:"vcpus"`
}

// A collection of values returned by getFlavor.
type LookupFlavorResult struct {
	Disk *int `pulumi:"disk"`
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs map[string]interface{} `pulumi:"extraSpecs"`
	FlavorId   *string                `pulumi:"flavorId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	IsPublic   *bool    `pulumi:"isPublic"`
	MinDisk    *int     `pulumi:"minDisk"`
	MinRam     *int     `pulumi:"minRam"`
	Name       *string  `pulumi:"name"`
	Ram        *int     `pulumi:"ram"`
	Region     string   `pulumi:"region"`
	RxTxFactor *float64 `pulumi:"rxTxFactor"`
	Swap       *int     `pulumi:"swap"`
	Vcpus      *int     `pulumi:"vcpus"`
}

func LookupFlavorOutput(ctx *pulumi.Context, args LookupFlavorOutputArgs, opts ...pulumi.InvokeOption) LookupFlavorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFlavorResult, error) {
			args := v.(LookupFlavorArgs)
			r, err := LookupFlavor(ctx, &args, opts...)
			return *r, err
		}).(LookupFlavorResultOutput)
}

// A collection of arguments for invoking getFlavor.
type LookupFlavorOutputArgs struct {
	// The exact amount of disk (in gigabytes).
	Disk pulumi.IntPtrInput `pulumi:"disk"`
	// The ID of the flavor. Conflicts with the `name`,
	// `minRam` and `minDisk`
	FlavorId pulumi.StringPtrInput `pulumi:"flavorId"`
	// The flavor visibility.
	IsPublic pulumi.BoolPtrInput `pulumi:"isPublic"`
	// The minimum amount of disk (in gigabytes). Conflicts
	// with the `flavorId`.
	MinDisk pulumi.IntPtrInput `pulumi:"minDisk"`
	// The minimum amount of RAM (in megabytes). Conflicts
	// with the `flavorId`.
	MinRam pulumi.IntPtrInput `pulumi:"minRam"`
	// The name of the flavor. Conflicts with the `flavorId`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The exact amount of RAM (in megabytes).
	Ram pulumi.IntPtrInput `pulumi:"ram"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The `rxTxFactor` of the flavor.
	RxTxFactor pulumi.Float64PtrInput `pulumi:"rxTxFactor"`
	// The amount of swap (in gigabytes).
	Swap pulumi.IntPtrInput `pulumi:"swap"`
	// The amount of VCPUs.
	Vcpus pulumi.IntPtrInput `pulumi:"vcpus"`
}

func (LookupFlavorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlavorArgs)(nil)).Elem()
}

// A collection of values returned by getFlavor.
type LookupFlavorResultOutput struct{ *pulumi.OutputState }

func (LookupFlavorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlavorResult)(nil)).Elem()
}

func (o LookupFlavorResultOutput) ToLookupFlavorResultOutput() LookupFlavorResultOutput {
	return o
}

func (o LookupFlavorResultOutput) ToLookupFlavorResultOutputWithContext(ctx context.Context) LookupFlavorResultOutput {
	return o
}

func (o LookupFlavorResultOutput) Disk() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *int { return v.Disk }).(pulumi.IntPtrOutput)
}

// Key/Value pairs of metadata for the flavor.
func (o LookupFlavorResultOutput) ExtraSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v LookupFlavorResult) map[string]interface{} { return v.ExtraSpecs }).(pulumi.MapOutput)
}

func (o LookupFlavorResultOutput) FlavorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *string { return v.FlavorId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFlavorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlavorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFlavorResultOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *bool { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

func (o LookupFlavorResultOutput) MinDisk() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *int { return v.MinDisk }).(pulumi.IntPtrOutput)
}

func (o LookupFlavorResultOutput) MinRam() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *int { return v.MinRam }).(pulumi.IntPtrOutput)
}

func (o LookupFlavorResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFlavorResultOutput) Ram() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *int { return v.Ram }).(pulumi.IntPtrOutput)
}

func (o LookupFlavorResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFlavorResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupFlavorResultOutput) RxTxFactor() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *float64 { return v.RxTxFactor }).(pulumi.Float64PtrOutput)
}

func (o LookupFlavorResultOutput) Swap() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *int { return v.Swap }).(pulumi.IntPtrOutput)
}

func (o LookupFlavorResultOutput) Vcpus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlavorResult) *int { return v.Vcpus }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFlavorResultOutput{})
}
