// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available Shared File System snapshot.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const snapshot1 = openstack.sharedfilesystem.getSnapshot({
 *     name: "snapshot1",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/sharedfilesystem_snapshot_v2.html.markdown.
 */
export function getSnapshot(args?: GetSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:sharedfilesystem/getSnapshot:getSnapshot", {
        "description": args.description,
        "name": args.name,
        "region": args.region,
        "shareId": args.shareId,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshot.
 */
export interface GetSnapshotArgs {
    /**
     * The human-readable description of the snapshot.
     */
    readonly description?: string;
    /**
     * The name of the snapshot.
     */
    readonly name?: string;
    /**
     * The region in which to obtain the V2 Shared File System client.
     */
    readonly region?: string;
    /**
     * The UUID of the source share that was used to create the snapshot.
     */
    readonly shareId?: string;
    /**
     * A snapshot status filter. A valid value is `available`, `error`,
     * `creating`, `deleting`, `manageStarting`, `manageError`, `unmanageStarting`,
     * `unmanageError` or `errorDeleting`.
     */
    readonly status?: string;
}

/**
 * A collection of values returned by getSnapshot.
 */
export interface GetSnapshotResult {
    /**
     * See Argument Reference above.
     */
    readonly description: string;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * See Argument Reference above.
     */
    readonly projectId: string;
    readonly region: string;
    /**
     * The UUID of the source share that was used to create the snapshot.
     */
    readonly shareId: string;
    /**
     * The file system protocol of a share snapshot.
     */
    readonly shareProto: string;
    /**
     * The share snapshot size, in GBs.
     */
    readonly shareSize: number;
    /**
     * The snapshot size, in GBs.
     */
    readonly size: number;
    /**
     * See Argument Reference above.
     */
    readonly status: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
