// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of availability zones from OpenStack
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const zones = pulumi.output(openstack.compute.getAvailabilityZones({ async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/compute_availability_zones_v2.html.markdown.
 */
export function getAvailabilityZones(args?: GetAvailabilityZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetAvailabilityZonesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:compute/getAvailabilityZones:getAvailabilityZones", {
        "region": args.region,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getAvailabilityZones.
 */
export interface GetAvailabilityZonesArgs {
    /**
     * The `region` to fetch availability zones from, defaults to the provider's `region`
     */
    readonly region?: string;
    /**
     * The `state` of the availability zones to match, default ("available").
     */
    readonly state?: string;
}

/**
 * A collection of values returned by getAvailabilityZones.
 */
export interface GetAvailabilityZonesResult {
    /**
     * The names of the availability zones, ordered alphanumerically, that match the queried `state`
     */
    readonly names: string[];
    readonly region: string;
    readonly state?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
