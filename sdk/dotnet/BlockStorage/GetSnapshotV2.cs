// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage
{
    public static class GetSnapshotV2
    {
        /// <summary>
        /// Use this data source to get information about an existing snapshot.
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
        ///     var snapshot1 = OpenStack.BlockStorage.GetSnapshotV2.Invoke(new()
        ///     {
        ///         MostRecent = true,
        ///         Name = "snapshot_1",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSnapshotV2Result> InvokeAsync(GetSnapshotV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotV2Result>("openstack:blockstorage/getSnapshotV2:getSnapshotV2", args ?? new GetSnapshotV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about an existing snapshot.
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
        ///     var snapshot1 = OpenStack.BlockStorage.GetSnapshotV2.Invoke(new()
        ///     {
        ///         MostRecent = true,
        ///         Name = "snapshot_1",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSnapshotV2Result> Invoke(GetSnapshotV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotV2Result>("openstack:blockstorage/getSnapshotV2:getSnapshotV2", args ?? new GetSnapshotV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotV2Args : global::Pulumi.InvokeArgs
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
        /// The region in which to obtain the V2 Block Storage
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

        public GetSnapshotV2Args()
        {
        }
        public static new GetSnapshotV2Args Empty => new GetSnapshotV2Args();
    }

    public sealed class GetSnapshotV2InvokeArgs : global::Pulumi.InvokeArgs
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
        /// The region in which to obtain the V2 Block Storage
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

        public GetSnapshotV2InvokeArgs()
        {
        }
        public static new GetSnapshotV2InvokeArgs Empty => new GetSnapshotV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotV2Result
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
        private GetSnapshotV2Result(
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
