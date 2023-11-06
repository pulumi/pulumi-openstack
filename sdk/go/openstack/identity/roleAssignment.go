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

// Manages a V3 Role assignment within OpenStack Keystone.
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
//			role1, err := identity.NewRole(ctx, "role1", nil)
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewRoleAssignment(ctx, "roleAssignment1", &identity.RoleAssignmentArgs{
//				UserId:    user1.ID(),
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
//
// ## Import
//
// Role assignments can be imported using a constructed id. The id should have the form of `domainID/projectID/groupID/userID/roleID`. When something is not used then leave blank.
//
// For example this will import the role assignment forprojectID014395cd-89fc-4c9b-96b7-13d1ee79dad2, userID4142e64b-1b35-44a0-9b1e-5affc7af1106, roleIDea257959-eeb1-4c10-8d33-26f0409a755d ( domainID and groupID are left blank)
//
// ```sh
//
//	$ pulumi import openstack:identity/roleAssignment:RoleAssignment role_assignment_1 /014395cd-89fc-4c9b-96b7-13d1ee79dad2//4142e64b-1b35-44a0-9b1e-5affc7af1106/ea257959-eeb1-4c10-8d33-26f0409a755d
//
// ```
type RoleAssignment struct {
	pulumi.CustomResourceState

	// The domain to assign the role in.
	DomainId pulumi.StringPtrOutput `pulumi:"domainId"`
	// The group to assign the role to.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// The project to assign the role in.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	Region    pulumi.StringOutput    `pulumi:"region"`
	// The role to assign.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// The user to assign the role to.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RoleAssignment
	err := ctx.RegisterResource("openstack:identity/roleAssignment:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAssignment gets an existing RoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("openstack:identity/roleAssignment:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssignment resources.
type roleAssignmentState struct {
	// The domain to assign the role in.
	DomainId *string `pulumi:"domainId"`
	// The group to assign the role to.
	GroupId *string `pulumi:"groupId"`
	// The project to assign the role in.
	ProjectId *string `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// The role to assign.
	RoleId *string `pulumi:"roleId"`
	// The user to assign the role to.
	UserId *string `pulumi:"userId"`
}

type RoleAssignmentState struct {
	// The domain to assign the role in.
	DomainId pulumi.StringPtrInput
	// The group to assign the role to.
	GroupId pulumi.StringPtrInput
	// The project to assign the role in.
	ProjectId pulumi.StringPtrInput
	Region    pulumi.StringPtrInput
	// The role to assign.
	RoleId pulumi.StringPtrInput
	// The user to assign the role to.
	UserId pulumi.StringPtrInput
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	// The domain to assign the role in.
	DomainId *string `pulumi:"domainId"`
	// The group to assign the role to.
	GroupId *string `pulumi:"groupId"`
	// The project to assign the role in.
	ProjectId *string `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// The role to assign.
	RoleId string `pulumi:"roleId"`
	// The user to assign the role to.
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// The domain to assign the role in.
	DomainId pulumi.StringPtrInput
	// The group to assign the role to.
	GroupId pulumi.StringPtrInput
	// The project to assign the role in.
	ProjectId pulumi.StringPtrInput
	Region    pulumi.StringPtrInput
	// The role to assign.
	RoleId pulumi.StringInput
	// The user to assign the role to.
	UserId pulumi.StringPtrInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}

type RoleAssignmentInput interface {
	pulumi.Input

	ToRoleAssignmentOutput() RoleAssignmentOutput
	ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput
}

func (*RoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (i *RoleAssignment) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return i.ToRoleAssignmentOutputWithContext(context.Background())
}

func (i *RoleAssignment) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentOutput)
}

// RoleAssignmentArrayInput is an input type that accepts RoleAssignmentArray and RoleAssignmentArrayOutput values.
// You can construct a concrete instance of `RoleAssignmentArrayInput` via:
//
//	RoleAssignmentArray{ RoleAssignmentArgs{...} }
type RoleAssignmentArrayInput interface {
	pulumi.Input

	ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput
	ToRoleAssignmentArrayOutputWithContext(context.Context) RoleAssignmentArrayOutput
}

type RoleAssignmentArray []RoleAssignmentInput

func (RoleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAssignment)(nil)).Elem()
}

func (i RoleAssignmentArray) ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput {
	return i.ToRoleAssignmentArrayOutputWithContext(context.Background())
}

func (i RoleAssignmentArray) ToRoleAssignmentArrayOutputWithContext(ctx context.Context) RoleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentArrayOutput)
}

// RoleAssignmentMapInput is an input type that accepts RoleAssignmentMap and RoleAssignmentMapOutput values.
// You can construct a concrete instance of `RoleAssignmentMapInput` via:
//
//	RoleAssignmentMap{ "key": RoleAssignmentArgs{...} }
type RoleAssignmentMapInput interface {
	pulumi.Input

	ToRoleAssignmentMapOutput() RoleAssignmentMapOutput
	ToRoleAssignmentMapOutputWithContext(context.Context) RoleAssignmentMapOutput
}

type RoleAssignmentMap map[string]RoleAssignmentInput

func (RoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAssignment)(nil)).Elem()
}

func (i RoleAssignmentMap) ToRoleAssignmentMapOutput() RoleAssignmentMapOutput {
	return i.ToRoleAssignmentMapOutputWithContext(context.Background())
}

func (i RoleAssignmentMap) ToRoleAssignmentMapOutputWithContext(ctx context.Context) RoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentMapOutput)
}

type RoleAssignmentOutput struct{ *pulumi.OutputState }

func (RoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return o
}

// The domain to assign the role in.
func (o RoleAssignmentOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.DomainId }).(pulumi.StringPtrOutput)
}

// The group to assign the role to.
func (o RoleAssignmentOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The project to assign the role in.
func (o RoleAssignmentOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role to assign.
func (o RoleAssignmentOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// The user to assign the role to.
func (o RoleAssignmentOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

type RoleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (RoleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentArrayOutput) ToRoleAssignmentArrayOutput() RoleAssignmentArrayOutput {
	return o
}

func (o RoleAssignmentArrayOutput) ToRoleAssignmentArrayOutputWithContext(ctx context.Context) RoleAssignmentArrayOutput {
	return o
}

func (o RoleAssignmentArrayOutput) Index(i pulumi.IntInput) RoleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoleAssignment {
		return vs[0].([]*RoleAssignment)[vs[1].(int)]
	}).(RoleAssignmentOutput)
}

type RoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (RoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentMapOutput) ToRoleAssignmentMapOutput() RoleAssignmentMapOutput {
	return o
}

func (o RoleAssignmentMapOutput) ToRoleAssignmentMapOutputWithContext(ctx context.Context) RoleAssignmentMapOutput {
	return o
}

func (o RoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) RoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoleAssignment {
		return vs[0].(map[string]*RoleAssignment)[vs[1].(string)]
	}).(RoleAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentInput)(nil)).Elem(), &RoleAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentArrayInput)(nil)).Elem(), RoleAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentMapInput)(nil)).Elem(), RoleAssignmentMap{})
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
	pulumi.RegisterOutputType(RoleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(RoleAssignmentMapOutput{})
}
