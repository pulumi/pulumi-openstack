// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetContainerPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContainerPlainArgs Empty = new GetContainerPlainArgs();

    /**
     * The Container name.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The Container name.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a container. If omitted, the `region`
     * argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a container. If omitted, the `region`
     * argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetContainerPlainArgs() {}

    private GetContainerPlainArgs(GetContainerPlainArgs $) {
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContainerPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContainerPlainArgs $;

        public Builder() {
            $ = new GetContainerPlainArgs();
        }

        public Builder(GetContainerPlainArgs defaults) {
            $ = new GetContainerPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The Container name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param region The region in which to obtain the V1 KeyManager client.
         * A KeyManager client is needed to fetch a container. If omitted, the `region`
         * argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetContainerPlainArgs build() {
            return $;
        }
    }

}
