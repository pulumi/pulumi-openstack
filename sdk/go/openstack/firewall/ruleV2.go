// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a v2 firewall rule resource within OpenStack.
//
// > **Note:** Firewall v2 has no support for OVN currently.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/firewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.NewRuleV2(ctx, "rule2", &firewall.RuleV2Args{
//				Action:          pulumi.String("deny"),
//				Description:     pulumi.String("drop TELNET traffic"),
//				DestinationPort: pulumi.String("23"),
//				Enabled:         pulumi.Bool(true),
//				Protocol:        pulumi.String("tcp"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Firewall Rules can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:firewall/ruleV2:RuleV2 rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
//
// ```
type RuleV2 struct {
	pulumi.CustomResourceState

	// Action to be taken (must be "allow", "deny" or "reject")
	// when the firewall rule matches. Changing this updates the `action` of an
	// existing firewall rule. Default is `deny`.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// A description for the firewall rule. Changing this
	// updates the `description` of an existing firewall rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination IP address on which the
	// firewall rule operates. Changing this updates the `destinationIpAddress`
	// of an existing firewall rule.
	DestinationIpAddress pulumi.StringPtrOutput `pulumi:"destinationIpAddress"`
	// The destination port on which the firewall
	// rule operates. Changing this updates the `destinationPort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	DestinationPort pulumi.StringPtrOutput `pulumi:"destinationPort"`
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// IP version, either 4 or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule. Default is `4`.
	IpVersion pulumi.IntPtrOutput `pulumi:"ipVersion"`
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Optional; Required if `sourcePort` or `destinationPort` is not
	// empty) The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule. Default is `any`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region pulumi.StringOutput `pulumi:"region"`
	// Sharing status of the firewall rule (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. On
	Shared pulumi.BoolPtrOutput `pulumi:"shared"`
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress pulumi.StringPtrOutput `pulumi:"sourceIpAddress"`
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	SourcePort pulumi.StringPtrOutput `pulumi:"sourcePort"`
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
}

