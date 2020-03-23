// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_dscp_marking_rule_v2.html.markdown.
func LookupQosDscpMarkingRule(ctx *pulumi.Context, args *LookupQosDscpMarkingRuleArgs, opts ...pulumi.InvokeOption) (*LookupQosDscpMarkingRuleResult, error) {
	var rv LookupQosDscpMarkingRuleResult
	err := ctx.Invoke("openstack:networking/getQosDscpMarkingRule:getQosDscpMarkingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQosDscpMarkingRule.
type LookupQosDscpMarkingRuleArgs struct {
	// The value of a DSCP mark.
	DscpMark *int `pulumi:"dscpMark"`
	// The QoS policy reference.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}


// A collection of values returned by getQosDscpMarkingRule.
type LookupQosDscpMarkingRuleResult struct {
	// See Argument Reference above.
	DscpMark int `pulumi:"dscpMark"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
}

