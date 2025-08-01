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

// Manages a V2 neutron security group rule resource within OpenStack.
// Unlike Nova security groups, neutron separates the group from the rules
// and also allows an admin to target a specific tenant_id.
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
//			secgroup1, err := networking.NewSecGroup(ctx, "secgroup_1", &networking.SecGroupArgs{
//				Name:        pulumi.String("secgroup_1"),
//				Description: pulumi.String("My neutron security group"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewSecGroupRule(ctx, "secgroup_rule_1", &networking.SecGroupRuleArgs{
//				Direction:       pulumi.String("ingress"),
//				Ethertype:       pulumi.String("IPv4"),
//				Protocol:        pulumi.String("tcp"),
//				PortRangeMin:    pulumi.Int(22),
//				PortRangeMax:    pulumi.Int(22),
//				RemoteIpPrefix:  pulumi.String("0.0.0.0/0"),
//				SecurityGroupId: secgroup1.ID(),
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
// > **Note:** To expose the full port-range 1:65535, use `0` for `portRangeMin`
// and `portRangeMax`.
//
// ## Import
//
// Security Group Rules can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:networking/secGroupRule:SecGroupRule secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745
// ```
type SecGroupRule struct {
	pulumi.CustomResourceState

	// A description of the rule. Changing this creates a new security group rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The direction of the rule, valid values are __ingress__
	// or __egress__. Changing this creates a new security group rule.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// The layer 3 protocol type, valid values are __IPv4__
	// or __IPv6__. Changing this creates a new security group rule.
	Ethertype pulumi.StringOutput `pulumi:"ethertype"`
	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMax pulumi.IntPtrOutput `pulumi:"portRangeMax"`
	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMin pulumi.IntPtrOutput `pulumi:"portRangeMin"`
	// The layer 4 protocol type, valid values are
	// following. Changing this creates a new security group rule. This is required
	// if you want to specify a port range.
	// * empty string or omitted (any protocol)
	// * integer value between 0 and 255 (valid IP protocol number)
	// * __tcp__
	// * __udp__
	// * __icmp__
	// * __ah__
	// * __dccp__
	// * __egp__
	// * __esp__
	// * __gre__
	// * __igmp__
	// * __ipv6-encap__
	// * __ipv6-frag__
	// * __ipv6-icmp__
	// * __ipv6-nonxt__
	// * __ipv6-opts__
	// * __ipv6-route__
	// * __ospf__
	// * __pgm__
	// * __rsvp__
	// * __sctp__
	// * __udplite__
	// * __vrrp__
	// * __ipip__
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group rule.
	Region pulumi.StringOutput `pulumi:"region"`
	// The remote address group id, the value
	// needs to be an OpenStack ID of an address group in the same tenant. Changing
	// this creates a new security group rule. This argument is mutually exclusive
	// with `remoteIpPrefix` and `remoteGroupId`.
	RemoteAddressGroupId pulumi.StringOutput `pulumi:"remoteAddressGroupId"`
	// The remote group id, the value needs to be an
	// Openstack ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	RemoteGroupId pulumi.StringOutput `pulumi:"remoteGroupId"`
	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	RemoteIpPrefix pulumi.StringOutput `pulumi:"remoteIpPrefix"`
	// The security group id the rule should belong
	// to, the value needs to be an Openstack ID of a security group in the same
	// tenant. Changing this creates a new security group rule.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group rule.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewSecGroupRule registers a new resource with the given unique name, arguments, and options.
func NewSecGroupRule(ctx *pulumi.Context,
	name string, args *SecGroupRuleArgs, opts ...pulumi.ResourceOption) (*SecGroupRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.Ethertype == nil {
		return nil, errors.New("invalid value for required argument 'Ethertype'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecGroupRule
	err := ctx.RegisterResource("openstack:networking/secGroupRule:SecGroupRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecGroupRule gets an existing SecGroupRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecGroupRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecGroupRuleState, opts ...pulumi.ResourceOption) (*SecGroupRule, error) {
	var resource SecGroupRule
	err := ctx.ReadResource("openstack:networking/secGroupRule:SecGroupRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecGroupRule resources.
type secGroupRuleState struct {
	// A description of the rule. Changing this creates a new security group rule.
	Description *string `pulumi:"description"`
	// The direction of the rule, valid values are __ingress__
	// or __egress__. Changing this creates a new security group rule.
	Direction *string `pulumi:"direction"`
	// The layer 3 protocol type, valid values are __IPv4__
	// or __IPv6__. Changing this creates a new security group rule.
	Ethertype *string `pulumi:"ethertype"`
	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMax *int `pulumi:"portRangeMax"`
	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMin *int `pulumi:"portRangeMin"`
	// The layer 4 protocol type, valid values are
	// following. Changing this creates a new security group rule. This is required
	// if you want to specify a port range.
	// * empty string or omitted (any protocol)
	// * integer value between 0 and 255 (valid IP protocol number)
	// * __tcp__
	// * __udp__
	// * __icmp__
	// * __ah__
	// * __dccp__
	// * __egp__
	// * __esp__
	// * __gre__
	// * __igmp__
	// * __ipv6-encap__
	// * __ipv6-frag__
	// * __ipv6-icmp__
	// * __ipv6-nonxt__
	// * __ipv6-opts__
	// * __ipv6-route__
	// * __ospf__
	// * __pgm__
	// * __rsvp__
	// * __sctp__
	// * __udplite__
	// * __vrrp__
	// * __ipip__
	Protocol *string `pulumi:"protocol"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group rule.
	Region *string `pulumi:"region"`
	// The remote address group id, the value
	// needs to be an OpenStack ID of an address group in the same tenant. Changing
	// this creates a new security group rule. This argument is mutually exclusive
	// with `remoteIpPrefix` and `remoteGroupId`.
	RemoteAddressGroupId *string `pulumi:"remoteAddressGroupId"`
	// The remote group id, the value needs to be an
	// Openstack ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	RemoteGroupId *string `pulumi:"remoteGroupId"`
	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	RemoteIpPrefix *string `pulumi:"remoteIpPrefix"`
	// The security group id the rule should belong
	// to, the value needs to be an Openstack ID of a security group in the same
	// tenant. Changing this creates a new security group rule.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group rule.
	TenantId *string `pulumi:"tenantId"`
}

type SecGroupRuleState struct {
	// A description of the rule. Changing this creates a new security group rule.
	Description pulumi.StringPtrInput
	// The direction of the rule, valid values are __ingress__
	// or __egress__. Changing this creates a new security group rule.
	Direction pulumi.StringPtrInput
	// The layer 3 protocol type, valid values are __IPv4__
	// or __IPv6__. Changing this creates a new security group rule.
	Ethertype pulumi.StringPtrInput
	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMax pulumi.IntPtrInput
	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMin pulumi.IntPtrInput
	// The layer 4 protocol type, valid values are
	// following. Changing this creates a new security group rule. This is required
	// if you want to specify a port range.
	// * empty string or omitted (any protocol)
	// * integer value between 0 and 255 (valid IP protocol number)
	// * __tcp__
	// * __udp__
	// * __icmp__
	// * __ah__
	// * __dccp__
	// * __egp__
	// * __esp__
	// * __gre__
	// * __igmp__
	// * __ipv6-encap__
	// * __ipv6-frag__
	// * __ipv6-icmp__
	// * __ipv6-nonxt__
	// * __ipv6-opts__
	// * __ipv6-route__
	// * __ospf__
	// * __pgm__
	// * __rsvp__
	// * __sctp__
	// * __udplite__
	// * __vrrp__
	// * __ipip__
	Protocol pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group rule.
	Region pulumi.StringPtrInput
	// The remote address group id, the value
	// needs to be an OpenStack ID of an address group in the same tenant. Changing
	// this creates a new security group rule. This argument is mutually exclusive
	// with `remoteIpPrefix` and `remoteGroupId`.
	RemoteAddressGroupId pulumi.StringPtrInput
	// The remote group id, the value needs to be an
	// Openstack ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	RemoteGroupId pulumi.StringPtrInput
	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	RemoteIpPrefix pulumi.StringPtrInput
	// The security group id the rule should belong
	// to, the value needs to be an Openstack ID of a security group in the same
	// tenant. Changing this creates a new security group rule.
	SecurityGroupId pulumi.StringPtrInput
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group rule.
	TenantId pulumi.StringPtrInput
}

func (SecGroupRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secGroupRuleState)(nil)).Elem()
}

type secGroupRuleArgs struct {
	// A description of the rule. Changing this creates a new security group rule.
	Description *string `pulumi:"description"`
	// The direction of the rule, valid values are __ingress__
	// or __egress__. Changing this creates a new security group rule.
	Direction string `pulumi:"direction"`
	// The layer 3 protocol type, valid values are __IPv4__
	// or __IPv6__. Changing this creates a new security group rule.
	Ethertype string `pulumi:"ethertype"`
	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMax *int `pulumi:"portRangeMax"`
	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMin *int `pulumi:"portRangeMin"`
	// The layer 4 protocol type, valid values are
	// following. Changing this creates a new security group rule. This is required
	// if you want to specify a port range.
	// * empty string or omitted (any protocol)
	// * integer value between 0 and 255 (valid IP protocol number)
	// * __tcp__
	// * __udp__
	// * __icmp__
	// * __ah__
	// * __dccp__
	// * __egp__
	// * __esp__
	// * __gre__
	// * __igmp__
	// * __ipv6-encap__
	// * __ipv6-frag__
	// * __ipv6-icmp__
	// * __ipv6-nonxt__
	// * __ipv6-opts__
	// * __ipv6-route__
	// * __ospf__
	// * __pgm__
	// * __rsvp__
	// * __sctp__
	// * __udplite__
	// * __vrrp__
	// * __ipip__
	Protocol *string `pulumi:"protocol"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group rule.
	Region *string `pulumi:"region"`
	// The remote address group id, the value
	// needs to be an OpenStack ID of an address group in the same tenant. Changing
	// this creates a new security group rule. This argument is mutually exclusive
	// with `remoteIpPrefix` and `remoteGroupId`.
	RemoteAddressGroupId *string `pulumi:"remoteAddressGroupId"`
	// The remote group id, the value needs to be an
	// Openstack ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	RemoteGroupId *string `pulumi:"remoteGroupId"`
	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	RemoteIpPrefix *string `pulumi:"remoteIpPrefix"`
	// The security group id the rule should belong
	// to, the value needs to be an Openstack ID of a security group in the same
	// tenant. Changing this creates a new security group rule.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group rule.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a SecGroupRule resource.
type SecGroupRuleArgs struct {
	// A description of the rule. Changing this creates a new security group rule.
	Description pulumi.StringPtrInput
	// The direction of the rule, valid values are __ingress__
	// or __egress__. Changing this creates a new security group rule.
	Direction pulumi.StringInput
	// The layer 3 protocol type, valid values are __IPv4__
	// or __IPv6__. Changing this creates a new security group rule.
	Ethertype pulumi.StringInput
	// The higher part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMax pulumi.IntPtrInput
	// The lower part of the allowed port range, valid
	// integer value needs to be between 1 and 65535. Changing this creates a new
	// security group rule.
	PortRangeMin pulumi.IntPtrInput
	// The layer 4 protocol type, valid values are
	// following. Changing this creates a new security group rule. This is required
	// if you want to specify a port range.
	// * empty string or omitted (any protocol)
	// * integer value between 0 and 255 (valid IP protocol number)
	// * __tcp__
	// * __udp__
	// * __icmp__
	// * __ah__
	// * __dccp__
	// * __egp__
	// * __esp__
	// * __gre__
	// * __igmp__
	// * __ipv6-encap__
	// * __ipv6-frag__
	// * __ipv6-icmp__
	// * __ipv6-nonxt__
	// * __ipv6-opts__
	// * __ipv6-route__
	// * __ospf__
	// * __pgm__
	// * __rsvp__
	// * __sctp__
	// * __udplite__
	// * __vrrp__
	// * __ipip__
	Protocol pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group rule.
	Region pulumi.StringPtrInput
	// The remote address group id, the value
	// needs to be an OpenStack ID of an address group in the same tenant. Changing
	// this creates a new security group rule. This argument is mutually exclusive
	// with `remoteIpPrefix` and `remoteGroupId`.
	RemoteAddressGroupId pulumi.StringPtrInput
	// The remote group id, the value needs to be an
	// Openstack ID of a security group in the same tenant. Changing this creates
	// a new security group rule.
	RemoteGroupId pulumi.StringPtrInput
	// The remote CIDR, the value needs to be a valid
	// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
	RemoteIpPrefix pulumi.StringPtrInput
	// The security group id the rule should belong
	// to, the value needs to be an Openstack ID of a security group in the same
	// tenant. Changing this creates a new security group rule.
	SecurityGroupId pulumi.StringInput
	// The owner of the security group. Required if admin
	// wants to create a port for another tenant. Changing this creates a new
	// security group rule.
	TenantId pulumi.StringPtrInput
}

func (SecGroupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secGroupRuleArgs)(nil)).Elem()
}

type SecGroupRuleInput interface {
	pulumi.Input

	ToSecGroupRuleOutput() SecGroupRuleOutput
	ToSecGroupRuleOutputWithContext(ctx context.Context) SecGroupRuleOutput
}

func (*SecGroupRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecGroupRule)(nil)).Elem()
}

func (i *SecGroupRule) ToSecGroupRuleOutput() SecGroupRuleOutput {
	return i.ToSecGroupRuleOutputWithContext(context.Background())
}

func (i *SecGroupRule) ToSecGroupRuleOutputWithContext(ctx context.Context) SecGroupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupRuleOutput)
}

// SecGroupRuleArrayInput is an input type that accepts SecGroupRuleArray and SecGroupRuleArrayOutput values.
// You can construct a concrete instance of `SecGroupRuleArrayInput` via:
//
//	SecGroupRuleArray{ SecGroupRuleArgs{...} }
type SecGroupRuleArrayInput interface {
	pulumi.Input

	ToSecGroupRuleArrayOutput() SecGroupRuleArrayOutput
	ToSecGroupRuleArrayOutputWithContext(context.Context) SecGroupRuleArrayOutput
}

type SecGroupRuleArray []SecGroupRuleInput

func (SecGroupRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecGroupRule)(nil)).Elem()
}

