// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage.Outputs
{

    [OutputType]
    public sealed class GetVolumeV3AttachmentResult
    {
        public readonly string Device;
        public readonly string Id;
        public readonly string InstanceId;

        [OutputConstructor]
        private GetVolumeV3AttachmentResult(
            string device,

            string id,

            string instanceId)
        {
            Device = device;
            Id = id;
            InstanceId = instanceId;
        }
    }
}
