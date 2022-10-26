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
    private @Nullable String charset;
    /**
     * @return Database collation. Changing this creates a new instance.
     * 
     */
    private @Nullable String collate;
    /**
     * @return Database to be created on new instance. Changing this creates a
     * new instance.
     * 
     */
    private String name;

    private InstanceDatabase() {}
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
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String charset;
        private @Nullable String collate;
        private String name;
        public Builder() {}
        public Builder(InstanceDatabase defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.charset = defaults.charset;
    	      this.collate = defaults.collate;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder charset(@Nullable String charset) {
            this.charset = charset;
            return this;
        }
        @CustomType.Setter
        public Builder collate(@Nullable String collate) {
            this.collate = collate;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public InstanceDatabase build() {
            final var o = new InstanceDatabase();
            o.charset = charset;
            o.collate = collate;
            o.name = name;
            return o;
        }
    }
}
