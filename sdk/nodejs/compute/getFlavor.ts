// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack flavor.
 */
export function getFlavor(args?: GetFlavorArgs, opts?: pulumi.InvokeOptions): Promise<GetFlavorResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:compute/getFlavor:getFlavor", {
        "description": args.description,
        "disk": args.disk,
        "flavorId": args.flavorId,
        "isPublic": args.isPublic,
        "minDisk": args.minDisk,
        "minRam": args.minRam,
        "name": args.name,
        "ram": args.ram,
        "region": args.region,
        "rxTxFactor": args.rxTxFactor,
        "swap": args.swap,
        "vcpus": args.vcpus,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlavor.
 */
export interface GetFlavorArgs {
    /**
     * The description of the flavor.
     */
    description?: string;
    /**
     * The exact amount of disk (in gigabytes).
     */
    disk?: number;
    /**
     * The ID of the flavor. Conflicts with the `name`,
     * `minRam` and `minDisk`
     */
    flavorId?: string;
    /**
     * The flavor visibility.
     */
    isPublic?: boolean;
    /**
     * The minimum amount of disk (in gigabytes). Conflicts
     * with the `flavorId`.
     */
    minDisk?: number;
    /**
     * The minimum amount of RAM (in megabytes). Conflicts
     * with the `flavorId`.
     */
    minRam?: number;
    /**
     * The name of the flavor. Conflicts with the `flavorId`.
     */
    name?: string;
    /**
     * The exact amount of RAM (in megabytes).
     */
    ram?: number;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The `rxTxFactor` of the flavor.
     */
    rxTxFactor?: number;
    /**
     * The amount of swap (in gigabytes).
     */
    swap?: number;
    /**
     * The amount of VCPUs.
     */
    vcpus?: number;
}

/**
 * A collection of values returned by getFlavor.
 */
export interface GetFlavorResult {
    readonly description?: string;
    readonly disk?: number;
    /**
     * Key/Value pairs of metadata for the flavor.
     */
    readonly extraSpecs: {[key: string]: any};
    readonly flavorId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isPublic?: boolean;
    readonly minDisk?: number;
    readonly minRam?: number;
    readonly name?: string;
    readonly ram?: number;
    readonly region: string;
    readonly rxTxFactor?: number;
    readonly swap?: number;
    readonly vcpus?: number;
}
/**
 * Use this data source to get the ID of an available OpenStack flavor.
 */
export function getFlavorOutput(args?: GetFlavorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlavorResult> {
    return pulumi.output(args).apply((a: any) => getFlavor(a, opts))
}

/**
 * A collection of arguments for invoking getFlavor.
 */
export interface GetFlavorOutputArgs {
    /**
     * The description of the flavor.
     */
    description?: pulumi.Input<string>;
    /**
     * The exact amount of disk (in gigabytes).
     */
    disk?: pulumi.Input<number>;
    /**
     * The ID of the flavor. Conflicts with the `name`,
     * `minRam` and `minDisk`
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The flavor visibility.
     */
    isPublic?: pulumi.Input<boolean>;
    /**
     * The minimum amount of disk (in gigabytes). Conflicts
     * with the `flavorId`.
     */
    minDisk?: pulumi.Input<number>;
    /**
     * The minimum amount of RAM (in megabytes). Conflicts
     * with the `flavorId`.
     */
    minRam?: pulumi.Input<number>;
    /**
     * The name of the flavor. Conflicts with the `flavorId`.
     */
    name?: pulumi.Input<string>;
    /**
     * The exact amount of RAM (in megabytes).
     */
    ram?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The `rxTxFactor` of the flavor.
     */
    rxTxFactor?: pulumi.Input<number>;
    /**
     * The amount of swap (in gigabytes).
     */
    swap?: pulumi.Input<number>;
    /**
     * The amount of VCPUs.
     */
    vcpus?: pulumi.Input<number>;
}
