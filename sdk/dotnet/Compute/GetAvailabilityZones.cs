// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    public static class GetAvailabilityZones
    {
        /// <summary>
        /// Use this data source to get a list of availability zones from OpenStack
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var zones = Output.Create(OpenStack.Compute.GetAvailabilityZones.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAvailabilityZonesResult> InvokeAsync(GetAvailabilityZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityZonesResult>("openstack:compute/getAvailabilityZones:getAvailabilityZones", args ?? new GetAvailabilityZonesArgs(), options.WithVersion());
    }


    public sealed class GetAvailabilityZonesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The `region` to fetch availability zones from, defaults to the provider's `region`
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The `state` of the availability zones to match, default ("available").
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetAvailabilityZonesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAvailabilityZonesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The names of the availability zones, ordered alphanumerically, that match the queried `state`
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string Region;
        public readonly string? State;

        [OutputConstructor]
        private GetAvailabilityZonesResult(
            string id,

            ImmutableArray<string> names,

            string region,

            string? state)
        {
            Id = id;
            Names = names;
            Region = region;
            State = state;
        }
    }
}
