// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.SharedFileSystem
{
    /// <summary>
    /// Use this resource to configure a share.
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
    ///     var network1 = new OpenStack.Networking.Network("network_1", new()
    ///     {
    ///         Name = "network_1",
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var subnet1 = new OpenStack.Networking.Subnet("subnet_1", new()
    ///     {
    ///         Name = "subnet_1",
    ///         Cidr = "192.168.199.0/24",
    ///         IpVersion = 4,
    ///         NetworkId = network1.Id,
    ///     });
    /// 
    ///     var sharenetwork1 = new OpenStack.SharedFileSystem.ShareNetwork("sharenetwork_1", new()
    ///     {
    ///         Name = "test_sharenetwork",
    ///         Description = "test share network with security services",
    ///         NeutronNetId = network1.Id,
    ///         NeutronSubnetId = subnet1.Id,
    ///     });
    /// 
    ///     var share1 = new OpenStack.SharedFileSystem.Share("share_1", new()
    ///     {
    ///         Name = "nfs_share",
    ///         Description = "test share description",
    ///         ShareProto = "NFS",
    ///         Size = 1,
    ///         ShareNetworkId = sharenetwork1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported by specifying the ID of the share:
    /// 
    /// ```sh
    /// $ pulumi import openstack:sharedfilesystem/share:Share share_1 id
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:sharedfilesystem/share:Share")]
    public partial class Share : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The map of metadata, assigned on the share, which has been
        /// explicitly and implicitly added.
        /// </summary>
        [Output("allMetadata")]
        public Output<ImmutableDictionary<string, object>> AllMetadata { get; private set; } = null!;

        /// <summary>
        /// The share availability zone. Changing this creates a
        /// new share.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The human-readable description for the share.
        /// Changing this updates the description of the existing share.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of export locations. For example, when a share server
        /// has more than one network interface, it can have multiple export locations.
        /// </summary>
        [Output("exportLocations")]
        public Output<ImmutableArray<Outputs.ShareExportLocation>> ExportLocations { get; private set; } = null!;

        /// <summary>
        /// Indicates whether a share has replicas or not.
        /// </summary>
        [Output("hasReplicas")]
        public Output<bool> HasReplicas { get; private set; } = null!;

        /// <summary>
        /// The share host name.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// The level of visibility for the share. Set to true to make
        /// share public. Set to false to make it private. Default value is false. Changing this
        /// updates the existing share.
        /// </summary>
        [Output("isPublic")]
        public Output<bool?> IsPublic { get; private set; } = null!;

        /// <summary>
        /// One or more metadata key and value pairs as a dictionary of
        /// strings.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The name of the share. Changing this updates the name
        /// of the existing share.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The owner of the Share.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a share. Changing this
        /// creates a new share.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The share replication type.
        /// </summary>
        [Output("replicationType")]
        public Output<string> ReplicationType { get; private set; } = null!;

        /// <summary>
        /// The UUID of a share network where the share server exists
        /// or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
        /// the share_network_id value from the snapshot is used. Changing this creates a new share.
        /// </summary>
        [Output("shareNetworkId")]
        public Output<string> ShareNetworkId { get; private set; } = null!;

        /// <summary>
        /// The share protocol - can either be NFS, CIFS,
        /// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
        /// </summary>
        [Output("shareProto")]
        public Output<string> ShareProto { get; private set; } = null!;

        /// <summary>
        /// The UUID of the share server.
        /// </summary>
        [Output("shareServerId")]
        public Output<string> ShareServerId { get; private set; } = null!;

        /// <summary>
        /// The share type name. If you omit this parameter, the default
        /// share type is used.
        /// </summary>
        [Output("shareType")]
        public Output<string> ShareType { get; private set; } = null!;

        /// <summary>
        /// The share size, in GBs. The requested share size cannot be greater
        /// than the allowed GB quota. Changing this resizes the existing share.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The UUID of the share's base snapshot. Changing this creates
        /// a new share.
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;


        /// <summary>
        /// Create a Share resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Share(string name, ShareArgs args, CustomResourceOptions? options = null)
            : base("openstack:sharedfilesystem/share:Share", name, args ?? new ShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Share(string name, Input<string> id, ShareState? state = null, CustomResourceOptions? options = null)
            : base("openstack:sharedfilesystem/share:Share", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Share resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Share Get(string name, Input<string> id, ShareState? state = null, CustomResourceOptions? options = null)
        {
            return new Share(name, id, state, options);
        }
    }

    public sealed class ShareArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The share availability zone. Changing this creates a
        /// new share.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The human-readable description for the share.
        /// Changing this updates the description of the existing share.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The level of visibility for the share. Set to true to make
        /// share public. Set to false to make it private. Default value is false. Changing this
        /// updates the existing share.
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
        /// The name of the share. Changing this updates the name
        /// of the existing share.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a share. Changing this
        /// creates a new share.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The UUID of a share network where the share server exists
        /// or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
        /// the share_network_id value from the snapshot is used. Changing this creates a new share.
        /// </summary>
        [Input("shareNetworkId")]
        public Input<string>? ShareNetworkId { get; set; }

        /// <summary>
        /// The share protocol - can either be NFS, CIFS,
        /// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
        /// </summary>
        [Input("shareProto", required: true)]
        public Input<string> ShareProto { get; set; } = null!;

        /// <summary>
        /// The share type name. If you omit this parameter, the default
        /// share type is used.
        /// </summary>
        [Input("shareType")]
        public Input<string>? ShareType { get; set; }

        /// <summary>
        /// The share size, in GBs. The requested share size cannot be greater
        /// than the allowed GB quota. Changing this resizes the existing share.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The UUID of the share's base snapshot. Changing this creates
        /// a new share.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public ShareArgs()
        {
        }
        public static new ShareArgs Empty => new ShareArgs();
    }

    public sealed class ShareState : global::Pulumi.ResourceArgs
    {
        [Input("allMetadata")]
        private InputMap<object>? _allMetadata;

        /// <summary>
        /// The map of metadata, assigned on the share, which has been
        /// explicitly and implicitly added.
        /// </summary>
        public InputMap<object> AllMetadata
        {
            get => _allMetadata ?? (_allMetadata = new InputMap<object>());
            set => _allMetadata = value;
        }

        /// <summary>
        /// The share availability zone. Changing this creates a
        /// new share.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The human-readable description for the share.
        /// Changing this updates the description of the existing share.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("exportLocations")]
        private InputList<Inputs.ShareExportLocationGetArgs>? _exportLocations;

        /// <summary>
        /// A list of export locations. For example, when a share server
        /// has more than one network interface, it can have multiple export locations.
        /// </summary>
        public InputList<Inputs.ShareExportLocationGetArgs> ExportLocations
        {
            get => _exportLocations ?? (_exportLocations = new InputList<Inputs.ShareExportLocationGetArgs>());
            set => _exportLocations = value;
        }

        /// <summary>
        /// Indicates whether a share has replicas or not.
        /// </summary>
        [Input("hasReplicas")]
        public Input<bool>? HasReplicas { get; set; }

        /// <summary>
        /// The share host name.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The level of visibility for the share. Set to true to make
        /// share public. Set to false to make it private. Default value is false. Changing this
        /// updates the existing share.
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
        /// The name of the share. Changing this updates the name
        /// of the existing share.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the Share.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a share. Changing this
        /// creates a new share.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The share replication type.
        /// </summary>
        [Input("replicationType")]
        public Input<string>? ReplicationType { get; set; }

        /// <summary>
        /// The UUID of a share network where the share server exists
        /// or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
        /// the share_network_id value from the snapshot is used. Changing this creates a new share.
        /// </summary>
        [Input("shareNetworkId")]
        public Input<string>? ShareNetworkId { get; set; }

        /// <summary>
        /// The share protocol - can either be NFS, CIFS,
        /// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
        /// </summary>
        [Input("shareProto")]
        public Input<string>? ShareProto { get; set; }

        /// <summary>
        /// The UUID of the share server.
        /// </summary>
        [Input("shareServerId")]
        public Input<string>? ShareServerId { get; set; }

        /// <summary>
        /// The share type name. If you omit this parameter, the default
        /// share type is used.
        /// </summary>
        [Input("shareType")]
        public Input<string>? ShareType { get; set; }

        /// <summary>
        /// The share size, in GBs. The requested share size cannot be greater
        /// than the allowed GB quota. Changing this resizes the existing share.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The UUID of the share's base snapshot. Changing this creates
        /// a new share.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public ShareState()
        {
        }
        public static new ShareState Empty => new ShareState();
    }
}
