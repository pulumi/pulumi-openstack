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
    /// Manages a V1 volume resource within OpenStack.
    /// 
    /// ## Import
    /// 
    /// Volumes can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:blockstorage/volumeV1:VolumeV1 volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:blockstorage/volumeV1:VolumeV1")]
    public partial class VolumeV1 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If a volume is attached to an instance, this attribute will
        /// display the Attachment ID, Instance ID, and the Device as the Instance
        /// sees it.
        /// </summary>
        [Output("attachments")]
        public Output<ImmutableArray<Outputs.VolumeV1Attachment>> Attachments { get; private set; } = null!;

        /// <summary>
        /// The availability zone for the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// A description of the volume. Changing this updates
        /// the volume's description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

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
        /// The size of the volume to create (in gigabytes). Changing
        /// this creates a new volume.
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
        /// Create a VolumeV1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VolumeV1(string name, VolumeV1Args args, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/volumeV1:VolumeV1", name, args ?? new VolumeV1Args(), MakeResourceOptions(options, ""))
        {
        }

        private VolumeV1(string name, Input<string> id, VolumeV1State? state = null, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/volumeV1:VolumeV1", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VolumeV1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VolumeV1 Get(string name, Input<string> id, VolumeV1State? state = null, CustomResourceOptions? options = null)
        {
            return new VolumeV1(name, id, state, options);
        }
    }

    public sealed class VolumeV1Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The availability zone for the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// A description of the volume. Changing this updates
        /// the volume's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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

        /// <summary>
        /// The size of the volume to create (in gigabytes). Changing
        /// this creates a new volume.
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

        public VolumeV1Args()
        {
        }
        public static new VolumeV1Args Empty => new VolumeV1Args();
    }

    public sealed class VolumeV1State : global::Pulumi.ResourceArgs
    {
        [Input("attachments")]
        private InputList<Inputs.VolumeV1AttachmentGetArgs>? _attachments;

        /// <summary>
        /// If a volume is attached to an instance, this attribute will
        /// display the Attachment ID, Instance ID, and the Device as the Instance
        /// sees it.
        /// </summary>
        public InputList<Inputs.VolumeV1AttachmentGetArgs> Attachments
        {
            get => _attachments ?? (_attachments = new InputList<Inputs.VolumeV1AttachmentGetArgs>());
            set => _attachments = value;
        }

        /// <summary>
        /// The availability zone for the volume.
        /// Changing this creates a new volume.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// A description of the volume. Changing this updates
        /// the volume's description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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

        /// <summary>
        /// The size of the volume to create (in gigabytes). Changing
        /// this creates a new volume.
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

        public VolumeV1State()
        {
        }
        public static new VolumeV1State Empty => new VolumeV1State();
    }
}
