// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack security group.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_secgroup_v2.html.markdown.
        /// </summary>
        public static Task<GetSecGroupResult> GetSecGroup(GetSecGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecGroupResult>("openstack:networking/getSecGroup:getSecGroup", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetSecGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-readable description the the subnet.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the security group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve security groups ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("secgroupId")]
        public Input<string>? SecgroupId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The list of security group tags to filter.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the security group.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public GetSecGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSecGroupResult
    {
        /// <summary>
        /// The set of string tags applied on the security group.
        /// </summary>
        public readonly ImmutableArray<string> AllTags;
        public readonly string? Description;
        /// <summary>
        /// See Argument Reference above.
        /// * `description`- See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        public readonly string? SecgroupId;
        public readonly ImmutableArray<string> Tags;
        public readonly string TenantId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSecGroupResult(
            ImmutableArray<string> allTags,
            string? description,
            string? name,
            string region,
            string? secgroupId,
            ImmutableArray<string> tags,
            string tenantId,
            string id)
        {
            AllTags = allTags;
            Description = description;
            Name = name;
            Region = region;
            SecgroupId = secgroupId;
            Tags = tags;
            TenantId = tenantId;
            Id = id;
        }
    }
}
