// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package blockstorage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V3 volume resource within OpenStack.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/blockstorage_volume_v3.html.markdown.
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
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &VolumeArgs{}
	}
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
