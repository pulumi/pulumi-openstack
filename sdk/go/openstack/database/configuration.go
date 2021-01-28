// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V1 DB configuration resource within OpenStack.
//
// ## Example Usage
// ### Configuration
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/database"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := database.NewConfiguration(ctx, "test", &database.ConfigurationArgs{
// 			Configurations: database.ConfigurationConfigurationArray{
// 				&database.ConfigurationConfigurationArgs{
// 					Name:  pulumi.String("max_connections"),
// 					Value: pulumi.String("200"),
// 				},
// 			},
// 			Datastore: &database.ConfigurationDatastoreArgs{
// 				Type:    pulumi.String("mysql"),
// 				Version: pulumi.String("mysql-5.7"),
// 			},
// 			Description: pulumi.String("description"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Configuration struct {
	pulumi.CustomResourceState

	// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
	Configurations ConfigurationConfigurationArrayOutput `pulumi:"configurations"`
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates resource.
	Datastore ConfigurationDatastoreOutput `pulumi:"datastore"`
	// Description of the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Configuration parameter name. Changing this creates a new resource.
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
	// Configuration parameter name. Changing this creates a new resource.
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
	// Configuration parameter name. Changing this creates a new resource.
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
	// Configuration parameter name. Changing this creates a new resource.
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
	// Configuration parameter name. Changing this creates a new resource.
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
	return reflect.TypeOf((*Configuration)(nil))
}

func (i *Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

type ConfigurationOutput struct {
	*pulumi.OutputState
}

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil))
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationOutput{})
}
