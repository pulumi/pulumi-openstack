// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openstack

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 BGP VPN router association resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := openstack.NewBgpvpnRouterAssociateV2(ctx, "association_1", &openstack.BgpvpnRouterAssociateV2Args{
//				BgpvpnId: pulumi.String("d57d39e1-dc63-44fd-8cbd-a4e1488100c5"),
//				RouterId: pulumi.String("423fa80f-e0d7-4d02-a9a5-8b8c05812bf6"),
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
// # BGP VPN router associations can be imported using the BGP VPN ID and BGP VPN
//
// router association ID separated by a slash, e.g.:
//
// hcl
//
// ```sh
// $ pulumi import openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2 association_1 e26d509e-fc2d-4fb5-8562-619911a9a6bc/3cc9df2d-80db-4536-8ba6-295d1d0f723f
// ```
type BgpvpnRouterAssociateV2 struct {
	pulumi.CustomResourceState

	// A boolean flag indicating whether extra
	// routes should be advertised. Defaults to true.
	AdvertiseExtraRoutes pulumi.BoolOutput `pulumi:"advertiseExtraRoutes"`
	// The ID of the BGP VPN to which the router will be
	// associated. Changing this creates a new BGP VPN router association.
	BgpvpnId pulumi.StringOutput `pulumi:"bgpvpnId"`
	// The ID of the project that owns the BGP VPN router
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN router
	// association.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN router association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN router association.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of the router to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN router association.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
}

// NewBgpvpnRouterAssociateV2 registers a new resource with the given unique name, arguments, and options.
func NewBgpvpnRouterAssociateV2(ctx *pulumi.Context,
	name string, args *BgpvpnRouterAssociateV2Args, opts ...pulumi.ResourceOption) (*BgpvpnRouterAssociateV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BgpvpnId == nil {
		return nil, errors.New("invalid value for required argument 'BgpvpnId'")
	}
	if args.RouterId == nil {
		return nil, errors.New("invalid value for required argument 'RouterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BgpvpnRouterAssociateV2
	err := ctx.RegisterResource("openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpvpnRouterAssociateV2 gets an existing BgpvpnRouterAssociateV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpvpnRouterAssociateV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpvpnRouterAssociateV2State, opts ...pulumi.ResourceOption) (*BgpvpnRouterAssociateV2, error) {
	var resource BgpvpnRouterAssociateV2
	err := ctx.ReadResource("openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpvpnRouterAssociateV2 resources.
type bgpvpnRouterAssociateV2State struct {
	// A boolean flag indicating whether extra
	// routes should be advertised. Defaults to true.
	AdvertiseExtraRoutes *bool `pulumi:"advertiseExtraRoutes"`
	// The ID of the BGP VPN to which the router will be
	// associated. Changing this creates a new BGP VPN router association.
	BgpvpnId *string `pulumi:"bgpvpnId"`
	// The ID of the project that owns the BGP VPN router
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN router
	// association.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN router association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN router association.
	Region *string `pulumi:"region"`
	// The ID of the router to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN router association.
	RouterId *string `pulumi:"routerId"`
}

type BgpvpnRouterAssociateV2State struct {
	// A boolean flag indicating whether extra
	// routes should be advertised. Defaults to true.
	AdvertiseExtraRoutes pulumi.BoolPtrInput
	// The ID of the BGP VPN to which the router will be
	// associated. Changing this creates a new BGP VPN router association.
	BgpvpnId pulumi.StringPtrInput
	// The ID of the project that owns the BGP VPN router
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN router
	// association.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN router association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN router association.
	Region pulumi.StringPtrInput
	// The ID of the router to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN router association.
	RouterId pulumi.StringPtrInput
}

func (BgpvpnRouterAssociateV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpvpnRouterAssociateV2State)(nil)).Elem()
}

type bgpvpnRouterAssociateV2Args struct {
	// A boolean flag indicating whether extra
	// routes should be advertised. Defaults to true.
	AdvertiseExtraRoutes *bool `pulumi:"advertiseExtraRoutes"`
	// The ID of the BGP VPN to which the router will be
	// associated. Changing this creates a new BGP VPN router association.
	BgpvpnId string `pulumi:"bgpvpnId"`
	// The ID of the project that owns the BGP VPN router
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN router
	// association.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN router association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN router association.
	Region *string `pulumi:"region"`
	// The ID of the router to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN router association.
	RouterId string `pulumi:"routerId"`
}

// The set of arguments for constructing a BgpvpnRouterAssociateV2 resource.
type BgpvpnRouterAssociateV2Args struct {
	// A boolean flag indicating whether extra
	// routes should be advertised. Defaults to true.
	AdvertiseExtraRoutes pulumi.BoolPtrInput
	// The ID of the BGP VPN to which the router will be
	// associated. Changing this creates a new BGP VPN router association.
	BgpvpnId pulumi.StringInput
	// The ID of the project that owns the BGP VPN router
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN router
	// association.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN router association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN router association.
	Region pulumi.StringPtrInput
	// The ID of the router to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN router association.
	RouterId pulumi.StringInput
}

func (BgpvpnRouterAssociateV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpvpnRouterAssociateV2Args)(nil)).Elem()
}

type BgpvpnRouterAssociateV2Input interface {
	pulumi.Input

	ToBgpvpnRouterAssociateV2Output() BgpvpnRouterAssociateV2Output
	ToBgpvpnRouterAssociateV2OutputWithContext(ctx context.Context) BgpvpnRouterAssociateV2Output
}

func (*BgpvpnRouterAssociateV2) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpvpnRouterAssociateV2)(nil)).Elem()
}

