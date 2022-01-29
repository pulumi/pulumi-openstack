// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage
{
    /// <summary>
    /// Manages a V3 block storage Qos Association resource within OpenStack.
    /// 
    /// &gt; **Note:** This usually requires admin privileges.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var qos = new OpenStack.BlockStorage.QosV3("qos", new OpenStack.BlockStorage.QosV3Args
    ///         {
    ///             Consumer = "front-end",
    ///             Specs = 
    ///             {
    ///                 { "read_iops_sec", "20000" },
    ///             },
    ///         });
    ///         var volumeType = new OpenStack.BlockStorage.VolumeTypeV3("volumeType", new OpenStack.BlockStorage.VolumeTypeV3Args
    ///         {
    ///         });
    ///         var qosAssociation = new OpenStack.BlockStorage.QosAssociationV3("qosAssociation", new OpenStack.BlockStorage.QosAssociationV3Args
    ///         {
    ///             QosId = qos.Id,
    ///             VolumeTypeId = volumeType.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Qos association can be imported using the `qos_id/volume_type_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:blockstorage/qosAssociationV3:QosAssociationV3 qos_association 941793f0-0a34-4bc4-b72e-a6326ae58283/ea257959-eeb1-4c10-8d33-26f0409a755d
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:blockstorage/qosAssociationV3:QosAssociationV3")]
    public partial class QosAssociationV3 : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the qos to associate. Changing this creates
        /// a new qos association.
        /// </summary>
        [Output("qosId")]
        public Output<string> QosId { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the qos association.
        /// If omitted, the `region` argument of the provider is used. Changing
        /// this creates a new qos association.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// ID of the volume_type to associate.
        /// Changing this creates a new qos association.
        /// </summary>
        [Output("volumeTypeId")]
        public Output<string> VolumeTypeId { get; private set; } = null!;


        /// <summary>
        /// Create a QosAssociationV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QosAssociationV3(string name, QosAssociationV3Args args, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/qosAssociationV3:QosAssociationV3", name, args ?? new QosAssociationV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private QosAssociationV3(string name, Input<string> id, QosAssociationV3State? state = null, CustomResourceOptions? options = null)
            : base("openstack:blockstorage/qosAssociationV3:QosAssociationV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QosAssociationV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QosAssociationV3 Get(string name, Input<string> id, QosAssociationV3State? state = null, CustomResourceOptions? options = null)
        {
            return new QosAssociationV3(name, id, state, options);
        }
    }

    public sealed class QosAssociationV3Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the qos to associate. Changing this creates
        /// a new qos association.
        /// </summary>
        [Input("qosId", required: true)]
        public Input<string> QosId { get; set; } = null!;

        /// <summary>
        /// The region in which to create the qos association.
        /// If omitted, the `region` argument of the provider is used. Changing
        /// this creates a new qos association.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// ID of the volume_type to associate.
        /// Changing this creates a new qos association.
        /// </summary>
        [Input("volumeTypeId", required: true)]
        public Input<string> VolumeTypeId { get; set; } = null!;

        public QosAssociationV3Args()
        {
        }
    }

    public sealed class QosAssociationV3State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the qos to associate. Changing this creates
        /// a new qos association.
        /// </summary>
        [Input("qosId")]
        public Input<string>? QosId { get; set; }

        /// <summary>
        /// The region in which to create the qos association.
        /// If omitted, the `region` argument of the provider is used. Changing
        /// this creates a new qos association.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// ID of the volume_type to associate.
        /// Changing this creates a new qos association.
        /// </summary>
        [Input("volumeTypeId")]
        public Input<string>? VolumeTypeId { get; set; }

        public QosAssociationV3State()
        {
        }
    }
}