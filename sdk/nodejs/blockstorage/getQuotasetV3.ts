// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the blockstorage quotaset v3 of an OpenStack project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const quota = openstack.blockstorage.getQuotasetV3({
 *     projectId: "2e367a3d29f94fd988e6ec54e305ec9d",
 * });
 * ```
 */
export function getQuotasetV3(args: GetQuotasetV3Args, opts?: pulumi.InvokeOptions): Promise<GetQuotasetV3Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:blockstorage/getQuotasetV3:getQuotasetV3", {
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getQuotasetV3.
 */
export interface GetQuotasetV3Args {
    /**
     * The id of the project to retrieve the quotaset.
     */
    projectId: string;
    /**
     * The region in which to obtain the V3 Blockstorage client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getQuotasetV3.
 */
export interface GetQuotasetV3Result {
    /**
     * The size (GB) of backups that are allowed.
     */
    readonly backupGigabytes: number;
    /**
     * The number of backups that are allowed.
     */
    readonly backups: number;
    /**
     * The size (GB) of volumes and snapshots that are allowed.
     */
    readonly gigabytes: number;
    /**
     * The number of groups that are allowed.
     */
    readonly groups: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The size (GB) of volumes that are allowed for each volume.
     */
    readonly perVolumeGigabytes: number;
    /**
     * See Argument Reference above.
     */
    readonly projectId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The number of snapshots that are allowed.
     */
    readonly snapshots: number;
    /**
     * Map with gigabytes_{volume_type}, snapshots_{volume_type}, volumes_{volume_type} for each volume type.
     */
    readonly volumeTypeQuota: {[key: string]: any};
    /**
     * The number of volumes that are allowed.
     */
    readonly volumes: number;
}
/**
 * Use this data source to get the blockstorage quotaset v3 of an OpenStack project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const quota = openstack.blockstorage.getQuotasetV3({
 *     projectId: "2e367a3d29f94fd988e6ec54e305ec9d",
 * });
 * ```
 */
export function getQuotasetV3Output(args: GetQuotasetV3OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQuotasetV3Result> {
    return pulumi.output(args).apply((a: any) => getQuotasetV3(a, opts))
}

/**
 * A collection of arguments for invoking getQuotasetV3.
 */
export interface GetQuotasetV3OutputArgs {
    /**
     * The id of the project to retrieve the quotaset.
     */
    projectId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Blockstorage client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
