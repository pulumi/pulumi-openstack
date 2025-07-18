// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bgpvpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 BGP VPN service resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/bgpvpn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bgpvpn.NewV2(ctx, "bgpvpn_1", &bgpvpn.V2Args{
//				Name: pulumi.String("bgpvpn1"),
//				RouteDistinguishers: pulumi.StringArray{
//					pulumi.String("64512:1"),
//				},
//				RouteTargets: pulumi.StringArray{
//					pulumi.String("64512:1"),
//				},
//				ImportTargets: pulumi.StringArray{
//					pulumi.String("64512:2"),
//				},
//				ExportTargets: pulumi.StringArray{
//					pulumi.String("64512:3"),
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
// BGP VPNs can be imported using the `id`, e.g.
//
// hcl
//
// ```sh
// $ pulumi import openstack:bgpvpn/v2:V2 bgpvpn_1 1eec2c66-6be2-4305-af3f-354c9b81f18c
// ```
type V2 struct {
	pulumi.CustomResourceState

	// A list of additional Route Targets that will be
	// used for export.
	ExportTargets pulumi.StringArrayOutput `pulumi:"exportTargets"`
	// A list of additional Route Targets that will be
	// imported.
	ImportTargets pulumi.StringArrayOutput `pulumi:"importTargets"`
	// The default BGP LOCAL\_PREF of routes that will be
	// advertised to the BGP VPN, unless overridden per-route.
	LocalPref pulumi.IntPtrOutput `pulumi:"localPref"`
	// The name of the BGP VPN. Changing this updates the name of
	// the existing BGP VPN.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of network IDs that are associated with the BGP VPN.
	Networks pulumi.StringArrayOutput `pulumi:"networks"`
	// A list of port IDs that are associated with the BGP VPN.
	Ports pulumi.StringArrayOutput `pulumi:"ports"`
	// The ID of the project that owns the BGPVPN. Only
	// administrative and users with `advsvc` role can specify a project ID other
	// than their own. Changing this creates a new BGP VPN.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// BGP VPN.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list of route distinguisher strings. If
	// specified, one of these RDs will be used to advertise VPN routes.
	RouteDistinguishers pulumi.StringArrayOutput `pulumi:"routeDistinguishers"`
	// A list of Route Targets that will be both
	// imported and used for export.
	RouteTargets pulumi.StringArrayOutput `pulumi:"routeTargets"`
	// A list of router IDs that are associated with the BGP VPN.
	Routers pulumi.StringArrayOutput `pulumi:"routers"`
	// Indicates whether the BGP VPN is shared across projects.
	Shared pulumi.BoolOutput `pulumi:"shared"`
	// The type of the BGP VPN (either `l2` or `l3`). Changing this
	// creates a new BGP VPN. Defaults to `l3`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The globally-assigned VXLAN VNI for the BGP VPN. Changing
	// this creates a new BGP VPN.
	Vni pulumi.IntPtrOutput `pulumi:"vni"`
}

