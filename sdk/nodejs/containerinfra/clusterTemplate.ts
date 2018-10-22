// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V1 Magnum cluster template resource within OpenStack.
 */
export class ClusterTemplate extends pulumi.CustomResource {
    /**
     * Get an existing ClusterTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterTemplateState): ClusterTemplate {
        return new ClusterTemplate(name, <any>state, { id });
    }

    public readonly apiserverPort: pulumi.Output<number | undefined>;
    public readonly clusterDistro: pulumi.Output<string>;
    public readonly coe: pulumi.Output<string>;
    public /*out*/ readonly createdAt: pulumi.Output<string>;
    public readonly dnsNameserver: pulumi.Output<string | undefined>;
    public readonly dockerStorageDriver: pulumi.Output<string | undefined>;
    public readonly dockerVolumeSize: pulumi.Output<number | undefined>;
    public readonly externalNetworkId: pulumi.Output<string | undefined>;
    public readonly fixedNetwork: pulumi.Output<string | undefined>;
    public readonly fixedSubnet: pulumi.Output<string | undefined>;
    public readonly flavor: pulumi.Output<string | undefined>;
    public readonly floatingIpEnabled: pulumi.Output<boolean | undefined>;
    public readonly httpProxy: pulumi.Output<string | undefined>;
    public readonly httpsProxy: pulumi.Output<string | undefined>;
    public readonly image: pulumi.Output<string>;
    public readonly insecureRegistry: pulumi.Output<string | undefined>;
    public readonly keypairId: pulumi.Output<string | undefined>;
    public readonly labels: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly masterFlavor: pulumi.Output<string | undefined>;
    public readonly masterLbEnabled: pulumi.Output<boolean | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly networkDriver: pulumi.Output<string>;
    public readonly noProxy: pulumi.Output<string | undefined>;
    public /*out*/ readonly projectId: pulumi.Output<string>;
    public readonly public: pulumi.Output<boolean | undefined>;
    public readonly region: pulumi.Output<string>;
    public readonly registryEnabled: pulumi.Output<boolean | undefined>;
    public readonly serverType: pulumi.Output<string>;
    public readonly tlsDisabled: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly updatedAt: pulumi.Output<string>;
    public /*out*/ readonly userId: pulumi.Output<string>;
    public readonly volumeDriver: pulumi.Output<string | undefined>;

