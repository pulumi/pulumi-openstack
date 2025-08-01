// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack QoS policy.
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
//			_, err := networking.LookupQosPolicy(ctx, &networking.LookupQosPolicyArgs{
//				Name: pulumi.StringRef("qos_policy_1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupQosPolicy(ctx *pulumi.Context, args *LookupQosPolicyArgs, opts ...pulumi.InvokeOption) (*LookupQosPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQosPolicyResult
	err := ctx.Invoke("openstack:networking/getQosPolicy:getQosPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQosPolicy.
type LookupQosPolicyArgs struct {
	// The human-readable description for the QoS policy.
	Description *string `pulumi:"description"`
	// Whether the QoS policy is default policy or not.
	IsDefault *bool `pulumi:"isDefault"`
	// The name of the QoS policy.
	Name *string `pulumi:"name"`
	// The owner of the QoS policy.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to retrieve a QoS policy ID. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// Whether this QoS policy is shared across all projects.
	Shared *bool `pulumi:"shared"`
	// The list of QoS policy tags to filter.
	Tags []string `pulumi:"tags"`
}

// A collection of values returned by getQosPolicy.
type LookupQosPolicyResult struct {
	// The set of string tags applied on the QoS policy.
	AllTags []string `pulumi:"allTags"`
	// The time at which QoS policy was created.
	CreatedAt string `pulumi:"createdAt"`
	// See Argument Reference above.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IsDefault bool `pulumi:"isDefault"`
	// See Argument Reference above.
	Name      string `pulumi:"name"`
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The revision number of the QoS policy.
	RevisionNumber int `pulumi:"revisionNumber"`
	// See Argument Reference above.
	Shared bool     `pulumi:"shared"`
	Tags   []string `pulumi:"tags"`
	// The time at which QoS policy was created.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupQosPolicyOutput(ctx *pulumi.Context, args LookupQosPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupQosPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupQosPolicyResultOutput, error) {
			args := v.(LookupQosPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:networking/getQosPolicy:getQosPolicy", args, LookupQosPolicyResultOutput{}, options).(LookupQosPolicyResultOutput), nil
		}).(LookupQosPolicyResultOutput)
}

// A collection of arguments for invoking getQosPolicy.
type LookupQosPolicyOutputArgs struct {
	// The human-readable description for the QoS policy.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Whether the QoS policy is default policy or not.
	IsDefault pulumi.BoolPtrInput `pulumi:"isDefault"`
	// The name of the QoS policy.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The owner of the QoS policy.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to retrieve a QoS policy ID. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Whether this QoS policy is shared across all projects.
	Shared pulumi.BoolPtrInput `pulumi:"shared"`
	// The list of QoS policy tags to filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
}

func (LookupQosPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQosPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getQosPolicy.
type LookupQosPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupQosPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQosPolicyResult)(nil)).Elem()
}

func (o LookupQosPolicyResultOutput) ToLookupQosPolicyResultOutput() LookupQosPolicyResultOutput {
	return o
}

func (o LookupQosPolicyResultOutput) ToLookupQosPolicyResultOutputWithContext(ctx context.Context) LookupQosPolicyResultOutput {
	return o
}

// The set of string tags applied on the QoS policy.
func (o LookupQosPolicyResultOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) []string { return v.AllTags }).(pulumi.StringArrayOutput)
}

// The time at which QoS policy was created.
func (o LookupQosPolicyResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQosPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosPolicyResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o LookupQosPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQosPolicyResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

// The revision number of the QoS policy.
func (o LookupQosPolicyResultOutput) RevisionNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) int { return v.RevisionNumber }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupQosPolicyResultOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) bool { return v.Shared }).(pulumi.BoolOutput)
}

func (o LookupQosPolicyResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The time at which QoS policy was created.
func (o LookupQosPolicyResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosPolicyResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQosPolicyResultOutput{})
}
