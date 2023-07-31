// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFwRuleV2Args extends com.pulumi.resources.InvokeArgs {

    public static final GetFwRuleV2Args Empty = new GetFwRuleV2Args();

    /**
     * Action to be taken when the firewall rule matches.
     * 
     */
    @Import(name="action")
    private @Nullable Output<String> action;

    /**
     * @return Action to be taken when the firewall rule matches.
     * 
     */
    public Optional<Output<String>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * The description of the firewall rule.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the firewall rule.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The destination IP address on which the
     * firewall rule operates.
     * 
     */
    @Import(name="destinationIpAddress")
    private @Nullable Output<String> destinationIpAddress;

    /**
     * @return The destination IP address on which the
     * firewall rule operates.
     * 
     */
    public Optional<Output<String>> destinationIpAddress() {
        return Optional.ofNullable(this.destinationIpAddress);
    }

    /**
     * The destination port on which the firewall
     * rule operates.
     * 
     */
    @Import(name="destinationPort")
    private @Nullable Output<String> destinationPort;

    /**
     * @return The destination port on which the firewall
     * rule operates.
     * 
     */
    public Optional<Output<String>> destinationPort() {
        return Optional.ofNullable(this.destinationPort);
    }

    /**
     * Enabled status for the firewall rule.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Enabled status for the firewall rule.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The ID of the firewall policy the rule belongs to.
     * 
     */
    @Import(name="firewallPolicyIds")
    private @Nullable Output<List<String>> firewallPolicyIds;

    /**
     * @return The ID of the firewall policy the rule belongs to.
     * 
     */
    public Optional<Output<List<String>>> firewallPolicyIds() {
        return Optional.ofNullable(this.firewallPolicyIds);
    }

    /**
     * IP version, either 4 (default) or 6.
     * 
     */
    @Import(name="ipVersion")
    private @Nullable Output<Integer> ipVersion;

    /**
     * @return IP version, either 4 (default) or 6.
     * 
     */
    public Optional<Output<Integer>> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }

    /**
     * The name of the firewall rule.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the firewall rule.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The protocol type on which the firewall rule operates.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return The protocol type on which the firewall rule operates.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve firewall policy ids. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve firewall policy ids. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The ID of the firewall rule.
     * 
     */
    @Import(name="ruleId")
    private @Nullable Output<String> ruleId;

    /**
     * @return The ID of the firewall rule.
     * 
     */
    public Optional<Output<String>> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }

    /**
     * The sharing status of the firewall policy.
     * 
     */
    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    /**
     * @return The sharing status of the firewall policy.
     * 
     */
    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    /**
     * The source IP address on which the firewall
     * rule operates.
     * 
     */
    @Import(name="sourceIpAddress")
    private @Nullable Output<String> sourceIpAddress;

    /**
     * @return The source IP address on which the firewall
     * rule operates.
     * 
     */
    public Optional<Output<String>> sourceIpAddress() {
        return Optional.ofNullable(this.sourceIpAddress);
    }

    /**
     * The source port on which the firewall
     * rule operates.
     * 
     */
    @Import(name="sourcePort")
    private @Nullable Output<String> sourcePort;

    /**
     * @return The source port on which the firewall
     * rule operates.
     * 
     */
    public Optional<Output<String>> sourcePort() {
        return Optional.ofNullable(this.sourcePort);
    }

    /**
     * The owner of the firewall policy.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the firewall policy.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private GetFwRuleV2Args() {}

    private GetFwRuleV2Args(GetFwRuleV2Args $) {
        this.action = $.action;
        this.description = $.description;
        this.destinationIpAddress = $.destinationIpAddress;
        this.destinationPort = $.destinationPort;
        this.enabled = $.enabled;
        this.firewallPolicyIds = $.firewallPolicyIds;
        this.ipVersion = $.ipVersion;
        this.name = $.name;
        this.protocol = $.protocol;
        this.region = $.region;
        this.ruleId = $.ruleId;
        this.shared = $.shared;
        this.sourceIpAddress = $.sourceIpAddress;
        this.sourcePort = $.sourcePort;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFwRuleV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFwRuleV2Args $;

        public Builder() {
            $ = new GetFwRuleV2Args();
        }

        public Builder(GetFwRuleV2Args defaults) {
            $ = new GetFwRuleV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param action Action to be taken when the firewall rule matches.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action Action to be taken when the firewall rule matches.
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param description The description of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param destinationIpAddress The destination IP address on which the
         * firewall rule operates.
         * 
         * @return builder
         * 
         */
        public Builder destinationIpAddress(@Nullable Output<String> destinationIpAddress) {
            $.destinationIpAddress = destinationIpAddress;
            return this;
        }

        /**
         * @param destinationIpAddress The destination IP address on which the
         * firewall rule operates.
         * 
         * @return builder
         * 
         */
        public Builder destinationIpAddress(String destinationIpAddress) {
            return destinationIpAddress(Output.of(destinationIpAddress));
        }

        /**
         * @param destinationPort The destination port on which the firewall
         * rule operates.
         * 
         * @return builder
         * 
         */
        public Builder destinationPort(@Nullable Output<String> destinationPort) {
            $.destinationPort = destinationPort;
            return this;
        }

        /**
         * @param destinationPort The destination port on which the firewall
         * rule operates.
         * 
         * @return builder
         * 
         */
        public Builder destinationPort(String destinationPort) {
            return destinationPort(Output.of(destinationPort));
        }

        /**
         * @param enabled Enabled status for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Enabled status for the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param firewallPolicyIds The ID of the firewall policy the rule belongs to.
         * 
         * @return builder
         * 
         */
        public Builder firewallPolicyIds(@Nullable Output<List<String>> firewallPolicyIds) {
            $.firewallPolicyIds = firewallPolicyIds;
            return this;
        }

        /**
         * @param firewallPolicyIds The ID of the firewall policy the rule belongs to.
         * 
         * @return builder
         * 
         */
        public Builder firewallPolicyIds(List<String> firewallPolicyIds) {
            return firewallPolicyIds(Output.of(firewallPolicyIds));
        }

        /**
         * @param firewallPolicyIds The ID of the firewall policy the rule belongs to.
         * 
         * @return builder
         * 
         */
        public Builder firewallPolicyIds(String... firewallPolicyIds) {
            return firewallPolicyIds(List.of(firewallPolicyIds));
        }

        /**
         * @param ipVersion IP version, either 4 (default) or 6.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(@Nullable Output<Integer> ipVersion) {
            $.ipVersion = ipVersion;
            return this;
        }

        /**
         * @param ipVersion IP version, either 4 (default) or 6.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(Integer ipVersion) {
            return ipVersion(Output.of(ipVersion));
        }

        /**
         * @param name The name of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param protocol The protocol type on which the firewall rule operates.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol type on which the firewall rule operates.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param region The region in which to obtain the V2 Neutron client.
         * A Neutron client is needed to retrieve firewall policy ids. If omitted, the
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
         * A Neutron client is needed to retrieve firewall policy ids. If omitted, the
         * `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param ruleId The ID of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(@Nullable Output<String> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param ruleId The ID of the firewall rule.
         * 
         * @return builder
         * 
         */
        public Builder ruleId(String ruleId) {
            return ruleId(Output.of(ruleId));
        }

        /**
         * @param shared The sharing status of the firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder shared(@Nullable Output<Boolean> shared) {
            $.shared = shared;
            return this;
        }

        /**
         * @param shared The sharing status of the firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        /**
         * @param sourceIpAddress The source IP address on which the firewall
         * rule operates.
         * 
         * @return builder
         * 
         */
        public Builder sourceIpAddress(@Nullable Output<String> sourceIpAddress) {
            $.sourceIpAddress = sourceIpAddress;
            return this;
        }

        /**
         * @param sourceIpAddress The source IP address on which the firewall
         * rule operates.
         * 
         * @return builder
         * 
         */
        public Builder sourceIpAddress(String sourceIpAddress) {
            return sourceIpAddress(Output.of(sourceIpAddress));
        }

        /**
         * @param sourcePort The source port on which the firewall
         * rule operates.
         * 
         * @return builder
         * 
         */
        public Builder sourcePort(@Nullable Output<String> sourcePort) {
            $.sourcePort = sourcePort;
            return this;
        }

        /**
         * @param sourcePort The source port on which the firewall
         * rule operates.
         * 
         * @return builder
         * 
         */
        public Builder sourcePort(String sourcePort) {
            return sourcePort(Output.of(sourcePort));
        }

        /**
         * @param tenantId The owner of the firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public GetFwRuleV2Args build() {
            return $;
        }
    }

}