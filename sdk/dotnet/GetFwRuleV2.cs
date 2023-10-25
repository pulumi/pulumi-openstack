// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack
{
    public static class GetFwRuleV2
    {
        /// <summary>
        /// Use this data source to get information of an available OpenStack firewall rule v2.
        /// </summary>
        public static Task<GetFwRuleV2Result> InvokeAsync(GetFwRuleV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFwRuleV2Result>("openstack:index/getFwRuleV2:getFwRuleV2", args ?? new GetFwRuleV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information of an available OpenStack firewall rule v2.
        /// </summary>
        public static Output<GetFwRuleV2Result> Invoke(GetFwRuleV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFwRuleV2Result>("openstack:index/getFwRuleV2:getFwRuleV2", args ?? new GetFwRuleV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFwRuleV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Action to be taken when the firewall rule matches.
        /// </summary>
        [Input("action")]
        public string? Action { get; set; }

        /// <summary>
        /// The description of the firewall rule.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The destination IP address on which the
        /// firewall rule operates.
        /// </summary>
        [Input("destinationIpAddress")]
        public string? DestinationIpAddress { get; set; }

        /// <summary>
        /// The destination port on which the firewall
        /// rule operates.
        /// </summary>
        [Input("destinationPort")]
        public string? DestinationPort { get; set; }

        /// <summary>
        /// Enabled status for the firewall rule.
        /// </summary>
        [Input("enabled")]
        public bool? Enabled { get; set; }

        [Input("firewallPolicyIds")]
        private List<string>? _firewallPolicyIds;

        /// <summary>
        /// The ID of the firewall policy the rule belongs to.
        /// </summary>
        public List<string> FirewallPolicyIds
        {
            get => _firewallPolicyIds ?? (_firewallPolicyIds = new List<string>());
            set => _firewallPolicyIds = value;
        }

        /// <summary>
        /// IP version, either 4 (default) or 6.
        /// </summary>
        [Input("ipVersion")]
        public int? IpVersion { get; set; }

        /// <summary>
        /// The name of the firewall rule.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The protocol type on which the firewall rule operates.
        /// </summary>
        [Input("protocol")]
        public string? Protocol { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve firewall policy ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The ID of the firewall rule.
        /// </summary>
        [Input("ruleId")]
        public string? RuleId { get; set; }

        /// <summary>
        /// The sharing status of the firewall policy.
        /// </summary>
        [Input("shared")]
        public bool? Shared { get; set; }

        /// <summary>
        /// The source IP address on which the firewall
        /// rule operates.
        /// </summary>
        [Input("sourceIpAddress")]
        public string? SourceIpAddress { get; set; }

        /// <summary>
        /// The source port on which the firewall
        /// rule operates.
        /// </summary>
        [Input("sourcePort")]
        public string? SourcePort { get; set; }

        /// <summary>
        /// The owner of the firewall policy.
        /// </summary>
        [Input("tenantId")]
        public string? TenantId { get; set; }

        public GetFwRuleV2Args()
        {
        }
        public static new GetFwRuleV2Args Empty => new GetFwRuleV2Args();
    }

    public sealed class GetFwRuleV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Action to be taken when the firewall rule matches.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The description of the firewall rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination IP address on which the
        /// firewall rule operates.
        /// </summary>
        [Input("destinationIpAddress")]
        public Input<string>? DestinationIpAddress { get; set; }

        /// <summary>
        /// The destination port on which the firewall
        /// rule operates.
        /// </summary>
        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        /// <summary>
        /// Enabled status for the firewall rule.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("firewallPolicyIds")]
        private InputList<string>? _firewallPolicyIds;

        /// <summary>
        /// The ID of the firewall policy the rule belongs to.
        /// </summary>
        public InputList<string> FirewallPolicyIds
        {
            get => _firewallPolicyIds ?? (_firewallPolicyIds = new InputList<string>());
            set => _firewallPolicyIds = value;
        }

        /// <summary>
        /// IP version, either 4 (default) or 6.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// The name of the firewall rule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol type on which the firewall rule operates.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve firewall policy ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the firewall rule.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// The sharing status of the firewall policy.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        /// <summary>
        /// The source IP address on which the firewall
        /// rule operates.
        /// </summary>
        [Input("sourceIpAddress")]
        public Input<string>? SourceIpAddress { get; set; }

        /// <summary>
        /// The source port on which the firewall
        /// rule operates.
        /// </summary>
        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        /// <summary>
        /// The owner of the firewall policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public GetFwRuleV2InvokeArgs()
        {
        }
        public static new GetFwRuleV2InvokeArgs Empty => new GetFwRuleV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetFwRuleV2Result
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? DestinationIpAddress;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? DestinationPort;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The ID of the firewall policy the rule belongs to.
        /// </summary>
        public readonly ImmutableArray<string> FirewallPolicyIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int? IpVersion;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? RuleId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool Shared;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? SourceIpAddress;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? SourcePort;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetFwRuleV2Result(
            string? action,

            string? description,

            string? destinationIpAddress,

            string? destinationPort,

            bool enabled,

            ImmutableArray<string> firewallPolicyIds,

            string id,

            int? ipVersion,

            string? name,

            string? protocol,

            string region,

            string? ruleId,

            bool shared,

            string? sourceIpAddress,

            string? sourcePort,

            string tenantId)
        {
            Action = action;
            Description = description;
            DestinationIpAddress = destinationIpAddress;
            DestinationPort = destinationPort;
            Enabled = enabled;
            FirewallPolicyIds = firewallPolicyIds;
            Id = id;
            IpVersion = ipVersion;
            Name = name;
            Protocol = protocol;
            Region = region;
            RuleId = ruleId;
            Shared = shared;
            SourceIpAddress = sourceIpAddress;
            SourcePort = sourcePort;
            TenantId = tenantId;
        }
    }
}
