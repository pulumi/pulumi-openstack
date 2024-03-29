// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetGroupPlainArgs Empty = new GetGroupPlainArgs();

    /**
     * The domain the group belongs to.
     * 
     */
    @Import(name="domainId")
    private @Nullable String domainId;

    /**
     * @return The domain the group belongs to.
     * 
     */
    public Optional<String> domainId() {
        return Optional.ofNullable(this.domainId);
    }

    /**
     * The name of the group.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the group.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetGroupPlainArgs() {}

    private GetGroupPlainArgs(GetGroupPlainArgs $) {
        this.domainId = $.domainId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetGroupPlainArgs $;

        public Builder() {
            $ = new GetGroupPlainArgs();
        }

        public Builder(GetGroupPlainArgs defaults) {
            $ = new GetGroupPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domainId The domain the group belongs to.
         * 
         * @return builder
         * 
         */
        public Builder domainId(@Nullable String domainId) {
            $.domainId = domainId;
            return this;
        }

        /**
         * @param name The name of the group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param region The region in which to obtain the V3 Keystone client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetGroupPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetGroupPlainArgs", "name");
            }
            return $;
        }
    }

}
