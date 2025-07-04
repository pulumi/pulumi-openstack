// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.ObjectStorage.Outputs
{

    [OutputType]
    public sealed class ContainerVersioningLegacy
    {
        /// <summary>
        /// Container in which versions will be stored.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Versioning type which can be `versions` or `history`
        /// according to [OpenStack
        /// documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ContainerVersioningLegacy(
            string location,

            string type)
        {
            Location = location;
            Type = type;
        }
    }
}
