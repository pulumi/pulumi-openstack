// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a networking V2 trunk resource within OpenStack.
type Trunk struct {
	pulumi.CustomResourceState

	// Administrative up/down status for the trunk
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing trunk.
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// The collection of tags assigned on the trunk, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// Human-readable description of the trunk. Changing this
	// updates the name of the existing trunk.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique name for the trunk. Changing this
	// updates the `name` of an existing trunk.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the port to be used as the parent port of the
	// trunk. This is the port that should be used as the compute instance network
	// port. Changing this creates a new trunk.
	PortId pulumi.StringOutput `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a trunk. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// trunk.
	Region pulumi.StringOutput `pulumi:"region"`
	// The set of ports that will be made subports of the trunk.
	// The structure of each subport is described below.
	SubPorts TrunkSubPortArrayOutput `pulumi:"subPorts"`
	// A set of string tags for the port.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The owner of the Trunk. Required if admin wants
	// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewTrunk registers a new resource with the given unique name, arguments, and options.
func NewTrunk(ctx *pulumi.Context,
	name string, args *TrunkArgs, opts ...pulumi.ResourceOption) (*Trunk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortId == nil {
		return nil, errors.New("invalid value for required argument 'PortId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Trunk
	err := ctx.RegisterResource("openstack:networking/trunk:Trunk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrunk gets an existing Trunk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrunk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrunkState, opts ...pulumi.ResourceOption) (*Trunk, error) {
	var resource Trunk
	err := ctx.ReadResource("openstack:networking/trunk:Trunk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trunk resources.
type trunkState struct {
	// Administrative up/down status for the trunk
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing trunk.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The collection of tags assigned on the trunk, which have been
	// explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// Human-readable description of the trunk. Changing this
	// updates the name of the existing trunk.
	Description *string `pulumi:"description"`
	// A unique name for the trunk. Changing this
	// updates the `name` of an existing trunk.
	Name *string `pulumi:"name"`
	// The ID of the port to be used as the parent port of the
	// trunk. This is the port that should be used as the compute instance network
	// port. Changing this creates a new trunk.
	PortId *string `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a trunk. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// trunk.
	Region *string `pulumi:"region"`
	// The set of ports that will be made subports of the trunk.
	// The structure of each subport is described below.
	SubPorts []TrunkSubPort `pulumi:"subPorts"`
	// A set of string tags for the port.
	Tags []string `pulumi:"tags"`
	// The owner of the Trunk. Required if admin wants
	// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
	TenantId *string `pulumi:"tenantId"`
}

type TrunkState struct {
	// Administrative up/down status for the trunk
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing trunk.
	AdminStateUp pulumi.BoolPtrInput
	// The collection of tags assigned on the trunk, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// Human-readable description of the trunk. Changing this
	// updates the name of the existing trunk.
	Description pulumi.StringPtrInput
	// A unique name for the trunk. Changing this
	// updates the `name` of an existing trunk.
	Name pulumi.StringPtrInput
	// The ID of the port to be used as the parent port of the
	// trunk. This is the port that should be used as the compute instance network
	// port. Changing this creates a new trunk.
	PortId pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a trunk. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// trunk.
	Region pulumi.StringPtrInput
	// The set of ports that will be made subports of the trunk.
	// The structure of each subport is described below.
	SubPorts TrunkSubPortArrayInput
	// A set of string tags for the port.
	Tags pulumi.StringArrayInput
	// The owner of the Trunk. Required if admin wants
	// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
	TenantId pulumi.StringPtrInput
}

func (TrunkState) ElementType() reflect.Type {
	return reflect.TypeOf((*trunkState)(nil)).Elem()
}

