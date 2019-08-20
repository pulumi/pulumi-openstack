// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const qosMinBwRule1 = openstack.networking.getQosMinimumBandwidthRule({
 *     minKbps: 2000,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_minimum_bandwidth_rule_v2.html.markdown.
 */
export function getQosMinimumBandwidthRule(args: GetQosMinimumBandwidthRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetQosMinimumBandwidthRuleResult> & GetQosMinimumBandwidthRuleResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetQosMinimumBandwidthRuleResult> = pulumi.runtime.invoke("openstack:networking/getQosMinimumBandwidthRule:getQosMinimumBandwidthRule", {
        "direction": args.direction,
        "minKbps": args.minKbps,
        "qosPolicyId": args.qosPolicyId,
        "region": args.region,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getQosMinimumBandwidthRule.
 */
export interface GetQosMinimumBandwidthRuleArgs {
    readonly direction?: string;
    /**
     * The value of a minimum kbps bandwidth.
     */
    readonly minKbps?: number;
    /**
     * The QoS policy reference.
     */
    readonly qosPolicyId: string;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getQosMinimumBandwidthRule.
 */
export interface GetQosMinimumBandwidthRuleResult {
    readonly direction: string;
    /**
     * See Argument Reference above.
     */
    readonly minKbps: number;
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
