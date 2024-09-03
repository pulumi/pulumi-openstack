// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack
{
    /// <summary>
    /// Manages a V2 loadbalancer resource within OpenStack.
    /// 
    /// &gt; **Note:** This resource has attributes that depend on octavia minor versions.
    /// Please ensure your Openstack cloud supports the required minor version.
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
    ///     var lb1 = new OpenStack.LbLoadbalancerV2("lb_1", new()
    ///     {
    ///         VipSubnetId = "d9415786-5f1a-428b-b35f-2f1523e146d2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer can be imported using the Load Balancer ID, e.g.:
    /// 
    /// ```sh
    /// $ pulumi import openstack:index/lbLoadbalancerV2:LbLoadbalancerV2 loadbalancer_1 19bcfdc7-c521-4a7e-9459-6750bd16df76
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:index/lbLoadbalancerV2:LbLoadbalancerV2")]
    public partial class LbLoadbalancerV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The administrative state of the Loadbalancer.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The availability zone of the Loadbalancer.
        /// Changing this creates a new loadbalancer. Available only for Octavia
        /// **minor version 2.14 or later**.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Human-readable description for the Loadbalancer.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The UUID of a flavor. Changing this creates a new
        /// loadbalancer.
        /// </summary>
        [Output("flavorId")]
        public Output<string> FlavorId { get; private set; } = null!;

        /// <summary>
        /// The name of the provider. Changing this
        /// creates a new loadbalancer.
        /// </summary>
        [Output("loadbalancerProvider")]
        public Output<string> LoadbalancerProvider { get; private set; } = null!;

        /// <summary>
        /// Human-readable name for the Loadbalancer. Does not have
        /// to be unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB member. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB member.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs to apply to the
        /// loadbalancer. The security groups must be specified by ID and not name (as
        /// opposed to how they are configured with the Compute Instance).
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// A list of simple strings assigned to the loadbalancer.
        /// Available only for Octavia **minor version 2.5 or later**.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Loadbalancer.  Only administrative users can specify a tenant UUID
        /// other than their own.  Changing this creates a new loadbalancer.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The ip address of the load balancer.
        /// Changing this creates a new loadbalancer.
        /// </summary>
        [Output("vipAddress")]
        public Output<string> VipAddress { get; private set; } = null!;

        /// <summary>
        /// The network on which to allocate the
        /// Loadbalancer's address. A tenant can only create Loadbalancers on networks
        /// authorized by policy (e.g. networks that belong to them or networks that
        /// are shared).  Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Output("vipNetworkId")]
        public Output<string> VipNetworkId { get; private set; } = null!;

        /// <summary>
        /// The port UUID that the loadbalancer will use.
        /// Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Output("vipPortId")]
        public Output<string> VipPortId { get; private set; } = null!;

        /// <summary>
        /// The ID of the QoS Policy which will
        /// be applied to the Virtual IP (VIP).
        /// </summary>
        [Output("vipQosPolicyId")]
        public Output<string?> VipQosPolicyId { get; private set; } = null!;

        /// <summary>
        /// The subnet on which to allocate the
        /// Loadbalancer's address. A tenant can only create Loadbalancers on networks
        /// authorized by policy (e.g. networks that belong to them or networks that
        /// are shared).  Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Output("vipSubnetId")]
        public Output<string> VipSubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a LbLoadbalancerV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LbLoadbalancerV2(string name, LbLoadbalancerV2Args? args = null, CustomResourceOptions? options = null)
            : base("openstack:index/lbLoadbalancerV2:LbLoadbalancerV2", name, args ?? new LbLoadbalancerV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private LbLoadbalancerV2(string name, Input<string> id, LbLoadbalancerV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:index/lbLoadbalancerV2:LbLoadbalancerV2", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "openstack:loadbalancer/loadBalancer:LoadBalancer" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LbLoadbalancerV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LbLoadbalancerV2 Get(string name, Input<string> id, LbLoadbalancerV2State? state = null, CustomResourceOptions? options = null)
        {
            return new LbLoadbalancerV2(name, id, state, options);
        }
    }

    public sealed class LbLoadbalancerV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the Loadbalancer.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The availability zone of the Loadbalancer.
        /// Changing this creates a new loadbalancer. Available only for Octavia
        /// **minor version 2.14 or later**.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Human-readable description for the Loadbalancer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The UUID of a flavor. Changing this creates a new
        /// loadbalancer.
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// The name of the provider. Changing this
        /// creates a new loadbalancer.
        /// </summary>
        [Input("loadbalancerProvider")]
        public Input<string>? LoadbalancerProvider { get; set; }

        /// <summary>
        /// Human-readable name for the Loadbalancer. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB member. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB member.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs to apply to the
        /// loadbalancer. The security groups must be specified by ID and not name (as
        /// opposed to how they are configured with the Compute Instance).
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of simple strings assigned to the loadbalancer.
        /// Available only for Octavia **minor version 2.5 or later**.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Loadbalancer.  Only administrative users can specify a tenant UUID
        /// other than their own.  Changing this creates a new loadbalancer.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The ip address of the load balancer.
        /// Changing this creates a new loadbalancer.
        /// </summary>
        [Input("vipAddress")]
        public Input<string>? VipAddress { get; set; }

        /// <summary>
        /// The network on which to allocate the
        /// Loadbalancer's address. A tenant can only create Loadbalancers on networks
        /// authorized by policy (e.g. networks that belong to them or networks that
        /// are shared).  Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Input("vipNetworkId")]
        public Input<string>? VipNetworkId { get; set; }

        /// <summary>
        /// The port UUID that the loadbalancer will use.
        /// Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Input("vipPortId")]
        public Input<string>? VipPortId { get; set; }

        /// <summary>
        /// The ID of the QoS Policy which will
        /// be applied to the Virtual IP (VIP).
        /// </summary>
        [Input("vipQosPolicyId")]
        public Input<string>? VipQosPolicyId { get; set; }

        /// <summary>
        /// The subnet on which to allocate the
        /// Loadbalancer's address. A tenant can only create Loadbalancers on networks
        /// authorized by policy (e.g. networks that belong to them or networks that
        /// are shared).  Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Input("vipSubnetId")]
        public Input<string>? VipSubnetId { get; set; }

        public LbLoadbalancerV2Args()
        {
        }
        public static new LbLoadbalancerV2Args Empty => new LbLoadbalancerV2Args();
    }

    public sealed class LbLoadbalancerV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the Loadbalancer.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The availability zone of the Loadbalancer.
        /// Changing this creates a new loadbalancer. Available only for Octavia
        /// **minor version 2.14 or later**.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Human-readable description for the Loadbalancer.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The UUID of a flavor. Changing this creates a new
        /// loadbalancer.
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// The name of the provider. Changing this
        /// creates a new loadbalancer.
        /// </summary>
        [Input("loadbalancerProvider")]
        public Input<string>? LoadbalancerProvider { get; set; }

        /// <summary>
        /// Human-readable name for the Loadbalancer. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB member. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB member.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security group IDs to apply to the
        /// loadbalancer. The security groups must be specified by ID and not name (as
        /// opposed to how they are configured with the Compute Instance).
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of simple strings assigned to the loadbalancer.
        /// Available only for Octavia **minor version 2.5 or later**.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the Loadbalancer.  Only administrative users can specify a tenant UUID
        /// other than their own.  Changing this creates a new loadbalancer.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The ip address of the load balancer.
        /// Changing this creates a new loadbalancer.
        /// </summary>
        [Input("vipAddress")]
        public Input<string>? VipAddress { get; set; }

        /// <summary>
        /// The network on which to allocate the
        /// Loadbalancer's address. A tenant can only create Loadbalancers on networks
        /// authorized by policy (e.g. networks that belong to them or networks that
        /// are shared).  Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Input("vipNetworkId")]
        public Input<string>? VipNetworkId { get; set; }

        /// <summary>
        /// The port UUID that the loadbalancer will use.
        /// Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Input("vipPortId")]
        public Input<string>? VipPortId { get; set; }

        /// <summary>
        /// The ID of the QoS Policy which will
        /// be applied to the Virtual IP (VIP).
        /// </summary>
        [Input("vipQosPolicyId")]
        public Input<string>? VipQosPolicyId { get; set; }

        /// <summary>
        /// The subnet on which to allocate the
        /// Loadbalancer's address. A tenant can only create Loadbalancers on networks
        /// authorized by policy (e.g. networks that belong to them or networks that
        /// are shared).  Changing this creates a new loadbalancer. Exactly one of
        /// `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
        /// </summary>
        [Input("vipSubnetId")]
        public Input<string>? VipSubnetId { get; set; }

        public LbLoadbalancerV2State()
        {
        }
        public static new LbLoadbalancerV2State Empty => new LbLoadbalancerV2State();
    }
}
