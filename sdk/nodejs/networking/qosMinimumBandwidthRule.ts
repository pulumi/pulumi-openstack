// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron QoS minimum bandwidth rule resource within OpenStack.
 *
 * ## Example Usage
 * ### Create a QoS Policy with some minimum bandwidth rule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qosPolicy1 = new openstack.networking.QosPolicy("qos_policy_1", {
 *     description: "min_kbps",
 * });
 * const minimumBandwidthRule1 = new openstack.networking.QosMinimumBandwidthRule("minimum_bandwidth_rule_1", {
 *     minKbps: 200,
 *     qosPolicyId: qosPolicy1.id,
 * });
 * ```
 *
 * ## Import
 *
 * QoS minimum bandwidth rules can be imported using the `qos_policy_id/minimum_bandwidth_rule_id` format, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule minimum_bandwidth_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
 * ```
 */
export class QosMinimumBandwidthRule extends pulumi.CustomResource {
    /**
     * Get an existing QosMinimumBandwidthRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QosMinimumBandwidthRuleState, opts?: pulumi.CustomResourceOptions): QosMinimumBandwidthRule {
        return new QosMinimumBandwidthRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule';

    /**
     * Returns true if the given object is an instance of QosMinimumBandwidthRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QosMinimumBandwidthRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QosMinimumBandwidthRule.__pulumiType;
    }

    /**
     * The direction of traffic. Defaults to "egress". Changing this updates the direction of the
     * existing QoS minimum bandwidth rule.
     */
    public readonly direction!: pulumi.Output<string | undefined>;
    /**
     * The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     */
    public readonly minKbps!: pulumi.Output<number>;
    /**
     * The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     */
    public readonly qosPolicyId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a QosMinimumBandwidthRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QosMinimumBandwidthRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QosMinimumBandwidthRuleArgs | QosMinimumBandwidthRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QosMinimumBandwidthRuleState | undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["minKbps"] = state ? state.minKbps : undefined;
            resourceInputs["qosPolicyId"] = state ? state.qosPolicyId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as QosMinimumBandwidthRuleArgs | undefined;
            if ((!args || args.minKbps === undefined) && !opts.urn) {
                throw new Error("Missing required property 'minKbps'");
            }
            if ((!args || args.qosPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'qosPolicyId'");
            }
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["minKbps"] = args ? args.minKbps : undefined;
            resourceInputs["qosPolicyId"] = args ? args.qosPolicyId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(QosMinimumBandwidthRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QosMinimumBandwidthRule resources.
 */
export interface QosMinimumBandwidthRuleState {
    /**
     * The direction of traffic. Defaults to "egress". Changing this updates the direction of the
     * existing QoS minimum bandwidth rule.
     */
    direction?: pulumi.Input<string>;
    /**
     * The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     */
    minKbps?: pulumi.Input<number>;
    /**
     * The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     */
    qosPolicyId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a QosMinimumBandwidthRule resource.
 */
export interface QosMinimumBandwidthRuleArgs {
    /**
     * The direction of traffic. Defaults to "egress". Changing this updates the direction of the
     * existing QoS minimum bandwidth rule.
     */
    direction?: pulumi.Input<string>;
    /**
     * The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     */
    minKbps: pulumi.Input<number>;
    /**
     * The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     */
    qosPolicyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     */
    region?: pulumi.Input<string>;
}
