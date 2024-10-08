// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QosPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final QosPolicyArgs Empty = new QosPolicyArgs();

    /**
     * The human-readable description for the QoS policy.
     * Changing this updates the description of the existing QoS policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The human-readable description for the QoS policy.
     * Changing this updates the description of the existing QoS policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Indicates whether the QoS policy is default
     * QoS policy or not. Changing this updates the default status of the existing
     * QoS policy.
     * 
     */
    @Import(name="isDefault")
    private @Nullable Output<Boolean> isDefault;

    /**
     * @return Indicates whether the QoS policy is default
     * QoS policy or not. Changing this updates the default status of the existing
     * QoS policy.
     * 
     */
    public Optional<Output<Boolean>> isDefault() {
        return Optional.ofNullable(this.isDefault);
    }

    /**
     * The name of the QoS policy. Changing this updates the name of
     * the existing QoS policy.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the QoS policy. Changing this updates the name of
     * the existing QoS policy.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The owner of the QoS policy. Required if admin wants to
     * create a QoS policy for another project. Changing this creates a new QoS policy.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The owner of the QoS policy. Required if admin wants to
     * create a QoS policy for another project. Changing this creates a new QoS policy.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron Qos policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * QoS policy.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron Qos policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * QoS policy.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Indicates whether this QoS policy is shared across
     * all projects. Changing this updates the shared status of the existing
     * QoS policy.
     * 
     */
    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    /**
     * @return Indicates whether this QoS policy is shared across
     * all projects. Changing this updates the shared status of the existing
     * QoS policy.
     * 
     */
    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    /**
     * A set of string tags for the QoS policy.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of string tags for the QoS policy.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
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

    private QosPolicyArgs() {}

    private QosPolicyArgs(QosPolicyArgs $) {
        this.description = $.description;
        this.isDefault = $.isDefault;
        this.name = $.name;
        this.projectId = $.projectId;
        this.region = $.region;
        this.shared = $.shared;
        this.tags = $.tags;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QosPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QosPolicyArgs $;

        public Builder() {
            $ = new QosPolicyArgs();
        }

        public Builder(QosPolicyArgs defaults) {
            $ = new QosPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The human-readable description for the QoS policy.
         * Changing this updates the description of the existing QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The human-readable description for the QoS policy.
         * Changing this updates the description of the existing QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param isDefault Indicates whether the QoS policy is default
         * QoS policy or not. Changing this updates the default status of the existing
         * QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder isDefault(@Nullable Output<Boolean> isDefault) {
            $.isDefault = isDefault;
            return this;
        }

        /**
         * @param isDefault Indicates whether the QoS policy is default
         * QoS policy or not. Changing this updates the default status of the existing
         * QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder isDefault(Boolean isDefault) {
            return isDefault(Output.of(isDefault));
        }

        /**
         * @param name The name of the QoS policy. Changing this updates the name of
         * the existing QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the QoS policy. Changing this updates the name of
         * the existing QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The owner of the QoS policy. Required if admin wants to
         * create a QoS policy for another project. Changing this creates a new QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The owner of the QoS policy. Required if admin wants to
         * create a QoS policy for another project. Changing this creates a new QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron Qos policy. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * QoS policy.
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
         * A Networking client is needed to create a Neutron Qos policy. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param shared Indicates whether this QoS policy is shared across
         * all projects. Changing this updates the shared status of the existing
         * QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder shared(@Nullable Output<Boolean> shared) {
            $.shared = shared;
            return this;
        }

        /**
         * @param shared Indicates whether this QoS policy is shared across
         * all projects. Changing this updates the shared status of the existing
         * QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        /**
         * @param tags A set of string tags for the QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of string tags for the QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of string tags for the QoS policy.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
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

        public QosPolicyArgs build() {
            return $;
        }
    }

}
