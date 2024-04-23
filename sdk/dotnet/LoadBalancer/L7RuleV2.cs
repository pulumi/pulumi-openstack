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
    /// Manages a V2 L7 Rule resource within OpenStack.
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
    ///     var network1 = new OpenStack.Networking.Network("network_1", new()
    ///     {
    ///         Name = "network_1",
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var subnet1 = new OpenStack.Networking.Subnet("subnet_1", new()
    ///     {
    ///         Name = "subnet_1",
    ///         Cidr = "192.168.199.0/24",
    ///         IpVersion = 4,
    ///         NetworkId = network1.Id,
    ///     });
    /// 
    ///     var loadbalancer1 = new OpenStack.LoadBalancer.LoadBalancer("loadbalancer_1", new()
    ///     {
    ///         Name = "loadbalancer_1",
    ///         VipSubnetId = subnet1.Id,
    ///     });
    /// 
    ///     var listener1 = new OpenStack.LoadBalancer.Listener("listener_1", new()
    ///     {
    ///         Name = "listener_1",
    ///         Protocol = "HTTP",
    ///         ProtocolPort = 8080,
    ///         LoadbalancerId = loadbalancer1.Id,
    ///     });
    /// 
    ///     var pool1 = new OpenStack.LoadBalancer.Pool("pool_1", new()
    ///     {
    ///         Name = "pool_1",
    ///         Protocol = "HTTP",
    ///         LbMethod = "ROUND_ROBIN",
    ///         LoadbalancerId = loadbalancer1.Id,
    ///     });
    /// 
    ///     var l7policy1 = new OpenStack.LoadBalancer.L7PolicyV2("l7policy_1", new()
    ///     {
    ///         Name = "test",
    ///         Action = "REDIRECT_TO_URL",
    ///         Description = "test description",
    ///         Position = 1,
    ///         ListenerId = listener1.Id,
    ///         RedirectUrl = "http://www.example.com",
    ///     });
    /// 
    ///     var l7rule1 = new OpenStack.LoadBalancer.L7RuleV2("l7rule_1", new()
    ///     {
    ///         L7policyId = l7policy1.Id,
    ///         Type = "PATH",
    ///         CompareType = "EQUAL_TO",
    ///         Value = "/api",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID
    /// separated by a slash, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import openstack:loadbalancer/l7RuleV2:L7RuleV2 l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:loadbalancer/l7RuleV2:L7RuleV2")]
    public partial class L7RuleV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the L7 Rule.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The comparison type for the L7 rule - can either be
        /// CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        /// </summary>
        [Output("compareType")]
        public Output<string> CompareType { get; private set; } = null!;

        /// <summary>
        /// When true the logic of the rule is inverted. For example, with invert
        /// true, equal to would become not equal to. Default is false.
        /// </summary>
        [Output("invert")]
        public Output<bool?> Invert { get; private set; } = null!;

        /// <summary>
        /// The key to use for the comparison. For example, the name of the cookie to
        /// evaluate. Valid when `type` is set to COOKIE or HEADER.
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        /// <summary>
        /// The ID of the L7 Policy to query. Changing this creates a new
        /// L7 Rule.
        /// </summary>
        [Output("l7policyId")]
        public Output<string> L7policyId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Listener owning this resource.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// L7 Rule.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the L7 Rule.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new L7 Rule.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
        /// HOST\_NAME or PATH.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The value to use for the comparison. For example, the file type to
        /// compare.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a L7RuleV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public L7RuleV2(string name, L7RuleV2Args args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/l7RuleV2:L7RuleV2", name, args ?? new L7RuleV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private L7RuleV2(string name, Input<string> id, L7RuleV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/l7RuleV2:L7RuleV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing L7RuleV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static L7RuleV2 Get(string name, Input<string> id, L7RuleV2State? state = null, CustomResourceOptions? options = null)
        {
            return new L7RuleV2(name, id, state, options);
        }
    }

    public sealed class L7RuleV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the L7 Rule.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The comparison type for the L7 rule - can either be
        /// CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        /// </summary>
        [Input("compareType", required: true)]
        public Input<string> CompareType { get; set; } = null!;

        /// <summary>
        /// When true the logic of the rule is inverted. For example, with invert
        /// true, equal to would become not equal to. Default is false.
        /// </summary>
        [Input("invert")]
        public Input<bool>? Invert { get; set; }

        /// <summary>
        /// The key to use for the comparison. For example, the name of the cookie to
        /// evaluate. Valid when `type` is set to COOKIE or HEADER.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The ID of the L7 Policy to query. Changing this creates a new
        /// L7 Rule.
        /// </summary>
        [Input("l7policyId", required: true)]
        public Input<string> L7policyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// L7 Rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the L7 Rule.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new L7 Rule.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
        /// HOST\_NAME or PATH.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The value to use for the comparison. For example, the file type to
        /// compare.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public L7RuleV2Args()
        {
        }
        public static new L7RuleV2Args Empty => new L7RuleV2Args();
    }

    public sealed class L7RuleV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the L7 Rule.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The comparison type for the L7 rule - can either be
        /// CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        /// </summary>
        [Input("compareType")]
        public Input<string>? CompareType { get; set; }

        /// <summary>
        /// When true the logic of the rule is inverted. For example, with invert
        /// true, equal to would become not equal to. Default is false.
        /// </summary>
        [Input("invert")]
        public Input<bool>? Invert { get; set; }

        /// <summary>
        /// The key to use for the comparison. For example, the name of the cookie to
        /// evaluate. Valid when `type` is set to COOKIE or HEADER.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The ID of the L7 Policy to query. Changing this creates a new
        /// L7 Rule.
        /// </summary>
        [Input("l7policyId")]
        public Input<string>? L7policyId { get; set; }

        /// <summary>
        /// The ID of the Listener owning this resource.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// L7 Rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the L7 Rule.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new L7 Rule.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
        /// HOST\_NAME or PATH.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The value to use for the comparison. For example, the file type to
        /// compare.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public L7RuleV2State()
        {
        }
        public static new L7RuleV2State Empty => new L7RuleV2State();
    }
}
