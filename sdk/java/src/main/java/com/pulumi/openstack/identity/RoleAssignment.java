// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.identity.RoleAssignmentArgs;
import com.pulumi.openstack.identity.inputs.RoleAssignmentState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V3 Role assignment within OpenStack Keystone.
 * 
 * &gt; **Note:** You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.identity.Project;
 * import com.pulumi.openstack.identity.ProjectArgs;
 * import com.pulumi.openstack.identity.User;
 * import com.pulumi.openstack.identity.UserArgs;
 * import com.pulumi.openstack.identity.Role;
 * import com.pulumi.openstack.identity.RoleArgs;
 * import com.pulumi.openstack.identity.RoleAssignment;
 * import com.pulumi.openstack.identity.RoleAssignmentArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var project1 = new Project("project1", ProjectArgs.builder()
 *             .name("project_1")
 *             .build());
 * 
 *         var user1 = new User("user1", UserArgs.builder()
 *             .name("user_1")
 *             .defaultProjectId(project1.id())
 *             .build());
 * 
 *         var role1 = new Role("role1", RoleArgs.builder()
 *             .name("role_1")
 *             .build());
 * 
 *         var roleAssignment1 = new RoleAssignment("roleAssignment1", RoleAssignmentArgs.builder()
 *             .userId(user1.id())
 *             .projectId(project1.id())
 *             .roleId(role1.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Role assignments can be imported using a constructed id. The id should have the form of
 * `domainID/projectID/groupID/userID/roleID`. When something is not used then leave blank.
 * 
 * For example this will import the role assignment for:
 * projectID: 014395cd-89fc-4c9b-96b7-13d1ee79dad2,
 * userID: 4142e64b-1b35-44a0-9b1e-5affc7af1106,
 * roleID: ea257959-eeb1-4c10-8d33-26f0409a755d
 * ( domainID and groupID are left blank)
 * 
 * ```sh
 * $ pulumi import openstack:identity/roleAssignment:RoleAssignment role_assignment_1 /014395cd-89fc-4c9b-96b7-13d1ee79dad2//4142e64b-1b35-44a0-9b1e-5affc7af1106/ea257959-eeb1-4c10-8d33-26f0409a755d
 * ```
 * 
 */
@ResourceType(type="openstack:identity/roleAssignment:RoleAssignment")
public class RoleAssignment extends com.pulumi.resources.CustomResource {
    /**
     * The domain to assign the role in.
     * 
     */
    @Export(name="domainId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> domainId;

    /**
     * @return The domain to assign the role in.
     * 
     */
    public Output<Optional<String>> domainId() {
        return Codegen.optional(this.domainId);
    }
    /**
     * The group to assign the role to.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupId;

    /**
     * @return The group to assign the role to.
     * 
     */
    public Output<Optional<String>> groupId() {
        return Codegen.optional(this.groupId);
    }
    /**
     * The project to assign the role in.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> projectId;

    /**
     * @return The project to assign the role in.
     * 
     */
    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new role assignment.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new role assignment.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The role to assign.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The role to assign.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }
    /**
     * The user to assign the role to.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userId;

    /**
     * @return The user to assign the role to.
     * 
     */
    public Output<Optional<String>> userId() {
        return Codegen.optional(this.userId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RoleAssignment(java.lang.String name) {
        this(name, RoleAssignmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RoleAssignment(java.lang.String name, RoleAssignmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RoleAssignment(java.lang.String name, RoleAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/roleAssignment:RoleAssignment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RoleAssignment(java.lang.String name, Output<java.lang.String> id, @Nullable RoleAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/roleAssignment:RoleAssignment", name, state, makeResourceOptions(options, id), false);
    }

    private static RoleAssignmentArgs makeArgs(RoleAssignmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RoleAssignmentArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RoleAssignment get(java.lang.String name, Output<java.lang.String> id, @Nullable RoleAssignmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RoleAssignment(name, id, state, options);
    }
}