func (i SecGroupRuleArray) ToSecGroupRuleArrayOutput() SecGroupRuleArrayOutput {
	return i.ToSecGroupRuleArrayOutputWithContext(context.Background())
}

func (i SecGroupRuleArray) ToSecGroupRuleArrayOutputWithContext(ctx context.Context) SecGroupRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupRuleArrayOutput)
}

// SecGroupRuleMapInput is an input type that accepts SecGroupRuleMap and SecGroupRuleMapOutput values.
// You can construct a concrete instance of `SecGroupRuleMapInput` via:
//
//	SecGroupRuleMap{ "key": SecGroupRuleArgs{...} }
type SecGroupRuleMapInput interface {
	pulumi.Input

	ToSecGroupRuleMapOutput() SecGroupRuleMapOutput
	ToSecGroupRuleMapOutputWithContext(context.Context) SecGroupRuleMapOutput
}

type SecGroupRuleMap map[string]SecGroupRuleInput

func (SecGroupRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecGroupRule)(nil)).Elem()
}

func (i SecGroupRuleMap) ToSecGroupRuleMapOutput() SecGroupRuleMapOutput {
	return i.ToSecGroupRuleMapOutputWithContext(context.Background())
}

func (i SecGroupRuleMap) ToSecGroupRuleMapOutputWithContext(ctx context.Context) SecGroupRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupRuleMapOutput)
}

