// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an available Barbican container.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/keymanager_container_v1.html.markdown.
        /// </summary>
        public static Task<GetContainerResult> GetContainer(GetContainerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetContainerResult>("openstack:keymanager/getContainer:getContainer", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Container name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to fetch a container. If omitted, the `region`
        /// argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetContainerArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetContainerResult
    {
        /// <summary>
        /// The list of the container consumers. The structure is described
        /// below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContainerConsumersResult> Consumers;
        /// <summary>
        /// The container reference / where to find the container.
        /// </summary>
        public readonly string ContainerRef;
        /// <summary>
        /// The date the container was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The creator of the container.
        /// </summary>
        public readonly string CreatorId;
        /// <summary>
        /// The name of the consumer.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// A set of dictionaries containing references to secrets. The
        /// structure is described below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContainerSecretRefsResult> SecretRefs;
        /// <summary>
        /// The status of the container.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The container type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The date the container was last updated.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetContainerResult(
            ImmutableArray<Outputs.GetContainerConsumersResult> consumers,
            string containerRef,
            string createdAt,
            string creatorId,
            string? name,
            string? region,
            ImmutableArray<Outputs.GetContainerSecretRefsResult> secretRefs,
            string status,
            string type,
            string updatedAt,
            string id)
        {
            Consumers = consumers;
            ContainerRef = containerRef;
            CreatedAt = createdAt;
            CreatorId = creatorId;
            Name = name;
            Region = region;
            SecretRefs = secretRefs;
            Status = status;
            Type = type;
            UpdatedAt = updatedAt;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetContainerConsumersResult
    {
        /// <summary>
        /// The Container name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The consumer URL.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private GetContainerConsumersResult(
            string? name,
            string? url)
        {
            Name = name;
            Url = url;
        }
    }

    [OutputType]
    public sealed class GetContainerSecretRefsResult
    {
        /// <summary>
        /// The Container name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The secret reference / where to find the secret, URL.
        /// </summary>
        public readonly string? SecretRef;

        [OutputConstructor]
        private GetContainerSecretRefsResult(
            string? name,
            string? secretRef)
        {
            Name = name;
            SecretRef = secretRef;
        }
    }
    }
}