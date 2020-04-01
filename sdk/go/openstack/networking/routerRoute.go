// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package networking

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a routing entry on a OpenStack V2 router.
//
// ## Notes
//
// The `nextHop` IP address must be directly reachable from the router at the ``networking.RouterRoute``
// resource creation time.  You can ensure that by explicitly specifying a dependency on the ``networking.RouterInterface``
// resource that connects the next hop to the router, as in the example above.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_router_route_v2.html.markdown.
type RouterRoute struct {
	pulumi.CustomResourceState

	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr pulumi.StringOutput `pulumi:"destinationCidr"`
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop pulumi.StringOutput `pulumi:"nextHop"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringOutput `pulumi:"region"`
	// ID of the router this routing entry belongs to. Changing
	// this creates a new routing entry.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
}

// NewRouterRoute registers a new resource with the given unique name, arguments, and options.
func NewRouterRoute(ctx *pulumi.Context,
	name string, args *RouterRouteArgs, opts ...pulumi.ResourceOption) (*RouterRoute, error) {
	if args == nil || args.DestinationCidr == nil {
		return nil, errors.New("missing required argument 'DestinationCidr'")
	}
	if args == nil || args.NextHop == nil {
		return nil, errors.New("missing required argument 'NextHop'")
	}
	if args == nil || args.RouterId == nil {
		return nil, errors.New("missing required argument 'RouterId'")
	}
	if args == nil {
		args = &RouterRouteArgs{}
	}
	var resource RouterRoute
	err := ctx.RegisterResource("openstack:networking/routerRoute:RouterRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterRoute gets an existing RouterRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterRouteState, opts ...pulumi.ResourceOption) (*RouterRoute, error) {
	var resource RouterRoute
	err := ctx.ReadResource("openstack:networking/routerRoute:RouterRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterRoute resources.
type routerRouteState struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr *string `pulumi:"destinationCidr"`
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop *string `pulumi:"nextHop"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `pulumi:"region"`
	// ID of the router this routing entry belongs to. Changing
	// this creates a new routing entry.
	RouterId *string `pulumi:"routerId"`
}

type RouterRouteState struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr pulumi.StringPtrInput
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringPtrInput
	// ID of the router this routing entry belongs to. Changing
	// this creates a new routing entry.
	RouterId pulumi.StringPtrInput
}

func (RouterRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerRouteState)(nil)).Elem()
}

type routerRouteArgs struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr string `pulumi:"destinationCidr"`
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop string `pulumi:"nextHop"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `pulumi:"region"`
	// ID of the router this routing entry belongs to. Changing
	// this creates a new routing entry.
	RouterId string `pulumi:"routerId"`
}

// The set of arguments for constructing a RouterRoute resource.
type RouterRouteArgs struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr pulumi.StringInput
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop pulumi.StringInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringPtrInput
	// ID of the router this routing entry belongs to. Changing
	// this creates a new routing entry.
	RouterId pulumi.StringInput
}

func (RouterRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerRouteArgs)(nil)).Elem()
}
