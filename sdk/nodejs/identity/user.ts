// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const project1 = new openstack.identity.Project("project_1", {});
 * const user1 = new openstack.identity.User("user_1", {
 *     defaultProjectId: project1.id,
 *     description: "A user",
 *     extra: {
 *         email: "user_1@foobar.com",
 *     },
 *     ignoreChangePasswordUponFirstUse: true,
 *     multiFactorAuthEnabled: true,
 *     multiFactorAuthRules: [
 *         {
 *             rules: [
 *                 "password",
 *                 "totp",
 *             ],
 *         },
 *         {
 *             rules: ["password"],
 *         },
 *     ],
 *     password: "password123",
 * });
 * ```
 *
 * ## Import
 *
 * Users can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:identity/user:User user_1 89c60255-9bd6-460c-822a-e2b959ede9d2
 * ```
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
    public static readonly __pulumiType = 'openstack:identity/user:User';

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
     * The default project this user belongs to.
     */
    public readonly defaultProjectId!: pulumi.Output<string>;
    /**
     * A description of the user.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The domain this user belongs to.
     */
    public readonly domainId!: pulumi.Output<string>;
    /**
     * Whether the user is enabled or disabled. Valid
     * values are `true` and `false`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Free-form key/value pairs of extra information.
     */
    public readonly extra!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * User will not have to
     * change their password upon first use. Valid values are `true` and `false`.
     */
    public readonly ignoreChangePasswordUponFirstUse!: pulumi.Output<boolean | undefined>;
    /**
     * User will not have a failure
     * lockout placed on their account. Valid values are `true` and `false`.
     */
    public readonly ignoreLockoutFailureAttempts!: pulumi.Output<boolean | undefined>;
    /**
     * User's password will not expire.
     * Valid values are `true` and `false`.
     */
    public readonly ignorePasswordExpiry!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable multi-factor
     * authentication. Valid values are `true` and `false`.
     */
    public readonly multiFactorAuthEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * A multi-factor authentication rule.
     * The structure is documented below. Please see the
     * [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
     * for more information on how to use mulit-factor rules.
     */
    public readonly multiFactorAuthRules!: pulumi.Output<outputs.identity.UserMultiFactorAuthRule[] | undefined>;
    /**
     * The name of the user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password for the user.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new User.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["defaultProjectId"] = state ? state.defaultProjectId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domainId"] = state ? state.domainId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["extra"] = state ? state.extra : undefined;
            resourceInputs["ignoreChangePasswordUponFirstUse"] = state ? state.ignoreChangePasswordUponFirstUse : undefined;
            resourceInputs["ignoreLockoutFailureAttempts"] = state ? state.ignoreLockoutFailureAttempts : undefined;
            resourceInputs["ignorePasswordExpiry"] = state ? state.ignorePasswordExpiry : undefined;
            resourceInputs["multiFactorAuthEnabled"] = state ? state.multiFactorAuthEnabled : undefined;
            resourceInputs["multiFactorAuthRules"] = state ? state.multiFactorAuthRules : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            resourceInputs["defaultProjectId"] = args ? args.defaultProjectId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["extra"] = args ? args.extra : undefined;
            resourceInputs["ignoreChangePasswordUponFirstUse"] = args ? args.ignoreChangePasswordUponFirstUse : undefined;
            resourceInputs["ignoreLockoutFailureAttempts"] = args ? args.ignoreLockoutFailureAttempts : undefined;
            resourceInputs["ignorePasswordExpiry"] = args ? args.ignorePasswordExpiry : undefined;
            resourceInputs["multiFactorAuthEnabled"] = args ? args.multiFactorAuthEnabled : undefined;
            resourceInputs["multiFactorAuthRules"] = args ? args.multiFactorAuthRules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The default project this user belongs to.
     */
    defaultProjectId?: pulumi.Input<string>;
    /**
     * A description of the user.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain this user belongs to.
     */
    domainId?: pulumi.Input<string>;
    /**
     * Whether the user is enabled or disabled. Valid
     * values are `true` and `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Free-form key/value pairs of extra information.
     */
    extra?: pulumi.Input<{[key: string]: any}>;
    /**
     * User will not have to
     * change their password upon first use. Valid values are `true` and `false`.
     */
    ignoreChangePasswordUponFirstUse?: pulumi.Input<boolean>;
    /**
     * User will not have a failure
     * lockout placed on their account. Valid values are `true` and `false`.
     */
    ignoreLockoutFailureAttempts?: pulumi.Input<boolean>;
    /**
     * User's password will not expire.
     * Valid values are `true` and `false`.
     */
    ignorePasswordExpiry?: pulumi.Input<boolean>;
    /**
     * Whether to enable multi-factor
     * authentication. Valid values are `true` and `false`.
     */
    multiFactorAuthEnabled?: pulumi.Input<boolean>;
    /**
     * A multi-factor authentication rule.
     * The structure is documented below. Please see the
     * [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
     * for more information on how to use mulit-factor rules.
     */
    multiFactorAuthRules?: pulumi.Input<pulumi.Input<inputs.identity.UserMultiFactorAuthRule>[]>;
    /**
     * The name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * The password for the user.
     */
    password?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new User.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The default project this user belongs to.
     */
    defaultProjectId?: pulumi.Input<string>;
    /**
     * A description of the user.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain this user belongs to.
     */
    domainId?: pulumi.Input<string>;
    /**
     * Whether the user is enabled or disabled. Valid
     * values are `true` and `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Free-form key/value pairs of extra information.
     */
    extra?: pulumi.Input<{[key: string]: any}>;
    /**
     * User will not have to
     * change their password upon first use. Valid values are `true` and `false`.
     */
    ignoreChangePasswordUponFirstUse?: pulumi.Input<boolean>;
    /**
     * User will not have a failure
     * lockout placed on their account. Valid values are `true` and `false`.
     */
    ignoreLockoutFailureAttempts?: pulumi.Input<boolean>;
    /**
     * User's password will not expire.
     * Valid values are `true` and `false`.
     */
    ignorePasswordExpiry?: pulumi.Input<boolean>;
    /**
     * Whether to enable multi-factor
     * authentication. Valid values are `true` and `false`.
     */
    multiFactorAuthEnabled?: pulumi.Input<boolean>;
    /**
     * A multi-factor authentication rule.
     * The structure is documented below. Please see the
     * [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
     * for more information on how to use mulit-factor rules.
     */
    multiFactorAuthRules?: pulumi.Input<pulumi.Input<inputs.identity.UserMultiFactorAuthRule>[]>;
    /**
     * The name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * The password for the user.
     */
    password?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new User.
     */
    region?: pulumi.Input<string>;
}
