// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.networking.inputs.PortAllowedAddressPairArgs;
import com.pulumi.openstack.networking.inputs.PortBindingArgs;
import com.pulumi.openstack.networking.inputs.PortExtraDhcpOptionArgs;
import com.pulumi.openstack.networking.inputs.PortFixedIpArgs;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PortArgs extends com.pulumi.resources.ResourceArgs {

    public static final PortArgs Empty = new PortArgs();

    /**
     * Administrative up/down status for the port
     * (must be `true` or `false` if provided). Changing this updates the
     * `admin_state_up` of an existing port.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the port
     * (must be `true` or `false` if provided). Changing this updates the
     * `admin_state_up` of an existing port.
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     * 
     */
    @Import(name="allowedAddressPairs")
    private @Nullable Output<List<PortAllowedAddressPairArgs>> allowedAddressPairs;

    /**
     * @return An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     * 
     */
    public Optional<Output<List<PortAllowedAddressPairArgs>>> allowedAddressPairs() {
        return Optional.ofNullable(this.allowedAddressPairs);
    }

    /**
     * The port binding allows to specify binding information
     * for the port. The structure is described below.
     * 
     */
    @Import(name="binding")
    private @Nullable Output<PortBindingArgs> binding;

    /**
     * @return The port binding allows to specify binding information
     * for the port. The structure is described below.
     * 
     */
    public Optional<Output<PortBindingArgs>> binding() {
        return Optional.ofNullable(this.binding);
    }

    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing port.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description of the port. Changing
     * this updates the `description` of an existing port.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     * 
     */
    @Import(name="deviceId")
    private @Nullable Output<String> deviceId;

    /**
     * @return The ID of the device attached to the port. Changing this
     * creates a new port.
     * 
     */
    public Optional<Output<String>> deviceId() {
        return Optional.ofNullable(this.deviceId);
    }

    /**
     * The device owner of the port. Changing this creates
     * a new port.
     * 
     */
    @Import(name="deviceOwner")
    private @Nullable Output<String> deviceOwner;

    /**
     * @return The device owner of the port. Changing this creates
     * a new port.
     * 
     */
    public Optional<Output<String>> deviceOwner() {
        return Optional.ofNullable(this.deviceOwner);
    }

    /**
     * The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     * 
     */
    @Import(name="dnsName")
    private @Nullable Output<String> dnsName;

    /**
     * @return The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     * 
     */
    public Optional<Output<String>> dnsName() {
        return Optional.ofNullable(this.dnsName);
    }

    /**
     * An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     * 
     */
    @Import(name="extraDhcpOptions")
    private @Nullable Output<List<PortExtraDhcpOptionArgs>> extraDhcpOptions;

    /**
     * @return An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     * 
     */
    public Optional<Output<List<PortExtraDhcpOptionArgs>>> extraDhcpOptions() {
        return Optional.ofNullable(this.extraDhcpOptions);
    }

    /**
     * An array of desired IPs for
     * this port. The structure is described below.
     * 
     */
    @Import(name="fixedIps")
    private @Nullable Output<List<PortFixedIpArgs>> fixedIps;

    /**
     * @return An array of desired IPs for
     * this port. The structure is described below.
     * 
     */
    public Optional<Output<List<PortFixedIpArgs>>> fixedIps() {
        return Optional.ofNullable(this.fixedIps);
    }

    /**
     * Specify a specific MAC address for the port. Changing
     * this creates a new port.
     * 
     */
    @Import(name="macAddress")
    private @Nullable Output<String> macAddress;

    /**
     * @return Specify a specific MAC address for the port. Changing
     * this creates a new port.
     * 
     */
    public Optional<Output<String>> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }

    /**
     * A unique name for the port. Changing this
     * updates the `name` of an existing port.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name for the port. Changing this
     * updates the `name` of an existing port.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the network to attach the port to. Changing
     * this creates a new port.
     * 
     */
    @Import(name="networkId", required=true)
    private Output<String> networkId;

    /**
     * @return The ID of the network to attach the port to. Changing
     * this creates a new port.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }

    /**
     * Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     * 
     */
    @Import(name="noFixedIp")
    private @Nullable Output<Boolean> noFixedIp;

    /**
     * @return Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     * 
     */
    public Optional<Output<Boolean>> noFixedIp() {
        return Optional.ofNullable(this.noFixedIp);
    }

    /**
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the port will yield to the default
     * behavior of the Networking service, which is to usually apply the &#34;default&#34;
     * security group.
     * 
     */
    @Import(name="noSecurityGroups")
    private @Nullable Output<Boolean> noSecurityGroups;

    /**
     * @return If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the port will yield to the default
     * behavior of the Networking service, which is to usually apply the &#34;default&#34;
     * security group.
     * 
     */
    public Optional<Output<Boolean>> noSecurityGroups() {
        return Optional.ofNullable(this.noSecurityGroups);
    }

    /**
     * Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of `true`. Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     * 
     */
    @Import(name="portSecurityEnabled")
    private @Nullable Output<Boolean> portSecurityEnabled;

    /**
     * @return Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of `true`. Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     * 
     */
    public Optional<Output<Boolean>> portSecurityEnabled() {
        return Optional.ofNullable(this.portSecurityEnabled);
    }

    /**
     * Reference to the associated QoS policy.
     * 
     */
    @Import(name="qosPolicyId")
    private @Nullable Output<String> qosPolicyId;

    /**
     * @return Reference to the associated QoS policy.
     * 
     */
    public Optional<Output<String>> qosPolicyId() {
        return Optional.ofNullable(this.qosPolicyId);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     * 
     */
    @Import(name="securityGroupIds")
    private @Nullable Output<List<String>> securityGroupIds;

    /**
     * @return A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     * 
     */
    public Optional<Output<List<String>>> securityGroupIds() {
        return Optional.ofNullable(this.securityGroupIds);
    }

    /**
     * A set of string tags for the port.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of string tags for the port.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The owner of the port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
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

    private PortArgs() {}

    private PortArgs(PortArgs $) {
        this.adminStateUp = $.adminStateUp;
        this.allowedAddressPairs = $.allowedAddressPairs;
        this.binding = $.binding;
        this.description = $.description;
        this.deviceId = $.deviceId;
        this.deviceOwner = $.deviceOwner;
        this.dnsName = $.dnsName;
        this.extraDhcpOptions = $.extraDhcpOptions;
        this.fixedIps = $.fixedIps;
        this.macAddress = $.macAddress;
        this.name = $.name;
        this.networkId = $.networkId;
        this.noFixedIp = $.noFixedIp;
        this.noSecurityGroups = $.noSecurityGroups;
        this.portSecurityEnabled = $.portSecurityEnabled;
        this.qosPolicyId = $.qosPolicyId;
        this.region = $.region;
        this.securityGroupIds = $.securityGroupIds;
        this.tags = $.tags;
        this.tenantId = $.tenantId;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PortArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PortArgs $;

        public Builder() {
            $ = new PortArgs();
        }

        public Builder(PortArgs defaults) {
            $ = new PortArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp Administrative up/down status for the port
         * (must be `true` or `false` if provided). Changing this updates the
         * `admin_state_up` of an existing port.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp Administrative up/down status for the port
         * (must be `true` or `false` if provided). Changing this updates the
         * `admin_state_up` of an existing port.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param allowedAddressPairs An IP/MAC Address pair of additional IP
         * addresses that can be active on this port. The structure is described
         * below.
         * 
         * @return builder
         * 
         */
        public Builder allowedAddressPairs(@Nullable Output<List<PortAllowedAddressPairArgs>> allowedAddressPairs) {
            $.allowedAddressPairs = allowedAddressPairs;
            return this;
        }

        /**
         * @param allowedAddressPairs An IP/MAC Address pair of additional IP
         * addresses that can be active on this port. The structure is described
         * below.
         * 
         * @return builder
         * 
         */
        public Builder allowedAddressPairs(List<PortAllowedAddressPairArgs> allowedAddressPairs) {
            return allowedAddressPairs(Output.of(allowedAddressPairs));
        }

        /**
         * @param allowedAddressPairs An IP/MAC Address pair of additional IP
         * addresses that can be active on this port. The structure is described
         * below.
         * 
         * @return builder
         * 
         */
        public Builder allowedAddressPairs(PortAllowedAddressPairArgs... allowedAddressPairs) {
            return allowedAddressPairs(List.of(allowedAddressPairs));
        }

        /**
         * @param binding The port binding allows to specify binding information
         * for the port. The structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder binding(@Nullable Output<PortBindingArgs> binding) {
            $.binding = binding;
            return this;
        }

        /**
         * @param binding The port binding allows to specify binding information
         * for the port. The structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder binding(PortBindingArgs binding) {
            return binding(Output.of(binding));
        }

        /**
         * @param description Human-readable description of the port. Changing
         * this updates the `description` of an existing port.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description of the port. Changing
         * this updates the `description` of an existing port.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param deviceId The ID of the device attached to the port. Changing this
         * creates a new port.
         * 
         * @return builder
         * 
         */
        public Builder deviceId(@Nullable Output<String> deviceId) {
            $.deviceId = deviceId;
            return this;
        }

        /**
         * @param deviceId The ID of the device attached to the port. Changing this
         * creates a new port.
         * 
         * @return builder
         * 
         */
        public Builder deviceId(String deviceId) {
            return deviceId(Output.of(deviceId));
        }

        /**
         * @param deviceOwner The device owner of the port. Changing this creates
         * a new port.
         * 
         * @return builder
         * 
         */
        public Builder deviceOwner(@Nullable Output<String> deviceOwner) {
            $.deviceOwner = deviceOwner;
            return this;
        }

        /**
         * @param deviceOwner The device owner of the port. Changing this creates
         * a new port.
         * 
         * @return builder
         * 
         */
        public Builder deviceOwner(String deviceOwner) {
            return deviceOwner(Output.of(deviceOwner));
        }

        /**
         * @param dnsName The port DNS name. Available, when Neutron DNS extension
         * is enabled.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(@Nullable Output<String> dnsName) {
            $.dnsName = dnsName;
            return this;
        }

        /**
         * @param dnsName The port DNS name. Available, when Neutron DNS extension
         * is enabled.
         * 
         * @return builder
         * 
         */
        public Builder dnsName(String dnsName) {
            return dnsName(Output.of(dnsName));
        }

        /**
         * @param extraDhcpOptions An extra DHCP option that needs to be configured
         * on the port. The structure is described below. Can be specified multiple
         * times.
         * 
         * @return builder
         * 
         */
        public Builder extraDhcpOptions(@Nullable Output<List<PortExtraDhcpOptionArgs>> extraDhcpOptions) {
            $.extraDhcpOptions = extraDhcpOptions;
            return this;
        }

        /**
         * @param extraDhcpOptions An extra DHCP option that needs to be configured
         * on the port. The structure is described below. Can be specified multiple
         * times.
         * 
         * @return builder
         * 
         */
        public Builder extraDhcpOptions(List<PortExtraDhcpOptionArgs> extraDhcpOptions) {
            return extraDhcpOptions(Output.of(extraDhcpOptions));
        }

        /**
         * @param extraDhcpOptions An extra DHCP option that needs to be configured
         * on the port. The structure is described below. Can be specified multiple
         * times.
         * 
         * @return builder
         * 
         */
        public Builder extraDhcpOptions(PortExtraDhcpOptionArgs... extraDhcpOptions) {
            return extraDhcpOptions(List.of(extraDhcpOptions));
        }

        /**
         * @param fixedIps An array of desired IPs for
         * this port. The structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder fixedIps(@Nullable Output<List<PortFixedIpArgs>> fixedIps) {
            $.fixedIps = fixedIps;
            return this;
        }

        /**
         * @param fixedIps An array of desired IPs for
         * this port. The structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder fixedIps(List<PortFixedIpArgs> fixedIps) {
            return fixedIps(Output.of(fixedIps));
        }

        /**
         * @param fixedIps An array of desired IPs for
         * this port. The structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder fixedIps(PortFixedIpArgs... fixedIps) {
            return fixedIps(List.of(fixedIps));
        }

        /**
         * @param macAddress Specify a specific MAC address for the port. Changing
         * this creates a new port.
         * 
         * @return builder
         * 
         */
        public Builder macAddress(@Nullable Output<String> macAddress) {
            $.macAddress = macAddress;
            return this;
        }

        /**
         * @param macAddress Specify a specific MAC address for the port. Changing
         * this creates a new port.
         * 
         * @return builder
         * 
         */
        public Builder macAddress(String macAddress) {
            return macAddress(Output.of(macAddress));
        }

        /**
         * @param name A unique name for the port. Changing this
         * updates the `name` of an existing port.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name for the port. Changing this
         * updates the `name` of an existing port.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param networkId The ID of the network to attach the port to. Changing
         * this creates a new port.
         * 
         * @return builder
         * 
         */
        public Builder networkId(Output<String> networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param networkId The ID of the network to attach the port to. Changing
         * this creates a new port.
         * 
         * @return builder
         * 
         */
        public Builder networkId(String networkId) {
            return networkId(Output.of(networkId));
        }

        /**
         * @param noFixedIp Create a port with no fixed
         * IP address. This will also remove any fixed IPs previously set on a port. `true`
         * is the only valid value for this argument.
         * 
         * @return builder
         * 
         */
        public Builder noFixedIp(@Nullable Output<Boolean> noFixedIp) {
            $.noFixedIp = noFixedIp;
            return this;
        }

        /**
         * @param noFixedIp Create a port with no fixed
         * IP address. This will also remove any fixed IPs previously set on a port. `true`
         * is the only valid value for this argument.
         * 
         * @return builder
         * 
         */
        public Builder noFixedIp(Boolean noFixedIp) {
            return noFixedIp(Output.of(noFixedIp));
        }

        /**
         * @param noSecurityGroups If set to
         * `true`, then no security groups are applied to the port. If set to `false` and
         * no `security_group_ids` are specified, then the port will yield to the default
         * behavior of the Networking service, which is to usually apply the &#34;default&#34;
         * security group.
         * 
         * @return builder
         * 
         */
        public Builder noSecurityGroups(@Nullable Output<Boolean> noSecurityGroups) {
            $.noSecurityGroups = noSecurityGroups;
            return this;
        }

        /**
         * @param noSecurityGroups If set to
         * `true`, then no security groups are applied to the port. If set to `false` and
         * no `security_group_ids` are specified, then the port will yield to the default
         * behavior of the Networking service, which is to usually apply the &#34;default&#34;
         * security group.
         * 
         * @return builder
         * 
         */
        public Builder noSecurityGroups(Boolean noSecurityGroups) {
            return noSecurityGroups(Output.of(noSecurityGroups));
        }

        /**
         * @param portSecurityEnabled Whether to explicitly enable or disable
         * port security on the port. Port Security is usually enabled by default, so
         * omitting argument will usually result in a value of `true`. Setting this
         * explicitly to `false` will disable port security. In order to disable port
         * security, the port must not have any security groups. Valid values are `true`
         * and `false`.
         * 
         * @return builder
         * 
         */
        public Builder portSecurityEnabled(@Nullable Output<Boolean> portSecurityEnabled) {
            $.portSecurityEnabled = portSecurityEnabled;
            return this;
        }

        /**
         * @param portSecurityEnabled Whether to explicitly enable or disable
         * port security on the port. Port Security is usually enabled by default, so
         * omitting argument will usually result in a value of `true`. Setting this
         * explicitly to `false` will disable port security. In order to disable port
         * security, the port must not have any security groups. Valid values are `true`
         * and `false`.
         * 
         * @return builder
         * 
         */
        public Builder portSecurityEnabled(Boolean portSecurityEnabled) {
            return portSecurityEnabled(Output.of(portSecurityEnabled));
        }

        /**
         * @param qosPolicyId Reference to the associated QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(@Nullable Output<String> qosPolicyId) {
            $.qosPolicyId = qosPolicyId;
            return this;
        }

        /**
         * @param qosPolicyId Reference to the associated QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(String qosPolicyId) {
            return qosPolicyId(Output.of(qosPolicyId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a port. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * port.
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
         * A Networking client is needed to create a port. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * port.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param securityGroupIds A list
         * of security group IDs to apply to the port. The security groups must be
         * specified by ID and not name (as opposed to how they are configured with
         * the Compute Instance).
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(@Nullable Output<List<String>> securityGroupIds) {
            $.securityGroupIds = securityGroupIds;
            return this;
        }

        /**
         * @param securityGroupIds A list
         * of security group IDs to apply to the port. The security groups must be
         * specified by ID and not name (as opposed to how they are configured with
         * the Compute Instance).
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(List<String> securityGroupIds) {
            return securityGroupIds(Output.of(securityGroupIds));
        }

        /**
         * @param securityGroupIds A list
         * of security group IDs to apply to the port. The security groups must be
         * specified by ID and not name (as opposed to how they are configured with
         * the Compute Instance).
         * 
         * @return builder
         * 
         */
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }

        /**
         * @param tags A set of string tags for the port.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of string tags for the port.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of string tags for the port.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param tenantId The owner of the port. Required if admin wants
         * to create a port for another tenant. Changing this creates a new port.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the port. Required if admin wants
         * to create a port for another tenant. Changing this creates a new port.
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

        public PortArgs build() {
            if ($.networkId == null) {
                throw new MissingRequiredPropertyException("PortArgs", "networkId");
            }
            return $;
        }
    }

}
