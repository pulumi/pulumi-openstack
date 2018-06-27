import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V1 DB database resource within OpenStack.
 */
export declare class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState): Database;
    /**
     * The ID for the database instance.
     */
    readonly instanceId: pulumi.Output<string>;
    /**
     * A unique name for the resource.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Openstack region resource is created in.
     */
    readonly region: pulumi.Output<string | undefined>;
    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * The ID for the database instance.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Openstack region resource is created in.
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The ID for the database instance.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Openstack region resource is created in.
     */
    readonly region?: pulumi.Input<string>;
}
