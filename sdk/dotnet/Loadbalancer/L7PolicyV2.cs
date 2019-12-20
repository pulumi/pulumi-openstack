// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Loadbalancer
{
    /// <summary>
    /// Manages a Load Balancer L7 Policy resource within OpenStack.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_l7policy_v2.html.markdown.
    /// </summary>
    public partial class L7PolicyV2 : Pulumi.CustomResource
    {
        /// <summary>
        /// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
        /// REDIRECT\_TO\_URL or REJECT.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The administrative state of the L7 Policy.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Output("adminStateUp")]
        public Output<bool?> AdminStateUp { get; private set; } = null!;

        /// <summary>
        /// Human-readable description for the L7 Policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Listener on which the L7 Policy will be associated with.
        /// Changing this creates a new L7 Policy.
        /// </summary>
        [Output("listenerId")]
        public Output<string> ListenerId { get; private set; } = null!;

        /// <summary>
        /// Human-readable name for the L7 Policy. Does not have
        /// to be unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The position of this policy on the listener. Positions start at 1.
        /// </summary>
        [Output("position")]
        public Output<int> Position { get; private set; } = null!;

        /// <summary>
        /// Requests matching this policy will be redirected to the
        /// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
        /// </summary>
        [Output("redirectPoolId")]
        public Output<string?> RedirectPoolId { get; private set; } = null!;

        /// <summary>
        /// Requests matching this policy will be redirected to this URL.
        /// Only valid if action is REDIRECT\_TO\_URL.
        /// </summary>
        [Output("redirectUrl")]
        public Output<string?> RedirectUrl { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// L7 Policy.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the L7 Policy.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new L7 Policy.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a L7PolicyV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public L7PolicyV2(string name, L7PolicyV2Args args, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/l7PolicyV2:L7PolicyV2", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private L7PolicyV2(string name, Input<string> id, L7PolicyV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:loadbalancer/l7PolicyV2:L7PolicyV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing L7PolicyV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static L7PolicyV2 Get(string name, Input<string> id, L7PolicyV2State? state = null, CustomResourceOptions? options = null)
        {
            return new L7PolicyV2(name, id, state, options);
        }
    }

    public sealed class L7PolicyV2Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
        /// REDIRECT\_TO\_URL or REJECT.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The administrative state of the L7 Policy.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// Human-readable description for the L7 Policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Listener on which the L7 Policy will be associated with.
        /// Changing this creates a new L7 Policy.
        /// </summary>
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        /// <summary>
        /// Human-readable name for the L7 Policy. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The position of this policy on the listener. Positions start at 1.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Requests matching this policy will be redirected to the
        /// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
        /// </summary>
        [Input("redirectPoolId")]
        public Input<string>? RedirectPoolId { get; set; }

        /// <summary>
        /// Requests matching this policy will be redirected to this URL.
        /// Only valid if action is REDIRECT\_TO\_URL.
        /// </summary>
        [Input("redirectUrl")]
        public Input<string>? RedirectUrl { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// L7 Policy.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the L7 Policy.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new L7 Policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public L7PolicyV2Args()
        {
        }
    }

    public sealed class L7PolicyV2State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
        /// REDIRECT\_TO\_URL or REJECT.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The administrative state of the L7 Policy.
        /// A valid value is true (UP) or false (DOWN).
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// Human-readable description for the L7 Policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Listener on which the L7 Policy will be associated with.
        /// Changing this creates a new L7 Policy.
        /// </summary>
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        /// <summary>
        /// Human-readable name for the L7 Policy. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The position of this policy on the listener. Positions start at 1.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Requests matching this policy will be redirected to the
        /// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
        /// </summary>
        [Input("redirectPoolId")]
        public Input<string>? RedirectPoolId { get; set; }

        /// <summary>
        /// Requests matching this policy will be redirected to this URL.
        /// Only valid if action is REDIRECT\_TO\_URL.
        /// </summary>
        [Input("redirectUrl")]
        public Input<string>? RedirectUrl { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an . If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// L7 Policy.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Required for admins. The UUID of the tenant who owns
        /// the L7 Policy.  Only administrative users can specify a tenant UUID
        /// other than their own. Changing this creates a new L7 Policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public L7PolicyV2State()
        {
        }
    }
}
