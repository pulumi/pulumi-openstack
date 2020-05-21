// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron QoS minimum bandwidth rule resource within OpenStack.
 *
 * ## Example Usage
 *
 * ### Create a QoS Policy with some minimum bandwidth rule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qosPolicy1 = new openstack.networking.QosPolicy("qosPolicy1", {
 *     description: "minKbps",
 * });
 * const minimumBandwidthRule1 = new openstack.networking.QosMinimumBandwidthRule("minimumBandwidthRule1", {
 *     minKbps: 200,
 *     qosPolicyId: qosPolicy1.id,
 * });
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as QosMinimumBandwidthRuleState | undefined;
            inputs["direction"] = state ? state.direction : undefined;
            inputs["minKbps"] = state ? state.minKbps : undefined;
            inputs["qosPolicyId"] = state ? state.qosPolicyId : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as QosMinimumBandwidthRuleArgs | undefined;
            if (!args || args.minKbps === undefined) {
                throw new Error("Missing required property 'minKbps'");
            }
            if (!args || args.qosPolicyId === undefined) {
                throw new Error("Missing required property 'qosPolicyId'");
            }
            inputs["direction"] = args ? args.direction : undefined;
            inputs["minKbps"] = args ? args.minKbps : undefined;
            inputs["qosPolicyId"] = args ? args.qosPolicyId : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(QosMinimumBandwidthRule.__pulumiType, name, inputs, opts);
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
    readonly direction?: pulumi.Input<string>;
    /**
     * The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     */
    readonly minKbps?: pulumi.Input<number>;
    /**
     * The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     */
    readonly qosPolicyId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a QosMinimumBandwidthRule resource.
 */
export interface QosMinimumBandwidthRuleArgs {
    /**
     * The direction of traffic. Defaults to "egress". Changing this updates the direction of the
     * existing QoS minimum bandwidth rule.
     */
    readonly direction?: pulumi.Input<string>;
    /**
     * The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     */
    readonly minKbps: pulumi.Input<number>;
    /**
     * The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     */
    readonly qosPolicyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     */
    readonly region?: pulumi.Input<string>;
}
