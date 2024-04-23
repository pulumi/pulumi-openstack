// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qosDscpMarkingRule1 = openstack.networking.getQosDscpMarkingRule({
 *     dscpMark: 26,
 * });
 * ```
 */
export function getQosDscpMarkingRule(args: GetQosDscpMarkingRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetQosDscpMarkingRuleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:networking/getQosDscpMarkingRule:getQosDscpMarkingRule", {
        "dscpMark": args.dscpMark,
        "qosPolicyId": args.qosPolicyId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getQosDscpMarkingRule.
 */
export interface GetQosDscpMarkingRuleArgs {
    /**
     * The value of a DSCP mark.
     */
    dscpMark?: number;
    /**
     * The QoS policy reference.
     */
    qosPolicyId: string;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getQosDscpMarkingRule.
 */
export interface GetQosDscpMarkingRuleResult {
    /**
     * See Argument Reference above.
     */
    readonly dscpMark: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
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
 * Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qosDscpMarkingRule1 = openstack.networking.getQosDscpMarkingRule({
 *     dscpMark: 26,
 * });
 * ```
 */
export function getQosDscpMarkingRuleOutput(args: GetQosDscpMarkingRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQosDscpMarkingRuleResult> {
    return pulumi.output(args).apply((a: any) => getQosDscpMarkingRule(a, opts))
}

/**
 * A collection of arguments for invoking getQosDscpMarkingRule.
 */
export interface GetQosDscpMarkingRuleOutputArgs {
    /**
     * The value of a DSCP mark.
     */
    dscpMark?: pulumi.Input<number>;
    /**
     * The QoS policy reference.
     */
    qosPolicyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
