// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 Image resource within OpenStack Glance.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const rancheros = new openstack.images.Image("rancheros", {
 *     containerFormat: "bare",
 *     diskFormat: "qcow2",
 *     imageSourceUrl: "https://releases.rancher.com/os/latest/rancheros-openstack.img",
 *     properties: {
 *         key: "value",
 *     },
 * });
 * ```
 * 
 * ## Notes
 * 
 * ### Properties
 * 
 * This resource supports the ability to add properties to a resource during
 * creation as well as add, update, and delete properties during an update of this
 * resource.
 * 
 * Newer versions of OpenStack are adding some read-only properties to each image.
 * These properties start with the prefix `os_`. If these properties are detected,
 * this resource will automatically reconcile these with the user-provided
 * properties.
 * 
 * In addition, the `direct_url` property is also automatically reconciled if the
 * Image Service set it.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/images_image_v2.html.markdown.
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageState, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:images/image:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * The checksum of the data associated with the image.
     */
    public /*out*/ readonly checksum!: pulumi.Output<string>;
    /**
     * The container format. Must be one of
     * "ami", "ari", "aki", "bare", "ovf".
     */
    public readonly containerFormat!: pulumi.Output<string>;
    /**
     * The date the image was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The disk format. Must be one of
     * "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
     */
    public readonly diskFormat!: pulumi.Output<string>;
    /**
     * the trailing path after the glance
     * endpoint that represent the location of the image
     * or the path to retrieve it.
     */
    public /*out*/ readonly file!: pulumi.Output<string>;
    public readonly imageCachePath!: pulumi.Output<string | undefined>;
    /**
     * This is the url of the raw image that will
     * be downloaded in the `image_cache_path` before being uploaded to Glance.
     * Glance is able to download image from internet but the `gophercloud` library
     * does not yet provide a way to do so.
     * Conflicts with `local_file_path`.
     */
    public readonly imageSourceUrl!: pulumi.Output<string | undefined>;
    /**
     * This is the filepath of the raw image file
     * that will be uploaded to Glance. Conflicts with `image_source_url`.
     */
    public readonly localFilePath!: pulumi.Output<string | undefined>;
    /**
     * The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
     */
    public /*out*/ readonly metadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * Amount of disk space (in GB) required to boot image.
     * Defaults to 0.
     */
    public readonly minDiskGb!: pulumi.Output<number | undefined>;
    /**
     * Amount of ram (in MB) required to boot image.
     * Defauts to 0.
     */
    public readonly minRamMb!: pulumi.Output<number | undefined>;
    /**
     * The name of the image.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the openstack user who owns the image.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * A map of key/value pairs to set freeform
     * information about an image. See the "Notes" section for further
     * information about properties.
     */
    public readonly properties!: pulumi.Output<{[key: string]: any}>;
    /**
     * If true, image will not be deletable.
     * Defaults to false.
     */
    public readonly protected!: pulumi.Output<boolean | undefined>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new Image.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The path to the JSON-schema that represent
     * the image or image
     */
    public /*out*/ readonly schema!: pulumi.Output<string>;
    /**
     * The size in bytes of the data associated with the image.
     */
    public /*out*/ readonly sizeBytes!: pulumi.Output<number>;
    /**
     * The status of the image. It can be "queued", "active"
     * or "saving".
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The tags of the image. It must be a list of strings.
     * At this time, it is not possible to delete all tags of an image.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * (**Deprecated** - use `updated_at` instead)
     */
    public /*out*/ readonly updateAt!: pulumi.Output<string>;
    /**
     * The date the image was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * If false, the checksum will not be verified
     * once the image is finished uploading. Defaults to true.
     */
    public readonly verifyChecksum!: pulumi.Output<boolean | undefined>;
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". The ability to set the
     * visibility depends upon the configuration of the OpenStack cloud.
     */
    public readonly visibility!: pulumi.Output<string | undefined>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageArgs | ImageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ImageState | undefined;
            inputs["checksum"] = state ? state.checksum : undefined;
            inputs["containerFormat"] = state ? state.containerFormat : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["diskFormat"] = state ? state.diskFormat : undefined;
            inputs["file"] = state ? state.file : undefined;
            inputs["imageCachePath"] = state ? state.imageCachePath : undefined;
            inputs["imageSourceUrl"] = state ? state.imageSourceUrl : undefined;
            inputs["localFilePath"] = state ? state.localFilePath : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["minDiskGb"] = state ? state.minDiskGb : undefined;
            inputs["minRamMb"] = state ? state.minRamMb : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["owner"] = state ? state.owner : undefined;
            inputs["properties"] = state ? state.properties : undefined;
            inputs["protected"] = state ? state.protected : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["schema"] = state ? state.schema : undefined;
            inputs["sizeBytes"] = state ? state.sizeBytes : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["updateAt"] = state ? state.updateAt : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
            inputs["verifyChecksum"] = state ? state.verifyChecksum : undefined;
            inputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ImageArgs | undefined;
            if (!args || args.containerFormat === undefined) {
                throw new Error("Missing required property 'containerFormat'");
            }
            if (!args || args.diskFormat === undefined) {
                throw new Error("Missing required property 'diskFormat'");
            }
            inputs["containerFormat"] = args ? args.containerFormat : undefined;
            inputs["diskFormat"] = args ? args.diskFormat : undefined;
            inputs["imageCachePath"] = args ? args.imageCachePath : undefined;
            inputs["imageSourceUrl"] = args ? args.imageSourceUrl : undefined;
            inputs["localFilePath"] = args ? args.localFilePath : undefined;
            inputs["minDiskGb"] = args ? args.minDiskGb : undefined;
            inputs["minRamMb"] = args ? args.minRamMb : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["protected"] = args ? args.protected : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["verifyChecksum"] = args ? args.verifyChecksum : undefined;
            inputs["visibility"] = args ? args.visibility : undefined;
            inputs["checksum"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["file"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["owner"] = undefined /*out*/;
            inputs["schema"] = undefined /*out*/;
            inputs["sizeBytes"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["updateAt"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        super(Image.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Image resources.
 */
export interface ImageState {
    /**
     * The checksum of the data associated with the image.
     */
    readonly checksum?: pulumi.Input<string>;
    /**
     * The container format. Must be one of
     * "ami", "ari", "aki", "bare", "ovf".
     */
    readonly containerFormat?: pulumi.Input<string>;
    /**
     * The date the image was created.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * The disk format. Must be one of
     * "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
     */
    readonly diskFormat?: pulumi.Input<string>;
    /**
     * the trailing path after the glance
     * endpoint that represent the location of the image
     * or the path to retrieve it.
     */
    readonly file?: pulumi.Input<string>;
    readonly imageCachePath?: pulumi.Input<string>;
    /**
     * This is the url of the raw image that will
     * be downloaded in the `image_cache_path` before being uploaded to Glance.
     * Glance is able to download image from internet but the `gophercloud` library
     * does not yet provide a way to do so.
     * Conflicts with `local_file_path`.
     */
    readonly imageSourceUrl?: pulumi.Input<string>;
    /**
     * This is the filepath of the raw image file
     * that will be uploaded to Glance. Conflicts with `image_source_url`.
     */
    readonly localFilePath?: pulumi.Input<string>;
    /**
     * The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Amount of disk space (in GB) required to boot image.
     * Defaults to 0.
     */
    readonly minDiskGb?: pulumi.Input<number>;
    /**
     * Amount of ram (in MB) required to boot image.
     * Defauts to 0.
     */
    readonly minRamMb?: pulumi.Input<number>;
    /**
     * The name of the image.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The id of the openstack user who owns the image.
     */
    readonly owner?: pulumi.Input<string>;
    /**
     * A map of key/value pairs to set freeform
     * information about an image. See the "Notes" section for further
     * information about properties.
     */
    readonly properties?: pulumi.Input<{[key: string]: any}>;
    /**
     * If true, image will not be deletable.
     * Defaults to false.
     */
    readonly protected?: pulumi.Input<boolean>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new Image.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The path to the JSON-schema that represent
     * the image or image
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * The size in bytes of the data associated with the image.
     */
    readonly sizeBytes?: pulumi.Input<number>;
    /**
     * The status of the image. It can be "queued", "active"
     * or "saving".
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The tags of the image. It must be a list of strings.
     * At this time, it is not possible to delete all tags of an image.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (**Deprecated** - use `updated_at` instead)
     */
    readonly updateAt?: pulumi.Input<string>;
    /**
     * The date the image was last updated.
     */
    readonly updatedAt?: pulumi.Input<string>;
    /**
     * If false, the checksum will not be verified
     * once the image is finished uploading. Defaults to true.
     */
    readonly verifyChecksum?: pulumi.Input<boolean>;
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". The ability to set the
     * visibility depends upon the configuration of the OpenStack cloud.
     */
    readonly visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * The container format. Must be one of
     * "ami", "ari", "aki", "bare", "ovf".
     */
    readonly containerFormat: pulumi.Input<string>;
    /**
     * The disk format. Must be one of
     * "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
     */
    readonly diskFormat: pulumi.Input<string>;
    readonly imageCachePath?: pulumi.Input<string>;
    /**
     * This is the url of the raw image that will
     * be downloaded in the `image_cache_path` before being uploaded to Glance.
     * Glance is able to download image from internet but the `gophercloud` library
     * does not yet provide a way to do so.
     * Conflicts with `local_file_path`.
     */
    readonly imageSourceUrl?: pulumi.Input<string>;
    /**
     * This is the filepath of the raw image file
     * that will be uploaded to Glance. Conflicts with `image_source_url`.
     */
    readonly localFilePath?: pulumi.Input<string>;
    /**
     * Amount of disk space (in GB) required to boot image.
     * Defaults to 0.
     */
    readonly minDiskGb?: pulumi.Input<number>;
    /**
     * Amount of ram (in MB) required to boot image.
     * Defauts to 0.
     */
    readonly minRamMb?: pulumi.Input<number>;
    /**
     * The name of the image.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of key/value pairs to set freeform
     * information about an image. See the "Notes" section for further
     * information about properties.
     */
    readonly properties?: pulumi.Input<{[key: string]: any}>;
    /**
     * If true, image will not be deletable.
     * Defaults to false.
     */
    readonly protected?: pulumi.Input<boolean>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new Image.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The tags of the image. It must be a list of strings.
     * At this time, it is not possible to delete all tags of an image.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If false, the checksum will not be verified
     * once the image is finished uploading. Defaults to true.
     */
    readonly verifyChecksum?: pulumi.Input<boolean>;
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". The ability to set the
     * visibility depends upon the configuration of the OpenStack cloud.
     */
    readonly visibility?: pulumi.Input<string>;
}