type SecGroupRuleOutput struct{ *pulumi.OutputState }

func (SecGroupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecGroupRule)(nil)).Elem()
}

func (o SecGroupRuleOutput) ToSecGroupRuleOutput() SecGroupRuleOutput {
	return o
}

func (o SecGroupRuleOutput) ToSecGroupRuleOutputWithContext(ctx context.Context) SecGroupRuleOutput {
	return o
}

// A description of the rule. Changing this creates a new security group rule.
func (o SecGroupRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The direction of the rule, valid values are __ingress__
// or __egress__. Changing this creates a new security group rule.
func (o SecGroupRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// The layer 3 protocol type, valid values are __IPv4__
// or __IPv6__. Changing this creates a new security group rule.
func (o SecGroupRuleOutput) Ethertype() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringOutput { return v.Ethertype }).(pulumi.StringOutput)
}

// The higher part of the allowed port range, valid
// integer value needs to be between 1 and 65535. Changing this creates a new
// security group rule.
func (o SecGroupRuleOutput) PortRangeMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.IntPtrOutput { return v.PortRangeMax }).(pulumi.IntPtrOutput)
}

// The lower part of the allowed port range, valid
// integer value needs to be between 1 and 65535. Changing this creates a new
// security group rule.
func (o SecGroupRuleOutput) PortRangeMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.IntPtrOutput { return v.PortRangeMin }).(pulumi.IntPtrOutput)
}

