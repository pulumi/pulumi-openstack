// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack Magnum cluster
 * template.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const clustertemplate1 = openstack.containerinfra.getClusterTemplate({
 *     name: "clustertemplate_1",
 * });
 * ```
 */
export function getClusterTemplate(args: GetClusterTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:containerinfra/getClusterTemplate:getClusterTemplate", {
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterTemplate.
 */
export interface GetClusterTemplateArgs {
    /**
     * The name of the cluster template.
     */
    name: string;
    /**
     * The region in which to obtain the V1 Container Infra
     * client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getClusterTemplate.
 */
export interface GetClusterTemplateResult {
    /**
     * The API server port for the Container Orchestration
     * Engine for this cluster template.
     */
    readonly apiserverPort: number;
    /**
     * The distro for the cluster (fedora-atomic, coreos, etc.).
     */
    readonly clusterDistro: string;
    /**
     * The Container Orchestration Engine for this cluster template.
     */
    readonly coe: string;
    /**
     * The time at which cluster template was created.
     */
    readonly createdAt: string;
    /**
     * Address of the DNS nameserver that is used in nodes of the
     * cluster.
     */
    readonly dnsNameserver: string;
    /**
     * Docker storage driver. Changing this updates the
     * Docker storage driver of the existing cluster template.
     */
    readonly dockerStorageDriver: string;
    /**
     * The size (in GB) of the Docker volume.
     */
    readonly dockerVolumeSize: number;
    /**
     * The ID of the external network that will be used for
     * the cluster.
     */
    readonly externalNetworkId: string;
    /**
     * The fixed network that will be attached to the cluster.
     */
    readonly fixedNetwork: string;
    /**
     * =The fixed subnet that will be attached to the cluster.
     */
    readonly fixedSubnet: string;
    /**
     * The flavor for the nodes of the cluster.
     */
    readonly flavor: string;
    /**
     * Indicates whether created cluster should create IP
     * floating IP for every node or not.
     */
    readonly floatingIpEnabled: boolean;
    /**
     * Indicates whether the ClusterTemplate is hidden or not.
     */
    readonly hidden: boolean;
    /**
     * The address of a proxy for receiving all HTTP requests and
     * relay them.
     */
    readonly httpProxy: string;
    /**
     * The address of a proxy for receiving all HTTPS requests and
     * relay them.
     */
    readonly httpsProxy: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The reference to an image that is used for nodes of the cluster.
     */
    readonly image: string;
    /**
     * The insecure registry URL for the cluster template.
     */
    readonly insecureRegistry: string;
    /**
     * The name of the Compute service SSH keypair.
     */
    readonly keypairId: string;
    /**
     * The list of key value pairs representing additional properties
     * of the cluster template.
     */
    readonly labels: {[key: string]: string};
    /**
     * The flavor for the master nodes.
     */
    readonly masterFlavor: string;
    /**
     * Indicates whether created cluster should has a
     * loadbalancer for master nodes or not.
     */
    readonly masterLbEnabled: boolean;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * The name of the driver for the container network.
     */
    readonly networkDriver: string;
    /**
     * A comma-separated list of IP addresses that shouldn't be used in
     * the cluster.
     */
    readonly noProxy: string;
    /**
     * The project of the cluster template.
     */
    readonly projectId: string;
    /**
     * Indicates whether cluster template should be public.
     */
    readonly public: boolean;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * Indicates whether Docker registry is enabled in the
     * cluster.
     */
    readonly registryEnabled: boolean;
    /**
     * The server type for the cluster template.
     */
    readonly serverType: string;
    /**
     * Indicates whether the TLS should be disabled in the cluster.
     */
    readonly tlsDisabled: boolean;
    /**
     * The time at which cluster template was updated.
     */
    readonly updatedAt: string;
    /**
     * The user of the cluster template.
     */
    readonly userId: string;
    /**
     * The name of the driver that is used for the volumes of the
     * cluster nodes.
     */
    readonly volumeDriver: string;
}
/**
 * Use this data source to get the ID of an available OpenStack Magnum cluster
 * template.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const clustertemplate1 = openstack.containerinfra.getClusterTemplate({
 *     name: "clustertemplate_1",
 * });
 * ```
 */
export function getClusterTemplateOutput(args: GetClusterTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClusterTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:containerinfra/getClusterTemplate:getClusterTemplate", {
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterTemplate.
 */
export interface GetClusterTemplateOutputArgs {
    /**
     * The name of the cluster template.
     */
    name: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 Container Infra
     * client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
