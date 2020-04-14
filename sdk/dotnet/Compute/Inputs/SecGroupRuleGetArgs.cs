// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Inputs
{

    public sealed class SecGroupRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required if `from_group_id` or `self` is empty. The IP range
        /// that will be the source of network traffic to the security group. Use 0.0.0.0/0
        /// to allow all IP addresses. Changing this creates a new security group rule. Cannot
        /// be combined with `from_group_id` or `self`.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// Required if `cidr` or `self` is empty. The ID of a
        /// group from which to forward traffic to the parent group. Changing this creates a
        /// new security group rule. Cannot be combined with `cidr` or `self`.
        /// </summary>
        [Input("fromGroupId")]
        public Input<string>? FromGroupId { get; set; }

        /// <summary>
        /// An integer representing the lower bound of the port
        /// range to open. Changing this creates a new security group rule.
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The protocol type that will be allowed. Changing
        /// this creates a new security group rule.
        /// </summary>
        [Input("ipProtocol", required: true)]
        public Input<string> IpProtocol { get; set; } = null!;

        /// <summary>
        /// Required if `cidr` and `from_group_id` is empty. If true,
        /// the security group itself will be added as a source to this ingress rule. Cannot
        /// be combined with `cidr` or `from_group_id`.
        /// </summary>
        [Input("self")]
        public Input<bool>? Self { get; set; }

        /// <summary>
        /// An integer representing the upper bound of the port
        /// range to open. Changing this creates a new security group rule.
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        public SecGroupRuleGetArgs()
        {
        }
    }
}
