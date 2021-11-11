// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V3 block storage volume type resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := blockstorage.NewVolumeTypeV3(ctx, "volumeType1", &blockstorage.VolumeTypeV3Args{
// 			Description: pulumi.String("Volume type 1"),
// 			ExtraSpecs: pulumi.AnyMap{
// 				"capabilities":        pulumi.Any("gpu"),
// 				"volume_backend_name": pulumi.Any("ssd"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Volume types can be imported using the `volume_type_id`, e.g.
//
// ```sh
//  $ pulumi import openstack:blockstorage/volumeTypeV3:VolumeTypeV3 volume_type_1 941793f0-0a34-4bc4-b72e-a6326ae58283
// ```
type VolumeTypeV3 struct {
	pulumi.CustomResourceState

	// Human-readable description of the port. Changing
	// this updates the `description` of an existing volume type.
	Description pulumi.StringOutput `pulumi:"description"`
	// Key/Value pairs of metadata for the volume type.
	ExtraSpecs pulumi.MapOutput `pulumi:"extraSpecs"`
	// Whether the volume type is public. Changing
	// this updates the `isPublic` of an existing volume type.
	IsPublic pulumi.BoolOutput `pulumi:"isPublic"`
	// Name of the volume type.  Changing this
	// updates the `name` of an existing volume type.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewVolumeTypeV3 registers a new resource with the given unique name, arguments, and options.
func NewVolumeTypeV3(ctx *pulumi.Context,
	name string, args *VolumeTypeV3Args, opts ...pulumi.ResourceOption) (*VolumeTypeV3, error) {
	if args == nil {
		args = &VolumeTypeV3Args{}
	}

	var resource VolumeTypeV3
	err := ctx.RegisterResource("openstack:blockstorage/volumeTypeV3:VolumeTypeV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeTypeV3 gets an existing VolumeTypeV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeTypeV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeTypeV3State, opts ...pulumi.ResourceOption) (*VolumeTypeV3, error) {
	var resource VolumeTypeV3
	err := ctx.ReadResource("openstack:blockstorage/volumeTypeV3:VolumeTypeV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeTypeV3 resources.
type volumeTypeV3State struct {
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing volume type.
	Description *string `pulumi:"description"`
	// Key/Value pairs of metadata for the volume type.
	ExtraSpecs map[string]interface{} `pulumi:"extraSpecs"`
	// Whether the volume type is public. Changing
	// this updates the `isPublic` of an existing volume type.
	IsPublic *bool `pulumi:"isPublic"`
	// Name of the volume type.  Changing this
	// updates the `name` of an existing volume type.
	Name *string `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `pulumi:"region"`
}

type VolumeTypeV3State struct {
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing volume type.
	Description pulumi.StringPtrInput
	// Key/Value pairs of metadata for the volume type.
	ExtraSpecs pulumi.MapInput
	// Whether the volume type is public. Changing
	// this updates the `isPublic` of an existing volume type.
	IsPublic pulumi.BoolPtrInput
	// Name of the volume type.  Changing this
	// updates the `name` of an existing volume type.
	Name pulumi.StringPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringPtrInput
}

func (VolumeTypeV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeTypeV3State)(nil)).Elem()
}

type volumeTypeV3Args struct {
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing volume type.
	Description *string `pulumi:"description"`
	// Key/Value pairs of metadata for the volume type.
	ExtraSpecs map[string]interface{} `pulumi:"extraSpecs"`
	// Whether the volume type is public. Changing
	// this updates the `isPublic` of an existing volume type.
	IsPublic *bool `pulumi:"isPublic"`
	// Name of the volume type.  Changing this
	// updates the `name` of an existing volume type.
	Name *string `pulumi:"name"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a VolumeTypeV3 resource.
type VolumeTypeV3Args struct {
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing volume type.
	Description pulumi.StringPtrInput
	// Key/Value pairs of metadata for the volume type.
	ExtraSpecs pulumi.MapInput
	// Whether the volume type is public. Changing
	// this updates the `isPublic` of an existing volume type.
	IsPublic pulumi.BoolPtrInput
	// Name of the volume type.  Changing this
	// updates the `name` of an existing volume type.
	Name pulumi.StringPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringPtrInput
}

func (VolumeTypeV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeTypeV3Args)(nil)).Elem()
}

type VolumeTypeV3Input interface {
	pulumi.Input

	ToVolumeTypeV3Output() VolumeTypeV3Output
	ToVolumeTypeV3OutputWithContext(ctx context.Context) VolumeTypeV3Output
}

func (*VolumeTypeV3) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeTypeV3)(nil))
}

