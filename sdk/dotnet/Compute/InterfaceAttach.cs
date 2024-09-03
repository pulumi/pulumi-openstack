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
    /// Attaches a Network Interface (a Port) to an Instance using the OpenStack
    /// Compute (Nova) v2 API.
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic Attachment
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network1 = new OpenStack.Networking.Network("network_1", new()
    ///     {
    ///         Name = "network_1",
    ///         AdminStateUp = true,
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
    ///     var ai1 = new OpenStack.Compute.InterfaceAttach("ai_1", new()
    ///     {
    ///         InstanceId = instance1.Id,
    ///         NetworkId = network1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Attachment Specifying a Fixed IP
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network1 = new OpenStack.Networking.Network("network_1", new()
    ///     {
    ///         Name = "network_1",
    ///         AdminStateUp = true,
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
    ///     var ai1 = new OpenStack.Compute.InterfaceAttach("ai_1", new()
    ///     {
    ///         InstanceId = instance1.Id,
    ///         NetworkId = network1OpenstackNetworkingPortV2.Id,
    ///         FixedIp = "10.0.10.10",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Attachment Using an Existing Port
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network1 = new OpenStack.Networking.Network("network_1", new()
    ///     {
    ///         Name = "network_1",
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var port1 = new OpenStack.Networking.Port("port_1", new()
    ///     {
    ///         Name = "port_1",
    ///         NetworkId = network1.Id,
    ///         AdminStateUp = true,
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
    ///     var ai1 = new OpenStack.Compute.InterfaceAttach("ai_1", new()
    ///     {
    ///         InstanceId = instance1.Id,
    ///         PortId = port1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Interface Attachments can be imported using the Instance ID and Port ID
    /// separated by a slash, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:compute/interfaceAttach:InterfaceAttach ai_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:compute/interfaceAttach:InterfaceAttach")]
    public partial class InterfaceAttach : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An IP address to assosciate with the port.
        /// *NOTE*: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        /// </summary>
        [Output("fixedIp")]
        public Output<string> FixedIp { get; private set; } = null!;

        /// <summary>
        /// The ID of the Instance to attach the Port or Network to.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Network to attach to an Instance. A port will be created automatically.
        /// *NOTE*: This option and `port_id` are mutually exclusive.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// The ID of the Port to attach to an Instance.
        /// *NOTE*: This option and `network_id` are mutually exclusive.
        /// </summary>
        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the interface attachment.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new attachment.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a InterfaceAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InterfaceAttach(string name, InterfaceAttachArgs args, CustomResourceOptions? options = null)
            : base("openstack:compute/interfaceAttach:InterfaceAttach", name, args ?? new InterfaceAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InterfaceAttach(string name, Input<string> id, InterfaceAttachState? state = null, CustomResourceOptions? options = null)
            : base("openstack:compute/interfaceAttach:InterfaceAttach", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InterfaceAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InterfaceAttach Get(string name, Input<string> id, InterfaceAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new InterfaceAttach(name, id, state, options);
        }
    }

    public sealed class InterfaceAttachArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An IP address to assosciate with the port.
        /// *NOTE*: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        /// </summary>
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// The ID of the Instance to attach the Port or Network to.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The ID of the Network to attach to an Instance. A port will be created automatically.
        /// *NOTE*: This option and `port_id` are mutually exclusive.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// The ID of the Port to attach to an Instance.
        /// *NOTE*: This option and `network_id` are mutually exclusive.
        /// </summary>
        [Input("portId")]
        public Input<string>? PortId { get; set; }

        /// <summary>
        /// The region in which to create the interface attachment.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new attachment.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public InterfaceAttachArgs()
        {
        }
        public static new InterfaceAttachArgs Empty => new InterfaceAttachArgs();
    }

    public sealed class InterfaceAttachState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An IP address to assosciate with the port.
        /// *NOTE*: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        /// </summary>
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// The ID of the Instance to attach the Port or Network to.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ID of the Network to attach to an Instance. A port will be created automatically.
        /// *NOTE*: This option and `port_id` are mutually exclusive.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// The ID of the Port to attach to an Instance.
        /// *NOTE*: This option and `network_id` are mutually exclusive.
        /// </summary>
        [Input("portId")]
        public Input<string>? PortId { get; set; }

        /// <summary>
        /// The region in which to create the interface attachment.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new attachment.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public InterfaceAttachState()
        {
        }
        public static new InterfaceAttachState Empty => new InterfaceAttachState();
    }
}
