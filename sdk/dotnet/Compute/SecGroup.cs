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
    /// Manages a V2 security group resource within OpenStack.
    /// 
    /// Please note that managing security groups through the OpenStack Compute API
    /// has been deprecated. Unless you are using an older OpenStack environment, it is
    /// recommended to use the `openstack.networking.SecGroup`
    /// and `openstack.networking.SecGroupRule`
    /// resources instead, which uses the OpenStack Networking API.
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
    ///     var secgroup1 = new OpenStack.Compute.SecGroup("secgroup1", new()
    ///     {
    ///         Description = "my security group",
    ///         Rules = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.SecGroupRuleArgs
    ///             {
    ///                 Cidr = "0.0.0.0/0",
    ///                 FromPort = 22,
    ///                 IpProtocol = "tcp",
    ///                 ToPort = 22,
    ///             },
    ///             new OpenStack.Compute.Inputs.SecGroupRuleArgs
    ///             {
    ///                 Cidr = "0.0.0.0/0",
    ///                 FromPort = 80,
    ///                 IpProtocol = "tcp",
    ///                 ToPort = 80,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Notes
    /// 
    /// ### ICMP Rules
    /// 
    /// When using ICMP as the `ip_protocol`, the `from_port` sets the ICMP _type_ and the `to_port` sets the ICMP _code_. To allow all ICMP types, set each value to `-1`, like so:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    /// });
    /// ```
    /// 
    /// A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages).
    /// 
    /// ### Referencing Security Groups
    /// 
    /// When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test_server = new OpenStack.Compute.Instance("test-server", new()
    ///     {
    ///         ImageId = "ad091b52-742f-469e-8f3c-fd81cadf0743",
    ///         FlavorId = "3",
    ///         KeyPair = "my_key_pair_name",
    ///         SecurityGroups = new[]
    ///         {
    ///             openstack_compute_secgroup_v2.Secgroup_1.Name,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Security Groups can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:compute/secGroup:SecGroup my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:compute/secGroup:SecGroup")]
    public partial class SecGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the security group. Changing this
        /// updates the `description` of an existing security group.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A unique name for the security group. Changing this
        /// updates the `name` of an existing security group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// A Compute client is needed to create a security group. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security group.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A rule describing how the security group operates. The
        /// rule object structure is documented below. Changing this updates the
        /// security group rules. As shown in the example above, multiple rule blocks
        /// may be used.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.SecGroupRule>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a SecGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecGroup(string name, SecGroupArgs args, CustomResourceOptions? options = null)
            : base("openstack:compute/secGroup:SecGroup", name, args ?? new SecGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecGroup(string name, Input<string> id, SecGroupState? state = null, CustomResourceOptions? options = null)
            : base("openstack:compute/secGroup:SecGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecGroup Get(string name, Input<string> id, SecGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SecGroup(name, id, state, options);
        }
    }

    public sealed class SecGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the security group. Changing this
        /// updates the `description` of an existing security group.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// A unique name for the security group. Changing this
        /// updates the `name` of an existing security group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// A Compute client is needed to create a security group. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("rules")]
        private InputList<Inputs.SecGroupRuleArgs>? _rules;

        /// <summary>
        /// A rule describing how the security group operates. The
        /// rule object structure is documented below. Changing this updates the
        /// security group rules. As shown in the example above, multiple rule blocks
        /// may be used.
        /// </summary>
        public InputList<Inputs.SecGroupRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecGroupRuleArgs>());
            set => _rules = value;
        }

        public SecGroupArgs()
        {
        }
        public static new SecGroupArgs Empty => new SecGroupArgs();
    }

    public sealed class SecGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the security group. Changing this
        /// updates the `description` of an existing security group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique name for the security group. Changing this
        /// updates the `name` of an existing security group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// A Compute client is needed to create a security group. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("rules")]
        private InputList<Inputs.SecGroupRuleGetArgs>? _rules;

        /// <summary>
        /// A rule describing how the security group operates. The
        /// rule object structure is documented below. Changing this updates the
        /// security group rules. As shown in the example above, multiple rule blocks
        /// may be used.
        /// </summary>
        public InputList<Inputs.SecGroupRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.SecGroupRuleGetArgs>());
            set => _rules = value;
        }

        public SecGroupState()
        {
        }
        public static new SecGroupState Empty => new SecGroupState();
    }
}
