// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get information about an existing volume.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/blockstorage_volume_v2.html.markdown.
        /// </summary>
        [Obsolete("Use GetVolumeV2.InvokeAsync() instead")]
        public static Task<GetVolumeV2Result> GetVolumeV2(GetVolumeV2Args? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeV2Result>("openstack:blockstorage/getVolumeV2:getVolumeV2", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetVolumeV2
    {
        /// <summary>
        /// Use this data source to get information about an existing volume.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/blockstorage_volume_v2.html.markdown.
        /// </summary>
        public static Task<GetVolumeV2Result> InvokeAsync(GetVolumeV2Args? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeV2Result>("openstack:blockstorage/getVolumeV2:getVolumeV2", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetVolumeV2Args : Pulumi.InvokeArgs
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
    }

    [OutputType]
    public sealed class GetVolumeV2Result
    {
        /// <summary>
        /// Indicates if the volume is bootable.
        /// </summary>
        public readonly string Bootable;
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
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetVolumeV2Result(
            string bootable,
            ImmutableDictionary<string, object> metadata,
            string name,
            string region,
            int size,
            string sourceVolumeId,
            string status,
            string volumeType,
            string id)
        {
            Bootable = bootable;
            Metadata = metadata;
            Name = name;
            Region = region;
            Size = size;
            SourceVolumeId = sourceVolumeId;
            Status = status;
            VolumeType = volumeType;
            Id = id;
        }
    }
}
