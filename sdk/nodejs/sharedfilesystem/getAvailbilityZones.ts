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
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const zones = pulumi.output(openstack.sharedfilesystem.getAvailbilityZones({}));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/sharedfilesystem_availability_zones_v2.html.markdown.
 */
export function getAvailbilityZones(args?: GetAvailbilityZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetAvailbilityZonesResult> & GetAvailbilityZonesResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAvailbilityZonesResult> = pulumi.runtime.invoke("openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones", {
        "region": args.region,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAvailbilityZones.
 */
export interface GetAvailbilityZonesArgs {
    /**
     * The region in which to obtain the V2 Shared File System
     * client. If omitted, the `region` argument of the provider is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getAvailbilityZones.
 */
export interface GetAvailbilityZonesResult {
    /**
     * The names of the availability zones, ordered alphanumerically.
     */
    readonly names: string[];
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
