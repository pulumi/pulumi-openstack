// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNetworkArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNetworkArgs Empty = new GetNetworkArgs();

    /**
     * Human-readable description of the network.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description of the network.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The external routing facility of the network.
     * 
     */
    @Import(name="external")
    private @Nullable Output<Boolean> external;

    /**
     * @return The external routing facility of the network.
     * 
     */
    public Optional<Output<Boolean>> external() {
        return Optional.ofNullable(this.external);
    }

    /**
     * The CIDR of a subnet within the network.
     * 
     */
    @Import(name="matchingSubnetCidr")
    private @Nullable Output<String> matchingSubnetCidr;

    /**
     * @return The CIDR of a subnet within the network.
     * 
     */
    public Optional<Output<String>> matchingSubnetCidr() {
        return Optional.ofNullable(this.matchingSubnetCidr);
    }

    /**
     * The network MTU to filter. Available, when Neutron `net-mtu`
     * extension is enabled.
     * 
     */
    @Import(name="mtu")
    private @Nullable Output<Integer> mtu;

    /**
     * @return The network MTU to filter. Available, when Neutron `net-mtu`
     * extension is enabled.
     * 
     */
    public Optional<Output<Integer>> mtu() {
        return Optional.ofNullable(this.mtu);
    }

    /**
     * The name of the network.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the network.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the network.
     * 
     */
    @Import(name="networkId")
    private @Nullable Output<String> networkId;

    /**
     * @return The ID of the network.
     * 
     */
    public Optional<Output<String>> networkId() {
        return Optional.ofNullable(this.networkId);
    }

    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve networks ids. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve networks ids. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The status of the network.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the network.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The list of network tags to filter.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The list of network tags to filter.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The owner of the network.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the network.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * The VLAN transparent attribute for the
     * network.
     * 
     */
    @Import(name="transparentVlan")
    private @Nullable Output<Boolean> transparentVlan;

    /**
     * @return The VLAN transparent attribute for the
     * network.
     * 
     */
    public Optional<Output<Boolean>> transparentVlan() {
        return Optional.ofNullable(this.transparentVlan);
    }

    private GetNetworkArgs() {}

    private GetNetworkArgs(GetNetworkArgs $) {
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
    public static Builder builder(GetNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNetworkArgs $;

        public Builder() {
            $ = new GetNetworkArgs();
        }

        public Builder(GetNetworkArgs defaults) {
            $ = new GetNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Human-readable description of the network.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description of the network.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param external The external routing facility of the network.
         * 
         * @return builder
         * 
         */
        public Builder external(@Nullable Output<Boolean> external) {
            $.external = external;
            return this;
        }

        /**
         * @param external The external routing facility of the network.
         * 
         * @return builder
         * 
         */
        public Builder external(Boolean external) {
            return external(Output.of(external));
        }

        /**
         * @param matchingSubnetCidr The CIDR of a subnet within the network.
         * 
         * @return builder
         * 
         */
        public Builder matchingSubnetCidr(@Nullable Output<String> matchingSubnetCidr) {
            $.matchingSubnetCidr = matchingSubnetCidr;
            return this;
        }

        /**
         * @param matchingSubnetCidr The CIDR of a subnet within the network.
         * 
         * @return builder
         * 
         */
        public Builder matchingSubnetCidr(String matchingSubnetCidr) {
            return matchingSubnetCidr(Output.of(matchingSubnetCidr));
        }

        /**
         * @param mtu The network MTU to filter. Available, when Neutron `net-mtu`
         * extension is enabled.
         * 
         * @return builder
         * 
         */
        public Builder mtu(@Nullable Output<Integer> mtu) {
            $.mtu = mtu;
            return this;
        }

        /**
         * @param mtu The network MTU to filter. Available, when Neutron `net-mtu`
         * extension is enabled.
         * 
         * @return builder
         * 
         */
        public Builder mtu(Integer mtu) {
            return mtu(Output.of(mtu));
        }

        /**
         * @param name The name of the network.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the network.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param networkId The ID of the network.
         * 
         * @return builder
         * 
         */
        public Builder networkId(@Nullable Output<String> networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param networkId The ID of the network.
         * 
         * @return builder
         * 
         */
        public Builder networkId(String networkId) {
            return networkId(Output.of(networkId));
        }

        /**
         * @param region The region in which to obtain the V2 Neutron client.
         * A Neutron client is needed to retrieve networks ids. If omitted, the
         * `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
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
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param status The status of the network.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the network.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The list of network tags to filter.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The list of network tags to filter.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
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
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the network.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param transparentVlan The VLAN transparent attribute for the
         * network.
         * 
         * @return builder
         * 
         */
        public Builder transparentVlan(@Nullable Output<Boolean> transparentVlan) {
            $.transparentVlan = transparentVlan;
            return this;
        }

        /**
         * @param transparentVlan The VLAN transparent attribute for the
         * network.
         * 
         * @return builder
         * 
         */
        public Builder transparentVlan(Boolean transparentVlan) {
            return transparentVlan(Output.of(transparentVlan));
        }

        public GetNetworkArgs build() {
            return $;
        }
    }

}
