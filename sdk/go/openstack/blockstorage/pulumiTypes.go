// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package blockstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type VolumeAttachment struct {
	Device *string `pulumi:"device"`
	Id *string `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
}

type VolumeAttachmentInput interface {
	pulumi.Input

	ToVolumeAttachmentOutput() VolumeAttachmentOutput
	ToVolumeAttachmentOutputWithContext(context.Context) VolumeAttachmentOutput
}

type VolumeAttachmentArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
	Id pulumi.StringPtrInput `pulumi:"id"`
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

type VolumeAttachmentOutput struct { *pulumi.OutputState }

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
	return o.ApplyT(func (v VolumeAttachment) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o VolumeAttachmentOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeAttachment) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VolumeAttachmentOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeAttachment) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

type VolumeAttachmentArrayOutput struct { *pulumi.OutputState}

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
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) VolumeAttachment {
		return vs[0].([]VolumeAttachment)[vs[1].(int)]
	}).(VolumeAttachmentOutput)
}

type VolumeV1Attachment struct {
	Device *string `pulumi:"device"`
	Id *string `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
}

type VolumeV1AttachmentInput interface {
	pulumi.Input

	ToVolumeV1AttachmentOutput() VolumeV1AttachmentOutput
	ToVolumeV1AttachmentOutputWithContext(context.Context) VolumeV1AttachmentOutput
}

type VolumeV1AttachmentArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
	Id pulumi.StringPtrInput `pulumi:"id"`
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

type VolumeV1AttachmentOutput struct { *pulumi.OutputState }

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
	return o.ApplyT(func (v VolumeV1Attachment) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o VolumeV1AttachmentOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeV1Attachment) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VolumeV1AttachmentOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeV1Attachment) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

type VolumeV1AttachmentArrayOutput struct { *pulumi.OutputState}

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
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) VolumeV1Attachment {
		return vs[0].([]VolumeV1Attachment)[vs[1].(int)]
	}).(VolumeV1AttachmentOutput)
}

type VolumeV2Attachment struct {
	Device *string `pulumi:"device"`
	Id *string `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
}

type VolumeV2AttachmentInput interface {
	pulumi.Input

	ToVolumeV2AttachmentOutput() VolumeV2AttachmentOutput
	ToVolumeV2AttachmentOutputWithContext(context.Context) VolumeV2AttachmentOutput
}

type VolumeV2AttachmentArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
	Id pulumi.StringPtrInput `pulumi:"id"`
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

type VolumeV2AttachmentOutput struct { *pulumi.OutputState }

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
	return o.ApplyT(func (v VolumeV2Attachment) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o VolumeV2AttachmentOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeV2Attachment) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VolumeV2AttachmentOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VolumeV2Attachment) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

type VolumeV2AttachmentArrayOutput struct { *pulumi.OutputState}

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
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) VolumeV2Attachment {
		return vs[0].([]VolumeV2Attachment)[vs[1].(int)]
	}).(VolumeV2AttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeAttachmentOutput{})
	pulumi.RegisterOutputType(VolumeAttachmentArrayOutput{})
	pulumi.RegisterOutputType(VolumeV1AttachmentOutput{})
	pulumi.RegisterOutputType(VolumeV1AttachmentArrayOutput{})
	pulumi.RegisterOutputType(VolumeV2AttachmentOutput{})
	pulumi.RegisterOutputType(VolumeV2AttachmentArrayOutput{})
}
