// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V3 block storage Qos Association resource within OpenStack.
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
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/blockstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			qos, err := blockstorage.NewQosV3(ctx, "qos", &blockstorage.QosV3Args{
//				Name:     pulumi.String("%s"),
//				Consumer: pulumi.String("front-end"),
//				Specs: pulumi.StringMap{
//					"read_iops_sec": pulumi.String("20000"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			volumeType, err := blockstorage.NewVolumeTypeV3(ctx, "volume_type", &blockstorage.VolumeTypeV3Args{
//				Name: pulumi.String("%s"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = blockstorage.NewQosAssociationV3(ctx, "qos_association", &blockstorage.QosAssociationV3Args{
//				QosId:        qos.ID(),
//				VolumeTypeId: volumeType.ID(),
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
// Qos association can be imported using the `qos_id/volume_type_id`, e.g.
//
// ```sh
// $ pulumi import openstack:blockstorage/qosAssociationV3:QosAssociationV3 qos_association 941793f0-0a34-4bc4-b72e-a6326ae58283/ea257959-eeb1-4c10-8d33-26f0409a755d
// ```
type QosAssociationV3 struct {
	pulumi.CustomResourceState

	// ID of the qos to associate. Changing this creates
	// a new qos association.
	QosId pulumi.StringOutput `pulumi:"qosId"`
	// The region in which to create the qos association.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new qos association.
	Region pulumi.StringOutput `pulumi:"region"`
	// ID of the volumeType to associate.
	// Changing this creates a new qos association.
	VolumeTypeId pulumi.StringOutput `pulumi:"volumeTypeId"`
}

// NewQosAssociationV3 registers a new resource with the given unique name, arguments, and options.
func NewQosAssociationV3(ctx *pulumi.Context,
	name string, args *QosAssociationV3Args, opts ...pulumi.ResourceOption) (*QosAssociationV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.QosId == nil {
		return nil, errors.New("invalid value for required argument 'QosId'")
	}
	if args.VolumeTypeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeTypeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QosAssociationV3
	err := ctx.RegisterResource("openstack:blockstorage/qosAssociationV3:QosAssociationV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQosAssociationV3 gets an existing QosAssociationV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQosAssociationV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QosAssociationV3State, opts ...pulumi.ResourceOption) (*QosAssociationV3, error) {
	var resource QosAssociationV3
	err := ctx.ReadResource("openstack:blockstorage/qosAssociationV3:QosAssociationV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QosAssociationV3 resources.
type qosAssociationV3State struct {
	// ID of the qos to associate. Changing this creates
	// a new qos association.
	QosId *string `pulumi:"qosId"`
	// The region in which to create the qos association.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new qos association.
	Region *string `pulumi:"region"`
	// ID of the volumeType to associate.
	// Changing this creates a new qos association.
	VolumeTypeId *string `pulumi:"volumeTypeId"`
}

type QosAssociationV3State struct {
	// ID of the qos to associate. Changing this creates
	// a new qos association.
	QosId pulumi.StringPtrInput
	// The region in which to create the qos association.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new qos association.
	Region pulumi.StringPtrInput
	// ID of the volumeType to associate.
	// Changing this creates a new qos association.
	VolumeTypeId pulumi.StringPtrInput
}

func (QosAssociationV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*qosAssociationV3State)(nil)).Elem()
}

type qosAssociationV3Args struct {
	// ID of the qos to associate. Changing this creates
	// a new qos association.
	QosId string `pulumi:"qosId"`
	// The region in which to create the qos association.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new qos association.
	Region *string `pulumi:"region"`
	// ID of the volumeType to associate.
	// Changing this creates a new qos association.
	VolumeTypeId string `pulumi:"volumeTypeId"`
}

// The set of arguments for constructing a QosAssociationV3 resource.
type QosAssociationV3Args struct {
	// ID of the qos to associate. Changing this creates
	// a new qos association.
	QosId pulumi.StringInput
	// The region in which to create the qos association.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new qos association.
	Region pulumi.StringPtrInput
	// ID of the volumeType to associate.
	// Changing this creates a new qos association.
	VolumeTypeId pulumi.StringInput
}

func (QosAssociationV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*qosAssociationV3Args)(nil)).Elem()
}

type QosAssociationV3Input interface {
	pulumi.Input

	ToQosAssociationV3Output() QosAssociationV3Output
	ToQosAssociationV3OutputWithContext(ctx context.Context) QosAssociationV3Output
}

func (*QosAssociationV3) ElementType() reflect.Type {
	return reflect.TypeOf((**QosAssociationV3)(nil)).Elem()
}

func (i *QosAssociationV3) ToQosAssociationV3Output() QosAssociationV3Output {
	return i.ToQosAssociationV3OutputWithContext(context.Background())
}

func (i *QosAssociationV3) ToQosAssociationV3OutputWithContext(ctx context.Context) QosAssociationV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(QosAssociationV3Output)
}

