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
    /// Manages a V3 block storage volume type resource within OpenStack.
    /// 
    /// &gt; **Note:** This usually requires admin privileges.
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
    ///     var volumeType1 = new OpenStack.BlockStorage.VolumeTypeV3("volume_type_1", new()
    ///     {
    ///         Name = "volume_type_1",
    ///         Description = "Volume type 1",
    ///         ExtraSpecs = 
    ///         {
    ///             { "capabilities", "gpu" },
    ///             { "volume_backend_name", "ssd" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Volume types can be imported using the `volume_type_id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:blockstorage/volumeTypeV3:VolumeTypeV3 volume_type_1 941793f0-0a34-4bc4-b72e-a6326ae58283
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:blockstorage/volumeTypeV3:VolumeTypeV3")]
    public partial class VolumeTypeV3 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Human-readable description of the port. Changing
        /// this updates the `description` of an existing volume type.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Key/Value pairs of metadata for the volume type.
        /// </summary>
        [Output("extraSpecs")]
        public Output<ImmutableDictionary<string, object>> ExtraSpecs { get; private set; } = null!;

        /// <summary>
        /// Whether the volume type is public. Changing
        /// this updates the `is_public` of an existing volume type.
        /// </summary>
        [Output("isPublic")]
        public Output<bool> IsPublic { get; private set; } = null!;

        /// <summary>
        /// Name of the volume type.  Changing this
        /// updates the `name` of an existing volume type.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a VolumeTypeV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VolumeTypeV3(string name, VolumeTypeV3Args? args = null, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/volumeTypeV3:VolumeTypeV3", name, args ?? new VolumeTypeV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private VolumeTypeV3(string name, Input<string> id, VolumeTypeV3State? state = null, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/volumeTypeV3:VolumeTypeV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VolumeTypeV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VolumeTypeV3 Get(string name, Input<string> id, VolumeTypeV3State? state = null, CustomResourceOptions? options = null)
        {
            return new VolumeTypeV3(name, id, state, options);
        }
    }

    public sealed class VolumeTypeV3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-readable description of the port. Changing
        /// this updates the `description` of an existing volume type.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("extraSpecs")]
        private InputMap<object>? _extraSpecs;

        /// <summary>
        /// Key/Value pairs of metadata for the volume type.
        /// </summary>
        public InputMap<object> ExtraSpecs
        {
            get => _extraSpecs ?? (_extraSpecs = new InputMap<object>());
            set => _extraSpecs = value;
        }

        /// <summary>
        /// Whether the volume type is public. Changing
        /// this updates the `is_public` of an existing volume type.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// Name of the volume type.  Changing this
        /// updates the `name` of an existing volume type.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public VolumeTypeV3Args()
        {
        }
        public static new VolumeTypeV3Args Empty => new VolumeTypeV3Args();
    }

    public sealed class VolumeTypeV3State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-readable description of the port. Changing
        /// this updates the `description` of an existing volume type.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("extraSpecs")]
        private InputMap<object>? _extraSpecs;

        /// <summary>
        /// Key/Value pairs of metadata for the volume type.
        /// </summary>
        public InputMap<object> ExtraSpecs
        {
            get => _extraSpecs ?? (_extraSpecs = new InputMap<object>());
            set => _extraSpecs = value;
        }

        /// <summary>
        /// Whether the volume type is public. Changing
        /// this updates the `is_public` of an existing volume type.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// Name of the volume type.  Changing this
        /// updates the `name` of an existing volume type.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public VolumeTypeV3State()
        {
        }
        public static new VolumeTypeV3State Empty => new VolumeTypeV3State();
    }
}
