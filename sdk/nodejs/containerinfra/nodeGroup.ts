// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V1 Magnum node group resource within OpenStack.
 *
 * ## Example Usage
 *
 * ### Create a Nodegroup
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const nodegroup1 = new openstack.containerinfra.NodeGroup("nodegroup_1", {
 *     name: "nodegroup_1",
 *     clusterId: "b9a45c5c-cd03-4958-82aa-b80bf93cb922",
 *     nodeCount: 5,
 * });
 * ```
 *
 * ## Attributes reference
 *
 * The following attributes are exported:
 *
 * * `region` - See Argument Reference above.
 * * `name` - See Argument Reference above.
 * * `projectId` - See Argument Reference above.
 * * `createdAt` - The time at which node group was created.
 * * `updatedAt` - The time at which node group was created.
 * * `dockerVolumeSize` - See Argument Reference above.
 * * `role` - See Argument Reference above.
 * * `imageId` - See Argument Reference above.
 * * `flavorId` - See Argument Reference above.
 * * `labels` - See Argument Reference above.
 * * `nodeCount` - See Argument Reference above.
 * * `minNodeCount` - See Argument Reference above.
 * * `maxNodeCount` - See Argument Reference above.
 * * `role` - See Argument Reference above.
 *
 * ## Import
 *
 * Node groups can be imported using the `id` (cluster_id/nodegroup_id), e.g.
 *
 * ```sh
 * $ pulumi import openstack:containerinfra/nodeGroup:NodeGroup nodegroup_1 b9a45c5c-cd03-4958-82aa-b80bf93cb922/ce0f9463-dd25-474b-9fe8-94de63e5e42b
 * ```
 */
export class NodeGroup extends pulumi.CustomResource {
    /**
     * Get an existing NodeGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeGroupState, opts?: pulumi.CustomResourceOptions): NodeGroup {
        return new NodeGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:containerinfra/nodeGroup:NodeGroup';

    /**
     * Returns true if the given object is an instance of NodeGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodeGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodeGroup.__pulumiType;
    }

    /**
     * The UUID of the V1 Container Infra cluster.
     * Changing this creates a new node group.
     */
    public readonly clusterId!: pulumi.Output<string>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     */
    public readonly dockerVolumeSize!: pulumi.Output<number>;
    /**
     * The flavor for the nodes of the node group. Can be set
     * via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * node group.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * The reference to an image that is used for nodes of the
     * node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
     * Changing this updates the image attribute of the existing node group.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The list of key value pairs representing additional
     * properties of the node group. Changing this creates a new node group.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any}>;
    /**
     * The maximum number of nodes for the node group.
     * Changing this update the maximum number of nodes of the node group.
     */
    public readonly maxNodeCount!: pulumi.Output<number | undefined>;
    /**
     * Indicates whether the provided labels should be
     * merged with cluster labels. Changing this creates a new nodegroup.
     */
    public readonly mergeLabels!: pulumi.Output<boolean | undefined>;
    /**
     * The minimum number of nodes for the node group.
     * Changing this update the minimum number of nodes of the node group.
     */
    public readonly minNodeCount!: pulumi.Output<number>;
    /**
     * The name of the node group. Changing this creates a new
     * node group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of nodes for the node group. Changing
     * this update the number of nodes of the node group.
     */
    public readonly nodeCount!: pulumi.Output<number | undefined>;
    /**
     * The project of the node group. Required if admin
     * wants to create a cluster in another project. Changing this creates a new
     * node group.
     */
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * node group.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The role of nodes in the node group. Changing this
     * creates a new node group.
     */
    public readonly role!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a NodeGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeGroupArgs | NodeGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodeGroupState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dockerVolumeSize"] = state ? state.dockerVolumeSize : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maxNodeCount"] = state ? state.maxNodeCount : undefined;
            resourceInputs["mergeLabels"] = state ? state.mergeLabels : undefined;
            resourceInputs["minNodeCount"] = state ? state.minNodeCount : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeCount"] = state ? state.nodeCount : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as NodeGroupArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["dockerVolumeSize"] = args ? args.dockerVolumeSize : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maxNodeCount"] = args ? args.maxNodeCount : undefined;
            resourceInputs["mergeLabels"] = args ? args.mergeLabels : undefined;
            resourceInputs["minNodeCount"] = args ? args.minNodeCount : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NodeGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodeGroup resources.
 */
export interface NodeGroupState {
    /**
     * The UUID of the V1 Container Infra cluster.
     * Changing this creates a new node group.
     */
    clusterId?: pulumi.Input<string>;
    createdAt?: pulumi.Input<string>;
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     */
    dockerVolumeSize?: pulumi.Input<number>;
    /**
     * The flavor for the nodes of the node group. Can be set
     * via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * node group.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The reference to an image that is used for nodes of the
     * node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
     * Changing this updates the image attribute of the existing node group.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The list of key value pairs representing additional
     * properties of the node group. Changing this creates a new node group.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The maximum number of nodes for the node group.
     * Changing this update the maximum number of nodes of the node group.
     */
    maxNodeCount?: pulumi.Input<number>;
    /**
     * Indicates whether the provided labels should be
     * merged with cluster labels. Changing this creates a new nodegroup.
     */
    mergeLabels?: pulumi.Input<boolean>;
    /**
     * The minimum number of nodes for the node group.
     * Changing this update the minimum number of nodes of the node group.
     */
    minNodeCount?: pulumi.Input<number>;
    /**
     * The name of the node group. Changing this creates a new
     * node group.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of nodes for the node group. Changing
     * this update the number of nodes of the node group.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * The project of the node group. Required if admin
     * wants to create a cluster in another project. Changing this creates a new
     * node group.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * node group.
     */
    region?: pulumi.Input<string>;
    /**
     * The role of nodes in the node group. Changing this
     * creates a new node group.
     */
    role?: pulumi.Input<string>;
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NodeGroup resource.
 */
export interface NodeGroupArgs {
    /**
     * The UUID of the V1 Container Infra cluster.
     * Changing this creates a new node group.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new node group.
     */
    dockerVolumeSize?: pulumi.Input<number>;
    /**
     * The flavor for the nodes of the node group. Can be set
     * via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * node group.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The reference to an image that is used for nodes of the
     * node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
     * Changing this updates the image attribute of the existing node group.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The list of key value pairs representing additional
     * properties of the node group. Changing this creates a new node group.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The maximum number of nodes for the node group.
     * Changing this update the maximum number of nodes of the node group.
     */
    maxNodeCount?: pulumi.Input<number>;
    /**
     * Indicates whether the provided labels should be
     * merged with cluster labels. Changing this creates a new nodegroup.
     */
    mergeLabels?: pulumi.Input<boolean>;
    /**
     * The minimum number of nodes for the node group.
     * Changing this update the minimum number of nodes of the node group.
     */
    minNodeCount?: pulumi.Input<number>;
    /**
     * The name of the node group. Changing this creates a new
     * node group.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of nodes for the node group. Changing
     * this update the number of nodes of the node group.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * node group.
     */
    region?: pulumi.Input<string>;
    /**
     * The role of nodes in the node group. Changing this
     * creates a new node group.
     */
    role?: pulumi.Input<string>;
}
