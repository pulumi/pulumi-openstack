// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack security group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const secgroup = pulumi.output(openstack.networking.getSecGroup({
 *     name: "tf_test_secgroup",
 * }));
 * ```
 */
export function getSecGroup(args?: GetSecGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSecGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("openstack:networking/getSecGroup:getSecGroup", {
        "description": args.description,
        "name": args.name,
        "region": args.region,
        "secgroupId": args.secgroupId,
        "tags": args.tags,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecGroup.
 */
export interface GetSecGroupArgs {
    /**
     * Human-readable description the the subnet.
     */
    description?: string;
    /**
     * The name of the security group.
     */
    name?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve security groups ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The ID of the security group.
     */
    secgroupId?: string;
    /**
     * The list of security group tags to filter.
     */
    tags?: string[];
    /**
     * The owner of the security group.
     */
    tenantId?: string;
}

/**
 * A collection of values returned by getSecGroup.
 */
export interface GetSecGroupResult {
    /**
     * The set of string tags applied on the security group.
     */
    readonly allTags: string[];
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     * * `description`- See Argument Reference above.
     */
    readonly name?: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    readonly secgroupId?: string;
    readonly tags?: string[];
    readonly tenantId: string;
}

export function getSecGroupOutput(args?: GetSecGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecGroupResult> {
    return pulumi.output(args).apply(a => getSecGroup(a, opts))
}

/**
 * A collection of arguments for invoking getSecGroup.
 */
export interface GetSecGroupOutputArgs {
    /**
     * Human-readable description the the subnet.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the security group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve security groups ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of the security group.
     */
    secgroupId?: pulumi.Input<string>;
    /**
     * The list of security group tags to filter.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the security group.
     */
    tenantId?: pulumi.Input<string>;
}
