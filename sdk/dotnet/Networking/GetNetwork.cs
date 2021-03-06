// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetNetwork
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack network.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var network = Output.Create(OpenStack.Networking.GetNetwork.InvokeAsync(new OpenStack.Networking.GetNetworkArgs
        ///         {
        ///             Name = "tf_test_network",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkResult> InvokeAsync(GetNetworkArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("openstack:networking/getNetwork:getNetwork", args ?? new GetNetworkArgs(), options.WithVersion());
    }


    public sealed class GetNetworkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Human-readable description of the network.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The external routing facility of the network.
        /// </summary>
        [Input("external")]
        public bool? External { get; set; }

        /// <summary>
        /// The CIDR of a subnet within the network.
        /// </summary>
        [Input("matchingSubnetCidr")]
        public string? MatchingSubnetCidr { get; set; }

        /// <summary>
        /// The network MTU to filter. Available, when Neutron `net-mtu`
        /// extension is enabled.
        /// </summary>
        [Input("mtu")]
        public int? Mtu { get; set; }

        /// <summary>
        /// The name of the network.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the network.
        /// </summary>
        [Input("networkId")]
        public string? NetworkId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve networks ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The status of the network.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The list of network tags to filter.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the network.
        /// </summary>
        [Input("tenantId")]
        public string? TenantId { get; set; }

        /// <summary>
        /// The VLAN transparent attribute for the
        /// network.
        /// </summary>
        [Input("transparentVlan")]
        public bool? TransparentVlan { get; set; }

        public GetNetworkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkResult
    {
        /// <summary>
        /// The administrative state of the network.
        /// </summary>
        public readonly string AdminStateUp;
        /// <summary>
        /// The set of string tags applied on the network.
        /// </summary>
        public readonly ImmutableArray<string> AllTags;
        /// <summary>
        /// The availability zone candidates for the network.
        /// </summary>
        public readonly ImmutableArray<string> AvailabilityZoneHints;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The network DNS domain. Available, when Neutron DNS extension
        /// is enabled
        /// </summary>
        public readonly string DnsDomain;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool? External;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? MatchingSubnetCidr;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int? Mtu;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        public readonly string? NetworkId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Specifies whether the network resource can be accessed by any
        /// tenant or not.
        /// </summary>
        public readonly string Shared;
        public readonly string? Status;
        /// <summary>
        /// A list of subnet IDs belonging to the network.
        /// </summary>
        public readonly ImmutableArray<string> Subnets;
        public readonly ImmutableArray<string> Tags;
        public readonly string? TenantId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool? TransparentVlan;

        [OutputConstructor]
        private GetNetworkResult(
            string adminStateUp,

            ImmutableArray<string> allTags,

            ImmutableArray<string> availabilityZoneHints,

            string? description,

            string dnsDomain,

            bool? external,

            string id,

            string? matchingSubnetCidr,

            int? mtu,

            string? name,

            string? networkId,

            string region,

            string shared,

            string? status,

            ImmutableArray<string> subnets,

            ImmutableArray<string> tags,

            string? tenantId,

            bool? transparentVlan)
        {
            AdminStateUp = adminStateUp;
            AllTags = allTags;
            AvailabilityZoneHints = availabilityZoneHints;
            Description = description;
            DnsDomain = dnsDomain;
            External = external;
            Id = id;
            MatchingSubnetCidr = matchingSubnetCidr;
            Mtu = mtu;
            Name = name;
            NetworkId = networkId;
            Region = region;
            Shared = shared;
            Status = status;
            Subnets = subnets;
            Tags = tags;
            TenantId = tenantId;
            TransparentVlan = transparentVlan;
        }
    }
}
