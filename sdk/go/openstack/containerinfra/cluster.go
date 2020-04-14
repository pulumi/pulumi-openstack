// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerinfra

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V1 Magnum cluster resource within OpenStack.
//
//
// ## Argument reference
//
// The following arguments are supported:
//
// * `region` - (Optional) The region in which to obtain the V1 Container Infra
//     client. A Container Infra client is needed to create a cluster. If omitted,
//     the `region` argument of the provider is used. Changing this creates a new
//     cluster.
//
// * `name` - (Required) The name of the cluster. Changing this updates the name
//     of the existing cluster template.
//
// * `projectId` - (Optional) The project of the cluster. Required if admin wants
//     to create a cluster in another project. Changing this creates a new
//     cluster.
//
// * `userId` - (Optional) The user of the cluster. Required if admin wants to
//     create a cluster template for another user. Changing this creates a new
//     cluster.
//
// * `clusterTemplateId` - (Required) The UUID of the V1 Container Infra cluster
//     template. Changing this creates a new cluster.
//
// * `createTimeout` - (Optional) The timeout (in minutes) for creating the
//     cluster. Changing this creates a new cluster.
//
// * `discoveryUrl` - (Optional) The URL used for cluster node discovery.
//     Changing this creates a new cluster.
//
// * `dockerVolumeSize` - (Optional) The size (in GB) of the Docker volume.
//     Changing this creates a new cluster.
//
// * `flavor` - (Optional) The flavor for the nodes of the cluster. Can be set via
//     the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
//     cluster.
//
// * `masterFlavor` - (Optional) The flavor for the master nodes. Can be set via
//     the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
//     new cluster.
//
// * `keypair` - (Optional) The name of the Compute service SSH keypair. Changing
//     this creates a new cluster.
//
// * `labels` - (Optional) The list of key value pairs representing additional
//     properties of the cluster. Changing this creates a new cluster.
//
// * `masterCount` - (Optional) The number of master nodes for the cluster.
//     Changing this creates a new cluster.
//
// * `nodeCount` - (Optional) The number of nodes for the cluster. Changing this
//     creates a new cluster.
//
// * `fixedNetwork` - (Optional) The fixed network that will be attached to the
//     cluster. Changing this creates a new cluster.
//
// * `fixedSubnet` - (Optional) The fixed subnet that will be attached to the
//     cluster. Changing this creates a new cluster.
//
// ## Attributes reference
//
// The following attributes are exported:
//
// * `region` - See Argument Reference above.
// * `name` - See Argument Reference above.
// * `projectId` - See Argument Reference above.
// * `createdAt` - The time at which cluster was created.
// * `updatedAt` - The time at which cluster was created.
// * `apiAddress` - COE API address.
// * `coeVersion` - COE software version.
// * `clusterTemplateId` - See Argument Reference above.
// * `containerVersion` - Container software version.
// * `createTimeout` - See Argument Reference above.
// * `discoveryUrl` - See Argument Reference above.
// * `dockerVolumeSize` - See Argument Reference above.
// * `flavor` - See Argument Reference above.
// * `masterFlavor` - See Argument Reference above.
// * `keypair` - See Argument Reference above.
// * `labels` - See Argument Reference above.
// * `masterCount` - See Argument Reference above.
// * `nodeCount` - See Argument Reference above.
// * `fixedNetwork` - See Argument Reference above.
// * `fixedSubnet` - See Argument Reference above.
// * `masterAddresses` - IP addresses of the master node of the cluster.
// * `nodeAddresses` - IP addresses of the node of the cluster.
// * `stackId` - UUID of the Orchestration service stack.
type Cluster struct {
	pulumi.CustomResourceState

	ApiAddress        pulumi.StringOutput `pulumi:"apiAddress"`
	ClusterTemplateId pulumi.StringOutput `pulumi:"clusterTemplateId"`
	CoeVersion        pulumi.StringOutput `pulumi:"coeVersion"`
	ContainerVersion  pulumi.StringOutput `pulumi:"containerVersion"`
	CreateTimeout     pulumi.IntOutput    `pulumi:"createTimeout"`
	CreatedAt         pulumi.StringOutput `pulumi:"createdAt"`
	DiscoveryUrl      pulumi.StringOutput `pulumi:"discoveryUrl"`
	DockerVolumeSize  pulumi.IntOutput    `pulumi:"dockerVolumeSize"`
	FixedNetwork      pulumi.StringOutput `pulumi:"fixedNetwork"`
	FixedSubnet       pulumi.StringOutput `pulumi:"fixedSubnet"`
	Flavor            pulumi.StringOutput `pulumi:"flavor"`
	Keypair           pulumi.StringOutput `pulumi:"keypair"`
	Labels            pulumi.MapOutput    `pulumi:"labels"`
	MasterAddresses   pulumi.StringOutput `pulumi:"masterAddresses"`
	MasterCount       pulumi.IntOutput    `pulumi:"masterCount"`
	MasterFlavor      pulumi.StringOutput `pulumi:"masterFlavor"`
	Name              pulumi.StringOutput `pulumi:"name"`
	NodeAddresses     pulumi.StringOutput `pulumi:"nodeAddresses"`
	NodeCount         pulumi.IntOutput    `pulumi:"nodeCount"`
	ProjectId         pulumi.StringOutput `pulumi:"projectId"`
	Region            pulumi.StringOutput `pulumi:"region"`
	StackId           pulumi.StringOutput `pulumi:"stackId"`
	UpdatedAt         pulumi.StringOutput `pulumi:"updatedAt"`
	UserId            pulumi.StringOutput `pulumi:"userId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.ClusterTemplateId == nil {
		return nil, errors.New("missing required argument 'ClusterTemplateId'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("openstack:containerinfra/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("openstack:containerinfra/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	ApiAddress        *string                `pulumi:"apiAddress"`
	ClusterTemplateId *string                `pulumi:"clusterTemplateId"`
	CoeVersion        *string                `pulumi:"coeVersion"`
	ContainerVersion  *string                `pulumi:"containerVersion"`
	CreateTimeout     *int                   `pulumi:"createTimeout"`
	CreatedAt         *string                `pulumi:"createdAt"`
	DiscoveryUrl      *string                `pulumi:"discoveryUrl"`
	DockerVolumeSize  *int                   `pulumi:"dockerVolumeSize"`
	FixedNetwork      *string                `pulumi:"fixedNetwork"`
	FixedSubnet       *string                `pulumi:"fixedSubnet"`
	Flavor            *string                `pulumi:"flavor"`
	Keypair           *string                `pulumi:"keypair"`
	Labels            map[string]interface{} `pulumi:"labels"`
	MasterAddresses   *string                `pulumi:"masterAddresses"`
	MasterCount       *int                   `pulumi:"masterCount"`
	MasterFlavor      *string                `pulumi:"masterFlavor"`
	Name              *string                `pulumi:"name"`
	NodeAddresses     *string                `pulumi:"nodeAddresses"`
	NodeCount         *int                   `pulumi:"nodeCount"`
	ProjectId         *string                `pulumi:"projectId"`
	Region            *string                `pulumi:"region"`
	StackId           *string                `pulumi:"stackId"`
	UpdatedAt         *string                `pulumi:"updatedAt"`
	UserId            *string                `pulumi:"userId"`
}

type ClusterState struct {
	ApiAddress        pulumi.StringPtrInput
	ClusterTemplateId pulumi.StringPtrInput
	CoeVersion        pulumi.StringPtrInput
	ContainerVersion  pulumi.StringPtrInput
	CreateTimeout     pulumi.IntPtrInput
	CreatedAt         pulumi.StringPtrInput
	DiscoveryUrl      pulumi.StringPtrInput
	DockerVolumeSize  pulumi.IntPtrInput
	FixedNetwork      pulumi.StringPtrInput
	FixedSubnet       pulumi.StringPtrInput
	Flavor            pulumi.StringPtrInput
	Keypair           pulumi.StringPtrInput
	Labels            pulumi.MapInput
	MasterAddresses   pulumi.StringPtrInput
	MasterCount       pulumi.IntPtrInput
	MasterFlavor      pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	NodeAddresses     pulumi.StringPtrInput
	NodeCount         pulumi.IntPtrInput
	ProjectId         pulumi.StringPtrInput
	Region            pulumi.StringPtrInput
	StackId           pulumi.StringPtrInput
	UpdatedAt         pulumi.StringPtrInput
	UserId            pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	ClusterTemplateId string                 `pulumi:"clusterTemplateId"`
	CreateTimeout     *int                   `pulumi:"createTimeout"`
	DiscoveryUrl      *string                `pulumi:"discoveryUrl"`
	DockerVolumeSize  *int                   `pulumi:"dockerVolumeSize"`
	FixedNetwork      *string                `pulumi:"fixedNetwork"`
	FixedSubnet       *string                `pulumi:"fixedSubnet"`
	Flavor            *string                `pulumi:"flavor"`
	Keypair           *string                `pulumi:"keypair"`
	Labels            map[string]interface{} `pulumi:"labels"`
	MasterCount       *int                   `pulumi:"masterCount"`
	MasterFlavor      *string                `pulumi:"masterFlavor"`
	Name              *string                `pulumi:"name"`
	NodeCount         *int                   `pulumi:"nodeCount"`
	Region            *string                `pulumi:"region"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	ClusterTemplateId pulumi.StringInput
	CreateTimeout     pulumi.IntPtrInput
	DiscoveryUrl      pulumi.StringPtrInput
	DockerVolumeSize  pulumi.IntPtrInput
	FixedNetwork      pulumi.StringPtrInput
	FixedSubnet       pulumi.StringPtrInput
	Flavor            pulumi.StringPtrInput
	Keypair           pulumi.StringPtrInput
	Labels            pulumi.MapInput
	MasterCount       pulumi.IntPtrInput
	MasterFlavor      pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	NodeCount         pulumi.IntPtrInput
	Region            pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
