// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available Barbican container.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const example = openstack.keymanager.getContainer({
 *     name: "my_container",
 * });
 * ```
 */
export function getContainer(args?: GetContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetContainerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:keymanager/getContainer:getContainer", {
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainer.
 */
export interface GetContainerArgs {
    /**
     * The Container name.
     */
    name?: string;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a container. If omitted, the `region`
     * argument of the provider is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getContainer.
 */
export interface GetContainerResult {
    /**
     * The list of ACLs assigned to a container. The `read` structure is
     * described below.
     */
    readonly acls: outputs.keymanager.GetContainerAcl[];
    /**
     * The list of the container consumers. The structure is described
     * below.
     */
    readonly consumers: outputs.keymanager.GetContainerConsumer[];
    /**
     * The container reference / where to find the container.
     */
    readonly containerRef: string;
    /**
     * The date the container ACL was created.
     */
    readonly createdAt: string;
    /**
     * The creator of the container.
     */
    readonly creatorId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the consumer.
     */
    readonly name?: string;
    /**
     * See Argument Reference above.
     */
    readonly region?: string;
    /**
     * A set of dictionaries containing references to secrets. The
     * structure is described below.
     */
    readonly secretRefs: outputs.keymanager.GetContainerSecretRef[];
    /**
     * The status of the container.
     */
    readonly status: string;
    /**
     * The container type.
     */
    readonly type: string;
    /**
     * The date the container ACL was last updated.
     */
    readonly updatedAt: string;
}
/**
 * Use this data source to get the ID of an available Barbican container.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const example = openstack.keymanager.getContainer({
 *     name: "my_container",
 * });
 * ```
 */
export function getContainerOutput(args?: GetContainerOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetContainerResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:keymanager/getContainer:getContainer", {
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getContainer.
 */
export interface GetContainerOutputArgs {
    /**
     * The Container name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a container. If omitted, the `region`
     * argument of the provider is used.
     */
    region?: pulumi.Input<string>;
}
