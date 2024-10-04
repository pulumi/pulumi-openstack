// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BGPVPN
{
    /// <summary>
    /// Manages a V2 BGP VPN service resource within OpenStack.
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
    ///     var bgpvpn1 = new OpenStack.BGPVPN.V2("bgpvpn_1", new()
    ///     {
    ///         Name = "bgpvpn1",
    ///         RouteDistinguishers = new[]
    ///         {
    ///             "64512:1",
    ///         },
    ///         RouteTargets = new[]
    ///         {
    ///             "64512:1",
    ///         },
    ///         ImportTargets = new[]
    ///         {
    ///             "64512:2",
    ///         },
    ///         ExportTargets = new[]
    ///         {
    ///             "64512:3",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BGP VPNs can be imported using the `id`, e.g.
    /// 
    /// hcl
    /// 
    /// ```sh
    /// $ pulumi import openstack:bgpvpn/v2:V2 bgpvpn_1 1eec2c66-6be2-4305-af3f-354c9b81f18c
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:bgpvpn/v2:V2")]
    public partial class V2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of additional Route Targets that will be
        /// used for export.
        /// </summary>
        [Output("exportTargets")]
        public Output<ImmutableArray<string>> ExportTargets { get; private set; } = null!;

        /// <summary>
        /// A list of additional Route Targets that will be
        /// imported.
        /// </summary>
        [Output("importTargets")]
        public Output<ImmutableArray<string>> ImportTargets { get; private set; } = null!;

        /// <summary>
        /// The default BGP LOCAL\_PREF of routes that will be
        /// advertised to the BGP VPN, unless overridden per-route.
        /// </summary>
        [Output("localPref")]
        public Output<int?> LocalPref { get; private set; } = null!;

        /// <summary>
        /// The name of the BGP VPN. Changing this updates the name of
        /// the existing BGP VPN.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of network IDs that are associated with the BGP VPN.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<string>> Networks { get; private set; } = null!;

        /// <summary>
        /// A list of port IDs that are associated with the BGP VPN.
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<string>> Ports { get; private set; } = null!;

        /// <summary>
        /// The ID of the project that owns the BGPVPN. Only
        /// administrative and users with `advsvc` role can specify a project ID other
        /// than their own. Changing this creates a new BGP VPN.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// BGP VPN.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A list of route distinguisher strings. If
        /// specified, one of these RDs will be used to advertise VPN routes.
        /// </summary>
        [Output("routeDistinguishers")]
        public Output<ImmutableArray<string>> RouteDistinguishers { get; private set; } = null!;

        /// <summary>
        /// A list of Route Targets that will be both
        /// imported and used for export.
        /// </summary>
        [Output("routeTargets")]
        public Output<ImmutableArray<string>> RouteTargets { get; private set; } = null!;

        /// <summary>
        /// A list of router IDs that are associated with the BGP VPN.
        /// </summary>
        [Output("routers")]
        public Output<ImmutableArray<string>> Routers { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the BGP VPN is shared across projects.
        /// </summary>
        [Output("shared")]
        public Output<bool> Shared { get; private set; } = null!;

        /// <summary>
        /// The type of the BGP VPN (either `l2` or `l3`). Changing this
        /// creates a new BGP VPN. Defaults to `l3`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The globally-assigned VXLAN VNI for the BGP VPN. Changing
        /// this creates a new BGP VPN.
        /// </summary>
        [Output("vni")]
        public Output<int?> Vni { get; private set; } = null!;


        /// <summary>
        /// Create a V2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public V2(string name, V2Args? args = null, CustomResourceOptions? options = null)
            : base("openstack:bgpvpn/v2:V2", name, args ?? new V2Args(), MakeResourceOptions(options, ""))
        {
        }

        private V2(string name, Input<string> id, V2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:bgpvpn/v2:V2", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "openstack:index/bgpvpnV2:BgpvpnV2" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing V2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static V2 Get(string name, Input<string> id, V2State? state = null, CustomResourceOptions? options = null)
        {
            return new V2(name, id, state, options);
        }
    }

    public sealed class V2Args : global::Pulumi.ResourceArgs
    {
        [Input("exportTargets")]
        private InputList<string>? _exportTargets;

        /// <summary>
        /// A list of additional Route Targets that will be
        /// used for export.
        /// </summary>
        public InputList<string> ExportTargets
        {
            get => _exportTargets ?? (_exportTargets = new InputList<string>());
            set => _exportTargets = value;
        }

        [Input("importTargets")]
        private InputList<string>? _importTargets;

        /// <summary>
        /// A list of additional Route Targets that will be
        /// imported.
        /// </summary>
        public InputList<string> ImportTargets
        {
            get => _importTargets ?? (_importTargets = new InputList<string>());
            set => _importTargets = value;
        }

        /// <summary>
        /// The default BGP LOCAL\_PREF of routes that will be
        /// advertised to the BGP VPN, unless overridden per-route.
        /// </summary>
        [Input("localPref")]
        public Input<int>? LocalPref { get; set; }

        /// <summary>
        /// The name of the BGP VPN. Changing this updates the name of
        /// the existing BGP VPN.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project that owns the BGPVPN. Only
        /// administrative and users with `advsvc` role can specify a project ID other
        /// than their own. Changing this creates a new BGP VPN.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// BGP VPN.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("routeDistinguishers")]
        private InputList<string>? _routeDistinguishers;

        /// <summary>
        /// A list of route distinguisher strings. If
        /// specified, one of these RDs will be used to advertise VPN routes.
        /// </summary>
        public InputList<string> RouteDistinguishers
        {
            get => _routeDistinguishers ?? (_routeDistinguishers = new InputList<string>());
            set => _routeDistinguishers = value;
        }

        [Input("routeTargets")]
        private InputList<string>? _routeTargets;

        /// <summary>
        /// A list of Route Targets that will be both
        /// imported and used for export.
        /// </summary>
        public InputList<string> RouteTargets
        {
            get => _routeTargets ?? (_routeTargets = new InputList<string>());
            set => _routeTargets = value;
        }

        /// <summary>
        /// The type of the BGP VPN (either `l2` or `l3`). Changing this
        /// creates a new BGP VPN. Defaults to `l3`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The globally-assigned VXLAN VNI for the BGP VPN. Changing
        /// this creates a new BGP VPN.
        /// </summary>
        [Input("vni")]
        public Input<int>? Vni { get; set; }

        public V2Args()
        {
        }
        public static new V2Args Empty => new V2Args();
    }

    public sealed class V2State : global::Pulumi.ResourceArgs
    {
        [Input("exportTargets")]
        private InputList<string>? _exportTargets;

        /// <summary>
        /// A list of additional Route Targets that will be
        /// used for export.
        /// </summary>
        public InputList<string> ExportTargets
        {
            get => _exportTargets ?? (_exportTargets = new InputList<string>());
            set => _exportTargets = value;
        }

        [Input("importTargets")]
        private InputList<string>? _importTargets;

        /// <summary>
        /// A list of additional Route Targets that will be
        /// imported.
        /// </summary>
        public InputList<string> ImportTargets
        {
            get => _importTargets ?? (_importTargets = new InputList<string>());
            set => _importTargets = value;
        }

        /// <summary>
        /// The default BGP LOCAL\_PREF of routes that will be
        /// advertised to the BGP VPN, unless overridden per-route.
        /// </summary>
        [Input("localPref")]
        public Input<int>? LocalPref { get; set; }

        /// <summary>
        /// The name of the BGP VPN. Changing this updates the name of
        /// the existing BGP VPN.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networks")]
        private InputList<string>? _networks;

        /// <summary>
        /// A list of network IDs that are associated with the BGP VPN.
        /// </summary>
        public InputList<string> Networks
        {
            get => _networks ?? (_networks = new InputList<string>());
            set => _networks = value;
        }

        [Input("ports")]
        private InputList<string>? _ports;

        /// <summary>
        /// A list of port IDs that are associated with the BGP VPN.
        /// </summary>
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        /// <summary>
        /// The ID of the project that owns the BGPVPN. Only
        /// administrative and users with `advsvc` role can specify a project ID other
        /// than their own. Changing this creates a new BGP VPN.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a BGP VPN service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// BGP VPN.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("routeDistinguishers")]
        private InputList<string>? _routeDistinguishers;

        /// <summary>
        /// A list of route distinguisher strings. If
        /// specified, one of these RDs will be used to advertise VPN routes.
        /// </summary>
        public InputList<string> RouteDistinguishers
        {
            get => _routeDistinguishers ?? (_routeDistinguishers = new InputList<string>());
            set => _routeDistinguishers = value;
        }

        [Input("routeTargets")]
        private InputList<string>? _routeTargets;

        /// <summary>
        /// A list of Route Targets that will be both
        /// imported and used for export.
        /// </summary>
        public InputList<string> RouteTargets
        {
            get => _routeTargets ?? (_routeTargets = new InputList<string>());
            set => _routeTargets = value;
        }

        [Input("routers")]
        private InputList<string>? _routers;

        /// <summary>
        /// A list of router IDs that are associated with the BGP VPN.
        /// </summary>
        public InputList<string> Routers
        {
            get => _routers ?? (_routers = new InputList<string>());
            set => _routers = value;
        }

        /// <summary>
        /// Indicates whether the BGP VPN is shared across projects.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        /// <summary>
        /// The type of the BGP VPN (either `l2` or `l3`). Changing this
        /// creates a new BGP VPN. Defaults to `l3`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The globally-assigned VXLAN VNI for the BGP VPN. Changing
        /// this creates a new BGP VPN.
        /// </summary>
        [Input("vni")]
        public Input<int>? Vni { get; set; }

        public V2State()
        {
        }
        public static new V2State Empty => new V2State();
    }
}
