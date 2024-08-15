// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FirewallArgs extends com.pulumi.resources.ResourceArgs {

    public static final FirewallArgs Empty = new FirewallArgs();

    /**
     * Administrative up/down status for the firewall
     * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
     * Changing this updates the `admin_state_up` of an existing firewall.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the firewall
     * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
     * Changing this updates the `admin_state_up` of an existing firewall.
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     * 
     */
    @Import(name="associatedRouters")
    private @Nullable Output<List<String>> associatedRouters;

    /**
     * @return Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     * 
     */
    public Optional<Output<List<String>>> associatedRouters() {
        return Optional.ofNullable(this.associatedRouters);
    }

    /**
     * A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Should this firewall not be associated with any routers
     * (must be &#34;true&#34; or &#34;false&#34; if provide - defaults to &#34;false&#34;).
     * Conflicts with `associated_routers`.
     * 
     */
    @Import(name="noRouters")
    private @Nullable Output<Boolean> noRouters;

    /**
     * @return Should this firewall not be associated with any routers
     * (must be &#34;true&#34; or &#34;false&#34; if provide - defaults to &#34;false&#34;).
     * Conflicts with `associated_routers`.
     * 
     */
    public Optional<Output<Boolean>> noRouters() {
        return Optional.ofNullable(this.noRouters);
    }

    /**
     * The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     * 
     */
    @Import(name="policyId", required=true)
    private Output<String> policyId;

    /**
     * @return The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     * 
     */
    public Output<String> policyId() {
        return this.policyId;
    }

    /**
     * The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
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

    private FirewallArgs() {}

    private FirewallArgs(FirewallArgs $) {
        this.adminStateUp = $.adminStateUp;
        this.associatedRouters = $.associatedRouters;
        this.description = $.description;
        this.name = $.name;
        this.noRouters = $.noRouters;
        this.policyId = $.policyId;
        this.region = $.region;
        this.tenantId = $.tenantId;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FirewallArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FirewallArgs $;

        public Builder() {
            $ = new FirewallArgs();
        }

        public Builder(FirewallArgs defaults) {
            $ = new FirewallArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp Administrative up/down status for the firewall
         * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
         * Changing this updates the `admin_state_up` of an existing firewall.
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
         * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
         * Changing this updates the `admin_state_up` of an existing firewall.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param associatedRouters Router(s) to associate this firewall instance
         * with. Must be a list of strings. Changing this updates the associated routers
         * of an existing firewall. Conflicts with `no_routers`.
         * 
         * @return builder
         * 
         */
        public Builder associatedRouters(@Nullable Output<List<String>> associatedRouters) {
            $.associatedRouters = associatedRouters;
            return this;
        }

        /**
         * @param associatedRouters Router(s) to associate this firewall instance
         * with. Must be a list of strings. Changing this updates the associated routers
         * of an existing firewall. Conflicts with `no_routers`.
         * 
         * @return builder
         * 
         */
        public Builder associatedRouters(List<String> associatedRouters) {
            return associatedRouters(Output.of(associatedRouters));
        }

        /**
         * @param associatedRouters Router(s) to associate this firewall instance
         * with. Must be a list of strings. Changing this updates the associated routers
         * of an existing firewall. Conflicts with `no_routers`.
         * 
         * @return builder
         * 
         */
        public Builder associatedRouters(String... associatedRouters) {
            return associatedRouters(List.of(associatedRouters));
        }

        /**
         * @param description A description for the firewall. Changing this
         * updates the `description` of an existing firewall.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the firewall. Changing this
         * updates the `description` of an existing firewall.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name A name for the firewall. Changing this
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
         * @param name A name for the firewall. Changing this
         * updates the `name` of an existing firewall.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param noRouters Should this firewall not be associated with any routers
         * (must be &#34;true&#34; or &#34;false&#34; if provide - defaults to &#34;false&#34;).
         * Conflicts with `associated_routers`.
         * 
         * @return builder
         * 
         */
        public Builder noRouters(@Nullable Output<Boolean> noRouters) {
            $.noRouters = noRouters;
            return this;
        }

        /**
         * @param noRouters Should this firewall not be associated with any routers
         * (must be &#34;true&#34; or &#34;false&#34; if provide - defaults to &#34;false&#34;).
         * Conflicts with `associated_routers`.
         * 
         * @return builder
         * 
         */
        public Builder noRouters(Boolean noRouters) {
            return noRouters(Output.of(noRouters));
        }

        /**
         * @param policyId The policy resource id for the firewall. Changing
         * this updates the `policy_id` of an existing firewall.
         * 
         * @return builder
         * 
         */
        public Builder policyId(Output<String> policyId) {
            $.policyId = policyId;
            return this;
        }

        /**
         * @param policyId The policy resource id for the firewall. Changing
         * this updates the `policy_id` of an existing firewall.
         * 
         * @return builder
         * 
         */
        public Builder policyId(String policyId) {
            return policyId(Output.of(policyId));
        }

        /**
         * @param region The region in which to obtain the v1 networking client.
         * A networking client is needed to create a firewall. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * firewall.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the v1 networking client.
         * A networking client is needed to create a firewall. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * firewall.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tenantId The owner of the floating IP. Required if admin wants
         * to create a firewall for another tenant. Changing this creates a new
         * firewall.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the floating IP. Required if admin wants
         * to create a firewall for another tenant. Changing this creates a new
         * firewall.
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

        public FirewallArgs build() {
            if ($.policyId == null) {
                throw new MissingRequiredPropertyException("FirewallArgs", "policyId");
            }
            return $;
        }
    }

}
