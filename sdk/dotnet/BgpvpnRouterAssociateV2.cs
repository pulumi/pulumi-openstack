// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack
{
    /// <summary>
    /// Manages a V2 BGP VPN router association resource within OpenStack.
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
    ///     var association1 = new OpenStack.BgpvpnRouterAssociateV2("association_1", new()
    ///     {
    ///         BgpvpnId = "d57d39e1-dc63-44fd-8cbd-a4e1488100c5",
    ///         RouterId = "423fa80f-e0d7-4d02-a9a5-8b8c05812bf6",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BGP VPN router associations can be imported using the BGP VPN ID and BGP VPN
    /// 
    /// router association ID separated by a slash, e.g.:
    /// 
    /// hcl
    /// 
    /// ```sh
    /// $ pulumi import openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2 association_1 e26d509e-fc2d-4fb5-8562-619911a9a6bc/3cc9df2d-80db-4536-8ba6-295d1d0f723f
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2")]
    public partial class BgpvpnRouterAssociateV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A boolean flag indicating whether extra
        /// routes should be advertised. Defaults to true.
        /// </summary>
        [Output("advertiseExtraRoutes")]
        public Output<bool> AdvertiseExtraRoutes { get; private set; } = null!;

        /// <summary>
        /// The ID of the BGP VPN to which the router will be
        /// associated. Changing this creates a new BGP VPN router association.
        /// </summary>
        [Output("bgpvpnId")]
        public Output<string> BgpvpnId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project that owns the BGP VPN router
        /// association. Only administrative and users with `advsvc` role can specify a
        /// project ID other than their own. Changing this creates a new BGP VPN router
        /// association.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN router association. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new BGP VPN router association.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The ID of the router to be associated with the BGP
        /// VPN. Changing this creates a new BGP VPN router association.
        /// </summary>
        [Output("routerId")]
        public Output<string> RouterId { get; private set; } = null!;


        /// <summary>
        /// Create a BgpvpnRouterAssociateV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BgpvpnRouterAssociateV2(string name, BgpvpnRouterAssociateV2Args args, CustomResourceOptions? options = null)
            : base("openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2", name, args ?? new BgpvpnRouterAssociateV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private BgpvpnRouterAssociateV2(string name, Input<string> id, BgpvpnRouterAssociateV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BgpvpnRouterAssociateV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BgpvpnRouterAssociateV2 Get(string name, Input<string> id, BgpvpnRouterAssociateV2State? state = null, CustomResourceOptions? options = null)
        {
            return new BgpvpnRouterAssociateV2(name, id, state, options);
        }
    }

    public sealed class BgpvpnRouterAssociateV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean flag indicating whether extra
        /// routes should be advertised. Defaults to true.
        /// </summary>
        [Input("advertiseExtraRoutes")]
        public Input<bool>? AdvertiseExtraRoutes { get; set; }

        /// <summary>
        /// The ID of the BGP VPN to which the router will be
        /// associated. Changing this creates a new BGP VPN router association.
        /// </summary>
        [Input("bgpvpnId", required: true)]
        public Input<string> BgpvpnId { get; set; } = null!;

        /// <summary>
        /// The ID of the project that owns the BGP VPN router
        /// association. Only administrative and users with `advsvc` role can specify a
        /// project ID other than their own. Changing this creates a new BGP VPN router
        /// association.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN router association. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new BGP VPN router association.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the router to be associated with the BGP
        /// VPN. Changing this creates a new BGP VPN router association.
        /// </summary>
        [Input("routerId", required: true)]
        public Input<string> RouterId { get; set; } = null!;

        public BgpvpnRouterAssociateV2Args()
        {
        }
        public static new BgpvpnRouterAssociateV2Args Empty => new BgpvpnRouterAssociateV2Args();
    }

    public sealed class BgpvpnRouterAssociateV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean flag indicating whether extra
        /// routes should be advertised. Defaults to true.
        /// </summary>
        [Input("advertiseExtraRoutes")]
        public Input<bool>? AdvertiseExtraRoutes { get; set; }

        /// <summary>
        /// The ID of the BGP VPN to which the router will be
        /// associated. Changing this creates a new BGP VPN router association.
        /// </summary>
        [Input("bgpvpnId")]
        public Input<string>? BgpvpnId { get; set; }

        /// <summary>
        /// The ID of the project that owns the BGP VPN router
        /// association. Only administrative and users with `advsvc` role can specify a
        /// project ID other than their own. Changing this creates a new BGP VPN router
        /// association.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN router association. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new BGP VPN router association.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the router to be associated with the BGP
        /// VPN. Changing this creates a new BGP VPN router association.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        public BgpvpnRouterAssociateV2State()
        {
        }
        public static new BgpvpnRouterAssociateV2State Empty => new BgpvpnRouterAssociateV2State();
    }
}