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
    /// Manages a V2 portforwarding resource within OpenStack.
    /// 
    /// ## Example Usage
    /// ### Simple portforwarding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var pf1 = new OpenStack.Networking.PortForwardingV2("pf1", new OpenStack.Networking.PortForwardingV2Args
    ///         {
    ///             ExternalPort = 7233,
    ///             FloatingipId = "7a52eb59-7d47-415d-a884-046666a6fbae",
    ///             InternalPort = 25,
    ///             InternalPortId = "b930d7f6-ceb7-40a0-8b81-a425dd994ccf",
    ///             Protocol = "tcp",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/portForwardingV2:PortForwardingV2")]
    public partial class PortForwardingV2 : Pulumi.CustomResource
    {
        /// <summary>
        /// A text describing the port forwarding. Changing this
        /// updates the `description` of an existing port forwarding.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The TCP/UDP/other protocol port number of the port forwarding. Changing this
        /// updates the `external_port` of an existing port forwarding.
        /// </summary>
        [Output("externalPort")]
        public Output<int> ExternalPort { get; private set; } = null!;

        /// <summary>
        /// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        /// </summary>
        [Output("floatingipId")]
        public Output<string> FloatingipId { get; private set; } = null!;

        /// <summary>
        /// The fixed IPv4 address of the Neutron port associated with the port forwarding.
        /// Changing this updates the `internal_ip_address` of an existing port forwarding.
        /// </summary>
        [Output("internalIpAddress")]
        public Output<string> InternalIpAddress { get; private set; } = null!;

        /// <summary>
        /// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
        /// port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        /// </summary>
        [Output("internalPort")]
        public Output<int> InternalPort { get; private set; } = null!;

        /// <summary>
        /// The ID of the Neutron port associated with the port forwarding. Changing
        /// this updates the `internal_port_id` of an existing port forwarding.
        /// </summary>
        [Output("internalPortId")]
        public Output<string> InternalPortId { get; private set; } = null!;

        /// <summary>
        /// The IP protocol used in the port forwarding. Changing this updates the `protocol`
        /// of an existing port forwarding.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a port forwarding. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// port forwarding.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a PortForwardingV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PortForwardingV2(string name, PortForwardingV2Args args, CustomResourceOptions? options = null)
            : base("openstack:networking/portForwardingV2:PortForwardingV2", name, args ?? new PortForwardingV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private PortForwardingV2(string name, Input<string> id, PortForwardingV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/portForwardingV2:PortForwardingV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PortForwardingV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PortForwardingV2 Get(string name, Input<string> id, PortForwardingV2State? state = null, CustomResourceOptions? options = null)
        {
            return new PortForwardingV2(name, id, state, options);
        }
    }

    public sealed class PortForwardingV2Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A text describing the port forwarding. Changing this
        /// updates the `description` of an existing port forwarding.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The TCP/UDP/other protocol port number of the port forwarding. Changing this
        /// updates the `external_port` of an existing port forwarding.
        /// </summary>
        [Input("externalPort", required: true)]
        public Input<int> ExternalPort { get; set; } = null!;

        /// <summary>
        /// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        /// </summary>
        [Input("floatingipId", required: true)]
        public Input<string> FloatingipId { get; set; } = null!;

        /// <summary>
        /// The fixed IPv4 address of the Neutron port associated with the port forwarding.
        /// Changing this updates the `internal_ip_address` of an existing port forwarding.
        /// </summary>
        [Input("internalIpAddress", required: true)]
        public Input<string> InternalIpAddress { get; set; } = null!;

        /// <summary>
        /// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
        /// port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        /// </summary>
        [Input("internalPort", required: true)]
        public Input<int> InternalPort { get; set; } = null!;

        /// <summary>
        /// The ID of the Neutron port associated with the port forwarding. Changing
        /// this updates the `internal_port_id` of an existing port forwarding.
        /// </summary>
        [Input("internalPortId", required: true)]
        public Input<string> InternalPortId { get; set; } = null!;

        /// <summary>
        /// The IP protocol used in the port forwarding. Changing this updates the `protocol`
        /// of an existing port forwarding.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a port forwarding. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// port forwarding.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public PortForwardingV2Args()
        {
        }
    }

    public sealed class PortForwardingV2State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A text describing the port forwarding. Changing this
        /// updates the `description` of an existing port forwarding.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The TCP/UDP/other protocol port number of the port forwarding. Changing this
        /// updates the `external_port` of an existing port forwarding.
        /// </summary>
        [Input("externalPort")]
        public Input<int>? ExternalPort { get; set; }

        /// <summary>
        /// The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        /// </summary>
        [Input("floatingipId")]
        public Input<string>? FloatingipId { get; set; }

        /// <summary>
        /// The fixed IPv4 address of the Neutron port associated with the port forwarding.
        /// Changing this updates the `internal_ip_address` of an existing port forwarding.
        /// </summary>
        [Input("internalIpAddress")]
        public Input<string>? InternalIpAddress { get; set; }

        /// <summary>
        /// The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
        /// port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        /// </summary>
        [Input("internalPort")]
        public Input<int>? InternalPort { get; set; }

        /// <summary>
        /// The ID of the Neutron port associated with the port forwarding. Changing
        /// this updates the `internal_port_id` of an existing port forwarding.
        /// </summary>
        [Input("internalPortId")]
        public Input<string>? InternalPortId { get; set; }

        /// <summary>
        /// The IP protocol used in the port forwarding. Changing this updates the `protocol`
        /// of an existing port forwarding.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a port forwarding. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// port forwarding.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public PortForwardingV2State()
        {
        }
    }
}