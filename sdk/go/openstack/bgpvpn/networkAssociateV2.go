// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bgpvpn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 BGP VPN network association resource within OpenStack.
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
//			_, err := bgpvpn.NewNetworkAssociateV2(ctx, "association_1", &bgpvpn.NetworkAssociateV2Args{
//				BgpvpnId:  pulumi.String("e7189337-5684-46ee-bcb1-44f1a57066c9"),
//				NetworkId: pulumi.String("de83d56c-4d2f-44f7-ac24-af393252204f"),
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
// # BGP VPN network associations can be imported using the BGP VPN ID and BGP VPN
//
// network association ID separated by a slash, e.g.:
//
// hcl
//
// ```sh
// $ pulumi import openstack:bgpvpn/networkAssociateV2:NetworkAssociateV2 association_1 2145aaa9-edaa-44fb-9815-e47a96677a72/67bb952a-f9d1-4fc8-ae84-082253a879d4
// ```
type NetworkAssociateV2 struct {
	pulumi.CustomResourceState

	// The ID of the BGP VPN to which the network will be
	// associated. Changing this creates a new BGP VPN network association
	BgpvpnId pulumi.StringOutput `pulumi:"bgpvpnId"`
	// The ID of the network to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN network association.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// The ID of the project that owns the BGP VPN network
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN network
	// association.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN network association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN network association.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewNetworkAssociateV2 registers a new resource with the given unique name, arguments, and options.
func NewNetworkAssociateV2(ctx *pulumi.Context,
	name string, args *NetworkAssociateV2Args, opts ...pulumi.ResourceOption) (*NetworkAssociateV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BgpvpnId == nil {
		return nil, errors.New("invalid value for required argument 'BgpvpnId'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("openstack:index/bgpvpnNetworkAssociateV2:BgpvpnNetworkAssociateV2"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkAssociateV2
	err := ctx.RegisterResource("openstack:bgpvpn/networkAssociateV2:NetworkAssociateV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAssociateV2 gets an existing NetworkAssociateV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAssociateV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAssociateV2State, opts ...pulumi.ResourceOption) (*NetworkAssociateV2, error) {
	var resource NetworkAssociateV2
	err := ctx.ReadResource("openstack:bgpvpn/networkAssociateV2:NetworkAssociateV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAssociateV2 resources.
type networkAssociateV2State struct {
	// The ID of the BGP VPN to which the network will be
	// associated. Changing this creates a new BGP VPN network association
	BgpvpnId *string `pulumi:"bgpvpnId"`
	// The ID of the network to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN network association.
	NetworkId *string `pulumi:"networkId"`
	// The ID of the project that owns the BGP VPN network
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN network
	// association.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN network association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN network association.
	Region *string `pulumi:"region"`
}

type NetworkAssociateV2State struct {
	// The ID of the BGP VPN to which the network will be
	// associated. Changing this creates a new BGP VPN network association
	BgpvpnId pulumi.StringPtrInput
	// The ID of the network to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN network association.
	NetworkId pulumi.StringPtrInput
	// The ID of the project that owns the BGP VPN network
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN network
	// association.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN network association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN network association.
	Region pulumi.StringPtrInput
}

func (NetworkAssociateV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAssociateV2State)(nil)).Elem()
}

type networkAssociateV2Args struct {
	// The ID of the BGP VPN to which the network will be
	// associated. Changing this creates a new BGP VPN network association
	BgpvpnId string `pulumi:"bgpvpnId"`
	// The ID of the network to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN network association.
	NetworkId string `pulumi:"networkId"`
	// The ID of the project that owns the BGP VPN network
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN network
	// association.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN network association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN network association.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a NetworkAssociateV2 resource.
type NetworkAssociateV2Args struct {
	// The ID of the BGP VPN to which the network will be
	// associated. Changing this creates a new BGP VPN network association
	BgpvpnId pulumi.StringInput
	// The ID of the network to be associated with the BGP
	// VPN. Changing this creates a new BGP VPN network association.
	NetworkId pulumi.StringInput
	// The ID of the project that owns the BGP VPN network
	// association. Only administrative and users with `advsvc` role can specify a
	// project ID other than their own. Changing this creates a new BGP VPN network
	// association.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a BGP VPN network association. If
	// omitted, the `region` argument of the provider is used. Changing this creates
	// a new BGP VPN network association.
	Region pulumi.StringPtrInput
}

func (NetworkAssociateV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAssociateV2Args)(nil)).Elem()
}

type NetworkAssociateV2Input interface {
	pulumi.Input

	ToNetworkAssociateV2Output() NetworkAssociateV2Output
	ToNetworkAssociateV2OutputWithContext(ctx context.Context) NetworkAssociateV2Output
}

func (*NetworkAssociateV2) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAssociateV2)(nil)).Elem()
}

func (i *NetworkAssociateV2) ToNetworkAssociateV2Output() NetworkAssociateV2Output {
	return i.ToNetworkAssociateV2OutputWithContext(context.Background())
}

func (i *NetworkAssociateV2) ToNetworkAssociateV2OutputWithContext(ctx context.Context) NetworkAssociateV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAssociateV2Output)
}

