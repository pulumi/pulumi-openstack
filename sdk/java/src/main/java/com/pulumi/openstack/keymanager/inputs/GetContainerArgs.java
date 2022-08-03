// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetContainerArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContainerArgs Empty = new GetContainerArgs();

    /**
     * The Container name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The Container name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a container. If omitted, the `region`
     * argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a container. If omitted, the `region`
     * argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetContainerArgs() {}

    private GetContainerArgs(GetContainerArgs $) {
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContainerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContainerArgs $;

        public Builder() {
            $ = new GetContainerArgs();
        }

        public Builder(GetContainerArgs defaults) {
            $ = new GetContainerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The Container name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The Container name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V1 KeyManager client.
         * A KeyManager client is needed to fetch a container. If omitted, the `region`
         * argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
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
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetContainerArgs build() {
            return $;
        }
    }

}
