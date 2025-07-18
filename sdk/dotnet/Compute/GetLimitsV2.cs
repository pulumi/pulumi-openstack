// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    public static class GetLimitsV2
    {
        /// <summary>
        /// Use this data source to get the compute limits of an OpenStack project.
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
        ///     var limits = OpenStack.Compute.GetLimitsV2.Invoke(new()
        ///     {
        ///         ProjectId = "2e367a3d29f94fd988e6ec54e305ec9d",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLimitsV2Result> InvokeAsync(GetLimitsV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLimitsV2Result>("openstack:compute/getLimitsV2:getLimitsV2", args ?? new GetLimitsV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the compute limits of an OpenStack project.
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
        ///     var limits = OpenStack.Compute.GetLimitsV2.Invoke(new()
        ///     {
        ///         ProjectId = "2e367a3d29f94fd988e6ec54e305ec9d",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLimitsV2Result> Invoke(GetLimitsV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLimitsV2Result>("openstack:compute/getLimitsV2:getLimitsV2", args ?? new GetLimitsV2InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the compute limits of an OpenStack project.
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
        ///     var limits = OpenStack.Compute.GetLimitsV2.Invoke(new()
        ///     {
        ///         ProjectId = "2e367a3d29f94fd988e6ec54e305ec9d",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLimitsV2Result> Invoke(GetLimitsV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLimitsV2Result>("openstack:compute/getLimitsV2:getLimitsV2", args ?? new GetLimitsV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLimitsV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the project to retrieve the limits.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetLimitsV2Args()
        {
        }
        public static new GetLimitsV2Args Empty => new GetLimitsV2Args();
    }

    public sealed class GetLimitsV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the project to retrieve the limits.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetLimitsV2InvokeArgs()
        {
        }
        public static new GetLimitsV2InvokeArgs Empty => new GetLimitsV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetLimitsV2Result
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The number of allowed metadata items for each image. Starting from version 2.39 this field is dropped from ‘os-limits’ response, because ‘image-metadata’ proxy API was deprecated. Available until version 2.38.
        /// </summary>
        public readonly int MaxImageMeta;
        /// <summary>
        /// The number of allowed injected files for the tenant. Available until version 2.56.
        /// </summary>
        public readonly int MaxPersonality;
        /// <summary>
        /// The number of allowed bytes of content for each injected file. Available until version 2.56.
        /// </summary>
        public readonly int MaxPersonalitySize;
        /// <summary>
        /// The number of allowed rules for each security group. Available until version 2.35.
        /// </summary>
        public readonly int MaxSecurityGroupRules;
        /// <summary>
        /// The number of allowed security groups for the tenant. Available until version 2.35.
        /// </summary>
        public readonly int MaxSecurityGroups;
        /// <summary>
        /// The number of allowed members for each server group.
        /// </summary>
        public readonly int MaxServerGroupMembers;
        /// <summary>
        /// The number of allowed server groups for the tenant.
        /// </summary>
        public readonly int MaxServerGroups;
        /// <summary>
        /// The number of allowed server groups for the tenant.
        /// </summary>
        public readonly int MaxServerMeta;
        /// <summary>
        /// The number of allowed server cores for the tenant.
        /// </summary>
        public readonly int MaxTotalCores;
        /// <summary>
        /// The number of allowed floating IP addresses for each tenant. Available until version 2.35.
        /// </summary>
        public readonly int MaxTotalFloatingIps;
        /// <summary>
        /// The number of allowed servers for the tenant.
        /// </summary>
        public readonly int MaxTotalInstances;
        /// <summary>
        /// The number of allowed key pairs for the user.
        /// </summary>
        public readonly int MaxTotalKeypairs;
        /// <summary>
        /// The number of allowed floating IP addresses for the tenant. Available until version 2.35.
        /// </summary>
        public readonly int MaxTotalRamSize;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The number of used server cores in the tenant.
        /// </summary>
        public readonly int TotalCoresUsed;
        /// <summary>
        /// The number of used floating IP addresses in the tenant.
        /// </summary>
        public readonly int TotalFloatingIpsUsed;
        /// <summary>
        /// The number of used server cores in the tenant.
        /// </summary>
        public readonly int TotalInstancesUsed;
        /// <summary>
        /// The amount of used server RAM in the tenant.
        /// </summary>
        public readonly int TotalRamUsed;
        /// <summary>
        /// The number of used security groups in the tenant. Available until version 2.35.
        /// </summary>
        public readonly int TotalSecurityGroupsUsed;
        /// <summary>
        /// The number of used server groups in each tenant.
        /// </summary>
        public readonly int TotalServerGroupsUsed;

        [OutputConstructor]
        private GetLimitsV2Result(
            string id,

            int maxImageMeta,

            int maxPersonality,

            int maxPersonalitySize,

            int maxSecurityGroupRules,

            int maxSecurityGroups,

            int maxServerGroupMembers,

            int maxServerGroups,

            int maxServerMeta,

            int maxTotalCores,

            int maxTotalFloatingIps,

            int maxTotalInstances,

            int maxTotalKeypairs,

            int maxTotalRamSize,

            string projectId,

            string region,

            int totalCoresUsed,

            int totalFloatingIpsUsed,

            int totalInstancesUsed,

            int totalRamUsed,

            int totalSecurityGroupsUsed,

            int totalServerGroupsUsed)
        {
            Id = id;
            MaxImageMeta = maxImageMeta;
            MaxPersonality = maxPersonality;
            MaxPersonalitySize = maxPersonalitySize;
            MaxSecurityGroupRules = maxSecurityGroupRules;
            MaxSecurityGroups = maxSecurityGroups;
            MaxServerGroupMembers = maxServerGroupMembers;
            MaxServerGroups = maxServerGroups;
            MaxServerMeta = maxServerMeta;
            MaxTotalCores = maxTotalCores;
            MaxTotalFloatingIps = maxTotalFloatingIps;
            MaxTotalInstances = maxTotalInstances;
            MaxTotalKeypairs = maxTotalKeypairs;
            MaxTotalRamSize = maxTotalRamSize;
            ProjectId = projectId;
            Region = region;
            TotalCoresUsed = totalCoresUsed;
            TotalFloatingIpsUsed = totalFloatingIpsUsed;
            TotalInstancesUsed = totalInstancesUsed;
            TotalRamUsed = totalRamUsed;
            TotalSecurityGroupsUsed = totalSecurityGroupsUsed;
            TotalServerGroupsUsed = totalServerGroupsUsed;
        }
    }
}
