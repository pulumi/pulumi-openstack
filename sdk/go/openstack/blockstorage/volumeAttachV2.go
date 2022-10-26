// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
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
//			volume1, err := blockstorage.NewVolumeV2(ctx, "volume1", &blockstorage.VolumeV2Args{
//				Size: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = blockstorage.NewVolumeAttachV2(ctx, "va1", &blockstorage.VolumeAttachV2Args{
//				Device:    pulumi.String("auto"),
//				HostName:  pulumi.String("devstack"),
//				Initiator: pulumi.String("iqn.1993-08.org.debian:01:e9861fb1859"),
//				IpAddress: pulumi.String("192.168.255.10"),
//				OsType:    pulumi.String("linux2"),
//				Platform:  pulumi.String("x86_64"),
//				VolumeId:  volume1.ID(),
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
// ## Import
//
// It is not possible to import this resource.
type VolumeAttachV2 struct {
	pulumi.CustomResourceState

	// Specify whether to attach the volume as Read-Only
	// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
	// If left unspecified, the Block Storage API will apply a default of `rw`.
	AttachMode pulumi.StringPtrOutput `pulumi:"attachMode"`
	// This is a map of key/value pairs that contain the connection
	// information. You will want to pass this information to a provisioner
	// script to finalize the connection. See below for more information.
	Data pulumi.MapOutput `pulumi:"data"`
	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify `auto` or a device such as `/dev/vdc`.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// The storage driver that the volume is based on.
	DriverVolumeType pulumi.StringOutput `pulumi:"driverVolumeType"`
	// The host to attach the volume to.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The iSCSI initiator string to make the connection.
	Initiator pulumi.StringPtrOutput `pulumi:"initiator"`
	// Deprecated: instance_id is no longer used in this resource
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// The IP address of the `hostName` above.
	IpAddress pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// A mount point base name for shared storage.
	MountPointBase pulumi.StringOutput `pulumi:"mountPointBase"`
	// Whether to connect to this volume via multipath.
	Multipath pulumi.BoolPtrOutput `pulumi:"multipath"`
	// The iSCSI initiator OS type.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The iSCSI initiator platform.
	Platform pulumi.StringPtrOutput `pulumi:"platform"`
	// The region in which to obtain the V2 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of the Volume to attach to an Instance.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
	// A wwnn name. Used for Fibre Channel connections.
	Wwnn pulumi.StringPtrOutput `pulumi:"wwnn"`
	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpns pulumi.StringArrayOutput `pulumi:"wwpns"`
}

