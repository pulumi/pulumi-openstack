// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 Role assignment within OpenStack Keystone.
 *
 * Note: You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const project1 = new openstack.identity.Project("project1", {});
 * const user1 = new openstack.identity.User("user1", {
 *     defaultProjectId: project1.id,
 * });
 * const role1 = new openstack.identity.Role("role1", {});
 * const roleAssignment1 = new openstack.identity.RoleAssignment("roleAssignment1", {
 *     projectId: project1.id,
 *     roleId: role1.id,
 *     userId: user1.id,
 * });
 * ```
 */
export class RoleAssignment extends pulumi.CustomResource {
    /**
     * Get an existing RoleAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleAssignmentState, opts?: pulumi.CustomResourceOptions): RoleAssignment {
        return new RoleAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:identity/roleAssignment:RoleAssignment';

    /**
     * Returns true if the given object is an instance of RoleAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleAssignment.__pulumiType;
    }

    /**
     * The domain to assign the role in.
     */
    public readonly domainId!: pulumi.Output<string | undefined>;
    /**
     * The group to assign the role to.
     */
    public readonly groupId!: pulumi.Output<string | undefined>;
    /**
     * The project to assign the role in.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string>;
    /**
     * The role to assign.
     */
    public readonly roleId!: pulumi.Output<string>;
    /**
     * The user to assign the role to.
     */
    public readonly userId!: pulumi.Output<string | undefined>;

    /**
     * Create a RoleAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleAssignmentArgs | RoleAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RoleAssignmentState | undefined;
            inputs["domainId"] = state ? state.domainId : undefined;
            inputs["groupId"] = state ? state.groupId : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["roleId"] = state ? state.roleId : undefined;
            inputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as RoleAssignmentArgs | undefined;
            if (!args || args.roleId === undefined) {
                throw new Error("Missing required property 'roleId'");
            }
            inputs["domainId"] = args ? args.domainId : undefined;
            inputs["groupId"] = args ? args.groupId : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["roleId"] = args ? args.roleId : undefined;
            inputs["userId"] = args ? args.userId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RoleAssignment.__pulumiType, name, inputs, opts);
    }
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
    readonly region?: pulumi.Input<string>;
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
    readonly region?: pulumi.Input<string>;
    /**
     * The role to assign.
     */
    readonly roleId: pulumi.Input<string>;
    /**
     * The user to assign the role to.
     */
    readonly userId?: pulumi.Input<string>;
}
