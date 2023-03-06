// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Inputs
{

    public sealed class InstanceVendorOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to try to detach all attached
        /// ports to the vm before destroying it to make sure the port state is correct
        /// after the vm destruction. This is helpful when the port is not deleted.
        /// </summary>
        [Input("detachPortsBeforeDestroy")]
        public Input<bool>? DetachPortsBeforeDestroy { get; set; }

        /// <summary>
        /// Boolean to control whether
        /// to ignore manual confirmation of the instance resizing. This can be helpful
        /// to work with some OpenStack clouds which automatically confirm resizing of
        /// instances after some timeout.
        /// </summary>
        [Input("ignoreResizeConfirmation")]
        public Input<bool>? IgnoreResizeConfirmation { get; set; }

        public InstanceVendorOptionsGetArgs()
        {
        }
        public static new InstanceVendorOptionsGetArgs Empty => new InstanceVendorOptionsGetArgs();
    }
}
