// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.identity.ApplicationCredentialArgs;
import com.pulumi.openstack.identity.inputs.ApplicationCredentialState;
import com.pulumi.openstack.identity.outputs.ApplicationCredentialAccessRule;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V3 Application Credential resource within OpenStack Keystone.
 * 
 * &gt; **Note:** All arguments including the application credential name and secret
 * will be stored in the raw state as plain-text. Read more about sensitive data
 * in state.
 * 
 * &gt; **Note:** An Application Credential is created within the authenticated user
 * project scope and is not visible by an admin or other accounts.
 * The Application Credential visibility is similar to
 * `openstack.compute.Keypair`.
 * 
 * ## Example Usage
 * 
 * ### Predefined secret
 * 
 * Application credential below will have only one `swiftoperator` role.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.identity.ApplicationCredential;
 * import com.pulumi.openstack.identity.ApplicationCredentialArgs;
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
 *         var swift = new ApplicationCredential("swift", ApplicationCredentialArgs.builder()
 *             .name("swift")
 *             .description("Swift technical application credential")
 *             .secret("supersecret")
 *             .roles("swiftoperator")
 *             .expiresAt("2019-02-13T12:12:12Z")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Unrestricted with autogenerated secret and unlimited TTL
 * 
 * Application credential below will inherit all the current user&#39;s roles.
 * 
 * !&gt; **WARNING:** Restrictions on these Identity operations are deliberately
 * imposed as a safeguard to prevent a compromised application credential from
 * regenerating itself. Disabling this restriction poses an inherent added risk.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.identity.ApplicationCredential;
 * import com.pulumi.openstack.identity.ApplicationCredentialArgs;
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
 *         var unrestricted = new ApplicationCredential("unrestricted", ApplicationCredentialArgs.builder()
 *             .name("unrestricted")
 *             .description("Unrestricted application credential")
 *             .unrestricted(true)
 *             .build());
 * 
 *         ctx.export("applicationCredentialSecret", unrestricted.secret());
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Application credential with access rules
 * 
 * &gt; **Note:** Application Credential access rules are supported only in Keystone
 * starting from [Train](https://releases.openstack.org/train/highlights.html#keystone-identity-service) release.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.identity.ApplicationCredential;
 * import com.pulumi.openstack.identity.ApplicationCredentialArgs;
 * import com.pulumi.openstack.identity.inputs.ApplicationCredentialAccessRuleArgs;
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
 *         var monitoring = new ApplicationCredential("monitoring", ApplicationCredentialArgs.builder()
 *             .name("monitoring")
 *             .expiresAt("2019-02-13T12:12:12Z")
 *             .accessRules(            
 *                 ApplicationCredentialAccessRuleArgs.builder()
 *                     .path("/v2.0/metrics")
 *                     .service("monitoring")
 *                     .method("GET")
 *                     .build(),
 *                 ApplicationCredentialAccessRuleArgs.builder()
 *                     .path("/v2.0/metrics")
 *                     .service("monitoring")
 *                     .method("PUT")
 *                     .build())
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
 * Application Credentials can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:identity/applicationCredential:ApplicationCredential application_credential_1 c17304b7-0953-4738-abb0-67005882b0a0
 * ```
 * 
 */
@ResourceType(type="openstack:identity/applicationCredential:ApplicationCredential")
public class ApplicationCredential extends com.pulumi.resources.CustomResource {
    /**
     * A collection of one or more access rules, which
     * this application credential allows to follow. The structure is described
     * below. Changing this creates a new application credential.
     * 
     */
    @Export(name="accessRules", refs={List.class,ApplicationCredentialAccessRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ApplicationCredentialAccessRule>> accessRules;

    /**
     * @return A collection of one or more access rules, which
     * this application credential allows to follow. The structure is described
     * below. Changing this creates a new application credential.
     * 
     */
    public Output<Optional<List<ApplicationCredentialAccessRule>>> accessRules() {
        return Codegen.optional(this.accessRules);
    }
    /**
     * A description of the application credential.
     * Changing this creates a new application credential.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the application credential.
     * Changing this creates a new application credential.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The expiration time of the application credential
     * in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
     * an application credential will never expire. Changing this creates a new
     * application credential.
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expiresAt;

    /**
     * @return The expiration time of the application credential
     * in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
     * an application credential will never expire. Changing this creates a new
     * application credential.
     * 
     */
    public Output<Optional<String>> expiresAt() {
        return Codegen.optional(this.expiresAt);
    }
    /**
     * A name of the application credential. Changing this
     * creates a new application credential.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name of the application credential. Changing this
     * creates a new application credential.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project the application credential was created
     * for and that authentication requests using this application credential will
     * be scoped to.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project the application credential was created
     * for and that authentication requests using this application credential will
     * be scoped to.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new application credential.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new application credential.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A collection of one or more role names, which this
     * application credential has to be associated with its project. If omitted,
     * all the current user&#39;s roles within the scoped project will be inherited by
     * a new application credential. Changing this creates a new application
     * credential.
     * 
     */
    @Export(name="roles", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> roles;

    /**
     * @return A collection of one or more role names, which this
     * application credential has to be associated with its project. If omitted,
     * all the current user&#39;s roles within the scoped project will be inherited by
     * a new application credential. Changing this creates a new application
     * credential.
     * 
     */
    public Output<List<String>> roles() {
        return this.roles;
    }
    /**
     * The secret for the application credential. If omitted,
     * it will be generated by the server. Changing this creates a new application
     * credential.
     * 
     */
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    /**
     * @return The secret for the application credential. If omitted,
     * it will be generated by the server. Changing this creates a new application
     * credential.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }
    /**
     * A flag indicating whether the application
     * credential may be used for creation or destruction of other application
     * credentials or trusts. Changing this creates a new application credential.
     * 
     */
    @Export(name="unrestricted", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> unrestricted;

    /**
     * @return A flag indicating whether the application
     * credential may be used for creation or destruction of other application
     * credentials or trusts. Changing this creates a new application credential.
     * 
     */
    public Output<Optional<Boolean>> unrestricted() {
        return Codegen.optional(this.unrestricted);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApplicationCredential(java.lang.String name) {
        this(name, ApplicationCredentialArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationCredential(java.lang.String name, @Nullable ApplicationCredentialArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationCredential(java.lang.String name, @Nullable ApplicationCredentialArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/applicationCredential:ApplicationCredential", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ApplicationCredential(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationCredentialState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/applicationCredential:ApplicationCredential", name, state, makeResourceOptions(options, id), false);
    }

    private static ApplicationCredentialArgs makeArgs(@Nullable ApplicationCredentialArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ApplicationCredentialArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secret"
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
    public static ApplicationCredential get(java.lang.String name, Output<java.lang.String> id, @Nullable ApplicationCredentialState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationCredential(name, id, state, options);
    }
}
