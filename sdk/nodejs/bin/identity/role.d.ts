import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V3 Role resource within OpenStack Keystone.
 *
 * Note: You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 */
export declare class Role extends pulumi.CustomResource {
    /**
     * Get an existing Role resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleState): Role;
    /**
     * The domain the role belongs to.
     */
    readonly domainId: pulumi.Output<string>;
    /**
     * The name of the role.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new Role.
     */
    readonly region: pulumi.Output<string>;
    /**
     * Create a Role resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RoleArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Role resources.
 */
export interface RoleState {
    /**
     * The domain the role belongs to.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new Role.
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Role resource.
 */
export interface RoleArgs {
    /**
     * The domain the role belongs to.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * The name of the role.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new Role.
     */
    readonly region?: pulumi.Input<string>;
}
