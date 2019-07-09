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
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/compute_keypair_v2.html.markdown.
 */
export function getKeypair(args: GetKeypairArgs, opts?: pulumi.InvokeOptions): Promise<GetKeypairResult> {
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
    readonly name: string;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: string;
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
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
