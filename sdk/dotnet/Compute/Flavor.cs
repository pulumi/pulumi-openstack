// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    /// <summary>
    /// Manages a V2 flavor resource within OpenStack.
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
    ///     var test_flavor = new OpenStack.Compute.Flavor("test-flavor", new()
    ///     {
    ///         Disk = 20,
    ///         ExtraSpecs = 
    ///         {
    ///             { "hw:cpu_policy", "CPU-POLICY" },
    ///             { "hw:cpu_thread_policy", "CPU-THREAD-POLICY" },
    ///         },
    ///         Ram = 8096,
    ///         Vcpus = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Flavors can be imported using the `ID`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:compute/flavor:Flavor my-flavor 4142e64b-1b35-44a0-9b1e-5affc7af1106
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:compute/flavor:Flavor")]
    public partial class Flavor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the flavor. Changing this
        /// updates the description of the flavor. Requires microversion &gt;= 2.55.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The amount of disk space in GiB to use for the root
        /// (/) partition. Changing this creates a new flavor.
        /// </summary>
        [Output("disk")]
        public Output<int> Disk { get; private set; } = null!;

        /// <summary>
        /// The amount of ephemeral in GiB. If unspecified,
        /// the default is 0. Changing this creates a new flavor.
        /// </summary>
        [Output("ephemeral")]
        public Output<int?> Ephemeral { get; private set; } = null!;

        /// <summary>
        /// Key/Value pairs of metadata for the flavor.
        /// </summary>
        [Output("extraSpecs")]
        public Output<ImmutableDictionary<string, object>> ExtraSpecs { get; private set; } = null!;

        /// <summary>
        /// Unique ID (integer or UUID) of flavor to create. Changing
        /// this creates a new flavor.
        /// </summary>
        [Output("flavorId")]
        public Output<string> FlavorId { get; private set; } = null!;

        /// <summary>
        /// Whether the flavor is public. Changing this creates
        /// a new flavor.
        /// </summary>
        [Output("isPublic")]
        public Output<bool?> IsPublic { get; private set; } = null!;

        /// <summary>
        /// A unique name for the flavor. Changing this creates a new
        /// flavor.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The amount of RAM to use, in megabytes. Changing this
        /// creates a new flavor.
        /// </summary>
        [Output("ram")]
        public Output<int> Ram { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Flavors are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new flavor.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// RX/TX bandwith factor. The default is 1. Changing
        /// this creates a new flavor.
        /// </summary>
        [Output("rxTxFactor")]
        public Output<double?> RxTxFactor { get; private set; } = null!;

        /// <summary>
        /// The amount of disk space in megabytes to use. If
        /// unspecified, the default is 0. Changing this creates a new flavor.
        /// </summary>
        [Output("swap")]
        public Output<int?> Swap { get; private set; } = null!;

        /// <summary>
        /// The number of virtual CPUs to use. Changing this creates
        /// a new flavor.
        /// </summary>
        [Output("vcpus")]
        public Output<int> Vcpus { get; private set; } = null!;


        /// <summary>
        /// Create a Flavor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Flavor(string name, FlavorArgs args, CustomResourceOptions? options = null)
            : base("openstack:compute/flavor:Flavor", name, args ?? new FlavorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Flavor(string name, Input<string> id, FlavorState? state = null, CustomResourceOptions? options = null)
            : base("openstack:compute/flavor:Flavor", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Flavor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Flavor Get(string name, Input<string> id, FlavorState? state = null, CustomResourceOptions? options = null)
        {
            return new Flavor(name, id, state, options);
        }
    }

    public sealed class FlavorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the flavor. Changing this
        /// updates the description of the flavor. Requires microversion &gt;= 2.55.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The amount of disk space in GiB to use for the root
        /// (/) partition. Changing this creates a new flavor.
        /// </summary>
        [Input("disk", required: true)]
        public Input<int> Disk { get; set; } = null!;

        /// <summary>
        /// The amount of ephemeral in GiB. If unspecified,
        /// the default is 0. Changing this creates a new flavor.
        /// </summary>
        [Input("ephemeral")]
        public Input<int>? Ephemeral { get; set; }

        [Input("extraSpecs")]
        private InputMap<object>? _extraSpecs;

        /// <summary>
        /// Key/Value pairs of metadata for the flavor.
        /// </summary>
        public InputMap<object> ExtraSpecs
        {
            get => _extraSpecs ?? (_extraSpecs = new InputMap<object>());
            set => _extraSpecs = value;
        }

        /// <summary>
        /// Unique ID (integer or UUID) of flavor to create. Changing
        /// this creates a new flavor.
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// Whether the flavor is public. Changing this creates
        /// a new flavor.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// A unique name for the flavor. Changing this creates a new
        /// flavor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The amount of RAM to use, in megabytes. Changing this
        /// creates a new flavor.
        /// </summary>
        [Input("ram", required: true)]
        public Input<int> Ram { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Flavors are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new flavor.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// RX/TX bandwith factor. The default is 1. Changing
        /// this creates a new flavor.
        /// </summary>
        [Input("rxTxFactor")]
        public Input<double>? RxTxFactor { get; set; }

        /// <summary>
        /// The amount of disk space in megabytes to use. If
        /// unspecified, the default is 0. Changing this creates a new flavor.
        /// </summary>
        [Input("swap")]
        public Input<int>? Swap { get; set; }

        /// <summary>
        /// The number of virtual CPUs to use. Changing this creates
        /// a new flavor.
        /// </summary>
        [Input("vcpus", required: true)]
        public Input<int> Vcpus { get; set; } = null!;

        public FlavorArgs()
        {
        }
        public static new FlavorArgs Empty => new FlavorArgs();
    }

    public sealed class FlavorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the flavor. Changing this
        /// updates the description of the flavor. Requires microversion &gt;= 2.55.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The amount of disk space in GiB to use for the root
        /// (/) partition. Changing this creates a new flavor.
        /// </summary>
        [Input("disk")]
        public Input<int>? Disk { get; set; }

        /// <summary>
        /// The amount of ephemeral in GiB. If unspecified,
        /// the default is 0. Changing this creates a new flavor.
        /// </summary>
        [Input("ephemeral")]
        public Input<int>? Ephemeral { get; set; }

        [Input("extraSpecs")]
        private InputMap<object>? _extraSpecs;

        /// <summary>
        /// Key/Value pairs of metadata for the flavor.
        /// </summary>
        public InputMap<object> ExtraSpecs
        {
            get => _extraSpecs ?? (_extraSpecs = new InputMap<object>());
            set => _extraSpecs = value;
        }

        /// <summary>
        /// Unique ID (integer or UUID) of flavor to create. Changing
        /// this creates a new flavor.
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// Whether the flavor is public. Changing this creates
        /// a new flavor.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// A unique name for the flavor. Changing this creates a new
        /// flavor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The amount of RAM to use, in megabytes. Changing this
        /// creates a new flavor.
        /// </summary>
        [Input("ram")]
        public Input<int>? Ram { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Flavors are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new flavor.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// RX/TX bandwith factor. The default is 1. Changing
        /// this creates a new flavor.
        /// </summary>
        [Input("rxTxFactor")]
        public Input<double>? RxTxFactor { get; set; }

        /// <summary>
        /// The amount of disk space in megabytes to use. If
        /// unspecified, the default is 0. Changing this creates a new flavor.
        /// </summary>
        [Input("swap")]
        public Input<int>? Swap { get; set; }

        /// <summary>
        /// The number of virtual CPUs to use. Changing this creates
        /// a new flavor.
        /// </summary>
        [Input("vcpus")]
        public Input<int>? Vcpus { get; set; }

        public FlavorState()
        {
        }
        public static new FlavorState Empty => new FlavorState();
    }
}
