// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Dns
{
    public static class GetDnsZone
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack DNS zone.
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
        ///     var zone1 = OpenStack.Dns.GetDnsZone.Invoke(new()
        ///     {
        ///         Name = "example.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDnsZoneResult> InvokeAsync(GetDnsZoneArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDnsZoneResult>("openstack:dns/getDnsZone:getDnsZone", args ?? new GetDnsZoneArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack DNS zone.
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
        ///     var zone1 = OpenStack.Dns.GetDnsZone.Invoke(new()
        ///     {
        ///         Name = "example.com",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDnsZoneResult> Invoke(GetDnsZoneInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDnsZoneResult>("openstack:dns/getDnsZone:getDnsZone", args ?? new GetDnsZoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDnsZoneArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Try to obtain zone ID by listing all projects
        /// (requires admin role by default, depends on your policy configuration)
        /// </summary>
        [Input("allProjects")]
        public string? AllProjects { get; set; }

        [Input("attributes")]
        private Dictionary<string, string>? _attributes;

        /// <summary>
        /// Attributes of the DNS Service scheduler.
        /// </summary>
        public Dictionary<string, string> Attributes
        {
            get => _attributes ?? (_attributes = new Dictionary<string, string>());
            set => _attributes = value;
        }

        /// <summary>
        /// The time the zone was created.
        /// </summary>
        [Input("createdAt")]
        public string? CreatedAt { get; set; }

        /// <summary>
        /// A description of the zone.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The email contact for the zone record.
        /// </summary>
        [Input("email")]
        public string? Email { get; set; }

        [Input("masters")]
        private List<string>? _masters;

        /// <summary>
        /// An array of master DNS servers. When `type` is  `SECONDARY`.
        /// </summary>
        public List<string> Masters
        {
            get => _masters ?? (_masters = new List<string>());
            set => _masters = value;
        }

        /// <summary>
        /// The name of the zone.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the pool hosting the zone.
        /// </summary>
        [Input("poolId")]
        public string? PoolId { get; set; }

        /// <summary>
        /// The ID of the project the DNS zone is obtained from,
        /// sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned user role in target project)
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 DNS client.
        /// A DNS client is needed to retrieve zone ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The serial number of the zone.
        /// </summary>
        [Input("serial")]
        public int? Serial { get; set; }

        /// <summary>
        /// The zone's status.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The time the zone was transferred.
        /// </summary>
        [Input("transferredAt")]
        public string? TransferredAt { get; set; }

        /// <summary>
        /// The time to live (TTL) of the zone.
        /// </summary>
        [Input("ttl")]
        public int? Ttl { get; set; }

        /// <summary>
        /// The type of the zone. Can either be `PRIMARY` or `SECONDARY`.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// The time the zone was last updated.
        /// </summary>
        [Input("updatedAt")]
        public string? UpdatedAt { get; set; }

        /// <summary>
        /// The version of the zone.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetDnsZoneArgs()
        {
        }
        public static new GetDnsZoneArgs Empty => new GetDnsZoneArgs();
    }

    public sealed class GetDnsZoneInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Try to obtain zone ID by listing all projects
        /// (requires admin role by default, depends on your policy configuration)
        /// </summary>
        [Input("allProjects")]
        public Input<string>? AllProjects { get; set; }

        [Input("attributes")]
        private InputMap<string>? _attributes;

        /// <summary>
        /// Attributes of the DNS Service scheduler.
        /// </summary>
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        /// <summary>
        /// The time the zone was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A description of the zone.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The email contact for the zone record.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("masters")]
        private InputList<string>? _masters;

        /// <summary>
        /// An array of master DNS servers. When `type` is  `SECONDARY`.
        /// </summary>
        public InputList<string> Masters
        {
            get => _masters ?? (_masters = new InputList<string>());
            set => _masters = value;
        }

        /// <summary>
        /// The name of the zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the pool hosting the zone.
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        /// <summary>
        /// The ID of the project the DNS zone is obtained from,
        /// sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned user role in target project)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 DNS client.
        /// A DNS client is needed to retrieve zone ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The serial number of the zone.
        /// </summary>
        [Input("serial")]
        public Input<int>? Serial { get; set; }

        /// <summary>
        /// The zone's status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The time the zone was transferred.
        /// </summary>
        [Input("transferredAt")]
        public Input<string>? TransferredAt { get; set; }

        /// <summary>
        /// The time to live (TTL) of the zone.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of the zone. Can either be `PRIMARY` or `SECONDARY`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The time the zone was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The version of the zone.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public GetDnsZoneInvokeArgs()
        {
        }
        public static new GetDnsZoneInvokeArgs Empty => new GetDnsZoneInvokeArgs();
    }


    [OutputType]
    public sealed class GetDnsZoneResult
    {
        public readonly string? AllProjects;
        /// <summary>
        /// Attributes of the DNS Service scheduler.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Attributes;
        /// <summary>
        /// The time the zone was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An array of master DNS servers. When `type` is  `SECONDARY`.
        /// </summary>
        public readonly ImmutableArray<string> Masters;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ID of the pool hosting the zone.
        /// </summary>
        public readonly string PoolId;
        /// <summary>
        /// The project ID that owns the zone.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The serial number of the zone.
        /// </summary>
        public readonly int Serial;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The time the zone was transferred.
        /// </summary>
        public readonly string TransferredAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int? Ttl;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The time the zone was last updated.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The version of the zone.
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GetDnsZoneResult(
            string? allProjects,

            ImmutableDictionary<string, string> attributes,

            string createdAt,

            string? description,

            string? email,

            string id,

            ImmutableArray<string> masters,

            string? name,

            string poolId,

            string projectId,

            string region,

            int serial,

            string? status,

            string transferredAt,

            int? ttl,

            string? type,

            string updatedAt,

            int version)
        {
            AllProjects = allProjects;
            Attributes = attributes;
            CreatedAt = createdAt;
            Description = description;
            Email = email;
            Id = id;
            Masters = masters;
            Name = name;
            PoolId = poolId;
            ProjectId = projectId;
            Region = region;
            Serial = serial;
            Status = status;
            TransferredAt = transferredAt;
            Ttl = ttl;
            Type = type;
            UpdatedAt = updatedAt;
            Version = version;
        }
    }
}
