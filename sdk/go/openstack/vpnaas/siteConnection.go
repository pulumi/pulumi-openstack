// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vpnaas

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Neutron IPSec site connection resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/vpnaas"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vpnaas.NewSiteConnection(ctx, "conn1", &vpnaas.SiteConnectionArgs{
// 			IkepolicyId:    pulumi.Any(openstack_vpnaas_ike_policy_v2.Policy_2.Id),
// 			IpsecpolicyId:  pulumi.Any(openstack_vpnaas_ipsec_policy_v2.Policy_1.Id),
// 			LocalEpGroupId: pulumi.Any(openstack_vpnaas_endpoint_group_v2.Group_2.Id),
// 			PeerAddress:    pulumi.String("192.168.10.1"),
// 			PeerEpGroupId:  pulumi.Any(openstack_vpnaas_endpoint_group_v2.Group_1.Id),
// 			Psk:            pulumi.String("secret"),
// 			VpnserviceId:   pulumi.Any(openstack_vpnaas_service_v2.Service_1.Id),
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
// Site Connections can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:vpnaas/siteConnection:SiteConnection conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
// ```
type SiteConnection struct {
	pulumi.CustomResourceState

	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing connection.
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// The human-readable description for the connection.
	// Changing this updates the description of the existing connection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A dictionary with dead peer detection (DPD) protocol controls.
	Dpds SiteConnectionDpdArrayOutput `pulumi:"dpds"`
	// The ID of the IKE policy. Changing this creates a new connection.
	IkepolicyId pulumi.StringOutput `pulumi:"ikepolicyId"`
	// A valid value is response-only or bi-directional. Default is bi-directional.
	Initiator pulumi.StringOutput `pulumi:"initiator"`
	// The ID of the IPsec policy. Changing this creates a new connection.
	IpsecpolicyId pulumi.StringOutput `pulumi:"ipsecpolicyId"`
	// The ID for the endpoint group that contains private subnets for the local side of the connection.
	// You must specify this parameter with the peerEpGroupId parameter unless
	// in backward- compatible mode where peerCidrs is provided with a subnetId for the VPN service.
	// Changing this updates the existing connection.
	LocalEpGroupId pulumi.StringPtrOutput `pulumi:"localEpGroupId"`
	// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
	// Most often, local ID would be domain name, email address, etc.
	// If this is not configured then the external IP address will be used as the ID.
	LocalId pulumi.StringPtrOutput `pulumi:"localId"`
	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	Mtu pulumi.IntOutput `pulumi:"mtu"`
	// The name of the connection. Changing this updates the name of
	// the existing connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// The peer gateway public IPv4 or IPv6 address or FQDN.
	PeerAddress pulumi.StringOutput `pulumi:"peerAddress"`
	// Unique list of valid peer private CIDRs in the form < netAddress > / < prefix > .
	PeerCidrs pulumi.StringArrayOutput `pulumi:"peerCidrs"`
	// The ID for the endpoint group that contains private CIDRs in the form < netAddress > / < prefix > for the peer side of the connection.
	// You must specify this parameter with the localEpGroupId parameter unless in backward-compatible mode
	// where peerCidrs is provided with a subnetId for the VPN service.
	PeerEpGroupId pulumi.StringPtrOutput `pulumi:"peerEpGroupId"`
	// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
	// Typically, this value matches the peerAddress value.
	// Changing this updates the existing policy.
	PeerId pulumi.StringOutput `pulumi:"peerId"`
	// The pre-shared key. A valid value is any string.
	Psk pulumi.StringOutput `pulumi:"psk"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec site connection. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// site connection.
	Region pulumi.StringOutput `pulumi:"region"`
	// The owner of the connection. Required if admin wants to
	// create a connection for another project. Changing this creates a new connection.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
	// The ID of the VPN service. Changing this creates a new connection.
	VpnserviceId pulumi.StringOutput `pulumi:"vpnserviceId"`
}

// NewSiteConnection registers a new resource with the given unique name, arguments, and options.
func NewSiteConnection(ctx *pulumi.Context,
	name string, args *SiteConnectionArgs, opts ...pulumi.ResourceOption) (*SiteConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IkepolicyId == nil {
		return nil, errors.New("invalid value for required argument 'IkepolicyId'")
	}
	if args.IpsecpolicyId == nil {
		return nil, errors.New("invalid value for required argument 'IpsecpolicyId'")
	}
	if args.PeerAddress == nil {
		return nil, errors.New("invalid value for required argument 'PeerAddress'")
	}
	if args.PeerId == nil {
		return nil, errors.New("invalid value for required argument 'PeerId'")
	}
	if args.Psk == nil {
		return nil, errors.New("invalid value for required argument 'Psk'")
	}
	if args.VpnserviceId == nil {
		return nil, errors.New("invalid value for required argument 'VpnserviceId'")
	}
	var resource SiteConnection
	err := ctx.RegisterResource("openstack:vpnaas/siteConnection:SiteConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteConnection gets an existing SiteConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteConnectionState, opts ...pulumi.ResourceOption) (*SiteConnection, error) {
	var resource SiteConnection
	err := ctx.ReadResource("openstack:vpnaas/siteConnection:SiteConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteConnection resources.
type siteConnectionState struct {
	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing connection.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The human-readable description for the connection.
	// Changing this updates the description of the existing connection.
	Description *string `pulumi:"description"`
	// A dictionary with dead peer detection (DPD) protocol controls.
	Dpds []SiteConnectionDpd `pulumi:"dpds"`
	// The ID of the IKE policy. Changing this creates a new connection.
	IkepolicyId *string `pulumi:"ikepolicyId"`
	// A valid value is response-only or bi-directional. Default is bi-directional.
	Initiator *string `pulumi:"initiator"`
	// The ID of the IPsec policy. Changing this creates a new connection.
	IpsecpolicyId *string `pulumi:"ipsecpolicyId"`
	// The ID for the endpoint group that contains private subnets for the local side of the connection.
	// You must specify this parameter with the peerEpGroupId parameter unless
	// in backward- compatible mode where peerCidrs is provided with a subnetId for the VPN service.
	// Changing this updates the existing connection.
	LocalEpGroupId *string `pulumi:"localEpGroupId"`
	// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
	// Most often, local ID would be domain name, email address, etc.
	// If this is not configured then the external IP address will be used as the ID.
	LocalId *string `pulumi:"localId"`
	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	Mtu *int `pulumi:"mtu"`
	// The name of the connection. Changing this updates the name of
	// the existing connection.
	Name *string `pulumi:"name"`
	// The peer gateway public IPv4 or IPv6 address or FQDN.
	PeerAddress *string `pulumi:"peerAddress"`
	// Unique list of valid peer private CIDRs in the form < netAddress > / < prefix > .
	PeerCidrs []string `pulumi:"peerCidrs"`
	// The ID for the endpoint group that contains private CIDRs in the form < netAddress > / < prefix > for the peer side of the connection.
	// You must specify this parameter with the localEpGroupId parameter unless in backward-compatible mode
	// where peerCidrs is provided with a subnetId for the VPN service.
	PeerEpGroupId *string `pulumi:"peerEpGroupId"`
	// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
	// Typically, this value matches the peerAddress value.
	// Changing this updates the existing policy.
	PeerId *string `pulumi:"peerId"`
	// The pre-shared key. A valid value is any string.
	Psk *string `pulumi:"psk"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec site connection. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// site connection.
	Region *string `pulumi:"region"`
	// The owner of the connection. Required if admin wants to
	// create a connection for another project. Changing this creates a new connection.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
	// The ID of the VPN service. Changing this creates a new connection.
	VpnserviceId *string `pulumi:"vpnserviceId"`
}

