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
    /// Manages a V1 Magnum cluster template resource within OpenStack.
    /// 
    /// 
    /// ## Argument reference
    /// 
    /// The following arguments are supported:
    /// 
    /// * `region` - (Optional) The region in which to obtain the V1 Container Infra
    ///     client. A Container Infra client is needed to create a cluster template. If
    ///     omitted,the `region` argument of the provider is used. Changing this
    ///     creates a new cluster template.
    /// 
    /// * `name` - (Required) The name of the cluster template. Changing this updates
    ///     the name of the existing cluster template.
    /// 
    /// * `project_id` - (Optional) The project of the cluster template. Required if
    ///     admin wants to create a cluster template in another project. Changing this
    ///     creates a new cluster template.
    /// 
    /// * `user_id` - (Optional) The user of the cluster template. Required if admin
    ///     wants to create a cluster template for another user. Changing this creates
    ///     a new cluster template.
    /// 
    /// * `apiserver_port` - (Optional) The API server port for the Container
    ///     Orchestration Engine for this cluster template. Changing this updates the
    ///     API server port of the existing cluster template.
    /// 
    /// * `coe` - (Required) The Container Orchestration Engine for this cluster
    ///     template. Changing this updates the engine of the existing cluster
    ///     template.
    /// 
    /// * `cluster_distro` - (Optional) The distro for the cluster (fedora-atomic,
    ///     coreos, etc.). Changing this updates the cluster distro of the existing
    ///     cluster template.
    /// 
    /// * `dns_nameserver` - (Optional) Address of the DNS nameserver that is used in
    ///     nodes of the cluster. Changing this updates the DNS nameserver of the
    ///     existing cluster template.
    /// 
    /// * `docker_storage_driver` - (Optional) Docker storage driver. Changing this
    ///     updates the Docker storage driver of the existing cluster template.
    /// 
    /// * `docker_volume_size` - (Optional) The size (in GB) of the Docker volume.
    ///     Changing this updates the Docker volume size of the existing cluster
    ///     template.
    /// 
    /// * `external_network_id` - (Optional) The ID of the external network that will
    ///     be used for the cluster. Changing this updates the external network ID of
    ///     the existing cluster template.
    /// 
    /// * `fixed_network` - (Optional) The fixed network that will be attached to the
    ///     cluster. Changing this updates the fixed network of the existing cluster
    ///     template.
    /// 
    /// * `fixed_subnet` - (Optional) The fixed subnet that will be attached to the
    ///     cluster. Changing this updates the fixed subnet of the existing cluster
    ///     template.
    /// 
    /// * `flavor` - (Optional) The flavor for the nodes of the cluster. Can be set via
    ///     the `OS_MAGNUM_FLAVOR` environment variable. Changing this updates the
    ///     flavor of the existing cluster template.
    /// 
    /// * `master_flavor` - (Optional) The flavor for the master nodes. Can be set via
    ///     the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this updates
    ///     the master flavor of the existing cluster template.
    /// 
    /// * `floating_ip_enabled` - (Optional) Indicates whether created cluster should
    ///     create floating IP for every node or not. Changing this updates the
    ///     floating IP enabled attribute of the existing cluster template.
    /// 
    /// * `http_proxy` - (Optional) The address of a proxy for receiving all HTTP
    ///     requests and relay them. Changing this updates the HTTP proxy address of
    ///     the existing cluster template.
    /// 
    /// * `https_proxy` - (Optional) The address of a proxy for receiving all HTTPS
    ///     requests and relay them. Changing this updates the HTTPS proxy address of
    ///     the existing cluster template.
    /// 
    /// * `image` - (Required) The reference to an image that is used for nodes of the
    ///     cluster. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
    ///     Changing this updates the image attribute of the existing cluster template.
    /// 
    /// * `insecure_registry` - (Optional) The insecure registry URL for the cluster
    ///     template. Changing this updates the insecure registry attribute of the
    ///     existing cluster template.
    /// 
    /// * `keypair_id` - (Optional) The name of the Compute service SSH keypair.
    ///     Changing this updates the keypair of the existing cluster template.
    /// 
    /// * `labels` - (Optional) The list of key value pairs representing additional
    ///     properties of the cluster template. Changing this updates the labels of the
    ///     existing cluster template.
    /// 
    /// * `master_lb_enabled` - (Optional) Indicates whether created cluster should
    ///     has a loadbalancer for master nodes or not. Changing this updates the
    ///     attribute of the existing cluster template.
    /// 
    /// * `network_driver` - (Optional) The name of the driver for the container
    ///     network. Changing this updates the network driver of the existing cluster
    ///     template.
    /// 
    /// * `no_proxy` - (Optional) A comma-separated list of IP addresses that shouldn't
    ///     be used in the cluster. Changing this updates the no proxy list of the
    ///     existing cluster template.
    /// 
    /// * `public` - (Optional) Indicates whether cluster template should be public.
    ///     Changing this updates the public attribute of the existing cluster
    ///     template.
    /// 
    /// * `registry_enabled` - (Optional) Indicates whether Docker registry is enabled
    ///     in the cluster. Changing this updates the registry enabled attribute of the
    ///     existing cluster template.
    /// 
    /// * `server_type` - (Optional) The server type for the cluster template. Changing
    ///     this updates the server type of the existing cluster template.
    /// 
    /// * `tls_disabled` - (Optional) Indicates whether the TLS should be disabled in
    ///     the cluster. Changing this updates the attribute of the existing cluster.
    /// 
    /// * `volume_driver` - (Optional) The name of the driver that is used for the
    ///     volumes of the cluster nodes. Changing this updates the volume driver of
    ///     the existing cluster template.
    /// 
    /// ## Attributes reference
    /// 
    /// The following attributes are exported:
    /// 
    /// * `region` - See Argument Reference above.
    /// * `name` - See Argument Reference above.
    /// * `project_id` - See Argument Reference above.
    /// * `created_at` - The time at which cluster template was created.
    /// * `updated_at` - The time at which cluster template was created.
    /// * `apiserver_port` - See Argument Reference above.
    /// * `coe` - See Argument Reference above.
    /// * `cluster_distro` - See Argument Reference above.
    /// * `dns_nameserver` - See Argument Reference above.
    /// * `docker_storage_driver` - See Argument Reference above.
    /// * `docker_volume_size` - See Argument Reference above.
    /// * `external_network_id` - See Argument Reference above.
    /// * `fixed_network` - See Argument Reference above.
    /// * `fixed_subnet` - See Argument Reference above.
    /// * `flavor` - See Argument Reference above.
    /// * `master_flavor` - See Argument Reference above.
    /// * `floating_ip_enabled` - See Argument Reference above.
    /// * `http_proxy` - See Argument Reference above.
    /// * `https_proxy` - See Argument Reference above.
    /// * `image` - See Argument Reference above.
    /// * `insecure_registry` - See Argument Reference above.
    /// * `keypair_id` - See Argument Reference above.
    /// * `labels` - See Argument Reference above.
    /// * `links` - A list containing associated cluster template links.
    /// * `master_lb_enabled` - See Argument Reference above.
    /// * `network_driver` - See Argument Reference above.
    /// * `no_proxy` - See Argument Reference above.
    /// * `public` - See Argument Reference above.
    /// * `registry_enabled` - See Argument Reference above.
    /// * `server_type` - See Argument Reference above.
    /// * `tls_disabled` - See Argument Reference above.
    /// * `volume_driver` - See Argument Reference above.
    /// </summary>
    public partial class ClusterTemplate : Pulumi.CustomResource
    {
        [Output("apiserverPort")]
        public Output<int?> ApiserverPort { get; private set; } = null!;

        [Output("clusterDistro")]
        public Output<string> ClusterDistro { get; private set; } = null!;

        [Output("coe")]
        public Output<string> Coe { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("dnsNameserver")]
        public Output<string?> DnsNameserver { get; private set; } = null!;

        [Output("dockerStorageDriver")]
        public Output<string?> DockerStorageDriver { get; private set; } = null!;

        [Output("dockerVolumeSize")]
        public Output<int?> DockerVolumeSize { get; private set; } = null!;

        [Output("externalNetworkId")]
        public Output<string?> ExternalNetworkId { get; private set; } = null!;

        [Output("fixedNetwork")]
        public Output<string?> FixedNetwork { get; private set; } = null!;

        [Output("fixedSubnet")]
        public Output<string?> FixedSubnet { get; private set; } = null!;

        [Output("flavor")]
        public Output<string?> Flavor { get; private set; } = null!;

        [Output("floatingIpEnabled")]
        public Output<bool?> FloatingIpEnabled { get; private set; } = null!;

        [Output("httpProxy")]
        public Output<string?> HttpProxy { get; private set; } = null!;

        [Output("httpsProxy")]
        public Output<string?> HttpsProxy { get; private set; } = null!;

        [Output("image")]
        public Output<string> Image { get; private set; } = null!;

        [Output("insecureRegistry")]
        public Output<string?> InsecureRegistry { get; private set; } = null!;

        [Output("keypairId")]
        public Output<string?> KeypairId { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableDictionary<string, object>?> Labels { get; private set; } = null!;

        [Output("masterFlavor")]
        public Output<string?> MasterFlavor { get; private set; } = null!;

        [Output("masterLbEnabled")]
        public Output<bool?> MasterLbEnabled { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkDriver")]
        public Output<string> NetworkDriver { get; private set; } = null!;

        [Output("noProxy")]
        public Output<string?> NoProxy { get; private set; } = null!;

        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        [Output("public")]
        public Output<bool?> Public { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("registryEnabled")]
        public Output<bool?> RegistryEnabled { get; private set; } = null!;

        [Output("serverType")]
        public Output<string> ServerType { get; private set; } = null!;

        [Output("tlsDisabled")]
        public Output<bool?> TlsDisabled { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        [Output("volumeDriver")]
        public Output<string?> VolumeDriver { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterTemplate(string name, ClusterTemplateArgs args, CustomResourceOptions? options = null)
            : base("openstack:containerinfra/clusterTemplate:ClusterTemplate", name, args ?? new ClusterTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterTemplate(string name, Input<string> id, ClusterTemplateState? state = null, CustomResourceOptions? options = null)
            : base("openstack:containerinfra/clusterTemplate:ClusterTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClusterTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterTemplate Get(string name, Input<string> id, ClusterTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterTemplate(name, id, state, options);
        }
    }

    public sealed class ClusterTemplateArgs : Pulumi.ResourceArgs
    {
        [Input("apiserverPort")]
        public Input<int>? ApiserverPort { get; set; }

        [Input("clusterDistro")]
        public Input<string>? ClusterDistro { get; set; }

        [Input("coe", required: true)]
        public Input<string> Coe { get; set; } = null!;

        [Input("dnsNameserver")]
        public Input<string>? DnsNameserver { get; set; }

        [Input("dockerStorageDriver")]
        public Input<string>? DockerStorageDriver { get; set; }

        [Input("dockerVolumeSize")]
        public Input<int>? DockerVolumeSize { get; set; }

        [Input("externalNetworkId")]
        public Input<string>? ExternalNetworkId { get; set; }

        [Input("fixedNetwork")]
        public Input<string>? FixedNetwork { get; set; }

        [Input("fixedSubnet")]
        public Input<string>? FixedSubnet { get; set; }

        [Input("flavor")]
        public Input<string>? Flavor { get; set; }

        [Input("floatingIpEnabled")]
        public Input<bool>? FloatingIpEnabled { get; set; }

        [Input("httpProxy")]
        public Input<string>? HttpProxy { get; set; }

        [Input("httpsProxy")]
        public Input<string>? HttpsProxy { get; set; }

        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        [Input("insecureRegistry")]
        public Input<string>? InsecureRegistry { get; set; }

        [Input("keypairId")]
        public Input<string>? KeypairId { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        [Input("masterFlavor")]
        public Input<string>? MasterFlavor { get; set; }

        [Input("masterLbEnabled")]
        public Input<bool>? MasterLbEnabled { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkDriver")]
        public Input<string>? NetworkDriver { get; set; }

        [Input("noProxy")]
        public Input<string>? NoProxy { get; set; }

        [Input("public")]
        public Input<bool>? Public { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("registryEnabled")]
        public Input<bool>? RegistryEnabled { get; set; }

        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        [Input("tlsDisabled")]
        public Input<bool>? TlsDisabled { get; set; }

        [Input("volumeDriver")]
        public Input<string>? VolumeDriver { get; set; }

        public ClusterTemplateArgs()
        {
        }
    }

    public sealed class ClusterTemplateState : Pulumi.ResourceArgs
    {
        [Input("apiserverPort")]
        public Input<int>? ApiserverPort { get; set; }

        [Input("clusterDistro")]
        public Input<string>? ClusterDistro { get; set; }

        [Input("coe")]
        public Input<string>? Coe { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("dnsNameserver")]
        public Input<string>? DnsNameserver { get; set; }

        [Input("dockerStorageDriver")]
        public Input<string>? DockerStorageDriver { get; set; }

        [Input("dockerVolumeSize")]
        public Input<int>? DockerVolumeSize { get; set; }

        [Input("externalNetworkId")]
        public Input<string>? ExternalNetworkId { get; set; }

        [Input("fixedNetwork")]
        public Input<string>? FixedNetwork { get; set; }

        [Input("fixedSubnet")]
        public Input<string>? FixedSubnet { get; set; }

        [Input("flavor")]
        public Input<string>? Flavor { get; set; }

        [Input("floatingIpEnabled")]
        public Input<bool>? FloatingIpEnabled { get; set; }

        [Input("httpProxy")]
        public Input<string>? HttpProxy { get; set; }

        [Input("httpsProxy")]
        public Input<string>? HttpsProxy { get; set; }

        [Input("image")]
        public Input<string>? Image { get; set; }

        [Input("insecureRegistry")]
        public Input<string>? InsecureRegistry { get; set; }

        [Input("keypairId")]
        public Input<string>? KeypairId { get; set; }

        [Input("labels")]
        private InputMap<object>? _labels;
        public InputMap<object> Labels
        {
            get => _labels ?? (_labels = new InputMap<object>());
            set => _labels = value;
        }

        [Input("masterFlavor")]
        public Input<string>? MasterFlavor { get; set; }

        [Input("masterLbEnabled")]
        public Input<bool>? MasterLbEnabled { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networkDriver")]
        public Input<string>? NetworkDriver { get; set; }

        [Input("noProxy")]
        public Input<string>? NoProxy { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("public")]
        public Input<bool>? Public { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("registryEnabled")]
        public Input<bool>? RegistryEnabled { get; set; }

        [Input("serverType")]
        public Input<string>? ServerType { get; set; }

        [Input("tlsDisabled")]
        public Input<bool>? TlsDisabled { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("userId")]
        public Input<string>? UserId { get; set; }

        [Input("volumeDriver")]
        public Input<string>? VolumeDriver { get; set; }

        public ClusterTemplateState()
        {
        }
    }
}
