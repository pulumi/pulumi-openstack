import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 Neutron IKE policy resource within OpenStack.
 */
export declare class IkePolicy extends pulumi.CustomResource {
    /**
     * Get an existing IkePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IkePolicyState): IkePolicy;
    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     */
    readonly authAlgorithm: pulumi.Output<string | undefined>;
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     */
    readonly encryptionAlgorithm: pulumi.Output<string | undefined>;
    /**
     * The IKE mode. A valid value is v1 or v2. Default is v1.
     * Changing this updates the existing policy.
     */
    readonly ikeVersion: pulumi.Output<string | undefined>;
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
     * Default is seconds.
     * - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
     * Default is 3600.
     */
    readonly lifetimes: pulumi.Output<{
        units: string;
        value: number;
    }[]>;
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
     * Changing this updates the existing policy.
     */
    readonly pfs: pulumi.Output<string | undefined>;
    /**
     * The IKE mode. A valid value is main, which is the default.
     * Changing this updates the existing policy.
     */
    readonly phase1NegotiationMode: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The owner of the policy. Required if admin wants to
     * create a service for another policy. Changing this creates a new policy.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a IkePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IkePolicyArgs, opts?: pulumi.ResourceOptions);
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
    readonly lifetimes?: pulumi.Input<{
        units?: pulumi.Input<string>;
        value?: pulumi.Input<number>;
    }[]>;
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
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    readonly lifetimes?: pulumi.Input<{
        units?: pulumi.Input<string>;
        value?: pulumi.Input<number>;
    }[]>;
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
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
