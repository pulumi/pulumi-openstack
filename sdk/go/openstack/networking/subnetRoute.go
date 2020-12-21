// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a routing entry on a OpenStack V2 subnet.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networking.NewRouter(ctx, "router1", &networking.RouterArgs{
// 			AdminStateUp: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		network1, err := networking.NewNetwork(ctx, "network1", &networking.NetworkArgs{
// 			AdminStateUp: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		subnet1, err := networking.NewSubnet(ctx, "subnet1", &networking.SubnetArgs{
// 			Cidr:      pulumi.String("192.168.199.0/24"),
// 			IpVersion: pulumi.Int(4),
// 			NetworkId: network1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = networking.NewSubnetRoute(ctx, "subnetRoute1", &networking.SubnetRouteArgs{
// 			DestinationCidr: pulumi.String("10.0.1.0/24"),
// 			NextHop:         pulumi.String("192.168.199.254"),
// 			SubnetId:        subnet1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Routing entries can be imported using a combined ID using the following format``<subnet_id>-route-<destination_cidr>-<next_hop>``
//
// ```sh
//  $ pulumi import openstack:networking/subnetRoute:SubnetRoute subnet_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidr == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidr'")
	}
	if args.NextHop == nil {
		return nil, errors.New("invalid value for required argument 'NextHop'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
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

type SubnetRouteInput interface {
	pulumi.Input

	ToSubnetRouteOutput() SubnetRouteOutput
	ToSubnetRouteOutputWithContext(ctx context.Context) SubnetRouteOutput
}

func (SubnetRoute) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetRoute)(nil)).Elem()
}

func (i SubnetRoute) ToSubnetRouteOutput() SubnetRouteOutput {
	return i.ToSubnetRouteOutputWithContext(context.Background())
}

func (i SubnetRoute) ToSubnetRouteOutputWithContext(ctx context.Context) SubnetRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetRouteOutput)
}

type SubnetRouteOutput struct {
	*pulumi.OutputState
}

func (SubnetRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetRouteOutput)(nil)).Elem()
}

func (o SubnetRouteOutput) ToSubnetRouteOutput() SubnetRouteOutput {
	return o
}

func (o SubnetRouteOutput) ToSubnetRouteOutputWithContext(ctx context.Context) SubnetRouteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetRouteOutput{})
}
