// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches a Block Storage Volume to an Instance using the OpenStack
// Compute (Nova) v2 API.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_volume_attach_v2.html.markdown.
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
	// The ID of the Volume to attach to an Instance.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewVolumeAttach registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttach(ctx *pulumi.Context,
	name string, args *VolumeAttachArgs, opts ...pulumi.ResourceOption) (*VolumeAttach, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil || args.VolumeId == nil {
		return nil, errors.New("missing required argument 'VolumeId'")
	}
	if args == nil {
		args = &VolumeAttachArgs{}
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
	// The ID of the Volume to attach to an Instance.
	VolumeId pulumi.StringInput
}

func (VolumeAttachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachArgs)(nil)).Elem()
}

