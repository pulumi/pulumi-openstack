// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of Openstack Image IDs matching the
 * specified criteria.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const images = pulumi.output(openstack.images.getImageIds({
 *     nameRegex: "^Ubuntu 16\\.04.*-amd64",
 *     properties: {
 *         key: "value",
 *     },
 *     sort: "updated_at",
 * }));
 * ```
 */
export function getImageIds(args?: GetImageIdsArgs, opts?: pulumi.InvokeOptions): Promise<GetImageIdsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:images/getImageIds:getImageIds", {
        "memberStatus": args.memberStatus,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "owner": args.owner,
        "properties": args.properties,
        "region": args.region,
        "sizeMax": args.sizeMax,
        "sizeMin": args.sizeMin,
        "sort": args.sort,
        "sortDirection": args.sortDirection,
        "sortKey": args.sortKey,
        "tag": args.tag,
        "visibility": args.visibility,
    }, opts);
}

/**
 * A collection of arguments for invoking getImageIds.
 */
export interface GetImageIdsArgs {
    /**
     * The status of the image. Must be one of
     * "accepted", "pending", "rejected", or "all".
     */
    memberStatus?: string;
    /**
     * The name of the image. Cannot be used simultaneously
     * with `nameRegex`.
     */
    name?: string;
    /**
     * The regular expressian of the name of the image.
     * Cannot be used simultaneously with `name`. Unlike filtering by `name` the
     * `nameRegex` filtering does by client on the result of OpenStack search
     * query.
     */
    nameRegex?: string;
    /**
     * The owner (UUID) of the image.
     */
    owner?: string;
    /**
     * a map of key/value pairs to match an image with.
     * All specified properties must be matched. Unlike other options filtering
     * by `properties` does by client on the result of OpenStack search query.
     */
    properties?: {[key: string]: any};
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used.
     */
    region?: string;
    /**
     * The maximum size (in bytes) of the image to return.
     */
    sizeMax?: number;
    /**
     * The minimum size (in bytes) of the image to return.
     */
    sizeMin?: number;
    /**
     * Sorts the response by one or more attribute and sort
     * direction combinations. You can also set multiple sort keys and directions.
     * Default direction is `desc`. Use the comma (,) character to separate
     * multiple values. For example expression `sort = "name:asc,status"`
     * sorts ascending by name and descending by status. `sort` cannot be used
     * simultaneously with `sortKey`. If both are present in a configuration
     * then only `sort` will be used.
     */
    sort?: string;
    /**
     * Order the results in either `asc` or `desc`.
     * Can be applied only with `sortKey`. Defaults to `asc`
     *
     * @deprecated Use option 'sort' instead.
     */
    sortDirection?: string;
    /**
     * Sort images based on a certain key. Defaults to
     * `name`. `sortKey` cannot be used simultaneously with `sort`. If both
     * are present in a configuration then only `sort` will be used.
     *
     * @deprecated Use option 'sort' instead.
     */
    sortKey?: string;
    /**
     * Search for images with a specific tag.
     */
    tag?: string;
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". Defaults to "private".
     */
    visibility?: string;
}

/**
 * A collection of values returned by getImageIds.
 */
export interface GetImageIdsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly memberStatus?: string;
    readonly name?: string;
    readonly nameRegex?: string;
    readonly owner?: string;
    readonly properties?: {[key: string]: any};
    readonly region: string;
    readonly sizeMax?: number;
    readonly sizeMin?: number;
    readonly sort?: string;
    /**
     * @deprecated Use option 'sort' instead.
     */
    readonly sortDirection?: string;
    /**
     * @deprecated Use option 'sort' instead.
     */
    readonly sortKey?: string;
    readonly tag?: string;
    readonly visibility?: string;
}

export function getImageIdsOutput(args?: GetImageIdsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageIdsResult> {
    return pulumi.output(args).apply(a => getImageIds(a, opts))
}

/**
 * A collection of arguments for invoking getImageIds.
 */
export interface GetImageIdsOutputArgs {
    /**
     * The status of the image. Must be one of
     * "accepted", "pending", "rejected", or "all".
     */
    memberStatus?: pulumi.Input<string>;
    /**
     * The name of the image. Cannot be used simultaneously
     * with `nameRegex`.
     */
    name?: pulumi.Input<string>;
    /**
     * The regular expressian of the name of the image.
     * Cannot be used simultaneously with `name`. Unlike filtering by `name` the
     * `nameRegex` filtering does by client on the result of OpenStack search
     * query.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The owner (UUID) of the image.
     */
    owner?: pulumi.Input<string>;
    /**
     * a map of key/value pairs to match an image with.
     * All specified properties must be matched. Unlike other options filtering
     * by `properties` does by client on the result of OpenStack search query.
     */
    properties?: pulumi.Input<{[key: string]: any}>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The maximum size (in bytes) of the image to return.
     */
    sizeMax?: pulumi.Input<number>;
    /**
     * The minimum size (in bytes) of the image to return.
     */
    sizeMin?: pulumi.Input<number>;
    /**
     * Sorts the response by one or more attribute and sort
     * direction combinations. You can also set multiple sort keys and directions.
     * Default direction is `desc`. Use the comma (,) character to separate
     * multiple values. For example expression `sort = "name:asc,status"`
     * sorts ascending by name and descending by status. `sort` cannot be used
     * simultaneously with `sortKey`. If both are present in a configuration
     * then only `sort` will be used.
     */
    sort?: pulumi.Input<string>;
    /**
     * Order the results in either `asc` or `desc`.
     * Can be applied only with `sortKey`. Defaults to `asc`
     *
     * @deprecated Use option 'sort' instead.
     */
    sortDirection?: pulumi.Input<string>;
    /**
     * Sort images based on a certain key. Defaults to
     * `name`. `sortKey` cannot be used simultaneously with `sort`. If both
     * are present in a configuration then only `sort` will be used.
     *
     * @deprecated Use option 'sort' instead.
     */
    sortKey?: pulumi.Input<string>;
    /**
     * Search for images with a specific tag.
     */
    tag?: pulumi.Input<string>;
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". Defaults to "private".
     */
    visibility?: pulumi.Input<string>;
}
