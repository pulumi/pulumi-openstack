// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Images
{
    public static class GetImageIds
    {
        /// <summary>
        /// Use this data source to get a list of Openstack Image IDs matching the
        /// specified criteria.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var images = OpenStack.Images.GetImageIds.Invoke(new()
        ///     {
        ///         NameRegex = "^Ubuntu 16\\.04.*-amd64",
        ///         Sort = "updated_at",
        ///         Properties = 
        ///         {
        ///             { "key", "value" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetImageIdsResult> InvokeAsync(GetImageIdsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImageIdsResult>("openstack:images/getImageIds:getImageIds", args ?? new GetImageIdsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a list of Openstack Image IDs matching the
        /// specified criteria.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var images = OpenStack.Images.GetImageIds.Invoke(new()
        ///     {
        ///         NameRegex = "^Ubuntu 16\\.04.*-amd64",
        ///         Sort = "updated_at",
        ///         Properties = 
        ///         {
        ///             { "key", "value" },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetImageIdsResult> Invoke(GetImageIdsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImageIdsResult>("openstack:images/getImageIds:getImageIds", args ?? new GetImageIdsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImageIdsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The container format of the image.
        /// </summary>
        [Input("containerFormat")]
        public string? ContainerFormat { get; set; }

        /// <summary>
        /// The disk format of the image.
        /// </summary>
        [Input("diskFormat")]
        public string? DiskFormat { get; set; }

        /// <summary>
        /// Whether or not the image is hidden from public list.
        /// </summary>
        [Input("hidden")]
        public bool? Hidden { get; set; }

        /// <summary>
        /// The status of the image. Must be one of
        /// "accepted", "pending", "rejected", or "all".
        /// </summary>
        [Input("memberStatus")]
        public string? MemberStatus { get; set; }

        /// <summary>
        /// The name of the image. Cannot be used simultaneously with
        /// `name_regex`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The regular expressian of the name of the image.
        /// Cannot be used simultaneously with `name`. Unlike filtering by `name` the
        /// `name_regex` filtering does by client on the result of OpenStack search
        /// query.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// The owner (UUID) of the image.
        /// </summary>
        [Input("owner")]
        public string? Owner { get; set; }

        [Input("properties")]
        private Dictionary<string, string>? _properties;

        /// <summary>
        /// a map of key/value pairs to match an image with.
        /// All specified properties must be matched. Unlike other options filtering by
        /// `properties` does by client on the result of OpenStack search query.
        /// </summary>
        public Dictionary<string, string> Properties
        {
            get => _properties ?? (_properties = new Dictionary<string, string>());
            set => _properties = value;
        }

        /// <summary>
        /// The region in which to obtain the V2 Glance client. A
        /// Glance client is needed to create an Image that can be used with a compute
        /// instance. If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The maximum size (in bytes) of the image to return.
        /// </summary>
        [Input("sizeMax")]
        public int? SizeMax { get; set; }

        /// <summary>
        /// The minimum size (in bytes) of the image to return.
        /// </summary>
        [Input("sizeMin")]
        public int? SizeMin { get; set; }

        /// <summary>
        /// Sorts the response by one or more attribute and sort
        /// direction combinations. You can also set multiple sort keys and directions.
        /// Default direction is `desc`. Use the comma (,) character to separate multiple
        /// values. For example expression `sort = "name:asc,status"` sorts ascending by
        /// name and descending by status.
        /// </summary>
        [Input("sort")]
        public string? Sort { get; set; }

        /// <summary>
        /// Search for images with a specific tag.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// A list of tags required to be set on the image (all
        /// specified tags must be in the images tag list for it to be matched).
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The visibility of the image. Must be one of
        /// "public", "private", "community", or "shared". Defaults to "private".
        /// </summary>
        [Input("visibility")]
        public string? Visibility { get; set; }

        public GetImageIdsArgs()
        {
        }
        public static new GetImageIdsArgs Empty => new GetImageIdsArgs();
    }

    public sealed class GetImageIdsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The container format of the image.
        /// </summary>
        [Input("containerFormat")]
        public Input<string>? ContainerFormat { get; set; }

        /// <summary>
        /// The disk format of the image.
        /// </summary>
        [Input("diskFormat")]
        public Input<string>? DiskFormat { get; set; }

        /// <summary>
        /// Whether or not the image is hidden from public list.
        /// </summary>
        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        /// <summary>
        /// The status of the image. Must be one of
        /// "accepted", "pending", "rejected", or "all".
        /// </summary>
        [Input("memberStatus")]
        public Input<string>? MemberStatus { get; set; }

        /// <summary>
        /// The name of the image. Cannot be used simultaneously with
        /// `name_regex`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The regular expressian of the name of the image.
        /// Cannot be used simultaneously with `name`. Unlike filtering by `name` the
        /// `name_regex` filtering does by client on the result of OpenStack search
        /// query.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// The owner (UUID) of the image.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// a map of key/value pairs to match an image with.
        /// All specified properties must be matched. Unlike other options filtering by
        /// `properties` does by client on the result of OpenStack search query.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The region in which to obtain the V2 Glance client. A
        /// Glance client is needed to create an Image that can be used with a compute
        /// instance. If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The maximum size (in bytes) of the image to return.
        /// </summary>
        [Input("sizeMax")]
        public Input<int>? SizeMax { get; set; }

        /// <summary>
        /// The minimum size (in bytes) of the image to return.
        /// </summary>
        [Input("sizeMin")]
        public Input<int>? SizeMin { get; set; }

        /// <summary>
        /// Sorts the response by one or more attribute and sort
        /// direction combinations. You can also set multiple sort keys and directions.
        /// Default direction is `desc`. Use the comma (,) character to separate multiple
        /// values. For example expression `sort = "name:asc,status"` sorts ascending by
        /// name and descending by status.
        /// </summary>
        [Input("sort")]
        public Input<string>? Sort { get; set; }

        /// <summary>
        /// Search for images with a specific tag.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags required to be set on the image (all
        /// specified tags must be in the images tag list for it to be matched).
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The visibility of the image. Must be one of
        /// "public", "private", "community", or "shared". Defaults to "private".
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public GetImageIdsInvokeArgs()
        {
        }
        public static new GetImageIdsInvokeArgs Empty => new GetImageIdsInvokeArgs();
    }


    [OutputType]
    public sealed class GetImageIdsResult
    {
        public readonly string? ContainerFormat;
        public readonly string? DiskFormat;
        public readonly bool? Hidden;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? MemberStatus;
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly string? Owner;
        public readonly ImmutableDictionary<string, string>? Properties;
        public readonly string Region;
        public readonly int? SizeMax;
        public readonly int? SizeMin;
        public readonly string? Sort;
        public readonly string? Tag;
        public readonly ImmutableArray<string> Tags;
        public readonly string? Visibility;

        [OutputConstructor]
        private GetImageIdsResult(
            string? containerFormat,

            string? diskFormat,

            bool? hidden,

            string id,

            ImmutableArray<string> ids,

            string? memberStatus,

            string? name,

            string? nameRegex,

            string? owner,

            ImmutableDictionary<string, string>? properties,

            string region,

            int? sizeMax,

            int? sizeMin,

            string? sort,

            string? tag,

            ImmutableArray<string> tags,

            string? visibility)
        {
            ContainerFormat = containerFormat;
            DiskFormat = diskFormat;
            Hidden = hidden;
            Id = id;
            Ids = ids;
            MemberStatus = memberStatus;
            Name = name;
            NameRegex = nameRegex;
            Owner = owner;
            Properties = properties;
            Region = region;
            SizeMax = sizeMax;
            SizeMin = sizeMin;
            Sort = sort;
            Tag = tag;
            Tags = tags;
            Visibility = visibility;
        }
    }
}
