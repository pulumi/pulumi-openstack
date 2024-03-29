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


public final class PolicyV2Args extends com.pulumi.resources.ResourceArgs {

    public static final PolicyV2Args Empty = new PolicyV2Args();

    /**
     * Audit status of the firewall policy
     * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;false&#34;).
     * This status is set to &#34;false&#34; whenever the firewall policy or any of its
     * rules are changed. Changing this updates the `audited` status of an existing
     * firewall policy.
     * 
     */
    @Import(name="audited")
    private @Nullable Output<Boolean> audited;

    /**
     * @return Audit status of the firewall policy
     * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;false&#34;).
     * This status is set to &#34;false&#34; whenever the firewall policy or any of its
     * rules are changed. Changing this updates the `audited` status of an existing
     * firewall policy.
     * 
     */
    public Optional<Output<Boolean>> audited() {
        return Optional.ofNullable(this.audited);
    }

    /**
     * A description for the firewall policy. Changing
     * this updates the `description` of an existing firewall policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the firewall policy. Changing
     * this updates the `description` of an existing firewall policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A name for the firewall policy. Changing this
     * updates the `name` of an existing firewall policy.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name for the firewall policy. Changing this
     * updates the `name` of an existing firewall policy.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * This argument conflicts and is interchangeable
     * with `tenant_id`. The owner of the firewall policy. Required if admin wants
     * to create a firewall policy for another project. Changing this creates a new
     * firewall policy.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return This argument conflicts and is interchangeable
     * with `tenant_id`. The owner of the firewall policy. Required if admin wants
     * to create a firewall policy for another project. Changing this creates a new
     * firewall policy.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall policy.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall policy.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * An array of one or more firewall rules that comprise
     * the policy. Changing this results in adding/removing rules from the
     * existing firewall policy.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<String>> rules;

    /**
     * @return An array of one or more firewall rules that comprise
     * the policy. Changing this results in adding/removing rules from the
     * existing firewall policy.
     * 
     */
    public Optional<Output<List<String>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    /**
     * Sharing status of the firewall policy (must be &#34;true&#34;
     * or &#34;false&#34; if provided). If this is &#34;true&#34; the policy is visible to, and
     * can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall policy. Only administrative users
     * can specify if the policy should be shared.
     * 
     */
    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    /**
     * @return Sharing status of the firewall policy (must be &#34;true&#34;
     * or &#34;false&#34; if provided). If this is &#34;true&#34; the policy is visible to, and
     * can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall policy. Only administrative users
     * can specify if the policy should be shared.
     * 
     */
    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    /**
     * This argument conflicts and is interchangeable
     * with `project_id`. The owner of the firewall policy. Required if admin wants
     * to create a firewall policy for another tenant. Changing this creates a new
     * firewall policy.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return This argument conflicts and is interchangeable
     * with `project_id`. The owner of the firewall policy. Required if admin wants
     * to create a firewall policy for another tenant. Changing this creates a new
     * firewall policy.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private PolicyV2Args() {}

    private PolicyV2Args(PolicyV2Args $) {
        this.audited = $.audited;
        this.description = $.description;
        this.name = $.name;
        this.projectId = $.projectId;
        this.region = $.region;
        this.rules = $.rules;
        this.shared = $.shared;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicyV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicyV2Args $;

        public Builder() {
            $ = new PolicyV2Args();
        }

        public Builder(PolicyV2Args defaults) {
            $ = new PolicyV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param audited Audit status of the firewall policy
         * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;false&#34;).
         * This status is set to &#34;false&#34; whenever the firewall policy or any of its
         * rules are changed. Changing this updates the `audited` status of an existing
         * firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder audited(@Nullable Output<Boolean> audited) {
            $.audited = audited;
            return this;
        }

        /**
         * @param audited Audit status of the firewall policy
         * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;false&#34;).
         * This status is set to &#34;false&#34; whenever the firewall policy or any of its
         * rules are changed. Changing this updates the `audited` status of an existing
         * firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder audited(Boolean audited) {
            return audited(Output.of(audited));
        }

        /**
         * @param description A description for the firewall policy. Changing
         * this updates the `description` of an existing firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the firewall policy. Changing
         * this updates the `description` of an existing firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name A name for the firewall policy. Changing this
         * updates the `name` of an existing firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name for the firewall policy. Changing this
         * updates the `name` of an existing firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId This argument conflicts and is interchangeable
         * with `tenant_id`. The owner of the firewall policy. Required if admin wants
         * to create a firewall policy for another project. Changing this creates a new
         * firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId This argument conflicts and is interchangeable
         * with `tenant_id`. The owner of the firewall policy. Required if admin wants
         * to create a firewall policy for another project. Changing this creates a new
         * firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the v2 networking client.
         * A networking client is needed to create a firewall policy. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * firewall policy.
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
         * A networking client is needed to create a firewall policy. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param rules An array of one or more firewall rules that comprise
         * the policy. Changing this results in adding/removing rules from the
         * existing firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<String>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules An array of one or more firewall rules that comprise
         * the policy. Changing this results in adding/removing rules from the
         * existing firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<String> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules An array of one or more firewall rules that comprise
         * the policy. Changing this results in adding/removing rules from the
         * existing firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder rules(String... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param shared Sharing status of the firewall policy (must be &#34;true&#34;
         * or &#34;false&#34; if provided). If this is &#34;true&#34; the policy is visible to, and
         * can be used in, firewalls in other tenants. Changing this updates the
         * `shared` status of an existing firewall policy. Only administrative users
         * can specify if the policy should be shared.
         * 
         * @return builder
         * 
         */
        public Builder shared(@Nullable Output<Boolean> shared) {
            $.shared = shared;
            return this;
        }

        /**
         * @param shared Sharing status of the firewall policy (must be &#34;true&#34;
         * or &#34;false&#34; if provided). If this is &#34;true&#34; the policy is visible to, and
         * can be used in, firewalls in other tenants. Changing this updates the
         * `shared` status of an existing firewall policy. Only administrative users
         * can specify if the policy should be shared.
         * 
         * @return builder
         * 
         */
        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        /**
         * @param tenantId This argument conflicts and is interchangeable
         * with `project_id`. The owner of the firewall policy. Required if admin wants
         * to create a firewall policy for another tenant. Changing this creates a new
         * firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId This argument conflicts and is interchangeable
         * with `project_id`. The owner of the firewall policy. Required if admin wants
         * to create a firewall policy for another tenant. Changing this creates a new
         * firewall policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public PolicyV2Args build() {
            return $;
        }
    }

}