func (i *VolumeTypeV3) ToVolumeTypeV3Output() VolumeTypeV3Output {
	return i.ToVolumeTypeV3OutputWithContext(context.Background())
}

func (i *VolumeTypeV3) ToVolumeTypeV3OutputWithContext(ctx context.Context) VolumeTypeV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeTypeV3Output)
}

func (i *VolumeTypeV3) ToVolumeTypeV3PtrOutput() VolumeTypeV3PtrOutput {
	return i.ToVolumeTypeV3PtrOutputWithContext(context.Background())
}

func (i *VolumeTypeV3) ToVolumeTypeV3PtrOutputWithContext(ctx context.Context) VolumeTypeV3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeTypeV3PtrOutput)
}

type VolumeTypeV3PtrInput interface {
	pulumi.Input

	ToVolumeTypeV3PtrOutput() VolumeTypeV3PtrOutput
	ToVolumeTypeV3PtrOutputWithContext(ctx context.Context) VolumeTypeV3PtrOutput
}

type volumeTypeV3PtrType VolumeTypeV3Args

func (*volumeTypeV3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeTypeV3)(nil))
}

func (i *volumeTypeV3PtrType) ToVolumeTypeV3PtrOutput() VolumeTypeV3PtrOutput {
	return i.ToVolumeTypeV3PtrOutputWithContext(context.Background())
}

func (i *volumeTypeV3PtrType) ToVolumeTypeV3PtrOutputWithContext(ctx context.Context) VolumeTypeV3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeTypeV3PtrOutput)
}

// VolumeTypeV3ArrayInput is an input type that accepts VolumeTypeV3Array and VolumeTypeV3ArrayOutput values.
// You can construct a concrete instance of `VolumeTypeV3ArrayInput` via:
//
//          VolumeTypeV3Array{ VolumeTypeV3Args{...} }
type VolumeTypeV3ArrayInput interface {
	pulumi.Input

	ToVolumeTypeV3ArrayOutput() VolumeTypeV3ArrayOutput
	ToVolumeTypeV3ArrayOutputWithContext(context.Context) VolumeTypeV3ArrayOutput
}

type VolumeTypeV3Array []VolumeTypeV3Input

func (VolumeTypeV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeTypeV3)(nil)).Elem()
}

func (i VolumeTypeV3Array) ToVolumeTypeV3ArrayOutput() VolumeTypeV3ArrayOutput {
	return i.ToVolumeTypeV3ArrayOutputWithContext(context.Background())
}

func (i VolumeTypeV3Array) ToVolumeTypeV3ArrayOutputWithContext(ctx context.Context) VolumeTypeV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeTypeV3ArrayOutput)
}

// VolumeTypeV3MapInput is an input type that accepts VolumeTypeV3Map and VolumeTypeV3MapOutput values.
// You can construct a concrete instance of `VolumeTypeV3MapInput` via:
//
//          VolumeTypeV3Map{ "key": VolumeTypeV3Args{...} }
type VolumeTypeV3MapInput interface {
	pulumi.Input

	ToVolumeTypeV3MapOutput() VolumeTypeV3MapOutput
	ToVolumeTypeV3MapOutputWithContext(context.Context) VolumeTypeV3MapOutput
}

type VolumeTypeV3Map map[string]VolumeTypeV3Input

func (VolumeTypeV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeTypeV3)(nil)).Elem()
}

func (i VolumeTypeV3Map) ToVolumeTypeV3MapOutput() VolumeTypeV3MapOutput {
	return i.ToVolumeTypeV3MapOutputWithContext(context.Background())
}

