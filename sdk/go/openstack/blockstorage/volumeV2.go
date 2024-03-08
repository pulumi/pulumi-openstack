// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 volume resource within OpenStack.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blockstorage.NewVolumeV2(ctx, "volume1", &blockstorage.VolumeV2Args{
//				Description: pulumi.String("first test volume"),
//				Region:      pulumi.String("RegionOne"),
//				Size:        pulumi.Int(3),
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
// Volumes can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:blockstorage/volumeV2:VolumeV2 volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
// ```
type VolumeV2 struct {
	pulumi.CustomResourceState

	// If a volume is attached to an instance, this attribute will
	// display the Attachment ID, Instance ID, and the Device as the Instance
	// sees it.
	Attachments VolumeV2AttachmentArrayOutput `pulumi:"attachments"`
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId pulumi.StringPtrOutput `pulumi:"consistencyGroupId"`
	// A description of the volume. Changing this updates
	// the volume's description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId pulumi.StringPtrOutput `pulumi:"imageId"`
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region pulumi.StringOutput `pulumi:"region"`
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints VolumeV2SchedulerHintArrayOutput `pulumi:"schedulerHints"`
	// The size of the volume to create (in gigabytes). Changing
	// this creates a new volume.
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

// NewVolumeV2 registers a new resource with the given unique name, arguments, and options.
func NewVolumeV2(ctx *pulumi.Context,
	name string, args *VolumeV2Args, opts ...pulumi.ResourceOption) (*VolumeV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VolumeV2
	err := ctx.RegisterResource("openstack:blockstorage/volumeV2:VolumeV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeV2 gets an existing VolumeV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeV2State, opts ...pulumi.ResourceOption) (*VolumeV2, error) {
	var resource VolumeV2
	err := ctx.ReadResource("openstack:blockstorage/volumeV2:VolumeV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeV2 resources.
type volumeV2State struct {
	// If a volume is attached to an instance, this attribute will
	// display the Attachment ID, Instance ID, and the Device as the Instance
	// sees it.
	Attachments []VolumeV2Attachment `pulumi:"attachments"`
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId *string `pulumi:"consistencyGroupId"`
	// A description of the volume. Changing this updates
	// the volume's description.
	Description *string `pulumi:"description"`
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId *string `pulumi:"imageId"`
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name *string `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region *string `pulumi:"region"`
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints []VolumeV2SchedulerHint `pulumi:"schedulerHints"`
	// The size of the volume to create (in gigabytes). Changing
	// this creates a new volume.
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

type VolumeV2State struct {
	// If a volume is attached to an instance, this attribute will
	// display the Attachment ID, Instance ID, and the Device as the Instance
	// sees it.
	Attachments VolumeV2AttachmentArrayInput
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone pulumi.StringPtrInput
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId pulumi.StringPtrInput
	// A description of the volume. Changing this updates
	// the volume's description.
	Description pulumi.StringPtrInput
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId pulumi.StringPtrInput
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata pulumi.MapInput
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name pulumi.StringPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region pulumi.StringPtrInput
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints VolumeV2SchedulerHintArrayInput
	// The size of the volume to create (in gigabytes). Changing
	// this creates a new volume.
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

func (VolumeV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeV2State)(nil)).Elem()
}

type volumeV2Args struct {
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId *string `pulumi:"consistencyGroupId"`
	// A description of the volume. Changing this updates
	// the volume's description.
	Description *string `pulumi:"description"`
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId *string `pulumi:"imageId"`
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name *string `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region *string `pulumi:"region"`
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints []VolumeV2SchedulerHint `pulumi:"schedulerHints"`
	// The size of the volume to create (in gigabytes). Changing
	// this creates a new volume.
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

// The set of arguments for constructing a VolumeV2 resource.
type VolumeV2Args struct {
	// The availability zone for the volume.
	// Changing this creates a new volume.
	AvailabilityZone pulumi.StringPtrInput
	// The consistency group to place the volume
	// in.
	ConsistencyGroupId pulumi.StringPtrInput
	// A description of the volume. Changing this updates
	// the volume's description.
	Description pulumi.StringPtrInput
	// The image ID from which to create the volume.
	// Changing this creates a new volume.
	ImageId pulumi.StringPtrInput
	// Metadata key/value pairs to associate with the volume.
	// Changing this updates the existing volume metadata.
	Metadata pulumi.MapInput
	// A unique name for the volume. Changing this updates the
	// volume's name.
	Name pulumi.StringPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume.
	Region pulumi.StringPtrInput
	// Provide the Cinder scheduler with hints on where
	// to instantiate a volume in the OpenStack cloud. The available hints are described below.
	SchedulerHints VolumeV2SchedulerHintArrayInput
	// The size of the volume to create (in gigabytes). Changing
	// this creates a new volume.
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

func (VolumeV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeV2Args)(nil)).Elem()
}

type VolumeV2Input interface {
	pulumi.Input

	ToVolumeV2Output() VolumeV2Output
	ToVolumeV2OutputWithContext(ctx context.Context) VolumeV2Output
}

func (*VolumeV2) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeV2)(nil)).Elem()
}

func (i *VolumeV2) ToVolumeV2Output() VolumeV2Output {
	return i.ToVolumeV2OutputWithContext(context.Background())
}

func (i *VolumeV2) ToVolumeV2OutputWithContext(ctx context.Context) VolumeV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV2Output)
}

