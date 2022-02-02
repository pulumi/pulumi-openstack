// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.SharedFileSystem
{
    public static class GetShare
    {
        /// <summary>
        /// Use this data source to get the ID of an available Shared File System share.
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
        ///         var share1 = Output.Create(OpenStack.SharedFileSystem.GetShare.InvokeAsync(new OpenStack.SharedFileSystem.GetShareArgs
        ///         {
        ///             Name = "share_1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetShareResult> InvokeAsync(GetShareArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetShareResult>("openstack:sharedfilesystem/getShare:getShare", args ?? new GetShareArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available Shared File System share.
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
        ///         var share1 = Output.Create(OpenStack.SharedFileSystem.GetShare.InvokeAsync(new OpenStack.SharedFileSystem.GetShareArgs
        ///         {
        ///             Name = "share_1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetShareResult> Invoke(GetShareInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetShareResult>("openstack:sharedfilesystem/getShare:getShare", args ?? new GetShareInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetShareArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The human-readable description for the share.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The export location path of the share. Available
        /// since Manila API version 2.35.
        /// </summary>
        [Input("exportLocationPath")]
        public string? ExportLocationPath { get; set; }

        /// <summary>
        /// The level of visibility for the share.
        /// length.
        /// </summary>
        [Input("isPublic")]
        public bool? IsPublic { get; set; }

        [Input("metadata")]
        private Dictionary<string, object>? _metadata;

        /// <summary>
        /// One or more metadata key and value pairs as a dictionary of
        /// strings.
        /// </summary>
        public Dictionary<string, object> Metadata
        {
            get => _metadata ?? (_metadata = new Dictionary<string, object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the share.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The UUID of the share's share network.
        /// </summary>
        [Input("shareNetworkId")]
        public string? ShareNetworkId { get; set; }

        /// <summary>
        /// The UUID of the share's base snapshot.
        /// </summary>
        [Input("snapshotId")]
        public string? SnapshotId { get; set; }

        /// <summary>
        /// A share status filter. A valid value is `creating`,
        /// `error`, `available`, `deleting`, `error_deleting`, `manage_starting`,
        /// `manage_error`, `unmanage_starting`, `unmanage_error`, `unmanaged`,
        /// `extending`, `extending_error`, `shrinking`, `shrinking_error`, or
        /// `shrinking_possible_data_loss_error`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetShareArgs()
        {
        }
    }

    public sealed class GetShareInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The human-readable description for the share.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The export location path of the share. Available
        /// since Manila API version 2.35.
        /// </summary>
        [Input("exportLocationPath")]
        public Input<string>? ExportLocationPath { get; set; }

        /// <summary>
        /// The level of visibility for the share.
        /// length.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// One or more metadata key and value pairs as a dictionary of
        /// strings.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the share.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The UUID of the share's share network.
        /// </summary>
        [Input("shareNetworkId")]
        public Input<string>? ShareNetworkId { get; set; }

        /// <summary>
        /// The UUID of the share's base snapshot.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// A share status filter. A valid value is `creating`,
        /// `error`, `available`, `deleting`, `error_deleting`, `manage_starting`,
        /// `manage_error`, `unmanage_starting`, `unmanage_error`, `unmanaged`,
        /// `extending`, `extending_error`, `shrinking`, `shrinking_error`, or
        /// `shrinking_possible_data_loss_error`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetShareInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetShareResult
    {
        /// <summary>
        /// The share availability zone.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? ExportLocationPath;
        /// <summary>
        /// A list of export locations. For example, when a share
        /// server has more than one network interface, it can have multiple export
        /// locations.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetShareExportLocationResult> ExportLocations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool IsPublic;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Metadata;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ShareNetworkId;
        /// <summary>
        /// The share protocol.
        /// </summary>
        public readonly string ShareProto;
        /// <summary>
        /// The share size, in GBs.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string SnapshotId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetShareResult(
            string availabilityZone,

            string description,

            string? exportLocationPath,

            ImmutableArray<Outputs.GetShareExportLocationResult> exportLocations,

            string id,

            bool isPublic,

            ImmutableDictionary<string, object> metadata,

            string name,

            string projectId,

            string region,

            string shareNetworkId,

            string shareProto,

            int size,

            string snapshotId,

            string status)
        {
            AvailabilityZone = availabilityZone;
            Description = description;
            ExportLocationPath = exportLocationPath;
            ExportLocations = exportLocations;
            Id = id;
            IsPublic = isPublic;
            Metadata = metadata;
            Name = name;
            ProjectId = projectId;
            Region = region;
            ShareNetworkId = shareNetworkId;
            ShareProto = shareProto;
            Size = size;
            SnapshotId = snapshotId;
            Status = status;
        }
    }
}
