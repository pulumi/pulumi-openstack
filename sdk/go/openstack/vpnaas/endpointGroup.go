// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpnaas

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 Neutron Endpoint Group resource within OpenStack.
type EndpointGroup struct {
	pulumi.CustomResourceState

	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints pulumi.StringArrayOutput `pulumi:"endpoints"`
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region pulumi.StringOutput `pulumi:"region"`
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type pulumi.StringOutput `pulumi:"type"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointGroup(ctx *pulumi.Context,
	name string, args *EndpointGroupArgs, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	if args == nil {
		args = &EndpointGroupArgs{}
	}
	var resource EndpointGroup
	err := ctx.RegisterResource("openstack:vpnaas/endpointGroup:EndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointGroup gets an existing EndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointGroupState, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	var resource EndpointGroup
	err := ctx.ReadResource("openstack:vpnaas/endpointGroup:EndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointGroup resources.
type endpointGroupState struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description *string `pulumi:"description"`
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints []string `pulumi:"endpoints"`
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region *string `pulumi:"region"`
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId *string `pulumi:"tenantId"`
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type *string `pulumi:"type"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type EndpointGroupState struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description pulumi.StringPtrInput
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints pulumi.StringArrayInput
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region pulumi.StringPtrInput
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId pulumi.StringPtrInput
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (EndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupState)(nil)).Elem()
}

type endpointGroupArgs struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description *string `pulumi:"description"`
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints []string `pulumi:"endpoints"`
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region *string `pulumi:"region"`
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId *string `pulumi:"tenantId"`
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type *string `pulumi:"type"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a EndpointGroup resource.
type EndpointGroupArgs struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description pulumi.StringPtrInput
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints pulumi.StringArrayInput
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region pulumi.StringPtrInput
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId pulumi.StringPtrInput
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (EndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupArgs)(nil)).Elem()
}
