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
    public sealed class NetworkSegment
    {
        /// <summary>
        /// The type of physical network.
        /// </summary>
        public readonly string? NetworkType;
        /// <summary>
        /// The physical network where this network is implemented.
        /// </summary>
        public readonly string? PhysicalNetwork;
        /// <summary>
        /// An isolated segment on the physical network.
        /// </summary>
        public readonly int? SegmentationId;

        [OutputConstructor]
        private NetworkSegment(
            string? networkType,

            string? physicalNetwork,

            int? segmentationId)
        {
            NetworkType = networkType;
            PhysicalNetwork = physicalNetwork;
            SegmentationId = segmentationId;
        }
    }
}
