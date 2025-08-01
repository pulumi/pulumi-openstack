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

// Manages a V2 portforwarding resource within OpenStack.
//
// ## Example Usage
//
// ### Simple portforwarding
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
//			_, err := networking.NewPortForwardingV2(ctx, "pf_1", &networking.PortForwardingV2Args{
//				FloatingipId:   pulumi.String("7a52eb59-7d47-415d-a884-046666a6fbae"),
//				ExternalPort:   pulumi.Int(7233),
//				InternalPort:   pulumi.Int(25),
//				InternalPortId: pulumi.String("b930d7f6-ceb7-40a0-8b81-a425dd994ccf"),
//				Protocol:       pulumi.String("tcp"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type PortForwardingV2 struct {
	pulumi.CustomResourceState

	// A text describing the port forwarding. Changing this
	// updates the `description` of an existing port forwarding.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The TCP/UDP/other protocol port number of the port forwarding. Changing this
	// updates the `externalPort` of an existing port forwarding.
	ExternalPort pulumi.IntOutput `pulumi:"externalPort"`
	// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
	FloatingipId pulumi.StringOutput `pulumi:"floatingipId"`
	// The fixed IPv4 address of the Neutron port associated with the port forwarding.
	// Changing this updates the `internalIpAddress` of an existing port forwarding.
	InternalIpAddress pulumi.StringOutput `pulumi:"internalIpAddress"`
	// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
	// port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
	InternalPort pulumi.IntOutput `pulumi:"internalPort"`
	// The ID of the Neutron port associated with the port forwarding. Changing
	// this updates the `internalPortId` of an existing port forwarding.
	InternalPortId pulumi.StringOutput `pulumi:"internalPortId"`
	// The IP protocol used in the port forwarding. Changing this updates the `protocol`
	// of an existing port forwarding.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port forwarding. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port forwarding.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewPortForwardingV2 registers a new resource with the given unique name, arguments, and options.
