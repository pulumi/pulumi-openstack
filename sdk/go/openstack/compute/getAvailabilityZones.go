// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of availability zones from OpenStack
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.GetAvailabilityZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAvailabilityZonesResult
	err := ctx.Invoke("openstack:compute/getAvailabilityZones:getAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// The `region` to fetch availability zones from, defaults to the provider's `region`
	Region *string `pulumi:"region"`
	// The `state` of the availability zones to match, default ("available").
	State *string `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The names of the availability zones, ordered alphanumerically, that match the queried `state`
	Names  []string `pulumi:"names"`
	Region string   `pulumi:"region"`
	State  *string  `pulumi:"state"`
}

func GetAvailabilityZonesOutput(ctx *pulumi.Context, args GetAvailabilityZonesOutputArgs, opts ...pulumi.InvokeOption) GetAvailabilityZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAvailabilityZonesResultOutput, error) {
			args := v.(GetAvailabilityZonesArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetAvailabilityZonesResult
			secret, err := ctx.InvokePackageRaw("openstack:compute/getAvailabilityZones:getAvailabilityZones", args, &rv, "", opts...)
			if err != nil {
				return GetAvailabilityZonesResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetAvailabilityZonesResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetAvailabilityZonesResultOutput), nil
			}
			return output, nil
		}).(GetAvailabilityZonesResultOutput)
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesOutputArgs struct {
	// The `region` to fetch availability zones from, defaults to the provider's `region`
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The `state` of the availability zones to match, default ("available").
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (GetAvailabilityZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesArgs)(nil)).Elem()
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResultOutput struct{ *pulumi.OutputState }

func (GetAvailabilityZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesResult)(nil)).Elem()
}

func (o GetAvailabilityZonesResultOutput) ToGetAvailabilityZonesResultOutput() GetAvailabilityZonesResultOutput {
	return o
}

func (o GetAvailabilityZonesResultOutput) ToGetAvailabilityZonesResultOutputWithContext(ctx context.Context) GetAvailabilityZonesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAvailabilityZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The names of the availability zones, ordered alphanumerically, that match the queried `state`
func (o GetAvailabilityZonesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAvailabilityZonesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetAvailabilityZonesResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAvailabilityZonesResultOutput{})
}
