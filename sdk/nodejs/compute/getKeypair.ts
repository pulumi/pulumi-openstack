// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
 * const kp = pulumi.output(openstack.compute.getKeypair({
 *     name: "sand",
 * }));
 * ```
 */
export function getKeypair(args: GetKeypairArgs, opts?: pulumi.InvokeOptions): Promise<GetKeypairResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:compute/getKeypair:getKeypair", {
        "name": args.name,
        "region": args.region,
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
}

export function getKeypairOutput(args: GetKeypairOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKeypairResult> {
    return pulumi.output(args).apply(a => getKeypair(a, opts))
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
}
