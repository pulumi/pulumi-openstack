// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron IKE policy resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const policy1 = new openstack.vpnaas.IkePolicy("policy1", {});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/vpnaas_ike_policy_v2.html.markdown.
 */
export class IkePolicy extends pulumi.CustomResource {
    /**
     * Get an existing IkePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IkePolicyState, opts?: pulumi.CustomResourceOptions): IkePolicy {
        return new IkePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:vpnaas/ikePolicy:IkePolicy';

    /**
     * Returns true if the given object is an instance of IkePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IkePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IkePolicy.__pulumiType;
    }

    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     */
    public readonly authAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    public readonly encryptionAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * The IKE mode. A valid value is v1 or v2. Default is v1.
     * Changing this updates the existing policy.
     */
    public readonly ikeVersion!: pulumi.Output<string | undefined>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
     * Default is seconds.
     * - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     */
    public readonly lifetimes!: pulumi.Output<outputs.vpnaas.IkePolicyLifetime[]>;
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
     * Changing this updates the existing policy.
     */
    public readonly pfs!: pulumi.Output<string | undefined>;
    /**
     * The IKE mode. A valid value is main, which is the default.
     * Changing this updates the existing policy.
     */
    public readonly phase1NegotiationMode!: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a service for another policy. Changing this creates a new policy.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a IkePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IkePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IkePolicyArgs | IkePolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IkePolicyState | undefined;
            inputs["authAlgorithm"] = state ? state.authAlgorithm : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encryptionAlgorithm"] = state ? state.encryptionAlgorithm : undefined;
            inputs["ikeVersion"] = state ? state.ikeVersion : undefined;
            inputs["lifetimes"] = state ? state.lifetimes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pfs"] = state ? state.pfs : undefined;
            inputs["phase1NegotiationMode"] = state ? state.phase1NegotiationMode : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as IkePolicyArgs | undefined;
            inputs["authAlgorithm"] = args ? args.authAlgorithm : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encryptionAlgorithm"] = args ? args.encryptionAlgorithm : undefined;
            inputs["ikeVersion"] = args ? args.ikeVersion : undefined;
            inputs["lifetimes"] = args ? args.lifetimes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pfs"] = args ? args.pfs : undefined;
            inputs["phase1NegotiationMode"] = args ? args.phase1NegotiationMode : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IkePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IkePolicy resources.
 */
export interface IkePolicyState {
    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     */
    readonly authAlgorithm?: pulumi.Input<string>;
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    readonly encryptionAlgorithm?: pulumi.Input<string>;
    /**
     * The IKE mode. A valid value is v1 or v2. Default is v1.
     * Changing this updates the existing policy.
     */
    readonly ikeVersion?: pulumi.Input<string>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
     * Default is seconds.
     * - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     */
    readonly lifetimes?: pulumi.Input<pulumi.Input<inputs.vpnaas.IkePolicyLifetime>[]>;
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
     * Changing this updates the existing policy.
     */
    readonly pfs?: pulumi.Input<string>;
    /**
     * The IKE mode. A valid value is main, which is the default.
     * Changing this updates the existing policy.
     */
    readonly phase1NegotiationMode?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a service for another policy. Changing this creates a new policy.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a IkePolicy resource.
 */
export interface IkePolicyArgs {
    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     */
    readonly authAlgorithm?: pulumi.Input<string>;
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    readonly encryptionAlgorithm?: pulumi.Input<string>;
    /**
     * The IKE mode. A valid value is v1 or v2. Default is v1.
     * Changing this updates the existing policy.
     */
    readonly ikeVersion?: pulumi.Input<string>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
     * Default is seconds.
     * - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     */
    readonly lifetimes?: pulumi.Input<pulumi.Input<inputs.vpnaas.IkePolicyLifetime>[]>;
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
     * Changing this updates the existing policy.
     */
    readonly pfs?: pulumi.Input<string>;
    /**
     * The IKE mode. A valid value is main, which is the default.
     * Changing this updates the existing policy.
     */
    readonly phase1NegotiationMode?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a service for another policy. Changing this creates a new policy.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
