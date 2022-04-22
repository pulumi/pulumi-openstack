// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack address-scope.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networking.LookupAddressScope(ctx, &networking.LookupAddressScopeArgs{
// 			IpVersion: pulumi.IntRef(4),
// 			Name:      pulumi.StringRef("public_addressscope"),
// 			Shared:    pulumi.BoolRef(true),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAddressScope(ctx *pulumi.Context, args *LookupAddressScopeArgs, opts ...pulumi.InvokeOption) (*LookupAddressScopeResult, error) {
	var rv LookupAddressScopeResult
	err := ctx.Invoke("openstack:networking/getAddressScope:getAddressScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAddressScope.
type LookupAddressScopeArgs struct {
	// IP version.
	IpVersion *int `pulumi:"ipVersion"`
	// Name of the address-scope.
	Name *string `pulumi:"name"`
	// The owner of the address-scope.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve address-scopes. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// Indicates whether this address-scope is shared across
	// all projects.
	Shared *bool `pulumi:"shared"`
}

// A collection of values returned by getAddressScope.
type LookupAddressScopeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IpVersion *int `pulumi:"ipVersion"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	ProjectId *string `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// See Argument Reference above.
	Shared *bool `pulumi:"shared"`
}

func LookupAddressScopeOutput(ctx *pulumi.Context, args LookupAddressScopeOutputArgs, opts ...pulumi.InvokeOption) LookupAddressScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAddressScopeResult, error) {
			args := v.(LookupAddressScopeArgs)
			r, err := LookupAddressScope(ctx, &args, opts...)
			var s LookupAddressScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAddressScopeResultOutput)
}

// A collection of arguments for invoking getAddressScope.
type LookupAddressScopeOutputArgs struct {
	// IP version.
	IpVersion pulumi.IntPtrInput `pulumi:"ipVersion"`
	// Name of the address-scope.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The owner of the address-scope.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve address-scopes. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Indicates whether this address-scope is shared across
	// all projects.
	Shared pulumi.BoolPtrInput `pulumi:"shared"`
}

func (LookupAddressScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddressScopeArgs)(nil)).Elem()
}

// A collection of values returned by getAddressScope.
type LookupAddressScopeResultOutput struct{ *pulumi.OutputState }

func (LookupAddressScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAddressScopeResult)(nil)).Elem()
}

func (o LookupAddressScopeResultOutput) ToLookupAddressScopeResultOutput() LookupAddressScopeResultOutput {
	return o
}

func (o LookupAddressScopeResultOutput) ToLookupAddressScopeResultOutputWithContext(ctx context.Context) LookupAddressScopeResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAddressScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAddressScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupAddressScopeResultOutput) IpVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAddressScopeResult) *int { return v.IpVersion }).(pulumi.IntPtrOutput)
}

// See Argument Reference above.
func (o LookupAddressScopeResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAddressScopeResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupAddressScopeResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAddressScopeResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupAddressScopeResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAddressScopeResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupAddressScopeResultOutput) Shared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAddressScopeResult) *bool { return v.Shared }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAddressScopeResultOutput{})
}
