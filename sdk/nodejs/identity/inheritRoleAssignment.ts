// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 Inherit Role assignment within OpenStack Keystone. This uses the
 * Openstack keystone `OS-INHERIT` api to created inherit roles within domains
 * and parent projects for users and groups.
 *
 * > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const user1 = new openstack.identity.User("user_1", {
 *     name: "user_1",
 *     domainId: "default",
 * });
 * const role1 = new openstack.identity.Role("role_1", {
 *     name: "role_1",
 *     domainId: "default",
 * });
 * const roleAssignment1 = new openstack.identity.InheritRoleAssignment("role_assignment_1", {
 *     userId: user1.id,
 *     domainId: "default",
 *     roleId: role1.id,
 * });
 * ```
 *
 * ## Import
 *
 * Inherit role assignments can be imported using a constructed id. The id should
 * have the form of `domainID/projectID/groupID/userID/roleID`. When something is
 * not used then leave blank.
 *
 * For example this will import the inherit role assignment for:
 * projectID: 014395cd-89fc-4c9b-96b7-13d1ee79dad2,
 * userID: 4142e64b-1b35-44a0-9b1e-5affc7af1106,
 * roleID: ea257959-eeb1-4c10-8d33-26f0409a755d
 * ( domainID and groupID are left blank)
 *
 * ```sh
 * $ pulumi import openstack:identity/inheritRoleAssignment:InheritRoleAssignment role_assignment_1 /014395cd-89fc-4c9b-96b7-13d1ee79dad2//4142e64b-1b35-44a0-9b1e-5affc7af1106/ea257959-eeb1-4c10-8d33-26f0409a755d
 * ```
 */
export class InheritRoleAssignment extends pulumi.CustomResource {
    /**
     * Get an existing InheritRoleAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InheritRoleAssignmentState, opts?: pulumi.CustomResourceOptions): InheritRoleAssignment {
        return new InheritRoleAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:identity/inheritRoleAssignment:InheritRoleAssignment';

    /**
     * Returns true if the given object is an instance of InheritRoleAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InheritRoleAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InheritRoleAssignment.__pulumiType;
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
     * The project should be able to containt child projects.
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
     * Create a InheritRoleAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InheritRoleAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InheritRoleAssignmentArgs | InheritRoleAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InheritRoleAssignmentState | undefined;
            resourceInputs["domainId"] = state ? state.domainId : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as InheritRoleAssignmentArgs | undefined;
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InheritRoleAssignment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InheritRoleAssignment resources.
 */
export interface InheritRoleAssignmentState {
    /**
     * The domain to assign the role in.
     */
    domainId?: pulumi.Input<string>;
    /**
     * The group to assign the role to.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The project to assign the role in.
     * The project should be able to containt child projects.
     */
    projectId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * The role to assign.
     */
    roleId?: pulumi.Input<string>;
    /**
     * The user to assign the role to.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InheritRoleAssignment resource.
 */
export interface InheritRoleAssignmentArgs {
    /**
     * The domain to assign the role in.
     */
    domainId?: pulumi.Input<string>;
    /**
     * The group to assign the role to.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The project to assign the role in.
     * The project should be able to containt child projects.
     */
    projectId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    /**
     * The role to assign.
     */
    roleId: pulumi.Input<string>;
    /**
     * The user to assign the role to.
     */
    userId?: pulumi.Input<string>;
}
