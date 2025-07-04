// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClusterArgs Empty = new ClusterArgs();

    /**
     * The UUID of the V1 Container Infra cluster
     * template. Changing this creates a new cluster.
     * 
     */
    @Import(name="clusterTemplateId")
    private @Nullable Output<String> clusterTemplateId;

    /**
     * @return The UUID of the V1 Container Infra cluster
     * template. Changing this creates a new cluster.
     * 
     */
    public Optional<Output<String>> clusterTemplateId() {
        return Optional.ofNullable(this.clusterTemplateId);
    }

    /**
     * The timeout (in minutes) for creating the
     * cluster. Changing this creates a new cluster.
     * 
     */
    @Import(name="createTimeout")
    private @Nullable Output<Integer> createTimeout;

    /**
     * @return The timeout (in minutes) for creating the
     * cluster. Changing this creates a new cluster.
     * 
     */
    public Optional<Output<Integer>> createTimeout() {
        return Optional.ofNullable(this.createTimeout);
    }

    /**
     * The URL used for cluster node discovery.
     * Changing this creates a new cluster.
     * 
     */
    @Import(name="discoveryUrl")
    private @Nullable Output<String> discoveryUrl;

    /**
     * @return The URL used for cluster node discovery.
     * Changing this creates a new cluster.
     * 
     */
    public Optional<Output<String>> discoveryUrl() {
        return Optional.ofNullable(this.discoveryUrl);
    }

    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new cluster.
     * 
     */
    @Import(name="dockerVolumeSize")
    private @Nullable Output<Integer> dockerVolumeSize;

    /**
     * @return The size (in GB) of the Docker volume.
     * Changing this creates a new cluster.
     * 
     */
    public Optional<Output<Integer>> dockerVolumeSize() {
        return Optional.ofNullable(this.dockerVolumeSize);
    }

    /**
     * The fixed network that will be attached to the
     * cluster. Changing this creates a new cluster.
     * 
     */
    @Import(name="fixedNetwork")
    private @Nullable Output<String> fixedNetwork;

    /**
     * @return The fixed network that will be attached to the
     * cluster. Changing this creates a new cluster.
     * 
     */
    public Optional<Output<String>> fixedNetwork() {
        return Optional.ofNullable(this.fixedNetwork);
    }

    /**
     * The fixed subnet that will be attached to the
     * cluster. Changing this creates a new cluster.
     * 
     */
    @Import(name="fixedSubnet")
    private @Nullable Output<String> fixedSubnet;

    /**
     * @return The fixed subnet that will be attached to the
     * cluster. Changing this creates a new cluster.
     * 
     */
    public Optional<Output<String>> fixedSubnet() {
        return Optional.ofNullable(this.fixedSubnet);
    }

    /**
     * The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * cluster.
     * 
     */
    @Import(name="flavor")
    private @Nullable Output<String> flavor;

    /**
     * @return The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * cluster.
     * 
     */
    public Optional<Output<String>> flavor() {
        return Optional.ofNullable(this.flavor);
    }

    /**
     * Indicates whether floating IP should be
     * created for every cluster node. Changing this creates a new cluster.
     * 
     */
    @Import(name="floatingIpEnabled")
    private @Nullable Output<Boolean> floatingIpEnabled;

    /**
     * @return Indicates whether floating IP should be
     * created for every cluster node. Changing this creates a new cluster.
     * 
     */
    public Optional<Output<Boolean>> floatingIpEnabled() {
        return Optional.ofNullable(this.floatingIpEnabled);
    }

    /**
     * The name of the Compute service SSH keypair. Changing
     * this creates a new cluster.
     * 
     */
    @Import(name="keypair")
    private @Nullable Output<String> keypair;

    /**
     * @return The name of the Compute service SSH keypair. Changing
     * this creates a new cluster.
     * 
     */
    public Optional<Output<String>> keypair() {
        return Optional.ofNullable(this.keypair);
    }

    /**
     * The list of key value pairs representing additional
     * properties of the cluster. Changing this creates a new cluster.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return The list of key value pairs representing additional
     * properties of the cluster. Changing this creates a new cluster.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The number of master nodes for the cluster.
     * Changing this creates a new cluster.
     * 
     */
    @Import(name="masterCount")
    private @Nullable Output<Integer> masterCount;

    /**
     * @return The number of master nodes for the cluster.
     * Changing this creates a new cluster.
     * 
     */
    public Optional<Output<Integer>> masterCount() {
        return Optional.ofNullable(this.masterCount);
    }

    /**
     * The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
     * new cluster.
     * 
     */
    @Import(name="masterFlavor")
    private @Nullable Output<String> masterFlavor;

    /**
     * @return The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
     * new cluster.
     * 
     */
    public Optional<Output<String>> masterFlavor() {
        return Optional.ofNullable(this.masterFlavor);
    }

    /**
     * Indicates whether to create a load balancer
     * for the master nodes. Changing this creates a new cluster.
     * 
     */
    @Import(name="masterLbEnabled")
    private @Nullable Output<Boolean> masterLbEnabled;

    /**
     * @return Indicates whether to create a load balancer
     * for the master nodes. Changing this creates a new cluster.
     * 
     */
    public Optional<Output<Boolean>> masterLbEnabled() {
        return Optional.ofNullable(this.masterLbEnabled);
    }

    /**
     * Indicates whether the provided labels should be
     * merged with cluster template labels. Changing this creates a new cluster.
     * 
     */
    @Import(name="mergeLabels")
    private @Nullable Output<Boolean> mergeLabels;

    /**
     * @return Indicates whether the provided labels should be
     * merged with cluster template labels. Changing this creates a new cluster.
     * 
     */
    public Optional<Output<Boolean>> mergeLabels() {
        return Optional.ofNullable(this.mergeLabels);
    }

    /**
     * The name of the cluster. Changing this creates a new
     * cluster.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the cluster. Changing this creates a new
     * cluster.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The number of nodes for the cluster.
     * 
     */
    @Import(name="nodeCount")
    private @Nullable Output<Integer> nodeCount;

    /**
     * @return The number of nodes for the cluster.
     * 
     */
    public Optional<Output<Integer>> nodeCount() {
        return Optional.ofNullable(this.nodeCount);
    }

    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * cluster.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * cluster.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private ClusterArgs() {}

    private ClusterArgs(ClusterArgs $) {
        this.clusterTemplateId = $.clusterTemplateId;
        this.createTimeout = $.createTimeout;
        this.discoveryUrl = $.discoveryUrl;
        this.dockerVolumeSize = $.dockerVolumeSize;
        this.fixedNetwork = $.fixedNetwork;
        this.fixedSubnet = $.fixedSubnet;
        this.flavor = $.flavor;
        this.floatingIpEnabled = $.floatingIpEnabled;
        this.keypair = $.keypair;
        this.labels = $.labels;
        this.masterCount = $.masterCount;
        this.masterFlavor = $.masterFlavor;
        this.masterLbEnabled = $.masterLbEnabled;
        this.mergeLabels = $.mergeLabels;
        this.name = $.name;
        this.nodeCount = $.nodeCount;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterArgs $;

        public Builder() {
            $ = new ClusterArgs();
        }

        public Builder(ClusterArgs defaults) {
            $ = new ClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterTemplateId The UUID of the V1 Container Infra cluster
         * template. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterTemplateId(@Nullable Output<String> clusterTemplateId) {
            $.clusterTemplateId = clusterTemplateId;
            return this;
        }

        /**
         * @param clusterTemplateId The UUID of the V1 Container Infra cluster
         * template. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder clusterTemplateId(String clusterTemplateId) {
            return clusterTemplateId(Output.of(clusterTemplateId));
        }

        /**
         * @param createTimeout The timeout (in minutes) for creating the
         * cluster. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder createTimeout(@Nullable Output<Integer> createTimeout) {
            $.createTimeout = createTimeout;
            return this;
        }

        /**
         * @param createTimeout The timeout (in minutes) for creating the
         * cluster. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder createTimeout(Integer createTimeout) {
            return createTimeout(Output.of(createTimeout));
        }

        /**
         * @param discoveryUrl The URL used for cluster node discovery.
         * Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder discoveryUrl(@Nullable Output<String> discoveryUrl) {
            $.discoveryUrl = discoveryUrl;
            return this;
        }

        /**
         * @param discoveryUrl The URL used for cluster node discovery.
         * Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder discoveryUrl(String discoveryUrl) {
            return discoveryUrl(Output.of(discoveryUrl));
        }

        /**
         * @param dockerVolumeSize The size (in GB) of the Docker volume.
         * Changing this creates a new cluster.
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
         * Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder dockerVolumeSize(Integer dockerVolumeSize) {
            return dockerVolumeSize(Output.of(dockerVolumeSize));
        }

        /**
         * @param fixedNetwork The fixed network that will be attached to the
         * cluster. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder fixedNetwork(@Nullable Output<String> fixedNetwork) {
            $.fixedNetwork = fixedNetwork;
            return this;
        }

        /**
         * @param fixedNetwork The fixed network that will be attached to the
         * cluster. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder fixedNetwork(String fixedNetwork) {
            return fixedNetwork(Output.of(fixedNetwork));
        }

        /**
         * @param fixedSubnet The fixed subnet that will be attached to the
         * cluster. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder fixedSubnet(@Nullable Output<String> fixedSubnet) {
            $.fixedSubnet = fixedSubnet;
            return this;
        }

        /**
         * @param fixedSubnet The fixed subnet that will be attached to the
         * cluster. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder fixedSubnet(String fixedSubnet) {
            return fixedSubnet(Output.of(fixedSubnet));
        }

        /**
         * @param flavor The flavor for the nodes of the cluster. Can be set via
         * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
         * cluster.
         * 
         * @return builder
         * 
         */
        public Builder flavor(@Nullable Output<String> flavor) {
            $.flavor = flavor;
            return this;
        }

        /**
         * @param flavor The flavor for the nodes of the cluster. Can be set via
         * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
         * cluster.
         * 
         * @return builder
         * 
         */
        public Builder flavor(String flavor) {
            return flavor(Output.of(flavor));
        }

        /**
         * @param floatingIpEnabled Indicates whether floating IP should be
         * created for every cluster node. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder floatingIpEnabled(@Nullable Output<Boolean> floatingIpEnabled) {
            $.floatingIpEnabled = floatingIpEnabled;
            return this;
        }

        /**
         * @param floatingIpEnabled Indicates whether floating IP should be
         * created for every cluster node. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder floatingIpEnabled(Boolean floatingIpEnabled) {
            return floatingIpEnabled(Output.of(floatingIpEnabled));
        }

        /**
         * @param keypair The name of the Compute service SSH keypair. Changing
         * this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder keypair(@Nullable Output<String> keypair) {
            $.keypair = keypair;
            return this;
        }

        /**
         * @param keypair The name of the Compute service SSH keypair. Changing
         * this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder keypair(String keypair) {
            return keypair(Output.of(keypair));
        }

        /**
         * @param labels The list of key value pairs representing additional
         * properties of the cluster. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels The list of key value pairs representing additional
         * properties of the cluster. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param masterCount The number of master nodes for the cluster.
         * Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder masterCount(@Nullable Output<Integer> masterCount) {
            $.masterCount = masterCount;
            return this;
        }

        /**
         * @param masterCount The number of master nodes for the cluster.
         * Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder masterCount(Integer masterCount) {
            return masterCount(Output.of(masterCount));
        }

        /**
         * @param masterFlavor The flavor for the master nodes. Can be set via
         * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
         * new cluster.
         * 
         * @return builder
         * 
         */
        public Builder masterFlavor(@Nullable Output<String> masterFlavor) {
            $.masterFlavor = masterFlavor;
            return this;
        }

        /**
         * @param masterFlavor The flavor for the master nodes. Can be set via
         * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
         * new cluster.
         * 
         * @return builder
         * 
         */
        public Builder masterFlavor(String masterFlavor) {
            return masterFlavor(Output.of(masterFlavor));
        }

        /**
         * @param masterLbEnabled Indicates whether to create a load balancer
         * for the master nodes. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder masterLbEnabled(@Nullable Output<Boolean> masterLbEnabled) {
            $.masterLbEnabled = masterLbEnabled;
            return this;
        }

        /**
         * @param masterLbEnabled Indicates whether to create a load balancer
         * for the master nodes. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder masterLbEnabled(Boolean masterLbEnabled) {
            return masterLbEnabled(Output.of(masterLbEnabled));
        }

        /**
         * @param mergeLabels Indicates whether the provided labels should be
         * merged with cluster template labels. Changing this creates a new cluster.
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
         * merged with cluster template labels. Changing this creates a new cluster.
         * 
         * @return builder
         * 
         */
        public Builder mergeLabels(Boolean mergeLabels) {
            return mergeLabels(Output.of(mergeLabels));
        }

        /**
         * @param name The name of the cluster. Changing this creates a new
         * cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the cluster. Changing this creates a new
         * cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nodeCount The number of nodes for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder nodeCount(@Nullable Output<Integer> nodeCount) {
            $.nodeCount = nodeCount;
            return this;
        }

        /**
         * @param nodeCount The number of nodes for the cluster.
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
         * cluster.
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
         * cluster.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public ClusterArgs build() {
            return $;
        }
    }

}
