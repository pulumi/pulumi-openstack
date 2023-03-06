// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage
{
    public static class GetVolumeV2
    {
        /// <summary>
        /// Use this data source to get information about an existing volume.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var volume1 = OpenStack.BlockStorage.GetVolumeV2.Invoke(new()
        ///     {
        ///         Name = "volume_1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVolumeV2Result> InvokeAsync(GetVolumeV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVolumeV2Result>("openstack:blockstorage/getVolumeV2:getVolumeV2", args ?? new GetVolumeV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about an existing volume.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var volume1 = OpenStack.BlockStorage.GetVolumeV2.Invoke(new()
        ///     {
        ///         Name = "volume_1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVolumeV2Result> Invoke(GetVolumeV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumeV2Result>("openstack:blockstorage/getVolumeV2:getVolumeV2", args ?? new GetVolumeV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumeV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Indicates if the volume is bootable.
        /// </summary>
        [Input("bootable")]
        public string? Bootable { get; set; }

        [Input("metadata")]
        private Dictionary<string, object>? _metadata;

        /// <summary>
        /// Metadata key/value pairs associated with the volume.
        /// </summary>
        public Dictionary<string, object> Metadata
        {
            get => _metadata ?? (_metadata = new Dictionary<string, object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the volume.
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
        /// The status of the volume.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The type of the volume.
        /// </summary>
        [Input("volumeType")]
        public string? VolumeType { get; set; }

        public GetVolumeV2Args()
        {
        }
        public static new GetVolumeV2Args Empty => new GetVolumeV2Args();
    }

    public sealed class GetVolumeV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Indicates if the volume is bootable.
        /// </summary>
        [Input("bootable")]
        public Input<string>? Bootable { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Metadata key/value pairs associated with the volume.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// The name of the volume.
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
        /// The status of the volume.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The type of the volume.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public GetVolumeV2InvokeArgs()
        {
        }
        public static new GetVolumeV2InvokeArgs Empty => new GetVolumeV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetVolumeV2Result
    {
        /// <summary>
        /// Indicates if the volume is bootable.
        /// </summary>
        public readonly string Bootable;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Metadata;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The size of the volume in GBs.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The ID of the volume from which the current volume was created.
        /// </summary>
        public readonly string SourceVolumeId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of the volume.
        /// </summary>
        public readonly string VolumeType;

        [OutputConstructor]
        private GetVolumeV2Result(
            string bootable,

            string id,

            ImmutableDictionary<string, object> metadata,

            string name,

            string region,

            int size,

            string sourceVolumeId,

            string status,

            string volumeType)
        {
            Bootable = bootable;
            Id = id;
            Metadata = metadata;
            Name = name;
            Region = region;
            Size = size;
            SourceVolumeId = sourceVolumeId;
            Status = status;
            VolumeType = volumeType;
        }
    }
}
