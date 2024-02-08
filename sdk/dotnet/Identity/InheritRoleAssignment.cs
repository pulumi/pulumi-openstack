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
    /// Manages a V3 Inherit Role assignment within OpenStack Keystone. This uses the
    /// Openstack keystone `OS-INHERIT` api to created inherit roles within domains
    /// and parent projects for users and groups.
    /// 
    /// &gt; **Note:** You _must_ have admin privileges in your OpenStack cloud to use
    /// this resource.
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
    ///     var user1 = new OpenStack.Identity.User("user1", new()
    ///     {
    ///         DomainId = "default",
    ///     });
    /// 
    ///     var role1 = new OpenStack.Identity.Role("role1", new()
    ///     {
    ///         DomainId = "default",
    ///     });
    /// 
    ///     var roleAssignment1 = new OpenStack.Identity.InheritRoleAssignment("roleAssignment1", new()
    ///     {
    ///         UserId = user1.Id,
    ///         DomainId = "default",
    ///         RoleId = role1.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Inherit role assignments can be imported using a constructed id. The id should have the form of `domainID/projectID/groupID/userID/roleID`. When something is not used then leave blank.
    /// 
    /// For example this will import the inherit role assignment for: projectID: 014395cd-89fc-4c9b-96b7-13d1ee79dad2, userID: 4142e64b-1b35-44a0-9b1e-5affc7af1106, roleID: ea257959-eeb1-4c10-8d33-26f0409a755d ( domainID and groupID are left blank)
    /// 
    /// ```sh
    ///  $ pulumi import openstack:identity/inheritRoleAssignment:InheritRoleAssignment role_assignment_1 /014395cd-89fc-4c9b-96b7-13d1ee79dad2//4142e64b-1b35-44a0-9b1e-5affc7af1106/ea257959-eeb1-4c10-8d33-26f0409a755d
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:identity/inheritRoleAssignment:InheritRoleAssignment")]
    public partial class InheritRoleAssignment : global::Pulumi.CustomResource
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
        /// The project should be able to containt child projects.
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
        /// Create a InheritRoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InheritRoleAssignment(string name, InheritRoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("openstack:identity/inheritRoleAssignment:InheritRoleAssignment", name, args ?? new InheritRoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InheritRoleAssignment(string name, Input<string> id, InheritRoleAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("openstack:identity/inheritRoleAssignment:InheritRoleAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InheritRoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InheritRoleAssignment Get(string name, Input<string> id, InheritRoleAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InheritRoleAssignment(name, id, state, options);
        }
    }

    public sealed class InheritRoleAssignmentArgs : global::Pulumi.ResourceArgs
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
        /// The project should be able to containt child projects.
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

        public InheritRoleAssignmentArgs()
        {
        }
        public static new InheritRoleAssignmentArgs Empty => new InheritRoleAssignmentArgs();
    }

    public sealed class InheritRoleAssignmentState : global::Pulumi.ResourceArgs
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
        /// The project should be able to containt child projects.
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

        public InheritRoleAssignmentState()
        {
        }
        public static new InheritRoleAssignmentState Empty => new InheritRoleAssignmentState();
    }
}
