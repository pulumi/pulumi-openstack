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
    /// Manages a V2 Neutron QoS bandwidth limit rule resource within OpenStack.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a QoS Policy with some bandwidth limit rule
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var qosPolicy1 = new OpenStack.Networking.QosPolicy("qos_policy_1", new()
    ///     {
    ///         Name = "qos_policy_1",
    ///         Description = "bw_limit",
    ///     });
    /// 
    ///     var bwLimitRule1 = new OpenStack.Networking.QosBandwidthLimitRule("bw_limit_rule_1", new()
    ///     {
    ///         QosPolicyId = qosPolicy1.Id,
    ///         MaxKbps = 3000,
    ///         MaxBurstKbps = 300,
    ///         Direction = "egress",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// QoS bandwidth limit rules can be imported using the `qos_policy_id/bandwidth_limit_rule` format, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule bw_limit_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule")]
    public partial class QosBandwidthLimitRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        /// existing QoS bandwidth limit rule.
        /// </summary>
        [Output("direction")]
        public Output<string?> Direction { get; private set; } = null!;

        /// <summary>
        /// The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
        /// maximum burst size in kilobits of the existing QoS bandwidth limit rule.
        /// </summary>
        [Output("maxBurstKbps")]
        public Output<int?> MaxBurstKbps { get; private set; } = null!;

        /// <summary>
        /// The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
        /// maximum kilobits per second of the existing QoS bandwidth limit rule.
        /// </summary>
        [Output("maxKbps")]
        public Output<int> MaxKbps { get; private set; } = null!;

        /// <summary>
        /// The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
        /// </summary>
        [Output("qosPolicyId")]
        public Output<string> QosPolicyId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a QosBandwidthLimitRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QosBandwidthLimitRule(string name, QosBandwidthLimitRuleArgs args, CustomResourceOptions? options = null)
            : base("openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule", name, args ?? new QosBandwidthLimitRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QosBandwidthLimitRule(string name, Input<string> id, QosBandwidthLimitRuleState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QosBandwidthLimitRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QosBandwidthLimitRule Get(string name, Input<string> id, QosBandwidthLimitRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new QosBandwidthLimitRule(name, id, state, options);
        }
    }

    public sealed class QosBandwidthLimitRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        /// existing QoS bandwidth limit rule.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
        /// maximum burst size in kilobits of the existing QoS bandwidth limit rule.
        /// </summary>
        [Input("maxBurstKbps")]
        public Input<int>? MaxBurstKbps { get; set; }

        /// <summary>
        /// The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
        /// maximum kilobits per second of the existing QoS bandwidth limit rule.
        /// </summary>
        [Input("maxKbps", required: true)]
        public Input<int> MaxKbps { get; set; } = null!;

        /// <summary>
        /// The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
        /// </summary>
        [Input("qosPolicyId", required: true)]
        public Input<string> QosPolicyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public QosBandwidthLimitRuleArgs()
        {
        }
        public static new QosBandwidthLimitRuleArgs Empty => new QosBandwidthLimitRuleArgs();
    }

    public sealed class QosBandwidthLimitRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        /// existing QoS bandwidth limit rule.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
        /// maximum burst size in kilobits of the existing QoS bandwidth limit rule.
        /// </summary>
        [Input("maxBurstKbps")]
        public Input<int>? MaxBurstKbps { get; set; }

        /// <summary>
        /// The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
        /// maximum kilobits per second of the existing QoS bandwidth limit rule.
        /// </summary>
        [Input("maxKbps")]
        public Input<int>? MaxKbps { get; set; }

        /// <summary>
        /// The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
        /// </summary>
        [Input("qosPolicyId")]
        public Input<string>? QosPolicyId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public QosBandwidthLimitRuleState()
        {
        }
        public static new QosBandwidthLimitRuleState Empty => new QosBandwidthLimitRuleState();
    }
}