// NewV2 registers a new resource with the given unique name, arguments, and options.
func NewV2(ctx *pulumi.Context,
	name string, args *V2Args, opts ...pulumi.ResourceOption) (*V2, error) {
	if args == nil {
		args = &V2Args{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("openstack:index/bgpvpnV2:BgpvpnV2"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource V2
	err := ctx.RegisterResource("openstack:bgpvpn/v2:V2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetV2 gets an existing V2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *V2State, opts ...pulumi.ResourceOption) (*V2, error) {
	var resource V2
	err := ctx.ReadResource("openstack:bgpvpn/v2:V2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering V2 resources.
type v2State struct {
	// A list of additional Route Targets that will be
	// used for export.
	ExportTargets []string `pulumi:"exportTargets"`
	// A list of additional Route Targets that will be
	// imported.
	ImportTargets []string `pulumi:"importTargets"`
	// The default BGP LOCAL\_PREF of routes that will be
	// advertised to the BGP VPN, unless overridden per-route.
	LocalPref *int `pulumi:"localPref"`
	// The name of the BGP VPN. Changing this updates the name of
	// the existing BGP VPN.
	Name *string `pulumi:"name"`
	// A list of network IDs that are associated with the BGP VPN.
	Networks []string `pulumi:"networks"`
	// A list of port IDs that are associated with the BGP VPN.
	Ports []string `pulumi:"ports"`
	// The ID of the project that owns the BGPVPN. Only
	// administrative and users with `advsvc` role can specify a project ID other
	// than their own. Changing this creates a new BGP VPN.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// BGP VPN.
	Region *string `pulumi:"region"`
	// A list of route distinguisher strings. If
	// specified, one of these RDs will be used to advertise VPN routes.
	RouteDistinguishers []string `pulumi:"routeDistinguishers"`
	// A list of Route Targets that will be both
	// imported and used for export.
	RouteTargets []string `pulumi:"routeTargets"`
	// A list of router IDs that are associated with the BGP VPN.
	Routers []string `pulumi:"routers"`
	// Indicates whether the BGP VPN is shared across projects.
	Shared *bool `pulumi:"shared"`
	// The type of the BGP VPN (either `l2` or `l3`). Changing this
	// creates a new BGP VPN. Defaults to `l3`.
	Type *string `pulumi:"type"`
	// The globally-assigned VXLAN VNI for the BGP VPN. Changing
	// this creates a new BGP VPN.
	Vni *int `pulumi:"vni"`
}

type V2State struct {
	// A list of additional Route Targets that will be
	// used for export.
	ExportTargets pulumi.StringArrayInput
	// A list of additional Route Targets that will be
	// imported.
	ImportTargets pulumi.StringArrayInput
	// The default BGP LOCAL\_PREF of routes that will be
	// advertised to the BGP VPN, unless overridden per-route.
	LocalPref pulumi.IntPtrInput
	// The name of the BGP VPN. Changing this updates the name of
	// the existing BGP VPN.
	Name pulumi.StringPtrInput
	// A list of network IDs that are associated with the BGP VPN.
	Networks pulumi.StringArrayInput
	// A list of port IDs that are associated with the BGP VPN.
	Ports pulumi.StringArrayInput
	// The ID of the project that owns the BGPVPN. Only
	// administrative and users with `advsvc` role can specify a project ID other
	// than their own. Changing this creates a new BGP VPN.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// BGP VPN.
	Region pulumi.StringPtrInput
	// A list of route distinguisher strings. If
	// specified, one of these RDs will be used to advertise VPN routes.
	RouteDistinguishers pulumi.StringArrayInput
	// A list of Route Targets that will be both
	// imported and used for export.
	RouteTargets pulumi.StringArrayInput
	// A list of router IDs that are associated with the BGP VPN.
	Routers pulumi.StringArrayInput
	// Indicates whether the BGP VPN is shared across projects.
	Shared pulumi.BoolPtrInput
	// The type of the BGP VPN (either `l2` or `l3`). Changing this
	// creates a new BGP VPN. Defaults to `l3`.
	Type pulumi.StringPtrInput
	// The globally-assigned VXLAN VNI for the BGP VPN. Changing
	// this creates a new BGP VPN.
	Vni pulumi.IntPtrInput
}

func (V2State) ElementType() reflect.Type {
	return reflect.TypeOf((*v2State)(nil)).Elem()
}

type v2Args struct {
	// A list of additional Route Targets that will be
	// used for export.
	ExportTargets []string `pulumi:"exportTargets"`
	// A list of additional Route Targets that will be
	// imported.
	ImportTargets []string `pulumi:"importTargets"`
	// The default BGP LOCAL\_PREF of routes that will be
	// advertised to the BGP VPN, unless overridden per-route.
	LocalPref *int `pulumi:"localPref"`
	// The name of the BGP VPN. Changing this updates the name of
	// the existing BGP VPN.
	Name *string `pulumi:"name"`
	// The ID of the project that owns the BGPVPN. Only
	// administrative and users with `advsvc` role can specify a project ID other
	// than their own. Changing this creates a new BGP VPN.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// BGP VPN.
	Region *string `pulumi:"region"`
	// A list of route distinguisher strings. If
	// specified, one of these RDs will be used to advertise VPN routes.
	RouteDistinguishers []string `pulumi:"routeDistinguishers"`
	// A list of Route Targets that will be both
	// imported and used for export.
	RouteTargets []string `pulumi:"routeTargets"`
	// The type of the BGP VPN (either `l2` or `l3`). Changing this
	// creates a new BGP VPN. Defaults to `l3`.
	Type *string `pulumi:"type"`
	// The globally-assigned VXLAN VNI for the BGP VPN. Changing
	// this creates a new BGP VPN.
	Vni *int `pulumi:"vni"`
}

// The set of arguments for constructing a V2 resource.
type V2Args struct {
	// A list of additional Route Targets that will be
	// used for export.
	ExportTargets pulumi.StringArrayInput
	// A list of additional Route Targets that will be
	// imported.
	ImportTargets pulumi.StringArrayInput
	// The default BGP LOCAL\_PREF of routes that will be
	// advertised to the BGP VPN, unless overridden per-route.
	LocalPref pulumi.IntPtrInput
	// The name of the BGP VPN. Changing this updates the name of
	// the existing BGP VPN.
	Name pulumi.StringPtrInput
	// The ID of the project that owns the BGPVPN. Only
	// administrative and users with `advsvc` role can specify a project ID other
	// than their own. Changing this creates a new BGP VPN.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// BGP VPN.
	Region pulumi.StringPtrInput
	// A list of route distinguisher strings. If
	// specified, one of these RDs will be used to advertise VPN routes.
	RouteDistinguishers pulumi.StringArrayInput
	// A list of Route Targets that will be both
	// imported and used for export.
	RouteTargets pulumi.StringArrayInput
	// The type of the BGP VPN (either `l2` or `l3`). Changing this
	// creates a new BGP VPN. Defaults to `l3`.
	Type pulumi.StringPtrInput
	// The globally-assigned VXLAN VNI for the BGP VPN. Changing
	// this creates a new BGP VPN.
	Vni pulumi.IntPtrInput
}

func (V2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*v2Args)(nil)).Elem()
}

type V2Input interface {
	pulumi.Input

	ToV2Output() V2Output
	ToV2OutputWithContext(ctx context.Context) V2Output
}

func (*V2) ElementType() reflect.Type {
	return reflect.TypeOf((**V2)(nil)).Elem()
}

func (i *V2) ToV2Output() V2Output {
	return i.ToV2OutputWithContext(context.Background())
}

func (i *V2) ToV2OutputWithContext(ctx context.Context) V2Output {
	return pulumi.ToOutputWithContext(ctx, i).(V2Output)
}

// V2ArrayInput is an input type that accepts V2Array and V2ArrayOutput values.
// You can construct a concrete instance of `V2ArrayInput` via:
//
//	V2Array{ V2Args{...} }
type V2ArrayInput interface {
	pulumi.Input

	ToV2ArrayOutput() V2ArrayOutput
	ToV2ArrayOutputWithContext(context.Context) V2ArrayOutput
}

type V2Array []V2Input

func (V2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*V2)(nil)).Elem()
}

func (i V2Array) ToV2ArrayOutput() V2ArrayOutput {
	return i.ToV2ArrayOutputWithContext(context.Background())
}

func (i V2Array) ToV2ArrayOutputWithContext(ctx context.Context) V2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(V2ArrayOutput)
}

// V2MapInput is an input type that accepts V2Map and V2MapOutput values.
// You can construct a concrete instance of `V2MapInput` via:
//
//	V2Map{ "key": V2Args{...} }
type V2MapInput interface {
	pulumi.Input

	ToV2MapOutput() V2MapOutput
	ToV2MapOutputWithContext(context.Context) V2MapOutput
}

type V2Map map[string]V2Input

func (V2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*V2)(nil)).Elem()
}

