// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.LookupQosDscpMarkingRule(ctx, &networking.LookupQosDscpMarkingRuleArgs{
//				DscpMark: pulumi.IntRef(26),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupQosDscpMarkingRule(ctx *pulumi.Context, args *LookupQosDscpMarkingRuleArgs, opts ...pulumi.InvokeOption) (*LookupQosDscpMarkingRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
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
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	QosPolicyId string `pulumi:"qosPolicyId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
}

func LookupQosDscpMarkingRuleOutput(ctx *pulumi.Context, args LookupQosDscpMarkingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupQosDscpMarkingRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQosDscpMarkingRuleResultOutput, error) {
			args := v.(LookupQosDscpMarkingRuleArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupQosDscpMarkingRuleResult
			secret, err := ctx.InvokePackageRaw("openstack:networking/getQosDscpMarkingRule:getQosDscpMarkingRule", args, &rv, "", opts...)
			if err != nil {
				return LookupQosDscpMarkingRuleResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupQosDscpMarkingRuleResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupQosDscpMarkingRuleResultOutput), nil
			}
			return output, nil
		}).(LookupQosDscpMarkingRuleResultOutput)
}

// A collection of arguments for invoking getQosDscpMarkingRule.
type LookupQosDscpMarkingRuleOutputArgs struct {
	// The value of a DSCP mark.
	DscpMark pulumi.IntPtrInput `pulumi:"dscpMark"`
	// The QoS policy reference.
	QosPolicyId pulumi.StringInput `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupQosDscpMarkingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQosDscpMarkingRuleArgs)(nil)).Elem()
}

// A collection of values returned by getQosDscpMarkingRule.
type LookupQosDscpMarkingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupQosDscpMarkingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQosDscpMarkingRuleResult)(nil)).Elem()
}

func (o LookupQosDscpMarkingRuleResultOutput) ToLookupQosDscpMarkingRuleResultOutput() LookupQosDscpMarkingRuleResultOutput {
	return o
}

func (o LookupQosDscpMarkingRuleResultOutput) ToLookupQosDscpMarkingRuleResultOutputWithContext(ctx context.Context) LookupQosDscpMarkingRuleResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupQosDscpMarkingRuleResultOutput) DscpMark() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQosDscpMarkingRuleResult) int { return v.DscpMark }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQosDscpMarkingRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosDscpMarkingRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosDscpMarkingRuleResultOutput) QosPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosDscpMarkingRuleResult) string { return v.QosPolicyId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupQosDscpMarkingRuleResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQosDscpMarkingRuleResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQosDscpMarkingRuleResultOutput{})
}
