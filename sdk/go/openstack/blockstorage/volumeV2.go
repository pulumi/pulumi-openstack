// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package blockstorage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 volume resource within OpenStack.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/blockstorage_volume_v2.html.markdown.
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
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &VolumeV2Args{}
	}
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
