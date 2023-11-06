// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type VolumeAttachment struct {
	Device     *string `pulumi:"device"`
	Id         *string `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
}

// VolumeAttachmentInput is an input type that accepts VolumeAttachmentArgs and VolumeAttachmentOutput values.
// You can construct a concrete instance of `VolumeAttachmentInput` via:
//
//	VolumeAttachmentArgs{...}
type VolumeAttachmentInput interface {
	pulumi.Input

	ToVolumeAttachmentOutput() VolumeAttachmentOutput
	ToVolumeAttachmentOutputWithContext(context.Context) VolumeAttachmentOutput
}

type VolumeAttachmentArgs struct {
	Device     pulumi.StringPtrInput `pulumi:"device"`
	Id         pulumi.StringPtrInput `pulumi:"id"`
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
}

func (VolumeAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachment)(nil)).Elem()
}

func (i VolumeAttachmentArgs) ToVolumeAttachmentOutput() VolumeAttachmentOutput {
	return i.ToVolumeAttachmentOutputWithContext(context.Background())
}

func (i VolumeAttachmentArgs) ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentOutput)
}

// VolumeAttachmentArrayInput is an input type that accepts VolumeAttachmentArray and VolumeAttachmentArrayOutput values.
// You can construct a concrete instance of `VolumeAttachmentArrayInput` via:
//
//	VolumeAttachmentArray{ VolumeAttachmentArgs{...} }
type VolumeAttachmentArrayInput interface {
	pulumi.Input

	ToVolumeAttachmentArrayOutput() VolumeAttachmentArrayOutput
	ToVolumeAttachmentArrayOutputWithContext(context.Context) VolumeAttachmentArrayOutput
}

type VolumeAttachmentArray []VolumeAttachmentInput

func (VolumeAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeAttachment)(nil)).Elem()
}

func (i VolumeAttachmentArray) ToVolumeAttachmentArrayOutput() VolumeAttachmentArrayOutput {
	return i.ToVolumeAttachmentArrayOutputWithContext(context.Background())
}

func (i VolumeAttachmentArray) ToVolumeAttachmentArrayOutputWithContext(ctx context.Context) VolumeAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeAttachmentArrayOutput)
}

type VolumeAttachmentOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeAttachment)(nil)).Elem()
}

func (o VolumeAttachmentOutput) ToVolumeAttachmentOutput() VolumeAttachmentOutput {
	return o
}

func (o VolumeAttachmentOutput) ToVolumeAttachmentOutputWithContext(ctx context.Context) VolumeAttachmentOutput {
	return o
}

func (o VolumeAttachmentOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeAttachment) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o VolumeAttachmentOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeAttachment) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VolumeAttachmentOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeAttachment) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

type VolumeAttachmentArrayOutput struct{ *pulumi.OutputState }

func (VolumeAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeAttachment)(nil)).Elem()
}

func (o VolumeAttachmentArrayOutput) ToVolumeAttachmentArrayOutput() VolumeAttachmentArrayOutput {
	return o
}

func (o VolumeAttachmentArrayOutput) ToVolumeAttachmentArrayOutputWithContext(ctx context.Context) VolumeAttachmentArrayOutput {
	return o
}

func (o VolumeAttachmentArrayOutput) Index(i pulumi.IntInput) VolumeAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeAttachment {
		return vs[0].([]VolumeAttachment)[vs[1].(int)]
	}).(VolumeAttachmentOutput)
}

type VolumeSchedulerHint struct {
	// Arbitrary key/value pairs of additional
	// properties to pass to the scheduler.
	AdditionalProperties map[string]interface{} `pulumi:"additionalProperties"`
	// The volume should be scheduled on a
	// different host from the set of volumes specified in the list provided.
	DifferentHosts []string `pulumi:"differentHosts"`
	// An instance UUID. The volume should be
	// scheduled on the same host as the instance.
	LocalToInstance *string `pulumi:"localToInstance"`
	// A conditional query that a back-end must pass in
	// order to host a volume. The query must use the `JsonFilter` syntax
	// which is described
	// [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
	// At this time, only simple queries are supported. Compound queries using
	// `and`, `or`, or `not` are not supported. An example of a simple query is:
	Query *string `pulumi:"query"`
	// A list of volume UUIDs. The volume should be
	// scheduled on the same host as another volume specified in the list provided.
	SameHosts []string `pulumi:"sameHosts"`
}

// VolumeSchedulerHintInput is an input type that accepts VolumeSchedulerHintArgs and VolumeSchedulerHintOutput values.
// You can construct a concrete instance of `VolumeSchedulerHintInput` via:
//
//	VolumeSchedulerHintArgs{...}
type VolumeSchedulerHintInput interface {
	pulumi.Input

	ToVolumeSchedulerHintOutput() VolumeSchedulerHintOutput
	ToVolumeSchedulerHintOutputWithContext(context.Context) VolumeSchedulerHintOutput
}

type VolumeSchedulerHintArgs struct {
	// Arbitrary key/value pairs of additional
	// properties to pass to the scheduler.
	AdditionalProperties pulumi.MapInput `pulumi:"additionalProperties"`
	// The volume should be scheduled on a
	// different host from the set of volumes specified in the list provided.
	DifferentHosts pulumi.StringArrayInput `pulumi:"differentHosts"`
	// An instance UUID. The volume should be
	// scheduled on the same host as the instance.
	LocalToInstance pulumi.StringPtrInput `pulumi:"localToInstance"`
	// A conditional query that a back-end must pass in
	// order to host a volume. The query must use the `JsonFilter` syntax
	// which is described
	// [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
	// At this time, only simple queries are supported. Compound queries using
	// `and`, `or`, or `not` are not supported. An example of a simple query is:
	Query pulumi.StringPtrInput `pulumi:"query"`
	// A list of volume UUIDs. The volume should be
	// scheduled on the same host as another volume specified in the list provided.
	SameHosts pulumi.StringArrayInput `pulumi:"sameHosts"`
}

func (VolumeSchedulerHintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSchedulerHint)(nil)).Elem()
}

func (i VolumeSchedulerHintArgs) ToVolumeSchedulerHintOutput() VolumeSchedulerHintOutput {
	return i.ToVolumeSchedulerHintOutputWithContext(context.Background())
}

func (i VolumeSchedulerHintArgs) ToVolumeSchedulerHintOutputWithContext(ctx context.Context) VolumeSchedulerHintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSchedulerHintOutput)
}

// VolumeSchedulerHintArrayInput is an input type that accepts VolumeSchedulerHintArray and VolumeSchedulerHintArrayOutput values.
// You can construct a concrete instance of `VolumeSchedulerHintArrayInput` via:
//
//	VolumeSchedulerHintArray{ VolumeSchedulerHintArgs{...} }
type VolumeSchedulerHintArrayInput interface {
	pulumi.Input

	ToVolumeSchedulerHintArrayOutput() VolumeSchedulerHintArrayOutput
	ToVolumeSchedulerHintArrayOutputWithContext(context.Context) VolumeSchedulerHintArrayOutput
}

type VolumeSchedulerHintArray []VolumeSchedulerHintInput

func (VolumeSchedulerHintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeSchedulerHint)(nil)).Elem()
}

func (i VolumeSchedulerHintArray) ToVolumeSchedulerHintArrayOutput() VolumeSchedulerHintArrayOutput {
	return i.ToVolumeSchedulerHintArrayOutputWithContext(context.Background())
}

func (i VolumeSchedulerHintArray) ToVolumeSchedulerHintArrayOutputWithContext(ctx context.Context) VolumeSchedulerHintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSchedulerHintArrayOutput)
}

type VolumeSchedulerHintOutput struct{ *pulumi.OutputState }

func (VolumeSchedulerHintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSchedulerHint)(nil)).Elem()
}

func (o VolumeSchedulerHintOutput) ToVolumeSchedulerHintOutput() VolumeSchedulerHintOutput {
	return o
}

func (o VolumeSchedulerHintOutput) ToVolumeSchedulerHintOutputWithContext(ctx context.Context) VolumeSchedulerHintOutput {
	return o
}

// Arbitrary key/value pairs of additional
// properties to pass to the scheduler.
func (o VolumeSchedulerHintOutput) AdditionalProperties() pulumi.MapOutput {
	return o.ApplyT(func(v VolumeSchedulerHint) map[string]interface{} { return v.AdditionalProperties }).(pulumi.MapOutput)
}

// The volume should be scheduled on a
// different host from the set of volumes specified in the list provided.
func (o VolumeSchedulerHintOutput) DifferentHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VolumeSchedulerHint) []string { return v.DifferentHosts }).(pulumi.StringArrayOutput)
}

// An instance UUID. The volume should be
// scheduled on the same host as the instance.
func (o VolumeSchedulerHintOutput) LocalToInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeSchedulerHint) *string { return v.LocalToInstance }).(pulumi.StringPtrOutput)
}

// A conditional query that a back-end must pass in
// order to host a volume. The query must use the `JsonFilter` syntax
// which is described
// [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
// At this time, only simple queries are supported. Compound queries using
// `and`, `or`, or `not` are not supported. An example of a simple query is:
func (o VolumeSchedulerHintOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeSchedulerHint) *string { return v.Query }).(pulumi.StringPtrOutput)
}

// A list of volume UUIDs. The volume should be
// scheduled on the same host as another volume specified in the list provided.
func (o VolumeSchedulerHintOutput) SameHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VolumeSchedulerHint) []string { return v.SameHosts }).(pulumi.StringArrayOutput)
}

type VolumeSchedulerHintArrayOutput struct{ *pulumi.OutputState }

func (VolumeSchedulerHintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeSchedulerHint)(nil)).Elem()
}

func (o VolumeSchedulerHintArrayOutput) ToVolumeSchedulerHintArrayOutput() VolumeSchedulerHintArrayOutput {
	return o
}

func (o VolumeSchedulerHintArrayOutput) ToVolumeSchedulerHintArrayOutputWithContext(ctx context.Context) VolumeSchedulerHintArrayOutput {
	return o
}

func (o VolumeSchedulerHintArrayOutput) Index(i pulumi.IntInput) VolumeSchedulerHintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeSchedulerHint {
		return vs[0].([]VolumeSchedulerHint)[vs[1].(int)]
	}).(VolumeSchedulerHintOutput)
}

type VolumeV1Attachment struct {
	Device     *string `pulumi:"device"`
	Id         *string `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
}

