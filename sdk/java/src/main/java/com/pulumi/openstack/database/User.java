// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.database.UserArgs;
import com.pulumi.openstack.database.inputs.UserState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 DB user resource within OpenStack.
 * 
 * &gt; **Note:** All arguments including the database password will be stored in the
 * raw state as plain-text. Read more about sensitive data in
 * state.
 * 
 * ## Example Usage
 * 
 * ### User
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.database.User;
 * import com.pulumi.openstack.database.UserArgs;
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
 *         var basic = new User(&#34;basic&#34;, UserArgs.builder()        
 *             .name(&#34;basic&#34;)
 *             .instanceId(basicOpenstackDbInstanceV1.id())
 *             .password(&#34;password&#34;)
 *             .databases(&#34;testdb&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="openstack:database/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * A list of database user should have access to.
     * 
     */
    @Export(name="databases", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> databases;

    /**
     * @return A list of database user should have access to.
     * 
     */
    public Output<List<String>> databases() {
        return this.databases;
    }
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> host;

    public Output<Optional<String>> host() {
        return Codegen.optional(this.host);
    }
    /**
     * The ID for the database instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID for the database instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * A unique name for the resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * User&#39;s password.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return User&#39;s password.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * Openstack region resource is created in.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Openstack region resource is created in.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:database/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:database/user:User", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
