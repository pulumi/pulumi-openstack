// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.
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
// 		_, err := networking.LookupQosMinimumBandwidthRule(ctx, &networking.LookupQosMinimumBandwidthRuleArgs{
// 			MinKbps: pulumi.IntRef(2000),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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

func LookupQosMinimumBandwidthRuleOutput(ctx *pulumi.Context, args LookupQosMinimumBandwidthRuleOutputArgs, opts ...pulumi.InvokeOption) LookupQosMinimumBandwidthRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQosMinimumBandwidthRuleResult, error) {
			args := v.(LookupQosMinimumBandwidthRuleArgs)
			r, err := LookupQosMinimumBandwidthRule(ctx, &args, opts...)
			var s LookupQosMinimumBandwidthRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQosMinimumBandwidthRuleResultOutput)
}

// A collection of arguments for invoking getQosMinimumBandwidthRule.
type LookupQosMinimumBandwidthRuleOutputArgs struct {
	Direction pulumi.StringPtrInput `pulumi:"direction"`
	// The value of a minimum kbps bandwidth.
	MinKbps pulumi.IntPtrInput `pulumi:"minKbps"`
	// The QoS policy reference.
	QosPolicyId pulumi.StringInput `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupQosMinimumBandwidthRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQosMinimumBandwidthRuleArgs)(nil)).Elem()
}

// A collection of values returned by getQosMinimumBandwidthRule.
type LookupQosMinimumBandwidthRuleResultOutput struct{ *pulumi.OutputState }

func (LookupQosMinimumBandwidthRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQosMinimumBandwidthRuleResult)(nil)).Elem()
}

func (o LookupQosMinimumBandwidthRuleResultOutput) ToLookupQosMinimumBandwidthRuleResultOutput() LookupQosMinimumBandwidthRuleResultOutput {
	return o
}

func (o LookupQosMinimumBandwidthRuleResultOutput) ToLookupQosMinimumBandwidthRuleResultOutputWithContext(ctx context.Context) LookupQosMinimumBandwidthRuleResultOutput {
	return o
}

func (o LookupQosMinimumBandwidthRuleResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosMinimumBandwidthRuleResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQosMinimumBandwidthRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosMinimumBandwidthRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosMinimumBandwidthRuleResultOutput) MinKbps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQosMinimumBandwidthRuleResult) int { return v.MinKbps }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupQosMinimumBandwidthRuleResultOutput) QosPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosMinimumBandwidthRuleResult) string { return v.QosPolicyId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosMinimumBandwidthRuleResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosMinimumBandwidthRuleResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQosMinimumBandwidthRuleResultOutput{})
}
