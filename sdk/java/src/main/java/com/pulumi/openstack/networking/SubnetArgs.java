// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.networking.inputs.SubnetAllocationPoolArgs;
import com.pulumi.openstack.networking.inputs.SubnetHostRouteArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SubnetArgs extends com.pulumi.resources.ResourceArgs {

    public static final SubnetArgs Empty = new SubnetArgs();

    /**
     * A block declaring the start and end range of the IP addresses available for
     * use with DHCP in this subnet.
     * The `allocation_pools` block is documented below.
     * 
     * @deprecated
     * use allocation_pool instead
     * 
     */
    @Deprecated /* use allocation_pool instead */
    @Import(name="allocationPools")
    private @Nullable Output<List<SubnetAllocationPoolArgs>> allocationPools;

    /**
     * @return A block declaring the start and end range of the IP addresses available for
     * use with DHCP in this subnet.
     * The `allocation_pools` block is documented below.
     * 
     * @deprecated
     * use allocation_pool instead
     * 
     */
    @Deprecated /* use allocation_pool instead */
    public Optional<Output<List<SubnetAllocationPoolArgs>>> allocationPools() {
        return Optional.ofNullable(this.allocationPools);
    }

    /**
     * CIDR representing IP range for this subnet, based on IP
     * version. You can omit this option if you are creating a subnet from a
     * subnet pool.
     * 
     */
    @Import(name="cidr")
    private @Nullable Output<String> cidr;

    /**
     * @return CIDR representing IP range for this subnet, based on IP
     * version. You can omit this option if you are creating a subnet from a
     * subnet pool.
     * 
     */
    public Optional<Output<String>> cidr() {
        return Optional.ofNullable(this.cidr);
    }

    /**
     * Human-readable description of the subnet. Changing this
     * updates the name of the existing subnet.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description of the subnet. Changing this
     * updates the name of the existing subnet.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * An array of DNS name server names used by hosts
     * in this subnet. Changing this updates the DNS name servers for the existing
     * subnet.
     * 
     */
    @Import(name="dnsNameservers")
    private @Nullable Output<List<String>> dnsNameservers;

    /**
     * @return An array of DNS name server names used by hosts
     * in this subnet. Changing this updates the DNS name servers for the existing
     * subnet.
     * 
     */
    public Optional<Output<List<String>>> dnsNameservers() {
        return Optional.ofNullable(this.dnsNameservers);
    }

    /**
     * The administrative state of the network.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value enables or
     * disables the DHCP capabilities of the existing subnet. Defaults to true.
     * 
     */
    @Import(name="enableDhcp")
    private @Nullable Output<Boolean> enableDhcp;

    /**
     * @return The administrative state of the network.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value enables or
     * disables the DHCP capabilities of the existing subnet. Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> enableDhcp() {
        return Optional.ofNullable(this.enableDhcp);
    }

    /**
     * Default gateway used by devices in this subnet.
     * Leaving this blank and not setting `no_gateway` will cause a default
     * gateway of `.1` to be used. Changing this updates the gateway IP of the
     * existing subnet.
     * 
     */
    @Import(name="gatewayIp")
    private @Nullable Output<String> gatewayIp;

    /**
     * @return Default gateway used by devices in this subnet.
     * Leaving this blank and not setting `no_gateway` will cause a default
     * gateway of `.1` to be used. Changing this updates the gateway IP of the
     * existing subnet.
     * 
     */
    public Optional<Output<String>> gatewayIp() {
        return Optional.ofNullable(this.gatewayIp);
    }

    /**
     * (**Deprecated** - use `openstack.networking.SubnetRoute`
     * instead) An array of routes that should be used by devices
     * with IPs from this subnet (not including local subnet route). The host_route
     * object structure is documented below. Changing this updates the host routes
     * for the existing subnet.
     * 
     * @deprecated
     * Use openstack.networking.SubnetRoute instead
     * 
     */
    @Deprecated /* Use openstack.networking.SubnetRoute instead */
    @Import(name="hostRoutes")
    private @Nullable Output<List<SubnetHostRouteArgs>> hostRoutes;

    /**
     * @return (**Deprecated** - use `openstack.networking.SubnetRoute`
     * instead) An array of routes that should be used by devices
     * with IPs from this subnet (not including local subnet route). The host_route
     * object structure is documented below. Changing this updates the host routes
     * for the existing subnet.
     * 
     * @deprecated
     * Use openstack.networking.SubnetRoute instead
     * 
     */
    @Deprecated /* Use openstack.networking.SubnetRoute instead */
    public Optional<Output<List<SubnetHostRouteArgs>>> hostRoutes() {
        return Optional.ofNullable(this.hostRoutes);
    }

    /**
     * IP version, either 4 (default) or 6. Changing this creates a
     * new subnet.
     * 
     */
    @Import(name="ipVersion")
    private @Nullable Output<Integer> ipVersion;

    /**
     * @return IP version, either 4 (default) or 6. Changing this creates a
     * new subnet.
     * 
     */
    public Optional<Output<Integer>> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }

    /**
     * The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     * 
     */
    @Import(name="ipv6AddressMode")
    private @Nullable Output<String> ipv6AddressMode;

    /**
     * @return The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     * 
     */
    public Optional<Output<String>> ipv6AddressMode() {
        return Optional.ofNullable(this.ipv6AddressMode);
    }

    /**
     * The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     * 
     */
    @Import(name="ipv6RaMode")
    private @Nullable Output<String> ipv6RaMode;

    /**
     * @return The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     * 
     */
    public Optional<Output<String>> ipv6RaMode() {
        return Optional.ofNullable(this.ipv6RaMode);
    }

    /**
     * The name of the subnet. Changing this updates the name of
     * the existing subnet.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the subnet. Changing this updates the name of
     * the existing subnet.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The UUID of the parent network. Changing this
     * creates a new subnet.
     * 
     */
    @Import(name="networkId", required=true)
    private Output<String> networkId;

    /**
     * @return The UUID of the parent network. Changing this
     * creates a new subnet.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }

    /**
     * Do not set a gateway IP on this subnet. Changing
     * this removes or adds a default gateway IP of the existing subnet.
     * 
     */
    @Import(name="noGateway")
    private @Nullable Output<Boolean> noGateway;

    /**
     * @return Do not set a gateway IP on this subnet. Changing
     * this removes or adds a default gateway IP of the existing subnet.
     * 
     */
    public Optional<Output<Boolean>> noGateway() {
        return Optional.ofNullable(this.noGateway);
    }

    /**
     * The prefix length to use when creating a subnet
     * from a subnet pool. The default subnet pool prefix length that was defined
     * when creating the subnet pool will be used if not provided. Changing this
     * creates a new subnet.
     * 
     */
    @Import(name="prefixLength")
    private @Nullable Output<Integer> prefixLength;

    /**
     * @return The prefix length to use when creating a subnet
     * from a subnet pool. The default subnet pool prefix length that was defined
     * when creating the subnet pool will be used if not provided. Changing this
     * creates a new subnet.
     * 
     */
    public Optional<Output<Integer>> prefixLength() {
        return Optional.ofNullable(this.prefixLength);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnet.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnet.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * An array of service types used by the subnet.
     * Changing this updates the service types for the existing subnet.
     * 
     */
    @Import(name="serviceTypes")
    private @Nullable Output<List<String>> serviceTypes;

    /**
     * @return An array of service types used by the subnet.
     * Changing this updates the service types for the existing subnet.
     * 
     */
    public Optional<Output<List<String>>> serviceTypes() {
        return Optional.ofNullable(this.serviceTypes);
    }

    /**
     * The ID of the subnetpool associated with the subnet.
     * 
     */
    @Import(name="subnetpoolId")
    private @Nullable Output<String> subnetpoolId;

    /**
     * @return The ID of the subnetpool associated with the subnet.
     * 
     */
    public Optional<Output<String>> subnetpoolId() {
        return Optional.ofNullable(this.subnetpoolId);
    }

    /**
     * A set of string tags for the subnet.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of string tags for the subnet.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The owner of the subnet. Required if admin wants to
     * create a subnet for another tenant. Changing this creates a new subnet.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the subnet. Required if admin wants to
     * create a subnet for another tenant. Changing this creates a new subnet.
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
    private @Nullable Output<Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional options.
     * 
     */
    public Optional<Output<Map<String,Object>>> valueSpecs() {
        return Optional.ofNullable(this.valueSpecs);
    }

    private SubnetArgs() {}

    private SubnetArgs(SubnetArgs $) {
        this.allocationPools = $.allocationPools;
        this.cidr = $.cidr;
        this.description = $.description;
        this.dnsNameservers = $.dnsNameservers;
        this.enableDhcp = $.enableDhcp;
        this.gatewayIp = $.gatewayIp;
        this.hostRoutes = $.hostRoutes;
        this.ipVersion = $.ipVersion;
        this.ipv6AddressMode = $.ipv6AddressMode;
        this.ipv6RaMode = $.ipv6RaMode;
        this.name = $.name;
        this.networkId = $.networkId;
        this.noGateway = $.noGateway;
        this.prefixLength = $.prefixLength;
        this.region = $.region;
        this.serviceTypes = $.serviceTypes;
        this.subnetpoolId = $.subnetpoolId;
        this.tags = $.tags;
        this.tenantId = $.tenantId;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubnetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubnetArgs $;

        public Builder() {
            $ = new SubnetArgs();
        }

        public Builder(SubnetArgs defaults) {
            $ = new SubnetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allocationPools A block declaring the start and end range of the IP addresses available for
         * use with DHCP in this subnet.
         * The `allocation_pools` block is documented below.
         * 
         * @return builder
         * 
         * @deprecated
         * use allocation_pool instead
         * 
         */
        @Deprecated /* use allocation_pool instead */
        public Builder allocationPools(@Nullable Output<List<SubnetAllocationPoolArgs>> allocationPools) {
            $.allocationPools = allocationPools;
            return this;
        }

        /**
         * @param allocationPools A block declaring the start and end range of the IP addresses available for
         * use with DHCP in this subnet.
         * The `allocation_pools` block is documented below.
         * 
         * @return builder
         * 
         * @deprecated
         * use allocation_pool instead
         * 
         */
        @Deprecated /* use allocation_pool instead */
        public Builder allocationPools(List<SubnetAllocationPoolArgs> allocationPools) {
            return allocationPools(Output.of(allocationPools));
        }

        /**
         * @param allocationPools A block declaring the start and end range of the IP addresses available for
         * use with DHCP in this subnet.
         * The `allocation_pools` block is documented below.
         * 
         * @return builder
         * 
         * @deprecated
         * use allocation_pool instead
         * 
         */
        @Deprecated /* use allocation_pool instead */
        public Builder allocationPools(SubnetAllocationPoolArgs... allocationPools) {
            return allocationPools(List.of(allocationPools));
        }

        /**
         * @param cidr CIDR representing IP range for this subnet, based on IP
         * version. You can omit this option if you are creating a subnet from a
         * subnet pool.
         * 
         * @return builder
         * 
         */
        public Builder cidr(@Nullable Output<String> cidr) {
            $.cidr = cidr;
            return this;
        }

        /**
         * @param cidr CIDR representing IP range for this subnet, based on IP
         * version. You can omit this option if you are creating a subnet from a
         * subnet pool.
         * 
         * @return builder
         * 
         */
        public Builder cidr(String cidr) {
            return cidr(Output.of(cidr));
        }

        /**
         * @param description Human-readable description of the subnet. Changing this
         * updates the name of the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description of the subnet. Changing this
         * updates the name of the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dnsNameservers An array of DNS name server names used by hosts
         * in this subnet. Changing this updates the DNS name servers for the existing
         * subnet.
         * 
         * @return builder
         * 
         */
        public Builder dnsNameservers(@Nullable Output<List<String>> dnsNameservers) {
            $.dnsNameservers = dnsNameservers;
            return this;
        }

        /**
         * @param dnsNameservers An array of DNS name server names used by hosts
         * in this subnet. Changing this updates the DNS name servers for the existing
         * subnet.
         * 
         * @return builder
         * 
         */
        public Builder dnsNameservers(List<String> dnsNameservers) {
            return dnsNameservers(Output.of(dnsNameservers));
        }

        /**
         * @param dnsNameservers An array of DNS name server names used by hosts
         * in this subnet. Changing this updates the DNS name servers for the existing
         * subnet.
         * 
         * @return builder
         * 
         */
        public Builder dnsNameservers(String... dnsNameservers) {
            return dnsNameservers(List.of(dnsNameservers));
        }

        /**
         * @param enableDhcp The administrative state of the network.
         * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value enables or
         * disables the DHCP capabilities of the existing subnet. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder enableDhcp(@Nullable Output<Boolean> enableDhcp) {
            $.enableDhcp = enableDhcp;
            return this;
        }

        /**
         * @param enableDhcp The administrative state of the network.
         * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value enables or
         * disables the DHCP capabilities of the existing subnet. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder enableDhcp(Boolean enableDhcp) {
            return enableDhcp(Output.of(enableDhcp));
        }

        /**
         * @param gatewayIp Default gateway used by devices in this subnet.
         * Leaving this blank and not setting `no_gateway` will cause a default
         * gateway of `.1` to be used. Changing this updates the gateway IP of the
         * existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder gatewayIp(@Nullable Output<String> gatewayIp) {
            $.gatewayIp = gatewayIp;
            return this;
        }

        /**
         * @param gatewayIp Default gateway used by devices in this subnet.
         * Leaving this blank and not setting `no_gateway` will cause a default
         * gateway of `.1` to be used. Changing this updates the gateway IP of the
         * existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder gatewayIp(String gatewayIp) {
            return gatewayIp(Output.of(gatewayIp));
        }

        /**
         * @param hostRoutes (**Deprecated** - use `openstack.networking.SubnetRoute`
         * instead) An array of routes that should be used by devices
         * with IPs from this subnet (not including local subnet route). The host_route
         * object structure is documented below. Changing this updates the host routes
         * for the existing subnet.
         * 
         * @return builder
         * 
         * @deprecated
         * Use openstack.networking.SubnetRoute instead
         * 
         */
        @Deprecated /* Use openstack.networking.SubnetRoute instead */
        public Builder hostRoutes(@Nullable Output<List<SubnetHostRouteArgs>> hostRoutes) {
            $.hostRoutes = hostRoutes;
            return this;
        }

        /**
         * @param hostRoutes (**Deprecated** - use `openstack.networking.SubnetRoute`
         * instead) An array of routes that should be used by devices
         * with IPs from this subnet (not including local subnet route). The host_route
         * object structure is documented below. Changing this updates the host routes
         * for the existing subnet.
         * 
         * @return builder
         * 
         * @deprecated
         * Use openstack.networking.SubnetRoute instead
         * 
         */
        @Deprecated /* Use openstack.networking.SubnetRoute instead */
        public Builder hostRoutes(List<SubnetHostRouteArgs> hostRoutes) {
            return hostRoutes(Output.of(hostRoutes));
        }

        /**
         * @param hostRoutes (**Deprecated** - use `openstack.networking.SubnetRoute`
         * instead) An array of routes that should be used by devices
         * with IPs from this subnet (not including local subnet route). The host_route
         * object structure is documented below. Changing this updates the host routes
         * for the existing subnet.
         * 
         * @return builder
         * 
         * @deprecated
         * Use openstack.networking.SubnetRoute instead
         * 
         */
        @Deprecated /* Use openstack.networking.SubnetRoute instead */
        public Builder hostRoutes(SubnetHostRouteArgs... hostRoutes) {
            return hostRoutes(List.of(hostRoutes));
        }

        /**
         * @param ipVersion IP version, either 4 (default) or 6. Changing this creates a
         * new subnet.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(@Nullable Output<Integer> ipVersion) {
            $.ipVersion = ipVersion;
            return this;
        }

        /**
         * @param ipVersion IP version, either 4 (default) or 6. Changing this creates a
         * new subnet.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(Integer ipVersion) {
            return ipVersion(Output.of(ipVersion));
        }

        /**
         * @param ipv6AddressMode The IPv6 address mode. Valid values are
         * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressMode(@Nullable Output<String> ipv6AddressMode) {
            $.ipv6AddressMode = ipv6AddressMode;
            return this;
        }

        /**
         * @param ipv6AddressMode The IPv6 address mode. Valid values are
         * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6AddressMode(String ipv6AddressMode) {
            return ipv6AddressMode(Output.of(ipv6AddressMode));
        }

        /**
         * @param ipv6RaMode The IPv6 Router Advertisement mode. Valid values
         * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6RaMode(@Nullable Output<String> ipv6RaMode) {
            $.ipv6RaMode = ipv6RaMode;
            return this;
        }

        /**
         * @param ipv6RaMode The IPv6 Router Advertisement mode. Valid values
         * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6RaMode(String ipv6RaMode) {
            return ipv6RaMode(Output.of(ipv6RaMode));
        }

        /**
         * @param name The name of the subnet. Changing this updates the name of
         * the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the subnet. Changing this updates the name of
         * the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param networkId The UUID of the parent network. Changing this
         * creates a new subnet.
         * 
         * @return builder
         * 
         */
        public Builder networkId(Output<String> networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param networkId The UUID of the parent network. Changing this
         * creates a new subnet.
         * 
         * @return builder
         * 
         */
        public Builder networkId(String networkId) {
            return networkId(Output.of(networkId));
        }

        /**
         * @param noGateway Do not set a gateway IP on this subnet. Changing
         * this removes or adds a default gateway IP of the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder noGateway(@Nullable Output<Boolean> noGateway) {
            $.noGateway = noGateway;
            return this;
        }

        /**
         * @param noGateway Do not set a gateway IP on this subnet. Changing
         * this removes or adds a default gateway IP of the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder noGateway(Boolean noGateway) {
            return noGateway(Output.of(noGateway));
        }

        /**
         * @param prefixLength The prefix length to use when creating a subnet
         * from a subnet pool. The default subnet pool prefix length that was defined
         * when creating the subnet pool will be used if not provided. Changing this
         * creates a new subnet.
         * 
         * @return builder
         * 
         */
        public Builder prefixLength(@Nullable Output<Integer> prefixLength) {
            $.prefixLength = prefixLength;
            return this;
        }

        /**
         * @param prefixLength The prefix length to use when creating a subnet
         * from a subnet pool. The default subnet pool prefix length that was defined
         * when creating the subnet pool will be used if not provided. Changing this
         * creates a new subnet.
         * 
         * @return builder
         * 
         */
        public Builder prefixLength(Integer prefixLength) {
            return prefixLength(Output.of(prefixLength));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron subnet. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * subnet.
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
         * A Networking client is needed to create a Neutron subnet. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * subnet.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param serviceTypes An array of service types used by the subnet.
         * Changing this updates the service types for the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder serviceTypes(@Nullable Output<List<String>> serviceTypes) {
            $.serviceTypes = serviceTypes;
            return this;
        }

        /**
         * @param serviceTypes An array of service types used by the subnet.
         * Changing this updates the service types for the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder serviceTypes(List<String> serviceTypes) {
            return serviceTypes(Output.of(serviceTypes));
        }

        /**
         * @param serviceTypes An array of service types used by the subnet.
         * Changing this updates the service types for the existing subnet.
         * 
         * @return builder
         * 
         */
        public Builder serviceTypes(String... serviceTypes) {
            return serviceTypes(List.of(serviceTypes));
        }

        /**
         * @param subnetpoolId The ID of the subnetpool associated with the subnet.
         * 
         * @return builder
         * 
         */
        public Builder subnetpoolId(@Nullable Output<String> subnetpoolId) {
            $.subnetpoolId = subnetpoolId;
            return this;
        }

        /**
         * @param subnetpoolId The ID of the subnetpool associated with the subnet.
         * 
         * @return builder
         * 
         */
        public Builder subnetpoolId(String subnetpoolId) {
            return subnetpoolId(Output.of(subnetpoolId));
        }

        /**
         * @param tags A set of string tags for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of string tags for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of string tags for the subnet.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param tenantId The owner of the subnet. Required if admin wants to
         * create a subnet for another tenant. Changing this creates a new subnet.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the subnet. Required if admin wants to
         * create a subnet for another tenant. Changing this creates a new subnet.
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
        public Builder valueSpecs(@Nullable Output<Map<String,Object>> valueSpecs) {
            $.valueSpecs = valueSpecs;
            return this;
        }

        /**
         * @param valueSpecs Map of additional options.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(Map<String,Object> valueSpecs) {
            return valueSpecs(Output.of(valueSpecs));
        }

        public SubnetArgs build() {
            if ($.networkId == null) {
                throw new MissingRequiredPropertyException("SubnetArgs", "networkId");
            }
            return $;
        }
    }

}
