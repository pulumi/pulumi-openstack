// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Instance struct {
	pulumi.CustomResourceState

	// A list of IP addresses assigned to the instance.
	Addresses pulumi.StringArrayOutput `pulumi:"addresses"`
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
	// Database to be created on new instance. Changing this creates a
	// new instance.
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Datastore == nil {
		return nil, errors.New("invalid value for required argument 'Datastore'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
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
	// A list of IP addresses assigned to the instance.
	Addresses []string `pulumi:"addresses"`
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
	// Database to be created on new instance. Changing this creates a
	// new instance.
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
	// A list of IP addresses assigned to the instance.
	Addresses pulumi.StringArrayInput
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
	// Database to be created on new instance. Changing this creates a
	// new instance.
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
	// Database to be created on new instance. Changing this creates a
	// new instance.
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
	// Database to be created on new instance. Changing this creates a
	// new instance.
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

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *Instance) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

type InstancePtrInput interface {
	pulumi.Input

	ToInstancePtrOutput() InstancePtrOutput
	ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput
}

type instancePtrType InstanceArgs

func (*instancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (i *instancePtrType) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *instancePtrType) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o.ToInstancePtrOutputWithContext(context.Background())
}

func (o InstanceOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Instance) *Instance {
		return &v
	}).(InstancePtrOutput)
}

type InstancePtrOutput struct{ *pulumi.OutputState }

func (InstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (o InstancePtrOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) Elem() InstanceOutput {
	return o.ApplyT(func(v *Instance) Instance {
		if v != nil {
			return *v
		}
		var ret Instance
		return ret
	}).(InstanceOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Instance)(nil))
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Instance {
		return vs[0].([]Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Instance)(nil))
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Instance {
		return vs[0].(map[string]Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancePtrInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstancePtrOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