// NewVolumeAttachV2 registers a new resource with the given unique name, arguments, and options.
func NewVolumeAttachV2(ctx *pulumi.Context,
	name string, args *VolumeAttachV2Args, opts ...pulumi.ResourceOption) (*VolumeAttachV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	var resource VolumeAttachV2
	err := ctx.RegisterResource("openstack:blockstorage/volumeAttachV2:VolumeAttachV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeAttachV2 gets an existing VolumeAttachV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeAttachV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeAttachV2State, opts ...pulumi.ResourceOption) (*VolumeAttachV2, error) {
	var resource VolumeAttachV2
	err := ctx.ReadResource("openstack:blockstorage/volumeAttachV2:VolumeAttachV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeAttachV2 resources.
type volumeAttachV2State struct {
	// Specify whether to attach the volume as Read-Only
	// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
	// If left unspecified, the Block Storage API will apply a default of `rw`.
	AttachMode *string `pulumi:"attachMode"`
	// This is a map of key/value pairs that contain the connection
	// information. You will want to pass this information to a provisioner
	// script to finalize the connection. See below for more information.
	Data map[string]interface{} `pulumi:"data"`
	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify `auto` or a device such as `/dev/vdc`.
	Device *string `pulumi:"device"`
	// The storage driver that the volume is based on.
	DriverVolumeType *string `pulumi:"driverVolumeType"`
	// The host to attach the volume to.
	HostName *string `pulumi:"hostName"`
	// The iSCSI initiator string to make the connection.
	Initiator *string `pulumi:"initiator"`
	// Deprecated: instance_id is no longer used in this resource
	InstanceId *string `pulumi:"instanceId"`
	// The IP address of the `hostName` above.
	IpAddress *string `pulumi:"ipAddress"`
	// A mount point base name for shared storage.
	MountPointBase *string `pulumi:"mountPointBase"`
	// Whether to connect to this volume via multipath.
	Multipath *bool `pulumi:"multipath"`
	// The iSCSI initiator OS type.
	OsType *string `pulumi:"osType"`
	// The iSCSI initiator platform.
	Platform *string `pulumi:"platform"`
	// The region in which to obtain the V2 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region *string `pulumi:"region"`
	// The ID of the Volume to attach to an Instance.
	VolumeId *string `pulumi:"volumeId"`
	// A wwnn name. Used for Fibre Channel connections.
	Wwnn *string `pulumi:"wwnn"`
	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpns []string `pulumi:"wwpns"`
}

type VolumeAttachV2State struct {
	// Specify whether to attach the volume as Read-Only
	// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
	// If left unspecified, the Block Storage API will apply a default of `rw`.
	AttachMode pulumi.StringPtrInput
	// This is a map of key/value pairs that contain the connection
	// information. You will want to pass this information to a provisioner
	// script to finalize the connection. See below for more information.
	Data pulumi.MapInput
	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify `auto` or a device such as `/dev/vdc`.
	Device pulumi.StringPtrInput
	// The storage driver that the volume is based on.
	DriverVolumeType pulumi.StringPtrInput
	// The host to attach the volume to.
	HostName pulumi.StringPtrInput
	// The iSCSI initiator string to make the connection.
	Initiator pulumi.StringPtrInput
	// Deprecated: instance_id is no longer used in this resource
	InstanceId pulumi.StringPtrInput
	// The IP address of the `hostName` above.
	IpAddress pulumi.StringPtrInput
	// A mount point base name for shared storage.
	MountPointBase pulumi.StringPtrInput
	// Whether to connect to this volume via multipath.
	Multipath pulumi.BoolPtrInput
	// The iSCSI initiator OS type.
	OsType pulumi.StringPtrInput
	// The iSCSI initiator platform.
	Platform pulumi.StringPtrInput
	// The region in which to obtain the V2 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region pulumi.StringPtrInput
	// The ID of the Volume to attach to an Instance.
	VolumeId pulumi.StringPtrInput
	// A wwnn name. Used for Fibre Channel connections.
	Wwnn pulumi.StringPtrInput
	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpns pulumi.StringArrayInput
}

func (VolumeAttachV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachV2State)(nil)).Elem()
}

type volumeAttachV2Args struct {
	// Specify whether to attach the volume as Read-Only
	// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
	// If left unspecified, the Block Storage API will apply a default of `rw`.
	AttachMode *string `pulumi:"attachMode"`
	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify `auto` or a device such as `/dev/vdc`.
	Device *string `pulumi:"device"`
	// The host to attach the volume to.
	HostName string `pulumi:"hostName"`
	// The iSCSI initiator string to make the connection.
	Initiator *string `pulumi:"initiator"`
	// Deprecated: instance_id is no longer used in this resource
	InstanceId *string `pulumi:"instanceId"`
	// The IP address of the `hostName` above.
	IpAddress *string `pulumi:"ipAddress"`
	// Whether to connect to this volume via multipath.
	Multipath *bool `pulumi:"multipath"`
	// The iSCSI initiator OS type.
	OsType *string `pulumi:"osType"`
	// The iSCSI initiator platform.
	Platform *string `pulumi:"platform"`
	// The region in which to obtain the V2 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region *string `pulumi:"region"`
	// The ID of the Volume to attach to an Instance.
	VolumeId string `pulumi:"volumeId"`
	// A wwnn name. Used for Fibre Channel connections.
	Wwnn *string `pulumi:"wwnn"`
	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpns []string `pulumi:"wwpns"`
}

// The set of arguments for constructing a VolumeAttachV2 resource.
type VolumeAttachV2Args struct {
	// Specify whether to attach the volume as Read-Only
	// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
	// If left unspecified, the Block Storage API will apply a default of `rw`.
	AttachMode pulumi.StringPtrInput
	// The device to tell the Block Storage service this
	// volume will be attached as. This is purely for informational purposes.
	// You can specify `auto` or a device such as `/dev/vdc`.
	Device pulumi.StringPtrInput
	// The host to attach the volume to.
	HostName pulumi.StringInput
	// The iSCSI initiator string to make the connection.
	Initiator pulumi.StringPtrInput
	// Deprecated: instance_id is no longer used in this resource
	InstanceId pulumi.StringPtrInput
	// The IP address of the `hostName` above.
	IpAddress pulumi.StringPtrInput
	// Whether to connect to this volume via multipath.
	Multipath pulumi.BoolPtrInput
	// The iSCSI initiator OS type.
	OsType pulumi.StringPtrInput
	// The iSCSI initiator platform.
	Platform pulumi.StringPtrInput
	// The region in which to obtain the V2 Block Storage
	// client. A Block Storage client is needed to create a volume attachment.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new volume attachment.
	Region pulumi.StringPtrInput
	// The ID of the Volume to attach to an Instance.
	VolumeId pulumi.StringInput
	// A wwnn name. Used for Fibre Channel connections.
	Wwnn pulumi.StringPtrInput
	// An array of wwpn strings. Used for Fibre Channel
	// connections.
	Wwpns pulumi.StringArrayInput
}

func (VolumeAttachV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeAttachV2Args)(nil)).Elem()
}

