// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Loadbalancer
{
    /// <summary>
    /// Manages a V2 listener resource within OpenStack.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_listener_v2.html.markdown.
    /// </summary>
    public partial class Listener : Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the Listener.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The maximum number of connections allowed
        /// for the Listener.
        /// </summary>
        [Output("connectionLimit")]
        public Output<int> ConnectionLimit { get; private set; } = null!;

        /// <summary>
        /// The ID of the default pool with which the
        /// Listener is associated.
        /// </summary>
        [Output("defaultPoolId")]
        public Output<string> DefaultPoolId { get; private set; } = null!;

        /// <summary>
        /// A reference to a Barbican Secrets
        /// container which stores TLS information. This is required if the protocol
        /// is `TERMINATED_HTTPS`. See
        /// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
        /// for more information.
        /// </summary>
        [Output("defaultTlsContainerRef")]
        public Output<string?> DefaultTlsContainerRef { get; private set; } = null!;

        /// <summary>
        /// Human-readable description for the Listener.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The load balancer on which to provision this
        /// Listener. Changing this creates a new Listener.
        /// </summary>
        [Output("loadbalancerId")]
        public Output<string> LoadbalancerId { get; private set; } = null!;

        /// <summary>
        /// Human-readable name for the Listener. Does not have
        /// to be unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The protocol - can either be TCP, HTTP, HTTPS,
        /// TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
        /// new Listener.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The port on which to listen for client traffic.
        /// Changing this creates a new Listener.
        /// </summary>
        [Output("protocolPort")]
        public Output<int> ProtocolPort { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// Listener.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A list of references to Barbican Secrets
        /// containers which store SNI information. See
        /// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
        /// for more information.
        /// </summary>
        [Output("sniContainerRefs")]
        public Output<ImmutableArray<string>> SniContainerRefs { get; private set; } = null!;

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Listener.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new Listener.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The client inactivity timeout in milliseconds.
        /// </summary>
        [Output("timeoutClientData")]
        public Output<int> TimeoutClientData { get; private set; } = null!;

        /// <summary>
        /// The member connection timeout in milliseconds.
        /// </summary>
        [Output("timeoutMemberConnect")]
        public Output<int> TimeoutMemberConnect { get; private set; } = null!;

        /// <summary>
        /// The member inactivity timeout in milliseconds.
        /// </summary>
        [Output("timeoutMemberData")]
        public Output<int> TimeoutMemberData { get; private set; } = null!;

        /// <summary>
        /// The time in milliseconds, to wait for additional
        /// TCP packets for content inspection.
        /// </summary>
        [Output("timeoutTcpInspect")]
        public Output<int> TimeoutTcpInspect { get; private set; } = null!;


        /// <summary>
        /// Create a Listener resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listener(string name, ListenerArgs args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/listener:Listener", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Listener(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/listener:Listener", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Listener resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Listener Get(string name, Input<string> id, ListenerState? state = null, CustomResourceOptions? options = null)
        {
            return new Listener(name, id, state, options);
        }
    }

    public sealed class ListenerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the Listener.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The maximum number of connections allowed
        /// for the Listener.
        /// </summary>
        [Input("connectionLimit")]
        public Input<int>? ConnectionLimit { get; set; }

        /// <summary>
        /// The ID of the default pool with which the
        /// Listener is associated.
        /// </summary>
        [Input("defaultPoolId")]
        public Input<string>? DefaultPoolId { get; set; }

        /// <summary>
        /// A reference to a Barbican Secrets
        /// container which stores TLS information. This is required if the protocol
        /// is `TERMINATED_HTTPS`. See
        /// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
        /// for more information.
        /// </summary>
        [Input("defaultTlsContainerRef")]
        public Input<string>? DefaultTlsContainerRef { get; set; }

        /// <summary>
        /// Human-readable description for the Listener.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The load balancer on which to provision this
        /// Listener. Changing this creates a new Listener.
        /// </summary>
        [Input("loadbalancerId", required: true)]
        public Input<string> LoadbalancerId { get; set; } = null!;

        /// <summary>
        /// Human-readable name for the Listener. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol - can either be TCP, HTTP, HTTPS,
        /// TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
        /// new Listener.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The port on which to listen for client traffic.
        /// Changing this creates a new Listener.
        /// </summary>
        [Input("protocolPort", required: true)]
        public Input<int> ProtocolPort { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// Listener.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sniContainerRefs")]
        private InputList<string>? _sniContainerRefs;

        /// <summary>
        /// A list of references to Barbican Secrets
        /// containers which store SNI information. See
        /// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
        /// for more information.
        /// </summary>
        public InputList<string> SniContainerRefs
        {
            get => _sniContainerRefs ?? (_sniContainerRefs = new InputList<string>());
            set => _sniContainerRefs = value;
        }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Listener.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new Listener.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The client inactivity timeout in milliseconds.
        /// </summary>
        [Input("timeoutClientData")]
        public Input<int>? TimeoutClientData { get; set; }

        /// <summary>
        /// The member connection timeout in milliseconds.
        /// </summary>
        [Input("timeoutMemberConnect")]
        public Input<int>? TimeoutMemberConnect { get; set; }

        /// <summary>
        /// The member inactivity timeout in milliseconds.
        /// </summary>
        [Input("timeoutMemberData")]
        public Input<int>? TimeoutMemberData { get; set; }

        /// <summary>
        /// The time in milliseconds, to wait for additional
        /// TCP packets for content inspection.
        /// </summary>
        [Input("timeoutTcpInspect")]
        public Input<int>? TimeoutTcpInspect { get; set; }

        public ListenerArgs()
        {
        }
    }

    public sealed class ListenerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the Listener.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The maximum number of connections allowed
        /// for the Listener.
        /// </summary>
        [Input("connectionLimit")]
        public Input<int>? ConnectionLimit { get; set; }

        /// <summary>
        /// The ID of the default pool with which the
        /// Listener is associated.
        /// </summary>
        [Input("defaultPoolId")]
        public Input<string>? DefaultPoolId { get; set; }

        /// <summary>
        /// A reference to a Barbican Secrets
        /// container which stores TLS information. This is required if the protocol
        /// is `TERMINATED_HTTPS`. See
        /// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
        /// for more information.
        /// </summary>
        [Input("defaultTlsContainerRef")]
        public Input<string>? DefaultTlsContainerRef { get; set; }

        /// <summary>
        /// Human-readable description for the Listener.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The load balancer on which to provision this
        /// Listener. Changing this creates a new Listener.
        /// </summary>
        [Input("loadbalancerId")]
        public Input<string>? LoadbalancerId { get; set; }

        /// <summary>
        /// Human-readable name for the Listener. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol - can either be TCP, HTTP, HTTPS,
        /// TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
        /// new Listener.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The port on which to listen for client traffic.
        /// Changing this creates a new Listener.
        /// </summary>
        [Input("protocolPort")]
        public Input<int>? ProtocolPort { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// Listener.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sniContainerRefs")]
        private InputList<string>? _sniContainerRefs;

        /// <summary>
        /// A list of references to Barbican Secrets
        /// containers which store SNI information. See
        /// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
        /// for more information.
        /// </summary>
        public InputList<string> SniContainerRefs
        {
            get => _sniContainerRefs ?? (_sniContainerRefs = new InputList<string>());
            set => _sniContainerRefs = value;
        }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Listener.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new Listener.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The client inactivity timeout in milliseconds.
        /// </summary>
        [Input("timeoutClientData")]
        public Input<int>? TimeoutClientData { get; set; }

        /// <summary>
        /// The member connection timeout in milliseconds.
        /// </summary>
        [Input("timeoutMemberConnect")]
        public Input<int>? TimeoutMemberConnect { get; set; }

        /// <summary>
        /// The member inactivity timeout in milliseconds.
        /// </summary>
        [Input("timeoutMemberData")]
        public Input<int>? TimeoutMemberData { get; set; }

        /// <summary>
        /// The time in milliseconds, to wait for additional
        /// TCP packets for content inspection.
        /// </summary>
        [Input("timeoutTcpInspect")]
        public Input<int>? TimeoutTcpInspect { get; set; }

        public ListenerState()
        {
        }
    }
}
