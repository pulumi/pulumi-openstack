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

// Manages a v2 firewall group resource within OpenStack.
//
// > **Note:** Firewall v2 has no support for OVN currently.
//
// ## Import
//
// Firewall groups can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:firewall/groupV2:GroupV2 group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97
//
// ```
type GroupV2 struct {
	pulumi.CustomResourceState

	// Administrative up/down status for the firewall
	// group (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall group.
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// A description for the firewall group. Changing this
	// updates the `description` of an existing firewall group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The egress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `egressFirewallPolicyId` of an existing firewall group.
	EgressFirewallPolicyId pulumi.StringPtrOutput `pulumi:"egressFirewallPolicyId"`
	// The ingress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `ingressFirewallPolicyId` of an existing firewall group.
	IngressFirewallPolicyId pulumi.StringPtrOutput `pulumi:"ingressFirewallPolicyId"`
	// A name for the firewall group. Changing this
	// updates the `name` of an existing firewall.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port(s) to associate this firewall group
	// with. Must be a list of strings. Changing this updates the associated ports
	// of an existing firewall group.
	Ports pulumi.StringArrayOutput `pulumi:"ports"`
	// This argument conflict and interchangeable with
	// `tenantId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another project. Changing this creates a new
	// firewall group.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall group.
	Region pulumi.StringOutput `pulumi:"region"`
	// Sharing status of the firewall group (must be "true"
	// or "false" if provided). If this is "true" the firewall group is visible to,
	// and can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall group. Only administrative users
	// can specify if the firewall group should be shared.
	Shared pulumi.BoolPtrOutput `pulumi:"shared"`
	// The status of the firewall group.
	Status pulumi.StringOutput `pulumi:"status"`
	// This argument conflict and interchangeable with
	// `projectId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another tenant. Changing this creates a new
	// firewall group.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewGroupV2 registers a new resource with the given unique name, arguments, and options.
func NewGroupV2(ctx *pulumi.Context,
	name string, args *GroupV2Args, opts ...pulumi.ResourceOption) (*GroupV2, error) {
	if args == nil {
		args = &GroupV2Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupV2
	err := ctx.RegisterResource("openstack:firewall/groupV2:GroupV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupV2 gets an existing GroupV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupV2State, opts ...pulumi.ResourceOption) (*GroupV2, error) {
	var resource GroupV2
	err := ctx.ReadResource("openstack:firewall/groupV2:GroupV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupV2 resources.
type groupV2State struct {
	// Administrative up/down status for the firewall
	// group (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall group.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// A description for the firewall group. Changing this
	// updates the `description` of an existing firewall group.
	Description *string `pulumi:"description"`
	// The egress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `egressFirewallPolicyId` of an existing firewall group.
	EgressFirewallPolicyId *string `pulumi:"egressFirewallPolicyId"`
	// The ingress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `ingressFirewallPolicyId` of an existing firewall group.
	IngressFirewallPolicyId *string `pulumi:"ingressFirewallPolicyId"`
	// A name for the firewall group. Changing this
	// updates the `name` of an existing firewall.
	Name *string `pulumi:"name"`
	// Port(s) to associate this firewall group
	// with. Must be a list of strings. Changing this updates the associated ports
	// of an existing firewall group.
	Ports []string `pulumi:"ports"`
	// This argument conflict and interchangeable with
	// `tenantId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another project. Changing this creates a new
	// firewall group.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall group.
	Region *string `pulumi:"region"`
	// Sharing status of the firewall group (must be "true"
	// or "false" if provided). If this is "true" the firewall group is visible to,
	// and can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall group. Only administrative users
	// can specify if the firewall group should be shared.
	Shared *bool `pulumi:"shared"`
	// The status of the firewall group.
	Status *string `pulumi:"status"`
	// This argument conflict and interchangeable with
	// `projectId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another tenant. Changing this creates a new
	// firewall group.
	TenantId *string `pulumi:"tenantId"`
}

type GroupV2State struct {
	// Administrative up/down status for the firewall
	// group (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall group.
	AdminStateUp pulumi.BoolPtrInput
	// A description for the firewall group. Changing this
	// updates the `description` of an existing firewall group.
	Description pulumi.StringPtrInput
	// The egress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `egressFirewallPolicyId` of an existing firewall group.
	EgressFirewallPolicyId pulumi.StringPtrInput
	// The ingress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `ingressFirewallPolicyId` of an existing firewall group.
	IngressFirewallPolicyId pulumi.StringPtrInput
	// A name for the firewall group. Changing this
	// updates the `name` of an existing firewall.
	Name pulumi.StringPtrInput
	// Port(s) to associate this firewall group
	// with. Must be a list of strings. Changing this updates the associated ports
	// of an existing firewall group.
	Ports pulumi.StringArrayInput
	// This argument conflict and interchangeable with
	// `tenantId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another project. Changing this creates a new
	// firewall group.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall group.
	Region pulumi.StringPtrInput
	// Sharing status of the firewall group (must be "true"
	// or "false" if provided). If this is "true" the firewall group is visible to,
	// and can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall group. Only administrative users
	// can specify if the firewall group should be shared.
	Shared pulumi.BoolPtrInput
	// The status of the firewall group.
	Status pulumi.StringPtrInput
	// This argument conflict and interchangeable with
	// `projectId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another tenant. Changing this creates a new
	// firewall group.
	TenantId pulumi.StringPtrInput
}

func (GroupV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*groupV2State)(nil)).Elem()
}

type groupV2Args struct {
	// Administrative up/down status for the firewall
	// group (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall group.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// A description for the firewall group. Changing this
	// updates the `description` of an existing firewall group.
	Description *string `pulumi:"description"`
	// The egress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `egressFirewallPolicyId` of an existing firewall group.
	EgressFirewallPolicyId *string `pulumi:"egressFirewallPolicyId"`
	// The ingress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `ingressFirewallPolicyId` of an existing firewall group.
	IngressFirewallPolicyId *string `pulumi:"ingressFirewallPolicyId"`
	// A name for the firewall group. Changing this
	// updates the `name` of an existing firewall.
	Name *string `pulumi:"name"`
	// Port(s) to associate this firewall group
	// with. Must be a list of strings. Changing this updates the associated ports
	// of an existing firewall group.
	Ports []string `pulumi:"ports"`
	// This argument conflict and interchangeable with
	// `tenantId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another project. Changing this creates a new
	// firewall group.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall group.
	Region *string `pulumi:"region"`
	// Sharing status of the firewall group (must be "true"
	// or "false" if provided). If this is "true" the firewall group is visible to,
	// and can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall group. Only administrative users
	// can specify if the firewall group should be shared.
	Shared *bool `pulumi:"shared"`
	// This argument conflict and interchangeable with
	// `projectId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another tenant. Changing this creates a new
	// firewall group.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a GroupV2 resource.
type GroupV2Args struct {
	// Administrative up/down status for the firewall
	// group (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the `adminStateUp` of an existing firewall group.
	AdminStateUp pulumi.BoolPtrInput
	// A description for the firewall group. Changing this
	// updates the `description` of an existing firewall group.
	Description pulumi.StringPtrInput
	// The egress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `egressFirewallPolicyId` of an existing firewall group.
	EgressFirewallPolicyId pulumi.StringPtrInput
	// The ingress firewall policy resource
	// id for the firewall group. Changing this updates the
	// `ingressFirewallPolicyId` of an existing firewall group.
	IngressFirewallPolicyId pulumi.StringPtrInput
	// A name for the firewall group. Changing this
	// updates the `name` of an existing firewall.
	Name pulumi.StringPtrInput
	// Port(s) to associate this firewall group
	// with. Must be a list of strings. Changing this updates the associated ports
	// of an existing firewall group.
	Ports pulumi.StringArrayInput
	// This argument conflict and interchangeable with
	// `tenantId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another project. Changing this creates a new
	// firewall group.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// firewall group.
	Region pulumi.StringPtrInput
	// Sharing status of the firewall group (must be "true"
	// or "false" if provided). If this is "true" the firewall group is visible to,
	// and can be used in, firewalls in other tenants. Changing this updates the
	// `shared` status of an existing firewall group. Only administrative users
	// can specify if the firewall group should be shared.
	Shared pulumi.BoolPtrInput
	// This argument conflict and interchangeable with
	// `projectId`. The owner of the firewall group. Required if admin wants to
	// create a firewall group for another tenant. Changing this creates a new
	// firewall group.
	TenantId pulumi.StringPtrInput
}

func (GroupV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*groupV2Args)(nil)).Elem()
}

type GroupV2Input interface {
	pulumi.Input

	ToGroupV2Output() GroupV2Output
	ToGroupV2OutputWithContext(ctx context.Context) GroupV2Output
}

func (*GroupV2) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupV2)(nil)).Elem()
}

