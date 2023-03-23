// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.compute.inputs.ServerGroupRulesArgs;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerGroupArgs Empty = new ServerGroupArgs();

    /**
     * A unique name for the server group. Changing this creates
     * a new server group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name for the server group. Changing this creates
     * a new server group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A list of exactly one policy name to associate with
     * the server group. See the Policies section for more information. Changing this
     * creates a new server group.
     * 
     */
    @Import(name="policies")
    private @Nullable Output<List<String>> policies;

    /**
     * @return A list of exactly one policy name to associate with
     * the server group. See the Policies section for more information. Changing this
     * creates a new server group.
     * 
     */
    public Optional<Output<List<String>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The rules which are applied to specified `policy`. Currently,
     * only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<ServerGroupRulesArgs> rules;

    /**
     * @return The rules which are applied to specified `policy`. Currently,
     * only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
     * 
     */
    public Optional<Output<ServerGroupRulesArgs>> rules() {
        return Optional.ofNullable(this.rules);
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

    private ServerGroupArgs() {}

    private ServerGroupArgs(ServerGroupArgs $) {
        this.name = $.name;
        this.policies = $.policies;
        this.region = $.region;
        this.rules = $.rules;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerGroupArgs $;

        public Builder() {
            $ = new ServerGroupArgs();
        }

        public Builder(ServerGroupArgs defaults) {
            $ = new ServerGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name A unique name for the server group. Changing this creates
         * a new server group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name for the server group. Changing this creates
         * a new server group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param policies A list of exactly one policy name to associate with
         * the server group. See the Policies section for more information. Changing this
         * creates a new server group.
         * 
         * @return builder
         * 
         */
        public Builder policies(@Nullable Output<List<String>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies A list of exactly one policy name to associate with
         * the server group. See the Policies section for more information. Changing this
         * creates a new server group.
         * 
         * @return builder
         * 
         */
        public Builder policies(List<String> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies A list of exactly one policy name to associate with
         * the server group. See the Policies section for more information. Changing this
         * creates a new server group.
         * 
         * @return builder
         * 
         */
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * If omitted, the `region` argument of the provider is used. Changing
         * this creates a new server group.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * If omitted, the `region` argument of the provider is used. Changing
         * this creates a new server group.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param rules The rules which are applied to specified `policy`. Currently,
         * only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<ServerGroupRulesArgs> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules The rules which are applied to specified `policy`. Currently,
         * only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
         * 
         * @return builder
         * 
         */
        public Builder rules(ServerGroupRulesArgs rules) {
            return rules(Output.of(rules));
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

        public ServerGroupArgs build() {
            return $;
        }
    }

}
