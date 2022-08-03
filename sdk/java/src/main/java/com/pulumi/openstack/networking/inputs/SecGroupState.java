// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecGroupState extends com.pulumi.resources.ResourceArgs {

    public static final SecGroupState Empty = new SecGroupState();

    /**
     * The collection of tags assigned on the security group, which have
     * been explicitly and implicitly added.
     * 
     */
    @Import(name="allTags")
    private @Nullable Output<List<String>> allTags;

    /**
     * @return The collection of tags assigned on the security group, which have
     * been explicitly and implicitly added.
     * 
     */
    public Optional<Output<List<String>>> allTags() {
        return Optional.ofNullable(this.allTags);
    }

    /**
     * Whether or not to delete the default
     * egress security rules. This is `false` by default. See the below note
     * for more information.
     * 
     */
    @Import(name="deleteDefaultRules")
    private @Nullable Output<Boolean> deleteDefaultRules;

    /**
     * @return Whether or not to delete the default
     * egress security rules. This is `false` by default. See the below note
     * for more information.
     * 
     */
    public Optional<Output<Boolean>> deleteDefaultRules() {
        return Optional.ofNullable(this.deleteDefaultRules);
    }

    /**
     * A unique name for the security group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A unique name for the security group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A unique name for the security group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name for the security group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * A set of string tags for the security group.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of string tags for the security group.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private SecGroupState() {}

    private SecGroupState(SecGroupState $) {
        this.allTags = $.allTags;
        this.deleteDefaultRules = $.deleteDefaultRules;
        this.description = $.description;
        this.name = $.name;
        this.region = $.region;
        this.tags = $.tags;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecGroupState $;

        public Builder() {
            $ = new SecGroupState();
        }

        public Builder(SecGroupState defaults) {
            $ = new SecGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param allTags The collection of tags assigned on the security group, which have
         * been explicitly and implicitly added.
         * 
         * @return builder
         * 
         */
        public Builder allTags(@Nullable Output<List<String>> allTags) {
            $.allTags = allTags;
            return this;
        }

        /**
         * @param allTags The collection of tags assigned on the security group, which have
         * been explicitly and implicitly added.
         * 
         * @return builder
         * 
         */
        public Builder allTags(List<String> allTags) {
            return allTags(Output.of(allTags));
        }

        /**
         * @param allTags The collection of tags assigned on the security group, which have
         * been explicitly and implicitly added.
         * 
         * @return builder
         * 
         */
        public Builder allTags(String... allTags) {
            return allTags(List.of(allTags));
        }

        /**
         * @param deleteDefaultRules Whether or not to delete the default
         * egress security rules. This is `false` by default. See the below note
         * for more information.
         * 
         * @return builder
         * 
         */
        public Builder deleteDefaultRules(@Nullable Output<Boolean> deleteDefaultRules) {
            $.deleteDefaultRules = deleteDefaultRules;
            return this;
        }

        /**
         * @param deleteDefaultRules Whether or not to delete the default
         * egress security rules. This is `false` by default. See the below note
         * for more information.
         * 
         * @return builder
         * 
         */
        public Builder deleteDefaultRules(Boolean deleteDefaultRules) {
            return deleteDefaultRules(Output.of(deleteDefaultRules));
        }

        /**
         * @param description A unique name for the security group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A unique name for the security group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name A unique name for the security group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name for the security group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V2 networking client.
         * A networking client is needed to create a port. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * security group.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 networking client.
         * A networking client is needed to create a port. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * security group.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tags A set of string tags for the security group.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of string tags for the security group.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of string tags for the security group.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param tenantId The owner of the security group. Required if admin
         * wants to create a port for another tenant. Changing this creates a new
         * security group.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the security group. Required if admin
         * wants to create a port for another tenant. Changing this creates a new
         * security group.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public SecGroupState build() {
            return $;
        }
    }

}
