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
    /// Manages memberships status for the shared OpenStack Glance V2 Image within the
    /// destination project, which has a member proposal.
    /// 
    /// ## Example Usage
    /// 
    /// Accept a shared image membershipship proposal within the current project.
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var rancheros = OpenStack.Images.GetImage.Invoke(new()
    ///     {
    ///         Name = "RancherOS",
    ///         Visibility = "shared",
    ///         MemberStatus = "all",
    ///     });
    /// 
    ///     var rancherosMember = new OpenStack.Images.ImageAccessAccept("rancherosMember", new()
    ///     {
    ///         ImageId = rancheros.Apply(getImageResult =&gt; getImageResult.Id),
    ///         Status = "accepted",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Image access acceptance status can be imported using the `image_id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:images/imageAccessAccept:ImageAccessAccept openstack_images_image_access_accept_v2 89c60255-9bd6-460c-822a-e2b959ede9d2
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:images/imageAccessAccept:ImageAccessAccept")]
    public partial class ImageAccessAccept : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date the image membership was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The proposed image ID.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// The member ID, e.g. the target project ID. Optional
        /// for admin accounts. Defaults to the current scope project ID.
        /// </summary>
        [Output("memberId")]
        public Output<string> MemberId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to manage Image memberships. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// membership.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The membership schema.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;

        /// <summary>
        /// The membership proposal status. Can either be
        /// `accepted`, `rejected` or `pending`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The date the image membership was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ImageAccessAccept resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageAccessAccept(string name, ImageAccessAcceptArgs args, CustomResourceOptions? options = null)
            : base("openstack:images/imageAccessAccept:ImageAccessAccept", name, args ?? new ImageAccessAcceptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageAccessAccept(string name, Input<string> id, ImageAccessAcceptState? state = null, CustomResourceOptions? options = null)
            : base("openstack:images/imageAccessAccept:ImageAccessAccept", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ImageAccessAccept resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageAccessAccept Get(string name, Input<string> id, ImageAccessAcceptState? state = null, CustomResourceOptions? options = null)
        {
            return new ImageAccessAccept(name, id, state, options);
        }
    }

    public sealed class ImageAccessAcceptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The proposed image ID.
        /// </summary>
        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        /// <summary>
        /// The member ID, e.g. the target project ID. Optional
        /// for admin accounts. Defaults to the current scope project ID.
        /// </summary>
        [Input("memberId")]
        public Input<string>? MemberId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to manage Image memberships. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// membership.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The membership proposal status. Can either be
        /// `accepted`, `rejected` or `pending`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public ImageAccessAcceptArgs()
        {
        }
        public static new ImageAccessAcceptArgs Empty => new ImageAccessAcceptArgs();
    }

    public sealed class ImageAccessAcceptState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date the image membership was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The proposed image ID.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// The member ID, e.g. the target project ID. Optional
        /// for admin accounts. Defaults to the current scope project ID.
        /// </summary>
        [Input("memberId")]
        public Input<string>? MemberId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Glance client.
        /// A Glance client is needed to manage Image memberships. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// membership.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The membership schema.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        /// <summary>
        /// The membership proposal status. Can either be
        /// `accepted`, `rejected` or `pending`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date the image membership was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ImageAccessAcceptState()
        {
        }
        public static new ImageAccessAcceptState Empty => new ImageAccessAcceptState();
    }
}
