// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V3 Endpoint resource within OpenStack Keystone.
// 
// > **Note:** This usually requires admin privileges.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_endpoint_v3.html.markdown.
type EndpointV3 struct {
	s *pulumi.ResourceState
}

// NewEndpointV3 registers a new resource with the given unique name, arguments, and options.
func NewEndpointV3(ctx *pulumi.Context,
	name string, args *EndpointV3Args, opts ...pulumi.ResourceOpt) (*EndpointV3, error) {
	if args == nil || args.EndpointRegion == nil {
		return nil, errors.New("missing required argument 'EndpointRegion'")
	}
	if args == nil || args.ServiceId == nil {
		return nil, errors.New("missing required argument 'ServiceId'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["endpointRegion"] = nil
		inputs["interface"] = nil
		inputs["name"] = nil
		inputs["region"] = nil
		inputs["serviceId"] = nil
		inputs["url"] = nil
	} else {
		inputs["endpointRegion"] = args.EndpointRegion
		inputs["interface"] = args.Interface
		inputs["name"] = args.Name
		inputs["region"] = args.Region
		inputs["serviceId"] = args.ServiceId
		inputs["url"] = args.Url
	}
	inputs["serviceName"] = nil
	inputs["serviceType"] = nil
	s, err := ctx.RegisterResource("openstack:identity/endpointV3:EndpointV3", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EndpointV3{s: s}, nil
}

// GetEndpointV3 gets an existing EndpointV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointV3(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EndpointV3State, opts ...pulumi.ResourceOpt) (*EndpointV3, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["endpointRegion"] = state.EndpointRegion
		inputs["interface"] = state.Interface
		inputs["name"] = state.Name
		inputs["region"] = state.Region
		inputs["serviceId"] = state.ServiceId
		inputs["serviceName"] = state.ServiceName
		inputs["serviceType"] = state.ServiceType
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("openstack:identity/endpointV3:EndpointV3", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &EndpointV3{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *EndpointV3) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *EndpointV3) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The endpoint region. The `region` and
// `endpointRegion` can be different.
func (r *EndpointV3) EndpointRegion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["endpointRegion"])
}

// The endpoint interface. Valid values are `public`,
// `internal` and `admin`. Default value is `public`
func (r *EndpointV3) Interface() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["interface"])
}

// The endpoint name.
func (r *EndpointV3) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used.
func (r *EndpointV3) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The endpoint service ID.
func (r *EndpointV3) ServiceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceId"])
}

// The service name of the endpoint.
func (r *EndpointV3) ServiceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceName"])
}

// The service type of the endpoint.
func (r *EndpointV3) ServiceType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceType"])
}

// The endpoint url.
func (r *EndpointV3) Url() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering EndpointV3 resources.
type EndpointV3State struct {
	// The endpoint region. The `region` and
	// `endpointRegion` can be different.
	EndpointRegion interface{}
	// The endpoint interface. Valid values are `public`,
	// `internal` and `admin`. Default value is `public`
	Interface interface{}
	// The endpoint name.
	Name interface{}
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region interface{}
	// The endpoint service ID.
	ServiceId interface{}
	// The service name of the endpoint.
	ServiceName interface{}
	// The service type of the endpoint.
	ServiceType interface{}
	// The endpoint url.
	Url interface{}
}

// The set of arguments for constructing a EndpointV3 resource.
type EndpointV3Args struct {
	// The endpoint region. The `region` and
	// `endpointRegion` can be different.
	EndpointRegion interface{}
	// The endpoint interface. Valid values are `public`,
	// `internal` and `admin`. Default value is `public`
	Interface interface{}
	// The endpoint name.
	Name interface{}
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region interface{}
	// The endpoint service ID.
	ServiceId interface{}
	// The endpoint url.
	Url interface{}
}