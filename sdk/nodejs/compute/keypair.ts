// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ### Import an Existing Public Key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const test_keypair = new openstack.compute.Keypair("test-keypair", {
 *     name: "my-keypair",
 *     publicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLotBCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAnOfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZqd9LvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TaIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIF61p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB",
 * });
 * ```
 *
 * ### Generate a Public/Private Key Pair
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const test_keypair = new openstack.compute.Keypair("test-keypair", {name: "my-keypair"});
 * ```
 *
 * ## Import
 *
 * Keypairs can be imported using the `name`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:compute/keypair:Keypair my-keypair test-keypair
 * ```
 */
export class Keypair extends pulumi.CustomResource {
    /**
     * Get an existing Keypair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeypairState, opts?: pulumi.CustomResourceOptions): Keypair {
        return new Keypair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:compute/keypair:Keypair';

    /**
     * Returns true if the given object is an instance of Keypair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Keypair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Keypair.__pulumiType;
    }

    /**
     * The fingerprint of the public key.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * A unique name for the keypair. Changing this creates a new
     * keypair.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The generated private key when no public key is specified.
     */
    public /*out*/ readonly privateKey!: pulumi.Output<string>;
    /**
     * A pregenerated OpenSSH-formatted public key.
     * Changing this creates a new keypair. If a public key is not specified, then
     * a public/private key pair will be automatically generated. If a pair is
     * created, then destroying this resource means you will lose access to that
     * keypair forever.
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new keypair.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * This allows administrative users to operate key-pairs
     * of specified user ID. For this feature your need to have openstack microversion
     * 2.10 (Liberty) or later.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Keypair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeypairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeypairArgs | KeypairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeypairState | undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as KeypairArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["privateKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Keypair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Keypair resources.
 */
export interface KeypairState {
    /**
     * The fingerprint of the public key.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * A unique name for the keypair. Changing this creates a new
     * keypair.
     */
    name?: pulumi.Input<string>;
    /**
     * The generated private key when no public key is specified.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * A pregenerated OpenSSH-formatted public key.
     * Changing this creates a new keypair. If a public key is not specified, then
     * a public/private key pair will be automatically generated. If a pair is
     * created, then destroying this resource means you will lose access to that
     * keypair forever.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new keypair.
     */
    region?: pulumi.Input<string>;
    /**
     * This allows administrative users to operate key-pairs
     * of specified user ID. For this feature your need to have openstack microversion
     * 2.10 (Liberty) or later.
     */
    userId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Keypair resource.
 */
export interface KeypairArgs {
    /**
     * A unique name for the keypair. Changing this creates a new
     * keypair.
     */
    name?: pulumi.Input<string>;
    /**
     * A pregenerated OpenSSH-formatted public key.
     * Changing this creates a new keypair. If a public key is not specified, then
     * a public/private key pair will be automatically generated. If a pair is
     * created, then destroying this resource means you will lose access to that
     * keypair forever.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new keypair.
     */
    region?: pulumi.Input<string>;
    /**
     * This allows administrative users to operate key-pairs
     * of specified user ID. For this feature your need to have openstack microversion
     * 2.10 (Liberty) or later.
     */
    userId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
