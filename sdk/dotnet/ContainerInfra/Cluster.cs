// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.ContainerInfra
{
    /// <summary>
    /// Manages a V1 Magnum cluster resource within OpenStack.
    /// 
    /// &gt; **Note:** All arguments including the `kubeconfig` computed attribute will be
    /// stored in the raw state as plain-text. Read more about sensitive data in
    /// state.
    /// 
    /// ## Example Usage
    /// ### Create a Cluster
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cluster1 = new OpenStack.ContainerInfra.Cluster("cluster1", new()
    ///     {
    ///         ClusterTemplateId = "b9a45c5c-cd03-4958-82aa-b80bf93cb922",
    ///         Keypair = "ssh_keypair",
    ///         MasterCount = 3,
    ///         NodeCount = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// ## Attributes reference
    /// 
    /// The following attributes are exported:
    /// 
    /// * `region` - See Argument Reference above.
    /// * `name` - See Argument Reference above.
    /// * `project_id` - See Argument Reference above.
    /// * `created_at` - The time at which cluster was created.
    /// * `updated_at` - The time at which cluster was created.
    /// * `api_address` - COE API address.
    /// * `coe_version` - COE software version.
    /// * `cluster_template_id` - See Argument Reference above.
    /// * `container_version` - Container software version.
    /// * `create_timeout` - See Argument Reference above.
    /// * `discovery_url` - See Argument Reference above.
    /// * `docker_volume_size` - See Argument Reference above.
    /// * `flavor` - See Argument Reference above.
    /// * `master_flavor` - See Argument Reference above.
    /// * `keypair` - See Argument Reference above.
    /// * `labels` - See Argument Reference above.
    /// * `merge_labels` - See Argument Reference above.
    /// * `master_count` - See Argument Reference above.
    /// * `node_count` - See Argument Reference above.
    /// * `fixed_network` - See Argument Reference above.
    /// * `fixed_subnet` - See Argument Reference above.
    /// * `floating_ip_enabled` - See Argument Reference above.
    /// * `master_addresses` - IP addresses of the master node of the cluster.
    /// * `node_addresses` - IP addresses of the node of the cluster.
    /// * `stack_id` - UUID of the Orchestration service stack.
    /// * `kubeconfig` - The Kubernetes cluster's credentials
    ///   * `raw_config` - The raw kubeconfig file
    ///   * `host` - The cluster's API server URL
    ///   * `cluster_ca_certificate` - The cluster's CA certificate
    ///   * `client_key` - The client's RSA key
    ///   * `client_certificate` - The client's certificate
    /// 
    /// ## Import
    /// 
    /// Clusters can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:containerinfra/cluster:Cluster cluster_1 ce0f9463-dd25-474b-9fe8-94de63e5e42b
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:containerinfra/cluster:Cluster")]
    public partial class Cluster : global::Pulumi.CustomResource
    {
        [Output("apiAddress")]
        public Output<string> ApiAddress { get; private set; } = null!;

        /// <summary>
        /// The UUID of the V1 Container Infra cluster
        /// template. Changing this creates a new cluster.
        /// </summary>
        [Output("clusterTemplateId")]
        public Output<string> ClusterTemplateId { get; private set; } = null!;

        [Output("coeVersion")]
        public Output<string> CoeVersion { get; private set; } = null!;

        [Output("containerVersion")]
        public Output<string> ContainerVersion { get; private set; } = null!;

        /// <summary>
        /// The timeout (in minutes) for creating the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Output("createTimeout")]
        public Output<int> CreateTimeout { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The URL used for cluster node discovery.
        /// Changing this creates a new cluster.
        /// </summary>
        [Output("discoveryUrl")]
        public Output<string> DiscoveryUrl { get; private set; } = null!;

        /// <summary>
        /// The size (in GB) of the Docker volume.
        /// Changing this creates a new cluster.
        /// </summary>
        [Output("dockerVolumeSize")]
        public Output<int> DockerVolumeSize { get; private set; } = null!;

        /// <summary>
        /// The fixed network that will be attached to the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Output("fixedNetwork")]
        public Output<string> FixedNetwork { get; private set; } = null!;

        /// <summary>
        /// The fixed subnet that will be attached to the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Output("fixedSubnet")]
        public Output<string> FixedSubnet { get; private set; } = null!;

        /// <summary>
        /// The flavor for the nodes of the cluster. Can be set via
        /// the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
        /// cluster.
        /// </summary>
        [Output("flavor")]
        public Output<string> Flavor { get; private set; } = null!;

        /// <summary>
        /// Indicates whether floating IP should be
        /// created for every cluster node. Changing this creates a new cluster.
        /// </summary>
        [Output("floatingIpEnabled")]
        public Output<bool> FloatingIpEnabled { get; private set; } = null!;

        /// <summary>
        /// The name of the Compute service SSH keypair. Changing
        /// this creates a new cluster.
        /// </summary>
        [Output("keypair")]
        public Output<string> Keypair { get; private set; } = null!;

        [Output("kubeconfig")]
        public Output<ImmutableDictionary<string, string>> Kubeconfig { get; private set; } = null!;

        /// <summary>
        /// The list of key value pairs representing additional
        /// properties of the cluster. Changing this creates a new cluster.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, object>> Labels { get; private set; } = null!;

        [Output("masterAddresses")]
        public Output<ImmutableArray<string>> MasterAddresses { get; private set; } = null!;

        /// <summary>
        /// The number of master nodes for the cluster.
        /// Changing this creates a new cluster.
        /// </summary>
        [Output("masterCount")]
        public Output<int> MasterCount { get; private set; } = null!;

        /// <summary>
        /// The flavor for the master nodes. Can be set via
        /// the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
        /// new cluster.
        /// </summary>
        [Output("masterFlavor")]
        public Output<string> MasterFlavor { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the provided labels should be
        /// merged with cluster template labels. Changing this creates a new cluster.
        /// </summary>
        [Output("mergeLabels")]
        public Output<bool?> MergeLabels { get; private set; } = null!;

        /// <summary>
        /// The name of the cluster. Changing this creates a new
        /// cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nodeAddresses")]
        public Output<ImmutableArray<string>> NodeAddresses { get; private set; } = null!;

        /// <summary>
        /// The number of nodes for the cluster.
        /// </summary>
        [Output("nodeCount")]
        public Output<int?> NodeCount { get; private set; } = null!;

        /// <summary>
        /// The project of the cluster. Required if admin wants
        /// to create a cluster in another project. Changing this creates a new
        /// cluster.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V1 Container Infra
        /// client. A Container Infra client is needed to create a cluster. If omitted,
        /// the `region` argument of the provider is used. Changing this creates a new
        /// cluster.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The user of the cluster. Required if admin wants to
        /// create a cluster template for another user. Changing this creates a new
        /// cluster.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a Cluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cluster(string name, ClusterArgs args, CustomResourceOptions? options = null)
            : base("openstack:containerinfra/cluster:Cluster", name, args ?? new ClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cluster(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
            : base("openstack:containerinfra/cluster:Cluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "kubeconfig",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cluster Get(string name, Input<string> id, ClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new Cluster(name, id, state, options);
        }
    }

    public sealed class ClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of the V1 Container Infra cluster
        /// template. Changing this creates a new cluster.
        /// </summary>
        [Input("clusterTemplateId", required: true)]
        public Input<string> ClusterTemplateId { get; set; } = null!;

        /// <summary>
        /// The timeout (in minutes) for creating the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Input("createTimeout")]
        public Input<int>? CreateTimeout { get; set; }

        /// <summary>
        /// The URL used for cluster node discovery.
        /// Changing this creates a new cluster.
        /// </summary>
        [Input("discoveryUrl")]
        public Input<string>? DiscoveryUrl { get; set; }

        /// <summary>
        /// The size (in GB) of the Docker volume.
        /// Changing this creates a new cluster.
        /// </summary>
        [Input("dockerVolumeSize")]
        public Input<int>? DockerVolumeSize { get; set; }

        /// <summary>
        /// The fixed network that will be attached to the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Input("fixedNetwork")]
        public Input<string>? FixedNetwork { get; set; }

        /// <summary>
        /// The fixed subnet that will be attached to the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Input("fixedSubnet")]
        public Input<string>? FixedSubnet { get; set; }

        /// <summary>
        /// The flavor for the nodes of the cluster. Can be set via
        /// the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
        /// cluster.
        /// </summary>
        [Input("flavor")]
        public Input<string>? Flavor { get; set; }

        /// <summary>
        /// Indicates whether floating IP should be
        /// created for every cluster node. Changing this creates a new cluster.
        /// </summary>
        [Input("floatingIpEnabled")]
        public Input<bool>? FloatingIpEnabled { get; set; }

        /// <summary>
        /// The name of the Compute service SSH keypair. Changing
        /// this creates a new cluster.
        /// </summary>
        [Input("keypair")]
        public Input<string>? Keypair { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// The list of key value pairs representing additional
        /// properties of the cluster. Changing this creates a new cluster.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        /// <summary>
        /// The number of master nodes for the cluster.
        /// Changing this creates a new cluster.
        /// </summary>
        [Input("masterCount")]
        public Input<int>? MasterCount { get; set; }

        /// <summary>
        /// The flavor for the master nodes. Can be set via
        /// the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
        /// new cluster.
        /// </summary>
        [Input("masterFlavor")]
        public Input<string>? MasterFlavor { get; set; }

        /// <summary>
        /// Indicates whether the provided labels should be
        /// merged with cluster template labels. Changing this creates a new cluster.
        /// </summary>
        [Input("mergeLabels")]
        public Input<bool>? MergeLabels { get; set; }

        /// <summary>
        /// The name of the cluster. Changing this creates a new
        /// cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of nodes for the cluster.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 Container Infra
        /// client. A Container Infra client is needed to create a cluster. If omitted,
        /// the `region` argument of the provider is used. Changing this creates a new
        /// cluster.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ClusterArgs()
        {
        }
        public static new ClusterArgs Empty => new ClusterArgs();
    }

    public sealed class ClusterState : global::Pulumi.ResourceArgs
    {
        [Input("apiAddress")]
        public Input<string>? ApiAddress { get; set; }

        /// <summary>
        /// The UUID of the V1 Container Infra cluster
        /// template. Changing this creates a new cluster.
        /// </summary>
        [Input("clusterTemplateId")]
        public Input<string>? ClusterTemplateId { get; set; }

        [Input("coeVersion")]
        public Input<string>? CoeVersion { get; set; }

        [Input("containerVersion")]
        public Input<string>? ContainerVersion { get; set; }

        /// <summary>
        /// The timeout (in minutes) for creating the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Input("createTimeout")]
        public Input<int>? CreateTimeout { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The URL used for cluster node discovery.
        /// Changing this creates a new cluster.
        /// </summary>
        [Input("discoveryUrl")]
        public Input<string>? DiscoveryUrl { get; set; }

        /// <summary>
        /// The size (in GB) of the Docker volume.
        /// Changing this creates a new cluster.
        /// </summary>
        [Input("dockerVolumeSize")]
        public Input<int>? DockerVolumeSize { get; set; }

        /// <summary>
        /// The fixed network that will be attached to the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Input("fixedNetwork")]
        public Input<string>? FixedNetwork { get; set; }

        /// <summary>
        /// The fixed subnet that will be attached to the
        /// cluster. Changing this creates a new cluster.
        /// </summary>
        [Input("fixedSubnet")]
        public Input<string>? FixedSubnet { get; set; }

        /// <summary>
        /// The flavor for the nodes of the cluster. Can be set via
        /// the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
        /// cluster.
        /// </summary>
        [Input("flavor")]
        public Input<string>? Flavor { get; set; }

        /// <summary>
        /// Indicates whether floating IP should be
        /// created for every cluster node. Changing this creates a new cluster.
        /// </summary>
        [Input("floatingIpEnabled")]
        public Input<bool>? FloatingIpEnabled { get; set; }

        /// <summary>
        /// The name of the Compute service SSH keypair. Changing
        /// this creates a new cluster.
        /// </summary>
        [Input("keypair")]
        public Input<string>? Keypair { get; set; }

        [Input("kubeconfig")]
        private InputMap<string>? _kubeconfig;
        public InputMap<string> Kubeconfig
        {
            get => _kubeconfig ?? (_kubeconfig = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _kubeconfig = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("labels")]
        private InputMap<object>? _labels;

        /// <summary>
        /// The list of key value pairs representing additional
        /// properties of the cluster. Changing this creates a new cluster.
        /// </summary>
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        [Input("masterAddresses")]
        private InputList<string>? _masterAddresses;
        public InputList<string> MasterAddresses
        {
            get => _masterAddresses ?? (_masterAddresses = new InputList<string>());
            set => _masterAddresses = value;
        }

        /// <summary>
        /// The number of master nodes for the cluster.
        /// Changing this creates a new cluster.
        /// </summary>
        [Input("masterCount")]
        public Input<int>? MasterCount { get; set; }

        /// <summary>
        /// The flavor for the master nodes. Can be set via
        /// the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
        /// new cluster.
        /// </summary>
        [Input("masterFlavor")]
        public Input<string>? MasterFlavor { get; set; }

        /// <summary>
        /// Indicates whether the provided labels should be
        /// merged with cluster template labels. Changing this creates a new cluster.
        /// </summary>
        [Input("mergeLabels")]
        public Input<bool>? MergeLabels { get; set; }

        /// <summary>
        /// The name of the cluster. Changing this creates a new
        /// cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeAddresses")]
        private InputList<string>? _nodeAddresses;
        public InputList<string> NodeAddresses
        {
            get => _nodeAddresses ?? (_nodeAddresses = new InputList<string>());
            set => _nodeAddresses = value;
        }

        /// <summary>
        /// The number of nodes for the cluster.
        /// </summary>
        [Input("nodeCount")]
        public Input<int>? NodeCount { get; set; }

        /// <summary>
        /// The project of the cluster. Required if admin wants
        /// to create a cluster in another project. Changing this creates a new
        /// cluster.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 Container Infra
        /// client. A Container Infra client is needed to create a cluster. If omitted,
        /// the `region` argument of the provider is used. Changing this creates a new
        /// cluster.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The user of the cluster. Required if admin wants to
        /// create a cluster template for another user. Changing this creates a new
        /// cluster.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public ClusterState()
        {
        }
        public static new ClusterState Empty => new ClusterState();
    }
}