// NetworkAssociateV2ArrayInput is an input type that accepts NetworkAssociateV2Array and NetworkAssociateV2ArrayOutput values.
// You can construct a concrete instance of `NetworkAssociateV2ArrayInput` via:
//
//	NetworkAssociateV2Array{ NetworkAssociateV2Args{...} }
type NetworkAssociateV2ArrayInput interface {
	pulumi.Input

	ToNetworkAssociateV2ArrayOutput() NetworkAssociateV2ArrayOutput
	ToNetworkAssociateV2ArrayOutputWithContext(context.Context) NetworkAssociateV2ArrayOutput
}

type NetworkAssociateV2Array []NetworkAssociateV2Input

func (NetworkAssociateV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAssociateV2)(nil)).Elem()
}

func (i NetworkAssociateV2Array) ToNetworkAssociateV2ArrayOutput() NetworkAssociateV2ArrayOutput {
	return i.ToNetworkAssociateV2ArrayOutputWithContext(context.Background())
}

func (i NetworkAssociateV2Array) ToNetworkAssociateV2ArrayOutputWithContext(ctx context.Context) NetworkAssociateV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAssociateV2ArrayOutput)
}

// NetworkAssociateV2MapInput is an input type that accepts NetworkAssociateV2Map and NetworkAssociateV2MapOutput values.
// You can construct a concrete instance of `NetworkAssociateV2MapInput` via:
//
//	NetworkAssociateV2Map{ "key": NetworkAssociateV2Args{...} }
type NetworkAssociateV2MapInput interface {
	pulumi.Input

	ToNetworkAssociateV2MapOutput() NetworkAssociateV2MapOutput
	ToNetworkAssociateV2MapOutputWithContext(context.Context) NetworkAssociateV2MapOutput
}

type NetworkAssociateV2Map map[string]NetworkAssociateV2Input

func (NetworkAssociateV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAssociateV2)(nil)).Elem()
}

func (i NetworkAssociateV2Map) ToNetworkAssociateV2MapOutput() NetworkAssociateV2MapOutput {
	return i.ToNetworkAssociateV2MapOutputWithContext(context.Background())
}

func (i NetworkAssociateV2Map) ToNetworkAssociateV2MapOutputWithContext(ctx context.Context) NetworkAssociateV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAssociateV2MapOutput)
}

type NetworkAssociateV2Output struct{ *pulumi.OutputState }

func (NetworkAssociateV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAssociateV2)(nil)).Elem()
}

func (o NetworkAssociateV2Output) ToNetworkAssociateV2Output() NetworkAssociateV2Output {
	return o
}

func (o NetworkAssociateV2Output) ToNetworkAssociateV2OutputWithContext(ctx context.Context) NetworkAssociateV2Output {
	return o
}

// The ID of the BGP VPN to which the network will be
// associated. Changing this creates a new BGP VPN network association
func (o NetworkAssociateV2Output) BgpvpnId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAssociateV2) pulumi.StringOutput { return v.BgpvpnId }).(pulumi.StringOutput)
}

// The ID of the network to be associated with the BGP
// VPN. Changing this creates a new BGP VPN network association.
func (o NetworkAssociateV2Output) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAssociateV2) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// The ID of the project that owns the BGP VPN network
// association. Only administrative and users with `advsvc` role can specify a
// project ID other than their own. Changing this creates a new BGP VPN network
// association.
func (o NetworkAssociateV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAssociateV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a BGP VPN network association. If
// omitted, the `region` argument of the provider is used. Changing this creates
// a new BGP VPN network association.
func (o NetworkAssociateV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAssociateV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type NetworkAssociateV2ArrayOutput struct{ *pulumi.OutputState }

func (NetworkAssociateV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAssociateV2)(nil)).Elem()
}

func (o NetworkAssociateV2ArrayOutput) ToNetworkAssociateV2ArrayOutput() NetworkAssociateV2ArrayOutput {
	return o
}

func (o NetworkAssociateV2ArrayOutput) ToNetworkAssociateV2ArrayOutputWithContext(ctx context.Context) NetworkAssociateV2ArrayOutput {
	return o
}

func (o NetworkAssociateV2ArrayOutput) Index(i pulumi.IntInput) NetworkAssociateV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkAssociateV2 {
		return vs[0].([]*NetworkAssociateV2)[vs[1].(int)]
	}).(NetworkAssociateV2Output)
}

type NetworkAssociateV2MapOutput struct{ *pulumi.OutputState }

func (NetworkAssociateV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAssociateV2)(nil)).Elem()
}

func (o NetworkAssociateV2MapOutput) ToNetworkAssociateV2MapOutput() NetworkAssociateV2MapOutput {
	return o
}

func (o NetworkAssociateV2MapOutput) ToNetworkAssociateV2MapOutputWithContext(ctx context.Context) NetworkAssociateV2MapOutput {
	return o
}

func (o NetworkAssociateV2MapOutput) MapIndex(k pulumi.StringInput) NetworkAssociateV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkAssociateV2 {
		return vs[0].(map[string]*NetworkAssociateV2)[vs[1].(string)]
	}).(NetworkAssociateV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAssociateV2Input)(nil)).Elem(), &NetworkAssociateV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAssociateV2ArrayInput)(nil)).Elem(), NetworkAssociateV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAssociateV2MapInput)(nil)).Elem(), NetworkAssociateV2Map{})
	pulumi.RegisterOutputType(NetworkAssociateV2Output{})
	pulumi.RegisterOutputType(NetworkAssociateV2ArrayOutput{})
	pulumi.RegisterOutputType(NetworkAssociateV2MapOutput{})
}
