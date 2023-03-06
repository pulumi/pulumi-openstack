// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking.Inputs
{

    public sealed class PortExtraDhcpOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP protocol version. Defaults to 4.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// Name of the DHCP option.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Value of the DHCP option.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public PortExtraDhcpOptionArgs()
        {
        }
        public static new PortExtraDhcpOptionArgs Empty => new PortExtraDhcpOptionArgs();
    }
}
