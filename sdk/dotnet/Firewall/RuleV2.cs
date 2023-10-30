// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Firewall
{
    /// <summary>
    /// Manages a v2 firewall rule resource within OpenStack.
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
    ///     var rule2 = new OpenStack.Firewall.RuleV2("rule2", new()
    ///     {
    ///         Action = "deny",
    ///         Description = "drop TELNET traffic",
    ///         DestinationPort = "23",
    ///         Enabled = true,
    ///         Protocol = "tcp",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Firewall Rules can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:firewall/ruleV2:RuleV2 rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:firewall/ruleV2:RuleV2")]
    public partial class RuleV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action to be taken (must be "allow", "deny" or "reject")
        /// when the firewall rule matches. Changing this updates the `action` of an
        /// existing firewall rule. Default is `deny`.
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// A description for the firewall rule. Changing this
        /// updates the `description` of an existing firewall rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The destination IP address on which the
        /// firewall rule operates. Changing this updates the `destination_ip_address`
        /// of an existing firewall rule.
        /// </summary>
        [Output("destinationIpAddress")]
        public Output<string?> DestinationIpAddress { get; private set; } = null!;

        /// <summary>
        /// The destination port on which the firewall
        /// rule operates. Changing this updates the `destination_port` of an existing
        /// firewall rule. Require not `any` or empty protocol.
        /// </summary>
        [Output("destinationPort")]
        public Output<string?> DestinationPort { get; private set; } = null!;

        /// <summary>
        /// Enabled status for the firewall rule (must be "true"
        /// or "false" if provided - defaults to "true"). Changing this updates the
        /// `enabled` status of an existing firewall rule.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// IP version, either 4 or 6. Changing this
        /// updates the `ip_version` of an existing firewall rule. Default is `4`.
        /// </summary>
        [Output("ipVersion")]
        public Output<int?> IpVersion { get; private set; } = null!;

        /// <summary>
        /// A unique name for the firewall rule. Changing this
        /// updates the `name` of an existing firewall rule.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall rule. Required if admin wants
        /// to create a firewall rule for another project. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// (Optional; Required if `source_port` or `destination_port` is not
        /// empty) The protocol type on which the firewall rule operates.
        /// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
        /// `protocol` of an existing firewall rule. Default is `any`.
        /// </summary>
        [Output("protocol")]
        public Output<string?> Protocol { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the v2 networking client.
        /// A networking client is needed to create a firewall rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Sharing status of the firewall rule (must be "true"
        /// or "false" if provided). If this is "true" the policy is visible to, and
        /// can be used in, firewalls in other tenants. Changing this updates the
        /// `shared` status of an existing firewall policy. On
        /// </summary>
        [Output("shared")]
        public Output<bool?> Shared { get; private set; } = null!;

        /// <summary>
        /// The source IP address on which the firewall
        /// rule operates. Changing this updates the `source_ip_address` of an existing
        /// firewall rule.
        /// </summary>
        [Output("sourceIpAddress")]
        public Output<string?> SourceIpAddress { get; private set; } = null!;

        /// <summary>
        /// The source port on which the firewall
        /// rule operates. Changing this updates the `source_port` of an existing
        /// firewall rule. Require not `any` or empty protocol.
        /// </summary>
        [Output("sourcePort")]
        public Output<string?> SourcePort { get; private set; } = null!;

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall rule. Required if admin wants
        /// to create a firewall rule for another tenant. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a RuleV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RuleV2(string name, RuleV2Args? args = null, CustomResourceOptions? options = null)
            : base("openstack:firewall/ruleV2:RuleV2", name, args ?? new RuleV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private RuleV2(string name, Input<string> id, RuleV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:firewall/ruleV2:RuleV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RuleV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RuleV2 Get(string name, Input<string> id, RuleV2State? state = null, CustomResourceOptions? options = null)
        {
            return new RuleV2(name, id, state, options);
        }
    }

    public sealed class RuleV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be taken (must be "allow", "deny" or "reject")
        /// when the firewall rule matches. Changing this updates the `action` of an
        /// existing firewall rule. Default is `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A description for the firewall rule. Changing this
        /// updates the `description` of an existing firewall rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination IP address on which the
        /// firewall rule operates. Changing this updates the `destination_ip_address`
        /// of an existing firewall rule.
        /// </summary>
        [Input("destinationIpAddress")]
        public Input<string>? DestinationIpAddress { get; set; }

        /// <summary>
        /// The destination port on which the firewall
        /// rule operates. Changing this updates the `destination_port` of an existing
        /// firewall rule. Require not `any` or empty protocol.
        /// </summary>
        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        /// <summary>
        /// Enabled status for the firewall rule (must be "true"
        /// or "false" if provided - defaults to "true"). Changing this updates the
        /// `enabled` status of an existing firewall rule.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// IP version, either 4 or 6. Changing this
        /// updates the `ip_version` of an existing firewall rule. Default is `4`.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// A unique name for the firewall rule. Changing this
        /// updates the `name` of an existing firewall rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall rule. Required if admin wants
        /// to create a firewall rule for another project. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// (Optional; Required if `source_port` or `destination_port` is not
        /// empty) The protocol type on which the firewall rule operates.
        /// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
        /// `protocol` of an existing firewall rule. Default is `any`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The region in which to obtain the v2 networking client.
        /// A networking client is needed to create a firewall rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Sharing status of the firewall rule (must be "true"
        /// or "false" if provided). If this is "true" the policy is visible to, and
        /// can be used in, firewalls in other tenants. Changing this updates the
        /// `shared` status of an existing firewall policy. On
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        /// <summary>
        /// The source IP address on which the firewall
        /// rule operates. Changing this updates the `source_ip_address` of an existing
        /// firewall rule.
        /// </summary>
        [Input("sourceIpAddress")]
        public Input<string>? SourceIpAddress { get; set; }

        /// <summary>
        /// The source port on which the firewall
        /// rule operates. Changing this updates the `source_port` of an existing
        /// firewall rule. Require not `any` or empty protocol.
        /// </summary>
        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall rule. Required if admin wants
        /// to create a firewall rule for another tenant. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public RuleV2Args()
        {
        }
        public static new RuleV2Args Empty => new RuleV2Args();
    }

    public sealed class RuleV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be taken (must be "allow", "deny" or "reject")
        /// when the firewall rule matches. Changing this updates the `action` of an
        /// existing firewall rule. Default is `deny`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A description for the firewall rule. Changing this
        /// updates the `description` of an existing firewall rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination IP address on which the
        /// firewall rule operates. Changing this updates the `destination_ip_address`
        /// of an existing firewall rule.
        /// </summary>
        [Input("destinationIpAddress")]
        public Input<string>? DestinationIpAddress { get; set; }

        /// <summary>
        /// The destination port on which the firewall
        /// rule operates. Changing this updates the `destination_port` of an existing
        /// firewall rule. Require not `any` or empty protocol.
        /// </summary>
        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        /// <summary>
        /// Enabled status for the firewall rule (must be "true"
        /// or "false" if provided - defaults to "true"). Changing this updates the
        /// `enabled` status of an existing firewall rule.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// IP version, either 4 or 6. Changing this
        /// updates the `ip_version` of an existing firewall rule. Default is `4`.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// A unique name for the firewall rule. Changing this
        /// updates the `name` of an existing firewall rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall rule. Required if admin wants
        /// to create a firewall rule for another project. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// (Optional; Required if `source_port` or `destination_port` is not
        /// empty) The protocol type on which the firewall rule operates.
        /// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
        /// `protocol` of an existing firewall rule. Default is `any`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The region in which to obtain the v2 networking client.
        /// A networking client is needed to create a firewall rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Sharing status of the firewall rule (must be "true"
        /// or "false" if provided). If this is "true" the policy is visible to, and
        /// can be used in, firewalls in other tenants. Changing this updates the
        /// `shared` status of an existing firewall policy. On
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        /// <summary>
        /// The source IP address on which the firewall
        /// rule operates. Changing this updates the `source_ip_address` of an existing
        /// firewall rule.
        /// </summary>
        [Input("sourceIpAddress")]
        public Input<string>? SourceIpAddress { get; set; }

        /// <summary>
        /// The source port on which the firewall
        /// rule operates. Changing this updates the `source_port` of an existing
        /// firewall rule. Require not `any` or empty protocol.
        /// </summary>
        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall rule. Required if admin wants
        /// to create a firewall rule for another tenant. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public RuleV2State()
        {
        }
        public static new RuleV2State Empty => new RuleV2State();
    }
}
