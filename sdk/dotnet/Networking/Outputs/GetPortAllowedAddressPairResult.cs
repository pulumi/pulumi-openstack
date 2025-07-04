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
    public sealed class GetPortAllowedAddressPairResult
    {
        /// <summary>
        /// The additional IP address.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The MAC address of the port.
        /// </summary>
        public readonly string MacAddress;

        [OutputConstructor]
        private GetPortAllowedAddressPairResult(
            string ipAddress,

            string macAddress)
        {
            IpAddress = ipAddress;
            MacAddress = macAddress;
        }
    }
}
