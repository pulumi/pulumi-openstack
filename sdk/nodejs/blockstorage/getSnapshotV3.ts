// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an existing snapshot.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const snapshot1 = openstack.blockstorage.getSnapshotV3({
 *     name: "snapshot_1",
 *     mostRecent: true,
 * });
 * ```
 */
export function getSnapshotV3(args?: GetSnapshotV3Args, opts?: pulumi.InvokeOptions): Promise<GetSnapshotV3Result> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:blockstorage/getSnapshotV3:getSnapshotV3", {
        "mostRecent": args.mostRecent,
        "name": args.name,
        "region": args.region,
        "status": args.status,
        "volumeId": args.volumeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshotV3.
 */
export interface GetSnapshotV3Args {
    /**
     * Pick the most recently created snapshot if there
     * are multiple results.
     */
    mostRecent?: boolean;
    /**
     * The name of the snapshot.
     */
    name?: string;
    /**
     * The region in which to obtain the V3 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The status of the snapshot.
     */
    status?: string;
    /**
     * The ID of the snapshot's volume.
     */
    volumeId?: string;
}

/**
 * A collection of values returned by getSnapshotV3.
 */
export interface GetSnapshotV3Result {
    /**
     * The snapshot's description.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The snapshot's metadata.
     */
    readonly metadata: {[key: string]: string};
    readonly mostRecent?: boolean;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The size of the snapshot.
     */
    readonly size: number;
    /**
     * See Argument Reference above.
     */
    readonly status: string;
    /**
     * See Argument Reference above.
     */
    readonly volumeId: string;
}
/**
 * Use this data source to get information about an existing snapshot.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const snapshot1 = openstack.blockstorage.getSnapshotV3({
 *     name: "snapshot_1",
 *     mostRecent: true,
 * });
 * ```
 */
export function getSnapshotV3Output(args?: GetSnapshotV3OutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSnapshotV3Result> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:blockstorage/getSnapshotV3:getSnapshotV3", {
        "mostRecent": args.mostRecent,
        "name": args.name,
        "region": args.region,
        "status": args.status,
        "volumeId": args.volumeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshotV3.
 */
export interface GetSnapshotV3OutputArgs {
    /**
     * Pick the most recently created snapshot if there
     * are multiple results.
     */
    mostRecent?: pulumi.Input<boolean>;
    /**
     * The name of the snapshot.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of the snapshot.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the snapshot's volume.
     */
    volumeId?: pulumi.Input<string>;
}