func (i *BgpvpnRouterAssociateV2) ToBgpvpnRouterAssociateV2Output() BgpvpnRouterAssociateV2Output {
	return i.ToBgpvpnRouterAssociateV2OutputWithContext(context.Background())
}

func (i *BgpvpnRouterAssociateV2) ToBgpvpnRouterAssociateV2OutputWithContext(ctx context.Context) BgpvpnRouterAssociateV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(BgpvpnRouterAssociateV2Output)
}

// BgpvpnRouterAssociateV2ArrayInput is an input type that accepts BgpvpnRouterAssociateV2Array and BgpvpnRouterAssociateV2ArrayOutput values.
// You can construct a concrete instance of `BgpvpnRouterAssociateV2ArrayInput` via:
//
//	BgpvpnRouterAssociateV2Array{ BgpvpnRouterAssociateV2Args{...} }
type BgpvpnRouterAssociateV2ArrayInput interface {
	pulumi.Input

	ToBgpvpnRouterAssociateV2ArrayOutput() BgpvpnRouterAssociateV2ArrayOutput
	ToBgpvpnRouterAssociateV2ArrayOutputWithContext(context.Context) BgpvpnRouterAssociateV2ArrayOutput
}

type BgpvpnRouterAssociateV2Array []BgpvpnRouterAssociateV2Input

func (BgpvpnRouterAssociateV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpvpnRouterAssociateV2)(nil)).Elem()
}

func (i BgpvpnRouterAssociateV2Array) ToBgpvpnRouterAssociateV2ArrayOutput() BgpvpnRouterAssociateV2ArrayOutput {
	return i.ToBgpvpnRouterAssociateV2ArrayOutputWithContext(context.Background())
}

func (i BgpvpnRouterAssociateV2Array) ToBgpvpnRouterAssociateV2ArrayOutputWithContext(ctx context.Context) BgpvpnRouterAssociateV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpvpnRouterAssociateV2ArrayOutput)
}

// BgpvpnRouterAssociateV2MapInput is an input type that accepts BgpvpnRouterAssociateV2Map and BgpvpnRouterAssociateV2MapOutput values.
// You can construct a concrete instance of `BgpvpnRouterAssociateV2MapInput` via:
//
//	BgpvpnRouterAssociateV2Map{ "key": BgpvpnRouterAssociateV2Args{...} }
type BgpvpnRouterAssociateV2MapInput interface {
	pulumi.Input

	ToBgpvpnRouterAssociateV2MapOutput() BgpvpnRouterAssociateV2MapOutput
	ToBgpvpnRouterAssociateV2MapOutputWithContext(context.Context) BgpvpnRouterAssociateV2MapOutput
}

type BgpvpnRouterAssociateV2Map map[string]BgpvpnRouterAssociateV2Input

func (BgpvpnRouterAssociateV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpvpnRouterAssociateV2)(nil)).Elem()
}

func (i BgpvpnRouterAssociateV2Map) ToBgpvpnRouterAssociateV2MapOutput() BgpvpnRouterAssociateV2MapOutput {
	return i.ToBgpvpnRouterAssociateV2MapOutputWithContext(context.Background())
}

func (i BgpvpnRouterAssociateV2Map) ToBgpvpnRouterAssociateV2MapOutputWithContext(ctx context.Context) BgpvpnRouterAssociateV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpvpnRouterAssociateV2MapOutput)
}

