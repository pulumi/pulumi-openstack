// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerinfra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack Magnum cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/containerinfra"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := containerinfra.LookupCluster(ctx, &containerinfra.LookupClusterArgs{
// 			Name: "cluster_1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("openstack:containerinfra/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// The name of the cluster.
	Name string `pulumi:"name"`
	// The region in which to obtain the V1 Container Infra
	// client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// COE API address.
	ApiAddress string `pulumi:"apiAddress"`
	// The UUID of the V1 Container Infra cluster template.
	ClusterTemplateId string `pulumi:"clusterTemplateId"`
	// COE software version.
	CoeVersion       string `pulumi:"coeVersion"`
	ContainerVersion string `pulumi:"containerVersion"`
	// The timeout (in minutes) for creating the cluster.
	CreateTimeout int `pulumi:"createTimeout"`
	// The time at which cluster was created.
	CreatedAt string `pulumi:"createdAt"`
	// The URL used for cluster node discovery.
	DiscoveryUrl string `pulumi:"discoveryUrl"`
	// The size (in GB) of the Docker volume.
	DockerVolumeSize int `pulumi:"dockerVolumeSize"`
	// The fixed network that is attached to the cluster.
	FixedNetwork string `pulumi:"fixedNetwork"`
	// The fixed subnet that is attached to the cluster.
	FixedSubnet string `pulumi:"fixedSubnet"`
	// The flavor for the nodes of the cluster.
	Flavor            string `pulumi:"flavor"`
	FloatingIpEnabled bool   `pulumi:"floatingIpEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Compute service SSH keypair.
	Keypair string `pulumi:"keypair"`
	// The list of key value pairs representing additional properties of
	// the cluster.
	Labels map[string]interface{} `pulumi:"labels"`
	// IP addresses of the master node of the cluster.
	MasterAddresses []string `pulumi:"masterAddresses"`
	// The number of master nodes for the cluster.
	MasterCount int `pulumi:"masterCount"`
	// The flavor for the master nodes.
	MasterFlavor string `pulumi:"masterFlavor"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// IP addresses of the node of the cluster.
	NodeAddresses []string `pulumi:"nodeAddresses"`
	// The number of nodes for the cluster.
	NodeCount int `pulumi:"nodeCount"`
	// The project of the cluster.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// UUID of the Orchestration service stack.
	StackId string `pulumi:"stackId"`
	// The time at which cluster was updated.
	UpdatedAt string `pulumi:"updatedAt"`
	// The user of the cluster.
	UserId string `pulumi:"userId"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			return *r, err
		}).(LookupClusterResultOutput)
}

// A collection of arguments for invoking getCluster.
type LookupClusterOutputArgs struct {
	// The name of the cluster.
	Name pulumi.StringInput `pulumi:"name"`
	// The region in which to obtain the V1 Container Infra
	// client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// COE API address.
func (o LookupClusterResultOutput) ApiAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ApiAddress }).(pulumi.StringOutput)
}

// The UUID of the V1 Container Infra cluster template.
func (o LookupClusterResultOutput) ClusterTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterTemplateId }).(pulumi.StringOutput)
}

// COE software version.
func (o LookupClusterResultOutput) CoeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CoeVersion }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ContainerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ContainerVersion }).(pulumi.StringOutput)
}

// The timeout (in minutes) for creating the cluster.
func (o LookupClusterResultOutput) CreateTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.CreateTimeout }).(pulumi.IntOutput)
}

// The time at which cluster was created.
func (o LookupClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The URL used for cluster node discovery.
func (o LookupClusterResultOutput) DiscoveryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.DiscoveryUrl }).(pulumi.StringOutput)
}

// The size (in GB) of the Docker volume.
func (o LookupClusterResultOutput) DockerVolumeSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.DockerVolumeSize }).(pulumi.IntOutput)
}

// The fixed network that is attached to the cluster.
func (o LookupClusterResultOutput) FixedNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.FixedNetwork }).(pulumi.StringOutput)
}

// The fixed subnet that is attached to the cluster.
func (o LookupClusterResultOutput) FixedSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.FixedSubnet }).(pulumi.StringOutput)
}

// The flavor for the nodes of the cluster.
func (o LookupClusterResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Flavor }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) FloatingIpEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupClusterResult) bool { return v.FloatingIpEnabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Compute service SSH keypair.
func (o LookupClusterResultOutput) Keypair() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Keypair }).(pulumi.StringOutput)
}

// The list of key value pairs representing additional properties of
// the cluster.
func (o LookupClusterResultOutput) Labels() pulumi.MapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]interface{} { return v.Labels }).(pulumi.MapOutput)
}

// IP addresses of the master node of the cluster.
func (o LookupClusterResultOutput) MasterAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.MasterAddresses }).(pulumi.StringArrayOutput)
}

// The number of master nodes for the cluster.
func (o LookupClusterResultOutput) MasterCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.MasterCount }).(pulumi.IntOutput)
}

// The flavor for the master nodes.
func (o LookupClusterResultOutput) MasterFlavor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.MasterFlavor }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// IP addresses of the node of the cluster.
func (o LookupClusterResultOutput) NodeAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []string { return v.NodeAddresses }).(pulumi.StringArrayOutput)
}

// The number of nodes for the cluster.
func (o LookupClusterResultOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.NodeCount }).(pulumi.IntOutput)
}

// The project of the cluster.
func (o LookupClusterResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupClusterResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Region }).(pulumi.StringOutput)
}

// UUID of the Orchestration service stack.
func (o LookupClusterResultOutput) StackId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.StackId }).(pulumi.StringOutput)
}

// The time at which cluster was updated.
func (o LookupClusterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The user of the cluster.
func (o LookupClusterResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
