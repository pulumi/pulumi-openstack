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
    /// ## Import
    /// 
    /// Floating IPs can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:networking/floatingIp:FloatingIp floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/floatingIp:FloatingIp")]
    public partial class FloatingIp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The actual/specific floating IP to obtain. By default,
        /// non-admin users are not able to specify a floating IP, so you must either be
        /// an admin user or have had a custom policy or role applied to your OpenStack
        /// user or project.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The collection of tags assigned on the floating IP, which have
        /// been explicitly and implicitly added.
        /// </summary>
        [Output("allTags")]
        public Output<ImmutableArray<string>> AllTags { get; private set; } = null!;

        /// <summary>
        /// Human-readable description for the floating IP.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The floating IP DNS domain. Available, when Neutron
        /// DNS extension is enabled. The data in this attribute will be published in an
        /// external DNS service when Neutron is configured to integrate with such a
        /// service. Changing this creates a new floating IP.
        /// </summary>
        [Output("dnsDomain")]
        public Output<string> DnsDomain { get; private set; } = null!;

        /// <summary>
        /// The floating IP DNS name. Available, when Neutron DNS
        /// extension is enabled. The data in this attribute will be published in an
        /// external DNS service when Neutron is configured to integrate with such a
        /// service. Changing this creates a new floating IP.
        /// </summary>
        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Fixed IP of the port to associate with this floating IP. Required if
        /// the port has multiple fixed IPs.
        /// </summary>
        [Output("fixedIp")]
        public Output<string> FixedIp { get; private set; } = null!;

        /// <summary>
        /// The name of the pool from which to obtain the floating
        /// IP. Changing this creates a new floating IP.
        /// </summary>
        [Output("pool")]
        public Output<string> Pool { get; private set; } = null!;

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
        /// The subnet ID of the floating IP pool. Specify this if
        /// the floating IP network has multiple subnets.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// A list of external subnet IDs to try over each to
        /// allocate a floating IP address. If a subnet ID in a list has exhausted
        /// floating IP pool, the next subnet ID will be tried. This argument is used only
        /// during the resource creation. Conflicts with a `subnet_id` argument.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A set of string tags for the floating IP.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The target tenant ID in which to allocate the floating
        /// IP, if you specify this together with a port_id, make sure the target port
        /// belongs to the same tenant. Changing this creates a new floating IP (which
        /// may or may not have a different address)
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, string>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a FloatingIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FloatingIp(string name, FloatingIpArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:networking/floatingIp:FloatingIp", name, args ?? new FloatingIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FloatingIp(string name, Input<string> id, FloatingIpState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/floatingIp:FloatingIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FloatingIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FloatingIp Get(string name, Input<string> id, FloatingIpState? state = null, CustomResourceOptions? options = null)
        {
            return new FloatingIp(name, id, state, options);
        }
    }

    public sealed class FloatingIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The actual/specific floating IP to obtain. By default,
        /// non-admin users are not able to specify a floating IP, so you must either be
        /// an admin user or have had a custom policy or role applied to your OpenStack
        /// user or project.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Human-readable description for the floating IP.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The floating IP DNS domain. Available, when Neutron
        /// DNS extension is enabled. The data in this attribute will be published in an
        /// external DNS service when Neutron is configured to integrate with such a
        /// service. Changing this creates a new floating IP.
        /// </summary>
        [Input("dnsDomain")]
        public Input<string>? DnsDomain { get; set; }

        /// <summary>
        /// The floating IP DNS name. Available, when Neutron DNS
        /// extension is enabled. The data in this attribute will be published in an
        /// external DNS service when Neutron is configured to integrate with such a
        /// service. Changing this creates a new floating IP.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// Fixed IP of the port to associate with this floating IP. Required if
        /// the port has multiple fixed IPs.
        /// </summary>
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// The name of the pool from which to obtain the floating
        /// IP. Changing this creates a new floating IP.
        /// </summary>
        [Input("pool")]
        public Input<string>? Pool { get; set; }

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

        /// <summary>
        /// The subnet ID of the floating IP pool. Specify this if
        /// the floating IP network has multiple subnets.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of external subnet IDs to try over each to
        /// allocate a floating IP address. If a subnet ID in a list has exhausted
        /// floating IP pool, the next subnet ID will be tried. This argument is used only
        /// during the resource creation. Conflicts with a `subnet_id` argument.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the floating IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The target tenant ID in which to allocate the floating
        /// IP, if you specify this together with a port_id, make sure the target port
        /// belongs to the same tenant. Changing this creates a new floating IP (which
        /// may or may not have a different address)
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public FloatingIpArgs()
        {
        }
        public static new FloatingIpArgs Empty => new FloatingIpArgs();
    }

    public sealed class FloatingIpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The actual/specific floating IP to obtain. By default,
        /// non-admin users are not able to specify a floating IP, so you must either be
        /// an admin user or have had a custom policy or role applied to your OpenStack
        /// user or project.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        [Input("allTags")]
        private InputList<string>? _allTags;

        /// <summary>
        /// The collection of tags assigned on the floating IP, which have
        /// been explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllTags
        {
            get => _allTags ?? (_allTags = new InputList<string>());
            set => _allTags = value;
        }

        /// <summary>
        /// Human-readable description for the floating IP.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The floating IP DNS domain. Available, when Neutron
        /// DNS extension is enabled. The data in this attribute will be published in an
        /// external DNS service when Neutron is configured to integrate with such a
        /// service. Changing this creates a new floating IP.
        /// </summary>
        [Input("dnsDomain")]
        public Input<string>? DnsDomain { get; set; }

        /// <summary>
        /// The floating IP DNS name. Available, when Neutron DNS
        /// extension is enabled. The data in this attribute will be published in an
        /// external DNS service when Neutron is configured to integrate with such a
        /// service. Changing this creates a new floating IP.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// Fixed IP of the port to associate with this floating IP. Required if
        /// the port has multiple fixed IPs.
        /// </summary>
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// The name of the pool from which to obtain the floating
        /// IP. Changing this creates a new floating IP.
        /// </summary>
        [Input("pool")]
        public Input<string>? Pool { get; set; }

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

        /// <summary>
        /// The subnet ID of the floating IP pool. Specify this if
        /// the floating IP network has multiple subnets.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of external subnet IDs to try over each to
        /// allocate a floating IP address. If a subnet ID in a list has exhausted
        /// floating IP pool, the next subnet ID will be tried. This argument is used only
        /// during the resource creation. Conflicts with a `subnet_id` argument.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the floating IP.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The target tenant ID in which to allocate the floating
        /// IP, if you specify this together with a port_id, make sure the target port
        /// belongs to the same tenant. Changing this creates a new floating IP (which
        /// may or may not have a different address)
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public FloatingIpState()
        {
        }
        public static new FloatingIpState Empty => new FloatingIpState();
    }
}
