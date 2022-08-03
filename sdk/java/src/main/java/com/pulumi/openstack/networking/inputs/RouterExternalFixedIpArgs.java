// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouterExternalFixedIpArgs extends com.pulumi.resources.ResourceArgs {

    public static final RouterExternalFixedIpArgs Empty = new RouterExternalFixedIpArgs();

    /**
     * The IP address to set on the router.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return The IP address to set on the router.
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * Subnet in which the fixed IP belongs to.
     * 
     */
    @Import(name="subnetId")
    private @Nullable Output<String> subnetId;

    /**
     * @return Subnet in which the fixed IP belongs to.
     * 
     */
    public Optional<Output<String>> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    private RouterExternalFixedIpArgs() {}

    private RouterExternalFixedIpArgs(RouterExternalFixedIpArgs $) {
        this.ipAddress = $.ipAddress;
        this.subnetId = $.subnetId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouterExternalFixedIpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouterExternalFixedIpArgs $;

        public Builder() {
            $ = new RouterExternalFixedIpArgs();
        }

        public Builder(RouterExternalFixedIpArgs defaults) {
            $ = new RouterExternalFixedIpArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipAddress The IP address to set on the router.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The IP address to set on the router.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param subnetId Subnet in which the fixed IP belongs to.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(@Nullable Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId Subnet in which the fixed IP belongs to.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        public RouterExternalFixedIpArgs build() {
            return $;
        }
    }

}
