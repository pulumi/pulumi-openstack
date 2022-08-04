// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ConfigurationDatastoreArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConfigurationDatastoreArgs Empty = new ConfigurationDatastoreArgs();

    /**
     * Database engine type to be used with this configuration. Changing this creates a new resource.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Database engine type to be used with this configuration. Changing this creates a new resource.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * Version of database engine type to be used with this configuration. Changing this creates a new resource.
     * 
     */
    @Import(name="version", required=true)
    private Output<String> version;

    /**
     * @return Version of database engine type to be used with this configuration. Changing this creates a new resource.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    private ConfigurationDatastoreArgs() {}

    private ConfigurationDatastoreArgs(ConfigurationDatastoreArgs $) {
        this.type = $.type;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConfigurationDatastoreArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConfigurationDatastoreArgs $;

        public Builder() {
            $ = new ConfigurationDatastoreArgs();
        }

        public Builder(ConfigurationDatastoreArgs defaults) {
            $ = new ConfigurationDatastoreArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param type Database engine type to be used with this configuration. Changing this creates a new resource.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Database engine type to be used with this configuration. Changing this creates a new resource.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param version Version of database engine type to be used with this configuration. Changing this creates a new resource.
         * 
         * @return builder
         * 
         */
        public Builder version(Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of database engine type to be used with this configuration. Changing this creates a new resource.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public ConfigurationDatastoreArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            $.version = Objects.requireNonNull($.version, "expected parameter 'version' to be non-null");
            return $;
        }
    }

}