func (i V2Map) ToV2MapOutput() V2MapOutput {
	return i.ToV2MapOutputWithContext(context.Background())
}

func (i V2Map) ToV2MapOutputWithContext(ctx context.Context) V2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(V2MapOutput)
}

type V2Output struct{ *pulumi.OutputState }

func (V2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**V2)(nil)).Elem()
}

func (o V2Output) ToV2Output() V2Output {
	return o
}

func (o V2Output) ToV2OutputWithContext(ctx context.Context) V2Output {
	return o
}

// A list of additional Route Targets that will be
// used for export.
func (o V2Output) ExportTargets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *V2) pulumi.StringArrayOutput { return v.ExportTargets }).(pulumi.StringArrayOutput)
}

// A list of additional Route Targets that will be
// imported.
func (o V2Output) ImportTargets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *V2) pulumi.StringArrayOutput { return v.ImportTargets }).(pulumi.StringArrayOutput)
}

// The default BGP LOCAL\_PREF of routes that will be
// advertised to the BGP VPN, unless overridden per-route.
func (o V2Output) LocalPref() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *V2) pulumi.IntPtrOutput { return v.LocalPref }).(pulumi.IntPtrOutput)
}

// The name of the BGP VPN. Changing this updates the name of
// the existing BGP VPN.
func (o V2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *V2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of network IDs that are associated with the BGP VPN.
func (o V2Output) Networks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *V2) pulumi.StringArrayOutput { return v.Networks }).(pulumi.StringArrayOutput)
}

