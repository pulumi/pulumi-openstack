// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches a Network Interface (a Port) to an Instance using the OpenStack
// Compute (Nova) v2 API.
//
// ## Example Usage
//
// ### Basic Attachment
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.NewNetwork(ctx, "network_1", &networking.NetworkArgs{
//				Name:         pulumi.String("network_1"),
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			instance1, err := compute.NewInstance(ctx, "instance_1", &compute.InstanceArgs{
//				Name: pulumi.String("instance_1"),
//				SecurityGroups: pulumi.StringArray{
//					pulumi.String("default"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInterfaceAttach(ctx, "ai_1", &compute.InterfaceAttachArgs{
//				InstanceId: instance1.ID(),
//				NetworkId:  pulumi.Any(network1OpenstackNetworkingPortV2.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Attachment Specifying a Fixed IP
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.NewNetwork(ctx, "network_1", &networking.NetworkArgs{
//				Name:         pulumi.String("network_1"),
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			instance1, err := compute.NewInstance(ctx, "instance_1", &compute.InstanceArgs{
//				Name: pulumi.String("instance_1"),
//				SecurityGroups: pulumi.StringArray{
//					pulumi.String("default"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInterfaceAttach(ctx, "ai_1", &compute.InterfaceAttachArgs{
//				InstanceId: instance1.ID(),
//				NetworkId:  pulumi.Any(network1OpenstackNetworkingPortV2.Id),
//				FixedIp:    pulumi.String("10.0.10.10"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Attachment Using an Existing Port
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network1, err := networking.NewNetwork(ctx, "network_1", &networking.NetworkArgs{
//				Name:         pulumi.String("network_1"),
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			port1, err := networking.NewPort(ctx, "port_1", &networking.PortArgs{
//				Name:         pulumi.String("port_1"),
//				NetworkId:    network1.ID(),
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			instance1, err := compute.NewInstance(ctx, "instance_1", &compute.InstanceArgs{
//				Name: pulumi.String("instance_1"),
//				SecurityGroups: pulumi.StringArray{
//					pulumi.String("default"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInterfaceAttach(ctx, "ai_1", &compute.InterfaceAttachArgs{
//				InstanceId: instance1.ID(),
//				PortId:     port1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Interface Attachments can be imported using the Instance ID and Port ID
// separated by a slash, e.g.
//
// ```sh
// $ pulumi import openstack:compute/interfaceAttach:InterfaceAttach ai_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
// ```
type InterfaceAttach struct {
	pulumi.CustomResourceState

	// An IP address to assosciate with the port.
	// _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
	FixedIp pulumi.StringOutput `pulumi:"fixedIp"`
	// The ID of the Instance to attach the Port or Network to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The ID of the Network to attach to an Instance. A port will be created automatically.
	// _NOTE_: This option and `portId` are mutually exclusive.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// The ID of the Port to attach to an Instance.
	// _NOTE_: This option and `networkId` are mutually exclusive.
	PortId pulumi.StringOutput `pulumi:"portId"`
	// The region in which to create the interface attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new attachment.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewInterfaceAttach registers a new resource with the given unique name, arguments, and options.
func NewInterfaceAttach(ctx *pulumi.Context,
	name string, args *InterfaceAttachArgs, opts ...pulumi.ResourceOption) (*InterfaceAttach, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InterfaceAttach
	err := ctx.RegisterResource("openstack:compute/interfaceAttach:InterfaceAttach", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterfaceAttach gets an existing InterfaceAttach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterfaceAttach(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterfaceAttachState, opts ...pulumi.ResourceOption) (*InterfaceAttach, error) {
	var resource InterfaceAttach
	err := ctx.ReadResource("openstack:compute/interfaceAttach:InterfaceAttach", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterfaceAttach resources.
type interfaceAttachState struct {
	// An IP address to assosciate with the port.
	// _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
	FixedIp *string `pulumi:"fixedIp"`
	// The ID of the Instance to attach the Port or Network to.
	InstanceId *string `pulumi:"instanceId"`
	// The ID of the Network to attach to an Instance. A port will be created automatically.
	// _NOTE_: This option and `portId` are mutually exclusive.
	NetworkId *string `pulumi:"networkId"`
	// The ID of the Port to attach to an Instance.
	// _NOTE_: This option and `networkId` are mutually exclusive.
	PortId *string `pulumi:"portId"`
	// The region in which to create the interface attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new attachment.
	Region *string `pulumi:"region"`
}

type InterfaceAttachState struct {
	// An IP address to assosciate with the port.
	// _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
	FixedIp pulumi.StringPtrInput
	// The ID of the Instance to attach the Port or Network to.
	InstanceId pulumi.StringPtrInput
	// The ID of the Network to attach to an Instance. A port will be created automatically.
	// _NOTE_: This option and `portId` are mutually exclusive.
	NetworkId pulumi.StringPtrInput
	// The ID of the Port to attach to an Instance.
	// _NOTE_: This option and `networkId` are mutually exclusive.
	PortId pulumi.StringPtrInput
	// The region in which to create the interface attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new attachment.
	Region pulumi.StringPtrInput
}

func (InterfaceAttachState) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceAttachState)(nil)).Elem()
}

type interfaceAttachArgs struct {
	// An IP address to assosciate with the port.
	// _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
	FixedIp *string `pulumi:"fixedIp"`
	// The ID of the Instance to attach the Port or Network to.
	InstanceId string `pulumi:"instanceId"`
	// The ID of the Network to attach to an Instance. A port will be created automatically.
	// _NOTE_: This option and `portId` are mutually exclusive.
	NetworkId *string `pulumi:"networkId"`
	// The ID of the Port to attach to an Instance.
	// _NOTE_: This option and `networkId` are mutually exclusive.
	PortId *string `pulumi:"portId"`
	// The region in which to create the interface attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new attachment.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a InterfaceAttach resource.
type InterfaceAttachArgs struct {
	// An IP address to assosciate with the port.
	// _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
	FixedIp pulumi.StringPtrInput
	// The ID of the Instance to attach the Port or Network to.
	InstanceId pulumi.StringInput
	// The ID of the Network to attach to an Instance. A port will be created automatically.
	// _NOTE_: This option and `portId` are mutually exclusive.
	NetworkId pulumi.StringPtrInput
	// The ID of the Port to attach to an Instance.
	// _NOTE_: This option and `networkId` are mutually exclusive.
	PortId pulumi.StringPtrInput
	// The region in which to create the interface attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new attachment.
	Region pulumi.StringPtrInput
}

func (InterfaceAttachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceAttachArgs)(nil)).Elem()
}

type InterfaceAttachInput interface {
	pulumi.Input

	ToInterfaceAttachOutput() InterfaceAttachOutput
	ToInterfaceAttachOutputWithContext(ctx context.Context) InterfaceAttachOutput
}

func (*InterfaceAttach) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceAttach)(nil)).Elem()
}

func (i *InterfaceAttach) ToInterfaceAttachOutput() InterfaceAttachOutput {
	return i.ToInterfaceAttachOutputWithContext(context.Background())
}

func (i *InterfaceAttach) ToInterfaceAttachOutputWithContext(ctx context.Context) InterfaceAttachOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceAttachOutput)
}

// InterfaceAttachArrayInput is an input type that accepts InterfaceAttachArray and InterfaceAttachArrayOutput values.
// You can construct a concrete instance of `InterfaceAttachArrayInput` via:
//
//	InterfaceAttachArray{ InterfaceAttachArgs{...} }
type InterfaceAttachArrayInput interface {
	pulumi.Input

	ToInterfaceAttachArrayOutput() InterfaceAttachArrayOutput
	ToInterfaceAttachArrayOutputWithContext(context.Context) InterfaceAttachArrayOutput
}

type InterfaceAttachArray []InterfaceAttachInput

func (InterfaceAttachArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InterfaceAttach)(nil)).Elem()
}

func (i InterfaceAttachArray) ToInterfaceAttachArrayOutput() InterfaceAttachArrayOutput {
	return i.ToInterfaceAttachArrayOutputWithContext(context.Background())
}

func (i InterfaceAttachArray) ToInterfaceAttachArrayOutputWithContext(ctx context.Context) InterfaceAttachArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceAttachArrayOutput)
}

// InterfaceAttachMapInput is an input type that accepts InterfaceAttachMap and InterfaceAttachMapOutput values.
// You can construct a concrete instance of `InterfaceAttachMapInput` via:
//
//	InterfaceAttachMap{ "key": InterfaceAttachArgs{...} }
type InterfaceAttachMapInput interface {
	pulumi.Input

	ToInterfaceAttachMapOutput() InterfaceAttachMapOutput
	ToInterfaceAttachMapOutputWithContext(context.Context) InterfaceAttachMapOutput
}

type InterfaceAttachMap map[string]InterfaceAttachInput

func (InterfaceAttachMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InterfaceAttach)(nil)).Elem()
}

func (i InterfaceAttachMap) ToInterfaceAttachMapOutput() InterfaceAttachMapOutput {
	return i.ToInterfaceAttachMapOutputWithContext(context.Background())
}

func (i InterfaceAttachMap) ToInterfaceAttachMapOutputWithContext(ctx context.Context) InterfaceAttachMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceAttachMapOutput)
}

type InterfaceAttachOutput struct{ *pulumi.OutputState }

func (InterfaceAttachOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceAttach)(nil)).Elem()
}