func (i *GroupV2) ToGroupV2Output() GroupV2Output {
	return i.ToGroupV2OutputWithContext(context.Background())
}

func (i *GroupV2) ToGroupV2OutputWithContext(ctx context.Context) GroupV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(GroupV2Output)
}

func (i *GroupV2) ToOutput(ctx context.Context) pulumix.Output[*GroupV2] {
	return pulumix.Output[*GroupV2]{
		OutputState: i.ToGroupV2OutputWithContext(ctx).OutputState,
	}
}

// GroupV2ArrayInput is an input type that accepts GroupV2Array and GroupV2ArrayOutput values.
// You can construct a concrete instance of `GroupV2ArrayInput` via:
//
//	GroupV2Array{ GroupV2Args{...} }
type GroupV2ArrayInput interface {
	pulumi.Input

	ToGroupV2ArrayOutput() GroupV2ArrayOutput
	ToGroupV2ArrayOutputWithContext(context.Context) GroupV2ArrayOutput
}

type GroupV2Array []GroupV2Input

func (GroupV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupV2)(nil)).Elem()
}

func (i GroupV2Array) ToGroupV2ArrayOutput() GroupV2ArrayOutput {
	return i.ToGroupV2ArrayOutputWithContext(context.Background())
}

func (i GroupV2Array) ToGroupV2ArrayOutputWithContext(ctx context.Context) GroupV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupV2ArrayOutput)
}

