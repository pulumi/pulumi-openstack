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

// Manages a V2 Neutron BGP Peer resource within OpenStack.
//
// This resource allows you to configure a BGP peer that can be associated with a
// BGP speaker to exchange routing information.
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
//			_, err := networking.NewBgpPeerV2(ctx, "peer_1", &networking.BgpPeerV2Args{
//				Name:     pulumi.String("bgp_peer_1"),
//				PeerIp:   pulumi.String("192.0.2.10"),
//				RemoteAs: pulumi.Int(65001),
//				AuthType: pulumi.String("md5"),
//				Password: pulumi.String("supersecret"),
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
// BGP peers can be imported using their ID:
//
// ```sh
// $ pulumi import openstack:networking/bgpPeerV2:BgpPeerV2 peer1 a1b2c3d4-e5f6-7890-abcd-1234567890ef
// ```
type BgpPeerV2 struct {
	pulumi.CustomResourceState

	// The authentication type to use. Can be one of `none`
	// or `md5`. Defaults to `none`. If set to not `none`, the `password` argument
	// must also be provided. Changing this creates a new BGP peer.
	AuthType pulumi.StringPtrOutput `pulumi:"authType"`
	// A name for the BGP peer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password used for MD5 authentication. Must be set
	// only when `authType` is not `none`.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The IP address of the BGP peer. Must be a valid IP
	// address. Changing this creates a new BGP peer.
	PeerIp pulumi.StringOutput `pulumi:"peerIp"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new BGP
	// peer.
	Region pulumi.StringOutput `pulumi:"region"`
	// The AS number of the BGP peer. Changing this
	// creates a new BGP peer.
	RemoteAs pulumi.IntOutput `pulumi:"remoteAs"`
	// The tenant/project ID. Required if admin privileges
	// are used. Changing this creates a new BGP peer.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewBgpPeerV2 registers a new resource with the given unique name, arguments, and options.
func NewBgpPeerV2(ctx *pulumi.Context,
	name string, args *BgpPeerV2Args, opts ...pulumi.ResourceOption) (*BgpPeerV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerIp == nil {
		return nil, errors.New("invalid value for required argument 'PeerIp'")
	}
	if args.RemoteAs == nil {
		return nil, errors.New("invalid value for required argument 'RemoteAs'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BgpPeerV2
	err := ctx.RegisterResource("openstack:networking/bgpPeerV2:BgpPeerV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpPeerV2 gets an existing BgpPeerV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpPeerV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpPeerV2State, opts ...pulumi.ResourceOption) (*BgpPeerV2, error) {
	var resource BgpPeerV2
	err := ctx.ReadResource("openstack:networking/bgpPeerV2:BgpPeerV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpPeerV2 resources.
type bgpPeerV2State struct {
	// The authentication type to use. Can be one of `none`
	// or `md5`. Defaults to `none`. If set to not `none`, the `password` argument
	// must also be provided. Changing this creates a new BGP peer.
	AuthType *string `pulumi:"authType"`
	// A name for the BGP peer.
	Name *string `pulumi:"name"`
	// The password used for MD5 authentication. Must be set
	// only when `authType` is not `none`.
	Password *string `pulumi:"password"`
	// The IP address of the BGP peer. Must be a valid IP
	// address. Changing this creates a new BGP peer.
	PeerIp *string `pulumi:"peerIp"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new BGP
	// peer.
	Region *string `pulumi:"region"`
	// The AS number of the BGP peer. Changing this
	// creates a new BGP peer.
	RemoteAs *int `pulumi:"remoteAs"`
	// The tenant/project ID. Required if admin privileges
	// are used. Changing this creates a new BGP peer.
	TenantId *string `pulumi:"tenantId"`
}

type BgpPeerV2State struct {
	// The authentication type to use. Can be one of `none`
	// or `md5`. Defaults to `none`. If set to not `none`, the `password` argument
	// must also be provided. Changing this creates a new BGP peer.
	AuthType pulumi.StringPtrInput
	// A name for the BGP peer.
	Name pulumi.StringPtrInput
	// The password used for MD5 authentication. Must be set
	// only when `authType` is not `none`.
	Password pulumi.StringPtrInput
	// The IP address of the BGP peer. Must be a valid IP
	// address. Changing this creates a new BGP peer.
	PeerIp pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new BGP
	// peer.
	Region pulumi.StringPtrInput
	// The AS number of the BGP peer. Changing this
	// creates a new BGP peer.
	RemoteAs pulumi.IntPtrInput
	// The tenant/project ID. Required if admin privileges
	// are used. Changing this creates a new BGP peer.
	TenantId pulumi.StringPtrInput
}

