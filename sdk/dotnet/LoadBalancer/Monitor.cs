// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.LoadBalancer
{
    /// <summary>
    /// Manages a V2 monitor resource within OpenStack.
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
    ///     var monitor1 = new OpenStack.LoadBalancer.Monitor("monitor_1", new()
    ///     {
    ///         PoolId = pool1.Id,
    ///         Type = "PING",
    ///         Delay = 20,
    ///         Timeout = 10,
    ///         MaxRetries = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer Pool Monitor can be imported using the Monitor ID, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2
    /// ```
    /// In case of using OpenContrail, the import may not work properly. If you face an issue, try to import the monitor providing its parent pool ID:
    /// 
    /// ```sh
    /// $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2/708bc224-0f8c-4981-ac82-97095fe051b6
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:loadbalancer/monitor:Monitor")]
    public partial class Monitor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the monitor.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The time, in seconds, between sending probes to members.
        /// </summary>
        [Output("delay")]
        public Output<int> Delay { get; private set; } = null!;

        /// <summary>
        /// The domain name to use in the HTTP host header
        /// health monitor requests. Supported in Octavia API version 2.10 or later.
        /// </summary>
        [Output("domainName")]
        public Output<string?> DomainName { get; private set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. Expected HTTP codes
        /// for a passing HTTP(S) monitor. You can either specify a single status like
        /// "200", a list like "200, 202" or a range like "200-202". Default is "200".
        /// </summary>
        [Output("expectedCodes")]
        public Output<string> ExpectedCodes { get; private set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. The HTTP method that 
        /// the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
        /// OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET.
        /// </summary>
        [Output("httpMethod")]
        public Output<string> HttpMethod { get; private set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. The HTTP version that
        /// the health monitor uses for requests. One of `1.0` or 1.1` is supported
        /// for HTTP(S) monitors. The default is `1.0`. Supported in Octavia API version
        /// 2.10 or later.
        /// </summary>
        [Output("httpVersion")]
        public Output<string?> HttpVersion { get; private set; } = null!;

        /// <summary>
        /// Number of permissible ping failures before
        /// changing the member's status to INACTIVE. Must be a number between 1
        /// and 10.
        /// </summary>
        [Output("maxRetries")]
        public Output<int> MaxRetries { get; private set; } = null!;

        /// <summary>
        /// Number of permissible ping failures before 
        /// changing the member's status to ERROR. Must be a number between 1 and 10.
        /// The default is 3. Changing this updates the max_retries_down of the
        /// existing monitor.
        /// </summary>
        [Output("maxRetriesDown")]
        public Output<int> MaxRetriesDown { get; private set; } = null!;

        /// <summary>
        /// The Name of the Monitor.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the pool that this monitor will be assigned to.
        /// </summary>
        [Output("poolId")]
        public Output<string> PoolId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a monitor. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// monitor.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the monitor.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new monitor.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Maximum number of seconds for a monitor to wait for a
        /// ping reply before it times out. The value must be less than the delay
        /// value.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// The type of probe, which is PING, TCP, HTTP, HTTPS,
        /// TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
        /// verify the member state. Changing this creates a new monitor.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. URI path that will be
        /// accessed if monitor type is HTTP or HTTPS. Default is `/`.
        /// </summary>
        [Output("urlPath")]
        public Output<string> UrlPath { get; private set; } = null!;


        /// <summary>
        /// Create a Monitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Monitor(string name, MonitorArgs args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/monitor:Monitor", name, args ?? new MonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Monitor(string name, Input<string> id, MonitorState? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/monitor:Monitor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Monitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Monitor Get(string name, Input<string> id, MonitorState? state = null, CustomResourceOptions? options = null)
        {
            return new Monitor(name, id, state, options);
        }
    }

    public sealed class MonitorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the monitor.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The time, in seconds, between sending probes to members.
        /// </summary>
        [Input("delay", required: true)]
        public Input<int> Delay { get; set; } = null!;

        /// <summary>
        /// The domain name to use in the HTTP host header
        /// health monitor requests. Supported in Octavia API version 2.10 or later.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. Expected HTTP codes
        /// for a passing HTTP(S) monitor. You can either specify a single status like
        /// "200", a list like "200, 202" or a range like "200-202". Default is "200".
        /// </summary>
        [Input("expectedCodes")]
        public Input<string>? ExpectedCodes { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. The HTTP method that 
        /// the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
        /// OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. The HTTP version that
        /// the health monitor uses for requests. One of `1.0` or 1.1` is supported
        /// for HTTP(S) monitors. The default is `1.0`. Supported in Octavia API version
        /// 2.10 or later.
        /// </summary>
        [Input("httpVersion")]
        public Input<string>? HttpVersion { get; set; }

        /// <summary>
        /// Number of permissible ping failures before
        /// changing the member's status to INACTIVE. Must be a number between 1
        /// and 10.
        /// </summary>
        [Input("maxRetries", required: true)]
        public Input<int> MaxRetries { get; set; } = null!;

        /// <summary>
        /// Number of permissible ping failures before 
        /// changing the member's status to ERROR. Must be a number between 1 and 10.
        /// The default is 3. Changing this updates the max_retries_down of the
        /// existing monitor.
        /// </summary>
        [Input("maxRetriesDown")]
        public Input<int>? MaxRetriesDown { get; set; }

        /// <summary>
        /// The Name of the Monitor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the pool that this monitor will be assigned to.
        /// </summary>
        [Input("poolId", required: true)]
        public Input<string> PoolId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a monitor. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// monitor.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the monitor.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new monitor.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Maximum number of seconds for a monitor to wait for a
        /// ping reply before it times out. The value must be less than the delay
        /// value.
        /// </summary>
        [Input("timeout", required: true)]
        public Input<int> Timeout { get; set; } = null!;

        /// <summary>
        /// The type of probe, which is PING, TCP, HTTP, HTTPS,
        /// TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
        /// verify the member state. Changing this creates a new monitor.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. URI path that will be
        /// accessed if monitor type is HTTP or HTTPS. Default is `/`.
        /// </summary>
        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        public MonitorArgs()
        {
        }
        public static new MonitorArgs Empty => new MonitorArgs();
    }

    public sealed class MonitorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the monitor.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The time, in seconds, between sending probes to members.
        /// </summary>
        [Input("delay")]
        public Input<int>? Delay { get; set; }

        /// <summary>
        /// The domain name to use in the HTTP host header
        /// health monitor requests. Supported in Octavia API version 2.10 or later.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. Expected HTTP codes
        /// for a passing HTTP(S) monitor. You can either specify a single status like
        /// "200", a list like "200, 202" or a range like "200-202". Default is "200".
        /// </summary>
        [Input("expectedCodes")]
        public Input<string>? ExpectedCodes { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. The HTTP method that 
        /// the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
        /// OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. The HTTP version that
        /// the health monitor uses for requests. One of `1.0` or 1.1` is supported
        /// for HTTP(S) monitors. The default is `1.0`. Supported in Octavia API version
        /// 2.10 or later.
        /// </summary>
        [Input("httpVersion")]
        public Input<string>? HttpVersion { get; set; }

        /// <summary>
        /// Number of permissible ping failures before
        /// changing the member's status to INACTIVE. Must be a number between 1
        /// and 10.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// Number of permissible ping failures before 
        /// changing the member's status to ERROR. Must be a number between 1 and 10.
        /// The default is 3. Changing this updates the max_retries_down of the
        /// existing monitor.
        /// </summary>
        [Input("maxRetriesDown")]
        public Input<int>? MaxRetriesDown { get; set; }

        /// <summary>
        /// The Name of the Monitor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the pool that this monitor will be assigned to.
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a monitor. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// monitor.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the monitor.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new monitor.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Maximum number of seconds for a monitor to wait for a
        /// ping reply before it times out. The value must be less than the delay
        /// value.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The type of probe, which is PING, TCP, HTTP, HTTPS,
        /// TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
        /// verify the member state. Changing this creates a new monitor.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. URI path that will be
        /// accessed if monitor type is HTTP or HTTPS. Default is `/`.
        /// </summary>
        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        public MonitorState()
        {
        }
        public static new MonitorState Empty => new MonitorState();
    }
}
