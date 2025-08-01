// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about hypervisors
 * by hostname.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const host01 = openstack.compute.getHypervisorV2({
 *     hostname: "host01",
 * });
 * ```
 */
export function getHypervisorV2(args?: GetHypervisorV2Args, opts?: pulumi.InvokeOptions): Promise<GetHypervisorV2Result> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:compute/getHypervisorV2:getHypervisorV2", {
        "hostname": args.hostname,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getHypervisorV2.
 */
export interface GetHypervisorV2Args {
    /**
     * The hostname of the hypervisor.
     */
    hostname?: string;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getHypervisorV2.
 */
export interface GetHypervisorV2Result {
    /**
     * The amount in GigaBytes of local storage the hypervisor can provide
     */
    readonly disk: number;
    /**
     * The IP address of the Hypervisor
     */
    readonly hostIp: string;
    /**
     * See Argument Reference above.
     */
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The number in MegaBytes of memory the hypervisor can provide
     */
    readonly memory: number;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The state of the hypervisor (`up` or `down`)
     */
    readonly state: string;
    /**
     * The status of the hypervisor (`enabled` or `disabled`)
     */
    readonly status: string;
    /**
     * The type of the hypervisor (example: `QEMU`)
     */
    readonly type: string;
    /**
     * The number of virtual CPUs the hypervisor can provide
     */
    readonly vcpus: number;
}
/**
 * Use this data source to get information about hypervisors
 * by hostname.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const host01 = openstack.compute.getHypervisorV2({
 *     hostname: "host01",
 * });
 * ```
 */
export function getHypervisorV2Output(args?: GetHypervisorV2OutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetHypervisorV2Result> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:compute/getHypervisorV2:getHypervisorV2", {
        "hostname": args.hostname,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getHypervisorV2.
 */
export interface GetHypervisorV2OutputArgs {
    /**
     * The hostname of the hypervisor.
     */
    hostname?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