func (BgpPeerV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerV2State)(nil)).Elem()
}

type bgpPeerV2Args struct {
	// The authentication type to use. Can be one of `none`
	// or `md5`. Defaults to `none`. If set to not `none`, the `password` argument
	// must also be provided. Changing this creates a new BGP peer.
	AuthType *string `pulumi:"authType"`
	// A name for the BGP peer.
	Name *string `pulumi:"name"`
	// The password used for MD5 authentication. Must be set
	// only when `authType` is not `none`.
	Password *string `pulumi:"password"`
	// The IP address of the BGP peer. Must be a valid IP
	// address. Changing this creates a new BGP peer.
	PeerIp string `pulumi:"peerIp"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new BGP
	// peer.
	Region *string `pulumi:"region"`
	// The AS number of the BGP peer. Changing this
	// creates a new BGP peer.
	RemoteAs int `pulumi:"remoteAs"`
	// The tenant/project ID. Required if admin privileges
	// are used. Changing this creates a new BGP peer.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a BgpPeerV2 resource.
type BgpPeerV2Args struct {
	// The authentication type to use. Can be one of `none`
	// or `md5`. Defaults to `none`. If set to not `none`, the `password` argument
	// must also be provided. Changing this creates a new BGP peer.
	AuthType pulumi.StringPtrInput
	// A name for the BGP peer.
	Name pulumi.StringPtrInput
	// The password used for MD5 authentication. Must be set
	// only when `authType` is not `none`.
	Password pulumi.StringPtrInput
	// The IP address of the BGP peer. Must be a valid IP
	// address. Changing this creates a new BGP peer.
	PeerIp pulumi.StringInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new BGP
	// peer.
	Region pulumi.StringPtrInput
	// The AS number of the BGP peer. Changing this
	// creates a new BGP peer.
	RemoteAs pulumi.IntInput
	// The tenant/project ID. Required if admin privileges
	// are used. Changing this creates a new BGP peer.
	TenantId pulumi.StringPtrInput
}

func (BgpPeerV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerV2Args)(nil)).Elem()
}

type BgpPeerV2Input interface {
	pulumi.Input

	ToBgpPeerV2Output() BgpPeerV2Output
	ToBgpPeerV2OutputWithContext(ctx context.Context) BgpPeerV2Output
}

func (*BgpPeerV2) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpPeerV2)(nil)).Elem()
}

func (i *BgpPeerV2) ToBgpPeerV2Output() BgpPeerV2Output {
	return i.ToBgpPeerV2OutputWithContext(context.Background())
}

func (i *BgpPeerV2) ToBgpPeerV2OutputWithContext(ctx context.Context) BgpPeerV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerV2Output)
}

// BgpPeerV2ArrayInput is an input type that accepts BgpPeerV2Array and BgpPeerV2ArrayOutput values.
// You can construct a concrete instance of `BgpPeerV2ArrayInput` via:
//
//	BgpPeerV2Array{ BgpPeerV2Args{...} }
type BgpPeerV2ArrayInput interface {
	pulumi.Input

	ToBgpPeerV2ArrayOutput() BgpPeerV2ArrayOutput
	ToBgpPeerV2ArrayOutputWithContext(context.Context) BgpPeerV2ArrayOutput
}

type BgpPeerV2Array []BgpPeerV2Input

func (BgpPeerV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpPeerV2)(nil)).Elem()
}

func (i BgpPeerV2Array) ToBgpPeerV2ArrayOutput() BgpPeerV2ArrayOutput {
	return i.ToBgpPeerV2ArrayOutputWithContext(context.Background())
}

func (i BgpPeerV2Array) ToBgpPeerV2ArrayOutputWithContext(ctx context.Context) BgpPeerV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerV2ArrayOutput)
}

// BgpPeerV2MapInput is an input type that accepts BgpPeerV2Map and BgpPeerV2MapOutput values.
// You can construct a concrete instance of `BgpPeerV2MapInput` via:
//
//	BgpPeerV2Map{ "key": BgpPeerV2Args{...} }
type BgpPeerV2MapInput interface {
	pulumi.Input

	ToBgpPeerV2MapOutput() BgpPeerV2MapOutput
	ToBgpPeerV2MapOutputWithContext(context.Context) BgpPeerV2MapOutput
}

type BgpPeerV2Map map[string]BgpPeerV2Input

func (BgpPeerV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpPeerV2)(nil)).Elem()
}

func (i BgpPeerV2Map) ToBgpPeerV2MapOutput() BgpPeerV2MapOutput {
	return i.ToBgpPeerV2MapOutputWithContext(context.Background())
}

func (i BgpPeerV2Map) ToBgpPeerV2MapOutputWithContext(ctx context.Context) BgpPeerV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerV2MapOutput)
}

type BgpPeerV2Output struct{ *pulumi.OutputState }

func (BgpPeerV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpPeerV2)(nil)).Elem()
}

func (o BgpPeerV2Output) ToBgpPeerV2Output() BgpPeerV2Output {
	return o
}

func (o BgpPeerV2Output) ToBgpPeerV2OutputWithContext(ctx context.Context) BgpPeerV2Output {
	return o
}

// The authentication type to use. Can be one of `none`
// or `md5`. Defaults to `none`. If set to not `none`, the `password` argument
// must also be provided. Changing this creates a new BGP peer.
func (o BgpPeerV2Output) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpPeerV2) pulumi.StringPtrOutput { return v.AuthType }).(pulumi.StringPtrOutput)
}

// A name for the BGP peer.
func (o BgpPeerV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password used for MD5 authentication. Must be set
// only when `authType` is not `none`.
func (o BgpPeerV2Output) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpPeerV2) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The IP address of the BGP peer. Must be a valid IP
// address. Changing this creates a new BGP peer.
func (o BgpPeerV2Output) PeerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerV2) pulumi.StringOutput { return v.PeerIp }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a Neutron network. If omitted, the
// `region` argument of the provider is used. Changing this creates a new BGP
// peer.
func (o BgpPeerV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The AS number of the BGP peer. Changing this
// creates a new BGP peer.
func (o BgpPeerV2Output) RemoteAs() pulumi.IntOutput {
	return o.ApplyT(func(v *BgpPeerV2) pulumi.IntOutput { return v.RemoteAs }).(pulumi.IntOutput)
}

// The tenant/project ID. Required if admin privileges
// are used. Changing this creates a new BGP peer.
func (o BgpPeerV2Output) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerV2) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type BgpPeerV2ArrayOutput struct{ *pulumi.OutputState }

func (BgpPeerV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpPeerV2)(nil)).Elem()
}

func (o BgpPeerV2ArrayOutput) ToBgpPeerV2ArrayOutput() BgpPeerV2ArrayOutput {
	return o
}

func (o BgpPeerV2ArrayOutput) ToBgpPeerV2ArrayOutputWithContext(ctx context.Context) BgpPeerV2ArrayOutput {
	return o
}

func (o BgpPeerV2ArrayOutput) Index(i pulumi.IntInput) BgpPeerV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BgpPeerV2 {
		return vs[0].([]*BgpPeerV2)[vs[1].(int)]
	}).(BgpPeerV2Output)
}

type BgpPeerV2MapOutput struct{ *pulumi.OutputState }

func (BgpPeerV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpPeerV2)(nil)).Elem()
}

func (o BgpPeerV2MapOutput) ToBgpPeerV2MapOutput() BgpPeerV2MapOutput {
	return o
}

func (o BgpPeerV2MapOutput) ToBgpPeerV2MapOutputWithContext(ctx context.Context) BgpPeerV2MapOutput {
	return o
}

func (o BgpPeerV2MapOutput) MapIndex(k pulumi.StringInput) BgpPeerV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BgpPeerV2 {
		return vs[0].(map[string]*BgpPeerV2)[vs[1].(string)]
	}).(BgpPeerV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerV2Input)(nil)).Elem(), &BgpPeerV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerV2ArrayInput)(nil)).Elem(), BgpPeerV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerV2MapInput)(nil)).Elem(), BgpPeerV2Map{})
	pulumi.RegisterOutputType(BgpPeerV2Output{})
	pulumi.RegisterOutputType(BgpPeerV2ArrayOutput{})
	pulumi.RegisterOutputType(BgpPeerV2MapOutput{})
}
