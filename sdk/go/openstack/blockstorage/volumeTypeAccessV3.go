// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V3 block storage volume type access resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project1, err := identity.NewProject(ctx, "project_1", &identity.ProjectArgs{
//				Name: pulumi.String("project_1"),
//			})
//			if err != nil {
//				return err
//			}
//			volumeType1, err := blockstorage.NewVolumeTypeV3(ctx, "volume_type_1", &blockstorage.VolumeTypeV3Args{
//				Name:     pulumi.String("volume_type_1"),
//				IsPublic: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = blockstorage.NewVolumeTypeAccessV3(ctx, "volume_type_access", &blockstorage.VolumeTypeAccessV3Args{
//				ProjectId:    project1.ID(),
//				VolumeTypeId: volumeType1.ID(),
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
// Volume types access can be imported using the `volume_type_id/project_id`, e.g.
//
// ```sh
// $ pulumi import openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3 volume_type_access 941793f0-0a34-4bc4-b72e-a6326ae58283/ed498e81f0cc448bae0ad4f8f21bf67f
// ```
type VolumeTypeAccessV3 struct {
	pulumi.CustomResourceState

	// ID of the project to give access to. Changing this
	// creates a new resource.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringOutput `pulumi:"region"`
	// ID of the volume type to give access to. Changing
	// this creates a new resource.
	VolumeTypeId pulumi.StringOutput `pulumi:"volumeTypeId"`
}

