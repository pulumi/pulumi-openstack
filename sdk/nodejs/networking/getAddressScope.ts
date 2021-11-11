// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack address-scope.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const publicAddressscope = pulumi.output(openstack.networking.getAddressScope({
 *     ipVersion: 4,
 *     name: "public_addressscope",
 *     shared: true,
 * }));
 * ```
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
    ipVersion?: number;
    /**
     * Name of the address-scope.
     */
    name?: string;
    /**
     * The owner of the address-scope.
     */
    projectId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve address-scopes. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * Indicates whether this address-scope is shared across
     * all projects.
     */
    shared?: boolean;
}

/**
 * A collection of values returned by getAddressScope.
 */
export interface GetAddressScopeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
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
}

export function getAddressScopeOutput(args?: GetAddressScopeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAddressScopeResult> {
    return pulumi.output(args).apply(a => getAddressScope(a, opts))
}

/**
 * A collection of arguments for invoking getAddressScope.
 */
export interface GetAddressScopeOutputArgs {
    /**
     * IP version.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Name of the address-scope.
     */
    name?: pulumi.Input<string>;
    /**
     * The owner of the address-scope.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve address-scopes. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * Indicates whether this address-scope is shared across
     * all projects.
     */
    shared?: pulumi.Input<boolean>;
}