type SiteConnectionState struct {
	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing connection.
	AdminStateUp pulumi.BoolPtrInput
	// The human-readable description for the connection.
	// Changing this updates the description of the existing connection.
	Description pulumi.StringPtrInput
	// A dictionary with dead peer detection (DPD) protocol controls.
	Dpds SiteConnectionDpdArrayInput
	// The ID of the IKE policy. Changing this creates a new connection.
	IkepolicyId pulumi.StringPtrInput
	// A valid value is response-only or bi-directional. Default is bi-directional.
	Initiator pulumi.StringPtrInput
	// The ID of the IPsec policy. Changing this creates a new connection.
	IpsecpolicyId pulumi.StringPtrInput
	// The ID for the endpoint group that contains private subnets for the local side of the connection.
	// You must specify this parameter with the peerEpGroupId parameter unless
	// in backward- compatible mode where peerCidrs is provided with a subnetId for the VPN service.
	// Changing this updates the existing connection.
	LocalEpGroupId pulumi.StringPtrInput
	// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
	// Most often, local ID would be domain name, email address, etc.
	// If this is not configured then the external IP address will be used as the ID.
	LocalId pulumi.StringPtrInput
	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	Mtu pulumi.IntPtrInput
	// The name of the connection. Changing this updates the name of
	// the existing connection.
	Name pulumi.StringPtrInput
	// The peer gateway public IPv4 or IPv6 address or FQDN.
	PeerAddress pulumi.StringPtrInput
	// Unique list of valid peer private CIDRs in the form < netAddress > / < prefix > .
	PeerCidrs pulumi.StringArrayInput
	// The ID for the endpoint group that contains private CIDRs in the form < netAddress > / < prefix > for the peer side of the connection.
	// You must specify this parameter with the localEpGroupId parameter unless in backward-compatible mode
	// where peerCidrs is provided with a subnetId for the VPN service.
	PeerEpGroupId pulumi.StringPtrInput
	// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
	// Typically, this value matches the peerAddress value.
	// Changing this updates the existing policy.
	PeerId pulumi.StringPtrInput
	// The pre-shared key. A valid value is any string.
	Psk pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec site connection. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// site connection.
	Region pulumi.StringPtrInput
	// The owner of the connection. Required if admin wants to
	// create a connection for another project. Changing this creates a new connection.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
	// The ID of the VPN service. Changing this creates a new connection.
	VpnserviceId pulumi.StringPtrInput
}