func (i VolumeTypeV3Map) ToVolumeTypeV3MapOutputWithContext(ctx context.Context) VolumeTypeV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeTypeV3MapOutput)
}

type VolumeTypeV3Output struct{ *pulumi.OutputState }

func (VolumeTypeV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeTypeV3)(nil))
}

func (o VolumeTypeV3Output) ToVolumeTypeV3Output() VolumeTypeV3Output {
	return o
}

func (o VolumeTypeV3Output) ToVolumeTypeV3OutputWithContext(ctx context.Context) VolumeTypeV3Output {
	return o
}

func (o VolumeTypeV3Output) ToVolumeTypeV3PtrOutput() VolumeTypeV3PtrOutput {
	return o.ToVolumeTypeV3PtrOutputWithContext(context.Background())
}

func (o VolumeTypeV3Output) ToVolumeTypeV3PtrOutputWithContext(ctx context.Context) VolumeTypeV3PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeTypeV3) *VolumeTypeV3 {
		return &v
	}).(VolumeTypeV3PtrOutput)
}

type VolumeTypeV3PtrOutput struct{ *pulumi.OutputState }

func (VolumeTypeV3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeTypeV3)(nil))
}

func (o VolumeTypeV3PtrOutput) ToVolumeTypeV3PtrOutput() VolumeTypeV3PtrOutput {
	return o
}

func (o VolumeTypeV3PtrOutput) ToVolumeTypeV3PtrOutputWithContext(ctx context.Context) VolumeTypeV3PtrOutput {
	return o
}

func (o VolumeTypeV3PtrOutput) Elem() VolumeTypeV3Output {
	return o.ApplyT(func(v *VolumeTypeV3) VolumeTypeV3 {
		if v != nil {
			return *v
		}
		var ret VolumeTypeV3
		return ret
	}).(VolumeTypeV3Output)
}

type VolumeTypeV3ArrayOutput struct{ *pulumi.OutputState }

func (VolumeTypeV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeTypeV3)(nil))
}

func (o VolumeTypeV3ArrayOutput) ToVolumeTypeV3ArrayOutput() VolumeTypeV3ArrayOutput {
	return o
}

func (o VolumeTypeV3ArrayOutput) ToVolumeTypeV3ArrayOutputWithContext(ctx context.Context) VolumeTypeV3ArrayOutput {
	return o
}

func (o VolumeTypeV3ArrayOutput) Index(i pulumi.IntInput) VolumeTypeV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeTypeV3 {
		return vs[0].([]VolumeTypeV3)[vs[1].(int)]
	}).(VolumeTypeV3Output)
}

type VolumeTypeV3MapOutput struct{ *pulumi.OutputState }

func (VolumeTypeV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VolumeTypeV3)(nil))
}

func (o VolumeTypeV3MapOutput) ToVolumeTypeV3MapOutput() VolumeTypeV3MapOutput {
	return o
}

func (o VolumeTypeV3MapOutput) ToVolumeTypeV3MapOutputWithContext(ctx context.Context) VolumeTypeV3MapOutput {
	return o
}

func (o VolumeTypeV3MapOutput) MapIndex(k pulumi.StringInput) VolumeTypeV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VolumeTypeV3 {
		return vs[0].(map[string]VolumeTypeV3)[vs[1].(string)]
	}).(VolumeTypeV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypeV3Input)(nil)).Elem(), &VolumeTypeV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypeV3PtrInput)(nil)).Elem(), &VolumeTypeV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypeV3ArrayInput)(nil)).Elem(), VolumeTypeV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypeV3MapInput)(nil)).Elem(), VolumeTypeV3Map{})
	pulumi.RegisterOutputType(VolumeTypeV3Output{})
	pulumi.RegisterOutputType(VolumeTypeV3PtrOutput{})
	pulumi.RegisterOutputType(VolumeTypeV3ArrayOutput{})
	pulumi.RegisterOutputType(VolumeTypeV3MapOutput{})
}
