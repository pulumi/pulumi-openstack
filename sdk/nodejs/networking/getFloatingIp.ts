// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack floating IP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const floatingip1 = openstack.networking.getFloatingIp({
 *     address: "192.168.0.4",
 * });
 * ```
 */
export function getFloatingIp(args?: GetFloatingIpArgs, opts?: pulumi.InvokeOptions): Promise<GetFloatingIpResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:networking/getFloatingIp:getFloatingIp", {
        "address": args.address,
        "description": args.description,
        "fixedIp": args.fixedIp,
        "pool": args.pool,
        "portId": args.portId,
        "region": args.region,
        "status": args.status,
        "tags": args.tags,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFloatingIp.
 */
export interface GetFloatingIpArgs {
    /**
     * The IP address of the floating IP.
     */
    address?: string;
    /**
     * Human-readable description of the floating IP.
     */
    description?: string;
    /**
     * The specific IP address of the internal port which should be associated with the floating IP.
     */
    fixedIp?: string;
    /**
     * The ID of the network from which the floating IP belongs to.
     */
    pool?: string;
    /**
     * The ID of the port the floating IP is attached.
     */
    portId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve floating IP ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * status of the floating IP (ACTIVE/DOWN).
     */
    status?: string;
    /**
     * The list of floating IP tags to filter.
     */
    tags?: string[];
    /**
     * The owner of the floating IP.
     */
    tenantId?: string;
}

/**
 * A collection of values returned by getFloatingIp.
 */
export interface GetFloatingIpResult {
    readonly address?: string;
    /**
     * A set of string tags applied on the floating IP.
     */
    readonly allTags: string[];
    readonly description?: string;
    /**
     * The floating IP DNS domain. Available, when Neutron DNS
     * extension is enabled.
     */
    readonly dnsDomain: string;
    /**
     * The floating IP DNS name. Available, when Neutron DNS extension
     * is enabled.
     */
    readonly dnsName: string;
    readonly fixedIp?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly pool?: string;
    readonly portId?: string;
    readonly region?: string;
    readonly status?: string;
    readonly tags?: string[];
    readonly tenantId?: string;
}
/**
 * Use this data source to get the ID of an available OpenStack floating IP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const floatingip1 = openstack.networking.getFloatingIp({
 *     address: "192.168.0.4",
 * });
 * ```
 */
export function getFloatingIpOutput(args?: GetFloatingIpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFloatingIpResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:networking/getFloatingIp:getFloatingIp", {
        "address": args.address,
        "description": args.description,
        "fixedIp": args.fixedIp,
        "pool": args.pool,
        "portId": args.portId,
        "region": args.region,
        "status": args.status,
        "tags": args.tags,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFloatingIp.
 */
export interface GetFloatingIpOutputArgs {
    /**
     * The IP address of the floating IP.
     */
    address?: pulumi.Input<string>;
    /**
     * Human-readable description of the floating IP.
     */
    description?: pulumi.Input<string>;
    /**
     * The specific IP address of the internal port which should be associated with the floating IP.
     */
    fixedIp?: pulumi.Input<string>;
    /**
     * The ID of the network from which the floating IP belongs to.
     */
    pool?: pulumi.Input<string>;
    /**
     * The ID of the port the floating IP is attached.
     */
    portId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve floating IP ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * status of the floating IP (ACTIVE/DOWN).
     */
    status?: pulumi.Input<string>;
    /**
     * The list of floating IP tags to filter.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the floating IP.
     */
    tenantId?: pulumi.Input<string>;
}
