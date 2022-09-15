// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.orchestration.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class StackV1StackOutput {
    /**
     * @return The description of the stack resource.
     * 
     */
    private final @Nullable String description;
    private final String outputKey;
    private final String outputValue;

    @CustomType.Constructor
    private StackV1StackOutput(
        @CustomType.Parameter("description") @Nullable String description,
        @CustomType.Parameter("outputKey") String outputKey,
        @CustomType.Parameter("outputValue") String outputValue) {
        this.description = description;
        this.outputKey = outputKey;
        this.outputValue = outputValue;
    }

    /**
     * @return The description of the stack resource.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public String outputKey() {
        return this.outputKey;
    }
    public String outputValue() {
        return this.outputValue;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(StackV1StackOutput defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String description;
        private String outputKey;
        private String outputValue;

        public Builder() {
    	      // Empty
        }

        public Builder(StackV1StackOutput defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.outputKey = defaults.outputKey;
    	      this.outputValue = defaults.outputValue;
        }

        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        public Builder outputKey(String outputKey) {
            this.outputKey = Objects.requireNonNull(outputKey);
            return this;
        }
        public Builder outputValue(String outputValue) {
            this.outputValue = Objects.requireNonNull(outputValue);
            return this;
        }        public StackV1StackOutput build() {
            return new StackV1StackOutput(description, outputKey, outputValue);
        }
    }
}