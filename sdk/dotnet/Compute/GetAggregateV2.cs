// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    public static class GetAggregateV2
    {
        /// <summary>
        /// Use this data source to get information about host aggregates
        /// by name.
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
        ///         var test = Output.Create(OpenStack.Compute.GetAggregateV2.InvokeAsync(new OpenStack.Compute.GetAggregateV2Args
        ///         {
        ///             Name = "test",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAggregateV2Result> InvokeAsync(GetAggregateV2Args args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAggregateV2Result>("openstack:compute/getAggregateV2:getAggregateV2", args ?? new GetAggregateV2Args(), options.WithVersion());
    }


    public sealed class GetAggregateV2Args : Pulumi.InvokeArgs
    {
        [Input("hosts")]
        private List<string>? _hosts;

        /// <summary>
        /// List of Hypervisors contained in the Host Aggregate
        /// </summary>
        public List<string> Hosts
        {
            get => _hosts ?? (_hosts = new List<string>());
            set => _hosts = value;
        }

        [Input("metadata")]
        private Dictionary<string, string>? _metadata;

        /// <summary>
        /// Metadata of the Host Aggregate
        /// </summary>
        public Dictionary<string, string> Metadata
        {
            get => _metadata ?? (_metadata = new Dictionary<string, string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the host aggregate
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetAggregateV2Args()
        {
        }
    }


    [OutputType]
    public sealed class GetAggregateV2Result
    {
        /// <summary>
        /// List of Hypervisors contained in the Host Aggregate
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Metadata of the Host Aggregate
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Availability zone of the Host Aggregate
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetAggregateV2Result(
            ImmutableArray<string> hosts,

            string id,

            ImmutableDictionary<string, string> metadata,

            string name,

            string zone)
        {
            Hosts = hosts;
            Id = id;
            Metadata = metadata;
            Name = name;
            Zone = zone;
        }
    }
}
