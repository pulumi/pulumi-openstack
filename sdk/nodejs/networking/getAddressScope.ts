// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack address-scope.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const publicAddressscope = openstack.networking.getAddressScope({
 *     ipVersion: 4,
 *     name: "publicAddressscope",
 *     shared: true,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_addressscope_v2.html.markdown.
 */
export function getAddressScope(args?: GetAddressScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetAddressScopeResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:networking/getAddressScope:getAddressScope", {
        "ipVersion": args.ipVersion,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
        "shared": args.shared,
    }, opts);
}

/**
 * A collection of arguments for invoking getAddressScope.
 */
export interface GetAddressScopeArgs {
    /**
     * IP version.
     */
    readonly ipVersion?: number;
    /**
     * Name of the address-scope.
     */
    readonly name?: string;
    /**
     * The owner of the address-scope.
     */
    readonly projectId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve address-scopes. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * Indicates whether this address-scope is shared across
     * all projects.
     */
    readonly shared?: boolean;
}

/**
 * A collection of values returned by getAddressScope.
 */
export interface GetAddressScopeResult {
    /**
     * See Argument Reference above.
     */
    readonly ipVersion?: number;
    /**
     * See Argument Reference above.
     */
    readonly name?: string;
    /**
     * See Argument Reference above.
     */
    readonly projectId?: string;
    readonly region?: string;
    /**
     * See Argument Reference above.
     */
    readonly shared?: boolean;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
