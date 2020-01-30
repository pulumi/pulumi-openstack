// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package networking

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 Neutron QoS minimum bandwidth rule resource within OpenStack.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_qos_minimum_bandwidth_rule_v2.html.markdown.
type QosMinimumBandwidthRule struct {
	pulumi.CustomResourceState

	// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
	// existing QoS minimum bandwidth rule.
	Direction pulumi.StringPtrOutput `pulumi:"direction"`
	// The minimum kilobits per second. Changing this updates the min kbps value of the existing
	// QoS minimum bandwidth rule.
	MinKbps pulumi.IntOutput `pulumi:"minKbps"`
	// The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
	QosPolicyId pulumi.StringOutput `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewQosMinimumBandwidthRule registers a new resource with the given unique name, arguments, and options.
func NewQosMinimumBandwidthRule(ctx *pulumi.Context,
	name string, args *QosMinimumBandwidthRuleArgs, opts ...pulumi.ResourceOption) (*QosMinimumBandwidthRule, error) {
	if args == nil || args.MinKbps == nil {
		return nil, errors.New("missing required argument 'MinKbps'")
	}
	if args == nil || args.QosPolicyId == nil {
		return nil, errors.New("missing required argument 'QosPolicyId'")
	}
	if args == nil {
		args = &QosMinimumBandwidthRuleArgs{}
	}
	var resource QosMinimumBandwidthRule
	err := ctx.RegisterResource("openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQosMinimumBandwidthRule gets an existing QosMinimumBandwidthRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQosMinimumBandwidthRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QosMinimumBandwidthRuleState, opts ...pulumi.ResourceOption) (*QosMinimumBandwidthRule, error) {
	var resource QosMinimumBandwidthRule
	err := ctx.ReadResource("openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QosMinimumBandwidthRule resources.
type qosMinimumBandwidthRuleState struct {
	// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
	// existing QoS minimum bandwidth rule.
	Direction *string `pulumi:"direction"`
	// The minimum kilobits per second. Changing this updates the min kbps value of the existing
	// QoS minimum bandwidth rule.
	MinKbps *int `pulumi:"minKbps"`
	// The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
	QosPolicyId *string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
	Region *string `pulumi:"region"`
}

type QosMinimumBandwidthRuleState struct {
	// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
	// existing QoS minimum bandwidth rule.
	Direction pulumi.StringPtrInput
	// The minimum kilobits per second. Changing this updates the min kbps value of the existing
	// QoS minimum bandwidth rule.
	MinKbps pulumi.IntPtrInput
	// The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
	QosPolicyId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
	Region pulumi.StringPtrInput
}

func (QosMinimumBandwidthRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*qosMinimumBandwidthRuleState)(nil)).Elem()
}

type qosMinimumBandwidthRuleArgs struct {
	// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
	// existing QoS minimum bandwidth rule.
	Direction *string `pulumi:"direction"`
	// The minimum kilobits per second. Changing this updates the min kbps value of the existing
	// QoS minimum bandwidth rule.
	MinKbps int `pulumi:"minKbps"`
	// The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a QosMinimumBandwidthRule resource.
type QosMinimumBandwidthRuleArgs struct {
	// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
	// existing QoS minimum bandwidth rule.
	Direction pulumi.StringPtrInput
	// The minimum kilobits per second. Changing this updates the min kbps value of the existing
	// QoS minimum bandwidth rule.
	MinKbps pulumi.IntInput
	// The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
	QosPolicyId pulumi.StringInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
	Region pulumi.StringPtrInput
}

func (QosMinimumBandwidthRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qosMinimumBandwidthRuleArgs)(nil)).Elem()
}
