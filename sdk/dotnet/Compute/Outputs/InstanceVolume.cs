// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Outputs
{

    [OutputType]
    public sealed class InstanceVolume
    {
        public readonly string? Device;
        public readonly string? Id;
        public readonly string VolumeId;

        [OutputConstructor]
        private InstanceVolume(
            string? device,

            string? id,

            string volumeId)
        {
            Device = device;
            Id = id;
            VolumeId = volumeId;
        }
    }
}
