// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package orchestration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type StackV1StackOutput struct {
	// The description of the stack resource.
	Description *string `pulumi:"description"`
	OutputKey   string  `pulumi:"outputKey"`
	OutputValue string  `pulumi:"outputValue"`
}

// StackV1StackOutputInput is an input type that accepts StackV1StackOutputArgs and StackV1StackOutputOutput values.
// You can construct a concrete instance of `StackV1StackOutputInput` via:
//
//	StackV1StackOutputArgs{...}
type StackV1StackOutputInput interface {
	pulumi.Input

	ToStackV1StackOutputOutput() StackV1StackOutputOutput
	ToStackV1StackOutputOutputWithContext(context.Context) StackV1StackOutputOutput
}

type StackV1StackOutputArgs struct {
	// The description of the stack resource.
	Description pulumi.StringPtrInput `pulumi:"description"`
	OutputKey   pulumi.StringInput    `pulumi:"outputKey"`
	OutputValue pulumi.StringInput    `pulumi:"outputValue"`
}

func (StackV1StackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StackV1StackOutput)(nil)).Elem()
}

func (i StackV1StackOutputArgs) ToStackV1StackOutputOutput() StackV1StackOutputOutput {
	return i.ToStackV1StackOutputOutputWithContext(context.Background())
}

func (i StackV1StackOutputArgs) ToStackV1StackOutputOutputWithContext(ctx context.Context) StackV1StackOutputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackV1StackOutputOutput)
}

func (i StackV1StackOutputArgs) ToOutput(ctx context.Context) pulumix.Output[StackV1StackOutput] {
	return pulumix.Output[StackV1StackOutput]{
		OutputState: i.ToStackV1StackOutputOutputWithContext(ctx).OutputState,
	}
}

// StackV1StackOutputArrayInput is an input type that accepts StackV1StackOutputArray and StackV1StackOutputArrayOutput values.
// You can construct a concrete instance of `StackV1StackOutputArrayInput` via:
//
//	StackV1StackOutputArray{ StackV1StackOutputArgs{...} }
type StackV1StackOutputArrayInput interface {
	pulumi.Input

	ToStackV1StackOutputArrayOutput() StackV1StackOutputArrayOutput
	ToStackV1StackOutputArrayOutputWithContext(context.Context) StackV1StackOutputArrayOutput
}

type StackV1StackOutputArray []StackV1StackOutputInput

func (StackV1StackOutputArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StackV1StackOutput)(nil)).Elem()
}

func (i StackV1StackOutputArray) ToStackV1StackOutputArrayOutput() StackV1StackOutputArrayOutput {
	return i.ToStackV1StackOutputArrayOutputWithContext(context.Background())
}

func (i StackV1StackOutputArray) ToStackV1StackOutputArrayOutputWithContext(ctx context.Context) StackV1StackOutputArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackV1StackOutputArrayOutput)
}

func (i StackV1StackOutputArray) ToOutput(ctx context.Context) pulumix.Output[[]StackV1StackOutput] {
	return pulumix.Output[[]StackV1StackOutput]{
		OutputState: i.ToStackV1StackOutputArrayOutputWithContext(ctx).OutputState,
	}
}

type StackV1StackOutputOutput struct{ *pulumi.OutputState }

func (StackV1StackOutputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StackV1StackOutput)(nil)).Elem()
}

func (o StackV1StackOutputOutput) ToStackV1StackOutputOutput() StackV1StackOutputOutput {
	return o
}

func (o StackV1StackOutputOutput) ToStackV1StackOutputOutputWithContext(ctx context.Context) StackV1StackOutputOutput {
	return o
}

func (o StackV1StackOutputOutput) ToOutput(ctx context.Context) pulumix.Output[StackV1StackOutput] {
	return pulumix.Output[StackV1StackOutput]{
		OutputState: o.OutputState,
	}
}

// The description of the stack resource.
func (o StackV1StackOutputOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StackV1StackOutput) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StackV1StackOutputOutput) OutputKey() pulumi.StringOutput {
	return o.ApplyT(func(v StackV1StackOutput) string { return v.OutputKey }).(pulumi.StringOutput)
}

func (o StackV1StackOutputOutput) OutputValue() pulumi.StringOutput {
	return o.ApplyT(func(v StackV1StackOutput) string { return v.OutputValue }).(pulumi.StringOutput)
}

type StackV1StackOutputArrayOutput struct{ *pulumi.OutputState }

func (StackV1StackOutputArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StackV1StackOutput)(nil)).Elem()
}

func (o StackV1StackOutputArrayOutput) ToStackV1StackOutputArrayOutput() StackV1StackOutputArrayOutput {
	return o
}

func (o StackV1StackOutputArrayOutput) ToStackV1StackOutputArrayOutputWithContext(ctx context.Context) StackV1StackOutputArrayOutput {
	return o
}

func (o StackV1StackOutputArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]StackV1StackOutput] {
	return pulumix.Output[[]StackV1StackOutput]{
		OutputState: o.OutputState,
	}
}

func (o StackV1StackOutputArrayOutput) Index(i pulumi.IntInput) StackV1StackOutputOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StackV1StackOutput {
		return vs[0].([]StackV1StackOutput)[vs[1].(int)]
	}).(StackV1StackOutputOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackV1StackOutputInput)(nil)).Elem(), StackV1StackOutputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackV1StackOutputArrayInput)(nil)).Elem(), StackV1StackOutputArray{})
	pulumi.RegisterOutputType(StackV1StackOutputOutput{})
	pulumi.RegisterOutputType(StackV1StackOutputArrayOutput{})
}
