// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.ContainerInfra
{
    public static class GetCluster
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack Magnum cluster.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var cluster1 = OpenStack.ContainerInfra.GetCluster.Invoke(new()
        ///     {
        ///         Name = "cluster_1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterResult> InvokeAsync(GetClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterResult>("openstack:containerinfra/getCluster:getCluster", args ?? new GetClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack Magnum cluster.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var cluster1 = OpenStack.ContainerInfra.GetCluster.Invoke(new()
        ///     {
        ///         Name = "cluster_1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetClusterResult> Invoke(GetClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterResult>("openstack:containerinfra/getCluster:getCluster", args ?? new GetClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V1 Container Infra
        /// client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetClusterArgs()
        {
        }
        public static new GetClusterArgs Empty => new GetClusterArgs();
    }

    public sealed class GetClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V1 Container Infra
        /// client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetClusterInvokeArgs()
        {
        }
        public static new GetClusterInvokeArgs Empty => new GetClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterResult
    {
        /// <summary>
        /// COE API address.
        /// </summary>
        public readonly string ApiAddress;
        /// <summary>
        /// The UUID of the V1 Container Infra cluster template.
        /// </summary>
        public readonly string ClusterTemplateId;
        /// <summary>
        /// COE software version.
        /// </summary>
        public readonly string CoeVersion;
        public readonly string ContainerVersion;
        /// <summary>
        /// The timeout (in minutes) for creating the cluster.
        /// </summary>
        public readonly int CreateTimeout;
        /// <summary>
        /// The time at which cluster was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The URL used for cluster node discovery.
        /// </summary>
        public readonly string DiscoveryUrl;
        /// <summary>
        /// The size (in GB) of the Docker volume.
        /// </summary>
        public readonly int DockerVolumeSize;
        /// <summary>
        /// The fixed network that is attached to the cluster.
        /// </summary>
        public readonly string FixedNetwork;
        /// <summary>
        /// The fixed subnet that is attached to the cluster.
        /// </summary>
        public readonly string FixedSubnet;
        /// <summary>
        /// The flavor for the nodes of the cluster.
        /// </summary>
        public readonly string Flavor;
        public readonly bool FloatingIpEnabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the Compute service SSH keypair.
        /// </summary>
        public readonly string Keypair;
        /// <summary>
        /// The Kubernetes cluster's credentials
        /// </summary>
        public readonly ImmutableDictionary<string, string> Kubeconfig;
        /// <summary>
        /// The list of key value pairs representing additional properties of
        /// the cluster.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        /// <summary>
        /// IP addresses of the master node of the cluster.
        /// </summary>
        public readonly ImmutableArray<string> MasterAddresses;
        /// <summary>
        /// The number of master nodes for the cluster.
        /// </summary>
        public readonly int MasterCount;
        /// <summary>
        /// The flavor for the master nodes.
        /// </summary>
        public readonly string MasterFlavor;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IP addresses of the node of the cluster.
        /// </summary>
        public readonly ImmutableArray<string> NodeAddresses;
        /// <summary>
        /// The number of nodes for the cluster.
        /// </summary>
        public readonly int NodeCount;
        /// <summary>
        /// The project of the cluster.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// UUID of the Orchestration service stack.
        /// </summary>
        public readonly string StackId;
        /// <summary>
        /// The time at which cluster was updated.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The user of the cluster.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private GetClusterResult(
            string apiAddress,

            string clusterTemplateId,

            string coeVersion,

            string containerVersion,

            int createTimeout,

            string createdAt,

            string discoveryUrl,

            int dockerVolumeSize,

            string fixedNetwork,

            string fixedSubnet,

            string flavor,

            bool floatingIpEnabled,

            string id,

            string keypair,

            ImmutableDictionary<string, string> kubeconfig,

            ImmutableDictionary<string, object> labels,

            ImmutableArray<string> masterAddresses,

            int masterCount,

            string masterFlavor,

            string name,

            ImmutableArray<string> nodeAddresses,

            int nodeCount,

            string projectId,

            string region,

            string stackId,

            string updatedAt,

            string userId)
        {
            ApiAddress = apiAddress;
            ClusterTemplateId = clusterTemplateId;
            CoeVersion = coeVersion;
            ContainerVersion = containerVersion;
            CreateTimeout = createTimeout;
            CreatedAt = createdAt;
            DiscoveryUrl = discoveryUrl;
            DockerVolumeSize = dockerVolumeSize;
            FixedNetwork = fixedNetwork;
            FixedSubnet = fixedSubnet;
            Flavor = flavor;
            FloatingIpEnabled = floatingIpEnabled;
            Id = id;
            Keypair = keypair;
            Kubeconfig = kubeconfig;
            Labels = labels;
            MasterAddresses = masterAddresses;
            MasterCount = masterCount;
            MasterFlavor = masterFlavor;
            Name = name;
            NodeAddresses = nodeAddresses;
            NodeCount = nodeCount;
            ProjectId = projectId;
            Region = region;
            StackId = stackId;
            UpdatedAt = updatedAt;
            UserId = userId;
        }
    }
}
