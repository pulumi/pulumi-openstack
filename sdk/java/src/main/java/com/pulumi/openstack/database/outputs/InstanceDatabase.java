// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceDatabase {
    /**
     * @return Database character set. Changing this creates a
     * new instance.
     * 
     */
    private final @Nullable String charset;
    /**
     * @return Database collation. Changing this creates a new instance.
     * 
     */
    private final @Nullable String collate;
    /**
     * @return Database to be created on new instance. Changing this creates a
     * new instance.
     * 
     */
    private final String name;

    @CustomType.Constructor
    private InstanceDatabase(
        @CustomType.Parameter("charset") @Nullable String charset,
        @CustomType.Parameter("collate") @Nullable String collate,
        @CustomType.Parameter("name") String name) {
        this.charset = charset;
        this.collate = collate;
        this.name = name;
    }

    /**
     * @return Database character set. Changing this creates a
     * new instance.
     * 
     */
    public Optional<String> charset() {
        return Optional.ofNullable(this.charset);
    }
    /**
     * @return Database collation. Changing this creates a new instance.
     * 
     */
    public Optional<String> collate() {
        return Optional.ofNullable(this.collate);
    }
    /**
     * @return Database to be created on new instance. Changing this creates a
     * new instance.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceDatabase defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String charset;
        private @Nullable String collate;
        private String name;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceDatabase defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.charset = defaults.charset;
    	      this.collate = defaults.collate;
    	      this.name = defaults.name;
        }

        public Builder charset(@Nullable String charset) {
            this.charset = charset;
            return this;
        }
        public Builder collate(@Nullable String collate) {
            this.collate = collate;
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }        public InstanceDatabase build() {
            return new InstanceDatabase(charset, collate, name);
        }
    }
}