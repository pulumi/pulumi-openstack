// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 Neutron QoS DSCP marking rule resource within OpenStack.
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
	if args == nil || args.DscpMark == nil {
		return nil, errors.New("missing required argument 'DscpMark'")
	}
	if args == nil || args.QosPolicyId == nil {
		return nil, errors.New("missing required argument 'QosPolicyId'")
	}
	if args == nil {
		args = &QosDscpMarkingRuleArgs{}
	}
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
