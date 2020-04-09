// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identity

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V3 Service resource within OpenStack Keystone.
//
// > **Note:** This usually requires admin privileges.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_service_v3.html.markdown.
type ServiceV3 struct {
	pulumi.CustomResourceState

	// The service description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The service status. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The service type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceV3 registers a new resource with the given unique name, arguments, and options.
func NewServiceV3(ctx *pulumi.Context,
	name string, args *ServiceV3Args, opts ...pulumi.ResourceOption) (*ServiceV3, error) {
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &ServiceV3Args{}
	}
	var resource ServiceV3
	err := ctx.RegisterResource("openstack:identity/serviceV3:ServiceV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceV3 gets an existing ServiceV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceV3State, opts ...pulumi.ResourceOption) (*ServiceV3, error) {
	var resource ServiceV3
	err := ctx.ReadResource("openstack:identity/serviceV3:ServiceV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceV3 resources.
type serviceV3State struct {
	// The service description.
	Description *string `pulumi:"description"`
	// The service status. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The service name.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The service type.
	Type *string `pulumi:"type"`
}

type ServiceV3State struct {
	// The service description.
	Description pulumi.StringPtrInput
	// The service status. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The service name.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput
	// The service type.
	Type pulumi.StringPtrInput
}

func (ServiceV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceV3State)(nil)).Elem()
}

type serviceV3Args struct {
	// The service description.
	Description *string `pulumi:"description"`
	// The service status. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The service name.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The service type.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ServiceV3 resource.
type ServiceV3Args struct {
	// The service description.
	Description pulumi.StringPtrInput
	// The service status. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The service name.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput
	// The service type.
	Type pulumi.StringInput
}

func (ServiceV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceV3Args)(nil)).Elem()
}
