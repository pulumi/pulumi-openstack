// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetSubnetPool
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack subnetpool.
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
        ///         var subnetpool1 = Output.Create(OpenStack.Networking.GetSubnetPool.InvokeAsync(new OpenStack.Networking.GetSubnetPoolArgs
        ///         {
        ///             Name = "subnetpool_1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSubnetPoolResult> InvokeAsync(GetSubnetPoolArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetPoolResult>("openstack:networking/getSubnetPool:getSubnetPool", args ?? new GetSubnetPoolArgs(), options.WithVersion());
    }


    public sealed class GetSubnetPoolArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Neutron address scope that subnetpools
        /// is assigned to.
        /// </summary>
        [Input("addressScopeId")]
        public string? AddressScopeId { get; set; }

        /// <summary>
        /// The size of the subnetpool default prefix
        /// length.
        /// </summary>
        [Input("defaultPrefixlen")]
        public int? DefaultPrefixlen { get; set; }

        /// <summary>
        /// The per-project quota on the prefix space that
        /// can be allocated from the subnetpool for project subnets.
        /// </summary>
        [Input("defaultQuota")]
        public int? DefaultQuota { get; set; }

        /// <summary>
        /// The human-readable description for the subnetpool.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The IP protocol version.
        /// </summary>
        [Input("ipVersion")]
        public int? IpVersion { get; set; }

        /// <summary>
        /// Whether the subnetpool is default subnetpool or not.
        /// </summary>
        [Input("isDefault")]
        public bool? IsDefault { get; set; }

        /// <summary>
        /// The size of the subnetpool max prefix length.
        /// </summary>
        [Input("maxPrefixlen")]
        public int? MaxPrefixlen { get; set; }

        /// <summary>
        /// The size of the subnetpool min prefix length.
        /// </summary>
        [Input("minPrefixlen")]
        public int? MinPrefixlen { get; set; }

        /// <summary>
        /// The name of the subnetpool.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The owner of the subnetpool.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to retrieve a subnetpool id. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Whether this subnetpool is shared across all projects.
        /// </summary>
        [Input("shared")]
        public bool? Shared { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The list of subnetpool tags to filter.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        public GetSubnetPoolArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubnetPoolResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// * `ip_version` -The IP protocol version.
        /// </summary>
        public readonly string AddressScopeId;
        /// <summary>
        /// The set of string tags applied on the subnetpool.
        /// </summary>
        public readonly ImmutableArray<string> AllTags;
        /// <summary>
        /// The time at which subnetpool was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int DefaultPrefixlen;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int DefaultQuota;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int IpVersion;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int MaxPrefixlen;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int MinPrefixlen;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly ImmutableArray<string> Prefixes;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The revision number of the subnetpool.
        /// </summary>
        public readonly int RevisionNumber;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool Shared;
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The time at which subnetpool was created.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetSubnetPoolResult(
            string addressScopeId,

            ImmutableArray<string> allTags,

            string createdAt,

            int defaultPrefixlen,

            int defaultQuota,

            string description,

            string id,

            int ipVersion,

            bool isDefault,

            int maxPrefixlen,

            int minPrefixlen,

            string name,

            ImmutableArray<string> prefixes,

            string projectId,

            string region,

            int revisionNumber,

            bool shared,

            ImmutableArray<string> tags,

            string updatedAt)
        {
            AddressScopeId = addressScopeId;
            AllTags = allTags;
            CreatedAt = createdAt;
            DefaultPrefixlen = defaultPrefixlen;
            DefaultQuota = defaultQuota;
            Description = description;
            Id = id;
            IpVersion = ipVersion;
            IsDefault = isDefault;
            MaxPrefixlen = maxPrefixlen;
            MinPrefixlen = minPrefixlen;
            Name = name;
            Prefixes = prefixes;
            ProjectId = projectId;
            Region = region;
            RevisionNumber = revisionNumber;
            Shared = shared;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
