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

// Manages a V3 Inherit Role assignment within OpenStack Keystone. This uses the
// Openstack keystone `OS-INHERIT` api to created inherit roles within domains
// and parent projects for users and groups.
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
//			user1, err := identity.NewUser(ctx, "user1", &identity.UserArgs{
//				DomainId: pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			role1, err := identity.NewRole(ctx, "role1", &identity.RoleArgs{
//				DomainId: pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewInheritRoleAssignment(ctx, "roleAssignment1", &identity.InheritRoleAssignmentArgs{
//				UserId:   user1.ID(),
//				DomainId: pulumi.String("default"),
//				RoleId:   role1.ID(),
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
// Inherit role assignments can be imported using a constructed id. The id should
//
// have the form of `domainID/projectID/groupID/userID/roleID`. When something is not used then leave blank. For example this will import the inherit role assignment for:
//
// projectID014395cd-89fc-4c9b-96b7-13d1ee79dad2, userID4142e64b-1b35-44a0-9b1e-5affc7af1106, roleIDea257959-eeb1-4c10-8d33-26f0409a755d ( domainID and groupID are left blank)
//
// ```sh
//
//	$ pulumi import openstack:identity/inheritRoleAssignment:InheritRoleAssignment role_assignment_1 /014395cd-89fc-4c9b-96b7-13d1ee79dad2//4142e64b-1b35-44a0-9b1e-5affc7af1106/ea257959-eeb1-4c10-8d33-26f0409a755d
//
// ```
type InheritRoleAssignment struct {
	pulumi.CustomResourceState

	// The domain to assign the role in.
	DomainId pulumi.StringPtrOutput `pulumi:"domainId"`
	// The group to assign the role to.
	GroupId pulumi.StringPtrOutput `pulumi:"groupId"`
	// The project to assign the role in.
	// The project should be able to containt child projects.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	Region    pulumi.StringOutput    `pulumi:"region"`
	// The role to assign.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// The user to assign the role to.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewInheritRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewInheritRoleAssignment(ctx *pulumi.Context,
	name string, args *InheritRoleAssignmentArgs, opts ...pulumi.ResourceOption) (*InheritRoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InheritRoleAssignment
	err := ctx.RegisterResource("openstack:identity/inheritRoleAssignment:InheritRoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInheritRoleAssignment gets an existing InheritRoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInheritRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InheritRoleAssignmentState, opts ...pulumi.ResourceOption) (*InheritRoleAssignment, error) {
	var resource InheritRoleAssignment
	err := ctx.ReadResource("openstack:identity/inheritRoleAssignment:InheritRoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InheritRoleAssignment resources.
type inheritRoleAssignmentState struct {
	// The domain to assign the role in.
	DomainId *string `pulumi:"domainId"`
	// The group to assign the role to.
	GroupId *string `pulumi:"groupId"`
	// The project to assign the role in.
	// The project should be able to containt child projects.
	ProjectId *string `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// The role to assign.
	RoleId *string `pulumi:"roleId"`
	// The user to assign the role to.
	UserId *string `pulumi:"userId"`
}

type InheritRoleAssignmentState struct {
	// The domain to assign the role in.
	DomainId pulumi.StringPtrInput
	// The group to assign the role to.
	GroupId pulumi.StringPtrInput
	// The project to assign the role in.
	// The project should be able to containt child projects.
	ProjectId pulumi.StringPtrInput
	Region    pulumi.StringPtrInput
	// The role to assign.
	RoleId pulumi.StringPtrInput
	// The user to assign the role to.
	UserId pulumi.StringPtrInput
}

func (InheritRoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*inheritRoleAssignmentState)(nil)).Elem()
}

type inheritRoleAssignmentArgs struct {
	// The domain to assign the role in.
	DomainId *string `pulumi:"domainId"`
	// The group to assign the role to.
	GroupId *string `pulumi:"groupId"`
	// The project to assign the role in.
	// The project should be able to containt child projects.
	ProjectId *string `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// The role to assign.
	RoleId string `pulumi:"roleId"`
	// The user to assign the role to.
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a InheritRoleAssignment resource.
type InheritRoleAssignmentArgs struct {
	// The domain to assign the role in.
	DomainId pulumi.StringPtrInput
	// The group to assign the role to.
	GroupId pulumi.StringPtrInput
	// The project to assign the role in.
	// The project should be able to containt child projects.
	ProjectId pulumi.StringPtrInput
	Region    pulumi.StringPtrInput
	// The role to assign.
	RoleId pulumi.StringInput
	// The user to assign the role to.
	UserId pulumi.StringPtrInput
}

func (InheritRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inheritRoleAssignmentArgs)(nil)).Elem()
}

type InheritRoleAssignmentInput interface {
	pulumi.Input

	ToInheritRoleAssignmentOutput() InheritRoleAssignmentOutput
	ToInheritRoleAssignmentOutputWithContext(ctx context.Context) InheritRoleAssignmentOutput
}

func (*InheritRoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**InheritRoleAssignment)(nil)).Elem()
}

func (i *InheritRoleAssignment) ToInheritRoleAssignmentOutput() InheritRoleAssignmentOutput {
	return i.ToInheritRoleAssignmentOutputWithContext(context.Background())
}

func (i *InheritRoleAssignment) ToInheritRoleAssignmentOutputWithContext(ctx context.Context) InheritRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InheritRoleAssignmentOutput)
}

