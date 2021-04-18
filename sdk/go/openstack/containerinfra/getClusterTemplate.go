// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerinfra

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack Magnum cluster
// template.
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
// 		_, err := containerinfra.LookupClusterTemplate(ctx, &containerinfra.LookupClusterTemplateArgs{
// 			Name: "clustertemplate_1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupClusterTemplate(ctx *pulumi.Context, args *LookupClusterTemplateArgs, opts ...pulumi.InvokeOption) (*LookupClusterTemplateResult, error) {
	var rv LookupClusterTemplateResult
	err := ctx.Invoke("openstack:containerinfra/getClusterTemplate:getClusterTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterTemplate.
type LookupClusterTemplateArgs struct {
	// The name of the cluster template.
	Name string `pulumi:"name"`
	// The region in which to obtain the V1 Container Infra
	// client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getClusterTemplate.
type LookupClusterTemplateResult struct {
	// The API server port for the Container Orchestration
	// Engine for this cluster template.
	ApiserverPort int `pulumi:"apiserverPort"`
	// The distro for the cluster (fedora-atomic, coreos, etc.).
	ClusterDistro string `pulumi:"clusterDistro"`
	// The Container Orchestration Engine for this cluster template.
	Coe string `pulumi:"coe"`
	// The time at which cluster template was created.
	CreatedAt string `pulumi:"createdAt"`
	// Address of the DNS nameserver that is used in nodes of the
	// cluster.
	DnsNameserver string `pulumi:"dnsNameserver"`
	// Docker storage driver. Changing this updates the
	// Docker storage driver of the existing cluster template.
	DockerStorageDriver string `pulumi:"dockerStorageDriver"`
	// The size (in GB) of the Docker volume.
	DockerVolumeSize int `pulumi:"dockerVolumeSize"`
	// The ID of the external network that will be used for
	// the cluster.
	ExternalNetworkId string `pulumi:"externalNetworkId"`
	// The fixed network that will be attached to the cluster.
	FixedNetwork string `pulumi:"fixedNetwork"`
	// =The fixed subnet that will be attached to the cluster.
	FixedSubnet string `pulumi:"fixedSubnet"`
	// The flavor for the nodes of the cluster.
	Flavor string `pulumi:"flavor"`
	// Indicates whether created cluster should create IP
	// floating IP for every node or not.
	FloatingIpEnabled bool `pulumi:"floatingIpEnabled"`
	// The address of a proxy for receiving all HTTP requests and
	// relay them.
	HttpProxy string `pulumi:"httpProxy"`
	// The address of a proxy for receiving all HTTPS requests and
	// relay them.
	HttpsProxy string `pulumi:"httpsProxy"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The reference to an image that is used for nodes of the cluster.
	Image string `pulumi:"image"`
	// The insecure registry URL for the cluster template.
	InsecureRegistry string `pulumi:"insecureRegistry"`
	// The name of the Compute service SSH keypair.
	KeypairId string `pulumi:"keypairId"`
	// The list of key value pairs representing additional properties
	// of the cluster template.
	Labels map[string]interface{} `pulumi:"labels"`
	// The flavor for the master nodes.
	MasterFlavor string `pulumi:"masterFlavor"`
	// Indicates whether created cluster should has a
	// loadbalancer for master nodes or not.
	MasterLbEnabled bool `pulumi:"masterLbEnabled"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// The name of the driver for the container network.
	NetworkDriver string `pulumi:"networkDriver"`
	// A comma-separated list of IP addresses that shouldn't be used in
	// the cluster.
	NoProxy string `pulumi:"noProxy"`
	// The project of the cluster template.
	ProjectId string `pulumi:"projectId"`
	// Indicates whether cluster template should be public.
	Public bool `pulumi:"public"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// Indicates whether Docker registry is enabled in the
	// cluster.
	RegistryEnabled bool `pulumi:"registryEnabled"`
	// The server type for the cluster template.
	ServerType string `pulumi:"serverType"`
	// Indicates whether the TLS should be disabled in the cluster.
	TlsDisabled bool `pulumi:"tlsDisabled"`
	// The time at which cluster template was updated.
	UpdatedAt string `pulumi:"updatedAt"`
	// The user of the cluster template.
	UserId string `pulumi:"userId"`
	// The name of the driver that is used for the volumes of the
	// cluster nodes.
	VolumeDriver string `pulumi:"volumeDriver"`
}
