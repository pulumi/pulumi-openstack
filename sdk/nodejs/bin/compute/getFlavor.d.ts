import * as pulumi from "@pulumi/pulumi";
/**
 * Use this data source to get the ID of an available OpenStack flavor.
 */
export declare function getFlavor(args?: GetFlavorArgs): Promise<GetFlavorResult>;
/**
 * A collection of arguments for invoking getFlavor.
 */
export interface GetFlavorArgs {
    /**
     * The exact amount of disk (in gigabytes).
     */
    readonly disk?: pulumi.Input<number>;
    /**
     * The minimum amount of disk (in gigabytes).
     */
    readonly minDisk?: pulumi.Input<number>;
    /**
     * The minimum amount of RAM (in megabytes).
     */
    readonly minRam?: pulumi.Input<number>;
    /**
     * The name of the flavor.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The exact amount of RAM (in megabytes).
     */
    readonly ram?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The `rx_tx_factor` of the flavor.
     */
    readonly rxTxFactor?: pulumi.Input<number>;
    /**
     * The amount of swap (in gigabytes).
     */
    readonly swap?: pulumi.Input<number>;
    /**
     * The amount of VCPUs.
     */
    readonly vcpus?: pulumi.Input<number>;
}
/**
 * A collection of values returned by getFlavor.
 */
export interface GetFlavorResult {
    /**
     * Whether the flavor is public or private.
     */
    readonly isPublic: boolean;
    readonly region: string;
}
