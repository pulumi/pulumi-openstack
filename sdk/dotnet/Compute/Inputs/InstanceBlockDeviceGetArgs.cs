// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Inputs
{

    public sealed class InstanceBlockDeviceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The boot index of the volume. It defaults to 0.
        /// Changing this creates a new server.
        /// </summary>
        [Input("bootIndex")]
        public Input<int>? BootIndex { get; set; }

        /// <summary>
        /// Delete the volume / block device upon
        /// termination of the instance. Defaults to false. Changing this creates a
        /// new server.
        /// </summary>
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        /// <summary>
        /// The type that gets created. Possible values
        /// are "volume" and "local". Changing this creates a new server.
        /// </summary>
        [Input("destinationType")]
        public Input<string>? DestinationType { get; set; }

        /// <summary>
        /// The low-level device type that will be used. Most
        /// common thing is to leave this empty. Changing this creates a new server.
        /// </summary>
        [Input("deviceType")]
        public Input<string>? DeviceType { get; set; }

        /// <summary>
        /// The low-level disk bus that will be used. Most common
        /// thing is to leave this empty. Changing this creates a new server.
        /// </summary>
        [Input("diskBus")]
        public Input<string>? DiskBus { get; set; }

        /// <summary>
        /// Specifies the guest server disk file system format,
        /// such as `ext2`, `ext3`, `ext4`, `xfs` or `swap`. Swap block device mappings
        /// have the following restrictions: source_type must be blank and destination_type
        /// must be local and only one swap disk per server and the size of the swap disk
        /// must be less than or equal to the swap size of the flavor. Changing this
        /// creates a new server.
        /// </summary>
        [Input("guestFormat")]
        public Input<string>? GuestFormat { get; set; }

        /// <summary>
        /// Enable the attachment of multiattach-capable
        /// volumes.
        /// </summary>
        [Input("multiattach")]
        public Input<bool>? Multiattach { get; set; }

        /// <summary>
        /// The source type of the device. Must be one of
        /// "blank", "image", "volume", or "snapshot". Changing this creates a new
        /// server.
        /// </summary>
        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        /// <summary>
        /// The UUID of
        /// the image, volume, or snapshot. Changing this creates a new server.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// The size of the volume to create (in gigabytes). Required
        /// in the following combinations: source=image and destination=volume,
        /// source=blank and destination=local, and source=blank and destination=volume.
        /// Changing this creates a new server.
        /// </summary>
        [Input("volumeSize")]
        public Input<int>? VolumeSize { get; set; }

        /// <summary>
        /// The volume type that will be used, for example SSD
        /// or HDD storage. The available options depend on how your specific OpenStack
        /// cloud is configured and what classes of storage are provided. Changing this
        /// creates a new server.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public InstanceBlockDeviceGetArgs()
        {
        }
        public static new InstanceBlockDeviceGetArgs Empty => new InstanceBlockDeviceGetArgs();
    }
}
