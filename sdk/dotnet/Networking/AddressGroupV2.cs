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
    /// Manages a V2 neutron address group resource within OpenStack.
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
    ///     var group1 = new OpenStack.Networking.AddressGroupV2("group_1", new()
    ///     {
    ///         Name = "group_1",
    ///         Description = "My neutron address group",
    ///         Addresses = new[]
    ///         {
    ///             "192.168.0.1/32",
    ///             "2001:db8::1/128",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Address Groups can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:networking/addressGroupV2:AddressGroupV2 group_1 782fef9c-d03c-400a-9735-2f9af5681cb3
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/addressGroupV2:AddressGroupV2")]
    public partial class AddressGroupV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of CIDR blocks that define the addresses in
        /// the address group. Each address must be a valid IPv4 or IPv6 CIDR block.
        /// </summary>
        [Output("addresses")]
        public Output<ImmutableArray<string>> Addresses { get; private set; } = null!;

        /// <summary>
        /// A description of the address group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A name of the address group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The owner of the address group. Required if admin
        /// wants to create a group for a specific project. Changing this creates a new
        /// address group.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new address group.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a AddressGroupV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AddressGroupV2(string name, AddressGroupV2Args args, CustomResourceOptions? options = null)
            : base("openstack:networking/addressGroupV2:AddressGroupV2", name, args ?? new AddressGroupV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private AddressGroupV2(string name, Input<string> id, AddressGroupV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/addressGroupV2:AddressGroupV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AddressGroupV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AddressGroupV2 Get(string name, Input<string> id, AddressGroupV2State? state = null, CustomResourceOptions? options = null)
        {
            return new AddressGroupV2(name, id, state, options);
        }
    }

    public sealed class AddressGroupV2Args : global::Pulumi.ResourceArgs
    {
        [Input("addresses", required: true)]
        private InputList<string>? _addresses;

        /// <summary>
        /// A list of CIDR blocks that define the addresses in
        /// the address group. Each address must be a valid IPv4 or IPv6 CIDR block.
        /// </summary>
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        /// <summary>
        /// A description of the address group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A name of the address group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the address group. Required if admin
        /// wants to create a group for a specific project. Changing this creates a new
        /// address group.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new address group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AddressGroupV2Args()
        {
        }
        public static new AddressGroupV2Args Empty => new AddressGroupV2Args();
    }

    public sealed class AddressGroupV2State : global::Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<string>? _addresses;

        /// <summary>
        /// A list of CIDR blocks that define the addresses in
        /// the address group. Each address must be a valid IPv4 or IPv6 CIDR block.
        /// </summary>
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        /// <summary>
        /// A description of the address group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A name of the address group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the address group. Required if admin
        /// wants to create a group for a specific project. Changing this creates a new
        /// address group.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new address group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AddressGroupV2State()
        {
        }
        public static new AddressGroupV2State Empty => new AddressGroupV2State();
    }
}
