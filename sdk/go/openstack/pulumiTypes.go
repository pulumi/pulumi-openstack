// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type BgpvpnPortAssociateV2Route struct {
	// The ID of the BGP VPN to be advertised. Required
	// if `type` is `bgpvpn`. Conflicts with `prefix`.
	BgpvpnId *string `pulumi:"bgpvpnId"`
	// The BGP LOCAL\_PREF value of the routes that will
	// be advertised.
	LocalPref *int `pulumi:"localPref"`
	// The CIDR prefix (v4 or v6) to be advertised. Required
	// if `type` is `prefix`. Conflicts with `bgpvpnId`.
	Prefix *string `pulumi:"prefix"`
	// Can be `prefix` or `bgpvpn`. For the `prefix` type, the
	// CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
	// `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpnId` key.
	Type string `pulumi:"type"`
}

// BgpvpnPortAssociateV2RouteInput is an input type that accepts BgpvpnPortAssociateV2RouteArgs and BgpvpnPortAssociateV2RouteOutput values.
// You can construct a concrete instance of `BgpvpnPortAssociateV2RouteInput` via:
//
//	BgpvpnPortAssociateV2RouteArgs{...}
type BgpvpnPortAssociateV2RouteInput interface {
	pulumi.Input

	ToBgpvpnPortAssociateV2RouteOutput() BgpvpnPortAssociateV2RouteOutput
	ToBgpvpnPortAssociateV2RouteOutputWithContext(context.Context) BgpvpnPortAssociateV2RouteOutput
}

type BgpvpnPortAssociateV2RouteArgs struct {
	// The ID of the BGP VPN to be advertised. Required
	// if `type` is `bgpvpn`. Conflicts with `prefix`.
	BgpvpnId pulumi.StringPtrInput `pulumi:"bgpvpnId"`
	// The BGP LOCAL\_PREF value of the routes that will
	// be advertised.
	LocalPref pulumi.IntPtrInput `pulumi:"localPref"`
	// The CIDR prefix (v4 or v6) to be advertised. Required
	// if `type` is `prefix`. Conflicts with `bgpvpnId`.
	Prefix pulumi.StringPtrInput `pulumi:"prefix"`
	// Can be `prefix` or `bgpvpn`. For the `prefix` type, the
	// CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
	// `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpnId` key.
	Type pulumi.StringInput `pulumi:"type"`
}

func (BgpvpnPortAssociateV2RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpvpnPortAssociateV2Route)(nil)).Elem()
}

func (i BgpvpnPortAssociateV2RouteArgs) ToBgpvpnPortAssociateV2RouteOutput() BgpvpnPortAssociateV2RouteOutput {
	return i.ToBgpvpnPortAssociateV2RouteOutputWithContext(context.Background())
}

func (i BgpvpnPortAssociateV2RouteArgs) ToBgpvpnPortAssociateV2RouteOutputWithContext(ctx context.Context) BgpvpnPortAssociateV2RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpvpnPortAssociateV2RouteOutput)
}

// BgpvpnPortAssociateV2RouteArrayInput is an input type that accepts BgpvpnPortAssociateV2RouteArray and BgpvpnPortAssociateV2RouteArrayOutput values.
// You can construct a concrete instance of `BgpvpnPortAssociateV2RouteArrayInput` via:
//
//	BgpvpnPortAssociateV2RouteArray{ BgpvpnPortAssociateV2RouteArgs{...} }
type BgpvpnPortAssociateV2RouteArrayInput interface {
	pulumi.Input

	ToBgpvpnPortAssociateV2RouteArrayOutput() BgpvpnPortAssociateV2RouteArrayOutput
	ToBgpvpnPortAssociateV2RouteArrayOutputWithContext(context.Context) BgpvpnPortAssociateV2RouteArrayOutput
}

type BgpvpnPortAssociateV2RouteArray []BgpvpnPortAssociateV2RouteInput

func (BgpvpnPortAssociateV2RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BgpvpnPortAssociateV2Route)(nil)).Elem()
}

func (i BgpvpnPortAssociateV2RouteArray) ToBgpvpnPortAssociateV2RouteArrayOutput() BgpvpnPortAssociateV2RouteArrayOutput {
	return i.ToBgpvpnPortAssociateV2RouteArrayOutputWithContext(context.Background())
}

func (i BgpvpnPortAssociateV2RouteArray) ToBgpvpnPortAssociateV2RouteArrayOutputWithContext(ctx context.Context) BgpvpnPortAssociateV2RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpvpnPortAssociateV2RouteArrayOutput)
}

type BgpvpnPortAssociateV2RouteOutput struct{ *pulumi.OutputState }

func (BgpvpnPortAssociateV2RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpvpnPortAssociateV2Route)(nil)).Elem()
}

func (o BgpvpnPortAssociateV2RouteOutput) ToBgpvpnPortAssociateV2RouteOutput() BgpvpnPortAssociateV2RouteOutput {
	return o
}

func (o BgpvpnPortAssociateV2RouteOutput) ToBgpvpnPortAssociateV2RouteOutputWithContext(ctx context.Context) BgpvpnPortAssociateV2RouteOutput {
	return o
}

// The ID of the BGP VPN to be advertised. Required
// if `type` is `bgpvpn`. Conflicts with `prefix`.
func (o BgpvpnPortAssociateV2RouteOutput) BgpvpnId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpvpnPortAssociateV2Route) *string { return v.BgpvpnId }).(pulumi.StringPtrOutput)
}

// The BGP LOCAL\_PREF value of the routes that will
// be advertised.
func (o BgpvpnPortAssociateV2RouteOutput) LocalPref() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpvpnPortAssociateV2Route) *int { return v.LocalPref }).(pulumi.IntPtrOutput)
}

// The CIDR prefix (v4 or v6) to be advertised. Required
// if `type` is `prefix`. Conflicts with `bgpvpnId`.
func (o BgpvpnPortAssociateV2RouteOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpvpnPortAssociateV2Route) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

// Can be `prefix` or `bgpvpn`. For the `prefix` type, the
// CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
// `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpnId` key.
func (o BgpvpnPortAssociateV2RouteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BgpvpnPortAssociateV2Route) string { return v.Type }).(pulumi.StringOutput)
}

type BgpvpnPortAssociateV2RouteArrayOutput struct{ *pulumi.OutputState }

func (BgpvpnPortAssociateV2RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BgpvpnPortAssociateV2Route)(nil)).Elem()
}

func (o BgpvpnPortAssociateV2RouteArrayOutput) ToBgpvpnPortAssociateV2RouteArrayOutput() BgpvpnPortAssociateV2RouteArrayOutput {
	return o
}

func (o BgpvpnPortAssociateV2RouteArrayOutput) ToBgpvpnPortAssociateV2RouteArrayOutputWithContext(ctx context.Context) BgpvpnPortAssociateV2RouteArrayOutput {
	return o
}

func (o BgpvpnPortAssociateV2RouteArrayOutput) Index(i pulumi.IntInput) BgpvpnPortAssociateV2RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BgpvpnPortAssociateV2Route {
		return vs[0].([]BgpvpnPortAssociateV2Route)[vs[1].(int)]
	}).(BgpvpnPortAssociateV2RouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpvpnPortAssociateV2RouteInput)(nil)).Elem(), BgpvpnPortAssociateV2RouteArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpvpnPortAssociateV2RouteArrayInput)(nil)).Elem(), BgpvpnPortAssociateV2RouteArray{})
	pulumi.RegisterOutputType(BgpvpnPortAssociateV2RouteOutput{})
	pulumi.RegisterOutputType(BgpvpnPortAssociateV2RouteArrayOutput{})
}
