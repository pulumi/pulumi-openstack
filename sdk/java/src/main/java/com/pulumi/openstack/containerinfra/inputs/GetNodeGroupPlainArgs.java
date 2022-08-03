// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetNodeGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNodeGroupPlainArgs Empty = new GetNodeGroupPlainArgs();

    /**
     * The name of the OpenStack Magnum cluster.
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return The name of the OpenStack Magnum cluster.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * The name of the node group.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the node group.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V1 Container Infra
     * client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V1 Container Infra
     * client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetNodeGroupPlainArgs() {}

    private GetNodeGroupPlainArgs(GetNodeGroupPlainArgs $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNodeGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNodeGroupPlainArgs $;

        public Builder() {
            $ = new GetNodeGroupPlainArgs();
        }

        public Builder(GetNodeGroupPlainArgs defaults) {
            $ = new GetNodeGroupPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId The name of the OpenStack Magnum cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param name The name of the node group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param region The region in which to obtain the V1 Container Infra
         * client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetNodeGroupPlainArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            return $;
        }
    }

}
