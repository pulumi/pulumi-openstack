// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LimitV3Args extends com.pulumi.resources.ResourceArgs {

    public static final LimitV3Args Empty = new LimitV3Args();

    /**
     * Description of the limit.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the limit.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The domain the limit applies to. Changing this
     * creates a new Limit.
     * 
     */
    @Import(name="domainId")
    private @Nullable Output<String> domainId;

    /**
     * @return The domain the limit applies to. Changing this
     * creates a new Limit.
     * 
     */
    public Optional<Output<String>> domainId() {
        return Optional.ofNullable(this.domainId);
    }

    /**
     * The project the limit applies to. Changing this
     * creates a new Limit.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The project the limit applies to. Changing this
     * creates a new Limit.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new Limit.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new Limit.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Integer for the actual limit.
     * 
     */
    @Import(name="resourceLimit", required=true)
    private Output<Integer> resourceLimit;

    /**
     * @return Integer for the actual limit.
     * 
     */
    public Output<Integer> resourceLimit() {
        return this.resourceLimit;
    }

    /**
     * The resource that the limit applies to. Changing
     * this creates a new Limit.
     * 
     */
    @Import(name="resourceName", required=true)
    private Output<String> resourceName;

    /**
     * @return The resource that the limit applies to. Changing
     * this creates a new Limit.
     * 
     */
    public Output<String> resourceName() {
        return this.resourceName;
    }

    /**
     * The service the limit applies to. Changing this
     * creates a new Limit.
     * 
     */
    @Import(name="serviceId", required=true)
    private Output<String> serviceId;

    /**
     * @return The service the limit applies to. Changing this
     * creates a new Limit.
     * 
     */
    public Output<String> serviceId() {
        return this.serviceId;
    }

    private LimitV3Args() {}

    private LimitV3Args(LimitV3Args $) {
        this.description = $.description;
        this.domainId = $.domainId;
        this.projectId = $.projectId;
        this.region = $.region;
        this.resourceLimit = $.resourceLimit;
        this.resourceName = $.resourceName;
        this.serviceId = $.serviceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LimitV3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LimitV3Args $;

        public Builder() {
            $ = new LimitV3Args();
        }

        public Builder(LimitV3Args defaults) {
            $ = new LimitV3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the limit.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the limit.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param domainId The domain the limit applies to. Changing this
         * creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder domainId(@Nullable Output<String> domainId) {
            $.domainId = domainId;
            return this;
        }

        /**
         * @param domainId The domain the limit applies to. Changing this
         * creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder domainId(String domainId) {
            return domainId(Output.of(domainId));
        }

        /**
         * @param projectId The project the limit applies to. Changing this
         * creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The project the limit applies to. Changing this
         * creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V3 Keystone client.
         * If omitted, the `region` argument of the provider is used. Changing this
         * creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V3 Keystone client.
         * If omitted, the `region` argument of the provider is used. Changing this
         * creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param resourceLimit Integer for the actual limit.
         * 
         * @return builder
         * 
         */
        public Builder resourceLimit(Output<Integer> resourceLimit) {
            $.resourceLimit = resourceLimit;
            return this;
        }

        /**
         * @param resourceLimit Integer for the actual limit.
         * 
         * @return builder
         * 
         */
        public Builder resourceLimit(Integer resourceLimit) {
            return resourceLimit(Output.of(resourceLimit));
        }

        /**
         * @param resourceName The resource that the limit applies to. Changing
         * this creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder resourceName(Output<String> resourceName) {
            $.resourceName = resourceName;
            return this;
        }

        /**
         * @param resourceName The resource that the limit applies to. Changing
         * this creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder resourceName(String resourceName) {
            return resourceName(Output.of(resourceName));
        }

        /**
         * @param serviceId The service the limit applies to. Changing this
         * creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(Output<String> serviceId) {
            $.serviceId = serviceId;
            return this;
        }

        /**
         * @param serviceId The service the limit applies to. Changing this
         * creates a new Limit.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(String serviceId) {
            return serviceId(Output.of(serviceId));
        }

        public LimitV3Args build() {
            if ($.resourceLimit == null) {
                throw new MissingRequiredPropertyException("LimitV3Args", "resourceLimit");
            }
            if ($.resourceName == null) {
                throw new MissingRequiredPropertyException("LimitV3Args", "resourceName");
            }
            if ($.serviceId == null) {
                throw new MissingRequiredPropertyException("LimitV3Args", "serviceId");
            }
            return $;
        }
    }

}
