// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 router interface resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network1, err := networking.NewNetwork(ctx, "network_1", &networking.NetworkArgs{
//				Name:         pulumi.String("tf_test_network"),
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			subnet1, err := networking.NewSubnet(ctx, "subnet_1", &networking.SubnetArgs{
//				NetworkId: network1.ID(),
//				Cidr:      pulumi.String("192.168.199.0/24"),
//				IpVersion: pulumi.Int(4),
//			})
//			if err != nil {
//				return err
//			}
//			router1, err := networking.NewRouter(ctx, "router_1", &networking.RouterArgs{
//				Name:              pulumi.String("my_router"),
//				ExternalNetworkId: pulumi.String("f67f0d72-0ddf-11e4-9d95-e1f29f417e2f"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewRouterInterface(ctx, "router_interface_1", &networking.RouterInterfaceArgs{
//				RouterId: router1.ID(),
//				SubnetId: subnet1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Router Interfaces can be imported using the port `id`, e.g.
//
// $ openstack port list --router <router name or id>
//
// ```sh
// $ pulumi import openstack:networking/routerInterface:RouterInterface int_1 port_id
// ```
type RouterInterface struct {
	pulumi.CustomResourceState

	// A boolean indicating whether the routes from the
	// corresponding router ID should be deleted so that the router interface can
	// be destroyed without any errors. The default value is `false`.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortId pulumi.StringOutput `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router interface.
	Region pulumi.StringOutput `pulumi:"region"`
	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewRouterInterface registers a new resource with the given unique name, arguments, and options.
func NewRouterInterface(ctx *pulumi.Context,
	name string, args *RouterInterfaceArgs, opts ...pulumi.ResourceOption) (*RouterInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouterId == nil {
		return nil, errors.New("invalid value for required argument 'RouterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterInterface
	err := ctx.RegisterResource("openstack:networking/routerInterface:RouterInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterInterface gets an existing RouterInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterInterfaceState, opts ...pulumi.ResourceOption) (*RouterInterface, error) {
	var resource RouterInterface
	err := ctx.ReadResource("openstack:networking/routerInterface:RouterInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterInterface resources.
type routerInterfaceState struct {
	// A boolean indicating whether the routes from the
	// corresponding router ID should be deleted so that the router interface can
	// be destroyed without any errors. The default value is `false`.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortId *string `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router interface.
	Region *string `pulumi:"region"`
	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	RouterId *string `pulumi:"routerId"`
	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	SubnetId *string `pulumi:"subnetId"`
}

type RouterInterfaceState struct {
	// A boolean indicating whether the routes from the
	// corresponding router ID should be deleted so that the router interface can
	// be destroyed without any errors. The default value is `false`.
	ForceDestroy pulumi.BoolPtrInput
	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortId pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router interface.
	Region pulumi.StringPtrInput
	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	RouterId pulumi.StringPtrInput
	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	SubnetId pulumi.StringPtrInput
}

func (RouterInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerInterfaceState)(nil)).Elem()
}

type routerInterfaceArgs struct {
	// A boolean indicating whether the routes from the
	// corresponding router ID should be deleted so that the router interface can
	// be destroyed without any errors. The default value is `false`.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortId *string `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router interface.
	Region *string `pulumi:"region"`
	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	RouterId string `pulumi:"routerId"`
	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	SubnetId *string `pulumi:"subnetId"`
}

// The set of arguments for constructing a RouterInterface resource.
type RouterInterfaceArgs struct {
	// A boolean indicating whether the routes from the
	// corresponding router ID should be deleted so that the router interface can
	// be destroyed without any errors. The default value is `false`.
	ForceDestroy pulumi.BoolPtrInput
	// ID of the port this interface connects to. Changing
	// this creates a new router interface.
	PortId pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router interface.
	Region pulumi.StringPtrInput
	// ID of the router this interface belongs to. Changing
	// this creates a new router interface.
	RouterId pulumi.StringInput
	// ID of the subnet this interface connects to. Changing
	// this creates a new router interface.
	SubnetId pulumi.StringPtrInput
}

func (RouterInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerInterfaceArgs)(nil)).Elem()
}

type RouterInterfaceInput interface {
	pulumi.Input

	ToRouterInterfaceOutput() RouterInterfaceOutput
	ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput
}

func (*RouterInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterInterface)(nil)).Elem()
}

func (i *RouterInterface) ToRouterInterfaceOutput() RouterInterfaceOutput {
	return i.ToRouterInterfaceOutputWithContext(context.Background())
}

func (i *RouterInterface) ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterInterfaceOutput)
}

