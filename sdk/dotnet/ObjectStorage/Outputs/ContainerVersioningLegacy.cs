// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        public readonly string Location;
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
