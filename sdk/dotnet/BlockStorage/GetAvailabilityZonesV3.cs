// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage
{
    public static class GetAvailabilityZonesV3
    {
        /// <summary>
        /// Use this data source to get a list of Block Storage availability zones from OpenStack
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
        ///         var zones = Output.Create(OpenStack.BlockStorage.GetAvailabilityZonesV3.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAvailabilityZonesV3Result> InvokeAsync(GetAvailabilityZonesV3Args? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityZonesV3Result>("openstack:blockstorage/getAvailabilityZonesV3:getAvailabilityZonesV3", args ?? new GetAvailabilityZonesV3Args(), options.WithVersion());
    }


    public sealed class GetAvailabilityZonesV3Args : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The region in which to obtain the Block Storage client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The `state` of the availability zones to match. Can
        /// either be `available` or `unavailable`. Default is `available`.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetAvailabilityZonesV3Args()
        {
        }
    }


    [OutputType]
    public sealed class GetAvailabilityZonesV3Result
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The names of the availability zones, ordered alphanumerically, that
        /// match the queried `state`.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetAvailabilityZonesV3Result(
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
