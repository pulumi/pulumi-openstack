// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the compute quotaset of an OpenStack project.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const quota = openstack.compute.getQuotaSetV2({
 *     projectId: "2e367a3d29f94fd988e6ec54e305ec9d",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getQuotaSetV2(args: GetQuotaSetV2Args, opts?: pulumi.InvokeOptions): Promise<GetQuotaSetV2Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:compute/getQuotaSetV2:getQuotaSetV2", {
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getQuotaSetV2.
 */
export interface GetQuotaSetV2Args {
    /**
     * The id of the project to retrieve the quotaset.
     */
    projectId: string;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getQuotaSetV2.
 */
export interface GetQuotaSetV2Result {
    /**
     * The number of allowed server cores.
     */
    readonly cores: number;
    /**
     * The number of allowed fixed IP addresses. Available until version 2.35.
     */
    readonly fixedIps: number;
    /**
     * The number of allowed floating IP addresses. Available until version 2.35.
     */
    readonly floatingIps: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The number of allowed bytes of content for each injected file. Available until version 2.56.
     */
    readonly injectedFileContentBytes: number;
    /**
     * The number of allowed bytes for each injected file path. Available until version 2.56.
     */
    readonly injectedFilePathBytes: number;
    /**
     * The number of allowed injected files. Available until version 2.56.
     */
    readonly injectedFiles: number;
    /**
     * The number of allowed servers.
     */
    readonly instances: number;
    /**
     * The number of allowed key pairs for each user.
     */
    readonly keyPairs: number;
    /**
     * The number of allowed metadata items for each server.
     */
    readonly metadataItems: number;
    /**
     * See Argument Reference above.
     */
    readonly projectId: string;
    /**
     * The amount of allowed server RAM, in MiB.
     */
    readonly ram: number;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The number of allowed rules for each security group. Available until version 2.35.
     */
    readonly securityGroupRules: number;
    /**
     * The number of allowed security groups. Available until version 2.35.
     */
    readonly securityGroups: number;
    /**
     * The number of allowed members for each server group.
     */
    readonly serverGroupMembers: number;
    /**
     * The number of allowed server groups.
     */
    readonly serverGroups: number;
}
/**
 * Use this data source to get the compute quotaset of an OpenStack project.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const quota = openstack.compute.getQuotaSetV2({
 *     projectId: "2e367a3d29f94fd988e6ec54e305ec9d",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getQuotaSetV2Output(args: GetQuotaSetV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQuotaSetV2Result> {
    return pulumi.output(args).apply((a: any) => getQuotaSetV2(a, opts))
}

/**
 * A collection of arguments for invoking getQuotaSetV2.
 */
export interface GetQuotaSetV2OutputArgs {
    /**
     * The id of the project to retrieve the quotaset.
     */
    projectId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