// NewRuleV2 registers a new resource with the given unique name, arguments, and options.
func NewRuleV2(ctx *pulumi.Context,
	name string, args *RuleV2Args, opts ...pulumi.ResourceOption) (*RuleV2, error) {
	if args == nil {
		args = &RuleV2Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RuleV2
	err := ctx.RegisterResource("openstack:firewall/ruleV2:RuleV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleV2 gets an existing RuleV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleV2State, opts ...pulumi.ResourceOption) (*RuleV2, error) {
	var resource RuleV2
	err := ctx.ReadResource("openstack:firewall/ruleV2:RuleV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleV2 resources.
type ruleV2State struct {
	// Action to be taken (must be "allow", "deny" or "reject")
	// when the firewall rule matches. Changing this updates the `action` of an
	// existing firewall rule. Default is `deny`.
	Action *string `pulumi:"action"`
	// A description for the firewall rule. Changing this
	// updates the `description` of an existing firewall rule.
	Description *string `pulumi:"description"`
	// The destination IP address on which the
	// firewall rule operates. Changing this updates the `destinationIpAddress`
	// of an existing firewall rule.
	DestinationIpAddress *string `pulumi:"destinationIpAddress"`
	// The destination port on which the firewall
	// rule operates. Changing this updates the `destinationPort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	DestinationPort *string `pulumi:"destinationPort"`
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled *bool `pulumi:"enabled"`
	// IP version, either 4 or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule. Default is `4`.
	IpVersion *int `pulumi:"ipVersion"`
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name *string `pulumi:"name"`
	// (Optional; Required if `sourcePort` or `destinationPort` is not
	// empty) The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule. Default is `any`.
	Protocol *string `pulumi:"protocol"`
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region *string `pulumi:"region"`
	// Sharing status of the firewall rule (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. On
	Shared *bool `pulumi:"shared"`
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	SourcePort *string `pulumi:"sourcePort"`
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId *string `pulumi:"tenantId"`
}

type RuleV2State struct {
	// Action to be taken (must be "allow", "deny" or "reject")
	// when the firewall rule matches. Changing this updates the `action` of an
	// existing firewall rule. Default is `deny`.
	Action pulumi.StringPtrInput
	// A description for the firewall rule. Changing this
	// updates the `description` of an existing firewall rule.
	Description pulumi.StringPtrInput
	// The destination IP address on which the
	// firewall rule operates. Changing this updates the `destinationIpAddress`
	// of an existing firewall rule.
	DestinationIpAddress pulumi.StringPtrInput
	// The destination port on which the firewall
	// rule operates. Changing this updates the `destinationPort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	DestinationPort pulumi.StringPtrInput
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled pulumi.BoolPtrInput
	// IP version, either 4 or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule. Default is `4`.
	IpVersion pulumi.IntPtrInput
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name pulumi.StringPtrInput
	// (Optional; Required if `sourcePort` or `destinationPort` is not
	// empty) The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule. Default is `any`.
	Protocol pulumi.StringPtrInput
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region pulumi.StringPtrInput
	// Sharing status of the firewall rule (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. On
	Shared pulumi.BoolPtrInput
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress pulumi.StringPtrInput
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	SourcePort pulumi.StringPtrInput
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId pulumi.StringPtrInput
}

func (RuleV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleV2State)(nil)).Elem()
}

type ruleV2Args struct {
	// Action to be taken (must be "allow", "deny" or "reject")
	// when the firewall rule matches. Changing this updates the `action` of an
	// existing firewall rule. Default is `deny`.
	Action *string `pulumi:"action"`
	// A description for the firewall rule. Changing this
	// updates the `description` of an existing firewall rule.
	Description *string `pulumi:"description"`
	// The destination IP address on which the
	// firewall rule operates. Changing this updates the `destinationIpAddress`
	// of an existing firewall rule.
	DestinationIpAddress *string `pulumi:"destinationIpAddress"`
	// The destination port on which the firewall
	// rule operates. Changing this updates the `destinationPort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	DestinationPort *string `pulumi:"destinationPort"`
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled *bool `pulumi:"enabled"`
	// IP version, either 4 or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule. Default is `4`.
	IpVersion *int `pulumi:"ipVersion"`
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name *string `pulumi:"name"`
	// (Optional; Required if `sourcePort` or `destinationPort` is not
	// empty) The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule. Default is `any`.
	Protocol *string `pulumi:"protocol"`
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region *string `pulumi:"region"`
	// Sharing status of the firewall rule (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. On
	Shared *bool `pulumi:"shared"`
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	SourcePort *string `pulumi:"sourcePort"`
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a RuleV2 resource.
type RuleV2Args struct {
	// Action to be taken (must be "allow", "deny" or "reject")
	// when the firewall rule matches. Changing this updates the `action` of an
	// existing firewall rule. Default is `deny`.
	Action pulumi.StringPtrInput
	// A description for the firewall rule. Changing this
	// updates the `description` of an existing firewall rule.
	Description pulumi.StringPtrInput
	// The destination IP address on which the
	// firewall rule operates. Changing this updates the `destinationIpAddress`
	// of an existing firewall rule.
	DestinationIpAddress pulumi.StringPtrInput
	// The destination port on which the firewall
	// rule operates. Changing this updates the `destinationPort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	DestinationPort pulumi.StringPtrInput
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled pulumi.BoolPtrInput
	// IP version, either 4 or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule. Default is `4`.
	IpVersion pulumi.IntPtrInput
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name pulumi.StringPtrInput
	// (Optional; Required if `sourcePort` or `destinationPort` is not
	// empty) The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule. Default is `any`.
	Protocol pulumi.StringPtrInput
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region pulumi.StringPtrInput
	// Sharing status of the firewall rule (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. On
	Shared pulumi.BoolPtrInput
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress pulumi.StringPtrInput
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule. Require not `any` or empty protocol.
	SourcePort pulumi.StringPtrInput
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId pulumi.StringPtrInput
}

func (RuleV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleV2Args)(nil)).Elem()
}

type RuleV2Input interface {
	pulumi.Input

	ToRuleV2Output() RuleV2Output
	ToRuleV2OutputWithContext(ctx context.Context) RuleV2Output
}

func (*RuleV2) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleV2)(nil)).Elem()
}

func (i *RuleV2) ToRuleV2Output() RuleV2Output {
	return i.ToRuleV2OutputWithContext(context.Background())
}

func (i *RuleV2) ToRuleV2OutputWithContext(ctx context.Context) RuleV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(RuleV2Output)
}

// RuleV2ArrayInput is an input type that accepts RuleV2Array and RuleV2ArrayOutput values.
// You can construct a concrete instance of `RuleV2ArrayInput` via:
//
//	RuleV2Array{ RuleV2Args{...} }
type RuleV2ArrayInput interface {
	pulumi.Input

	ToRuleV2ArrayOutput() RuleV2ArrayOutput
	ToRuleV2ArrayOutputWithContext(context.Context) RuleV2ArrayOutput
}

type RuleV2Array []RuleV2Input

func (RuleV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleV2)(nil)).Elem()
}

func (i RuleV2Array) ToRuleV2ArrayOutput() RuleV2ArrayOutput {
	return i.ToRuleV2ArrayOutputWithContext(context.Background())
}

func (i RuleV2Array) ToRuleV2ArrayOutputWithContext(ctx context.Context) RuleV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleV2ArrayOutput)
}

// RuleV2MapInput is an input type that accepts RuleV2Map and RuleV2MapOutput values.
// You can construct a concrete instance of `RuleV2MapInput` via:
//
//	RuleV2Map{ "key": RuleV2Args{...} }
type RuleV2MapInput interface {
	pulumi.Input

	ToRuleV2MapOutput() RuleV2MapOutput
	ToRuleV2MapOutputWithContext(context.Context) RuleV2MapOutput
}

type RuleV2Map map[string]RuleV2Input

func (RuleV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleV2)(nil)).Elem()
}

func (i RuleV2Map) ToRuleV2MapOutput() RuleV2MapOutput {
	return i.ToRuleV2MapOutputWithContext(context.Background())
}

