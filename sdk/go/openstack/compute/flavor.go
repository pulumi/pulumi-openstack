// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a V2 flavor resource within OpenStack.
//
// ## Import
//
// Flavors can be imported using the `ID`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:compute/flavor:Flavor my-flavor 4142e64b-1b35-44a0-9b1e-5affc7af1106
//
// ```
type Flavor struct {
	pulumi.CustomResourceState

	// The description of the flavor. Changing this
	// updates the description of the flavor. Requires microversion >= 2.55.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The amount of disk space in GiB to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk pulumi.IntOutput `pulumi:"disk"`
	// The amount of ephemeral in GiB. If unspecified,
	// the default is 0. Changing this creates a new flavor.
	Ephemeral pulumi.IntPtrOutput `pulumi:"ephemeral"`
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs pulumi.MapOutput `pulumi:"extraSpecs"`
	// Unique ID (integer or UUID) of flavor to create. Changing
	// this creates a new flavor.
	FlavorId pulumi.StringOutput `pulumi:"flavorId"`
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram pulumi.IntOutput `pulumi:"ram"`
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region pulumi.StringOutput `pulumi:"region"`
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor pulumi.Float64PtrOutput `pulumi:"rxTxFactor"`
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap pulumi.IntPtrOutput `pulumi:"swap"`
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus pulumi.IntOutput `pulumi:"vcpus"`
}

// NewFlavor registers a new resource with the given unique name, arguments, and options.
func NewFlavor(ctx *pulumi.Context,
	name string, args *FlavorArgs, opts ...pulumi.ResourceOption) (*Flavor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disk == nil {
		return nil, errors.New("invalid value for required argument 'Disk'")
	}
	if args.Ram == nil {
		return nil, errors.New("invalid value for required argument 'Ram'")
	}
	if args.Vcpus == nil {
		return nil, errors.New("invalid value for required argument 'Vcpus'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Flavor
	err := ctx.RegisterResource("openstack:compute/flavor:Flavor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlavor gets an existing Flavor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlavor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlavorState, opts ...pulumi.ResourceOption) (*Flavor, error) {
	var resource Flavor
	err := ctx.ReadResource("openstack:compute/flavor:Flavor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flavor resources.
type flavorState struct {
	// The description of the flavor. Changing this
	// updates the description of the flavor. Requires microversion >= 2.55.
	Description *string `pulumi:"description"`
	// The amount of disk space in GiB to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk *int `pulumi:"disk"`
	// The amount of ephemeral in GiB. If unspecified,
	// the default is 0. Changing this creates a new flavor.
	Ephemeral *int `pulumi:"ephemeral"`
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs map[string]interface{} `pulumi:"extraSpecs"`
	// Unique ID (integer or UUID) of flavor to create. Changing
	// this creates a new flavor.
	FlavorId *string `pulumi:"flavorId"`
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic *bool `pulumi:"isPublic"`
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name *string `pulumi:"name"`
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram *int `pulumi:"ram"`
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region *string `pulumi:"region"`
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor *float64 `pulumi:"rxTxFactor"`
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap *int `pulumi:"swap"`
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus *int `pulumi:"vcpus"`
}

type FlavorState struct {
	// The description of the flavor. Changing this
	// updates the description of the flavor. Requires microversion >= 2.55.
	Description pulumi.StringPtrInput
	// The amount of disk space in GiB to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk pulumi.IntPtrInput
	// The amount of ephemeral in GiB. If unspecified,
	// the default is 0. Changing this creates a new flavor.
	Ephemeral pulumi.IntPtrInput
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs pulumi.MapInput
	// Unique ID (integer or UUID) of flavor to create. Changing
	// this creates a new flavor.
	FlavorId pulumi.StringPtrInput
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic pulumi.BoolPtrInput
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name pulumi.StringPtrInput
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram pulumi.IntPtrInput
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region pulumi.StringPtrInput
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor pulumi.Float64PtrInput
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap pulumi.IntPtrInput
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus pulumi.IntPtrInput
}

func (FlavorState) ElementType() reflect.Type {
	return reflect.TypeOf((*flavorState)(nil)).Elem()
}

type flavorArgs struct {
	// The description of the flavor. Changing this
	// updates the description of the flavor. Requires microversion >= 2.55.
	Description *string `pulumi:"description"`
	// The amount of disk space in GiB to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk int `pulumi:"disk"`
	// The amount of ephemeral in GiB. If unspecified,
	// the default is 0. Changing this creates a new flavor.
	Ephemeral *int `pulumi:"ephemeral"`
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs map[string]interface{} `pulumi:"extraSpecs"`
	// Unique ID (integer or UUID) of flavor to create. Changing
	// this creates a new flavor.
	FlavorId *string `pulumi:"flavorId"`
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic *bool `pulumi:"isPublic"`
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name *string `pulumi:"name"`
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram int `pulumi:"ram"`
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region *string `pulumi:"region"`
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor *float64 `pulumi:"rxTxFactor"`
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap *int `pulumi:"swap"`
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus int `pulumi:"vcpus"`
}

// The set of arguments for constructing a Flavor resource.
type FlavorArgs struct {
	// The description of the flavor. Changing this
	// updates the description of the flavor. Requires microversion >= 2.55.
	Description pulumi.StringPtrInput
	// The amount of disk space in GiB to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk pulumi.IntInput
	// The amount of ephemeral in GiB. If unspecified,
	// the default is 0. Changing this creates a new flavor.
	Ephemeral pulumi.IntPtrInput
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs pulumi.MapInput
	// Unique ID (integer or UUID) of flavor to create. Changing
	// this creates a new flavor.
	FlavorId pulumi.StringPtrInput
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic pulumi.BoolPtrInput
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name pulumi.StringPtrInput
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram pulumi.IntInput
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region pulumi.StringPtrInput
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor pulumi.Float64PtrInput
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap pulumi.IntPtrInput
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus pulumi.IntInput
}

func (FlavorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flavorArgs)(nil)).Elem()
}

type FlavorInput interface {
	pulumi.Input

	ToFlavorOutput() FlavorOutput
	ToFlavorOutputWithContext(ctx context.Context) FlavorOutput
}

func (*Flavor) ElementType() reflect.Type {
	return reflect.TypeOf((**Flavor)(nil)).Elem()
}

func (i *Flavor) ToFlavorOutput() FlavorOutput {
	return i.ToFlavorOutputWithContext(context.Background())
}

func (i *Flavor) ToFlavorOutputWithContext(ctx context.Context) FlavorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlavorOutput)
}

func (i *Flavor) ToOutput(ctx context.Context) pulumix.Output[*Flavor] {
	return pulumix.Output[*Flavor]{
		OutputState: i.ToFlavorOutputWithContext(ctx).OutputState,
	}
}

// FlavorArrayInput is an input type that accepts FlavorArray and FlavorArrayOutput values.
// You can construct a concrete instance of `FlavorArrayInput` via:
//
//	FlavorArray{ FlavorArgs{...} }
type FlavorArrayInput interface {
	pulumi.Input

	ToFlavorArrayOutput() FlavorArrayOutput
	ToFlavorArrayOutputWithContext(context.Context) FlavorArrayOutput
}

type FlavorArray []FlavorInput

func (FlavorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flavor)(nil)).Elem()
}

