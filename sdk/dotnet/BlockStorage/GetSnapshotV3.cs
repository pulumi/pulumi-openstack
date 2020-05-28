// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage
{
    public static class GetSnapshotV3
    {
        /// <summary>
        /// Use this data source to get information about an existing snapshot.
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
        ///         var snapshot1 = Output.Create(OpenStack.BlockStorage.GetSnapshotV3.InvokeAsync(new OpenStack.BlockStorage.GetSnapshotV3Args
        ///         {
        ///             MostRecent = true,
        ///             Name = "snapshot_1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSnapshotV3Result> InvokeAsync(GetSnapshotV3Args? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotV3Result>("openstack:blockstorage/getSnapshotV3:getSnapshotV3", args ?? new GetSnapshotV3Args(), options.WithVersion());
    }


    public sealed class GetSnapshotV3Args : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Pick the most recently created snapshot if there
        /// are multiple results.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Block Storage
        /// client. If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The status of the snapshot.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The ID of the snapshot's volume.
        /// </summary>
        [Input("volumeId")]
        public string? VolumeId { get; set; }

        public GetSnapshotV3Args()
        {
        }
    }


    [OutputType]
    public sealed class GetSnapshotV3Result
    {
        /// <summary>
        /// The snapshot's description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The snapshot's metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly bool? MostRecent;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The size of the snapshot.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string VolumeId;

        [OutputConstructor]
        private GetSnapshotV3Result(
            string description,

            string id,

            ImmutableDictionary<string, object> metadata,

            bool? mostRecent,

            string name,

            string region,

            int size,

            string status,

            string volumeId)
        {
            Description = description;
            Id = id;
            Metadata = metadata;
            MostRecent = mostRecent;
            Name = name;
            Region = region;
            Size = size;
            Status = status;
            VolumeId = volumeId;
        }
    }
}
