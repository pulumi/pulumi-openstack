// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.networking.inputs.NetworkSegmentArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkArgs Empty = new NetworkArgs();

    /**
     * The administrative state of the network.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
     * state of the existing network.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the network.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
     * state of the existing network.
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new network.
     * 
     */
    @Import(name="availabilityZoneHints")
    private @Nullable Output<List<String>> availabilityZoneHints;

    /**
     * @return An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new network.
     * 
     */
    public Optional<Output<List<String>>> availabilityZoneHints() {
        return Optional.ofNullable(this.availabilityZoneHints);
    }

    /**
     * Human-readable description of the network. Changing this
     * updates the name of the existing network.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description of the network. Changing this
     * updates the name of the existing network.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The network DNS domain. Available, when Neutron DNS
     * extension is enabled. The `dns_domain` of a network in conjunction with the
     * `dns_name` attribute of its ports will be published in an external DNS
     * service when Neutron is configured to integrate with such a service.
     * 
     */
    @Import(name="dnsDomain")
    private @Nullable Output<String> dnsDomain;

    /**
     * @return The network DNS domain. Available, when Neutron DNS
     * extension is enabled. The `dns_domain` of a network in conjunction with the
     * `dns_name` attribute of its ports will be published in an external DNS
     * service when Neutron is configured to integrate with such a service.
     * 
     */
    public Optional<Output<String>> dnsDomain() {
        return Optional.ofNullable(this.dnsDomain);
    }

    /**
     * Specifies whether the network resource has the
     * external routing facility. Valid values are true and false. Defaults to
     * false. Changing this updates the external attribute of the existing network.
     * 
     */
    @Import(name="external")
    private @Nullable Output<Boolean> external;

    /**
     * @return Specifies whether the network resource has the
     * external routing facility. Valid values are true and false. Defaults to
     * false. Changing this updates the external attribute of the existing network.
     * 
     */
    public Optional<Output<Boolean>> external() {
        return Optional.ofNullable(this.external);
    }

    /**
     * The network MTU. Available for read-only, when Neutron
     * `net-mtu` extension is enabled. Available for the modification, when
     * Neutron `net-mtu-writable` extension is enabled.
     * 
     */
    @Import(name="mtu")
    private @Nullable Output<Integer> mtu;

    /**
     * @return The network MTU. Available for read-only, when Neutron
     * `net-mtu` extension is enabled. Available for the modification, when
     * Neutron `net-mtu-writable` extension is enabled.
     * 
     */
    public Optional<Output<Integer>> mtu() {
        return Optional.ofNullable(this.mtu);
    }

    /**
     * The name of the network. Changing this updates the name of
     * the existing network.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the network. Changing this updates the name of
     * the existing network.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Whether to explicitly enable or disable
     * port security on the network. Port Security is usually enabled by default, so
     * omitting this argument will usually result in a value of &#34;true&#34;. Setting this
     * explicitly to `false` will disable port security. Valid values are `true` and
     * `false`.
     * 
     */
    @Import(name="portSecurityEnabled")
    private @Nullable Output<Boolean> portSecurityEnabled;

    /**
     * @return Whether to explicitly enable or disable
     * port security on the network. Port Security is usually enabled by default, so
     * omitting this argument will usually result in a value of &#34;true&#34;. Setting this
     * explicitly to `false` will disable port security. Valid values are `true` and
     * `false`.
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
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * network.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * network.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * An array of one or more provider segment objects.
     * Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
     * updating any provider related segments attributes. Check your plug-in whether
     * it supports updating.
     * 
     */
    @Import(name="segments")
    private @Nullable Output<List<NetworkSegmentArgs>> segments;

    /**
     * @return An array of one or more provider segment objects.
     * Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
     * updating any provider related segments attributes. Check your plug-in whether
     * it supports updating.
     * 
     */
    public Optional<Output<List<NetworkSegmentArgs>>> segments() {
        return Optional.ofNullable(this.segments);
    }

    /**
     * Specifies whether the network resource can be accessed
     * by any tenant or not. Changing this updates the sharing capabilities of the
     * existing network.
     * 
     */
    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    /**
     * @return Specifies whether the network resource can be accessed
     * by any tenant or not. Changing this updates the sharing capabilities of the
     * existing network.
     * 
     */
    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    /**
     * A set of string tags for the network.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of string tags for the network.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The owner of the network. Required if admin wants to
     * create a network for another tenant. Changing this creates a new network.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the network. Required if admin wants to
     * create a network for another tenant. Changing this creates a new network.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Specifies whether the network resource has the
     * VLAN transparent attribute set. Valid values are true and false. Defaults to
     * false. Changing this updates the `transparent_vlan` attribute of the existing
     * network.
     * 
     */
    @Import(name="transparentVlan")
    private @Nullable Output<Boolean> transparentVlan;

    /**
     * @return Specifies whether the network resource has the
     * VLAN transparent attribute set. Valid values are true and false. Defaults to
     * false. Changing this updates the `transparent_vlan` attribute of the existing
     * network.
     * 
     */
    public Optional<Output<Boolean>> transparentVlan() {
        return Optional.ofNullable(this.transparentVlan);
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

    private NetworkArgs() {}

    private NetworkArgs(NetworkArgs $) {
        this.adminStateUp = $.adminStateUp;
        this.availabilityZoneHints = $.availabilityZoneHints;
        this.description = $.description;
        this.dnsDomain = $.dnsDomain;
        this.external = $.external;
        this.mtu = $.mtu;
        this.name = $.name;
        this.portSecurityEnabled = $.portSecurityEnabled;
        this.qosPolicyId = $.qosPolicyId;
        this.region = $.region;
        this.segments = $.segments;
        this.shared = $.shared;
        this.tags = $.tags;
        this.tenantId = $.tenantId;
        this.transparentVlan = $.transparentVlan;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkArgs $;

        public Builder() {
            $ = new NetworkArgs();
        }

        public Builder(NetworkArgs defaults) {
            $ = new NetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp The administrative state of the network.
         * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
         * state of the existing network.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp The administrative state of the network.
         * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
         * state of the existing network.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param availabilityZoneHints An availability zone is used to make
         * network resources highly available. Used for resources with high availability
         * so that they are scheduled on different availability zones. Changing this
         * creates a new network.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZoneHints(@Nullable Output<List<String>> availabilityZoneHints) {
            $.availabilityZoneHints = availabilityZoneHints;
            return this;
        }

        /**
         * @param availabilityZoneHints An availability zone is used to make
         * network resources highly available. Used for resources with high availability
         * so that they are scheduled on different availability zones. Changing this
         * creates a new network.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZoneHints(List<String> availabilityZoneHints) {
            return availabilityZoneHints(Output.of(availabilityZoneHints));
        }

        /**
         * @param availabilityZoneHints An availability zone is used to make
         * network resources highly available. Used for resources with high availability
         * so that they are scheduled on different availability zones. Changing this
         * creates a new network.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZoneHints(String... availabilityZoneHints) {
            return availabilityZoneHints(List.of(availabilityZoneHints));
        }

        /**
         * @param description Human-readable description of the network. Changing this
         * updates the name of the existing network.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description of the network. Changing this
         * updates the name of the existing network.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dnsDomain The network DNS domain. Available, when Neutron DNS
         * extension is enabled. The `dns_domain` of a network in conjunction with the
         * `dns_name` attribute of its ports will be published in an external DNS
         * service when Neutron is configured to integrate with such a service.
         * 
         * @return builder
         * 
         */
        public Builder dnsDomain(@Nullable Output<String> dnsDomain) {
            $.dnsDomain = dnsDomain;
            return this;
        }

        /**
         * @param dnsDomain The network DNS domain. Available, when Neutron DNS
         * extension is enabled. The `dns_domain` of a network in conjunction with the
         * `dns_name` attribute of its ports will be published in an external DNS
         * service when Neutron is configured to integrate with such a service.
         * 
         * @return builder
         * 
         */
        public Builder dnsDomain(String dnsDomain) {
            return dnsDomain(Output.of(dnsDomain));
        }

        /**
         * @param external Specifies whether the network resource has the
         * external routing facility. Valid values are true and false. Defaults to
         * false. Changing this updates the external attribute of the existing network.
         * 
         * @return builder
         * 
         */
        public Builder external(@Nullable Output<Boolean> external) {
            $.external = external;
            return this;
        }

        /**
         * @param external Specifies whether the network resource has the
         * external routing facility. Valid values are true and false. Defaults to
         * false. Changing this updates the external attribute of the existing network.
         * 
         * @return builder
         * 
         */
        public Builder external(Boolean external) {
            return external(Output.of(external));
        }

        /**
         * @param mtu The network MTU. Available for read-only, when Neutron
         * `net-mtu` extension is enabled. Available for the modification, when
         * Neutron `net-mtu-writable` extension is enabled.
         * 
         * @return builder
         * 
         */
        public Builder mtu(@Nullable Output<Integer> mtu) {
            $.mtu = mtu;
            return this;
        }

        /**
         * @param mtu The network MTU. Available for read-only, when Neutron
         * `net-mtu` extension is enabled. Available for the modification, when
         * Neutron `net-mtu-writable` extension is enabled.
         * 
         * @return builder
         * 
         */
        public Builder mtu(Integer mtu) {
            return mtu(Output.of(mtu));
        }

        /**
         * @param name The name of the network. Changing this updates the name of
         * the existing network.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the network. Changing this updates the name of
         * the existing network.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param portSecurityEnabled Whether to explicitly enable or disable
         * port security on the network. Port Security is usually enabled by default, so
         * omitting this argument will usually result in a value of &#34;true&#34;. Setting this
         * explicitly to `false` will disable port security. Valid values are `true` and
         * `false`.
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
         * port security on the network. Port Security is usually enabled by default, so
         * omitting this argument will usually result in a value of &#34;true&#34;. Setting this
         * explicitly to `false` will disable port security. Valid values are `true` and
         * `false`.
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
         * A Networking client is needed to create a Neutron network. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * network.
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
         * A Networking client is needed to create a Neutron network. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * network.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param segments An array of one or more provider segment objects.
         * Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
         * updating any provider related segments attributes. Check your plug-in whether
         * it supports updating.
         * 
         * @return builder
         * 
         */
        public Builder segments(@Nullable Output<List<NetworkSegmentArgs>> segments) {
            $.segments = segments;
            return this;
        }

        /**
         * @param segments An array of one or more provider segment objects.
         * Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
         * updating any provider related segments attributes. Check your plug-in whether
         * it supports updating.
         * 
         * @return builder
         * 
         */
        public Builder segments(List<NetworkSegmentArgs> segments) {
            return segments(Output.of(segments));
        }

        /**
         * @param segments An array of one or more provider segment objects.
         * Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
         * updating any provider related segments attributes. Check your plug-in whether
         * it supports updating.
         * 
         * @return builder
         * 
         */
        public Builder segments(NetworkSegmentArgs... segments) {
            return segments(List.of(segments));
        }

        /**
         * @param shared Specifies whether the network resource can be accessed
         * by any tenant or not. Changing this updates the sharing capabilities of the
         * existing network.
         * 
         * @return builder
         * 
         */
        public Builder shared(@Nullable Output<Boolean> shared) {
            $.shared = shared;
            return this;
        }

        /**
         * @param shared Specifies whether the network resource can be accessed
         * by any tenant or not. Changing this updates the sharing capabilities of the
         * existing network.
         * 
         * @return builder
         * 
         */
        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        /**
         * @param tags A set of string tags for the network.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of string tags for the network.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of string tags for the network.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param tenantId The owner of the network. Required if admin wants to
         * create a network for another tenant. Changing this creates a new network.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the network. Required if admin wants to
         * create a network for another tenant. Changing this creates a new network.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param transparentVlan Specifies whether the network resource has the
         * VLAN transparent attribute set. Valid values are true and false. Defaults to
         * false. Changing this updates the `transparent_vlan` attribute of the existing
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
         * @param transparentVlan Specifies whether the network resource has the
         * VLAN transparent attribute set. Valid values are true and false. Defaults to
         * false. Changing this updates the `transparent_vlan` attribute of the existing
         * network.
         * 
         * @return builder
         * 
         */
        public Builder transparentVlan(Boolean transparentVlan) {
            return transparentVlan(Output.of(transparentVlan));
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

        public NetworkArgs build() {
            return $;
        }
    }

}
