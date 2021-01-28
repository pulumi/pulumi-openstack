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
    /// Manages a V2 members resource within OpenStack (batch members update).
    /// 
    /// &gt; **Note:** This resource works only within Octavia API. For
    /// legacy Neutron LBaaS v2 extension please use
    /// openstack.loadbalancer.Member resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var members1 = new OpenStack.LoadBalancer.Members("members1", new OpenStack.LoadBalancer.MembersArgs
    ///         {
    ///             Members = 
    ///             {
    ///                 new OpenStack.LoadBalancer.Inputs.MembersMemberArgs
    ///                 {
    ///                     Address = "192.168.199.23",
    ///                     ProtocolPort = 8080,
    ///                 },
    ///                 new OpenStack.LoadBalancer.Inputs.MembersMemberArgs
    ///                 {
    ///                     Address = "192.168.199.24",
    ///                     ProtocolPort = 8080,
    ///                 },
    ///             },
    ///             PoolId = "935685fb-a896-40f9-9ff4-ae531a3a00fe",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Load Balancer Pool Members can be imported using the Pool ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:loadbalancer/members:Members members_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:loadbalancer/members:Members")]
    public partial class Members : Pulumi.CustomResource
    {
        /// <summary>
        /// A set of dictionaries containing member parameters. The
        /// structure is described below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.MembersMember>> MemberList { get; private set; } = null!;

        /// <summary>
        /// The id of the pool that members will be assigned to.
        /// Changing this creates a new members resource.
        /// </summary>
        [Output("poolId")]
        public Output<string> PoolId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create pool members. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// members resource.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a Members resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Members(string name, MembersArgs args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/members:Members", name, args ?? new MembersArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Members(string name, Input<string> id, MembersState? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/members:Members", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Members resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Members Get(string name, Input<string> id, MembersState? state = null, CustomResourceOptions? options = null)
        {
            return new Members(name, id, state, options);
        }
    }

    public sealed class MembersArgs : Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<Inputs.MembersMemberArgs>? _members;

        /// <summary>
        /// A set of dictionaries containing member parameters. The
        /// structure is described below.
        /// </summary>
        public InputList<Inputs.MembersMemberArgs> MemberList
        {
            get => _members ?? (_members = new InputList<Inputs.MembersMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// The id of the pool that members will be assigned to.
        /// Changing this creates a new members resource.
        /// </summary>
        [Input("poolId", required: true)]
        public Input<string> PoolId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create pool members. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// members resource.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MembersArgs()
        {
        }
    }

    public sealed class MembersState : Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<Inputs.MembersMemberGetArgs>? _members;

        /// <summary>
        /// A set of dictionaries containing member parameters. The
        /// structure is described below.
        /// </summary>
        public InputList<Inputs.MembersMemberGetArgs> MemberList
        {
            get => _members ?? (_members = new InputList<Inputs.MembersMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// The id of the pool that members will be assigned to.
        /// Changing this creates a new members resource.
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create pool members. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// members resource.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MembersState()
        {
        }
    }
}
