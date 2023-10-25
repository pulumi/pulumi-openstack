// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to get the compute limits of an OpenStack project.
func GetLimitsV2(ctx *pulumi.Context, args *GetLimitsV2Args, opts ...pulumi.InvokeOption) (*GetLimitsV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLimitsV2Result
	err := ctx.Invoke("openstack:compute/getLimitsV2:getLimitsV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLimitsV2.
type GetLimitsV2Args struct {
	// The id of the project to retrieve the limits.
	ProjectId string `pulumi:"projectId"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getLimitsV2.
type GetLimitsV2Result struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The number of allowed metadata items for each image. Starting from version 2.39 this field is dropped from ‘os-limits’ response, because ‘image-metadata’ proxy API was deprecated. Available until version 2.38.
	MaxImageMeta int `pulumi:"maxImageMeta"`
	// The number of allowed injected files for the tenant. Available until version 2.56.
	MaxPersonality int `pulumi:"maxPersonality"`
	// The number of allowed bytes of content for each injected file. Available until version 2.56.
	MaxPersonalitySize int `pulumi:"maxPersonalitySize"`
	// The number of allowed rules for each security group. Available until version 2.35.
	MaxSecurityGroupRules int `pulumi:"maxSecurityGroupRules"`
	// The number of allowed security groups for the tenant. Available until version 2.35.
	MaxSecurityGroups int `pulumi:"maxSecurityGroups"`
	// The number of allowed members for each server group.
	MaxServerGroupMembers int `pulumi:"maxServerGroupMembers"`
	// The number of allowed server groups for the tenant.
	MaxServerGroups int `pulumi:"maxServerGroups"`
	// The number of allowed server groups for the tenant.
	MaxServerMeta int `pulumi:"maxServerMeta"`
	// The number of allowed server cores for the tenant.
	MaxTotalCores int `pulumi:"maxTotalCores"`
	// The number of allowed floating IP addresses for each tenant. Available until version 2.35.
	MaxTotalFloatingIps int `pulumi:"maxTotalFloatingIps"`
	// The number of allowed servers for the tenant.
	MaxTotalInstances int `pulumi:"maxTotalInstances"`
	// The number of allowed key pairs for the user.
	MaxTotalKeypairs int `pulumi:"maxTotalKeypairs"`
	// The number of allowed floating IP addresses for the tenant. Available until version 2.35.
	MaxTotalRamSize int `pulumi:"maxTotalRamSize"`
	// See Argument Reference above.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The number of used server cores in the tenant.
	TotalCoresUsed int `pulumi:"totalCoresUsed"`
	// The number of used floating IP addresses in the tenant.
	TotalFloatingIpsUsed int `pulumi:"totalFloatingIpsUsed"`
	// The number of used server cores in the tenant.
	TotalInstancesUsed int `pulumi:"totalInstancesUsed"`
	// The amount of used server RAM in the tenant.
	TotalRamUsed int `pulumi:"totalRamUsed"`
	// The number of used security groups in the tenant. Available until version 2.35.
	TotalSecurityGroupsUsed int `pulumi:"totalSecurityGroupsUsed"`
	// The number of used server groups in each tenant.
	TotalServerGroupsUsed int `pulumi:"totalServerGroupsUsed"`
}

func GetLimitsV2Output(ctx *pulumi.Context, args GetLimitsV2OutputArgs, opts ...pulumi.InvokeOption) GetLimitsV2ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLimitsV2Result, error) {
			args := v.(GetLimitsV2Args)
			r, err := GetLimitsV2(ctx, &args, opts...)
			var s GetLimitsV2Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLimitsV2ResultOutput)
}

// A collection of arguments for invoking getLimitsV2.
type GetLimitsV2OutputArgs struct {
	// The id of the project to retrieve the limits.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetLimitsV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLimitsV2Args)(nil)).Elem()
}

// A collection of values returned by getLimitsV2.
type GetLimitsV2ResultOutput struct{ *pulumi.OutputState }

func (GetLimitsV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLimitsV2Result)(nil)).Elem()
}

func (o GetLimitsV2ResultOutput) ToGetLimitsV2ResultOutput() GetLimitsV2ResultOutput {
	return o
}

func (o GetLimitsV2ResultOutput) ToGetLimitsV2ResultOutputWithContext(ctx context.Context) GetLimitsV2ResultOutput {
	return o
}

func (o GetLimitsV2ResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetLimitsV2Result] {
	return pulumix.Output[GetLimitsV2Result]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetLimitsV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLimitsV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// The number of allowed metadata items for each image. Starting from version 2.39 this field is dropped from ‘os-limits’ response, because ‘image-metadata’ proxy API was deprecated. Available until version 2.38.
func (o GetLimitsV2ResultOutput) MaxImageMeta() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxImageMeta }).(pulumi.IntOutput)
}

// The number of allowed injected files for the tenant. Available until version 2.56.
func (o GetLimitsV2ResultOutput) MaxPersonality() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxPersonality }).(pulumi.IntOutput)
}

// The number of allowed bytes of content for each injected file. Available until version 2.56.
func (o GetLimitsV2ResultOutput) MaxPersonalitySize() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxPersonalitySize }).(pulumi.IntOutput)
}

// The number of allowed rules for each security group. Available until version 2.35.
func (o GetLimitsV2ResultOutput) MaxSecurityGroupRules() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxSecurityGroupRules }).(pulumi.IntOutput)
}

// The number of allowed security groups for the tenant. Available until version 2.35.
func (o GetLimitsV2ResultOutput) MaxSecurityGroups() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxSecurityGroups }).(pulumi.IntOutput)
}

// The number of allowed members for each server group.
func (o GetLimitsV2ResultOutput) MaxServerGroupMembers() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxServerGroupMembers }).(pulumi.IntOutput)
}

// The number of allowed server groups for the tenant.
func (o GetLimitsV2ResultOutput) MaxServerGroups() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxServerGroups }).(pulumi.IntOutput)
}

// The number of allowed server groups for the tenant.
func (o GetLimitsV2ResultOutput) MaxServerMeta() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxServerMeta }).(pulumi.IntOutput)
}

// The number of allowed server cores for the tenant.
func (o GetLimitsV2ResultOutput) MaxTotalCores() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxTotalCores }).(pulumi.IntOutput)
}

// The number of allowed floating IP addresses for each tenant. Available until version 2.35.
func (o GetLimitsV2ResultOutput) MaxTotalFloatingIps() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxTotalFloatingIps }).(pulumi.IntOutput)
}

// The number of allowed servers for the tenant.
func (o GetLimitsV2ResultOutput) MaxTotalInstances() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxTotalInstances }).(pulumi.IntOutput)
}

// The number of allowed key pairs for the user.
func (o GetLimitsV2ResultOutput) MaxTotalKeypairs() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxTotalKeypairs }).(pulumi.IntOutput)
}

// The number of allowed floating IP addresses for the tenant. Available until version 2.35.
func (o GetLimitsV2ResultOutput) MaxTotalRamSize() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.MaxTotalRamSize }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o GetLimitsV2ResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLimitsV2Result) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetLimitsV2ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetLimitsV2Result) string { return v.Region }).(pulumi.StringOutput)
}

// The number of used server cores in the tenant.
func (o GetLimitsV2ResultOutput) TotalCoresUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.TotalCoresUsed }).(pulumi.IntOutput)
}

// The number of used floating IP addresses in the tenant.
func (o GetLimitsV2ResultOutput) TotalFloatingIpsUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.TotalFloatingIpsUsed }).(pulumi.IntOutput)
}

// The number of used server cores in the tenant.
func (o GetLimitsV2ResultOutput) TotalInstancesUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.TotalInstancesUsed }).(pulumi.IntOutput)
}

// The amount of used server RAM in the tenant.
func (o GetLimitsV2ResultOutput) TotalRamUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.TotalRamUsed }).(pulumi.IntOutput)
}

// The number of used security groups in the tenant. Available until version 2.35.
func (o GetLimitsV2ResultOutput) TotalSecurityGroupsUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.TotalSecurityGroupsUsed }).(pulumi.IntOutput)
}

// The number of used server groups in each tenant.
func (o GetLimitsV2ResultOutput) TotalServerGroupsUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsV2Result) int { return v.TotalServerGroupsUsed }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLimitsV2ResultOutput{})
}
