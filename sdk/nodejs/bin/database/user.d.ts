import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V1 DB user resource within OpenStack.
 */
export declare class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState): User;
    /**
     * A list of database user should have access to.
     */
    readonly databases: pulumi.Output<string[]>;
    readonly host: pulumi.Output<string | undefined>;
    readonly instanceId: pulumi.Output<string>;
    /**
     * A unique name for the resource.
     */
    readonly name: pulumi.Output<string>;
    /**
     * User's password.
     */
    readonly password: pulumi.Output<string>;
    /**
     * Openstack region resource is created in.
     */
    readonly region: pulumi.Output<string>;
    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * A list of database user should have access to.
     */
    readonly databases?: pulumi.Input<pulumi.Input<string>[]>;
    readonly host?: pulumi.Input<string>;
    readonly instanceId?: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * User's password.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Openstack region resource is created in.
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * A list of database user should have access to.
     */
    readonly databases?: pulumi.Input<pulumi.Input<string>[]>;
    readonly host?: pulumi.Input<string>;
    readonly instanceId: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * User's password.
     */
    readonly password: pulumi.Input<string>;
    /**
     * Openstack region resource is created in.
     */
    readonly region: pulumi.Input<string>;
}
