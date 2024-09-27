// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openstack

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 BGP VPN port association resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := openstack.NewBgpvpnPortAssociateV2(ctx, "association_1", &openstack.BgpvpnPortAssociateV2Args{
//				BgpvpnId: pulumi.String("19382ec5-8098-47d9-a9c6-6270c91103f4"),
//				PortId:   pulumi.String("b83a95b8-c2c8-4eac-9a9e-ddc85bd1266f"),
//				Routes: openstack.BgpvpnPortAssociateV2RouteArray{
//					&openstack.BgpvpnPortAssociateV2RouteArgs{
//						Type:   pulumi.String("prefix"),
//						Prefix: pulumi.String("192.168.170.1/32"),
//					},
//					&openstack.BgpvpnPortAssociateV2RouteArgs{
//						Type:     pulumi.String("bgpvpn"),
//						BgpvpnId: pulumi.String("35af1cc6-3d0f-4c5d-86f8-8cdb508d3f0c"),
//					},
//				},
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
// # BGP VPN port associations can be imported using the BGP VPN ID and BGP VPN port
//
// association ID separated by a slash, e.g.:
//
// hcl
//
// ```sh
// $ pulumi import openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2 association_1 5bb44ecf-f8fe-4d75-8fc5-313f96ee2696/8f8fc660-3f28-414e-896a-0c7c51162fcf
// ```
type BgpvpnPortAssociateV2 struct {
	pulumi.CustomResourceState

	// A boolean flag indicating whether fixed
	// IPs should be advertised. Defaults to true.
	AdvertiseFixedIps pulumi.BoolOutput `pulumi:"advertiseFixedIps"`
	// The ID of the BGP VPN to which the port will be
	// associated. Changing this creates a new BGP VPN port association.
	BgpvpnId pulumi.StringOutput `pulumi:"bgpvpnId"`
	// The ID of the port to be associated with the BGP VPN.
	// Changing this creates a new BGP VPN port association.
	PortId pulumi.StringOutput `pulumi:"portId"`
	// The ID of the project that owns the port
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN port
	// association.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN port association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN port association.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list of dictionaries containing the following keys:
	Routes BgpvpnPortAssociateV2RouteArrayOutput `pulumi:"routes"`
}

// NewBgpvpnPortAssociateV2 registers a new resource with the given unique name, arguments, and options.
func NewBgpvpnPortAssociateV2(ctx *pulumi.Context,
	name string, args *BgpvpnPortAssociateV2Args, opts ...pulumi.ResourceOption) (*BgpvpnPortAssociateV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BgpvpnId == nil {
		return nil, errors.New("invalid value for required argument 'BgpvpnId'")
	}
	if args.PortId == nil {
		return nil, errors.New("invalid value for required argument 'PortId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BgpvpnPortAssociateV2
	err := ctx.RegisterResource("openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpvpnPortAssociateV2 gets an existing BgpvpnPortAssociateV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpvpnPortAssociateV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpvpnPortAssociateV2State, opts ...pulumi.ResourceOption) (*BgpvpnPortAssociateV2, error) {
	var resource BgpvpnPortAssociateV2
	err := ctx.ReadResource("openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpvpnPortAssociateV2 resources.
type bgpvpnPortAssociateV2State struct {
	// A boolean flag indicating whether fixed
	// IPs should be advertised. Defaults to true.
	AdvertiseFixedIps *bool `pulumi:"advertiseFixedIps"`
	// The ID of the BGP VPN to which the port will be
	// associated. Changing this creates a new BGP VPN port association.
	BgpvpnId *string `pulumi:"bgpvpnId"`
	// The ID of the port to be associated with the BGP VPN.
	// Changing this creates a new BGP VPN port association.
	PortId *string `pulumi:"portId"`
	// The ID of the project that owns the port
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN port
	// association.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN port association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN port association.
	Region *string `pulumi:"region"`
	// A list of dictionaries containing the following keys:
	Routes []BgpvpnPortAssociateV2Route `pulumi:"routes"`
}

type BgpvpnPortAssociateV2State struct {
	// A boolean flag indicating whether fixed
	// IPs should be advertised. Defaults to true.
	AdvertiseFixedIps pulumi.BoolPtrInput
	// The ID of the BGP VPN to which the port will be
	// associated. Changing this creates a new BGP VPN port association.
	BgpvpnId pulumi.StringPtrInput
	// The ID of the port to be associated with the BGP VPN.
	// Changing this creates a new BGP VPN port association.
	PortId pulumi.StringPtrInput
	// The ID of the project that owns the port
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN port
	// association.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN port association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN port association.
	Region pulumi.StringPtrInput
	// A list of dictionaries containing the following keys:
	Routes BgpvpnPortAssociateV2RouteArrayInput
}

func (BgpvpnPortAssociateV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpvpnPortAssociateV2State)(nil)).Elem()
}

type bgpvpnPortAssociateV2Args struct {
	// A boolean flag indicating whether fixed
	// IPs should be advertised. Defaults to true.
	AdvertiseFixedIps *bool `pulumi:"advertiseFixedIps"`
	// The ID of the BGP VPN to which the port will be
	// associated. Changing this creates a new BGP VPN port association.
	BgpvpnId string `pulumi:"bgpvpnId"`
	// The ID of the port to be associated with the BGP VPN.
	// Changing this creates a new BGP VPN port association.
	PortId string `pulumi:"portId"`
	// The ID of the project that owns the port
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN port
	// association.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN port association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN port association.
	Region *string `pulumi:"region"`
	// A list of dictionaries containing the following keys:
	Routes []BgpvpnPortAssociateV2Route `pulumi:"routes"`
}

// The set of arguments for constructing a BgpvpnPortAssociateV2 resource.
type BgpvpnPortAssociateV2Args struct {
	// A boolean flag indicating whether fixed
	// IPs should be advertised. Defaults to true.
	AdvertiseFixedIps pulumi.BoolPtrInput
	// The ID of the BGP VPN to which the port will be
	// associated. Changing this creates a new BGP VPN port association.
	BgpvpnId pulumi.StringInput
	// The ID of the port to be associated with the BGP VPN.
	// Changing this creates a new BGP VPN port association.
	PortId pulumi.StringInput
	// The ID of the project that owns the port
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN port
	// association.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN port association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN port association.
	Region pulumi.StringPtrInput
	// A list of dictionaries containing the following keys:
	Routes BgpvpnPortAssociateV2RouteArrayInput
}

func (BgpvpnPortAssociateV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpvpnPortAssociateV2Args)(nil)).Elem()
}

type BgpvpnPortAssociateV2Input interface {
	pulumi.Input

	ToBgpvpnPortAssociateV2Output() BgpvpnPortAssociateV2Output
	ToBgpvpnPortAssociateV2OutputWithContext(ctx context.Context) BgpvpnPortAssociateV2Output
}

func (*BgpvpnPortAssociateV2) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpvpnPortAssociateV2)(nil)).Elem()
}

