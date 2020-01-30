// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package database

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 DB database resource within OpenStack.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/db_database_v1.html.markdown.
type Database struct {
	pulumi.CustomResourceState

	// The ID for the database instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// A unique name for the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Openstack region resource is created in.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &DatabaseArgs{}
	}
	var resource Database
	err := ctx.RegisterResource("openstack:database/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("openstack:database/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// The ID for the database instance.
	InstanceId *string `pulumi:"instanceId"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// Openstack region resource is created in.
	Region *string `pulumi:"region"`
}

type DatabaseState struct {
	// The ID for the database instance.
	InstanceId pulumi.StringPtrInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// Openstack region resource is created in.
	Region pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The ID for the database instance.
	InstanceId string `pulumi:"instanceId"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// Openstack region resource is created in.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The ID for the database instance.
	InstanceId pulumi.StringInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// Openstack region resource is created in.
	Region pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

