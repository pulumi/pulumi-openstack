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
    /// ## Import
    /// 
    /// This resource can be imported by specifying the ID of the share and the ID of the share access, separated by a slash, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:sharedfilesystem/shareAccess:ShareAccess share_access_1 &lt;share id&gt;/&lt;share access id&gt;
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:sharedfilesystem/shareAccess:ShareAccess")]
    public partial class ShareAccess : Pulumi.CustomResource
    {
        /// <summary>
        /// The access credential of the entity granted access.
        /// </summary>
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// The access level to the share. Can either be `rw` or `ro`.
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// The value that defines the access. Can either be an IP
        /// address or a username verified by configured Security Service of the Share Network.
        /// </summary>
        [Output("accessTo")]
        public Output<string> AccessTo { get; private set; } = null!;

        /// <summary>
        /// The access rule type. Can either be an ip, user,
        /// cert, or cephx. cephx support requires an OpenStack environment that supports
        /// Shared Filesystem microversion 2.13 (Mitaka) or later.
        /// </summary>
        [Output("accessType")]
        public Output<string> AccessType { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a share access. Changing this
        /// creates a new share access.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The UUID of the share to which you are granted access.
        /// </summary>
        [Output("shareId")]
        public Output<string> ShareId { get; private set; } = null!;


        /// <summary>
        /// Create a ShareAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ShareAccess(string name, ShareAccessArgs args, CustomResourceOptions? options = null)
            : base("openstack:sharedfilesystem/shareAccess:ShareAccess", name, args ?? new ShareAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ShareAccess(string name, Input<string> id, ShareAccessState? state = null, CustomResourceOptions? options = null)
            : base("openstack:sharedfilesystem/shareAccess:ShareAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ShareAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ShareAccess Get(string name, Input<string> id, ShareAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new ShareAccess(name, id, state, options);
        }
    }

    public sealed class ShareAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access level to the share. Can either be `rw` or `ro`.
        /// </summary>
        [Input("accessLevel", required: true)]
        public Input<string> AccessLevel { get; set; } = null!;

        /// <summary>
        /// The value that defines the access. Can either be an IP
        /// address or a username verified by configured Security Service of the Share Network.
        /// </summary>
        [Input("accessTo", required: true)]
        public Input<string> AccessTo { get; set; } = null!;

        /// <summary>
        /// The access rule type. Can either be an ip, user,
        /// cert, or cephx. cephx support requires an OpenStack environment that supports
        /// Shared Filesystem microversion 2.13 (Mitaka) or later.
        /// </summary>
        [Input("accessType", required: true)]
        public Input<string> AccessType { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a share access. Changing this
        /// creates a new share access.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The UUID of the share to which you are granted access.
        /// </summary>
        [Input("shareId", required: true)]
        public Input<string> ShareId { get; set; } = null!;

        public ShareAccessArgs()
        {
        }
    }

    public sealed class ShareAccessState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access credential of the entity granted access.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// The access level to the share. Can either be `rw` or `ro`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The value that defines the access. Can either be an IP
        /// address or a username verified by configured Security Service of the Share Network.
        /// </summary>
        [Input("accessTo")]
        public Input<string>? AccessTo { get; set; }

        /// <summary>
        /// The access rule type. Can either be an ip, user,
        /// cert, or cephx. cephx support requires an OpenStack environment that supports
        /// Shared Filesystem microversion 2.13 (Mitaka) or later.
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a share access. Changing this
        /// creates a new share access.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The UUID of the share to which you are granted access.
        /// </summary>
        [Input("shareId")]
        public Input<string>? ShareId { get; set; }

        public ShareAccessState()
        {
        }
    }
}
