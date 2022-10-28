// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PortAllowedAddressPair {
    /**
     * @return The additional IP address.
     * 
     */
    private String ipAddress;
    /**
     * @return The additional MAC address.
     * 
     */
    private @Nullable String macAddress;

    private PortAllowedAddressPair() {}
    /**
     * @return The additional IP address.
     * 
     */
    public String ipAddress() {
        return this.ipAddress;
    }
    /**
     * @return The additional MAC address.
     * 
     */
    public Optional<String> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PortAllowedAddressPair defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ipAddress;
        private @Nullable String macAddress;
        public Builder() {}
        public Builder(PortAllowedAddressPair defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipAddress = defaults.ipAddress;
    	      this.macAddress = defaults.macAddress;
        }

        @CustomType.Setter
        public Builder ipAddress(String ipAddress) {
            this.ipAddress = Objects.requireNonNull(ipAddress);
            return this;
        }
        @CustomType.Setter
        public Builder macAddress(@Nullable String macAddress) {
            this.macAddress = macAddress;
            return this;
        }
        public PortAllowedAddressPair build() {
            final var o = new PortAllowedAddressPair();
            o.ipAddress = ipAddress;
            o.macAddress = macAddress;
            return o;
        }
    }
}
