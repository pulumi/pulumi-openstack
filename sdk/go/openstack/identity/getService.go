// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to get the ID of an OpenStack service.
//
// > **Note:** This usually requires admin privileges.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.GetService(ctx, &identity.GetServiceArgs{
//				Name: pulumi.StringRef("keystone"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetService(ctx *pulumi.Context, args *GetServiceArgs, opts ...pulumi.InvokeOption) (*GetServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceResult
	err := ctx.Invoke("openstack:identity/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// The service status.
	Enabled *bool `pulumi:"enabled"`
	// The service name.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The service type.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getService.
type GetServiceResult struct {
	// The service description.
	Description string `pulumi:"description"`
	// See Argument Reference above.
	Enabled *bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	Type *string `pulumi:"type"`
}

func GetServiceOutput(ctx *pulumi.Context, args GetServiceOutputArgs, opts ...pulumi.InvokeOption) GetServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceResult, error) {
			args := v.(GetServiceArgs)
			r, err := GetService(ctx, &args, opts...)
			var s GetServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceResultOutput)
}

// A collection of arguments for invoking getService.
type GetServiceOutputArgs struct {
	// The service status.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// The service name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The service type.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceArgs)(nil)).Elem()
}

// A collection of values returned by getService.
type GetServiceResultOutput struct{ *pulumi.OutputState }

func (GetServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceResult)(nil)).Elem()
}

func (o GetServiceResultOutput) ToGetServiceResultOutput() GetServiceResultOutput {
	return o
}

func (o GetServiceResultOutput) ToGetServiceResultOutputWithContext(ctx context.Context) GetServiceResultOutput {
	return o
}

func (o GetServiceResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetServiceResult] {
	return pulumix.Output[GetServiceResult]{
		OutputState: o.OutputState,
	}
}

// The service description.
func (o GetServiceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Description }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetServiceResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetServiceResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetServiceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetServiceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetServiceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceResultOutput{})
}
