// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.ContainerInfra
{
    public static class GetClusterTemplate
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack Magnum cluster
        /// template.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetClusterTemplateResult> InvokeAsync(GetClusterTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterTemplateResult>("openstack:containerinfra/getClusterTemplate:getClusterTemplate", args ?? new GetClusterTemplateArgs(), options.WithVersion());
    }


    public sealed class GetClusterTemplateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster template.
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

        public GetClusterTemplateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterTemplateResult
    {
        /// <summary>
        /// The API server port for the Container Orchestration
        /// Engine for this cluster template.
        /// </summary>
        public readonly int ApiserverPort;
        /// <summary>
        /// The distro for the cluster (fedora-atomic, coreos, etc.).
        /// </summary>
        public readonly string ClusterDistro;
        /// <summary>
        /// The Container Orchestration Engine for this cluster template.
        /// </summary>
        public readonly string Coe;
        /// <summary>
        /// The time at which cluster template was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Address of the DNS nameserver that is used in nodes of the
        /// cluster.
        /// </summary>
        public readonly string DnsNameserver;
        /// <summary>
        /// Docker storage driver. Changing this updates the
        /// Docker storage driver of the existing cluster template.
        /// </summary>
        public readonly string DockerStorageDriver;
        /// <summary>
        /// The size (in GB) of the Docker volume.
        /// </summary>
        public readonly int DockerVolumeSize;
        /// <summary>
        /// The ID of the external network that will be used for
        /// the cluster.
        /// </summary>
        public readonly string ExternalNetworkId;
        /// <summary>
        /// The fixed network that will be attached to the cluster.
        /// </summary>
        public readonly string FixedNetwork;
        /// <summary>
        /// =The fixed subnet that will be attached to the cluster.
        /// </summary>
        public readonly string FixedSubnet;
        /// <summary>
        /// The flavor for the nodes of the cluster.
        /// </summary>
        public readonly string Flavor;
        /// <summary>
        /// Indicates whether created cluster should create IP
        /// floating IP for every node or not.
        /// </summary>
        public readonly bool FloatingIpEnabled;
        /// <summary>
        /// The address of a proxy for receiving all HTTP requests and
        /// relay them.
        /// </summary>
        public readonly string HttpProxy;
        /// <summary>
        /// The address of a proxy for receiving all HTTPS requests and
        /// relay them.
        /// </summary>
        public readonly string HttpsProxy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The reference to an image that is used for nodes of the cluster.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// The insecure registry URL for the cluster template.
        /// </summary>
        public readonly string InsecureRegistry;
        /// <summary>
        /// The name of the Compute service SSH keypair.
        /// </summary>
        public readonly string KeypairId;
        /// <summary>
        /// The list of key value pairs representing additional properties
        /// of the cluster template.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        /// <summary>
        /// The flavor for the master nodes.
        /// </summary>
        public readonly string MasterFlavor;
        /// <summary>
        /// Indicates whether created cluster should has a
        /// loadbalancer for master nodes or not.
        /// </summary>
        public readonly bool MasterLbEnabled;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the driver for the container network.
        /// </summary>
        public readonly string NetworkDriver;
        /// <summary>
        /// A comma-separated list of IP addresses that shouldn't be used in
        /// the cluster.
        /// </summary>
        public readonly string NoProxy;
        /// <summary>
        /// The project of the cluster template.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Indicates whether cluster template should be public.
        /// </summary>
        public readonly bool Public;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Indicates whether Docker registry is enabled in the
        /// cluster.
        /// </summary>
        public readonly bool RegistryEnabled;
        /// <summary>
        /// The server type for the cluster template.
        /// </summary>
        public readonly string ServerType;
        /// <summary>
        /// Indicates whether the TLS should be disabled in the cluster.
        /// </summary>
        public readonly bool TlsDisabled;
        /// <summary>
        /// The time at which cluster template was updated.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The user of the cluster template.
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// The name of the driver that is used for the volumes of the
        /// cluster nodes.
        /// </summary>
        public readonly string VolumeDriver;

        [OutputConstructor]
        private GetClusterTemplateResult(
            int apiserverPort,

            string clusterDistro,

            string coe,

            string createdAt,

            string dnsNameserver,

            string dockerStorageDriver,

            int dockerVolumeSize,

            string externalNetworkId,

            string fixedNetwork,

            string fixedSubnet,

            string flavor,

            bool floatingIpEnabled,

            string httpProxy,

            string httpsProxy,

            string id,

            string image,

            string insecureRegistry,

            string keypairId,

            ImmutableDictionary<string, object> labels,

            string masterFlavor,

            bool masterLbEnabled,

            string name,

            string networkDriver,

            string noProxy,

            string projectId,

            bool @public,

            string region,

            bool registryEnabled,

            string serverType,

            bool tlsDisabled,

            string updatedAt,

            string userId,

            string volumeDriver)
        {
            ApiserverPort = apiserverPort;
            ClusterDistro = clusterDistro;
            Coe = coe;
            CreatedAt = createdAt;
            DnsNameserver = dnsNameserver;
            DockerStorageDriver = dockerStorageDriver;
            DockerVolumeSize = dockerVolumeSize;
            ExternalNetworkId = externalNetworkId;
            FixedNetwork = fixedNetwork;
            FixedSubnet = fixedSubnet;
            Flavor = flavor;
            FloatingIpEnabled = floatingIpEnabled;
            HttpProxy = httpProxy;
            HttpsProxy = httpsProxy;
            Id = id;
            Image = image;
            InsecureRegistry = insecureRegistry;
            KeypairId = keypairId;
            Labels = labels;
            MasterFlavor = masterFlavor;
            MasterLbEnabled = masterLbEnabled;
            Name = name;
            NetworkDriver = networkDriver;
            NoProxy = noProxy;
            ProjectId = projectId;
            Public = @public;
            Region = region;
            RegistryEnabled = registryEnabled;
            ServerType = serverType;
            TlsDisabled = tlsDisabled;
            UpdatedAt = updatedAt;
            UserId = userId;
            VolumeDriver = volumeDriver;
        }
    }
}
