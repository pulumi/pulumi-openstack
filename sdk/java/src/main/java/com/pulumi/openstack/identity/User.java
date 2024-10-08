// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.identity.UserArgs;
import com.pulumi.openstack.identity.inputs.UserState;
import com.pulumi.openstack.identity.outputs.UserMultiFactorAuthRule;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V3 User resource within OpenStack Keystone.
 * 
 * &gt; **Note:** All arguments including the user password will be stored in the
 * raw state as plain-text. Read more about sensitive data in
 * state.
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
 * import com.pulumi.openstack.identity.inputs.UserMultiFactorAuthRuleArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var project1 = new Project("project1", ProjectArgs.builder()
 *             .name("project_1")
 *             .build());
 * 
 *         var user1 = new User("user1", UserArgs.builder()
 *             .defaultProjectId(project1.id())
 *             .name("user_1")
 *             .description("A user")
 *             .password("password123")
 *             .ignoreChangePasswordUponFirstUse(true)
 *             .multiFactorAuthEnabled(true)
 *             .multiFactorAuthRules(            
 *                 UserMultiFactorAuthRuleArgs.builder()
 *                     .rules(                    
 *                         "password",
 *                         "totp")
 *                     .build(),
 *                 UserMultiFactorAuthRuleArgs.builder()
 *                     .rules("password")
 *                     .build())
 *             .extra(Map.of("email", "user_1}{@literal @}{@code foobar.com"))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Users can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:identity/user:User user_1 89c60255-9bd6-460c-822a-e2b959ede9d2
 * ```
 * 
 */
@ResourceType(type="openstack:identity/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * The default project this user belongs to.
     * 
     */
    @Export(name="defaultProjectId", refs={String.class}, tree="[0]")
    private Output<String> defaultProjectId;

    /**
     * @return The default project this user belongs to.
     * 
     */
    public Output<String> defaultProjectId() {
        return this.defaultProjectId;
    }
    /**
     * A description of the user.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the user.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The domain this user belongs to.
     * 
     */
    @Export(name="domainId", refs={String.class}, tree="[0]")
    private Output<String> domainId;

    /**
     * @return The domain this user belongs to.
     * 
     */
    public Output<String> domainId() {
        return this.domainId;
    }
    /**
     * Whether the user is enabled or disabled. Valid
     * values are `true` and `false`.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether the user is enabled or disabled. Valid
     * values are `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * Free-form key/value pairs of extra information.
     * 
     */
    @Export(name="extra", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> extra;

    /**
     * @return Free-form key/value pairs of extra information.
     * 
     */
    public Output<Optional<Map<String,String>>> extra() {
        return Codegen.optional(this.extra);
    }
    /**
     * User will not have to
     * change their password upon first use. Valid values are `true` and `false`.
     * 
     */
    @Export(name="ignoreChangePasswordUponFirstUse", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreChangePasswordUponFirstUse;

    /**
     * @return User will not have to
     * change their password upon first use. Valid values are `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> ignoreChangePasswordUponFirstUse() {
        return Codegen.optional(this.ignoreChangePasswordUponFirstUse);
    }
    /**
     * User will not have a failure
     * lockout placed on their account. Valid values are `true` and `false`.
     * 
     */
    @Export(name="ignoreLockoutFailureAttempts", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreLockoutFailureAttempts;

    /**
     * @return User will not have a failure
     * lockout placed on their account. Valid values are `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> ignoreLockoutFailureAttempts() {
        return Codegen.optional(this.ignoreLockoutFailureAttempts);
    }
    /**
     * User&#39;s password will not expire.
     * Valid values are `true` and `false`.
     * 
     */
    @Export(name="ignorePasswordExpiry", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignorePasswordExpiry;

    /**
     * @return User&#39;s password will not expire.
     * Valid values are `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> ignorePasswordExpiry() {
        return Codegen.optional(this.ignorePasswordExpiry);
    }
    /**
     * Whether to enable multi-factor
     * authentication. Valid values are `true` and `false`.
     * 
     */
    @Export(name="multiFactorAuthEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> multiFactorAuthEnabled;

    /**
     * @return Whether to enable multi-factor
     * authentication. Valid values are `true` and `false`.
     * 
     */
    public Output<Optional<Boolean>> multiFactorAuthEnabled() {
        return Codegen.optional(this.multiFactorAuthEnabled);
    }
    /**
     * A multi-factor authentication rule.
     * The structure is documented below. Please see the
     * [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
     * for more information on how to use mulit-factor rules.
     * 
     */
    @Export(name="multiFactorAuthRules", refs={List.class,UserMultiFactorAuthRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<UserMultiFactorAuthRule>> multiFactorAuthRules;

    /**
     * @return A multi-factor authentication rule.
     * The structure is documented below. Please see the
     * [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
     * for more information on how to use mulit-factor rules.
     * 
     */
    public Output<Optional<List<UserMultiFactorAuthRule>>> multiFactorAuthRules() {
        return Codegen.optional(this.multiFactorAuthRules);
    }
    /**
     * The name of the user.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the user.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The password for the user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The password for the user.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new User.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new User.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(java.lang.String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(java.lang.String name, @Nullable UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(java.lang.String name, @Nullable UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/user:User", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private User(java.lang.String name, Output<java.lang.String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/user:User", name, state, makeResourceOptions(options, id), false);
    }

    private static UserArgs makeArgs(@Nullable UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UserArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static User get(java.lang.String name, Output<java.lang.String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
