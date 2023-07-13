// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron IPSec policy resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const policy1 = new openstack.vpnaas.IpSecPolicy("policy1", {});
 * ```
 *
 * ## Import
 *
 * Policies can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:vpnaas/ipSecPolicy:IpSecPolicy policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 */
export class IpSecPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IpSecPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpSecPolicyState, opts?: pulumi.CustomResourceOptions): IpSecPolicy {
        return new IpSecPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:vpnaas/ipSecPolicy:IpSecPolicy';

    /**
     * Returns true if the given object is an instance of IpSecPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpSecPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpSecPolicy.__pulumiType;
    }

    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     */
    public readonly authAlgorithm!: pulumi.Output<string>;
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     */
    public readonly encapsulationMode!: pulumi.Output<string>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    public readonly encryptionAlgorithm!: pulumi.Output<string>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     */
    public readonly lifetimes!: pulumi.Output<outputs.vpnaas.IpSecPolicyLifetime[]>;
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
     * is group5. Changing this updates the existing policy.
     */
    public readonly pfs!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The transform protocol. Valid values are esp, ah and ah-esp.
     * Changing this updates the existing policy. Default is ESP.
     */
    public readonly transformProtocol!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a IpSecPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpSecPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpSecPolicyArgs | IpSecPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpSecPolicyState | undefined;
            resourceInputs["authAlgorithm"] = state ? state.authAlgorithm : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encapsulationMode"] = state ? state.encapsulationMode : undefined;
            resourceInputs["encryptionAlgorithm"] = state ? state.encryptionAlgorithm : undefined;
            resourceInputs["lifetimes"] = state ? state.lifetimes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pfs"] = state ? state.pfs : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["transformProtocol"] = state ? state.transformProtocol : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as IpSecPolicyArgs | undefined;
            resourceInputs["authAlgorithm"] = args ? args.authAlgorithm : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encapsulationMode"] = args ? args.encapsulationMode : undefined;
            resourceInputs["encryptionAlgorithm"] = args ? args.encryptionAlgorithm : undefined;
            resourceInputs["lifetimes"] = args ? args.lifetimes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pfs"] = args ? args.pfs : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["transformProtocol"] = args ? args.transformProtocol : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpSecPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpSecPolicy resources.
 */
export interface IpSecPolicyState {
    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     */
    authAlgorithm?: pulumi.Input<string>;
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     */
    encapsulationMode?: pulumi.Input<string>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    encryptionAlgorithm?: pulumi.Input<string>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     */
    lifetimes?: pulumi.Input<pulumi.Input<inputs.vpnaas.IpSecPolicyLifetime>[]>;
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     */
    name?: pulumi.Input<string>;
    /**
     * The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
     * is group5. Changing this updates the existing policy.
     */
    pfs?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     */
    region?: pulumi.Input<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The transform protocol. Valid values are esp, ah and ah-esp.
     * Changing this updates the existing policy. Default is ESP.
     */
    transformProtocol?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a IpSecPolicy resource.
 */
export interface IpSecPolicyArgs {
    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     */
    authAlgorithm?: pulumi.Input<string>;
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     */
    encapsulationMode?: pulumi.Input<string>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    encryptionAlgorithm?: pulumi.Input<string>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     */
    lifetimes?: pulumi.Input<pulumi.Input<inputs.vpnaas.IpSecPolicyLifetime>[]>;
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     */
    name?: pulumi.Input<string>;
    /**
     * The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
     * is group5. Changing this updates the existing policy.
     */
    pfs?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     */
    region?: pulumi.Input<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The transform protocol. Valid values are esp, ah and ah-esp.
     * Changing this updates the existing policy. Default is ESP.
     */
    transformProtocol?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
