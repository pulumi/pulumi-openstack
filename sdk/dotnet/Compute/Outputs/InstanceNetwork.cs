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
    public sealed class InstanceNetwork
    {
        /// <summary>
        /// Specifies if this network should be used for
        /// provisioning access. Accepts true or false. Defaults to false.
        /// </summary>
        public readonly bool? AccessNetwork;
        /// <summary>
        /// Specifies a fixed IPv4 address to be used on this
        /// network. Changing this creates a new server.
        /// </summary>
        public readonly string? FixedIpV4;
        public readonly string? FixedIpV6;
        public readonly string? Mac;
        /// <summary>
        /// The human-readable
        /// name of the network. Changing this creates a new server.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The port UUID of a
        /// network to attach to the server. Changing this creates a new server.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// The network UUID to
        /// attach to the server. Changing this creates a new server.
        /// </summary>
        public readonly string? Uuid;

        [OutputConstructor]
        private InstanceNetwork(
            bool? accessNetwork,

            string? fixedIpV4,

            string? fixedIpV6,

            string? mac,

            string? name,

            string? port,

            string? uuid)
        {
            AccessNetwork = accessNetwork;
            FixedIpV4 = fixedIpV4;
            FixedIpV6 = fixedIpV6;
            Mac = mac;
            Name = name;
            Port = port;
            Uuid = uuid;
        }
    }
}
