// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FlavorAccessArgs extends com.pulumi.resources.ResourceArgs {

    public static final FlavorAccessArgs Empty = new FlavorAccessArgs();

    /**
     * The UUID of flavor to use. Changing this creates a new flavor access.
     * 
     */
    @Import(name="flavorId", required=true)
    private Output<String> flavorId;

    /**
     * @return The UUID of flavor to use. Changing this creates a new flavor access.
     * 
     */
    public Output<String> flavorId() {
        return this.flavorId;
    }

    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor access.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor access.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The UUID of tenant which is allowed to use the flavor.
     * Changing this creates a new flavor access.
     * 
     */
    @Import(name="tenantId", required=true)
    private Output<String> tenantId;

    /**
     * @return The UUID of tenant which is allowed to use the flavor.
     * Changing this creates a new flavor access.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    private FlavorAccessArgs() {}

    private FlavorAccessArgs(FlavorAccessArgs $) {
        this.flavorId = $.flavorId;
        this.region = $.region;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FlavorAccessArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FlavorAccessArgs $;

        public Builder() {
            $ = new FlavorAccessArgs();
        }

        public Builder(FlavorAccessArgs defaults) {
            $ = new FlavorAccessArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param flavorId The UUID of flavor to use. Changing this creates a new flavor access.
         * 
         * @return builder
         * 
         */
        public Builder flavorId(Output<String> flavorId) {
            $.flavorId = flavorId;
            return this;
        }

        /**
         * @param flavorId The UUID of flavor to use. Changing this creates a new flavor access.
         * 
         * @return builder
         * 
         */
        public Builder flavorId(String flavorId) {
            return flavorId(Output.of(flavorId));
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * If omitted, the `region` argument of the provider is used.
         * Changing this creates a new flavor access.
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
         * If omitted, the `region` argument of the provider is used.
         * Changing this creates a new flavor access.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tenantId The UUID of tenant which is allowed to use the flavor.
         * Changing this creates a new flavor access.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The UUID of tenant which is allowed to use the flavor.
         * Changing this creates a new flavor access.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public FlavorAccessArgs build() {
            if ($.flavorId == null) {
                throw new MissingRequiredPropertyException("FlavorAccessArgs", "flavorId");
            }
            if ($.tenantId == null) {
                throw new MissingRequiredPropertyException("FlavorAccessArgs", "tenantId");
            }
            return $;
        }
    }

}