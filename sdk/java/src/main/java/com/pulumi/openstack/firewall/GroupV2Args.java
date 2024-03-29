// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupV2Args extends com.pulumi.resources.ResourceArgs {

    public static final GroupV2Args Empty = new GroupV2Args();

    /**
     * Administrative up/down status for the firewall
     * group (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
     * Changing this updates the `admin_state_up` of an existing firewall group.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the firewall
     * group (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
     * Changing this updates the `admin_state_up` of an existing firewall group.
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * A description for the firewall group. Changing this
     * updates the `description` of an existing firewall group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the firewall group. Changing this
     * updates the `description` of an existing firewall group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The egress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `egress_firewall_policy_id` of an existing firewall group.
     * 
     */
    @Import(name="egressFirewallPolicyId")
    private @Nullable Output<String> egressFirewallPolicyId;

    /**
     * @return The egress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `egress_firewall_policy_id` of an existing firewall group.
     * 
     */
    public Optional<Output<String>> egressFirewallPolicyId() {
        return Optional.ofNullable(this.egressFirewallPolicyId);
    }

    /**
     * The ingress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `ingress_firewall_policy_id` of an existing firewall group.
     * 
     */
    @Import(name="ingressFirewallPolicyId")
    private @Nullable Output<String> ingressFirewallPolicyId;

    /**
     * @return The ingress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `ingress_firewall_policy_id` of an existing firewall group.
     * 
     */
    public Optional<Output<String>> ingressFirewallPolicyId() {
        return Optional.ofNullable(this.ingressFirewallPolicyId);
    }

    /**
     * A name for the firewall group. Changing this
     * updates the `name` of an existing firewall.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name for the firewall group. Changing this
     * updates the `name` of an existing firewall.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Port(s) to associate this firewall group
     * with. Must be a list of strings. Changing this updates the associated ports
     * of an existing firewall group.
     * 
     */
    @Import(name="ports")
    private @Nullable Output<List<String>> ports;

    /**
     * @return Port(s) to associate this firewall group
     * with. Must be a list of strings. Changing this updates the associated ports
     * of an existing firewall group.
     * 
     */
    public Optional<Output<List<String>>> ports() {
        return Optional.ofNullable(this.ports);
    }

    /**
     * This argument conflicts and  is interchangeable
     * with `tenant_id`. The owner of the firewall group. Required if admin wants
     * to create a firewall group for another project. Changing this creates a new
     * firewall group.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return This argument conflicts and  is interchangeable
     * with `tenant_id`. The owner of the firewall group. Required if admin wants
     * to create a firewall group for another project. Changing this creates a new
     * firewall group.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall group.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall group.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Sharing status of the firewall group (must be &#34;true&#34;
     * or &#34;false&#34; if provided). If this is &#34;true&#34; the firewall group is visible to,
     * and can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall group. Only administrative users
     * can specify if the firewall group should be shared.
     * 
     */
    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    /**
     * @return Sharing status of the firewall group (must be &#34;true&#34;
     * or &#34;false&#34; if provided). If this is &#34;true&#34; the firewall group is visible to,
     * and can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall group. Only administrative users
     * can specify if the firewall group should be shared.
     * 
     */
    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    /**
     * This argument conflicts and is interchangeable with
     * `project_id`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another tenant. Changing this creates a new
     * firewall group.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return This argument conflicts and is interchangeable with
     * `project_id`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another tenant. Changing this creates a new
     * firewall group.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private GroupV2Args() {}

    private GroupV2Args(GroupV2Args $) {
        this.adminStateUp = $.adminStateUp;
        this.description = $.description;
        this.egressFirewallPolicyId = $.egressFirewallPolicyId;
        this.ingressFirewallPolicyId = $.ingressFirewallPolicyId;
        this.name = $.name;
        this.ports = $.ports;
        this.projectId = $.projectId;
        this.region = $.region;
        this.shared = $.shared;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupV2Args $;

        public Builder() {
            $ = new GroupV2Args();
        }

        public Builder(GroupV2Args defaults) {
            $ = new GroupV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp Administrative up/down status for the firewall
         * group (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
         * Changing this updates the `admin_state_up` of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp Administrative up/down status for the firewall
         * group (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
         * Changing this updates the `admin_state_up` of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param description A description for the firewall group. Changing this
         * updates the `description` of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the firewall group. Changing this
         * updates the `description` of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param egressFirewallPolicyId The egress firewall policy resource
         * id for the firewall group. Changing this updates the
         * `egress_firewall_policy_id` of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder egressFirewallPolicyId(@Nullable Output<String> egressFirewallPolicyId) {
            $.egressFirewallPolicyId = egressFirewallPolicyId;
            return this;
        }

        /**
         * @param egressFirewallPolicyId The egress firewall policy resource
         * id for the firewall group. Changing this updates the
         * `egress_firewall_policy_id` of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder egressFirewallPolicyId(String egressFirewallPolicyId) {
            return egressFirewallPolicyId(Output.of(egressFirewallPolicyId));
        }

        /**
         * @param ingressFirewallPolicyId The ingress firewall policy resource
         * id for the firewall group. Changing this updates the
         * `ingress_firewall_policy_id` of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder ingressFirewallPolicyId(@Nullable Output<String> ingressFirewallPolicyId) {
            $.ingressFirewallPolicyId = ingressFirewallPolicyId;
            return this;
        }

        /**
         * @param ingressFirewallPolicyId The ingress firewall policy resource
         * id for the firewall group. Changing this updates the
         * `ingress_firewall_policy_id` of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder ingressFirewallPolicyId(String ingressFirewallPolicyId) {
            return ingressFirewallPolicyId(Output.of(ingressFirewallPolicyId));
        }

        /**
         * @param name A name for the firewall group. Changing this
         * updates the `name` of an existing firewall.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name for the firewall group. Changing this
         * updates the `name` of an existing firewall.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param ports Port(s) to associate this firewall group
         * with. Must be a list of strings. Changing this updates the associated ports
         * of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder ports(@Nullable Output<List<String>> ports) {
            $.ports = ports;
            return this;
        }

        /**
         * @param ports Port(s) to associate this firewall group
         * with. Must be a list of strings. Changing this updates the associated ports
         * of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder ports(List<String> ports) {
            return ports(Output.of(ports));
        }

        /**
         * @param ports Port(s) to associate this firewall group
         * with. Must be a list of strings. Changing this updates the associated ports
         * of an existing firewall group.
         * 
         * @return builder
         * 
         */
        public Builder ports(String... ports) {
            return ports(List.of(ports));
        }

        /**
         * @param projectId This argument conflicts and  is interchangeable
         * with `tenant_id`. The owner of the firewall group. Required if admin wants
         * to create a firewall group for another project. Changing this creates a new
         * firewall group.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId This argument conflicts and  is interchangeable
         * with `tenant_id`. The owner of the firewall group. Required if admin wants
         * to create a firewall group for another project. Changing this creates a new
         * firewall group.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the v2 networking client.
         * A networking client is needed to create a firewall group. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * firewall group.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the v2 networking client.
         * A networking client is needed to create a firewall group. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * firewall group.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param shared Sharing status of the firewall group (must be &#34;true&#34;
         * or &#34;false&#34; if provided). If this is &#34;true&#34; the firewall group is visible to,
         * and can be used in, firewalls in other tenants. Changing this updates the
         * `shared` status of an existing firewall group. Only administrative users
         * can specify if the firewall group should be shared.
         * 
         * @return builder
         * 
         */
        public Builder shared(@Nullable Output<Boolean> shared) {
            $.shared = shared;
            return this;
        }

        /**
         * @param shared Sharing status of the firewall group (must be &#34;true&#34;
         * or &#34;false&#34; if provided). If this is &#34;true&#34; the firewall group is visible to,
         * and can be used in, firewalls in other tenants. Changing this updates the
         * `shared` status of an existing firewall group. Only administrative users
         * can specify if the firewall group should be shared.
         * 
         * @return builder
         * 
         */
        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        /**
         * @param tenantId This argument conflicts and is interchangeable with
         * `project_id`. The owner of the firewall group. Required if admin wants to
         * create a firewall group for another tenant. Changing this creates a new
         * firewall group.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId This argument conflicts and is interchangeable with
         * `project_id`. The owner of the firewall group. Required if admin wants to
         * create a firewall group for another tenant. Changing this creates a new
         * firewall group.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public GroupV2Args build() {
            return $;
        }
    }

}
