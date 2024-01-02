// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VipArgs extends com.pulumi.resources.ResourceArgs {

    public static final VipArgs Empty = new VipArgs();

    /**
     * The IP address of the vip. Changing this creates a new
     * vip.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return The IP address of the vip. Changing this creates a new
     * vip.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * The administrative state of the vip.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
     * state of the existing vip.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the vip.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
     * state of the existing vip.
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * The maximum number of connections allowed for the
     * vip. Default is -1, meaning no limit. Changing this updates the conn_limit
     * of the existing vip.
     * 
     */
    @Import(name="connLimit")
    private @Nullable Output<Integer> connLimit;

    /**
     * @return The maximum number of connections allowed for the
     * vip. Default is -1, meaning no limit. Changing this updates the conn_limit
     * of the existing vip.
     * 
     */
    public Optional<Output<Integer>> connLimit() {
        return Optional.ofNullable(this.connLimit);
    }

    /**
     * Human-readable description for the vip. Changing
     * this updates the description of the existing vip.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description for the vip. Changing
     * this updates the description of the existing vip.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A *Networking* Floating IP that will be associated
     * with the vip. The Floating IP must be provisioned already.
     * 
     */
    @Import(name="floatingIp")
    private @Nullable Output<String> floatingIp;

    /**
     * @return A *Networking* Floating IP that will be associated
     * with the vip. The Floating IP must be provisioned already.
     * 
     */
    public Optional<Output<String>> floatingIp() {
        return Optional.ofNullable(this.floatingIp);
    }

    /**
     * The name of the vip. Changing this updates the name of
     * the existing vip.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the vip. Changing this updates the name of
     * the existing vip.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Omit this field to prevent session persistence.
     * The persistence object structure is documented below. Changing this updates
     * the persistence of the existing vip.
     * 
     */
    @Import(name="persistence")
    private @Nullable Output<Map<String,Object>> persistence;

    /**
     * @return Omit this field to prevent session persistence.
     * The persistence object structure is documented below. Changing this updates
     * the persistence of the existing vip.
     * 
     */
    public Optional<Output<Map<String,Object>>> persistence() {
        return Optional.ofNullable(this.persistence);
    }

    /**
     * The ID of the pool with which the vip is associated.
     * Changing this updates the pool_id of the existing vip.
     * 
     */
    @Import(name="poolId", required=true)
    private Output<String> poolId;

    /**
     * @return The ID of the pool with which the vip is associated.
     * Changing this updates the pool_id of the existing vip.
     * 
     */
    public Output<String> poolId() {
        return this.poolId;
    }

    /**
     * The port on which to listen for client traffic. Changing
     * this creates a new vip.
     * 
     */
    @Import(name="port", required=true)
    private Output<Integer> port;

    /**
     * @return The port on which to listen for client traffic. Changing
     * this creates a new vip.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }

    /**
     * The protocol - can be either &#39;TCP, &#39;HTTP&#39;, or
     * HTTPS&#39;. Changing this creates a new vip.
     * 
     */
    @Import(name="protocol", required=true)
    private Output<String> protocol;

    /**
     * @return The protocol - can be either &#39;TCP, &#39;HTTP&#39;, or
     * HTTPS&#39;. Changing this creates a new vip.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VIP. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * VIP.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VIP. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * VIP.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The network on which to allocate the vip&#39;s address. A
     * tenant can only create vips on networks authorized by policy (e.g. networks
     * that belong to them or networks that are shared). Changing this creates a
     * new vip.
     * 
     */
    @Import(name="subnetId", required=true)
    private Output<String> subnetId;

    /**
     * @return The network on which to allocate the vip&#39;s address. A
     * tenant can only create vips on networks authorized by policy (e.g. networks
     * that belong to them or networks that are shared). Changing this creates a
     * new vip.
     * 
     */
    public Output<String> subnetId() {
        return this.subnetId;
    }

    /**
     * The owner of the vip. Required if admin wants to
     * create a vip member for another tenant. Changing this creates a new vip.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the vip. Required if admin wants to
     * create a vip member for another tenant. Changing this creates a new vip.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private VipArgs() {}

    private VipArgs(VipArgs $) {
        this.address = $.address;
        this.adminStateUp = $.adminStateUp;
        this.connLimit = $.connLimit;
        this.description = $.description;
        this.floatingIp = $.floatingIp;
        this.name = $.name;
        this.persistence = $.persistence;
        this.poolId = $.poolId;
        this.port = $.port;
        this.protocol = $.protocol;
        this.region = $.region;
        this.subnetId = $.subnetId;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VipArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VipArgs $;

        public Builder() {
            $ = new VipArgs();
        }

        public Builder(VipArgs defaults) {
            $ = new VipArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The IP address of the vip. Changing this creates a new
         * vip.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address The IP address of the vip. Changing this creates a new
         * vip.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param adminStateUp The administrative state of the vip.
         * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
         * state of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp The administrative state of the vip.
         * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
         * state of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param connLimit The maximum number of connections allowed for the
         * vip. Default is -1, meaning no limit. Changing this updates the conn_limit
         * of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder connLimit(@Nullable Output<Integer> connLimit) {
            $.connLimit = connLimit;
            return this;
        }

        /**
         * @param connLimit The maximum number of connections allowed for the
         * vip. Default is -1, meaning no limit. Changing this updates the conn_limit
         * of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder connLimit(Integer connLimit) {
            return connLimit(Output.of(connLimit));
        }

        /**
         * @param description Human-readable description for the vip. Changing
         * this updates the description of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description for the vip. Changing
         * this updates the description of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param floatingIp A *Networking* Floating IP that will be associated
         * with the vip. The Floating IP must be provisioned already.
         * 
         * @return builder
         * 
         */
        public Builder floatingIp(@Nullable Output<String> floatingIp) {
            $.floatingIp = floatingIp;
            return this;
        }

        /**
         * @param floatingIp A *Networking* Floating IP that will be associated
         * with the vip. The Floating IP must be provisioned already.
         * 
         * @return builder
         * 
         */
        public Builder floatingIp(String floatingIp) {
            return floatingIp(Output.of(floatingIp));
        }

        /**
         * @param name The name of the vip. Changing this updates the name of
         * the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the vip. Changing this updates the name of
         * the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param persistence Omit this field to prevent session persistence.
         * The persistence object structure is documented below. Changing this updates
         * the persistence of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder persistence(@Nullable Output<Map<String,Object>> persistence) {
            $.persistence = persistence;
            return this;
        }

        /**
         * @param persistence Omit this field to prevent session persistence.
         * The persistence object structure is documented below. Changing this updates
         * the persistence of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder persistence(Map<String,Object> persistence) {
            return persistence(Output.of(persistence));
        }

        /**
         * @param poolId The ID of the pool with which the vip is associated.
         * Changing this updates the pool_id of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder poolId(Output<String> poolId) {
            $.poolId = poolId;
            return this;
        }

        /**
         * @param poolId The ID of the pool with which the vip is associated.
         * Changing this updates the pool_id of the existing vip.
         * 
         * @return builder
         * 
         */
        public Builder poolId(String poolId) {
            return poolId(Output.of(poolId));
        }

        /**
         * @param port The port on which to listen for client traffic. Changing
         * this creates a new vip.
         * 
         * @return builder
         * 
         */
        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port on which to listen for client traffic. Changing
         * this creates a new vip.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol The protocol - can be either &#39;TCP, &#39;HTTP&#39;, or
         * HTTPS&#39;. Changing this creates a new vip.
         * 
         * @return builder
         * 
         */
        public Builder protocol(Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol - can be either &#39;TCP, &#39;HTTP&#39;, or
         * HTTPS&#39;. Changing this creates a new vip.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a VIP. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * VIP.
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
         * A Networking client is needed to create a VIP. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * VIP.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param subnetId The network on which to allocate the vip&#39;s address. A
         * tenant can only create vips on networks authorized by policy (e.g. networks
         * that belong to them or networks that are shared). Changing this creates a
         * new vip.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId The network on which to allocate the vip&#39;s address. A
         * tenant can only create vips on networks authorized by policy (e.g. networks
         * that belong to them or networks that are shared). Changing this creates a
         * new vip.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        /**
         * @param tenantId The owner of the vip. Required if admin wants to
         * create a vip member for another tenant. Changing this creates a new vip.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the vip. Required if admin wants to
         * create a vip member for another tenant. Changing this creates a new vip.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public VipArgs build() {
            if ($.poolId == null) {
                throw new MissingRequiredPropertyException("VipArgs", "poolId");
            }
            if ($.port == null) {
                throw new MissingRequiredPropertyException("VipArgs", "port");
            }
            if ($.protocol == null) {
                throw new MissingRequiredPropertyException("VipArgs", "protocol");
            }
            if ($.subnetId == null) {
                throw new MissingRequiredPropertyException("VipArgs", "subnetId");
            }
            return $;
        }
    }

}