func (i RuleV2Map) ToRuleV2MapOutputWithContext(ctx context.Context) RuleV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleV2MapOutput)
}

type RuleV2Output struct{ *pulumi.OutputState }

func (RuleV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleV2)(nil)).Elem()
}

func (o RuleV2Output) ToRuleV2Output() RuleV2Output {
	return o
}

func (o RuleV2Output) ToRuleV2OutputWithContext(ctx context.Context) RuleV2Output {
	return o
}

// Action to be taken (must be "allow", "deny" or "reject")
// when the firewall rule matches. Changing this updates the `action` of an
// existing firewall rule. Default is `deny`.
func (o RuleV2Output) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringPtrOutput { return v.Action }).(pulumi.StringPtrOutput)
}

// A description for the firewall rule. Changing this
// updates the `description` of an existing firewall rule.
func (o RuleV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination IP address on which the
// firewall rule operates. Changing this updates the `destinationIpAddress`
// of an existing firewall rule.
func (o RuleV2Output) DestinationIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringPtrOutput { return v.DestinationIpAddress }).(pulumi.StringPtrOutput)
}

// The destination port on which the firewall
// rule operates. Changing this updates the `destinationPort` of an existing
// firewall rule. Require not `any` or empty protocol.
func (o RuleV2Output) DestinationPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringPtrOutput { return v.DestinationPort }).(pulumi.StringPtrOutput)
}

// Enabled status for the firewall rule (must be "true"
// or "false" if provided - defaults to "true"). Changing this updates the
// `enabled` status of an existing firewall rule.
func (o RuleV2Output) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// IP version, either 4 or 6. Changing this
// updates the `ipVersion` of an existing firewall rule. Default is `4`.
func (o RuleV2Output) IpVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.IntPtrOutput { return v.IpVersion }).(pulumi.IntPtrOutput)
}

// A unique name for the firewall rule. Changing this
// updates the `name` of an existing firewall rule.
func (o RuleV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Optional; Required if `sourcePort` or `destinationPort` is not
// empty) The protocol type on which the firewall rule operates.
// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
// `protocol` of an existing firewall rule. Default is `any`.
func (o RuleV2Output) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The region in which to obtain the v2 networking client.
// A networking client is needed to create a firewall rule. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// firewall rule.
func (o RuleV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Sharing status of the firewall rule (must be "true"
// or "false" if provided). If this is "true" the policy is visible to, and
// can be used in, firewalls in other tenants. Changing this updates the
// `shared` status of an existing firewall policy. On
func (o RuleV2Output) Shared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.BoolPtrOutput { return v.Shared }).(pulumi.BoolPtrOutput)
}

// The source IP address on which the firewall
// rule operates. Changing this updates the `sourceIpAddress` of an existing
// firewall rule.
func (o RuleV2Output) SourceIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringPtrOutput { return v.SourceIpAddress }).(pulumi.StringPtrOutput)
}

// The source port on which the firewall
// rule operates. Changing this updates the `sourcePort` of an existing
// firewall rule. Require not `any` or empty protocol.
func (o RuleV2Output) SourcePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringPtrOutput { return v.SourcePort }).(pulumi.StringPtrOutput)
}

// The owner of the firewall rule. Required if admin
// wants to create a firewall rule for another tenant. Changing this creates a
// new firewall rule.
func (o RuleV2Output) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuleV2) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

type RuleV2ArrayOutput struct{ *pulumi.OutputState }

func (RuleV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleV2)(nil)).Elem()
}

func (o RuleV2ArrayOutput) ToRuleV2ArrayOutput() RuleV2ArrayOutput {
	return o
}

func (o RuleV2ArrayOutput) ToRuleV2ArrayOutputWithContext(ctx context.Context) RuleV2ArrayOutput {
	return o
}

func (o RuleV2ArrayOutput) Index(i pulumi.IntInput) RuleV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuleV2 {
		return vs[0].([]*RuleV2)[vs[1].(int)]
	}).(RuleV2Output)
}

type RuleV2MapOutput struct{ *pulumi.OutputState }

func (RuleV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleV2)(nil)).Elem()
}

func (o RuleV2MapOutput) ToRuleV2MapOutput() RuleV2MapOutput {
	return o
}

func (o RuleV2MapOutput) ToRuleV2MapOutputWithContext(ctx context.Context) RuleV2MapOutput {
	return o
}

func (o RuleV2MapOutput) MapIndex(k pulumi.StringInput) RuleV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuleV2 {
		return vs[0].(map[string]*RuleV2)[vs[1].(string)]
	}).(RuleV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleV2Input)(nil)).Elem(), &RuleV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleV2ArrayInput)(nil)).Elem(), RuleV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleV2MapInput)(nil)).Elem(), RuleV2Map{})
	pulumi.RegisterOutputType(RuleV2Output{})
	pulumi.RegisterOutputType(RuleV2ArrayOutput{})
	pulumi.RegisterOutputType(RuleV2MapOutput{})
}
