// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PortFixedIp {
    /**
     * @return The additional IP address.
     * 
     */
    private final @Nullable String ipAddress;
    /**
     * @return Subnet in which to allocate IP address for
     * this port.
     * 
     */
    private final String subnetId;

    @CustomType.Constructor
    private PortFixedIp(
        @CustomType.Parameter("ipAddress") @Nullable String ipAddress,
        @CustomType.Parameter("subnetId") String subnetId) {
        this.ipAddress = ipAddress;
        this.subnetId = subnetId;
    }

    /**
     * @return The additional IP address.
     * 
     */
    public Optional<String> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }
    /**
     * @return Subnet in which to allocate IP address for
     * this port.
     * 
     */
    public String subnetId() {
        return this.subnetId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PortFixedIp defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String ipAddress;
        private String subnetId;

        public Builder() {
    	      // Empty
        }

        public Builder(PortFixedIp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipAddress = defaults.ipAddress;
    	      this.subnetId = defaults.subnetId;
        }

        public Builder ipAddress(@Nullable String ipAddress) {
            this.ipAddress = ipAddress;
            return this;
        }
        public Builder subnetId(String subnetId) {
            this.subnetId = Objects.requireNonNull(subnetId);
            return this;
        }        public PortFixedIp build() {
            return new PortFixedIp(ipAddress, subnetId);
        }
    }
}
