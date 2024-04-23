// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    /// <summary>
    /// Attaches a Block Storage Volume to an Instance using the OpenStack
    /// Compute (Nova) v2 API.
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic attachment of a single volume to a single instance
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var volume1 = new OpenStack.BlockStorage.Volume("volume_1", new()
    ///     {
    ///         Name = "volume_1",
    ///         Size = 1,
    ///     });
    /// 
    ///     var instance1 = new OpenStack.Compute.Instance("instance_1", new()
    ///     {
    ///         Name = "instance_1",
    ///         SecurityGroups = new[]
    ///         {
    ///             "default",
    ///         },
    ///     });
    /// 
    ///     var va1 = new OpenStack.Compute.VolumeAttach("va_1", new()
    ///     {
    ///         InstanceId = instance1.Id,
    ///         VolumeId = volume1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Using Multiattach-enabled volumes
    /// 
    /// Multiattach Volumes are dependent upon your OpenStack cloud and not all
    /// clouds support multiattach.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var volume1 = new OpenStack.BlockStorage.Volume("volume_1", new()
    ///     {
    ///         Name = "volume_1",
    ///         Size = 1,
    ///         Multiattach = true,
    ///     });
    /// 
    ///     var instance1 = new OpenStack.Compute.Instance("instance_1", new()
    ///     {
    ///         Name = "instance_1",
    ///         SecurityGroups = new[]
    ///         {
    ///             "default",
    ///         },
    ///     });
    /// 
    ///     var instance2 = new OpenStack.Compute.Instance("instance_2", new()
    ///     {
    ///         Name = "instance_2",
    ///         SecurityGroups = new[]
    ///         {
    ///             "default",
    ///         },
    ///     });
    /// 
    ///     var va1 = new OpenStack.Compute.VolumeAttach("va_1", new()
    ///     {
    ///         InstanceId = instance1.Id,
    ///         VolumeId = volume1.Id,
    ///         Multiattach = true,
    ///     });
    /// 
    ///     var va2 = new OpenStack.Compute.VolumeAttach("va_2", new()
    ///     {
    ///         InstanceId = instance2.Id,
    ///         VolumeId = volume1.Id,
    ///         Multiattach = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             va1,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// It is recommended to use `depends_on` for the attach resources
    /// to enforce the volume attachments to happen one at a time.
    /// 
    /// ## Import
    /// 
    /// Volume Attachments can be imported using the Instance ID and Volume ID
    /// separated by a slash, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:compute/volumeAttach:VolumeAttach va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:compute/volumeAttach:VolumeAttach")]
    public partial class VolumeAttach : global::Pulumi.CustomResource
    {
        [Output("device")]
        public Output<string> Device { get; private set; } = null!;

        /// <summary>
        /// The ID of the Instance to attach the Volume to.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Enable attachment of multiattach-capable volumes.
        /// </summary>
        [Output("multiattach")]
        public Output<bool?> Multiattach { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// A Compute client is needed to create a volume attachment. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a
        /// new volume attachment.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Map of additional vendor-specific options.
        /// Supported options are described below.
        /// </summary>
        [Output("vendorOptions")]
        public Output<Outputs.VolumeAttachVendorOptions?> VendorOptions { get; private set; } = null!;

        /// <summary>
        /// The ID of the Volume to attach to an Instance.
        /// </summary>
        [Output("volumeId")]
        public Output<string> VolumeId { get; private set; } = null!;


        /// <summary>
        /// Create a VolumeAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VolumeAttach(string name, VolumeAttachArgs args, CustomResourceOptions? options = null)
            : base("openstack:compute/volumeAttach:VolumeAttach", name, args ?? new VolumeAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VolumeAttach(string name, Input<string> id, VolumeAttachState? state = null, CustomResourceOptions? options = null)
            : base("openstack:compute/volumeAttach:VolumeAttach", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VolumeAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VolumeAttach Get(string name, Input<string> id, VolumeAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new VolumeAttach(name, id, state, options);
        }
    }

    public sealed class VolumeAttachArgs : global::Pulumi.ResourceArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// The ID of the Instance to attach the Volume to.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Enable attachment of multiattach-capable volumes.
        /// </summary>
        [Input("multiattach")]
        public Input<bool>? Multiattach { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// A Compute client is needed to create a volume attachment. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a
        /// new volume attachment.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Map of additional vendor-specific options.
        /// Supported options are described below.
        /// </summary>
        [Input("vendorOptions")]
        public Input<Inputs.VolumeAttachVendorOptionsArgs>? VendorOptions { get; set; }

        /// <summary>
        /// The ID of the Volume to attach to an Instance.
        /// </summary>
        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public VolumeAttachArgs()
        {
        }
        public static new VolumeAttachArgs Empty => new VolumeAttachArgs();
    }

    public sealed class VolumeAttachState : global::Pulumi.ResourceArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// The ID of the Instance to attach the Volume to.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Enable attachment of multiattach-capable volumes.
        /// </summary>
        [Input("multiattach")]
        public Input<bool>? Multiattach { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// A Compute client is needed to create a volume attachment. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a
        /// new volume attachment.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Map of additional vendor-specific options.
        /// Supported options are described below.
        /// </summary>
        [Input("vendorOptions")]
        public Input<Inputs.VolumeAttachVendorOptionsGetArgs>? VendorOptions { get; set; }

        /// <summary>
        /// The ID of the Volume to attach to an Instance.
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        public VolumeAttachState()
        {
        }
        public static new VolumeAttachState Empty => new VolumeAttachState();
    }
}
