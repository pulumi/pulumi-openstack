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
    /// Manages a V2 port resource within OpenStack.
    /// 
    /// &gt; **Note:** Ports do not get an IP if the network they are attached
    /// to does not have a subnet. If you create the subnet resource in the
    /// same run as the port, make sure to use `fixed_ip.subnet_id` or
    /// `depends_on` to enforce the subnet resource creation before the port
    /// creation is triggered. More info here
    /// 
    /// ## Example Usage
    /// ### Simple port
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network1 = new OpenStack.Networking.Network("network1", new()
    ///     {
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var port1 = new OpenStack.Networking.Port("port1", new()
    ///     {
    ///         NetworkId = network1.Id,
    ///         AdminStateUp = true,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Port defining fixed_ip.subnet_id
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network1 = new OpenStack.Networking.Network("network1", new()
    ///     {
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var subnet1 = new OpenStack.Networking.Subnet("subnet1", new()
    ///     {
    ///         NetworkId = network1.Id,
    ///         Cidr = "192.168.199.0/24",
    ///     });
    /// 
    ///     var port1 = new OpenStack.Networking.Port("port1", new()
    ///     {
    ///         NetworkId = network1.Id,
    ///         AdminStateUp = true,
    ///         FixedIps = new[]
    ///         {
    ///             new OpenStack.Networking.Inputs.PortFixedIpArgs
    ///             {
    ///                 SubnetId = subnet1.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Port with physical binding information
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var network1 = new OpenStack.Networking.Network("network1", new()
    ///     {
    ///         AdminStateUp = true,
    ///     });
    /// 
    ///     var port1 = new OpenStack.Networking.Port("port1", new()
    ///     {
    ///         NetworkId = network1.Id,
    ///         DeviceId = "cdf70fcf-c161-4f24-9c70-96b3f5a54b71",
    ///         DeviceOwner = "baremetal:none",
    ///         AdminStateUp = true,
    ///         Binding = new OpenStack.Networking.Inputs.PortBindingArgs
    ///         {
    ///             HostId = "b080b9cf-46e0-4ce8-ad47-0fd4accc872b",
    ///             VnicType = "baremetal",
    ///             Profile = @"{
    ///   ""local_link_information"": [
    ///     {
    ///       ""switch_info"": ""info1"",
    ///       ""port_id"": ""Ethernet3/4"",
    ///       ""switch_id"": ""12:34:56:78:9A:BC""
    ///     },
    ///     {
    ///       ""switch_info"": ""info2"",
    ///       ""port_id"": ""Ethernet3/4"",
    ///       ""switch_id"": ""12:34:56:78:9A:BD""
    ///     }
    ///   ],
    ///   ""vlan_type"": ""allowed""
    /// }
    /// ",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Notes
    /// 
    /// ### Ports and Instances
    /// 
    /// There are some notes to consider when connecting Instances to networks using
    /// Ports. Please see the `openstack.compute.Instance` documentation for further
    /// documentation.
    /// 
    /// ## Import
    /// 
    /// Ports can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:networking/port:Port port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/port:Port")]
    public partial class Port : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative up/down status for the port
        /// (must be `true` or `false` if provided). Changing this updates the
        /// `admin_state_up` of an existing port.
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The collection of Fixed IP addresses on the port in the
        /// order returned by the Network v2 API.
        /// </summary>
        [Output("allFixedIps")]
        public Output<ImmutableArray<string>> AllFixedIps { get; private set; } = null!;

        /// <summary>
        /// The collection of Security Group IDs on the port
        /// which have been explicitly and implicitly added.
        /// </summary>
        [Output("allSecurityGroupIds")]
        public Output<ImmutableArray<string>> AllSecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The collection of tags assigned on the port, which have been
        /// explicitly and implicitly added.
        /// </summary>
        [Output("allTags")]
        public Output<ImmutableArray<string>> AllTags { get; private set; } = null!;

        /// <summary>
        /// An IP/MAC Address pair of additional IP
        /// addresses that can be active on this port. The structure is described
        /// below.
        /// </summary>
        [Output("allowedAddressPairs")]
        public Output<ImmutableArray<Outputs.PortAllowedAddressPair>> AllowedAddressPairs { get; private set; } = null!;

        /// <summary>
        /// The port binding allows to specify binding information
        /// for the port. The structure is described below.
        /// </summary>
        [Output("binding")]
        public Output<Outputs.PortBinding> Binding { get; private set; } = null!;

        /// <summary>
        /// Human-readable description of the port. Changing
        /// this updates the `description` of an existing port.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the device attached to the port. Changing this
        /// creates a new port.
        /// </summary>
        [Output("deviceId")]
        public Output<string> DeviceId { get; private set; } = null!;

        /// <summary>
        /// The device owner of the port. Changing this creates
        /// a new port.
        /// </summary>
        [Output("deviceOwner")]
        public Output<string> DeviceOwner { get; private set; } = null!;

        /// <summary>
        /// The list of maps representing port DNS assignments.
        /// </summary>
        [Output("dnsAssignments")]
        public Output<ImmutableArray<ImmutableDictionary<string, object>>> DnsAssignments { get; private set; } = null!;

        /// <summary>
        /// The port DNS name. Available, when Neutron DNS extension
        /// is enabled.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// An extra DHCP option that needs to be configured
        /// on the port. The structure is described below. Can be specified multiple
        /// times.
        /// </summary>
        [Output("extraDhcpOptions")]
        public Output<ImmutableArray<Outputs.PortExtraDhcpOption>> ExtraDhcpOptions { get; private set; } = null!;

        /// <summary>
        /// An array of desired IPs for
        /// this port. The structure is described below.
        /// </summary>
        [Output("fixedIps")]
        public Output<ImmutableArray<Outputs.PortFixedIp>> FixedIps { get; private set; } = null!;

        /// <summary>
        /// Specify a specific MAC address for the port. Changing
        /// this creates a new port.
        /// </summary>
        [Output("macAddress")]
        public Output<string> MacAddress { get; private set; } = null!;

        /// <summary>
        /// A unique name for the port. Changing this
        /// updates the `name` of an existing port.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the network to attach the port to. Changing
        /// this creates a new port.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Create a port with no fixed
        /// IP address. This will also remove any fixed IPs previously set on a port. `true`
        /// is the only valid value for this argument.
        /// </summary>
        [Output("noFixedIp")]
        public Output<bool?> NoFixedIp { get; private set; } = null!;

        /// <summary>
        /// If set to
        /// `true`, then no security groups are applied to the port. If set to `false` and
        /// no `security_group_ids` are specified, then the port will yield to the default
        /// behavior of the Networking service, which is to usually apply the "default"
        /// security group.
        /// </summary>
        [Output("noSecurityGroups")]
        public Output<bool?> NoSecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Whether to explicitly enable or disable
        /// port security on the port. Port Security is usually enabled by default, so
        /// omitting argument will usually result in a value of `true`. Setting this
        /// explicitly to `false` will disable port security. In order to disable port
        /// security, the port must not have any security groups. Valid values are `true`
        /// and `false`.
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
        /// A Networking client is needed to create a port. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// port.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A list
        /// of security group IDs to apply to the port. The security groups must be
        /// specified by ID and not name (as opposed to how they are configured with
        /// the Compute Instance).
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// A set of string tags for the port.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The owner of the port. Required if admin wants
        /// to create a port for another tenant. Changing this creates a new port.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a Port resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Port(string name, PortArgs args, CustomResourceOptions? options = null)
            : base("openstack:networking/port:Port", name, args ?? new PortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Port(string name, Input<string> id, PortState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/port:Port", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Port resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Port Get(string name, Input<string> id, PortState? state = null, CustomResourceOptions? options = null)
        {
            return new Port(name, id, state, options);
        }
    }

    public sealed class PortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative up/down status for the port
        /// (must be `true` or `false` if provided). Changing this updates the
        /// `admin_state_up` of an existing port.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("allowedAddressPairs")]
        private InputList<Inputs.PortAllowedAddressPairArgs>? _allowedAddressPairs;

        /// <summary>
        /// An IP/MAC Address pair of additional IP
        /// addresses that can be active on this port. The structure is described
        /// below.
        /// </summary>
        public InputList<Inputs.PortAllowedAddressPairArgs> AllowedAddressPairs
        {
            get => _allowedAddressPairs ?? (_allowedAddressPairs = new InputList<Inputs.PortAllowedAddressPairArgs>());
            set => _allowedAddressPairs = value;
        }

        /// <summary>
        /// The port binding allows to specify binding information
        /// for the port. The structure is described below.
        /// </summary>
        [Input("binding")]
        public Input<Inputs.PortBindingArgs>? Binding { get; set; }

        /// <summary>
        /// Human-readable description of the port. Changing
        /// this updates the `description` of an existing port.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the device attached to the port. Changing this
        /// creates a new port.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// The device owner of the port. Changing this creates
        /// a new port.
        /// </summary>
        [Input("deviceOwner")]
        public Input<string>? DeviceOwner { get; set; }

        /// <summary>
        /// The port DNS name. Available, when Neutron DNS extension
        /// is enabled.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        [Input("extraDhcpOptions")]
        private InputList<Inputs.PortExtraDhcpOptionArgs>? _extraDhcpOptions;

        /// <summary>
        /// An extra DHCP option that needs to be configured
        /// on the port. The structure is described below. Can be specified multiple
        /// times.
        /// </summary>
        public InputList<Inputs.PortExtraDhcpOptionArgs> ExtraDhcpOptions
        {
            get => _extraDhcpOptions ?? (_extraDhcpOptions = new InputList<Inputs.PortExtraDhcpOptionArgs>());
            set => _extraDhcpOptions = value;
        }

        [Input("fixedIps")]
        private InputList<Inputs.PortFixedIpArgs>? _fixedIps;

        /// <summary>
        /// An array of desired IPs for
        /// this port. The structure is described below.
        /// </summary>
        public InputList<Inputs.PortFixedIpArgs> FixedIps
        {
            get => _fixedIps ?? (_fixedIps = new InputList<Inputs.PortFixedIpArgs>());
            set => _fixedIps = value;
        }

        /// <summary>
        /// Specify a specific MAC address for the port. Changing
        /// this creates a new port.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// A unique name for the port. Changing this
        /// updates the `name` of an existing port.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the network to attach the port to. Changing
        /// this creates a new port.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Create a port with no fixed
        /// IP address. This will also remove any fixed IPs previously set on a port. `true`
        /// is the only valid value for this argument.
        /// </summary>
        [Input("noFixedIp")]
        public Input<bool>? NoFixedIp { get; set; }

        /// <summary>
        /// If set to
        /// `true`, then no security groups are applied to the port. If set to `false` and
        /// no `security_group_ids` are specified, then the port will yield to the default
        /// behavior of the Networking service, which is to usually apply the "default"
        /// security group.
        /// </summary>
        [Input("noSecurityGroups")]
        public Input<bool>? NoSecurityGroups { get; set; }

        /// <summary>
        /// Whether to explicitly enable or disable
        /// port security on the port. Port Security is usually enabled by default, so
        /// omitting argument will usually result in a value of `true`. Setting this
        /// explicitly to `false` will disable port security. In order to disable port
        /// security, the port must not have any security groups. Valid values are `true`
        /// and `false`.
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
        /// A Networking client is needed to create a port. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// port.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list
        /// of security group IDs to apply to the port. The security groups must be
        /// specified by ID and not name (as opposed to how they are configured with
        /// the Compute Instance).
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
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
        /// The owner of the port. Required if admin wants
        /// to create a port for another tenant. Changing this creates a new port.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public PortArgs()
        {
        }
        public static new PortArgs Empty => new PortArgs();
    }

    public sealed class PortState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative up/down status for the port
        /// (must be `true` or `false` if provided). Changing this updates the
        /// `admin_state_up` of an existing port.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        [Input("allFixedIps")]
        private InputList<string>? _allFixedIps;

        /// <summary>
        /// The collection of Fixed IP addresses on the port in the
        /// order returned by the Network v2 API.
        /// </summary>
        public InputList<string> AllFixedIps
        {
            get => _allFixedIps ?? (_allFixedIps = new InputList<string>());
            set => _allFixedIps = value;
        }

        [Input("allSecurityGroupIds")]
        private InputList<string>? _allSecurityGroupIds;

        /// <summary>
        /// The collection of Security Group IDs on the port
        /// which have been explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllSecurityGroupIds
        {
            get => _allSecurityGroupIds ?? (_allSecurityGroupIds = new InputList<string>());
            set => _allSecurityGroupIds = value;
        }

        [Input("allTags")]
        private InputList<string>? _allTags;

        /// <summary>
        /// The collection of tags assigned on the port, which have been
        /// explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllTags
        {
            get => _allTags ?? (_allTags = new InputList<string>());
            set => _allTags = value;
        }

        [Input("allowedAddressPairs")]
        private InputList<Inputs.PortAllowedAddressPairGetArgs>? _allowedAddressPairs;

        /// <summary>
        /// An IP/MAC Address pair of additional IP
        /// addresses that can be active on this port. The structure is described
        /// below.
        /// </summary>
        public InputList<Inputs.PortAllowedAddressPairGetArgs> AllowedAddressPairs
        {
            get => _allowedAddressPairs ?? (_allowedAddressPairs = new InputList<Inputs.PortAllowedAddressPairGetArgs>());
            set => _allowedAddressPairs = value;
        }

        /// <summary>
        /// The port binding allows to specify binding information
        /// for the port. The structure is described below.
        /// </summary>
        [Input("binding")]
        public Input<Inputs.PortBindingGetArgs>? Binding { get; set; }

        /// <summary>
        /// Human-readable description of the port. Changing
        /// this updates the `description` of an existing port.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the device attached to the port. Changing this
        /// creates a new port.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// The device owner of the port. Changing this creates
        /// a new port.
        /// </summary>
        [Input("deviceOwner")]
        public Input<string>? DeviceOwner { get; set; }

        [Input("dnsAssignments")]
        private InputList<ImmutableDictionary<string, object>>? _dnsAssignments;

        /// <summary>
        /// The list of maps representing port DNS assignments.
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> DnsAssignments
        {
            get => _dnsAssignments ?? (_dnsAssignments = new InputList<ImmutableDictionary<string, object>>());
            set => _dnsAssignments = value;
        }

        /// <summary>
        /// The port DNS name. Available, when Neutron DNS extension
        /// is enabled.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        [Input("extraDhcpOptions")]
        private InputList<Inputs.PortExtraDhcpOptionGetArgs>? _extraDhcpOptions;

        /// <summary>
        /// An extra DHCP option that needs to be configured
        /// on the port. The structure is described below. Can be specified multiple
        /// times.
        /// </summary>
        public InputList<Inputs.PortExtraDhcpOptionGetArgs> ExtraDhcpOptions
        {
            get => _extraDhcpOptions ?? (_extraDhcpOptions = new InputList<Inputs.PortExtraDhcpOptionGetArgs>());
            set => _extraDhcpOptions = value;
        }

        [Input("fixedIps")]
        private InputList<Inputs.PortFixedIpGetArgs>? _fixedIps;

        /// <summary>
        /// An array of desired IPs for
        /// this port. The structure is described below.
        /// </summary>
        public InputList<Inputs.PortFixedIpGetArgs> FixedIps
        {
            get => _fixedIps ?? (_fixedIps = new InputList<Inputs.PortFixedIpGetArgs>());
            set => _fixedIps = value;
        }

        /// <summary>
        /// Specify a specific MAC address for the port. Changing
        /// this creates a new port.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// A unique name for the port. Changing this
        /// updates the `name` of an existing port.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the network to attach the port to. Changing
        /// this creates a new port.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Create a port with no fixed
        /// IP address. This will also remove any fixed IPs previously set on a port. `true`
        /// is the only valid value for this argument.
        /// </summary>
        [Input("noFixedIp")]
        public Input<bool>? NoFixedIp { get; set; }

        /// <summary>
        /// If set to
        /// `true`, then no security groups are applied to the port. If set to `false` and
        /// no `security_group_ids` are specified, then the port will yield to the default
        /// behavior of the Networking service, which is to usually apply the "default"
        /// security group.
        /// </summary>
        [Input("noSecurityGroups")]
        public Input<bool>? NoSecurityGroups { get; set; }

        /// <summary>
        /// Whether to explicitly enable or disable
        /// port security on the port. Port Security is usually enabled by default, so
        /// omitting argument will usually result in a value of `true`. Setting this
        /// explicitly to `false` will disable port security. In order to disable port
        /// security, the port must not have any security groups. Valid values are `true`
        /// and `false`.
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
        /// A Networking client is needed to create a port. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// port.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list
        /// of security group IDs to apply to the port. The security groups must be
        /// specified by ID and not name (as opposed to how they are configured with
        /// the Compute Instance).
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
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
        /// The owner of the port. Required if admin wants
        /// to create a port for another tenant. Changing this creates a new port.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public PortState()
        {
        }
        public static new PortState Empty => new PortState();
    }
}
