// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFlavorV2PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFlavorV2PlainArgs Empty = new GetFlavorV2PlainArgs();

    @Import(name="flavorId")
    private @Nullable String flavorId;

    public Optional<String> flavorId() {
        return Optional.ofNullable(this.flavorId);
    }

    @Import(name="name")
    private @Nullable String name;

    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="region")
    private @Nullable String region;

    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetFlavorV2PlainArgs() {}

    private GetFlavorV2PlainArgs(GetFlavorV2PlainArgs $) {
        this.flavorId = $.flavorId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFlavorV2PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFlavorV2PlainArgs $;

        public Builder() {
            $ = new GetFlavorV2PlainArgs();
        }

        public Builder(GetFlavorV2PlainArgs defaults) {
            $ = new GetFlavorV2PlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder flavorId(@Nullable String flavorId) {
            $.flavorId = flavorId;
            return this;
        }

        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetFlavorV2PlainArgs build() {
            return $;
        }
    }

}