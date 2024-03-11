// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Orchestration
{
    /// <summary>
    /// Manages a V1 stack resource within OpenStack.
    /// 
    /// ## Example Usage
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
    ///     var stack1 = new OpenStack.Orchestration.StackV1("stack1", new()
    ///     {
    ///         DisableRollback = true,
    ///         EnvironmentOpts = 
    ///         {
    ///             { "Bin", @"
    /// 
    /// " },
    ///         },
    ///         Parameters = 
    ///         {
    ///             { "length", 4 },
    ///         },
    ///         TemplateOpts = 
    ///         {
    ///             { "Bin", @"heat_template_version: 2013-05-23
    /// parameters:
    ///   length:
    ///     type: number
    /// resources:
    ///   test_res:
    ///     type: OS::Heat::TestResource
    ///   random:
    ///     type: OS::Heat::RandomString
    ///     properties:
    ///       length: {get_param: length}
    /// 
    /// " },
    ///         },
    ///         Timeout = 30,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// stacks can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:orchestration/stackV1:StackV1 stack_1 ea257959-eeb1-4c10-8d33-26f0409a755d
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:orchestration/stackV1:StackV1")]
    public partial class StackV1 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of stack outputs.
        /// </summary>
        [Output("StackOutputs")]
        public Output<ImmutableArray<Outputs.StackV1StackOutput>> StackOutputs { get; private set; } = null!;

        /// <summary>
        /// List of stack capabilities for stack.
        /// </summary>
        [Output("capabilities")]
        public Output<ImmutableArray<string>> Capabilities { get; private set; } = null!;

        /// <summary>
        /// The date and time when the resource was created. The date
        /// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
        /// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
        /// is the time zone as an offset from UTC.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The description of the stack resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Enables or disables deletion of all stack
        /// resources when a stack creation fails. Default is true, meaning all
        /// resources are not deleted when stack creation fails.
        /// </summary>
        [Output("disableRollback")]
        public Output<bool> DisableRollback { get; private set; } = null!;

        /// <summary>
        /// Environment key/value pairs to associate with
        /// the stack which contains details for the environment of the stack.
        /// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
        /// Environment Opts.
        /// </summary>
        [Output("environmentOpts")]
        public Output<ImmutableDictionary<string, object>?> EnvironmentOpts { get; private set; } = null!;

        /// <summary>
        /// A unique name for the stack. It must start with an
        /// alphabetic character. Changing this updates the stack's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of notification topics for stack.
        /// </summary>
        [Output("notificationTopics")]
        public Output<ImmutableArray<string>> NotificationTopics { get; private set; } = null!;

        /// <summary>
        /// User-defined key/value pairs as parameters to pass
        /// to the template. Changing this updates the existing stack parameters.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, object>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the stack. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new stack.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The status of the stack.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The reason for the current status of the stack.
        /// </summary>
        [Output("statusReason")]
        public Output<string> StatusReason { get; private set; } = null!;

        /// <summary>
        /// A list of tags to assosciate with the Stack
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The description of the stack template.
        /// </summary>
        [Output("templateDescription")]
        public Output<string> TemplateDescription { get; private set; } = null!;

        /// <summary>
        /// Template key/value pairs to associate with the
        /// stack which contains either the template file or url.
        /// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
        /// Template Opts.
        /// </summary>
        [Output("templateOpts")]
        public Output<ImmutableDictionary<string, object>> TemplateOpts { get; private set; } = null!;

        /// <summary>
        /// The timeout for stack action in minutes.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// The date and time when the resource was updated. The date
        /// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
        /// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
        /// is the time zone as an offset from UTC.
        /// </summary>
        [Output("updatedTime")]
        public Output<string> UpdatedTime { get; private set; } = null!;


        /// <summary>
        /// Create a StackV1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackV1(string name, StackV1Args args, CustomResourceOptions? options = null)
            : base("openstack:orchestration/stackV1:StackV1", name, args ?? new StackV1Args(), MakeResourceOptions(options, ""))
        {
        }

        private StackV1(string name, Input<string> id, StackV1State? state = null, CustomResourceOptions? options = null)
            : base("openstack:orchestration/stackV1:StackV1", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StackV1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StackV1 Get(string name, Input<string> id, StackV1State? state = null, CustomResourceOptions? options = null)
        {
            return new StackV1(name, id, state, options);
        }
    }

    public sealed class StackV1Args : global::Pulumi.ResourceArgs
    {
        [Input("StackOutputs")]
        private InputList<Inputs.StackV1StackOutputArgs>? _StackOutputs;

        /// <summary>
        /// A list of stack outputs.
        /// </summary>
        public InputList<Inputs.StackV1StackOutputArgs> StackOutputs
        {
            get => _StackOutputs ?? (_StackOutputs = new InputList<Inputs.StackV1StackOutputArgs>());
            set => _StackOutputs = value;
        }

        [Input("capabilities")]
        private InputList<string>? _capabilities;

        /// <summary>
        /// List of stack capabilities for stack.
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        /// <summary>
        /// The date and time when the resource was created. The date
        /// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
        /// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
        /// is the time zone as an offset from UTC.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The description of the stack resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enables or disables deletion of all stack
        /// resources when a stack creation fails. Default is true, meaning all
        /// resources are not deleted when stack creation fails.
        /// </summary>
        [Input("disableRollback")]
        public Input<bool>? DisableRollback { get; set; }

        [Input("environmentOpts")]
        private InputMap<object>? _environmentOpts;

        /// <summary>
        /// Environment key/value pairs to associate with
        /// the stack which contains details for the environment of the stack.
        /// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
        /// Environment Opts.
        /// </summary>
        public InputMap<object> EnvironmentOpts
        {
            get => _environmentOpts ?? (_environmentOpts = new InputMap<object>());
            set => _environmentOpts = value;
        }

        /// <summary>
        /// A unique name for the stack. It must start with an
        /// alphabetic character. Changing this updates the stack's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationTopics")]
        private InputList<string>? _notificationTopics;

        /// <summary>
        /// List of notification topics for stack.
        /// </summary>
        public InputList<string> NotificationTopics
        {
            get => _notificationTopics ?? (_notificationTopics = new InputList<string>());
            set => _notificationTopics = value;
        }

        [Input("parameters")]
        private InputMap<object>? _parameters;

        /// <summary>
        /// User-defined key/value pairs as parameters to pass
        /// to the template. Changing this updates the existing stack parameters.
        /// </summary>
        public InputMap<object> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<object>());
            set => _parameters = value;
        }

        /// <summary>
        /// The region in which to create the stack. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new stack.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The status of the stack.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The reason for the current status of the stack.
        /// </summary>
        [Input("statusReason")]
        public Input<string>? StatusReason { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to assosciate with the Stack
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The description of the stack template.
        /// </summary>
        [Input("templateDescription")]
        public Input<string>? TemplateDescription { get; set; }

        [Input("templateOpts", required: true)]
        private InputMap<object>? _templateOpts;

        /// <summary>
        /// Template key/value pairs to associate with the
        /// stack which contains either the template file or url.
        /// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
        /// Template Opts.
        /// </summary>
        public InputMap<object> TemplateOpts
        {
            get => _templateOpts ?? (_templateOpts = new InputMap<object>());
            set => _templateOpts = value;
        }

        /// <summary>
        /// The timeout for stack action in minutes.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The date and time when the resource was updated. The date
        /// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
        /// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
        /// is the time zone as an offset from UTC.
        /// </summary>
        [Input("updatedTime")]
        public Input<string>? UpdatedTime { get; set; }

        public StackV1Args()
        {
        }
        public static new StackV1Args Empty => new StackV1Args();
    }

    public sealed class StackV1State : global::Pulumi.ResourceArgs
    {
        [Input("StackOutputs")]
        private InputList<Inputs.StackV1StackOutputGetArgs>? _StackOutputs;

        /// <summary>
        /// A list of stack outputs.
        /// </summary>
        public InputList<Inputs.StackV1StackOutputGetArgs> StackOutputs
        {
            get => _StackOutputs ?? (_StackOutputs = new InputList<Inputs.StackV1StackOutputGetArgs>());
            set => _StackOutputs = value;
        }

        [Input("capabilities")]
        private InputList<string>? _capabilities;

        /// <summary>
        /// List of stack capabilities for stack.
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        /// <summary>
        /// The date and time when the resource was created. The date
        /// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
        /// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
        /// is the time zone as an offset from UTC.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The description of the stack resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Enables or disables deletion of all stack
        /// resources when a stack creation fails. Default is true, meaning all
        /// resources are not deleted when stack creation fails.
        /// </summary>
        [Input("disableRollback")]
        public Input<bool>? DisableRollback { get; set; }

        [Input("environmentOpts")]
        private InputMap<object>? _environmentOpts;

        /// <summary>
        /// Environment key/value pairs to associate with
        /// the stack which contains details for the environment of the stack.
        /// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
        /// Environment Opts.
        /// </summary>
        public InputMap<object> EnvironmentOpts
        {
            get => _environmentOpts ?? (_environmentOpts = new InputMap<object>());
            set => _environmentOpts = value;
        }

        /// <summary>
        /// A unique name for the stack. It must start with an
        /// alphabetic character. Changing this updates the stack's name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationTopics")]
        private InputList<string>? _notificationTopics;

        /// <summary>
        /// List of notification topics for stack.
        /// </summary>
        public InputList<string> NotificationTopics
        {
            get => _notificationTopics ?? (_notificationTopics = new InputList<string>());
            set => _notificationTopics = value;
        }

        [Input("parameters")]
        private InputMap<object>? _parameters;

        /// <summary>
        /// User-defined key/value pairs as parameters to pass
        /// to the template. Changing this updates the existing stack parameters.
        /// </summary>
        public InputMap<object> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<object>());
            set => _parameters = value;
        }

        /// <summary>
        /// The region in which to create the stack. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new stack.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The status of the stack.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The reason for the current status of the stack.
        /// </summary>
        [Input("statusReason")]
        public Input<string>? StatusReason { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to assosciate with the Stack
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The description of the stack template.
        /// </summary>
        [Input("templateDescription")]
        public Input<string>? TemplateDescription { get; set; }

        [Input("templateOpts")]
        private InputMap<object>? _templateOpts;

        /// <summary>
        /// Template key/value pairs to associate with the
        /// stack which contains either the template file or url.
        /// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
        /// Template Opts.
        /// </summary>
        public InputMap<object> TemplateOpts
        {
            get => _templateOpts ?? (_templateOpts = new InputMap<object>());
            set => _templateOpts = value;
        }

        /// <summary>
        /// The timeout for stack action in minutes.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The date and time when the resource was updated. The date
        /// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
        /// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
        /// is the time zone as an offset from UTC.
        /// </summary>
        [Input("updatedTime")]
        public Input<string>? UpdatedTime { get; set; }

        public StackV1State()
        {
        }
        public static new StackV1State Empty => new StackV1State();
    }
}
