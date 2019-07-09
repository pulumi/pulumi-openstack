// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a routing entry on a OpenStack V2 subnet.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_subnet_route_v2.html.markdown.
type SubnetRoute struct {
	s *pulumi.ResourceState
}

// NewSubnetRoute registers a new resource with the given unique name, arguments, and options.
func NewSubnetRoute(ctx *pulumi.Context,
	name string, args *SubnetRouteArgs, opts ...pulumi.ResourceOpt) (*SubnetRoute, error) {
	if args == nil || args.DestinationCidr == nil {
		return nil, errors.New("missing required argument 'DestinationCidr'")
	}
	if args == nil || args.NextHop == nil {
		return nil, errors.New("missing required argument 'NextHop'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["destinationCidr"] = nil
		inputs["nextHop"] = nil
		inputs["region"] = nil
		inputs["subnetId"] = nil
	} else {
		inputs["destinationCidr"] = args.DestinationCidr
		inputs["nextHop"] = args.NextHop
		inputs["region"] = args.Region
		inputs["subnetId"] = args.SubnetId
	}
	s, err := ctx.RegisterResource("openstack:networking/subnetRoute:SubnetRoute", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubnetRoute{s: s}, nil
}

// GetSubnetRoute gets an existing SubnetRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetRoute(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubnetRouteState, opts ...pulumi.ResourceOpt) (*SubnetRoute, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["destinationCidr"] = state.DestinationCidr
		inputs["nextHop"] = state.NextHop
		inputs["region"] = state.Region
		inputs["subnetId"] = state.SubnetId
	}
	s, err := ctx.ReadResource("openstack:networking/subnetRoute:SubnetRoute", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubnetRoute{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SubnetRoute) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SubnetRoute) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// CIDR block to match on the packet’s destination IP. Changing
// this creates a new routing entry.
func (r *SubnetRoute) DestinationCidr() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["destinationCidr"])
}

// IP address of the next hop gateway.  Changing
// this creates a new routing entry.
func (r *SubnetRoute) NextHop() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["nextHop"])
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to configure a routing entry on a subnet. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// routing entry.
func (r *SubnetRoute) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// ID of the subnet this routing entry belongs to. Changing
// this creates a new routing entry.
func (r *SubnetRoute) SubnetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetId"])
}

// Input properties used for looking up and filtering SubnetRoute resources.
type SubnetRouteState struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr interface{}
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop interface{}
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region interface{}
	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetId interface{}
}

// The set of arguments for constructing a SubnetRoute resource.
type SubnetRouteArgs struct {
	// CIDR block to match on the packet’s destination IP. Changing
	// this creates a new routing entry.
	DestinationCidr interface{}
	// IP address of the next hop gateway.  Changing
	// this creates a new routing entry.
	NextHop interface{}
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region interface{}
	// ID of the subnet this routing entry belongs to. Changing
	// this creates a new routing entry.
	SubnetId interface{}
}