// VolumeV2ArrayInput is an input type that accepts VolumeV2Array and VolumeV2ArrayOutput values.
// You can construct a concrete instance of `VolumeV2ArrayInput` via:
//
//	VolumeV2Array{ VolumeV2Args{...} }
type VolumeV2ArrayInput interface {
	pulumi.Input

	ToVolumeV2ArrayOutput() VolumeV2ArrayOutput
	ToVolumeV2ArrayOutputWithContext(context.Context) VolumeV2ArrayOutput
}

type VolumeV2Array []VolumeV2Input

func (VolumeV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeV2)(nil)).Elem()
}

func (i VolumeV2Array) ToVolumeV2ArrayOutput() VolumeV2ArrayOutput {
	return i.ToVolumeV2ArrayOutputWithContext(context.Background())
}

func (i VolumeV2Array) ToVolumeV2ArrayOutputWithContext(ctx context.Context) VolumeV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV2ArrayOutput)
}

// VolumeV2MapInput is an input type that accepts VolumeV2Map and VolumeV2MapOutput values.
// You can construct a concrete instance of `VolumeV2MapInput` via:
//
//	VolumeV2Map{ "key": VolumeV2Args{...} }
type VolumeV2MapInput interface {
	pulumi.Input

	ToVolumeV2MapOutput() VolumeV2MapOutput
	ToVolumeV2MapOutputWithContext(context.Context) VolumeV2MapOutput
}

type VolumeV2Map map[string]VolumeV2Input

func (VolumeV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeV2)(nil)).Elem()
}

func (i VolumeV2Map) ToVolumeV2MapOutput() VolumeV2MapOutput {
	return i.ToVolumeV2MapOutputWithContext(context.Background())
}

func (i VolumeV2Map) ToVolumeV2MapOutputWithContext(ctx context.Context) VolumeV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV2MapOutput)
}

type VolumeV2Output struct{ *pulumi.OutputState }

func (VolumeV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeV2)(nil)).Elem()
}

func (o VolumeV2Output) ToVolumeV2Output() VolumeV2Output {
	return o
}

func (o VolumeV2Output) ToVolumeV2OutputWithContext(ctx context.Context) VolumeV2Output {
	return o
}

// If a volume is attached to an instance, this attribute will
// display the Attachment ID, Instance ID, and the Device as the Instance
// sees it.
func (o VolumeV2Output) Attachments() VolumeV2AttachmentArrayOutput {
	return o.ApplyT(func(v *VolumeV2) VolumeV2AttachmentArrayOutput { return v.Attachments }).(VolumeV2AttachmentArrayOutput)
}

