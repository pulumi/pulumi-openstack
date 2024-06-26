// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a v1 firewall rule resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/firewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.NewRule(ctx, "rule_1", &firewall.RuleArgs{
//				Name:            pulumi.String("my_rule"),
//				Description:     pulumi.String("drop TELNET traffic"),
//				Action:          pulumi.String("deny"),
//				Protocol:        pulumi.String("tcp"),
//				DestinationPort: pulumi.String("23"),
//				Enabled:         pulumi.Bool(true),
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
// $ pulumi import openstack:firewall/rule:Rule rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
// ```
type Rule struct {
	pulumi.CustomResourceState

	// Action to be taken ( must be "allow" or "deny") when the
	// firewall rule matches. Changing this updates the `action` of an existing
	// firewall rule.
	Action pulumi.StringOutput `pulumi:"action"`
	// A description for the firewall rule. Changing this
	// updates the `description` of an existing firewall rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination IP address on which the
	// firewall rule operates. Changing this updates the `destinationIpAddress`
	// of an existing firewall rule.
	DestinationIpAddress pulumi.StringPtrOutput `pulumi:"destinationIpAddress"`
	// The destination port on which the firewall
	// rule operates. Changing this updates the `destinationPort` of an existing
	// firewall rule.
	DestinationPort pulumi.StringPtrOutput `pulumi:"destinationPort"`
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// IP version, either 4 (default) or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule.
	IpVersion pulumi.IntPtrOutput `pulumi:"ipVersion"`
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The region in which to obtain the v1 Compute client.
	// A Compute client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region pulumi.StringOutput `pulumi:"region"`
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress pulumi.StringPtrOutput `pulumi:"sourceIpAddress"`
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule.
	SourcePort pulumi.StringPtrOutput `pulumi:"sourcePort"`
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rule
	err := ctx.RegisterResource("openstack:firewall/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("openstack:firewall/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// Action to be taken ( must be "allow" or "deny") when the
	// firewall rule matches. Changing this updates the `action` of an existing
	// firewall rule.
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
	// firewall rule.
	DestinationPort *string `pulumi:"destinationPort"`
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled *bool `pulumi:"enabled"`
	// IP version, either 4 (default) or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule.
	IpVersion *int `pulumi:"ipVersion"`
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name *string `pulumi:"name"`
	// The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule.
	Protocol *string `pulumi:"protocol"`
	// The region in which to obtain the v1 Compute client.
	// A Compute client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region *string `pulumi:"region"`
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule.
	SourcePort *string `pulumi:"sourcePort"`
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type RuleState struct {
	// Action to be taken ( must be "allow" or "deny") when the
	// firewall rule matches. Changing this updates the `action` of an existing
	// firewall rule.
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
	// firewall rule.
	DestinationPort pulumi.StringPtrInput
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled pulumi.BoolPtrInput
	// IP version, either 4 (default) or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule.
	IpVersion pulumi.IntPtrInput
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name pulumi.StringPtrInput
	// The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule.
	Protocol pulumi.StringPtrInput
	// The region in which to obtain the v1 Compute client.
	// A Compute client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region pulumi.StringPtrInput
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress pulumi.StringPtrInput
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule.
	SourcePort pulumi.StringPtrInput
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// Action to be taken ( must be "allow" or "deny") when the
	// firewall rule matches. Changing this updates the `action` of an existing
	// firewall rule.
	Action string `pulumi:"action"`
	// A description for the firewall rule. Changing this
	// updates the `description` of an existing firewall rule.
	Description *string `pulumi:"description"`
	// The destination IP address on which the
	// firewall rule operates. Changing this updates the `destinationIpAddress`
	// of an existing firewall rule.
	DestinationIpAddress *string `pulumi:"destinationIpAddress"`
	// The destination port on which the firewall
	// rule operates. Changing this updates the `destinationPort` of an existing
	// firewall rule.
	DestinationPort *string `pulumi:"destinationPort"`
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled *bool `pulumi:"enabled"`
	// IP version, either 4 (default) or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule.
	IpVersion *int `pulumi:"ipVersion"`
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name *string `pulumi:"name"`
	// The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule.
	Protocol string `pulumi:"protocol"`
	// The region in which to obtain the v1 Compute client.
	// A Compute client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region *string `pulumi:"region"`
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress *string `pulumi:"sourceIpAddress"`
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule.
	SourcePort *string `pulumi:"sourcePort"`
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// Action to be taken ( must be "allow" or "deny") when the
	// firewall rule matches. Changing this updates the `action` of an existing
	// firewall rule.
	Action pulumi.StringInput
	// A description for the firewall rule. Changing this
	// updates the `description` of an existing firewall rule.
	Description pulumi.StringPtrInput
	// The destination IP address on which the
	// firewall rule operates. Changing this updates the `destinationIpAddress`
	// of an existing firewall rule.
	DestinationIpAddress pulumi.StringPtrInput
	// The destination port on which the firewall
	// rule operates. Changing this updates the `destinationPort` of an existing
	// firewall rule.
	DestinationPort pulumi.StringPtrInput
	// Enabled status for the firewall rule (must be "true"
	// or "false" if provided - defaults to "true"). Changing this updates the
	// `enabled` status of an existing firewall rule.
	Enabled pulumi.BoolPtrInput
	// IP version, either 4 (default) or 6. Changing this
	// updates the `ipVersion` of an existing firewall rule.
	IpVersion pulumi.IntPtrInput
	// A unique name for the firewall rule. Changing this
	// updates the `name` of an existing firewall rule.
	Name pulumi.StringPtrInput
	// The protocol type on which the firewall rule operates.
	// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
	// `protocol` of an existing firewall rule.
	Protocol pulumi.StringInput
	// The region in which to obtain the v1 Compute client.
	// A Compute client is needed to create a firewall rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall rule.
	Region pulumi.StringPtrInput
	// The source IP address on which the firewall
	// rule operates. Changing this updates the `sourceIpAddress` of an existing
	// firewall rule.
	SourceIpAddress pulumi.StringPtrInput
	// The source port on which the firewall
	// rule operates. Changing this updates the `sourcePort` of an existing
	// firewall rule.
	SourcePort pulumi.StringPtrInput
	// The owner of the firewall rule. Required if admin
	// wants to create a firewall rule for another tenant. Changing this creates a
	// new firewall rule.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//	RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//	RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

// Action to be taken ( must be "allow" or "deny") when the
// firewall rule matches. Changing this updates the `action` of an existing
// firewall rule.
func (o RuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// A description for the firewall rule. Changing this
// updates the `description` of an existing firewall rule.
func (o RuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination IP address on which the
// firewall rule operates. Changing this updates the `destinationIpAddress`
// of an existing firewall rule.
func (o RuleOutput) DestinationIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.DestinationIpAddress }).(pulumi.StringPtrOutput)
}

// The destination port on which the firewall
// rule operates. Changing this updates the `destinationPort` of an existing
// firewall rule.
func (o RuleOutput) DestinationPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.DestinationPort }).(pulumi.StringPtrOutput)
}

