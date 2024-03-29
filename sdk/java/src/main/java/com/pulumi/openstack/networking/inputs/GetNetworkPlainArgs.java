// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNetworkPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkPlainArgs Empty = new GetNetworkPlainArgs();

    /**
     * Human-readable description of the network.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return Human-readable description of the network.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The external routing facility of the network.
     * 
     */
    @Import(name="external")
    private @Nullable Boolean external;

    /**
     * @return The external routing facility of the network.
     * 
     */
    public Optional<Boolean> external() {
        return Optional.ofNullable(this.external);
    }

    /**
     * The CIDR of a subnet within the network.
     * 
     */
    @Import(name="matchingSubnetCidr")
    private @Nullable String matchingSubnetCidr;

    /**
     * @return The CIDR of a subnet within the network.
     * 
     */
    public Optional<String> matchingSubnetCidr() {
        return Optional.ofNullable(this.matchingSubnetCidr);
    }

    /**
     * The network MTU to filter. Available, when Neutron `net-mtu`
     * extension is enabled.
     * 
     */
    @Import(name="mtu")
    private @Nullable Integer mtu;

    /**
     * @return The network MTU to filter. Available, when Neutron `net-mtu`
     * extension is enabled.
     * 
     */
    public Optional<Integer> mtu() {
        return Optional.ofNullable(this.mtu);
    }

    /**
     * The name of the network.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the network.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the network.
     * 
     */
    @Import(name="networkId")
    private @Nullable String networkId;

    /**
     * @return The ID of the network.
     * 
     */
    public Optional<String> networkId() {
        return Optional.ofNullable(this.networkId);
    }

    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve networks ids. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve networks ids. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The status of the network.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the network.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The list of network tags to filter.
     * 
     */
    @Import(name="tags")
    private @Nullable List<String> tags;

    /**
     * @return The list of network tags to filter.
     * 
     */
    public Optional<List<String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The owner of the network.
     * 
     */
    @Import(name="tenantId")
    private @Nullable String tenantId;

    /**
     * @return The owner of the network.
     * 
     */
    public Optional<String> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * The VLAN transparent attribute for the
     * network.
     * 
     */
    @Import(name="transparentVlan")
    private @Nullable Boolean transparentVlan;

    /**
     * @return The VLAN transparent attribute for the
     * network.
     * 
     */
    public Optional<Boolean> transparentVlan() {
        return Optional.ofNullable(this.transparentVlan);
    }

    private GetNetworkPlainArgs() {}

    private GetNetworkPlainArgs(GetNetworkPlainArgs $) {
        this.description = $.description;
        this.external = $.external;
        this.matchingSubnetCidr = $.matchingSubnetCidr;
        this.mtu = $.mtu;
        this.name = $.name;
        this.networkId = $.networkId;
        this.region = $.region;
        this.status = $.status;
        this.tags = $.tags;
        this.tenantId = $.tenantId;
        this.transparentVlan = $.transparentVlan;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNetworkPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkPlainArgs $;

        public Builder() {
            $ = new GetNetworkPlainArgs();
        }

        public Builder(GetNetworkPlainArgs defaults) {
            $ = new GetNetworkPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Human-readable description of the network.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param external The external routing facility of the network.
         * 
         * @return builder
         * 
         */
        public Builder external(@Nullable Boolean external) {
            $.external = external;
            return this;
        }

        /**
         * @param matchingSubnetCidr The CIDR of a subnet within the network.
         * 
         * @return builder
         * 
         */
        public Builder matchingSubnetCidr(@Nullable String matchingSubnetCidr) {
            $.matchingSubnetCidr = matchingSubnetCidr;
            return this;
        }

        /**
         * @param mtu The network MTU to filter. Available, when Neutron `net-mtu`
         * extension is enabled.
         * 
         * @return builder
         * 
         */
        public Builder mtu(@Nullable Integer mtu) {
            $.mtu = mtu;
            return this;
        }

        /**
         * @param name The name of the network.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param networkId The ID of the network.
         * 
         * @return builder
         * 
         */
        public Builder networkId(@Nullable String networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Neutron client.
         * A Neutron client is needed to retrieve networks ids. If omitted, the
         * `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        /**
         * @param status The status of the network.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param tags The list of network tags to filter.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable List<String> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The list of network tags to filter.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param tenantId The owner of the network.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable String tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param transparentVlan The VLAN transparent attribute for the
         * network.
         * 
         * @return builder
         * 
         */
        public Builder transparentVlan(@Nullable Boolean transparentVlan) {
            $.transparentVlan = transparentVlan;
            return this;
        }

        public GetNetworkPlainArgs build() {
            return $;
        }
    }

}