// The availability zone for the volume.
// Changing this creates a new volume.
func (o VolumeV2Output) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The consistency group to place the volume
// in.
func (o VolumeV2Output) ConsistencyGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringPtrOutput { return v.ConsistencyGroupId }).(pulumi.StringPtrOutput)
}

// A description of the volume. Changing this updates
// the volume's description.
func (o VolumeV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The image ID from which to create the volume.
// Changing this creates a new volume.
func (o VolumeV2Output) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringPtrOutput { return v.ImageId }).(pulumi.StringPtrOutput)
}

// Metadata key/value pairs to associate with the volume.
// Changing this updates the existing volume metadata.
func (o VolumeV2Output) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// A unique name for the volume. Changing this updates the
// volume's name.
func (o VolumeV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to create the volume. If
// omitted, the `region` argument of the provider is used. Changing this
// creates a new volume.
func (o VolumeV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Provide the Cinder scheduler with hints on where
// to instantiate a volume in the OpenStack cloud. The available hints are described below.
func (o VolumeV2Output) SchedulerHints() VolumeV2SchedulerHintArrayOutput {
	return o.ApplyT(func(v *VolumeV2) VolumeV2SchedulerHintArrayOutput { return v.SchedulerHints }).(VolumeV2SchedulerHintArrayOutput)
}

// The size of the volume to create (in gigabytes). Changing
// this creates a new volume.
func (o VolumeV2Output) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The snapshot ID from which to create the volume.
// Changing this creates a new volume.
func (o VolumeV2Output) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// The volume ID to replicate with.
func (o VolumeV2Output) SourceReplica() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringPtrOutput { return v.SourceReplica }).(pulumi.StringPtrOutput)
}

// The volume ID from which to create the volume.
// Changing this creates a new volume.
func (o VolumeV2Output) SourceVolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringPtrOutput { return v.SourceVolId }).(pulumi.StringPtrOutput)
}

// The type of volume to create.
// Changing this creates a new volume.
func (o VolumeV2Output) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeV2) pulumi.StringOutput { return v.VolumeType }).(pulumi.StringOutput)
}

type VolumeV2ArrayOutput struct{ *pulumi.OutputState }

func (VolumeV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeV2)(nil)).Elem()
}

func (o VolumeV2ArrayOutput) ToVolumeV2ArrayOutput() VolumeV2ArrayOutput {
	return o
}

func (o VolumeV2ArrayOutput) ToVolumeV2ArrayOutputWithContext(ctx context.Context) VolumeV2ArrayOutput {
	return o
}

func (o VolumeV2ArrayOutput) Index(i pulumi.IntInput) VolumeV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VolumeV2 {
		return vs[0].([]*VolumeV2)[vs[1].(int)]
	}).(VolumeV2Output)
}

type VolumeV2MapOutput struct{ *pulumi.OutputState }

func (VolumeV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeV2)(nil)).Elem()
}

func (o VolumeV2MapOutput) ToVolumeV2MapOutput() VolumeV2MapOutput {
	return o
}

func (o VolumeV2MapOutput) ToVolumeV2MapOutputWithContext(ctx context.Context) VolumeV2MapOutput {
	return o
}

func (o VolumeV2MapOutput) MapIndex(k pulumi.StringInput) VolumeV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VolumeV2 {
		return vs[0].(map[string]*VolumeV2)[vs[1].(string)]
	}).(VolumeV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV2Input)(nil)).Elem(), &VolumeV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV2ArrayInput)(nil)).Elem(), VolumeV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV2MapInput)(nil)).Elem(), VolumeV2Map{})
	pulumi.RegisterOutputType(VolumeV2Output{})
	pulumi.RegisterOutputType(VolumeV2ArrayOutput{})
	pulumi.RegisterOutputType(VolumeV2MapOutput{})
}
