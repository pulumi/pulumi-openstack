// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a v1 firewall policy resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/firewall"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		rule1, err := firewall.NewRule(ctx, "rule1", &firewall.RuleArgs{
// 			Action:          pulumi.String("deny"),
// 			Description:     pulumi.String("drop TELNET traffic"),
// 			DestinationPort: pulumi.String("23"),
// 			Enabled:         pulumi.Bool(true),
// 			Protocol:        pulumi.String("tcp"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		rule2, err := firewall.NewRule(ctx, "rule2", &firewall.RuleArgs{
// 			Action:          pulumi.String("deny"),
// 			Description:     pulumi.String("drop NTP traffic"),
// 			DestinationPort: pulumi.String("123"),
// 			Enabled:         pulumi.Bool(false),
// 			Protocol:        pulumi.String("udp"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = firewall.NewPolicy(ctx, "policy1", &firewall.PolicyArgs{
// 			Rules: pulumi.StringArray{
// 				rule1.ID(),
// 				rule2.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Firewall Policies can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:firewall/policy:Policy policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0
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
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct {
	*pulumi.OutputState
}

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
}
