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
    /// Manages a v1 firewall rule resource within OpenStack.
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
    ///     var rule1 = new OpenStack.Firewall.Rule("rule_1", new()
    ///     {
    ///         Name = "my_rule",
    ///         Description = "drop TELNET traffic",
    ///         Action = "deny",
    ///         Protocol = "tcp",
    ///         DestinationPort = "23",
    ///         Enabled = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Firewall Rules can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:firewall/rule:Rule rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:firewall/rule:Rule")]
    public partial class Rule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action to be taken ( must be "allow" or "deny") when the
        /// firewall rule matches. Changing this updates the `action` of an existing
        /// firewall rule.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

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
        /// firewall rule.
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
        /// IP version, either 4 (default) or 6. Changing this
        /// updates the `ip_version` of an existing firewall rule.
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
        /// The protocol type on which the firewall rule operates.
        /// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
        /// `protocol` of an existing firewall rule.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the v1 Compute client.
        /// A Compute client is needed to create a firewall rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

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
        /// firewall rule.
        /// </summary>
        [Output("sourcePort")]
        public Output<string?> SourcePort { get; private set; } = null!;

        /// <summary>
        /// The owner of the firewall rule. Required if admin
        /// wants to create a firewall rule for another tenant. Changing this creates a
        /// new firewall rule.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a Rule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Rule(string name, RuleArgs args, CustomResourceOptions? options = null)
            : base("openstack:firewall/rule:Rule", name, args ?? new RuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Rule(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
            : base("openstack:firewall/rule:Rule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Rule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Rule Get(string name, Input<string> id, RuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Rule(name, id, state, options);
        }
    }

    public sealed class RuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be taken ( must be "allow" or "deny") when the
        /// firewall rule matches. Changing this updates the `action` of an existing
        /// firewall rule.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

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
        /// firewall rule.
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
        /// IP version, either 4 (default) or 6. Changing this
        /// updates the `ip_version` of an existing firewall rule.
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
        /// The protocol type on which the firewall rule operates.
        /// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
        /// `protocol` of an existing firewall rule.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the v1 Compute client.
        /// A Compute client is needed to create a firewall rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

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
        /// firewall rule.
        /// </summary>
        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        /// <summary>
        /// The owner of the firewall rule. Required if admin
        /// wants to create a firewall rule for another tenant. Changing this creates a
        /// new firewall rule.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public RuleArgs()
        {
        }
        public static new RuleArgs Empty => new RuleArgs();
    }

    public sealed class RuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action to be taken ( must be "allow" or "deny") when the
        /// firewall rule matches. Changing this updates the `action` of an existing
        /// firewall rule.
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
        /// firewall rule.
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
        /// IP version, either 4 (default) or 6. Changing this
        /// updates the `ip_version` of an existing firewall rule.
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
        /// The protocol type on which the firewall rule operates.
        /// Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
        /// `protocol` of an existing firewall rule.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The region in which to obtain the v1 Compute client.
        /// A Compute client is needed to create a firewall rule. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// firewall rule.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

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
        /// firewall rule.
        /// </summary>
        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        /// <summary>
        /// The owner of the firewall rule. Required if admin
        /// wants to create a firewall rule for another tenant. Changing this creates a
        /// new firewall rule.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public RuleState()
        {
        }
        public static new RuleState Empty => new RuleState();
    }
}
