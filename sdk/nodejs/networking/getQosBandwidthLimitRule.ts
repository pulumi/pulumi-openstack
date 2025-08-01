// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack QoS bandwidth limit rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qosBandwidthLimitRule1 = openstack.networking.getQosBandwidthLimitRule({
 *     maxKbps: 300,
 * });
 * ```
 */
export function getQosBandwidthLimitRule(args: GetQosBandwidthLimitRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetQosBandwidthLimitRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:networking/getQosBandwidthLimitRule:getQosBandwidthLimitRule", {
        "maxBurstKbps": args.maxBurstKbps,
        "maxKbps": args.maxKbps,
        "qosPolicyId": args.qosPolicyId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getQosBandwidthLimitRule.
 */
export interface GetQosBandwidthLimitRuleArgs {
    /**
     * The maximum burst size in kilobits of a QoS bandwidth limit rule.
     */
    maxBurstKbps?: number;
    /**
     * The maximum kilobits per second of a QoS bandwidth limit rule.
     */
    maxKbps?: number;
    /**
     * The QoS policy reference.
     */
    qosPolicyId: string;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getQosBandwidthLimitRule.
 */
export interface GetQosBandwidthLimitRuleResult {
    /**
     * See Argument Reference above.
     */
    readonly direction: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly maxBurstKbps: number;
    /**
     * See Argument Reference above.
     */
    readonly maxKbps: number;
    /**
     * See Argument Reference above.
     */
    readonly qosPolicyId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
}
/**
 * Use this data source to get the ID of an available OpenStack QoS bandwidth limit rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qosBandwidthLimitRule1 = openstack.networking.getQosBandwidthLimitRule({
 *     maxKbps: 300,
 * });
 * ```
 */
export function getQosBandwidthLimitRuleOutput(args: GetQosBandwidthLimitRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetQosBandwidthLimitRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:networking/getQosBandwidthLimitRule:getQosBandwidthLimitRule", {
        "maxBurstKbps": args.maxBurstKbps,
        "maxKbps": args.maxKbps,
        "qosPolicyId": args.qosPolicyId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getQosBandwidthLimitRule.
 */
export interface GetQosBandwidthLimitRuleOutputArgs {
    /**
     * The maximum burst size in kilobits of a QoS bandwidth limit rule.
     */
    maxBurstKbps?: pulumi.Input<number>;
    /**
     * The maximum kilobits per second of a QoS bandwidth limit rule.
     */
    maxKbps?: pulumi.Input<number>;
    /**
     * The QoS policy reference.
     */
    qosPolicyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
