// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a user membership to group V3 resource within OpenStack.
//
// > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
// this resource.
//
// ***
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
//			project1, err := identity.NewProject(ctx, "project1", nil)
//			if err != nil {
//				return err
//			}
//			user1, err := identity.NewUser(ctx, "user1", &identity.UserArgs{
//				DefaultProjectId: project1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			group1, err := identity.NewGroupV3(ctx, "group1", &identity.GroupV3Args{
//				Description: pulumi.String("group 1"),
//			})
//			if err != nil {
//				return err
//			}
//			role1, err := identity.NewRole(ctx, "role1", nil)
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewUserMembershipV3(ctx, "userMembership1", &identity.UserMembershipV3Args{
//				UserId:  user1.ID(),
//				GroupId: group1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewRoleAssignment(ctx, "roleAssignment1", &identity.RoleAssignmentArgs{
//				GroupId:   group1.ID(),
//				ProjectId: project1.ID(),
//				RoleId:    role1.ID(),
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
// This resource can be imported by specifying all two arguments, separated
// by a forward slash:
//
// ```sh
// $ pulumi import openstack:identity/userMembershipV3:UserMembershipV3 user_membership_1 user_id/group_id
// ```
type UserMembershipV3 struct {
	pulumi.CustomResourceState

	// The UUID of group to which the user will be added.
	// Changing this creates a new user membership.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The region in which to obtain the V3 Identity client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new user membership.
	Region pulumi.StringOutput `pulumi:"region"`
	// The UUID of user to use. Changing this creates a new user membership.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUserMembershipV3 registers a new resource with the given unique name, arguments, and options.
func NewUserMembershipV3(ctx *pulumi.Context,
	name string, args *UserMembershipV3Args, opts ...pulumi.ResourceOption) (*UserMembershipV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserMembershipV3
	err := ctx.RegisterResource("openstack:identity/userMembershipV3:UserMembershipV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserMembershipV3 gets an existing UserMembershipV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserMembershipV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserMembershipV3State, opts ...pulumi.ResourceOption) (*UserMembershipV3, error) {
	var resource UserMembershipV3
	err := ctx.ReadResource("openstack:identity/userMembershipV3:UserMembershipV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserMembershipV3 resources.
type userMembershipV3State struct {
	// The UUID of group to which the user will be added.
	// Changing this creates a new user membership.
	GroupId *string `pulumi:"groupId"`
	// The region in which to obtain the V3 Identity client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new user membership.
	Region *string `pulumi:"region"`
	// The UUID of user to use. Changing this creates a new user membership.
	UserId *string `pulumi:"userId"`
}

type UserMembershipV3State struct {
	// The UUID of group to which the user will be added.
	// Changing this creates a new user membership.
	GroupId pulumi.StringPtrInput
	// The region in which to obtain the V3 Identity client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new user membership.
	Region pulumi.StringPtrInput
	// The UUID of user to use. Changing this creates a new user membership.
	UserId pulumi.StringPtrInput
}

func (UserMembershipV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*userMembershipV3State)(nil)).Elem()
}

type userMembershipV3Args struct {
	// The UUID of group to which the user will be added.
	// Changing this creates a new user membership.
	GroupId string `pulumi:"groupId"`
	// The region in which to obtain the V3 Identity client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new user membership.
	Region *string `pulumi:"region"`
	// The UUID of user to use. Changing this creates a new user membership.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a UserMembershipV3 resource.
type UserMembershipV3Args struct {
	// The UUID of group to which the user will be added.
	// Changing this creates a new user membership.
	GroupId pulumi.StringInput
	// The region in which to obtain the V3 Identity client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new user membership.
	Region pulumi.StringPtrInput
	// The UUID of user to use. Changing this creates a new user membership.
	UserId pulumi.StringInput
}

func (UserMembershipV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*userMembershipV3Args)(nil)).Elem()
}

type UserMembershipV3Input interface {
	pulumi.Input

	ToUserMembershipV3Output() UserMembershipV3Output
	ToUserMembershipV3OutputWithContext(ctx context.Context) UserMembershipV3Output
}

func (*UserMembershipV3) ElementType() reflect.Type {
	return reflect.TypeOf((**UserMembershipV3)(nil)).Elem()
}

func (i *UserMembershipV3) ToUserMembershipV3Output() UserMembershipV3Output {
	return i.ToUserMembershipV3OutputWithContext(context.Background())
}

func (i *UserMembershipV3) ToUserMembershipV3OutputWithContext(ctx context.Context) UserMembershipV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(UserMembershipV3Output)
}

// UserMembershipV3ArrayInput is an input type that accepts UserMembershipV3Array and UserMembershipV3ArrayOutput values.
// You can construct a concrete instance of `UserMembershipV3ArrayInput` via:
//
//	UserMembershipV3Array{ UserMembershipV3Args{...} }
type UserMembershipV3ArrayInput interface {
	pulumi.Input

	ToUserMembershipV3ArrayOutput() UserMembershipV3ArrayOutput
	ToUserMembershipV3ArrayOutputWithContext(context.Context) UserMembershipV3ArrayOutput
}

type UserMembershipV3Array []UserMembershipV3Input

func (UserMembershipV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserMembershipV3)(nil)).Elem()
}

