// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerinfra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Create a Cluster
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
// 		_, err := containerinfra.NewCluster(ctx, "cluster1", &containerinfra.ClusterArgs{
// 			ClusterTemplateId: pulumi.String("b9a45c5c-cd03-4958-82aa-b80bf93cb922"),
// 			Keypair:           pulumi.String("ssh_keypair"),
// 			MasterCount:       pulumi.Int(3),
// 			NodeCount:         pulumi.Int(5),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
// * `mergeLabels` - See Argument Reference above.
// * `masterCount` - See Argument Reference above.
// * `nodeCount` - See Argument Reference above.
// * `fixedNetwork` - See Argument Reference above.
// * `fixedSubnet` - See Argument Reference above.
// * `floatingIpEnabled` - See Argument Reference above.
// * `masterAddresses` - IP addresses of the master node of the cluster.
// * `nodeAddresses` - IP addresses of the node of the cluster.
// * `stackId` - UUID of the Orchestration service stack.
// * `kubeconfig` - The Kubernetes cluster's credentials
//   * `rawConfig` - The raw kubeconfig file
//   * `host` - The cluster's API server URL
//   * `clusterCaCertificate` - The cluster's CA certificate
//   * `clientKey` - The client's RSA key
//   * `clientCertificate` - The client's certificate
//
// ## Import
//
// Clusters can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:containerinfra/cluster:Cluster cluster_1 ce0f9463-dd25-474b-9fe8-94de63e5e42b
// ```
type Cluster struct {
	pulumi.CustomResourceState

	ApiAddress pulumi.StringOutput `pulumi:"apiAddress"`
	// The UUID of the V1 Container Infra cluster
	// template. Changing this creates a new cluster.
	ClusterTemplateId pulumi.StringOutput `pulumi:"clusterTemplateId"`
	CoeVersion        pulumi.StringOutput `pulumi:"coeVersion"`
	ContainerVersion  pulumi.StringOutput `pulumi:"containerVersion"`
	// The timeout (in minutes) for creating the
	// cluster. Changing this creates a new cluster.
	CreateTimeout pulumi.IntOutput    `pulumi:"createTimeout"`
	CreatedAt     pulumi.StringOutput `pulumi:"createdAt"`
	// The URL used for cluster node discovery.
	// Changing this creates a new cluster.
	DiscoveryUrl pulumi.StringOutput `pulumi:"discoveryUrl"`
	// The size (in GB) of the Docker volume.
	// Changing this creates a new cluster.
	DockerVolumeSize pulumi.IntOutput `pulumi:"dockerVolumeSize"`
	// The fixed network that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedNetwork pulumi.StringOutput `pulumi:"fixedNetwork"`
	// The fixed subnet that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedSubnet pulumi.StringOutput `pulumi:"fixedSubnet"`
	// The flavor for the nodes of the cluster. Can be set via
	// the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// cluster.
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// Indicates whether floating IP should be
	// created for every cluster node. Changing this creates a new cluster.
	FloatingIpEnabled pulumi.BoolOutput `pulumi:"floatingIpEnabled"`
	// The name of the Compute service SSH keypair. Changing
	// this creates a new cluster.
	Keypair    pulumi.StringOutput    `pulumi:"keypair"`
	Kubeconfig pulumi.StringMapOutput `pulumi:"kubeconfig"`
	// The list of key value pairs representing additional
	// properties of the cluster. Changing this creates a new cluster.
	Labels          pulumi.MapOutput         `pulumi:"labels"`
	MasterAddresses pulumi.StringArrayOutput `pulumi:"masterAddresses"`
	// The number of master nodes for the cluster.
	// Changing this creates a new cluster.
	MasterCount pulumi.IntOutput `pulumi:"masterCount"`
	// The flavor for the master nodes. Can be set via
	// the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
	// new cluster.
	MasterFlavor pulumi.StringOutput `pulumi:"masterFlavor"`
	// Indicates whether the provided labels should be
	// merged with cluster template labels. Changing this creates a new cluster.
	MergeLabels pulumi.BoolPtrOutput `pulumi:"mergeLabels"`
	// The name of the cluster. Changing this updates the name
	// of the existing cluster template.
	Name          pulumi.StringOutput      `pulumi:"name"`
	NodeAddresses pulumi.StringArrayOutput `pulumi:"nodeAddresses"`
	// The number of nodes for the cluster. Changing this
	// creates a new cluster.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// The project of the cluster. Required if admin wants
	// to create a cluster in another project. Changing this creates a new
	// cluster.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// cluster.
	Region    pulumi.StringOutput `pulumi:"region"`
	StackId   pulumi.StringOutput `pulumi:"stackId"`
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The user of the cluster. Required if admin wants to
	// create a cluster template for another user. Changing this creates a new
	// cluster.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterTemplateId'")
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
	ApiAddress *string `pulumi:"apiAddress"`
	// The UUID of the V1 Container Infra cluster
	// template. Changing this creates a new cluster.
	ClusterTemplateId *string `pulumi:"clusterTemplateId"`
	CoeVersion        *string `pulumi:"coeVersion"`
	ContainerVersion  *string `pulumi:"containerVersion"`
	// The timeout (in minutes) for creating the
	// cluster. Changing this creates a new cluster.
	CreateTimeout *int    `pulumi:"createTimeout"`
	CreatedAt     *string `pulumi:"createdAt"`
	// The URL used for cluster node discovery.
	// Changing this creates a new cluster.
	DiscoveryUrl *string `pulumi:"discoveryUrl"`
	// The size (in GB) of the Docker volume.
	// Changing this creates a new cluster.
	DockerVolumeSize *int `pulumi:"dockerVolumeSize"`
	// The fixed network that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedNetwork *string `pulumi:"fixedNetwork"`
	// The fixed subnet that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedSubnet *string `pulumi:"fixedSubnet"`
	// The flavor for the nodes of the cluster. Can be set via
	// the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// cluster.
	Flavor *string `pulumi:"flavor"`
	// Indicates whether floating IP should be
	// created for every cluster node. Changing this creates a new cluster.
	FloatingIpEnabled *bool `pulumi:"floatingIpEnabled"`
	// The name of the Compute service SSH keypair. Changing
	// this creates a new cluster.
	Keypair    *string           `pulumi:"keypair"`
	Kubeconfig map[string]string `pulumi:"kubeconfig"`
	// The list of key value pairs representing additional
	// properties of the cluster. Changing this creates a new cluster.
	Labels          map[string]interface{} `pulumi:"labels"`
	MasterAddresses []string               `pulumi:"masterAddresses"`
	// The number of master nodes for the cluster.
	// Changing this creates a new cluster.
	MasterCount *int `pulumi:"masterCount"`
	// The flavor for the master nodes. Can be set via
	// the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
	// new cluster.
	MasterFlavor *string `pulumi:"masterFlavor"`
	// Indicates whether the provided labels should be
	// merged with cluster template labels. Changing this creates a new cluster.
	MergeLabels *bool `pulumi:"mergeLabels"`
	// The name of the cluster. Changing this updates the name
	// of the existing cluster template.
	Name          *string  `pulumi:"name"`
	NodeAddresses []string `pulumi:"nodeAddresses"`
	// The number of nodes for the cluster. Changing this
	// creates a new cluster.
	NodeCount *int `pulumi:"nodeCount"`
	// The project of the cluster. Required if admin wants
	// to create a cluster in another project. Changing this creates a new
	// cluster.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// cluster.
	Region    *string `pulumi:"region"`
	StackId   *string `pulumi:"stackId"`
	UpdatedAt *string `pulumi:"updatedAt"`
	// The user of the cluster. Required if admin wants to
	// create a cluster template for another user. Changing this creates a new
	// cluster.
	UserId *string `pulumi:"userId"`
}