// A list of port IDs that are associated with the BGP VPN.
func (o V2Output) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *V2) pulumi.StringArrayOutput { return v.Ports }).(pulumi.StringArrayOutput)
}

// The ID of the project that owns the BGPVPN. Only
// administrative and users with `advsvc` role can specify a project ID other
// than their own. Changing this creates a new BGP VPN.
func (o V2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *V2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a BGP VPN service. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// BGP VPN.
func (o V2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *V2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A list of route distinguisher strings. If
// specified, one of these RDs will be used to advertise VPN routes.
func (o V2Output) RouteDistinguishers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *V2) pulumi.StringArrayOutput { return v.RouteDistinguishers }).(pulumi.StringArrayOutput)
}

// A list of Route Targets that will be both
// imported and used for export.
func (o V2Output) RouteTargets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *V2) pulumi.StringArrayOutput { return v.RouteTargets }).(pulumi.StringArrayOutput)
}

// A list of router IDs that are associated with the BGP VPN.
func (o V2Output) Routers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *V2) pulumi.StringArrayOutput { return v.Routers }).(pulumi.StringArrayOutput)
}

// Indicates whether the BGP VPN is shared across projects.
func (o V2Output) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v *V2) pulumi.BoolOutput { return v.Shared }).(pulumi.BoolOutput)
}

// The type of the BGP VPN (either `l2` or `l3`). Changing this
// creates a new BGP VPN. Defaults to `l3`.
func (o V2Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *V2) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The globally-assigned VXLAN VNI for the BGP VPN. Changing
// this creates a new BGP VPN.
func (o V2Output) Vni() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *V2) pulumi.IntPtrOutput { return v.Vni }).(pulumi.IntPtrOutput)
}

type V2ArrayOutput struct{ *pulumi.OutputState }

func (V2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*V2)(nil)).Elem()
}

func (o V2ArrayOutput) ToV2ArrayOutput() V2ArrayOutput {
	return o
}

func (o V2ArrayOutput) ToV2ArrayOutputWithContext(ctx context.Context) V2ArrayOutput {
	return o
}

func (o V2ArrayOutput) Index(i pulumi.IntInput) V2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *V2 {
		return vs[0].([]*V2)[vs[1].(int)]
	}).(V2Output)
}

type V2MapOutput struct{ *pulumi.OutputState }

func (V2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*V2)(nil)).Elem()
}

func (o V2MapOutput) ToV2MapOutput() V2MapOutput {
	return o
}

func (o V2MapOutput) ToV2MapOutputWithContext(ctx context.Context) V2MapOutput {
	return o
}

func (o V2MapOutput) MapIndex(k pulumi.StringInput) V2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *V2 {
		return vs[0].(map[string]*V2)[vs[1].(string)]
	}).(V2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*V2Input)(nil)).Elem(), &V2{})
	pulumi.RegisterInputType(reflect.TypeOf((*V2ArrayInput)(nil)).Elem(), V2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*V2MapInput)(nil)).Elem(), V2Map{})
	pulumi.RegisterOutputType(V2Output{})
	pulumi.RegisterOutputType(V2ArrayOutput{})
	pulumi.RegisterOutputType(V2MapOutput{})
}
