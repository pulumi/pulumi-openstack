// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IkePolicyLifetime {
    private final @Nullable String units;
    /**
     * @return The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     * 
     */
    private final @Nullable Integer value;

    @CustomType.Constructor
    private IkePolicyLifetime(
        @CustomType.Parameter("units") @Nullable String units,
        @CustomType.Parameter("value") @Nullable Integer value) {
        this.units = units;
        this.value = value;
    }

    public Optional<String> units() {
        return Optional.ofNullable(this.units);
    }
    /**
     * @return The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     * 
     */
    public Optional<Integer> value() {
        return Optional.ofNullable(this.value);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IkePolicyLifetime defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String units;
        private @Nullable Integer value;

        public Builder() {
    	      // Empty
        }

        public Builder(IkePolicyLifetime defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.units = defaults.units;
    	      this.value = defaults.value;
        }

        public Builder units(@Nullable String units) {
            this.units = units;
            return this;
        }
        public Builder value(@Nullable Integer value) {
            this.value = value;
            return this;
        }        public IkePolicyLifetime build() {
            return new IkePolicyLifetime(units, value);
        }
    }
}
