// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get firewall policy information of an available OpenStack firewall policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const policy = pulumi.output(openstack.firewall.getPolicy({
 *     name: "tf_test_policy",
 * }, { async: true }));
 * ```
 */
export function getPolicy(args?: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:firewall/getPolicy:getPolicy", {
        "name": args.name,
        "policyId": args.policyId,
        "region": args.region,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    /**
     * The name of the firewall policy.
     */
    readonly name?: string;
    /**
     * The ID of the firewall policy.
     */
    readonly policyId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve firewall policy ids. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
    /**
     * The owner of the firewall policy.
     */
    readonly tenantId?: string;
}

/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    /**
     * The audit status of the firewall policy.
     */
    readonly audited: boolean;
    /**
     * The description of the firewall policy.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly name?: string;
    /**
     * See Argument Reference above.
     */
    readonly policyId?: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * The array of one or more firewall rules that comprise the policy.
     */
    readonly rules: string[];
    /**
     * The sharing status of the firewall policy.
     */
    readonly shared: boolean;
    /**
     * See Argument Reference above.
     */
    readonly tenantId: string;
}