// VolumeV1AttachmentInput is an input type that accepts VolumeV1AttachmentArgs and VolumeV1AttachmentOutput values.
// You can construct a concrete instance of `VolumeV1AttachmentInput` via:
//
//	VolumeV1AttachmentArgs{...}
type VolumeV1AttachmentInput interface {
	pulumi.Input

	ToVolumeV1AttachmentOutput() VolumeV1AttachmentOutput
	ToVolumeV1AttachmentOutputWithContext(context.Context) VolumeV1AttachmentOutput
}

type VolumeV1AttachmentArgs struct {
	Device     pulumi.StringPtrInput `pulumi:"device"`
	Id         pulumi.StringPtrInput `pulumi:"id"`
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
}

func (VolumeV1AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeV1Attachment)(nil)).Elem()
}

func (i VolumeV1AttachmentArgs) ToVolumeV1AttachmentOutput() VolumeV1AttachmentOutput {
	return i.ToVolumeV1AttachmentOutputWithContext(context.Background())
}

func (i VolumeV1AttachmentArgs) ToVolumeV1AttachmentOutputWithContext(ctx context.Context) VolumeV1AttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV1AttachmentOutput)
}

// VolumeV1AttachmentArrayInput is an input type that accepts VolumeV1AttachmentArray and VolumeV1AttachmentArrayOutput values.
// You can construct a concrete instance of `VolumeV1AttachmentArrayInput` via:
//
//	VolumeV1AttachmentArray{ VolumeV1AttachmentArgs{...} }
type VolumeV1AttachmentArrayInput interface {
	pulumi.Input

	ToVolumeV1AttachmentArrayOutput() VolumeV1AttachmentArrayOutput
	ToVolumeV1AttachmentArrayOutputWithContext(context.Context) VolumeV1AttachmentArrayOutput
}

