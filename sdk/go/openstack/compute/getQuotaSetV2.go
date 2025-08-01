// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the compute quotaset of an OpenStack project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.LookupQuotaSetV2(ctx, &compute.LookupQuotaSetV2Args{
//				ProjectId: "2e367a3d29f94fd988e6ec54e305ec9d",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupQuotaSetV2(ctx *pulumi.Context, args *LookupQuotaSetV2Args, opts ...pulumi.InvokeOption) (*LookupQuotaSetV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQuotaSetV2Result
	err := ctx.Invoke("openstack:compute/getQuotaSetV2:getQuotaSetV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQuotaSetV2.
type LookupQuotaSetV2Args struct {
	// The id of the project to retrieve the quotaset.
	ProjectId string `pulumi:"projectId"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getQuotaSetV2.
type LookupQuotaSetV2Result struct {
	// The number of allowed server cores.
	Cores int `pulumi:"cores"`
	// The number of allowed fixed IP addresses. Available until version 2.35.
	FixedIps int `pulumi:"fixedIps"`
	// The number of allowed floating IP addresses. Available until version 2.35.
	FloatingIps int `pulumi:"floatingIps"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The number of allowed bytes of content for each injected file. Available until version 2.56.
	InjectedFileContentBytes int `pulumi:"injectedFileContentBytes"`
	// The number of allowed bytes for each injected file path. Available until version 2.56.
	InjectedFilePathBytes int `pulumi:"injectedFilePathBytes"`
	// The number of allowed injected files. Available until version 2.56.
	InjectedFiles int `pulumi:"injectedFiles"`
	// The number of allowed servers.
	Instances int `pulumi:"instances"`
	// The number of allowed key pairs for each user.
	KeyPairs int `pulumi:"keyPairs"`
	// The number of allowed metadata items for each server.
	MetadataItems int `pulumi:"metadataItems"`
	// See Argument Reference above.
	ProjectId string `pulumi:"projectId"`
	// The amount of allowed server RAM, in MiB.
	Ram int `pulumi:"ram"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The number of allowed rules for each security group. Available until version 2.35.
	SecurityGroupRules int `pulumi:"securityGroupRules"`
	// The number of allowed security groups. Available until version 2.35.
	SecurityGroups int `pulumi:"securityGroups"`
	// The number of allowed members for each server group.
	ServerGroupMembers int `pulumi:"serverGroupMembers"`
	// The number of allowed server groups.
	ServerGroups int `pulumi:"serverGroups"`
}

func LookupQuotaSetV2Output(ctx *pulumi.Context, args LookupQuotaSetV2OutputArgs, opts ...pulumi.InvokeOption) LookupQuotaSetV2ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupQuotaSetV2ResultOutput, error) {
			args := v.(LookupQuotaSetV2Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:compute/getQuotaSetV2:getQuotaSetV2", args, LookupQuotaSetV2ResultOutput{}, options).(LookupQuotaSetV2ResultOutput), nil
		}).(LookupQuotaSetV2ResultOutput)
}

// A collection of arguments for invoking getQuotaSetV2.
type LookupQuotaSetV2OutputArgs struct {
	// The id of the project to retrieve the quotaset.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupQuotaSetV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQuotaSetV2Args)(nil)).Elem()
}

// A collection of values returned by getQuotaSetV2.
type LookupQuotaSetV2ResultOutput struct{ *pulumi.OutputState }

func (LookupQuotaSetV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQuotaSetV2Result)(nil)).Elem()
}

func (o LookupQuotaSetV2ResultOutput) ToLookupQuotaSetV2ResultOutput() LookupQuotaSetV2ResultOutput {
	return o
}

func (o LookupQuotaSetV2ResultOutput) ToLookupQuotaSetV2ResultOutputWithContext(ctx context.Context) LookupQuotaSetV2ResultOutput {
	return o
}

// The number of allowed server cores.
func (o LookupQuotaSetV2ResultOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.Cores }).(pulumi.IntOutput)
}

// The number of allowed fixed IP addresses. Available until version 2.35.
func (o LookupQuotaSetV2ResultOutput) FixedIps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.FixedIps }).(pulumi.IntOutput)
}

// The number of allowed floating IP addresses. Available until version 2.35.
func (o LookupQuotaSetV2ResultOutput) FloatingIps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.FloatingIps }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQuotaSetV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// The number of allowed bytes of content for each injected file. Available until version 2.56.
func (o LookupQuotaSetV2ResultOutput) InjectedFileContentBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.InjectedFileContentBytes }).(pulumi.IntOutput)
}

// The number of allowed bytes for each injected file path. Available until version 2.56.
func (o LookupQuotaSetV2ResultOutput) InjectedFilePathBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.InjectedFilePathBytes }).(pulumi.IntOutput)
}

// The number of allowed injected files. Available until version 2.56.
func (o LookupQuotaSetV2ResultOutput) InjectedFiles() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.InjectedFiles }).(pulumi.IntOutput)
}

// The number of allowed servers.
func (o LookupQuotaSetV2ResultOutput) Instances() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.Instances }).(pulumi.IntOutput)
}

// The number of allowed key pairs for each user.
func (o LookupQuotaSetV2ResultOutput) KeyPairs() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.KeyPairs }).(pulumi.IntOutput)
}

// The number of allowed metadata items for each server.
func (o LookupQuotaSetV2ResultOutput) MetadataItems() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.MetadataItems }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupQuotaSetV2ResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The amount of allowed server RAM, in MiB.
func (o LookupQuotaSetV2ResultOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.Ram }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupQuotaSetV2ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) string { return v.Region }).(pulumi.StringOutput)
}

// The number of allowed rules for each security group. Available until version 2.35.
func (o LookupQuotaSetV2ResultOutput) SecurityGroupRules() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.SecurityGroupRules }).(pulumi.IntOutput)
}

// The number of allowed security groups. Available until version 2.35.
func (o LookupQuotaSetV2ResultOutput) SecurityGroups() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.SecurityGroups }).(pulumi.IntOutput)
}

// The number of allowed members for each server group.
func (o LookupQuotaSetV2ResultOutput) ServerGroupMembers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.ServerGroupMembers }).(pulumi.IntOutput)
}

// The number of allowed server groups.
func (o LookupQuotaSetV2ResultOutput) ServerGroups() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuotaSetV2Result) int { return v.ServerGroups }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQuotaSetV2ResultOutput{})
}
