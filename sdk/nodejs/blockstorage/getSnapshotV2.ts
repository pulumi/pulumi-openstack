// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
 * const snapshot1 = openstack.blockstorage.getSnapshotV2({
 *     mostRecent: true,
 *     name: "snapshot_1",
 * });
 * ```
 */
export function getSnapshotV2(args?: GetSnapshotV2Args, opts?: pulumi.InvokeOptions): Promise<GetSnapshotV2Result> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:blockstorage/getSnapshotV2:getSnapshotV2", {
        "mostRecent": args.mostRecent,
        "name": args.name,
        "region": args.region,
        "status": args.status,
        "volumeId": args.volumeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshotV2.
 */
export interface GetSnapshotV2Args {
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
     * The region in which to obtain the V2 Block Storage
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
 * A collection of values returned by getSnapshotV2.
 */
export interface GetSnapshotV2Result {
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
    readonly metadata: {[key: string]: any};
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
 * const snapshot1 = openstack.blockstorage.getSnapshotV2({
 *     mostRecent: true,
 *     name: "snapshot_1",
 * });
 * ```
 */
export function getSnapshotV2Output(args?: GetSnapshotV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSnapshotV2Result> {
    return pulumi.output(args).apply((a: any) => getSnapshotV2(a, opts))
}

/**
 * A collection of arguments for invoking getSnapshotV2.
 */
export interface GetSnapshotV2OutputArgs {
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
     * The region in which to obtain the V2 Block Storage
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
