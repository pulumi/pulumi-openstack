// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack image.
 */
export function getImage(args?: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:images/getImage:getImage", {
        "hidden": args.hidden,
        "memberStatus": args.memberStatus,
        "mostRecent": args.mostRecent,
        "name": args.name,
        "nameRegex": args.nameRegex,
        "owner": args.owner,
        "properties": args.properties,
        "region": args.region,
        "sizeMax": args.sizeMax,
        "sizeMin": args.sizeMin,
        "sortDirection": args.sortDirection,
        "sortKey": args.sortKey,
        "tag": args.tag,
        "tags": args.tags,
        "visibility": args.visibility,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * Whether or not the image is hidden from public list.
     */
    hidden?: boolean;
    /**
     * The status of the image. Must be one of
     * "accepted", "pending", "rejected", or "all".
     */
    memberStatus?: string;
    /**
     * If more than one result is returned, use the most
     * recent image.
     */
    mostRecent?: boolean;
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
     * Filtering is applied if server responce contains at least 2 images. In
     * case there is only one image the `properties` ignores.
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
     * Order the results in either `asc` or `desc`.
     */
    sortDirection?: string;
    /**
     * Sort images based on a certain key. Defaults to `name`.
     */
    sortKey?: string;
    /**
     * Search for images with a specific tag.
     */
    tag?: string;
    /**
     * A list of tags required to be set on the image 
     * (all specified tags must be in the images tag list for it to be matched).
     */
    tags?: string[];
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". Defaults to "private".
     */
    visibility?: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    /**
     * The checksum of the data associated with the image.
     */
    readonly checksum: string;
    /**
     * The format of the image's container.
     */
    readonly containerFormat: string;
    /**
     * The date the image was created.
     */
    readonly createdAt: string;
    /**
     * The format of the image's disk.
     */
    readonly diskFormat: string;
    /**
     * the trailing path after the glance endpoint that represent the
     * location of the image or the path to retrieve it.
     */
    readonly file: string;
    readonly hidden?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly memberStatus?: string;
    /**
     * The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
     */
    readonly metadata: {[key: string]: any};
    /**
     * The minimum amount of disk space required to use the image.
     */
    readonly minDiskGb: number;
    /**
     * The minimum amount of ram required to use the image.
     */
    readonly minRamMb: number;
    readonly mostRecent?: boolean;
    readonly name?: string;
    readonly nameRegex?: string;
    readonly owner?: string;
    /**
     * Freeform information about the image.
     */
    readonly properties?: {[key: string]: any};
    /**
     * Whether or not the image is protected.
     */
    readonly protected: boolean;
    readonly region: string;
    /**
     * The path to the JSON-schema that represent
     * the image or image
     */
    readonly schema: string;
    /**
     * The size of the image (in bytes).
     */
    readonly sizeBytes: number;
    readonly sizeMax?: number;
    readonly sizeMin?: number;
    readonly sortDirection?: string;
    readonly sortKey?: string;
    readonly tag?: string;
    /**
     * The tags list of the image.
     */
    readonly tags: string[];
    /**
     * The date the image was last updated.
     */
    readonly updatedAt: string;
    readonly visibility?: string;
}
/**
 * Use this data source to get the ID of an available OpenStack image.
 */
export function getImageOutput(args?: GetImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageResult> {
    return pulumi.output(args).apply((a: any) => getImage(a, opts))
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageOutputArgs {
    /**
     * Whether or not the image is hidden from public list.
     */
    hidden?: pulumi.Input<boolean>;
    /**
     * The status of the image. Must be one of
     * "accepted", "pending", "rejected", or "all".
     */
    memberStatus?: pulumi.Input<string>;
    /**
     * If more than one result is returned, use the most
     * recent image.
     */
    mostRecent?: pulumi.Input<boolean>;
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
     * Filtering is applied if server responce contains at least 2 images. In
     * case there is only one image the `properties` ignores.
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
     * Order the results in either `asc` or `desc`.
     */
    sortDirection?: pulumi.Input<string>;
    /**
     * Sort images based on a certain key. Defaults to `name`.
     */
    sortKey?: pulumi.Input<string>;
    /**
     * Search for images with a specific tag.
     */
    tag?: pulumi.Input<string>;
    /**
     * A list of tags required to be set on the image 
     * (all specified tags must be in the images tag list for it to be matched).
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". Defaults to "private".
     */
    visibility?: pulumi.Input<string>;
}
