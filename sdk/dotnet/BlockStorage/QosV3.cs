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
    /// Manages a V3 block storage Quality-Of-Servirce (qos) resource within OpenStack.
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
    ///     var qos = new OpenStack.BlockStorage.QosV3("qos", new()
    ///     {
    ///         Consumer = "back-end",
    ///         Specs = 
    ///         {
    ///             { "read_iops_sec", "40000" },
    ///             { "write_iops_sec", "40000" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Qos can be imported using the `qos_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:blockstorage/qosV3:QosV3 qos 941793f0-0a34-4bc4-b72e-a6326ae58283
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:blockstorage/qosV3:QosV3")]
    public partial class QosV3 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The consumer of qos. Can be one of `front-end`,
        /// `back-end` or `both`. Changing this updates the `consumer` of an
        /// existing qos.
        /// </summary>
        [Output("consumer")]
        public Output<string?> Consumer { get; private set; } = null!;

        /// <summary>
        /// Name of the qos.  Changing this creates a new qos.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the qos. If omitted,
        /// the `region` argument of the provider is used. Changing this creates
        /// a new qos.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Key/Value pairs of specs for the qos.
        /// </summary>
        [Output("specs")]
        public Output<ImmutableDictionary<string, object>?> Specs { get; private set; } = null!;


        /// <summary>
        /// Create a QosV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QosV3(string name, QosV3Args? args = null, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/qosV3:QosV3", name, args ?? new QosV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private QosV3(string name, Input<string> id, QosV3State? state = null, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/qosV3:QosV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QosV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QosV3 Get(string name, Input<string> id, QosV3State? state = null, CustomResourceOptions? options = null)
        {
            return new QosV3(name, id, state, options);
        }
    }

    public sealed class QosV3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The consumer of qos. Can be one of `front-end`,
        /// `back-end` or `both`. Changing this updates the `consumer` of an
        /// existing qos.
        /// </summary>
        [Input("consumer")]
        public Input<string>? Consumer { get; set; }

        /// <summary>
        /// Name of the qos.  Changing this creates a new qos.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to create the qos. If omitted,
        /// the `region` argument of the provider is used. Changing this creates
        /// a new qos.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("specs")]
        private InputMap<object>? _specs;

        /// <summary>
        /// Key/Value pairs of specs for the qos.
        /// </summary>
        public InputMap<object> Specs
        {
            get => _specs ?? (_specs = new InputMap<object>());
            set => _specs = value;
        }

        public QosV3Args()
        {
        }
        public static new QosV3Args Empty => new QosV3Args();
    }

    public sealed class QosV3State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The consumer of qos. Can be one of `front-end`,
        /// `back-end` or `both`. Changing this updates the `consumer` of an
        /// existing qos.
        /// </summary>
        [Input("consumer")]
        public Input<string>? Consumer { get; set; }

        /// <summary>
        /// Name of the qos.  Changing this creates a new qos.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to create the qos. If omitted,
        /// the `region` argument of the provider is used. Changing this creates
        /// a new qos.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("specs")]
        private InputMap<object>? _specs;

        /// <summary>
        /// Key/Value pairs of specs for the qos.
        /// </summary>
        public InputMap<object> Specs
        {
            get => _specs ?? (_specs = new InputMap<object>());
            set => _specs = value;
        }

        public QosV3State()
        {
        }
        public static new QosV3State Empty => new QosV3State();
    }
}
