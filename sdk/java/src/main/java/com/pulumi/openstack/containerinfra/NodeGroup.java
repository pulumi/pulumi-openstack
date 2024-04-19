// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.containerinfra.NodeGroupArgs;
import com.pulumi.openstack.containerinfra.inputs.NodeGroupState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 Magnum node group resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ### Create a Nodegroup
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.containerinfra.NodeGroup;
 * import com.pulumi.openstack.containerinfra.NodeGroupArgs;
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
 *         var nodegroup1 = new NodeGroup(&#34;nodegroup1&#34;, NodeGroupArgs.builder()        
 *             .name(&#34;nodegroup_1&#34;)
 *             .clusterId(&#34;b9a45c5c-cd03-4958-82aa-b80bf93cb922&#34;)
 *             .nodeCount(5)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Attributes reference
 * 
 * The following attributes are exported:
 * 
 * * `region` - See Argument Reference above.
 * * `name` - See Argument Reference above.
 * * `project_id` - See Argument Reference above.
 * * `created_at` - The time at which node group was created.
 * * `updated_at` - The time at which node group was created.
 * * `docker_volume_size` - See Argument Reference above.
 * * `role` - See Argument Reference above.
 * * `image_id` - See Argument Reference above.
 * * `flavor_id` - See Argument Reference above.
 * * `labels` - See Argument Reference above.
 * * `node_count` - See Argument Reference above.
 * * `min_node_count` - See Argument Reference above.
 * * `max_node_count` - See Argument Reference above.
 * * `role` - See Argument Reference above.
 * 
 * ## Import
 * 
 * Node groups can be imported using the `id` (cluster_id/nodegroup_id), e.g.
 * 
 * ```sh
 * $ pulumi import openstack:containerinfra/nodeGroup:NodeGroup nodegroup_1 b9a45c5c-cd03-4958-82aa-b80bf93cb922/ce0f9463-dd25-474b-9fe8-94de63e5e42b
 * ```
 * 
 */
@ResourceType(type="openstack:containerinfra/nodeGroup:NodeGroup")
public class NodeGroup extends com.pulumi.resources.CustomResource {
    /**
     * The UUID of the V1 Container Infra cluster.
     * Changing this creates a new node group.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The UUID of the V1 Container Infra cluster.
     * Changing this creates a new node group.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     * 
     */
    @Export(name="dockerVolumeSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> dockerVolumeSize;

    /**
     * @return The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     * 
     */
    public Output<Integer> dockerVolumeSize() {
        return this.dockerVolumeSize;
    }
    /**
     * The flavor for the nodes of the node group. Can be set
     * via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * node group.
     * 
     */
    @Export(name="flavorId", refs={String.class}, tree="[0]")
    private Output<String> flavorId;

    /**
     * @return The flavor for the nodes of the node group. Can be set
     * via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * node group.
     * 
     */
    public Output<String> flavorId() {
        return this.flavorId;
    }
    /**
     * The reference to an image that is used for nodes of the
     * node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
     * Changing this updates the image attribute of the existing node group.
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output<String> imageId;

    /**
     * @return The reference to an image that is used for nodes of the
     * node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
     * Changing this updates the image attribute of the existing node group.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * The list of key value pairs representing additional
     * properties of the node group. Changing this creates a new node group.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> labels;

    /**
     * @return The list of key value pairs representing additional
     * properties of the node group. Changing this creates a new node group.
     * 
     */
    public Output<Map<String,Object>> labels() {
        return this.labels;
    }
    /**
     * The maximum number of nodes for the node group.
     * Changing this update the maximum number of nodes of the node group.
     * 
     */
    @Export(name="maxNodeCount", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxNodeCount;

    /**
     * @return The maximum number of nodes for the node group.
     * Changing this update the maximum number of nodes of the node group.
     * 
     */
    public Output<Optional<Integer>> maxNodeCount() {
        return Codegen.optional(this.maxNodeCount);
    }
    /**
     * Indicates whether the provided labels should be
     * merged with cluster labels. Changing this creates a new nodegroup.
     * 
     */
    @Export(name="mergeLabels", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mergeLabels;

    /**
     * @return Indicates whether the provided labels should be
     * merged with cluster labels. Changing this creates a new nodegroup.
     * 
     */
    public Output<Optional<Boolean>> mergeLabels() {
        return Codegen.optional(this.mergeLabels);
    }
    /**
     * The minimum number of nodes for the node group.
     * Changing this update the minimum number of nodes of the node group.
     * 
     */
    @Export(name="minNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> minNodeCount;

    /**
     * @return The minimum number of nodes for the node group.
     * Changing this update the minimum number of nodes of the node group.
     * 
     */
    public Output<Integer> minNodeCount() {
        return this.minNodeCount;
    }
    /**
     * The name of the node group. Changing this creates a new
     * node group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the node group. Changing this creates a new
     * node group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The number of nodes for the node group. Changing
     * this update the number of nodes of the node group.
     * 
     */
    @Export(name="nodeCount", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> nodeCount;

    /**
     * @return The number of nodes for the node group. Changing
     * this update the number of nodes of the node group.
     * 
     */
    public Output<Optional<Integer>> nodeCount() {
        return Codegen.optional(this.nodeCount);
    }
    /**
     * The project of the node group. Required if admin
     * wants to create a cluster in another project. Changing this creates a new
     * node group.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The project of the node group. Required if admin
     * wants to create a cluster in another project. Changing this creates a new
     * node group.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * node group.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * node group.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The role of nodes in the node group. Changing this
     * creates a new node group.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The role of nodes in the node group. Changing this
     * creates a new node group.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NodeGroup(String name) {
        this(name, NodeGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NodeGroup(String name, NodeGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NodeGroup(String name, NodeGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:containerinfra/nodeGroup:NodeGroup", name, args == null ? NodeGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NodeGroup(String name, Output<String> id, @Nullable NodeGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:containerinfra/nodeGroup:NodeGroup", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static NodeGroup get(String name, Output<String> id, @Nullable NodeGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NodeGroup(name, id, state, options);
    }
}
