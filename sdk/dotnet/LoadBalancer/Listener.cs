// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.LoadBalancer
{
    /// <summary>
    /// Manages a V2 listener resource within OpenStack.
    /// 
    /// &gt; **Note:** This resource has attributes that depend on octavia minor versions.
    /// Please ensure your Openstack cloud supports the required minor version.
    /// 
    /// ## Example Usage
    /// 
    /// ### Simple listener
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var listener1 = new OpenStack.LoadBalancer.Listener("listener_1", new()
    ///     {
    ///         Protocol = "HTTP",
    ///         ProtocolPort = 8080,
    ///         LoadbalancerId = "d9415786-5f1a-428b-b35f-2f1523e146d2",
    ///         InsertHeaders = 
    ///         {
    ///             { "X-Forwarded-For", "true" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Listener with TLS and client certificate authentication
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// using Std = Pulumi.Std;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var certificate1 = new OpenStack.KeyManager.SecretV1("certificate_1", new()
    ///     {
    ///         Name = "certificate",
    ///         Payload = Std.Filebase64.Invoke(new()
    ///         {
    ///             Input = "snakeoil.p12",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         PayloadContentEncoding = "base64",
    ///         PayloadContentType = "application/octet-stream",
    ///     });
    /// 
    ///     var caCertificate1 = new OpenStack.KeyManager.SecretV1("ca_certificate_1", new()
    ///     {
    ///         Name = "certificate",
    ///         Payload = Std.File.Invoke(new()
    ///         {
    ///             Input = "CA.pem",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         SecretType = "certificate",
    ///         PayloadContentType = "text/plain",
    ///     });
    /// 
    ///     var subnet1 = OpenStack.Networking.GetSubnet.Invoke(new()
    ///     {
    ///         Name = "my-subnet",
    ///     });
    /// 
    ///     var lb1 = new OpenStack.LbLoadbalancerV2("lb_1", new()
    ///     {
    ///         Name = "loadbalancer",
    ///         VipSubnetId = subnet1.Apply(getSubnetResult =&gt; getSubnetResult.Id),
    ///     });
    /// 
    ///     var listener1 = new OpenStack.LoadBalancer.Listener("listener_1", new()
    ///     {
    ///         Name = "https",
    ///         Protocol = "TERMINATED_HTTPS",
    ///         ProtocolPort = 443,
    ///         LoadbalancerId = lb1.Id,
    ///         DefaultTlsContainerRef = certificate1,
    ///         ClientAuthentication = "OPTIONAL",
    ///         ClientCaTlsContainerRef = caCertificate2.SecretRef,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer Listener can be imported using the Listener ID, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import openstack:loadbalancer/listener:Listener listener_1 b67ce64e-8b26-405d-afeb-4a078901f15a
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:loadbalancer/listener:Listener")]
    public partial class Listener : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the Listener. A
        /// valid value is true (UP) or false (DOWN).
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// A list of CIDR blocks that are permitted to
        /// connect to this listener, denying all other source addresses. If not present,
        /// defaults to allow all.
        /// </summary>
        [Output("allowedCidrs")]
        public Output<ImmutableArray<string>> AllowedCidrs { get; private set; } = null!;

        /// <summary>
        /// A list of ALPN protocols. Available protocols:
        /// `http/1.0`, `http/1.1`, `h2`. Supported only in **Octavia minor version &gt;=
        /// 2.20**.
        /// </summary>
        [Output("alpnProtocols")]
        public Output<ImmutableArray<string>> AlpnProtocols { get; private set; } = null!;

        /// <summary>
        /// The TLS client authentication mode.
        /// Available options: `NONE`, `OPTIONAL` or `MANDATORY`. Requires
        /// `TERMINATED_HTTPS` listener protocol and the `client_ca_tls_container_ref`.
        /// Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Output("clientAuthentication")]
        public Output<string?> ClientAuthentication { get; private set; } = null!;

        /// <summary>
        /// The ref of the key manager service
        /// secret containing a PEM format client CA certificate bundle for
        /// `TERMINATED_HTTPS` listeners. Required if `client_authentication` is
        /// `OPTIONAL` or `MANDATORY`. Supported only in **Octavia minor version &gt;=
        /// 2.8**.
        /// </summary>
        [Output("clientCaTlsContainerRef")]
        public Output<string?> ClientCaTlsContainerRef { get; private set; } = null!;

        /// <summary>
        /// The URI of the key manager service
        /// secret containing a PEM format CA revocation list file for `TERMINATED_HTTPS`
        /// listeners. Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Output("clientCrlContainerRef")]
        public Output<string?> ClientCrlContainerRef { get; private set; } = null!;

        /// <summary>
        /// The maximum number of connections allowed for
        /// the Listener.
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
        /// container which stores TLS information. This is required if the protocol is
        /// `TERMINATED_HTTPS`. See
        /// [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
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
        /// Defines whether the
        /// **includeSubDomains** directive should be added to the
        /// Strict-Transport-Security HTTP response header. This requires setting the
        /// `hsts_max_age` option as well in order to become effective. Requires
        /// `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
        /// version &gt;= 2.27**.
        /// </summary>
        [Output("hstsIncludeSubdomains")]
        public Output<bool?> HstsIncludeSubdomains { get; private set; } = null!;

        /// <summary>
        /// The value of the **max_age** directive for the
        /// Strict-Transport-Security HTTP response header. Setting this enables HTTP
        /// Strict Transport Security (HSTS) for the TLS-terminated listener. Requires
        /// `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
        /// version &gt;= 2.27**.
        /// </summary>
        [Output("hstsMaxAge")]
        public Output<int?> HstsMaxAge { get; private set; } = null!;

        /// <summary>
        /// Defines whether the **preload** directive should
        /// be added to the Strict-Transport-Security HTTP response header. This requires
        /// setting the `hsts_max_age` option as well in order to become effective.
        /// Requires `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia
        /// minor version &gt;= 2.27**.
        /// </summary>
        [Output("hstsPreload")]
        public Output<bool?> HstsPreload { get; private set; } = null!;

        /// <summary>
        /// The list of key value pairs representing
        /// headers to insert into the request before it is sent to the backend members.
        /// Changing this updates the headers of the existing listener.
        /// </summary>
        [Output("insertHeaders")]
        public Output<ImmutableDictionary<string, string>?> InsertHeaders { get; private set; } = null!;

        /// <summary>
        /// The load balancer on which to provision this
        /// Listener. Changing this creates a new Listener.
        /// </summary>
        [Output("loadbalancerId")]
        public Output<string> LoadbalancerId { get; private set; } = null!;

        /// <summary>
        /// Human-readable name for the Listener. Does not have to be
        /// unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The protocol can be either `TCP`, `HTTP`, `HTTPS`,
        /// `TERMINATED_HTTPS`, `UDP`, `SCTP` (supported only in **Octavia minor version
        /// \&gt;= 2.23**), or `PROMETHEUS` (supported only in **Octavia minor version &gt;=
        /// 2.25**). Changing this creates a new Listener.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The port on which to listen for client traffic.
        /// * Changing this creates a new Listener.
        /// </summary>
        [Output("protocolPort")]
        public Output<int> ProtocolPort { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a listener. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new Listener.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A list of references to Barbican Secrets
        /// containers which store SNI information. See
        /// [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
        /// for more information.
        /// </summary>
        [Output("sniContainerRefs")]
        public Output<ImmutableArray<string>> SniContainerRefs { get; private set; } = null!;

        /// <summary>
        /// A list of simple strings assigned to the pool. Available
        /// for Octavia **minor version 2.5 or later**.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Listener.  Only administrative users can specify a tenant UUID other than
        /// their own. Changing this creates a new Listener.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The client inactivity timeout in
        /// milliseconds.
        /// </summary>
        [Output("timeoutClientData")]
        public Output<int> TimeoutClientData { get; private set; } = null!;

        /// <summary>
        /// The member connection timeout in
        /// milliseconds.
        /// </summary>
        [Output("timeoutMemberConnect")]
        public Output<int> TimeoutMemberConnect { get; private set; } = null!;

        /// <summary>
        /// The member inactivity timeout in
        /// milliseconds.
        /// </summary>
        [Output("timeoutMemberData")]
        public Output<int> TimeoutMemberData { get; private set; } = null!;

        /// <summary>
        /// The time in milliseconds, to wait for
        /// additional TCP packets for content inspection.
        /// </summary>
        [Output("timeoutTcpInspect")]
        public Output<int> TimeoutTcpInspect { get; private set; } = null!;

        /// <summary>
        /// List of ciphers in OpenSSL format
        /// (colon-separated). See
        /// https://www.openssl.org/docs/man1.1.1/man1/ciphers.html for more information.
        /// Supported only in **Octavia minor version &gt;= 2.15**.
        /// </summary>
        [Output("tlsCiphers")]
        public Output<string> TlsCiphers { get; private set; } = null!;

        /// <summary>
        /// A list of TLS protocol versions. Available
        /// versions: `TLSv1`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Supported only in
        /// **Octavia minor version &gt;= 2.17**.
        /// </summary>
        [Output("tlsVersions")]
        public Output<ImmutableArray<string>> TlsVersions { get; private set; } = null!;


        /// <summary>
        /// Create a Listener resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listener(string name, ListenerArgs args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/listener:Listener", name, args ?? new ListenerArgs(), MakeResourceOptions(options, ""))
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

    public sealed class ListenerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the Listener. A
        /// valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("allowedCidrs")]
        private InputList<string>? _allowedCidrs;

        /// <summary>
        /// A list of CIDR blocks that are permitted to
        /// connect to this listener, denying all other source addresses. If not present,
        /// defaults to allow all.
        /// </summary>
        public InputList<string> AllowedCidrs
        {
            get => _allowedCidrs ?? (_allowedCidrs = new InputList<string>());
            set => _allowedCidrs = value;
        }

        [Input("alpnProtocols")]
        private InputList<string>? _alpnProtocols;

        /// <summary>
        /// A list of ALPN protocols. Available protocols:
        /// `http/1.0`, `http/1.1`, `h2`. Supported only in **Octavia minor version &gt;=
        /// 2.20**.
        /// </summary>
        public InputList<string> AlpnProtocols
        {
            get => _alpnProtocols ?? (_alpnProtocols = new InputList<string>());
            set => _alpnProtocols = value;
        }

        /// <summary>
        /// The TLS client authentication mode.
        /// Available options: `NONE`, `OPTIONAL` or `MANDATORY`. Requires
        /// `TERMINATED_HTTPS` listener protocol and the `client_ca_tls_container_ref`.
        /// Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("clientAuthentication")]
        public Input<string>? ClientAuthentication { get; set; }

        /// <summary>
        /// The ref of the key manager service
        /// secret containing a PEM format client CA certificate bundle for
        /// `TERMINATED_HTTPS` listeners. Required if `client_authentication` is
        /// `OPTIONAL` or `MANDATORY`. Supported only in **Octavia minor version &gt;=
        /// 2.8**.
        /// </summary>
        [Input("clientCaTlsContainerRef")]
        public Input<string>? ClientCaTlsContainerRef { get; set; }

        /// <summary>
        /// The URI of the key manager service
        /// secret containing a PEM format CA revocation list file for `TERMINATED_HTTPS`
        /// listeners. Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("clientCrlContainerRef")]
        public Input<string>? ClientCrlContainerRef { get; set; }

        /// <summary>
        /// The maximum number of connections allowed for
        /// the Listener.
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
        /// container which stores TLS information. This is required if the protocol is
        /// `TERMINATED_HTTPS`. See
        /// [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
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
        /// Defines whether the
        /// **includeSubDomains** directive should be added to the
        /// Strict-Transport-Security HTTP response header. This requires setting the
        /// `hsts_max_age` option as well in order to become effective. Requires
        /// `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
        /// version &gt;= 2.27**.
        /// </summary>
        [Input("hstsIncludeSubdomains")]
        public Input<bool>? HstsIncludeSubdomains { get; set; }

        /// <summary>
        /// The value of the **max_age** directive for the
        /// Strict-Transport-Security HTTP response header. Setting this enables HTTP
        /// Strict Transport Security (HSTS) for the TLS-terminated listener. Requires
        /// `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
        /// version &gt;= 2.27**.
        /// </summary>
        [Input("hstsMaxAge")]
        public Input<int>? HstsMaxAge { get; set; }

        /// <summary>
        /// Defines whether the **preload** directive should
        /// be added to the Strict-Transport-Security HTTP response header. This requires
        /// setting the `hsts_max_age` option as well in order to become effective.
        /// Requires `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia
        /// minor version &gt;= 2.27**.
        /// </summary>
        [Input("hstsPreload")]
        public Input<bool>? HstsPreload { get; set; }

        [Input("insertHeaders")]
        private InputMap<string>? _insertHeaders;

        /// <summary>
        /// The list of key value pairs representing
        /// headers to insert into the request before it is sent to the backend members.
        /// Changing this updates the headers of the existing listener.
        /// </summary>
        public InputMap<string> InsertHeaders
        {
            get => _insertHeaders ?? (_insertHeaders = new InputMap<string>());
            set => _insertHeaders = value;
        }

        /// <summary>
        /// The load balancer on which to provision this
        /// Listener. Changing this creates a new Listener.
        /// </summary>
        [Input("loadbalancerId", required: true)]
        public Input<string> LoadbalancerId { get; set; } = null!;

        /// <summary>
        /// Human-readable name for the Listener. Does not have to be
        /// unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol can be either `TCP`, `HTTP`, `HTTPS`,
        /// `TERMINATED_HTTPS`, `UDP`, `SCTP` (supported only in **Octavia minor version
        /// \&gt;= 2.23**), or `PROMETHEUS` (supported only in **Octavia minor version &gt;=
        /// 2.25**). Changing this creates a new Listener.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The port on which to listen for client traffic.
        /// * Changing this creates a new Listener.
        /// </summary>
        [Input("protocolPort", required: true)]
        public Input<int> ProtocolPort { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a listener. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new Listener.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sniContainerRefs")]
        private InputList<string>? _sniContainerRefs;

        /// <summary>
        /// A list of references to Barbican Secrets
        /// containers which store SNI information. See
        /// [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
        /// for more information.
        /// </summary>
        public InputList<string> SniContainerRefs
        {
            get => _sniContainerRefs ?? (_sniContainerRefs = new InputList<string>());
            set => _sniContainerRefs = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of simple strings assigned to the pool. Available
        /// for Octavia **minor version 2.5 or later**.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Listener.  Only administrative users can specify a tenant UUID other than
        /// their own. Changing this creates a new Listener.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The client inactivity timeout in
        /// milliseconds.
        /// </summary>
        [Input("timeoutClientData")]
        public Input<int>? TimeoutClientData { get; set; }

        /// <summary>
        /// The member connection timeout in
        /// milliseconds.
        /// </summary>
        [Input("timeoutMemberConnect")]
        public Input<int>? TimeoutMemberConnect { get; set; }

        /// <summary>
        /// The member inactivity timeout in
        /// milliseconds.
        /// </summary>
        [Input("timeoutMemberData")]
        public Input<int>? TimeoutMemberData { get; set; }

        /// <summary>
        /// The time in milliseconds, to wait for
        /// additional TCP packets for content inspection.
        /// </summary>
        [Input("timeoutTcpInspect")]
        public Input<int>? TimeoutTcpInspect { get; set; }

        /// <summary>
        /// List of ciphers in OpenSSL format
        /// (colon-separated). See
        /// https://www.openssl.org/docs/man1.1.1/man1/ciphers.html for more information.
        /// Supported only in **Octavia minor version &gt;= 2.15**.
        /// </summary>
        [Input("tlsCiphers")]
        public Input<string>? TlsCiphers { get; set; }

        [Input("tlsVersions")]
        private InputList<string>? _tlsVersions;

        /// <summary>
        /// A list of TLS protocol versions. Available
        /// versions: `TLSv1`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Supported only in
        /// **Octavia minor version &gt;= 2.17**.
        /// </summary>
        public InputList<string> TlsVersions
        {
            get => _tlsVersions ?? (_tlsVersions = new InputList<string>());
            set => _tlsVersions = value;
        }

        public ListenerArgs()
        {
        }
        public static new ListenerArgs Empty => new ListenerArgs();
    }

    public sealed class ListenerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the Listener. A
        /// valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("allowedCidrs")]
        private InputList<string>? _allowedCidrs;

        /// <summary>
        /// A list of CIDR blocks that are permitted to
        /// connect to this listener, denying all other source addresses. If not present,
        /// defaults to allow all.
        /// </summary>
        public InputList<string> AllowedCidrs
        {
            get => _allowedCidrs ?? (_allowedCidrs = new InputList<string>());
            set => _allowedCidrs = value;
        }

        [Input("alpnProtocols")]
        private InputList<string>? _alpnProtocols;

        /// <summary>
        /// A list of ALPN protocols. Available protocols:
        /// `http/1.0`, `http/1.1`, `h2`. Supported only in **Octavia minor version &gt;=
        /// 2.20**.
        /// </summary>
        public InputList<string> AlpnProtocols
        {
            get => _alpnProtocols ?? (_alpnProtocols = new InputList<string>());
            set => _alpnProtocols = value;
        }

        /// <summary>
        /// The TLS client authentication mode.
        /// Available options: `NONE`, `OPTIONAL` or `MANDATORY`. Requires
        /// `TERMINATED_HTTPS` listener protocol and the `client_ca_tls_container_ref`.
        /// Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("clientAuthentication")]
        public Input<string>? ClientAuthentication { get; set; }

        /// <summary>
        /// The ref of the key manager service
        /// secret containing a PEM format client CA certificate bundle for
        /// `TERMINATED_HTTPS` listeners. Required if `client_authentication` is
        /// `OPTIONAL` or `MANDATORY`. Supported only in **Octavia minor version &gt;=
        /// 2.8**.
        /// </summary>
        [Input("clientCaTlsContainerRef")]
        public Input<string>? ClientCaTlsContainerRef { get; set; }

        /// <summary>
        /// The URI of the key manager service
        /// secret containing a PEM format CA revocation list file for `TERMINATED_HTTPS`
        /// listeners. Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("clientCrlContainerRef")]
        public Input<string>? ClientCrlContainerRef { get; set; }

        /// <summary>
        /// The maximum number of connections allowed for
        /// the Listener.
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
        /// container which stores TLS information. This is required if the protocol is
        /// `TERMINATED_HTTPS`. See
        /// [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
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
        /// Defines whether the
        /// **includeSubDomains** directive should be added to the
        /// Strict-Transport-Security HTTP response header. This requires setting the
        /// `hsts_max_age` option as well in order to become effective. Requires
        /// `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
        /// version &gt;= 2.27**.
        /// </summary>
        [Input("hstsIncludeSubdomains")]
        public Input<bool>? HstsIncludeSubdomains { get; set; }

        /// <summary>
        /// The value of the **max_age** directive for the
        /// Strict-Transport-Security HTTP response header. Setting this enables HTTP
        /// Strict Transport Security (HSTS) for the TLS-terminated listener. Requires
        /// `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
        /// version &gt;= 2.27**.
        /// </summary>
        [Input("hstsMaxAge")]
        public Input<int>? HstsMaxAge { get; set; }

        /// <summary>
        /// Defines whether the **preload** directive should
        /// be added to the Strict-Transport-Security HTTP response header. This requires
        /// setting the `hsts_max_age` option as well in order to become effective.
        /// Requires `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia
        /// minor version &gt;= 2.27**.
        /// </summary>
        [Input("hstsPreload")]
        public Input<bool>? HstsPreload { get; set; }

        [Input("insertHeaders")]
        private InputMap<string>? _insertHeaders;

        /// <summary>
        /// The list of key value pairs representing
        /// headers to insert into the request before it is sent to the backend members.
        /// Changing this updates the headers of the existing listener.
        /// </summary>
        public InputMap<string> InsertHeaders
        {
            get => _insertHeaders ?? (_insertHeaders = new InputMap<string>());
            set => _insertHeaders = value;
        }

        /// <summary>
        /// The load balancer on which to provision this
        /// Listener. Changing this creates a new Listener.
        /// </summary>
        [Input("loadbalancerId")]
        public Input<string>? LoadbalancerId { get; set; }

        /// <summary>
        /// Human-readable name for the Listener. Does not have to be
        /// unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol can be either `TCP`, `HTTP`, `HTTPS`,
        /// `TERMINATED_HTTPS`, `UDP`, `SCTP` (supported only in **Octavia minor version
        /// \&gt;= 2.23**), or `PROMETHEUS` (supported only in **Octavia minor version &gt;=
        /// 2.25**). Changing this creates a new Listener.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The port on which to listen for client traffic.
        /// * Changing this creates a new Listener.
        /// </summary>
        [Input("protocolPort")]
        public Input<int>? ProtocolPort { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a listener. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new Listener.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("sniContainerRefs")]
        private InputList<string>? _sniContainerRefs;

        /// <summary>
        /// A list of references to Barbican Secrets
        /// containers which store SNI information. See
        /// [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
        /// for more information.
        /// </summary>
        public InputList<string> SniContainerRefs
        {
            get => _sniContainerRefs ?? (_sniContainerRefs = new InputList<string>());
            set => _sniContainerRefs = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of simple strings assigned to the pool. Available
        /// for Octavia **minor version 2.5 or later**.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Listener.  Only administrative users can specify a tenant UUID other than
        /// their own. Changing this creates a new Listener.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The client inactivity timeout in
        /// milliseconds.
        /// </summary>
        [Input("timeoutClientData")]
        public Input<int>? TimeoutClientData { get; set; }

        /// <summary>
        /// The member connection timeout in
        /// milliseconds.
        /// </summary>
        [Input("timeoutMemberConnect")]
        public Input<int>? TimeoutMemberConnect { get; set; }

        /// <summary>
        /// The member inactivity timeout in
        /// milliseconds.
        /// </summary>
        [Input("timeoutMemberData")]
        public Input<int>? TimeoutMemberData { get; set; }

        /// <summary>
        /// The time in milliseconds, to wait for
        /// additional TCP packets for content inspection.
        /// </summary>
        [Input("timeoutTcpInspect")]
        public Input<int>? TimeoutTcpInspect { get; set; }

        /// <summary>
        /// List of ciphers in OpenSSL format
        /// (colon-separated). See
        /// https://www.openssl.org/docs/man1.1.1/man1/ciphers.html for more information.
        /// Supported only in **Octavia minor version &gt;= 2.15**.
        /// </summary>
        [Input("tlsCiphers")]
        public Input<string>? TlsCiphers { get; set; }

        [Input("tlsVersions")]
        private InputList<string>? _tlsVersions;

        /// <summary>
        /// A list of TLS protocol versions. Available
        /// versions: `TLSv1`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Supported only in
        /// **Octavia minor version &gt;= 2.17**.
        /// </summary>
        public InputList<string> TlsVersions
        {
            get => _tlsVersions ?? (_tlsVersions = new InputList<string>());
            set => _tlsVersions = value;
        }

        public ListenerState()
        {
        }
        public static new ListenerState Empty => new ListenerState();
    }
}
