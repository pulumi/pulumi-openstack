// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    /// <summary>
    /// Associates a floating IP to a port. This is useful for situations
    /// where you have a pre-allocated floating IP or are unable to use the
    /// `openstack.networking.FloatingIp` resource to create a floating IP.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var port1 = new OpenStack.Networking.Port("port_1", new()
    ///     {
    ///         NetworkId = "a5bbd213-e1d3-49b6-aed1-9df60ea94b9a",
    ///     });
    /// 
    ///     var fip1 = new OpenStack.Networking.FloatingIpAssociate("fip_1", new()
    ///     {
    ///         FloatingIp = "1.2.3.4",
    ///         PortId = port1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Floating IP associations can be imported using the `id` of the floating IP, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:networking/floatingIpAssociate:FloatingIpAssociate fip 2c7f39f3-702b-48d1-940c-b50384177ee1
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/floatingIpAssociate:FloatingIpAssociate")]
    public partial class FloatingIpAssociate : global::Pulumi.CustomResource
    {
        [Output("fixedIp")]
        public Output<string> FixedIp { get; private set; } = null!;

        /// <summary>
        /// IP Address of an existing floating IP.
        /// </summary>
        [Output("floatingIp")]
        public Output<string> FloatingIp { get; private set; } = null!;

        /// <summary>
        /// ID of an existing port with at least one IP address to
        /// associate with this floating IP.
        /// </summary>
        [Output("portId")]
        public Output<string> PortId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a floating IP that can be used with
        /// another networking resource, such as a load balancer. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// floating IP (which may or may not have a different address).
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a FloatingIpAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FloatingIpAssociate(string name, FloatingIpAssociateArgs args, CustomResourceOptions? options = null)
            : base("openstack:networking/floatingIpAssociate:FloatingIpAssociate", name, args ?? new FloatingIpAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FloatingIpAssociate(string name, Input<string> id, FloatingIpAssociateState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/floatingIpAssociate:FloatingIpAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FloatingIpAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FloatingIpAssociate Get(string name, Input<string> id, FloatingIpAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new FloatingIpAssociate(name, id, state, options);
        }
    }

    public sealed class FloatingIpAssociateArgs : global::Pulumi.ResourceArgs
    {
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// IP Address of an existing floating IP.
        /// </summary>
        [Input("floatingIp", required: true)]
        public Input<string> FloatingIp { get; set; } = null!;

        /// <summary>
        /// ID of an existing port with at least one IP address to
        /// associate with this floating IP.
        /// </summary>
        [Input("portId", required: true)]
        public Input<string> PortId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a floating IP that can be used with
        /// another networking resource, such as a load balancer. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// floating IP (which may or may not have a different address).
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public FloatingIpAssociateArgs()
        {
        }
        public static new FloatingIpAssociateArgs Empty => new FloatingIpAssociateArgs();
    }

    public sealed class FloatingIpAssociateState : global::Pulumi.ResourceArgs
    {
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// IP Address of an existing floating IP.
        /// </summary>
        [Input("floatingIp")]
        public Input<string>? FloatingIp { get; set; }

        /// <summary>
        /// ID of an existing port with at least one IP address to
        /// associate with this floating IP.
        /// </summary>
        [Input("portId")]
        public Input<string>? PortId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a floating IP that can be used with
        /// another networking resource, such as a load balancer. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// floating IP (which may or may not have a different address).
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public FloatingIpAssociateState()
        {
        }
        public static new FloatingIpAssociateState Empty => new FloatingIpAssociateState();
    }
}
