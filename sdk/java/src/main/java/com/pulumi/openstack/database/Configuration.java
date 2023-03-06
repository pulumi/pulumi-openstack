// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.database.ConfigurationArgs;
import com.pulumi.openstack.database.inputs.ConfigurationState;
import com.pulumi.openstack.database.outputs.ConfigurationConfiguration;
import com.pulumi.openstack.database.outputs.ConfigurationDatastore;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="openstack:database/configuration:Configuration")
public class Configuration extends com.pulumi.resources.CustomResource {
    /**
     * An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
     * 
     */
    @Export(name="configurations", type=List.class, parameters={ConfigurationConfiguration.class})
    private Output</* @Nullable */ List<ConfigurationConfiguration>> configurations;

    /**
     * @return An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
     * 
     */
    public Output<Optional<List<ConfigurationConfiguration>>> configurations() {
        return Codegen.optional(this.configurations);
    }
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates resource.
     * 
     */
    @Export(name="datastore", type=ConfigurationDatastore.class, parameters={})
    private Output<ConfigurationDatastore> datastore;

    /**
     * @return An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates resource.
     * 
     */
    public Output<ConfigurationDatastore> datastore() {
        return this.datastore;
    }
    /**
     * Description of the resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return Description of the resource.
     * 
     */
    public Output<String> description() {
        return this.description;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Configuration(String name) {
        this(name, ConfigurationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Configuration(String name, ConfigurationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Configuration(String name, ConfigurationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:database/configuration:Configuration", name, args == null ? ConfigurationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Configuration(String name, Output<String> id, @Nullable ConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:database/configuration:Configuration", name, state, makeResourceOptions(options, id));
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
    public static Configuration get(String name, Output<String> id, @Nullable ConfigurationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Configuration(name, id, state, options);
    }
}
