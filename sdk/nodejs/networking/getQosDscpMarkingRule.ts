// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const qosDscpMarkingRule1 = pulumi.output(openstack.networking.getQosDscpMarkingRule({
 *     dscpMark: 26,
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_dscp_marking_rule_v2.html.markdown.
 */
export function getQosDscpMarkingRule(args: GetQosDscpMarkingRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetQosDscpMarkingRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
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
    readonly dscpMark?: number;
    /**
     * The QoS policy reference.
     */
    readonly qosPolicyId: string;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
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
     * See Argument Reference above.
     */
    readonly qosPolicyId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
