// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LbFlavorprofileV2Args extends com.pulumi.resources.ResourceArgs {

    public static final LbFlavorprofileV2Args Empty = new LbFlavorprofileV2Args();

    /**
     * String that passes the flavor_data for the flavorprofile.
     * The data that are allowed depend on the `provider_name` that is passed. jsonencode
     * can be used for readability as shown in the example above.
     * Changing this updates the existing flavorprofile.
     * 
     */
    @Import(name="flavorData", required=true)
    private Output<String> flavorData;

    /**
     * @return String that passes the flavor_data for the flavorprofile.
     * The data that are allowed depend on the `provider_name` that is passed. jsonencode
     * can be used for readability as shown in the example above.
     * Changing this updates the existing flavorprofile.
     * 
     */
    public Output<String> flavorData() {
        return this.flavorData;
    }

    /**
     * Name of the flavorprofile. Changing this updates the existing
     * flavorprofile.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the flavorprofile. Changing this updates the existing
     * flavorprofile.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The provider_name that the flavor_profile will use.
     * Changing this updates the existing flavorprofile.
     * 
     */
    @Import(name="providerName", required=true)
    private Output<String> providerName;

    /**
     * @return The provider_name that the flavor_profile will use.
     * Changing this updates the existing flavorprofile.
     * 
     */
    public Output<String> providerName() {
        return this.providerName;
    }

    @Import(name="region")
    private @Nullable Output<String> region;

    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private LbFlavorprofileV2Args() {}

    private LbFlavorprofileV2Args(LbFlavorprofileV2Args $) {
        this.flavorData = $.flavorData;
        this.name = $.name;
        this.providerName = $.providerName;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LbFlavorprofileV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LbFlavorprofileV2Args $;

        public Builder() {
            $ = new LbFlavorprofileV2Args();
        }

        public Builder(LbFlavorprofileV2Args defaults) {
            $ = new LbFlavorprofileV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param flavorData String that passes the flavor_data for the flavorprofile.
         * The data that are allowed depend on the `provider_name` that is passed. jsonencode
         * can be used for readability as shown in the example above.
         * Changing this updates the existing flavorprofile.
         * 
         * @return builder
         * 
         */
        public Builder flavorData(Output<String> flavorData) {
            $.flavorData = flavorData;
            return this;
        }

        /**
         * @param flavorData String that passes the flavor_data for the flavorprofile.
         * The data that are allowed depend on the `provider_name` that is passed. jsonencode
         * can be used for readability as shown in the example above.
         * Changing this updates the existing flavorprofile.
         * 
         * @return builder
         * 
         */
        public Builder flavorData(String flavorData) {
            return flavorData(Output.of(flavorData));
        }

        /**
         * @param name Name of the flavorprofile. Changing this updates the existing
         * flavorprofile.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the flavorprofile. Changing this updates the existing
         * flavorprofile.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param providerName The provider_name that the flavor_profile will use.
         * Changing this updates the existing flavorprofile.
         * 
         * @return builder
         * 
         */
        public Builder providerName(Output<String> providerName) {
            $.providerName = providerName;
            return this;
        }

        /**
         * @param providerName The provider_name that the flavor_profile will use.
         * Changing this updates the existing flavorprofile.
         * 
         * @return builder
         * 
         */
        public Builder providerName(String providerName) {
            return providerName(Output.of(providerName));
        }

        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        public Builder region(String region) {
            return region(Output.of(region));
        }

        public LbFlavorprofileV2Args build() {
            if ($.flavorData == null) {
                throw new MissingRequiredPropertyException("LbFlavorprofileV2Args", "flavorData");
            }
            if ($.providerName == null) {
                throw new MissingRequiredPropertyException("LbFlavorprofileV2Args", "providerName");
            }
            return $;
        }
    }

}