// The layer 4 protocol type, valid values are
// following. Changing this creates a new security group rule. This is required
// if you want to specify a port range.
// * empty string or omitted (any protocol)
// * integer value between 0 and 255 (valid IP protocol number)
// * __tcp__
// * __udp__
// * __icmp__
// * __ah__
// * __dccp__
// * __egp__
// * __esp__
// * __gre__
// * __igmp__
// * __ipv6-encap__
// * __ipv6-frag__
// * __ipv6-icmp__
// * __ipv6-nonxt__
// * __ipv6-opts__
// * __ipv6-route__
// * __ospf__
// * __pgm__
// * __rsvp__
// * __sctp__
// * __udplite__
// * __vrrp__
// * __ipip__
func (o SecGroupRuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to create a port. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// security group rule.
func (o SecGroupRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The remote address group id, the value
// needs to be an OpenStack ID of an address group in the same tenant. Changing
// this creates a new security group rule. This argument is mutually exclusive
// with `remoteIpPrefix` and `remoteGroupId`.
func (o SecGroupRuleOutput) RemoteAddressGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringOutput { return v.RemoteAddressGroupId }).(pulumi.StringOutput)
}

// The remote group id, the value needs to be an
// Openstack ID of a security group in the same tenant. Changing this creates
// a new security group rule.
func (o SecGroupRuleOutput) RemoteGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringOutput { return v.RemoteGroupId }).(pulumi.StringOutput)
}

