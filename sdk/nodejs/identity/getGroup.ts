// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an OpenStack group.
 *
 * Note: This usually requires admin privileges.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const admins = pulumi.output(openstack.identity.getGroup({
 *     name: "admins",
 * }, { async: true }));
 * ```
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:identity/getGroup:getGroup", {
        "domainId": args.domainId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The domain the group belongs to.
     */
    readonly domainId?: string;
    /**
     * The name of the group.
     */
    readonly name: string;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * A description of the group.
     */
    readonly description: string;
    /**
     * See Argument Reference above.
     */
    readonly domainId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
}