// Enabled status for the firewall rule (must be "true"
// or "false" if provided - defaults to "true"). Changing this updates the
// `enabled` status of an existing firewall rule.
func (o RuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// IP version, either 4 (default) or 6. Changing this
// updates the `ipVersion` of an existing firewall rule.
func (o RuleOutput) IpVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.IntPtrOutput { return v.IpVersion }).(pulumi.IntPtrOutput)
}

// A unique name for the firewall rule. Changing this
// updates the `name` of an existing firewall rule.
func (o RuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protocol type on which the firewall rule operates.
// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
// `protocol` of an existing firewall rule.
func (o RuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The region in which to obtain the v1 Compute client.
// A Compute client is needed to create a firewall rule. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// firewall rule.
func (o RuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The source IP address on which the firewall
// rule operates. Changing this updates the `sourceIpAddress` of an existing
// firewall rule.
func (o RuleOutput) SourceIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.SourceIpAddress }).(pulumi.StringPtrOutput)
}

// The source port on which the firewall
// rule operates. Changing this updates the `sourcePort` of an existing
// firewall rule.
func (o RuleOutput) SourcePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.SourcePort }).(pulumi.StringPtrOutput)
}

// The owner of the firewall rule. Required if admin
// wants to create a firewall rule for another tenant. Changing this creates a
// new firewall rule.
func (o RuleOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Map of additional options.
func (o RuleOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *Rule) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].([]*Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].(map[string]*Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleArrayInput)(nil)).Elem(), RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleMapInput)(nil)).Elem(), RuleMap{})
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}
