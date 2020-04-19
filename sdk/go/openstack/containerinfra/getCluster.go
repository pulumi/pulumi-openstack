// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerinfra

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack Magnum cluster.
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
	Flavor string `pulumi:"flavor"`
	// id is the provider-assigned unique ID for this managed resource.
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
