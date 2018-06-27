import * as pulumi from "@pulumi/pulumi";
/**
 * Use this data source to get the ID and public key of an OpenStack keypair.
 */
export declare function getKeypair(args: GetKeypairArgs): Promise<GetKeypairResult>;
/**
 * A collection of arguments for invoking getKeypair.
 */
export interface GetKeypairArgs {
    /**
     * The unique name of the keypair.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * A collection of values returned by getKeypair.
 */
export interface GetKeypairResult {
    /**
     * The OpenSSH-formatted public key of the keypair.
     */
    readonly publicKey: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
}
