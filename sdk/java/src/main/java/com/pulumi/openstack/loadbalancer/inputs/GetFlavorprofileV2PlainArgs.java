// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFlavorprofileV2PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFlavorprofileV2PlainArgs Empty = new GetFlavorprofileV2PlainArgs();

    /**
     * The ID of the flavorprofile. Conflicts with `name` and
     * `provider_name`.
     * 
     */
    @Import(name="flavorprofileId")
    private @Nullable String flavorprofileId;

    /**
     * @return The ID of the flavorprofile. Conflicts with `name` and
     * `provider_name`.
     * 
     */
    public Optional<String> flavorprofileId() {
        return Optional.ofNullable(this.flavorprofileId);
    }

    /**
     * The name of the flavorprofile. Conflicts with `flavorprofile_id`.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the flavorprofile. Conflicts with `flavorprofile_id`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the provider that the flavorprofile uses. Conflicts
     * with `flavorprofile_id`.
     * 
     */
    @Import(name="providerName")
    private @Nullable String providerName;

    /**
     * @return The name of the provider that the flavorprofile uses. Conflicts
     * with `flavorprofile_id`.
     * 
     */
    public Optional<String> providerName() {
        return Optional.ofNullable(this.providerName);
    }

    /**
     * The region in which to obtain the V2 Load Balancer client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V2 Load Balancer client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetFlavorprofileV2PlainArgs() {}

    private GetFlavorprofileV2PlainArgs(GetFlavorprofileV2PlainArgs $) {
        this.flavorprofileId = $.flavorprofileId;
        this.name = $.name;
        this.providerName = $.providerName;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFlavorprofileV2PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFlavorprofileV2PlainArgs $;

        public Builder() {
            $ = new GetFlavorprofileV2PlainArgs();
        }

        public Builder(GetFlavorprofileV2PlainArgs defaults) {
            $ = new GetFlavorprofileV2PlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param flavorprofileId The ID of the flavorprofile. Conflicts with `name` and
         * `provider_name`.
         * 
         * @return builder
         * 
         */
        public Builder flavorprofileId(@Nullable String flavorprofileId) {
            $.flavorprofileId = flavorprofileId;
            return this;
        }

        /**
         * @param name The name of the flavorprofile. Conflicts with `flavorprofile_id`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param providerName The name of the provider that the flavorprofile uses. Conflicts
         * with `flavorprofile_id`.
         * 
         * @return builder
         * 
         */
        public Builder providerName(@Nullable String providerName) {
            $.providerName = providerName;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Load Balancer client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetFlavorprofileV2PlainArgs build() {
            return $;
        }
    }

}
