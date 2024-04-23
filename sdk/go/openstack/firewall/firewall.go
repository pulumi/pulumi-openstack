// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a v1 firewall resource within OpenStack.
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
//			rule1, err := firewall.NewRule(ctx, "rule_1", &firewall.RuleArgs{
//				Name:            pulumi.String("my-rule-1"),
//				Description:     pulumi.String("drop TELNET traffic"),
//				Action:          pulumi.String("deny"),
//				Protocol:        pulumi.String("tcp"),
//				DestinationPort: pulumi.String("23"),
//				Enabled:         pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			rule2, err := firewall.NewRule(ctx, "rule_2", &firewall.RuleArgs{
//				Name:            pulumi.String("my-rule-2"),
//				Description:     pulumi.String("drop NTP traffic"),
//				Action:          pulumi.String("deny"),
//				Protocol:        pulumi.String("udp"),
//				DestinationPort: pulumi.String("123"),
//				Enabled:         pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			policy1, err := firewall.NewPolicy(ctx, "policy_1", &firewall.PolicyArgs{
//				Name: pulumi.String("my-policy"),
//				Rules: pulumi.StringArray{
//					rule1.ID(),
//					rule2.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = firewall.NewFirewall(ctx, "firewall_1", &firewall.FirewallArgs{
//				Name:     pulumi.String("my-firewall"),
//				PolicyId: policy1.ID(),
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
// Firewalls can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:firewall/firewall:Firewall firewall_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97
// ```
type Firewall struct {
	pulumi.CustomResourceState

	// Administrative up/down status for the firewall
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall.
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// Router(s) to associate this firewall instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall. Conflicts with `noRouters`.
	AssociatedRouters pulumi.StringArrayOutput `pulumi:"associatedRouters"`
	// A description for the firewall. Changing this
	// updates the `description` of an existing firewall.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A name for the firewall. Changing this
	// updates the `name` of an existing firewall.
	Name pulumi.StringOutput `pulumi:"name"`
	// Should this firewall not be associated with any routers
	// (must be "true" or "false" if provide - defaults to "false").
	// Conflicts with `associatedRouters`.
	NoRouters pulumi.BoolPtrOutput `pulumi:"noRouters"`
	// The policy resource id for the firewall. Changing
	// this updates the `policyId` of an existing firewall.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall.
	Region pulumi.StringOutput `pulumi:"region"`
	// The owner of the floating IP. Required if admin wants
	// to create a firewall for another tenant. Changing this creates a new
	// firewall.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOption) (*Firewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Firewall
	err := ctx.RegisterResource("openstack:firewall/firewall:Firewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallState, opts ...pulumi.ResourceOption) (*Firewall, error) {
	var resource Firewall
	err := ctx.ReadResource("openstack:firewall/firewall:Firewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewall resources.
type firewallState struct {
	// Administrative up/down status for the firewall
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Router(s) to associate this firewall instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall. Conflicts with `noRouters`.
	AssociatedRouters []string `pulumi:"associatedRouters"`
	// A description for the firewall. Changing this
	// updates the `description` of an existing firewall.
	Description *string `pulumi:"description"`
	// A name for the firewall. Changing this
	// updates the `name` of an existing firewall.
	Name *string `pulumi:"name"`
	// Should this firewall not be associated with any routers
	// (must be "true" or "false" if provide - defaults to "false").
	// Conflicts with `associatedRouters`.
	NoRouters *bool `pulumi:"noRouters"`
	// The policy resource id for the firewall. Changing
	// this updates the `policyId` of an existing firewall.
	PolicyId *string `pulumi:"policyId"`
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall.
	Region *string `pulumi:"region"`
	// The owner of the floating IP. Required if admin wants
	// to create a firewall for another tenant. Changing this creates a new
	// firewall.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type FirewallState struct {
	// Administrative up/down status for the firewall
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall.
	AdminStateUp pulumi.BoolPtrInput
	// Router(s) to associate this firewall instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall. Conflicts with `noRouters`.
	AssociatedRouters pulumi.StringArrayInput
	// A description for the firewall. Changing this
	// updates the `description` of an existing firewall.
	Description pulumi.StringPtrInput
	// A name for the firewall. Changing this
	// updates the `name` of an existing firewall.
	Name pulumi.StringPtrInput
	// Should this firewall not be associated with any routers
	// (must be "true" or "false" if provide - defaults to "false").
	// Conflicts with `associatedRouters`.
	NoRouters pulumi.BoolPtrInput
	// The policy resource id for the firewall. Changing
	// this updates the `policyId` of an existing firewall.
	PolicyId pulumi.StringPtrInput
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall.
	Region pulumi.StringPtrInput
	// The owner of the floating IP. Required if admin wants
	// to create a firewall for another tenant. Changing this creates a new
	// firewall.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallState)(nil)).Elem()
}

type firewallArgs struct {
	// Administrative up/down status for the firewall
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Router(s) to associate this firewall instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall. Conflicts with `noRouters`.
	AssociatedRouters []string `pulumi:"associatedRouters"`
	// A description for the firewall. Changing this
	// updates the `description` of an existing firewall.
	Description *string `pulumi:"description"`
	// A name for the firewall. Changing this
	// updates the `name` of an existing firewall.
	Name *string `pulumi:"name"`
	// Should this firewall not be associated with any routers
	// (must be "true" or "false" if provide - defaults to "false").
	// Conflicts with `associatedRouters`.
	NoRouters *bool `pulumi:"noRouters"`
	// The policy resource id for the firewall. Changing
	// this updates the `policyId` of an existing firewall.
	PolicyId string `pulumi:"policyId"`
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall.
	Region *string `pulumi:"region"`
	// The owner of the floating IP. Required if admin wants
	// to create a firewall for another tenant. Changing this creates a new
	// firewall.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	// Administrative up/down status for the firewall
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall.
	AdminStateUp pulumi.BoolPtrInput
	// Router(s) to associate this firewall instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall. Conflicts with `noRouters`.
	AssociatedRouters pulumi.StringArrayInput
	// A description for the firewall. Changing this
	// updates the `description` of an existing firewall.
	Description pulumi.StringPtrInput
	// A name for the firewall. Changing this
	// updates the `name` of an existing firewall.
	Name pulumi.StringPtrInput
	// Should this firewall not be associated with any routers
	// (must be "true" or "false" if provide - defaults to "false").
	// Conflicts with `associatedRouters`.
	NoRouters pulumi.BoolPtrInput
	// The policy resource id for the firewall. Changing
	// this updates the `policyId` of an existing firewall.
	PolicyId pulumi.StringInput
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall.
	Region pulumi.StringPtrInput
	// The owner of the floating IP. Required if admin wants
	// to create a firewall for another tenant. Changing this creates a new
	// firewall.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (FirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallArgs)(nil)).Elem()
}

type FirewallInput interface {
	pulumi.Input

	ToFirewallOutput() FirewallOutput
	ToFirewallOutputWithContext(ctx context.Context) FirewallOutput
}

func (*Firewall) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (i *Firewall) ToFirewallOutput() FirewallOutput {
	return i.ToFirewallOutputWithContext(context.Background())
}

func (i *Firewall) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOutput)
}

// FirewallArrayInput is an input type that accepts FirewallArray and FirewallArrayOutput values.
// You can construct a concrete instance of `FirewallArrayInput` via:
//
//	FirewallArray{ FirewallArgs{...} }
type FirewallArrayInput interface {
	pulumi.Input

	ToFirewallArrayOutput() FirewallArrayOutput
	ToFirewallArrayOutputWithContext(context.Context) FirewallArrayOutput
}

type FirewallArray []FirewallInput

func (FirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (i FirewallArray) ToFirewallArrayOutput() FirewallArrayOutput {
	return i.ToFirewallArrayOutputWithContext(context.Background())
}

func (i FirewallArray) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallArrayOutput)
}

