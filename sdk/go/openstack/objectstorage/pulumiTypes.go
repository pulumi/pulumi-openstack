// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package objectstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ContainerVersioningLegacy struct {
	// Container in which versions will be stored.
	Location string `pulumi:"location"`
	// Versioning type which can be `versions` or `history`
	// according to [OpenStack
	// documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
	Type string `pulumi:"type"`
}

// ContainerVersioningLegacyInput is an input type that accepts ContainerVersioningLegacyArgs and ContainerVersioningLegacyOutput values.
// You can construct a concrete instance of `ContainerVersioningLegacyInput` via:
//
//	ContainerVersioningLegacyArgs{...}
type ContainerVersioningLegacyInput interface {
	pulumi.Input

	ToContainerVersioningLegacyOutput() ContainerVersioningLegacyOutput
	ToContainerVersioningLegacyOutputWithContext(context.Context) ContainerVersioningLegacyOutput
}

type ContainerVersioningLegacyArgs struct {
	// Container in which versions will be stored.
	Location pulumi.StringInput `pulumi:"location"`
	// Versioning type which can be `versions` or `history`
	// according to [OpenStack
	// documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
	Type pulumi.StringInput `pulumi:"type"`
}

func (ContainerVersioningLegacyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerVersioningLegacy)(nil)).Elem()
}

func (i ContainerVersioningLegacyArgs) ToContainerVersioningLegacyOutput() ContainerVersioningLegacyOutput {
	return i.ToContainerVersioningLegacyOutputWithContext(context.Background())
}

func (i ContainerVersioningLegacyArgs) ToContainerVersioningLegacyOutputWithContext(ctx context.Context) ContainerVersioningLegacyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerVersioningLegacyOutput)
}

func (i ContainerVersioningLegacyArgs) ToContainerVersioningLegacyPtrOutput() ContainerVersioningLegacyPtrOutput {
	return i.ToContainerVersioningLegacyPtrOutputWithContext(context.Background())
}

func (i ContainerVersioningLegacyArgs) ToContainerVersioningLegacyPtrOutputWithContext(ctx context.Context) ContainerVersioningLegacyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerVersioningLegacyOutput).ToContainerVersioningLegacyPtrOutputWithContext(ctx)
}

// ContainerVersioningLegacyPtrInput is an input type that accepts ContainerVersioningLegacyArgs, ContainerVersioningLegacyPtr and ContainerVersioningLegacyPtrOutput values.
// You can construct a concrete instance of `ContainerVersioningLegacyPtrInput` via:
//
//	        ContainerVersioningLegacyArgs{...}
//
//	or:
//
//	        nil
type ContainerVersioningLegacyPtrInput interface {
	pulumi.Input

	ToContainerVersioningLegacyPtrOutput() ContainerVersioningLegacyPtrOutput
	ToContainerVersioningLegacyPtrOutputWithContext(context.Context) ContainerVersioningLegacyPtrOutput
}

type containerVersioningLegacyPtrType ContainerVersioningLegacyArgs

func ContainerVersioningLegacyPtr(v *ContainerVersioningLegacyArgs) ContainerVersioningLegacyPtrInput {
	return (*containerVersioningLegacyPtrType)(v)
}

func (*containerVersioningLegacyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerVersioningLegacy)(nil)).Elem()
}

func (i *containerVersioningLegacyPtrType) ToContainerVersioningLegacyPtrOutput() ContainerVersioningLegacyPtrOutput {
	return i.ToContainerVersioningLegacyPtrOutputWithContext(context.Background())
}

func (i *containerVersioningLegacyPtrType) ToContainerVersioningLegacyPtrOutputWithContext(ctx context.Context) ContainerVersioningLegacyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerVersioningLegacyPtrOutput)
}

type ContainerVersioningLegacyOutput struct{ *pulumi.OutputState }

func (ContainerVersioningLegacyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerVersioningLegacy)(nil)).Elem()
}

func (o ContainerVersioningLegacyOutput) ToContainerVersioningLegacyOutput() ContainerVersioningLegacyOutput {
	return o
}

func (o ContainerVersioningLegacyOutput) ToContainerVersioningLegacyOutputWithContext(ctx context.Context) ContainerVersioningLegacyOutput {
	return o
}

func (o ContainerVersioningLegacyOutput) ToContainerVersioningLegacyPtrOutput() ContainerVersioningLegacyPtrOutput {
	return o.ToContainerVersioningLegacyPtrOutputWithContext(context.Background())
}

func (o ContainerVersioningLegacyOutput) ToContainerVersioningLegacyPtrOutputWithContext(ctx context.Context) ContainerVersioningLegacyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerVersioningLegacy) *ContainerVersioningLegacy {
		return &v
	}).(ContainerVersioningLegacyPtrOutput)
}

// Container in which versions will be stored.
func (o ContainerVersioningLegacyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerVersioningLegacy) string { return v.Location }).(pulumi.StringOutput)
}

// Versioning type which can be `versions` or `history`
// according to [OpenStack
// documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
func (o ContainerVersioningLegacyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerVersioningLegacy) string { return v.Type }).(pulumi.StringOutput)
}

type ContainerVersioningLegacyPtrOutput struct{ *pulumi.OutputState }

func (ContainerVersioningLegacyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerVersioningLegacy)(nil)).Elem()
}

func (o ContainerVersioningLegacyPtrOutput) ToContainerVersioningLegacyPtrOutput() ContainerVersioningLegacyPtrOutput {
	return o
}

func (o ContainerVersioningLegacyPtrOutput) ToContainerVersioningLegacyPtrOutputWithContext(ctx context.Context) ContainerVersioningLegacyPtrOutput {
	return o
}

func (o ContainerVersioningLegacyPtrOutput) Elem() ContainerVersioningLegacyOutput {
	return o.ApplyT(func(v *ContainerVersioningLegacy) ContainerVersioningLegacy {
		if v != nil {
			return *v
		}
		var ret ContainerVersioningLegacy
		return ret
	}).(ContainerVersioningLegacyOutput)
}

// Container in which versions will be stored.
func (o ContainerVersioningLegacyPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerVersioningLegacy) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

// Versioning type which can be `versions` or `history`
// according to [OpenStack
// documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
func (o ContainerVersioningLegacyPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerVersioningLegacy) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerVersioningLegacyInput)(nil)).Elem(), ContainerVersioningLegacyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerVersioningLegacyPtrInput)(nil)).Elem(), ContainerVersioningLegacyArgs{})
	pulumi.RegisterOutputType(ContainerVersioningLegacyOutput{})
	pulumi.RegisterOutputType(ContainerVersioningLegacyPtrOutput{})
}
