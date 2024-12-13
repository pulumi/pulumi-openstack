// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerinfra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information of an available OpenStack Magnum node group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/containerinfra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containerinfra.LookupNodeGroup(ctx, &containerinfra.LookupNodeGroupArgs{
//				ClusterId: "cluster_1",
//				Name:      pulumi.StringRef("nodegroup_1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNodeGroup(ctx *pulumi.Context, args *LookupNodeGroupArgs, opts ...pulumi.InvokeOption) (*LookupNodeGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNodeGroupResult
	err := ctx.Invoke("openstack:containerinfra/getNodeGroup:getNodeGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNodeGroup.
type LookupNodeGroupArgs struct {
	// The name of the OpenStack Magnum cluster.
	ClusterId string `pulumi:"clusterId"`
	// The name of the node group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V1 Container Infra
	// client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getNodeGroup.
type LookupNodeGroupResult struct {
	ClusterId string `pulumi:"clusterId"`
	// The time at which the node group was created.
	CreatedAt string `pulumi:"createdAt"`
	// The size (in GB) of the Docker volume.
	DockerVolumeSize int `pulumi:"dockerVolumeSize"`
	// The flavor for the nodes of the node group.
	Flavor string `pulumi:"flavor"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The reference to an image that is used for nodes of the node group.
	Image string `pulumi:"image"`
	// The list of key value pairs representing additional properties of
	// the node group.
	Labels map[string]string `pulumi:"labels"`
	// The maximum number of nodes for the node group.
	MaxNodeCount int `pulumi:"maxNodeCount"`
	// The minimum number of nodes for the node group.
	MinNodeCount int `pulumi:"minNodeCount"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// The number of nodes for the node group.
	NodeCount int `pulumi:"nodeCount"`
	// The project of the node group.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The role of the node group.
	Role string `pulumi:"role"`
	// The time at which the node group was updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupNodeGroupOutput(ctx *pulumi.Context, args LookupNodeGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNodeGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNodeGroupResultOutput, error) {
			args := v.(LookupNodeGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:containerinfra/getNodeGroup:getNodeGroup", args, LookupNodeGroupResultOutput{}, options).(LookupNodeGroupResultOutput), nil
		}).(LookupNodeGroupResultOutput)
}

// A collection of arguments for invoking getNodeGroup.
type LookupNodeGroupOutputArgs struct {
	// The name of the OpenStack Magnum cluster.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The name of the node group.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V1 Container Infra
	// client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupNodeGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeGroupArgs)(nil)).Elem()
}

// A collection of values returned by getNodeGroup.
type LookupNodeGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNodeGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeGroupResult)(nil)).Elem()
}

func (o LookupNodeGroupResultOutput) ToLookupNodeGroupResultOutput() LookupNodeGroupResultOutput {
	return o
}

func (o LookupNodeGroupResultOutput) ToLookupNodeGroupResultOutputWithContext(ctx context.Context) LookupNodeGroupResultOutput {
	return o
}

func (o LookupNodeGroupResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The time at which the node group was created.
func (o LookupNodeGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The size (in GB) of the Docker volume.
func (o LookupNodeGroupResultOutput) DockerVolumeSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) int { return v.DockerVolumeSize }).(pulumi.IntOutput)
}

// The flavor for the nodes of the node group.
func (o LookupNodeGroupResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.Flavor }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNodeGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The reference to an image that is used for nodes of the node group.
func (o LookupNodeGroupResultOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.Image }).(pulumi.StringOutput)
}

// The list of key value pairs representing additional properties of
// the node group.
func (o LookupNodeGroupResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The maximum number of nodes for the node group.
func (o LookupNodeGroupResultOutput) MaxNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) int { return v.MaxNodeCount }).(pulumi.IntOutput)
}

// The minimum number of nodes for the node group.
func (o LookupNodeGroupResultOutput) MinNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) int { return v.MinNodeCount }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupNodeGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of nodes for the node group.
func (o LookupNodeGroupResultOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) int { return v.NodeCount }).(pulumi.IntOutput)
}

// The project of the node group.
func (o LookupNodeGroupResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupNodeGroupResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.Region }).(pulumi.StringOutput)
}

// The role of the node group.
func (o LookupNodeGroupResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.Role }).(pulumi.StringOutput)
}

// The time at which the node group was updated.
func (o LookupNodeGroupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeGroupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNodeGroupResultOutput{})
}
