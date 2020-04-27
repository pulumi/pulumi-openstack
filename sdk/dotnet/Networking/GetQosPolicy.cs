// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetQosPolicy
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS policy.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetQosPolicyResult> InvokeAsync(GetQosPolicyArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQosPolicyResult>("openstack:networking/getQosPolicy:getQosPolicy", args ?? new GetQosPolicyArgs(), options.WithVersion());
    }


    public sealed class GetQosPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The human-readable description for the QoS policy.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Whether the QoS policy is default policy or not.
        /// </summary>
        [Input("isDefault")]
        public bool? IsDefault { get; set; }

        /// <summary>
        /// The name of the QoS policy.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The owner of the QoS policy.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to retrieve a QoS policy ID. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Whether this QoS policy is shared across all projects.
        /// </summary>
        [Input("shared")]
        public bool? Shared { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The list of QoS policy tags to filter.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        public GetQosPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetQosPolicyResult
    {
        /// <summary>
        /// The set of string tags applied on the QoS policy.
        /// </summary>
        public readonly ImmutableArray<string> AllTags;
        /// <summary>
        /// The time at which QoS policy was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The revision number of the QoS policy.
        /// </summary>
        public readonly int RevisionNumber;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool Shared;
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The time at which QoS policy was created.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetQosPolicyResult(
            ImmutableArray<string> allTags,

            string createdAt,

            string description,

            string id,

            bool isDefault,

            string name,

            string projectId,

            string region,

            int revisionNumber,

            bool shared,

            ImmutableArray<string> tags,

            string updatedAt)
        {
            AllTags = allTags;
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            IsDefault = isDefault;
            Name = name;
            ProjectId = projectId;
            Region = region;
            RevisionNumber = revisionNumber;
            Shared = shared;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