func (i UserMembershipV3Array) ToUserMembershipV3ArrayOutput() UserMembershipV3ArrayOutput {
	return i.ToUserMembershipV3ArrayOutputWithContext(context.Background())
}

func (i UserMembershipV3Array) ToUserMembershipV3ArrayOutputWithContext(ctx context.Context) UserMembershipV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMembershipV3ArrayOutput)
}

// UserMembershipV3MapInput is an input type that accepts UserMembershipV3Map and UserMembershipV3MapOutput values.
// You can construct a concrete instance of `UserMembershipV3MapInput` via:
//
//	UserMembershipV3Map{ "key": UserMembershipV3Args{...} }
type UserMembershipV3MapInput interface {
	pulumi.Input

	ToUserMembershipV3MapOutput() UserMembershipV3MapOutput
	ToUserMembershipV3MapOutputWithContext(context.Context) UserMembershipV3MapOutput
}

type UserMembershipV3Map map[string]UserMembershipV3Input

func (UserMembershipV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserMembershipV3)(nil)).Elem()
}

func (i UserMembershipV3Map) ToUserMembershipV3MapOutput() UserMembershipV3MapOutput {
	return i.ToUserMembershipV3MapOutputWithContext(context.Background())
}

func (i UserMembershipV3Map) ToUserMembershipV3MapOutputWithContext(ctx context.Context) UserMembershipV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMembershipV3MapOutput)
}

type UserMembershipV3Output struct{ *pulumi.OutputState }

func (UserMembershipV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**UserMembershipV3)(nil)).Elem()
}

func (o UserMembershipV3Output) ToUserMembershipV3Output() UserMembershipV3Output {
	return o
}

func (o UserMembershipV3Output) ToUserMembershipV3OutputWithContext(ctx context.Context) UserMembershipV3Output {
	return o
}

// The UUID of group to which the user will be added.
// Changing this creates a new user membership.
func (o UserMembershipV3Output) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserMembershipV3) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The region in which to obtain the V3 Identity client.
// If omitted, the `region` argument of the provider is used.
// Changing this creates a new user membership.
func (o UserMembershipV3Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *UserMembershipV3) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The UUID of user to use. Changing this creates a new user membership.
func (o UserMembershipV3Output) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserMembershipV3) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type UserMembershipV3ArrayOutput struct{ *pulumi.OutputState }

func (UserMembershipV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserMembershipV3)(nil)).Elem()
}

func (o UserMembershipV3ArrayOutput) ToUserMembershipV3ArrayOutput() UserMembershipV3ArrayOutput {
	return o
}

func (o UserMembershipV3ArrayOutput) ToUserMembershipV3ArrayOutputWithContext(ctx context.Context) UserMembershipV3ArrayOutput {
	return o
}

func (o UserMembershipV3ArrayOutput) Index(i pulumi.IntInput) UserMembershipV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserMembershipV3 {
		return vs[0].([]*UserMembershipV3)[vs[1].(int)]
	}).(UserMembershipV3Output)
}

type UserMembershipV3MapOutput struct{ *pulumi.OutputState }

func (UserMembershipV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserMembershipV3)(nil)).Elem()
}

func (o UserMembershipV3MapOutput) ToUserMembershipV3MapOutput() UserMembershipV3MapOutput {
	return o
}

func (o UserMembershipV3MapOutput) ToUserMembershipV3MapOutputWithContext(ctx context.Context) UserMembershipV3MapOutput {
	return o
}

func (o UserMembershipV3MapOutput) MapIndex(k pulumi.StringInput) UserMembershipV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserMembershipV3 {
		return vs[0].(map[string]*UserMembershipV3)[vs[1].(string)]
	}).(UserMembershipV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserMembershipV3Input)(nil)).Elem(), &UserMembershipV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMembershipV3ArrayInput)(nil)).Elem(), UserMembershipV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMembershipV3MapInput)(nil)).Elem(), UserMembershipV3Map{})
	pulumi.RegisterOutputType(UserMembershipV3Output{})
	pulumi.RegisterOutputType(UserMembershipV3ArrayOutput{})
	pulumi.RegisterOutputType(UserMembershipV3MapOutput{})
}
