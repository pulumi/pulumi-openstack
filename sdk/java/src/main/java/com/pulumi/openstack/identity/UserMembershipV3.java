// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.identity.UserMembershipV3Args;
import com.pulumi.openstack.identity.inputs.UserMembershipV3State;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a user membership to group V3 resource within OpenStack.
 * 
 * &gt; **Note:** You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 * 
 * ***
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
 * import com.pulumi.openstack.identity.GroupV3;
 * import com.pulumi.openstack.identity.GroupV3Args;
 * import com.pulumi.openstack.identity.Role;
 * import com.pulumi.openstack.identity.RoleArgs;
 * import com.pulumi.openstack.identity.UserMembershipV3;
 * import com.pulumi.openstack.identity.UserMembershipV3Args;
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
 *         var group1 = new GroupV3("group1", GroupV3Args.builder()
 *             .name("group_1")
 *             .description("group 1")
 *             .build());
 * 
 *         var role1 = new Role("role1", RoleArgs.builder()
 *             .name("role_1")
 *             .build());
 * 
 *         var userMembership1 = new UserMembershipV3("userMembership1", UserMembershipV3Args.builder()
 *             .userId(user1.id())
 *             .groupId(group1.id())
 *             .build());
 * 
 *         var roleAssignment1 = new RoleAssignment("roleAssignment1", RoleAssignmentArgs.builder()
 *             .groupId(group1.id())
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
 * This resource can be imported by specifying all two arguments, separated
 * by a forward slash:
 * 
 * ```sh
 * $ pulumi import openstack:identity/userMembershipV3:UserMembershipV3 user_membership_1 user_id/group_id
 * ```
 * 
 */
@ResourceType(type="openstack:identity/userMembershipV3:UserMembershipV3")
public class UserMembershipV3 extends com.pulumi.resources.CustomResource {
    /**
     * The UUID of group to which the user will be added.
     * Changing this creates a new user membership.
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    /**
     * @return The UUID of group to which the user will be added.
     * Changing this creates a new user membership.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }
    /**
     * The region in which to obtain the V3 Identity client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new user membership.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V3 Identity client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new user membership.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The UUID of user to use. Changing this creates a new user membership.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return The UUID of user to use. Changing this creates a new user membership.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserMembershipV3(java.lang.String name) {
        this(name, UserMembershipV3Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserMembershipV3(java.lang.String name, UserMembershipV3Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserMembershipV3(java.lang.String name, UserMembershipV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/userMembershipV3:UserMembershipV3", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private UserMembershipV3(java.lang.String name, Output<java.lang.String> id, @Nullable UserMembershipV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/userMembershipV3:UserMembershipV3", name, state, makeResourceOptions(options, id), false);
    }

    private static UserMembershipV3Args makeArgs(UserMembershipV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UserMembershipV3Args.Empty : args;
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
    public static UserMembershipV3 get(java.lang.String name, Output<java.lang.String> id, @Nullable UserMembershipV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserMembershipV3(name, id, state, options);
    }
}
