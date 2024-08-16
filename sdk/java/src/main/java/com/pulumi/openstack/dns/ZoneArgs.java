// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.dns;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final ZoneArgs Empty = new ZoneArgs();

    /**
     * Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     * 
     */
    @Import(name="attributes")
    private @Nullable Output<Map<String,String>> attributes;

    /**
     * @return Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     * 
     */
    public Optional<Output<Map<String,String>>> attributes() {
        return Optional.ofNullable(this.attributes);
    }

    /**
     * A description of the zone.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the zone.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     * 
     */
    @Import(name="disableStatusCheck")
    private @Nullable Output<Boolean> disableStatusCheck;

    /**
     * @return Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     * 
     */
    public Optional<Output<Boolean>> disableStatusCheck() {
        return Optional.ofNullable(this.disableStatusCheck);
    }

    /**
     * The email contact for the zone record.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return The email contact for the zone record.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * An array of master DNS servers. For when `type` is
     * `SECONDARY`.
     * 
     */
    @Import(name="masters")
    private @Nullable Output<List<String>> masters;

    /**
     * @return An array of master DNS servers. For when `type` is
     * `SECONDARY`.
     * 
     */
    public Optional<Output<List<String>>> masters() {
        return Optional.ofNullable(this.masters);
    }

    /**
     * The name of the zone. Note the `.` at the end of the name.
     * Changing this creates a new DNS zone.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the zone. Note the `.` at the end of the name.
     * Changing this creates a new DNS zone.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project DNS zone is created
     * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
     * user role in target project)
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project DNS zone is created
     * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
     * user role in target project)
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The time to live (TTL) of the zone.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return The time to live (TTL) of the zone.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * The type of zone. Can either be `PRIMARY` or `SECONDARY`.
     * Changing this creates a new zone.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of zone. Can either be `PRIMARY` or `SECONDARY`.
     * Changing this creates a new zone.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Map of additional options. Changing this creates a
     * new zone.
     * 
     */
    @Import(name="valueSpecs")
    private @Nullable Output<Map<String,String>> valueSpecs;

    /**
     * @return Map of additional options. Changing this creates a
     * new zone.
     * 
     */
    public Optional<Output<Map<String,String>>> valueSpecs() {
        return Optional.ofNullable(this.valueSpecs);
    }

    private ZoneArgs() {}

    private ZoneArgs(ZoneArgs $) {
        this.attributes = $.attributes;
        this.description = $.description;
        this.disableStatusCheck = $.disableStatusCheck;
        this.email = $.email;
        this.masters = $.masters;
        this.name = $.name;
        this.projectId = $.projectId;
        this.region = $.region;
        this.ttl = $.ttl;
        this.type = $.type;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneArgs $;

        public Builder() {
            $ = new ZoneArgs();
        }

        public Builder(ZoneArgs defaults) {
            $ = new ZoneArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attributes Attributes for the DNS Service scheduler.
         * Changing this creates a new zone.
         * 
         * @return builder
         * 
         */
        public Builder attributes(@Nullable Output<Map<String,String>> attributes) {
            $.attributes = attributes;
            return this;
        }

        /**
         * @param attributes Attributes for the DNS Service scheduler.
         * Changing this creates a new zone.
         * 
         * @return builder
         * 
         */
        public Builder attributes(Map<String,String> attributes) {
            return attributes(Output.of(attributes));
        }

        /**
         * @param description A description of the zone.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the zone.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableStatusCheck Disable wait for zone to reach ACTIVE
         * status. The check is enabled by default. If this argument is true, zone
         * will be considered as created/updated if OpenStack request returned success.
         * 
         * @return builder
         * 
         */
        public Builder disableStatusCheck(@Nullable Output<Boolean> disableStatusCheck) {
            $.disableStatusCheck = disableStatusCheck;
            return this;
        }

        /**
         * @param disableStatusCheck Disable wait for zone to reach ACTIVE
         * status. The check is enabled by default. If this argument is true, zone
         * will be considered as created/updated if OpenStack request returned success.
         * 
         * @return builder
         * 
         */
        public Builder disableStatusCheck(Boolean disableStatusCheck) {
            return disableStatusCheck(Output.of(disableStatusCheck));
        }

        /**
         * @param email The email contact for the zone record.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The email contact for the zone record.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param masters An array of master DNS servers. For when `type` is
         * `SECONDARY`.
         * 
         * @return builder
         * 
         */
        public Builder masters(@Nullable Output<List<String>> masters) {
            $.masters = masters;
            return this;
        }

        /**
         * @param masters An array of master DNS servers. For when `type` is
         * `SECONDARY`.
         * 
         * @return builder
         * 
         */
        public Builder masters(List<String> masters) {
            return masters(Output.of(masters));
        }

        /**
         * @param masters An array of master DNS servers. For when `type` is
         * `SECONDARY`.
         * 
         * @return builder
         * 
         */
        public Builder masters(String... masters) {
            return masters(List.of(masters));
        }

        /**
         * @param name The name of the zone. Note the `.` at the end of the name.
         * Changing this creates a new DNS zone.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the zone. Note the `.` at the end of the name.
         * Changing this creates a new DNS zone.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The ID of the project DNS zone is created
         * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
         * user role in target project)
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project DNS zone is created
         * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
         * user role in target project)
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * Keypairs are associated with accounts, but a Compute client is needed to
         * create one. If omitted, the `region` argument of the provider is used.
         * Changing this creates a new DNS zone.
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
         * Keypairs are associated with accounts, but a Compute client is needed to
         * create one. If omitted, the `region` argument of the provider is used.
         * Changing this creates a new DNS zone.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param ttl The time to live (TTL) of the zone.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The time to live (TTL) of the zone.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param type The type of zone. Can either be `PRIMARY` or `SECONDARY`.
         * Changing this creates a new zone.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of zone. Can either be `PRIMARY` or `SECONDARY`.
         * Changing this creates a new zone.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param valueSpecs Map of additional options. Changing this creates a
         * new zone.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(@Nullable Output<Map<String,String>> valueSpecs) {
            $.valueSpecs = valueSpecs;
            return this;
        }

        /**
         * @param valueSpecs Map of additional options. Changing this creates a
         * new zone.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(Map<String,String> valueSpecs) {
            return valueSpecs(Output.of(valueSpecs));
        }

        public ZoneArgs build() {
            return $;
        }
    }

}
