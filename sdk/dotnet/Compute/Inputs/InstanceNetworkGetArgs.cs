// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Inputs
{

    public sealed class InstanceNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies if this network should be used for
        /// provisioning access. Accepts true or false. Defaults to false.
        /// </summary>
        [Input("accessNetwork")]
        public Input<bool>? AccessNetwork { get; set; }

        /// <summary>
        /// Specifies a fixed IPv4 address to be used on this
        /// network. Changing this creates a new server.
        /// </summary>
        [Input("fixedIpV4")]
        public Input<string>? FixedIpV4 { get; set; }

        [Input("fixedIpV6")]
        public Input<string>? FixedIpV6 { get; set; }

        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// The human-readable
        /// name of the network. Changing this creates a new server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port UUID of a
        /// network to attach to the server. Changing this creates a new server.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The network UUID to
        /// attach to the server. Changing this creates a new server.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public InstanceNetworkGetArgs()
        {
        }
        public static new InstanceNetworkGetArgs Empty => new InstanceNetworkGetArgs();
    }
}
