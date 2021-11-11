// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 router resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networking.NewRouter(ctx, "router1", &networking.RouterArgs{
// 			AdminStateUp:      pulumi.Bool(true),
// 			ExternalNetworkId: pulumi.String("f67f0d72-0ddf-11e4-9d95-e1f29f417e2f"),
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
// Routers can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:networking/router:Router router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2
// ```
type Router struct {
	pulumi.CustomResourceState

	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing router.
	AdminStateUp pulumi.BoolOutput `pulumi:"adminStateUp"`
	// The collection of tags assigned on the router, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new router.
	AvailabilityZoneHints pulumi.StringArrayOutput `pulumi:"availabilityZoneHints"`
	// Human-readable description for the router.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	Distributed pulumi.BoolOutput `pulumi:"distributed"`
	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An `externalNetworkId` has to be set in order to
	// set this property. Changing this updates the `enableSnat` of the router.
	// Setting this value **requires** an **ext-gw-mode** extension to be enabled
	// in OpenStack Neutron.
	EnableSnat pulumi.BoolOutput `pulumi:"enableSnat"`
	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An `externalNetworkId`
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	ExternalFixedIps RouterExternalFixedIpArrayOutput `pulumi:"externalFixedIps"`
	// The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	//
	// Deprecated: use external_network_id instead
	ExternalGateway pulumi.StringOutput `pulumi:"externalGateway"`
	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	ExternalNetworkId pulumi.StringOutput `pulumi:"externalNetworkId"`
	// A list of external subnet IDs to try over
	// each to obtain a fixed IP for the router. If a subnet ID in a list has
	// exhausted floating IP pool, the next subnet ID will be tried. This argument is
	// used only during the router creation and allows to set only one external fixed
	// IP. Conflicts with an `externalFixedIp` argument.
	ExternalSubnetIds pulumi.StringArrayOutput `pulumi:"externalSubnetIds"`
	// A unique name for the router. Changing this
	// updates the `name` of an existing router.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router.
	Region pulumi.StringOutput `pulumi:"region"`
	// A set of string tags for the router.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional driver-specific options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions RouterVendorOptionsPtrOutput `pulumi:"vendorOptions"`
}