// RouterInterfaceArrayInput is an input type that accepts RouterInterfaceArray and RouterInterfaceArrayOutput values.
// You can construct a concrete instance of `RouterInterfaceArrayInput` via:
//
//	RouterInterfaceArray{ RouterInterfaceArgs{...} }
type RouterInterfaceArrayInput interface {
	pulumi.Input

	ToRouterInterfaceArrayOutput() RouterInterfaceArrayOutput
	ToRouterInterfaceArrayOutputWithContext(context.Context) RouterInterfaceArrayOutput
}

type RouterInterfaceArray []RouterInterfaceInput

func (RouterInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterInterface)(nil)).Elem()
}

func (i RouterInterfaceArray) ToRouterInterfaceArrayOutput() RouterInterfaceArrayOutput {
	return i.ToRouterInterfaceArrayOutputWithContext(context.Background())
}

func (i RouterInterfaceArray) ToRouterInterfaceArrayOutputWithContext(ctx context.Context) RouterInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterInterfaceArrayOutput)
}

// RouterInterfaceMapInput is an input type that accepts RouterInterfaceMap and RouterInterfaceMapOutput values.
// You can construct a concrete instance of `RouterInterfaceMapInput` via:
//
//	RouterInterfaceMap{ "key": RouterInterfaceArgs{...} }
type RouterInterfaceMapInput interface {
	pulumi.Input

	ToRouterInterfaceMapOutput() RouterInterfaceMapOutput
	ToRouterInterfaceMapOutputWithContext(context.Context) RouterInterfaceMapOutput
}

type RouterInterfaceMap map[string]RouterInterfaceInput

func (RouterInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterInterface)(nil)).Elem()
}

func (i RouterInterfaceMap) ToRouterInterfaceMapOutput() RouterInterfaceMapOutput {
	return i.ToRouterInterfaceMapOutputWithContext(context.Background())
}

func (i RouterInterfaceMap) ToRouterInterfaceMapOutputWithContext(ctx context.Context) RouterInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterInterfaceMapOutput)
}

type RouterInterfaceOutput struct{ *pulumi.OutputState }

func (RouterInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterInterface)(nil)).Elem()
}

func (o RouterInterfaceOutput) ToRouterInterfaceOutput() RouterInterfaceOutput {
	return o
}

func (o RouterInterfaceOutput) ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput {
	return o
}

// A boolean indicating whether the routes from the
// corresponding router ID should be deleted so that the router interface can
// be destroyed without any errors. The default value is `false`.
func (o RouterInterfaceOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// ID of the port this interface connects to. Changing
// this creates a new router interface.
func (o RouterInterfaceOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.PortId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to create a router. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// router interface.
func (o RouterInterfaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// ID of the router this interface belongs to. Changing
// this creates a new router interface.
func (o RouterInterfaceOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.RouterId }).(pulumi.StringOutput)
}

// ID of the subnet this interface connects to. Changing
// this creates a new router interface.
func (o RouterInterfaceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

type RouterInterfaceArrayOutput struct{ *pulumi.OutputState }

func (RouterInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterInterface)(nil)).Elem()
}

func (o RouterInterfaceArrayOutput) ToRouterInterfaceArrayOutput() RouterInterfaceArrayOutput {
	return o
}

func (o RouterInterfaceArrayOutput) ToRouterInterfaceArrayOutputWithContext(ctx context.Context) RouterInterfaceArrayOutput {
	return o
}

func (o RouterInterfaceArrayOutput) Index(i pulumi.IntInput) RouterInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterInterface {
		return vs[0].([]*RouterInterface)[vs[1].(int)]
	}).(RouterInterfaceOutput)
}

type RouterInterfaceMapOutput struct{ *pulumi.OutputState }

func (RouterInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterInterface)(nil)).Elem()
}

func (o RouterInterfaceMapOutput) ToRouterInterfaceMapOutput() RouterInterfaceMapOutput {
	return o
}

func (o RouterInterfaceMapOutput) ToRouterInterfaceMapOutputWithContext(ctx context.Context) RouterInterfaceMapOutput {
	return o
}

func (o RouterInterfaceMapOutput) MapIndex(k pulumi.StringInput) RouterInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterInterface {
		return vs[0].(map[string]*RouterInterface)[vs[1].(string)]
	}).(RouterInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInterfaceInput)(nil)).Elem(), &RouterInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInterfaceArrayInput)(nil)).Elem(), RouterInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInterfaceMapInput)(nil)).Elem(), RouterInterfaceMap{})
	pulumi.RegisterOutputType(RouterInterfaceOutput{})
	pulumi.RegisterOutputType(RouterInterfaceArrayOutput{})
	pulumi.RegisterOutputType(RouterInterfaceMapOutput{})
}
