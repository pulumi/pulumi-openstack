import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 neutron security group resource within OpenStack.
 * Unlike Nova security groups, neutron separates the group from the rules
 * and also allows an admin to target a specific tenant_id.
 */
export declare class SecGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecGroupState): SecGroup;
    /**
     * Whether or not to delete the default
     * egress security rules. This is `false` by default. See the below note
     * for more information.
     */
    readonly deleteDefaultRules: pulumi.Output<boolean | undefined>;
    /**
     * A unique name for the security group.
     */
    readonly description: pulumi.Output<string>;
    /**
     * A unique name for the security group.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Create a SecGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecGroupArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering SecGroup resources.
 */
export interface SecGroupState {
    /**
     * Whether or not to delete the default
     * egress security rules. This is `false` by default. See the below note
     * for more information.
     */
    readonly deleteDefaultRules?: pulumi.Input<boolean>;
    /**
     * A unique name for the security group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique name for the security group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group.
     */
    readonly tenantId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a SecGroup resource.
 */
export interface SecGroupArgs {
    /**
     * Whether or not to delete the default
     * egress security rules. This is `false` by default. See the below note
     * for more information.
     */
    readonly deleteDefaultRules?: pulumi.Input<boolean>;
    /**
     * A unique name for the security group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique name for the security group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group.
     */
    readonly tenantId?: pulumi.Input<string>;
}
