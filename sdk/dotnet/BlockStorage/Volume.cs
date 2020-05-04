// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage
{
    /// <summary>
    /// Manages a V3 volume resource within OpenStack.
    /// </summary>
    public partial class Volume : Pulumi.CustomResource
    {
        /// <summary>
        /// If a volume is attached to an instance, this attribute will
        /// display the Attachment ID, Instance ID, and the Device as the Instance
        /// sees it.
        /// </summary>
        [Output("attachments")]
        public Output<ImmutableArray<Outputs.VolumeAttachment>> Attachments { get; private set; } = null!;

        /// <summary>
        /// The availability zone for the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The consistency group to place the volume
        /// in.
        /// </summary>
        [Output("consistencyGroupId")]
        public Output<string?> ConsistencyGroupId { get; private set; } = null!;

        /// <summary>
        /// A description of the volume. Changing this updates
        /// the volume's description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// When this option is set it allows extending
        /// attached volumes. Note: updating size of an attached volume requires Cinder
        /// support for version 3.42 and a compatible storage driver.
        /// </summary>
        [Output("enableOnlineResize")]
        public Output<bool?> EnableOnlineResize { get; private set; } = null!;

        /// <summary>
        /// The image ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        /// <summary>
        /// Metadata key/value pairs to associate with the volume.
        /// Changing this updates the existing volume metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Allow the volume to be attached to more than one Compute instance.
        /// </summary>
        [Output("multiattach")]
        public Output<bool?> Multiattach { get; private set; } = null!;

        /// <summary>
        /// A unique name for the volume. Changing this updates the
        /// volume's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new volume.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Provide the Cinder scheduler with hints on where
        /// to instantiate a volume in the OpenStack cloud. The available hints are described below.
        /// </summary>
        [Output("schedulerHints")]
        public Output<ImmutableArray<Outputs.VolumeSchedulerHint>> SchedulerHints { get; private set; } = null!;

        /// <summary>
        /// The size of the volume to create (in gigabytes).
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The snapshot ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// The volume ID to replicate with.
        /// </summary>
        [Output("sourceReplica")]
        public Output<string?> SourceReplica { get; private set; } = null!;

        /// <summary>
        /// The volume ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Output("sourceVolId")]
        public Output<string?> SourceVolId { get; private set; } = null!;

        /// <summary>
        /// The type of volume to create.
        /// Changing this creates a new volume.
        /// </summary>
        [Output("volumeType")]
        public Output<string> VolumeType { get; private set; } = null!;


        /// <summary>
        /// Create a Volume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Volume(string name, VolumeArgs args, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/volume:Volume", name, args ?? new VolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Volume(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/volume:Volume", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Volume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Volume Get(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new Volume(name, id, state, options);
        }
    }

    public sealed class VolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The availability zone for the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The consistency group to place the volume
        /// in.
        /// </summary>
        [Input("consistencyGroupId")]
        public Input<string>? ConsistencyGroupId { get; set; }

        /// <summary>
        /// A description of the volume. Changing this updates
        /// the volume's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When this option is set it allows extending
        /// attached volumes. Note: updating size of an attached volume requires Cinder
        /// support for version 3.42 and a compatible storage driver.
        /// </summary>
        [Input("enableOnlineResize")]
        public Input<bool>? EnableOnlineResize { get; set; }

        /// <summary>
        /// The image ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Metadata key/value pairs to associate with the volume.
        /// Changing this updates the existing volume metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// Allow the volume to be attached to more than one Compute instance.
        /// </summary>
        [Input("multiattach")]
        public Input<bool>? Multiattach { get; set; }

        /// <summary>
        /// A unique name for the volume. Changing this updates the
        /// volume's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new volume.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("schedulerHints")]
        private InputList<Inputs.VolumeSchedulerHintArgs>? _schedulerHints;

        /// <summary>
        /// Provide the Cinder scheduler with hints on where
        /// to instantiate a volume in the OpenStack cloud. The available hints are described below.
        /// </summary>
        public InputList<Inputs.VolumeSchedulerHintArgs> SchedulerHints
        {
            get => _schedulerHints ?? (_schedulerHints = new InputList<Inputs.VolumeSchedulerHintArgs>());
            set => _schedulerHints = value;
        }

        /// <summary>
        /// The size of the volume to create (in gigabytes).
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// The snapshot ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// The volume ID to replicate with.
        /// </summary>
        [Input("sourceReplica")]
        public Input<string>? SourceReplica { get; set; }

        /// <summary>
        /// The volume ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("sourceVolId")]
        public Input<string>? SourceVolId { get; set; }

        /// <summary>
        /// The type of volume to create.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public VolumeArgs()
        {
        }
    }

    public sealed class VolumeState : Pulumi.ResourceArgs
    {
        [Input("attachments")]
        private InputList<Inputs.VolumeAttachmentGetArgs>? _attachments;

        /// <summary>
        /// If a volume is attached to an instance, this attribute will
        /// display the Attachment ID, Instance ID, and the Device as the Instance
        /// sees it.
        /// </summary>
        public InputList<Inputs.VolumeAttachmentGetArgs> Attachments
        {
            get => _attachments ?? (_attachments = new InputList<Inputs.VolumeAttachmentGetArgs>());
            set => _attachments = value;
        }

        /// <summary>
        /// The availability zone for the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The consistency group to place the volume
        /// in.
        /// </summary>
        [Input("consistencyGroupId")]
        public Input<string>? ConsistencyGroupId { get; set; }

        /// <summary>
        /// A description of the volume. Changing this updates
        /// the volume's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When this option is set it allows extending
        /// attached volumes. Note: updating size of an attached volume requires Cinder
        /// support for version 3.42 and a compatible storage driver.
        /// </summary>
        [Input("enableOnlineResize")]
        public Input<bool>? EnableOnlineResize { get; set; }

        /// <summary>
        /// The image ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Metadata key/value pairs to associate with the volume.
        /// Changing this updates the existing volume metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// Allow the volume to be attached to more than one Compute instance.
        /// </summary>
        [Input("multiattach")]
        public Input<bool>? Multiattach { get; set; }

        /// <summary>
        /// A unique name for the volume. Changing this updates the
        /// volume's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new volume.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("schedulerHints")]
        private InputList<Inputs.VolumeSchedulerHintGetArgs>? _schedulerHints;

        /// <summary>
        /// Provide the Cinder scheduler with hints on where
        /// to instantiate a volume in the OpenStack cloud. The available hints are described below.
        /// </summary>
        public InputList<Inputs.VolumeSchedulerHintGetArgs> SchedulerHints
        {
            get => _schedulerHints ?? (_schedulerHints = new InputList<Inputs.VolumeSchedulerHintGetArgs>());
            set => _schedulerHints = value;
        }

        /// <summary>
        /// The size of the volume to create (in gigabytes).
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The snapshot ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        /// <summary>
        /// The volume ID to replicate with.
        /// </summary>
        [Input("sourceReplica")]
        public Input<string>? SourceReplica { get; set; }

        /// <summary>
        /// The volume ID from which to create the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("sourceVolId")]
        public Input<string>? SourceVolId { get; set; }

        /// <summary>
        /// The type of volume to create.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public VolumeState()
        {
        }
    }
}
