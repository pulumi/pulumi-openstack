// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 Neutron network resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/compute"
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		network1, err := networking.NewNetwork(ctx, "network1", &networking.NetworkArgs{
// 			AdminStateUp: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		subnet1, err := networking.NewSubnet(ctx, "subnet1", &networking.SubnetArgs{
// 			Cidr:      pulumi.String("192.168.199.0/24"),
// 			IpVersion: pulumi.Int(4),
// 			NetworkId: network1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secgroup1, err := compute.NewSecGroup(ctx, "secgroup1", &compute.SecGroupArgs{
// 			Description: pulumi.String("a security group"),
// 			Rules: compute.SecGroupRuleArray{
// 				&compute.SecGroupRuleArgs{
// 					Cidr:       pulumi.String("0.0.0.0/0"),
// 					FromPort:   pulumi.Int(22),
// 					IpProtocol: pulumi.String("tcp"),
// 					ToPort:     pulumi.Int(22),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		port1, err := networking.NewPort(ctx, "port1", &networking.PortArgs{
// 			AdminStateUp: pulumi.Bool(true),
// 			FixedIps: networking.PortFixedIpArray{
// 				&networking.PortFixedIpArgs{
// 					IpAddress: pulumi.String("192.168.199.10"),
// 					SubnetId:  subnet1.ID(),
// 				},
// 			},
// 			NetworkId: network1.ID(),
// 			SecurityGroupIds: pulumi.StringArray{
// 				secgroup1.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewInstance(ctx, "instance1", &compute.InstanceArgs{
// 			Networks: compute.InstanceNetworkArray{
// 				&compute.InstanceNetworkArgs{
// 					Port: port1.ID(),
// 				},
// 			},
// 			SecurityGroups: pulumi.StringArray{
// 				secgroup1.Name,
// 			},
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
// Networks can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:networking/network:Network network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9
// ```
type Network struct {
	pulumi.CustomResourceState

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	AdminStateUp pulumi.BoolOutput `pulumi:"adminStateUp"`
	// The collection of tags assigned on the network, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	AvailabilityZoneHints pulumi.StringArrayOutput `pulumi:"availabilityZoneHints"`
	// Human-readable description of the network. Changing this
	// updates the name of the existing network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The network DNS domain. Available, when Neutron DNS
	// extension is enabled. The `dnsDomain` of a network in conjunction with the
	// `dnsName` attribute of its ports will be published in an external DNS
	// service when Neutron is configured to integrate with such a service.
	DnsDomain pulumi.StringOutput `pulumi:"dnsDomain"`
	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	External pulumi.BoolOutput `pulumi:"external"`
	// The network MTU. Available for read-only, when Neutron
	// `net-mtu` extension is enabled. Available for the modification, when
	// Neutron `net-mtu-writable` extension is enabled.
	Mtu pulumi.IntOutput `pulumi:"mtu"`
	// The name of the network. Changing this updates the name of
	// the existing network.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether to explicitly enable or disable
	// port security on the network. Port Security is usually enabled by default, so
	// omitting this argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. Valid values are `true` and
	// `false`.
	PortSecurityEnabled pulumi.BoolOutput `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringOutput `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// network.
	Region pulumi.StringOutput `pulumi:"region"`
	// An array of one or more provider segment objects.
	Segments NetworkSegmentArrayOutput `pulumi:"segments"`
	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabilities of the
	// existing network.
	Shared pulumi.BoolOutput `pulumi:"shared"`
	// A set of string tags for the network.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Specifies whether the network resource has the
	// VLAN transparent attribute set. Valid values are true and false. Defaults to
	// false. Changing this updates the `transparentVlan` attribute of the existing
	// network.
	TransparentVlan pulumi.BoolOutput `pulumi:"transparentVlan"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil {
		args = &NetworkArgs{}
	}

	var resource Network
	err := ctx.RegisterResource("openstack:networking/network:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("openstack:networking/network:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The collection of tags assigned on the network, which have been
	// explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	AvailabilityZoneHints []string `pulumi:"availabilityZoneHints"`
	// Human-readable description of the network. Changing this
	// updates the name of the existing network.
	Description *string `pulumi:"description"`
	// The network DNS domain. Available, when Neutron DNS
	// extension is enabled. The `dnsDomain` of a network in conjunction with the
	// `dnsName` attribute of its ports will be published in an external DNS
	// service when Neutron is configured to integrate with such a service.
	DnsDomain *string `pulumi:"dnsDomain"`
	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	External *bool `pulumi:"external"`
	// The network MTU. Available for read-only, when Neutron
	// `net-mtu` extension is enabled. Available for the modification, when
	// Neutron `net-mtu-writable` extension is enabled.
	Mtu *int `pulumi:"mtu"`
	// The name of the network. Changing this updates the name of
	// the existing network.
	Name *string `pulumi:"name"`
	// Whether to explicitly enable or disable
	// port security on the network. Port Security is usually enabled by default, so
	// omitting this argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. Valid values are `true` and
	// `false`.
	PortSecurityEnabled *bool `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId *string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// network.
	Region *string `pulumi:"region"`
	// An array of one or more provider segment objects.
	Segments []NetworkSegment `pulumi:"segments"`
	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabilities of the
	// existing network.
	Shared *bool `pulumi:"shared"`
	// A set of string tags for the network.
	Tags []string `pulumi:"tags"`
	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	TenantId *string `pulumi:"tenantId"`
	// Specifies whether the network resource has the
	// VLAN transparent attribute set. Valid values are true and false. Defaults to
	// false. Changing this updates the `transparentVlan` attribute of the existing
	// network.
	TransparentVlan *bool `pulumi:"transparentVlan"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type NetworkState struct {
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	AdminStateUp pulumi.BoolPtrInput
	// The collection of tags assigned on the network, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	AvailabilityZoneHints pulumi.StringArrayInput
	// Human-readable description of the network. Changing this
	// updates the name of the existing network.
	Description pulumi.StringPtrInput
	// The network DNS domain. Available, when Neutron DNS
	// extension is enabled. The `dnsDomain` of a network in conjunction with the
	// `dnsName` attribute of its ports will be published in an external DNS
	// service when Neutron is configured to integrate with such a service.
	DnsDomain pulumi.StringPtrInput
	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	External pulumi.BoolPtrInput
	// The network MTU. Available for read-only, when Neutron
	// `net-mtu` extension is enabled. Available for the modification, when
	// Neutron `net-mtu-writable` extension is enabled.
	Mtu pulumi.IntPtrInput
	// The name of the network. Changing this updates the name of
	// the existing network.
	Name pulumi.StringPtrInput
	// Whether to explicitly enable or disable
	// port security on the network. Port Security is usually enabled by default, so
	// omitting this argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. Valid values are `true` and
	// `false`.
	PortSecurityEnabled pulumi.BoolPtrInput
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// network.
	Region pulumi.StringPtrInput
	// An array of one or more provider segment objects.
	Segments NetworkSegmentArrayInput
	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabilities of the
	// existing network.
	Shared pulumi.BoolPtrInput
	// A set of string tags for the network.
	Tags pulumi.StringArrayInput
	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	TenantId pulumi.StringPtrInput
	// Specifies whether the network resource has the
	// VLAN transparent attribute set. Valid values are true and false. Defaults to
	// false. Changing this updates the `transparentVlan` attribute of the existing
	// network.
	TransparentVlan pulumi.BoolPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	AvailabilityZoneHints []string `pulumi:"availabilityZoneHints"`
	// Human-readable description of the network. Changing this
	// updates the name of the existing network.
	Description *string `pulumi:"description"`
	// The network DNS domain. Available, when Neutron DNS
	// extension is enabled. The `dnsDomain` of a network in conjunction with the
	// `dnsName` attribute of its ports will be published in an external DNS
	// service when Neutron is configured to integrate with such a service.
	DnsDomain *string `pulumi:"dnsDomain"`
	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	External *bool `pulumi:"external"`
	// The network MTU. Available for read-only, when Neutron
	// `net-mtu` extension is enabled. Available for the modification, when
	// Neutron `net-mtu-writable` extension is enabled.
	Mtu *int `pulumi:"mtu"`
	// The name of the network. Changing this updates the name of
	// the existing network.
	Name *string `pulumi:"name"`
	// Whether to explicitly enable or disable
	// port security on the network. Port Security is usually enabled by default, so
	// omitting this argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. Valid values are `true` and
	// `false`.
	PortSecurityEnabled *bool `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId *string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// network.
	Region *string `pulumi:"region"`
	// An array of one or more provider segment objects.
	Segments []NetworkSegment `pulumi:"segments"`
	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabilities of the
	// existing network.
	Shared *bool `pulumi:"shared"`
	// A set of string tags for the network.
	Tags []string `pulumi:"tags"`
	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	TenantId *string `pulumi:"tenantId"`
	// Specifies whether the network resource has the
	// VLAN transparent attribute set. Valid values are true and false. Defaults to
	// false. Changing this updates the `transparentVlan` attribute of the existing
	// network.
	TransparentVlan *bool `pulumi:"transparentVlan"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	AdminStateUp pulumi.BoolPtrInput
	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	AvailabilityZoneHints pulumi.StringArrayInput
	// Human-readable description of the network. Changing this
	// updates the name of the existing network.
	Description pulumi.StringPtrInput
	// The network DNS domain. Available, when Neutron DNS
	// extension is enabled. The `dnsDomain` of a network in conjunction with the
	// `dnsName` attribute of its ports will be published in an external DNS
	// service when Neutron is configured to integrate with such a service.
	DnsDomain pulumi.StringPtrInput
	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	External pulumi.BoolPtrInput
	// The network MTU. Available for read-only, when Neutron
	// `net-mtu` extension is enabled. Available for the modification, when
	// Neutron `net-mtu-writable` extension is enabled.
	Mtu pulumi.IntPtrInput
	// The name of the network. Changing this updates the name of
	// the existing network.
	Name pulumi.StringPtrInput
	// Whether to explicitly enable or disable
	// port security on the network. Port Security is usually enabled by default, so
	// omitting this argument will usually result in a value of "true". Setting this
	// explicitly to `false` will disable port security. Valid values are `true` and
	// `false`.
	PortSecurityEnabled pulumi.BoolPtrInput
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// network.
	Region pulumi.StringPtrInput
	// An array of one or more provider segment objects.
	Segments NetworkSegmentArrayInput
	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabilities of the
	// existing network.
	Shared pulumi.BoolPtrInput
	// A set of string tags for the network.
	Tags pulumi.StringArrayInput
	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	TenantId pulumi.StringPtrInput
	// Specifies whether the network resource has the
	// VLAN transparent attribute set. Valid values are true and false. Defaults to
	// false. Changing this updates the `transparentVlan` attribute of the existing
	// network.
	TransparentVlan pulumi.BoolPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}

type NetworkInput interface {
	pulumi.Input

	ToNetworkOutput() NetworkOutput
	ToNetworkOutputWithContext(ctx context.Context) NetworkOutput
}

func (*Network) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (i *Network) ToNetworkOutput() NetworkOutput {
	return i.ToNetworkOutputWithContext(context.Background())
}

func (i *Network) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkOutput)
}

func (i *Network) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i *Network) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPtrOutput)
}

type NetworkPtrInput interface {
	pulumi.Input

	ToNetworkPtrOutput() NetworkPtrOutput
	ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput
}

type networkPtrType NetworkArgs

func (*networkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil))
}

func (i *networkPtrType) ToNetworkPtrOutput() NetworkPtrOutput {
	return i.ToNetworkPtrOutputWithContext(context.Background())
}

func (i *networkPtrType) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPtrOutput)
}

// NetworkArrayInput is an input type that accepts NetworkArray and NetworkArrayOutput values.
// You can construct a concrete instance of `NetworkArrayInput` via:
//
//          NetworkArray{ NetworkArgs{...} }
type NetworkArrayInput interface {
	pulumi.Input

	ToNetworkArrayOutput() NetworkArrayOutput
	ToNetworkArrayOutputWithContext(context.Context) NetworkArrayOutput
}

type NetworkArray []NetworkInput

func (NetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Network)(nil))
}

func (i NetworkArray) ToNetworkArrayOutput() NetworkArrayOutput {
	return i.ToNetworkArrayOutputWithContext(context.Background())
}

func (i NetworkArray) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkArrayOutput)
}

// NetworkMapInput is an input type that accepts NetworkMap and NetworkMapOutput values.
// You can construct a concrete instance of `NetworkMapInput` via:
//
//          NetworkMap{ "key": NetworkArgs{...} }
type NetworkMapInput interface {
	pulumi.Input

	ToNetworkMapOutput() NetworkMapOutput
	ToNetworkMapOutputWithContext(context.Context) NetworkMapOutput
}

type NetworkMap map[string]NetworkInput

func (NetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Network)(nil))
}

func (i NetworkMap) ToNetworkMapOutput() NetworkMapOutput {
	return i.ToNetworkMapOutputWithContext(context.Background())
}

func (i NetworkMap) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMapOutput)
}

type NetworkOutput struct {
	*pulumi.OutputState
}

func (NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Network)(nil))
}

func (o NetworkOutput) ToNetworkOutput() NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkOutputWithContext(ctx context.Context) NetworkOutput {
	return o
}

func (o NetworkOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o.ToNetworkPtrOutputWithContext(context.Background())
}

func (o NetworkOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o.ApplyT(func(v Network) *Network {
		return &v
	}).(NetworkPtrOutput)
}

type NetworkPtrOutput struct {
	*pulumi.OutputState
}

func (NetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Network)(nil))
}

func (o NetworkPtrOutput) ToNetworkPtrOutput() NetworkPtrOutput {
	return o
}

func (o NetworkPtrOutput) ToNetworkPtrOutputWithContext(ctx context.Context) NetworkPtrOutput {
	return o
}

type NetworkArrayOutput struct{ *pulumi.OutputState }

func (NetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Network)(nil))
}

func (o NetworkArrayOutput) ToNetworkArrayOutput() NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) ToNetworkArrayOutputWithContext(ctx context.Context) NetworkArrayOutput {
	return o
}

func (o NetworkArrayOutput) Index(i pulumi.IntInput) NetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Network {
		return vs[0].([]Network)[vs[1].(int)]
	}).(NetworkOutput)
}

type NetworkMapOutput struct{ *pulumi.OutputState }

func (NetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Network)(nil))
}

func (o NetworkMapOutput) ToNetworkMapOutput() NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) ToNetworkMapOutputWithContext(ctx context.Context) NetworkMapOutput {
	return o
}

func (o NetworkMapOutput) MapIndex(k pulumi.StringInput) NetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Network {
		return vs[0].(map[string]Network)[vs[1].(string)]
	}).(NetworkOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkOutput{})
	pulumi.RegisterOutputType(NetworkPtrOutput{})
	pulumi.RegisterOutputType(NetworkArrayOutput{})
	pulumi.RegisterOutputType(NetworkMapOutput{})
}
