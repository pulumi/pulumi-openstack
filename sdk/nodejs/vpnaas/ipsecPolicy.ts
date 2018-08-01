// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Manages a V2 Neutron IPSec policy resource within OpenStack.
 */
export class IpsecPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IpsecPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpsecPolicyState): IpsecPolicy {
        return new IpsecPolicy(name, <any>state, { id });
    }

    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     */
    public readonly authAlgorithm: pulumi.Output<string>;
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     */
    public readonly encapsulationMode: pulumi.Output<string>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    public readonly encryptionAlgorithm: pulumi.Output<string>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
     * Default is seconds.
     * - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     */
    public readonly lifetimes: pulumi.Output<{ units: string, value: number }[]>;
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
     * Changing this updates the existing policy.
     */
    public readonly pfs: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     */
    public readonly region: pulumi.Output<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     */
    public readonly tenantId: pulumi.Output<string>;
    /**
     * The transform protocol. Valid values are ESP, AH and AH-ESP.
     * Changing this updates the existing policy. Default is ESP.
     */
    public readonly transformProtocol: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a IpsecPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpsecPolicyArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: IpsecPolicyArgs | IpsecPolicyState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: IpsecPolicyState = argsOrState as IpsecPolicyState | undefined;
            inputs["authAlgorithm"] = state ? state.authAlgorithm : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encapsulationMode"] = state ? state.encapsulationMode : undefined;
            inputs["encryptionAlgorithm"] = state ? state.encryptionAlgorithm : undefined;
            inputs["lifetimes"] = state ? state.lifetimes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pfs"] = state ? state.pfs : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["transformProtocol"] = state ? state.transformProtocol : undefined;
            inputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as IpsecPolicyArgs | undefined;
            inputs["authAlgorithm"] = args ? args.authAlgorithm : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encapsulationMode"] = args ? args.encapsulationMode : undefined;
            inputs["encryptionAlgorithm"] = args ? args.encryptionAlgorithm : undefined;
            inputs["lifetimes"] = args ? args.lifetimes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pfs"] = args ? args.pfs : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["transformProtocol"] = args ? args.transformProtocol : undefined;
            inputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        super("openstack:vpnaas/ipsecPolicy:IpsecPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpsecPolicy resources.
 */
export interface IpsecPolicyState {
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
     * The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     */
    readonly encapsulationMode?: pulumi.Input<string>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    readonly encryptionAlgorithm?: pulumi.Input<string>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
     * Default is seconds.
     * - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     */
    readonly lifetimes?: pulumi.Input<pulumi.Input<{ units?: pulumi.Input<string>, value?: pulumi.Input<number> }>[]>;
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
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * The transform protocol. Valid values are ESP, AH and AH-ESP.
     * Changing this updates the existing policy. Default is ESP.
     */
    readonly transformProtocol?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a IpsecPolicy resource.
 */
export interface IpsecPolicyArgs {
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
     * The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     */
    readonly encapsulationMode?: pulumi.Input<string>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    readonly encryptionAlgorithm?: pulumi.Input<string>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
     * Default is seconds.
     * - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     */
    readonly lifetimes?: pulumi.Input<pulumi.Input<{ units?: pulumi.Input<string>, value?: pulumi.Input<number> }>[]>;
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
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * The transform protocol. Valid values are ESP, AH and AH-ESP.
     * Changing this updates the existing policy. Default is ESP.
     */
    readonly transformProtocol?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
