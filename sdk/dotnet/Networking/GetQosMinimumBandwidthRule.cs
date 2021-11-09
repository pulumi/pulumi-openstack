// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.OpenStack.Networking
{
    public static class GetQosMinimumBandwidthRule
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var qosMinBwRule1 = Output.Create(OpenStack.Networking.GetQosMinimumBandwidthRule.InvokeAsync(new OpenStack.Networking.GetQosMinimumBandwidthRuleArgs
        ///         {
        ///             MinKbps = 2000,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetQosMinimumBandwidthRuleResult> InvokeAsync(GetQosMinimumBandwidthRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQosMinimumBandwidthRuleResult>("openstack:networking/getQosMinimumBandwidthRule:getQosMinimumBandwidthRule", args ?? new GetQosMinimumBandwidthRuleArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var qosMinBwRule1 = Output.Create(OpenStack.Networking.GetQosMinimumBandwidthRule.InvokeAsync(new OpenStack.Networking.GetQosMinimumBandwidthRuleArgs
        ///         {
        ///             MinKbps = 2000,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetQosMinimumBandwidthRuleResult> Invoke(GetQosMinimumBandwidthRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetQosMinimumBandwidthRuleResult>("openstack:networking/getQosMinimumBandwidthRule:getQosMinimumBandwidthRule", args ?? new GetQosMinimumBandwidthRuleInvokeArgs(), options.WithVersion());
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

    public sealed class GetQosMinimumBandwidthRuleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The value of a minimum kbps bandwidth.
        /// </summary>
        [Input("minKbps")]
        public Input<int>? MinKbps { get; set; }

        /// <summary>
        /// The QoS policy reference.
        /// </summary>
        [Input("qosPolicyId", required: true)]
        public Input<string> QosPolicyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetQosMinimumBandwidthRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetQosMinimumBandwidthRuleResult
    {
        public readonly string Direction;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
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

        [OutputConstructor]
        private GetQosMinimumBandwidthRuleResult(
            string direction,

            string id,

            int minKbps,

            string qosPolicyId,

            string region)
        {
            Direction = direction;
            Id = id;
            MinKbps = minKbps;
            QosPolicyId = qosPolicyId;
            Region = region;
        }
    }
}
