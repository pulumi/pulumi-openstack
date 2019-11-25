// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Networking
{
    /// <summary>
    /// Manages a V2 Neutron QoS policy resource within OpenStack.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_qos_policy_v2.html.markdown.
    /// </summary>
    public partial class QosPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The collection of tags assigned on the QoS policy, which have been
        /// explicitly and implicitly added.
        /// </summary>
        [Output("allTags")]
        public Output<ImmutableArray<string>> AllTags { get; private set; } = null!;

        /// <summary>
        /// The time at which QoS policy was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The human-readable description for the QoS policy.
        /// Changing this updates the description of the existing QoS policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the QoS policy is default
        /// QoS policy or not. Changing this updates the default status of the existing
        /// QoS policy.
        /// </summary>
        [Output("isDefault")]
        public Output<bool?> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The name of the QoS policy. Changing this updates the name of
        /// the existing QoS policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The owner of the QoS policy. Required if admin wants to
        /// create a QoS policy for another project. Changing this creates a new QoS policy.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron Qos policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// QoS policy.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The revision number of the QoS policy.
        /// </summary>
        [Output("revisionNumber")]
        public Output<int> RevisionNumber { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this QoS policy is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// QoS policy.
        /// </summary>
        [Output("shared")]
        public Output<bool?> Shared { get; private set; } = null!;

        /// <summary>
        /// A set of string tags for the QoS policy.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The time at which QoS policy was created.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a QosPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QosPolicy(string name, QosPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:networking/qosPolicy:QosPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private QosPolicy(string name, Input<string> id, QosPolicyState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/qosPolicy:QosPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QosPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QosPolicy Get(string name, Input<string> id, QosPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new QosPolicy(name, id, state, options);
        }
    }

    public sealed class QosPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description for the QoS policy.
        /// Changing this updates the description of the existing QoS policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether the QoS policy is default
        /// QoS policy or not. Changing this updates the default status of the existing
        /// QoS policy.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of the QoS policy. Changing this updates the name of
        /// the existing QoS policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the QoS policy. Required if admin wants to
        /// create a QoS policy for another project. Changing this creates a new QoS policy.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron Qos policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// QoS policy.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates whether this QoS policy is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// QoS policy.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the QoS policy.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        public QosPolicyArgs()
        {
        }
    }

    public sealed class QosPolicyState : Pulumi.ResourceArgs
    {
        [Input("allTags")]
        private InputList<string>? _allTags;

        /// <summary>
        /// The collection of tags assigned on the QoS policy, which have been
        /// explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllTags
        {
            get => _allTags ?? (_allTags = new InputList<string>());
            set => _allTags = value;
        }

        /// <summary>
        /// The time at which QoS policy was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The human-readable description for the QoS policy.
        /// Changing this updates the description of the existing QoS policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether the QoS policy is default
        /// QoS policy or not. Changing this updates the default status of the existing
        /// QoS policy.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The name of the QoS policy. Changing this updates the name of
        /// the existing QoS policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the QoS policy. Required if admin wants to
        /// create a QoS policy for another project. Changing this creates a new QoS policy.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron Qos policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// QoS policy.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The revision number of the QoS policy.
        /// </summary>
        [Input("revisionNumber")]
        public Input<int>? RevisionNumber { get; set; }

        /// <summary>
        /// Indicates whether this QoS policy is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// QoS policy.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the QoS policy.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The time at which QoS policy was created.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        public QosPolicyState()
        {
        }
    }
}
