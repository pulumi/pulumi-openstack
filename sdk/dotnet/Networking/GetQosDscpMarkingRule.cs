// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetQosDscpMarkingRule
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var qosDscpMarkingRule1 = OpenStack.Networking.GetQosDscpMarkingRule.Invoke(new()
        ///     {
        ///         DscpMark = 26,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetQosDscpMarkingRuleResult> InvokeAsync(GetQosDscpMarkingRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQosDscpMarkingRuleResult>("openstack:networking/getQosDscpMarkingRule:getQosDscpMarkingRule", args ?? new GetQosDscpMarkingRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var qosDscpMarkingRule1 = OpenStack.Networking.GetQosDscpMarkingRule.Invoke(new()
        ///     {
        ///         DscpMark = 26,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetQosDscpMarkingRuleResult> Invoke(GetQosDscpMarkingRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQosDscpMarkingRuleResult>("openstack:networking/getQosDscpMarkingRule:getQosDscpMarkingRule", args ?? new GetQosDscpMarkingRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQosDscpMarkingRuleArgs : global::Pulumi.InvokeArgs
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
        public static new GetQosDscpMarkingRuleArgs Empty => new GetQosDscpMarkingRuleArgs();
    }

    public sealed class GetQosDscpMarkingRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The value of a DSCP mark.
        /// </summary>
        [Input("dscpMark")]
        public Input<int>? DscpMark { get; set; }

        /// <summary>
        /// The QoS policy reference.
        /// </summary>
        [Input("qosPolicyId", required: true)]
        public Input<string> QosPolicyId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetQosDscpMarkingRuleInvokeArgs()
        {
        }
        public static new GetQosDscpMarkingRuleInvokeArgs Empty => new GetQosDscpMarkingRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetQosDscpMarkingRuleResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int DscpMark;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string QosPolicyId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetQosDscpMarkingRuleResult(
            int dscpMark,

            string id,

            string qosPolicyId,

            string region)
        {
            DscpMark = dscpMark;
            Id = id;
            QosPolicyId = qosPolicyId;
            Region = region;
        }
    }
}
