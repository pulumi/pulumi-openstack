// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package database

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 DB user resource within OpenStack.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/db_user_v1.html.markdown.
type User struct {
	pulumi.CustomResourceState

	// A list of database user should have access to.
	Databases pulumi.StringArrayOutput `pulumi:"databases"`
	Host pulumi.StringPtrOutput `pulumi:"host"`
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// A unique name for the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// User's password.
	Password pulumi.StringOutput `pulumi:"password"`
	// Openstack region resource is created in.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("openstack:database/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("openstack:database/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// A list of database user should have access to.
	Databases []string `pulumi:"databases"`
	Host *string `pulumi:"host"`
	InstanceId *string `pulumi:"instanceId"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// User's password.
	Password *string `pulumi:"password"`
	// Openstack region resource is created in.
	Region *string `pulumi:"region"`
}

type UserState struct {
	// A list of database user should have access to.
	Databases pulumi.StringArrayInput
	Host pulumi.StringPtrInput
	InstanceId pulumi.StringPtrInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// User's password.
	Password pulumi.StringPtrInput
	// Openstack region resource is created in.
	Region pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// A list of database user should have access to.
	Databases []string `pulumi:"databases"`
	Host *string `pulumi:"host"`
	InstanceId string `pulumi:"instanceId"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// User's password.
	Password string `pulumi:"password"`
	// Openstack region resource is created in.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// A list of database user should have access to.
	Databases pulumi.StringArrayInput
	Host pulumi.StringPtrInput
	InstanceId pulumi.StringInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// User's password.
	Password pulumi.StringInput
	// Openstack region resource is created in.
	Region pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