func (i *BgpvpnPortAssociateV2) ToBgpvpnPortAssociateV2Output() BgpvpnPortAssociateV2Output {
	return i.ToBgpvpnPortAssociateV2OutputWithContext(context.Background())
}

func (i *BgpvpnPortAssociateV2) ToBgpvpnPortAssociateV2OutputWithContext(ctx context.Context) BgpvpnPortAssociateV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(BgpvpnPortAssociateV2Output)
}

// BgpvpnPortAssociateV2ArrayInput is an input type that accepts BgpvpnPortAssociateV2Array and BgpvpnPortAssociateV2ArrayOutput values.
// You can construct a concrete instance of `BgpvpnPortAssociateV2ArrayInput` via:
//
//	BgpvpnPortAssociateV2Array{ BgpvpnPortAssociateV2Args{...} }
type BgpvpnPortAssociateV2ArrayInput interface {
	pulumi.Input

	ToBgpvpnPortAssociateV2ArrayOutput() BgpvpnPortAssociateV2ArrayOutput
	ToBgpvpnPortAssociateV2ArrayOutputWithContext(context.Context) BgpvpnPortAssociateV2ArrayOutput
}

type BgpvpnPortAssociateV2Array []BgpvpnPortAssociateV2Input

func (BgpvpnPortAssociateV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpvpnPortAssociateV2)(nil)).Elem()
}

func (i BgpvpnPortAssociateV2Array) ToBgpvpnPortAssociateV2ArrayOutput() BgpvpnPortAssociateV2ArrayOutput {
	return i.ToBgpvpnPortAssociateV2ArrayOutputWithContext(context.Background())
}

func (i BgpvpnPortAssociateV2Array) ToBgpvpnPortAssociateV2ArrayOutputWithContext(ctx context.Context) BgpvpnPortAssociateV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpvpnPortAssociateV2ArrayOutput)
}

// BgpvpnPortAssociateV2MapInput is an input type that accepts BgpvpnPortAssociateV2Map and BgpvpnPortAssociateV2MapOutput values.
// You can construct a concrete instance of `BgpvpnPortAssociateV2MapInput` via:
//
//	BgpvpnPortAssociateV2Map{ "key": BgpvpnPortAssociateV2Args{...} }
type BgpvpnPortAssociateV2MapInput interface {
	pulumi.Input

	ToBgpvpnPortAssociateV2MapOutput() BgpvpnPortAssociateV2MapOutput
	ToBgpvpnPortAssociateV2MapOutputWithContext(context.Context) BgpvpnPortAssociateV2MapOutput
}

type BgpvpnPortAssociateV2Map map[string]BgpvpnPortAssociateV2Input

func (BgpvpnPortAssociateV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpvpnPortAssociateV2)(nil)).Elem()
}

func (i BgpvpnPortAssociateV2Map) ToBgpvpnPortAssociateV2MapOutput() BgpvpnPortAssociateV2MapOutput {
	return i.ToBgpvpnPortAssociateV2MapOutputWithContext(context.Background())
}

