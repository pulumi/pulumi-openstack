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
    /// ## Example Usage
    /// 
    /// ### NFS
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network1 = new OpenStack.Networking.Network("network1", new()
    ///     {
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var subnet1 = new OpenStack.Networking.Subnet("subnet1", new()
    ///     {
    ///         Cidr = "192.168.199.0/24",
    ///         IpVersion = 4,
    ///         NetworkId = network1.Id,
    ///     });
    /// 
    ///     var sharenetwork1 = new OpenStack.SharedFileSystem.ShareNetwork("sharenetwork1", new()
    ///     {
    ///         Description = "test share network with security services",
    ///         NeutronNetId = network1.Id,
    ///         NeutronSubnetId = subnet1.Id,
    ///     });
    /// 
    ///     var share1 = new OpenStack.SharedFileSystem.Share("share1", new()
    ///     {
    ///         Description = "test share description",
    ///         ShareProto = "NFS",
    ///         Size = 1,
    ///         ShareNetworkId = sharenetwork1.Id,
    ///     });
    /// 
    ///     var shareAccess1 = new OpenStack.SharedFileSystem.ShareAccess("shareAccess1", new()
    ///     {
    ///         ShareId = share1.Id,
    ///         AccessType = "ip",
    ///         AccessTo = "192.168.199.10",
    ///         AccessLevel = "rw",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### CIFS
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network1 = new OpenStack.Networking.Network("network1", new()
    ///     {
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var subnet1 = new OpenStack.Networking.Subnet("subnet1", new()
    ///     {
    ///         Cidr = "192.168.199.0/24",
    ///         IpVersion = 4,
    ///         NetworkId = network1.Id,
    ///     });
    /// 
    ///     var securityservice1 = new OpenStack.SharedFileSystem.SecurityService("securityservice1", new()
    ///     {
    ///         Description = "created by terraform",
    ///         Type = "active_directory",
    ///         Server = "192.168.199.10",
    ///         DnsIp = "192.168.199.10",
    ///         Domain = "example.com",
    ///         Ou = "CN=Computers,DC=example,DC=com",
    ///         User = "joinDomainUser",
    ///         Password = "s8cret",
    ///     });
    /// 
    ///     var sharenetwork1 = new OpenStack.SharedFileSystem.ShareNetwork("sharenetwork1", new()
    ///     {
    ///         Description = "share the secure love",
    ///         NeutronNetId = network1.Id,
    ///         NeutronSubnetId = subnet1.Id,
    ///         SecurityServiceIds = new[]
    ///         {
    ///             securityservice1.Id,
    ///         },
    ///     });
    /// 
    ///     var share1 = new OpenStack.SharedFileSystem.Share("share1", new()
    ///     {
    ///         ShareProto = "CIFS",
    ///         Size = 1,
    ///         ShareNetworkId = sharenetwork1.Id,
    ///     });
    /// 
    ///     var shareAccess1 = new OpenStack.SharedFileSystem.ShareAccess("shareAccess1", new()
    ///     {
    ///         ShareId = share1.Id,
    ///         AccessType = "user",
    ///         AccessTo = "windows",
    ///         AccessLevel = "ro",
    ///     });
    /// 
    ///     var shareAccess2 = new OpenStack.SharedFileSystem.ShareAccess("shareAccess2", new()
    ///     {
    ///         ShareId = share1.Id,
    ///         AccessType = "user",
    ///         AccessTo = "linux",
    ///         AccessLevel = "rw",
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["exportLocations"] = share1.ExportLocations,
    ///     };
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported by specifying the ID of the share and the ID of the
    /// share access, separated by a slash, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import openstack:sharedfilesystem/shareAccess:ShareAccess share_access_1 share_id/share_access_id
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:sharedfilesystem/shareAccess:ShareAccess")]
    public partial class ShareAccess : global::Pulumi.CustomResource
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
        /// The share access state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


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
                AdditionalSecretOutputs =
                {
                    "accessKey",
                },
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

    public sealed class ShareAccessArgs : global::Pulumi.ResourceArgs
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
        public static new ShareAccessArgs Empty => new ShareAccessArgs();
    }

    public sealed class ShareAccessState : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// The access credential of the entity granted access.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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

        /// <summary>
        /// The share access state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public ShareAccessState()
        {
        }
        public static new ShareAccessState Empty => new ShareAccessState();
    }
}
