// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of Block Storage availability zones from OpenStack
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const zones = openstack.blockstorage.getAvailabilityZonesV3({});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/blockstorage_availability_zones_v3.html.markdown.
 */
export function getAvailabilityZonesV3(args?: GetAvailabilityZonesV3Args, opts?: pulumi.InvokeOptions): Promise<GetAvailabilityZonesV3Result> & GetAvailabilityZonesV3Result {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAvailabilityZonesV3Result> = pulumi.runtime.invoke("openstack:blockstorage/getAvailabilityZonesV3:getAvailabilityZonesV3", {
        "region": args.region,
        "state": args.state,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAvailabilityZonesV3.
 */
export interface GetAvailabilityZonesV3Args {
    /**
     * The region in which to obtain the Block Storage client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The `state` of the availability zones to match. Can
     * either be `available` or `unavailable`. Default is `available`.
     */
    readonly state?: string;
}

/**
 * A collection of values returned by getAvailabilityZonesV3.
 */
export interface GetAvailabilityZonesV3Result {
    /**
     * The names of the availability zones, ordered alphanumerically, that
     * match the queried `state`.
     */
    readonly names: string[];
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * See Argument Reference above.
     */
    readonly state?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
