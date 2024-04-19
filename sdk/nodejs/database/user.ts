// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V1 DB user resource within OpenStack.
 *
 * > **Note:** All arguments including the database password will be stored in the
 * raw state as plain-text. Read more about sensitive data in
 * state.
 *
 * ## Example Usage
 *
 * ### User
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const basic = new openstack.database.User("basic", {
 *     name: "basic",
 *     instanceId: basicOpenstackDbInstanceV1.id,
 *     password: "password",
 *     databases: ["testdb"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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
    /**
     * The ID for the database instance.
     */
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["databases"] = state ? state.databases : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["databases"] = args ? args.databases : undefined;
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * A list of database user should have access to.
     */
    databases?: pulumi.Input<pulumi.Input<string>[]>;
    host?: pulumi.Input<string>;
    /**
     * The ID for the database instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * User's password.
     */
    password?: pulumi.Input<string>;
    /**
     * Openstack region resource is created in.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * A list of database user should have access to.
     */
    databases?: pulumi.Input<pulumi.Input<string>[]>;
    host?: pulumi.Input<string>;
    /**
     * The ID for the database instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * User's password.
     */
    password: pulumi.Input<string>;
    /**
     * Openstack region resource is created in.
     */
    region?: pulumi.Input<string>;
}
