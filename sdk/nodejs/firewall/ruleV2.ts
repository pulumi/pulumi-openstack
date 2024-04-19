// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a v2 firewall rule resource within OpenStack.
 *
 * > **Note:** Firewall v2 has no support for OVN currently.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const rule2 = new openstack.firewall.RuleV2("rule2", {
 *     action: "deny",
 *     description: "drop TELNET traffic",
 *     destinationPort: "23",
 *     enabled: true,
 *     protocol: "tcp",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall Rules can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:firewall/ruleV2:RuleV2 rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
 * ```
 */
export class RuleV2 extends pulumi.CustomResource {
    /**
     * Get an existing RuleV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleV2State, opts?: pulumi.CustomResourceOptions): RuleV2 {
        return new RuleV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:firewall/ruleV2:RuleV2';

    /**
     * Returns true if the given object is an instance of RuleV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RuleV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RuleV2.__pulumiType;
    }

    /**
     * Action to be taken (must be "allow", "deny" or "reject")
     * when the firewall rule matches. Changing this updates the `action` of an
     * existing firewall rule. Default is `deny`.
     */
    public readonly action!: pulumi.Output<string | undefined>;
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
     * firewall rule. Require not `any` or empty protocol.
     */
    public readonly destinationPort!: pulumi.Output<string | undefined>;
    /**
     * Enabled status for the firewall rule (must be "true"
     * or "false" if provided - defaults to "true"). Changing this updates the
     * `enabled` status of an existing firewall rule.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * IP version, either 4 or 6. Changing this
     * updates the `ipVersion` of an existing firewall rule. Default is `4`.
     */
    public readonly ipVersion!: pulumi.Output<number | undefined>;
    /**
     * A unique name for the firewall rule. Changing this
     * updates the `name` of an existing firewall rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `tenantId`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another project. Changing this creates a new
     * firewall rule.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * (Optional; Required if `sourcePort` or `destinationPort` is not
     * empty) The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule. Default is `any`.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Sharing status of the firewall rule (must be "true"
     * or "false" if provided). If this is "true" the policy is visible to, and
     * can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall policy. On
     */
    public readonly shared!: pulumi.Output<boolean | undefined>;
    /**
     * The source IP address on which the firewall
     * rule operates. Changing this updates the `sourceIpAddress` of an existing
     * firewall rule.
     */
    public readonly sourceIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The source port on which the firewall
     * rule operates. Changing this updates the `sourcePort` of an existing
     * firewall rule. Require not `any` or empty protocol.
     */
    public readonly sourcePort!: pulumi.Output<string | undefined>;
    /**
     * This argument conflicts and is interchangeable
     * with `projectId`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another tenant. Changing this creates a new
     * firewall rule.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a RuleV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RuleV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleV2Args | RuleV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuleV2State | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationIpAddress"] = state ? state.destinationIpAddress : undefined;
            resourceInputs["destinationPort"] = state ? state.destinationPort : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["sourceIpAddress"] = state ? state.sourceIpAddress : undefined;
            resourceInputs["sourcePort"] = state ? state.sourcePort : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as RuleV2Args | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationIpAddress"] = args ? args.destinationIpAddress : undefined;
            resourceInputs["destinationPort"] = args ? args.destinationPort : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["sourceIpAddress"] = args ? args.sourceIpAddress : undefined;
            resourceInputs["sourcePort"] = args ? args.sourcePort : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RuleV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RuleV2 resources.
 */
export interface RuleV2State {
    /**
     * Action to be taken (must be "allow", "deny" or "reject")
     * when the firewall rule matches. Changing this updates the `action` of an
     * existing firewall rule. Default is `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * A description for the firewall rule. Changing this
     * updates the `description` of an existing firewall rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination IP address on which the
     * firewall rule operates. Changing this updates the `destinationIpAddress`
     * of an existing firewall rule.
     */
    destinationIpAddress?: pulumi.Input<string>;
    /**
     * The destination port on which the firewall
     * rule operates. Changing this updates the `destinationPort` of an existing
     * firewall rule. Require not `any` or empty protocol.
     */
    destinationPort?: pulumi.Input<string>;
    /**
     * Enabled status for the firewall rule (must be "true"
     * or "false" if provided - defaults to "true"). Changing this updates the
     * `enabled` status of an existing firewall rule.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * IP version, either 4 or 6. Changing this
     * updates the `ipVersion` of an existing firewall rule. Default is `4`.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * A unique name for the firewall rule. Changing this
     * updates the `name` of an existing firewall rule.
     */
    name?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `tenantId`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another project. Changing this creates a new
     * firewall rule.
     */
    projectId?: pulumi.Input<string>;
    /**
     * (Optional; Required if `sourcePort` or `destinationPort` is not
     * empty) The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule. Default is `any`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     */
    region?: pulumi.Input<string>;
    /**
     * Sharing status of the firewall rule (must be "true"
     * or "false" if provided). If this is "true" the policy is visible to, and
     * can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall policy. On
     */
    shared?: pulumi.Input<boolean>;
    /**
     * The source IP address on which the firewall
     * rule operates. Changing this updates the `sourceIpAddress` of an existing
     * firewall rule.
     */
    sourceIpAddress?: pulumi.Input<string>;
    /**
     * The source port on which the firewall
     * rule operates. Changing this updates the `sourcePort` of an existing
     * firewall rule. Require not `any` or empty protocol.
     */
    sourcePort?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `projectId`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another tenant. Changing this creates a new
     * firewall rule.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RuleV2 resource.
 */
export interface RuleV2Args {
    /**
     * Action to be taken (must be "allow", "deny" or "reject")
     * when the firewall rule matches. Changing this updates the `action` of an
     * existing firewall rule. Default is `deny`.
     */
    action?: pulumi.Input<string>;
    /**
     * A description for the firewall rule. Changing this
     * updates the `description` of an existing firewall rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination IP address on which the
     * firewall rule operates. Changing this updates the `destinationIpAddress`
     * of an existing firewall rule.
     */
    destinationIpAddress?: pulumi.Input<string>;
    /**
     * The destination port on which the firewall
     * rule operates. Changing this updates the `destinationPort` of an existing
     * firewall rule. Require not `any` or empty protocol.
     */
    destinationPort?: pulumi.Input<string>;
    /**
     * Enabled status for the firewall rule (must be "true"
     * or "false" if provided - defaults to "true"). Changing this updates the
     * `enabled` status of an existing firewall rule.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * IP version, either 4 or 6. Changing this
     * updates the `ipVersion` of an existing firewall rule. Default is `4`.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * A unique name for the firewall rule. Changing this
     * updates the `name` of an existing firewall rule.
     */
    name?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `tenantId`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another project. Changing this creates a new
     * firewall rule.
     */
    projectId?: pulumi.Input<string>;
    /**
     * (Optional; Required if `sourcePort` or `destinationPort` is not
     * empty) The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule. Default is `any`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     */
    region?: pulumi.Input<string>;
    /**
     * Sharing status of the firewall rule (must be "true"
     * or "false" if provided). If this is "true" the policy is visible to, and
     * can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall policy. On
     */
    shared?: pulumi.Input<boolean>;
    /**
     * The source IP address on which the firewall
     * rule operates. Changing this updates the `sourceIpAddress` of an existing
     * firewall rule.
     */
    sourceIpAddress?: pulumi.Input<string>;
    /**
     * The source port on which the firewall
     * rule operates. Changing this updates the `sourcePort` of an existing
     * firewall rule. Require not `any` or empty protocol.
     */
    sourcePort?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `projectId`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another tenant. Changing this creates a new
     * firewall rule.
     */
    tenantId?: pulumi.Input<string>;
}
