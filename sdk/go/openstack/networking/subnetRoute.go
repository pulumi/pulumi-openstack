// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package networking

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a routing entry on a OpenStack V2 subnet.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_subnet_route_v2.html.markdown.
type SubnetRoute struct {
	pulumi.CustomResourceState

	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr pulumi.StringOutput `pulumi:"destinationCidr"`
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop pulumi.StringOutput `pulumi:"nextHop"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringOutput `pulumi:"region"`
	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewSubnetRoute registers a new resource with the given unique name, arguments, and options.
func NewSubnetRoute(ctx *pulumi.Context,
	name string, args *SubnetRouteArgs, opts ...pulumi.ResourceOption) (*SubnetRoute, error) {
	if args == nil || args.DestinationCidr == nil {
		return nil, errors.New("missing required argument 'DestinationCidr'")
	}
	if args == nil || args.NextHop == nil {
		return nil, errors.New("missing required argument 'NextHop'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil {
		args = &SubnetRouteArgs{}
	}
	var resource SubnetRoute
	err := ctx.RegisterResource("openstack:networking/subnetRoute:SubnetRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetRoute gets an existing SubnetRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetRouteState, opts ...pulumi.ResourceOption) (*SubnetRoute, error) {
	var resource SubnetRoute
	err := ctx.ReadResource("openstack:networking/subnetRoute:SubnetRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetRoute resources.
type subnetRouteState struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr *string `pulumi:"destinationCidr"`
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop *string `pulumi:"nextHop"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `pulumi:"region"`
	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetId *string `pulumi:"subnetId"`
}

type SubnetRouteState struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr pulumi.StringPtrInput
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringPtrInput
	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetId pulumi.StringPtrInput
}

func (SubnetRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetRouteState)(nil)).Elem()
}

type subnetRouteArgs struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr string `pulumi:"destinationCidr"`
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop string `pulumi:"nextHop"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `pulumi:"region"`
	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a SubnetRoute resource.
type SubnetRouteArgs struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr pulumi.StringInput
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop pulumi.StringInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringPtrInput
	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetId pulumi.StringInput
}

func (SubnetRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetRouteArgs)(nil)).Elem()
}
