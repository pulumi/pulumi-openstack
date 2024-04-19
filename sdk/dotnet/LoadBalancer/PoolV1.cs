// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.LoadBalancer
{
    /// <summary>
    /// Manages a V1 load balancer pool resource within OpenStack.
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
    ///     var pool1 = new OpenStack.LoadBalancer.PoolV1("pool_1", new()
    ///     {
    ///         Name = "tf_test_lb_pool",
    ///         Protocol = "HTTP",
    ///         SubnetId = "12345",
    ///         LbMethod = "ROUND_ROBIN",
    ///         LbProvider = "haproxy",
    ///         MonitorIds = new[]
    ///         {
    ///             "67890",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Complete Load Balancing Stack Example
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
    ///     var subnet1 = new OpenStack.Networking.Subnet("subnet_1", new()
    ///     {
    ///         NetworkId = network1.Id,
    ///         Cidr = "192.168.199.0/24",
    ///         IpVersion = 4,
    ///     });
    /// 
    ///     var secgroup1 = new OpenStack.Compute.SecGroup("secgroup_1", new()
    ///     {
    ///         Name = "secgroup_1",
    ///         Description = "Rules for secgroup_1",
    ///         Rules = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.SecGroupRuleArgs
    ///             {
    ///                 FromPort = -1,
    ///                 ToPort = -1,
    ///                 IpProtocol = "icmp",
    ///                 Cidr = "0.0.0.0/0",
    ///             },
    ///             new OpenStack.Compute.Inputs.SecGroupRuleArgs
    ///             {
    ///                 FromPort = 80,
    ///                 ToPort = 80,
    ///                 IpProtocol = "tcp",
    ///                 Cidr = "0.0.0.0/0",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var instance1 = new OpenStack.Compute.Instance("instance_1", new()
    ///     {
    ///         Name = "instance_1",
    ///         SecurityGroups = new[]
    ///         {
    ///             "default",
    ///             secgroup1.Name,
    ///         },
    ///         Networks = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.InstanceNetworkArgs
    ///             {
    ///                 Uuid = network1.Id,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var instance2 = new OpenStack.Compute.Instance("instance_2", new()
    ///     {
    ///         Name = "instance_2",
    ///         SecurityGroups = new[]
    ///         {
    ///             "default",
    ///             secgroup1.Name,
    ///         },
    ///         Networks = new[]
    ///         {
    ///             new OpenStack.Compute.Inputs.InstanceNetworkArgs
    ///             {
    ///                 Uuid = network1.Id,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var monitor1 = new OpenStack.LoadBalancer.MonitorV1("monitor_1", new()
    ///     {
    ///         Type = "TCP",
    ///         Delay = 30,
    ///         Timeout = 5,
    ///         MaxRetries = 3,
    ///         AdminStateUp = "true",
    ///     });
    /// 
    ///     var pool1 = new OpenStack.LoadBalancer.PoolV1("pool_1", new()
    ///     {
    ///         Name = "pool_1",
    ///         Protocol = "TCP",
    ///         SubnetId = subnet1.Id,
    ///         LbMethod = "ROUND_ROBIN",
    ///         MonitorIds = new[]
    ///         {
    ///             monitor1.Id,
    ///         },
    ///     });
    /// 
    ///     var member1 = new OpenStack.LoadBalancer.MemberV1("member_1", new()
    ///     {
    ///         PoolId = pool1.Id,
    ///         Address = instance1.AccessIpV4,
    ///         Port = 80,
    ///     });
    /// 
    ///     var member2 = new OpenStack.LoadBalancer.MemberV1("member_2", new()
    ///     {
    ///         PoolId = pool1.Id,
    ///         Address = instance2.AccessIpV4,
    ///         Port = 80,
    ///     });
    /// 
    ///     var vip1 = new OpenStack.LoadBalancer.Vip("vip_1", new()
    ///     {
    ///         Name = "vip_1",
    ///         SubnetId = subnet1.Id,
    ///         Protocol = "TCP",
    ///         Port = 80,
    ///         PoolId = pool1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Notes
    /// 
    /// The `member` block is deprecated in favor of the `openstack.loadbalancer.MemberV1` resource.
    /// 
    /// ## Import
    /// 
    /// Load Balancer Pools can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:loadbalancer/poolV1:PoolV1 pool_1 b255e6ba-02ad-43e6-8951-3428ca26b713
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:loadbalancer/poolV1:PoolV1")]
    public partial class PoolV1 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The algorithm used to distribute load between the
        /// members of the pool. The current specification supports 'ROUND_ROBIN' and
        /// 'LEAST_CONNECTIONS' as valid values for this attribute.
        /// </summary>
        [Output("lbMethod")]
        public Output<string> LbMethod { get; private set; } = null!;

        /// <summary>
        /// The backend load balancing provider. For example:
        /// `haproxy`, `F5`, etc.
        /// </summary>
        [Output("lbProvider")]
        public Output<string> LbProvider { get; private set; } = null!;

        /// <summary>
        /// An existing node to add to the pool. Changing this
        /// updates the members of the pool. The member object structure is documented
        /// below. Please note that the `member` block is deprecated in favor of the
        /// `openstack.loadbalancer.MemberV1` resource.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// A list of IDs of monitors to associate with the
        /// pool.
        /// </summary>
        [Output("monitorIds")]
        public Output<ImmutableArray<string>> MonitorIds { get; private set; } = null!;

        /// <summary>
        /// The name of the pool. Changing this updates the name of
        /// the existing pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The protocol used by the pool members, you can use
        /// either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB pool. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB pool.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The network on which the members of the pool will be
        /// located. Only members that are on this network can be added to the pool.
        /// Changing this creates a new pool.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// The owner of the pool. Required if admin wants to
        /// create a pool member for another tenant. Changing this creates a new pool.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a PoolV1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PoolV1(string name, PoolV1Args args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/poolV1:PoolV1", name, args ?? new PoolV1Args(), MakeResourceOptions(options, ""))
        {
        }

        private PoolV1(string name, Input<string> id, PoolV1State? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/poolV1:PoolV1", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PoolV1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PoolV1 Get(string name, Input<string> id, PoolV1State? state = null, CustomResourceOptions? options = null)
        {
            return new PoolV1(name, id, state, options);
        }
    }

    public sealed class PoolV1Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm used to distribute load between the
        /// members of the pool. The current specification supports 'ROUND_ROBIN' and
        /// 'LEAST_CONNECTIONS' as valid values for this attribute.
        /// </summary>
        [Input("lbMethod", required: true)]
        public Input<string> LbMethod { get; set; } = null!;

        /// <summary>
        /// The backend load balancing provider. For example:
        /// `haproxy`, `F5`, etc.
        /// </summary>
        [Input("lbProvider")]
        public Input<string>? LbProvider { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// An existing node to add to the pool. Changing this
        /// updates the members of the pool. The member object structure is documented
        /// below. Please note that the `member` block is deprecated in favor of the
        /// `openstack.loadbalancer.MemberV1` resource.
        /// </summary>
        [Obsolete(@"Use openstack.loadbalancer.MemberV1 instead")]
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("monitorIds")]
        private InputList<string>? _monitorIds;

        /// <summary>
        /// A list of IDs of monitors to associate with the
        /// pool.
        /// </summary>
        public InputList<string> MonitorIds
        {
            get => _monitorIds ?? (_monitorIds = new InputList<string>());
            set => _monitorIds = value;
        }

        /// <summary>
        /// The name of the pool. Changing this updates the name of
        /// the existing pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol used by the pool members, you can use
        /// either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB pool. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB pool.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The network on which the members of the pool will be
        /// located. Only members that are on this network can be added to the pool.
        /// Changing this creates a new pool.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// The owner of the pool. Required if admin wants to
        /// create a pool member for another tenant. Changing this creates a new pool.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public PoolV1Args()
        {
        }
        public static new PoolV1Args Empty => new PoolV1Args();
    }

    public sealed class PoolV1State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm used to distribute load between the
        /// members of the pool. The current specification supports 'ROUND_ROBIN' and
        /// 'LEAST_CONNECTIONS' as valid values for this attribute.
        /// </summary>
        [Input("lbMethod")]
        public Input<string>? LbMethod { get; set; }

        /// <summary>
        /// The backend load balancing provider. For example:
        /// `haproxy`, `F5`, etc.
        /// </summary>
        [Input("lbProvider")]
        public Input<string>? LbProvider { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// An existing node to add to the pool. Changing this
        /// updates the members of the pool. The member object structure is documented
        /// below. Please note that the `member` block is deprecated in favor of the
        /// `openstack.loadbalancer.MemberV1` resource.
        /// </summary>
        [Obsolete(@"Use openstack.loadbalancer.MemberV1 instead")]
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        [Input("monitorIds")]
        private InputList<string>? _monitorIds;

        /// <summary>
        /// A list of IDs of monitors to associate with the
        /// pool.
        /// </summary>
        public InputList<string> MonitorIds
        {
            get => _monitorIds ?? (_monitorIds = new InputList<string>());
            set => _monitorIds = value;
        }

        /// <summary>
        /// The name of the pool. Changing this updates the name of
        /// the existing pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol used by the pool members, you can use
        /// either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB pool. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB pool.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The network on which the members of the pool will be
        /// located. Only members that are on this network can be added to the pool.
        /// Changing this creates a new pool.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The owner of the pool. Required if admin wants to
        /// create a pool member for another tenant. Changing this creates a new pool.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public PoolV1State()
        {
        }
        public static new PoolV1State Empty => new PoolV1State();
    }
}
