// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceState Empty = new ServiceState();

    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing service.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing service.
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * The human-readable description for the service.
     * Changing this updates the description of the existing service.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The human-readable description for the service.
     * Changing this updates the description of the existing service.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The read-only external (public) IPv4 address that is used for the VPN service.
     * 
     */
    @Import(name="externalV4Ip")
    private @Nullable Output<String> externalV4Ip;

    /**
     * @return The read-only external (public) IPv4 address that is used for the VPN service.
     * 
     */
    public Optional<Output<String>> externalV4Ip() {
        return Optional.ofNullable(this.externalV4Ip);
    }

    /**
     * The read-only external (public) IPv6 address that is used for the VPN service.
     * 
     */
    @Import(name="externalV6Ip")
    private @Nullable Output<String> externalV6Ip;

    /**
     * @return The read-only external (public) IPv6 address that is used for the VPN service.
     * 
     */
    public Optional<Output<String>> externalV6Ip() {
        return Optional.ofNullable(this.externalV6Ip);
    }

    /**
     * The name of the service. Changing this updates the name of
     * the existing service.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the service. Changing this updates the name of
     * the existing service.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The ID of the router. Changing this creates a new service.
     * 
     */
    @Import(name="routerId")
    private @Nullable Output<String> routerId;

    /**
     * @return The ID of the router. Changing this creates a new service.
     * 
     */
    public Optional<Output<String>> routerId() {
        return Optional.ofNullable(this.routerId);
    }

    /**
     * Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * SubnetID is the ID of the subnet. Default is null.
     * 
     */
    @Import(name="subnetId")
    private @Nullable Output<String> subnetId;

    /**
     * @return SubnetID is the ID of the subnet. Default is null.
     * 
     */
    public Optional<Output<String>> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    /**
     * The owner of the service. Required if admin wants to
     * create a service for another project. Changing this creates a new service.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the service. Required if admin wants to
     * create a service for another project. Changing this creates a new service.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Map of additional options.
     * 
     */
    @Import(name="valueSpecs")
    private @Nullable Output<Map<String,String>> valueSpecs;

    /**
     * @return Map of additional options.
     * 
     */
    public Optional<Output<Map<String,String>>> valueSpecs() {
        return Optional.ofNullable(this.valueSpecs);
    }

    private ServiceState() {}

    private ServiceState(ServiceState $) {
        this.adminStateUp = $.adminStateUp;
        this.description = $.description;
        this.externalV4Ip = $.externalV4Ip;
        this.externalV6Ip = $.externalV6Ip;
        this.name = $.name;
        this.region = $.region;
        this.routerId = $.routerId;
        this.status = $.status;
        this.subnetId = $.subnetId;
        this.tenantId = $.tenantId;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceState $;

        public Builder() {
            $ = new ServiceState();
        }

        public Builder(ServiceState defaults) {
            $ = new ServiceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp The administrative state of the resource. Can either be up(true) or down(false).
         * Changing this updates the administrative state of the existing service.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp The administrative state of the resource. Can either be up(true) or down(false).
         * Changing this updates the administrative state of the existing service.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param description The human-readable description for the service.
         * Changing this updates the description of the existing service.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The human-readable description for the service.
         * Changing this updates the description of the existing service.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param externalV4Ip The read-only external (public) IPv4 address that is used for the VPN service.
         * 
         * @return builder
         * 
         */
        public Builder externalV4Ip(@Nullable Output<String> externalV4Ip) {
            $.externalV4Ip = externalV4Ip;
            return this;
        }

        /**
         * @param externalV4Ip The read-only external (public) IPv4 address that is used for the VPN service.
         * 
         * @return builder
         * 
         */
        public Builder externalV4Ip(String externalV4Ip) {
            return externalV4Ip(Output.of(externalV4Ip));
        }

        /**
         * @param externalV6Ip The read-only external (public) IPv6 address that is used for the VPN service.
         * 
         * @return builder
         * 
         */
        public Builder externalV6Ip(@Nullable Output<String> externalV6Ip) {
            $.externalV6Ip = externalV6Ip;
            return this;
        }

        /**
         * @param externalV6Ip The read-only external (public) IPv6 address that is used for the VPN service.
         * 
         * @return builder
         * 
         */
        public Builder externalV6Ip(String externalV6Ip) {
            return externalV6Ip(Output.of(externalV6Ip));
        }

        /**
         * @param name The name of the service. Changing this updates the name of
         * the existing service.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the service. Changing this updates the name of
         * the existing service.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a VPN service. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * service.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a VPN service. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * service.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param routerId The ID of the router. Changing this creates a new service.
         * 
         * @return builder
         * 
         */
        public Builder routerId(@Nullable Output<String> routerId) {
            $.routerId = routerId;
            return this;
        }

        /**
         * @param routerId The ID of the router. Changing this creates a new service.
         * 
         * @return builder
         * 
         */
        public Builder routerId(String routerId) {
            return routerId(Output.of(routerId));
        }

        /**
         * @param status Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param subnetId SubnetID is the ID of the subnet. Default is null.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(@Nullable Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId SubnetID is the ID of the subnet. Default is null.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        /**
         * @param tenantId The owner of the service. Required if admin wants to
         * create a service for another project. Changing this creates a new service.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the service. Required if admin wants to
         * create a service for another project. Changing this creates a new service.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param valueSpecs Map of additional options.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(@Nullable Output<Map<String,String>> valueSpecs) {
            $.valueSpecs = valueSpecs;
            return this;
        }

        /**
         * @param valueSpecs Map of additional options.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(Map<String,String> valueSpecs) {
            return valueSpecs(Output.of(valueSpecs));
        }

        public ServiceState build() {
            return $;
        }
    }

}