// FirewallMapInput is an input type that accepts FirewallMap and FirewallMapOutput values.
// You can construct a concrete instance of `FirewallMapInput` via:
//
//	FirewallMap{ "key": FirewallArgs{...} }
type FirewallMapInput interface {
	pulumi.Input

	ToFirewallMapOutput() FirewallMapOutput
	ToFirewallMapOutputWithContext(context.Context) FirewallMapOutput
}

type FirewallMap map[string]FirewallInput

func (FirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (i FirewallMap) ToFirewallMapOutput() FirewallMapOutput {
	return i.ToFirewallMapOutputWithContext(context.Background())
}

func (i FirewallMap) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMapOutput)
}

type FirewallOutput struct{ *pulumi.OutputState }

func (FirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (o FirewallOutput) ToFirewallOutput() FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return o
}

// Administrative up/down status for the firewall
// (must be "true" or "false" if provided - defaults to "true").
// Changing this updates the `adminStateUp` of an existing firewall.
func (o FirewallOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// Router(s) to associate this firewall instance
// with. Must be a list of strings. Changing this updates the associated routers
// of an existing firewall. Conflicts with `noRouters`.
func (o FirewallOutput) AssociatedRouters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringArrayOutput { return v.AssociatedRouters }).(pulumi.StringArrayOutput)
}

// A description for the firewall. Changing this
// updates the `description` of an existing firewall.
func (o FirewallOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A name for the firewall. Changing this
// updates the `name` of an existing firewall.
func (o FirewallOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Should this firewall not be associated with any routers
// (must be "true" or "false" if provide - defaults to "false").
// Conflicts with `associatedRouters`.
func (o FirewallOutput) NoRouters() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.BoolPtrOutput { return v.NoRouters }).(pulumi.BoolPtrOutput)
}

// The policy resource id for the firewall. Changing
// this updates the `policyId` of an existing firewall.
func (o FirewallOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The region in which to obtain the v1 networking client.
// A networking client is needed to create a firewall. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// firewall.
func (o FirewallOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The owner of the floating IP. Required if admin wants
// to create a firewall for another tenant. Changing this creates a new
// firewall.
func (o FirewallOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Map of additional options.
func (o FirewallOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *Firewall) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type FirewallArrayOutput struct{ *pulumi.OutputState }

func (FirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (o FirewallArrayOutput) ToFirewallArrayOutput() FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) Index(i pulumi.IntInput) FirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Firewall {
		return vs[0].([]*Firewall)[vs[1].(int)]
	}).(FirewallOutput)
}

type FirewallMapOutput struct{ *pulumi.OutputState }

func (FirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (o FirewallMapOutput) ToFirewallMapOutput() FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) MapIndex(k pulumi.StringInput) FirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Firewall {
		return vs[0].(map[string]*Firewall)[vs[1].(string)]
	}).(FirewallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInput)(nil)).Elem(), &Firewall{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallArrayInput)(nil)).Elem(), FirewallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMapInput)(nil)).Elem(), FirewallMap{})
	pulumi.RegisterOutputType(FirewallOutput{})
	pulumi.RegisterOutputType(FirewallArrayOutput{})
	pulumi.RegisterOutputType(FirewallMapOutput{})
}