func NewPortForwardingV2(ctx *pulumi.Context,
	name string, args *PortForwardingV2Args, opts ...pulumi.ResourceOption) (*PortForwardingV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExternalPort == nil {
		return nil, errors.New("invalid value for required argument 'ExternalPort'")
	}
	if args.FloatingipId == nil {
		return nil, errors.New("invalid value for required argument 'FloatingipId'")
	}
	if args.InternalIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'InternalIpAddress'")
	}
	if args.InternalPort == nil {
		return nil, errors.New("invalid value for required argument 'InternalPort'")
	}
	if args.InternalPortId == nil {
		return nil, errors.New("invalid value for required argument 'InternalPortId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PortForwardingV2
	err := ctx.RegisterResource("openstack:networking/portForwardingV2:PortForwardingV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortForwardingV2 gets an existing PortForwardingV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortForwardingV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortForwardingV2State, opts ...pulumi.ResourceOption) (*PortForwardingV2, error) {
	var resource PortForwardingV2
	err := ctx.ReadResource("openstack:networking/portForwardingV2:PortForwardingV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortForwardingV2 resources.
type portForwardingV2State struct {
	// A text describing the port forwarding. Changing this
	// updates the `description` of an existing port forwarding.
	Description *string `pulumi:"description"`
	// The TCP/UDP/other protocol port number of the port forwarding. Changing this
	// updates the `externalPort` of an existing port forwarding.
	ExternalPort *int `pulumi:"externalPort"`
	// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
	FloatingipId *string `pulumi:"floatingipId"`
	// The fixed IPv4 address of the Neutron port associated with the port forwarding.
	// Changing this updates the `internalIpAddress` of an existing port forwarding.
	InternalIpAddress *string `pulumi:"internalIpAddress"`
	// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
	// port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
	InternalPort *int `pulumi:"internalPort"`
	// The ID of the Neutron port associated with the port forwarding. Changing
	// this updates the `internalPortId` of an existing port forwarding.
	InternalPortId *string `pulumi:"internalPortId"`
	// The IP protocol used in the port forwarding. Changing this updates the `protocol`
	// of an existing port forwarding.
	Protocol *string `pulumi:"protocol"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port forwarding. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port forwarding.
	Region *string `pulumi:"region"`
}

type PortForwardingV2State struct {
	// A text describing the port forwarding. Changing this
	// updates the `description` of an existing port forwarding.
	Description pulumi.StringPtrInput
	// The TCP/UDP/other protocol port number of the port forwarding. Changing this
	// updates the `externalPort` of an existing port forwarding.
	ExternalPort pulumi.IntPtrInput
	// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
	FloatingipId pulumi.StringPtrInput
	// The fixed IPv4 address of the Neutron port associated with the port forwarding.
	// Changing this updates the `internalIpAddress` of an existing port forwarding.
	InternalIpAddress pulumi.StringPtrInput
	// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
	// port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
	InternalPort pulumi.IntPtrInput
	// The ID of the Neutron port associated with the port forwarding. Changing
	// this updates the `internalPortId` of an existing port forwarding.
	InternalPortId pulumi.StringPtrInput
	// The IP protocol used in the port forwarding. Changing this updates the `protocol`
	// of an existing port forwarding.
	Protocol pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port forwarding. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port forwarding.
	Region pulumi.StringPtrInput
}

func (PortForwardingV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*portForwardingV2State)(nil)).Elem()
}

type portForwardingV2Args struct {
	// A text describing the port forwarding. Changing this
	// updates the `description` of an existing port forwarding.
	Description *string `pulumi:"description"`
	// The TCP/UDP/other protocol port number of the port forwarding. Changing this
	// updates the `externalPort` of an existing port forwarding.
	ExternalPort int `pulumi:"externalPort"`
	// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
	FloatingipId string `pulumi:"floatingipId"`
	// The fixed IPv4 address of the Neutron port associated with the port forwarding.
	// Changing this updates the `internalIpAddress` of an existing port forwarding.
	InternalIpAddress string `pulumi:"internalIpAddress"`
	// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
	// port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
	InternalPort int `pulumi:"internalPort"`
	// The ID of the Neutron port associated with the port forwarding. Changing
	// this updates the `internalPortId` of an existing port forwarding.
	InternalPortId string `pulumi:"internalPortId"`
	// The IP protocol used in the port forwarding. Changing this updates the `protocol`
	// of an existing port forwarding.
	Protocol string `pulumi:"protocol"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port forwarding. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port forwarding.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a PortForwardingV2 resource.
type PortForwardingV2Args struct {
	// A text describing the port forwarding. Changing this
	// updates the `description` of an existing port forwarding.
	Description pulumi.StringPtrInput
	// The TCP/UDP/other protocol port number of the port forwarding. Changing this
	// updates the `externalPort` of an existing port forwarding.
	ExternalPort pulumi.IntInput
	// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
	FloatingipId pulumi.StringInput
	// The fixed IPv4 address of the Neutron port associated with the port forwarding.
	// Changing this updates the `internalIpAddress` of an existing port forwarding.
	InternalIpAddress pulumi.StringInput
	// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
	// port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
	InternalPort pulumi.IntInput
	// The ID of the Neutron port associated with the port forwarding. Changing
	// this updates the `internalPortId` of an existing port forwarding.
	InternalPortId pulumi.StringInput
	// The IP protocol used in the port forwarding. Changing this updates the `protocol`
	// of an existing port forwarding.
	Protocol pulumi.StringInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a port forwarding. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port forwarding.
	Region pulumi.StringPtrInput
}

func (PortForwardingV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*portForwardingV2Args)(nil)).Elem()
}

type PortForwardingV2Input interface {
	pulumi.Input

	ToPortForwardingV2Output() PortForwardingV2Output
	ToPortForwardingV2OutputWithContext(ctx context.Context) PortForwardingV2Output
}

func (*PortForwardingV2) ElementType() reflect.Type {
	return reflect.TypeOf((**PortForwardingV2)(nil)).Elem()
}

func (i *PortForwardingV2) ToPortForwardingV2Output() PortForwardingV2Output {
	return i.ToPortForwardingV2OutputWithContext(context.Background())
}

func (i *PortForwardingV2) ToPortForwardingV2OutputWithContext(ctx context.Context) PortForwardingV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(PortForwardingV2Output)
}

// PortForwardingV2ArrayInput is an input type that accepts PortForwardingV2Array and PortForwardingV2ArrayOutput values.
// You can construct a concrete instance of `PortForwardingV2ArrayInput` via:
//
//	PortForwardingV2Array{ PortForwardingV2Args{...} }
type PortForwardingV2ArrayInput interface {
	pulumi.Input

	ToPortForwardingV2ArrayOutput() PortForwardingV2ArrayOutput
	ToPortForwardingV2ArrayOutputWithContext(context.Context) PortForwardingV2ArrayOutput
}

type PortForwardingV2Array []PortForwardingV2Input

func (PortForwardingV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortForwardingV2)(nil)).Elem()
}

func (i PortForwardingV2Array) ToPortForwardingV2ArrayOutput() PortForwardingV2ArrayOutput {
	return i.ToPortForwardingV2ArrayOutputWithContext(context.Background())
}

func (i PortForwardingV2Array) ToPortForwardingV2ArrayOutputWithContext(ctx context.Context) PortForwardingV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortForwardingV2ArrayOutput)
}

// PortForwardingV2MapInput is an input type that accepts PortForwardingV2Map and PortForwardingV2MapOutput values.
// You can construct a concrete instance of `PortForwardingV2MapInput` via:
//
//	PortForwardingV2Map{ "key": PortForwardingV2Args{...} }
type PortForwardingV2MapInput interface {
	pulumi.Input

	ToPortForwardingV2MapOutput() PortForwardingV2MapOutput
	ToPortForwardingV2MapOutputWithContext(context.Context) PortForwardingV2MapOutput
}

