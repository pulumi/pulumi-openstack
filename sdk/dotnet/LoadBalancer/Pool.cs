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
    /// Manages a V2 pool resource within OpenStack.
    /// 
    /// &gt; **Note:** This resource has attributes that depend on octavia minor versions.
    /// Please ensure your Openstack cloud supports the required minor version.
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
    ///     var pool1 = new OpenStack.LoadBalancer.Pool("pool_1", new()
    ///     {
    ///         Protocol = "HTTP",
    ///         LbMethod = "ROUND_ROBIN",
    ///         ListenerId = "d9415786-5f1a-428b-b35f-2f1523e146d2",
    ///         Persistence = new OpenStack.LoadBalancer.Inputs.PoolPersistenceArgs
    ///         {
    ///             Type = "APP_COOKIE",
    ///             CookieName = "testCookie",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer Pool can be imported using the Pool ID, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import openstack:loadbalancer/pool:Pool pool_1 60ad9ee4-249a-4d60-a45b-aa60e046c513
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:loadbalancer/pool:Pool")]
    public partial class Pool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the pool. A valid
        /// value is true (UP) or false (DOWN).
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// A list of ALPN protocols. Available protocols:
        /// `http/1.0`, `http/1.1`, `h2`. Supported only in **Octavia minor version &gt;=
        /// 2.24**.
        /// </summary>
        [Output("alpnProtocols")]
        public Output<ImmutableArray<string>> AlpnProtocols { get; private set; } = null!;

        /// <summary>
        /// The reference of the key manager service
        /// secret containing a PEM format CA certificate bundle for `tls_enabled` pools.
        /// Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Output("caTlsContainerRef")]
        public Output<string?> CaTlsContainerRef { get; private set; } = null!;

        /// <summary>
        /// The reference of the key manager service
        /// secret containing a PEM format CA revocation list file for `tls_enabled`
        /// pools. Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Output("crlContainerRef")]
        public Output<string?> CrlContainerRef { get; private set; } = null!;

        /// <summary>
        /// Human-readable description for the pool.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The load balancing algorithm to distribute traffic
        /// to the pool's members. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS,
        /// SOURCE_IP, or SOURCE_IP_PORT.
        /// </summary>
        [Output("lbMethod")]
        public Output<string> LbMethod { get; private set; } = null!;

        /// <summary>
        /// The Listener on which the members of the pool will
        /// be associated with. Changing this creates a new pool. Note: One of
        /// LoadbalancerID or ListenerID must be provided.
        /// </summary>
        [Output("listenerId")]
        public Output<string?> ListenerId { get; private set; } = null!;

        /// <summary>
        /// The load balancer on which to provision this
        /// pool. Changing this creates a new pool. Note: One of LoadbalancerID or
        /// ListenerID must be provided.
        /// </summary>
        [Output("loadbalancerId")]
        public Output<string?> LoadbalancerId { get; private set; } = null!;

        /// <summary>
        /// Human-readable name for the pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Omit this field to prevent session persistence.
        /// Indicates whether connections in the same session will be processed by the
        /// same Pool member or not. Changing this creates a new pool.
        /// </summary>
        [Output("persistence")]
        public Output<Outputs.PoolPersistence?> Persistence { get; private set; } = null!;

        /// <summary>
        /// The protocol - can either be TCP, HTTP, HTTPS, PROXY,
        /// UDP, PROXYV2 (**Octavia minor version &gt;= 2.22**) or SCTP (**Octavia minor
        /// version &gt;= 2.23**). Changing this creates a new pool.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a pool. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new pool.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the pool.  Only administrative users can specify a tenant UUID other than
        /// their own. Changing this creates a new pool.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// List of ciphers in OpenSSL format
        /// (colon-separated). See
        /// https://www.openssl.org/docs/man1.1.1/man1/ciphers.html for more information.
        /// Supported only in **Octavia minor version &gt;= 2.15**.
        /// </summary>
        [Output("tlsCiphers")]
        public Output<string> TlsCiphers { get; private set; } = null!;

        /// <summary>
        /// The reference to the key manager service
        /// secret containing a PKCS12 format certificate/key bundle for `tls_enabled`
        /// pools for TLS client authentication to the member servers. Supported only in
        /// **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Output("tlsContainerRef")]
        public Output<string?> TlsContainerRef { get; private set; } = null!;

        /// <summary>
        /// When true connections to backend member servers
        /// will use TLS encryption. Default is false. Supported only in **Octavia minor
        /// version &gt;= 2.8**.
        /// </summary>
        [Output("tlsEnabled")]
        public Output<bool?> TlsEnabled { get; private set; } = null!;

        /// <summary>
        /// A list of TLS protocol versions. Available
        /// versions: `TLSv1`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Supported only in
        /// **Octavia minor version &gt;= 2.17**.
        /// </summary>
        [Output("tlsVersions")]
        public Output<ImmutableArray<string>> TlsVersions { get; private set; } = null!;


        /// <summary>
        /// Create a Pool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pool(string name, PoolArgs args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/pool:Pool", name, args ?? new PoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pool(string name, Input<string> id, PoolState? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/pool:Pool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Pool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pool Get(string name, Input<string> id, PoolState? state = null, CustomResourceOptions? options = null)
        {
            return new Pool(name, id, state, options);
        }
    }

    public sealed class PoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the pool. A valid
        /// value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("alpnProtocols")]
        private InputList<string>? _alpnProtocols;

        /// <summary>
        /// A list of ALPN protocols. Available protocols:
        /// `http/1.0`, `http/1.1`, `h2`. Supported only in **Octavia minor version &gt;=
        /// 2.24**.
        /// </summary>
        public InputList<string> AlpnProtocols
        {
            get => _alpnProtocols ?? (_alpnProtocols = new InputList<string>());
            set => _alpnProtocols = value;
        }

        /// <summary>
        /// The reference of the key manager service
        /// secret containing a PEM format CA certificate bundle for `tls_enabled` pools.
        /// Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("caTlsContainerRef")]
        public Input<string>? CaTlsContainerRef { get; set; }

        /// <summary>
        /// The reference of the key manager service
        /// secret containing a PEM format CA revocation list file for `tls_enabled`
        /// pools. Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("crlContainerRef")]
        public Input<string>? CrlContainerRef { get; set; }

        /// <summary>
        /// Human-readable description for the pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The load balancing algorithm to distribute traffic
        /// to the pool's members. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS,
        /// SOURCE_IP, or SOURCE_IP_PORT.
        /// </summary>
        [Input("lbMethod", required: true)]
        public Input<string> LbMethod { get; set; } = null!;

        /// <summary>
        /// The Listener on which the members of the pool will
        /// be associated with. Changing this creates a new pool. Note: One of
        /// LoadbalancerID or ListenerID must be provided.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// The load balancer on which to provision this
        /// pool. Changing this creates a new pool. Note: One of LoadbalancerID or
        /// ListenerID must be provided.
        /// </summary>
        [Input("loadbalancerId")]
        public Input<string>? LoadbalancerId { get; set; }

        /// <summary>
        /// Human-readable name for the pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Omit this field to prevent session persistence.
        /// Indicates whether connections in the same session will be processed by the
        /// same Pool member or not. Changing this creates a new pool.
        /// </summary>
        [Input("persistence")]
        public Input<Inputs.PoolPersistenceArgs>? Persistence { get; set; }

        /// <summary>
        /// The protocol - can either be TCP, HTTP, HTTPS, PROXY,
        /// UDP, PROXYV2 (**Octavia minor version &gt;= 2.22**) or SCTP (**Octavia minor
        /// version &gt;= 2.23**). Changing this creates a new pool.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a pool. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new pool.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the pool.  Only administrative users can specify a tenant UUID other than
        /// their own. Changing this creates a new pool.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// List of ciphers in OpenSSL format
        /// (colon-separated). See
        /// https://www.openssl.org/docs/man1.1.1/man1/ciphers.html for more information.
        /// Supported only in **Octavia minor version &gt;= 2.15**.
        /// </summary>
        [Input("tlsCiphers")]
        public Input<string>? TlsCiphers { get; set; }

        /// <summary>
        /// The reference to the key manager service
        /// secret containing a PKCS12 format certificate/key bundle for `tls_enabled`
        /// pools for TLS client authentication to the member servers. Supported only in
        /// **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("tlsContainerRef")]
        public Input<string>? TlsContainerRef { get; set; }

        /// <summary>
        /// When true connections to backend member servers
        /// will use TLS encryption. Default is false. Supported only in **Octavia minor
        /// version &gt;= 2.8**.
        /// </summary>
        [Input("tlsEnabled")]
        public Input<bool>? TlsEnabled { get; set; }

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

        public PoolArgs()
        {
        }
        public static new PoolArgs Empty => new PoolArgs();
    }

    public sealed class PoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the pool. A valid
        /// value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("alpnProtocols")]
        private InputList<string>? _alpnProtocols;

        /// <summary>
        /// A list of ALPN protocols. Available protocols:
        /// `http/1.0`, `http/1.1`, `h2`. Supported only in **Octavia minor version &gt;=
        /// 2.24**.
        /// </summary>
        public InputList<string> AlpnProtocols
        {
            get => _alpnProtocols ?? (_alpnProtocols = new InputList<string>());
            set => _alpnProtocols = value;
        }

        /// <summary>
        /// The reference of the key manager service
        /// secret containing a PEM format CA certificate bundle for `tls_enabled` pools.
        /// Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("caTlsContainerRef")]
        public Input<string>? CaTlsContainerRef { get; set; }

        /// <summary>
        /// The reference of the key manager service
        /// secret containing a PEM format CA revocation list file for `tls_enabled`
        /// pools. Supported only in **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("crlContainerRef")]
        public Input<string>? CrlContainerRef { get; set; }

        /// <summary>
        /// Human-readable description for the pool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The load balancing algorithm to distribute traffic
        /// to the pool's members. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS,
        /// SOURCE_IP, or SOURCE_IP_PORT.
        /// </summary>
        [Input("lbMethod")]
        public Input<string>? LbMethod { get; set; }

        /// <summary>
        /// The Listener on which the members of the pool will
        /// be associated with. Changing this creates a new pool. Note: One of
        /// LoadbalancerID or ListenerID must be provided.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// The load balancer on which to provision this
        /// pool. Changing this creates a new pool. Note: One of LoadbalancerID or
        /// ListenerID must be provided.
        /// </summary>
        [Input("loadbalancerId")]
        public Input<string>? LoadbalancerId { get; set; }

        /// <summary>
        /// Human-readable name for the pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Omit this field to prevent session persistence.
        /// Indicates whether connections in the same session will be processed by the
        /// same Pool member or not. Changing this creates a new pool.
        /// </summary>
        [Input("persistence")]
        public Input<Inputs.PoolPersistenceGetArgs>? Persistence { get; set; }

        /// <summary>
        /// The protocol - can either be TCP, HTTP, HTTPS, PROXY,
        /// UDP, PROXYV2 (**Octavia minor version &gt;= 2.22**) or SCTP (**Octavia minor
        /// version &gt;= 2.23**). Changing this creates a new pool.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a pool. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new pool.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the pool.  Only administrative users can specify a tenant UUID other than
        /// their own. Changing this creates a new pool.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// List of ciphers in OpenSSL format
        /// (colon-separated). See
        /// https://www.openssl.org/docs/man1.1.1/man1/ciphers.html for more information.
        /// Supported only in **Octavia minor version &gt;= 2.15**.
        /// </summary>
        [Input("tlsCiphers")]
        public Input<string>? TlsCiphers { get; set; }

        /// <summary>
        /// The reference to the key manager service
        /// secret containing a PKCS12 format certificate/key bundle for `tls_enabled`
        /// pools for TLS client authentication to the member servers. Supported only in
        /// **Octavia minor version &gt;= 2.8**.
        /// </summary>
        [Input("tlsContainerRef")]
        public Input<string>? TlsContainerRef { get; set; }

        /// <summary>
        /// When true connections to backend member servers
        /// will use TLS encryption. Default is false. Supported only in **Octavia minor
        /// version &gt;= 2.8**.
        /// </summary>
        [Input("tlsEnabled")]
        public Input<bool>? TlsEnabled { get; set; }

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

        public PoolState()
        {
        }
        public static new PoolState Empty => new PoolState();
    }
}