    /**
     * Create a ClusterTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterTemplateArgs | ClusterTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ClusterTemplateState = argsOrState as ClusterTemplateState | undefined;
            inputs["apiserverPort"] = state ? state.apiserverPort : undefined;
            inputs["clusterDistro"] = state ? state.clusterDistro : undefined;
            inputs["coe"] = state ? state.coe : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["dnsNameserver"] = state ? state.dnsNameserver : undefined;
            inputs["dockerStorageDriver"] = state ? state.dockerStorageDriver : undefined;
            inputs["dockerVolumeSize"] = state ? state.dockerVolumeSize : undefined;
            inputs["externalNetworkId"] = state ? state.externalNetworkId : undefined;
            inputs["fixedNetwork"] = state ? state.fixedNetwork : undefined;
            inputs["fixedSubnet"] = state ? state.fixedSubnet : undefined;
            inputs["flavor"] = state ? state.flavor : undefined;
            inputs["floatingIpEnabled"] = state ? state.floatingIpEnabled : undefined;
            inputs["httpProxy"] = state ? state.httpProxy : undefined;
            inputs["httpsProxy"] = state ? state.httpsProxy : undefined;
            inputs["image"] = state ? state.image : undefined;
            inputs["insecureRegistry"] = state ? state.insecureRegistry : undefined;
            inputs["keypairId"] = state ? state.keypairId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["masterFlavor"] = state ? state.masterFlavor : undefined;
            inputs["masterLbEnabled"] = state ? state.masterLbEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkDriver"] = state ? state.networkDriver : undefined;
            inputs["noProxy"] = state ? state.noProxy : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["public"] = state ? state.public : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["registryEnabled"] = state ? state.registryEnabled : undefined;
            inputs["serverType"] = state ? state.serverType : undefined;
            inputs["tlsDisabled"] = state ? state.tlsDisabled : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
            inputs["userId"] = state ? state.userId : undefined;
            inputs["volumeDriver"] = state ? state.volumeDriver : undefined;
        } else {
            const args = argsOrState as ClusterTemplateArgs | undefined;
            if (!args || args.coe === undefined) {
                throw new Error("Missing required property 'coe'");
            }
            if (!args || args.image === undefined) {
                throw new Error("Missing required property 'image'");
            }
            inputs["apiserverPort"] = args ? args.apiserverPort : undefined;
            inputs["clusterDistro"] = args ? args.clusterDistro : undefined;
            inputs["coe"] = args ? args.coe : undefined;
            inputs["dnsNameserver"] = args ? args.dnsNameserver : undefined;
            inputs["dockerStorageDriver"] = args ? args.dockerStorageDriver : undefined;
            inputs["dockerVolumeSize"] = args ? args.dockerVolumeSize : undefined;
            inputs["externalNetworkId"] = args ? args.externalNetworkId : undefined;
            inputs["fixedNetwork"] = args ? args.fixedNetwork : undefined;
            inputs["fixedSubnet"] = args ? args.fixedSubnet : undefined;
            inputs["flavor"] = args ? args.flavor : undefined;
            inputs["floatingIpEnabled"] = args ? args.floatingIpEnabled : undefined;
            inputs["httpProxy"] = args ? args.httpProxy : undefined;
            inputs["httpsProxy"] = args ? args.httpsProxy : undefined;
            inputs["image"] = args ? args.image : undefined;
            inputs["insecureRegistry"] = args ? args.insecureRegistry : undefined;
            inputs["keypairId"] = args ? args.keypairId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["masterFlavor"] = args ? args.masterFlavor : undefined;
            inputs["masterLbEnabled"] = args ? args.masterLbEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkDriver"] = args ? args.networkDriver : undefined;
            inputs["noProxy"] = args ? args.noProxy : undefined;
            inputs["public"] = args ? args.public : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["registryEnabled"] = args ? args.registryEnabled : undefined;
            inputs["serverType"] = args ? args.serverType : undefined;
            inputs["tlsDisabled"] = args ? args.tlsDisabled : undefined;
            inputs["volumeDriver"] = args ? args.volumeDriver : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["projectId"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
            inputs["userId"] = undefined /*out*/;
        }
        super("openstack:containerinfra/clusterTemplate:ClusterTemplate", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterTemplate resources.
 */
export interface ClusterTemplateState {
    readonly apiserverPort?: pulumi.Input<number>;
    readonly clusterDistro?: pulumi.Input<string>;
    readonly coe?: pulumi.Input<string>;
    readonly createdAt?: pulumi.Input<string>;
    readonly dnsNameserver?: pulumi.Input<string>;
    readonly dockerStorageDriver?: pulumi.Input<string>;
    readonly dockerVolumeSize?: pulumi.Input<number>;
    readonly externalNetworkId?: pulumi.Input<string>;
    readonly fixedNetwork?: pulumi.Input<string>;
    readonly fixedSubnet?: pulumi.Input<string>;
    readonly flavor?: pulumi.Input<string>;
    readonly floatingIpEnabled?: pulumi.Input<boolean>;
    readonly httpProxy?: pulumi.Input<string>;
    readonly httpsProxy?: pulumi.Input<string>;
    readonly image?: pulumi.Input<string>;
    readonly insecureRegistry?: pulumi.Input<string>;
    readonly keypairId?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    readonly masterFlavor?: pulumi.Input<string>;
    readonly masterLbEnabled?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly networkDriver?: pulumi.Input<string>;
    readonly noProxy?: pulumi.Input<string>;
    readonly projectId?: pulumi.Input<string>;
    readonly public?: pulumi.Input<boolean>;
    readonly region?: pulumi.Input<string>;
    readonly registryEnabled?: pulumi.Input<boolean>;
    readonly serverType?: pulumi.Input<string>;
    readonly tlsDisabled?: pulumi.Input<boolean>;
    readonly updatedAt?: pulumi.Input<string>;
    readonly userId?: pulumi.Input<string>;
    readonly volumeDriver?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClusterTemplate resource.
 */
export interface ClusterTemplateArgs {
    readonly apiserverPort?: pulumi.Input<number>;
    readonly clusterDistro?: pulumi.Input<string>;
    readonly coe: pulumi.Input<string>;
    readonly dnsNameserver?: pulumi.Input<string>;
    readonly dockerStorageDriver?: pulumi.Input<string>;
    readonly dockerVolumeSize?: pulumi.Input<number>;
    readonly externalNetworkId?: pulumi.Input<string>;
    readonly fixedNetwork?: pulumi.Input<string>;
    readonly fixedSubnet?: pulumi.Input<string>;
    readonly flavor?: pulumi.Input<string>;
    readonly floatingIpEnabled?: pulumi.Input<boolean>;
    readonly httpProxy?: pulumi.Input<string>;
    readonly httpsProxy?: pulumi.Input<string>;
    readonly image: pulumi.Input<string>;
    readonly insecureRegistry?: pulumi.Input<string>;
    readonly keypairId?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    readonly masterFlavor?: pulumi.Input<string>;
    readonly masterLbEnabled?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly networkDriver?: pulumi.Input<string>;
    readonly noProxy?: pulumi.Input<string>;
    readonly public?: pulumi.Input<boolean>;
    readonly region?: pulumi.Input<string>;
    readonly registryEnabled?: pulumi.Input<boolean>;
    readonly serverType?: pulumi.Input<string>;
    readonly tlsDisabled?: pulumi.Input<boolean>;
    readonly volumeDriver?: pulumi.Input<string>;
}
