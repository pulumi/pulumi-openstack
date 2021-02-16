// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
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
    public readonly clusterTemplateId!: pulumi.Output<string>;
    public /*out*/ readonly coeVersion!: pulumi.Output<string>;
    public /*out*/ readonly containerVersion!: pulumi.Output<string>;
    public readonly createTimeout!: pulumi.Output<number>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly discoveryUrl!: pulumi.Output<string>;
    public readonly dockerVolumeSize!: pulumi.Output<number>;
    public readonly fixedNetwork!: pulumi.Output<string>;
    public readonly fixedSubnet!: pulumi.Output<string>;
    public readonly flavor!: pulumi.Output<string>;
    public readonly floatingIpEnabled!: pulumi.Output<boolean>;
    public readonly keypair!: pulumi.Output<string>;
    public /*out*/ readonly kubeconfig!: pulumi.Output<outputs.containerinfra.ClusterKubeconfig>;
    public readonly labels!: pulumi.Output<{[key: string]: any}>;
    public /*out*/ readonly masterAddresses!: pulumi.Output<string[]>;
    public readonly masterCount!: pulumi.Output<number>;
    public readonly masterFlavor!: pulumi.Output<string>;
    public readonly mergeLabels!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly nodeAddresses!: pulumi.Output<string[]>;
    public readonly nodeCount!: pulumi.Output<number>;
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    public /*out*/ readonly stackId!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["apiAddress"] = state ? state.apiAddress : undefined;
            inputs["clusterTemplateId"] = state ? state.clusterTemplateId : undefined;
            inputs["coeVersion"] = state ? state.coeVersion : undefined;
            inputs["containerVersion"] = state ? state.containerVersion : undefined;
            inputs["createTimeout"] = state ? state.createTimeout : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["discoveryUrl"] = state ? state.discoveryUrl : undefined;
            inputs["dockerVolumeSize"] = state ? state.dockerVolumeSize : undefined;
            inputs["fixedNetwork"] = state ? state.fixedNetwork : undefined;
            inputs["fixedSubnet"] = state ? state.fixedSubnet : undefined;
            inputs["flavor"] = state ? state.flavor : undefined;
            inputs["floatingIpEnabled"] = state ? state.floatingIpEnabled : undefined;
            inputs["keypair"] = state ? state.keypair : undefined;
            inputs["kubeconfig"] = state ? state.kubeconfig : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["masterAddresses"] = state ? state.masterAddresses : undefined;
            inputs["masterCount"] = state ? state.masterCount : undefined;
            inputs["masterFlavor"] = state ? state.masterFlavor : undefined;
            inputs["mergeLabels"] = state ? state.mergeLabels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeAddresses"] = state ? state.nodeAddresses : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
            inputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.clusterTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterTemplateId'");
            }
            inputs["clusterTemplateId"] = args ? args.clusterTemplateId : undefined;
            inputs["createTimeout"] = args ? args.createTimeout : undefined;
            inputs["discoveryUrl"] = args ? args.discoveryUrl : undefined;
            inputs["dockerVolumeSize"] = args ? args.dockerVolumeSize : undefined;
            inputs["fixedNetwork"] = args ? args.fixedNetwork : undefined;
            inputs["fixedSubnet"] = args ? args.fixedSubnet : undefined;
            inputs["flavor"] = args ? args.flavor : undefined;
            inputs["floatingIpEnabled"] = args ? args.floatingIpEnabled : undefined;
            inputs["keypair"] = args ? args.keypair : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["masterCount"] = args ? args.masterCount : undefined;
            inputs["masterFlavor"] = args ? args.masterFlavor : undefined;
            inputs["mergeLabels"] = args ? args.mergeLabels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["apiAddress"] = undefined /*out*/;
            inputs["coeVersion"] = undefined /*out*/;
            inputs["containerVersion"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["kubeconfig"] = undefined /*out*/;
            inputs["masterAddresses"] = undefined /*out*/;
            inputs["nodeAddresses"] = undefined /*out*/;
            inputs["projectId"] = undefined /*out*/;
            inputs["stackId"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
            inputs["userId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    readonly apiAddress?: pulumi.Input<string>;
    readonly clusterTemplateId?: pulumi.Input<string>;
    readonly coeVersion?: pulumi.Input<string>;
    readonly containerVersion?: pulumi.Input<string>;
    readonly createTimeout?: pulumi.Input<number>;
    readonly createdAt?: pulumi.Input<string>;
    readonly discoveryUrl?: pulumi.Input<string>;
    readonly dockerVolumeSize?: pulumi.Input<number>;
    readonly fixedNetwork?: pulumi.Input<string>;
    readonly fixedSubnet?: pulumi.Input<string>;
    readonly flavor?: pulumi.Input<string>;
    readonly floatingIpEnabled?: pulumi.Input<boolean>;
    readonly keypair?: pulumi.Input<string>;
    readonly kubeconfig?: pulumi.Input<inputs.containerinfra.ClusterKubeconfig>;
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    readonly masterAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly masterCount?: pulumi.Input<number>;
    readonly masterFlavor?: pulumi.Input<string>;
    readonly mergeLabels?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly nodeAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly nodeCount?: pulumi.Input<number>;
    readonly projectId?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly stackId?: pulumi.Input<string>;
    readonly updatedAt?: pulumi.Input<string>;
    readonly userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    readonly clusterTemplateId: pulumi.Input<string>;
    readonly createTimeout?: pulumi.Input<number>;
    readonly discoveryUrl?: pulumi.Input<string>;
    readonly dockerVolumeSize?: pulumi.Input<number>;
    readonly fixedNetwork?: pulumi.Input<string>;
    readonly fixedSubnet?: pulumi.Input<string>;
    readonly flavor?: pulumi.Input<string>;
    readonly floatingIpEnabled?: pulumi.Input<boolean>;
    readonly keypair?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    readonly masterCount?: pulumi.Input<number>;
    readonly masterFlavor?: pulumi.Input<string>;
    readonly mergeLabels?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly nodeCount?: pulumi.Input<number>;
    readonly region?: pulumi.Input<string>;
}
