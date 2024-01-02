// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ConfigurationConfiguration {
    /**
     * @return Configuration parameter name. Changing this creates a new resource.
     * 
     */
    private String name;
    /**
     * @return Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
     * 
     */
    private @Nullable Boolean stringType;
    /**
     * @return Configuration parameter value. Changing this creates a new resource.
     * 
     */
    private String value;

    private ConfigurationConfiguration() {}
    /**
     * @return Configuration parameter name. Changing this creates a new resource.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
     * 
     */
    public Optional<Boolean> stringType() {
        return Optional.ofNullable(this.stringType);
    }
    /**
     * @return Configuration parameter value. Changing this creates a new resource.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConfigurationConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private @Nullable Boolean stringType;
        private String value;
        public Builder() {}
        public Builder(ConfigurationConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.stringType = defaults.stringType;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("ConfigurationConfiguration", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder stringType(@Nullable Boolean stringType) {

            this.stringType = stringType;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("ConfigurationConfiguration", "value");
            }
            this.value = value;
            return this;
        }
        public ConfigurationConfiguration build() {
            final var _resultValue = new ConfigurationConfiguration();
            _resultValue.name = name;
            _resultValue.stringType = stringType;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
