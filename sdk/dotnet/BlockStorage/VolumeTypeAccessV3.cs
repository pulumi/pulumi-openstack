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
    /// Manages a V3 block storage volume type access resource within OpenStack.
    /// 
    /// &gt; **Note:** This usually requires admin privileges.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project1 = new OpenStack.Identity.Project("project1");
    /// 
    ///     var volumeType1 = new OpenStack.BlockStorage.VolumeTypeV3("volumeType1", new()
    ///     {
    ///         IsPublic = false,
    ///     });
    /// 
    ///     var volumeTypeAccess = new OpenStack.BlockStorage.VolumeTypeAccessV3("volumeTypeAccess", new()
    ///     {
    ///         ProjectId = project1.Id,
    ///         VolumeTypeId = volumeType1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Volume types access can be imported using the `volume_type_id/project_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3 volume_type_access 941793f0-0a34-4bc4-b72e-a6326ae58283/ed498e81f0cc448bae0ad4f8f21bf67f
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3")]
    public partial class VolumeTypeAccessV3 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the project to give access to. Changing this
        /// creates a new resource.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// ID of the volume type to give access to. Changing
        /// this creates a new resource.
        /// </summary>
        [Output("volumeTypeId")]
        public Output<string> VolumeTypeId { get; private set; } = null!;


        /// <summary>
        /// Create a VolumeTypeAccessV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VolumeTypeAccessV3(string name, VolumeTypeAccessV3Args args, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3", name, args ?? new VolumeTypeAccessV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private VolumeTypeAccessV3(string name, Input<string> id, VolumeTypeAccessV3State? state = null, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VolumeTypeAccessV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VolumeTypeAccessV3 Get(string name, Input<string> id, VolumeTypeAccessV3State? state = null, CustomResourceOptions? options = null)
        {
            return new VolumeTypeAccessV3(name, id, state, options);
        }
    }

    public sealed class VolumeTypeAccessV3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the project to give access to. Changing this
        /// creates a new resource.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// ID of the volume type to give access to. Changing
        /// this creates a new resource.
        /// </summary>
        [Input("volumeTypeId", required: true)]
        public Input<string> VolumeTypeId { get; set; } = null!;

        public VolumeTypeAccessV3Args()
        {
        }
        public static new VolumeTypeAccessV3Args Empty => new VolumeTypeAccessV3Args();
    }

    public sealed class VolumeTypeAccessV3State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the project to give access to. Changing this
        /// creates a new resource.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// ID of the volume type to give access to. Changing
        /// this creates a new resource.
        /// </summary>
        [Input("volumeTypeId")]
        public Input<string>? VolumeTypeId { get; set; }

        public VolumeTypeAccessV3State()
        {
        }
        public static new VolumeTypeAccessV3State Empty => new VolumeTypeAccessV3State();
    }
}
