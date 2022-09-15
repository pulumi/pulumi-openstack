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
 * ### Create a Cluster
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
 *             .clusterId(&#34;b9a45c5c-cd03-4958-82aa-b80bf93cb922&#34;)
 *             .nodeCount(5)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Attributes reference
 * 
 * The following attributes are exported:
 * 
 * * `region` - See Argument Reference above.
 * * `name` - See Argument Reference above.
 * * `project_id` - See Argument Reference above.
 * * `created_at` - The time at which cluster was created.
 * * `updated_at` - The time at which cluster was created.
 * * `docker_volume_size` - See Argument Reference above.
 * * `role` - See Argument Reference above.
 * * `image` - See Argument Reference above.
 * * `flavor` - See Argument Reference above.
 * * `labels` - See Argument Reference above.
 * * `node_count` - See Argument Reference above.
 * * `min_node_count` - See Argument Reference above.
 * * `max_node_count` - See Argument Reference above.
 * * `role` - See Argument Reference above.
 * 
 * ## Import
 * 
 * Node groups can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:containerinfra/nodeGroup:NodeGroup nodegroup_1 ce0f9463-dd25-474b-9fe8-94de63e5e42b
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
    @Export(name="clusterId", type=String.class, parameters={})
    private Output<String> clusterId;

    /**
     * @return The UUID of the V1 Container Infra cluster.
     * Changing this creates a new node group.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     * 
     */
    @Export(name="dockerVolumeSize", type=Integer.class, parameters={})
    private Output<Integer> dockerVolumeSize;

    /**
     * @return The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     * 
     */
    public Output<Integer> dockerVolumeSize() {
        return this.dockerVolumeSize;
    }
    @Export(name="flavorId", type=String.class, parameters={})
    private Output<String> flavorId;

    public Output<String> flavorId() {
        return this.flavorId;
    }
    @Export(name="imageId", type=String.class, parameters={})
    private Output<String> imageId;

    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * The list of key value pairs representing additional
     * properties of the node group. Changing this creates a new node group.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, Object.class})
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
    @Export(name="maxNodeCount", type=Integer.class, parameters={})
    private Output<Integer> maxNodeCount;

    /**
     * @return The maximum number of nodes for the node group.
     * Changing this update the maximum number of nodes of the node group.
     * 
     */
    public Output<Integer> maxNodeCount() {
        return this.maxNodeCount;
    }
    /**
     * Indicates whether the provided labels should be
     * merged with cluster labels. Changing this creates a new nodegroup.
     * 
     */
    @Export(name="mergeLabels", type=Boolean.class, parameters={})
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
    @Export(name="minNodeCount", type=Integer.class, parameters={})
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
    @Export(name="name", type=String.class, parameters={})
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
    @Export(name="nodeCount", type=Integer.class, parameters={})
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
    @Export(name="projectId", type=String.class, parameters={})
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
    @Export(name="region", type=String.class, parameters={})
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
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role of nodes in the node group. Changing this
     * creates a new node group.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    @Export(name="updatedAt", type=String.class, parameters={})
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