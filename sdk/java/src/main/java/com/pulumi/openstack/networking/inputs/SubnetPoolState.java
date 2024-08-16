// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

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


public final class SubnetPoolState extends com.pulumi.resources.ResourceArgs {

    public static final SubnetPoolState Empty = new SubnetPoolState();

    /**
     * The Neutron address scope to assign to the
     * subnetpool. Changing this updates the address scope id of the existing
     * subnetpool.
     * 
     */
    @Import(name="addressScopeId")
    private @Nullable Output<String> addressScopeId;

    /**
     * @return The Neutron address scope to assign to the
     * subnetpool. Changing this updates the address scope id of the existing
     * subnetpool.
     * 
     */
    public Optional<Output<String>> addressScopeId() {
        return Optional.ofNullable(this.addressScopeId);
    }

    /**
     * The collection of tags assigned on the subnetpool, which have been
     * explicitly and implicitly added.
     * 
     */
    @Import(name="allTags")
    private @Nullable Output<List<String>> allTags;

    /**
     * @return The collection of tags assigned on the subnetpool, which have been
     * explicitly and implicitly added.
     * 
     */
    public Optional<Output<List<String>>> allTags() {
        return Optional.ofNullable(this.allTags);
    }

    /**
     * The time at which subnetpool was created.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The time at which subnetpool was created.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * The size of the prefix to allocate when the cidr
     * or prefixlen attributes are omitted when you create the subnet. Defaults to the
     * MinPrefixLen. Changing this updates the default prefixlen of the existing
     * subnetpool.
     * 
     */
    @Import(name="defaultPrefixlen")
    private @Nullable Output<Integer> defaultPrefixlen;

    /**
     * @return The size of the prefix to allocate when the cidr
     * or prefixlen attributes are omitted when you create the subnet. Defaults to the
     * MinPrefixLen. Changing this updates the default prefixlen of the existing
     * subnetpool.
     * 
     */
    public Optional<Output<Integer>> defaultPrefixlen() {
        return Optional.ofNullable(this.defaultPrefixlen);
    }

    /**
     * The per-project quota on the prefix space that can be
     * allocated from the subnetpool for project subnets. Changing this updates the
     * default quota of the existing subnetpool.
     * 
     */
    @Import(name="defaultQuota")
    private @Nullable Output<Integer> defaultQuota;

    /**
     * @return The per-project quota on the prefix space that can be
     * allocated from the subnetpool for project subnets. Changing this updates the
     * default quota of the existing subnetpool.
     * 
     */
    public Optional<Output<Integer>> defaultQuota() {
        return Optional.ofNullable(this.defaultQuota);
    }

    /**
     * The human-readable description for the subnetpool.
     * Changing this updates the description of the existing subnetpool.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The human-readable description for the subnetpool.
     * Changing this updates the description of the existing subnetpool.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The IP protocol version.
     * 
     */
    @Import(name="ipVersion")
    private @Nullable Output<Integer> ipVersion;

