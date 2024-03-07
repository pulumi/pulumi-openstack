// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack
{
    public static class GetFwGroupV2
    {
        /// <summary>
        /// Use this data source to get information of an available OpenStack firewall group v2.
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
        ///     var @group = OpenStack.GetFwGroupV2.Invoke(new()
        ///     {
        ///         Name = "tf_test_group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetFwGroupV2Result> InvokeAsync(GetFwGroupV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFwGroupV2Result>("openstack:index/getFwGroupV2:getFwGroupV2", args ?? new GetFwGroupV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information of an available OpenStack firewall group v2.
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
        ///     var @group = OpenStack.GetFwGroupV2.Invoke(new()
        ///     {
        ///         Name = "tf_test_group",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetFwGroupV2Result> Invoke(GetFwGroupV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFwGroupV2Result>("openstack:index/getFwGroupV2:getFwGroupV2", args ?? new GetFwGroupV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFwGroupV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Administrative up/down status for the firewall group.
        /// </summary>
        [Input("adminStateUp")]
        public bool? AdminStateUp { get; set; }

        /// <summary>
        /// Human-readable description of the firewall group.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The egress policy ID of the firewall group.
        /// </summary>
        [Input("egressFirewallPolicyId")]
        public string? EgressFirewallPolicyId { get; set; }

        /// <summary>
        /// The ID of the firewall group.
        /// </summary>
        [Input("groupId")]
        public string? GroupId { get; set; }

        /// <summary>
        /// The ingress policy ID of the firewall group.
        /// </summary>
        [Input("ingressFirewallPolicyId")]
        public string? IngressFirewallPolicyId { get; set; }

        /// <summary>
        /// The name of the firewall group.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall group.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve firewall group ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The sharing status of the firewall group.
        /// </summary>
        [Input("shared")]
        public bool? Shared { get; set; }

        /// <summary>
        /// Enabled status for the firewall group.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall group.
        /// </summary>
        [Input("tenantId")]
        public string? TenantId { get; set; }

        public GetFwGroupV2Args()
        {
        }
        public static new GetFwGroupV2Args Empty => new GetFwGroupV2Args();
    }

    public sealed class GetFwGroupV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Administrative up/down status for the firewall group.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// Human-readable description of the firewall group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The egress policy ID of the firewall group.
        /// </summary>
        [Input("egressFirewallPolicyId")]
        public Input<string>? EgressFirewallPolicyId { get; set; }

        /// <summary>
        /// The ID of the firewall group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The ingress policy ID of the firewall group.
        /// </summary>
        [Input("ingressFirewallPolicyId")]
        public Input<string>? IngressFirewallPolicyId { get; set; }

        /// <summary>
        /// The name of the firewall group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall group.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve firewall group ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The sharing status of the firewall group.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        /// <summary>
        /// Enabled status for the firewall group.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall group.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public GetFwGroupV2InvokeArgs()
        {
        }
        public static new GetFwGroupV2InvokeArgs Empty => new GetFwGroupV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetFwGroupV2Result
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool AdminStateUp;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? EgressFirewallPolicyId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? GroupId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? IngressFirewallPolicyId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Ports associated with the firewall group.
        /// </summary>
        public readonly ImmutableArray<string> Ports;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool Shared;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetFwGroupV2Result(
            bool adminStateUp,

            string? description,

            string? egressFirewallPolicyId,

            string? groupId,

            string id,

            string? ingressFirewallPolicyId,

            string? name,

            ImmutableArray<string> ports,

            string projectId,

            string region,

            bool shared,

            string status,

            string tenantId)
        {
            AdminStateUp = adminStateUp;
            Description = description;
            EgressFirewallPolicyId = egressFirewallPolicyId;
            GroupId = groupId;
            Id = id;
            IngressFirewallPolicyId = ingressFirewallPolicyId;
            Name = name;
            Ports = ports;
            ProjectId = projectId;
            Region = region;
            Shared = shared;
            Status = status;
            TenantId = tenantId;
        }
    }
}
