// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.ObjectStorage.Inputs
{

    public sealed class ContainerVersioningLegacyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container in which versions will be stored.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Versioning type which can be `versions` or `history`
        /// according to [OpenStack
        /// documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ContainerVersioningLegacyArgs()
        {
        }
        public static new ContainerVersioningLegacyArgs Empty => new ContainerVersioningLegacyArgs();
    }
}
