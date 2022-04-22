// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Security Groups can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:networking/secGroup:SecGroup secgroup_1 38809219-5e8a-4852-9139-6f461c90e8bc
// ```
type SecGroup struct {
	pulumi.CustomResourceState

	// The collection of tags assigned on the security group, which have
	// been explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// Whether or not to delete the default
	// egress security rules. This is `false` by default. See the below note
	// for more information.
	DeleteDefaultRules pulumi.BoolPtrOutput `pulumi:"deleteDefaultRules"`
	// A unique name for the security group.
	Description pulumi.StringOutput `pulumi:"description"`
	// A unique name for the security group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region pulumi.StringOutput `pulumi:"region"`
	// A set of string tags for the security group.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewSecGroup registers a new resource with the given unique name, arguments, and options.
func NewSecGroup(ctx *pulumi.Context,
	name string, args *SecGroupArgs, opts ...pulumi.ResourceOption) (*SecGroup, error) {
	if args == nil {
		args = &SecGroupArgs{}
	}

	var resource SecGroup
	err := ctx.RegisterResource("openstack:networking/secGroup:SecGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecGroup gets an existing SecGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecGroupState, opts ...pulumi.ResourceOption) (*SecGroup, error) {
	var resource SecGroup
	err := ctx.ReadResource("openstack:networking/secGroup:SecGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecGroup resources.
type secGroupState struct {
	// The collection of tags assigned on the security group, which have
	// been explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// Whether or not to delete the default
	// egress security rules. This is `false` by default. See the below note
	// for more information.
	DeleteDefaultRules *bool `pulumi:"deleteDefaultRules"`
	// A unique name for the security group.
	Description *string `pulumi:"description"`
	// A unique name for the security group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region *string `pulumi:"region"`
	// A set of string tags for the security group.
	Tags []string `pulumi:"tags"`
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group.
	TenantId *string `pulumi:"tenantId"`
}

type SecGroupState struct {
	// The collection of tags assigned on the security group, which have
	// been explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// Whether or not to delete the default
	// egress security rules. This is `false` by default. See the below note
	// for more information.
	DeleteDefaultRules pulumi.BoolPtrInput
	// A unique name for the security group.
	Description pulumi.StringPtrInput
	// A unique name for the security group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region pulumi.StringPtrInput
	// A set of string tags for the security group.
	Tags pulumi.StringArrayInput
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group.
	TenantId pulumi.StringPtrInput
}

func (SecGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*secGroupState)(nil)).Elem()
}

type secGroupArgs struct {
	// Whether or not to delete the default
	// egress security rules. This is `false` by default. See the below note
	// for more information.
	DeleteDefaultRules *bool `pulumi:"deleteDefaultRules"`
	// A unique name for the security group.
	Description *string `pulumi:"description"`
	// A unique name for the security group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region *string `pulumi:"region"`
	// A set of string tags for the security group.
	Tags []string `pulumi:"tags"`
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a SecGroup resource.
type SecGroupArgs struct {
	// Whether or not to delete the default
	// egress security rules. This is `false` by default. See the below note
	// for more information.
	DeleteDefaultRules pulumi.BoolPtrInput
	// A unique name for the security group.
	Description pulumi.StringPtrInput
	// A unique name for the security group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region pulumi.StringPtrInput
	// A set of string tags for the security group.
	Tags pulumi.StringArrayInput
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group.
	TenantId pulumi.StringPtrInput
}

func (SecGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secGroupArgs)(nil)).Elem()
}

type SecGroupInput interface {
	pulumi.Input

	ToSecGroupOutput() SecGroupOutput
	ToSecGroupOutputWithContext(ctx context.Context) SecGroupOutput
}

func (*SecGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SecGroup)(nil)).Elem()
}

func (i *SecGroup) ToSecGroupOutput() SecGroupOutput {
	return i.ToSecGroupOutputWithContext(context.Background())
}

func (i *SecGroup) ToSecGroupOutputWithContext(ctx context.Context) SecGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupOutput)
}

// SecGroupArrayInput is an input type that accepts SecGroupArray and SecGroupArrayOutput values.
// You can construct a concrete instance of `SecGroupArrayInput` via:
//
//          SecGroupArray{ SecGroupArgs{...} }
type SecGroupArrayInput interface {
	pulumi.Input

	ToSecGroupArrayOutput() SecGroupArrayOutput
	ToSecGroupArrayOutputWithContext(context.Context) SecGroupArrayOutput
}

type SecGroupArray []SecGroupInput

func (SecGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecGroup)(nil)).Elem()
}

func (i SecGroupArray) ToSecGroupArrayOutput() SecGroupArrayOutput {
	return i.ToSecGroupArrayOutputWithContext(context.Background())
}

func (i SecGroupArray) ToSecGroupArrayOutputWithContext(ctx context.Context) SecGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupArrayOutput)
}

// SecGroupMapInput is an input type that accepts SecGroupMap and SecGroupMapOutput values.
// You can construct a concrete instance of `SecGroupMapInput` via:
//
//          SecGroupMap{ "key": SecGroupArgs{...} }
type SecGroupMapInput interface {
	pulumi.Input

	ToSecGroupMapOutput() SecGroupMapOutput
	ToSecGroupMapOutputWithContext(context.Context) SecGroupMapOutput
}

type SecGroupMap map[string]SecGroupInput

func (SecGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecGroup)(nil)).Elem()
}

func (i SecGroupMap) ToSecGroupMapOutput() SecGroupMapOutput {
	return i.ToSecGroupMapOutputWithContext(context.Background())
}

func (i SecGroupMap) ToSecGroupMapOutputWithContext(ctx context.Context) SecGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupMapOutput)
}

type SecGroupOutput struct{ *pulumi.OutputState }

func (SecGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecGroup)(nil)).Elem()
}

func (o SecGroupOutput) ToSecGroupOutput() SecGroupOutput {
	return o
}

func (o SecGroupOutput) ToSecGroupOutputWithContext(ctx context.Context) SecGroupOutput {
	return o
}

type SecGroupArrayOutput struct{ *pulumi.OutputState }

func (SecGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecGroup)(nil)).Elem()
}

func (o SecGroupArrayOutput) ToSecGroupArrayOutput() SecGroupArrayOutput {
	return o
}

func (o SecGroupArrayOutput) ToSecGroupArrayOutputWithContext(ctx context.Context) SecGroupArrayOutput {
	return o
}

func (o SecGroupArrayOutput) Index(i pulumi.IntInput) SecGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecGroup {
		return vs[0].([]*SecGroup)[vs[1].(int)]
	}).(SecGroupOutput)
}

type SecGroupMapOutput struct{ *pulumi.OutputState }

func (SecGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecGroup)(nil)).Elem()
}

func (o SecGroupMapOutput) ToSecGroupMapOutput() SecGroupMapOutput {
	return o
}

func (o SecGroupMapOutput) ToSecGroupMapOutputWithContext(ctx context.Context) SecGroupMapOutput {
	return o
}

func (o SecGroupMapOutput) MapIndex(k pulumi.StringInput) SecGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecGroup {
		return vs[0].(map[string]*SecGroup)[vs[1].(string)]
	}).(SecGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupInput)(nil)).Elem(), &SecGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupArrayInput)(nil)).Elem(), SecGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupMapInput)(nil)).Elem(), SecGroupMap{})
	pulumi.RegisterOutputType(SecGroupOutput{})
	pulumi.RegisterOutputType(SecGroupArrayOutput{})
	pulumi.RegisterOutputType(SecGroupMapOutput{})
}
