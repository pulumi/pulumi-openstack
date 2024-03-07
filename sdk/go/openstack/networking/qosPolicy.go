// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Neutron QoS policy resource within OpenStack.
//
// ## Example Usage
//
// ### Create a QoS Policy
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.NewQosPolicy(ctx, "qosPolicy1", &networking.QosPolicyArgs{
//				Description: pulumi.String("bw_limit"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// QoS Policies can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:networking/qosPolicy:QosPolicy qos_policy_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae
// ```
type QosPolicy struct {
	pulumi.CustomResourceState

	// The collection of tags assigned on the QoS policy, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// The time at which QoS policy was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The human-readable description for the QoS policy.
	// Changing this updates the description of the existing QoS policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether the QoS policy is default
	// QoS policy or not. Changing this updates the default status of the existing
	// QoS policy.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// The name of the QoS policy. Changing this updates the name of
	// the existing QoS policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner of the QoS policy. Required if admin wants to
	// create a QoS policy for another project. Changing this creates a new QoS policy.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron Qos policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// QoS policy.
	Region pulumi.StringOutput `pulumi:"region"`
	// The revision number of the QoS policy.
	RevisionNumber pulumi.IntOutput `pulumi:"revisionNumber"`
	// Indicates whether this QoS policy is shared across
	// all projects. Changing this updates the shared status of the existing
	// QoS policy.
	Shared pulumi.BoolPtrOutput `pulumi:"shared"`
	// A set of string tags for the QoS policy.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The time at which QoS policy was created.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewQosPolicy registers a new resource with the given unique name, arguments, and options.
func NewQosPolicy(ctx *pulumi.Context,
	name string, args *QosPolicyArgs, opts ...pulumi.ResourceOption) (*QosPolicy, error) {
	if args == nil {
		args = &QosPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QosPolicy
	err := ctx.RegisterResource("openstack:networking/qosPolicy:QosPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQosPolicy gets an existing QosPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQosPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QosPolicyState, opts ...pulumi.ResourceOption) (*QosPolicy, error) {
	var resource QosPolicy
	err := ctx.ReadResource("openstack:networking/qosPolicy:QosPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QosPolicy resources.
type qosPolicyState struct {
	// The collection of tags assigned on the QoS policy, which have been
	// explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// The time at which QoS policy was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The human-readable description for the QoS policy.
	// Changing this updates the description of the existing QoS policy.
	Description *string `pulumi:"description"`
	// Indicates whether the QoS policy is default
	// QoS policy or not. Changing this updates the default status of the existing
	// QoS policy.
	IsDefault *bool `pulumi:"isDefault"`
	// The name of the QoS policy. Changing this updates the name of
	// the existing QoS policy.
	Name *string `pulumi:"name"`
	// The owner of the QoS policy. Required if admin wants to
	// create a QoS policy for another project. Changing this creates a new QoS policy.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron Qos policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// QoS policy.
	Region *string `pulumi:"region"`
	// The revision number of the QoS policy.
	RevisionNumber *int `pulumi:"revisionNumber"`
	// Indicates whether this QoS policy is shared across
	// all projects. Changing this updates the shared status of the existing
	// QoS policy.
	Shared *bool `pulumi:"shared"`
	// A set of string tags for the QoS policy.
	Tags []string `pulumi:"tags"`
	// The time at which QoS policy was created.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type QosPolicyState struct {
	// The collection of tags assigned on the QoS policy, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// The time at which QoS policy was created.
	CreatedAt pulumi.StringPtrInput
	// The human-readable description for the QoS policy.
	// Changing this updates the description of the existing QoS policy.
	Description pulumi.StringPtrInput
	// Indicates whether the QoS policy is default
	// QoS policy or not. Changing this updates the default status of the existing
	// QoS policy.
	IsDefault pulumi.BoolPtrInput
	// The name of the QoS policy. Changing this updates the name of
	// the existing QoS policy.
	Name pulumi.StringPtrInput
	// The owner of the QoS policy. Required if admin wants to
	// create a QoS policy for another project. Changing this creates a new QoS policy.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron Qos policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// QoS policy.
	Region pulumi.StringPtrInput
	// The revision number of the QoS policy.
	RevisionNumber pulumi.IntPtrInput
	// Indicates whether this QoS policy is shared across
	// all projects. Changing this updates the shared status of the existing
	// QoS policy.
	Shared pulumi.BoolPtrInput
	// A set of string tags for the QoS policy.
	Tags pulumi.StringArrayInput
	// The time at which QoS policy was created.
	UpdatedAt pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (QosPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*qosPolicyState)(nil)).Elem()
}

type qosPolicyArgs struct {
	// The human-readable description for the QoS policy.
	// Changing this updates the description of the existing QoS policy.
	Description *string `pulumi:"description"`
	// Indicates whether the QoS policy is default
	// QoS policy or not. Changing this updates the default status of the existing
	// QoS policy.
	IsDefault *bool `pulumi:"isDefault"`
	// The name of the QoS policy. Changing this updates the name of
	// the existing QoS policy.
	Name *string `pulumi:"name"`
	// The owner of the QoS policy. Required if admin wants to
	// create a QoS policy for another project. Changing this creates a new QoS policy.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron Qos policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// QoS policy.
	Region *string `pulumi:"region"`
	// Indicates whether this QoS policy is shared across
	// all projects. Changing this updates the shared status of the existing
	// QoS policy.
	Shared *bool `pulumi:"shared"`
	// A set of string tags for the QoS policy.
	Tags []string `pulumi:"tags"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a QosPolicy resource.
type QosPolicyArgs struct {
	// The human-readable description for the QoS policy.
	// Changing this updates the description of the existing QoS policy.
	Description pulumi.StringPtrInput
	// Indicates whether the QoS policy is default
	// QoS policy or not. Changing this updates the default status of the existing
	// QoS policy.
	IsDefault pulumi.BoolPtrInput
	// The name of the QoS policy. Changing this updates the name of
	// the existing QoS policy.
	Name pulumi.StringPtrInput
	// The owner of the QoS policy. Required if admin wants to
	// create a QoS policy for another project. Changing this creates a new QoS policy.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron Qos policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// QoS policy.
	Region pulumi.StringPtrInput
	// Indicates whether this QoS policy is shared across
	// all projects. Changing this updates the shared status of the existing
	// QoS policy.
	Shared pulumi.BoolPtrInput
	// A set of string tags for the QoS policy.
	Tags pulumi.StringArrayInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (QosPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qosPolicyArgs)(nil)).Elem()
}

type QosPolicyInput interface {
	pulumi.Input

	ToQosPolicyOutput() QosPolicyOutput
	ToQosPolicyOutputWithContext(ctx context.Context) QosPolicyOutput
}

func (*QosPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**QosPolicy)(nil)).Elem()
}

func (i *QosPolicy) ToQosPolicyOutput() QosPolicyOutput {
	return i.ToQosPolicyOutputWithContext(context.Background())
}

func (i *QosPolicy) ToQosPolicyOutputWithContext(ctx context.Context) QosPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyOutput)
}

// QosPolicyArrayInput is an input type that accepts QosPolicyArray and QosPolicyArrayOutput values.
// You can construct a concrete instance of `QosPolicyArrayInput` via:
//
//	QosPolicyArray{ QosPolicyArgs{...} }
type QosPolicyArrayInput interface {
	pulumi.Input

	ToQosPolicyArrayOutput() QosPolicyArrayOutput
	ToQosPolicyArrayOutputWithContext(context.Context) QosPolicyArrayOutput
}

type QosPolicyArray []QosPolicyInput

func (QosPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QosPolicy)(nil)).Elem()
}

func (i QosPolicyArray) ToQosPolicyArrayOutput() QosPolicyArrayOutput {
	return i.ToQosPolicyArrayOutputWithContext(context.Background())
}

func (i QosPolicyArray) ToQosPolicyArrayOutputWithContext(ctx context.Context) QosPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyArrayOutput)
}

// QosPolicyMapInput is an input type that accepts QosPolicyMap and QosPolicyMapOutput values.
// You can construct a concrete instance of `QosPolicyMapInput` via:
//
//	QosPolicyMap{ "key": QosPolicyArgs{...} }
type QosPolicyMapInput interface {
	pulumi.Input

	ToQosPolicyMapOutput() QosPolicyMapOutput
	ToQosPolicyMapOutputWithContext(context.Context) QosPolicyMapOutput
}

type QosPolicyMap map[string]QosPolicyInput

func (QosPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QosPolicy)(nil)).Elem()
}

func (i QosPolicyMap) ToQosPolicyMapOutput() QosPolicyMapOutput {
	return i.ToQosPolicyMapOutputWithContext(context.Background())
}

func (i QosPolicyMap) ToQosPolicyMapOutputWithContext(ctx context.Context) QosPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosPolicyMapOutput)
}

type QosPolicyOutput struct{ *pulumi.OutputState }

func (QosPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QosPolicy)(nil)).Elem()
}

func (o QosPolicyOutput) ToQosPolicyOutput() QosPolicyOutput {
	return o
}

func (o QosPolicyOutput) ToQosPolicyOutputWithContext(ctx context.Context) QosPolicyOutput {
	return o
}

// The collection of tags assigned on the QoS policy, which have been
// explicitly and implicitly added.
func (o QosPolicyOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.StringArrayOutput { return v.AllTags }).(pulumi.StringArrayOutput)
}

// The time at which QoS policy was created.
func (o QosPolicyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The human-readable description for the QoS policy.
// Changing this updates the description of the existing QoS policy.
func (o QosPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether the QoS policy is default
// QoS policy or not. Changing this updates the default status of the existing
// QoS policy.
func (o QosPolicyOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// The name of the QoS policy. Changing this updates the name of
// the existing QoS policy.
func (o QosPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The owner of the QoS policy. Required if admin wants to
// create a QoS policy for another project. Changing this creates a new QoS policy.
func (o QosPolicyOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a Neutron Qos policy. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// QoS policy.
func (o QosPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The revision number of the QoS policy.
func (o QosPolicyOutput) RevisionNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.IntOutput { return v.RevisionNumber }).(pulumi.IntOutput)
}

// Indicates whether this QoS policy is shared across
// all projects. Changing this updates the shared status of the existing
// QoS policy.
func (o QosPolicyOutput) Shared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.BoolPtrOutput { return v.Shared }).(pulumi.BoolPtrOutput)
}

// A set of string tags for the QoS policy.
func (o QosPolicyOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The time at which QoS policy was created.
func (o QosPolicyOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Map of additional options.
func (o QosPolicyOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *QosPolicy) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type QosPolicyArrayOutput struct{ *pulumi.OutputState }

func (QosPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QosPolicy)(nil)).Elem()
}

func (o QosPolicyArrayOutput) ToQosPolicyArrayOutput() QosPolicyArrayOutput {
	return o
}

func (o QosPolicyArrayOutput) ToQosPolicyArrayOutputWithContext(ctx context.Context) QosPolicyArrayOutput {
	return o
}

func (o QosPolicyArrayOutput) Index(i pulumi.IntInput) QosPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QosPolicy {
		return vs[0].([]*QosPolicy)[vs[1].(int)]
	}).(QosPolicyOutput)
}

type QosPolicyMapOutput struct{ *pulumi.OutputState }

func (QosPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QosPolicy)(nil)).Elem()
}

func (o QosPolicyMapOutput) ToQosPolicyMapOutput() QosPolicyMapOutput {
	return o
}

func (o QosPolicyMapOutput) ToQosPolicyMapOutputWithContext(ctx context.Context) QosPolicyMapOutput {
	return o
}

func (o QosPolicyMapOutput) MapIndex(k pulumi.StringInput) QosPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QosPolicy {
		return vs[0].(map[string]*QosPolicy)[vs[1].(string)]
	}).(QosPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QosPolicyInput)(nil)).Elem(), &QosPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosPolicyArrayInput)(nil)).Elem(), QosPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosPolicyMapInput)(nil)).Elem(), QosPolicyMap{})
	pulumi.RegisterOutputType(QosPolicyOutput{})
	pulumi.RegisterOutputType(QosPolicyArrayOutput{})
	pulumi.RegisterOutputType(QosPolicyMapOutput{})
}
