// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking.Outputs
{

    [OutputType]
    public sealed class PortExtraDhcpOption
    {
        /// <summary>
        /// IP protocol version. Defaults to 4.
        /// </summary>
        public readonly int? IpVersion;
        /// <summary>
        /// Name of the DHCP option.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Value of the DHCP option.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private PortExtraDhcpOption(
            int? ipVersion,

            string name,

            string value)
        {
            IpVersion = ipVersion;
            Name = name;
            Value = value;
        }
    }
}
