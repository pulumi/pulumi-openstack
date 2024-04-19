// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a v1 firewall policy resource within OpenStack.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err = firewall.NewPolicy(ctx, "policy_1", &firewall.PolicyArgs{
//				Name: pulumi.String("my-policy"),
//				Rules: pulumi.StringArray{
//					rule1.ID(),
//					rule2.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Firewall Policies can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:firewall/policy:Policy policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Audit status of the firewall policy
	// (must be "true" or "false" if provided - defaults to "false").
	// This status is set to "false" whenever the firewall policy or any of its
	// rules are changed. Changing this updates the `audited` status of an existing
	// firewall policy.
	Audited pulumi.BoolPtrOutput `pulumi:"audited"`
	// A description for the firewall policy. Changing
	// this updates the `description` of an existing firewall policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A name for the firewall policy. Changing this
	// updates the `name` of an existing firewall policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall policy.
	Region pulumi.StringOutput `pulumi:"region"`
	// An array of one or more firewall rules that comprise
	// the policy. Changing this results in adding/removing rules from the
	// existing firewall policy.
	Rules pulumi.StringArrayOutput `pulumi:"rules"`
	// Sharing status of the firewall policy (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. Only administrative users
	// can specify if the policy should be shared.
	Shared   pulumi.BoolPtrOutput `pulumi:"shared"`
	TenantId pulumi.StringOutput  `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("openstack:firewall/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("openstack:firewall/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Audit status of the firewall policy
	// (must be "true" or "false" if provided - defaults to "false").
	// This status is set to "false" whenever the firewall policy or any of its
	// rules are changed. Changing this updates the `audited` status of an existing
	// firewall policy.
	Audited *bool `pulumi:"audited"`
	// A description for the firewall policy. Changing
	// this updates the `description` of an existing firewall policy.
	Description *string `pulumi:"description"`
	// A name for the firewall policy. Changing this
	// updates the `name` of an existing firewall policy.
	Name *string `pulumi:"name"`
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall policy.
	Region *string `pulumi:"region"`
	// An array of one or more firewall rules that comprise
	// the policy. Changing this results in adding/removing rules from the
	// existing firewall policy.
	Rules []string `pulumi:"rules"`
	// Sharing status of the firewall policy (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. Only administrative users
	// can specify if the policy should be shared.
	Shared   *bool   `pulumi:"shared"`
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type PolicyState struct {
	// Audit status of the firewall policy
	// (must be "true" or "false" if provided - defaults to "false").
	// This status is set to "false" whenever the firewall policy or any of its
	// rules are changed. Changing this updates the `audited` status of an existing
	// firewall policy.
	Audited pulumi.BoolPtrInput
	// A description for the firewall policy. Changing
	// this updates the `description` of an existing firewall policy.
	Description pulumi.StringPtrInput
	// A name for the firewall policy. Changing this
	// updates the `name` of an existing firewall policy.
	Name pulumi.StringPtrInput
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall policy.
	Region pulumi.StringPtrInput
	// An array of one or more firewall rules that comprise
	// the policy. Changing this results in adding/removing rules from the
	// existing firewall policy.
	Rules pulumi.StringArrayInput
	// Sharing status of the firewall policy (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. Only administrative users
	// can specify if the policy should be shared.
	Shared   pulumi.BoolPtrInput
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Audit status of the firewall policy
	// (must be "true" or "false" if provided - defaults to "false").
	// This status is set to "false" whenever the firewall policy or any of its
	// rules are changed. Changing this updates the `audited` status of an existing
	// firewall policy.
	Audited *bool `pulumi:"audited"`
	// A description for the firewall policy. Changing
	// this updates the `description` of an existing firewall policy.
	Description *string `pulumi:"description"`
	// A name for the firewall policy. Changing this
	// updates the `name` of an existing firewall policy.
	Name *string `pulumi:"name"`
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall policy.
	Region *string `pulumi:"region"`
	// An array of one or more firewall rules that comprise
	// the policy. Changing this results in adding/removing rules from the
	// existing firewall policy.
	Rules []string `pulumi:"rules"`
	// Sharing status of the firewall policy (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. Only administrative users
	// can specify if the policy should be shared.
	Shared   *bool   `pulumi:"shared"`
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Audit status of the firewall policy
	// (must be "true" or "false" if provided - defaults to "false").
	// This status is set to "false" whenever the firewall policy or any of its
	// rules are changed. Changing this updates the `audited` status of an existing
	// firewall policy.
	Audited pulumi.BoolPtrInput
	// A description for the firewall policy. Changing
	// this updates the `description` of an existing firewall policy.
	Description pulumi.StringPtrInput
	// A name for the firewall policy. Changing this
	// updates the `name` of an existing firewall policy.
	Name pulumi.StringPtrInput
	// The region in which to obtain the v1 networking client.
	// A networking client is needed to create a firewall policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall policy.
	Region pulumi.StringPtrInput
	// An array of one or more firewall rules that comprise
	// the policy. Changing this results in adding/removing rules from the
	// existing firewall policy.
	Rules pulumi.StringArrayInput
	// Sharing status of the firewall policy (must be "true"
	// or "false" if provided). If this is "true" the policy is visible to, and
	// can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall policy. Only administrative users
	// can specify if the policy should be shared.
	Shared   pulumi.BoolPtrInput
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//	PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//	PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// Audit status of the firewall policy
// (must be "true" or "false" if provided - defaults to "false").
// This status is set to "false" whenever the firewall policy or any of its
// rules are changed. Changing this updates the `audited` status of an existing
// firewall policy.
func (o PolicyOutput) Audited() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolPtrOutput { return v.Audited }).(pulumi.BoolPtrOutput)
}

// A description for the firewall policy. Changing
// this updates the `description` of an existing firewall policy.
func (o PolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A name for the firewall policy. Changing this
// updates the `name` of an existing firewall policy.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to obtain the v1 networking client.
// A networking client is needed to create a firewall policy. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// firewall policy.
func (o PolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An array of one or more firewall rules that comprise
// the policy. Changing this results in adding/removing rules from the
// existing firewall policy.
func (o PolicyOutput) Rules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringArrayOutput { return v.Rules }).(pulumi.StringArrayOutput)
}

// Sharing status of the firewall policy (must be "true"
// or "false" if provided). If this is "true" the policy is visible to, and
// can be used in, firewalls in other tenants. Changing this updates the
// `shared` status of an existing firewall policy. Only administrative users
// can specify if the policy should be shared.
func (o PolicyOutput) Shared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolPtrOutput { return v.Shared }).(pulumi.BoolPtrOutput)
}

func (o PolicyOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Map of additional options.
func (o PolicyOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *Policy) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].([]*Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].(map[string]*Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
