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
        /// Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_dscp_marking_rule_v2.html.markdown.
        /// </summary>
        public static Task<GetQosDscpMarkingRuleResult> GetQosDscpMarkingRule(GetQosDscpMarkingRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQosDscpMarkingRuleResult>("openstack:networking/getQosDscpMarkingRule:getQosDscpMarkingRule", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetQosDscpMarkingRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The value of a DSCP mark.
        /// </summary>
        [Input("dscpMark")]
        public int? DscpMark { get; set; }

        /// <summary>
        /// The QoS policy reference.
        /// </summary>
        [Input("qosPolicyId", required: true)]
        public string QosPolicyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetQosDscpMarkingRuleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetQosDscpMarkingRuleResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int DscpMark;
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
        private GetQosDscpMarkingRuleResult(
            int dscpMark,
            string qosPolicyId,
            string region,
            string id)
        {
            DscpMark = dscpMark;
            QosPolicyId = qosPolicyId;
            Region = region;
            Id = id;
        }
    }
}
