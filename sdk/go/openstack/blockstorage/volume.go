// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a V3 volume resource within OpenStack.
//
// ## Import
//
// Volumes can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:blockstorage/volume:Volume volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
//
// ```
type Volume struct {
	pulumi.CustomResourceState

	// If a volume is attached to an instance, this attribute will
	// display the Attachment ID, Instance ID, and the Device as the Instance
	// sees it.
	Attachments VolumeAttachmentArrayOutput `pulumi:"attachments"`
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId pulumi.StringPtrOutput `pulumi:"consistencyGroupId"`
	// A description of the volume. Changing this updates
	// the volume's description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When this option is set it allows extending
	// attached volumes. Note: updating size of an attached volume requires Cinder
	// support for version 3.42 and a compatible storage driver.
	EnableOnlineResize pulumi.BoolPtrOutput `pulumi:"enableOnlineResize"`
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Allow the volume to be attached to more than one Compute instance.
	Multiattach pulumi.BoolPtrOutput `pulumi:"multiattach"`
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region pulumi.StringOutput `pulumi:"region"`
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints VolumeSchedulerHintArrayOutput `pulumi:"schedulerHints"`
	// The size of the volume to create (in gigabytes).
	Size pulumi.IntOutput `pulumi:"size"`
	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// The volume ID to replicate with.
	SourceReplica pulumi.StringPtrOutput `pulumi:"sourceReplica"`
	// The volume ID from which to create the volume.
	// Changing this creates a new volume.
	SourceVolId pulumi.StringPtrOutput `pulumi:"sourceVolId"`
	// The type of volume to create.
	// Changing this creates a new volume.
	VolumeType pulumi.StringOutput `pulumi:"volumeType"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("openstack:blockstorage/volume:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("openstack:blockstorage/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// If a volume is attached to an instance, this attribute will
	// display the Attachment ID, Instance ID, and the Device as the Instance
	// sees it.
	Attachments []VolumeAttachment `pulumi:"attachments"`
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId *string `pulumi:"consistencyGroupId"`
	// A description of the volume. Changing this updates
	// the volume's description.
	Description *string `pulumi:"description"`
	// When this option is set it allows extending
	// attached volumes. Note: updating size of an attached volume requires Cinder
	// support for version 3.42 and a compatible storage driver.
	EnableOnlineResize *bool `pulumi:"enableOnlineResize"`
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId *string `pulumi:"imageId"`
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Allow the volume to be attached to more than one Compute instance.
	Multiattach *bool `pulumi:"multiattach"`
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name *string `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region *string `pulumi:"region"`
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints []VolumeSchedulerHint `pulumi:"schedulerHints"`
	// The size of the volume to create (in gigabytes).
	Size *int `pulumi:"size"`
	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	SnapshotId *string `pulumi:"snapshotId"`
	// The volume ID to replicate with.
	SourceReplica *string `pulumi:"sourceReplica"`
	// The volume ID from which to create the volume.
	// Changing this creates a new volume.
	SourceVolId *string `pulumi:"sourceVolId"`
	// The type of volume to create.
	// Changing this creates a new volume.
	VolumeType *string `pulumi:"volumeType"`
}

