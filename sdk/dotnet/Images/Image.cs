// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Images
{
    /// <summary>
    /// Manages a V2 Image resource within OpenStack Glance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rancheros = new OpenStack.Images.Image("rancheros", new OpenStack.Images.ImageArgs
    ///         {
    ///             ContainerFormat = "bare",
    ///             DiskFormat = "qcow2",
    ///             ImageSourceUrl = "https://releases.rancher.com/os/latest/rancheros-openstack.img",
    ///             Properties = 
    ///             {
    ///                 { "key", "value" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Notes
    /// 
    /// ### Properties
    /// 
    /// This resource supports the ability to add properties to a resource during
    /// creation as well as add, update, and delete properties during an update of this
    /// resource.
    /// 
    /// Newer versions of OpenStack are adding some read-only properties to each image.
    /// These properties start with the prefix `os_`. If these properties are detected,
    /// this resource will automatically reconcile these with the user-provided
    /// properties.
    /// 
    /// In addition, the `direct_url` property is also automatically reconciled if the
    /// Image Service set it.
    /// </summary>
    public partial class Image : Pulumi.CustomResource
    {
        /// <summary>
        /// The checksum of the data associated with the image.
        /// </summary>
        [Output("checksum")]
        public Output<string> Checksum { get; private set; } = null!;

        /// <summary>
        /// The container format. Must be one of
        /// "ami", "ari", "aki", "bare", "ovf".
        /// </summary>
        [Output("containerFormat")]
        public Output<string> ContainerFormat { get; private set; } = null!;

        /// <summary>
        /// The date the image was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The disk format. Must be one of
        /// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
        /// </summary>
        [Output("diskFormat")]
        public Output<string> DiskFormat { get; private set; } = null!;

        /// <summary>
        /// the trailing path after the glance
        /// endpoint that represent the location of the image
        /// or the path to retrieve it.
        /// </summary>
        [Output("file")]
        public Output<string> File { get; private set; } = null!;

        [Output("imageCachePath")]
        public Output<string?> ImageCachePath { get; private set; } = null!;

        /// <summary>
        /// This is the url of the raw image. If `web_download`
        /// is not used, then the image will be downloaded in the `image_cache_path` before
        /// being uploaded to Glance.
        /// Conflicts with `local_file_path`.
        /// </summary>
        [Output("imageSourceUrl")]
        public Output<string?> ImageSourceUrl { get; private set; } = null!;

        /// <summary>
        /// This is the filepath of the raw image file
        /// that will be uploaded to Glance. Conflicts with `image_source_url` and
        /// `web_download`.
        /// </summary>
        [Output("localFilePath")]
        public Output<string?> LocalFilePath { get; private set; } = null!;

        /// <summary>
        /// The metadata associated with the image.
        /// Image metadata allow for meaningfully define the image properties
        /// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Amount of disk space (in GB) required to boot image.
        /// Defaults to 0.
        /// </summary>
        [Output("minDiskGb")]
        public Output<int?> MinDiskGb { get; private set; } = null!;

        /// <summary>
        /// Amount of ram (in MB) required to boot image.
        /// Defauts to 0.
        /// </summary>
        [Output("minRamMb")]
        public Output<int?> MinRamMb { get; private set; } = null!;

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the openstack user who owns the image.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// A map of key/value pairs to set freeform
        /// information about an image. See the "Notes" section for further
        /// information about properties.
        /// </summary>
        [Output("properties")]
        public Output<ImmutableDictionary<string, object>> Properties { get; private set; } = null!;

        /// <summary>
        /// If true, image will not be deletable.
        /// Defaults to false.
        /// </summary>
        [Output("protected")]
        public Output<bool?> Protected { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to create an Image that can be used with
        /// a compute instance. If omitted, the `region` argument of the provider
        /// is used. Changing this creates a new Image.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The path to the JSON-schema that represent
        /// the image or image
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// The size in bytes of the data associated with the image.
        /// </summary>
        [Output("sizeBytes")]
        public Output<int> SizeBytes { get; private set; } = null!;

        /// <summary>
        /// The status of the image. It can be "queued", "active"
        /// or "saving".
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tags of the image. It must be a list of strings.
        /// At this time, it is not possible to delete all tags of an image.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// (**Deprecated** - use `updated_at` instead)
        /// </summary>
        [Output("updateAt")]
        public Output<string> UpdateAt { get; private set; } = null!;

        /// <summary>
        /// The date the image was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// If false, the checksum will not be verified
        /// once the image is finished uploading. Conflicts with `web_download`.
        /// Defaults to true when not using `web_download`.
        /// </summary>
        [Output("verifyChecksum")]
        public Output<bool?> VerifyChecksum { get; private set; } = null!;

        /// <summary>
        /// The visibility of the image. Must be one of
        /// "public", "private", "community", or "shared". The ability to set the
        /// visibility depends upon the configuration of the OpenStack cloud.
        /// </summary>
        [Output("visibility")]
        public Output<string?> Visibility { get; private set; } = null!;

        /// <summary>
        /// If true, the "web-download" import method will
        /// be used to let Openstack download the image directly from the remote source.
        /// Conflicts with `local_file_path`. Defaults to false.
        /// </summary>
        [Output("webDownload")]
        public Output<bool?> WebDownload { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs args, CustomResourceOptions? options = null)
            : base("openstack:images/image:Image", name, args ?? new ImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
            : base("openstack:images/image:Image", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, ImageState? state = null, CustomResourceOptions? options = null)
        {
            return new Image(name, id, state, options);
        }
    }

    public sealed class ImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The container format. Must be one of
        /// "ami", "ari", "aki", "bare", "ovf".
        /// </summary>
        [Input("containerFormat", required: true)]
        public Input<string> ContainerFormat { get; set; } = null!;

        /// <summary>
        /// The disk format. Must be one of
        /// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
        /// </summary>
        [Input("diskFormat", required: true)]
        public Input<string> DiskFormat { get; set; } = null!;

        [Input("imageCachePath")]
        public Input<string>? ImageCachePath { get; set; }

        /// <summary>
        /// This is the url of the raw image. If `web_download`
        /// is not used, then the image will be downloaded in the `image_cache_path` before
        /// being uploaded to Glance.
        /// Conflicts with `local_file_path`.
        /// </summary>
        [Input("imageSourceUrl")]
        public Input<string>? ImageSourceUrl { get; set; }

        /// <summary>
        /// This is the filepath of the raw image file
        /// that will be uploaded to Glance. Conflicts with `image_source_url` and
        /// `web_download`.
        /// </summary>
        [Input("localFilePath")]
        public Input<string>? LocalFilePath { get; set; }

        /// <summary>
        /// Amount of disk space (in GB) required to boot image.
        /// Defaults to 0.
        /// </summary>
        [Input("minDiskGb")]
        public Input<int>? MinDiskGb { get; set; }

        /// <summary>
        /// Amount of ram (in MB) required to boot image.
        /// Defauts to 0.
        /// </summary>
        [Input("minRamMb")]
        public Input<int>? MinRamMb { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("properties")]
        private InputMap<object>? _properties;

        /// <summary>
        /// A map of key/value pairs to set freeform
        /// information about an image. See the "Notes" section for further
        /// information about properties.
        /// </summary>
        public InputMap<object> Properties
        {
            get => _properties ?? (_properties = new InputMap<object>());
            set => _properties = value;
        }

        /// <summary>
        /// If true, image will not be deletable.
        /// Defaults to false.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to create an Image that can be used with
        /// a compute instance. If omitted, the `region` argument of the provider
        /// is used. Changing this creates a new Image.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags of the image. It must be a list of strings.
        /// At this time, it is not possible to delete all tags of an image.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// If false, the checksum will not be verified
        /// once the image is finished uploading. Conflicts with `web_download`.
        /// Defaults to true when not using `web_download`.
        /// </summary>
        [Input("verifyChecksum")]
        public Input<bool>? VerifyChecksum { get; set; }

        /// <summary>
        /// The visibility of the image. Must be one of
        /// "public", "private", "community", or "shared". The ability to set the
        /// visibility depends upon the configuration of the OpenStack cloud.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// If true, the "web-download" import method will
        /// be used to let Openstack download the image directly from the remote source.
        /// Conflicts with `local_file_path`. Defaults to false.
        /// </summary>
        [Input("webDownload")]
        public Input<bool>? WebDownload { get; set; }

        public ImageArgs()
        {
        }
    }

    public sealed class ImageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The checksum of the data associated with the image.
        /// </summary>
        [Input("checksum")]
        public Input<string>? Checksum { get; set; }

        /// <summary>
        /// The container format. Must be one of
        /// "ami", "ari", "aki", "bare", "ovf".
        /// </summary>
        [Input("containerFormat")]
        public Input<string>? ContainerFormat { get; set; }

        /// <summary>
        /// The date the image was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The disk format. Must be one of
        /// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
        /// </summary>
        [Input("diskFormat")]
        public Input<string>? DiskFormat { get; set; }

        /// <summary>
        /// the trailing path after the glance
        /// endpoint that represent the location of the image
        /// or the path to retrieve it.
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        [Input("imageCachePath")]
        public Input<string>? ImageCachePath { get; set; }

        /// <summary>
        /// This is the url of the raw image. If `web_download`
        /// is not used, then the image will be downloaded in the `image_cache_path` before
        /// being uploaded to Glance.
        /// Conflicts with `local_file_path`.
        /// </summary>
        [Input("imageSourceUrl")]
        public Input<string>? ImageSourceUrl { get; set; }

        /// <summary>
        /// This is the filepath of the raw image file
        /// that will be uploaded to Glance. Conflicts with `image_source_url` and
        /// `web_download`.
        /// </summary>
        [Input("localFilePath")]
        public Input<string>? LocalFilePath { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// The metadata associated with the image.
        /// Image metadata allow for meaningfully define the image properties
        /// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// Amount of disk space (in GB) required to boot image.
        /// Defaults to 0.
        /// </summary>
        [Input("minDiskGb")]
        public Input<int>? MinDiskGb { get; set; }

        /// <summary>
        /// Amount of ram (in MB) required to boot image.
        /// Defauts to 0.
        /// </summary>
        [Input("minRamMb")]
        public Input<int>? MinRamMb { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the openstack user who owns the image.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("properties")]
        private InputMap<object>? _properties;

        /// <summary>
        /// A map of key/value pairs to set freeform
        /// information about an image. See the "Notes" section for further
        /// information about properties.
        /// </summary>
        public InputMap<object> Properties
        {
            get => _properties ?? (_properties = new InputMap<object>());
            set => _properties = value;
        }

        /// <summary>
        /// If true, image will not be deletable.
        /// Defaults to false.
        /// </summary>
        [Input("protected")]
        public Input<bool>? Protected { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to create an Image that can be used with
        /// a compute instance. If omitted, the `region` argument of the provider
        /// is used. Changing this creates a new Image.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The path to the JSON-schema that represent
        /// the image or image
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// The size in bytes of the data associated with the image.
        /// </summary>
        [Input("sizeBytes")]
        public Input<int>? SizeBytes { get; set; }

        /// <summary>
        /// The status of the image. It can be "queued", "active"
        /// or "saving".
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags of the image. It must be a list of strings.
        /// At this time, it is not possible to delete all tags of an image.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// (**Deprecated** - use `updated_at` instead)
        /// </summary>
        [Input("updateAt")]
        public Input<string>? UpdateAt { get; set; }

        /// <summary>
        /// The date the image was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// If false, the checksum will not be verified
        /// once the image is finished uploading. Conflicts with `web_download`.
        /// Defaults to true when not using `web_download`.
        /// </summary>
        [Input("verifyChecksum")]
        public Input<bool>? VerifyChecksum { get; set; }

        /// <summary>
        /// The visibility of the image. Must be one of
        /// "public", "private", "community", or "shared". The ability to set the
        /// visibility depends upon the configuration of the OpenStack cloud.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// If true, the "web-download" import method will
        /// be used to let Openstack download the image directly from the remote source.
        /// Conflicts with `local_file_path`. Defaults to false.
        /// </summary>
        [Input("webDownload")]
        public Input<bool>? WebDownload { get; set; }

        public ImageState()
        {
        }
    }
}
