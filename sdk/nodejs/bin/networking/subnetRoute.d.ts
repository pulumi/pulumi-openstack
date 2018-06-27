import * as pulumi from "@pulumi/pulumi";
/**
 * Creates a routing entry on a OpenStack V2 subnet.
 */
export declare class SubnetRoute extends pulumi.CustomResource {
    /**
     * Get an existing SubnetRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetRouteState): SubnetRoute;
    /**
     * CIDR block to match on the packet’s destination IP. Changing
     * this creates a new routing entry.
     */
    readonly destinationCidr: pulumi.Output<string>;
    /**
     * IP address of the next hop gateway.  Changing
     * this creates a new routing entry.
     */
    readonly nextHop: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     */
    readonly region: pulumi.Output<string>;
    /**
     * ID of the subnet this routing entry belongs to. Changing
     * this creates a new routing entry.
     */
    readonly subnetId: pulumi.Output<string>;
    /**
     * Create a SubnetRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetRouteArgs, opts?: pulumi.ResourceOptions);
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
