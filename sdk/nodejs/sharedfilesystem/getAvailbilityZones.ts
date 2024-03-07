// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of Shared File System availability zones
 * from OpenStack
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const zones = openstack.sharedfilesystem.getAvailbilityZones({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAvailbilityZones(args?: GetAvailbilityZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetAvailbilityZonesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getAvailbilityZones.
 */
export interface GetAvailbilityZonesArgs {
    /**
     * The region in which to obtain the V2 Shared File System
     * client. If omitted, the `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getAvailbilityZones.
 */
export interface GetAvailbilityZonesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The names of the availability zones, ordered alphanumerically.
     */
    readonly names: string[];
    /**
     * See Argument Reference above.
     */
    readonly region: string;
}
/**
 * Use this data source to get a list of Shared File System availability zones
 * from OpenStack
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const zones = openstack.sharedfilesystem.getAvailbilityZones({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getAvailbilityZonesOutput(args?: GetAvailbilityZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAvailbilityZonesResult> {
    return pulumi.output(args).apply((a: any) => getAvailbilityZones(a, opts))
}

/**
 * A collection of arguments for invoking getAvailbilityZones.
 */
export interface GetAvailbilityZonesOutputArgs {
    /**
     * The region in which to obtain the V2 Shared File System
     * client. If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
