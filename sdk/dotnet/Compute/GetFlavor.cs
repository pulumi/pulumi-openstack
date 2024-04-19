// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    public static class GetFlavor
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack flavor.
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
        ///     var small = OpenStack.Compute.GetFlavor.Invoke(new()
        ///     {
        ///         Vcpus = 1,
        ///         Ram = 512,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetFlavorResult> InvokeAsync(GetFlavorArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFlavorResult>("openstack:compute/getFlavor:getFlavor", args ?? new GetFlavorArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack flavor.
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
        ///     var small = OpenStack.Compute.GetFlavor.Invoke(new()
        ///     {
        ///         Vcpus = 1,
        ///         Ram = 512,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetFlavorResult> Invoke(GetFlavorInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFlavorResult>("openstack:compute/getFlavor:getFlavor", args ?? new GetFlavorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFlavorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of the flavor.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The exact amount of disk (in gigabytes).
        /// </summary>
        [Input("disk")]
        public int? Disk { get; set; }

        /// <summary>
        /// The ID of the flavor. Conflicts with the `name`,
        /// `min_ram` and `min_disk`
        /// </summary>
        [Input("flavorId")]
        public string? FlavorId { get; set; }

        /// <summary>
        /// The flavor visibility.
        /// </summary>
        [Input("isPublic")]
        public bool? IsPublic { get; set; }

        /// <summary>
        /// The minimum amount of disk (in gigabytes). Conflicts
        /// with the `flavor_id`.
        /// </summary>
        [Input("minDisk")]
        public int? MinDisk { get; set; }

        /// <summary>
        /// The minimum amount of RAM (in megabytes). Conflicts
        /// with the `flavor_id`.
        /// </summary>
        [Input("minRam")]
        public int? MinRam { get; set; }

        /// <summary>
        /// The name of the flavor. Conflicts with the `flavor_id`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The exact amount of RAM (in megabytes).
        /// </summary>
        [Input("ram")]
        public int? Ram { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The `rx_tx_factor` of the flavor.
        /// </summary>
        [Input("rxTxFactor")]
        public double? RxTxFactor { get; set; }

        /// <summary>
        /// The amount of swap (in gigabytes).
        /// </summary>
        [Input("swap")]
        public int? Swap { get; set; }

        /// <summary>
        /// The amount of VCPUs.
        /// </summary>
        [Input("vcpus")]
        public int? Vcpus { get; set; }

        public GetFlavorArgs()
        {
        }
        public static new GetFlavorArgs Empty => new GetFlavorArgs();
    }

    public sealed class GetFlavorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The description of the flavor.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The exact amount of disk (in gigabytes).
        /// </summary>
        [Input("disk")]
        public Input<int>? Disk { get; set; }

        /// <summary>
        /// The ID of the flavor. Conflicts with the `name`,
        /// `min_ram` and `min_disk`
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// The flavor visibility.
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// The minimum amount of disk (in gigabytes). Conflicts
        /// with the `flavor_id`.
        /// </summary>
        [Input("minDisk")]
        public Input<int>? MinDisk { get; set; }

        /// <summary>
        /// The minimum amount of RAM (in megabytes). Conflicts
        /// with the `flavor_id`.
        /// </summary>
        [Input("minRam")]
        public Input<int>? MinRam { get; set; }

        /// <summary>
        /// The name of the flavor. Conflicts with the `flavor_id`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The exact amount of RAM (in megabytes).
        /// </summary>
        [Input("ram")]
        public Input<int>? Ram { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The `rx_tx_factor` of the flavor.
        /// </summary>
        [Input("rxTxFactor")]
        public Input<double>? RxTxFactor { get; set; }

        /// <summary>
        /// The amount of swap (in gigabytes).
        /// </summary>
        [Input("swap")]
        public Input<int>? Swap { get; set; }

        /// <summary>
        /// The amount of VCPUs.
        /// </summary>
        [Input("vcpus")]
        public Input<int>? Vcpus { get; set; }

        public GetFlavorInvokeArgs()
        {
        }
        public static new GetFlavorInvokeArgs Empty => new GetFlavorInvokeArgs();
    }


    [OutputType]
    public sealed class GetFlavorResult
    {
        public readonly string? Description;
        public readonly int? Disk;
        /// <summary>
        /// Key/Value pairs of metadata for the flavor.
        /// </summary>
        public readonly ImmutableDictionary<string, object> ExtraSpecs;
        public readonly string? FlavorId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IsPublic;
        public readonly int? MinDisk;
        public readonly int? MinRam;
        public readonly string? Name;
        public readonly int? Ram;
        public readonly string Region;
        public readonly double? RxTxFactor;
        public readonly int? Swap;
        public readonly int? Vcpus;

        [OutputConstructor]
        private GetFlavorResult(
            string? description,

            int? disk,

            ImmutableDictionary<string, object> extraSpecs,

            string? flavorId,

            string id,

            bool? isPublic,

            int? minDisk,

            int? minRam,

            string? name,

            int? ram,

            string region,

            double? rxTxFactor,

            int? swap,

            int? vcpus)
        {
            Description = description;
            Disk = disk;
            ExtraSpecs = extraSpecs;
            FlavorId = flavorId;
            Id = id;
            IsPublic = isPublic;
            MinDisk = minDisk;
            MinRam = minRam;
            Name = name;
            Ram = ram;
            Region = region;
            RxTxFactor = rxTxFactor;
            Swap = swap;
            Vcpus = vcpus;
        }
    }
}
