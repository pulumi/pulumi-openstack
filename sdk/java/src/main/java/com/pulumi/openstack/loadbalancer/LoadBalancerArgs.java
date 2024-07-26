// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LoadBalancerArgs extends com.pulumi.resources.ResourceArgs {

    public static final LoadBalancerArgs Empty = new LoadBalancerArgs();

    /**
     * The administrative state of the Loadbalancer.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the Loadbalancer.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * The availability zone of the Loadbalancer.
     * Changing this creates a new loadbalancer. Available only for Octavia
     * **minor version 2.14 or later**.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return The availability zone of the Loadbalancer.
     * Changing this creates a new loadbalancer. Available only for Octavia
     * **minor version 2.14 or later**.
     * 
     */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * Human-readable description for the Loadbalancer.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description for the Loadbalancer.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The UUID of a flavor. Changing this creates a new
     * loadbalancer.
     * 
     */
    @Import(name="flavorId")
    private @Nullable Output<String> flavorId;

    /**
     * @return The UUID of a flavor. Changing this creates a new
     * loadbalancer.
     * 
     */
    public Optional<Output<String>> flavorId() {
        return Optional.ofNullable(this.flavorId);
    }

    /**
     * The name of the provider. Changing this
     * creates a new loadbalancer.
     * 
     */
    @Import(name="loadbalancerProvider")
    private @Nullable Output<String> loadbalancerProvider;

    /**
     * @return The name of the provider. Changing this
     * creates a new loadbalancer.
     * 
     */
    public Optional<Output<String>> loadbalancerProvider() {
        return Optional.ofNullable(this.loadbalancerProvider);
    }

    /**
     * Human-readable name for the Loadbalancer. Does not have
     * to be unique.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Human-readable name for the Loadbalancer. Does not have
     * to be unique.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * A list of security group IDs to apply to the
     * loadbalancer. The security groups must be specified by ID and not name (as
     * opposed to how they are configured with the Compute Instance).
     * 
     */
    @Import(name="securityGroupIds")
    private @Nullable Output<List<String>> securityGroupIds;

    /**
     * @return A list of security group IDs to apply to the
     * loadbalancer. The security groups must be specified by ID and not name (as
     * opposed to how they are configured with the Compute Instance).
     * 
     */
    public Optional<Output<List<String>>> securityGroupIds() {
        return Optional.ofNullable(this.securityGroupIds);
    }

    /**
     * A list of simple strings assigned to the loadbalancer.
     * Available only for Octavia **minor version 2.5 or later**.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of simple strings assigned to the loadbalancer.
     * Available only for Octavia **minor version 2.5 or later**.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Required for admins. The UUID of the tenant who owns
     * the Loadbalancer.  Only administrative users can specify a tenant UUID
     * other than their own.  Changing this creates a new loadbalancer.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the Loadbalancer.  Only administrative users can specify a tenant UUID
     * other than their own.  Changing this creates a new loadbalancer.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * The ip address of the load balancer.
     * Changing this creates a new loadbalancer.
     * 
     */
    @Import(name="vipAddress")
    private @Nullable Output<String> vipAddress;

    /**
     * @return The ip address of the load balancer.
     * Changing this creates a new loadbalancer.
     * 
     */
    public Optional<Output<String>> vipAddress() {
        return Optional.ofNullable(this.vipAddress);
    }

    /**
     * The network on which to allocate the
     * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer. Exactly one of
     * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
     * 
     */
    @Import(name="vipNetworkId")
    private @Nullable Output<String> vipNetworkId;

    /**
     * @return The network on which to allocate the
     * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer. Exactly one of
     * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
     * 
     */
    public Optional<Output<String>> vipNetworkId() {
        return Optional.ofNullable(this.vipNetworkId);
    }

    /**
     * The port UUID that the loadbalancer will use.
     * Changing this creates a new loadbalancer. Exactly one of
     * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
     * 
     */
    @Import(name="vipPortId")
    private @Nullable Output<String> vipPortId;

    /**
     * @return The port UUID that the loadbalancer will use.
     * Changing this creates a new loadbalancer. Exactly one of
     * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
     * 
     */
    public Optional<Output<String>> vipPortId() {
        return Optional.ofNullable(this.vipPortId);
    }

    /**
     * The ID of the QoS Policy which will
     * be applied to the Virtual IP (VIP).
     * 
     */
    @Import(name="vipQosPolicyId")
    private @Nullable Output<String> vipQosPolicyId;

    /**
     * @return The ID of the QoS Policy which will
     * be applied to the Virtual IP (VIP).
     * 
     */
    public Optional<Output<String>> vipQosPolicyId() {
        return Optional.ofNullable(this.vipQosPolicyId);
    }

    /**
     * The subnet on which to allocate the
     * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer. Exactly one of
     * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
     * 
     */
    @Import(name="vipSubnetId")
    private @Nullable Output<String> vipSubnetId;

    /**
     * @return The subnet on which to allocate the
     * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer. Exactly one of
     * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
     * 
     */
    public Optional<Output<String>> vipSubnetId() {
        return Optional.ofNullable(this.vipSubnetId);
    }

    private LoadBalancerArgs() {}

    private LoadBalancerArgs(LoadBalancerArgs $) {
        this.adminStateUp = $.adminStateUp;
        this.availabilityZone = $.availabilityZone;
        this.description = $.description;
        this.flavorId = $.flavorId;
        this.loadbalancerProvider = $.loadbalancerProvider;
        this.name = $.name;
        this.region = $.region;
        this.securityGroupIds = $.securityGroupIds;
        this.tags = $.tags;
        this.tenantId = $.tenantId;
        this.vipAddress = $.vipAddress;
        this.vipNetworkId = $.vipNetworkId;
        this.vipPortId = $.vipPortId;
        this.vipQosPolicyId = $.vipQosPolicyId;
        this.vipSubnetId = $.vipSubnetId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LoadBalancerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LoadBalancerArgs $;

        public Builder() {
            $ = new LoadBalancerArgs();
        }

        public Builder(LoadBalancerArgs defaults) {
            $ = new LoadBalancerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp The administrative state of the Loadbalancer.
         * A valid value is true (UP) or false (DOWN).
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp The administrative state of the Loadbalancer.
         * A valid value is true (UP) or false (DOWN).
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param availabilityZone The availability zone of the Loadbalancer.
         * Changing this creates a new loadbalancer. Available only for Octavia
         * **minor version 2.14 or later**.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The availability zone of the Loadbalancer.
         * Changing this creates a new loadbalancer. Available only for Octavia
         * **minor version 2.14 or later**.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param description Human-readable description for the Loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description for the Loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param flavorId The UUID of a flavor. Changing this creates a new
         * loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder flavorId(@Nullable Output<String> flavorId) {
            $.flavorId = flavorId;
            return this;
        }

        /**
         * @param flavorId The UUID of a flavor. Changing this creates a new
         * loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder flavorId(String flavorId) {
            return flavorId(Output.of(flavorId));
        }

        /**
         * @param loadbalancerProvider The name of the provider. Changing this
         * creates a new loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder loadbalancerProvider(@Nullable Output<String> loadbalancerProvider) {
            $.loadbalancerProvider = loadbalancerProvider;
            return this;
        }

        /**
         * @param loadbalancerProvider The name of the provider. Changing this
         * creates a new loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder loadbalancerProvider(String loadbalancerProvider) {
            return loadbalancerProvider(Output.of(loadbalancerProvider));
        }

        /**
         * @param name Human-readable name for the Loadbalancer. Does not have
         * to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Human-readable name for the Loadbalancer. Does not have
         * to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create an LB member. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * LB member.
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
         * A Networking client is needed to create an LB member. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * LB member.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param securityGroupIds A list of security group IDs to apply to the
         * loadbalancer. The security groups must be specified by ID and not name (as
         * opposed to how they are configured with the Compute Instance).
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(@Nullable Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        /**
         * @param securityGroupIds A list of security group IDs to apply to the
         * loadbalancer. The security groups must be specified by ID and not name (as
         * opposed to how they are configured with the Compute Instance).
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        /**
         * @param securityGroupIds A list of security group IDs to apply to the
         * loadbalancer. The security groups must be specified by ID and not name (as
         * opposed to how they are configured with the Compute Instance).
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param tags A list of simple strings assigned to the loadbalancer.
         * Available only for Octavia **minor version 2.5 or later**.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of simple strings assigned to the loadbalancer.
         * Available only for Octavia **minor version 2.5 or later**.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of simple strings assigned to the loadbalancer.
         * Available only for Octavia **minor version 2.5 or later**.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param tenantId Required for admins. The UUID of the tenant who owns
         * the Loadbalancer.  Only administrative users can specify a tenant UUID
         * other than their own.  Changing this creates a new loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId Required for admins. The UUID of the tenant who owns
         * the Loadbalancer.  Only administrative users can specify a tenant UUID
         * other than their own.  Changing this creates a new loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param vipAddress The ip address of the load balancer.
         * Changing this creates a new loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder vipAddress(@Nullable Output<String> vipAddress) {
            $.vipAddress = vipAddress;
            return this;
        }

        /**
         * @param vipAddress The ip address of the load balancer.
         * Changing this creates a new loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder vipAddress(String vipAddress) {
            return vipAddress(Output.of(vipAddress));
        }

        /**
         * @param vipNetworkId The network on which to allocate the
         * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
         * authorized by policy (e.g. networks that belong to them or networks that
         * are shared).  Changing this creates a new loadbalancer. Exactly one of
         * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
         * 
         * @return builder
         * 
         */
        public Builder vipNetworkId(@Nullable Output<String> vipNetworkId) {
            $.vipNetworkId = vipNetworkId;
            return this;
        }

        /**
         * @param vipNetworkId The network on which to allocate the
         * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
         * authorized by policy (e.g. networks that belong to them or networks that
         * are shared).  Changing this creates a new loadbalancer. Exactly one of
         * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
         * 
         * @return builder
         * 
         */
        public Builder vipNetworkId(String vipNetworkId) {
            return vipNetworkId(Output.of(vipNetworkId));
        }

        /**
         * @param vipPortId The port UUID that the loadbalancer will use.
         * Changing this creates a new loadbalancer. Exactly one of
         * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
         * 
         * @return builder
         * 
         */
        public Builder vipPortId(@Nullable Output<String> vipPortId) {
            $.vipPortId = vipPortId;
            return this;
        }

        /**
         * @param vipPortId The port UUID that the loadbalancer will use.
         * Changing this creates a new loadbalancer. Exactly one of
         * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
         * 
         * @return builder
         * 
         */
        public Builder vipPortId(String vipPortId) {
            return vipPortId(Output.of(vipPortId));
        }

        /**
         * @param vipQosPolicyId The ID of the QoS Policy which will
         * be applied to the Virtual IP (VIP).
         * 
         * @return builder
         * 
         */
        public Builder vipQosPolicyId(@Nullable Output<String> vipQosPolicyId) {
            $.vipQosPolicyId = vipQosPolicyId;
            return this;
        }

        /**
         * @param vipQosPolicyId The ID of the QoS Policy which will
         * be applied to the Virtual IP (VIP).
         * 
         * @return builder
         * 
         */
        public Builder vipQosPolicyId(String vipQosPolicyId) {
            return vipQosPolicyId(Output.of(vipQosPolicyId));
        }

        /**
         * @param vipSubnetId The subnet on which to allocate the
         * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
         * authorized by policy (e.g. networks that belong to them or networks that
         * are shared).  Changing this creates a new loadbalancer. Exactly one of
         * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
         * 
         * @return builder
         * 
         */
        public Builder vipSubnetId(@Nullable Output<String> vipSubnetId) {
            $.vipSubnetId = vipSubnetId;
            return this;
        }

        /**
         * @param vipSubnetId The subnet on which to allocate the
         * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
         * authorized by policy (e.g. networks that belong to them or networks that
         * are shared).  Changing this creates a new loadbalancer. Exactly one of
         * `vip_subnet_id`, `vip_network_id` or `vip_port_id` has to be defined.
         * 
         * @return builder
         * 
         */
        public Builder vipSubnetId(String vipSubnetId) {
            return vipSubnetId(Output.of(vipSubnetId));
        }

        public LoadBalancerArgs build() {
            return $;
        }
    }

}
