// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.database.InstanceArgs;
import com.pulumi.openstack.database.inputs.InstanceState;
import com.pulumi.openstack.database.outputs.InstanceDatabase;
import com.pulumi.openstack.database.outputs.InstanceDatastore;
import com.pulumi.openstack.database.outputs.InstanceNetwork;
import com.pulumi.openstack.database.outputs.InstanceUser;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 DB instance resource within OpenStack.
 * 
 * &gt; **Note:** All arguments including the instance user password will be stored
 * in the raw state as plain-text. Read more about sensitive data in
 * state.
 * 
 * ## Example Usage
 * ### Instance
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.database.Instance;
 * import com.pulumi.openstack.database.InstanceArgs;
 * import com.pulumi.openstack.database.inputs.InstanceDatastoreArgs;
 * import com.pulumi.openstack.database.inputs.InstanceNetworkArgs;
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
 *         var test = new Instance(&#34;test&#34;, InstanceArgs.builder()        
 *             .datastore(InstanceDatastoreArgs.builder()
 *                 .type(&#34;mysql&#34;)
 *                 .version(&#34;mysql-5.7&#34;)
 *                 .build())
 *             .flavorId(&#34;31792d21-c355-4587-9290-56c1ed0ca376&#34;)
 *             .networks(InstanceNetworkArgs.builder()
 *                 .uuid(&#34;c0612505-caf2-4fb0-b7cb-56a0240a2b12&#34;)
 *                 .build())
 *             .region(&#34;region-test&#34;)
 *             .size(8)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="openstack:database/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * A list of IP addresses assigned to the instance.
     * 
     */
    @Export(name="addresses", type=List.class, parameters={String.class})
    private Output<List<String>> addresses;

    /**
     * @return A list of IP addresses assigned to the instance.
     * 
     */
    public Output<List<String>> addresses() {
        return this.addresses;
    }
    /**
     * Configuration ID to be attached to the instance. Database instance
     * will be rebooted when configuration is detached.
     * 
     */
    @Export(name="configurationId", type=String.class, parameters={})
    private Output</* @Nullable */ String> configurationId;

    /**
     * @return Configuration ID to be attached to the instance. Database instance
     * will be rebooted when configuration is detached.
     * 
     */
    public Output<Optional<String>> configurationId() {
        return Codegen.optional(this.configurationId);
    }
    /**
     * An array of database name, charset and collate. The database
     * object structure is documented below.
     * 
     */
    @Export(name="databases", type=List.class, parameters={InstanceDatabase.class})
    private Output</* @Nullable */ List<InstanceDatabase>> databases;

    /**
     * @return An array of database name, charset and collate. The database
     * object structure is documented below.
     * 
     */
    public Output<Optional<List<InstanceDatabase>>> databases() {
        return Codegen.optional(this.databases);
    }
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates a new instance.
     * 
     */
    @Export(name="datastore", type=InstanceDatastore.class, parameters={})
    private Output<InstanceDatastore> datastore;

    /**
     * @return An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates a new instance.
     * 
     */
    public Output<InstanceDatastore> datastore() {
        return this.datastore;
    }
    /**
     * The flavor ID of the desired flavor for the instance.
     * Changing this creates new instance.
     * 
     */
    @Export(name="flavorId", type=String.class, parameters={})
    private Output<String> flavorId;

    /**
     * @return The flavor ID of the desired flavor for the instance.
     * Changing this creates new instance.
     * 
     */
    public Output<String> flavorId() {
        return this.flavorId;
    }
    /**
     * A unique name for the resource.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A unique name for the resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new instance.
     * 
     */
    @Export(name="networks", type=List.class, parameters={InstanceNetwork.class})
    private Output</* @Nullable */ List<InstanceNetwork>> networks;

    /**
     * @return An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new instance.
     * 
     */
    public Output<Optional<List<InstanceNetwork>>> networks() {
        return Codegen.optional(this.networks);
    }
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to create the db instance. Changing this
     * creates a new instance.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Specifies the volume size in GB. Changing this creates new instance.
     * 
     */
    @Export(name="size", type=Integer.class, parameters={})
    private Output<Integer> size;

    /**
     * @return Specifies the volume size in GB. Changing this creates new instance.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * An array of username, password, host and databases. The user
     * object structure is documented below.
     * 
     */
    @Export(name="users", type=List.class, parameters={InstanceUser.class})
    private Output</* @Nullable */ List<InstanceUser>> users;

    /**
     * @return An array of username, password, host and databases. The user
     * object structure is documented below.
     * 
     */
    public Output<Optional<List<InstanceUser>>> users() {
        return Codegen.optional(this.users);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:database/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:database/instance:Instance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
