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
    /// Manages a V3 Role assignment within OpenStack Keystone.
    /// 
    /// Note: You _must_ have admin privileges in your OpenStack cloud to use
    /// this resource.
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
    ///         var role1 = new OpenStack.Identity.Role("role1", new OpenStack.Identity.RoleArgs
    ///         {
    ///         });
    ///         var roleAssignment1 = new OpenStack.Identity.RoleAssignment("roleAssignment1", new OpenStack.Identity.RoleAssignmentArgs
    ///         {
    ///             ProjectId = project1.Id,
    ///             RoleId = role1.Id,
    ///             UserId = user1.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:identity/roleAssignment:RoleAssignment")]
    public partial class RoleAssignment : Pulumi.CustomResource
    {
        /// <summary>
        /// The domain to assign the role in.
        /// </summary>
        [Output("domainId")]
        public Output<string?> DomainId { get; private set; } = null!;

        /// <summary>
        /// The group to assign the role to.
        /// </summary>
        [Output("groupId")]
        public Output<string?> GroupId { get; private set; } = null!;

        /// <summary>
        /// The project to assign the role in.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The role to assign.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;

        /// <summary>
        /// The user to assign the role to.
        /// </summary>
        [Output("userId")]
        public Output<string?> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a RoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoleAssignment(string name, RoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("openstack:identity/roleAssignment:RoleAssignment", name, args ?? new RoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoleAssignment(string name, Input<string> id, RoleAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("openstack:identity/roleAssignment:RoleAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoleAssignment Get(string name, Input<string> id, RoleAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new RoleAssignment(name, id, state, options);
        }
    }

    public sealed class RoleAssignmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to assign the role in.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// The group to assign the role to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The project to assign the role in.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role to assign.
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        /// <summary>
        /// The user to assign the role to.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public RoleAssignmentArgs()
        {
        }
    }

    public sealed class RoleAssignmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to assign the role in.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// The group to assign the role to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The project to assign the role in.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role to assign.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// The user to assign the role to.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public RoleAssignmentState()
        {
        }
    }
}