func (o InterfaceAttachOutput) ToInterfaceAttachOutput() InterfaceAttachOutput {
	return o
}

func (o InterfaceAttachOutput) ToInterfaceAttachOutputWithContext(ctx context.Context) InterfaceAttachOutput {
	return o
}

// An IP address to assosciate with the port.
// _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
func (o InterfaceAttachOutput) FixedIp() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceAttach) pulumi.StringOutput { return v.FixedIp }).(pulumi.StringOutput)
}

// The ID of the Instance to attach the Port or Network to.
func (o InterfaceAttachOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceAttach) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The ID of the Network to attach to an Instance. A port will be created automatically.
// _NOTE_: This option and `portId` are mutually exclusive.
func (o InterfaceAttachOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceAttach) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// The ID of the Port to attach to an Instance.
// _NOTE_: This option and `networkId` are mutually exclusive.
func (o InterfaceAttachOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceAttach) pulumi.StringOutput { return v.PortId }).(pulumi.StringOutput)
}

// The region in which to create the interface attachment.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new attachment.
func (o InterfaceAttachOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceAttach) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type InterfaceAttachArrayOutput struct{ *pulumi.OutputState }

func (InterfaceAttachArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InterfaceAttach)(nil)).Elem()
}

func (o InterfaceAttachArrayOutput) ToInterfaceAttachArrayOutput() InterfaceAttachArrayOutput {
	return o
}

