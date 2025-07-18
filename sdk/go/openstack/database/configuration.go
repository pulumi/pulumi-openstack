// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Configuration struct {
	pulumi.CustomResourceState

	// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
	Configurations ConfigurationConfigurationArrayOutput `pulumi:"configurations"`
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates resource.
	Datastore ConfigurationDatastoreOutput `pulumi:"datastore"`
	// Description of the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// A unique name for the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Datastore == nil {
		return nil, errors.New("invalid value for required argument 'Datastore'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Configuration
	err := ctx.RegisterResource("openstack:database/configuration:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("openstack:database/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
	Configurations []ConfigurationConfiguration `pulumi:"configurations"`
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates resource.
	Datastore *ConfigurationDatastore `pulumi:"datastore"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region *string `pulumi:"region"`
}

type ConfigurationState struct {
	// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
	Configurations ConfigurationConfigurationArrayInput
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates resource.
	Datastore ConfigurationDatastorePtrInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region pulumi.StringPtrInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
	Configurations []ConfigurationConfiguration `pulumi:"configurations"`
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates resource.
	Datastore ConfigurationDatastore `pulumi:"datastore"`
	// Description of the resource.
	Description string `pulumi:"description"`
	// A unique name for the resource.
	Name *string `pulumi:"name"`
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
	Configurations ConfigurationConfigurationArrayInput
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates resource.
	Datastore ConfigurationDatastoreInput
	// Description of the resource.
	Description pulumi.StringInput
	// A unique name for the resource.
	Name pulumi.StringPtrInput
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region pulumi.StringPtrInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}

type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput
}

func (*Configuration) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (i *Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

// ConfigurationArrayInput is an input type that accepts ConfigurationArray and ConfigurationArrayOutput values.
// You can construct a concrete instance of `ConfigurationArrayInput` via:
//
//	ConfigurationArray{ ConfigurationArgs{...} }
type ConfigurationArrayInput interface {
	pulumi.Input

	ToConfigurationArrayOutput() ConfigurationArrayOutput
	ToConfigurationArrayOutputWithContext(context.Context) ConfigurationArrayOutput
}

type ConfigurationArray []ConfigurationInput

func (ConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Configuration)(nil)).Elem()
}

func (i ConfigurationArray) ToConfigurationArrayOutput() ConfigurationArrayOutput {
	return i.ToConfigurationArrayOutputWithContext(context.Background())
}

func (i ConfigurationArray) ToConfigurationArrayOutputWithContext(ctx context.Context) ConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationArrayOutput)
}

// ConfigurationMapInput is an input type that accepts ConfigurationMap and ConfigurationMapOutput values.
// You can construct a concrete instance of `ConfigurationMapInput` via:
//
//	ConfigurationMap{ "key": ConfigurationArgs{...} }
type ConfigurationMapInput interface {
	pulumi.Input

	ToConfigurationMapOutput() ConfigurationMapOutput
	ToConfigurationMapOutputWithContext(context.Context) ConfigurationMapOutput
}

type ConfigurationMap map[string]ConfigurationInput

func (ConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Configuration)(nil)).Elem()
}

func (i ConfigurationMap) ToConfigurationMapOutput() ConfigurationMapOutput {
	return i.ToConfigurationMapOutputWithContext(context.Background())
}

func (i ConfigurationMap) ToConfigurationMapOutputWithContext(ctx context.Context) ConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationMapOutput)
}

type ConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
func (o ConfigurationOutput) Configurations() ConfigurationConfigurationArrayOutput {
	return o.ApplyT(func(v *Configuration) ConfigurationConfigurationArrayOutput { return v.Configurations }).(ConfigurationConfigurationArrayOutput)
}

// An array of database engine type and version. The datastore
// object structure is documented below. Changing this creates resource.
func (o ConfigurationOutput) Datastore() ConfigurationDatastoreOutput {
	return o.ApplyT(func(v *Configuration) ConfigurationDatastoreOutput { return v.Datastore }).(ConfigurationDatastoreOutput)
}

// Description of the resource.
func (o ConfigurationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A unique name for the resource.
func (o ConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to create the db instance. Changing this
// creates a new instance.
func (o ConfigurationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Configuration)(nil)).Elem()
}

func (o ConfigurationArrayOutput) ToConfigurationArrayOutput() ConfigurationArrayOutput {
	return o
}

func (o ConfigurationArrayOutput) ToConfigurationArrayOutputWithContext(ctx context.Context) ConfigurationArrayOutput {
	return o
}

func (o ConfigurationArrayOutput) Index(i pulumi.IntInput) ConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Configuration {
		return vs[0].([]*Configuration)[vs[1].(int)]
	}).(ConfigurationOutput)
}

type ConfigurationMapOutput struct{ *pulumi.OutputState }

func (ConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Configuration)(nil)).Elem()
}

func (o ConfigurationMapOutput) ToConfigurationMapOutput() ConfigurationMapOutput {
	return o
}

func (o ConfigurationMapOutput) ToConfigurationMapOutputWithContext(ctx context.Context) ConfigurationMapOutput {
	return o
}

func (o ConfigurationMapOutput) MapIndex(k pulumi.StringInput) ConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Configuration {
		return vs[0].(map[string]*Configuration)[vs[1].(string)]
	}).(ConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationInput)(nil)).Elem(), &Configuration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationArrayInput)(nil)).Elem(), ConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationMapInput)(nil)).Elem(), ConfigurationMap{})
	pulumi.RegisterOutputType(ConfigurationOutput{})
	pulumi.RegisterOutputType(ConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationMapOutput{})
}
