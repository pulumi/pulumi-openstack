// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an OpenStack role.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const admin = openstack.identity.getRole({
 *     name: "admin",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_role_v3.html.markdown.
 */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:identity/getRole:getRole", {
        "domainId": args.domainId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleArgs {
    /**
     * The domain the role belongs to.
     */
    readonly domainId?: string;
    /**
     * The name of the role.
     */
    readonly name: string;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getRole.
 */
export interface GetRoleResult {
    /**
     * See Argument Reference above.
     */
    readonly domainId: string;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
