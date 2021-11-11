// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an existing volume.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const volume1 = pulumi.output(openstack.blockstorage.getVolumeV2({
 *     name: "volume_1",
 * }));
 * ```
 */
export function getVolumeV2(args?: GetVolumeV2Args, opts?: pulumi.InvokeOptions): Promise<GetVolumeV2Result> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:blockstorage/getVolumeV2:getVolumeV2", {
        "bootable": args.bootable,
        "metadata": args.metadata,
        "name": args.name,
        "region": args.region,
        "status": args.status,
        "volumeType": args.volumeType,
    }, opts);
}

/**
 * A collection of arguments for invoking getVolumeV2.
 */
export interface GetVolumeV2Args {
    /**
     * Indicates if the volume is bootable.
     */
    bootable?: string;
    /**
     * Metadata key/value pairs associated with the volume.
     */
    metadata?: {[key: string]: any};
    /**
     * The name of the volume.
     */
    name?: string;
    /**
     * The region in which to obtain the V2 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The status of the volume.
     */
    status?: string;
    /**
     * The type of the volume.
     */
    volumeType?: string;
}

/**
 * A collection of values returned by getVolumeV2.
 */
export interface GetVolumeV2Result {
    /**
     * Indicates if the volume is bootable.
     */
    readonly bootable: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly metadata: {[key: string]: any};
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The size of the volume in GBs.
     */
    readonly size: number;
    /**
     * The ID of the volume from which the current volume was created.
     */
    readonly sourceVolumeId: string;
    /**
     * See Argument Reference above.
     */
    readonly status: string;
    /**
     * The type of the volume.
     */
    readonly volumeType: string;
}

export function getVolumeV2Output(args?: GetVolumeV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeV2Result> {
    return pulumi.output(args).apply(a => getVolumeV2(a, opts))
}

/**
 * A collection of arguments for invoking getVolumeV2.
 */
export interface GetVolumeV2OutputArgs {
    /**
     * Indicates if the volume is bootable.
     */
    bootable?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs associated with the volume.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the volume.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of the volume.
     */
    status?: pulumi.Input<string>;
    /**
     * The type of the volume.
     */
    volumeType?: pulumi.Input<string>;
}
