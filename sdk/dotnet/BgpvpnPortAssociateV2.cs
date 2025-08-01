// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack
{
    /// <summary>
    /// Manages a V2 BGP VPN port association resource within OpenStack.
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
    ///     var association1 = new OpenStack.BGPVPN.PortAssociateV2("association_1", new()
    ///     {
    ///         BgpvpnId = "19382ec5-8098-47d9-a9c6-6270c91103f4",
    ///         PortId = "b83a95b8-c2c8-4eac-9a9e-ddc85bd1266f",
    ///         Routes = new[]
    ///         {
    ///             new OpenStack.BGPVPN.Inputs.PortAssociateV2RouteArgs
    ///             {
    ///                 Type = "prefix",
    ///                 Prefix = "192.168.170.1/32",
    ///             },
    ///             new OpenStack.BGPVPN.Inputs.PortAssociateV2RouteArgs
    ///             {
    ///                 Type = "bgpvpn",
    ///                 BgpvpnId = "35af1cc6-3d0f-4c5d-86f8-8cdb508d3f0c",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BGP VPN port associations can be imported using the BGP VPN ID and BGP VPN port
    /// 
    /// association ID separated by a slash, e.g.:
    /// 
    /// hcl
    /// 
    /// ```sh
    /// $ pulumi import openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2 association_1 5bb44ecf-f8fe-4d75-8fc5-313f96ee2696/8f8fc660-3f28-414e-896a-0c7c51162fcf
    /// ```
    /// </summary>
    [Obsolete(@"openstack.index/bgpvpnportassociatev2.BgpvpnPortAssociateV2 has been deprecated in favor of openstack.bgpvpn/portassociatev2.PortAssociateV2")]
    [OpenStackResourceType("openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2")]
    public partial class BgpvpnPortAssociateV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A boolean flag indicating whether fixed
        /// IPs should be advertised. Defaults to true.
        /// </summary>
        [Output("advertiseFixedIps")]
        public Output<bool> AdvertiseFixedIps { get; private set; } = null!;

        /// <summary>
        /// The ID of the BGP VPN to which the port will be
        /// associated. Changing this creates a new BGP VPN port association.
        /// </summary>
        [Output("bgpvpnId")]
        public Output<string> BgpvpnId { get; private set; } = null!;

        /// <summary>
        /// The ID of the port to be associated with the BGP VPN.
        /// Changing this creates a new BGP VPN port association.
        /// </summary>
        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project that owns the port
        /// association. Only administrative and users with `advsvc` role can specify a
        /// project ID other than their own. Changing this creates a new BGP VPN port
        /// association.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN port association. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new BGP VPN port association.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A list of dictionaries containing the following keys:
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<Outputs.BgpvpnPortAssociateV2Route>> Routes { get; private set; } = null!;


        /// <summary>
        /// Create a BgpvpnPortAssociateV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BgpvpnPortAssociateV2(string name, BgpvpnPortAssociateV2Args args, CustomResourceOptions? options = null)
            : base("openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2", name, args ?? new BgpvpnPortAssociateV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private BgpvpnPortAssociateV2(string name, Input<string> id, BgpvpnPortAssociateV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BgpvpnPortAssociateV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BgpvpnPortAssociateV2 Get(string name, Input<string> id, BgpvpnPortAssociateV2State? state = null, CustomResourceOptions? options = null)
        {
            return new BgpvpnPortAssociateV2(name, id, state, options);
        }
    }

    public sealed class BgpvpnPortAssociateV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean flag indicating whether fixed
        /// IPs should be advertised. Defaults to true.
        /// </summary>
        [Input("advertiseFixedIps")]
        public Input<bool>? AdvertiseFixedIps { get; set; }

        /// <summary>
        /// The ID of the BGP VPN to which the port will be
        /// associated. Changing this creates a new BGP VPN port association.
        /// </summary>
        [Input("bgpvpnId", required: true)]
        public Input<string> BgpvpnId { get; set; } = null!;

        /// <summary>
        /// The ID of the port to be associated with the BGP VPN.
        /// Changing this creates a new BGP VPN port association.
        /// </summary>
        [Input("portId", required: true)]
        public Input<string> PortId { get; set; } = null!;

        /// <summary>
        /// The ID of the project that owns the port
        /// association. Only administrative and users with `advsvc` role can specify a
        /// project ID other than their own. Changing this creates a new BGP VPN port
        /// association.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN port association. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new BGP VPN port association.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("routes")]
        private InputList<Inputs.BgpvpnPortAssociateV2RouteArgs>? _routes;

        /// <summary>
        /// A list of dictionaries containing the following keys:
        /// </summary>
        public InputList<Inputs.BgpvpnPortAssociateV2RouteArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.BgpvpnPortAssociateV2RouteArgs>());
            set => _routes = value;
        }

        public BgpvpnPortAssociateV2Args()
        {
        }
        public static new BgpvpnPortAssociateV2Args Empty => new BgpvpnPortAssociateV2Args();
    }

    public sealed class BgpvpnPortAssociateV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean flag indicating whether fixed
        /// IPs should be advertised. Defaults to true.
        /// </summary>
        [Input("advertiseFixedIps")]
        public Input<bool>? AdvertiseFixedIps { get; set; }

        /// <summary>
        /// The ID of the BGP VPN to which the port will be
        /// associated. Changing this creates a new BGP VPN port association.
        /// </summary>
        [Input("bgpvpnId")]
        public Input<string>? BgpvpnId { get; set; }

        /// <summary>
        /// The ID of the port to be associated with the BGP VPN.
        /// Changing this creates a new BGP VPN port association.
        /// </summary>
        [Input("portId")]
        public Input<string>? PortId { get; set; }

        /// <summary>
        /// The ID of the project that owns the port
        /// association. Only administrative and users with `advsvc` role can specify a
        /// project ID other than their own. Changing this creates a new BGP VPN port
        /// association.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN port association. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new BGP VPN port association.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("routes")]
        private InputList<Inputs.BgpvpnPortAssociateV2RouteGetArgs>? _routes;

        /// <summary>
        /// A list of dictionaries containing the following keys:
        /// </summary>
        public InputList<Inputs.BgpvpnPortAssociateV2RouteGetArgs> Routes
        {
            get => _routes ?? (_routes = new InputList<Inputs.BgpvpnPortAssociateV2RouteGetArgs>());
            set => _routes = value;
        }

        public BgpvpnPortAssociateV2State()
        {
        }
        public static new BgpvpnPortAssociateV2State Empty => new BgpvpnPortAssociateV2State();
    }
}
