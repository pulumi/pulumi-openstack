// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of Block Storage availability zones from OpenStack
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
//			_, err := blockstorage.GetAvailabilityZonesV3(ctx, &blockstorage.GetAvailabilityZonesV3Args{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAvailabilityZonesV3(ctx *pulumi.Context, args *GetAvailabilityZonesV3Args, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesV3Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAvailabilityZonesV3Result
	err := ctx.Invoke("openstack:blockstorage/getAvailabilityZonesV3:getAvailabilityZonesV3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZonesV3.
type GetAvailabilityZonesV3Args struct {
	// The region in which to obtain the Block Storage client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The `state` of the availability zones to match. Can
	// either be `available` or `unavailable`. Default is `available`.
	State *string `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZonesV3.
type GetAvailabilityZonesV3Result struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The names of the availability zones, ordered alphanumerically, that
	// match the queried `state`.
	Names []string `pulumi:"names"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	State *string `pulumi:"state"`
}

func GetAvailabilityZonesV3Output(ctx *pulumi.Context, args GetAvailabilityZonesV3OutputArgs, opts ...pulumi.InvokeOption) GetAvailabilityZonesV3ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAvailabilityZonesV3ResultOutput, error) {
			args := v.(GetAvailabilityZonesV3Args)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetAvailabilityZonesV3Result
			secret, err := ctx.InvokePackageRaw("openstack:blockstorage/getAvailabilityZonesV3:getAvailabilityZonesV3", args, &rv, "", opts...)
			if err != nil {
				return GetAvailabilityZonesV3ResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetAvailabilityZonesV3ResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetAvailabilityZonesV3ResultOutput), nil
			}
			return output, nil
		}).(GetAvailabilityZonesV3ResultOutput)
}

// A collection of arguments for invoking getAvailabilityZonesV3.
type GetAvailabilityZonesV3OutputArgs struct {
	// The region in which to obtain the Block Storage client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The `state` of the availability zones to match. Can
	// either be `available` or `unavailable`. Default is `available`.
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (GetAvailabilityZonesV3OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesV3Args)(nil)).Elem()
}

// A collection of values returned by getAvailabilityZonesV3.
type GetAvailabilityZonesV3ResultOutput struct{ *pulumi.OutputState }

func (GetAvailabilityZonesV3ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesV3Result)(nil)).Elem()
}

func (o GetAvailabilityZonesV3ResultOutput) ToGetAvailabilityZonesV3ResultOutput() GetAvailabilityZonesV3ResultOutput {
	return o
}

func (o GetAvailabilityZonesV3ResultOutput) ToGetAvailabilityZonesV3ResultOutputWithContext(ctx context.Context) GetAvailabilityZonesV3ResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAvailabilityZonesV3ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZonesV3Result) string { return v.Id }).(pulumi.StringOutput)
}

// The names of the availability zones, ordered alphanumerically, that
// match the queried `state`.
func (o GetAvailabilityZonesV3ResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesV3Result) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o GetAvailabilityZonesV3ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZonesV3Result) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetAvailabilityZonesV3ResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAvailabilityZonesV3Result) *string { return v.State }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAvailabilityZonesV3ResultOutput{})
}
