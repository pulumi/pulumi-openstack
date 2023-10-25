// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetQosBandwidthLimitRule
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS bandwidth limit rule.
        /// </summary>
        public static Task<GetQosBandwidthLimitRuleResult> InvokeAsync(GetQosBandwidthLimitRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQosBandwidthLimitRuleResult>("openstack:networking/getQosBandwidthLimitRule:getQosBandwidthLimitRule", args ?? new GetQosBandwidthLimitRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS bandwidth limit rule.
        /// </summary>
        public static Output<GetQosBandwidthLimitRuleResult> Invoke(GetQosBandwidthLimitRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQosBandwidthLimitRuleResult>("openstack:networking/getQosBandwidthLimitRule:getQosBandwidthLimitRule", args ?? new GetQosBandwidthLimitRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQosBandwidthLimitRuleArgs : global::Pulumi.InvokeArgs
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
        public static new GetQosBandwidthLimitRuleArgs Empty => new GetQosBandwidthLimitRuleArgs();
    }

    public sealed class GetQosBandwidthLimitRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The maximum burst size in kilobits of a QoS bandwidth limit rule.
        /// </summary>
        [Input("maxBurstKbps")]
        public Input<int>? MaxBurstKbps { get; set; }

        /// <summary>
        /// The maximum kilobits per second of a QoS bandwidth limit rule.
        /// </summary>
        [Input("maxKbps")]
        public Input<int>? MaxKbps { get; set; }

        /// <summary>
        /// The QoS policy reference.
        /// </summary>
        [Input("qosPolicyId", required: true)]
        public Input<string> QosPolicyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetQosBandwidthLimitRuleInvokeArgs()
        {
        }
        public static new GetQosBandwidthLimitRuleInvokeArgs Empty => new GetQosBandwidthLimitRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetQosBandwidthLimitRuleResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
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

        [OutputConstructor]
        private GetQosBandwidthLimitRuleResult(
            string direction,

            string id,

            int maxBurstKbps,

            int maxKbps,

            string qosPolicyId,

            string region)
        {
            Direction = direction;
            Id = id;
            MaxBurstKbps = maxBurstKbps;
            MaxKbps = maxKbps;
            QosPolicyId = qosPolicyId;
            Region = region;
        }
    }
}
