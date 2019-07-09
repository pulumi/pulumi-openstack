// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a v1 firewall resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const rule1 = new openstack.firewall.Rule("rule_1", {
 *     action: "deny",
 *     description: "drop TELNET traffic",
 *     destinationPort: "23",
 *     enabled: true,
 *     protocol: "tcp",
 * });
 * const rule2 = new openstack.firewall.Rule("rule_2", {
 *     action: "deny",
 *     description: "drop NTP traffic",
 *     destinationPort: "123",
 *     enabled: false,
 *     protocol: "udp",
 * });
 * const policy1 = new openstack.firewall.Policy("policy_1", {
 *     rules: [
 *         rule1.id,
 *         rule2.id,
 *     ],
 * });
 * const firewall1 = new openstack.firewall.Firewall("firewall_1", {
 *     policyId: policy1.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/fw_firewall_v1.html.markdown.
 */
export class Firewall extends pulumi.CustomResource {
    /**
     * Get an existing Firewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallState, opts?: pulumi.CustomResourceOptions): Firewall {
        return new Firewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:firewall/firewall:Firewall';

    /**
     * Returns true if the given object is an instance of Firewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Firewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Firewall.__pulumiType;
    }

    /**
     * Administrative up/down status for the firewall
     * (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `admin_state_up` of an existing firewall.
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     */
    public readonly associatedRouters!: pulumi.Output<string[]>;
    /**
     * A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Should this firewall not be associated with any routers
     * (must be "true" or "false" if provide - defaults to "false").
     * Conflicts with `associated_routers`.
     */
    public readonly noRouters!: pulumi.Output<boolean | undefined>;
    /**
     * The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Firewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallArgs | FirewallState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FirewallState | undefined;
            inputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            inputs["associatedRouters"] = state ? state.associatedRouters : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["noRouters"] = state ? state.noRouters : undefined;
            inputs["policyId"] = state ? state.policyId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as FirewallArgs | undefined;
            if (!args || args.policyId === undefined) {
                throw new Error("Missing required property 'policyId'");
            }
            inputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            inputs["associatedRouters"] = args ? args.associatedRouters : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["noRouters"] = args ? args.noRouters : undefined;
            inputs["policyId"] = args ? args.policyId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        super(Firewall.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Firewall resources.
 */
export interface FirewallState {
    /**
     * Administrative up/down status for the firewall
     * (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `admin_state_up` of an existing firewall.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     */
    readonly associatedRouters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Should this firewall not be associated with any routers
     * (must be "true" or "false" if provide - defaults to "false").
     * Conflicts with `associated_routers`.
     */
    readonly noRouters?: pulumi.Input<boolean>;
    /**
     * The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     */
    readonly policyId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Firewall resource.
 */
export interface FirewallArgs {
    /**
     * Administrative up/down status for the firewall
     * (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `admin_state_up` of an existing firewall.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     */
    readonly associatedRouters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Should this firewall not be associated with any routers
     * (must be "true" or "false" if provide - defaults to "false").
     * Conflicts with `associated_routers`.
     */
    readonly noRouters?: pulumi.Input<boolean>;
    /**
     * The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     */
    readonly policyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