type VolumeState struct {
	// If a volume is attached to an instance, this attribute will
	// display the Attachment ID, Instance ID, and the Device as the Instance
	// sees it.
	Attachments VolumeAttachmentArrayInput
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone pulumi.StringPtrInput
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId pulumi.StringPtrInput
	// A description of the volume. Changing this updates
	// the volume's description.
	Description pulumi.StringPtrInput
	// When this option is set it allows extending
	// attached volumes. Note: updating size of an attached volume requires Cinder
	// support for version 3.42 and a compatible storage driver.
	EnableOnlineResize pulumi.BoolPtrInput
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId pulumi.StringPtrInput
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata pulumi.MapInput
	// Allow the volume to be attached to more than one Compute instance.
	Multiattach pulumi.BoolPtrInput
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name pulumi.StringPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region pulumi.StringPtrInput
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints VolumeSchedulerHintArrayInput
	// The size of the volume to create (in gigabytes).
	Size pulumi.IntPtrInput
	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	SnapshotId pulumi.StringPtrInput
	// The volume ID to replicate with.
	SourceReplica pulumi.StringPtrInput
	// The volume ID from which to create the volume.
	// Changing this creates a new volume.
	SourceVolId pulumi.StringPtrInput
	// The type of volume to create.
	// Changing this creates a new volume.
	VolumeType pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId *string `pulumi:"consistencyGroupId"`
	// A description of the volume. Changing this updates
	// the volume's description.
	Description *string `pulumi:"description"`
	// When this option is set it allows extending
	// attached volumes. Note: updating size of an attached volume requires Cinder
	// support for version 3.42 and a compatible storage driver.
	EnableOnlineResize *bool `pulumi:"enableOnlineResize"`
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId *string `pulumi:"imageId"`
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Allow the volume to be attached to more than one Compute instance.
	Multiattach *bool `pulumi:"multiattach"`
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name *string `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region *string `pulumi:"region"`
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints []VolumeSchedulerHint `pulumi:"schedulerHints"`
	// The size of the volume to create (in gigabytes).
	Size int `pulumi:"size"`
	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	SnapshotId *string `pulumi:"snapshotId"`
	// The volume ID to replicate with.
	SourceReplica *string `pulumi:"sourceReplica"`
	// The volume ID from which to create the volume.
	// Changing this creates a new volume.
	SourceVolId *string `pulumi:"sourceVolId"`
	// The type of volume to create.
	// Changing this creates a new volume.
	VolumeType *string `pulumi:"volumeType"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone pulumi.StringPtrInput
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId pulumi.StringPtrInput
	// A description of the volume. Changing this updates
	// the volume's description.
	Description pulumi.StringPtrInput
	// When this option is set it allows extending
	// attached volumes. Note: updating size of an attached volume requires Cinder
	// support for version 3.42 and a compatible storage driver.
	EnableOnlineResize pulumi.BoolPtrInput
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId pulumi.StringPtrInput
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata pulumi.MapInput
	// Allow the volume to be attached to more than one Compute instance.
	Multiattach pulumi.BoolPtrInput
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name pulumi.StringPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region pulumi.StringPtrInput
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints VolumeSchedulerHintArrayInput
	// The size of the volume to create (in gigabytes).
	Size pulumi.IntInput
	// The snapshot ID from which to create the volume.
	// Changing this creates a new volume.
	SnapshotId pulumi.StringPtrInput
	// The volume ID to replicate with.
	SourceReplica pulumi.StringPtrInput
	// The volume ID from which to create the volume.
	// Changing this creates a new volume.
	SourceVolId pulumi.StringPtrInput
	// The type of volume to create.
	// Changing this creates a new volume.
	VolumeType pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

func (i *Volume) ToOutput(ctx context.Context) pulumix.Output[*Volume] {
	return pulumix.Output[*Volume]{
		OutputState: i.ToVolumeOutputWithContext(ctx).OutputState,
	}
}

// VolumeArrayInput is an input type that accepts VolumeArray and VolumeArrayOutput values.
// You can construct a concrete instance of `VolumeArrayInput` via:
//
//	VolumeArray{ VolumeArgs{...} }
type VolumeArrayInput interface {
	pulumi.Input

	ToVolumeArrayOutput() VolumeArrayOutput
	ToVolumeArrayOutputWithContext(context.Context) VolumeArrayOutput
}

type VolumeArray []VolumeInput

func (VolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (i VolumeArray) ToVolumeArrayOutput() VolumeArrayOutput {
	return i.ToVolumeArrayOutputWithContext(context.Background())
}

func (i VolumeArray) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArrayOutput)
}

func (i VolumeArray) ToOutput(ctx context.Context) pulumix.Output[[]*Volume] {
	return pulumix.Output[[]*Volume]{
		OutputState: i.ToVolumeArrayOutputWithContext(ctx).OutputState,
	}
}

// VolumeMapInput is an input type that accepts VolumeMap and VolumeMapOutput values.
// You can construct a concrete instance of `VolumeMapInput` via:
//
//	VolumeMap{ "key": VolumeArgs{...} }
type VolumeMapInput interface {
	pulumi.Input

	ToVolumeMapOutput() VolumeMapOutput
	ToVolumeMapOutputWithContext(context.Context) VolumeMapOutput
}

type VolumeMap map[string]VolumeInput

func (VolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (i VolumeMap) ToVolumeMapOutput() VolumeMapOutput {
	return i.ToVolumeMapOutputWithContext(context.Background())
}

func (i VolumeMap) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMapOutput)
}

