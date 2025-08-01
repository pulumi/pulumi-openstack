// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an OpenStack endpoint.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const endpoint1 = openstack.identity.getEndpoint({
 *     serviceName: "demo",
 * });
 * ```
 */
export function getEndpoint(args?: GetEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:identity/getEndpoint:getEndpoint", {
        "endpointRegion": args.endpointRegion,
        "interface": args.interface,
        "name": args.name,
        "region": args.region,
        "serviceId": args.serviceId,
        "serviceName": args.serviceName,
        "serviceType": args.serviceType,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpoint.
 */
export interface GetEndpointArgs {
    /**
     * The region the endpoint is assigned to. The
     * `region` and `endpointRegion` can be different.
     */
    endpointRegion?: string;
    /**
     * The endpoint interface. Valid values are `public`,
     * `internal`, and `admin`. Default value is `public`
     */
    interface?: string;
    /**
     * The name of the endpoint.
     */
    name?: string;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The service id this endpoint belongs to.
     */
    serviceId?: string;
    /**
     * The service name of the endpoint.
     */
    serviceName?: string;
    /**
     * The service type of the endpoint.
     */
    serviceType?: string;
}

/**
 * A collection of values returned by getEndpoint.
 */
export interface GetEndpointResult {
    /**
     * See Argument Reference above.
     */
    readonly endpointRegion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly interface?: string;
    /**
     * See Argument Reference above.
     */
    readonly name?: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceId?: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName?: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceType?: string;
    /**
     * The endpoint URL.
     */
    readonly url: string;
}
/**
 * Use this data source to get the ID of an OpenStack endpoint.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const endpoint1 = openstack.identity.getEndpoint({
 *     serviceName: "demo",
 * });
 * ```
 */
export function getEndpointOutput(args?: GetEndpointOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEndpointResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:identity/getEndpoint:getEndpoint", {
        "endpointRegion": args.endpointRegion,
        "interface": args.interface,
        "name": args.name,
        "region": args.region,
        "serviceId": args.serviceId,
        "serviceName": args.serviceName,
        "serviceType": args.serviceType,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpoint.
 */
export interface GetEndpointOutputArgs {
    /**
     * The region the endpoint is assigned to. The
     * `region` and `endpointRegion` can be different.
     */
    endpointRegion?: pulumi.Input<string>;
    /**
     * The endpoint interface. Valid values are `public`,
     * `internal`, and `admin`. Default value is `public`
     */
    interface?: pulumi.Input<string>;
    /**
     * The name of the endpoint.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The service id this endpoint belongs to.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The service name of the endpoint.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The service type of the endpoint.
     */
    serviceType?: pulumi.Input<string>;
}
