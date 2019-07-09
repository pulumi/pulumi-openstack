// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
 * const trunk1 = pulumi.output(openstack.networking.getTrunk({
 *     name: "trunk_1",
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_trunk_v2.html.markdown.
 */
export function getTrunk(args?: GetTrunkArgs, opts?: pulumi.InvokeOptions): Promise<GetTrunkResult> {
    args = args || {};
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
    readonly adminStateUp?: boolean;
    /**
     * Human-readable description of the trunk.
     */
    readonly description?: string;
    /**
     * The name of the trunk.
     */
    readonly name?: string;
    /**
     * The ID of the trunk parent port.
     */
    readonly portId?: string;
    /**
     * The owner of the trunk.
     */
    readonly projectId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve trunk ids. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The status of the trunk.
     */
    readonly status?: string;
    /**
     * The list of trunk tags to filter.
     */
    readonly tags?: string[];
    /**
     * The ID of the trunk.
     */
    readonly trunkId?: string;
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
    readonly subPorts: { portId: string, segmentationId: number, segmentationType: string }[];
    readonly tags?: string[];
    readonly trunkId?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
