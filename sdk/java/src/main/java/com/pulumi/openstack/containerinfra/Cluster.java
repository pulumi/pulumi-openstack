// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.containerinfra.ClusterArgs;
import com.pulumi.openstack.containerinfra.inputs.ClusterState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 Magnum cluster resource within OpenStack.
 * 
 * &gt; **Note:** All arguments including the `kubeconfig` computed attribute will be
 * stored in the raw state as plain-text. Read more about sensitive data in
 * state.
 * 
 * ## Example Usage
 * 
 * ### Create a Cluster
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.containerinfra.Cluster;
 * import com.pulumi.openstack.containerinfra.ClusterArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var cluster1 = new Cluster("cluster1", ClusterArgs.builder()
 *             .name("cluster_1")
 *             .clusterTemplateId("b9a45c5c-cd03-4958-82aa-b80bf93cb922")
 *             .masterCount(3)
 *             .nodeCount(5)
 *             .keypair("ssh_keypair")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Attributes reference
 * 
 * The following attributes are exported:
 * 
 * * `region` - See Argument Reference above.
 * * `name` - See Argument Reference above.
 * * `project_id` - See Argument Reference above.
 * * `created_at` - The time at which cluster was created.
 * * `updated_at` - The time at which cluster was created.
 * * `api_address` - COE API address.
 * * `coe_version` - COE software version.
 * * `cluster_template_id` - See Argument Reference above.
 * * `container_version` - Container software version.
 * * `create_timeout` - See Argument Reference above.
 * * `discovery_url` - See Argument Reference above.
 * * `docker_volume_size` - See Argument Reference above.
 * * `flavor` - See Argument Reference above.
 * * `master_flavor` - See Argument Reference above.
 * * `keypair` - See Argument Reference above.
 * * `labels` - See Argument Reference above.
 * * `merge_labels` - See Argument Reference above.
 * * `master_count` - See Argument Reference above.
 * * `node_count` - See Argument Reference above.
 * * `fixed_network` - See Argument Reference above.
 * * `fixed_subnet` - See Argument Reference above.
 * * `floating_ip_enabled` - See Argument Reference above.
 * * `master_addresses` - IP addresses of the master node of the cluster.
 * * `node_addresses` - IP addresses of the node of the cluster.
 * * `stack_id` - UUID of the Orchestration service stack.
 * * `kubeconfig` - The Kubernetes cluster&#39;s credentials
 *   * `raw_config` - The raw kubeconfig file
 *   * `host` - The cluster&#39;s API server URL
 *   * `cluster_ca_certificate` - The cluster&#39;s CA certificate
 *   * `client_key` - The client&#39;s RSA key
 *   * `client_certificate` - The client&#39;s certificate
 * 
 * ## Import
 * 
 * Clusters can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:containerinfra/cluster:Cluster cluster_1 ce0f9463-dd25-474b-9fe8-94de63e5e42b
 * ```
 * 
 */
@ResourceType(type="openstack:containerinfra/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    @Export(name="apiAddress", refs={String.class}, tree="[0]")
    private Output<String> apiAddress;

    public Output<String> apiAddress() {
        return this.apiAddress;
    }
    /**
     * The UUID of the V1 Container Infra cluster
     * template. Changing this creates a new cluster.
     * 
     */
    @Export(name="clusterTemplateId", refs={String.class}, tree="[0]")
    private Output<String> clusterTemplateId;

    /**
     * @return The UUID of the V1 Container Infra cluster
     * template. Changing this creates a new cluster.
     * 
     */
    public Output<String> clusterTemplateId() {
        return this.clusterTemplateId;
    }
    @Export(name="coeVersion", refs={String.class}, tree="[0]")
    private Output<String> coeVersion;

    public Output<String> coeVersion() {
        return this.coeVersion;
    }
    @Export(name="containerVersion", refs={String.class}, tree="[0]")
    private Output<String> containerVersion;

    public Output<String> containerVersion() {
        return this.containerVersion;
    }
    /**
     * The timeout (in minutes) for creating the
     * cluster. Changing this creates a new cluster.
     * 
     */
    @Export(name="createTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> createTimeout;

    /**
     * @return The timeout (in minutes) for creating the
     * cluster. Changing this creates a new cluster.
     * 
     */
    public Output<Integer> createTimeout() {
        return this.createTimeout;
    }
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The URL used for cluster node discovery.
     * Changing this creates a new cluster.
     * 
     */
    @Export(name="discoveryUrl", refs={String.class}, tree="[0]")
    private Output<String> discoveryUrl;

    /**
     * @return The URL used for cluster node discovery.
     * Changing this creates a new cluster.
     * 
     */
    public Output<String> discoveryUrl() {
        return this.discoveryUrl;
    }
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new cluster.
     * 
     */
    @Export(name="dockerVolumeSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> dockerVolumeSize;

    /**
     * @return The size (in GB) of the Docker volume.
     * Changing this creates a new cluster.
     * 
     */
    public Output<Integer> dockerVolumeSize() {
        return this.dockerVolumeSize;
    }
    /**
     * The fixed network that will be attached to the
     * cluster. Changing this creates a new cluster.
     * 
     */
    @Export(name="fixedNetwork", refs={String.class}, tree="[0]")
    private Output<String> fixedNetwork;

    /**
     * @return The fixed network that will be attached to the
     * cluster. Changing this creates a new cluster.
     * 
     */
    public Output<String> fixedNetwork() {
        return this.fixedNetwork;
    }
    /**
     * The fixed subnet that will be attached to the
     * cluster. Changing this creates a new cluster.
     * 
     */
    @Export(name="fixedSubnet", refs={String.class}, tree="[0]")
    private Output<String> fixedSubnet;

    /**
     * @return The fixed subnet that will be attached to the
     * cluster. Changing this creates a new cluster.
     * 
     */
    public Output<String> fixedSubnet() {
        return this.fixedSubnet;
    }
    /**
     * The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * cluster.
     * 
     */
    @Export(name="flavor", refs={String.class}, tree="[0]")
    private Output<String> flavor;

    /**
     * @return The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * cluster.
     * 
     */
    public Output<String> flavor() {
        return this.flavor;
    }
    /**
     * Indicates whether floating IP should be
     * created for every cluster node. Changing this creates a new cluster.
     * 
     */
    @Export(name="floatingIpEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> floatingIpEnabled;

    /**
     * @return Indicates whether floating IP should be
     * created for every cluster node. Changing this creates a new cluster.
     * 
     */
    public Output<Boolean> floatingIpEnabled() {
        return this.floatingIpEnabled;
    }
    /**
     * The name of the Compute service SSH keypair. Changing
     * this creates a new cluster.
     * 
     */
    @Export(name="keypair", refs={String.class}, tree="[0]")
    private Output<String> keypair;

    /**
     * @return The name of the Compute service SSH keypair. Changing
     * this creates a new cluster.
     * 
     */
    public Output<String> keypair() {
        return this.keypair;
    }
    @Export(name="kubeconfig", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> kubeconfig;

    public Output<Map<String,String>> kubeconfig() {
        return this.kubeconfig;
    }
    /**
     * The list of key value pairs representing additional
     * properties of the cluster. Changing this creates a new cluster.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> labels;

    /**
     * @return The list of key value pairs representing additional
     * properties of the cluster. Changing this creates a new cluster.
     * 
     */
    public Output<Map<String,Object>> labels() {
        return this.labels;
    }
    @Export(name="masterAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> masterAddresses;

    public Output<List<String>> masterAddresses() {
        return this.masterAddresses;
    }
    /**
     * The number of master nodes for the cluster.
     * Changing this creates a new cluster.
     * 
     */
    @Export(name="masterCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> masterCount;

    /**
     * @return The number of master nodes for the cluster.
     * Changing this creates a new cluster.
     * 
     */
    public Output<Integer> masterCount() {
        return this.masterCount;
    }
    /**
     * The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
     * new cluster.
     * 
     */
    @Export(name="masterFlavor", refs={String.class}, tree="[0]")
    private Output<String> masterFlavor;

    /**
     * @return The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
     * new cluster.
     * 
     */
    public Output<String> masterFlavor() {
        return this.masterFlavor;
    }
    /**
     * Indicates whether the provided labels should be
     * merged with cluster template labels. Changing this creates a new cluster.
     * 
     */
    @Export(name="mergeLabels", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mergeLabels;

    /**
     * @return Indicates whether the provided labels should be
     * merged with cluster template labels. Changing this creates a new cluster.
     * 
     */
    public Output<Optional<Boolean>> mergeLabels() {
        return Codegen.optional(this.mergeLabels);
    }
    /**
     * The name of the cluster. Changing this creates a new
     * cluster.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the cluster. Changing this creates a new
     * cluster.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="nodeAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> nodeAddresses;

    public Output<List<String>> nodeAddresses() {
        return this.nodeAddresses;
    }
    /**
     * The number of nodes for the cluster.
     * 
     */
    @Export(name="nodeCount", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> nodeCount;

    /**
     * @return The number of nodes for the cluster.
     * 
     */
    public Output<Optional<Integer>> nodeCount() {
        return Codegen.optional(this.nodeCount);
    }
    /**
     * The project of the cluster. Required if admin wants
     * to create a cluster in another project. Changing this creates a new
     * cluster.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project of the cluster. Required if admin wants
     * to create a cluster in another project. Changing this creates a new
     * cluster.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * cluster.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * cluster.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    @Export(name="stackId", refs={String.class}, tree="[0]")
    private Output<String> stackId;

    public Output<String> stackId() {
        return this.stackId;
    }
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * The user of the cluster. Required if admin wants to
     * create a cluster template for another user. Changing this creates a new
     * cluster.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return The user of the cluster. Required if admin wants to
     * create a cluster template for another user. Changing this creates a new
     * cluster.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:containerinfra/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:containerinfra/cluster:Cluster", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "kubeconfig"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
