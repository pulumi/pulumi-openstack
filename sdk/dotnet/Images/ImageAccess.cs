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
    /// Manages members for the shared OpenStack Glance V2 Image within the source
    /// project, which owns the Image.
    /// 
    /// ## Example Usage
    /// ### Unprivileged user
    /// 
    /// Create a shared image and propose a membership to the
    /// `bed6b6cbb86a4e2d8dc2735c2f1000e4` project ID.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var rancheros = new OpenStack.Images.Image("rancheros", new()
    ///     {
    ///         ContainerFormat = "bare",
    ///         DiskFormat = "qcow2",
    ///         ImageSourceUrl = "https://releases.rancher.com/os/latest/rancheros-openstack.img",
    ///         Properties = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///         Visibility = "shared",
    ///     });
    /// 
    ///     var rancherosMember = new OpenStack.Images.ImageAccess("rancherosMember", new()
    ///     {
    ///         ImageId = rancheros.Id,
    ///         MemberId = "bed6b6cbb86a4e2d8dc2735c2f1000e4",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Privileged user
    /// 
    /// Create a shared image and set a membership to the
    /// `bed6b6cbb86a4e2d8dc2735c2f1000e4` project ID.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var rancheros = new OpenStack.Images.Image("rancheros", new()
    ///     {
    ///         ContainerFormat = "bare",
    ///         DiskFormat = "qcow2",
    ///         ImageSourceUrl = "https://releases.rancher.com/os/latest/rancheros-openstack.img",
    ///         Properties = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///         Visibility = "shared",
    ///     });
    /// 
    ///     var rancherosMember = new OpenStack.Images.ImageAccess("rancherosMember", new()
    ///     {
    ///         ImageId = rancheros.Id,
    ///         MemberId = "bed6b6cbb86a4e2d8dc2735c2f1000e4",
    ///         Status = "accepted",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Image access can be imported using the `image_id` and the `member_id`, separated by a slash, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:images/imageAccess:ImageAccess openstack_images_image_access_v2 89c60255-9bd6-460c-822a-e2b959ede9d2/bed6b6cbb86a4e2d8dc2735c2f1000e4
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:images/imageAccess:ImageAccess")]
    public partial class ImageAccess : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date the image access was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The image ID.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// The member ID, e.g. the target project ID.
        /// </summary>
        [Output("memberId")]
        public Output<string> MemberId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to manage Image members. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new resource.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The member schema.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// The member proposal status. Optional if admin wants to
        /// force the member proposal acceptance. Can either be `accepted`, `rejected` or
        /// `pending`. Defaults to `pending`. Foridden for non-admin users.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The date the image access was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ImageAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageAccess(string name, ImageAccessArgs args, CustomResourceOptions? options = null)
            : base("openstack:images/imageAccess:ImageAccess", name, args ?? new ImageAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageAccess(string name, Input<string> id, ImageAccessState? state = null, CustomResourceOptions? options = null)
            : base("openstack:images/imageAccess:ImageAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ImageAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageAccess Get(string name, Input<string> id, ImageAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new ImageAccess(name, id, state, options);
        }
    }

    public sealed class ImageAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The image ID.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// The member ID, e.g. the target project ID.
        /// </summary>
        [Input("memberId", required: true)]
        public Input<string> MemberId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to manage Image members. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new resource.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The member proposal status. Optional if admin wants to
        /// force the member proposal acceptance. Can either be `accepted`, `rejected` or
        /// `pending`. Defaults to `pending`. Foridden for non-admin users.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ImageAccessArgs()
        {
        }
        public static new ImageAccessArgs Empty => new ImageAccessArgs();
    }

    public sealed class ImageAccessState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date the image access was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The image ID.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The member ID, e.g. the target project ID.
        /// </summary>
        [Input("memberId")]
        public Input<string>? MemberId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to manage Image members. If omitted, the `region`
        /// argument of the provider is used. Changing this creates a new resource.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The member schema.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// The member proposal status. Optional if admin wants to
        /// force the member proposal acceptance. Can either be `accepted`, `rejected` or
        /// `pending`. Defaults to `pending`. Foridden for non-admin users.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date the image access was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ImageAccessState()
        {
        }
        public static new ImageAccessState Empty => new ImageAccessState();
    }
}
