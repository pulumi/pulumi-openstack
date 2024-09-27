// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V3 group resource within OpenStack Keystone.
//
// > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
// this resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewGroupV3(ctx, "group_1", &identity.GroupV3Args{
//				Name:        pulumi.String("group_1"),
//				Description: pulumi.String("group 1"),
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
// groups can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:identity/groupV3:GroupV3 group_1 89c60255-9bd6-460c-822a-e2b959ede9d2
// ```
type GroupV3 struct {
	pulumi.CustomResourceState

	// A description of the group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain the group belongs to.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// The name of the group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new group.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewGroupV3 registers a new resource with the given unique name, arguments, and options.
func NewGroupV3(ctx *pulumi.Context,
	name string, args *GroupV3Args, opts ...pulumi.ResourceOption) (*GroupV3, error) {
	if args == nil {
		args = &GroupV3Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GroupV3
	err := ctx.RegisterResource("openstack:identity/groupV3:GroupV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupV3 gets an existing GroupV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupV3State, opts ...pulumi.ResourceOption) (*GroupV3, error) {
	var resource GroupV3
	err := ctx.ReadResource("openstack:identity/groupV3:GroupV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupV3 resources.
type groupV3State struct {
	// A description of the group.
	Description *string `pulumi:"description"`
	// The domain the group belongs to.
	DomainId *string `pulumi:"domainId"`
	// The name of the group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new group.
	Region *string `pulumi:"region"`
}

type GroupV3State struct {
	// A description of the group.
	Description pulumi.StringPtrInput
	// The domain the group belongs to.
	DomainId pulumi.StringPtrInput
	// The name of the group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new group.
	Region pulumi.StringPtrInput
}

func (GroupV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*groupV3State)(nil)).Elem()
}

type groupV3Args struct {
	// A description of the group.
	Description *string `pulumi:"description"`
	// The domain the group belongs to.
	DomainId *string `pulumi:"domainId"`
	// The name of the group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new group.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a GroupV3 resource.
type GroupV3Args struct {
	// A description of the group.
	Description pulumi.StringPtrInput
	// The domain the group belongs to.
	DomainId pulumi.StringPtrInput
	// The name of the group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new group.
	Region pulumi.StringPtrInput
}

func (GroupV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*groupV3Args)(nil)).Elem()
}

type GroupV3Input interface {
	pulumi.Input

	ToGroupV3Output() GroupV3Output
	ToGroupV3OutputWithContext(ctx context.Context) GroupV3Output
}

func (*GroupV3) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupV3)(nil)).Elem()
}

func (i *GroupV3) ToGroupV3Output() GroupV3Output {
	return i.ToGroupV3OutputWithContext(context.Background())
}

func (i *GroupV3) ToGroupV3OutputWithContext(ctx context.Context) GroupV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(GroupV3Output)
}

// GroupV3ArrayInput is an input type that accepts GroupV3Array and GroupV3ArrayOutput values.
// You can construct a concrete instance of `GroupV3ArrayInput` via:
//
//	GroupV3Array{ GroupV3Args{...} }
type GroupV3ArrayInput interface {
	pulumi.Input

	ToGroupV3ArrayOutput() GroupV3ArrayOutput
	ToGroupV3ArrayOutputWithContext(context.Context) GroupV3ArrayOutput
}

type GroupV3Array []GroupV3Input

func (GroupV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupV3)(nil)).Elem()
}

func (i GroupV3Array) ToGroupV3ArrayOutput() GroupV3ArrayOutput {
	return i.ToGroupV3ArrayOutputWithContext(context.Background())
}

func (i GroupV3Array) ToGroupV3ArrayOutputWithContext(ctx context.Context) GroupV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupV3ArrayOutput)
}

// GroupV3MapInput is an input type that accepts GroupV3Map and GroupV3MapOutput values.
// You can construct a concrete instance of `GroupV3MapInput` via:
//
//	GroupV3Map{ "key": GroupV3Args{...} }
type GroupV3MapInput interface {
	pulumi.Input

	ToGroupV3MapOutput() GroupV3MapOutput
	ToGroupV3MapOutputWithContext(context.Context) GroupV3MapOutput
}

type GroupV3Map map[string]GroupV3Input

func (GroupV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupV3)(nil)).Elem()
}

func (i GroupV3Map) ToGroupV3MapOutput() GroupV3MapOutput {
	return i.ToGroupV3MapOutputWithContext(context.Background())
}

func (i GroupV3Map) ToGroupV3MapOutputWithContext(ctx context.Context) GroupV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupV3MapOutput)
}

type GroupV3Output struct{ *pulumi.OutputState }

func (GroupV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupV3)(nil)).Elem()
}

func (o GroupV3Output) ToGroupV3Output() GroupV3Output {
	return o
}

func (o GroupV3Output) ToGroupV3OutputWithContext(ctx context.Context) GroupV3Output {
	return o
}

// A description of the group.
func (o GroupV3Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupV3) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain the group belongs to.
func (o GroupV3Output) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupV3) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// The name of the group.
func (o GroupV3Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupV3) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new group.
func (o GroupV3Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupV3) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type GroupV3ArrayOutput struct{ *pulumi.OutputState }

func (GroupV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupV3)(nil)).Elem()
}

func (o GroupV3ArrayOutput) ToGroupV3ArrayOutput() GroupV3ArrayOutput {
	return o
}

func (o GroupV3ArrayOutput) ToGroupV3ArrayOutputWithContext(ctx context.Context) GroupV3ArrayOutput {
	return o
}

func (o GroupV3ArrayOutput) Index(i pulumi.IntInput) GroupV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupV3 {
		return vs[0].([]*GroupV3)[vs[1].(int)]
	}).(GroupV3Output)
}

type GroupV3MapOutput struct{ *pulumi.OutputState }

func (GroupV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupV3)(nil)).Elem()
}

func (o GroupV3MapOutput) ToGroupV3MapOutput() GroupV3MapOutput {
	return o
}

func (o GroupV3MapOutput) ToGroupV3MapOutputWithContext(ctx context.Context) GroupV3MapOutput {
	return o
}

func (o GroupV3MapOutput) MapIndex(k pulumi.StringInput) GroupV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupV3 {
		return vs[0].(map[string]*GroupV3)[vs[1].(string)]
	}).(GroupV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupV3Input)(nil)).Elem(), &GroupV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupV3ArrayInput)(nil)).Elem(), GroupV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupV3MapInput)(nil)).Elem(), GroupV3Map{})
	pulumi.RegisterOutputType(GroupV3Output{})
	pulumi.RegisterOutputType(GroupV3ArrayOutput{})
	pulumi.RegisterOutputType(GroupV3MapOutput{})
}
