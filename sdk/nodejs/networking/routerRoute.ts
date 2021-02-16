// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a routing entry on a OpenStack V2 router.
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
 * const int1 = new openstack.networking.RouterInterface("int_1", {
 *     routerId: router1.id,
 *     subnetId: subnet1.id,
 * });
 * const routerRoute1 = new openstack.networking.RouterRoute("router_route_1", {
 *     destinationCidr: "10.0.1.0/24",
 *     nextHop: "192.168.199.254",
 *     routerId: router1.id,
 * }, { dependsOn: [int1] });
 * ```
 * ## Notes
 *
 * The `nextHop` IP address must be directly reachable from the router at the ``openstack.networking.RouterRoute``
 * resource creation time.  You can ensure that by explicitly specifying a dependency on the ``openstack.networking.RouterInterface``
 * resource that connects the next hop to the router, as in the example above.
 *
 * ## Import
 *
 * Routing entries can be imported using a combined ID using the following format``<router_id>-route-<destination_cidr>-<next_hop>``
 *
 * ```sh
 *  $ pulumi import openstack:networking/routerRoute:RouterRoute router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
 * ```
 */
export class RouterRoute extends pulumi.CustomResource {
    /**
     * Get an existing RouterRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterRouteState, opts?: pulumi.CustomResourceOptions): RouterRoute {
        return new RouterRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/routerRoute:RouterRoute';

    /**
     * Returns true if the given object is an instance of RouterRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouterRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouterRoute.__pulumiType;
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
     * A networking client is needed to configure a routing entry on a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * ID of the router this routing entry belongs to. Changing
     * this creates a new routing entry.
     */
    public readonly routerId!: pulumi.Output<string>;

    /**
     * Create a RouterRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterRouteArgs | RouterRouteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterRouteState | undefined;
            inputs["destinationCidr"] = state ? state.destinationCidr : undefined;
            inputs["nextHop"] = state ? state.nextHop : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["routerId"] = state ? state.routerId : undefined;
        } else {
            const args = argsOrState as RouterRouteArgs | undefined;
            if ((!args || args.destinationCidr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationCidr'");
            }
            if ((!args || args.nextHop === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nextHop'");
            }
            if ((!args || args.routerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routerId'");
            }
            inputs["destinationCidr"] = args ? args.destinationCidr : undefined;
            inputs["nextHop"] = args ? args.nextHop : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["routerId"] = args ? args.routerId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RouterRoute.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterRoute resources.
 */
export interface RouterRouteState {
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
     * A networking client is needed to configure a routing entry on a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * ID of the router this routing entry belongs to. Changing
     * this creates a new routing entry.
     */
    readonly routerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouterRoute resource.
 */
export interface RouterRouteArgs {
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
     * A networking client is needed to configure a routing entry on a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * ID of the router this routing entry belongs to. Changing
     * this creates a new routing entry.
     */
    readonly routerId: pulumi.Input<string>;
}