func (i BgpvpnPortAssociateV2Map) ToBgpvpnPortAssociateV2MapOutputWithContext(ctx context.Context) BgpvpnPortAssociateV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpvpnPortAssociateV2MapOutput)
}

type BgpvpnPortAssociateV2Output struct{ *pulumi.OutputState }

func (BgpvpnPortAssociateV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpvpnPortAssociateV2)(nil)).Elem()
}

func (o BgpvpnPortAssociateV2Output) ToBgpvpnPortAssociateV2Output() BgpvpnPortAssociateV2Output {
	return o
}

func (o BgpvpnPortAssociateV2Output) ToBgpvpnPortAssociateV2OutputWithContext(ctx context.Context) BgpvpnPortAssociateV2Output {
	return o
}

// A boolean flag indicating whether fixed
// IPs should be advertised. Defaults to true.
func (o BgpvpnPortAssociateV2Output) AdvertiseFixedIps() pulumi.BoolOutput {
	return o.ApplyT(func(v *BgpvpnPortAssociateV2) pulumi.BoolOutput { return v.AdvertiseFixedIps }).(pulumi.BoolOutput)
}

// The ID of the BGP VPN to which the port will be
// associated. Changing this creates a new BGP VPN port association.
func (o BgpvpnPortAssociateV2Output) BgpvpnId() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpvpnPortAssociateV2) pulumi.StringOutput { return v.BgpvpnId }).(pulumi.StringOutput)
}

// The ID of the port to be associated with the BGP VPN.
// Changing this creates a new BGP VPN port association.
func (o BgpvpnPortAssociateV2Output) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpvpnPortAssociateV2) pulumi.StringOutput { return v.PortId }).(pulumi.StringOutput)
}

// The ID of the project that owns the port
// association. Only administrative and users with `advsvc` role can specify a
// project ID other than their own. Changing this creates a new BGP VPN port
// association.
func (o BgpvpnPortAssociateV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpvpnPortAssociateV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a BGP VPN port association. If
// omitted, the `region` argument of the provider is used. Changing this creates
// a new BGP VPN port association.
func (o BgpvpnPortAssociateV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpvpnPortAssociateV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A list of dictionaries containing the following keys:
func (o BgpvpnPortAssociateV2Output) Routes() BgpvpnPortAssociateV2RouteArrayOutput {
	return o.ApplyT(func(v *BgpvpnPortAssociateV2) BgpvpnPortAssociateV2RouteArrayOutput { return v.Routes }).(BgpvpnPortAssociateV2RouteArrayOutput)
}

type BgpvpnPortAssociateV2ArrayOutput struct{ *pulumi.OutputState }

func (BgpvpnPortAssociateV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpvpnPortAssociateV2)(nil)).Elem()
}

func (o BgpvpnPortAssociateV2ArrayOutput) ToBgpvpnPortAssociateV2ArrayOutput() BgpvpnPortAssociateV2ArrayOutput {
	return o
}

func (o BgpvpnPortAssociateV2ArrayOutput) ToBgpvpnPortAssociateV2ArrayOutputWithContext(ctx context.Context) BgpvpnPortAssociateV2ArrayOutput {
	return o
}

func (o BgpvpnPortAssociateV2ArrayOutput) Index(i pulumi.IntInput) BgpvpnPortAssociateV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BgpvpnPortAssociateV2 {
		return vs[0].([]*BgpvpnPortAssociateV2)[vs[1].(int)]
	}).(BgpvpnPortAssociateV2Output)
}

type BgpvpnPortAssociateV2MapOutput struct{ *pulumi.OutputState }

func (BgpvpnPortAssociateV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpvpnPortAssociateV2)(nil)).Elem()
}

func (o BgpvpnPortAssociateV2MapOutput) ToBgpvpnPortAssociateV2MapOutput() BgpvpnPortAssociateV2MapOutput {
	return o
}

func (o BgpvpnPortAssociateV2MapOutput) ToBgpvpnPortAssociateV2MapOutputWithContext(ctx context.Context) BgpvpnPortAssociateV2MapOutput {
	return o
}

func (o BgpvpnPortAssociateV2MapOutput) MapIndex(k pulumi.StringInput) BgpvpnPortAssociateV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BgpvpnPortAssociateV2 {
		return vs[0].(map[string]*BgpvpnPortAssociateV2)[vs[1].(string)]
	}).(BgpvpnPortAssociateV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpvpnPortAssociateV2Input)(nil)).Elem(), &BgpvpnPortAssociateV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpvpnPortAssociateV2ArrayInput)(nil)).Elem(), BgpvpnPortAssociateV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpvpnPortAssociateV2MapInput)(nil)).Elem(), BgpvpnPortAssociateV2Map{})
	pulumi.RegisterOutputType(BgpvpnPortAssociateV2Output{})
	pulumi.RegisterOutputType(BgpvpnPortAssociateV2ArrayOutput{})
	pulumi.RegisterOutputType(BgpvpnPortAssociateV2MapOutput{})
}
