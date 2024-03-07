// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information of an available OpenStack firewall rule v2.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const rule = openstack.getFwRuleV2({
 *     name: "tf_test_rule",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getFwRuleV2(args?: GetFwRuleV2Args, opts?: pulumi.InvokeOptions): Promise<GetFwRuleV2Result> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:index/getFwRuleV2:getFwRuleV2", {
        "action": args.action,
        "description": args.description,
        "destinationIpAddress": args.destinationIpAddress,
        "destinationPort": args.destinationPort,
        "enabled": args.enabled,
        "firewallPolicyIds": args.firewallPolicyIds,
        "ipVersion": args.ipVersion,
        "name": args.name,
        "projectId": args.projectId,
        "protocol": args.protocol,
        "region": args.region,
        "ruleId": args.ruleId,
        "shared": args.shared,
        "sourceIpAddress": args.sourceIpAddress,
        "sourcePort": args.sourcePort,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFwRuleV2.
 */
export interface GetFwRuleV2Args {
    /**
     * Action to be taken when the firewall rule matches.
     */
    action?: string;
    /**
     * The description of the firewall rule.
     */
    description?: string;
    /**
     * The destination IP address on which the
     * firewall rule operates.
     */
    destinationIpAddress?: string;
    /**
     * The destination port on which the firewall
     * rule operates.
     */
    destinationPort?: string;
    /**
     * Enabled status for the firewall rule.
     */
    enabled?: boolean;
    /**
     * The ID of the firewall policy the rule belongs to.
     */
    firewallPolicyIds?: string[];
    /**
     * IP version, either 4 (default) or 6.
     */
    ipVersion?: number;
    /**
     * The name of the firewall rule.
     */
    name?: string;
    /**
     * This argument conflicts and is interchangeable
     * with `tenantId`. The owner of the firewall rule.
     */
    projectId?: string;
    /**
     * The protocol type on which the firewall rule operates.
     */
    protocol?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve firewall policy ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The ID of the firewall rule.
     */
    ruleId?: string;
    /**
     * The sharing status of the firewall policy.
     */
    shared?: boolean;
    /**
     * The source IP address on which the firewall
     * rule operates.
     */
    sourceIpAddress?: string;
    /**
     * The source port on which the firewall
     * rule operates.
     */
    sourcePort?: string;
    /**
     * This argument conflicts and is interchangeable
     * with `projectId`. The owner of the firewall rule.
     */
    tenantId?: string;
}

/**
 * A collection of values returned by getFwRuleV2.
 */
export interface GetFwRuleV2Result {
    /**
     * See Argument Reference above.
     */
    readonly action?: string;
    /**
     * See Argument Reference above.
     */
    readonly description?: string;
    /**
     * See Argument Reference above.
     */
    readonly destinationIpAddress?: string;
    /**
     * See Argument Reference above.
     */
    readonly destinationPort?: string;
    /**
     * See Argument Reference above.
     */
    readonly enabled: boolean;
    /**
     * The ID of the firewall policy the rule belongs to.
     */
    readonly firewallPolicyIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly ipVersion?: number;
    /**
     * See Argument Reference above.
     */
    readonly name?: string;
    /**
     * See Argument Reference above.
     */
    readonly projectId: string;
    /**
     * See Argument Reference above.
     */
    readonly protocol?: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * See Argument Reference above.
     */
    readonly ruleId?: string;
    /**
     * See Argument Reference above.
     */
    readonly shared: boolean;
    /**
     * See Argument Reference above.
     */
    readonly sourceIpAddress?: string;
    /**
     * See Argument Reference above.
     */
    readonly sourcePort?: string;
    /**
     * See Argument Reference above.
     */
    readonly tenantId: string;
}
/**
 * Use this data source to get information of an available OpenStack firewall rule v2.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const rule = openstack.getFwRuleV2({
 *     name: "tf_test_rule",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getFwRuleV2Output(args?: GetFwRuleV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFwRuleV2Result> {
    return pulumi.output(args).apply((a: any) => getFwRuleV2(a, opts))
}

/**
 * A collection of arguments for invoking getFwRuleV2.
 */
export interface GetFwRuleV2OutputArgs {
    /**
     * Action to be taken when the firewall rule matches.
     */
    action?: pulumi.Input<string>;
    /**
     * The description of the firewall rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The destination IP address on which the
     * firewall rule operates.
     */
    destinationIpAddress?: pulumi.Input<string>;
    /**
     * The destination port on which the firewall
     * rule operates.
     */
    destinationPort?: pulumi.Input<string>;
    /**
     * Enabled status for the firewall rule.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The ID of the firewall policy the rule belongs to.
     */
    firewallPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IP version, either 4 (default) or 6.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * The name of the firewall rule.
     */
    name?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `tenantId`. The owner of the firewall rule.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The protocol type on which the firewall rule operates.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve firewall policy ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of the firewall rule.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * The sharing status of the firewall policy.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * The source IP address on which the firewall
     * rule operates.
     */
    sourceIpAddress?: pulumi.Input<string>;
    /**
     * The source port on which the firewall
     * rule operates.
     */
    sourcePort?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `projectId`. The owner of the firewall rule.
     */
    tenantId?: pulumi.Input<string>;
}
