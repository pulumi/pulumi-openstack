// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information of an available OpenStack firewall rule v2.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/firewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.LookupRuleV2(ctx, &firewall.LookupRuleV2Args{
//				Name: pulumi.StringRef("tf_test_rule"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2
func GetFwRuleV2(ctx *pulumi.Context, args *GetFwRuleV2Args, opts ...pulumi.InvokeOption) (*GetFwRuleV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFwRuleV2Result
	err := ctx.Invoke("openstack:index/getFwRuleV2:getFwRuleV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFwRuleV2.
type GetFwRuleV2Args struct {
	// Action to be taken when the firewall rule matches.
	Action *string `pulumi:"action"`
	// The description of the firewall rule.
	Description *string `pulumi:"description"`
	// The destination IP address on which the
	// firewall rule operates.
	DestinationIpAddress *string `pulumi:"destinationIpAddress"`
	// The destination port on which the firewall
	// rule operates.
	DestinationPort *string `pulumi:"destinationPort"`
	// Enabled status for the firewall rule.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the firewall policy the rule belongs to.
	FirewallPolicyIds []string `pulumi:"firewallPolicyIds"`
	// IP version, either 4 (default) or 6.
	IpVersion *int `pulumi:"ipVersion"`
	// The name of the firewall rule.
	Name *string `pulumi:"name"`
	// This argument conflicts and is interchangeable
	// with `tenantId`. The owner of the firewall rule.
	ProjectId *string `pulumi:"projectId"`
	// The protocol type on which the firewall rule operates.
	Protocol *string `pulumi:"protocol"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve firewall policy ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The ID of the firewall rule.
	RuleId *string `pulumi:"ruleId"`
	// The sharing status of the firewall policy.
	Shared *bool `pulumi:"shared"`
	// The source IP address on which the firewall
	// rule operates.
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// The source port on which the firewall
	// rule operates.
	SourcePort *string `pulumi:"sourcePort"`
	// This argument conflicts and is interchangeable
	// with `projectId`. The owner of the firewall rule.
	TenantId *string `pulumi:"tenantId"`
}

// A collection of values returned by getFwRuleV2.
type GetFwRuleV2Result struct {
	// See Argument Reference above.
	Action *string `pulumi:"action"`
	// See Argument Reference above.
	Description *string `pulumi:"description"`
	// See Argument Reference above.
	DestinationIpAddress *string `pulumi:"destinationIpAddress"`
	// See Argument Reference above.
	DestinationPort *string `pulumi:"destinationPort"`
	// See Argument Reference above.
	Enabled bool `pulumi:"enabled"`
	// The ID of the firewall policy the rule belongs to.
	FirewallPolicyIds []string `pulumi:"firewallPolicyIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IpVersion *int `pulumi:"ipVersion"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Protocol *string `pulumi:"protocol"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	RuleId *string `pulumi:"ruleId"`
	// See Argument Reference above.
	Shared bool `pulumi:"shared"`
	// See Argument Reference above.
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// See Argument Reference above.
	SourcePort *string `pulumi:"sourcePort"`
	// See Argument Reference above.
	TenantId string `pulumi:"tenantId"`
}

func GetFwRuleV2Output(ctx *pulumi.Context, args GetFwRuleV2OutputArgs, opts ...pulumi.InvokeOption) GetFwRuleV2ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetFwRuleV2ResultOutput, error) {
			args := v.(GetFwRuleV2Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:index/getFwRuleV2:getFwRuleV2", args, GetFwRuleV2ResultOutput{}, options).(GetFwRuleV2ResultOutput), nil
		}).(GetFwRuleV2ResultOutput)
}

// A collection of arguments for invoking getFwRuleV2.
type GetFwRuleV2OutputArgs struct {
	// Action to be taken when the firewall rule matches.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// The description of the firewall rule.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The destination IP address on which the
	// firewall rule operates.
	DestinationIpAddress pulumi.StringPtrInput `pulumi:"destinationIpAddress"`
	// The destination port on which the firewall
	// rule operates.
	DestinationPort pulumi.StringPtrInput `pulumi:"destinationPort"`
	// Enabled status for the firewall rule.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// The ID of the firewall policy the rule belongs to.
	FirewallPolicyIds pulumi.StringArrayInput `pulumi:"firewallPolicyIds"`
	// IP version, either 4 (default) or 6.
	IpVersion pulumi.IntPtrInput `pulumi:"ipVersion"`
	// The name of the firewall rule.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// This argument conflicts and is interchangeable
	// with `tenantId`. The owner of the firewall rule.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The protocol type on which the firewall rule operates.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve firewall policy ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The ID of the firewall rule.
	RuleId pulumi.StringPtrInput `pulumi:"ruleId"`
	// The sharing status of the firewall policy.
	Shared pulumi.BoolPtrInput `pulumi:"shared"`
	// The source IP address on which the firewall
	// rule operates.
	SourceIpAddress pulumi.StringPtrInput `pulumi:"sourceIpAddress"`
	// The source port on which the firewall
	// rule operates.
	SourcePort pulumi.StringPtrInput `pulumi:"sourcePort"`
	// This argument conflicts and is interchangeable
	// with `projectId`. The owner of the firewall rule.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (GetFwRuleV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFwRuleV2Args)(nil)).Elem()
}

// A collection of values returned by getFwRuleV2.
type GetFwRuleV2ResultOutput struct{ *pulumi.OutputState }

func (GetFwRuleV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFwRuleV2Result)(nil)).Elem()
}

func (o GetFwRuleV2ResultOutput) ToGetFwRuleV2ResultOutput() GetFwRuleV2ResultOutput {
	return o
}

func (o GetFwRuleV2ResultOutput) ToGetFwRuleV2ResultOutputWithContext(ctx context.Context) GetFwRuleV2ResultOutput {
	return o
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) DestinationIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.DestinationIpAddress }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) DestinationPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.DestinationPort }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The ID of the firewall policy the rule belongs to.
func (o GetFwRuleV2ResultOutput) FirewallPolicyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) []string { return v.FirewallPolicyIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFwRuleV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) IpVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *int { return v.IpVersion }).(pulumi.IntPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) bool { return v.Shared }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) SourceIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.SourceIpAddress }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) SourcePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) *string { return v.SourcePort }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetFwRuleV2ResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFwRuleV2Result) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFwRuleV2ResultOutput{})
}
