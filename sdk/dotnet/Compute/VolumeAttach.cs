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
    /// ### Basic attachment of a single volume to a single instance
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var volume1 = new OpenStack.BlockStorage.VolumeV2("volume1", new OpenStack.BlockStorage.VolumeV2Args
    ///         {
    ///             Size = 1,
    ///         });
    ///         var instance1 = new OpenStack.Compute.Instance("instance1", new OpenStack.Compute.InstanceArgs
    ///         {
    ///             SecurityGroups = 
    ///             {
    ///                 "default",
    ///             },
    ///         });
    ///         var va1 = new OpenStack.Compute.VolumeAttach("va1", new OpenStack.Compute.VolumeAttachArgs
    ///         {
    ///             InstanceId = instance1.Id,
    ///             VolumeId = volume1.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Using Multiattach-enabled volumes
    /// 
    /// Multiattach Volumes are dependent upon your OpenStack cloud and not all
    /// clouds support multiattach.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var volume1 = new OpenStack.BlockStorage.Volume("volume1", new OpenStack.BlockStorage.VolumeArgs
    ///         {
    ///             Multiattach = true,
    ///             Size = 1,
    ///         });
    ///         var instance1 = new OpenStack.Compute.Instance("instance1", new OpenStack.Compute.InstanceArgs
    ///         {
    ///             SecurityGroups = 
    ///             {
    ///                 "default",
    ///             },
    ///         });
    ///         var instance2 = new OpenStack.Compute.Instance("instance2", new OpenStack.Compute.InstanceArgs
    ///         {
    ///             SecurityGroups = 
    ///             {
    ///                 "default",
    ///             },
    ///         });
    ///         var va1 = new OpenStack.Compute.VolumeAttach("va1", new OpenStack.Compute.VolumeAttachArgs
    ///         {
    ///             InstanceId = instance1.Id,
    ///             Multiattach = true,
    ///             VolumeId = openstack_blockstorage_volume_v2.Volume_1.Id,
    ///         });
    ///         var va2 = new OpenStack.Compute.VolumeAttach("va2", new OpenStack.Compute.VolumeAttachArgs
    ///         {
    ///             InstanceId = instance2.Id,
    ///             Multiattach = true,
    ///             VolumeId = openstack_blockstorage_volume_v2.Volume_1.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 "openstack_compute_volume_attach_v2.va_1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// It is recommended to use `depends_on` for the attach resources
    /// to enforce the volume attachments to happen one at a time.
    /// </summary>
    public partial class VolumeAttach : Pulumi.CustomResource
    {
        /// <summary>
        /// See Argument Reference above. _NOTE_: The correctness of this
        /// information is dependent upon the hypervisor in use. In some cases, this
        /// should not be used as an authoritative piece of information.
        /// </summary>
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

    public sealed class VolumeAttachArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// See Argument Reference above. _NOTE_: The correctness of this
        /// information is dependent upon the hypervisor in use. In some cases, this
        /// should not be used as an authoritative piece of information.
        /// </summary>
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
        /// The ID of the Volume to attach to an Instance.
        /// </summary>
        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public VolumeAttachArgs()
        {
        }
    }

    public sealed class VolumeAttachState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// See Argument Reference above. _NOTE_: The correctness of this
        /// information is dependent upon the hypervisor in use. In some cases, this
        /// should not be used as an authoritative piece of information.
        /// </summary>
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
        /// The ID of the Volume to attach to an Instance.
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        public VolumeAttachState()
        {
        }
    }
}
