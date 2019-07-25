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
 * const snapshot1 = pulumi.output(openstack.blockstorage.getSnapshotV2({
 *     mostRecent: true,
 *     name: "snapshot_1",
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/blockstorage_snapshot_v2.html.markdown.
 */
export function getSnapshotV2(args?: GetSnapshotV2Args, opts?: pulumi.InvokeOptions): Promise<GetSnapshotV2Result> & GetSnapshotV2Result {
    args = args || {};
    const promise: Promise<GetSnapshotV2Result> = pulumi.runtime.invoke("openstack:blockstorage/getSnapshotV2:getSnapshotV2", {
        "mostRecent": args.mostRecent,
        "name": args.name,
        "region": args.region,
        "status": args.status,
        "volumeId": args.volumeId,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getSnapshotV2.
 */
export interface GetSnapshotV2Args {
    /**
     * Pick the most recently created snapshot if there
     * are multiple results.
     */
    readonly mostRecent?: boolean;
    /**
     * The name of the snapshot.
     */
    readonly name?: string;
    /**
     * The region in which to obtain the V2 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The status of the snapshot.
     */
    readonly status?: string;
    /**
     * The ID of the snapshot's volume.
     */
    readonly volumeId?: string;
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
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
