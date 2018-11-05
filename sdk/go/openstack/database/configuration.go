// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 DB configuration resource within OpenStack.
type Configuration struct {
	s *pulumi.ResourceState
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOpt) (*Configuration, error) {
	if args == nil || args.Datastore == nil {
		return nil, errors.New("missing required argument 'Datastore'")
	}
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["configurations"] = nil
		inputs["datastore"] = nil
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["region"] = nil
	} else {
		inputs["configurations"] = args.Configurations
		inputs["datastore"] = args.Datastore
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["region"] = args.Region
	}
	s, err := ctx.RegisterResource("openstack:database/configuration:Configuration", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Configuration{s: s}, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConfigurationState, opts ...pulumi.ResourceOpt) (*Configuration, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["configurations"] = state.Configurations
		inputs["datastore"] = state.Datastore
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("openstack:database/configuration:Configuration", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Configuration{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Configuration) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Configuration) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
func (r *Configuration) Configurations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["configurations"])
}

// An array of database engine type and version. The datastore
// object structure is documented below. Changing this creates resource.
func (r *Configuration) Datastore() *pulumi.Output {
	return r.s.State["datastore"]
}

// Description of the resource.
func (r *Configuration) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// A unique name for the resource.
func (r *Configuration) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The region in which to create the db instance. Changing this
// creates a new instance.
func (r *Configuration) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering Configuration resources.
type ConfigurationState struct {
	// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
	Configurations interface{}
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates resource.
	Datastore interface{}
	// Description of the resource.
	Description interface{}
	// A unique name for the resource.
	Name interface{}
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region interface{}
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
	Configurations interface{}
	// An array of database engine type and version. The datastore
	// object structure is documented below. Changing this creates resource.
	Datastore interface{}
	// Description of the resource.
	Description interface{}
	// A unique name for the resource.
	Name interface{}
	// The region in which to create the db instance. Changing this
	// creates a new instance.
	Region interface{}
}
