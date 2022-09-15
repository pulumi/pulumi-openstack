// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NodeGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final NodeGroupArgs Empty = new NodeGroupArgs();

    /**
     * The UUID of the V1 Container Infra cluster.
     * Changing this creates a new node group.
     * 
     */
    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    /**
     * @return The UUID of the V1 Container Infra cluster.
     * Changing this creates a new node group.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }

    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     * 
     */
    @Import(name="dockerVolumeSize")
    private @Nullable Output<Integer> dockerVolumeSize;

    /**
     * @return The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     * 
     */
    public Optional<Output<Integer>> dockerVolumeSize() {
        return Optional.ofNullable(this.dockerVolumeSize);
    }

    @Import(name="flavorId")
    private @Nullable Output<String> flavorId;

    public Optional<Output<String>> flavorId() {
        return Optional.ofNullable(this.flavorId);
    }

    @Import(name="imageId")
    private @Nullable Output<String> imageId;

    public Optional<Output<String>> imageId() {
        return Optional.ofNullable(this.imageId);
    }

    /**
     * The list of key value pairs representing additional
     * properties of the node group. Changing this creates a new node group.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,Object>> labels;

    /**
     * @return The list of key value pairs representing additional
     * properties of the node group. Changing this creates a new node group.
     * 
     */
    public Optional<Output<Map<String,Object>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The maximum number of nodes for the node group.
     * Changing this update the maximum number of nodes of the node group.
     * 
     */
    @Import(name="maxNodeCount")
    private @Nullable Output<Integer> maxNodeCount;

    /**
     * @return The maximum number of nodes for the node group.
     * Changing this update the maximum number of nodes of the node group.
     * 
     */
    public Optional<Output<Integer>> maxNodeCount() {
        return Optional.ofNullable(this.maxNodeCount);
    }

    /**
     * Indicates whether the provided labels should be
     * merged with cluster labels. Changing this creates a new nodegroup.
     * 
     */
    @Import(name="mergeLabels")
    private @Nullable Output<Boolean> mergeLabels;

    /**
     * @return Indicates whether the provided labels should be
     * merged with cluster labels. Changing this creates a new nodegroup.
     * 
     */
    public Optional<Output<Boolean>> mergeLabels() {
        return Optional.ofNullable(this.mergeLabels);
    }

    /**
     * The minimum number of nodes for the node group.
     * Changing this update the minimum number of nodes of the node group.
     * 
     */
    @Import(name="minNodeCount")
    private @Nullable Output<Integer> minNodeCount;

    /**
     * @return The minimum number of nodes for the node group.
     * Changing this update the minimum number of nodes of the node group.
     * 
     */
    public Optional<Output<Integer>> minNodeCount() {
        return Optional.ofNullable(this.minNodeCount);
    }

    /**
     * The name of the node group. Changing this creates a new
     * node group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the node group. Changing this creates a new
     * node group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The number of nodes for the node group. Changing
     * this update the number of nodes of the node group.
     * 
     */
    @Import(name="nodeCount")
    private @Nullable Output<Integer> nodeCount;

    /**
     * @return The number of nodes for the node group. Changing
     * this update the number of nodes of the node group.
     * 
     */
    public Optional<Output<Integer>> nodeCount() {
        return Optional.ofNullable(this.nodeCount);
    }

    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * node group.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * node group.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private NodeGroupArgs() {}

    private NodeGroupArgs(NodeGroupArgs $) {
        this.clusterId = $.clusterId;
        this.dockerVolumeSize = $.dockerVolumeSize;
        this.flavorId = $.flavorId;
        this.imageId = $.imageId;
        this.labels = $.labels;
        this.maxNodeCount = $.maxNodeCount;
        this.mergeLabels = $.mergeLabels;
        this.minNodeCount = $.minNodeCount;
        this.name = $.name;
        this.nodeCount = $.nodeCount;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NodeGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NodeGroupArgs $;

        public Builder() {
            $ = new NodeGroupArgs();
        }

        public Builder(NodeGroupArgs defaults) {
            $ = new NodeGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId The UUID of the V1 Container Infra cluster.
         * Changing this creates a new node group.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId The UUID of the V1 Container Infra cluster.
         * Changing this creates a new node group.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param dockerVolumeSize The size (in GB) of the Docker volume.
         * Changing this creates a new node group.
         * 
         * @return builder
         * 
         */
        public Builder dockerVolumeSize(@Nullable Output<Integer> dockerVolumeSize) {
            $.dockerVolumeSize = dockerVolumeSize;
            return this;
        }

        /**
         * @param dockerVolumeSize The size (in GB) of the Docker volume.
         * Changing this creates a new node group.
         * 
         * @return builder
         * 
         */
        public Builder dockerVolumeSize(Integer dockerVolumeSize) {
            return dockerVolumeSize(Output.of(dockerVolumeSize));
        }

        public Builder flavorId(@Nullable Output<String> flavorId) {
            $.flavorId = flavorId;
            return this;
        }

        public Builder flavorId(String flavorId) {
            return flavorId(Output.of(flavorId));
        }

        public Builder imageId(@Nullable Output<String> imageId) {
            $.imageId = imageId;
            return this;
        }

        public Builder imageId(String imageId) {
            return imageId(Output.of(imageId));
        }

        /**
         * @param labels The list of key value pairs representing additional
         * properties of the node group. Changing this creates a new node group.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,Object>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels The list of key value pairs representing additional
         * properties of the node group. Changing this creates a new node group.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,Object> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param maxNodeCount The maximum number of nodes for the node group.
         * Changing this update the maximum number of nodes of the node group.
         * 
         * @return builder
         * 
         */
        public Builder maxNodeCount(@Nullable Output<Integer> maxNodeCount) {
            $.maxNodeCount = maxNodeCount;
            return this;
        }

        /**
         * @param maxNodeCount The maximum number of nodes for the node group.
         * Changing this update the maximum number of nodes of the node group.
         * 
         * @return builder
         * 
         */
        public Builder maxNodeCount(Integer maxNodeCount) {
            return maxNodeCount(Output.of(maxNodeCount));
        }

        /**
         * @param mergeLabels Indicates whether the provided labels should be
         * merged with cluster labels. Changing this creates a new nodegroup.
         * 
         * @return builder
         * 
         */
        public Builder mergeLabels(@Nullable Output<Boolean> mergeLabels) {
            $.mergeLabels = mergeLabels;
            return this;
        }

        /**
         * @param mergeLabels Indicates whether the provided labels should be
         * merged with cluster labels. Changing this creates a new nodegroup.
         * 
         * @return builder
         * 
         */
        public Builder mergeLabels(Boolean mergeLabels) {
            return mergeLabels(Output.of(mergeLabels));
        }

        /**
         * @param minNodeCount The minimum number of nodes for the node group.
         * Changing this update the minimum number of nodes of the node group.
         * 
         * @return builder
         * 
         */
        public Builder minNodeCount(@Nullable Output<Integer> minNodeCount) {
            $.minNodeCount = minNodeCount;
            return this;
        }

        /**
         * @param minNodeCount The minimum number of nodes for the node group.
         * Changing this update the minimum number of nodes of the node group.
         * 
         * @return builder
         * 
         */
        public Builder minNodeCount(Integer minNodeCount) {
            return minNodeCount(Output.of(minNodeCount));
        }

        /**
         * @param name The name of the node group. Changing this creates a new
         * node group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the node group. Changing this creates a new
         * node group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nodeCount The number of nodes for the node group. Changing
         * this update the number of nodes of the node group.
         * 
         * @return builder
         * 
         */
        public Builder nodeCount(@Nullable Output<Integer> nodeCount) {
            $.nodeCount = nodeCount;
            return this;
        }

        /**
         * @param nodeCount The number of nodes for the node group. Changing
         * this update the number of nodes of the node group.
         * 
         * @return builder
         * 
         */
        public Builder nodeCount(Integer nodeCount) {
            return nodeCount(Output.of(nodeCount));
        }

        /**
         * @param region The region in which to obtain the V1 Container Infra
         * client. A Container Infra client is needed to create a cluster. If omitted,
         * the `region` argument of the provider is used. Changing this creates a new
         * node group.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V1 Container Infra
         * client. A Container Infra client is needed to create a cluster. If omitted,
         * the `region` argument of the provider is used. Changing this creates a new
         * node group.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public NodeGroupArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            return $;
        }
    }

}