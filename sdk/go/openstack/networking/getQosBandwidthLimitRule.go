// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack QoS bandwidth limit rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networking.LookupQosBandwidthLimitRule(ctx, &networking.LookupQosBandwidthLimitRuleArgs{
// 			MaxKbps: pulumi.IntRef(300),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupQosBandwidthLimitRule(ctx *pulumi.Context, args *LookupQosBandwidthLimitRuleArgs, opts ...pulumi.InvokeOption) (*LookupQosBandwidthLimitRuleResult, error) {
	var rv LookupQosBandwidthLimitRuleResult
	err := ctx.Invoke("openstack:networking/getQosBandwidthLimitRule:getQosBandwidthLimitRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQosBandwidthLimitRule.
type LookupQosBandwidthLimitRuleArgs struct {
	// The maximum burst size in kilobits of a QoS bandwidth limit rule.
	MaxBurstKbps *int `pulumi:"maxBurstKbps"`
	// The maximum kilobits per second of a QoS bandwidth limit rule.
	MaxKbps *int `pulumi:"maxKbps"`
	// The QoS policy reference.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getQosBandwidthLimitRule.
type LookupQosBandwidthLimitRuleResult struct {
	// See Argument Reference above.
	Direction string `pulumi:"direction"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	MaxBurstKbps int `pulumi:"maxBurstKbps"`
	// See Argument Reference above.
	MaxKbps int `pulumi:"maxKbps"`
	// See Argument Reference above.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
}

func LookupQosBandwidthLimitRuleOutput(ctx *pulumi.Context, args LookupQosBandwidthLimitRuleOutputArgs, opts ...pulumi.InvokeOption) LookupQosBandwidthLimitRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQosBandwidthLimitRuleResult, error) {
			args := v.(LookupQosBandwidthLimitRuleArgs)
			r, err := LookupQosBandwidthLimitRule(ctx, &args, opts...)
			return *r, err
		}).(LookupQosBandwidthLimitRuleResultOutput)
}

// A collection of arguments for invoking getQosBandwidthLimitRule.
type LookupQosBandwidthLimitRuleOutputArgs struct {
	// The maximum burst size in kilobits of a QoS bandwidth limit rule.
	MaxBurstKbps pulumi.IntPtrInput `pulumi:"maxBurstKbps"`
	// The maximum kilobits per second of a QoS bandwidth limit rule.
	MaxKbps pulumi.IntPtrInput `pulumi:"maxKbps"`
	// The QoS policy reference.
	QosPolicyId pulumi.StringInput `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupQosBandwidthLimitRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQosBandwidthLimitRuleArgs)(nil)).Elem()
}

// A collection of values returned by getQosBandwidthLimitRule.
type LookupQosBandwidthLimitRuleResultOutput struct{ *pulumi.OutputState }

func (LookupQosBandwidthLimitRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQosBandwidthLimitRuleResult)(nil)).Elem()
}

func (o LookupQosBandwidthLimitRuleResultOutput) ToLookupQosBandwidthLimitRuleResultOutput() LookupQosBandwidthLimitRuleResultOutput {
	return o
}

func (o LookupQosBandwidthLimitRuleResultOutput) ToLookupQosBandwidthLimitRuleResultOutputWithContext(ctx context.Context) LookupQosBandwidthLimitRuleResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupQosBandwidthLimitRuleResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosBandwidthLimitRuleResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQosBandwidthLimitRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosBandwidthLimitRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosBandwidthLimitRuleResultOutput) MaxBurstKbps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQosBandwidthLimitRuleResult) int { return v.MaxBurstKbps }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupQosBandwidthLimitRuleResultOutput) MaxKbps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQosBandwidthLimitRuleResult) int { return v.MaxKbps }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupQosBandwidthLimitRuleResultOutput) QosPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosBandwidthLimitRuleResult) string { return v.QosPolicyId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosBandwidthLimitRuleResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosBandwidthLimitRuleResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQosBandwidthLimitRuleResultOutput{})
}
