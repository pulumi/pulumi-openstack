// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.compute.inputs.GetInstanceV2NetworkArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceV2Args extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceV2Args Empty = new GetInstanceV2Args();

    /**
     * The UUID of the instance
     * 
     */
    @Import(name="id", required=true)
    private Output<String> id;

    /**
     * @return The UUID of the instance
     * 
     */
    public Output<String> id() {
        return this.id;
    }

    /**
     * An array of maps, detailed below.
     * 
     */
    @Import(name="networks")
    private @Nullable Output<List<GetInstanceV2NetworkArgs>> networks;

    /**
     * @return An array of maps, detailed below.
     * 
     */
    public Optional<Output<List<GetInstanceV2NetworkArgs>>> networks() {
        return Optional.ofNullable(this.networks);
    }

    @Import(name="region")
    private @Nullable Output<String> region;

    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The user data added when the server was created.
     * 
     */
    @Import(name="userData")
    private @Nullable Output<String> userData;

    /**
     * @return The user data added when the server was created.
     * 
     */
    public Optional<Output<String>> userData() {
        return Optional.ofNullable(this.userData);
    }

    private GetInstanceV2Args() {}

    private GetInstanceV2Args(GetInstanceV2Args $) {
        this.id = $.id;
        this.networks = $.networks;
        this.region = $.region;
        this.userData = $.userData;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceV2Args $;

        public Builder() {
            $ = new GetInstanceV2Args();
        }

        public Builder(GetInstanceV2Args defaults) {
            $ = new GetInstanceV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The UUID of the instance
         * 
         * @return builder
         * 
         */
        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The UUID of the instance
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param networks An array of maps, detailed below.
         * 
         * @return builder
         * 
         */
        public Builder networks(@Nullable Output<List<GetInstanceV2NetworkArgs>> networks) {
            $.networks = networks;
            return this;
        }

        /**
         * @param networks An array of maps, detailed below.
         * 
         * @return builder
         * 
         */
        public Builder networks(List<GetInstanceV2NetworkArgs> networks) {
            return networks(Output.of(networks));
        }

        /**
         * @param networks An array of maps, detailed below.
         * 
         * @return builder
         * 
         */
        public Builder networks(GetInstanceV2NetworkArgs... networks) {
            return networks(List.of(networks));
        }

        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param userData The user data added when the server was created.
         * 
         * @return builder
         * 
         */
        public Builder userData(@Nullable Output<String> userData) {
            $.userData = userData;
            return this;
        }

        /**
         * @param userData The user data added when the server was created.
         * 
         * @return builder
         * 
         */
        public Builder userData(String userData) {
            return userData(Output.of(userData));
        }

        public GetInstanceV2Args build() {
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetInstanceV2Args", "id");
            }
            return $;
        }
    }

}
