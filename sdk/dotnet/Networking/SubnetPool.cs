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
    /// Manages a V2 Neutron subnetpool resource within OpenStack.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a Subnet Pool
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
    ///     var subnetpool1 = new OpenStack.Networking.SubnetPool("subnetpool_1", new()
    ///     {
    ///         Name = "subnetpool_1",
    ///         IpVersion = 6,
    ///         Prefixes = new[]
    ///         {
    ///             "fdf7:b13d:dead:beef::/64",
    ///             "fd65:86cc:a334:39b7::/64",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Create a Subnet from a Subnet Pool
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
    ///     var network1 = new OpenStack.Networking.Network("network_1", new()
    ///     {
    ///         Name = "network_1",
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var subnetpool1 = new OpenStack.Networking.SubnetPool("subnetpool_1", new()
    ///     {
    ///         Name = "subnetpool_1",
    ///         Prefixes = new[]
    ///         {
    ///             "10.11.12.0/24",
    ///         },
    ///     });
    /// 
    ///     var subnet1 = new OpenStack.Networking.Subnet("subnet_1", new()
    ///     {
    ///         Name = "subnet_1",
    ///         Cidr = "10.11.12.0/25",
    ///         NetworkId = network1.Id,
    ///         SubnetpoolId = subnetpool1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Subnetpools can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:networking/subnetPool:SubnetPool subnetpool_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/subnetPool:SubnetPool")]
    public partial class SubnetPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Neutron address scope to assign to the
        /// subnetpool. Changing this updates the address scope id of the existing
        /// subnetpool.
        /// </summary>
        [Output("addressScopeId")]
        public Output<string?> AddressScopeId { get; private set; } = null!;

        /// <summary>
        /// The collection of tags assigned on the subnetpool, which have been
        /// explicitly and implicitly added.
        /// </summary>
        [Output("allTags")]
        public Output<ImmutableArray<string>> AllTags { get; private set; } = null!;

        /// <summary>
        /// The time at which subnetpool was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The size of the prefix to allocate when the cidr
        /// or prefixlen attributes are omitted when you create the subnet. Defaults to the
        /// MinPrefixLen. Changing this updates the default prefixlen of the existing
        /// subnetpool.
        /// </summary>
        [Output("defaultPrefixlen")]
        public Output<int> DefaultPrefixlen { get; private set; } = null!;

        /// <summary>
        /// The per-project quota on the prefix space that can be
        /// allocated from the subnetpool for project subnets. Changing this updates the
        /// default quota of the existing subnetpool.
        /// </summary>
        [Output("defaultQuota")]
        public Output<int?> DefaultQuota { get; private set; } = null!;

        /// <summary>
        /// The human-readable description for the subnetpool.
        /// Changing this updates the description of the existing subnetpool.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The IP protocol version.
        /// </summary>
        [Output("ipVersion")]
        public Output<int> IpVersion { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the subnetpool is default
        /// subnetpool or not. Changing this updates the default status of the existing
        /// subnetpool.
        /// </summary>
        [Output("isDefault")]
        public Output<bool?> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The maximum prefix size that can be allocated from
        /// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
        /// default is 128. Changing this updates the max prefixlen of the existing
        /// subnetpool.
        /// </summary>
        [Output("maxPrefixlen")]
        public Output<int> MaxPrefixlen { get; private set; } = null!;

        /// <summary>
        /// The smallest prefix that can be allocated from a
        /// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
        /// is 64. Changing this updates the min prefixlen of the existing subnetpool.
        /// </summary>
        [Output("minPrefixlen")]
        public Output<int> MinPrefixlen { get; private set; } = null!;

        /// <summary>
        /// The name of the subnetpool. Changing this updates the name of
        /// the existing subnetpool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of subnet prefixes to assign to the subnetpool.
        /// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
        /// subnet prefix must be unique among all subnet prefixes in all subnetpools that
        /// are associated with the address scope. Changing this updates the prefixes list
        /// of the existing subnetpool.
        /// </summary>
        [Output("prefixes")]
        public Output<ImmutableArray<string>> Prefixes { get; private set; } = null!;

        /// <summary>
        /// The owner of the subnetpool. Required if admin wants to
        /// create a subnetpool for another project. Changing this creates a new subnetpool.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron subnetpool. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// subnetpool.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The revision number of the subnetpool.
        /// </summary>
        [Output("revisionNumber")]
        public Output<int> RevisionNumber { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this subnetpool is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// subnetpool.
        /// </summary>
        [Output("shared")]
        public Output<bool?> Shared { get; private set; } = null!;

        /// <summary>
        /// A set of string tags for the subnetpool.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The time at which subnetpool was created.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a SubnetPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubnetPool(string name, SubnetPoolArgs args, CustomResourceOptions? options = null)
            : base("openstack:networking/subnetPool:SubnetPool", name, args ?? new SubnetPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubnetPool(string name, Input<string> id, SubnetPoolState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/subnetPool:SubnetPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SubnetPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubnetPool Get(string name, Input<string> id, SubnetPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new SubnetPool(name, id, state, options);
        }
    }

    public sealed class SubnetPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Neutron address scope to assign to the
        /// subnetpool. Changing this updates the address scope id of the existing
        /// subnetpool.
        /// </summary>
        [Input("addressScopeId")]
        public Input<string>? AddressScopeId { get; set; }

        /// <summary>
        /// The size of the prefix to allocate when the cidr
        /// or prefixlen attributes are omitted when you create the subnet. Defaults to the
        /// MinPrefixLen. Changing this updates the default prefixlen of the existing
        /// subnetpool.
        /// </summary>
        [Input("defaultPrefixlen")]
        public Input<int>? DefaultPrefixlen { get; set; }

        /// <summary>
        /// The per-project quota on the prefix space that can be
        /// allocated from the subnetpool for project subnets. Changing this updates the
        /// default quota of the existing subnetpool.
        /// </summary>
        [Input("defaultQuota")]
        public Input<int>? DefaultQuota { get; set; }

        /// <summary>
        /// The human-readable description for the subnetpool.
        /// Changing this updates the description of the existing subnetpool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP protocol version.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// Indicates whether the subnetpool is default
        /// subnetpool or not. Changing this updates the default status of the existing
        /// subnetpool.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The maximum prefix size that can be allocated from
        /// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
        /// default is 128. Changing this updates the max prefixlen of the existing
        /// subnetpool.
        /// </summary>
        [Input("maxPrefixlen")]
        public Input<int>? MaxPrefixlen { get; set; }

        /// <summary>
        /// The smallest prefix that can be allocated from a
        /// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
        /// is 64. Changing this updates the min prefixlen of the existing subnetpool.
        /// </summary>
        [Input("minPrefixlen")]
        public Input<int>? MinPrefixlen { get; set; }

        /// <summary>
        /// The name of the subnetpool. Changing this updates the name of
        /// the existing subnetpool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prefixes", required: true)]
        private InputList<string>? _prefixes;

        /// <summary>
        /// A list of subnet prefixes to assign to the subnetpool.
        /// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
        /// subnet prefix must be unique among all subnet prefixes in all subnetpools that
        /// are associated with the address scope. Changing this updates the prefixes list
        /// of the existing subnetpool.
        /// </summary>
        public InputList<string> Prefixes
        {
            get => _prefixes ?? (_prefixes = new InputList<string>());
            set => _prefixes = value;
        }

        /// <summary>
        /// The owner of the subnetpool. Required if admin wants to
        /// create a subnetpool for another project. Changing this creates a new subnetpool.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron subnetpool. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// subnetpool.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates whether this subnetpool is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// subnetpool.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the subnetpool.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

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

        public SubnetPoolArgs()
        {
        }
        public static new SubnetPoolArgs Empty => new SubnetPoolArgs();
    }

    public sealed class SubnetPoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Neutron address scope to assign to the
        /// subnetpool. Changing this updates the address scope id of the existing
        /// subnetpool.
        /// </summary>
        [Input("addressScopeId")]
        public Input<string>? AddressScopeId { get; set; }

        [Input("allTags")]
        private InputList<string>? _allTags;

        /// <summary>
        /// The collection of tags assigned on the subnetpool, which have been
        /// explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllTags
        {
            get => _allTags ?? (_allTags = new InputList<string>());
            set => _allTags = value;
        }

        /// <summary>
        /// The time at which subnetpool was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The size of the prefix to allocate when the cidr
        /// or prefixlen attributes are omitted when you create the subnet. Defaults to the
        /// MinPrefixLen. Changing this updates the default prefixlen of the existing
        /// subnetpool.
        /// </summary>
        [Input("defaultPrefixlen")]
        public Input<int>? DefaultPrefixlen { get; set; }

        /// <summary>
        /// The per-project quota on the prefix space that can be
        /// allocated from the subnetpool for project subnets. Changing this updates the
        /// default quota of the existing subnetpool.
        /// </summary>
        [Input("defaultQuota")]
        public Input<int>? DefaultQuota { get; set; }

        /// <summary>
        /// The human-readable description for the subnetpool.
        /// Changing this updates the description of the existing subnetpool.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP protocol version.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// Indicates whether the subnetpool is default
        /// subnetpool or not. Changing this updates the default status of the existing
        /// subnetpool.
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// The maximum prefix size that can be allocated from
        /// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
        /// default is 128. Changing this updates the max prefixlen of the existing
        /// subnetpool.
        /// </summary>
        [Input("maxPrefixlen")]
        public Input<int>? MaxPrefixlen { get; set; }

        /// <summary>
        /// The smallest prefix that can be allocated from a
        /// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
        /// is 64. Changing this updates the min prefixlen of the existing subnetpool.
        /// </summary>
        [Input("minPrefixlen")]
        public Input<int>? MinPrefixlen { get; set; }

        /// <summary>
        /// The name of the subnetpool. Changing this updates the name of
        /// the existing subnetpool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("prefixes")]
        private InputList<string>? _prefixes;

        /// <summary>
        /// A list of subnet prefixes to assign to the subnetpool.
        /// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
        /// subnet prefix must be unique among all subnet prefixes in all subnetpools that
        /// are associated with the address scope. Changing this updates the prefixes list
        /// of the existing subnetpool.
        /// </summary>
        public InputList<string> Prefixes
        {
            get => _prefixes ?? (_prefixes = new InputList<string>());
            set => _prefixes = value;
        }

        /// <summary>
        /// The owner of the subnetpool. Required if admin wants to
        /// create a subnetpool for another project. Changing this creates a new subnetpool.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron subnetpool. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// subnetpool.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The revision number of the subnetpool.
        /// </summary>
        [Input("revisionNumber")]
        public Input<int>? RevisionNumber { get; set; }

        /// <summary>
        /// Indicates whether this subnetpool is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// subnetpool.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the subnetpool.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The time at which subnetpool was created.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

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

        public SubnetPoolState()
        {
        }
        public static new SubnetPoolState Empty => new SubnetPoolState();
    }
}