// NewRouter registers a new resource with the given unique name, arguments, and options.
func NewRouter(ctx *pulumi.Context,
	name string, args *RouterArgs, opts ...pulumi.ResourceOption) (*Router, error) {
	if args == nil {
		args = &RouterArgs{}
	}

	var resource Router
	err := ctx.RegisterResource("openstack:networking/router:Router", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouter gets an existing Router resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterState, opts ...pulumi.ResourceOption) (*Router, error) {
	var resource Router
	err := ctx.ReadResource("openstack:networking/router:Router", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Router resources.
type routerState struct {
	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing router.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The collection of tags assigned on the router, which have been
	// explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new router.
	AvailabilityZoneHints []string `pulumi:"availabilityZoneHints"`
	// Human-readable description for the router.
	Description *string `pulumi:"description"`
	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	Distributed *bool `pulumi:"distributed"`
	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An `externalNetworkId` has to be set in order to
	// set this property. Changing this updates the `enableSnat` of the router.
	// Setting this value **requires** an **ext-gw-mode** extension to be enabled
	// in OpenStack Neutron.
	EnableSnat *bool `pulumi:"enableSnat"`
	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An `externalNetworkId`
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	ExternalFixedIps []RouterExternalFixedIp `pulumi:"externalFixedIps"`
	// The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	//
	// Deprecated: use external_network_id instead
	ExternalGateway *string `pulumi:"externalGateway"`
	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	ExternalNetworkId *string `pulumi:"externalNetworkId"`
	// A list of external subnet IDs to try over
	// each to obtain a fixed IP for the router. If a subnet ID in a list has
	// exhausted floating IP pool, the next subnet ID will be tried. This argument is
	// used only during the router creation and allows to set only one external fixed
	// IP. Conflicts with an `externalFixedIp` argument.
	ExternalSubnetIds []string `pulumi:"externalSubnetIds"`
	// A unique name for the router. Changing this
	// updates the `name` of an existing router.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router.
	Region *string `pulumi:"region"`
	// A set of string tags for the router.
	Tags []string `pulumi:"tags"`
	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional driver-specific options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions *RouterVendorOptions `pulumi:"vendorOptions"`
}

type RouterState struct {
	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing router.
	AdminStateUp pulumi.BoolPtrInput
	// The collection of tags assigned on the router, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new router.
	AvailabilityZoneHints pulumi.StringArrayInput
	// Human-readable description for the router.
	Description pulumi.StringPtrInput
	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	Distributed pulumi.BoolPtrInput
	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An `externalNetworkId` has to be set in order to
	// set this property. Changing this updates the `enableSnat` of the router.
	// Setting this value **requires** an **ext-gw-mode** extension to be enabled
	// in OpenStack Neutron.
	EnableSnat pulumi.BoolPtrInput
	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An `externalNetworkId`
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	ExternalFixedIps RouterExternalFixedIpArrayInput
	// The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	//
	// Deprecated: use external_network_id instead
	ExternalGateway pulumi.StringPtrInput
	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	ExternalNetworkId pulumi.StringPtrInput
	// A list of external subnet IDs to try over
	// each to obtain a fixed IP for the router. If a subnet ID in a list has
	// exhausted floating IP pool, the next subnet ID will be tried. This argument is
	// used only during the router creation and allows to set only one external fixed
	// IP. Conflicts with an `externalFixedIp` argument.
	ExternalSubnetIds pulumi.StringArrayInput
	// A unique name for the router. Changing this
	// updates the `name` of an existing router.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router.
	Region pulumi.StringPtrInput
	// A set of string tags for the router.
	Tags pulumi.StringArrayInput
	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	TenantId pulumi.StringPtrInput
	// Map of additional driver-specific options.
	ValueSpecs pulumi.MapInput
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions RouterVendorOptionsPtrInput
}

func (RouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerState)(nil)).Elem()
}

type routerArgs struct {
	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing router.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new router.
	AvailabilityZoneHints []string `pulumi:"availabilityZoneHints"`
	// Human-readable description for the router.
	Description *string `pulumi:"description"`
	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	Distributed *bool `pulumi:"distributed"`
	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An `externalNetworkId` has to be set in order to
	// set this property. Changing this updates the `enableSnat` of the router.
	// Setting this value **requires** an **ext-gw-mode** extension to be enabled
	// in OpenStack Neutron.
	EnableSnat *bool `pulumi:"enableSnat"`
	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An `externalNetworkId`
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	ExternalFixedIps []RouterExternalFixedIp `pulumi:"externalFixedIps"`
	// The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	//
	// Deprecated: use external_network_id instead
	ExternalGateway *string `pulumi:"externalGateway"`
	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	ExternalNetworkId *string `pulumi:"externalNetworkId"`
	// A list of external subnet IDs to try over
	// each to obtain a fixed IP for the router. If a subnet ID in a list has
	// exhausted floating IP pool, the next subnet ID will be tried. This argument is
	// used only during the router creation and allows to set only one external fixed
	// IP. Conflicts with an `externalFixedIp` argument.
	ExternalSubnetIds []string `pulumi:"externalSubnetIds"`
	// A unique name for the router. Changing this
	// updates the `name` of an existing router.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router.
	Region *string `pulumi:"region"`
	// A set of string tags for the router.
	Tags []string `pulumi:"tags"`
	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional driver-specific options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions *RouterVendorOptions `pulumi:"vendorOptions"`
}

// The set of arguments for constructing a Router resource.
type RouterArgs struct {
	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing router.
	AdminStateUp pulumi.BoolPtrInput
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new router.
	AvailabilityZoneHints pulumi.StringArrayInput
	// Human-readable description for the router.
	Description pulumi.StringPtrInput
	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	Distributed pulumi.BoolPtrInput
	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An `externalNetworkId` has to be set in order to
	// set this property. Changing this updates the `enableSnat` of the router.
	// Setting this value **requires** an **ext-gw-mode** extension to be enabled
	// in OpenStack Neutron.
	EnableSnat pulumi.BoolPtrInput
	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An `externalNetworkId`
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	ExternalFixedIps RouterExternalFixedIpArrayInput
	// The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	//
	// Deprecated: use external_network_id instead
	ExternalGateway pulumi.StringPtrInput
	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	ExternalNetworkId pulumi.StringPtrInput
	// A list of external subnet IDs to try over
	// each to obtain a fixed IP for the router. If a subnet ID in a list has
	// exhausted floating IP pool, the next subnet ID will be tried. This argument is
	// used only during the router creation and allows to set only one external fixed
	// IP. Conflicts with an `externalFixedIp` argument.
	ExternalSubnetIds pulumi.StringArrayInput
	// A unique name for the router. Changing this
	// updates the `name` of an existing router.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// router.
	Region pulumi.StringPtrInput
	// A set of string tags for the router.
	Tags pulumi.StringArrayInput
	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	TenantId pulumi.StringPtrInput
	// Map of additional driver-specific options.
	ValueSpecs pulumi.MapInput
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions RouterVendorOptionsPtrInput
}

func (RouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerArgs)(nil)).Elem()
}

type RouterInput interface {
	pulumi.Input

	ToRouterOutput() RouterOutput
	ToRouterOutputWithContext(ctx context.Context) RouterOutput
}

func (*Router) ElementType() reflect.Type {
	return reflect.TypeOf((*Router)(nil))
}

func (i *Router) ToRouterOutput() RouterOutput {
	return i.ToRouterOutputWithContext(context.Background())
}

func (i *Router) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterOutput)
}

