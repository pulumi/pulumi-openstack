// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches a Block Storage Volume to an Instance using the OpenStack
// Compute (Nova) v2 API.
//
// ## Example Usage
// ### Basic attachment of a single volume to a single instance
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		volume1, err := blockstorage.NewVolumeV2(ctx, "volume1", &blockstorage.VolumeV2Args{
// 			Size: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		instance1, err := compute.NewInstance(ctx, "instance1", &compute.InstanceArgs{
// 			SecurityGroups: pulumi.StringArray{
// 				pulumi.String("default"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewVolumeAttach(ctx, "va1", &compute.VolumeAttachArgs{
// 			InstanceId: instance1.ID(),
// 			VolumeId:   volume1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Using Multiattach-enabled volumes
//
// Multiattach Volumes are dependent upon your OpenStack cloud and not all
// clouds support multiattach.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := blockstorage.NewVolume(ctx, "volume1", &blockstorage.VolumeArgs{
// 			Multiattach: pulumi.Bool(true),
// 			Size:        pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		instance1, err := compute.NewInstance(ctx, "instance1", &compute.InstanceArgs{
// 			SecurityGroups: pulumi.StringArray{
// 				pulumi.String("default"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		instance2, err := compute.NewInstance(ctx, "instance2", &compute.InstanceArgs{
// 			SecurityGroups: pulumi.StringArray{
// 				pulumi.String("default"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewVolumeAttach(ctx, "va1", &compute.VolumeAttachArgs{
// 			InstanceId:  instance1.ID(),
// 			Multiattach: pulumi.Bool(true),
// 			VolumeId:    pulumi.Any(openstack_blockstorage_volume_v2.Volume_1.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewVolumeAttach(ctx, "va2", &compute.VolumeAttachArgs{
// 			InstanceId:  instance2.ID(),
// 			Multiattach: pulumi.Bool(true),
// 			VolumeId:    pulumi.Any(openstack_blockstorage_volume_v2.Volume_1.Id),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			pulumi.Resource("openstack_compute_volume_attach_v2.va_1"),
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// It is recommended to use `dependsOn` for the attach resources
// to enforce the volume attachments to happen one at a time.
//
// ## Import
//
// Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g.
//
// ```sh
//  $ pulumi import openstack:compute/volumeAttach:VolumeAttach va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
// ```
type VolumeAttach struct {
	pulumi.CustomResourceState

	// See Argument Reference above. _NOTE_: The correctness of this
	// information is dependent upon the hypervisor in use. In some cases, this
	// should not be used as an authoritative piece of information.
	Device pulumi.StringOutput `pulumi:"device"`
	// The ID of the Instance to attach the Volume to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Enable attachment of multiattach-capable volumes.
	Multiattach pulumi.BoolPtrOutput `pulumi:"multiattach"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a volume attachment. If omitted, the
	// `region` argument of the provider is used. Changing this creates a
	// new volume attachment.
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions VolumeAttachVendorOptionsPtrOutput `pulumi:"vendorOptions"`
	// The ID of the Volume to attach to an Instance.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewVolumeAttach registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttach(ctx *pulumi.Context,
	name string, args *VolumeAttachArgs, opts ...pulumi.ResourceOption) (*VolumeAttach, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	var resource VolumeAttach
	err := ctx.RegisterResource("openstack:compute/volumeAttach:VolumeAttach", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttach gets an existing VolumeAttach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttach(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeAttachState, opts ...pulumi.ResourceOption) (*VolumeAttach, error) {
	var resource VolumeAttach
	err := ctx.ReadResource("openstack:compute/volumeAttach:VolumeAttach", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeAttach resources.
type volumeAttachState struct {
	// See Argument Reference above. _NOTE_: The correctness of this
	// information is dependent upon the hypervisor in use. In some cases, this
	// should not be used as an authoritative piece of information.
	Device *string `pulumi:"device"`
	// The ID of the Instance to attach the Volume to.
	InstanceId *string `pulumi:"instanceId"`
	// Enable attachment of multiattach-capable volumes.
	Multiattach *bool `pulumi:"multiattach"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a volume attachment. If omitted, the
	// `region` argument of the provider is used. Changing this creates a
	// new volume attachment.
	Region *string `pulumi:"region"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions *VolumeAttachVendorOptions `pulumi:"vendorOptions"`
	// The ID of the Volume to attach to an Instance.
	VolumeId *string `pulumi:"volumeId"`
}

type VolumeAttachState struct {
	// See Argument Reference above. _NOTE_: The correctness of this
	// information is dependent upon the hypervisor in use. In some cases, this
	// should not be used as an authoritative piece of information.
	Device pulumi.StringPtrInput
	// The ID of the Instance to attach the Volume to.
	InstanceId pulumi.StringPtrInput
	// Enable attachment of multiattach-capable volumes.
	Multiattach pulumi.BoolPtrInput
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a volume attachment. If omitted, the
	// `region` argument of the provider is used. Changing this creates a
	// new volume attachment.
	Region pulumi.StringPtrInput
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions VolumeAttachVendorOptionsPtrInput
	// The ID of the Volume to attach to an Instance.
	VolumeId pulumi.StringPtrInput
}

func (VolumeAttachState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachState)(nil)).Elem()
}

type volumeAttachArgs struct {
	// See Argument Reference above. _NOTE_: The correctness of this
	// information is dependent upon the hypervisor in use. In some cases, this
	// should not be used as an authoritative piece of information.
	Device *string `pulumi:"device"`
	// The ID of the Instance to attach the Volume to.
	InstanceId string `pulumi:"instanceId"`
	// Enable attachment of multiattach-capable volumes.
	Multiattach *bool `pulumi:"multiattach"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a volume attachment. If omitted, the
	// `region` argument of the provider is used. Changing this creates a
	// new volume attachment.
	Region *string `pulumi:"region"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions *VolumeAttachVendorOptions `pulumi:"vendorOptions"`
	// The ID of the Volume to attach to an Instance.
	VolumeId string `pulumi:"volumeId"`
}

// The set of arguments for constructing a VolumeAttach resource.
type VolumeAttachArgs struct {
	// See Argument Reference above. _NOTE_: The correctness of this
	// information is dependent upon the hypervisor in use. In some cases, this
	// should not be used as an authoritative piece of information.
	Device pulumi.StringPtrInput
	// The ID of the Instance to attach the Volume to.
	InstanceId pulumi.StringInput
	// Enable attachment of multiattach-capable volumes.
	Multiattach pulumi.BoolPtrInput
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a volume attachment. If omitted, the
	// `region` argument of the provider is used. Changing this creates a
	// new volume attachment.
	Region pulumi.StringPtrInput
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions VolumeAttachVendorOptionsPtrInput
	// The ID of the Volume to attach to an Instance.
	VolumeId pulumi.StringInput
}

func (VolumeAttachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachArgs)(nil)).Elem()
}

type VolumeAttachInput interface {
	pulumi.Input

	ToVolumeAttachOutput() VolumeAttachOutput
	ToVolumeAttachOutputWithContext(ctx context.Context) VolumeAttachOutput
}

func (*VolumeAttach) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttach)(nil)).Elem()
}

func (i *VolumeAttach) ToVolumeAttachOutput() VolumeAttachOutput {
	return i.ToVolumeAttachOutputWithContext(context.Background())
}

func (i *VolumeAttach) ToVolumeAttachOutputWithContext(ctx context.Context) VolumeAttachOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachOutput)
}

// VolumeAttachArrayInput is an input type that accepts VolumeAttachArray and VolumeAttachArrayOutput values.
// You can construct a concrete instance of `VolumeAttachArrayInput` via:
//
//          VolumeAttachArray{ VolumeAttachArgs{...} }
type VolumeAttachArrayInput interface {
	pulumi.Input

	ToVolumeAttachArrayOutput() VolumeAttachArrayOutput
	ToVolumeAttachArrayOutputWithContext(context.Context) VolumeAttachArrayOutput
}

type VolumeAttachArray []VolumeAttachInput

func (VolumeAttachArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeAttach)(nil)).Elem()
}

func (i VolumeAttachArray) ToVolumeAttachArrayOutput() VolumeAttachArrayOutput {
	return i.ToVolumeAttachArrayOutputWithContext(context.Background())
}

func (i VolumeAttachArray) ToVolumeAttachArrayOutputWithContext(ctx context.Context) VolumeAttachArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachArrayOutput)
}

// VolumeAttachMapInput is an input type that accepts VolumeAttachMap and VolumeAttachMapOutput values.
// You can construct a concrete instance of `VolumeAttachMapInput` via:
//
//          VolumeAttachMap{ "key": VolumeAttachArgs{...} }
type VolumeAttachMapInput interface {
	pulumi.Input

	ToVolumeAttachMapOutput() VolumeAttachMapOutput
	ToVolumeAttachMapOutputWithContext(context.Context) VolumeAttachMapOutput
}

type VolumeAttachMap map[string]VolumeAttachInput

func (VolumeAttachMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeAttach)(nil)).Elem()
}

func (i VolumeAttachMap) ToVolumeAttachMapOutput() VolumeAttachMapOutput {
	return i.ToVolumeAttachMapOutputWithContext(context.Background())
}

func (i VolumeAttachMap) ToVolumeAttachMapOutputWithContext(ctx context.Context) VolumeAttachMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachMapOutput)
}

type VolumeAttachOutput struct{ *pulumi.OutputState }

func (VolumeAttachOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttach)(nil)).Elem()
}

func (o VolumeAttachOutput) ToVolumeAttachOutput() VolumeAttachOutput {
	return o
}

func (o VolumeAttachOutput) ToVolumeAttachOutputWithContext(ctx context.Context) VolumeAttachOutput {
	return o
}

type VolumeAttachArrayOutput struct{ *pulumi.OutputState }

func (VolumeAttachArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeAttach)(nil)).Elem()
}

func (o VolumeAttachArrayOutput) ToVolumeAttachArrayOutput() VolumeAttachArrayOutput {
	return o
}

func (o VolumeAttachArrayOutput) ToVolumeAttachArrayOutputWithContext(ctx context.Context) VolumeAttachArrayOutput {
	return o
}

func (o VolumeAttachArrayOutput) Index(i pulumi.IntInput) VolumeAttachOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VolumeAttach {
		return vs[0].([]*VolumeAttach)[vs[1].(int)]
	}).(VolumeAttachOutput)
}

type VolumeAttachMapOutput struct{ *pulumi.OutputState }

func (VolumeAttachMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeAttach)(nil)).Elem()
}

func (o VolumeAttachMapOutput) ToVolumeAttachMapOutput() VolumeAttachMapOutput {
	return o
}

func (o VolumeAttachMapOutput) ToVolumeAttachMapOutputWithContext(ctx context.Context) VolumeAttachMapOutput {
	return o
}

func (o VolumeAttachMapOutput) MapIndex(k pulumi.StringInput) VolumeAttachOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VolumeAttach {
		return vs[0].(map[string]*VolumeAttach)[vs[1].(string)]
	}).(VolumeAttachOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachInput)(nil)).Elem(), &VolumeAttach{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachArrayInput)(nil)).Elem(), VolumeAttachArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachMapInput)(nil)).Elem(), VolumeAttachMap{})
	pulumi.RegisterOutputType(VolumeAttachOutput{})
	pulumi.RegisterOutputType(VolumeAttachArrayOutput{})
	pulumi.RegisterOutputType(VolumeAttachMapOutput{})
}
