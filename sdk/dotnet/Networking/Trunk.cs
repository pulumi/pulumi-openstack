// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    /// <summary>
    /// Manages a networking V2 trunk resource within OpenStack.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_trunk_v2.html.markdown.
    /// </summary>
    public partial class Trunk : Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative up/down status for the trunk
        /// (must be "true" or "false" if provided). Changing this updates the
        /// `admin_state_up` of an existing trunk.
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The collection of tags assigned on the trunk, which have been
        /// explicitly and implicitly added.
        /// </summary>
        [Output("allTags")]
        public Output<ImmutableArray<string>> AllTags { get; private set; } = null!;

        /// <summary>
        /// Human-readable description of the trunk. Changing this
        /// updates the name of the existing trunk.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A unique name for the trunk. Changing this
        /// updates the `name` of an existing trunk.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the port to be made a subport of the trunk.
        /// </summary>
        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a trunk. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// trunk.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The set of ports that will be made subports of the trunk.
        /// The structure of each subport is described below.
        /// </summary>
        [Output("subPorts")]
        public Output<ImmutableArray<Outputs.TrunkSubPorts>> SubPorts { get; private set; } = null!;

        /// <summary>
        /// A set of string tags for the port.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The owner of the Trunk. Required if admin wants
        /// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a Trunk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trunk(string name, TrunkArgs args, CustomResourceOptions? options = null)
            : base("openstack:networking/trunk:Trunk", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Trunk(string name, Input<string> id, TrunkState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/trunk:Trunk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trunk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trunk Get(string name, Input<string> id, TrunkState? state = null, CustomResourceOptions? options = null)
        {
            return new Trunk(name, id, state, options);
        }
    }

    public sealed class TrunkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative up/down status for the trunk
        /// (must be "true" or "false" if provided). Changing this updates the
        /// `admin_state_up` of an existing trunk.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// Human-readable description of the trunk. Changing this
        /// updates the name of the existing trunk.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique name for the trunk. Changing this
        /// updates the `name` of an existing trunk.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the port to be made a subport of the trunk.
        /// </summary>
        [Input("portId", required: true)]
        public Input<string> PortId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a trunk. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// trunk.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("subPorts")]
        private InputList<Inputs.TrunkSubPortsArgs>? _subPorts;

        /// <summary>
        /// The set of ports that will be made subports of the trunk.
        /// The structure of each subport is described below.
        /// </summary>
        public InputList<Inputs.TrunkSubPortsArgs> SubPorts
        {
            get => _subPorts ?? (_subPorts = new InputList<Inputs.TrunkSubPortsArgs>());
            set => _subPorts = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the port.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the Trunk. Required if admin wants
        /// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public TrunkArgs()
        {
        }
    }

    public sealed class TrunkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative up/down status for the trunk
        /// (must be "true" or "false" if provided). Changing this updates the
        /// `admin_state_up` of an existing trunk.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("allTags")]
        private InputList<string>? _allTags;

        /// <summary>
        /// The collection of tags assigned on the trunk, which have been
        /// explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllTags
        {
            get => _allTags ?? (_allTags = new InputList<string>());
            set => _allTags = value;
        }

        /// <summary>
        /// Human-readable description of the trunk. Changing this
        /// updates the name of the existing trunk.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique name for the trunk. Changing this
        /// updates the `name` of an existing trunk.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the port to be made a subport of the trunk.
        /// </summary>
        [Input("portId")]
        public Input<string>? PortId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a trunk. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// trunk.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("subPorts")]
        private InputList<Inputs.TrunkSubPortsGetArgs>? _subPorts;

        /// <summary>
        /// The set of ports that will be made subports of the trunk.
        /// The structure of each subport is described below.
        /// </summary>
        public InputList<Inputs.TrunkSubPortsGetArgs> SubPorts
        {
            get => _subPorts ?? (_subPorts = new InputList<Inputs.TrunkSubPortsGetArgs>());
            set => _subPorts = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the port.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the Trunk. Required if admin wants
        /// to create a trunk on behalf of another tenant. Changing this creates a new trunk.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public TrunkState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class TrunkSubPortsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the port to be made a subport of the trunk.
        /// </summary>
        [Input("portId", required: true)]
        public Input<string> PortId { get; set; } = null!;

        /// <summary>
        /// The numeric id of the subport segment.
        /// </summary>
        [Input("segmentationId", required: true)]
        public Input<int> SegmentationId { get; set; } = null!;

        /// <summary>
        /// The segmentation technology to use, e.g., "vlan".
        /// </summary>
        [Input("segmentationType", required: true)]
        public Input<string> SegmentationType { get; set; } = null!;

        public TrunkSubPortsArgs()
        {
        }
    }

    public sealed class TrunkSubPortsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the port to be made a subport of the trunk.
        /// </summary>
        [Input("portId", required: true)]
        public Input<string> PortId { get; set; } = null!;

        /// <summary>
        /// The numeric id of the subport segment.
        /// </summary>
        [Input("segmentationId", required: true)]
        public Input<int> SegmentationId { get; set; } = null!;

        /// <summary>
        /// The segmentation technology to use, e.g., "vlan".
        /// </summary>
        [Input("segmentationType", required: true)]
        public Input<string> SegmentationType { get; set; } = null!;

        public TrunkSubPortsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class TrunkSubPorts
    {
        /// <summary>
        /// The ID of the port to be made a subport of the trunk.
        /// </summary>
        public readonly string PortId;
        /// <summary>
        /// The numeric id of the subport segment.
        /// </summary>
        public readonly int SegmentationId;
        /// <summary>
        /// The segmentation technology to use, e.g., "vlan".
        /// </summary>
        public readonly string SegmentationType;

        [OutputConstructor]
        private TrunkSubPorts(
            string portId,
            int segmentationId,
            string segmentationType)
        {
            PortId = portId;
            SegmentationId = segmentationId;
            SegmentationType = segmentationType;
        }
    }
    }
}