func (i GroupV2Array) ToOutput(ctx context.Context) pulumix.Output[[]*GroupV2] {
	return pulumix.Output[[]*GroupV2]{
		OutputState: i.ToGroupV2ArrayOutputWithContext(ctx).OutputState,
	}
}

// GroupV2MapInput is an input type that accepts GroupV2Map and GroupV2MapOutput values.
// You can construct a concrete instance of `GroupV2MapInput` via:
//
//	GroupV2Map{ "key": GroupV2Args{...} }
type GroupV2MapInput interface {
	pulumi.Input

	ToGroupV2MapOutput() GroupV2MapOutput
	ToGroupV2MapOutputWithContext(context.Context) GroupV2MapOutput
}

type GroupV2Map map[string]GroupV2Input

func (GroupV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupV2)(nil)).Elem()
}

func (i GroupV2Map) ToGroupV2MapOutput() GroupV2MapOutput {
	return i.ToGroupV2MapOutputWithContext(context.Background())
}

func (i GroupV2Map) ToGroupV2MapOutputWithContext(ctx context.Context) GroupV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupV2MapOutput)
}

func (i GroupV2Map) ToOutput(ctx context.Context) pulumix.Output[map[string]*GroupV2] {
	return pulumix.Output[map[string]*GroupV2]{
		OutputState: i.ToGroupV2MapOutputWithContext(ctx).OutputState,
	}
}

type GroupV2Output struct{ *pulumi.OutputState }

func (GroupV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupV2)(nil)).Elem()
}

func (o GroupV2Output) ToGroupV2Output() GroupV2Output {
	return o
}

func (o GroupV2Output) ToGroupV2OutputWithContext(ctx context.Context) GroupV2Output {
	return o
}

func (o GroupV2Output) ToOutput(ctx context.Context) pulumix.Output[*GroupV2] {
	return pulumix.Output[*GroupV2]{
		OutputState: o.OutputState,
	}
}