type VolumeAttachV2Input interface {
	pulumi.Input

	ToVolumeAttachV2Output() VolumeAttachV2Output
	ToVolumeAttachV2OutputWithContext(ctx context.Context) VolumeAttachV2Output
}

func (*VolumeAttachV2) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachV2)(nil)).Elem()
}

func (i *VolumeAttachV2) ToVolumeAttachV2Output() VolumeAttachV2Output {
	return i.ToVolumeAttachV2OutputWithContext(context.Background())
}

func (i *VolumeAttachV2) ToVolumeAttachV2OutputWithContext(ctx context.Context) VolumeAttachV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachV2Output)
}

// VolumeAttachV2ArrayInput is an input type that accepts VolumeAttachV2Array and VolumeAttachV2ArrayOutput values.
// You can construct a concrete instance of `VolumeAttachV2ArrayInput` via:
//
//	VolumeAttachV2Array{ VolumeAttachV2Args{...} }
type VolumeAttachV2ArrayInput interface {
	pulumi.Input

	ToVolumeAttachV2ArrayOutput() VolumeAttachV2ArrayOutput
	ToVolumeAttachV2ArrayOutputWithContext(context.Context) VolumeAttachV2ArrayOutput
}

type VolumeAttachV2Array []VolumeAttachV2Input

func (VolumeAttachV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeAttachV2)(nil)).Elem()
}

func (i VolumeAttachV2Array) ToVolumeAttachV2ArrayOutput() VolumeAttachV2ArrayOutput {
	return i.ToVolumeAttachV2ArrayOutputWithContext(context.Background())
}

func (i VolumeAttachV2Array) ToVolumeAttachV2ArrayOutputWithContext(ctx context.Context) VolumeAttachV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachV2ArrayOutput)
}

// VolumeAttachV2MapInput is an input type that accepts VolumeAttachV2Map and VolumeAttachV2MapOutput values.
// You can construct a concrete instance of `VolumeAttachV2MapInput` via:
//
//	VolumeAttachV2Map{ "key": VolumeAttachV2Args{...} }
type VolumeAttachV2MapInput interface {
	pulumi.Input

	ToVolumeAttachV2MapOutput() VolumeAttachV2MapOutput
	ToVolumeAttachV2MapOutputWithContext(context.Context) VolumeAttachV2MapOutput
}

type VolumeAttachV2Map map[string]VolumeAttachV2Input

func (VolumeAttachV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeAttachV2)(nil)).Elem()
}

func (i VolumeAttachV2Map) ToVolumeAttachV2MapOutput() VolumeAttachV2MapOutput {
	return i.ToVolumeAttachV2MapOutputWithContext(context.Background())
}

func (i VolumeAttachV2Map) ToVolumeAttachV2MapOutputWithContext(ctx context.Context) VolumeAttachV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachV2MapOutput)
}

type VolumeAttachV2Output struct{ *pulumi.OutputState }

func (VolumeAttachV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeAttachV2)(nil)).Elem()
}

func (o VolumeAttachV2Output) ToVolumeAttachV2Output() VolumeAttachV2Output {
	return o
}

func (o VolumeAttachV2Output) ToVolumeAttachV2OutputWithContext(ctx context.Context) VolumeAttachV2Output {
	return o
}

// Specify whether to attach the volume as Read-Only
// (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
// If left unspecified, the Block Storage API will apply a default of `rw`.
func (o VolumeAttachV2Output) AttachMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringPtrOutput { return v.AttachMode }).(pulumi.StringPtrOutput)
}

