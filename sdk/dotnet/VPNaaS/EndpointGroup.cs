// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.VPNaaS
{
    /// <summary>
    /// Manages a V2 Neutron Endpoint Group resource within OpenStack.
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
    ///     var group1 = new OpenStack.VPNaaS.EndpointGroup("group_1", new()
    ///     {
    ///         Name = "Group 1",
    ///         Type = "cidr",
    ///         Endpoints = new[]
    ///         {
    ///             "10.2.0.0/24",
    ///             "10.3.0.0/24",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Groups can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:vpnaas/endpointGroup:EndpointGroup group_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:vpnaas/endpointGroup:EndpointGroup")]
    public partial class EndpointGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The human-readable description for the group.
        /// Changing this updates the description of the existing group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
        /// Changing this creates a new group.
        /// </summary>
        [Output("endpoints")]
        public Output<ImmutableArray<string>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// The name of the group. Changing this updates the name of
        /// the existing group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an endpoint group. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// group.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The owner of the group. Required if admin wants to
        /// create an endpoint group for another project. Changing this creates a new group.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
        /// Changing this creates a new group.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointGroup(string name, EndpointGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:vpnaas/endpointGroup:EndpointGroup", name, args ?? new EndpointGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointGroup(string name, Input<string> id, EndpointGroupState? state = null, CustomResourceOptions? options = null)
            : base("openstack:vpnaas/endpointGroup:EndpointGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EndpointGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointGroup Get(string name, Input<string> id, EndpointGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointGroup(name, id, state, options);
        }
    }

    public sealed class EndpointGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description for the group.
        /// Changing this updates the description of the existing group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("endpoints")]
        private InputList<string>? _endpoints;

        /// <summary>
        /// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
        /// Changing this creates a new group.
        /// </summary>
        public InputList<string> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<string>());
            set => _endpoints = value;
        }

        /// <summary>
        /// The name of the group. Changing this updates the name of
        /// the existing group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an endpoint group. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the group. Required if admin wants to
        /// create an endpoint group for another project. Changing this creates a new group.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
        /// Changing this creates a new group.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        public EndpointGroupArgs()
        {
        }
        public static new EndpointGroupArgs Empty => new EndpointGroupArgs();
    }

    public sealed class EndpointGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description for the group.
        /// Changing this updates the description of the existing group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("endpoints")]
        private InputList<string>? _endpoints;

        /// <summary>
        /// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
        /// Changing this creates a new group.
        /// </summary>
        public InputList<string> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<string>());
            set => _endpoints = value;
        }

        /// <summary>
        /// The name of the group. Changing this updates the name of
        /// the existing group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an endpoint group. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the group. Required if admin wants to
        /// create an endpoint group for another project. Changing this creates a new group.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
        /// Changing this creates a new group.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        public EndpointGroupState()
        {
        }
        public static new EndpointGroupState Empty => new EndpointGroupState();
    }
}
