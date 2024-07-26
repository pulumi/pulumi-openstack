// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Outputs
{

    [OutputType]
    public sealed class SecGroupRule
    {
        /// <summary>
        /// Required if `from_group_id` or `self` is empty. The IP range
        /// that will be the source of network traffic to the security group. Use 0.0.0.0/0
        /// to allow all IP addresses. Changing this creates a new security group rule. Cannot
        /// be combined with `from_group_id` or `self`.
        /// </summary>
        public readonly string? Cidr;
        /// <summary>
        /// Required if `cidr` or `self` is empty. The ID of a
        /// group from which to forward traffic to the parent group. Changing this creates a
        /// new security group rule. Cannot be combined with `cidr` or `self`.
        /// </summary>
        public readonly string? FromGroupId;
        /// <summary>
        /// An integer representing the lower bound of the port
        /// range to open. Changing this creates a new security group rule.
        /// </summary>
        public readonly int FromPort;
        public readonly string? Id;
        /// <summary>
        /// The protocol type that will be allowed. Changing
        /// this creates a new security group rule.
        /// </summary>
        public readonly string IpProtocol;
        /// <summary>
        /// Required if `cidr` and `from_group_id` is empty. If true,
        /// the security group itself will be added as a source to this ingress rule. Cannot
        /// be combined with `cidr` or `from_group_id`.
        /// </summary>
        public readonly bool? Self;
        /// <summary>
        /// An integer representing the upper bound of the port
        /// range to open. Changing this creates a new security group rule.
        /// </summary>
        public readonly int ToPort;

        [OutputConstructor]
        private SecGroupRule(
            string? cidr,

            string? fromGroupId,

            int fromPort,

            string? id,

            string ipProtocol,

            bool? self,

            int toPort)
        {
            Cidr = cidr;
            FromGroupId = fromGroupId;
            FromPort = fromPort;
            Id = id;
            IpProtocol = ipProtocol;
            Self = self;
            ToPort = toPort;
        }
    }
}