// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource is experimental and may be removed in the future! Feedback
// is requested if you find this resource useful or if you find any problems
// with it.
// 
// Creates a general purpose attachment connection to a Block
// Storage volume using the OpenStack Block Storage (Cinder) v3 API.
// Depending on your Block Storage service configuration, this
// resource can assist in attaching a volume to a non-OpenStack resource
// such as a bare-metal server or a remote virtual machine in a
// different cloud provider.
// 
// This does not actually attach a volume to an instance. Please use
// the `compute.VolumeAttach` resource for that.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/blockstorage_volume_attach_v3.html.markdown.
type VolumeAttach struct {
	s *pulumi.ResourceState
}

// NewVolumeAttach registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttach(ctx *pulumi.Context,
	name string, args *VolumeAttachArgs, opts ...pulumi.ResourceOpt) (*VolumeAttach, error) {
	if args == nil || args.HostName == nil {
		return nil, errors.New("missing required argument 'HostName'")
	}
	if args == nil || args.VolumeId == nil {
		return nil, errors.New("missing required argument 'VolumeId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attachMode"] = nil
		inputs["device"] = nil
		inputs["hostName"] = nil
		inputs["initiator"] = nil
		inputs["ipAddress"] = nil
		inputs["multipath"] = nil
		inputs["osType"] = nil
		inputs["platform"] = nil
		inputs["region"] = nil
		inputs["volumeId"] = nil
		inputs["wwnn"] = nil
		inputs["wwpns"] = nil
	} else {
		inputs["attachMode"] = args.AttachMode
		inputs["device"] = args.Device
		inputs["hostName"] = args.HostName
		inputs["initiator"] = args.Initiator
		inputs["ipAddress"] = args.IpAddress
		inputs["multipath"] = args.Multipath
		inputs["osType"] = args.OsType
		inputs["platform"] = args.Platform
		inputs["region"] = args.Region
		inputs["volumeId"] = args.VolumeId
		inputs["wwnn"] = args.Wwnn
		inputs["wwpns"] = args.Wwpns
	}
	inputs["data"] = nil
	inputs["driverVolumeType"] = nil
	inputs["mountPointBase"] = nil
	s, err := ctx.RegisterResource("openstack:blockstorage/volumeAttach:VolumeAttach", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VolumeAttach{s: s}, nil
}

// GetVolumeAttach gets an existing VolumeAttach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttach(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VolumeAttachState, opts ...pulumi.ResourceOpt) (*VolumeAttach, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attachMode"] = state.AttachMode
		inputs["data"] = state.Data
		inputs["device"] = state.Device
		inputs["driverVolumeType"] = state.DriverVolumeType
		inputs["hostName"] = state.HostName
		inputs["initiator"] = state.Initiator
		inputs["ipAddress"] = state.IpAddress
		inputs["mountPointBase"] = state.MountPointBase
		inputs["multipath"] = state.Multipath
		inputs["osType"] = state.OsType
		inputs["platform"] = state.Platform
		inputs["region"] = state.Region
		inputs["volumeId"] = state.VolumeId
		inputs["wwnn"] = state.Wwnn
		inputs["wwpns"] = state.Wwpns
	}
	s, err := ctx.ReadResource("openstack:blockstorage/volumeAttach:VolumeAttach", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VolumeAttach{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VolumeAttach) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VolumeAttach) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specify whether to attach the volume as Read-Only
// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
// If left unspecified, the Block Storage API will apply a default of `rw`.
func (r *VolumeAttach) AttachMode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["attachMode"])
}

// This is a map of key/value pairs that contain the connection
// information. You will want to pass this information to a provisioner
// script to finalize the connection. See below for more information.
func (r *VolumeAttach) Data() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["data"])
}

// The device to tell the Block Storage service this
// volume will be attached as. This is purely for informational purposes.
// You can specify `auto` or a device such as `/dev/vdc`.
func (r *VolumeAttach) Device() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["device"])
}

// The storage driver that the volume is based on.
func (r *VolumeAttach) DriverVolumeType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["driverVolumeType"])
}

// The host to attach the volume to.
func (r *VolumeAttach) HostName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hostName"])
}

// The iSCSI initiator string to make the connection.
func (r *VolumeAttach) Initiator() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["initiator"])
}

// The IP address of the `hostName` above.
func (r *VolumeAttach) IpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ipAddress"])
}

// A mount point base name for shared storage.
func (r *VolumeAttach) MountPointBase() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["mountPointBase"])
}

// Whether to connect to this volume via multipath.
func (r *VolumeAttach) Multipath() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["multipath"])
}

// The iSCSI initiator OS type.
func (r *VolumeAttach) OsType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["osType"])
}

// The iSCSI initiator platform.
func (r *VolumeAttach) Platform() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["platform"])
}

// The region in which to obtain the V3 Block Storage
// client. A Block Storage client is needed to create a volume attachment.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new volume attachment.
func (r *VolumeAttach) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The ID of the Volume to attach to an Instance.
func (r *VolumeAttach) VolumeId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["volumeId"])
}

// A wwnn name. Used for Fibre Channel connections.
func (r *VolumeAttach) Wwnn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["wwnn"])
}

// An array of wwpn strings. Used for Fibre Channel
// connections.
func (r *VolumeAttach) Wwpns() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["wwpns"])
}

// Input properties used for looking up and filtering VolumeAttach resources.
type VolumeAttachState struct {
	// Specify whether to attach the volume as Read-Only
	// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
	// If left unspecified, the Block Storage API will apply a default of `rw`.
	AttachMode interface{}
	// This is a map of key/value pairs that contain the connection
	// information. You will want to pass this information to a provisioner
	// script to finalize the connection. See below for more information.
	Data interface{}
	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify `auto` or a device such as `/dev/vdc`.
	Device interface{}
	// The storage driver that the volume is based on.
	DriverVolumeType interface{}
	// The host to attach the volume to.
	HostName interface{}
	// The iSCSI initiator string to make the connection.
	Initiator interface{}
	// The IP address of the `hostName` above.
	IpAddress interface{}
	// A mount point base name for shared storage.
	MountPointBase interface{}
	// Whether to connect to this volume via multipath.
	Multipath interface{}
	// The iSCSI initiator OS type.
	OsType interface{}
	// The iSCSI initiator platform.
	Platform interface{}
	// The region in which to obtain the V3 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region interface{}
	// The ID of the Volume to attach to an Instance.
	VolumeId interface{}
	// A wwnn name. Used for Fibre Channel connections.
	Wwnn interface{}
	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpns interface{}
}

// The set of arguments for constructing a VolumeAttach resource.
type VolumeAttachArgs struct {
	// Specify whether to attach the volume as Read-Only
	// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
	// If left unspecified, the Block Storage API will apply a default of `rw`.
	AttachMode interface{}
	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify `auto` or a device such as `/dev/vdc`.
	Device interface{}
	// The host to attach the volume to.
	HostName interface{}
	// The iSCSI initiator string to make the connection.
	Initiator interface{}
	// The IP address of the `hostName` above.
	IpAddress interface{}
	// Whether to connect to this volume via multipath.
	Multipath interface{}
	// The iSCSI initiator OS type.
	OsType interface{}
	// The iSCSI initiator platform.
	Platform interface{}
	// The region in which to obtain the V3 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region interface{}
	// The ID of the Volume to attach to an Instance.
	VolumeId interface{}
	// A wwnn name. Used for Fibre Channel connections.
	Wwnn interface{}
	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpns interface{}
}