// The remote CIDR, the value needs to be a valid
// CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
func (o SecGroupRuleOutput) RemoteIpPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringOutput { return v.RemoteIpPrefix }).(pulumi.StringOutput)
}

// The security group id the rule should belong
// to, the value needs to be an Openstack ID of a security group in the same
// tenant. Changing this creates a new security group rule.
func (o SecGroupRuleOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// The owner of the security group. Required if admin
// wants to create a port for another tenant. Changing this creates a new
// security group rule.
func (o SecGroupRuleOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroupRule) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type SecGroupRuleArrayOutput struct{ *pulumi.OutputState }

func (SecGroupRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecGroupRule)(nil)).Elem()
}

func (o SecGroupRuleArrayOutput) ToSecGroupRuleArrayOutput() SecGroupRuleArrayOutput {
	return o
}

func (o SecGroupRuleArrayOutput) ToSecGroupRuleArrayOutputWithContext(ctx context.Context) SecGroupRuleArrayOutput {
	return o
}

func (o SecGroupRuleArrayOutput) Index(i pulumi.IntInput) SecGroupRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecGroupRule {
		return vs[0].([]*SecGroupRule)[vs[1].(int)]
	}).(SecGroupRuleOutput)
}

type SecGroupRuleMapOutput struct{ *pulumi.OutputState }

func (SecGroupRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecGroupRule)(nil)).Elem()
}

func (o SecGroupRuleMapOutput) ToSecGroupRuleMapOutput() SecGroupRuleMapOutput {
	return o
}

func (o SecGroupRuleMapOutput) ToSecGroupRuleMapOutputWithContext(ctx context.Context) SecGroupRuleMapOutput {
	return o
}

func (o SecGroupRuleMapOutput) MapIndex(k pulumi.StringInput) SecGroupRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecGroupRule {
		return vs[0].(map[string]*SecGroupRule)[vs[1].(string)]
	}).(SecGroupRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupRuleInput)(nil)).Elem(), &SecGroupRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupRuleArrayInput)(nil)).Elem(), SecGroupRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupRuleMapInput)(nil)).Elem(), SecGroupRuleMap{})
	pulumi.RegisterOutputType(SecGroupRuleOutput{})
	pulumi.RegisterOutputType(SecGroupRuleArrayOutput{})
	pulumi.RegisterOutputType(SecGroupRuleMapOutput{})
}