// Administrative up/down status for the firewall
// group (must be "true" or "false" if provided - defaults to "true").
// Changing this updates the `adminStateUp` of an existing firewall group.
func (o GroupV2Output) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// A description for the firewall group. Changing this
// updates the `description` of an existing firewall group.
func (o GroupV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The egress firewall policy resource
// id for the firewall group. Changing this updates the
// `egressFirewallPolicyId` of an existing firewall group.
func (o GroupV2Output) EgressFirewallPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringPtrOutput { return v.EgressFirewallPolicyId }).(pulumi.StringPtrOutput)
}

// The ingress firewall policy resource
// id for the firewall group. Changing this updates the
// `ingressFirewallPolicyId` of an existing firewall group.
func (o GroupV2Output) IngressFirewallPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringPtrOutput { return v.IngressFirewallPolicyId }).(pulumi.StringPtrOutput)
}

// A name for the firewall group. Changing this
// updates the `name` of an existing firewall.
func (o GroupV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Port(s) to associate this firewall group
// with. Must be a list of strings. Changing this updates the associated ports
// of an existing firewall group.
func (o GroupV2Output) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringArrayOutput { return v.Ports }).(pulumi.StringArrayOutput)
}

// This argument conflict and interchangeable with
// `tenantId`. The owner of the firewall group. Required if admin wants to
// create a firewall group for another project. Changing this creates a new
// firewall group.
func (o GroupV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the v2 networking client.
// A networking client is needed to create a firewall group. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// firewall group.
func (o GroupV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Sharing status of the firewall group (must be "true"
// or "false" if provided). If this is "true" the firewall group is visible to,
// and can be used in, firewalls in other tenants. Changing this updates the
// `shared` status of an existing firewall group. Only administrative users
// can specify if the firewall group should be shared.
func (o GroupV2Output) Shared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.BoolPtrOutput { return v.Shared }).(pulumi.BoolPtrOutput)
}

// The status of the firewall group.
func (o GroupV2Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// This argument conflict and interchangeable with
// `projectId`. The owner of the firewall group. Required if admin wants to
// create a firewall group for another tenant. Changing this creates a new
// firewall group.
func (o GroupV2Output) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupV2) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type GroupV2ArrayOutput struct{ *pulumi.OutputState }

func (GroupV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupV2)(nil)).Elem()
}

func (o GroupV2ArrayOutput) ToGroupV2ArrayOutput() GroupV2ArrayOutput {
	return o
}

func (o GroupV2ArrayOutput) ToGroupV2ArrayOutputWithContext(ctx context.Context) GroupV2ArrayOutput {
	return o
}

func (o GroupV2ArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GroupV2] {
	return pulumix.Output[[]*GroupV2]{
		OutputState: o.OutputState,
	}
}

func (o GroupV2ArrayOutput) Index(i pulumi.IntInput) GroupV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupV2 {
		return vs[0].([]*GroupV2)[vs[1].(int)]
	}).(GroupV2Output)
}

type GroupV2MapOutput struct{ *pulumi.OutputState }

func (GroupV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupV2)(nil)).Elem()
}

func (o GroupV2MapOutput) ToGroupV2MapOutput() GroupV2MapOutput {
	return o
}

func (o GroupV2MapOutput) ToGroupV2MapOutputWithContext(ctx context.Context) GroupV2MapOutput {
	return o
}

func (o GroupV2MapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GroupV2] {
	return pulumix.Output[map[string]*GroupV2]{
		OutputState: o.OutputState,
	}
}

func (o GroupV2MapOutput) MapIndex(k pulumi.StringInput) GroupV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupV2 {
		return vs[0].(map[string]*GroupV2)[vs[1].(string)]
	}).(GroupV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupV2Input)(nil)).Elem(), &GroupV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupV2ArrayInput)(nil)).Elem(), GroupV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupV2MapInput)(nil)).Elem(), GroupV2Map{})
	pulumi.RegisterOutputType(GroupV2Output{})
	pulumi.RegisterOutputType(GroupV2ArrayOutput{})
	pulumi.RegisterOutputType(GroupV2MapOutput{})
}
