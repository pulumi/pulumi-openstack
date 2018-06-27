import * as pulumi from "@pulumi/pulumi";
/**
 * Use this data source to get the ID of an OpenStack user.
 */
export declare function getUser(args?: GetUserArgs): Promise<GetUserResult>;
/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
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
     * The identity provider ID of the user.
     */
    readonly idpId?: pulumi.Input<string>;
    /**
     * The name of the user.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
     */
    readonly passwordExpiresAt?: pulumi.Input<string>;
    /**
     * The protocol ID of the user.
     */
    readonly protocolId?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The unique ID of the user.
     */
    readonly uniqueId?: pulumi.Input<string>;
}
/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * See Argument Reference above.
     */
    readonly defaultProjectId: string;
    /**
     * See Argument Reference above.
     */
    readonly domainId: string;
    /**
     * The region the user is located in.
     */
    readonly region: string;
}