// This is a map of key/value pairs that contain the connection
// information. You will want to pass this information to a provisioner
// script to finalize the connection. See below for more information.
func (o VolumeAttachV2Output) Data() pulumi.MapOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.MapOutput { return v.Data }).(pulumi.MapOutput)
}

// The device to tell the Block Storage service this
// volume will be attached as. This is purely for informational purposes.
// You can specify `auto` or a device such as `/dev/vdc`.
func (o VolumeAttachV2Output) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// The storage driver that the volume is based on.
func (o VolumeAttachV2Output) DriverVolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringOutput { return v.DriverVolumeType }).(pulumi.StringOutput)
}

// The host to attach the volume to.
func (o VolumeAttachV2Output) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The iSCSI initiator string to make the connection.
func (o VolumeAttachV2Output) Initiator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringPtrOutput { return v.Initiator }).(pulumi.StringPtrOutput)
}

// Deprecated: instance_id is no longer used in this resource
func (o VolumeAttachV2Output) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringPtrOutput { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// The IP address of the `hostName` above.
func (o VolumeAttachV2Output) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringPtrOutput { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// A mount point base name for shared storage.
func (o VolumeAttachV2Output) MountPointBase() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringOutput { return v.MountPointBase }).(pulumi.StringOutput)
}

// Whether to connect to this volume via multipath.
func (o VolumeAttachV2Output) Multipath() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.BoolPtrOutput { return v.Multipath }).(pulumi.BoolPtrOutput)
}

// The iSCSI initiator OS type.
func (o VolumeAttachV2Output) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// The iSCSI initiator platform.
func (o VolumeAttachV2Output) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringPtrOutput { return v.Platform }).(pulumi.StringPtrOutput)
}

// The region in which to obtain the V2 Block Storage
// client. A Block Storage client is needed to create a volume attachment.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new volume attachment.
func (o VolumeAttachV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The ID of the Volume to attach to an Instance.
func (o VolumeAttachV2Output) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

// A wwnn name. Used for Fibre Channel connections.
func (o VolumeAttachV2Output) Wwnn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringPtrOutput { return v.Wwnn }).(pulumi.StringPtrOutput)
}

// An array of wwpn strings. Used for Fibre Channel
// connections.
func (o VolumeAttachV2Output) Wwpns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VolumeAttachV2) pulumi.StringArrayOutput { return v.Wwpns }).(pulumi.StringArrayOutput)
}

type VolumeAttachV2ArrayOutput struct{ *pulumi.OutputState }

func (VolumeAttachV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeAttachV2)(nil)).Elem()
}

func (o VolumeAttachV2ArrayOutput) ToVolumeAttachV2ArrayOutput() VolumeAttachV2ArrayOutput {
	return o
}

func (o VolumeAttachV2ArrayOutput) ToVolumeAttachV2ArrayOutputWithContext(ctx context.Context) VolumeAttachV2ArrayOutput {
	return o
}

func (o VolumeAttachV2ArrayOutput) Index(i pulumi.IntInput) VolumeAttachV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VolumeAttachV2 {
		return vs[0].([]*VolumeAttachV2)[vs[1].(int)]
	}).(VolumeAttachV2Output)
}

type VolumeAttachV2MapOutput struct{ *pulumi.OutputState }

func (VolumeAttachV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeAttachV2)(nil)).Elem()
}

func (o VolumeAttachV2MapOutput) ToVolumeAttachV2MapOutput() VolumeAttachV2MapOutput {
	return o
}

func (o VolumeAttachV2MapOutput) ToVolumeAttachV2MapOutputWithContext(ctx context.Context) VolumeAttachV2MapOutput {
	return o
}

func (o VolumeAttachV2MapOutput) MapIndex(k pulumi.StringInput) VolumeAttachV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VolumeAttachV2 {
		return vs[0].(map[string]*VolumeAttachV2)[vs[1].(string)]
	}).(VolumeAttachV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachV2Input)(nil)).Elem(), &VolumeAttachV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachV2ArrayInput)(nil)).Elem(), VolumeAttachV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachV2MapInput)(nil)).Elem(), VolumeAttachV2Map{})
	pulumi.RegisterOutputType(VolumeAttachV2Output{})
	pulumi.RegisterOutputType(VolumeAttachV2ArrayOutput{})
	pulumi.RegisterOutputType(VolumeAttachV2MapOutput{})
}
