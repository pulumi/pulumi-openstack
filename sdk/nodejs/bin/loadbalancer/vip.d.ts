import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V1 load balancer vip resource within OpenStack.
 */
export declare class Vip extends pulumi.CustomResource {
    /**
     * Get an existing Vip resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VipState): Vip;
    /**
     * The IP address of the vip. Changing this creates a new
     * vip.
     */
    readonly address: pulumi.Output<string>;
    /**
     * The administrative state of the vip.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing vip.
     */
    readonly adminStateUp: pulumi.Output<boolean>;
    /**
     * The maximum number of connections allowed for the
     * vip. Default is -1, meaning no limit. Changing this updates the conn_limit
     * of the existing vip.
     */
    readonly connLimit: pulumi.Output<number>;
    /**
     * Human-readable description for the vip. Changing
     * this updates the description of the existing vip.
     */
    readonly description: pulumi.Output<string>;
    /**
     * A *Networking* Floating IP that will be associated
     * with the vip. The Floating IP must be provisioned already.
     */
    readonly floatingIp: pulumi.Output<string | undefined>;
    /**
     * The name of the vip. Changing this updates the name of
     * the existing vip.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Omit this field to prevent session persistence.
     * The persistence object structure is documented below. Changing this updates
     * the persistence of the existing vip.
     */
    readonly persistence: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The ID of the pool with which the vip is associated.
     * Changing this updates the pool_id of the existing vip.
     */
    readonly poolId: pulumi.Output<string>;
    /**
     * The port on which to listen for client traffic. Changing
     * this creates a new vip.
     */
    readonly port: pulumi.Output<number>;
    /**
     * Port UUID for this VIP at associated floating IP (if any).
     */
    readonly portId: pulumi.Output<string>;
    /**
     * The protocol - can be either 'TCP, 'HTTP', or
     * HTTPS'. Changing this creates a new vip.
     */
    readonly protocol: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VIP. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * VIP.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The network on which to allocate the vip's address. A
     * tenant can only create vips on networks authorized by policy (e.g. networks
     * that belong to them or networks that are shared). Changing this creates a
     * new vip.
     */
    readonly subnetId: pulumi.Output<string>;
    /**
     * The owner of the vip. Required if admin wants to
     * create a vip member for another tenant. Changing this creates a new vip.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Create a Vip resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VipArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Vip resources.
 */
export interface VipState {
    /**
     * The IP address of the vip. Changing this creates a new
     * vip.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The administrative state of the vip.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing vip.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The maximum number of connections allowed for the
     * vip. Default is -1, meaning no limit. Changing this updates the conn_limit
     * of the existing vip.
     */
    readonly connLimit?: pulumi.Input<number>;
    /**
     * Human-readable description for the vip. Changing
     * this updates the description of the existing vip.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A *Networking* Floating IP that will be associated
     * with the vip. The Floating IP must be provisioned already.
     */
    readonly floatingIp?: pulumi.Input<string>;
    /**
     * The name of the vip. Changing this updates the name of
     * the existing vip.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Omit this field to prevent session persistence.
     * The persistence object structure is documented below. Changing this updates
     * the persistence of the existing vip.
     */
    readonly persistence?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the pool with which the vip is associated.
     * Changing this updates the pool_id of the existing vip.
     */
    readonly poolId?: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic. Changing
     * this creates a new vip.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Port UUID for this VIP at associated floating IP (if any).
     */
    readonly portId?: pulumi.Input<string>;
    /**
     * The protocol - can be either 'TCP, 'HTTP', or
     * HTTPS'. Changing this creates a new vip.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VIP. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * VIP.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The network on which to allocate the vip's address. A
     * tenant can only create vips on networks authorized by policy (e.g. networks
     * that belong to them or networks that are shared). Changing this creates a
     * new vip.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * The owner of the vip. Required if admin wants to
     * create a vip member for another tenant. Changing this creates a new vip.
     */
    readonly tenantId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Vip resource.
 */
export interface VipArgs {
    /**
     * The IP address of the vip. Changing this creates a new
     * vip.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The administrative state of the vip.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing vip.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The maximum number of connections allowed for the
     * vip. Default is -1, meaning no limit. Changing this updates the conn_limit
     * of the existing vip.
     */
    readonly connLimit?: pulumi.Input<number>;
    /**
     * Human-readable description for the vip. Changing
     * this updates the description of the existing vip.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A *Networking* Floating IP that will be associated
     * with the vip. The Floating IP must be provisioned already.
     */
    readonly floatingIp?: pulumi.Input<string>;
    /**
     * The name of the vip. Changing this updates the name of
     * the existing vip.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Omit this field to prevent session persistence.
     * The persistence object structure is documented below. Changing this updates
     * the persistence of the existing vip.
     */
    readonly persistence?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the pool with which the vip is associated.
     * Changing this updates the pool_id of the existing vip.
     */
    readonly poolId: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic. Changing
     * this creates a new vip.
     */
    readonly port: pulumi.Input<number>;
    /**
     * The protocol - can be either 'TCP, 'HTTP', or
     * HTTPS'. Changing this creates a new vip.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VIP. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * VIP.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The network on which to allocate the vip's address. A
     * tenant can only create vips on networks authorized by policy (e.g. networks
     * that belong to them or networks that are shared). Changing this creates a
     * new vip.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * The owner of the vip. Required if admin wants to
     * create a vip member for another tenant. Changing this creates a new vip.
     */
    readonly tenantId?: pulumi.Input<string>;
}
