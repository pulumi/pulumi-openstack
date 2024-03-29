// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRouterExternalFixedIp {
    /**
     * @return The IP address to set on the router.
     * 
     */
    private @Nullable String ipAddress;
    /**
     * @return Subnet in which the fixed IP belongs to.
     * 
     */
    private @Nullable String subnetId;

    private GetRouterExternalFixedIp() {}
    /**
     * @return The IP address to set on the router.
     * 
     */
    public Optional<String> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }
    /**
     * @return Subnet in which the fixed IP belongs to.
     * 
     */
    public Optional<String> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRouterExternalFixedIp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String ipAddress;
        private @Nullable String subnetId;
        public Builder() {}
        public Builder(GetRouterExternalFixedIp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipAddress = defaults.ipAddress;
    	      this.subnetId = defaults.subnetId;
        }

        @CustomType.Setter
        public Builder ipAddress(@Nullable String ipAddress) {

            this.ipAddress = ipAddress;
            return this;
        }
        @CustomType.Setter
        public Builder subnetId(@Nullable String subnetId) {

            this.subnetId = subnetId;
            return this;
        }
        public GetRouterExternalFixedIp build() {
            final var _resultValue = new GetRouterExternalFixedIp();
            _resultValue.ipAddress = ipAddress;
            _resultValue.subnetId = subnetId;
            return _resultValue;
        }
    }
}
