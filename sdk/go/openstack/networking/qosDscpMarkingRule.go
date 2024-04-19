// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Neutron QoS DSCP marking rule resource within OpenStack.
//
// ## Example Usage
//
// ### Create a QoS Policy with some DSCP marking rule
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
//			qosPolicy1, err := networking.NewQosPolicy(ctx, "qos_policy_1", &networking.QosPolicyArgs{
//				Name:        pulumi.String("qos_policy_1"),
//				Description: pulumi.String("dscp_mark"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewQosDscpMarkingRule(ctx, "dscp_marking_rule_1", &networking.QosDscpMarkingRuleArgs{
//				QosPolicyId: qosPolicy1.ID(),
//				DscpMark:    pulumi.Int(26),
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
// QoS DSCP marking rules can be imported using the `qos_policy_id/dscp_marking_rule_id` format, e.g.
//
// ```sh
// $ pulumi import openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule dscp_marking_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
// ```
type QosDscpMarkingRule struct {
	pulumi.CustomResourceState

	// The value of DSCP mark. Changing this updates the DSCP mark value existing
	// QoS DSCP marking rule.
	DscpMark pulumi.IntOutput `pulumi:"dscpMark"`
	// The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
	QosPolicyId pulumi.StringOutput `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewQosDscpMarkingRule registers a new resource with the given unique name, arguments, and options.
func NewQosDscpMarkingRule(ctx *pulumi.Context,
	name string, args *QosDscpMarkingRuleArgs, opts ...pulumi.ResourceOption) (*QosDscpMarkingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DscpMark == nil {
		return nil, errors.New("invalid value for required argument 'DscpMark'")
	}
	if args.QosPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'QosPolicyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QosDscpMarkingRule
	err := ctx.RegisterResource("openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQosDscpMarkingRule gets an existing QosDscpMarkingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQosDscpMarkingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QosDscpMarkingRuleState, opts ...pulumi.ResourceOption) (*QosDscpMarkingRule, error) {
	var resource QosDscpMarkingRule
	err := ctx.ReadResource("openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QosDscpMarkingRule resources.
type qosDscpMarkingRuleState struct {
	// The value of DSCP mark. Changing this updates the DSCP mark value existing
	// QoS DSCP marking rule.
	DscpMark *int `pulumi:"dscpMark"`
	// The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
	QosPolicyId *string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
	Region *string `pulumi:"region"`
}

type QosDscpMarkingRuleState struct {
	// The value of DSCP mark. Changing this updates the DSCP mark value existing
	// QoS DSCP marking rule.
	DscpMark pulumi.IntPtrInput
	// The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
	QosPolicyId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
	Region pulumi.StringPtrInput
}

func (QosDscpMarkingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*qosDscpMarkingRuleState)(nil)).Elem()
}

type qosDscpMarkingRuleArgs struct {
	// The value of DSCP mark. Changing this updates the DSCP mark value existing
	// QoS DSCP marking rule.
	DscpMark int `pulumi:"dscpMark"`
	// The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a QosDscpMarkingRule resource.
type QosDscpMarkingRuleArgs struct {
	// The value of DSCP mark. Changing this updates the DSCP mark value existing
	// QoS DSCP marking rule.
	DscpMark pulumi.IntInput
	// The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
	QosPolicyId pulumi.StringInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
	Region pulumi.StringPtrInput
}

func (QosDscpMarkingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qosDscpMarkingRuleArgs)(nil)).Elem()
}

type QosDscpMarkingRuleInput interface {
	pulumi.Input

	ToQosDscpMarkingRuleOutput() QosDscpMarkingRuleOutput
	ToQosDscpMarkingRuleOutputWithContext(ctx context.Context) QosDscpMarkingRuleOutput
}

func (*QosDscpMarkingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**QosDscpMarkingRule)(nil)).Elem()
}

func (i *QosDscpMarkingRule) ToQosDscpMarkingRuleOutput() QosDscpMarkingRuleOutput {
	return i.ToQosDscpMarkingRuleOutputWithContext(context.Background())
}

func (i *QosDscpMarkingRule) ToQosDscpMarkingRuleOutputWithContext(ctx context.Context) QosDscpMarkingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosDscpMarkingRuleOutput)
}

// QosDscpMarkingRuleArrayInput is an input type that accepts QosDscpMarkingRuleArray and QosDscpMarkingRuleArrayOutput values.
// You can construct a concrete instance of `QosDscpMarkingRuleArrayInput` via:
//
//	QosDscpMarkingRuleArray{ QosDscpMarkingRuleArgs{...} }
type QosDscpMarkingRuleArrayInput interface {
	pulumi.Input

	ToQosDscpMarkingRuleArrayOutput() QosDscpMarkingRuleArrayOutput
	ToQosDscpMarkingRuleArrayOutputWithContext(context.Context) QosDscpMarkingRuleArrayOutput
}

type QosDscpMarkingRuleArray []QosDscpMarkingRuleInput

func (QosDscpMarkingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QosDscpMarkingRule)(nil)).Elem()
}

func (i QosDscpMarkingRuleArray) ToQosDscpMarkingRuleArrayOutput() QosDscpMarkingRuleArrayOutput {
	return i.ToQosDscpMarkingRuleArrayOutputWithContext(context.Background())
}

func (i QosDscpMarkingRuleArray) ToQosDscpMarkingRuleArrayOutputWithContext(ctx context.Context) QosDscpMarkingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosDscpMarkingRuleArrayOutput)
}

// QosDscpMarkingRuleMapInput is an input type that accepts QosDscpMarkingRuleMap and QosDscpMarkingRuleMapOutput values.
// You can construct a concrete instance of `QosDscpMarkingRuleMapInput` via:
//
//	QosDscpMarkingRuleMap{ "key": QosDscpMarkingRuleArgs{...} }
type QosDscpMarkingRuleMapInput interface {
	pulumi.Input

	ToQosDscpMarkingRuleMapOutput() QosDscpMarkingRuleMapOutput
	ToQosDscpMarkingRuleMapOutputWithContext(context.Context) QosDscpMarkingRuleMapOutput
}

type QosDscpMarkingRuleMap map[string]QosDscpMarkingRuleInput

func (QosDscpMarkingRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QosDscpMarkingRule)(nil)).Elem()
}

func (i QosDscpMarkingRuleMap) ToQosDscpMarkingRuleMapOutput() QosDscpMarkingRuleMapOutput {
	return i.ToQosDscpMarkingRuleMapOutputWithContext(context.Background())
}

func (i QosDscpMarkingRuleMap) ToQosDscpMarkingRuleMapOutputWithContext(ctx context.Context) QosDscpMarkingRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QosDscpMarkingRuleMapOutput)
}

type QosDscpMarkingRuleOutput struct{ *pulumi.OutputState }

func (QosDscpMarkingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QosDscpMarkingRule)(nil)).Elem()
}

func (o QosDscpMarkingRuleOutput) ToQosDscpMarkingRuleOutput() QosDscpMarkingRuleOutput {
	return o
}

func (o QosDscpMarkingRuleOutput) ToQosDscpMarkingRuleOutputWithContext(ctx context.Context) QosDscpMarkingRuleOutput {
	return o
}

// The value of DSCP mark. Changing this updates the DSCP mark value existing
// QoS DSCP marking rule.
func (o QosDscpMarkingRuleOutput) DscpMark() pulumi.IntOutput {
	return o.ApplyT(func(v *QosDscpMarkingRule) pulumi.IntOutput { return v.DscpMark }).(pulumi.IntOutput)
}

// The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
func (o QosDscpMarkingRuleOutput) QosPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *QosDscpMarkingRule) pulumi.StringOutput { return v.QosPolicyId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
// `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
func (o QosDscpMarkingRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *QosDscpMarkingRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type QosDscpMarkingRuleArrayOutput struct{ *pulumi.OutputState }

func (QosDscpMarkingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QosDscpMarkingRule)(nil)).Elem()
}

