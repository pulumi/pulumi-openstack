// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to get firewall policy information of an available OpenStack firewall policy.
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("openstack:firewall/getPolicy:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyArgs struct {
	// The name of the firewall policy.
	Name *string `pulumi:"name"`
	// The ID of the firewall policy.
	PolicyId *string `pulumi:"policyId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve firewall policy ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The owner of the firewall policy.
	TenantId *string `pulumi:"tenantId"`
}

// A collection of values returned by getPolicy.
type LookupPolicyResult struct {
	// The audit status of the firewall policy.
	Audited bool `pulumi:"audited"`
	// The description of the firewall policy.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	PolicyId *string `pulumi:"policyId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The array of one or more firewall rules that comprise the policy.
	Rules []string `pulumi:"rules"`
	// The sharing status of the firewall policy.
	Shared bool `pulumi:"shared"`
	// See Argument Reference above.
	TenantId string `pulumi:"tenantId"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyOutputArgs struct {
	// The name of the firewall policy.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the firewall policy.
	PolicyId pulumi.StringPtrInput `pulumi:"policyId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve firewall policy ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The owner of the firewall policy.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy.
type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyResult] {
	return pulumix.Output[LookupPolicyResult]{
		OutputState: o.OutputState,
	}
}

// The audit status of the firewall policy.
func (o LookupPolicyResultOutput) Audited() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPolicyResult) bool { return v.Audited }).(pulumi.BoolOutput)
}

// The description of the firewall policy.
func (o LookupPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupPolicyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupPolicyResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

// The array of one or more firewall rules that comprise the policy.
func (o LookupPolicyResultOutput) Rules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []string { return v.Rules }).(pulumi.StringArrayOutput)
}

// The sharing status of the firewall policy.
func (o LookupPolicyResultOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPolicyResult) bool { return v.Shared }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o LookupPolicyResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
