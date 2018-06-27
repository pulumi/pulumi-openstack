import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 floating IP resource within OpenStack Nova (compute)
 * that can be used for compute instances.
 *
 * Please note that managing floating IPs through the OpenStack Compute API has
 * been deprecated. Unless you are using an older OpenStack environment, it is
 * recommended to use the [`openstack_networking_floatingip_v2`](networking_floatingip_v2.html)
 * resource instead, which uses the OpenStack Networking API.
 */
export declare class FloatingIP extends pulumi.CustomResource {
    /**
     * Get an existing FloatingIP resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FloatingIPState): FloatingIP;
    /**
     * The actual floating IP address itself.
     */
    readonly address: pulumi.Output<string>;
    /**
     * The fixed IP address corresponding to the floating IP.
     */
    readonly fixedIp: pulumi.Output<string>;
    /**
     * UUID of the compute instance associated with the floating IP.
     */
    readonly instanceId: pulumi.Output<string>;
    /**
     * The name of the pool from which to obtain the floating
     * IP. Changing this creates a new floating IP.
     */
    readonly pool: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a floating IP that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new floating IP (which may or may not
     * have a different address).
     */
    readonly region: pulumi.Output<string>;
    /**
     * Create a FloatingIP resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FloatingIPArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering FloatingIP resources.
 */
export interface FloatingIPState {
    /**
     * The actual floating IP address itself.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The fixed IP address corresponding to the floating IP.
     */
    readonly fixedIp?: pulumi.Input<string>;
    /**
     * UUID of the compute instance associated with the floating IP.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The name of the pool from which to obtain the floating
     * IP. Changing this creates a new floating IP.
     */
    readonly pool?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a floating IP that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new floating IP (which may or may not
     * have a different address).
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a FloatingIP resource.
 */
export interface FloatingIPArgs {
    /**
     * The name of the pool from which to obtain the floating
     * IP. Changing this creates a new floating IP.
     */
    readonly pool: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a floating IP that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new floating IP (which may or may not
     * have a different address).
     */
    readonly region?: pulumi.Input<string>;
}