func (o InterfaceAttachArrayOutput) ToInterfaceAttachArrayOutputWithContext(ctx context.Context) InterfaceAttachArrayOutput {
	return o
}

func (o InterfaceAttachArrayOutput) Index(i pulumi.IntInput) InterfaceAttachOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InterfaceAttach {
		return vs[0].([]*InterfaceAttach)[vs[1].(int)]
	}).(InterfaceAttachOutput)
}

type InterfaceAttachMapOutput struct{ *pulumi.OutputState }

func (InterfaceAttachMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InterfaceAttach)(nil)).Elem()
}

func (o InterfaceAttachMapOutput) ToInterfaceAttachMapOutput() InterfaceAttachMapOutput {
	return o
}

func (o InterfaceAttachMapOutput) ToInterfaceAttachMapOutputWithContext(ctx context.Context) InterfaceAttachMapOutput {
	return o
}

func (o InterfaceAttachMapOutput) MapIndex(k pulumi.StringInput) InterfaceAttachOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InterfaceAttach {
		return vs[0].(map[string]*InterfaceAttach)[vs[1].(string)]
	}).(InterfaceAttachOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceAttachInput)(nil)).Elem(), &InterfaceAttach{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceAttachArrayInput)(nil)).Elem(), InterfaceAttachArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InterfaceAttachMapInput)(nil)).Elem(), InterfaceAttachMap{})
	pulumi.RegisterOutputType(InterfaceAttachOutput{})
	pulumi.RegisterOutputType(InterfaceAttachArrayOutput{})
	pulumi.RegisterOutputType(InterfaceAttachMapOutput{})
}
