// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack
{
    [Obsolete(@"openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2")]
    public static class GetFwPolicyV2
    {
        /// <summary>
        /// Use this data source to get information of an available OpenStack firewall policy v2.
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
        ///     var policy = OpenStack.Firewall.GetPolicyV2.Invoke(new()
        ///     {
        ///         Name = "tf_test_policy",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFwPolicyV2Result> InvokeAsync(GetFwPolicyV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFwPolicyV2Result>("openstack:index/getFwPolicyV2:getFwPolicyV2", args ?? new GetFwPolicyV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information of an available OpenStack firewall policy v2.
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
        ///     var policy = OpenStack.Firewall.GetPolicyV2.Invoke(new()
        ///     {
        ///         Name = "tf_test_policy",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFwPolicyV2Result> Invoke(GetFwPolicyV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFwPolicyV2Result>("openstack:index/getFwPolicyV2:getFwPolicyV2", args ?? new GetFwPolicyV2InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information of an available OpenStack firewall policy v2.
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
        ///     var policy = OpenStack.Firewall.GetPolicyV2.Invoke(new()
        ///     {
        ///         Name = "tf_test_policy",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFwPolicyV2Result> Invoke(GetFwPolicyV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFwPolicyV2Result>("openstack:index/getFwPolicyV2:getFwPolicyV2", args ?? new GetFwPolicyV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFwPolicyV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether this policy has been audited.
        /// </summary>
        [Input("audited")]
        public bool? Audited { get; set; }

        /// <summary>
        /// Human-readable description of the policy.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The name of the firewall policy.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the firewall policy.
        /// </summary>
        [Input("policyId")]
        public string? PolicyId { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall policy.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve firewall policy ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Whether this policy is shared across all projects.
        /// </summary>
        [Input("shared")]
        public bool? Shared { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall policy.
        /// </summary>
        [Input("tenantId")]
        public string? TenantId { get; set; }

        public GetFwPolicyV2Args()
        {
        }
        public static new GetFwPolicyV2Args Empty => new GetFwPolicyV2Args();
    }

    public sealed class GetFwPolicyV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether this policy has been audited.
        /// </summary>
        [Input("audited")]
        public Input<bool>? Audited { get; set; }

        /// <summary>
        /// Human-readable description of the policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the firewall policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the firewall policy.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `tenant_id`. The owner of the firewall policy.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve firewall policy ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Whether this policy is shared across all projects.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        /// <summary>
        /// This argument conflicts and is interchangeable
        /// with `project_id`. The owner of the firewall policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public GetFwPolicyV2InvokeArgs()
        {
        }
        public static new GetFwPolicyV2InvokeArgs Empty => new GetFwPolicyV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetFwPolicyV2Result
    {
        /// <summary>
        /// The audit status of the firewall policy.
        /// </summary>
        public readonly bool Audited;
        public readonly string? Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? PolicyId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The array of one or more firewall rules that comprise the policy.
        /// </summary>
        public readonly ImmutableArray<string> Rules;
        /// <summary>
        /// The sharing status of the firewall policy.
        /// </summary>
        public readonly bool Shared;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetFwPolicyV2Result(
            bool audited,

            string? description,

            string id,

            string? name,

            string? policyId,

            string projectId,

            string region,

            ImmutableArray<string> rules,

            bool shared,

            string tenantId)
        {
            Audited = audited;
            Description = description;
            Id = id;
            Name = name;
            PolicyId = policyId;
            ProjectId = projectId;
            Region = region;
            Rules = rules;
            Shared = shared;
            TenantId = tenantId;
        }
    }
}
