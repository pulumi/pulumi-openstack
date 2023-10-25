// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.VPNaaS
{
    /// <summary>
    /// Manages a V2 Neutron IPSec site connection resource within OpenStack.
    /// 
    /// ## Import
    /// 
    /// Site Connections can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:vpnaas/siteConnection:SiteConnection conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:vpnaas/siteConnection:SiteConnection")]
    public partial class SiteConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the resource. Can either be up(true) or down(false).
        /// Changing this updates the administrative state of the existing connection.
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The human-readable description for the connection.
        /// Changing this updates the description of the existing connection.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A dictionary with dead peer detection (DPD) protocol controls.
        /// </summary>
        [Output("dpds")]
        public Output<ImmutableArray<Outputs.SiteConnectionDpd>> Dpds { get; private set; } = null!;

        /// <summary>
        /// The ID of the IKE policy. Changing this creates a new connection.
        /// </summary>
        [Output("ikepolicyId")]
        public Output<string> IkepolicyId { get; private set; } = null!;

        /// <summary>
        /// A valid value is response-only or bi-directional. Default is bi-directional.
        /// </summary>
        [Output("initiator")]
        public Output<string> Initiator { get; private set; } = null!;

        /// <summary>
        /// The ID of the IPsec policy. Changing this creates a new connection.
        /// </summary>
        [Output("ipsecpolicyId")]
        public Output<string> IpsecpolicyId { get; private set; } = null!;

        /// <summary>
        /// The ID for the endpoint group that contains private subnets for the local side of the connection.
        /// You must specify this parameter with the peer_ep_group_id parameter unless
        /// in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
        /// Changing this updates the existing connection.
        /// </summary>
        [Output("localEpGroupId")]
        public Output<string?> LocalEpGroupId { get; private set; } = null!;

        /// <summary>
        /// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
        /// Most often, local ID would be domain name, email address, etc.
        /// If this is not configured then the external IP address will be used as the ID.
        /// </summary>
        [Output("localId")]
        public Output<string?> LocalId { get; private set; } = null!;

        /// <summary>
        /// The maximum transmission unit (MTU) value to address fragmentation.
        /// Minimum value is 68 for IPv4, and 1280 for IPv6.
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;

        /// <summary>
        /// The name of the connection. Changing this updates the name of
        /// the existing connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The peer gateway public IPv4 or IPv6 address or FQDN.
        /// </summary>
        [Output("peerAddress")]
        public Output<string> PeerAddress { get; private set; } = null!;

        /// <summary>
        /// Unique list of valid peer private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; .
        /// </summary>
        [Output("peerCidrs")]
        public Output<ImmutableArray<string>> PeerCidrs { get; private set; } = null!;

        /// <summary>
        /// The ID for the endpoint group that contains private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; for the peer side of the connection.
        /// You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
        /// where peer_cidrs is provided with a subnet_id for the VPN service.
        /// </summary>
        [Output("peerEpGroupId")]
        public Output<string?> PeerEpGroupId { get; private set; } = null!;

        /// <summary>
        /// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
        /// Typically, this value matches the peer_address value.
        /// Changing this updates the existing policy.
        /// </summary>
        [Output("peerId")]
        public Output<string> PeerId { get; private set; } = null!;

        /// <summary>
        /// The pre-shared key. A valid value is any string.
        /// </summary>
        [Output("psk")]
        public Output<string> Psk { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an IPSec site connection. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// site connection.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The owner of the connection. Required if admin wants to
        /// create a connection for another project. Changing this creates a new connection.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPN service. Changing this creates a new connection.
        /// </summary>
        [Output("vpnserviceId")]
        public Output<string> VpnserviceId { get; private set; } = null!;


        /// <summary>
        /// Create a SiteConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SiteConnection(string name, SiteConnectionArgs args, CustomResourceOptions? options = null)
            : base("openstack:vpnaas/siteConnection:SiteConnection", name, args ?? new SiteConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SiteConnection(string name, Input<string> id, SiteConnectionState? state = null, CustomResourceOptions? options = null)
            : base("openstack:vpnaas/siteConnection:SiteConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SiteConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SiteConnection Get(string name, Input<string> id, SiteConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new SiteConnection(name, id, state, options);
        }
    }

    public sealed class SiteConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the resource. Can either be up(true) or down(false).
        /// Changing this updates the administrative state of the existing connection.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The human-readable description for the connection.
        /// Changing this updates the description of the existing connection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dpds")]
        private InputList<Inputs.SiteConnectionDpdArgs>? _dpds;

        /// <summary>
        /// A dictionary with dead peer detection (DPD) protocol controls.
        /// </summary>
        public InputList<Inputs.SiteConnectionDpdArgs> Dpds
        {
            get => _dpds ?? (_dpds = new InputList<Inputs.SiteConnectionDpdArgs>());
            set => _dpds = value;
        }

        /// <summary>
        /// The ID of the IKE policy. Changing this creates a new connection.
        /// </summary>
        [Input("ikepolicyId", required: true)]
        public Input<string> IkepolicyId { get; set; } = null!;

        /// <summary>
        /// A valid value is response-only or bi-directional. Default is bi-directional.
        /// </summary>
        [Input("initiator")]
        public Input<string>? Initiator { get; set; }

        /// <summary>
        /// The ID of the IPsec policy. Changing this creates a new connection.
        /// </summary>
        [Input("ipsecpolicyId", required: true)]
        public Input<string> IpsecpolicyId { get; set; } = null!;

        /// <summary>
        /// The ID for the endpoint group that contains private subnets for the local side of the connection.
        /// You must specify this parameter with the peer_ep_group_id parameter unless
        /// in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
        /// Changing this updates the existing connection.
        /// </summary>
        [Input("localEpGroupId")]
        public Input<string>? LocalEpGroupId { get; set; }

        /// <summary>
        /// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
        /// Most often, local ID would be domain name, email address, etc.
        /// If this is not configured then the external IP address will be used as the ID.
        /// </summary>
        [Input("localId")]
        public Input<string>? LocalId { get; set; }

        /// <summary>
        /// The maximum transmission unit (MTU) value to address fragmentation.
        /// Minimum value is 68 for IPv4, and 1280 for IPv6.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// The name of the connection. Changing this updates the name of
        /// the existing connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The peer gateway public IPv4 or IPv6 address or FQDN.
        /// </summary>
        [Input("peerAddress", required: true)]
        public Input<string> PeerAddress { get; set; } = null!;

        [Input("peerCidrs")]
        private InputList<string>? _peerCidrs;

        /// <summary>
        /// Unique list of valid peer private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; .
        /// </summary>
        public InputList<string> PeerCidrs
        {
            get => _peerCidrs ?? (_peerCidrs = new InputList<string>());
            set => _peerCidrs = value;
        }

        /// <summary>
        /// The ID for the endpoint group that contains private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; for the peer side of the connection.
        /// You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
        /// where peer_cidrs is provided with a subnet_id for the VPN service.
        /// </summary>
        [Input("peerEpGroupId")]
        public Input<string>? PeerEpGroupId { get; set; }

        /// <summary>
        /// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
        /// Typically, this value matches the peer_address value.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("peerId", required: true)]
        public Input<string> PeerId { get; set; } = null!;

        /// <summary>
        /// The pre-shared key. A valid value is any string.
        /// </summary>
        [Input("psk", required: true)]
        public Input<string> Psk { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an IPSec site connection. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// site connection.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the connection. Required if admin wants to
        /// create a connection for another project. Changing this creates a new connection.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// The ID of the VPN service. Changing this creates a new connection.
        /// </summary>
        [Input("vpnserviceId", required: true)]
        public Input<string> VpnserviceId { get; set; } = null!;

        public SiteConnectionArgs()
        {
        }
        public static new SiteConnectionArgs Empty => new SiteConnectionArgs();
    }

    public sealed class SiteConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the resource. Can either be up(true) or down(false).
        /// Changing this updates the administrative state of the existing connection.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The human-readable description for the connection.
        /// Changing this updates the description of the existing connection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dpds")]
        private InputList<Inputs.SiteConnectionDpdGetArgs>? _dpds;

        /// <summary>
        /// A dictionary with dead peer detection (DPD) protocol controls.
        /// </summary>
        public InputList<Inputs.SiteConnectionDpdGetArgs> Dpds
        {
            get => _dpds ?? (_dpds = new InputList<Inputs.SiteConnectionDpdGetArgs>());
            set => _dpds = value;
        }

        /// <summary>
        /// The ID of the IKE policy. Changing this creates a new connection.
        /// </summary>
        [Input("ikepolicyId")]
        public Input<string>? IkepolicyId { get; set; }

        /// <summary>
        /// A valid value is response-only or bi-directional. Default is bi-directional.
        /// </summary>
        [Input("initiator")]
        public Input<string>? Initiator { get; set; }

        /// <summary>
        /// The ID of the IPsec policy. Changing this creates a new connection.
        /// </summary>
        [Input("ipsecpolicyId")]
        public Input<string>? IpsecpolicyId { get; set; }

        /// <summary>
        /// The ID for the endpoint group that contains private subnets for the local side of the connection.
        /// You must specify this parameter with the peer_ep_group_id parameter unless
        /// in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
        /// Changing this updates the existing connection.
        /// </summary>
        [Input("localEpGroupId")]
        public Input<string>? LocalEpGroupId { get; set; }

        /// <summary>
        /// An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
        /// Most often, local ID would be domain name, email address, etc.
        /// If this is not configured then the external IP address will be used as the ID.
        /// </summary>
        [Input("localId")]
        public Input<string>? LocalId { get; set; }

        /// <summary>
        /// The maximum transmission unit (MTU) value to address fragmentation.
        /// Minimum value is 68 for IPv4, and 1280 for IPv6.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// The name of the connection. Changing this updates the name of
        /// the existing connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The peer gateway public IPv4 or IPv6 address or FQDN.
        /// </summary>
        [Input("peerAddress")]
        public Input<string>? PeerAddress { get; set; }

        [Input("peerCidrs")]
        private InputList<string>? _peerCidrs;

        /// <summary>
        /// Unique list of valid peer private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; .
        /// </summary>
        public InputList<string> PeerCidrs
        {
            get => _peerCidrs ?? (_peerCidrs = new InputList<string>());
            set => _peerCidrs = value;
        }

        /// <summary>
        /// The ID for the endpoint group that contains private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; for the peer side of the connection.
        /// You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
        /// where peer_cidrs is provided with a subnet_id for the VPN service.
        /// </summary>
        [Input("peerEpGroupId")]
        public Input<string>? PeerEpGroupId { get; set; }

        /// <summary>
        /// The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
        /// Typically, this value matches the peer_address value.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("peerId")]
        public Input<string>? PeerId { get; set; }

        /// <summary>
        /// The pre-shared key. A valid value is any string.
        /// </summary>
        [Input("psk")]
        public Input<string>? Psk { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an IPSec site connection. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// site connection.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the connection. Required if admin wants to
        /// create a connection for another project. Changing this creates a new connection.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// The ID of the VPN service. Changing this creates a new connection.
        /// </summary>
        [Input("vpnserviceId")]
        public Input<string>? VpnserviceId { get; set; }

        public SiteConnectionState()
        {
        }
        public static new SiteConnectionState Empty => new SiteConnectionState();
    }
}
