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
    /// Manages a V1 load balancer member resource within OpenStack.
    /// </summary>
    public partial class MemberV1 : Pulumi.CustomResource
    {
        /// <summary>
        /// The IP address of the member. Changing this creates a
        /// new member.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The administrative state of the member.
        /// Acceptable values are 'true' and 'false'. Changing this value updates the
        /// state of the existing member.
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// The ID of the LB pool. Changing this creates a new
        /// member.
        /// </summary>
        [Output("poolId")]
        public Output<string> PoolId { get; private set; } = null!;

        /// <summary>
        /// An integer representing the port on which the member is
        /// hosted. Changing this creates a new member.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB member. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB member.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The owner of the member. Required if admin wants to
        /// create a member for another tenant. Changing this creates a new member.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;

        [Output("weight")]
        public Output<int> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a MemberV1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MemberV1(string name, MemberV1Args args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/memberV1:MemberV1", name, args ?? new MemberV1Args(), MakeResourceOptions(options, ""))
        {
        }

        private MemberV1(string name, Input<string> id, MemberV1State? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/memberV1:MemberV1", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MemberV1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MemberV1 Get(string name, Input<string> id, MemberV1State? state = null, CustomResourceOptions? options = null)
        {
            return new MemberV1(name, id, state, options);
        }
    }

    public sealed class MemberV1Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address of the member. Changing this creates a
        /// new member.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// The administrative state of the member.
        /// Acceptable values are 'true' and 'false'. Changing this value updates the
        /// state of the existing member.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The ID of the LB pool. Changing this creates a new
        /// member.
        /// </summary>
        [Input("poolId", required: true)]
        public Input<string> PoolId { get; set; } = null!;

        /// <summary>
        /// An integer representing the port on which the member is
        /// hosted. Changing this creates a new member.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB member. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB member.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the member. Required if admin wants to
        /// create a member for another tenant. Changing this creates a new member.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public MemberV1Args()
        {
        }
    }

    public sealed class MemberV1State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address of the member. Changing this creates a
        /// new member.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The administrative state of the member.
        /// Acceptable values are 'true' and 'false'. Changing this value updates the
        /// state of the existing member.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// The ID of the LB pool. Changing this creates a new
        /// member.
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        /// <summary>
        /// An integer representing the port on which the member is
        /// hosted. Changing this creates a new member.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an LB member. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// LB member.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the member. Required if admin wants to
        /// create a member for another tenant. Changing this creates a new member.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public MemberV1State()
        {
        }
    }
}