func (i *Router) ToRouterPtrOutput() RouterPtrOutput {
	return i.ToRouterPtrOutputWithContext(context.Background())
}

func (i *Router) ToRouterPtrOutputWithContext(ctx context.Context) RouterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPtrOutput)
}

type RouterPtrInput interface {
	pulumi.Input

	ToRouterPtrOutput() RouterPtrOutput
	ToRouterPtrOutputWithContext(ctx context.Context) RouterPtrOutput
}

type routerPtrType RouterArgs

func (*routerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Router)(nil))
}

func (i *routerPtrType) ToRouterPtrOutput() RouterPtrOutput {
	return i.ToRouterPtrOutputWithContext(context.Background())
}

func (i *routerPtrType) ToRouterPtrOutputWithContext(ctx context.Context) RouterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPtrOutput)
}

// RouterArrayInput is an input type that accepts RouterArray and RouterArrayOutput values.
// You can construct a concrete instance of `RouterArrayInput` via:
//
//          RouterArray{ RouterArgs{...} }
type RouterArrayInput interface {
	pulumi.Input

	ToRouterArrayOutput() RouterArrayOutput
	ToRouterArrayOutputWithContext(context.Context) RouterArrayOutput
}

type RouterArray []RouterInput

func (RouterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Router)(nil)).Elem()
}

func (i RouterArray) ToRouterArrayOutput() RouterArrayOutput {
	return i.ToRouterArrayOutputWithContext(context.Background())
}

func (i RouterArray) ToRouterArrayOutputWithContext(ctx context.Context) RouterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterArrayOutput)
}

// RouterMapInput is an input type that accepts RouterMap and RouterMapOutput values.
// You can construct a concrete instance of `RouterMapInput` via:
//
//          RouterMap{ "key": RouterArgs{...} }
type RouterMapInput interface {
	pulumi.Input

	ToRouterMapOutput() RouterMapOutput
	ToRouterMapOutputWithContext(context.Context) RouterMapOutput
}

type RouterMap map[string]RouterInput

func (RouterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Router)(nil)).Elem()
}

func (i RouterMap) ToRouterMapOutput() RouterMapOutput {
	return i.ToRouterMapOutputWithContext(context.Background())
}

func (i RouterMap) ToRouterMapOutputWithContext(ctx context.Context) RouterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMapOutput)
}

type RouterOutput struct{ *pulumi.OutputState }

func (RouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Router)(nil))
}

func (o RouterOutput) ToRouterOutput() RouterOutput {
	return o
}

func (o RouterOutput) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return o
}

func (o RouterOutput) ToRouterPtrOutput() RouterPtrOutput {
	return o.ToRouterPtrOutputWithContext(context.Background())
}

func (o RouterOutput) ToRouterPtrOutputWithContext(ctx context.Context) RouterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Router) *Router {
		return &v
	}).(RouterPtrOutput)
}

type RouterPtrOutput struct{ *pulumi.OutputState }

func (RouterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Router)(nil))
}

func (o RouterPtrOutput) ToRouterPtrOutput() RouterPtrOutput {
	return o
}

func (o RouterPtrOutput) ToRouterPtrOutputWithContext(ctx context.Context) RouterPtrOutput {
	return o
}

func (o RouterPtrOutput) Elem() RouterOutput {
	return o.ApplyT(func(v *Router) Router {
		if v != nil {
			return *v
		}
		var ret Router
		return ret
	}).(RouterOutput)
}

type RouterArrayOutput struct{ *pulumi.OutputState }

func (RouterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Router)(nil))
}

func (o RouterArrayOutput) ToRouterArrayOutput() RouterArrayOutput {
	return o
}

func (o RouterArrayOutput) ToRouterArrayOutputWithContext(ctx context.Context) RouterArrayOutput {
	return o
}

func (o RouterArrayOutput) Index(i pulumi.IntInput) RouterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Router {
		return vs[0].([]Router)[vs[1].(int)]
	}).(RouterOutput)
}

type RouterMapOutput struct{ *pulumi.OutputState }

func (RouterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Router)(nil))
}

func (o RouterMapOutput) ToRouterMapOutput() RouterMapOutput {
	return o
}

func (o RouterMapOutput) ToRouterMapOutputWithContext(ctx context.Context) RouterMapOutput {
	return o
}

func (o RouterMapOutput) MapIndex(k pulumi.StringInput) RouterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Router {
		return vs[0].(map[string]Router)[vs[1].(string)]
	}).(RouterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInput)(nil)).Elem(), &Router{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterPtrInput)(nil)).Elem(), &Router{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterArrayInput)(nil)).Elem(), RouterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterMapInput)(nil)).Elem(), RouterMap{})
	pulumi.RegisterOutputType(RouterOutput{})
	pulumi.RegisterOutputType(RouterPtrOutput{})
	pulumi.RegisterOutputType(RouterArrayOutput{})
	pulumi.RegisterOutputType(RouterMapOutput{})
}
