// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network = pulumi.output(openstack.networking.getNetwork({
 *     name: "tf_test_network",
 * }, { async: true }));
 * ```
 */
export function getNetwork(args?: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:networking/getNetwork:getNetwork", {
        "description": args.description,
        "external": args.external,
        "matchingSubnetCidr": args.matchingSubnetCidr,
        "mtu": args.mtu,
        "name": args.name,
        "networkId": args.networkId,
        "region": args.region,
        "status": args.status,
        "tags": args.tags,
        "tenantId": args.tenantId,
        "transparentVlan": args.transparentVlan,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    /**
     * Human-readable description of the network.
     */
    readonly description?: string;
    /**
     * The external routing facility of the network.
     */
    readonly external?: boolean;
    /**
     * The CIDR of a subnet within the network.
     */
    readonly matchingSubnetCidr?: string;
    /**
     * The network MTU to filter. Available, when Neutron `net-mtu`
     * extension is enabled.
     */
    readonly mtu?: number;
    /**
     * The name of the network.
     */
    readonly name?: string;
    /**
     * The ID of the network.
     */
    readonly networkId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve networks ids. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The status of the network.
     */
    readonly status?: string;
    /**
     * The list of network tags to filter.
     */
    readonly tags?: string[];
    /**
     * The owner of the network.
     */
    readonly tenantId?: string;
    /**
     * The VLAN transparent attribute for the
     * network.
     */
    readonly transparentVlan?: boolean;
}

/**
 * A collection of values returned by getNetwork.
 */
export interface GetNetworkResult {
    /**
     * The administrative state of the network.
     */
    readonly adminStateUp: string;
    /**
     * The set of string tags applied on the network.
     */
    readonly allTags: string[];
    /**
     * The availability zone candidates for the network.
     */
    readonly availabilityZoneHints: string[];
    /**
     * See Argument Reference above.
     */
    readonly description?: string;
    /**
     * The network DNS domain. Available, when Neutron DNS extension
     * is enabled
     */
    readonly dnsDomain: string;
    /**
     * See Argument Reference above.
     */
    readonly external?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly matchingSubnetCidr?: string;
    /**
     * See Argument Reference above.
     */
    readonly mtu?: number;
    /**
     * See Argument Reference above.
     */
    readonly name?: string;
    readonly networkId?: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * Specifies whether the network resource can be accessed by any
     * tenant or not.
     */
    readonly shared: string;
    readonly status?: string;
    /**
     * A list of subnet IDs belonging to the network.
     */
    readonly subnets: string[];
    readonly tags?: string[];
    readonly tenantId?: string;
    /**
     * See Argument Reference above.
     */
    readonly transparentVlan?: boolean;
}
