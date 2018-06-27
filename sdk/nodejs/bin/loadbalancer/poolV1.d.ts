import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V1 load balancer pool resource within OpenStack.
 */
export declare class PoolV1 extends pulumi.CustomResource {
    /**
     * Get an existing PoolV1 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolV1State): PoolV1;
    /**
     * The algorithm used to distribute load between the
     * members of the pool. The current specification supports 'ROUND_ROBIN' and
     * 'LEAST_CONNECTIONS' as valid values for this attribute.
     */
    readonly lbMethod: pulumi.Output<string>;
    /**
     * The backend load balancing provider. For example:
     * `haproxy`, `F5`, etc.
     */
    readonly lbProvider: pulumi.Output<string>;
    /**
     * A list of IDs of monitors to associate with the
     * pool.
     */
    readonly monitorIds: pulumi.Output<string[] | undefined>;
    /**
     * The name of the pool. Changing this updates the name of
     * the existing pool.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The protocol used by the pool members, you can use
     * either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
     */
    readonly protocol: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB pool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB pool.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The network on which the members of the pool will be
     * located. Only members that are on this network can be added to the pool.
     * Changing this creates a new pool.
     */
    readonly subnetId: pulumi.Output<string>;
    /**
     * The owner of the member. Required if admin wants to
     * create a pool member for another tenant. Changing this creates a new member.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Create a PoolV1 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PoolV1Args, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering PoolV1 resources.
 */
export interface PoolV1State {
    /**
     * The algorithm used to distribute load between the
     * members of the pool. The current specification supports 'ROUND_ROBIN' and
     * 'LEAST_CONNECTIONS' as valid values for this attribute.
     */
    readonly lbMethod?: pulumi.Input<string>;
    /**
     * The backend load balancing provider. For example:
     * `haproxy`, `F5`, etc.
     */
    readonly lbProvider?: pulumi.Input<string>;
    /**
     * A list of IDs of monitors to associate with the
     * pool.
     */
    readonly monitorIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the pool. Changing this updates the name of
     * the existing pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol used by the pool members, you can use
     * either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB pool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB pool.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The network on which the members of the pool will be
     * located. Only members that are on this network can be added to the pool.
     * Changing this creates a new pool.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * The owner of the member. Required if admin wants to
     * create a pool member for another tenant. Changing this creates a new member.
     */
    readonly tenantId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a PoolV1 resource.
 */
export interface PoolV1Args {
    /**
     * The algorithm used to distribute load between the
     * members of the pool. The current specification supports 'ROUND_ROBIN' and
     * 'LEAST_CONNECTIONS' as valid values for this attribute.
     */
    readonly lbMethod: pulumi.Input<string>;
    /**
     * The backend load balancing provider. For example:
     * `haproxy`, `F5`, etc.
     */
    readonly lbProvider?: pulumi.Input<string>;
    /**
     * A list of IDs of monitors to associate with the
     * pool.
     */
    readonly monitorIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the pool. Changing this updates the name of
     * the existing pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol used by the pool members, you can use
     * either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB pool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB pool.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The network on which the members of the pool will be
     * located. Only members that are on this network can be added to the pool.
     * Changing this creates a new pool.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * The owner of the member. Required if admin wants to
     * create a pool member for another tenant. Changing this creates a new member.
     */
    readonly tenantId?: pulumi.Input<string>;
}
