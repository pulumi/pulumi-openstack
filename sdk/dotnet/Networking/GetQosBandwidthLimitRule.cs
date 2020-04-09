// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS bandwidth limit rule.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_bandwidth_limit_rule_v2.html.markdown.
        /// </summary>
        [Obsolete("Use GetQosBandwidthLimitRule.InvokeAsync() instead")]
        public static Task<GetQosBandwidthLimitRuleResult> GetQosBandwidthLimitRule(GetQosBandwidthLimitRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQosBandwidthLimitRuleResult>("openstack:networking/getQosBandwidthLimitRule:getQosBandwidthLimitRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetQosBandwidthLimitRule
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS bandwidth limit rule.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_bandwidth_limit_rule_v2.html.markdown.
        /// </summary>
        public static Task<GetQosBandwidthLimitRuleResult> InvokeAsync(GetQosBandwidthLimitRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQosBandwidthLimitRuleResult>("openstack:networking/getQosBandwidthLimitRule:getQosBandwidthLimitRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetQosBandwidthLimitRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The maximum burst size in kilobits of a QoS bandwidth limit rule.
        /// </summary>
        [Input("maxBurstKbps")]
        public int? MaxBurstKbps { get; set; }

        /// <summary>
        /// The maximum kilobits per second of a QoS bandwidth limit rule.
        /// </summary>
        [Input("maxKbps")]
        public int? MaxKbps { get; set; }

        /// <summary>
        /// The QoS policy reference.
        /// </summary>
        [Input("qosPolicyId", required: true)]
        public string QosPolicyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetQosBandwidthLimitRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetQosBandwidthLimitRuleResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int MaxBurstKbps;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int MaxKbps;
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
        private GetQosBandwidthLimitRuleResult(
            string direction,
            int maxBurstKbps,
            int maxKbps,
            string qosPolicyId,
            string region,
            string id)
        {
            Direction = direction;
            MaxBurstKbps = maxBurstKbps;
            MaxKbps = maxKbps;
            QosPolicyId = qosPolicyId;
            Region = region;
            Id = id;
        }
    }
}
