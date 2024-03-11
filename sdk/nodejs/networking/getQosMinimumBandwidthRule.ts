// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qosMinBwRule1 = openstack.networking.getQosMinimumBandwidthRule({
 *     minKbps: 2000,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getQosMinimumBandwidthRule(args: GetQosMinimumBandwidthRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetQosMinimumBandwidthRuleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:networking/getQosMinimumBandwidthRule:getQosMinimumBandwidthRule", {
        "direction": args.direction,
        "minKbps": args.minKbps,
        "qosPolicyId": args.qosPolicyId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getQosMinimumBandwidthRule.
 */
export interface GetQosMinimumBandwidthRuleArgs {
    direction?: string;
    /**
     * The value of a minimum kbps bandwidth.
     */
    minKbps?: number;
    /**
     * The QoS policy reference.
     */
    qosPolicyId: string;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getQosMinimumBandwidthRule.
 */
export interface GetQosMinimumBandwidthRuleResult {
    readonly direction: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
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
}
/**
 * Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qosMinBwRule1 = openstack.networking.getQosMinimumBandwidthRule({
 *     minKbps: 2000,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getQosMinimumBandwidthRuleOutput(args: GetQosMinimumBandwidthRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQosMinimumBandwidthRuleResult> {
    return pulumi.output(args).apply((a: any) => getQosMinimumBandwidthRule(a, opts))
}

/**
 * A collection of arguments for invoking getQosMinimumBandwidthRule.
 */
export interface GetQosMinimumBandwidthRuleOutputArgs {
    direction?: pulumi.Input<string>;
    /**
     * The value of a minimum kbps bandwidth.
     */
    minKbps?: pulumi.Input<number>;
    /**
     * The QoS policy reference.
     */
    qosPolicyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
