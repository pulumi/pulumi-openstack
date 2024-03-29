// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceDatabaseArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceDatabaseArgs Empty = new InstanceDatabaseArgs();

    /**
     * Database character set. Changing this creates a
     * new instance.
     * 
     */
    @Import(name="charset")
    private @Nullable Output<String> charset;

    /**
     * @return Database character set. Changing this creates a
     * new instance.
     * 
     */
    public Optional<Output<String>> charset() {
        return Optional.ofNullable(this.charset);
    }

    /**
     * Database collation. Changing this creates a new instance.
     * 
     */
    @Import(name="collate")
    private @Nullable Output<String> collate;

    /**
     * @return Database collation. Changing this creates a new instance.
     * 
     */
    public Optional<Output<String>> collate() {
        return Optional.ofNullable(this.collate);
    }

    /**
     * Database to be created on new instance. Changing this creates a
     * new instance.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Database to be created on new instance. Changing this creates a
     * new instance.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private InstanceDatabaseArgs() {}

    private InstanceDatabaseArgs(InstanceDatabaseArgs $) {
        this.charset = $.charset;
        this.collate = $.collate;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceDatabaseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceDatabaseArgs $;

        public Builder() {
            $ = new InstanceDatabaseArgs();
        }

        public Builder(InstanceDatabaseArgs defaults) {
            $ = new InstanceDatabaseArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param charset Database character set. Changing this creates a
         * new instance.
         * 
         * @return builder
         * 
         */
        public Builder charset(@Nullable Output<String> charset) {
            $.charset = charset;
            return this;
        }

        /**
         * @param charset Database character set. Changing this creates a
         * new instance.
         * 
         * @return builder
         * 
         */
        public Builder charset(String charset) {
            return charset(Output.of(charset));
        }

        /**
         * @param collate Database collation. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder collate(@Nullable Output<String> collate) {
            $.collate = collate;
            return this;
        }

        /**
         * @param collate Database collation. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder collate(String collate) {
            return collate(Output.of(collate));
        }

        /**
         * @param name Database to be created on new instance. Changing this creates a
         * new instance.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Database to be created on new instance. Changing this creates a
         * new instance.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public InstanceDatabaseArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("InstanceDatabaseArgs", "name");
            }
            return $;
        }
    }

}
