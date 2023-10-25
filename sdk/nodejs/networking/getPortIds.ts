// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of Openstack Port IDs matching the
 * specified criteria.
 */
export function getPortIds(args?: GetPortIdsArgs, opts?: pulumi.InvokeOptions): Promise<GetPortIdsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:networking/getPortIds:getPortIds", {
        "adminStateUp": args.adminStateUp,
        "description": args.description,
        "deviceId": args.deviceId,
        "deviceOwner": args.deviceOwner,
        "dnsName": args.dnsName,
        "fixedIp": args.fixedIp,
        "macAddress": args.macAddress,
        "name": args.name,
        "networkId": args.networkId,
        "projectId": args.projectId,
        "region": args.region,
        "securityGroupIds": args.securityGroupIds,
        "sortDirection": args.sortDirection,
        "sortKey": args.sortKey,
        "status": args.status,
        "tags": args.tags,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPortIds.
 */
export interface GetPortIdsArgs {
    /**
     * The administrative state of the port.
     */
    adminStateUp?: boolean;
    /**
     * Human-readable description of the port.
     */
    description?: string;
    /**
     * The ID of the device the port belongs to.
     */
    deviceId?: string;
    /**
     * The device owner of the port.
     */
    deviceOwner?: string;
    dnsName?: string;
    /**
     * The port IP address filter.
     */
    fixedIp?: string;
    /**
     * The MAC address of the port.
     */
    macAddress?: string;
    /**
     * The name of the port.
     */
    name?: string;
    /**
     * The ID of the network the port belongs to.
     */
    networkId?: string;
    /**
     * The owner of the port.
     */
    projectId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve port ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The list of port security group IDs to filter.
     */
    securityGroupIds?: string[];
    /**
     * Order the results in either `asc` or `desc`.
     * Defaults to none.
     */
    sortDirection?: string;
    /**
     * Sort ports based on a certain key. Defaults to none.
     */
    sortKey?: string;
    /**
     * The status of the port.
     */
    status?: string;
    /**
     * The list of port tags to filter.
     */
    tags?: string[];
    tenantId?: string;
}

/**
 * A collection of values returned by getPortIds.
 */
export interface GetPortIdsResult {
    readonly adminStateUp?: boolean;
    readonly description?: string;
    readonly deviceId?: string;
    readonly deviceOwner?: string;
    readonly dnsName?: string;
    readonly fixedIp?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly macAddress?: string;
    readonly name?: string;
    readonly networkId?: string;
    readonly projectId?: string;
    readonly region?: string;
    readonly securityGroupIds?: string[];
    readonly sortDirection?: string;
    readonly sortKey?: string;
    readonly status?: string;
    readonly tags?: string[];
    readonly tenantId?: string;
}
/**
 * Use this data source to get a list of Openstack Port IDs matching the
 * specified criteria.
 */
export function getPortIdsOutput(args?: GetPortIdsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPortIdsResult> {
    return pulumi.output(args).apply((a: any) => getPortIds(a, opts))
}

/**
 * A collection of arguments for invoking getPortIds.
 */
export interface GetPortIdsOutputArgs {
    /**
     * The administrative state of the port.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description of the port.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the device the port belongs to.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * The device owner of the port.
     */
    deviceOwner?: pulumi.Input<string>;
    dnsName?: pulumi.Input<string>;
    /**
     * The port IP address filter.
     */
    fixedIp?: pulumi.Input<string>;
    /**
     * The MAC address of the port.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The name of the port.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the network the port belongs to.
     */
    networkId?: pulumi.Input<string>;
    /**
     * The owner of the port.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve port ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The list of port security group IDs to filter.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Order the results in either `asc` or `desc`.
     * Defaults to none.
     */
    sortDirection?: pulumi.Input<string>;
    /**
     * Sort ports based on a certain key. Defaults to none.
     */
    sortKey?: pulumi.Input<string>;
    /**
     * The status of the port.
     */
    status?: pulumi.Input<string>;
    /**
     * The list of port tags to filter.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    tenantId?: pulumi.Input<string>;
}