type VolumeV1AttachmentArray []VolumeV1AttachmentInput

func (VolumeV1AttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeV1Attachment)(nil)).Elem()
}

func (i VolumeV1AttachmentArray) ToVolumeV1AttachmentArrayOutput() VolumeV1AttachmentArrayOutput {
	return i.ToVolumeV1AttachmentArrayOutputWithContext(context.Background())
}

func (i VolumeV1AttachmentArray) ToVolumeV1AttachmentArrayOutputWithContext(ctx context.Context) VolumeV1AttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV1AttachmentArrayOutput)
}

type VolumeV1AttachmentOutput struct{ *pulumi.OutputState }

func (VolumeV1AttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeV1Attachment)(nil)).Elem()
}

func (o VolumeV1AttachmentOutput) ToVolumeV1AttachmentOutput() VolumeV1AttachmentOutput {
	return o
}

func (o VolumeV1AttachmentOutput) ToVolumeV1AttachmentOutputWithContext(ctx context.Context) VolumeV1AttachmentOutput {
	return o
}

func (o VolumeV1AttachmentOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeV1Attachment) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o VolumeV1AttachmentOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeV1Attachment) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VolumeV1AttachmentOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeV1Attachment) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

type VolumeV1AttachmentArrayOutput struct{ *pulumi.OutputState }

