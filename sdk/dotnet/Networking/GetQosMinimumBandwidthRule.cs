// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Networking
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_minimum_bandwidth_rule_v2.html.markdown.
        /// </summary>
        public static Task<GetQosMinimumBandwidthRuleResult> GetQosMinimumBandwidthRule(GetQosMinimumBandwidthRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQosMinimumBandwidthRuleResult>("openstack:networking/getQosMinimumBandwidthRule:getQosMinimumBandwidthRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetQosMinimumBandwidthRuleArgs : Pulumi.InvokeArgs
    {
        [Input("direction")]
        public string? Direction { get; set; }

        /// <summary>
        /// The value of a minimum kbps bandwidth.
        /// </summary>
        [Input("minKbps")]
        public int? MinKbps { get; set; }

        /// <summary>
        /// The QoS policy reference.
        /// </summary>
        [Input("qosPolicyId", required: true)]
        public string QosPolicyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetQosMinimumBandwidthRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetQosMinimumBandwidthRuleResult
    {
        public readonly string Direction;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int MinKbps;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string QosPolicyId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetQosMinimumBandwidthRuleResult(
            string direction,
            int minKbps,
            string qosPolicyId,
            string region,
            string id)
        {
            Direction = direction;
            MinKbps = minKbps;
            QosPolicyId = qosPolicyId;
            Region = region;
            Id = id;
        }
    }
}
