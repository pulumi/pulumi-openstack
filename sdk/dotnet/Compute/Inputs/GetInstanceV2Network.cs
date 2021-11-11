// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Inputs
{

    public sealed class GetInstanceV2NetworkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IPv4 address assigned to this network port.
        /// </summary>
        [Input("fixedIpV4", required: true)]
        public string FixedIpV4 { get; set; } = null!;

        /// <summary>
        /// The IPv6 address assigned to this network port.
        /// </summary>
        [Input("fixedIpV6", required: true)]
        public string FixedIpV6 { get; set; } = null!;

        /// <summary>
        /// The MAC address assigned to this network interface.
        /// </summary>
        [Input("mac", required: true)]
        public string Mac { get; set; } = null!;

        /// <summary>
        /// The name of the network
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The port UUID for this network
        /// </summary>
        [Input("port", required: true)]
        public string Port { get; set; } = null!;

        /// <summary>
        /// The UUID of the network
        /// </summary>
        [Input("uuid", required: true)]
        public string Uuid { get; set; } = null!;

        public GetInstanceV2NetworkArgs()
        {
        }
    }
}