func (i FlavorArray) ToFlavorArrayOutput() FlavorArrayOutput {
	return i.ToFlavorArrayOutputWithContext(context.Background())
}

func (i FlavorArray) ToFlavorArrayOutputWithContext(ctx context.Context) FlavorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlavorArrayOutput)
}

func (i FlavorArray) ToOutput(ctx context.Context) pulumix.Output[[]*Flavor] {
	return pulumix.Output[[]*Flavor]{
		OutputState: i.ToFlavorArrayOutputWithContext(ctx).OutputState,
	}
}

// FlavorMapInput is an input type that accepts FlavorMap and FlavorMapOutput values.
// You can construct a concrete instance of `FlavorMapInput` via:
//
//	FlavorMap{ "key": FlavorArgs{...} }
type FlavorMapInput interface {
	pulumi.Input

	ToFlavorMapOutput() FlavorMapOutput
	ToFlavorMapOutputWithContext(context.Context) FlavorMapOutput
}

type FlavorMap map[string]FlavorInput

func (FlavorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flavor)(nil)).Elem()
}

func (i FlavorMap) ToFlavorMapOutput() FlavorMapOutput {
	return i.ToFlavorMapOutputWithContext(context.Background())
}

func (i FlavorMap) ToFlavorMapOutputWithContext(ctx context.Context) FlavorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlavorMapOutput)
}

func (i FlavorMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Flavor] {
	return pulumix.Output[map[string]*Flavor]{
		OutputState: i.ToFlavorMapOutputWithContext(ctx).OutputState,
	}
}

type FlavorOutput struct{ *pulumi.OutputState }

func (FlavorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Flavor)(nil)).Elem()
}

func (o FlavorOutput) ToFlavorOutput() FlavorOutput {
	return o
}

func (o FlavorOutput) ToFlavorOutputWithContext(ctx context.Context) FlavorOutput {
	return o
}

