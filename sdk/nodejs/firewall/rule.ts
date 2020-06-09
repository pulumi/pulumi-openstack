// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a v1 firewall rule resource within OpenStack.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const rule1 = new openstack.firewall.Rule("rule1", {
 *     action: "deny",
 *     description: "drop TELNET traffic",
 *     destinationPort: "23",
 *     enabled: true,
 *     protocol: "tcp",
 * });
 * ```
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:firewall/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * Action to be taken ( must be "allow" or "deny") when the
     * firewall rule matches. Changing this updates the `action` of an existing
     * firewall rule.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * A description for the firewall rule. Changing this
     * updates the `description` of an existing firewall rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination IP address on which the
     * firewall rule operates. Changing this updates the `destinationIpAddress`
     * of an existing firewall rule.
     */
    public readonly destinationIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The destination port on which the firewall
     * rule operates. Changing this updates the `destinationPort` of an existing
     * firewall rule.
     */
    public readonly destinationPort!: pulumi.Output<string | undefined>;
    /**
     * Enabled status for the firewall rule (must be "true"
     * or "false" if provided - defaults to "true"). Changing this updates the
     * `enabled` status of an existing firewall rule.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * IP version, either 4 (default) or 6. Changing this
     * updates the `ipVersion` of an existing firewall rule.
     */
    public readonly ipVersion!: pulumi.Output<number | undefined>;
    /**
     * A unique name for the firewall rule. Changing this
     * updates the `name` of an existing firewall rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The region in which to obtain the v1 Compute client.
     * A Compute client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The source IP address on which the firewall
     * rule operates. Changing this updates the `sourceIpAddress` of an existing
     * firewall rule.
     */
    public readonly sourceIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The source port on which the firewall
     * rule operates. Changing this updates the `sourcePort` of an existing
     * firewall rule.
     */
    public readonly sourcePort!: pulumi.Output<string | undefined>;
    /**
     * The owner of the firewall rule. Required if admin
     * wants to create a firewall rule for another tenant. Changing this creates a
     * new firewall rule.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RuleState | undefined;
            inputs["action"] = state ? state.action : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destinationIpAddress"] = state ? state.destinationIpAddress : undefined;
            inputs["destinationPort"] = state ? state.destinationPort : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["ipVersion"] = state ? state.ipVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["sourceIpAddress"] = state ? state.sourceIpAddress : undefined;
            inputs["sourcePort"] = state ? state.sourcePort : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if (!args || args.action === undefined) {
                throw new Error("Missing required property 'action'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationIpAddress"] = args ? args.destinationIpAddress : undefined;
            inputs["destinationPort"] = args ? args.destinationPort : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["sourceIpAddress"] = args ? args.sourceIpAddress : undefined;
            inputs["sourcePort"] = args ? args.sourcePort : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Rule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * Action to be taken ( must be "allow" or "deny") when the
     * firewall rule matches. Changing this updates the `action` of an existing
     * firewall rule.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * A description for the firewall rule. Changing this
     * updates the `description` of an existing firewall rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination IP address on which the
     * firewall rule operates. Changing this updates the `destinationIpAddress`
     * of an existing firewall rule.
     */
    readonly destinationIpAddress?: pulumi.Input<string>;
    /**
     * The destination port on which the firewall
     * rule operates. Changing this updates the `destinationPort` of an existing
     * firewall rule.
     */
    readonly destinationPort?: pulumi.Input<string>;
    /**
     * Enabled status for the firewall rule (must be "true"
     * or "false" if provided - defaults to "true"). Changing this updates the
     * `enabled` status of an existing firewall rule.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * IP version, either 4 (default) or 6. Changing this
     * updates the `ipVersion` of an existing firewall rule.
     */
    readonly ipVersion?: pulumi.Input<number>;
    /**
     * A unique name for the firewall rule. Changing this
     * updates the `name` of an existing firewall rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the v1 Compute client.
     * A Compute client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The source IP address on which the firewall
     * rule operates. Changing this updates the `sourceIpAddress` of an existing
     * firewall rule.
     */
    readonly sourceIpAddress?: pulumi.Input<string>;
    /**
     * The source port on which the firewall
     * rule operates. Changing this updates the `sourcePort` of an existing
     * firewall rule.
     */
    readonly sourcePort?: pulumi.Input<string>;
    /**
     * The owner of the firewall rule. Required if admin
     * wants to create a firewall rule for another tenant. Changing this creates a
     * new firewall rule.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * Action to be taken ( must be "allow" or "deny") when the
     * firewall rule matches. Changing this updates the `action` of an existing
     * firewall rule.
     */
    readonly action: pulumi.Input<string>;
    /**
     * A description for the firewall rule. Changing this
     * updates the `description` of an existing firewall rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination IP address on which the
     * firewall rule operates. Changing this updates the `destinationIpAddress`
     * of an existing firewall rule.
     */
    readonly destinationIpAddress?: pulumi.Input<string>;
    /**
     * The destination port on which the firewall
     * rule operates. Changing this updates the `destinationPort` of an existing
     * firewall rule.
     */
    readonly destinationPort?: pulumi.Input<string>;
    /**
     * Enabled status for the firewall rule (must be "true"
     * or "false" if provided - defaults to "true"). Changing this updates the
     * `enabled` status of an existing firewall rule.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * IP version, either 4 (default) or 6. Changing this
     * updates the `ipVersion` of an existing firewall rule.
     */
    readonly ipVersion?: pulumi.Input<number>;
    /**
     * A unique name for the firewall rule. Changing this
     * updates the `name` of an existing firewall rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The region in which to obtain the v1 Compute client.
     * A Compute client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The source IP address on which the firewall
     * rule operates. Changing this updates the `sourceIpAddress` of an existing
     * firewall rule.
     */
    readonly sourceIpAddress?: pulumi.Input<string>;
    /**
     * The source port on which the firewall
     * rule operates. Changing this updates the `sourcePort` of an existing
     * firewall rule.
     */
    readonly sourcePort?: pulumi.Input<string>;
    /**
     * The owner of the firewall rule. Required if admin
     * wants to create a firewall rule for another tenant. Changing this creates a
     * new firewall rule.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
