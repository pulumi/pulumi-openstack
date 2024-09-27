// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack router.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.LookupRouter(ctx, &networking.LookupRouterArgs{
//				Name: pulumi.StringRef("router_1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRouter(ctx *pulumi.Context, args *LookupRouterArgs, opts ...pulumi.InvokeOption) (*LookupRouterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouterResult
	err := ctx.Invoke("openstack:networking/getRouter:getRouter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouter.
type LookupRouterArgs struct {
	// Administrative up/down status for the router (must be "true" or "false" if provided).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description of the router.
	Description *string `pulumi:"description"`
	// Indicates whether or not to get a distributed router.
	Distributed *bool `pulumi:"distributed"`
	// The value that points out if the Source NAT is enabled on the router.
	EnableSnat *bool `pulumi:"enableSnat"`
	// The name of the router.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve router ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The UUID of the router resource.
	RouterId *string `pulumi:"routerId"`
	// The status of the router (ACTIVE/DOWN).
	Status *string `pulumi:"status"`
	// The list of router tags to filter.
	Tags []string `pulumi:"tags"`
	// The owner of the router.
	TenantId *string `pulumi:"tenantId"`
}

// A collection of values returned by getRouter.
type LookupRouterResult struct {
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The set of string tags applied on the router.
	AllTags []string `pulumi:"allTags"`
	// The availability zone that is used to make router resources highly available.
	AvailabilityZoneHints []string `pulumi:"availabilityZoneHints"`
	Description           *string  `pulumi:"description"`
	Distributed           *bool    `pulumi:"distributed"`
	// The value that points out if the Source NAT is enabled on the router.
	EnableSnat bool `pulumi:"enableSnat"`
	// The external fixed IPs of the router.
	ExternalFixedIps []GetRouterExternalFixedIp `pulumi:"externalFixedIps"`
	// The network UUID of an external gateway for the router.
	ExternalNetworkId string `pulumi:"externalNetworkId"`
	// The provider-assigned unique ID for this managed resource.
	Id       string   `pulumi:"id"`
	Name     *string  `pulumi:"name"`
	Region   *string  `pulumi:"region"`
	RouterId *string  `pulumi:"routerId"`
	Status   *string  `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *string  `pulumi:"tenantId"`
}

func LookupRouterOutput(ctx *pulumi.Context, args LookupRouterOutputArgs, opts ...pulumi.InvokeOption) LookupRouterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterResultOutput, error) {
			args := v.(LookupRouterArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupRouterResult
			secret, err := ctx.InvokePackageRaw("openstack:networking/getRouter:getRouter", args, &rv, "", opts...)
			if err != nil {
				return LookupRouterResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupRouterResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupRouterResultOutput), nil
			}
			return output, nil
		}).(LookupRouterResultOutput)
}

// A collection of arguments for invoking getRouter.
type LookupRouterOutputArgs struct {
	// Administrative up/down status for the router (must be "true" or "false" if provided).
	AdminStateUp pulumi.BoolPtrInput `pulumi:"adminStateUp"`
	// Human-readable description of the router.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Indicates whether or not to get a distributed router.
	Distributed pulumi.BoolPtrInput `pulumi:"distributed"`
	// The value that points out if the Source NAT is enabled on the router.
	EnableSnat pulumi.BoolPtrInput `pulumi:"enableSnat"`
	// The name of the router.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve router ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The UUID of the router resource.
	RouterId pulumi.StringPtrInput `pulumi:"routerId"`
	// The status of the router (ACTIVE/DOWN).
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The list of router tags to filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The owner of the router.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupRouterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterArgs)(nil)).Elem()
}

// A collection of values returned by getRouter.
type LookupRouterResultOutput struct{ *pulumi.OutputState }

func (LookupRouterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterResult)(nil)).Elem()
}

func (o LookupRouterResultOutput) ToLookupRouterResultOutput() LookupRouterResultOutput {
	return o
}

func (o LookupRouterResultOutput) ToLookupRouterResultOutputWithContext(ctx context.Context) LookupRouterResultOutput {
	return o
}

func (o LookupRouterResultOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRouterResult) *bool { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// The set of string tags applied on the router.
func (o LookupRouterResultOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouterResult) []string { return v.AllTags }).(pulumi.StringArrayOutput)
}

// The availability zone that is used to make router resources highly available.
func (o LookupRouterResultOutput) AvailabilityZoneHints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouterResult) []string { return v.AvailabilityZoneHints }).(pulumi.StringArrayOutput)
}

func (o LookupRouterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRouterResultOutput) Distributed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRouterResult) *bool { return v.Distributed }).(pulumi.BoolPtrOutput)
}

// The value that points out if the Source NAT is enabled on the router.
func (o LookupRouterResultOutput) EnableSnat() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRouterResult) bool { return v.EnableSnat }).(pulumi.BoolOutput)
}

// The external fixed IPs of the router.
func (o LookupRouterResultOutput) ExternalFixedIps() GetRouterExternalFixedIpArrayOutput {
	return o.ApplyT(func(v LookupRouterResult) []GetRouterExternalFixedIp { return v.ExternalFixedIps }).(GetRouterExternalFixedIpArrayOutput)
}

// The network UUID of an external gateway for the router.
func (o LookupRouterResultOutput) ExternalNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.ExternalNetworkId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRouterResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupRouterResultOutput) RouterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterResult) *string { return v.RouterId }).(pulumi.StringPtrOutput)
}

func (o LookupRouterResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupRouterResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouterResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupRouterResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterResultOutput{})
}