func (o FlavorOutput) ToOutput(ctx context.Context) pulumix.Output[*Flavor] {
	return pulumix.Output[*Flavor]{
		OutputState: o.OutputState,
	}
}

// The description of the flavor. Changing this
// updates the description of the flavor. Requires microversion >= 2.55.
func (o FlavorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Flavor) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The amount of disk space in GiB to use for the root
// (/) partition. Changing this creates a new flavor.
func (o FlavorOutput) Disk() pulumi.IntOutput {
	return o.ApplyT(func(v *Flavor) pulumi.IntOutput { return v.Disk }).(pulumi.IntOutput)
}

// The amount of ephemeral in GiB. If unspecified,
// the default is 0. Changing this creates a new flavor.
func (o FlavorOutput) Ephemeral() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Flavor) pulumi.IntPtrOutput { return v.Ephemeral }).(pulumi.IntPtrOutput)
}

// Key/Value pairs of metadata for the flavor.
func (o FlavorOutput) ExtraSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *Flavor) pulumi.MapOutput { return v.ExtraSpecs }).(pulumi.MapOutput)
}

// Unique ID (integer or UUID) of flavor to create. Changing
// this creates a new flavor.
func (o FlavorOutput) FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Flavor) pulumi.StringOutput { return v.FlavorId }).(pulumi.StringOutput)
}

// Whether the flavor is public. Changing this creates
// a new flavor.
func (o FlavorOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Flavor) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// A unique name for the flavor. Changing this creates a new
// flavor.
func (o FlavorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Flavor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The amount of RAM to use, in megabytes. Changing this
// creates a new flavor.
func (o FlavorOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v *Flavor) pulumi.IntOutput { return v.Ram }).(pulumi.IntOutput)
}

// The region in which to obtain the V2 Compute client.
// Flavors are associated with accounts, but a Compute client is needed to
// create one. If omitted, the `region` argument of the provider is used.
// Changing this creates a new flavor.
func (o FlavorOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Flavor) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// RX/TX bandwith factor. The default is 1. Changing
// this creates a new flavor.
func (o FlavorOutput) RxTxFactor() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Flavor) pulumi.Float64PtrOutput { return v.RxTxFactor }).(pulumi.Float64PtrOutput)
}

// The amount of disk space in megabytes to use. If
// unspecified, the default is 0. Changing this creates a new flavor.
func (o FlavorOutput) Swap() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Flavor) pulumi.IntPtrOutput { return v.Swap }).(pulumi.IntPtrOutput)
}

// The number of virtual CPUs to use. Changing this creates
// a new flavor.
func (o FlavorOutput) Vcpus() pulumi.IntOutput {
	return o.ApplyT(func(v *Flavor) pulumi.IntOutput { return v.Vcpus }).(pulumi.IntOutput)
}

type FlavorArrayOutput struct{ *pulumi.OutputState }

func (FlavorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Flavor)(nil)).Elem()
}

func (o FlavorArrayOutput) ToFlavorArrayOutput() FlavorArrayOutput {
	return o
}

func (o FlavorArrayOutput) ToFlavorArrayOutputWithContext(ctx context.Context) FlavorArrayOutput {
	return o
}

func (o FlavorArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Flavor] {
	return pulumix.Output[[]*Flavor]{
		OutputState: o.OutputState,
	}
}

func (o FlavorArrayOutput) Index(i pulumi.IntInput) FlavorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Flavor {
		return vs[0].([]*Flavor)[vs[1].(int)]
	}).(FlavorOutput)
}

type FlavorMapOutput struct{ *pulumi.OutputState }

func (FlavorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Flavor)(nil)).Elem()
}

func (o FlavorMapOutput) ToFlavorMapOutput() FlavorMapOutput {
	return o
}

func (o FlavorMapOutput) ToFlavorMapOutputWithContext(ctx context.Context) FlavorMapOutput {
	return o
}

func (o FlavorMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Flavor] {
	return pulumix.Output[map[string]*Flavor]{
		OutputState: o.OutputState,
	}
}

func (o FlavorMapOutput) MapIndex(k pulumi.StringInput) FlavorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Flavor {
		return vs[0].(map[string]*Flavor)[vs[1].(string)]
	}).(FlavorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlavorInput)(nil)).Elem(), &Flavor{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlavorArrayInput)(nil)).Elem(), FlavorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FlavorMapInput)(nil)).Elem(), FlavorMap{})
	pulumi.RegisterOutputType(FlavorOutput{})
	pulumi.RegisterOutputType(FlavorArrayOutput{})
	pulumi.RegisterOutputType(FlavorMapOutput{})
}
