// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage.Inputs
{

    public sealed class VolumeV2AttachmentArgs : Pulumi.ResourceArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public VolumeV2AttachmentArgs()
        {
        }
    }
}
