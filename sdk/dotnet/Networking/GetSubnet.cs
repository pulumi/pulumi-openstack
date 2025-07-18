// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetSubnet
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack subnet.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var subnet1 = OpenStack.Networking.GetSubnet.Invoke(new()
        ///     {
        ///         Name = "subnet_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSubnetResult> InvokeAsync(GetSubnetArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubnetResult>("openstack:networking/getSubnet:getSubnet", args ?? new GetSubnetArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack subnet.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var subnet1 = OpenStack.Networking.GetSubnet.Invoke(new()
        ///     {
        ///         Name = "subnet_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSubnetResult> Invoke(GetSubnetInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubnetResult>("openstack:networking/getSubnet:getSubnet", args ?? new GetSubnetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack subnet.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var subnet1 = OpenStack.Networking.GetSubnet.Invoke(new()
        ///     {
        ///         Name = "subnet_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSubnetResult> Invoke(GetSubnetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubnetResult>("openstack:networking/getSubnet:getSubnet", args ?? new GetSubnetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubnetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CIDR of the subnet.
        /// </summary>
        [Input("cidr")]
        public string? Cidr { get; set; }

        /// <summary>
        /// Human-readable description of the subnet.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// If the subnet has DHCP enabled.
        /// </summary>
        [Input("dhcpEnabled")]
        public bool? DhcpEnabled { get; set; }

        /// <summary>
        /// If the subnet publishes DNS records.
        /// </summary>
        [Input("dnsPublishFixedIp")]
        public bool? DnsPublishFixedIp { get; set; }

        /// <summary>
        /// The IP of the subnet's gateway.
        /// </summary>
        [Input("gatewayIp")]
        public string? GatewayIp { get; set; }

        /// <summary>
        /// The IP version of the subnet (either 4 or 6).
        /// </summary>
        [Input("ipVersion")]
        public int? IpVersion { get; set; }

        /// <summary>
        /// The IPv6 address mode. Valid values are
        /// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
        /// </summary>
        [Input("ipv6AddressMode")]
        public string? Ipv6AddressMode { get; set; }

        /// <summary>
        /// The IPv6 Router Advertisement mode. Valid values
        /// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
        /// </summary>
        [Input("ipv6RaMode")]
        public string? Ipv6RaMode { get; set; }

        /// <summary>
        /// The name of the subnet.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the network the subnet belongs to.
        /// </summary>
        [Input("networkId")]
        public string? NetworkId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve subnet ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The ID of the segment the subnet belongs to.
        /// Available when neutron segment extension is enabled.
        /// </summary>
        [Input("segmentId")]
        public string? SegmentId { get; set; }

        /// <summary>
        /// The ID of the subnet.
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        /// <summary>
        /// The ID of the subnetpool associated with the subnet.
        /// </summary>
        [Input("subnetpoolId")]
        public string? SubnetpoolId { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The list of subnet tags to filter.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the subnet.
        /// </summary>
        [Input("tenantId")]
        public string? TenantId { get; set; }

        public GetSubnetArgs()
        {
        }
        public static new GetSubnetArgs Empty => new GetSubnetArgs();
    }

    public sealed class GetSubnetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CIDR of the subnet.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// Human-readable description of the subnet.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If the subnet has DHCP enabled.
        /// </summary>
        [Input("dhcpEnabled")]
        public Input<bool>? DhcpEnabled { get; set; }

        /// <summary>
        /// If the subnet publishes DNS records.
        /// </summary>
        [Input("dnsPublishFixedIp")]
        public Input<bool>? DnsPublishFixedIp { get; set; }

        /// <summary>
        /// The IP of the subnet's gateway.
        /// </summary>
        [Input("gatewayIp")]
        public Input<string>? GatewayIp { get; set; }

        /// <summary>
        /// The IP version of the subnet (either 4 or 6).
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// The IPv6 address mode. Valid values are
        /// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
        /// </summary>
        [Input("ipv6AddressMode")]
        public Input<string>? Ipv6AddressMode { get; set; }

        /// <summary>
        /// The IPv6 Router Advertisement mode. Valid values
        /// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
        /// </summary>
        [Input("ipv6RaMode")]
        public Input<string>? Ipv6RaMode { get; set; }

        /// <summary>
        /// The name of the subnet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the network the subnet belongs to.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve subnet ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the segment the subnet belongs to.
        /// Available when neutron segment extension is enabled.
        /// </summary>
        [Input("segmentId")]
        public Input<string>? SegmentId { get; set; }

        /// <summary>
        /// The ID of the subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The ID of the subnetpool associated with the subnet.
        /// </summary>
        [Input("subnetpoolId")]
        public Input<string>? SubnetpoolId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The list of subnet tags to filter.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the subnet.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public GetSubnetInvokeArgs()
        {
        }
        public static new GetSubnetInvokeArgs Empty => new GetSubnetInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubnetResult
    {
        /// <summary>
        /// A set of string tags applied on the subnet.
        /// </summary>
        public readonly ImmutableArray<string> AllTags;
        /// <summary>
        /// Allocation pools of the subnet.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSubnetAllocationPoolResult> AllocationPools;
        public readonly string Cidr;
        public readonly string Description;
        public readonly bool? DhcpEnabled;
        /// <summary>
        /// DNS Nameservers of the subnet.
        /// </summary>
        public readonly ImmutableArray<string> DnsNameservers;
        public readonly bool? DnsPublishFixedIp;
        /// <summary>
        /// Whether the subnet has DHCP enabled or not.
        /// </summary>
        public readonly bool EnableDhcp;
        public readonly string GatewayIp;
        /// <summary>
        /// Host Routes of the subnet.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSubnetHostRouteResult> HostRoutes;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int IpVersion;
        public readonly string Ipv6AddressMode;
        public readonly string Ipv6RaMode;
        public readonly string Name;
        public readonly string NetworkId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        public readonly string SegmentId;
        /// <summary>
        /// Service types of the subnet.
        /// </summary>
        public readonly ImmutableArray<string> ServiceTypes;
        public readonly string SubnetId;
        public readonly string SubnetpoolId;
        public readonly ImmutableArray<string> Tags;
        public readonly string TenantId;

        [OutputConstructor]
        private GetSubnetResult(
            ImmutableArray<string> allTags,

            ImmutableArray<Outputs.GetSubnetAllocationPoolResult> allocationPools,

            string cidr,

            string description,

            bool? dhcpEnabled,

            ImmutableArray<string> dnsNameservers,

            bool? dnsPublishFixedIp,

            bool enableDhcp,

            string gatewayIp,

            ImmutableArray<Outputs.GetSubnetHostRouteResult> hostRoutes,

            string id,

            int ipVersion,

            string ipv6AddressMode,

            string ipv6RaMode,

            string name,

            string networkId,

            string region,

            string segmentId,

            ImmutableArray<string> serviceTypes,

            string subnetId,

            string subnetpoolId,

            ImmutableArray<string> tags,

            string tenantId)
        {
            AllTags = allTags;
            AllocationPools = allocationPools;
            Cidr = cidr;
            Description = description;
            DhcpEnabled = dhcpEnabled;
            DnsNameservers = dnsNameservers;
            DnsPublishFixedIp = dnsPublishFixedIp;
            EnableDhcp = enableDhcp;
            GatewayIp = gatewayIp;
            HostRoutes = hostRoutes;
            Id = id;
            IpVersion = ipVersion;
            Ipv6AddressMode = ipv6AddressMode;
            Ipv6RaMode = ipv6RaMode;
            Name = name;
            NetworkId = networkId;
            Region = region;
            SegmentId = segmentId;
            ServiceTypes = serviceTypes;
            SubnetId = subnetId;
            SubnetpoolId = subnetpoolId;
            Tags = tags;
            TenantId = tenantId;
        }
    }
}
