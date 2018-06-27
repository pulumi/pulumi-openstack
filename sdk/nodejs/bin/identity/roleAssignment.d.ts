import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V3 Role assignment within OpenStack Keystone.
 *
 * Note: You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 */
export declare class RoleAssignment extends pulumi.CustomResource {
    /**
     * Get an existing RoleAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleAssignmentState): RoleAssignment;
    /**
     * The domain to assign the role in.
     */
    readonly domainId: pulumi.Output<string | undefined>;
    /**
     * The group to assign the role to.
     */
    readonly groupId: pulumi.Output<string | undefined>;
    /**
     * The project to assign the role in.
     */
    readonly projectId: pulumi.Output<string | undefined>;
    /**
     * The role to assign.
     */
    readonly roleId: pulumi.Output<string>;
    /**
     * The user to assign the role to.
     */
    readonly userId: pulumi.Output<string | undefined>;
    /**
     * Create a RoleAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleAssignmentArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering RoleAssignment resources.
 */
export interface RoleAssignmentState {
    /**
     * The domain to assign the role in.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * The group to assign the role to.
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The project to assign the role in.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The role to assign.
     */
    readonly roleId?: pulumi.Input<string>;
    /**
     * The user to assign the role to.
     */
    readonly userId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a RoleAssignment resource.
 */
export interface RoleAssignmentArgs {
    /**
     * The domain to assign the role in.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * The group to assign the role to.
     */
    readonly groupId?: pulumi.Input<string>;
    /**
     * The project to assign the role in.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The role to assign.
     */
    readonly roleId: pulumi.Input<string>;
    /**
     * The user to assign the role to.
     */
    readonly userId?: pulumi.Input<string>;
}