type BgpvpnRouterAssociateV2Output struct{ *pulumi.OutputState }

func (BgpvpnRouterAssociateV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpvpnRouterAssociateV2)(nil)).Elem()
}

func (o BgpvpnRouterAssociateV2Output) ToBgpvpnRouterAssociateV2Output() BgpvpnRouterAssociateV2Output {
	return o
}

func (o BgpvpnRouterAssociateV2Output) ToBgpvpnRouterAssociateV2OutputWithContext(ctx context.Context) BgpvpnRouterAssociateV2Output {
	return o
}

// A boolean flag indicating whether extra
// routes should be advertised. Defaults to true.
func (o BgpvpnRouterAssociateV2Output) AdvertiseExtraRoutes() pulumi.BoolOutput {
	return o.ApplyT(func(v *BgpvpnRouterAssociateV2) pulumi.BoolOutput { return v.AdvertiseExtraRoutes }).(pulumi.BoolOutput)
}

// The ID of the BGP VPN to which the router will be
// associated. Changing this creates a new BGP VPN router association.
func (o BgpvpnRouterAssociateV2Output) BgpvpnId() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpvpnRouterAssociateV2) pulumi.StringOutput { return v.BgpvpnId }).(pulumi.StringOutput)
}

// The ID of the project that owns the BGP VPN router
// association. Only administrative and users with `advsvc` role can specify a
// project ID other than their own. Changing this creates a new BGP VPN router
// association.
func (o BgpvpnRouterAssociateV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpvpnRouterAssociateV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a BGP VPN router association. If
// omitted, the `region` argument of the provider is used. Changing this creates
// a new BGP VPN router association.
func (o BgpvpnRouterAssociateV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpvpnRouterAssociateV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The ID of the router to be associated with the BGP
// VPN. Changing this creates a new BGP VPN router association.
func (o BgpvpnRouterAssociateV2Output) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpvpnRouterAssociateV2) pulumi.StringOutput { return v.RouterId }).(pulumi.StringOutput)
}

type BgpvpnRouterAssociateV2ArrayOutput struct{ *pulumi.OutputState }

func (BgpvpnRouterAssociateV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpvpnRouterAssociateV2)(nil)).Elem()
}

func (o BgpvpnRouterAssociateV2ArrayOutput) ToBgpvpnRouterAssociateV2ArrayOutput() BgpvpnRouterAssociateV2ArrayOutput {
	return o
}

func (o BgpvpnRouterAssociateV2ArrayOutput) ToBgpvpnRouterAssociateV2ArrayOutputWithContext(ctx context.Context) BgpvpnRouterAssociateV2ArrayOutput {
	return o
}

func (o BgpvpnRouterAssociateV2ArrayOutput) Index(i pulumi.IntInput) BgpvpnRouterAssociateV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BgpvpnRouterAssociateV2 {
		return vs[0].([]*BgpvpnRouterAssociateV2)[vs[1].(int)]
	}).(BgpvpnRouterAssociateV2Output)
}

type BgpvpnRouterAssociateV2MapOutput struct{ *pulumi.OutputState }

func (BgpvpnRouterAssociateV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpvpnRouterAssociateV2)(nil)).Elem()
}

func (o BgpvpnRouterAssociateV2MapOutput) ToBgpvpnRouterAssociateV2MapOutput() BgpvpnRouterAssociateV2MapOutput {
	return o
}

func (o BgpvpnRouterAssociateV2MapOutput) ToBgpvpnRouterAssociateV2MapOutputWithContext(ctx context.Context) BgpvpnRouterAssociateV2MapOutput {
	return o
}

func (o BgpvpnRouterAssociateV2MapOutput) MapIndex(k pulumi.StringInput) BgpvpnRouterAssociateV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BgpvpnRouterAssociateV2 {
		return vs[0].(map[string]*BgpvpnRouterAssociateV2)[vs[1].(string)]
	}).(BgpvpnRouterAssociateV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpvpnRouterAssociateV2Input)(nil)).Elem(), &BgpvpnRouterAssociateV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpvpnRouterAssociateV2ArrayInput)(nil)).Elem(), BgpvpnRouterAssociateV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpvpnRouterAssociateV2MapInput)(nil)).Elem(), BgpvpnRouterAssociateV2Map{})
	pulumi.RegisterOutputType(BgpvpnRouterAssociateV2Output{})
	pulumi.RegisterOutputType(BgpvpnRouterAssociateV2ArrayOutput{})
	pulumi.RegisterOutputType(BgpvpnRouterAssociateV2MapOutput{})
}
