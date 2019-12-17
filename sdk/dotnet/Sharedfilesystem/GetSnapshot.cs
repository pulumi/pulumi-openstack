// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.SharedFileSystem
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an available Shared File System snapshot.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/sharedfilesystem_snapshot_v2.html.markdown.
        /// </summary>
        public static Task<GetSnapshotResult> GetSnapshot(GetSnapshotArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("openstack:sharedfilesystem/getSnapshot:getSnapshot", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetSnapshotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description of the snapshot.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("shareId")]
        public Input<string>? ShareId { get; set; }

        /// <summary>
        /// A snapshot status filter. A valid value is `available`, `error`,
        /// `creating`, `deleting`, `manage_starting`, `manage_error`, `unmanage_starting`,
        /// `unmanage_error` or `error_deleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetSnapshotArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSnapshotResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ProjectId;
        public readonly string Region;
        /// <summary>
        /// The UUID of the source share that was used to create the snapshot.
        /// </summary>
        public readonly string ShareId;
        /// <summary>
        /// The file system protocol of a share snapshot.
        /// </summary>
        public readonly string ShareProto;
        /// <summary>
        /// The share snapshot size, in GBs.
        /// </summary>
        public readonly int ShareSize;
        /// <summary>
        /// The snapshot size, in GBs.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSnapshotResult(
            string description,
            string name,
            string projectId,
            string region,
            string shareId,
            string shareProto,
            int shareSize,
            int size,
            string status,
            string id)
        {
            Description = description;
            Name = name;
            ProjectId = projectId;
            Region = region;
            ShareId = shareId;
            ShareProto = shareProto;
            ShareSize = shareSize;
            Size = size;
            Status = status;
            Id = id;
        }
    }
}