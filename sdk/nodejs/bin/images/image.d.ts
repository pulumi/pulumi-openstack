import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 Image resource within OpenStack Glance.
 */
export declare class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageState): Image;
    /**
     * The checksum of the data associated with the image.
     */
    readonly checksum: pulumi.Output<string>;
    /**
     * The container format. Must be one of
     * "ami", "ari", "aki", "bare", "ovf".
     */
    readonly containerFormat: pulumi.Output<string>;
    /**
     * The date the image was created.
     */
    readonly createdAt: pulumi.Output<string>;
    /**
     * The disk format. Must be one of
     * "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
     */
    readonly diskFormat: pulumi.Output<string>;
    /**
     * the trailing path after the glance
     * endpoint that represent the location of the image
     * or the path to retrieve it.
     */
    readonly file: pulumi.Output<string>;
    /**
     * This is the directory where the images will
     * be downloaded. Images will be stored with a filename corresponding to
     * the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"
     */
    readonly imageCachePath: pulumi.Output<string | undefined>;
    /**
     * This is the url of the raw image that will
     * be downloaded in the `image_cache_path` before being uploaded to Glance.
     * Glance is able to download image from internet but the `gophercloud` library
     * does not yet provide a way to do so.
     * Conflicts with `local_file_path`.
     */
    readonly imageSourceUrl: pulumi.Output<string | undefined>;
    /**
     * This is the filepath of the raw image file
     * that will be uploaded to Glance. Conflicts with `image_source_url`.
     */
    readonly localFilePath: pulumi.Output<string | undefined>;
    /**
     * The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
     */
    readonly metadata: pulumi.Output<{
        [key: string]: any;
    }>;
    /**
     * Amount of disk space (in GB) required to boot image.
     * Defaults to 0.
     */
    readonly minDiskGb: pulumi.Output<number | undefined>;
    /**
     * Amount of ram (in MB) required to boot image.
     * Defauts to 0.
     */
    readonly minRamMb: pulumi.Output<number | undefined>;
    /**
     * The name of the image.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The id of the openstack user who owns the image.
     */
    readonly owner: pulumi.Output<string>;
    /**
     * A map of key/value pairs to set freeform
     * information about an image.
     */
    readonly properties: pulumi.Output<{
        [key: string]: any;
    }>;
    /**
     * If true, image will not be deletable.
     * Defaults to false.
     */
    readonly protected: pulumi.Output<boolean | undefined>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new Image.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The path to the JSON-schema that represent
     * the image or image
     */
    readonly schema: pulumi.Output<string>;
    /**
     * The size in bytes of the data associated with the image.
     */
    readonly sizeBytes: pulumi.Output<number>;
    /**
     * The status of the image. It can be "queued", "active"
     * or "saving".
     */
    readonly status: pulumi.Output<string>;
    /**
     * The tags of the image. It must be a list of strings.
     * At this time, it is not possible to delete all tags of an image.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * The date the image was last updated.
     */
    readonly updateAt: pulumi.Output<string>;
    /**
     * If false, the checksum will not be verified
     * once the image is finished uploading. Defaults to true.
     */
    readonly verifyChecksum: pulumi.Output<boolean | undefined>;
    /**
     * The visibility of the image. Must be one of
     * "public", "private", "community", or "shared". The ability to set the
     * visibility depends upon the configuration of the OpenStack cloud.
     */
    readonly visibility: pulumi.Output<string | undefined>;
    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.ResourceOptions);
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
    /**
     * This is the directory where the images will
     * be downloaded. Images will be stored with a filename corresponding to
     * the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"
     */
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
     * and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.
     */
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
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
     * information about an image.
     */
    readonly properties?: pulumi.Input<{
        [key: string]: any;
    }>;
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
     * The date the image was last updated.
     */
    readonly updateAt?: pulumi.Input<string>;
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
    /**
     * This is the directory where the images will
     * be downloaded. Images will be stored with a filename corresponding to
     * the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"
     */
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
     * information about an image.
     */
    readonly properties?: pulumi.Input<{
        [key: string]: any;
    }>;
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
