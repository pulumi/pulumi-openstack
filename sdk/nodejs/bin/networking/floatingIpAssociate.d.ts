import * as pulumi from "@pulumi/pulumi";
/**
 * Associates a floating IP to a port. This is useful for situations
 * where you have a pre-allocated floating IP or are unable to use the
 * `openstack_networking_floatingip_v2` resource to create a floating IP.
 */
export declare class FloatingIpAssociate extends pulumi.CustomResource {
    /**
     * Get an existing FloatingIpAssociate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FloatingIpAssociateState): FloatingIpAssociate;
    /**
     * IP Address of an existing floating IP.
     */
    readonly floatingIp: pulumi.Output<string>;
    /**
     * ID of an existing port with at least one IP address to
     * associate with this floating IP.
     */
    readonly portId: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a floating IP that can be used with
     * another networking resource, such as a load balancer. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * floating IP (which may or may not have a different address).
     */
    readonly region: pulumi.Output<string>;
    /**
     * Create a FloatingIpAssociate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FloatingIpAssociateArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering FloatingIpAssociate resources.
 */
export interface FloatingIpAssociateState {
    /**
     * IP Address of an existing floating IP.
     */
    readonly floatingIp?: pulumi.Input<string>;
    /**
     * ID of an existing port with at least one IP address to
     * associate with this floating IP.
     */
    readonly portId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a floating IP that can be used with
     * another networking resource, such as a load balancer. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * floating IP (which may or may not have a different address).
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a FloatingIpAssociate resource.
 */
export interface FloatingIpAssociateArgs {
    /**
     * IP Address of an existing floating IP.
     */
    readonly floatingIp: pulumi.Input<string>;
    /**
     * ID of an existing port with at least one IP address to
     * associate with this floating IP.
     */
    readonly portId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a floating IP that can be used with
     * another networking resource, such as a load balancer. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * floating IP (which may or may not have a different address).
     */
    readonly region?: pulumi.Input<string>;
}
