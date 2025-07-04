// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID and public key of an OpenStack keypair.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const kp = openstack.compute.getKeypair({
 *     name: "sand",
 * });
 * ```
 */
export function getKeypair(args: GetKeypairArgs, opts?: pulumi.InvokeOptions): Promise<GetKeypairResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:compute/getKeypair:getKeypair", {
        "name": args.name,
        "region": args.region,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getKeypair.
 */
export interface GetKeypairArgs {
    /**
     * The unique name of the keypair.
     */
    name: string;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The user id of the owner of the key pair.
     * This parameter can be specified only if the provider is configured to use
     * the credentials of an OpenStack administrator.
     */
    userId?: string;
}

/**
 * A collection of values returned by getKeypair.
 */
export interface GetKeypairResult {
    /**
     * The fingerprint of the OpenSSH key.
     */
    readonly fingerprint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * The OpenSSH-formatted public key of the keypair.
     */
    readonly publicKey: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * See Argument Reference above.
     */
    readonly userId: string;
}
/**
 * Use this data source to get the ID and public key of an OpenStack keypair.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const kp = openstack.compute.getKeypair({
 *     name: "sand",
 * });
 * ```
 */
export function getKeypairOutput(args: GetKeypairOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetKeypairResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:compute/getKeypair:getKeypair", {
        "name": args.name,
        "region": args.region,
        "userId": args.userId,
    }, opts);
}

/**
 * A collection of arguments for invoking getKeypair.
 */
export interface GetKeypairOutputArgs {
    /**
     * The unique name of the keypair.
     */
    name: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The user id of the owner of the key pair.
     * This parameter can be specified only if the provider is configured to use
     * the credentials of an OpenStack administrator.
     */
    userId?: pulumi.Input<string>;
}
