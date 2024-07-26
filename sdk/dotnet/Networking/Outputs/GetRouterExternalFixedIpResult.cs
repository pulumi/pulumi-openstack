// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking.Outputs
{

    [OutputType]
    public sealed class GetRouterExternalFixedIpResult
    {
        /// <summary>
        /// The IP address to set on the router.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// Subnet in which the fixed IP belongs to.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private GetRouterExternalFixedIpResult(
            string? ipAddress,

            string? subnetId)
        {
            IpAddress = ipAddress;
            SubnetId = subnetId;
        }
    }
}