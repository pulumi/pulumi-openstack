// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Outputs
{

    [OutputType]
    public sealed class InstanceBlockDevice
    {
        /// <summary>
        /// The boot index of the volume. It defaults to 0.
        /// Changing this creates a new server.
        /// </summary>
        public readonly int? BootIndex;
        /// <summary>
        /// Delete the volume / block device upon
        /// termination of the instance. Defaults to false. Changing this creates a
        /// new server.
        /// </summary>
        public readonly bool? DeleteOnTermination;
        /// <summary>
        /// The type that gets created. Possible values
        /// are "volume" and "local". Changing this creates a new server.
        /// </summary>
        public readonly string? DestinationType;
        /// <summary>
        /// The low-level device type that will be used. Most
        /// common thing is to leave this empty. Changing this creates a new server.
        /// </summary>
        public readonly string? DeviceType;
        /// <summary>
        /// The low-level disk bus that will be used. Most common
        /// thing is to leave this empty. Changing this creates a new server.
        /// </summary>
        public readonly string? DiskBus;
        /// <summary>
        /// Specifies the guest server disk file system format,
        /// such as `ext2`, `ext3`, `ext4`, `xfs` or `swap`. Swap block device mappings
        /// have the following restrictions: source_type must be blank and destination_type
        /// must be local and only one swap disk per server and the size of the swap disk
        /// must be less than or equal to the swap size of the flavor. Changing this
        /// creates a new server.
        /// </summary>
        public readonly string? GuestFormat;
        /// <summary>
        /// Enable the attachment of multiattach-capable
        /// volumes.
        /// </summary>
        public readonly bool? Multiattach;
        /// <summary>
        /// The source type of the device. Must be one of
        /// "blank", "image", "volume", or "snapshot". Changing this creates a new
        /// server.
        /// </summary>
        public readonly string SourceType;
        /// <summary>
        /// The UUID of
        /// the image, volume, or snapshot. Changing this creates a new server.
        /// </summary>
        public readonly string? Uuid;
        /// <summary>
        /// The size of the volume to create (in gigabytes). Required
        /// in the following combinations: source=image and destination=volume,
        /// source=blank and destination=local, and source=blank and destination=volume.
        /// Changing this creates a new server.
        /// </summary>
        public readonly int? VolumeSize;
        /// <summary>
        /// The volume type that will be used, for example SSD
        /// or HDD storage. The available options depend on how your specific OpenStack
        /// cloud is configured and what classes of storage are provided. Changing this
        /// creates a new server.
        /// </summary>
        public readonly string? VolumeType;

        [OutputConstructor]
        private InstanceBlockDevice(
            int? bootIndex,

            bool? deleteOnTermination,

            string? destinationType,

            string? deviceType,

            string? diskBus,

            string? guestFormat,

            bool? multiattach,

            string sourceType,

            string? uuid,

            int? volumeSize,

            string? volumeType)
        {
            BootIndex = bootIndex;
            DeleteOnTermination = deleteOnTermination;
            DestinationType = destinationType;
            DeviceType = deviceType;
            DiskBus = diskBus;
            GuestFormat = guestFormat;
            Multiattach = multiattach;
            SourceType = sourceType;
            Uuid = uuid;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }
}
