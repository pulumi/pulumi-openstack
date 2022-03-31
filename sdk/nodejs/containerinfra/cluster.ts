// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Create a Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const cluster1 = new openstack.containerinfra.Cluster("cluster_1", {
 *     clusterTemplateId: "b9a45c5c-cd03-4958-82aa-b80bf93cb922",
 *     keypair: "ssh_keypair",
 *     masterCount: 3,
 *     nodeCount: 5,
 * });
 * ```
 * ## Attributes reference
 *
 * The following attributes are exported:
 *
 * * `region` - See Argument Reference above.
 * * `name` - See Argument Reference above.
 * * `projectId` - See Argument Reference above.
 * * `createdAt` - The time at which cluster was created.
 * * `updatedAt` - The time at which cluster was created.
 * * `apiAddress` - COE API address.
 * * `coeVersion` - COE software version.
 * * `clusterTemplateId` - See Argument Reference above.
 * * `containerVersion` - Container software version.
 * * `createTimeout` - See Argument Reference above.
 * * `discoveryUrl` - See Argument Reference above.
 * * `dockerVolumeSize` - See Argument Reference above.
 * * `flavor` - See Argument Reference above.
 * * `masterFlavor` - See Argument Reference above.
 * * `keypair` - See Argument Reference above.
 * * `labels` - See Argument Reference above.
 * * `mergeLabels` - See Argument Reference above.
 * * `masterCount` - See Argument Reference above.
 * * `nodeCount` - See Argument Reference above.
 * * `fixedNetwork` - See Argument Reference above.
 * * `fixedSubnet` - See Argument Reference above.
 * * `floatingIpEnabled` - See Argument Reference above.
 * * `masterAddresses` - IP addresses of the master node of the cluster.
 * * `nodeAddresses` - IP addresses of the node of the cluster.
 * * `stackId` - UUID of the Orchestration service stack.
 * * `kubeconfig` - The Kubernetes cluster's credentials
 *   * `rawConfig` - The raw kubeconfig file
 *   * `host` - The cluster's API server URL
 *   * `clusterCaCertificate` - The cluster's CA certificate
 *   * `clientKey` - The client's RSA key
 *   * `clientCertificate` - The client's certificate
 *
 * ## Import
 *
 * Clusters can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:containerinfra/cluster:Cluster cluster_1 ce0f9463-dd25-474b-9fe8-94de63e5e42b
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:containerinfra/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    public /*out*/ readonly apiAddress!: pulumi.Output<string>;
    /**
     * The UUID of the V1 Container Infra cluster
     * template. Changing this creates a new cluster.
     */
    public readonly clusterTemplateId!: pulumi.Output<string>;
    public /*out*/ readonly coeVersion!: pulumi.Output<string>;
    public /*out*/ readonly containerVersion!: pulumi.Output<string>;
    /**
     * The timeout (in minutes) for creating the
     * cluster. Changing this creates a new cluster.
     */
    public readonly createTimeout!: pulumi.Output<number>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The URL used for cluster node discovery.
     * Changing this creates a new cluster.
     */
    public readonly discoveryUrl!: pulumi.Output<string>;
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new cluster.
     */
    public readonly dockerVolumeSize!: pulumi.Output<number>;
    /**
     * The fixed network that will be attached to the
     * cluster. Changing this creates a new cluster.
     */
    public readonly fixedNetwork!: pulumi.Output<string>;
    /**
     * The fixed subnet that will be attached to the
     * cluster. Changing this creates a new cluster.
     */
    public readonly fixedSubnet!: pulumi.Output<string>;
    /**
     * The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * cluster.
     */
    public readonly flavor!: pulumi.Output<string>;
    /**
     * Indicates whether floating IP should be
     * created for every cluster node. Changing this creates a new cluster.
     */
    public readonly floatingIpEnabled!: pulumi.Output<boolean>;
    /**
     * The name of the Compute service SSH keypair. Changing
     * this creates a new cluster.
     */
    public readonly keypair!: pulumi.Output<string>;
    public /*out*/ readonly kubeconfig!: pulumi.Output<{[key: string]: string}>;
    /**
     * The list of key value pairs representing additional
     * properties of the cluster. Changing this creates a new cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any}>;
    public /*out*/ readonly masterAddresses!: pulumi.Output<string[]>;
    /**
     * The number of master nodes for the cluster.
     * Changing this creates a new cluster.
     */
    public readonly masterCount!: pulumi.Output<number>;
    /**
     * The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
     * new cluster.
     */
    public readonly masterFlavor!: pulumi.Output<string>;
    /**
     * Indicates whether the provided labels should be
     * merged with cluster template labels. Changing this creates a new cluster.
     */
    public readonly mergeLabels!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the cluster. Changing this updates the name
     * of the existing cluster template.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly nodeAddresses!: pulumi.Output<string[]>;
    /**
     * The number of nodes for the cluster. Changing this
     * creates a new cluster.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * The project of the cluster. Required if admin wants
     * to create a cluster in another project. Changing this creates a new
     * cluster.
     */
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * cluster.
     */
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly stackId!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The user of the cluster. Required if admin wants to
     * create a cluster template for another user. Changing this creates a new
     * cluster.
     */
    public /*out*/ readonly userId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["apiAddress"] = state ? state.apiAddress : undefined;
            resourceInputs["clusterTemplateId"] = state ? state.clusterTemplateId : undefined;
            resourceInputs["coeVersion"] = state ? state.coeVersion : undefined;
            resourceInputs["containerVersion"] = state ? state.containerVersion : undefined;
            resourceInputs["createTimeout"] = state ? state.createTimeout : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["discoveryUrl"] = state ? state.discoveryUrl : undefined;
            resourceInputs["dockerVolumeSize"] = state ? state.dockerVolumeSize : undefined;
            resourceInputs["fixedNetwork"] = state ? state.fixedNetwork : undefined;
            resourceInputs["fixedSubnet"] = state ? state.fixedSubnet : undefined;
            resourceInputs["flavor"] = state ? state.flavor : undefined;
            resourceInputs["floatingIpEnabled"] = state ? state.floatingIpEnabled : undefined;
            resourceInputs["keypair"] = state ? state.keypair : undefined;
            resourceInputs["kubeconfig"] = state ? state.kubeconfig : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["masterAddresses"] = state ? state.masterAddresses : undefined;
            resourceInputs["masterCount"] = state ? state.masterCount : undefined;
            resourceInputs["masterFlavor"] = state ? state.masterFlavor : undefined;
            resourceInputs["mergeLabels"] = state ? state.mergeLabels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeAddresses"] = state ? state.nodeAddresses : undefined;
            resourceInputs["nodeCount"] = state ? state.nodeCount : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["stackId"] = state ? state.stackId : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.clusterTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterTemplateId'");
            }
            resourceInputs["clusterTemplateId"] = args ? args.clusterTemplateId : undefined;
            resourceInputs["createTimeout"] = args ? args.createTimeout : undefined;
            resourceInputs["discoveryUrl"] = args ? args.discoveryUrl : undefined;
            resourceInputs["dockerVolumeSize"] = args ? args.dockerVolumeSize : undefined;
            resourceInputs["fixedNetwork"] = args ? args.fixedNetwork : undefined;
            resourceInputs["fixedSubnet"] = args ? args.fixedSubnet : undefined;
            resourceInputs["flavor"] = args ? args.flavor : undefined;
            resourceInputs["floatingIpEnabled"] = args ? args.floatingIpEnabled : undefined;
            resourceInputs["keypair"] = args ? args.keypair : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["masterCount"] = args ? args.masterCount : undefined;
            resourceInputs["masterFlavor"] = args ? args.masterFlavor : undefined;
            resourceInputs["mergeLabels"] = args ? args.mergeLabels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeCount"] = args ? args.nodeCount : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["apiAddress"] = undefined /*out*/;
            resourceInputs["coeVersion"] = undefined /*out*/;
            resourceInputs["containerVersion"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["kubeconfig"] = undefined /*out*/;
            resourceInputs["masterAddresses"] = undefined /*out*/;
            resourceInputs["nodeAddresses"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["stackId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    apiAddress?: pulumi.Input<string>;
    /**
     * The UUID of the V1 Container Infra cluster
     * template. Changing this creates a new cluster.
     */
    clusterTemplateId?: pulumi.Input<string>;
    coeVersion?: pulumi.Input<string>;
    containerVersion?: pulumi.Input<string>;
    /**
     * The timeout (in minutes) for creating the
     * cluster. Changing this creates a new cluster.
     */
    createTimeout?: pulumi.Input<number>;
    createdAt?: pulumi.Input<string>;
    /**
     * The URL used for cluster node discovery.
     * Changing this creates a new cluster.
     */
    discoveryUrl?: pulumi.Input<string>;
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new cluster.
     */
    dockerVolumeSize?: pulumi.Input<number>;
    /**
     * The fixed network that will be attached to the
     * cluster. Changing this creates a new cluster.
     */
    fixedNetwork?: pulumi.Input<string>;
    /**
     * The fixed subnet that will be attached to the
     * cluster. Changing this creates a new cluster.
     */
    fixedSubnet?: pulumi.Input<string>;
    /**
     * The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * cluster.
     */
    flavor?: pulumi.Input<string>;
    /**
     * Indicates whether floating IP should be
     * created for every cluster node. Changing this creates a new cluster.
     */
    floatingIpEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Compute service SSH keypair. Changing
     * this creates a new cluster.
     */
    keypair?: pulumi.Input<string>;
    kubeconfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of key value pairs representing additional
     * properties of the cluster. Changing this creates a new cluster.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    masterAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of master nodes for the cluster.
     * Changing this creates a new cluster.
     */
    masterCount?: pulumi.Input<number>;
    /**
     * The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
     * new cluster.
     */
    masterFlavor?: pulumi.Input<string>;
    /**
     * Indicates whether the provided labels should be
     * merged with cluster template labels. Changing this creates a new cluster.
     */
    mergeLabels?: pulumi.Input<boolean>;
    /**
     * The name of the cluster. Changing this updates the name
     * of the existing cluster template.
     */
    name?: pulumi.Input<string>;
    nodeAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of nodes for the cluster. Changing this
     * creates a new cluster.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * The project of the cluster. Required if admin wants
     * to create a cluster in another project. Changing this creates a new
     * cluster.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * cluster.
     */
    region?: pulumi.Input<string>;
    stackId?: pulumi.Input<string>;
    updatedAt?: pulumi.Input<string>;
    /**
     * The user of the cluster. Required if admin wants to
     * create a cluster template for another user. Changing this creates a new
     * cluster.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The UUID of the V1 Container Infra cluster
     * template. Changing this creates a new cluster.
     */
    clusterTemplateId: pulumi.Input<string>;
    /**
     * The timeout (in minutes) for creating the
     * cluster. Changing this creates a new cluster.
     */
    createTimeout?: pulumi.Input<number>;
    /**
     * The URL used for cluster node discovery.
     * Changing this creates a new cluster.
     */
    discoveryUrl?: pulumi.Input<string>;
    /**
     * The size (in GB) of the Docker volume.
     * Changing this creates a new cluster.
     */
    dockerVolumeSize?: pulumi.Input<number>;
    /**
     * The fixed network that will be attached to the
     * cluster. Changing this creates a new cluster.
     */
    fixedNetwork?: pulumi.Input<string>;
    /**
     * The fixed subnet that will be attached to the
     * cluster. Changing this creates a new cluster.
     */
    fixedSubnet?: pulumi.Input<string>;
    /**
     * The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
     * cluster.
     */
    flavor?: pulumi.Input<string>;
    /**
     * Indicates whether floating IP should be
     * created for every cluster node. Changing this creates a new cluster.
     */
    floatingIpEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Compute service SSH keypair. Changing
     * this creates a new cluster.
     */
    keypair?: pulumi.Input<string>;
    /**
     * The list of key value pairs representing additional
     * properties of the cluster. Changing this creates a new cluster.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The number of master nodes for the cluster.
     * Changing this creates a new cluster.
     */
    masterCount?: pulumi.Input<number>;
    /**
     * The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this creates a
     * new cluster.
     */
    masterFlavor?: pulumi.Input<string>;
    /**
     * Indicates whether the provided labels should be
     * merged with cluster template labels. Changing this creates a new cluster.
     */
    mergeLabels?: pulumi.Input<boolean>;
    /**
     * The name of the cluster. Changing this updates the name
     * of the existing cluster template.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of nodes for the cluster. Changing this
     * creates a new cluster.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * cluster.
     */
    region?: pulumi.Input<string>;
}
