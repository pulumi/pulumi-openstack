// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V3 Role resource within OpenStack Keystone.
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
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/identity"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := identity.NewRole(ctx, "role1", nil)
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
// Roles can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:identity/role:Role role_1 89c60255-9bd6-460c-822a-e2b959ede9d2
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
	return reflect.TypeOf((*Role)(nil))
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

type RoleOutput struct {
	*pulumi.OutputState
}

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Role)(nil))
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoleOutput{})
}
