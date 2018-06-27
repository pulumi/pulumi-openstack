import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V3 User resource within OpenStack Keystone.
 *
 * Note: You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
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
     * The default project this user belongs to.
     */
    readonly defaultProjectId: pulumi.Output<string>;
    /**
     * A description of the user.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * The domain this user belongs to.
     */
    readonly domainId: pulumi.Output<string>;
    /**
     * Whether the user is enabled or disabled. Valid
     * values are `true` and `false`.
     */
    readonly enabled: pulumi.Output<boolean | undefined>;
    /**
     * Free-form key/value pairs of extra information.
     */
    readonly extra: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * User will not have to
     * change their password upon first use. Valid values are `true` and `false`.
     */
    readonly ignoreChangePasswordUponFirstUse: pulumi.Output<boolean | undefined>;
    /**
     * User will not have a failure
     * lockout placed on their account. Valid values are `true` and `false`.
     */
    readonly ignoreLockoutFailureAttempts: pulumi.Output<boolean | undefined>;
    /**
     * User's password will not expire.
     * Valid values are `true` and `false`.
     */
    readonly ignorePasswordExpiry: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable multi-factor
     * authentication. Valid values are `true` and `false`.
     */
    readonly multiFactorAuthEnabled: pulumi.Output<boolean | undefined>;
    /**
     * A multi-factor authentication rule.
     * The structure is documented below. Please see the
     * [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
     * for more information on how to use mulit-factor rules.
     */
    readonly multiFactorAuthRules: pulumi.Output<{
        rules: string[];
    }[] | undefined>;
    /**
     * The name of the user.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The password for the user.
     */
    readonly password: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new User.
     */
    readonly region: pulumi.Output<string>;
    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The default project this user belongs to.
     */
    readonly defaultProjectId?: pulumi.Input<string>;
    /**
     * A description of the user.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain this user belongs to.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * Whether the user is enabled or disabled. Valid
     * values are `true` and `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Free-form key/value pairs of extra information.
     */
    readonly extra?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * User will not have to
     * change their password upon first use. Valid values are `true` and `false`.
     */
    readonly ignoreChangePasswordUponFirstUse?: pulumi.Input<boolean>;
    /**
     * User will not have a failure
     * lockout placed on their account. Valid values are `true` and `false`.
     */
    readonly ignoreLockoutFailureAttempts?: pulumi.Input<boolean>;
    /**
     * User's password will not expire.
     * Valid values are `true` and `false`.
     */
    readonly ignorePasswordExpiry?: pulumi.Input<boolean>;
    /**
     * Whether to enable multi-factor
     * authentication. Valid values are `true` and `false`.
     */
    readonly multiFactorAuthEnabled?: pulumi.Input<boolean>;
    /**
     * A multi-factor authentication rule.
     * The structure is documented below. Please see the
     * [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
     * for more information on how to use mulit-factor rules.
     */
    readonly multiFactorAuthRules?: pulumi.Input<{
        rules: pulumi.Input<pulumi.Input<string>[]>;
    }[]>;
    /**
     * The name of the user.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password for the user.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new User.
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * The default project this user belongs to.
     */
    readonly defaultProjectId?: pulumi.Input<string>;
    /**
     * A description of the user.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain this user belongs to.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * Whether the user is enabled or disabled. Valid
     * values are `true` and `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Free-form key/value pairs of extra information.
     */
    readonly extra?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * User will not have to
     * change their password upon first use. Valid values are `true` and `false`.
     */
    readonly ignoreChangePasswordUponFirstUse?: pulumi.Input<boolean>;
    /**
     * User will not have a failure
     * lockout placed on their account. Valid values are `true` and `false`.
     */
    readonly ignoreLockoutFailureAttempts?: pulumi.Input<boolean>;
    /**
     * User's password will not expire.
     * Valid values are `true` and `false`.
     */
    readonly ignorePasswordExpiry?: pulumi.Input<boolean>;
    /**
     * Whether to enable multi-factor
     * authentication. Valid values are `true` and `false`.
     */
    readonly multiFactorAuthEnabled?: pulumi.Input<boolean>;
    /**
     * A multi-factor authentication rule.
     * The structure is documented below. Please see the
     * [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
     * for more information on how to use mulit-factor rules.
     */
    readonly multiFactorAuthRules?: pulumi.Input<{
        rules: pulumi.Input<pulumi.Input<string>[]>;
    }[]>;
    /**
     * The name of the user.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password for the user.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new User.
     */
    readonly region?: pulumi.Input<string>;
}
