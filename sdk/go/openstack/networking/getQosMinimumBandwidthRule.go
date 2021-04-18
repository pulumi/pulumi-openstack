// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.
func LookupQosMinimumBandwidthRule(ctx *pulumi.Context, args *LookupQosMinimumBandwidthRuleArgs, opts ...pulumi.InvokeOption) (*LookupQosMinimumBandwidthRuleResult, error) {
	var rv LookupQosMinimumBandwidthRuleResult
	err := ctx.Invoke("openstack:networking/getQosMinimumBandwidthRule:getQosMinimumBandwidthRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQosMinimumBandwidthRule.
type LookupQosMinimumBandwidthRuleArgs struct {
	Direction *string `pulumi:"direction"`
	// The value of a minimum kbps bandwidth.
	MinKbps *int `pulumi:"minKbps"`
	// The QoS policy reference.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getQosMinimumBandwidthRule.
type LookupQosMinimumBandwidthRuleResult struct {
	Direction string `pulumi:"direction"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	MinKbps int `pulumi:"minKbps"`
	// See Argument Reference above.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
}