func (SiteConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteConnectionState)(nil)).Elem()
}

type siteConnectionArgs struct {
	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing connection.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The human-readable description for the connection.
	// Changing this updates the description of the existing connection.
	Description *string `pulumi:"description"`
	// A dictionary with dead peer detection (DPD) protocol controls.
	Dpds []SiteConnectionDpd `pulumi:"dpds"`
	// The ID of the IKE policy. Changing this creates a new connection.
	IkepolicyId string `pulumi:"ikepolicyId"`
	// A valid value is response-only or bi-directional. Default is bi-directional.
	Initiator *string `pulumi:"initiator"`
	// The ID of the IPsec policy. Changing this creates a new connection.
	IpsecpolicyId string `pulumi:"ipsecpolicyId"`
	// The ID for the endpoint group that contains private subnets for the local side of the connection.
	// You must specify this parameter with the peerEpGroupId parameter unless
	// in backward- compatible mode where peerCidrs is provided with a subnetId for the VPN service.
	// Changing this updates the existing connection.
	LocalEpGroupId *string `pulumi:"localEpGroupId"`
	// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
	// Most often, local ID would be domain name, email address, etc.
	// If this is not configured then the external IP address will be used as the ID.
	LocalId *string `pulumi:"localId"`
	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	Mtu *int `pulumi:"mtu"`
	// The name of the connection. Changing this updates the name of
	// the existing connection.
	Name *string `pulumi:"name"`
	// The peer gateway public IPv4 or IPv6 address or FQDN.
	PeerAddress string `pulumi:"peerAddress"`
	// Unique list of valid peer private CIDRs in the form < netAddress > / < prefix > .
	PeerCidrs []string `pulumi:"peerCidrs"`
	// The ID for the endpoint group that contains private CIDRs in the form < netAddress > / < prefix > for the peer side of the connection.
	// You must specify this parameter with the localEpGroupId parameter unless in backward-compatible mode
	// where peerCidrs is provided with a subnetId for the VPN service.
	PeerEpGroupId *string `pulumi:"peerEpGroupId"`
	// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
	// Typically, this value matches the peerAddress value.
	// Changing this updates the existing policy.
	PeerId string `pulumi:"peerId"`
	// The pre-shared key. A valid value is any string.
	Psk string `pulumi:"psk"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec site connection. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// site connection.
	Region *string `pulumi:"region"`
	// The owner of the connection. Required if admin wants to
	// create a connection for another project. Changing this creates a new connection.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
	// The ID of the VPN service. Changing this creates a new connection.
	VpnserviceId string `pulumi:"vpnserviceId"`
}

// The set of arguments for constructing a SiteConnection resource.
type SiteConnectionArgs struct {
	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing connection.
	AdminStateUp pulumi.BoolPtrInput
	// The human-readable description for the connection.
	// Changing this updates the description of the existing connection.
	Description pulumi.StringPtrInput
	// A dictionary with dead peer detection (DPD) protocol controls.
	Dpds SiteConnectionDpdArrayInput
	// The ID of the IKE policy. Changing this creates a new connection.
	IkepolicyId pulumi.StringInput
	// A valid value is response-only or bi-directional. Default is bi-directional.
	Initiator pulumi.StringPtrInput
	// The ID of the IPsec policy. Changing this creates a new connection.
	IpsecpolicyId pulumi.StringInput
	// The ID for the endpoint group that contains private subnets for the local side of the connection.
	// You must specify this parameter with the peerEpGroupId parameter unless
	// in backward- compatible mode where peerCidrs is provided with a subnetId for the VPN service.
	// Changing this updates the existing connection.
	LocalEpGroupId pulumi.StringPtrInput
	// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
	// Most often, local ID would be domain name, email address, etc.
	// If this is not configured then the external IP address will be used as the ID.
	LocalId pulumi.StringPtrInput
	// The maximum transmission unit (MTU) value to address fragmentation.
	// Minimum value is 68 for IPv4, and 1280 for IPv6.
	Mtu pulumi.IntPtrInput
	// The name of the connection. Changing this updates the name of
	// the existing connection.
	Name pulumi.StringPtrInput
	// The peer gateway public IPv4 or IPv6 address or FQDN.
	PeerAddress pulumi.StringInput
	// Unique list of valid peer private CIDRs in the form < netAddress > / < prefix > .
	PeerCidrs pulumi.StringArrayInput
	// The ID for the endpoint group that contains private CIDRs in the form < netAddress > / < prefix > for the peer side of the connection.
	// You must specify this parameter with the localEpGroupId parameter unless in backward-compatible mode
	// where peerCidrs is provided with a subnetId for the VPN service.
	PeerEpGroupId pulumi.StringPtrInput
	// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
	// Typically, this value matches the peerAddress value.
	// Changing this updates the existing policy.
	PeerId pulumi.StringInput
	// The pre-shared key. A valid value is any string.
	Psk pulumi.StringInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec site connection. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// site connection.
	Region pulumi.StringPtrInput
	// The owner of the connection. Required if admin wants to
	// create a connection for another project. Changing this creates a new connection.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
	// The ID of the VPN service. Changing this creates a new connection.
	VpnserviceId pulumi.StringInput
}

func (SiteConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteConnectionArgs)(nil)).Elem()
}

type SiteConnectionInput interface {
	pulumi.Input

	ToSiteConnectionOutput() SiteConnectionOutput
	ToSiteConnectionOutputWithContext(ctx context.Context) SiteConnectionOutput
}

func (*SiteConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConnection)(nil))
}

func (i *SiteConnection) ToSiteConnectionOutput() SiteConnectionOutput {
	return i.ToSiteConnectionOutputWithContext(context.Background())
}

func (i *SiteConnection) ToSiteConnectionOutputWithContext(ctx context.Context) SiteConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionOutput)
}

func (i *SiteConnection) ToSiteConnectionPtrOutput() SiteConnectionPtrOutput {
	return i.ToSiteConnectionPtrOutputWithContext(context.Background())
}

func (i *SiteConnection) ToSiteConnectionPtrOutputWithContext(ctx context.Context) SiteConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionPtrOutput)
}

type SiteConnectionPtrInput interface {
	pulumi.Input

	ToSiteConnectionPtrOutput() SiteConnectionPtrOutput
	ToSiteConnectionPtrOutputWithContext(ctx context.Context) SiteConnectionPtrOutput
}

type siteConnectionPtrType SiteConnectionArgs

func (*siteConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConnection)(nil))
}

func (i *siteConnectionPtrType) ToSiteConnectionPtrOutput() SiteConnectionPtrOutput {
	return i.ToSiteConnectionPtrOutputWithContext(context.Background())
}

func (i *siteConnectionPtrType) ToSiteConnectionPtrOutputWithContext(ctx context.Context) SiteConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionPtrOutput)
}

// SiteConnectionArrayInput is an input type that accepts SiteConnectionArray and SiteConnectionArrayOutput values.
// You can construct a concrete instance of `SiteConnectionArrayInput` via:
//
//          SiteConnectionArray{ SiteConnectionArgs{...} }
type SiteConnectionArrayInput interface {
	pulumi.Input

	ToSiteConnectionArrayOutput() SiteConnectionArrayOutput
	ToSiteConnectionArrayOutputWithContext(context.Context) SiteConnectionArrayOutput
}

type SiteConnectionArray []SiteConnectionInput

func (SiteConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SiteConnection)(nil)).Elem()
}

func (i SiteConnectionArray) ToSiteConnectionArrayOutput() SiteConnectionArrayOutput {
	return i.ToSiteConnectionArrayOutputWithContext(context.Background())
}

func (i SiteConnectionArray) ToSiteConnectionArrayOutputWithContext(ctx context.Context) SiteConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionArrayOutput)
}

// SiteConnectionMapInput is an input type that accepts SiteConnectionMap and SiteConnectionMapOutput values.
// You can construct a concrete instance of `SiteConnectionMapInput` via:
//
//          SiteConnectionMap{ "key": SiteConnectionArgs{...} }
type SiteConnectionMapInput interface {
	pulumi.Input

	ToSiteConnectionMapOutput() SiteConnectionMapOutput
	ToSiteConnectionMapOutputWithContext(context.Context) SiteConnectionMapOutput
}

type SiteConnectionMap map[string]SiteConnectionInput

func (SiteConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SiteConnection)(nil)).Elem()
}

func (i SiteConnectionMap) ToSiteConnectionMapOutput() SiteConnectionMapOutput {
	return i.ToSiteConnectionMapOutputWithContext(context.Background())
}

func (i SiteConnectionMap) ToSiteConnectionMapOutputWithContext(ctx context.Context) SiteConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionMapOutput)
}

type SiteConnectionOutput struct{ *pulumi.OutputState }

func (SiteConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConnection)(nil))
}

func (o SiteConnectionOutput) ToSiteConnectionOutput() SiteConnectionOutput {
	return o
}

func (o SiteConnectionOutput) ToSiteConnectionOutputWithContext(ctx context.Context) SiteConnectionOutput {
	return o
}

func (o SiteConnectionOutput) ToSiteConnectionPtrOutput() SiteConnectionPtrOutput {
	return o.ToSiteConnectionPtrOutputWithContext(context.Background())
}

func (o SiteConnectionOutput) ToSiteConnectionPtrOutputWithContext(ctx context.Context) SiteConnectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteConnection) *SiteConnection {
		return &v
	}).(SiteConnectionPtrOutput)
}

type SiteConnectionPtrOutput struct{ *pulumi.OutputState }

func (SiteConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConnection)(nil))
}

func (o SiteConnectionPtrOutput) ToSiteConnectionPtrOutput() SiteConnectionPtrOutput {
	return o
}

func (o SiteConnectionPtrOutput) ToSiteConnectionPtrOutputWithContext(ctx context.Context) SiteConnectionPtrOutput {
	return o
}

func (o SiteConnectionPtrOutput) Elem() SiteConnectionOutput {
	return o.ApplyT(func(v *SiteConnection) SiteConnection {
		if v != nil {
			return *v
		}
		var ret SiteConnection
		return ret
	}).(SiteConnectionOutput)
}

type SiteConnectionArrayOutput struct{ *pulumi.OutputState }

func (SiteConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SiteConnection)(nil))
}

func (o SiteConnectionArrayOutput) ToSiteConnectionArrayOutput() SiteConnectionArrayOutput {
	return o
}

func (o SiteConnectionArrayOutput) ToSiteConnectionArrayOutputWithContext(ctx context.Context) SiteConnectionArrayOutput {
	return o
}

func (o SiteConnectionArrayOutput) Index(i pulumi.IntInput) SiteConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SiteConnection {
		return vs[0].([]SiteConnection)[vs[1].(int)]
	}).(SiteConnectionOutput)
}

type SiteConnectionMapOutput struct{ *pulumi.OutputState }

func (SiteConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SiteConnection)(nil))
}

func (o SiteConnectionMapOutput) ToSiteConnectionMapOutput() SiteConnectionMapOutput {
	return o
}

func (o SiteConnectionMapOutput) ToSiteConnectionMapOutputWithContext(ctx context.Context) SiteConnectionMapOutput {
	return o
}

func (o SiteConnectionMapOutput) MapIndex(k pulumi.StringInput) SiteConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SiteConnection {
		return vs[0].(map[string]SiteConnection)[vs[1].(string)]
	}).(SiteConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteConnectionInput)(nil)).Elem(), &SiteConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteConnectionPtrInput)(nil)).Elem(), &SiteConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteConnectionArrayInput)(nil)).Elem(), SiteConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteConnectionMapInput)(nil)).Elem(), SiteConnectionMap{})
	pulumi.RegisterOutputType(SiteConnectionOutput{})
	pulumi.RegisterOutputType(SiteConnectionPtrOutput{})
	pulumi.RegisterOutputType(SiteConnectionArrayOutput{})
	pulumi.RegisterOutputType(SiteConnectionMapOutput{})
}