// NewVolumeTypeAccessV3 registers a new resource with the given unique name, arguments, and options.
func NewVolumeTypeAccessV3(ctx *pulumi.Context,
	name string, args *VolumeTypeAccessV3Args, opts ...pulumi.ResourceOption) (*VolumeTypeAccessV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.VolumeTypeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeTypeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VolumeTypeAccessV3
	err := ctx.RegisterResource("openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeTypeAccessV3 gets an existing VolumeTypeAccessV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeTypeAccessV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeTypeAccessV3State, opts ...pulumi.ResourceOption) (*VolumeTypeAccessV3, error) {
	var resource VolumeTypeAccessV3
	err := ctx.ReadResource("openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeTypeAccessV3 resources.
type volumeTypeAccessV3State struct {
	// ID of the project to give access to. Changing this
	// creates a new resource.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `pulumi:"region"`
	// ID of the volume type to give access to. Changing
	// this creates a new resource.
	VolumeTypeId *string `pulumi:"volumeTypeId"`
}

type VolumeTypeAccessV3State struct {
	// ID of the project to give access to. Changing this
	// creates a new resource.
	ProjectId pulumi.StringPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringPtrInput
	// ID of the volume type to give access to. Changing
	// this creates a new resource.
	VolumeTypeId pulumi.StringPtrInput
}

func (VolumeTypeAccessV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeTypeAccessV3State)(nil)).Elem()
}

type volumeTypeAccessV3Args struct {
	// ID of the project to give access to. Changing this
	// creates a new resource.
	ProjectId string `pulumi:"projectId"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `pulumi:"region"`
	// ID of the volume type to give access to. Changing
	// this creates a new resource.
	VolumeTypeId string `pulumi:"volumeTypeId"`
}

// The set of arguments for constructing a VolumeTypeAccessV3 resource.
type VolumeTypeAccessV3Args struct {
	// ID of the project to give access to. Changing this
	// creates a new resource.
	ProjectId pulumi.StringInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringPtrInput
	// ID of the volume type to give access to. Changing
	// this creates a new resource.
	VolumeTypeId pulumi.StringInput
}

func (VolumeTypeAccessV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeTypeAccessV3Args)(nil)).Elem()
}

type VolumeTypeAccessV3Input interface {
	pulumi.Input

	ToVolumeTypeAccessV3Output() VolumeTypeAccessV3Output
	ToVolumeTypeAccessV3OutputWithContext(ctx context.Context) VolumeTypeAccessV3Output
}

func (*VolumeTypeAccessV3) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeTypeAccessV3)(nil)).Elem()
}

func (i *VolumeTypeAccessV3) ToVolumeTypeAccessV3Output() VolumeTypeAccessV3Output {
	return i.ToVolumeTypeAccessV3OutputWithContext(context.Background())
}

func (i *VolumeTypeAccessV3) ToVolumeTypeAccessV3OutputWithContext(ctx context.Context) VolumeTypeAccessV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeTypeAccessV3Output)
}

// VolumeTypeAccessV3ArrayInput is an input type that accepts VolumeTypeAccessV3Array and VolumeTypeAccessV3ArrayOutput values.
// You can construct a concrete instance of `VolumeTypeAccessV3ArrayInput` via:
//
//	VolumeTypeAccessV3Array{ VolumeTypeAccessV3Args{...} }
type VolumeTypeAccessV3ArrayInput interface {
	pulumi.Input

	ToVolumeTypeAccessV3ArrayOutput() VolumeTypeAccessV3ArrayOutput
	ToVolumeTypeAccessV3ArrayOutputWithContext(context.Context) VolumeTypeAccessV3ArrayOutput
}

type VolumeTypeAccessV3Array []VolumeTypeAccessV3Input

func (VolumeTypeAccessV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeTypeAccessV3)(nil)).Elem()
}

func (i VolumeTypeAccessV3Array) ToVolumeTypeAccessV3ArrayOutput() VolumeTypeAccessV3ArrayOutput {
	return i.ToVolumeTypeAccessV3ArrayOutputWithContext(context.Background())
}

func (i VolumeTypeAccessV3Array) ToVolumeTypeAccessV3ArrayOutputWithContext(ctx context.Context) VolumeTypeAccessV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeTypeAccessV3ArrayOutput)
}

// VolumeTypeAccessV3MapInput is an input type that accepts VolumeTypeAccessV3Map and VolumeTypeAccessV3MapOutput values.
// You can construct a concrete instance of `VolumeTypeAccessV3MapInput` via:
//
//	VolumeTypeAccessV3Map{ "key": VolumeTypeAccessV3Args{...} }
type VolumeTypeAccessV3MapInput interface {
	pulumi.Input

	ToVolumeTypeAccessV3MapOutput() VolumeTypeAccessV3MapOutput
	ToVolumeTypeAccessV3MapOutputWithContext(context.Context) VolumeTypeAccessV3MapOutput
}

type VolumeTypeAccessV3Map map[string]VolumeTypeAccessV3Input

func (VolumeTypeAccessV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeTypeAccessV3)(nil)).Elem()
}

func (i VolumeTypeAccessV3Map) ToVolumeTypeAccessV3MapOutput() VolumeTypeAccessV3MapOutput {
	return i.ToVolumeTypeAccessV3MapOutputWithContext(context.Background())
}

func (i VolumeTypeAccessV3Map) ToVolumeTypeAccessV3MapOutputWithContext(ctx context.Context) VolumeTypeAccessV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeTypeAccessV3MapOutput)
}

type VolumeTypeAccessV3Output struct{ *pulumi.OutputState }

func (VolumeTypeAccessV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeTypeAccessV3)(nil)).Elem()
}

func (o VolumeTypeAccessV3Output) ToVolumeTypeAccessV3Output() VolumeTypeAccessV3Output {
	return o
}

func (o VolumeTypeAccessV3Output) ToVolumeTypeAccessV3OutputWithContext(ctx context.Context) VolumeTypeAccessV3Output {
	return o
}

// ID of the project to give access to. Changing this
// creates a new resource.
func (o VolumeTypeAccessV3Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeTypeAccessV3) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to create the volume. If
// omitted, the `region` argument of the provider is used. Changing this
// creates a new quotaset.
func (o VolumeTypeAccessV3Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeTypeAccessV3) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// ID of the volume type to give access to. Changing
// this creates a new resource.
func (o VolumeTypeAccessV3Output) VolumeTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeTypeAccessV3) pulumi.StringOutput { return v.VolumeTypeId }).(pulumi.StringOutput)
}

type VolumeTypeAccessV3ArrayOutput struct{ *pulumi.OutputState }

func (VolumeTypeAccessV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VolumeTypeAccessV3)(nil)).Elem()
}

func (o VolumeTypeAccessV3ArrayOutput) ToVolumeTypeAccessV3ArrayOutput() VolumeTypeAccessV3ArrayOutput {
	return o
}

func (o VolumeTypeAccessV3ArrayOutput) ToVolumeTypeAccessV3ArrayOutputWithContext(ctx context.Context) VolumeTypeAccessV3ArrayOutput {
	return o
}

func (o VolumeTypeAccessV3ArrayOutput) Index(i pulumi.IntInput) VolumeTypeAccessV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VolumeTypeAccessV3 {
		return vs[0].([]*VolumeTypeAccessV3)[vs[1].(int)]
	}).(VolumeTypeAccessV3Output)
}

type VolumeTypeAccessV3MapOutput struct{ *pulumi.OutputState }

func (VolumeTypeAccessV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VolumeTypeAccessV3)(nil)).Elem()
}

func (o VolumeTypeAccessV3MapOutput) ToVolumeTypeAccessV3MapOutput() VolumeTypeAccessV3MapOutput {
	return o
}

func (o VolumeTypeAccessV3MapOutput) ToVolumeTypeAccessV3MapOutputWithContext(ctx context.Context) VolumeTypeAccessV3MapOutput {
	return o
}

func (o VolumeTypeAccessV3MapOutput) MapIndex(k pulumi.StringInput) VolumeTypeAccessV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VolumeTypeAccessV3 {
		return vs[0].(map[string]*VolumeTypeAccessV3)[vs[1].(string)]
	}).(VolumeTypeAccessV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypeAccessV3Input)(nil)).Elem(), &VolumeTypeAccessV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypeAccessV3ArrayInput)(nil)).Elem(), VolumeTypeAccessV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeTypeAccessV3MapInput)(nil)).Elem(), VolumeTypeAccessV3Map{})
	pulumi.RegisterOutputType(VolumeTypeAccessV3Output{})
	pulumi.RegisterOutputType(VolumeTypeAccessV3ArrayOutput{})
	pulumi.RegisterOutputType(VolumeTypeAccessV3MapOutput{})
}
