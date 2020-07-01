// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available Shared File System share network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const sharenetwork1 = pulumi.output(openstack.sharedfilesystem.getShareNetwork({
 *     name: "sharenetwork_1",
 * }, { async: true }));
 * ```
 */
export function getShareNetwork(args?: GetShareNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetShareNetworkResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:sharedfilesystem/getShareNetwork:getShareNetwork", {
        "description": args.description,
        "ipVersion": args.ipVersion,
        "name": args.name,
        "networkType": args.networkType,
        "neutronNetId": args.neutronNetId,
        "neutronSubnetId": args.neutronSubnetId,
        "region": args.region,
        "securityServiceId": args.securityServiceId,
        "segmentationId": args.segmentationId,
    }, opts);
}

/**
 * A collection of arguments for invoking getShareNetwork.
 */
export interface GetShareNetworkArgs {
    /**
     * The human-readable description of the share network.
     */
    readonly description?: string;
    /**
     * The IP version of the share network. Can either be 4 or 6.
     */
    readonly ipVersion?: number;
    /**
     * The name of the share network.
     */
    readonly name?: string;
    /**
     * The share network type. Can either be VLAN, VXLAN,
     * GRE, or flat.
     */
    readonly networkType?: string;
    /**
     * The neutron network UUID of the share network.
     */
    readonly neutronNetId?: string;
    /**
     * The neutron subnet UUID of the share network.
     */
    readonly neutronSubnetId?: string;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to read a share network. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The security service IDs associated with
     * the share network.
     */
    readonly securityServiceId?: string;
    /**
     * The share network segmentation ID.
     */
    readonly segmentationId?: number;
}

/**
 * A collection of values returned by getShareNetwork.
 */
export interface GetShareNetworkResult {
    /**
     * See Argument Reference above.
     */
    readonly cidr: string;
    /**
     * See Argument Reference above.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly ipVersion: number;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * See Argument Reference above.
     */
    readonly networkType: string;
    /**
     * See Argument Reference above.
     */
    readonly neutronNetId: string;
    /**
     * See Argument Reference above.
     */
    readonly neutronSubnetId: string;
    /**
     * The owner of the Share Network.
     */
    readonly projectId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * See Argument Reference above.
     */
    readonly securityServiceId?: string;
    /**
     * The list of security service IDs associated with
     * the share network.
     */
    readonly securityServiceIds: string[];
    /**
     * See Argument Reference above.
     */
    readonly segmentationId: number;
}
