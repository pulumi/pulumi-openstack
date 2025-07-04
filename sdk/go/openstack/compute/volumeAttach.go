// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches a Block Storage Volume to an Instance using the OpenStack
// Compute (Nova) v2 API.
//
// ## Example Usage
//
// ### Basic attachment of a single volume to a single instance
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/blockstorage"
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			volume1, err := blockstorage.NewVolume(ctx, "volume_1", &blockstorage.VolumeArgs{
//				Name: pulumi.String("volume_1"),
//				Size: pulumi.Int(1),
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
//			_, err = compute.NewVolumeAttach(ctx, "va_1", &compute.VolumeAttachArgs{
//				InstanceId: instance1.ID(),
//				VolumeId:   volume1.ID(),
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
// ### Using Multiattach-enabled volumes
//
// Multiattach Volumes are dependent upon your OpenStack cloud and not all
// clouds support multiattach. Multiattach volumes require a volumeType that has [multiattach enabled](https://docs.openstack.org/cinder/latest/admin/volume-multiattach.html#multiattach-volume-type).
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/blockstorage"
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			volume1, err := blockstorage.NewVolume(ctx, "volume_1", &blockstorage.VolumeArgs{
//				Name:       pulumi.String("volume_1"),
//				Size:       pulumi.Int(1),
//				VolumeType: pulumi.String("multiattach"),
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
//			instance2, err := compute.NewInstance(ctx, "instance_2", &compute.InstanceArgs{
//				Name: pulumi.String("instance_2"),
//				SecurityGroups: pulumi.StringArray{
//					pulumi.String("default"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			va1, err := compute.NewVolumeAttach(ctx, "va_1", &compute.VolumeAttachArgs{
//				InstanceId:  instance1.ID(),
//				VolumeId:    volume1.ID(),
//				Multiattach: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewVolumeAttach(ctx, "va_2", &compute.VolumeAttachArgs{
//				InstanceId:  instance2.ID(),
//				VolumeId:    volume1.ID(),
//				Multiattach: pulumi.Bool(true),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				va1,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// It is recommended to use `dependsOn` for the attach resources
// to enforce the volume attachments to happen one at a time.
//
// ## Import
//
// Volume Attachments can be imported using the Instance ID and Volume ID
// separated by a slash, e.g.
//
// ```sh
// $ pulumi import openstack:compute/volumeAttach:VolumeAttach va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
// ```
type VolumeAttach struct {
	pulumi.CustomResourceState

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
	// Add a device role tag that is applied to the volume when
	// attaching it to the VM. Changing this creates a new volume attachment with
	// the new tag. Requires microversion >= 2.49.
	Tag pulumi.StringPtrOutput `pulumi:"tag"`
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Add a device role tag that is applied to the volume when
	// attaching it to the VM. Changing this creates a new volume attachment with
	// the new tag. Requires microversion >= 2.49.
	Tag *string `pulumi:"tag"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions *VolumeAttachVendorOptions `pulumi:"vendorOptions"`
	// The ID of the Volume to attach to an Instance.
	VolumeId *string `pulumi:"volumeId"`
}

type VolumeAttachState struct {
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
	// Add a device role tag that is applied to the volume when
	// attaching it to the VM. Changing this creates a new volume attachment with
	// the new tag. Requires microversion >= 2.49.
	Tag pulumi.StringPtrInput
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
	// Add a device role tag that is applied to the volume when
	// attaching it to the VM. Changing this creates a new volume attachment with
	// the new tag. Requires microversion >= 2.49.
	Tag *string `pulumi:"tag"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions *VolumeAttachVendorOptions `pulumi:"vendorOptions"`
	// The ID of the Volume to attach to an Instance.
	VolumeId string `pulumi:"volumeId"`
}

// The set of arguments for constructing a VolumeAttach resource.
type VolumeAttachArgs struct {
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
	// Add a device role tag that is applied to the volume when
	// attaching it to the VM. Changing this creates a new volume attachment with
	// the new tag. Requires microversion >= 2.49.
	Tag pulumi.StringPtrInput
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
//	VolumeAttachArray{ VolumeAttachArgs{...} }
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
//	VolumeAttachMap{ "key": VolumeAttachArgs{...} }
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

func (o VolumeAttachOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttach) pulumi.StringOutput { return v.Device }).(pulumi.StringOutput)
}

// The ID of the Instance to attach the Volume to.
func (o VolumeAttachOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttach) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Enable attachment of multiattach-capable volumes.
func (o VolumeAttachOutput) Multiattach() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VolumeAttach) pulumi.BoolPtrOutput { return v.Multiattach }).(pulumi.BoolPtrOutput)
}

// The region in which to obtain the V2 Compute client.
// A Compute client is needed to create a volume attachment. If omitted, the
// `region` argument of the provider is used. Changing this creates a
// new volume attachment.
func (o VolumeAttachOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttach) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Add a device role tag that is applied to the volume when
// attaching it to the VM. Changing this creates a new volume attachment with
// the new tag. Requires microversion >= 2.49.
func (o VolumeAttachOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttach) pulumi.StringPtrOutput { return v.Tag }).(pulumi.StringPtrOutput)
}

// Map of additional vendor-specific options.
// Supported options are described below.
func (o VolumeAttachOutput) VendorOptions() VolumeAttachVendorOptionsPtrOutput {
	return o.ApplyT(func(v *VolumeAttach) VolumeAttachVendorOptionsPtrOutput { return v.VendorOptions }).(VolumeAttachVendorOptionsPtrOutput)
}

// The ID of the Volume to attach to an Instance.
func (o VolumeAttachOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttach) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
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
