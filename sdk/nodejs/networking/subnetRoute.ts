// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a routing entry on a OpenStack V2 subnet.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const router1 = new openstack.networking.Router("router_1", {
 *     adminStateUp: true,
 * });
 * const network1 = new openstack.networking.Network("network_1", {
 *     adminStateUp: true,
 * });
 * const subnet1 = new openstack.networking.Subnet("subnet_1", {
 *     cidr: "192.168.199.0/24",
 *     ipVersion: 4,
 *     networkId: network1.id,
 * });
 * const subnetRoute1 = new openstack.networking.SubnetRoute("subnet_route_1", {
 *     destinationCidr: "10.0.1.0/24",
 *     nextHop: "192.168.199.254",
 *     subnetId: subnet1.id,
 * });
 * ```
 *
 * ## Import
 *
 * Routing entries can be imported using a combined ID using the following format``<subnet_id>-route-<destination_cidr>-<next_hop>``
 *
 * ```sh
 *  $ pulumi import openstack:networking/subnetRoute:SubnetRoute subnet_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
 * ```
 */
export class SubnetRoute extends pulumi.CustomResource {
    /**
     * Get an existing SubnetRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetRouteState, opts?: pulumi.CustomResourceOptions): SubnetRoute {
        return new SubnetRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/subnetRoute:SubnetRoute';

    /**
     * Returns true if the given object is an instance of SubnetRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubnetRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubnetRoute.__pulumiType;
    }

    /**
     * CIDR block to match on the packet’s destination IP. Changing
     * this creates a new routing entry.
     */
    public readonly destinationCidr!: pulumi.Output<string>;
    /**
     * IP address of the next hop gateway.  Changing
     * this creates a new routing entry.
     */
    public readonly nextHop!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * ID of the subnet this routing entry belongs to. Changing
     * this creates a new routing entry.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a SubnetRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetRouteArgs | SubnetRouteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubnetRouteState | undefined;
            inputs["destinationCidr"] = state ? state.destinationCidr : undefined;
            inputs["nextHop"] = state ? state.nextHop : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as SubnetRouteArgs | undefined;
            if ((!args || args.destinationCidr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationCidr'");
            }
            if ((!args || args.nextHop === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nextHop'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["destinationCidr"] = args ? args.destinationCidr : undefined;
            inputs["nextHop"] = args ? args.nextHop : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SubnetRoute.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubnetRoute resources.
 */
export interface SubnetRouteState {
    /**
     * CIDR block to match on the packet’s destination IP. Changing
     * this creates a new routing entry.
     */
    readonly destinationCidr?: pulumi.Input<string>;
    /**
     * IP address of the next hop gateway.  Changing
     * this creates a new routing entry.
     */
    readonly nextHop?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * ID of the subnet this routing entry belongs to. Changing
     * this creates a new routing entry.
     */
    readonly subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SubnetRoute resource.
 */
export interface SubnetRouteArgs {
    /**
     * CIDR block to match on the packet’s destination IP. Changing
     * this creates a new routing entry.
     */
    readonly destinationCidr: pulumi.Input<string>;
    /**
     * IP address of the next hop gateway.  Changing
     * this creates a new routing entry.
     */
    readonly nextHop: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * ID of the subnet this routing entry belongs to. Changing
     * this creates a new routing entry.
     */
    readonly subnetId: pulumi.Input<string>;
}
