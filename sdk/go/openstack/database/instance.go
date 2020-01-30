// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package database

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 DB instance resource within OpenStack.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/db_instance_v1.html.markdown.
type Instance struct {
	pulumi.CustomResourceState

	// Configuration ID to be attached to the instance. Database instance
	// will be rebooted when configuration is detached.
	ConfigurationId pulumi.StringPtrOutput `pulumi:"configurationId"`
	// An array of database name, charset and collate. The database
	// object structure is documented below.
	Databases InstanceDatabaseArrayOutput `pulumi:"databases"`
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates a new instance.
	Datastore InstanceDatastoreOutput `pulumi:"datastore"`
	// The flavor ID of the desired flavor for the instance.
	// Changing this creates new instance.
	FlavorId pulumi.StringOutput `pulumi:"flavorId"`
	// A unique name for the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new instance.
	Networks InstanceNetworkArrayOutput `pulumi:"networks"`
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the volume size in GB. Changing this creates new instance.
	Size pulumi.IntOutput `pulumi:"size"`
	// An array of username, password, host and databases. The user
	// object structure is documented below.
	Users InstanceUserArrayOutput `pulumi:"users"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil || args.Datastore == nil {
		return nil, errors.New("missing required argument 'Datastore'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &InstanceArgs{}
	}
	var resource Instance
	err := ctx.RegisterResource("openstack:database/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("openstack:database/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Configuration ID to be attached to the instance. Database instance
	// will be rebooted when configuration is detached.
	ConfigurationId *string `pulumi:"configurationId"`
	// An array of database name, charset and collate. The database
	// object structure is documented below.
	Databases []InstanceDatabase `pulumi:"databases"`
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates a new instance.
	Datastore *InstanceDatastore `pulumi:"datastore"`
	// The flavor ID of the desired flavor for the instance.
	// Changing this creates new instance.
	FlavorId *string `pulumi:"flavorId"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new instance.
	Networks []InstanceNetwork `pulumi:"networks"`
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region *string `pulumi:"region"`
	// Specifies the volume size in GB. Changing this creates new instance.
	Size *int `pulumi:"size"`
	// An array of username, password, host and databases. The user
	// object structure is documented below.
	Users []InstanceUser `pulumi:"users"`
}

type InstanceState struct {
	// Configuration ID to be attached to the instance. Database instance
	// will be rebooted when configuration is detached.
	ConfigurationId pulumi.StringPtrInput
	// An array of database name, charset and collate. The database
	// object structure is documented below.
	Databases InstanceDatabaseArrayInput
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates a new instance.
	Datastore InstanceDatastorePtrInput
	// The flavor ID of the desired flavor for the instance.
	// Changing this creates new instance.
	FlavorId pulumi.StringPtrInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new instance.
	Networks InstanceNetworkArrayInput
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region pulumi.StringPtrInput
	// Specifies the volume size in GB. Changing this creates new instance.
	Size pulumi.IntPtrInput
	// An array of username, password, host and databases. The user
	// object structure is documented below.
	Users InstanceUserArrayInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Configuration ID to be attached to the instance. Database instance
	// will be rebooted when configuration is detached.
	ConfigurationId *string `pulumi:"configurationId"`
	// An array of database name, charset and collate. The database
	// object structure is documented below.
	Databases []InstanceDatabase `pulumi:"databases"`
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates a new instance.
	Datastore InstanceDatastore `pulumi:"datastore"`
	// The flavor ID of the desired flavor for the instance.
	// Changing this creates new instance.
	FlavorId *string `pulumi:"flavorId"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new instance.
	Networks []InstanceNetwork `pulumi:"networks"`
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region *string `pulumi:"region"`
	// Specifies the volume size in GB. Changing this creates new instance.
	Size int `pulumi:"size"`
	// An array of username, password, host and databases. The user
	// object structure is documented below.
	Users []InstanceUser `pulumi:"users"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Configuration ID to be attached to the instance. Database instance
	// will be rebooted when configuration is detached.
	ConfigurationId pulumi.StringPtrInput
	// An array of database name, charset and collate. The database
	// object structure is documented below.
	Databases InstanceDatabaseArrayInput
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates a new instance.
	Datastore InstanceDatastoreInput
	// The flavor ID of the desired flavor for the instance.
	// Changing this creates new instance.
	FlavorId pulumi.StringPtrInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new instance.
	Networks InstanceNetworkArrayInput
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region pulumi.StringPtrInput
	// Specifies the volume size in GB. Changing this creates new instance.
	Size pulumi.IntInput
	// An array of username, password, host and databases. The user
	// object structure is documented below.
	Users InstanceUserArrayInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

