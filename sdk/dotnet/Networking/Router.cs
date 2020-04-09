// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    /// <summary>
    /// Manages a V2 router resource within OpenStack.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_router_v2.html.markdown.
    /// </summary>
    public partial class Router : Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative up/down status for the router
        /// (must be "true" or "false" if provided). Changing this updates the
        /// `admin_state_up` of an existing router.
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The collection of tags assigned on the router, which have been
        /// explicitly and implicitly added.
        /// </summary>
        [Output("allTags")]
        public Output<ImmutableArray<string>> AllTags { get; private set; } = null!;

        /// <summary>
        /// An availability zone is used to make 
        /// network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
        /// this creates a new router.
        /// </summary>
        [Output("availabilityZoneHints")]
        public Output<ImmutableArray<string>> AvailabilityZoneHints { get; private set; } = null!;

        /// <summary>
        /// Human-readable description for the router.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates whether or not to create a
        /// distributed router. The default policy setting in Neutron restricts
        /// usage of this property to administrative users only.
        /// </summary>
        [Output("distributed")]
        public Output<bool> Distributed { get; private set; } = null!;

        /// <summary>
        /// Enable Source NAT for the router. Valid values are
        /// "true" or "false". An `external_network_id` has to be set in order to
        /// set this property. Changing this updates the `enable_snat` of the router.
        /// Setting this value **requires** an **ext-gw-mode** extension to be enabled
        /// in OpenStack Neutron.
        /// </summary>
        [Output("enableSnat")]
        public Output<bool> EnableSnat { get; private set; } = null!;

        /// <summary>
        /// An external fixed IP for the router. This
        /// can be repeated. The structure is described below. An `external_network_id`
        /// has to be set in order to set this property. Changing this updates the
        /// external fixed IPs of the router.
        /// </summary>
        [Output("externalFixedIps")]
        public Output<ImmutableArray<Outputs.RouterExternalFixedIps>> ExternalFixedIps { get; private set; } = null!;

        /// <summary>
        /// The
        /// network UUID of an external gateway for the router. A router with an
        /// external gateway is required if any compute instances or load balancers
        /// will be using floating IPs. Changing this updates the external gateway
        /// of an existing router.
        /// </summary>
        [Output("externalGateway")]
        public Output<string> ExternalGateway { get; private set; } = null!;

        /// <summary>
        /// The network UUID of an external gateway
        /// for the router. A router with an external gateway is required if any
        /// compute instances or load balancers will be using floating IPs. Changing
        /// this updates the external gateway of the router.
        /// </summary>
        [Output("externalNetworkId")]
        public Output<string> ExternalNetworkId { get; private set; } = null!;

        /// <summary>
        /// A unique name for the router. Changing this
        /// updates the `name` of an existing router.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a router. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// router.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A set of string tags for the router.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The owner of the floating IP. Required if admin wants
        /// to create a router for another tenant. Changing this creates a new router.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Map of additional driver-specific options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;

        /// <summary>
        /// Map of additional vendor-specific options.
        /// Supported options are described below.
        /// </summary>
        [Output("vendorOptions")]
        public Output<Outputs.RouterVendorOptions?> VendorOptions { get; private set; } = null!;


        /// <summary>
        /// Create a Router resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Router(string name, RouterArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:networking/router:Router", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Router(string name, Input<string> id, RouterState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/router:Router", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Router resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Router Get(string name, Input<string> id, RouterState? state = null, CustomResourceOptions? options = null)
        {
            return new Router(name, id, state, options);
        }
    }

    public sealed class RouterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative up/down status for the router
        /// (must be "true" or "false" if provided). Changing this updates the
        /// `admin_state_up` of an existing router.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("availabilityZoneHints")]
        private InputList<string>? _availabilityZoneHints;

        /// <summary>
        /// An availability zone is used to make 
        /// network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
        /// this creates a new router.
        /// </summary>
        public InputList<string> AvailabilityZoneHints
        {
            get => _availabilityZoneHints ?? (_availabilityZoneHints = new InputList<string>());
            set => _availabilityZoneHints = value;
        }

        /// <summary>
        /// Human-readable description for the router.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether or not to create a
        /// distributed router. The default policy setting in Neutron restricts
        /// usage of this property to administrative users only.
        /// </summary>
        [Input("distributed")]
        public Input<bool>? Distributed { get; set; }

        /// <summary>
        /// Enable Source NAT for the router. Valid values are
        /// "true" or "false". An `external_network_id` has to be set in order to
        /// set this property. Changing this updates the `enable_snat` of the router.
        /// Setting this value **requires** an **ext-gw-mode** extension to be enabled
        /// in OpenStack Neutron.
        /// </summary>
        [Input("enableSnat")]
        public Input<bool>? EnableSnat { get; set; }

        [Input("externalFixedIps")]
        private InputList<Inputs.RouterExternalFixedIpsArgs>? _externalFixedIps;

        /// <summary>
        /// An external fixed IP for the router. This
        /// can be repeated. The structure is described below. An `external_network_id`
        /// has to be set in order to set this property. Changing this updates the
        /// external fixed IPs of the router.
        /// </summary>
        public InputList<Inputs.RouterExternalFixedIpsArgs> ExternalFixedIps
        {
            get => _externalFixedIps ?? (_externalFixedIps = new InputList<Inputs.RouterExternalFixedIpsArgs>());
            set => _externalFixedIps = value;
        }

        /// <summary>
        /// The
        /// network UUID of an external gateway for the router. A router with an
        /// external gateway is required if any compute instances or load balancers
        /// will be using floating IPs. Changing this updates the external gateway
        /// of an existing router.
        /// </summary>
        [Input("externalGateway")]
        public Input<string>? ExternalGateway { get; set; }

        /// <summary>
        /// The network UUID of an external gateway
        /// for the router. A router with an external gateway is required if any
        /// compute instances or load balancers will be using floating IPs. Changing
        /// this updates the external gateway of the router.
        /// </summary>
        [Input("externalNetworkId")]
        public Input<string>? ExternalNetworkId { get; set; }

        /// <summary>
        /// A unique name for the router. Changing this
        /// updates the `name` of an existing router.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a router. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// router.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the router.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the floating IP. Required if admin wants
        /// to create a router for another tenant. Changing this creates a new router.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional driver-specific options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// Map of additional vendor-specific options.
        /// Supported options are described below.
        /// </summary>
        [Input("vendorOptions")]
        public Input<Inputs.RouterVendorOptionsArgs>? VendorOptions { get; set; }

        public RouterArgs()
        {
        }
    }

    public sealed class RouterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative up/down status for the router
        /// (must be "true" or "false" if provided). Changing this updates the
        /// `admin_state_up` of an existing router.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("allTags")]
        private InputList<string>? _allTags;

        /// <summary>
        /// The collection of tags assigned on the router, which have been
        /// explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllTags
        {
            get => _allTags ?? (_allTags = new InputList<string>());
            set => _allTags = value;
        }

        [Input("availabilityZoneHints")]
        private InputList<string>? _availabilityZoneHints;

        /// <summary>
        /// An availability zone is used to make 
        /// network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
        /// this creates a new router.
        /// </summary>
        public InputList<string> AvailabilityZoneHints
        {
            get => _availabilityZoneHints ?? (_availabilityZoneHints = new InputList<string>());
            set => _availabilityZoneHints = value;
        }

        /// <summary>
        /// Human-readable description for the router.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether or not to create a
        /// distributed router. The default policy setting in Neutron restricts
        /// usage of this property to administrative users only.
        /// </summary>
        [Input("distributed")]
        public Input<bool>? Distributed { get; set; }

        /// <summary>
        /// Enable Source NAT for the router. Valid values are
        /// "true" or "false". An `external_network_id` has to be set in order to
        /// set this property. Changing this updates the `enable_snat` of the router.
        /// Setting this value **requires** an **ext-gw-mode** extension to be enabled
        /// in OpenStack Neutron.
        /// </summary>
        [Input("enableSnat")]
        public Input<bool>? EnableSnat { get; set; }

        [Input("externalFixedIps")]
        private InputList<Inputs.RouterExternalFixedIpsGetArgs>? _externalFixedIps;

        /// <summary>
        /// An external fixed IP for the router. This
        /// can be repeated. The structure is described below. An `external_network_id`
        /// has to be set in order to set this property. Changing this updates the
        /// external fixed IPs of the router.
        /// </summary>
        public InputList<Inputs.RouterExternalFixedIpsGetArgs> ExternalFixedIps
        {
            get => _externalFixedIps ?? (_externalFixedIps = new InputList<Inputs.RouterExternalFixedIpsGetArgs>());
            set => _externalFixedIps = value;
        }

        /// <summary>
        /// The
        /// network UUID of an external gateway for the router. A router with an
        /// external gateway is required if any compute instances or load balancers
        /// will be using floating IPs. Changing this updates the external gateway
        /// of an existing router.
        /// </summary>
        [Input("externalGateway")]
        public Input<string>? ExternalGateway { get; set; }

        /// <summary>
        /// The network UUID of an external gateway
        /// for the router. A router with an external gateway is required if any
        /// compute instances or load balancers will be using floating IPs. Changing
        /// this updates the external gateway of the router.
        /// </summary>
        [Input("externalNetworkId")]
        public Input<string>? ExternalNetworkId { get; set; }

        /// <summary>
        /// A unique name for the router. Changing this
        /// updates the `name` of an existing router.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a router. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// router.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the router.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the floating IP. Required if admin wants
        /// to create a router for another tenant. Changing this creates a new router.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional driver-specific options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// Map of additional vendor-specific options.
        /// Supported options are described below.
        /// </summary>
        [Input("vendorOptions")]
        public Input<Inputs.RouterVendorOptionsGetArgs>? VendorOptions { get; set; }

        public RouterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class RouterExternalFixedIpsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address to set on the router.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Subnet in which the fixed IP belongs to.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public RouterExternalFixedIpsArgs()
        {
        }
    }

    public sealed class RouterExternalFixedIpsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address to set on the router.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Subnet in which the fixed IP belongs to.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public RouterExternalFixedIpsGetArgs()
        {
        }
    }

    public sealed class RouterVendorOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean to control whether
        /// the Router gateway is assigned during creation or updated after creation.
        /// </summary>
        [Input("setRouterGatewayAfterCreate")]
        public Input<bool>? SetRouterGatewayAfterCreate { get; set; }

        public RouterVendorOptionsArgs()
        {
        }
    }

    public sealed class RouterVendorOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean to control whether
        /// the Router gateway is assigned during creation or updated after creation.
        /// </summary>
        [Input("setRouterGatewayAfterCreate")]
        public Input<bool>? SetRouterGatewayAfterCreate { get; set; }

        public RouterVendorOptionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class RouterExternalFixedIps
    {
        /// <summary>
        /// The IP address to set on the router.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// Subnet in which the fixed IP belongs to.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private RouterExternalFixedIps(
            string ipAddress,
            string? subnetId)
        {
            IpAddress = ipAddress;
            SubnetId = subnetId;
        }
    }

    [OutputType]
    public sealed class RouterVendorOptions
    {
        /// <summary>
        /// Boolean to control whether
        /// the Router gateway is assigned during creation or updated after creation.
        /// </summary>
        public readonly bool? SetRouterGatewayAfterCreate;

        [OutputConstructor]
        private RouterVendorOptions(bool? setRouterGatewayAfterCreate)
        {
            SetRouterGatewayAfterCreate = setRouterGatewayAfterCreate;
        }
    }
    }
}
