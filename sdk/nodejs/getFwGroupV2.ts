// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get information of an available OpenStack firewall group v2.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const group = openstack.firewall.getGroupV2({
 *     name: "tf_test_group",
 * });
 * ```
 */
/** @deprecated openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2 */
export function getFwGroupV2(args?: GetFwGroupV2Args, opts?: pulumi.InvokeOptions): Promise<GetFwGroupV2Result> {
    pulumi.log.warn("getFwGroupV2 is deprecated: openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:index/getFwGroupV2:getFwGroupV2", {
        "adminStateUp": args.adminStateUp,
        "description": args.description,
        "egressFirewallPolicyId": args.egressFirewallPolicyId,
        "groupId": args.groupId,
        "ingressFirewallPolicyId": args.ingressFirewallPolicyId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
        "shared": args.shared,
        "status": args.status,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFwGroupV2.
 */
export interface GetFwGroupV2Args {
    /**
     * Administrative up/down status for the firewall group.
     */
    adminStateUp?: boolean;
    /**
     * Human-readable description of the firewall group.
     */
    description?: string;
    /**
     * The egress policy ID of the firewall group.
     */
    egressFirewallPolicyId?: string;
    /**
     * The ID of the firewall group.
     */
    groupId?: string;
    /**
     * The ingress policy ID of the firewall group.
     */
    ingressFirewallPolicyId?: string;
    /**
     * The name of the firewall group.
     */
    name?: string;
    /**
     * This argument conflicts and is interchangeable
     * with `tenantId`. The owner of the firewall group.
     */
    projectId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve firewall group ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The sharing status of the firewall group.
     */
    shared?: boolean;
    /**
     * Enabled status for the firewall group.
     */
    status?: string;
    /**
     * This argument conflicts and is interchangeable
     * with `projectId`. The owner of the firewall group.
     */
    tenantId?: string;
}

/**
 * A collection of values returned by getFwGroupV2.
 */
export interface GetFwGroupV2Result {
    /**
     * See Argument Reference above.
     */
    readonly adminStateUp: boolean;
    /**
     * See Argument Reference above.
     */
    readonly description?: string;
    /**
     * See Argument Reference above.
     */
    readonly egressFirewallPolicyId?: string;
    /**
     * See Argument Reference above.
     */
    readonly groupId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly ingressFirewallPolicyId?: string;
    /**
     * See Argument Reference above.
     */
    readonly name?: string;
    /**
     * Ports associated with the firewall group.
     */
    readonly ports: string[];
    /**
     * See Argument Reference above.
     */
    readonly projectId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * See Argument Reference above.
     */
    readonly shared: boolean;
    /**
     * See Argument Reference above.
     */
    readonly status: string;
    /**
     * See Argument Reference above.
     */
    readonly tenantId: string;
}
/**
 * Use this data source to get information of an available OpenStack firewall group v2.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const group = openstack.firewall.getGroupV2({
 *     name: "tf_test_group",
 * });
 * ```
 */
/** @deprecated openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2 */
export function getFwGroupV2Output(args?: GetFwGroupV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFwGroupV2Result> {
    pulumi.log.warn("getFwGroupV2 is deprecated: openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2")
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:index/getFwGroupV2:getFwGroupV2", {
        "adminStateUp": args.adminStateUp,
        "description": args.description,
        "egressFirewallPolicyId": args.egressFirewallPolicyId,
        "groupId": args.groupId,
        "ingressFirewallPolicyId": args.ingressFirewallPolicyId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
        "shared": args.shared,
        "status": args.status,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getFwGroupV2.
 */
export interface GetFwGroupV2OutputArgs {
    /**
     * Administrative up/down status for the firewall group.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description of the firewall group.
     */
    description?: pulumi.Input<string>;
    /**
     * The egress policy ID of the firewall group.
     */
    egressFirewallPolicyId?: pulumi.Input<string>;
    /**
     * The ID of the firewall group.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The ingress policy ID of the firewall group.
     */
    ingressFirewallPolicyId?: pulumi.Input<string>;
    /**
     * The name of the firewall group.
     */
    name?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `tenantId`. The owner of the firewall group.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve firewall group ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The sharing status of the firewall group.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * Enabled status for the firewall group.
     */
    status?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable
     * with `projectId`. The owner of the firewall group.
     */
    tenantId?: pulumi.Input<string>;
}
