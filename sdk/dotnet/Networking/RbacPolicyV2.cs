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
    /// The RBAC policy resource contains functionality for working with Neutron RBAC
    /// Policies. Role-Based Access Control (RBAC) policy framework enables both
    /// operators and users to grant access to resources for specific projects.
    /// 
    /// Sharing an object with a specific project is accomplished by creating a
    /// policy entry that permits the target project the `access_as_shared` action
    /// on that object.
    /// 
    /// To make a network available as an external network for specific projects
    /// rather than all projects, use the `access_as_external` action.
    /// If a network is marked as external during creation, it now implicitly creates
    /// a wildcard RBAC policy granting everyone access to preserve previous behavior
    /// before this feature was added.
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
    ///         var network1 = new OpenStack.Networking.Network("network1", new OpenStack.Networking.NetworkArgs
    ///         {
    ///             AdminStateUp = true,
    ///         });
    ///         var rbacPolicy1 = new OpenStack.Networking.RbacPolicyV2("rbacPolicy1", new OpenStack.Networking.RbacPolicyV2Args
    ///         {
    ///             Action = "access_as_shared",
    ///             ObjectId = network1.Id,
    ///             ObjectType = "network",
    ///             TargetTenant = "20415a973c9e45d3917f078950644697",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// RBAC policies can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:networking/rbacPolicyV2:RbacPolicyV2 rbac_policy_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/rbacPolicyV2:RbacPolicyV2")]
    public partial class RbacPolicyV2 : Pulumi.CustomResource
    {
        /// <summary>
        /// Action for the RBAC policy. Can either be
        /// `access_as_external` or `access_as_shared`.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The ID of the `object_type` resource. An
        /// `object_type` of `network` returns a network ID and an `object_type` of
        /// `qos_policy` returns a QoS ID.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// The type of the object that the RBAC policy
        /// affects. Can either be `qos-policy` or `network`.
        /// </summary>
        [Output("objectType")]
        public Output<string> ObjectType { get; private set; } = null!;

        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to configure a routing entry on a subnet. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// routing entry.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The ID of the tenant to which the RBAC policy
        /// will be enforced.
        /// </summary>
        [Output("targetTenant")]
        public Output<string> TargetTenant { get; private set; } = null!;


        /// <summary>
        /// Create a RbacPolicyV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RbacPolicyV2(string name, RbacPolicyV2Args args, CustomResourceOptions? options = null)
            : base("openstack:networking/rbacPolicyV2:RbacPolicyV2", name, args ?? new RbacPolicyV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private RbacPolicyV2(string name, Input<string> id, RbacPolicyV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/rbacPolicyV2:RbacPolicyV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RbacPolicyV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RbacPolicyV2 Get(string name, Input<string> id, RbacPolicyV2State? state = null, CustomResourceOptions? options = null)
        {
            return new RbacPolicyV2(name, id, state, options);
        }
    }

    public sealed class RbacPolicyV2Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action for the RBAC policy. Can either be
        /// `access_as_external` or `access_as_shared`.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The ID of the `object_type` resource. An
        /// `object_type` of `network` returns a network ID and an `object_type` of
        /// `qos_policy` returns a QoS ID.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        /// <summary>
        /// The type of the object that the RBAC policy
        /// affects. Can either be `qos-policy` or `network`.
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to configure a routing entry on a subnet. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// routing entry.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the tenant to which the RBAC policy
        /// will be enforced.
        /// </summary>
        [Input("targetTenant", required: true)]
        public Input<string> TargetTenant { get; set; } = null!;

        public RbacPolicyV2Args()
        {
        }
    }

    public sealed class RbacPolicyV2State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action for the RBAC policy. Can either be
        /// `access_as_external` or `access_as_shared`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The ID of the `object_type` resource. An
        /// `object_type` of `network` returns a network ID and an `object_type` of
        /// `qos_policy` returns a QoS ID.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// The type of the object that the RBAC policy
        /// affects. Can either be `qos-policy` or `network`.
        /// </summary>
        [Input("objectType")]
        public Input<string>? ObjectType { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to configure a routing entry on a subnet. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// routing entry.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the tenant to which the RBAC policy
        /// will be enforced.
        /// </summary>
        [Input("targetTenant")]
        public Input<string>? TargetTenant { get; set; }

        public RbacPolicyV2State()
        {
        }
    }
}
