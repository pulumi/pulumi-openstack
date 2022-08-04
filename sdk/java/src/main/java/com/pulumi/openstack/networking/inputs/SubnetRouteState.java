// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SubnetRouteState extends com.pulumi.resources.ResourceArgs {

    public static final SubnetRouteState Empty = new SubnetRouteState();

    /**
     * CIDR block to match on the packet’s destination IP. Changing
     * this creates a new routing entry.
     * 
     */
    @Import(name="destinationCidr")
    private @Nullable Output<String> destinationCidr;

    /**
     * @return CIDR block to match on the packet’s destination IP. Changing
     * this creates a new routing entry.
     * 
     */
    public Optional<Output<String>> destinationCidr() {
        return Optional.ofNullable(this.destinationCidr);
    }

    /**
     * IP address of the next hop gateway.  Changing
     * this creates a new routing entry.
     * 
     */
    @Import(name="nextHop")
    private @Nullable Output<String> nextHop;

    /**
     * @return IP address of the next hop gateway.  Changing
     * this creates a new routing entry.
     * 
     */
    public Optional<Output<String>> nextHop() {
        return Optional.ofNullable(this.nextHop);
    }

    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * ID of the subnet this routing entry belongs to. Changing
     * this creates a new routing entry.
     * 
     */
    @Import(name="subnetId")
    private @Nullable Output<String> subnetId;

    /**
     * @return ID of the subnet this routing entry belongs to. Changing
     * this creates a new routing entry.
     * 
     */
    public Optional<Output<String>> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    private SubnetRouteState() {}

    private SubnetRouteState(SubnetRouteState $) {
        this.destinationCidr = $.destinationCidr;
        this.nextHop = $.nextHop;
        this.region = $.region;
        this.subnetId = $.subnetId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubnetRouteState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubnetRouteState $;

        public Builder() {
            $ = new SubnetRouteState();
        }

        public Builder(SubnetRouteState defaults) {
            $ = new SubnetRouteState(Objects.requireNonNull(defaults));
        }

        /**
         * @param destinationCidr CIDR block to match on the packet’s destination IP. Changing
         * this creates a new routing entry.
         * 
         * @return builder
         * 
         */
        public Builder destinationCidr(@Nullable Output<String> destinationCidr) {
            $.destinationCidr = destinationCidr;
            return this;
        }

        /**
         * @param destinationCidr CIDR block to match on the packet’s destination IP. Changing
         * this creates a new routing entry.
         * 
         * @return builder
         * 
         */
        public Builder destinationCidr(String destinationCidr) {
            return destinationCidr(Output.of(destinationCidr));
        }

        /**
         * @param nextHop IP address of the next hop gateway.  Changing
         * this creates a new routing entry.
         * 
         * @return builder
         * 
         */
        public Builder nextHop(@Nullable Output<String> nextHop) {
            $.nextHop = nextHop;
            return this;
        }

        /**
         * @param nextHop IP address of the next hop gateway.  Changing
         * this creates a new routing entry.
         * 
         * @return builder
         * 
         */
        public Builder nextHop(String nextHop) {
            return nextHop(Output.of(nextHop));
        }

        /**
         * @param region The region in which to obtain the V2 networking client.
         * A networking client is needed to configure a routing entry on a subnet. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * routing entry.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 networking client.
         * A networking client is needed to configure a routing entry on a subnet. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * routing entry.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param subnetId ID of the subnet this routing entry belongs to. Changing
         * this creates a new routing entry.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(@Nullable Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId ID of the subnet this routing entry belongs to. Changing
         * this creates a new routing entry.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        public SubnetRouteState build() {
            return $;
        }
    }

}
