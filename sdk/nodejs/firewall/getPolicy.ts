// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
 * const policy = openstack.firewall.getPolicy({
 *     name: "tfTestPolicy",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/fw_policy_v1.html.markdown.
 */
export function getPolicy(args?: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> & GetPolicyResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetPolicyResult> = pulumi.runtime.invoke("openstack:firewall/getPolicy:getPolicy", {
        "name": args.name,
        "policyId": args.policyId,
        "region": args.region,
        "tenantId": args.tenantId,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
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
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
