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
    /// Manages a V1 load balancer monitor resource within OpenStack.
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
    ///     var monitor1 = new OpenStack.LoadBalancer.MonitorV1("monitor_1", new()
    ///     {
    ///         Type = "PING",
    ///         Delay = 30,
    ///         Timeout = 5,
    ///         MaxRetries = 3,
    ///         AdminStateUp = "true",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer Members can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:loadbalancer/monitorV1:MonitorV1 monitor_1 119d7530-72e9-449a-aa97-124a5ef1992c
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:loadbalancer/monitorV1:MonitorV1")]
    public partial class MonitorV1 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the monitor.
        /// Acceptable values are "true" and "false". Changing this value updates the
        /// state of the existing monitor.
        /// </summary>
        [Output("adminStateUp")]
        public Output<string> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The time, in seconds, between sending probes to members.
        /// Changing this creates a new monitor.
        /// </summary>
        [Output("delay")]
        public Output<int> Delay { get; private set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. Expected HTTP codes
        /// for a passing HTTP(S) monitor. You can either specify a single status like
        /// "200", or a range like "200-202". Changing this updates the expected_codes
        /// of the existing monitor.
        /// </summary>
        [Output("expectedCodes")]
        public Output<string?> ExpectedCodes { get; private set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. The HTTP method used
        /// for requests by the monitor. If this attribute is not specified, it defaults
        /// to "GET". Changing this updates the http_method of the existing monitor.
        /// </summary>
        [Output("httpMethod")]
        public Output<string?> HttpMethod { get; private set; } = null!;

        /// <summary>
        /// Number of permissible ping failures before changing
        /// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
        /// this updates the max_retries of the existing monitor.
        /// </summary>
        [Output("maxRetries")]
        public Output<int> MaxRetries { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB monitor. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB monitor.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The owner of the monitor. Required if admin wants to
        /// create a monitor for another tenant. Changing this creates a new monitor.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Maximum number of seconds for a monitor to wait for a
        /// ping reply before it times out. The value must be less than the delay value.
        /// Changing this updates the timeout of the existing monitor.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// The type of probe, which is PING, TCP, HTTP, or HTTPS,
        /// that is sent by the monitor to verify the member state. Changing this
        /// creates a new monitor.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. URI path that will be
        /// accessed if monitor type is HTTP or HTTPS. Changing this updates the
        /// url_path of the existing monitor.
        /// </summary>
        [Output("urlPath")]
        public Output<string?> UrlPath { get; private set; } = null!;


        /// <summary>
        /// Create a MonitorV1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MonitorV1(string name, MonitorV1Args args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/monitorV1:MonitorV1", name, args ?? new MonitorV1Args(), MakeResourceOptions(options, ""))
        {
        }

        private MonitorV1(string name, Input<string> id, MonitorV1State? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/monitorV1:MonitorV1", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MonitorV1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MonitorV1 Get(string name, Input<string> id, MonitorV1State? state = null, CustomResourceOptions? options = null)
        {
            return new MonitorV1(name, id, state, options);
        }
    }

    public sealed class MonitorV1Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the monitor.
        /// Acceptable values are "true" and "false". Changing this value updates the
        /// state of the existing monitor.
        /// </summary>
        [Input("adminStateUp")]
        public Input<string>? AdminStateUp { get; set; }

        /// <summary>
        /// The time, in seconds, between sending probes to members.
        /// Changing this creates a new monitor.
        /// </summary>
        [Input("delay", required: true)]
        public Input<int> Delay { get; set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. Expected HTTP codes
        /// for a passing HTTP(S) monitor. You can either specify a single status like
        /// "200", or a range like "200-202". Changing this updates the expected_codes
        /// of the existing monitor.
        /// </summary>
        [Input("expectedCodes")]
        public Input<string>? ExpectedCodes { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. The HTTP method used
        /// for requests by the monitor. If this attribute is not specified, it defaults
        /// to "GET". Changing this updates the http_method of the existing monitor.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// Number of permissible ping failures before changing
        /// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
        /// this updates the max_retries of the existing monitor.
        /// </summary>
        [Input("maxRetries", required: true)]
        public Input<int> MaxRetries { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB monitor. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB monitor.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the monitor. Required if admin wants to
        /// create a monitor for another tenant. Changing this creates a new monitor.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Maximum number of seconds for a monitor to wait for a
        /// ping reply before it times out. The value must be less than the delay value.
        /// Changing this updates the timeout of the existing monitor.
        /// </summary>
        [Input("timeout", required: true)]
        public Input<int> Timeout { get; set; } = null!;

        /// <summary>
        /// The type of probe, which is PING, TCP, HTTP, or HTTPS,
        /// that is sent by the monitor to verify the member state. Changing this
        /// creates a new monitor.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Required for HTTP(S) types. URI path that will be
        /// accessed if monitor type is HTTP or HTTPS. Changing this updates the
        /// url_path of the existing monitor.
        /// </summary>
        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        public MonitorV1Args()
        {
        }
        public static new MonitorV1Args Empty => new MonitorV1Args();
    }

    public sealed class MonitorV1State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the monitor.
        /// Acceptable values are "true" and "false". Changing this value updates the
        /// state of the existing monitor.
        /// </summary>
        [Input("adminStateUp")]
        public Input<string>? AdminStateUp { get; set; }

        /// <summary>
        /// The time, in seconds, between sending probes to members.
        /// Changing this creates a new monitor.
        /// </summary>
        [Input("delay")]
        public Input<int>? Delay { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. Expected HTTP codes
        /// for a passing HTTP(S) monitor. You can either specify a single status like
        /// "200", or a range like "200-202". Changing this updates the expected_codes
        /// of the existing monitor.
        /// </summary>
        [Input("expectedCodes")]
        public Input<string>? ExpectedCodes { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. The HTTP method used
        /// for requests by the monitor. If this attribute is not specified, it defaults
        /// to "GET". Changing this updates the http_method of the existing monitor.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// Number of permissible ping failures before changing
        /// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
        /// this updates the max_retries of the existing monitor.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB monitor. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB monitor.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the monitor. Required if admin wants to
        /// create a monitor for another tenant. Changing this creates a new monitor.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Maximum number of seconds for a monitor to wait for a
        /// ping reply before it times out. The value must be less than the delay value.
        /// Changing this updates the timeout of the existing monitor.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The type of probe, which is PING, TCP, HTTP, or HTTPS,
        /// that is sent by the monitor to verify the member state. Changing this
        /// creates a new monitor.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Required for HTTP(S) types. URI path that will be
        /// accessed if monitor type is HTTP or HTTPS. Changing this updates the
        /// url_path of the existing monitor.
        /// </summary>
        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        public MonitorV1State()
        {
        }
        public static new MonitorV1State Empty => new MonitorV1State();
    }
}