type ClusterState struct {
	ApiAddress pulumi.StringPtrInput
	// The UUID of the V1 Container Infra cluster
	// template. Changing this creates a new cluster.
	ClusterTemplateId pulumi.StringPtrInput
	CoeVersion        pulumi.StringPtrInput
	ContainerVersion  pulumi.StringPtrInput
	// The timeout (in minutes) for creating the
	// cluster. Changing this creates a new cluster.
	CreateTimeout pulumi.IntPtrInput
	CreatedAt     pulumi.StringPtrInput
	// The URL used for cluster node discovery.
	// Changing this creates a new cluster.
	DiscoveryUrl pulumi.StringPtrInput
	// The size (in GB) of the Docker volume.
	// Changing this creates a new cluster.
	DockerVolumeSize pulumi.IntPtrInput
	// The fixed network that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedNetwork pulumi.StringPtrInput
	// The fixed subnet that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedSubnet pulumi.StringPtrInput
	// The flavor for the nodes of the cluster. Can be set via
	// the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// cluster.
	Flavor pulumi.StringPtrInput
	// Indicates whether floating IP should be
	// created for every cluster node. Changing this creates a new cluster.
	FloatingIpEnabled pulumi.BoolPtrInput
	// The name of the Compute service SSH keypair. Changing
	// this creates a new cluster.
	Keypair    pulumi.StringPtrInput
	Kubeconfig pulumi.StringMapInput
	// The list of key value pairs representing additional
	// properties of the cluster. Changing this creates a new cluster.
	Labels          pulumi.MapInput
	MasterAddresses pulumi.StringArrayInput
	// The number of master nodes for the cluster.
	// Changing this creates a new cluster.
	MasterCount pulumi.IntPtrInput
	// The flavor for the master nodes. Can be set via
	// the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
	// new cluster.
	MasterFlavor pulumi.StringPtrInput
	// Indicates whether the provided labels should be
	// merged with cluster template labels. Changing this creates a new cluster.
	MergeLabels pulumi.BoolPtrInput
	// The name of the cluster. Changing this updates the name
	// of the existing cluster template.
	Name          pulumi.StringPtrInput
	NodeAddresses pulumi.StringArrayInput
	// The number of nodes for the cluster. Changing this
	// creates a new cluster.
	NodeCount pulumi.IntPtrInput
	// The project of the cluster. Required if admin wants
	// to create a cluster in another project. Changing this creates a new
	// cluster.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// cluster.
	Region    pulumi.StringPtrInput
	StackId   pulumi.StringPtrInput
	UpdatedAt pulumi.StringPtrInput
	// The user of the cluster. Required if admin wants to
	// create a cluster template for another user. Changing this creates a new
	// cluster.
	UserId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The UUID of the V1 Container Infra cluster
	// template. Changing this creates a new cluster.
	ClusterTemplateId string `pulumi:"clusterTemplateId"`
	// The timeout (in minutes) for creating the
	// cluster. Changing this creates a new cluster.
	CreateTimeout *int `pulumi:"createTimeout"`
	// The URL used for cluster node discovery.
	// Changing this creates a new cluster.
	DiscoveryUrl *string `pulumi:"discoveryUrl"`
	// The size (in GB) of the Docker volume.
	// Changing this creates a new cluster.
	DockerVolumeSize *int `pulumi:"dockerVolumeSize"`
	// The fixed network that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedNetwork *string `pulumi:"fixedNetwork"`
	// The fixed subnet that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedSubnet *string `pulumi:"fixedSubnet"`
	// The flavor for the nodes of the cluster. Can be set via
	// the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// cluster.
	Flavor *string `pulumi:"flavor"`
	// Indicates whether floating IP should be
	// created for every cluster node. Changing this creates a new cluster.
	FloatingIpEnabled *bool `pulumi:"floatingIpEnabled"`
	// The name of the Compute service SSH keypair. Changing
	// this creates a new cluster.
	Keypair *string `pulumi:"keypair"`
	// The list of key value pairs representing additional
	// properties of the cluster. Changing this creates a new cluster.
	Labels map[string]interface{} `pulumi:"labels"`
	// The number of master nodes for the cluster.
	// Changing this creates a new cluster.
	MasterCount *int `pulumi:"masterCount"`
	// The flavor for the master nodes. Can be set via
	// the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
	// new cluster.
	MasterFlavor *string `pulumi:"masterFlavor"`
	// Indicates whether the provided labels should be
	// merged with cluster template labels. Changing this creates a new cluster.
	MergeLabels *bool `pulumi:"mergeLabels"`
	// The name of the cluster. Changing this updates the name
	// of the existing cluster template.
	Name *string `pulumi:"name"`
	// The number of nodes for the cluster. Changing this
	// creates a new cluster.
	NodeCount *int `pulumi:"nodeCount"`
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// cluster.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The UUID of the V1 Container Infra cluster
	// template. Changing this creates a new cluster.
	ClusterTemplateId pulumi.StringInput
	// The timeout (in minutes) for creating the
	// cluster. Changing this creates a new cluster.
	CreateTimeout pulumi.IntPtrInput
	// The URL used for cluster node discovery.
	// Changing this creates a new cluster.
	DiscoveryUrl pulumi.StringPtrInput
	// The size (in GB) of the Docker volume.
	// Changing this creates a new cluster.
	DockerVolumeSize pulumi.IntPtrInput
	// The fixed network that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedNetwork pulumi.StringPtrInput
	// The fixed subnet that will be attached to the
	// cluster. Changing this creates a new cluster.
	FixedSubnet pulumi.StringPtrInput
	// The flavor for the nodes of the cluster. Can be set via
	// the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// cluster.
	Flavor pulumi.StringPtrInput
	// Indicates whether floating IP should be
	// created for every cluster node. Changing this creates a new cluster.
	FloatingIpEnabled pulumi.BoolPtrInput
	// The name of the Compute service SSH keypair. Changing
	// this creates a new cluster.
	Keypair pulumi.StringPtrInput
	// The list of key value pairs representing additional
	// properties of the cluster. Changing this creates a new cluster.
	Labels pulumi.MapInput
	// The number of master nodes for the cluster.
	// Changing this creates a new cluster.
	MasterCount pulumi.IntPtrInput
	// The flavor for the master nodes. Can be set via
	// the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
	// new cluster.
	MasterFlavor pulumi.StringPtrInput
	// Indicates whether the provided labels should be
	// merged with cluster template labels. Changing this creates a new cluster.
	MergeLabels pulumi.BoolPtrInput
	// The name of the cluster. Changing this updates the name
	// of the existing cluster template.
	Name pulumi.StringPtrInput
	// The number of nodes for the cluster. Changing this
	// creates a new cluster.
	NodeCount pulumi.IntPtrInput
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// cluster.
	Region pulumi.StringPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//          ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//          ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
