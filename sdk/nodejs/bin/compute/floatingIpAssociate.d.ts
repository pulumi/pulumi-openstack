import * as pulumi from "@pulumi/pulumi";
/**
 * Associate a floating IP to an instance. This can be used instead of the
 * `floating_ip` options in `openstack_compute_instance_v2`.
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
     * The specific IP address to direct traffic to.
     */
    readonly fixedIp: pulumi.Output<string | undefined>;
    /**
     * The floating IP to associate.
     */
    readonly floatingIp: pulumi.Output<string>;
    /**
     * The instance to associte the floating IP with.
     */
    readonly instanceId: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new floatingip_associate.
     */
    readonly region: pulumi.Output<string>;
    /**
     * In cases where the OpenStack environment
     * does not automatically wait until the association has finished, set this
     * option to have Terraform poll the instance until the floating IP has been
     * associated. Defaults to false.
     */
    readonly waitUntilAssociated: pulumi.Output<boolean | undefined>;
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
     * The specific IP address to direct traffic to.
     */
    readonly fixedIp?: pulumi.Input<string>;
    /**
     * The floating IP to associate.
     */
    readonly floatingIp?: pulumi.Input<string>;
    /**
     * The instance to associte the floating IP with.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new floatingip_associate.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * In cases where the OpenStack environment
     * does not automatically wait until the association has finished, set this
     * option to have Terraform poll the instance until the floating IP has been
     * associated. Defaults to false.
     */
    readonly waitUntilAssociated?: pulumi.Input<boolean>;
}
/**
 * The set of arguments for constructing a FloatingIpAssociate resource.
 */
export interface FloatingIpAssociateArgs {
    /**
     * The specific IP address to direct traffic to.
     */
    readonly fixedIp?: pulumi.Input<string>;
    /**
     * The floating IP to associate.
     */
    readonly floatingIp: pulumi.Input<string>;
    /**
     * The instance to associte the floating IP with.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new floatingip_associate.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * In cases where the OpenStack environment
     * does not automatically wait until the association has finished, set this
     * option to have Terraform poll the instance until the floating IP has been
     * associated. Defaults to false.
     */
    readonly waitUntilAssociated?: pulumi.Input<boolean>;
}