type PortForwardingV2Map map[string]PortForwardingV2Input

func (PortForwardingV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortForwardingV2)(nil)).Elem()
}

func (i PortForwardingV2Map) ToPortForwardingV2MapOutput() PortForwardingV2MapOutput {
	return i.ToPortForwardingV2MapOutputWithContext(context.Background())
}

func (i PortForwardingV2Map) ToPortForwardingV2MapOutputWithContext(ctx context.Context) PortForwardingV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortForwardingV2MapOutput)
}

type PortForwardingV2Output struct{ *pulumi.OutputState }

func (PortForwardingV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**PortForwardingV2)(nil)).Elem()
}

func (o PortForwardingV2Output) ToPortForwardingV2Output() PortForwardingV2Output {
	return o
}

func (o PortForwardingV2Output) ToPortForwardingV2OutputWithContext(ctx context.Context) PortForwardingV2Output {
	return o
}

// A text describing the port forwarding. Changing this
// updates the `description` of an existing port forwarding.
func (o PortForwardingV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PortForwardingV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The TCP/UDP/other protocol port number of the port forwarding. Changing this
// updates the `externalPort` of an existing port forwarding.
func (o PortForwardingV2Output) ExternalPort() pulumi.IntOutput {
	return o.ApplyT(func(v *PortForwardingV2) pulumi.IntOutput { return v.ExternalPort }).(pulumi.IntOutput)
}

// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
func (o PortForwardingV2Output) FloatingipId() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForwardingV2) pulumi.StringOutput { return v.FloatingipId }).(pulumi.StringOutput)
}

// The fixed IPv4 address of the Neutron port associated with the port forwarding.
// Changing this updates the `internalIpAddress` of an existing port forwarding.
func (o PortForwardingV2Output) InternalIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForwardingV2) pulumi.StringOutput { return v.InternalIpAddress }).(pulumi.StringOutput)
}

// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
// port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
func (o PortForwardingV2Output) InternalPort() pulumi.IntOutput {
	return o.ApplyT(func(v *PortForwardingV2) pulumi.IntOutput { return v.InternalPort }).(pulumi.IntOutput)
}

// The ID of the Neutron port associated with the port forwarding. Changing
// this updates the `internalPortId` of an existing port forwarding.
func (o PortForwardingV2Output) InternalPortId() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForwardingV2) pulumi.StringOutput { return v.InternalPortId }).(pulumi.StringOutput)
}

// The IP protocol used in the port forwarding. Changing this updates the `protocol`
// of an existing port forwarding.
func (o PortForwardingV2Output) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForwardingV2) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to create a port forwarding. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// port forwarding.
func (o PortForwardingV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PortForwardingV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type PortForwardingV2ArrayOutput struct{ *pulumi.OutputState }

func (PortForwardingV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortForwardingV2)(nil)).Elem()
}

func (o PortForwardingV2ArrayOutput) ToPortForwardingV2ArrayOutput() PortForwardingV2ArrayOutput {
	return o
}

func (o PortForwardingV2ArrayOutput) ToPortForwardingV2ArrayOutputWithContext(ctx context.Context) PortForwardingV2ArrayOutput {
	return o
}

func (o PortForwardingV2ArrayOutput) Index(i pulumi.IntInput) PortForwardingV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PortForwardingV2 {
		return vs[0].([]*PortForwardingV2)[vs[1].(int)]
	}).(PortForwardingV2Output)
}

type PortForwardingV2MapOutput struct{ *pulumi.OutputState }

func (PortForwardingV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortForwardingV2)(nil)).Elem()
}

func (o PortForwardingV2MapOutput) ToPortForwardingV2MapOutput() PortForwardingV2MapOutput {
	return o
}

func (o PortForwardingV2MapOutput) ToPortForwardingV2MapOutputWithContext(ctx context.Context) PortForwardingV2MapOutput {
	return o
}

func (o PortForwardingV2MapOutput) MapIndex(k pulumi.StringInput) PortForwardingV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PortForwardingV2 {
		return vs[0].(map[string]*PortForwardingV2)[vs[1].(string)]
	}).(PortForwardingV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortForwardingV2Input)(nil)).Elem(), &PortForwardingV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortForwardingV2ArrayInput)(nil)).Elem(), PortForwardingV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortForwardingV2MapInput)(nil)).Elem(), PortForwardingV2Map{})
	pulumi.RegisterOutputType(PortForwardingV2Output{})
	pulumi.RegisterOutputType(PortForwardingV2ArrayOutput{})
	pulumi.RegisterOutputType(PortForwardingV2MapOutput{})
}
