// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V3 Role resource within OpenStack Keystone.
//
// > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
// this resource.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewRole(ctx, "role1", nil)
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
// Roles can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:identity/role:Role role_1 89c60255-9bd6-460c-822a-e2b959ede9d2
// ```
type Role struct {
	pulumi.CustomResourceState

	// The domain the role belongs to.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// The name of the role.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new Role.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		args = &RoleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Role
	err := ctx.RegisterResource("openstack:identity/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("openstack:identity/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// The domain the role belongs to.
	DomainId *string `pulumi:"domainId"`
	// The name of the role.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new Role.
	Region *string `pulumi:"region"`
}

type RoleState struct {
	// The domain the role belongs to.
	DomainId pulumi.StringPtrInput
	// The name of the role.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new Role.
	Region pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// The domain the role belongs to.
	DomainId *string `pulumi:"domainId"`
	// The name of the role.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new Role.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// The domain the role belongs to.
	DomainId pulumi.StringPtrInput
	// The name of the role.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new Role.
	Region pulumi.StringPtrInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//	RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//	RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

// The domain the role belongs to.
func (o RoleOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// The name of the role.
func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new Role.
func (o RoleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
