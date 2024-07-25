// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.LoadBalancer
{
    public static class GetFlavorV2
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack Load Balancer flavor.
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
        ///     var flavor1 = OpenStack.LoadBalancer.GetFlavorV2.Invoke(new()
        ///     {
        ///         Name = "flavor_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFlavorV2Result> InvokeAsync(GetFlavorV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFlavorV2Result>("openstack:loadbalancer/getFlavorV2:getFlavorV2", args ?? new GetFlavorV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an OpenStack Load Balancer flavor.
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
        ///     var flavor1 = OpenStack.LoadBalancer.GetFlavorV2.Invoke(new()
        ///     {
        ///         Name = "flavor_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFlavorV2Result> Invoke(GetFlavorV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlavorV2Result>("openstack:loadbalancer/getFlavorV2:getFlavorV2", args ?? new GetFlavorV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFlavorV2Args : global::Pulumi.InvokeArgs
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

        public GetFlavorV2Args()
        {
        }
        public static new GetFlavorV2Args Empty => new GetFlavorV2Args();
    }

    public sealed class GetFlavorV2InvokeArgs : global::Pulumi.InvokeArgs
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

        public GetFlavorV2InvokeArgs()
        {
        }
        public static new GetFlavorV2InvokeArgs Empty => new GetFlavorV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetFlavorV2Result
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
        public readonly string? Region;

        [OutputConstructor]
        private GetFlavorV2Result(
            string description,

            bool enabled,

            string flavorId,

            string flavorProfileId,

            string id,

            string name,

            string? region)
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
