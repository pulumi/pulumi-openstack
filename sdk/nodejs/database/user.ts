// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V1 DB user resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ### User
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const basic = new openstack.database.User("basic", {
 *     databases: ["testdb"],
 *     instance: openstack_db_instance_v1_basic.id,
 *     password: "password",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/db_user_v1.html.markdown.
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:database/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * A list of database user should have access to.
     */
    public readonly databases!: pulumi.Output<string[]>;
    public readonly host!: pulumi.Output<string | undefined>;
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * A unique name for the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * User's password.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Openstack region resource is created in.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["databases"] = state ? state.databases : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            if (!args || args.password === undefined) {
                throw new Error("Missing required property 'password'");
            }
            inputs["databases"] = args ? args.databases : undefined;
            inputs["host"] = args ? args.host : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(User.__pulumiType, name, inputs, opts);
    }
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
    readonly region?: pulumi.Input<string>;
}
