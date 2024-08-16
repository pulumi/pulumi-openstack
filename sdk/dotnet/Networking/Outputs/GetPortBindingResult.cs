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
    public sealed class GetPortBindingResult
    {
        /// <summary>
        /// The ID of the host, which has the allocatee port.
        /// </summary>
        public readonly string HostId;
        /// <summary>
        /// A JSON string containing the binding profile information.
        /// </summary>
        public readonly string Profile;
        /// <summary>
        /// A map of JSON strings containing additional details for this
        /// specific binding.
        /// </summary>
        public readonly ImmutableDictionary<string, string> VifDetails;
        /// <summary>
        /// The VNIC type of the port binding.
        /// </summary>
        public readonly string VifType;
        /// <summary>
        /// VNIC type for the port.
        /// </summary>
        public readonly string VnicType;

        [OutputConstructor]
        private GetPortBindingResult(
            string hostId,

            string profile,

            ImmutableDictionary<string, string> vifDetails,

            string vifType,

            string vnicType)
        {
            HostId = hostId;
            Profile = profile;
            VifDetails = vifDetails;
            VifType = vifType;
            VnicType = vnicType;
        }
    }
}
