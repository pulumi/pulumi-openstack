// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const example = openstack.keymanager.getSecret({
 *     mode: "cbc",
 *     secretType: "passphrase",
 * });
 * ```
 * ## Date Filters
 *
 * The values for the `expirationFilter`, `createdAtFilter`, and
 * `updatedAtFilter` parameters are comma-separated lists of time stamps in
 * RFC3339 format. The time stamps can be prefixed with any of these comparison
 * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
 * (less-than), *lte:* (less-than-or-equal).
 *
 * For example, to get a passphrase a Secret with CBC moda, that will expire in
 * January of 2020:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const dateFilterExample = openstack.keymanager.getSecret({
 *     expirationFilter: "gt:2020-01-01T00:00:00Z",
 *     mode: "cbc",
 *     secretType: "passphrase",
 * });
 * ```
 */
export function getSecret(args?: GetSecretArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:keymanager/getSecret:getSecret", {
        "aclOnly": args.aclOnly,
        "algorithm": args.algorithm,
        "bitLength": args.bitLength,
        "createdAtFilter": args.createdAtFilter,
        "expirationFilter": args.expirationFilter,
        "mode": args.mode,
        "name": args.name,
        "region": args.region,
        "secretType": args.secretType,
        "updatedAtFilter": args.updatedAtFilter,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecret.
 */
export interface GetSecretArgs {
    /**
     * Select the Secret with an ACL that contains the user.
     * Project scope is ignored. Defaults to `false`.
     */
    aclOnly?: boolean;
    /**
     * The Secret algorithm.
     */
    algorithm?: string;
    /**
     * The Secret bit length.
     */
    bitLength?: number;
    /**
     * Date filter to select the Secret with
     * created matching the specified criteria. See Date Filters below for more
     * detail.
     */
    createdAtFilter?: string;
    /**
     * Date filter to select the Secret with
     * expiration matching the specified criteria. See Date Filters below for more
     * detail.
     */
    expirationFilter?: string;
    /**
     * The Secret mode.
     */
    mode?: string;
    /**
     * The Secret name.
     */
    name?: string;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a secret. If omitted, the `region`
     * argument of the provider is used.
     */
    region?: string;
    /**
     * The Secret type. For more information see
     * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
     */
    secretType?: string;
    /**
     * Date filter to select the Secret with
     * updated matching the specified criteria. See Date Filters below for more
     * detail.
     */
    updatedAtFilter?: string;
}

/**
 * A collection of values returned by getSecret.
 */
export interface GetSecretResult {
    /**
     * See Argument Reference above.
     */
    readonly aclOnly?: boolean;
    /**
     * The list of ACLs assigned to a secret. The `read` structure is described below.
     */
    readonly acls: outputs.keymanager.GetSecretAcl[];
    /**
     * See Argument Reference above.
     */
    readonly algorithm?: string;
    /**
     * See Argument Reference above.
     */
    readonly bitLength?: number;
    /**
     * The map of the content types, assigned on the secret.
     */
    readonly contentTypes: {[key: string]: any};
    /**
     * The date the secret ACL was created.
     */
    readonly createdAt: string;
    /**
     * See Argument Reference above.
     */
    readonly createdAtFilter?: string;
    /**
     * The creator of the secret.
     */
    readonly creatorId: string;
    /**
     * The date the secret will expire.
     */
    readonly expiration: string;
    /**
     * See Argument Reference above.
     */
    readonly expirationFilter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The map of metadata, assigned on the secret, which has been
     * explicitly and implicitly added.
     */
    readonly metadata: {[key: string]: any};
    /**
     * See Argument Reference above.
     */
    readonly mode?: string;
    /**
     * See Argument Reference above.
     */
    readonly name?: string;
    /**
     * The secret payload.
     */
    readonly payload: string;
    /**
     * The Secret encoding.
     */
    readonly payloadContentEncoding: string;
    /**
     * The Secret content type.
     */
    readonly payloadContentType: string;
    /**
     * See Argument Reference above.
     */
    readonly region?: string;
    /**
     * The secret reference / where to find the secret.
     */
    readonly secretRef: string;
    /**
     * See Argument Reference above.
     */
    readonly secretType?: string;
    /**
     * The status of the secret.
     */
    readonly status: string;
    /**
     * The date the secret ACL was last updated.
     */
    readonly updatedAt: string;
    /**
     * See Argument Reference above.
     */
    readonly updatedAtFilter?: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const example = openstack.keymanager.getSecret({
 *     mode: "cbc",
 *     secretType: "passphrase",
 * });
 * ```
 * ## Date Filters
 *
 * The values for the `expirationFilter`, `createdAtFilter`, and
 * `updatedAtFilter` parameters are comma-separated lists of time stamps in
 * RFC3339 format. The time stamps can be prefixed with any of these comparison
 * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
 * (less-than), *lte:* (less-than-or-equal).
 *
 * For example, to get a passphrase a Secret with CBC moda, that will expire in
 * January of 2020:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const dateFilterExample = openstack.keymanager.getSecret({
 *     expirationFilter: "gt:2020-01-01T00:00:00Z",
 *     mode: "cbc",
 *     secretType: "passphrase",
 * });
 * ```
 */
export function getSecretOutput(args?: GetSecretOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretResult> {
    return pulumi.output(args).apply((a: any) => getSecret(a, opts))
}

/**
 * A collection of arguments for invoking getSecret.
 */
export interface GetSecretOutputArgs {
    /**
     * Select the Secret with an ACL that contains the user.
     * Project scope is ignored. Defaults to `false`.
     */
    aclOnly?: pulumi.Input<boolean>;
    /**
     * The Secret algorithm.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * The Secret bit length.
     */
    bitLength?: pulumi.Input<number>;
    /**
     * Date filter to select the Secret with
     * created matching the specified criteria. See Date Filters below for more
     * detail.
     */
    createdAtFilter?: pulumi.Input<string>;
    /**
     * Date filter to select the Secret with
     * expiration matching the specified criteria. See Date Filters below for more
     * detail.
     */
    expirationFilter?: pulumi.Input<string>;
    /**
     * The Secret mode.
     */
    mode?: pulumi.Input<string>;
    /**
     * The Secret name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a secret. If omitted, the `region`
     * argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The Secret type. For more information see
     * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
     */
    secretType?: pulumi.Input<string>;
    /**
     * Date filter to select the Secret with
     * updated matching the specified criteria. See Date Filters below for more
     * detail.
     */
    updatedAtFilter?: pulumi.Input<string>;
}
