// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetSubnetIdsV2
    {
        /// <summary>
        /// Use this data source to get a list of Openstack Subnet IDs matching the
        /// specified criteria.
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
        ///     var subnets = OpenStack.Networking.GetSubnetIdsV2.Invoke(new()
        ///     {
        ///         NameRegex = "public",
        ///         Tags = new[]
        ///         {
        ///             "public",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSubnetIdsV2Result> InvokeAsync(GetSubnetIdsV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubnetIdsV2Result>("openstack:networking/getSubnetIdsV2:getSubnetIdsV2", args ?? new GetSubnetIdsV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a list of Openstack Subnet IDs matching the
        /// specified criteria.
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
        ///     var subnets = OpenStack.Networking.GetSubnetIdsV2.Invoke(new()
        ///     {
        ///         NameRegex = "public",
        ///         Tags = new[]
        ///         {
        ///             "public",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSubnetIdsV2Result> Invoke(GetSubnetIdsV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubnetIdsV2Result>("openstack:networking/getSubnetIdsV2:getSubnetIdsV2", args ?? new GetSubnetIdsV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubnetIdsV2Args : global::Pulumi.InvokeArgs
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

        [Input("nameRegex")]
        public string? NameRegex { get; set; }

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
        /// Order the results in either `asc` or `desc`.
        /// Defaults to none.
        /// </summary>
        [Input("sortDirection")]
        public string? SortDirection { get; set; }

        /// <summary>
        /// Sort subnets based on a certain key. Defaults to none.
        /// </summary>
        [Input("sortKey")]
        public string? SortKey { get; set; }

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

        public GetSubnetIdsV2Args()
        {
        }
        public static new GetSubnetIdsV2Args Empty => new GetSubnetIdsV2Args();
    }

    public sealed class GetSubnetIdsV2InvokeArgs : global::Pulumi.InvokeArgs
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

        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

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
        /// Order the results in either `asc` or `desc`.
        /// Defaults to none.
        /// </summary>
        [Input("sortDirection")]
        public Input<string>? SortDirection { get; set; }

        /// <summary>
        /// Sort subnets based on a certain key. Defaults to none.
        /// </summary>
        [Input("sortKey")]
        public Input<string>? SortKey { get; set; }

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

        public GetSubnetIdsV2InvokeArgs()
        {
        }
        public static new GetSubnetIdsV2InvokeArgs Empty => new GetSubnetIdsV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetSubnetIdsV2Result
    {
        public readonly string? Cidr;
        public readonly string? Description;
        public readonly bool? DhcpEnabled;
        public readonly string? GatewayIp;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly int? IpVersion;
        public readonly string? Ipv6AddressMode;
        public readonly string Ipv6RaMode;
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly string? NetworkId;
        public readonly string Region;
        public readonly string? SortDirection;
        public readonly string? SortKey;
        public readonly string? SubnetpoolId;
        public readonly ImmutableArray<string> Tags;
        public readonly string? TenantId;

        [OutputConstructor]
        private GetSubnetIdsV2Result(
            string? cidr,

            string? description,

            bool? dhcpEnabled,

            string? gatewayIp,

            string id,

            ImmutableArray<string> ids,

            int? ipVersion,

            string? ipv6AddressMode,

            string ipv6RaMode,

            string? name,

            string? nameRegex,

            string? networkId,

            string region,

            string? sortDirection,

            string? sortKey,

            string? subnetpoolId,

            ImmutableArray<string> tags,

            string? tenantId)
        {
            Cidr = cidr;
            Description = description;
            DhcpEnabled = dhcpEnabled;
            GatewayIp = gatewayIp;
            Id = id;
            Ids = ids;
            IpVersion = ipVersion;
            Ipv6AddressMode = ipv6AddressMode;
            Ipv6RaMode = ipv6RaMode;
            Name = name;
            NameRegex = nameRegex;
            NetworkId = networkId;
            Region = region;
            SortDirection = sortDirection;
            SortKey = sortKey;
            SubnetpoolId = subnetpoolId;
            Tags = tags;
            TenantId = tenantId;
        }
    }
}
