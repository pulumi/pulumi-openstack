// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
 * const qosBandwidthLimitRule1 = pulumi.output(openstack.networking.getQosBandwidthLimitRule({
 *     maxKbps: 300,
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_bandwidth_limit_rule_v2.html.markdown.
 */
export function getQosBandwidthLimitRule(args: GetQosBandwidthLimitRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetQosBandwidthLimitRuleResult> {
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
    readonly maxBurstKbps?: number;
    /**
     * The maximum kilobits per second of a QoS bandwidth limit rule.
     */
    readonly maxKbps?: number;
    /**
     * The QoS policy reference.
     */
    readonly qosPolicyId: string;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
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
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
