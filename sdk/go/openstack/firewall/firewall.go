// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package firewall

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a v1 firewall resource within OpenStack.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/fw_firewall_v1.html.markdown.
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
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil {
		args = &FirewallArgs{}
	}
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

