// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAuthScopeArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthScopeArgs Empty = new GetAuthScopeArgs();

    /**
     * The name of the scope. This is an arbitrary name which is
     * only used as a unique identifier so an actual token isn&#39;t used as the ID.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the scope. This is an arbitrary name which is
     * only used as a unique identifier so an actual token isn&#39;t used as the ID.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The region in which to obtain the V3 Identity client.
     * A Identity client is needed to retrieve tokens IDs. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V3 Identity client.
     * A Identity client is needed to retrieve tokens IDs. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetAuthScopeArgs() {}

    private GetAuthScopeArgs(GetAuthScopeArgs $) {
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthScopeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthScopeArgs $;

        public Builder() {
            $ = new GetAuthScopeArgs();
        }

        public Builder(GetAuthScopeArgs defaults) {
            $ = new GetAuthScopeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the scope. This is an arbitrary name which is
         * only used as a unique identifier so an actual token isn&#39;t used as the ID.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the scope. This is an arbitrary name which is
         * only used as a unique identifier so an actual token isn&#39;t used as the ID.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V3 Identity client.
         * A Identity client is needed to retrieve tokens IDs. If omitted, the
         * `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V3 Identity client.
         * A Identity client is needed to retrieve tokens IDs. If omitted, the
         * `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetAuthScopeArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
