// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PortAllowedAddressPairArgs extends com.pulumi.resources.ResourceArgs {

    public static final PortAllowedAddressPairArgs Empty = new PortAllowedAddressPairArgs();

    /**
     * The additional IP address.
     * 
     */
    @Import(name="ipAddress", required=true)
    private Output<String> ipAddress;

    /**
     * @return The additional IP address.
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }

    /**
     * The additional MAC address.
     * 
     */
    @Import(name="macAddress")
    private @Nullable Output<String> macAddress;

    /**
     * @return The additional MAC address.
     * 
     */
    public Optional<Output<String>> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }

    private PortAllowedAddressPairArgs() {}

    private PortAllowedAddressPairArgs(PortAllowedAddressPairArgs $) {
        this.ipAddress = $.ipAddress;
        this.macAddress = $.macAddress;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PortAllowedAddressPairArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PortAllowedAddressPairArgs $;

        public Builder() {
            $ = new PortAllowedAddressPairArgs();
        }

        public Builder(PortAllowedAddressPairArgs defaults) {
            $ = new PortAllowedAddressPairArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipAddress The additional IP address.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The additional IP address.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param macAddress The additional MAC address.
         * 
         * @return builder
         * 
         */
        public Builder macAddress(@Nullable Output<String> macAddress) {
            $.macAddress = macAddress;
            return this;
        }

        /**
         * @param macAddress The additional MAC address.
         * 
         * @return builder
         * 
         */
        public Builder macAddress(String macAddress) {
            return macAddress(Output.of(macAddress));
        }

        public PortAllowedAddressPairArgs build() {
            if ($.ipAddress == null) {
                throw new MissingRequiredPropertyException("PortAllowedAddressPairArgs", "ipAddress");
            }
            return $;
        }
    }

}