type trunkArgs struct {
	// Administrative up/down status for the trunk
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing trunk.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description of the trunk. Changing this
	// updates the name of the existing trunk.
	Description *string `pulumi:"description"`
	// A unique name for the trunk. Changing this
	// updates the `name` of an existing trunk.
	Name *string `pulumi:"name"`
	// The ID of the port to be used as the parent port of the
	// trunk. This is the port that should be used as the compute instance network
	// port. Changing this creates a new trunk.
	PortId string `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a trunk. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// trunk.
	Region *string `pulumi:"region"`
	// The set of ports that will be made subports of the trunk.
	// The structure of each subport is described below.
	SubPorts []TrunkSubPort `pulumi:"subPorts"`
	// A set of string tags for the port.
	Tags []string `pulumi:"tags"`
	// The owner of the Trunk. Required if admin wants
	// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a Trunk resource.
type TrunkArgs struct {
	// Administrative up/down status for the trunk
	// (must be "true" or "false" if provided). Changing this updates the
	// `adminStateUp` of an existing trunk.
	AdminStateUp pulumi.BoolPtrInput
	// Human-readable description of the trunk. Changing this
	// updates the name of the existing trunk.
	Description pulumi.StringPtrInput
	// A unique name for the trunk. Changing this
	// updates the `name` of an existing trunk.
	Name pulumi.StringPtrInput
	// The ID of the port to be used as the parent port of the
	// trunk. This is the port that should be used as the compute instance network
	// port. Changing this creates a new trunk.
	PortId pulumi.StringInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a trunk. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// trunk.
	Region pulumi.StringPtrInput
	// The set of ports that will be made subports of the trunk.
	// The structure of each subport is described below.
	SubPorts TrunkSubPortArrayInput
	// A set of string tags for the port.
	Tags pulumi.StringArrayInput
	// The owner of the Trunk. Required if admin wants
	// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
	TenantId pulumi.StringPtrInput
}

func (TrunkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trunkArgs)(nil)).Elem()
}

type TrunkInput interface {
	pulumi.Input

	ToTrunkOutput() TrunkOutput
	ToTrunkOutputWithContext(ctx context.Context) TrunkOutput
}

func (*Trunk) ElementType() reflect.Type {
	return reflect.TypeOf((**Trunk)(nil)).Elem()
}

func (i *Trunk) ToTrunkOutput() TrunkOutput {
	return i.ToTrunkOutputWithContext(context.Background())
}

func (i *Trunk) ToTrunkOutputWithContext(ctx context.Context) TrunkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrunkOutput)
}

func (i *Trunk) ToOutput(ctx context.Context) pulumix.Output[*Trunk] {
	return pulumix.Output[*Trunk]{
		OutputState: i.ToTrunkOutputWithContext(ctx).OutputState,
	}
}

// TrunkArrayInput is an input type that accepts TrunkArray and TrunkArrayOutput values.
// You can construct a concrete instance of `TrunkArrayInput` via:
//
//	TrunkArray{ TrunkArgs{...} }
type TrunkArrayInput interface {
	pulumi.Input

	ToTrunkArrayOutput() TrunkArrayOutput
	ToTrunkArrayOutputWithContext(context.Context) TrunkArrayOutput
}

type TrunkArray []TrunkInput

func (TrunkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trunk)(nil)).Elem()
}

func (i TrunkArray) ToTrunkArrayOutput() TrunkArrayOutput {
	return i.ToTrunkArrayOutputWithContext(context.Background())
}

func (i TrunkArray) ToTrunkArrayOutputWithContext(ctx context.Context) TrunkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrunkArrayOutput)
}

func (i TrunkArray) ToOutput(ctx context.Context) pulumix.Output[[]*Trunk] {
	return pulumix.Output[[]*Trunk]{
		OutputState: i.ToTrunkArrayOutputWithContext(ctx).OutputState,
	}
}

// TrunkMapInput is an input type that accepts TrunkMap and TrunkMapOutput values.
// You can construct a concrete instance of `TrunkMapInput` via:
//
//	TrunkMap{ "key": TrunkArgs{...} }
type TrunkMapInput interface {
	pulumi.Input

	ToTrunkMapOutput() TrunkMapOutput
	ToTrunkMapOutputWithContext(context.Context) TrunkMapOutput
}

type TrunkMap map[string]TrunkInput

func (TrunkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trunk)(nil)).Elem()
}

func (i TrunkMap) ToTrunkMapOutput() TrunkMapOutput {
	return i.ToTrunkMapOutputWithContext(context.Background())
}

func (i TrunkMap) ToTrunkMapOutputWithContext(ctx context.Context) TrunkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrunkMapOutput)
}

func (i TrunkMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Trunk] {
	return pulumix.Output[map[string]*Trunk]{
		OutputState: i.ToTrunkMapOutputWithContext(ctx).OutputState,
	}
}

type TrunkOutput struct{ *pulumi.OutputState }

func (TrunkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trunk)(nil)).Elem()
}

func (o TrunkOutput) ToTrunkOutput() TrunkOutput {
	return o
}

func (o TrunkOutput) ToTrunkOutputWithContext(ctx context.Context) TrunkOutput {
	return o
}

func (o TrunkOutput) ToOutput(ctx context.Context) pulumix.Output[*Trunk] {
	return pulumix.Output[*Trunk]{
		OutputState: o.OutputState,
	}
}

// Administrative up/down status for the trunk
// (must be "true" or "false" if provided). Changing this updates the
// `adminStateUp` of an existing trunk.
func (o TrunkOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Trunk) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// The collection of tags assigned on the trunk, which have been
// explicitly and implicitly added.
func (o TrunkOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Trunk) pulumi.StringArrayOutput { return v.AllTags }).(pulumi.StringArrayOutput)
}

// Human-readable description of the trunk. Changing this
// updates the name of the existing trunk.
func (o TrunkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trunk) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique name for the trunk. Changing this
// updates the `name` of an existing trunk.
func (o TrunkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Trunk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the port to be used as the parent port of the
// trunk. This is the port that should be used as the compute instance network
// port. Changing this creates a new trunk.
func (o TrunkOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *Trunk) pulumi.StringOutput { return v.PortId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to create a trunk. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// trunk.
func (o TrunkOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Trunk) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The set of ports that will be made subports of the trunk.
// The structure of each subport is described below.
func (o TrunkOutput) SubPorts() TrunkSubPortArrayOutput {
	return o.ApplyT(func(v *Trunk) TrunkSubPortArrayOutput { return v.SubPorts }).(TrunkSubPortArrayOutput)
}

// A set of string tags for the port.
func (o TrunkOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Trunk) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The owner of the Trunk. Required if admin wants
// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
func (o TrunkOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Trunk) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type TrunkArrayOutput struct{ *pulumi.OutputState }

func (TrunkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trunk)(nil)).Elem()
}

func (o TrunkArrayOutput) ToTrunkArrayOutput() TrunkArrayOutput {
	return o
}

func (o TrunkArrayOutput) ToTrunkArrayOutputWithContext(ctx context.Context) TrunkArrayOutput {
	return o
}

func (o TrunkArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Trunk] {
	return pulumix.Output[[]*Trunk]{
		OutputState: o.OutputState,
	}
}

func (o TrunkArrayOutput) Index(i pulumi.IntInput) TrunkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trunk {
		return vs[0].([]*Trunk)[vs[1].(int)]
	}).(TrunkOutput)
}

type TrunkMapOutput struct{ *pulumi.OutputState }

func (TrunkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trunk)(nil)).Elem()
}

func (o TrunkMapOutput) ToTrunkMapOutput() TrunkMapOutput {
	return o
}

func (o TrunkMapOutput) ToTrunkMapOutputWithContext(ctx context.Context) TrunkMapOutput {
	return o
}

func (o TrunkMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Trunk] {
	return pulumix.Output[map[string]*Trunk]{
		OutputState: o.OutputState,
	}
}

func (o TrunkMapOutput) MapIndex(k pulumi.StringInput) TrunkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trunk {
		return vs[0].(map[string]*Trunk)[vs[1].(string)]
	}).(TrunkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrunkInput)(nil)).Elem(), &Trunk{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrunkArrayInput)(nil)).Elem(), TrunkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrunkMapInput)(nil)).Elem(), TrunkMap{})
	pulumi.RegisterOutputType(TrunkOutput{})
	pulumi.RegisterOutputType(TrunkArrayOutput{})
	pulumi.RegisterOutputType(TrunkMapOutput{})
}
