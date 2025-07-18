// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.LoadBalancer
{
    public static class GetLbFlavorDeprecated
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack Load Balancer flavor.
        /// 
        /// &gt; **Note:** This data source is deprecated, please use `openstack.loadbalancer.FlavorV2` instead.
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
        ///     var flavor1 = OpenStack.LoadBalancer.GetLbFlavorDeprecated.Invoke(new()
        ///     {
        ///         Name = "flavor_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLbFlavorDeprecatedResult> InvokeAsync(GetLbFlavorDeprecatedArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLbFlavorDeprecatedResult>("openstack:loadbalancer/getLbFlavorDeprecated:getLbFlavorDeprecated", args ?? new GetLbFlavorDeprecatedArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an OpenStack Load Balancer flavor.
        /// 
        /// &gt; **Note:** This data source is deprecated, please use `openstack.loadbalancer.FlavorV2` instead.
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
        ///     var flavor1 = OpenStack.LoadBalancer.GetLbFlavorDeprecated.Invoke(new()
        ///     {
        ///         Name = "flavor_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLbFlavorDeprecatedResult> Invoke(GetLbFlavorDeprecatedInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbFlavorDeprecatedResult>("openstack:loadbalancer/getLbFlavorDeprecated:getLbFlavorDeprecated", args ?? new GetLbFlavorDeprecatedInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an OpenStack Load Balancer flavor.
        /// 
        /// &gt; **Note:** This data source is deprecated, please use `openstack.loadbalancer.FlavorV2` instead.
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
        ///     var flavor1 = OpenStack.LoadBalancer.GetLbFlavorDeprecated.Invoke(new()
        ///     {
        ///         Name = "flavor_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLbFlavorDeprecatedResult> Invoke(GetLbFlavorDeprecatedInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLbFlavorDeprecatedResult>("openstack:loadbalancer/getLbFlavorDeprecated:getLbFlavorDeprecated", args ?? new GetLbFlavorDeprecatedInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbFlavorDeprecatedArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the flavor. Exactly one of `name`, `flavor_id` is required to be set.
        /// </summary>
        [Input("flavorId")]
        public string? FlavorId { get; set; }

        /// <summary>
        /// The name of the flavor. Exactly one of `name`, `flavor_id` is required to be set.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Load Balancer client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetLbFlavorDeprecatedArgs()
        {
        }
        public static new GetLbFlavorDeprecatedArgs Empty => new GetLbFlavorDeprecatedArgs();
    }

    public sealed class GetLbFlavorDeprecatedInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the flavor. Exactly one of `name`, `flavor_id` is required to be set.
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// The name of the flavor. Exactly one of `name`, `flavor_id` is required to be set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Load Balancer client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetLbFlavorDeprecatedInvokeArgs()
        {
        }
        public static new GetLbFlavorDeprecatedInvokeArgs Empty => new GetLbFlavorDeprecatedInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbFlavorDeprecatedResult
    {
        /// <summary>
        /// The description of the flavor.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Is the flavor enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The ID of the flavor.
        /// </summary>
        public readonly string FlavorId;
        /// <summary>
        /// The ID of the flavor profile.
        /// </summary>
        public readonly string FlavorProfileId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the flavor.
        /// </summary>
        public readonly string Name;
        public readonly string Region;

        [OutputConstructor]
        private GetLbFlavorDeprecatedResult(
            string description,

            bool enabled,

            string flavorId,

            string flavorProfileId,

            string id,

            string name,

            string region)
        {
            Description = description;
            Enabled = enabled;
            FlavorId = flavorId;
            FlavorProfileId = flavorProfileId;
            Id = id;
            Name = name;
            Region = region;
        }
    }
}