// InheritRoleAssignmentArrayInput is an input type that accepts InheritRoleAssignmentArray and InheritRoleAssignmentArrayOutput values.
// You can construct a concrete instance of `InheritRoleAssignmentArrayInput` via:
//
//	InheritRoleAssignmentArray{ InheritRoleAssignmentArgs{...} }
type InheritRoleAssignmentArrayInput interface {
	pulumi.Input

	ToInheritRoleAssignmentArrayOutput() InheritRoleAssignmentArrayOutput
	ToInheritRoleAssignmentArrayOutputWithContext(context.Context) InheritRoleAssignmentArrayOutput
}

type InheritRoleAssignmentArray []InheritRoleAssignmentInput

func (InheritRoleAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InheritRoleAssignment)(nil)).Elem()
}

func (i InheritRoleAssignmentArray) ToInheritRoleAssignmentArrayOutput() InheritRoleAssignmentArrayOutput {
	return i.ToInheritRoleAssignmentArrayOutputWithContext(context.Background())
}

func (i InheritRoleAssignmentArray) ToInheritRoleAssignmentArrayOutputWithContext(ctx context.Context) InheritRoleAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InheritRoleAssignmentArrayOutput)
}

// InheritRoleAssignmentMapInput is an input type that accepts InheritRoleAssignmentMap and InheritRoleAssignmentMapOutput values.
// You can construct a concrete instance of `InheritRoleAssignmentMapInput` via:
//
//	InheritRoleAssignmentMap{ "key": InheritRoleAssignmentArgs{...} }
type InheritRoleAssignmentMapInput interface {
	pulumi.Input

	ToInheritRoleAssignmentMapOutput() InheritRoleAssignmentMapOutput
	ToInheritRoleAssignmentMapOutputWithContext(context.Context) InheritRoleAssignmentMapOutput
}

type InheritRoleAssignmentMap map[string]InheritRoleAssignmentInput

func (InheritRoleAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InheritRoleAssignment)(nil)).Elem()
}

func (i InheritRoleAssignmentMap) ToInheritRoleAssignmentMapOutput() InheritRoleAssignmentMapOutput {
	return i.ToInheritRoleAssignmentMapOutputWithContext(context.Background())
}

func (i InheritRoleAssignmentMap) ToInheritRoleAssignmentMapOutputWithContext(ctx context.Context) InheritRoleAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InheritRoleAssignmentMapOutput)
}

type InheritRoleAssignmentOutput struct{ *pulumi.OutputState }

func (InheritRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InheritRoleAssignment)(nil)).Elem()
}

func (o InheritRoleAssignmentOutput) ToInheritRoleAssignmentOutput() InheritRoleAssignmentOutput {
	return o
}

func (o InheritRoleAssignmentOutput) ToInheritRoleAssignmentOutputWithContext(ctx context.Context) InheritRoleAssignmentOutput {
	return o
}

// The domain to assign the role in.
func (o InheritRoleAssignmentOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InheritRoleAssignment) pulumi.StringPtrOutput { return v.DomainId }).(pulumi.StringPtrOutput)
}

// The group to assign the role to.
func (o InheritRoleAssignmentOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InheritRoleAssignment) pulumi.StringPtrOutput { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The project to assign the role in.
// The project should be able to containt child projects.
func (o InheritRoleAssignmentOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InheritRoleAssignment) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o InheritRoleAssignmentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *InheritRoleAssignment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role to assign.
func (o InheritRoleAssignmentOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *InheritRoleAssignment) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// The user to assign the role to.
func (o InheritRoleAssignmentOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InheritRoleAssignment) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

type InheritRoleAssignmentArrayOutput struct{ *pulumi.OutputState }

func (InheritRoleAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InheritRoleAssignment)(nil)).Elem()
}

func (o InheritRoleAssignmentArrayOutput) ToInheritRoleAssignmentArrayOutput() InheritRoleAssignmentArrayOutput {
	return o
}

func (o InheritRoleAssignmentArrayOutput) ToInheritRoleAssignmentArrayOutputWithContext(ctx context.Context) InheritRoleAssignmentArrayOutput {
	return o
}

func (o InheritRoleAssignmentArrayOutput) Index(i pulumi.IntInput) InheritRoleAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InheritRoleAssignment {
		return vs[0].([]*InheritRoleAssignment)[vs[1].(int)]
	}).(InheritRoleAssignmentOutput)
}

type InheritRoleAssignmentMapOutput struct{ *pulumi.OutputState }

func (InheritRoleAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InheritRoleAssignment)(nil)).Elem()
}

func (o InheritRoleAssignmentMapOutput) ToInheritRoleAssignmentMapOutput() InheritRoleAssignmentMapOutput {
	return o
}

func (o InheritRoleAssignmentMapOutput) ToInheritRoleAssignmentMapOutputWithContext(ctx context.Context) InheritRoleAssignmentMapOutput {
	return o
}

func (o InheritRoleAssignmentMapOutput) MapIndex(k pulumi.StringInput) InheritRoleAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InheritRoleAssignment {
		return vs[0].(map[string]*InheritRoleAssignment)[vs[1].(string)]
	}).(InheritRoleAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InheritRoleAssignmentInput)(nil)).Elem(), &InheritRoleAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*InheritRoleAssignmentArrayInput)(nil)).Elem(), InheritRoleAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InheritRoleAssignmentMapInput)(nil)).Elem(), InheritRoleAssignmentMap{})
	pulumi.RegisterOutputType(InheritRoleAssignmentOutput{})
	pulumi.RegisterOutputType(InheritRoleAssignmentArrayOutput{})
	pulumi.RegisterOutputType(InheritRoleAssignmentMapOutput{})
}
