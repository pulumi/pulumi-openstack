// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking.Inputs
{

    public sealed class SubnetHostRouteGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR.
        /// </summary>
        [Input("destinationCidr", required: true)]
        public Input<string> DestinationCidr { get; set; } = null!;

        /// <summary>
        /// The next hop in the route.
        /// </summary>
        [Input("nextHop", required: true)]
        public Input<string> NextHop { get; set; } = null!;

        public SubnetHostRouteGetArgs()
        {
        }
        public static new SubnetHostRouteGetArgs Empty => new SubnetHostRouteGetArgs();
    }
}
