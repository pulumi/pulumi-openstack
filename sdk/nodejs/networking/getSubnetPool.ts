// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get the ID of an available OpenStack subnetpool.
 */
export function getSubnetPool(args?: GetSubnetPoolArgs): Promise<GetSubnetPoolResult> {
    args = args || {};
    return pulumi.runtime.invoke("openstack:networking/getSubnetPool:getSubnetPool", {
        "addressScopeId": args.addressScopeId,
        "defaultPrefixlen": args.defaultPrefixlen,
        "defaultQuota": args.defaultQuota,
        "description": args.description,
        "ipVersion": args.ipVersion,
        "isDefault": args.isDefault,
        "maxPrefixlen": args.maxPrefixlen,
        "minPrefixlen": args.minPrefixlen,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
        "shared": args.shared,
    });
}

/**
 * A collection of arguments for invoking getSubnetPool.
 */
export interface GetSubnetPoolArgs {
    /**
     * The Neutron address scope that subnetpools
     * is assigned to.
     */
    readonly addressScopeId?: string;
    /**
     * The size of the subnetpool default prefix
     * length.
     */
    readonly defaultPrefixlen?: number;
    /**
     * The per-project quota on the prefix space that
     * can be allocated from the subnetpool for project subnets.
     */
    readonly defaultQuota?: number;
    /**
     * The human-readable description for the subnetpool.
     */
    readonly description?: string;
    /**
     * The IP protocol version.
     */
    readonly ipVersion?: number;
    /**
     * Whether the subnetpool is default subnetpool or not.
     */
    readonly isDefault?: boolean;
    /**
     * The size of the subnetpool max prefix length.
     */
    readonly maxPrefixlen?: number;
    /**
     * The size of the subnetpool min prefix length.
     */
    readonly minPrefixlen?: number;
    /**
     * The name of the subnetpool.
     */
    readonly name?: string;
    /**
     * The owner of the subnetpool.
     */
    readonly projectId?: string;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to retrieve a subnetpool id. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * Whether this subnetpool is shared across all projects.
     */
    readonly shared?: boolean;
}

/**
 * A collection of values returned by getSubnetPool.
 */
export interface GetSubnetPoolResult {
    /**
     * See Argument Reference above.
     * * `ip_version` -The IP protocol version.
     */
    readonly addressScopeId: string;
    /**
     * The time at which subnetpool was created.
     */
    readonly createdAt: string;
    /**
     * See Argument Reference above.
     */
    readonly defaultPrefixlen: number;
    /**
     * See Argument Reference above.
     */
    readonly defaultQuota: number;
    /**
     * See Argument Reference above.
     */
    readonly description: string;
    readonly ipVersion: number;
    /**
     * See Argument Reference above.
     */
    readonly isDefault: boolean;
    /**
     * See Argument Reference above.
     */
    readonly maxPrefixlen: number;
    /**
     * See Argument Reference above.
     */
    readonly minPrefixlen: number;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * See Argument Reference above.
     */
    readonly prefixes: string[];
    /**
     * See Argument Reference above.
     */
    readonly projectId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The revision number of the subnetpool.
     */
    readonly revisionNumber: number;
    /**
     * See Argument Reference above.
     */
    readonly shared: boolean;
    /**
     * The time at which subnetpool was created.
     */
    readonly updatedAt: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
