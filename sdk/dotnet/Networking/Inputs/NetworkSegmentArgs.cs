// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking.Inputs
{

    public sealed class NetworkSegmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of physical network.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The physical network where this network is implemented.
        /// </summary>
        [Input("physicalNetwork")]
        public Input<string>? PhysicalNetwork { get; set; }

        /// <summary>
        /// An isolated segment on the physical network.
        /// </summary>
        [Input("segmentationId")]
        public Input<int>? SegmentationId { get; set; }

        public NetworkSegmentArgs()
        {
        }
    }
}