func (o QosDscpMarkingRuleArrayOutput) ToQosDscpMarkingRuleArrayOutput() QosDscpMarkingRuleArrayOutput {
	return o
}

func (o QosDscpMarkingRuleArrayOutput) ToQosDscpMarkingRuleArrayOutputWithContext(ctx context.Context) QosDscpMarkingRuleArrayOutput {
	return o
}

func (o QosDscpMarkingRuleArrayOutput) Index(i pulumi.IntInput) QosDscpMarkingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QosDscpMarkingRule {
		return vs[0].([]*QosDscpMarkingRule)[vs[1].(int)]
	}).(QosDscpMarkingRuleOutput)
}

type QosDscpMarkingRuleMapOutput struct{ *pulumi.OutputState }

func (QosDscpMarkingRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QosDscpMarkingRule)(nil)).Elem()
}

func (o QosDscpMarkingRuleMapOutput) ToQosDscpMarkingRuleMapOutput() QosDscpMarkingRuleMapOutput {
	return o
}

func (o QosDscpMarkingRuleMapOutput) ToQosDscpMarkingRuleMapOutputWithContext(ctx context.Context) QosDscpMarkingRuleMapOutput {
	return o
}

func (o QosDscpMarkingRuleMapOutput) MapIndex(k pulumi.StringInput) QosDscpMarkingRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QosDscpMarkingRule {
		return vs[0].(map[string]*QosDscpMarkingRule)[vs[1].(string)]
	}).(QosDscpMarkingRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QosDscpMarkingRuleInput)(nil)).Elem(), &QosDscpMarkingRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosDscpMarkingRuleArrayInput)(nil)).Elem(), QosDscpMarkingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QosDscpMarkingRuleMapInput)(nil)).Elem(), QosDscpMarkingRuleMap{})
	pulumi.RegisterOutputType(QosDscpMarkingRuleOutput{})
	pulumi.RegisterOutputType(QosDscpMarkingRuleArrayOutput{})
	pulumi.RegisterOutputType(QosDscpMarkingRuleMapOutput{})
}
