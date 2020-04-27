// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack flavor.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const small = pulumi.output(openstack.compute.getFlavor({
 *     ram: 512,
 *     vcpus: 1,
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/compute_flavor_v2.html.markdown.
 */
export function getFlavor(args?: GetFlavorArgs, opts?: pulumi.InvokeOptions): Promise<GetFlavorResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:compute/getFlavor:getFlavor", {
        "disk": args.disk,
        "flavorId": args.flavorId,
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
     * The exact amount of disk (in gigabytes).
     */
    readonly disk?: number;
    /**
     * The ID of the flavor. Conflicts with the `name`,
     * `minRam` and `minDisk`
     */
    readonly flavorId?: string;
    /**
     * The minimum amount of disk (in gigabytes). Conflicts
     * with the `flavorId`.
     */
    readonly minDisk?: number;
    /**
     * The minimum amount of RAM (in megabytes). Conflicts
     * with the `flavorId`.
     */
    readonly minRam?: number;
    /**
     * The name of the flavor. Conflicts with the `flavorId`.
     */
    readonly name?: string;
    /**
     * The exact amount of RAM (in megabytes).
     */
    readonly ram?: number;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The `rxTxFactor` of the flavor.
     */
    readonly rxTxFactor?: number;
    /**
     * The amount of swap (in gigabytes).
     */
    readonly swap?: number;
    /**
     * The amount of VCPUs.
     */
    readonly vcpus?: number;
}

/**
 * A collection of values returned by getFlavor.
 */
export interface GetFlavorResult {
    readonly disk?: number;
    /**
     * Key/Value pairs of metadata for the flavor.
     */
    readonly extraSpecs: {[key: string]: any};
    readonly flavorId?: string;
    /**
     * Whether the flavor is public or private.
     */
    readonly isPublic: boolean;
    readonly minDisk?: number;
    readonly minRam?: number;
    readonly name?: string;
    readonly ram?: number;
    readonly region: string;
    readonly rxTxFactor?: number;
    readonly swap?: number;
    readonly vcpus?: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
