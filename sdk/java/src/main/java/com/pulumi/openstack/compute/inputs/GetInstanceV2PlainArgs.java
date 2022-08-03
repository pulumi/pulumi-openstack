// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.compute.inputs.GetInstanceV2Network;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetInstanceV2PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstanceV2PlainArgs Empty = new GetInstanceV2PlainArgs();

    /**
     * The UUID of the instance
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return The UUID of the instance
     * 
     */
    public String id() {
        return this.id;
    }

    /**
     * An array of maps, detailed below.
     * 
     */
    @Import(name="networks")
    private @Nullable List<GetInstanceV2Network> networks;

    /**
     * @return An array of maps, detailed below.
     * 
     */
    public Optional<List<GetInstanceV2Network>> networks() {
        return Optional.ofNullable(this.networks);
    }

    @Import(name="region")
    private @Nullable String region;

    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The user data added when the server was created.
     * 
     */
    @Import(name="userData")
    private @Nullable String userData;

    /**
     * @return The user data added when the server was created.
     * 
     */
    public Optional<String> userData() {
        return Optional.ofNullable(this.userData);
    }

    private GetInstanceV2PlainArgs() {}

    private GetInstanceV2PlainArgs(GetInstanceV2PlainArgs $) {
        this.id = $.id;
        this.networks = $.networks;
        this.region = $.region;
        this.userData = $.userData;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstanceV2PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstanceV2PlainArgs $;

        public Builder() {
            $ = new GetInstanceV2PlainArgs();
        }

        public Builder(GetInstanceV2PlainArgs defaults) {
            $ = new GetInstanceV2PlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The UUID of the instance
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        /**
         * @param networks An array of maps, detailed below.
         * 
         * @return builder
         * 
         */
        public Builder networks(@Nullable List<GetInstanceV2Network> networks) {
            $.networks = networks;
            return this;
        }

        /**
         * @param networks An array of maps, detailed below.
         * 
         * @return builder
         * 
         */
        public Builder networks(GetInstanceV2Network... networks) {
            return networks(List.of(networks));
        }

        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        /**
         * @param userData The user data added when the server was created.
         * 
         * @return builder
         * 
         */
        public Builder userData(@Nullable String userData) {
            $.userData = userData;
            return this;
        }

        public GetInstanceV2PlainArgs build() {
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            return $;
        }
    }

}