func (VolumeV1AttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeV1Attachment)(nil)).Elem()
}

func (o VolumeV1AttachmentArrayOutput) ToVolumeV1AttachmentArrayOutput() VolumeV1AttachmentArrayOutput {
	return o
}

func (o VolumeV1AttachmentArrayOutput) ToVolumeV1AttachmentArrayOutputWithContext(ctx context.Context) VolumeV1AttachmentArrayOutput {
	return o
}

func (o VolumeV1AttachmentArrayOutput) Index(i pulumi.IntInput) VolumeV1AttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeV1Attachment {
		return vs[0].([]VolumeV1Attachment)[vs[1].(int)]
	}).(VolumeV1AttachmentOutput)
}

type VolumeV2Attachment struct {
	Device     *string `pulumi:"device"`
	Id         *string `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
}

// VolumeV2AttachmentInput is an input type that accepts VolumeV2AttachmentArgs and VolumeV2AttachmentOutput values.
// You can construct a concrete instance of `VolumeV2AttachmentInput` via:
//
//	VolumeV2AttachmentArgs{...}
type VolumeV2AttachmentInput interface {
	pulumi.Input

	ToVolumeV2AttachmentOutput() VolumeV2AttachmentOutput
	ToVolumeV2AttachmentOutputWithContext(context.Context) VolumeV2AttachmentOutput
}

type VolumeV2AttachmentArgs struct {
	Device     pulumi.StringPtrInput `pulumi:"device"`
	Id         pulumi.StringPtrInput `pulumi:"id"`
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
}

func (VolumeV2AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeV2Attachment)(nil)).Elem()
}

func (i VolumeV2AttachmentArgs) ToVolumeV2AttachmentOutput() VolumeV2AttachmentOutput {
	return i.ToVolumeV2AttachmentOutputWithContext(context.Background())
}

func (i VolumeV2AttachmentArgs) ToVolumeV2AttachmentOutputWithContext(ctx context.Context) VolumeV2AttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV2AttachmentOutput)
}

// VolumeV2AttachmentArrayInput is an input type that accepts VolumeV2AttachmentArray and VolumeV2AttachmentArrayOutput values.
// You can construct a concrete instance of `VolumeV2AttachmentArrayInput` via:
//
//	VolumeV2AttachmentArray{ VolumeV2AttachmentArgs{...} }
type VolumeV2AttachmentArrayInput interface {
	pulumi.Input

	ToVolumeV2AttachmentArrayOutput() VolumeV2AttachmentArrayOutput
	ToVolumeV2AttachmentArrayOutputWithContext(context.Context) VolumeV2AttachmentArrayOutput
}

type VolumeV2AttachmentArray []VolumeV2AttachmentInput

func (VolumeV2AttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeV2Attachment)(nil)).Elem()
}

func (i VolumeV2AttachmentArray) ToVolumeV2AttachmentArrayOutput() VolumeV2AttachmentArrayOutput {
	return i.ToVolumeV2AttachmentArrayOutputWithContext(context.Background())
}

func (i VolumeV2AttachmentArray) ToVolumeV2AttachmentArrayOutputWithContext(ctx context.Context) VolumeV2AttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV2AttachmentArrayOutput)
}

type VolumeV2AttachmentOutput struct{ *pulumi.OutputState }

func (VolumeV2AttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeV2Attachment)(nil)).Elem()
}

func (o VolumeV2AttachmentOutput) ToVolumeV2AttachmentOutput() VolumeV2AttachmentOutput {
	return o
}

func (o VolumeV2AttachmentOutput) ToVolumeV2AttachmentOutputWithContext(ctx context.Context) VolumeV2AttachmentOutput {
	return o
}

func (o VolumeV2AttachmentOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeV2Attachment) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o VolumeV2AttachmentOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeV2Attachment) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VolumeV2AttachmentOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeV2Attachment) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

type VolumeV2AttachmentArrayOutput struct{ *pulumi.OutputState }

func (VolumeV2AttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeV2Attachment)(nil)).Elem()
}

func (o VolumeV2AttachmentArrayOutput) ToVolumeV2AttachmentArrayOutput() VolumeV2AttachmentArrayOutput {
	return o
}

func (o VolumeV2AttachmentArrayOutput) ToVolumeV2AttachmentArrayOutputWithContext(ctx context.Context) VolumeV2AttachmentArrayOutput {
	return o
}

func (o VolumeV2AttachmentArrayOutput) Index(i pulumi.IntInput) VolumeV2AttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeV2Attachment {
		return vs[0].([]VolumeV2Attachment)[vs[1].(int)]
	}).(VolumeV2AttachmentOutput)
}

type VolumeV2SchedulerHint struct {
	// Arbitrary key/value pairs of additional
	// properties to pass to the scheduler.
	AdditionalProperties map[string]interface{} `pulumi:"additionalProperties"`
	// The volume should be scheduled on a
	// different host from the set of volumes specified in the list provided.
	DifferentHosts []string `pulumi:"differentHosts"`
	// An instance UUID. The volume should be
	// scheduled on the same host as the instance.
	LocalToInstance *string `pulumi:"localToInstance"`
	// A conditional query that a back-end must pass in
	// order to host a volume. The query must use the `JsonFilter` syntax
	// which is described
	// [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
	// At this time, only simple queries are supported. Compound queries using
	// `and`, `or`, or `not` are not supported. An example of a simple query is:
	Query *string `pulumi:"query"`
	// A list of volume UUIDs. The volume should be
	// scheduled on the same host as another volume specified in the list provided.
	SameHosts []string `pulumi:"sameHosts"`
}

// VolumeV2SchedulerHintInput is an input type that accepts VolumeV2SchedulerHintArgs and VolumeV2SchedulerHintOutput values.
// You can construct a concrete instance of `VolumeV2SchedulerHintInput` via:
//
//	VolumeV2SchedulerHintArgs{...}
type VolumeV2SchedulerHintInput interface {
	pulumi.Input

	ToVolumeV2SchedulerHintOutput() VolumeV2SchedulerHintOutput
	ToVolumeV2SchedulerHintOutputWithContext(context.Context) VolumeV2SchedulerHintOutput
}

type VolumeV2SchedulerHintArgs struct {
	// Arbitrary key/value pairs of additional
	// properties to pass to the scheduler.
	AdditionalProperties pulumi.MapInput `pulumi:"additionalProperties"`
	// The volume should be scheduled on a
	// different host from the set of volumes specified in the list provided.
	DifferentHosts pulumi.StringArrayInput `pulumi:"differentHosts"`
	// An instance UUID. The volume should be
	// scheduled on the same host as the instance.
	LocalToInstance pulumi.StringPtrInput `pulumi:"localToInstance"`
	// A conditional query that a back-end must pass in
	// order to host a volume. The query must use the `JsonFilter` syntax
	// which is described
	// [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
	// At this time, only simple queries are supported. Compound queries using
	// `and`, `or`, or `not` are not supported. An example of a simple query is:
	Query pulumi.StringPtrInput `pulumi:"query"`
	// A list of volume UUIDs. The volume should be
	// scheduled on the same host as another volume specified in the list provided.
	SameHosts pulumi.StringArrayInput `pulumi:"sameHosts"`
}

func (VolumeV2SchedulerHintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeV2SchedulerHint)(nil)).Elem()
}

func (i VolumeV2SchedulerHintArgs) ToVolumeV2SchedulerHintOutput() VolumeV2SchedulerHintOutput {
	return i.ToVolumeV2SchedulerHintOutputWithContext(context.Background())
}

func (i VolumeV2SchedulerHintArgs) ToVolumeV2SchedulerHintOutputWithContext(ctx context.Context) VolumeV2SchedulerHintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV2SchedulerHintOutput)
}

// VolumeV2SchedulerHintArrayInput is an input type that accepts VolumeV2SchedulerHintArray and VolumeV2SchedulerHintArrayOutput values.
// You can construct a concrete instance of `VolumeV2SchedulerHintArrayInput` via:
//
//	VolumeV2SchedulerHintArray{ VolumeV2SchedulerHintArgs{...} }
type VolumeV2SchedulerHintArrayInput interface {
	pulumi.Input

	ToVolumeV2SchedulerHintArrayOutput() VolumeV2SchedulerHintArrayOutput
	ToVolumeV2SchedulerHintArrayOutputWithContext(context.Context) VolumeV2SchedulerHintArrayOutput
}

type VolumeV2SchedulerHintArray []VolumeV2SchedulerHintInput

func (VolumeV2SchedulerHintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeV2SchedulerHint)(nil)).Elem()
}

func (i VolumeV2SchedulerHintArray) ToVolumeV2SchedulerHintArrayOutput() VolumeV2SchedulerHintArrayOutput {
	return i.ToVolumeV2SchedulerHintArrayOutputWithContext(context.Background())
}

func (i VolumeV2SchedulerHintArray) ToVolumeV2SchedulerHintArrayOutputWithContext(ctx context.Context) VolumeV2SchedulerHintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeV2SchedulerHintArrayOutput)
}

type VolumeV2SchedulerHintOutput struct{ *pulumi.OutputState }

func (VolumeV2SchedulerHintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeV2SchedulerHint)(nil)).Elem()
}

func (o VolumeV2SchedulerHintOutput) ToVolumeV2SchedulerHintOutput() VolumeV2SchedulerHintOutput {
	return o
}

func (o VolumeV2SchedulerHintOutput) ToVolumeV2SchedulerHintOutputWithContext(ctx context.Context) VolumeV2SchedulerHintOutput {
	return o
}

// Arbitrary key/value pairs of additional
// properties to pass to the scheduler.
func (o VolumeV2SchedulerHintOutput) AdditionalProperties() pulumi.MapOutput {
	return o.ApplyT(func(v VolumeV2SchedulerHint) map[string]interface{} { return v.AdditionalProperties }).(pulumi.MapOutput)
}

// The volume should be scheduled on a
// different host from the set of volumes specified in the list provided.
func (o VolumeV2SchedulerHintOutput) DifferentHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VolumeV2SchedulerHint) []string { return v.DifferentHosts }).(pulumi.StringArrayOutput)
}

// An instance UUID. The volume should be
// scheduled on the same host as the instance.
func (o VolumeV2SchedulerHintOutput) LocalToInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeV2SchedulerHint) *string { return v.LocalToInstance }).(pulumi.StringPtrOutput)
}

// A conditional query that a back-end must pass in
// order to host a volume. The query must use the `JsonFilter` syntax
// which is described
// [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
// At this time, only simple queries are supported. Compound queries using
// `and`, `or`, or `not` are not supported. An example of a simple query is:
func (o VolumeV2SchedulerHintOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeV2SchedulerHint) *string { return v.Query }).(pulumi.StringPtrOutput)
}

// A list of volume UUIDs. The volume should be
// scheduled on the same host as another volume specified in the list provided.
func (o VolumeV2SchedulerHintOutput) SameHosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VolumeV2SchedulerHint) []string { return v.SameHosts }).(pulumi.StringArrayOutput)
}

type VolumeV2SchedulerHintArrayOutput struct{ *pulumi.OutputState }

func (VolumeV2SchedulerHintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeV2SchedulerHint)(nil)).Elem()
}

func (o VolumeV2SchedulerHintArrayOutput) ToVolumeV2SchedulerHintArrayOutput() VolumeV2SchedulerHintArrayOutput {
	return o
}

func (o VolumeV2SchedulerHintArrayOutput) ToVolumeV2SchedulerHintArrayOutputWithContext(ctx context.Context) VolumeV2SchedulerHintArrayOutput {
	return o
}

func (o VolumeV2SchedulerHintArrayOutput) Index(i pulumi.IntInput) VolumeV2SchedulerHintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeV2SchedulerHint {
		return vs[0].([]VolumeV2SchedulerHint)[vs[1].(int)]
	}).(VolumeV2SchedulerHintOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentInput)(nil)).Elem(), VolumeAttachmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeAttachmentArrayInput)(nil)).Elem(), VolumeAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeSchedulerHintInput)(nil)).Elem(), VolumeSchedulerHintArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeSchedulerHintArrayInput)(nil)).Elem(), VolumeSchedulerHintArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV1AttachmentInput)(nil)).Elem(), VolumeV1AttachmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV1AttachmentArrayInput)(nil)).Elem(), VolumeV1AttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV2AttachmentInput)(nil)).Elem(), VolumeV2AttachmentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV2AttachmentArrayInput)(nil)).Elem(), VolumeV2AttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV2SchedulerHintInput)(nil)).Elem(), VolumeV2SchedulerHintArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeV2SchedulerHintArrayInput)(nil)).Elem(), VolumeV2SchedulerHintArray{})
	pulumi.RegisterOutputType(VolumeAttachmentOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentArrayOutput{})
	pulumi.RegisterOutputType(VolumeSchedulerHintOutput{})
	pulumi.RegisterOutputType(VolumeSchedulerHintArrayOutput{})
	pulumi.RegisterOutputType(VolumeV1AttachmentOutput{})
	pulumi.RegisterOutputType(VolumeV1AttachmentArrayOutput{})
	pulumi.RegisterOutputType(VolumeV2AttachmentOutput{})
	pulumi.RegisterOutputType(VolumeV2AttachmentArrayOutput{})
	pulumi.RegisterOutputType(VolumeV2SchedulerHintOutput{})
	pulumi.RegisterOutputType(VolumeV2SchedulerHintArrayOutput{})
}
