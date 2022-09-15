// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the compute limits of an OpenStack project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const limits = pulumi.output(openstack.compute.getLimitsV2({
 *     projectId: "2e367a3d29f94fd988e6ec54e305ec9d",
 * }));
 * ```
 */
export function getLimitsV2(args: GetLimitsV2Args, opts?: pulumi.InvokeOptions): Promise<GetLimitsV2Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("openstack:compute/getLimitsV2:getLimitsV2", {
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getLimitsV2.
 */
export interface GetLimitsV2Args {
    /**
     * The id of the project to retrieve the limits.
     */
    projectId: string;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getLimitsV2.
 */
export interface GetLimitsV2Result {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The number of allowed metadata items for each image. Starting from version 2.39 this field is dropped from ‘os-limits’ response, because ‘image-metadata’ proxy API was deprecated. Available until version 2.38.
     */
    readonly maxImageMeta: number;
    /**
     * The number of allowed injected files for the tenant. Available until version 2.56.
     */
    readonly maxPersonality: number;
    /**
     * The number of allowed bytes of content for each injected file. Available until version 2.56.
     */
    readonly maxPersonalitySize: number;
    /**
     * The number of allowed rules for each security group. Available until version 2.35.
     */
    readonly maxSecurityGroupRules: number;
    /**
     * The number of allowed security groups for the tenant. Available until version 2.35.
     */
    readonly maxSecurityGroups: number;
    /**
     * The number of allowed members for each server group.
     */
    readonly maxServerGroupMembers: number;
    /**
     * The number of allowed server groups for the tenant.
     */
    readonly maxServerGroups: number;
    /**
     * The number of allowed server groups for the tenant.
     */
    readonly maxServerMeta: number;
    /**
     * The number of allowed server cores for the tenant.
     */
    readonly maxTotalCores: number;
    /**
     * The number of allowed floating IP addresses for each tenant. Available until version 2.35.
     */
    readonly maxTotalFloatingIps: number;
    /**
     * The number of allowed servers for the tenant.
     */
    readonly maxTotalInstances: number;
    /**
     * The number of allowed key pairs for the user.
     */
    readonly maxTotalKeypairs: number;
    /**
     * The number of allowed floating IP addresses for the tenant. Available until version 2.35.
     */
    readonly maxTotalRamSize: number;
    /**
     * See Argument Reference above.
     */
    readonly projectId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The number of used server cores in the tenant.
     */
    readonly totalCoresUsed: number;
    /**
     * The number of used floating IP addresses in the tenant.
     */
    readonly totalFloatingIpsUsed: number;
    /**
     * The number of used server cores in the tenant.
     */
    readonly totalInstancesUsed: number;
    /**
     * The amount of used server RAM in the tenant.
     */
    readonly totalRamUsed: number;
    /**
     * The number of used security groups in the tenant. Available until version 2.35.
     */
    readonly totalSecurityGroupsUsed: number;
    /**
     * The number of used server groups in each tenant.
     */
    readonly totalServerGroupsUsed: number;
}

export function getLimitsV2Output(args: GetLimitsV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLimitsV2Result> {
    return pulumi.output(args).apply(a => getLimitsV2(a, opts))
}

/**
 * A collection of arguments for invoking getLimitsV2.
 */
export interface GetLimitsV2OutputArgs {
    /**
     * The id of the project to retrieve the limits.
     */
    projectId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}