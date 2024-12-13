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
        ///     var snapshot1 = OpenStack.BlockStorage.GetSnapshotV3.Invoke(new()
        ///     {
        ///         Name = "snapshot_1",
        ///         MostRecent = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSnapshotV3Result> InvokeAsync(GetSnapshotV3Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotV3Result>("openstack:blockstorage/getSnapshotV3:getSnapshotV3", args ?? new GetSnapshotV3Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about an existing snapshot.
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
        ///     var snapshot1 = OpenStack.BlockStorage.GetSnapshotV3.Invoke(new()
        ///     {
        ///         Name = "snapshot_1",
        ///         MostRecent = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSnapshotV3Result> Invoke(GetSnapshotV3InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotV3Result>("openstack:blockstorage/getSnapshotV3:getSnapshotV3", args ?? new GetSnapshotV3InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about an existing snapshot.
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
        ///     var snapshot1 = OpenStack.BlockStorage.GetSnapshotV3.Invoke(new()
        ///     {
        ///         Name = "snapshot_1",
        ///         MostRecent = true,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSnapshotV3Result> Invoke(GetSnapshotV3InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotV3Result>("openstack:blockstorage/getSnapshotV3:getSnapshotV3", args ?? new GetSnapshotV3InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotV3Args : global::Pulumi.InvokeArgs
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
        public static new GetSnapshotV3Args Empty => new GetSnapshotV3Args();
    }

    public sealed class GetSnapshotV3InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Pick the most recently created snapshot if there
        /// are multiple results.
        /// </summary>
        [Input("mostRecent")]
        public Input<bool>? MostRecent { get; set; }

        /// <summary>
        /// The name of the snapshot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Block Storage
        /// client. If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The status of the snapshot.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The ID of the snapshot's volume.
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        public GetSnapshotV3InvokeArgs()
        {
        }
        public static new GetSnapshotV3InvokeArgs Empty => new GetSnapshotV3InvokeArgs();
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
        public readonly ImmutableDictionary<string, string> Metadata;
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

            ImmutableDictionary<string, string> metadata,

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
