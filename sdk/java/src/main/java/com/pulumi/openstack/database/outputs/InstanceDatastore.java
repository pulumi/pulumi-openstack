// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class InstanceDatastore {
    /**
     * @return Database engine type to be used in new instance. Changing this
     * creates a new instance.
     * 
     */
    private String type;
    /**
     * @return Version of database engine type to be used in new instance.
     * Changing this creates a new instance.
     * 
     */
    private String version;

    private InstanceDatastore() {}
    /**
     * @return Database engine type to be used in new instance. Changing this
     * creates a new instance.
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return Version of database engine type to be used in new instance.
     * Changing this creates a new instance.
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceDatastore defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String type;
        private String version;
        public Builder() {}
        public Builder(InstanceDatastore defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.type = defaults.type;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("InstanceDatastore", "type");
            }
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("InstanceDatastore", "version");
            }
            this.version = version;
            return this;
        }
        public InstanceDatastore build() {
            final var _resultValue = new InstanceDatastore();
            _resultValue.type = type;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}