// QosAssociationV3ArrayInput is an input type that accepts QosAssociationV3Array and QosAssociationV3ArrayOutput values.
// You can construct a concrete instance of `QosAssociationV3ArrayInput` via:
//
//	QosAssociationV3Array{ QosAssociationV3Args{...} }
type QosAssociationV3ArrayInput interface {
	pulumi.Input

	ToQosAssociationV3ArrayOutput() QosAssociationV3ArrayOutput
	ToQosAssociationV3ArrayOutputWithContext(context.Context) QosAssociationV3ArrayOutput
}

type QosAssociationV3Array []QosAssociationV3Input

func (QosAssociationV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QosAssociationV3)(nil)).Elem()
}

func (i QosAssociationV3Array) ToQosAssociationV3ArrayOutput() QosAssociationV3ArrayOutput {
	return i.ToQosAssociationV3ArrayOutputWithContext(context.Background())
}

func (i QosAssociationV3Array) ToQosAssociationV3ArrayOutputWithContext(ctx context.Context) QosAssociationV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosAssociationV3ArrayOutput)
}

// QosAssociationV3MapInput is an input type that accepts QosAssociationV3Map and QosAssociationV3MapOutput values.
// You can construct a concrete instance of `QosAssociationV3MapInput` via:
//
//	QosAssociationV3Map{ "key": QosAssociationV3Args{...} }
type QosAssociationV3MapInput interface {
	pulumi.Input

	ToQosAssociationV3MapOutput() QosAssociationV3MapOutput
	ToQosAssociationV3MapOutputWithContext(context.Context) QosAssociationV3MapOutput
}

type QosAssociationV3Map map[string]QosAssociationV3Input

func (QosAssociationV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QosAssociationV3)(nil)).Elem()
}

func (i QosAssociationV3Map) ToQosAssociationV3MapOutput() QosAssociationV3MapOutput {
	return i.ToQosAssociationV3MapOutputWithContext(context.Background())
}

func (i QosAssociationV3Map) ToQosAssociationV3MapOutputWithContext(ctx context.Context) QosAssociationV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosAssociationV3MapOutput)
}

type QosAssociationV3Output struct{ *pulumi.OutputState }

func (QosAssociationV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**QosAssociationV3)(nil)).Elem()
}

func (o QosAssociationV3Output) ToQosAssociationV3Output() QosAssociationV3Output {
	return o
}

func (o QosAssociationV3Output) ToQosAssociationV3OutputWithContext(ctx context.Context) QosAssociationV3Output {
	return o
}

// ID of the qos to associate. Changing this creates
// a new qos association.
func (o QosAssociationV3Output) QosId() pulumi.StringOutput {
	return o.ApplyT(func(v *QosAssociationV3) pulumi.StringOutput { return v.QosId }).(pulumi.StringOutput)
}

// The region in which to create the qos association.
// If omitted, the `region` argument of the provider is used. Changing
// this creates a new qos association.
func (o QosAssociationV3Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *QosAssociationV3) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// ID of the volumeType to associate.
// Changing this creates a new qos association.
func (o QosAssociationV3Output) VolumeTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *QosAssociationV3) pulumi.StringOutput { return v.VolumeTypeId }).(pulumi.StringOutput)
}

type QosAssociationV3ArrayOutput struct{ *pulumi.OutputState }

func (QosAssociationV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QosAssociationV3)(nil)).Elem()
}

func (o QosAssociationV3ArrayOutput) ToQosAssociationV3ArrayOutput() QosAssociationV3ArrayOutput {
	return o
}

func (o QosAssociationV3ArrayOutput) ToQosAssociationV3ArrayOutputWithContext(ctx context.Context) QosAssociationV3ArrayOutput {
	return o
}

func (o QosAssociationV3ArrayOutput) Index(i pulumi.IntInput) QosAssociationV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QosAssociationV3 {
		return vs[0].([]*QosAssociationV3)[vs[1].(int)]
	}).(QosAssociationV3Output)
}

type QosAssociationV3MapOutput struct{ *pulumi.OutputState }

func (QosAssociationV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QosAssociationV3)(nil)).Elem()
}

func (o QosAssociationV3MapOutput) ToQosAssociationV3MapOutput() QosAssociationV3MapOutput {
	return o
}

func (o QosAssociationV3MapOutput) ToQosAssociationV3MapOutputWithContext(ctx context.Context) QosAssociationV3MapOutput {
	return o
}

func (o QosAssociationV3MapOutput) MapIndex(k pulumi.StringInput) QosAssociationV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QosAssociationV3 {
		return vs[0].(map[string]*QosAssociationV3)[vs[1].(string)]
	}).(QosAssociationV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QosAssociationV3Input)(nil)).Elem(), &QosAssociationV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosAssociationV3ArrayInput)(nil)).Elem(), QosAssociationV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosAssociationV3MapInput)(nil)).Elem(), QosAssociationV3Map{})
	pulumi.RegisterOutputType(QosAssociationV3Output{})
	pulumi.RegisterOutputType(QosAssociationV3ArrayOutput{})
	pulumi.RegisterOutputType(QosAssociationV3MapOutput{})
}