func (i VolumeMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Volume] {
	return pulumix.Output[map[string]*Volume]{
		OutputState: i.ToVolumeMapOutputWithContext(ctx).OutputState,
	}
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

func (o VolumeOutput) ToOutput(ctx context.Context) pulumix.Output[*Volume] {
	return pulumix.Output[*Volume]{
		OutputState: o.OutputState,
	}
}

// If a volume is attached to an instance, this attribute will
// display the Attachment ID, Instance ID, and the Device as the Instance
// sees it.
func (o VolumeOutput) Attachments() VolumeAttachmentArrayOutput {
	return o.ApplyT(func(v *Volume) VolumeAttachmentArrayOutput { return v.Attachments }).(VolumeAttachmentArrayOutput)
}

// The availability zone for the volume.
// Changing this creates a new volume.
func (o VolumeOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The consistency group to place the volume
// in.
func (o VolumeOutput) ConsistencyGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.ConsistencyGroupId }).(pulumi.StringPtrOutput)
}

// A description of the volume. Changing this updates
// the volume's description.
func (o VolumeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When this option is set it allows extending
// attached volumes. Note: updating size of an attached volume requires Cinder
// support for version 3.42 and a compatible storage driver.
func (o VolumeOutput) EnableOnlineResize() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.EnableOnlineResize }).(pulumi.BoolPtrOutput)
}

// The image ID from which to create the volume.
// Changing this creates a new volume.
func (o VolumeOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.ImageId }).(pulumi.StringPtrOutput)
}

// Metadata key/value pairs to associate with the volume.
// Changing this updates the existing volume metadata.
func (o VolumeOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *Volume) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// Allow the volume to be attached to more than one Compute instance.
func (o VolumeOutput) Multiattach() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.Multiattach }).(pulumi.BoolPtrOutput)
}

// A unique name for the volume. Changing this updates the
// volume's name.
func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to create the volume. If
// omitted, the `region` argument of the provider is used. Changing this
// creates a new volume.
func (o VolumeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Provide the Cinder scheduler with hints on where
// to instantiate a volume in the OpenStack cloud. The available hints are described below.
func (o VolumeOutput) SchedulerHints() VolumeSchedulerHintArrayOutput {
	return o.ApplyT(func(v *Volume) VolumeSchedulerHintArrayOutput { return v.SchedulerHints }).(VolumeSchedulerHintArrayOutput)
}

// The size of the volume to create (in gigabytes).
func (o VolumeOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The snapshot ID from which to create the volume.
// Changing this creates a new volume.
func (o VolumeOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// The volume ID to replicate with.
func (o VolumeOutput) SourceReplica() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SourceReplica }).(pulumi.StringPtrOutput)
}

// The volume ID from which to create the volume.
// Changing this creates a new volume.
func (o VolumeOutput) SourceVolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SourceVolId }).(pulumi.StringPtrOutput)
}

// The type of volume to create.
// Changing this creates a new volume.
func (o VolumeOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeType }).(pulumi.StringOutput)
}

type VolumeArrayOutput struct{ *pulumi.OutputState }

func (VolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (o VolumeArrayOutput) ToVolumeArrayOutput() VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Volume] {
	return pulumix.Output[[]*Volume]{
		OutputState: o.OutputState,
	}
}

func (o VolumeArrayOutput) Index(i pulumi.IntInput) VolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].([]*Volume)[vs[1].(int)]
	}).(VolumeOutput)
}

type VolumeMapOutput struct{ *pulumi.OutputState }

func (VolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (o VolumeMapOutput) ToVolumeMapOutput() VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Volume] {
	return pulumix.Output[map[string]*Volume]{
		OutputState: o.OutputState,
	}
}

func (o VolumeMapOutput) MapIndex(k pulumi.StringInput) VolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].(map[string]*Volume)[vs[1].(string)]
	}).(VolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeInput)(nil)).Elem(), &Volume{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeArrayInput)(nil)).Elem(), VolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMapInput)(nil)).Elem(), VolumeMap{})
	pulumi.RegisterOutputType(VolumeOutput{})
	pulumi.RegisterOutputType(VolumeArrayOutput{})
	pulumi.RegisterOutputType(VolumeMapOutput{})
}
