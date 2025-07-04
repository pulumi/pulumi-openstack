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
    /// Manages a V2 Neutron network resource within OpenStack.
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
    ///     var network1 = new OpenStack.Networking.Network("network_1", new()
    ///     {
    ///         Name = "network_1",
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var subnet1 = new OpenStack.Networking.Subnet("subnet_1", new()
    ///     {
    ///         Name = "subnet_1",
    ///         NetworkId = network1.Id,
    ///         Cidr = "192.168.199.0/24",
    ///         IpVersion = 4,
    ///     });
    /// 
    ///     var secgroup1 = new OpenStack.Networking.SecGroup("secgroup_1", new()
    ///     {
    ///         Name = "secgroup_1",
    ///         Description = "a security group",
    ///     });
    /// 
    ///     var secgroupRule1 = new OpenStack.Networking.SecGroupRule("secgroup_rule_1", new()
    ///     {
    ///         Direction = "ingress",
    ///         Ethertype = "IPv4",
    ///         Protocol = "tcp",
    ///         PortRangeMin = 22,
    ///         PortRangeMax = 22,
    ///         RemoteIpPrefix = "0.0.0.0/0",
    ///         SecurityGroupId = secgroup1.Id,
    ///     });
    /// 
    ///     var port1 = new OpenStack.Networking.Port("port_1", new()
    ///     {
    ///         Name = "port_1",
    ///         NetworkId = network1.Id,
    ///         AdminStateUp = true,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             secgroup1.Id,
    ///         },
    ///         FixedIps = new[]
    ///         {
    ///             new OpenStack.Networking.Inputs.PortFixedIpArgs
    ///             {
    ///                 SubnetId = subnet1.Id,
    ///                 IpAddress = "192.168.199.10",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var instance1 = new OpenStack.Compute.Instance("instance_1", new()
    ///     {
    ///         Name = "instance_1",
    ///         SecurityGroups = new[]
    ///         {
    ///             secgroup1.Name,
    ///         },
    ///         Networks = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.InstanceNetworkArgs
    ///             {
    ///                 Port = port1.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Networks can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:networking/network:Network network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/network:Network")]
    public partial class Network : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the network.
        /// Acceptable values are "true" and "false". Changing this value updates the
        /// state of the existing network.
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The collection of tags assigned on the network, which have been
        /// explicitly and implicitly added.
        /// </summary>
        [Output("allTags")]
        public Output<ImmutableArray<string>> AllTags { get; private set; } = null!;

        /// <summary>
        /// An availability zone is used to make
        /// network resources highly available. Used for resources with high availability
        /// so that they are scheduled on different availability zones. Changing this
        /// creates a new network.
        /// </summary>
        [Output("availabilityZoneHints")]
        public Output<ImmutableArray<string>> AvailabilityZoneHints { get; private set; } = null!;

        /// <summary>
        /// Human-readable description of the network. Changing this
        /// updates the name of the existing network.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The network DNS domain. Available, when Neutron DNS
        /// extension is enabled. The `dns_domain` of a network in conjunction with the
        /// `dns_name` attribute of its ports will be published in an external DNS
        /// service when Neutron is configured to integrate with such a service.
        /// </summary>
        [Output("dnsDomain")]
        public Output<string> DnsDomain { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the network resource has the
        /// external routing facility. Valid values are true and false. Defaults to
        /// false. Changing this updates the external attribute of the existing network.
        /// </summary>
        [Output("external")]
        public Output<bool> External { get; private set; } = null!;

        /// <summary>
        /// The network MTU. Available for read-only, when Neutron
        /// `net-mtu` extension is enabled. Available for the modification, when
        /// Neutron `net-mtu-writable` extension is enabled.
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;

        /// <summary>
        /// The name of the network. Changing this updates the name of
        /// the existing network.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether to explicitly enable or disable
        /// port security on the network. Port Security is usually enabled by default, so
        /// omitting this argument will usually result in a value of "true". Setting this
        /// explicitly to `false` will disable port security. Valid values are `true` and
        /// `false`.
        /// </summary>
        [Output("portSecurityEnabled")]
        public Output<bool> PortSecurityEnabled { get; private set; } = null!;

        /// <summary>
        /// Reference to the associated QoS policy.
        /// </summary>
        [Output("qosPolicyId")]
        public Output<string> QosPolicyId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron network. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// network.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// An array of one or more provider segment objects.
        /// Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
        /// updating any provider related segments attributes. Check your plug-in whether
        /// it supports updating.
        /// </summary>
        [Output("segments")]
        public Output<ImmutableArray<Outputs.NetworkSegment>> Segments { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the network resource can be accessed
        /// by any tenant or not. Changing this updates the sharing capabilities of the
        /// existing network.
        /// </summary>
        [Output("shared")]
        public Output<bool> Shared { get; private set; } = null!;

        /// <summary>
        /// A set of string tags for the network.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The owner of the network. Required if admin wants to
        /// create a network for another tenant. Changing this creates a new network.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the network resource has the
        /// VLAN transparent attribute set. Valid values are true and false. Defaults to
        /// false. Changing this updates the `transparent_vlan` attribute of the existing
        /// network.
        /// </summary>
        [Output("transparentVlan")]
        public Output<bool> TransparentVlan { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, string>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a Network resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Network(string name, NetworkArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:networking/network:Network", name, args ?? new NetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Network(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/network:Network", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Network resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Network Get(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Network(name, id, state, options);
        }
    }

    public sealed class NetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the network.
        /// Acceptable values are "true" and "false". Changing this value updates the
        /// state of the existing network.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("availabilityZoneHints")]
        private InputList<string>? _availabilityZoneHints;

        /// <summary>
        /// An availability zone is used to make
        /// network resources highly available. Used for resources with high availability
        /// so that they are scheduled on different availability zones. Changing this
        /// creates a new network.
        /// </summary>
        public InputList<string> AvailabilityZoneHints
        {
            get => _availabilityZoneHints ?? (_availabilityZoneHints = new InputList<string>());
            set => _availabilityZoneHints = value;
        }

        /// <summary>
        /// Human-readable description of the network. Changing this
        /// updates the name of the existing network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The network DNS domain. Available, when Neutron DNS
        /// extension is enabled. The `dns_domain` of a network in conjunction with the
        /// `dns_name` attribute of its ports will be published in an external DNS
        /// service when Neutron is configured to integrate with such a service.
        /// </summary>
        [Input("dnsDomain")]
        public Input<string>? DnsDomain { get; set; }

        /// <summary>
        /// Specifies whether the network resource has the
        /// external routing facility. Valid values are true and false. Defaults to
        /// false. Changing this updates the external attribute of the existing network.
        /// </summary>
        [Input("external")]
        public Input<bool>? External { get; set; }

        /// <summary>
        /// The network MTU. Available for read-only, when Neutron
        /// `net-mtu` extension is enabled. Available for the modification, when
        /// Neutron `net-mtu-writable` extension is enabled.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// The name of the network. Changing this updates the name of
        /// the existing network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to explicitly enable or disable
        /// port security on the network. Port Security is usually enabled by default, so
        /// omitting this argument will usually result in a value of "true". Setting this
        /// explicitly to `false` will disable port security. Valid values are `true` and
        /// `false`.
        /// </summary>
        [Input("portSecurityEnabled")]
        public Input<bool>? PortSecurityEnabled { get; set; }

        /// <summary>
        /// Reference to the associated QoS policy.
        /// </summary>
        [Input("qosPolicyId")]
        public Input<string>? QosPolicyId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron network. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// network.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("segments")]
        private InputList<Inputs.NetworkSegmentArgs>? _segments;

        /// <summary>
        /// An array of one or more provider segment objects.
        /// Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
        /// updating any provider related segments attributes. Check your plug-in whether
        /// it supports updating.
        /// </summary>
        public InputList<Inputs.NetworkSegmentArgs> Segments
        {
            get => _segments ?? (_segments = new InputList<Inputs.NetworkSegmentArgs>());
            set => _segments = value;
        }

        /// <summary>
        /// Specifies whether the network resource can be accessed
        /// by any tenant or not. Changing this updates the sharing capabilities of the
        /// existing network.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the network.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the network. Required if admin wants to
        /// create a network for another tenant. Changing this creates a new network.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Specifies whether the network resource has the
        /// VLAN transparent attribute set. Valid values are true and false. Defaults to
        /// false. Changing this updates the `transparent_vlan` attribute of the existing
        /// network.
        /// </summary>
        [Input("transparentVlan")]
        public Input<bool>? TransparentVlan { get; set; }

        [Input("valueSpecs")]
        private InputMap<string>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<string> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<string>());
            set => _valueSpecs = value;
        }

        public NetworkArgs()
        {
        }
        public static new NetworkArgs Empty => new NetworkArgs();
    }

    public sealed class NetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the network.
        /// Acceptable values are "true" and "false". Changing this value updates the
        /// state of the existing network.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("allTags")]
        private InputList<string>? _allTags;

        /// <summary>
        /// The collection of tags assigned on the network, which have been
        /// explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllTags
        {
            get => _allTags ?? (_allTags = new InputList<string>());
            set => _allTags = value;
        }

        [Input("availabilityZoneHints")]
        private InputList<string>? _availabilityZoneHints;

        /// <summary>
        /// An availability zone is used to make
        /// network resources highly available. Used for resources with high availability
        /// so that they are scheduled on different availability zones. Changing this
        /// creates a new network.
        /// </summary>
        public InputList<string> AvailabilityZoneHints
        {
            get => _availabilityZoneHints ?? (_availabilityZoneHints = new InputList<string>());
            set => _availabilityZoneHints = value;
        }

        /// <summary>
        /// Human-readable description of the network. Changing this
        /// updates the name of the existing network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The network DNS domain. Available, when Neutron DNS
        /// extension is enabled. The `dns_domain` of a network in conjunction with the
        /// `dns_name` attribute of its ports will be published in an external DNS
        /// service when Neutron is configured to integrate with such a service.
        /// </summary>
        [Input("dnsDomain")]
        public Input<string>? DnsDomain { get; set; }

        /// <summary>
        /// Specifies whether the network resource has the
        /// external routing facility. Valid values are true and false. Defaults to
        /// false. Changing this updates the external attribute of the existing network.
        /// </summary>
        [Input("external")]
        public Input<bool>? External { get; set; }

        /// <summary>
        /// The network MTU. Available for read-only, when Neutron
        /// `net-mtu` extension is enabled. Available for the modification, when
        /// Neutron `net-mtu-writable` extension is enabled.
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        /// <summary>
        /// The name of the network. Changing this updates the name of
        /// the existing network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether to explicitly enable or disable
        /// port security on the network. Port Security is usually enabled by default, so
        /// omitting this argument will usually result in a value of "true". Setting this
        /// explicitly to `false` will disable port security. Valid values are `true` and
        /// `false`.
        /// </summary>
        [Input("portSecurityEnabled")]
        public Input<bool>? PortSecurityEnabled { get; set; }

        /// <summary>
        /// Reference to the associated QoS policy.
        /// </summary>
        [Input("qosPolicyId")]
        public Input<string>? QosPolicyId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron network. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// network.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("segments")]
        private InputList<Inputs.NetworkSegmentGetArgs>? _segments;

        /// <summary>
        /// An array of one or more provider segment objects.
        /// Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
        /// updating any provider related segments attributes. Check your plug-in whether
        /// it supports updating.
        /// </summary>
        public InputList<Inputs.NetworkSegmentGetArgs> Segments
        {
            get => _segments ?? (_segments = new InputList<Inputs.NetworkSegmentGetArgs>());
            set => _segments = value;
        }

        /// <summary>
        /// Specifies whether the network resource can be accessed
        /// by any tenant or not. Changing this updates the sharing capabilities of the
        /// existing network.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the network.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the network. Required if admin wants to
        /// create a network for another tenant. Changing this creates a new network.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Specifies whether the network resource has the
        /// VLAN transparent attribute set. Valid values are true and false. Defaults to
        /// false. Changing this updates the `transparent_vlan` attribute of the existing
        /// network.
        /// </summary>
        [Input("transparentVlan")]
        public Input<bool>? TransparentVlan { get; set; }

        [Input("valueSpecs")]
        private InputMap<string>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<string> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<string>());
            set => _valueSpecs = value;
        }

        public NetworkState()
        {
        }
        public static new NetworkState Empty => new NetworkState();
    }
}
