// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about host aggregates
// by name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.LookupAggregateV2(ctx, &compute.LookupAggregateV2Args{
// 			Name: "test",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAggregateV2(ctx *pulumi.Context, args *LookupAggregateV2Args, opts ...pulumi.InvokeOption) (*LookupAggregateV2Result, error) {
	var rv LookupAggregateV2Result
	err := ctx.Invoke("openstack:compute/getAggregateV2:getAggregateV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAggregateV2.
type LookupAggregateV2Args struct {
	// List of Hypervisors contained in the Host Aggregate
	Hosts []string `pulumi:"hosts"`
	// Metadata of the Host Aggregate
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the host aggregate
	Name string `pulumi:"name"`
}

// A collection of values returned by getAggregateV2.
type LookupAggregateV2Result struct {
	// List of Hypervisors contained in the Host Aggregate
	Hosts []string `pulumi:"hosts"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Metadata of the Host Aggregate
	Metadata map[string]string `pulumi:"metadata"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// Availability zone of the Host Aggregate
	Zone string `pulumi:"zone"`
}

func LookupAggregateV2Output(ctx *pulumi.Context, args LookupAggregateV2OutputArgs, opts ...pulumi.InvokeOption) LookupAggregateV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAggregateV2Result, error) {
			args := v.(LookupAggregateV2Args)
			r, err := LookupAggregateV2(ctx, &args, opts...)
			return *r, err
		}).(LookupAggregateV2ResultOutput)
}

// A collection of arguments for invoking getAggregateV2.
type LookupAggregateV2OutputArgs struct {
	// List of Hypervisors contained in the Host Aggregate
	Hosts pulumi.StringArrayInput `pulumi:"hosts"`
	// Metadata of the Host Aggregate
	Metadata pulumi.StringMapInput `pulumi:"metadata"`
	// The name of the host aggregate
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupAggregateV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAggregateV2Args)(nil)).Elem()
}

// A collection of values returned by getAggregateV2.
type LookupAggregateV2ResultOutput struct{ *pulumi.OutputState }

func (LookupAggregateV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAggregateV2Result)(nil)).Elem()
}

func (o LookupAggregateV2ResultOutput) ToLookupAggregateV2ResultOutput() LookupAggregateV2ResultOutput {
	return o
}

func (o LookupAggregateV2ResultOutput) ToLookupAggregateV2ResultOutputWithContext(ctx context.Context) LookupAggregateV2ResultOutput {
	return o
}

// List of Hypervisors contained in the Host Aggregate
func (o LookupAggregateV2ResultOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAggregateV2Result) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAggregateV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAggregateV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// Metadata of the Host Aggregate
func (o LookupAggregateV2ResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAggregateV2Result) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// See Argument Reference above.
func (o LookupAggregateV2ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAggregateV2Result) string { return v.Name }).(pulumi.StringOutput)
}

// Availability zone of the Host Aggregate
func (o LookupAggregateV2ResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAggregateV2Result) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAggregateV2ResultOutput{})
}
