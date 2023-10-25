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
    /// Manages a V2 Neutron QoS minimum bandwidth rule resource within OpenStack.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// QoS minimum bandwidth rules can be imported using the `qos_policy_id/minimum_bandwidth_rule_id` format, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule minimum_bandwidth_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule")]
    public partial class QosMinimumBandwidthRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        /// existing QoS minimum bandwidth rule.
        /// </summary>
        [Output("direction")]
        public Output<string?> Direction { get; private set; } = null!;

        /// <summary>
        /// The minimum kilobits per second. Changing this updates the min kbps value of the existing
        /// QoS minimum bandwidth rule.
        /// </summary>
        [Output("minKbps")]
        public Output<int> MinKbps { get; private set; } = null!;

        /// <summary>
        /// The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        /// </summary>
        [Output("qosPolicyId")]
        public Output<string> QosPolicyId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a QosMinimumBandwidthRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QosMinimumBandwidthRule(string name, QosMinimumBandwidthRuleArgs args, CustomResourceOptions? options = null)
            : base("openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule", name, args ?? new QosMinimumBandwidthRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QosMinimumBandwidthRule(string name, Input<string> id, QosMinimumBandwidthRuleState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QosMinimumBandwidthRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QosMinimumBandwidthRule Get(string name, Input<string> id, QosMinimumBandwidthRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new QosMinimumBandwidthRule(name, id, state, options);
        }
    }

    public sealed class QosMinimumBandwidthRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        /// existing QoS minimum bandwidth rule.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The minimum kilobits per second. Changing this updates the min kbps value of the existing
        /// QoS minimum bandwidth rule.
        /// </summary>
        [Input("minKbps", required: true)]
        public Input<int> MinKbps { get; set; } = null!;

        /// <summary>
        /// The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        /// </summary>
        [Input("qosPolicyId", required: true)]
        public Input<string> QosPolicyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public QosMinimumBandwidthRuleArgs()
        {
        }
        public static new QosMinimumBandwidthRuleArgs Empty => new QosMinimumBandwidthRuleArgs();
    }

    public sealed class QosMinimumBandwidthRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        /// existing QoS minimum bandwidth rule.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The minimum kilobits per second. Changing this updates the min kbps value of the existing
        /// QoS minimum bandwidth rule.
        /// </summary>
        [Input("minKbps")]
        public Input<int>? MinKbps { get; set; }

        /// <summary>
        /// The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        /// </summary>
        [Input("qosPolicyId")]
        public Input<string>? QosPolicyId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public QosMinimumBandwidthRuleState()
        {
        }
        public static new QosMinimumBandwidthRuleState Empty => new QosMinimumBandwidthRuleState();
    }
}
