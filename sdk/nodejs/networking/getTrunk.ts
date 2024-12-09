// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack trunk.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const trunk1 = openstack.networking.getTrunk({
 *     name: "trunk_1",
 * });
 * ```
 */
export function getTrunk(args?: GetTrunkArgs, opts?: pulumi.InvokeOptions): Promise<GetTrunkResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:networking/getTrunk:getTrunk", {
        "adminStateUp": args.adminStateUp,
        "description": args.description,
        "name": args.name,
        "portId": args.portId,
        "projectId": args.projectId,
        "region": args.region,
        "status": args.status,
        "tags": args.tags,
        "trunkId": args.trunkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrunk.
 */
export interface GetTrunkArgs {
    /**
     * The administrative state of the trunk.
     */
    adminStateUp?: boolean;
    /**
     * Human-readable description of the trunk.
     */
    description?: string;
    /**
     * The name of the trunk.
     */
    name?: string;
    /**
     * The ID of the trunk parent port.
     */
    portId?: string;
    /**
     * The owner of the trunk.
     */
    projectId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve trunk ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The status of the trunk.
     */
    status?: string;
    /**
     * The list of trunk tags to filter.
     */
    tags?: string[];
    /**
     * The ID of the trunk.
     */
    trunkId?: string;
}

/**
 * A collection of values returned by getTrunk.
 */
export interface GetTrunkResult {
    readonly adminStateUp?: boolean;
    /**
     * The set of string tags applied on the trunk.
     */
    readonly allTags: string[];
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    /**
     * The ID of the trunk subport.
     */
    readonly portId?: string;
    readonly projectId: string;
    readonly region: string;
    readonly status?: string;
    /**
     * The set of the trunk subports. The structure of each subport is
     * described below.
     */
    readonly subPorts: outputs.networking.GetTrunkSubPort[];
    readonly tags?: string[];
    readonly trunkId?: string;
}
/**
 * Use this data source to get the ID of an available OpenStack trunk.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const trunk1 = openstack.networking.getTrunk({
 *     name: "trunk_1",
 * });
 * ```
 */
export function getTrunkOutput(args?: GetTrunkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTrunkResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:networking/getTrunk:getTrunk", {
        "adminStateUp": args.adminStateUp,
        "description": args.description,
        "name": args.name,
        "portId": args.portId,
        "projectId": args.projectId,
        "region": args.region,
        "status": args.status,
        "tags": args.tags,
        "trunkId": args.trunkId,
    }, opts);
}

/**
 * A collection of arguments for invoking getTrunk.
 */
export interface GetTrunkOutputArgs {
    /**
     * The administrative state of the trunk.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description of the trunk.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the trunk.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the trunk parent port.
     */
    portId?: pulumi.Input<string>;
    /**
     * The owner of the trunk.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve trunk ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of the trunk.
     */
    status?: pulumi.Input<string>;
    /**
     * The list of trunk tags to filter.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the trunk.
     */
    trunkId?: pulumi.Input<string>;
}
