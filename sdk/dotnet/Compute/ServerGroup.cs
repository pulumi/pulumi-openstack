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
    /// Manages a V2 Server Group resource within OpenStack.
    /// 
    /// ## Example Usage
    /// 
    /// ### Compute service API version 2.63 or below:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test_sg = new OpenStack.Compute.ServerGroup("test-sg", new()
    ///     {
    ///         Name = "my-sg",
    ///         Policies = "anti-affinity",
    ///     });
    /// 
    ///     var test_instance = new OpenStack.Compute.Instance("test-instance", new()
    ///     {
    ///         Name = "my-instance",
    ///         ImageId = "ad091b52-742f-469e-8f3c-fd81cadf0743",
    ///         FlavorId = "3",
    ///         SchedulerHints = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.InstanceSchedulerHintArgs
    ///             {
    ///                 Group = test_sg.Id,
    ///             },
    ///         },
    ///         Networks = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.InstanceNetworkArgs
    ///             {
    ///                 Name = "my_network",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Compute service API version 2.64 or above:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test_sg = new OpenStack.Compute.ServerGroup("test-sg", new()
    ///     {
    ///         Name = "my-sg",
    ///         Policies = "anti-affinity",
    ///         Rules = new OpenStack.Compute.Inputs.ServerGroupRulesArgs
    ///         {
    ///             MaxServerPerHost = 3,
    ///         },
    ///     });
    /// 
    ///     var test_instance = new OpenStack.Compute.Instance("test-instance", new()
    ///     {
    ///         Name = "my-instance",
    ///         ImageId = "ad091b52-742f-469e-8f3c-fd81cadf0743",
    ///         FlavorId = "3",
    ///         SchedulerHints = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.InstanceSchedulerHintArgs
    ///             {
    ///                 Group = test_sg.Id,
    ///             },
    ///         },
    ///         Networks = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.InstanceNetworkArgs
    ///             {
    ///                 Name = "my_network",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Policies
    /// 
    /// * `affinity` - All instances/servers launched in this group will be hosted on
    ///   the same compute node.
    /// 
    /// * `anti-affinity` - All instances/servers launched in this group will be
    ///   hosted on different compute nodes.
    /// 
    /// * `soft-affinity` - All instances/servers launched in this group will be hosted
    ///   on the same compute node if possible, but if not possible they
    ///   still will be scheduled instead of failure. To use this policy your
    ///   OpenStack environment should support Compute service API 2.15 or above.
    /// 
    /// * `soft-anti-affinity` - All instances/servers launched in this group will be
    ///   hosted on different compute nodes if possible, but if not possible they
    ///   still will be scheduled instead of failure. To use this policy your
    ///   OpenStack environment should support Compute service API 2.15 or above.
    /// 
    /// ## Import
    /// 
    /// Server Groups can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:compute/serverGroup:ServerGroup test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:compute/serverGroup:ServerGroup")]
    public partial class ServerGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The instances that are part of this server group.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// A unique name for the server group. Changing this creates
        /// a new server group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of exactly one policy name to associate with
        /// the server group. See the Policies section for more information. Changing this
        /// creates a new server group.
        /// </summary>
        [Output("policies")]
        public Output<string?> Policies { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used. Changing
        /// this creates a new server group.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The rules which are applied to specified `policy`. Currently,
        /// only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        /// </summary>
        [Output("rules")]
        public Output<Outputs.ServerGroupRules> Rules { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, string>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a ServerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerGroup(string name, ServerGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:compute/serverGroup:ServerGroup", name, args ?? new ServerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerGroup(string name, Input<string> id, ServerGroupState? state = null, CustomResourceOptions? options = null)
            : base("openstack:compute/serverGroup:ServerGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerGroup Get(string name, Input<string> id, ServerGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerGroup(name, id, state, options);
        }
    }

    public sealed class ServerGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique name for the server group. Changing this creates
        /// a new server group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A list of exactly one policy name to associate with
        /// the server group. See the Policies section for more information. Changing this
        /// creates a new server group.
        /// </summary>
        [Input("policies")]
        public Input<string>? Policies { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used. Changing
        /// this creates a new server group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The rules which are applied to specified `policy`. Currently,
        /// only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        /// </summary>
        [Input("rules")]
        public Input<Inputs.ServerGroupRulesArgs>? Rules { get; set; }

        [Input("valueSpecs")]
        private InputMap<string>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<string> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<string>());
            set => _valueSpecs = value;
        }

        public ServerGroupArgs()
        {
        }
        public static new ServerGroupArgs Empty => new ServerGroupArgs();
    }

    public sealed class ServerGroupState : global::Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// The instances that are part of this server group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// A unique name for the server group. Changing this creates
        /// a new server group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A list of exactly one policy name to associate with
        /// the server group. See the Policies section for more information. Changing this
        /// creates a new server group.
        /// </summary>
        [Input("policies")]
        public Input<string>? Policies { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used. Changing
        /// this creates a new server group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The rules which are applied to specified `policy`. Currently,
        /// only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        /// </summary>
        [Input("rules")]
        public Input<Inputs.ServerGroupRulesGetArgs>? Rules { get; set; }

        [Input("valueSpecs")]
        private InputMap<string>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<string> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<string>());
            set => _valueSpecs = value;
        }

        public ServerGroupState()
        {
        }
        public static new ServerGroupState Empty => new ServerGroupState();
    }
}
