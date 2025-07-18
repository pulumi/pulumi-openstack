// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Firewall
{
    /// <summary>
    /// Manages a v2 firewall policy resource within OpenStack.
    /// 
    /// &gt; **Note:** Firewall v2 has no support for OVN currently.
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
    ///     var rule1 = new OpenStack.Firewall.RuleV2("rule_1", new()
    ///     {
    ///         Name = "firewall_rule_1",
    ///         Description = "drop TELNET traffic",
    ///         Action = "deny",
    ///         Protocol = "tcp",
    ///         DestinationPort = "23",
    ///         Enabled = true,
    ///     });
    /// 
    ///     var rule2 = new OpenStack.Firewall.RuleV2("rule_2", new()
    ///     {
    ///         Name = "firewall_rule_2",
    ///         Description = "drop NTP traffic",
    ///         Action = "deny",
    ///         Protocol = "udp",
    ///         DestinationPort = "123",
    ///         Enabled = false,
    ///     });
    /// 
    ///     var policy1 = new OpenStack.Firewall.PolicyV2("policy_1", new()
    ///     {
    ///         Name = "firewall_policy",
    ///         Rules = new[]
    ///         {
    ///             rule1.Id,
    ///             rule2.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Policies can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:firewall/policyV2:PolicyV2 policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:firewall/policyV2:PolicyV2")]
    public partial class PolicyV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Audit status of the firewall policy
        /// (must be "true" or "false" if provided - defaults to "false").
        /// This status is set to "false" whenever the firewall policy or any of its
        /// rules are changed. Changing this updates the `audited` status of an existing
        /// firewall policy.
        /// </summary>
        [Output("audited")]
        public Output<bool?> Audited { get; private set; } = null!;

        /// <summary>
        /// A description for the firewall policy. Changing
        /// this updates the `description` of an existing firewall policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A name for the firewall policy. Changing this
        /// updates the `name` of an existing firewall policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall policy. Required if admin wants
        /// to create a firewall policy for another project. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the v2 networking client.
        /// A networking client is needed to create a firewall policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// An array of one or more firewall rules that comprise
        /// the policy. Changing this results in adding/removing rules from the
        /// existing firewall policy.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<string>> Rules { get; private set; } = null!;

        /// <summary>
        /// Sharing status of the firewall policy (must be "true"
        /// or "false" if provided). If this is "true" the policy is visible to, and
        /// can be used in, firewalls in other tenants. Changing this updates the
        /// `shared` status of an existing firewall policy. Only administrative users
        /// can specify if the policy should be shared.
        /// </summary>
        [Output("shared")]
        public Output<bool?> Shared { get; private set; } = null!;

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall policy. Required if admin wants
        /// to create a firewall policy for another tenant. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyV2(string name, PolicyV2Args? args = null, CustomResourceOptions? options = null)
            : base("openstack:firewall/policyV2:PolicyV2", name, args ?? new PolicyV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyV2(string name, Input<string> id, PolicyV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:firewall/policyV2:PolicyV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyV2 Get(string name, Input<string> id, PolicyV2State? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyV2(name, id, state, options);
        }
    }

    public sealed class PolicyV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Audit status of the firewall policy
        /// (must be "true" or "false" if provided - defaults to "false").
        /// This status is set to "false" whenever the firewall policy or any of its
        /// rules are changed. Changing this updates the `audited` status of an existing
        /// firewall policy.
        /// </summary>
        [Input("audited")]
        public Input<bool>? Audited { get; set; }

        /// <summary>
        /// A description for the firewall policy. Changing
        /// this updates the `description` of an existing firewall policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A name for the firewall policy. Changing this
        /// updates the `name` of an existing firewall policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall policy. Required if admin wants
        /// to create a firewall policy for another project. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the v2 networking client.
        /// A networking client is needed to create a firewall policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("rules")]
        private InputList<string>? _rules;

        /// <summary>
        /// An array of one or more firewall rules that comprise
        /// the policy. Changing this results in adding/removing rules from the
        /// existing firewall policy.
        /// </summary>
        public InputList<string> Rules
        {
            get => _rules ?? (_rules = new InputList<string>());
            set => _rules = value;
        }

        /// <summary>
        /// Sharing status of the firewall policy (must be "true"
        /// or "false" if provided). If this is "true" the policy is visible to, and
        /// can be used in, firewalls in other tenants. Changing this updates the
        /// `shared` status of an existing firewall policy. Only administrative users
        /// can specify if the policy should be shared.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall policy. Required if admin wants
        /// to create a firewall policy for another tenant. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public PolicyV2Args()
        {
        }
        public static new PolicyV2Args Empty => new PolicyV2Args();
    }

    public sealed class PolicyV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Audit status of the firewall policy
        /// (must be "true" or "false" if provided - defaults to "false").
        /// This status is set to "false" whenever the firewall policy or any of its
        /// rules are changed. Changing this updates the `audited` status of an existing
        /// firewall policy.
        /// </summary>
        [Input("audited")]
        public Input<bool>? Audited { get; set; }

        /// <summary>
        /// A description for the firewall policy. Changing
        /// this updates the `description` of an existing firewall policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A name for the firewall policy. Changing this
        /// updates the `name` of an existing firewall policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall policy. Required if admin wants
        /// to create a firewall policy for another project. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the v2 networking client.
        /// A networking client is needed to create a firewall policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("rules")]
        private InputList<string>? _rules;

        /// <summary>
        /// An array of one or more firewall rules that comprise
        /// the policy. Changing this results in adding/removing rules from the
        /// existing firewall policy.
        /// </summary>
        public InputList<string> Rules
        {
            get => _rules ?? (_rules = new InputList<string>());
            set => _rules = value;
        }

        /// <summary>
        /// Sharing status of the firewall policy (must be "true"
        /// or "false" if provided). If this is "true" the policy is visible to, and
        /// can be used in, firewalls in other tenants. Changing this updates the
        /// `shared` status of an existing firewall policy. Only administrative users
        /// can specify if the policy should be shared.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall policy. Required if admin wants
        /// to create a firewall policy for another tenant. Changing this creates a new
        /// firewall policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public PolicyV2State()
        {
        }
        public static new PolicyV2State Empty => new PolicyV2State();
    }
}