    /**
     * @return The IP protocol version.
     * 
     */
    public Optional<Output<Integer>> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }

    /**
     * Indicates whether the subnetpool is default
     * subnetpool or not. Changing this updates the default status of the existing
     * subnetpool.
     * 
     */
    @Import(name="isDefault")
    private @Nullable Output<Boolean> isDefault;

    /**
     * @return Indicates whether the subnetpool is default
     * subnetpool or not. Changing this updates the default status of the existing
     * subnetpool.
     * 
     */
    public Optional<Output<Boolean>> isDefault() {
        return Optional.ofNullable(this.isDefault);
    }

    /**
     * The maximum prefix size that can be allocated from
     * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
     * default is 128. Changing this updates the max prefixlen of the existing
     * subnetpool.
     * 
     */
    @Import(name="maxPrefixlen")
    private @Nullable Output<Integer> maxPrefixlen;

    /**
     * @return The maximum prefix size that can be allocated from
     * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
     * default is 128. Changing this updates the max prefixlen of the existing
     * subnetpool.
     * 
     */
    public Optional<Output<Integer>> maxPrefixlen() {
        return Optional.ofNullable(this.maxPrefixlen);
    }

    /**
     * The smallest prefix that can be allocated from a
     * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
     * is 64. Changing this updates the min prefixlen of the existing subnetpool.
     * 
     */
    @Import(name="minPrefixlen")
    private @Nullable Output<Integer> minPrefixlen;

    /**
     * @return The smallest prefix that can be allocated from a
     * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
     * is 64. Changing this updates the min prefixlen of the existing subnetpool.
     * 
     */
    public Optional<Output<Integer>> minPrefixlen() {
        return Optional.ofNullable(this.minPrefixlen);
    }

    /**
     * The name of the subnetpool. Changing this updates the name of
     * the existing subnetpool.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the subnetpool. Changing this updates the name of
     * the existing subnetpool.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A list of subnet prefixes to assign to the subnetpool.
     * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
     * subnet prefix must be unique among all subnet prefixes in all subnetpools that
     * are associated with the address scope. Changing this updates the prefixes list
     * of the existing subnetpool.
     * 
     */
    @Import(name="prefixes")
    private @Nullable Output<List<String>> prefixes;

    /**
     * @return A list of subnet prefixes to assign to the subnetpool.
     * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
     * subnet prefix must be unique among all subnet prefixes in all subnetpools that
     * are associated with the address scope. Changing this updates the prefixes list
     * of the existing subnetpool.
     * 
     */
    public Optional<Output<List<String>>> prefixes() {
        return Optional.ofNullable(this.prefixes);
    }

    /**
     * The owner of the subnetpool. Required if admin wants to
     * create a subnetpool for another project. Changing this creates a new subnetpool.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The owner of the subnetpool. Required if admin wants to
     * create a subnetpool for another project. Changing this creates a new subnetpool.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnetpool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnetpool.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnetpool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnetpool.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The revision number of the subnetpool.
     * 
     */
    @Import(name="revisionNumber")
    private @Nullable Output<Integer> revisionNumber;

    /**
     * @return The revision number of the subnetpool.
     * 
     */
    public Optional<Output<Integer>> revisionNumber() {
        return Optional.ofNullable(this.revisionNumber);
    }

    /**
     * Indicates whether this subnetpool is shared across
     * all projects. Changing this updates the shared status of the existing
     * subnetpool.
     * 
     */
    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    /**
     * @return Indicates whether this subnetpool is shared across
     * all projects. Changing this updates the shared status of the existing
     * subnetpool.
     * 
     */
    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    /**
     * A set of string tags for the subnetpool.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of string tags for the subnetpool.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The time at which subnetpool was created.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return The time at which subnetpool was created.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
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

    private SubnetPoolState() {}

    private SubnetPoolState(SubnetPoolState $) {
        this.addressScopeId = $.addressScopeId;
        this.allTags = $.allTags;
        this.createdAt = $.createdAt;
        this.defaultPrefixlen = $.defaultPrefixlen;
        this.defaultQuota = $.defaultQuota;
        this.description = $.description;
        this.ipVersion = $.ipVersion;
        this.isDefault = $.isDefault;
        this.maxPrefixlen = $.maxPrefixlen;
        this.minPrefixlen = $.minPrefixlen;
        this.name = $.name;
        this.prefixes = $.prefixes;
        this.projectId = $.projectId;
        this.region = $.region;
        this.revisionNumber = $.revisionNumber;
        this.shared = $.shared;
        this.tags = $.tags;
        this.updatedAt = $.updatedAt;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubnetPoolState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubnetPoolState $;

        public Builder() {
            $ = new SubnetPoolState();
        }

        public Builder(SubnetPoolState defaults) {
            $ = new SubnetPoolState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressScopeId The Neutron address scope to assign to the
         * subnetpool. Changing this updates the address scope id of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder addressScopeId(@Nullable Output<String> addressScopeId) {
            $.addressScopeId = addressScopeId;
            return this;
        }

        /**
         * @param addressScopeId The Neutron address scope to assign to the
         * subnetpool. Changing this updates the address scope id of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder addressScopeId(String addressScopeId) {
            return addressScopeId(Output.of(addressScopeId));
        }

        /**
         * @param allTags The collection of tags assigned on the subnetpool, which have been
         * explicitly and implicitly added.
         * 
         * @return builder
         * 
         */
        public Builder allTags(@Nullable Output<List<String>> allTags) {
            $.allTags = allTags;
            return this;
        }

        /**
         * @param allTags The collection of tags assigned on the subnetpool, which have been
         * explicitly and implicitly added.
         * 
         * @return builder
         * 
         */
        public Builder allTags(List<String> allTags) {
            return allTags(Output.of(allTags));
        }

        /**
         * @param allTags The collection of tags assigned on the subnetpool, which have been
         * explicitly and implicitly added.
         * 
         * @return builder
         * 
         */
        public Builder allTags(String... allTags) {
            return allTags(List.of(allTags));
        }

        /**
         * @param createdAt The time at which subnetpool was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The time at which subnetpool was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param defaultPrefixlen The size of the prefix to allocate when the cidr
         * or prefixlen attributes are omitted when you create the subnet. Defaults to the
         * MinPrefixLen. Changing this updates the default prefixlen of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder defaultPrefixlen(@Nullable Output<Integer> defaultPrefixlen) {
            $.defaultPrefixlen = defaultPrefixlen;
            return this;
        }

        /**
         * @param defaultPrefixlen The size of the prefix to allocate when the cidr
         * or prefixlen attributes are omitted when you create the subnet. Defaults to the
         * MinPrefixLen. Changing this updates the default prefixlen of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder defaultPrefixlen(Integer defaultPrefixlen) {
            return defaultPrefixlen(Output.of(defaultPrefixlen));
        }

        /**
         * @param defaultQuota The per-project quota on the prefix space that can be
         * allocated from the subnetpool for project subnets. Changing this updates the
         * default quota of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder defaultQuota(@Nullable Output<Integer> defaultQuota) {
            $.defaultQuota = defaultQuota;
            return this;
        }

        /**
         * @param defaultQuota The per-project quota on the prefix space that can be
         * allocated from the subnetpool for project subnets. Changing this updates the
         * default quota of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder defaultQuota(Integer defaultQuota) {
            return defaultQuota(Output.of(defaultQuota));
        }

        /**
         * @param description The human-readable description for the subnetpool.
         * Changing this updates the description of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The human-readable description for the subnetpool.
         * Changing this updates the description of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ipVersion The IP protocol version.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(@Nullable Output<Integer> ipVersion) {
            $.ipVersion = ipVersion;
            return this;
        }

        /**
         * @param ipVersion The IP protocol version.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(Integer ipVersion) {
            return ipVersion(Output.of(ipVersion));
        }

        /**
         * @param isDefault Indicates whether the subnetpool is default
         * subnetpool or not. Changing this updates the default status of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder isDefault(@Nullable Output<Boolean> isDefault) {
            $.isDefault = isDefault;
            return this;
        }

        /**
         * @param isDefault Indicates whether the subnetpool is default
         * subnetpool or not. Changing this updates the default status of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder isDefault(Boolean isDefault) {
            return isDefault(Output.of(isDefault));
        }

        /**
         * @param maxPrefixlen The maximum prefix size that can be allocated from
         * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
         * default is 128. Changing this updates the max prefixlen of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder maxPrefixlen(@Nullable Output<Integer> maxPrefixlen) {
            $.maxPrefixlen = maxPrefixlen;
            return this;
        }

        /**
         * @param maxPrefixlen The maximum prefix size that can be allocated from
         * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
         * default is 128. Changing this updates the max prefixlen of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder maxPrefixlen(Integer maxPrefixlen) {
            return maxPrefixlen(Output.of(maxPrefixlen));
        }

        /**
         * @param minPrefixlen The smallest prefix that can be allocated from a
         * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
         * is 64. Changing this updates the min prefixlen of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder minPrefixlen(@Nullable Output<Integer> minPrefixlen) {
            $.minPrefixlen = minPrefixlen;
            return this;
        }

        /**
         * @param minPrefixlen The smallest prefix that can be allocated from a
         * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
         * is 64. Changing this updates the min prefixlen of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder minPrefixlen(Integer minPrefixlen) {
            return minPrefixlen(Output.of(minPrefixlen));
        }

        /**
         * @param name The name of the subnetpool. Changing this updates the name of
         * the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the subnetpool. Changing this updates the name of
         * the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param prefixes A list of subnet prefixes to assign to the subnetpool.
         * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
         * subnet prefix must be unique among all subnet prefixes in all subnetpools that
         * are associated with the address scope. Changing this updates the prefixes list
         * of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder prefixes(@Nullable Output<List<String>> prefixes) {
            $.prefixes = prefixes;
            return this;
        }

        /**
         * @param prefixes A list of subnet prefixes to assign to the subnetpool.
         * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
         * subnet prefix must be unique among all subnet prefixes in all subnetpools that
         * are associated with the address scope. Changing this updates the prefixes list
         * of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder prefixes(List<String> prefixes) {
            return prefixes(Output.of(prefixes));
        }

        /**
         * @param prefixes A list of subnet prefixes to assign to the subnetpool.
         * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
         * subnet prefix must be unique among all subnet prefixes in all subnetpools that
         * are associated with the address scope. Changing this updates the prefixes list
         * of the existing subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder prefixes(String... prefixes) {
            return prefixes(List.of(prefixes));
        }

        /**
         * @param projectId The owner of the subnetpool. Required if admin wants to
         * create a subnetpool for another project. Changing this creates a new subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The owner of the subnetpool. Required if admin wants to
         * create a subnetpool for another project. Changing this creates a new subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron subnetpool. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * subnetpool.
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
         * A Networking client is needed to create a Neutron subnetpool. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param revisionNumber The revision number of the subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder revisionNumber(@Nullable Output<Integer> revisionNumber) {
            $.revisionNumber = revisionNumber;
            return this;
        }

        /**
         * @param revisionNumber The revision number of the subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder revisionNumber(Integer revisionNumber) {
            return revisionNumber(Output.of(revisionNumber));
        }

        /**
         * @param shared Indicates whether this subnetpool is shared across
         * all projects. Changing this updates the shared status of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder shared(@Nullable Output<Boolean> shared) {
            $.shared = shared;
            return this;
        }

        /**
         * @param shared Indicates whether this subnetpool is shared across
         * all projects. Changing this updates the shared status of the existing
         * subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        /**
         * @param tags A set of string tags for the subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of string tags for the subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of string tags for the subnetpool.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param updatedAt The time at which subnetpool was created.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt The time at which subnetpool was created.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
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

        public SubnetPoolState build() {
            return $;
        }
    }

}
