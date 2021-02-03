// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity
{
    /// <summary>
    /// Manages a user membership to group V3 resource within OpenStack.
    /// 
    /// Note: You _must_ have admin privileges in your OpenStack cloud to use
    /// this resource.
    /// 
    /// ***
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
    ///         var project1 = new OpenStack.Identity.Project("project1", new OpenStack.Identity.ProjectArgs
    ///         {
    ///         });
    ///         var user1 = new OpenStack.Identity.User("user1", new OpenStack.Identity.UserArgs
    ///         {
    ///             DefaultProjectId = project1.Id,
    ///         });
    ///         var group1 = new OpenStack.Identity.GroupV3("group1", new OpenStack.Identity.GroupV3Args
    ///         {
    ///             Description = "group 1",
    ///         });
    ///         var role1 = new OpenStack.Identity.Role("role1", new OpenStack.Identity.RoleArgs
    ///         {
    ///         });
    ///         var userMembership1 = new OpenStack.Identity.UserMembershipV3("userMembership1", new OpenStack.Identity.UserMembershipV3Args
    ///         {
    ///             GroupId = group1.Id,
    ///             UserId = user1.Id,
    ///         });
    ///         var roleAssignment1 = new OpenStack.Identity.RoleAssignment("roleAssignment1", new OpenStack.Identity.RoleAssignmentArgs
    ///         {
    ///             GroupId = group1.Id,
    ///             ProjectId = project1.Id,
    ///             RoleId = role1.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported by specifying all two arguments, separated by a forward slash
    /// 
    /// ```sh
    ///  $ pulumi import openstack:identity/userMembershipV3:UserMembershipV3 user_membership_1 &lt;user_id&gt;/&lt;group_id&gt;
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:identity/userMembershipV3:UserMembershipV3")]
    public partial class UserMembershipV3 : Pulumi.CustomResource
    {
        /// <summary>
        /// The UUID of group to which the user will be added.
        /// Changing this creates a new user membership.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Identity client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new user membership.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The UUID of user to use. Changing this creates a new user membership.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserMembershipV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserMembershipV3(string name, UserMembershipV3Args args, CustomResourceOptions? options = null)
            : base("openstack:identity/userMembershipV3:UserMembershipV3", name, args ?? new UserMembershipV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private UserMembershipV3(string name, Input<string> id, UserMembershipV3State? state = null, CustomResourceOptions? options = null)
            : base("openstack:identity/userMembershipV3:UserMembershipV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserMembershipV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserMembershipV3 Get(string name, Input<string> id, UserMembershipV3State? state = null, CustomResourceOptions? options = null)
        {
            return new UserMembershipV3(name, id, state, options);
        }
    }

    public sealed class UserMembershipV3Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of group to which the user will be added.
        /// Changing this creates a new user membership.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Identity client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new user membership.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The UUID of user to use. Changing this creates a new user membership.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserMembershipV3Args()
        {
        }
    }

    public sealed class UserMembershipV3State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of group to which the user will be added.
        /// Changing this creates a new user membership.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Identity client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new user membership.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The UUID of user to use. Changing this creates a new user membership.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public UserMembershipV3State()
        {
        }
    }
}
