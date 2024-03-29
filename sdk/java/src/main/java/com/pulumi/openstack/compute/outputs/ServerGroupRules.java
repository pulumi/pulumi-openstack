// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServerGroupRules {
    private @Nullable Integer maxServerPerHost;

    private ServerGroupRules() {}
    public Optional<Integer> maxServerPerHost() {
        return Optional.ofNullable(this.maxServerPerHost);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServerGroupRules defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer maxServerPerHost;
        public Builder() {}
        public Builder(ServerGroupRules defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxServerPerHost = defaults.maxServerPerHost;
        }

        @CustomType.Setter
        public Builder maxServerPerHost(@Nullable Integer maxServerPerHost) {

            this.maxServerPerHost = maxServerPerHost;
            return this;
        }
        public ServerGroupRules build() {
            final var _resultValue = new ServerGroupRules();
            _resultValue.maxServerPerHost = maxServerPerHost;
            return _resultValue;
        }
    }
}
