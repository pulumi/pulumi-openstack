// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Inputs
{

    public sealed class VolumeAttachVendorOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean to control whether
        /// to ignore volume status confirmation of the attached volume. This can be helpful
        /// to work with some OpenStack clouds which don't have the Block Storage V3 API available.
        /// </summary>
        [Input("ignoreVolumeConfirmation")]
        public Input<bool>? IgnoreVolumeConfirmation { get; set; }

        public VolumeAttachVendorOptionsGetArgs()
        {
        }
        public static new VolumeAttachVendorOptionsGetArgs Empty => new VolumeAttachVendorOptionsGetArgs();
    }
}