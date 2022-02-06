// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.LoadBalancer.Outputs
{

    [OutputType]
    public sealed class MembersMember
    {
        /// <summary>
        /// The IP address of the members to receive traffic from
        /// the load balancer.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The administrative state of the member.
        /// A valid value is true (UP) or false (DOWN). Defaults to true.
        /// </summary>
        public readonly bool? AdminStateUp;
        /// <summary>
        /// A bool that indicates whether the member is
        /// backup. **Requires octavia minor version 2.1 or later**.
        /// </summary>
        public readonly bool? Backup;
        /// <summary>
        /// The unique ID for the members.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Human-readable name for the member.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The port on which to listen for client traffic.
        /// </summary>
        public readonly int ProtocolPort;
        /// <summary>
        /// The subnet in which to access the member.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// A positive integer value that indicates the relative
        /// portion of traffic that this members should receive from the pool. For
        /// example, a member with a weight of 10 receives five times as much traffic
        /// as a member with a weight of 2. Defaults to 1.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private MembersMember(
            string address,

            bool? adminStateUp,

            bool? backup,

            string? id,

            string? name,

            int protocolPort,

            string? subnetId,

            int? weight)
        {
            Address = address;
            AdminStateUp = adminStateUp;
            Backup = backup;
            Id = id;
            Name = name;
            ProtocolPort = protocolPort;
            SubnetId = subnetId;
            Weight = weight;
        }
    }
}
