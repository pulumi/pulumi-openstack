// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.images.ImageArgs;
import com.pulumi.openstack.images.inputs.ImageState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Image resource within OpenStack Glance.
 * 
 * &gt; **Note:** All arguments including the source image URL password will be
 * stored in the raw state as plain-text. Read more about sensitive data in
 * state.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.images.Image;
 * import com.pulumi.openstack.images.ImageArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var rancheros = new Image(&#34;rancheros&#34;, ImageArgs.builder()        
 *             .containerFormat(&#34;bare&#34;)
 *             .diskFormat(&#34;qcow2&#34;)
 *             .imageSourceUrl(&#34;https://releases.rancher.com/os/latest/rancheros-openstack.img&#34;)
 *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
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
 * In addition, the `direct_url` and `stores` properties are also automatically reconciled if the
 * Image Service set it.
 * 
 * ## Import
 * 
 * Images can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:images/image:Image rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2
 * ```
 * 
 */
@ResourceType(type="openstack:images/image:Image")
public class Image extends com.pulumi.resources.CustomResource {
    /**
     * The checksum of the data associated with the image.
     * 
     */
    @Export(name="checksum", type=String.class, parameters={})
    private Output<String> checksum;

    /**
     * @return The checksum of the data associated with the image.
     * 
     */
    public Output<String> checksum() {
        return this.checksum;
    }
    /**
     * The container format. Must be one of
     * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;bare&#34;, &#34;ovf&#34;.
     * 
     */
    @Export(name="containerFormat", type=String.class, parameters={})
    private Output<String> containerFormat;

    /**
     * @return The container format. Must be one of
     * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;bare&#34;, &#34;ovf&#34;.
     * 
     */
    public Output<String> containerFormat() {
        return this.containerFormat;
    }
    /**
     * The date the image was created.
     * 
     */
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    /**
     * @return The date the image was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * If true, this provider will decompress downloaded
     * image before uploading it to OpenStack. Decompression algorithm is chosen by
     * checking &#34;Content-Type&#34; header, supported algorithm are: gzip, bzip2 and xz.
     * Defaults to false. Changing this creates a new Image.
     * 
     */
    @Export(name="decompress", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> decompress;

    /**
     * @return If true, this provider will decompress downloaded
     * image before uploading it to OpenStack. Decompression algorithm is chosen by
     * checking &#34;Content-Type&#34; header, supported algorithm are: gzip, bzip2 and xz.
     * Defaults to false. Changing this creates a new Image.
     * 
     */
    public Output<Optional<Boolean>> decompress() {
        return Codegen.optional(this.decompress);
    }
    /**
     * The disk format. Must be one of
     * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;vhd&#34;, &#34;vmdk&#34;, &#34;raw&#34;, &#34;qcow2&#34;, &#34;vdi&#34;, &#34;iso&#34;.
     * 
     */
    @Export(name="diskFormat", type=String.class, parameters={})
    private Output<String> diskFormat;

    /**
     * @return The disk format. Must be one of
     * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;vhd&#34;, &#34;vmdk&#34;, &#34;raw&#34;, &#34;qcow2&#34;, &#34;vdi&#34;, &#34;iso&#34;.
     * 
     */
    public Output<String> diskFormat() {
        return this.diskFormat;
    }
    /**
     * the trailing path after the glance
     * endpoint that represent the location of the image
     * or the path to retrieve it.
     * 
     */
    @Export(name="file", type=String.class, parameters={})
    private Output<String> file;

    /**
     * @return the trailing path after the glance
     * endpoint that represent the location of the image
     * or the path to retrieve it.
     * 
     */
    public Output<String> file() {
        return this.file;
    }
    /**
     * If true, image will be hidden from public list.
     * Defaults to false.
     * 
     */
    @Export(name="hidden", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> hidden;

    /**
     * @return If true, image will be hidden from public list.
     * Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> hidden() {
        return Codegen.optional(this.hidden);
    }
    @Export(name="imageCachePath", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageCachePath;

    public Output<Optional<String>> imageCachePath() {
        return Codegen.optional(this.imageCachePath);
    }
    /**
     * Unique ID (valid UUID) of image to create. Changing
     * this creates a new image.
     * 
     */
    @Export(name="imageId", type=String.class, parameters={})
    private Output<String> imageId;

    /**
     * @return Unique ID (valid UUID) of image to create. Changing
     * this creates a new image.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * The password of basic auth to download `image_source_url`.
     * 
     */
    @Export(name="imageSourcePassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageSourcePassword;

    /**
     * @return The password of basic auth to download `image_source_url`.
     * 
     */
    public Output<Optional<String>> imageSourcePassword() {
        return Codegen.optional(this.imageSourcePassword);
    }
    /**
     * This is the url of the raw image. If `web_download`
     * is not used, then the image will be downloaded in the `image_cache_path` before
     * being uploaded to Glance.
     * Conflicts with `local_file_path`.
     * 
     */
    @Export(name="imageSourceUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageSourceUrl;

    /**
     * @return This is the url of the raw image. If `web_download`
     * is not used, then the image will be downloaded in the `image_cache_path` before
     * being uploaded to Glance.
     * Conflicts with `local_file_path`.
     * 
     */
    public Output<Optional<String>> imageSourceUrl() {
        return Codegen.optional(this.imageSourceUrl);
    }
    /**
     * The username of basic auth to download `image_source_url`.
     * 
     */
    @Export(name="imageSourceUsername", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageSourceUsername;

    /**
     * @return The username of basic auth to download `image_source_url`.
     * 
     */
    public Output<Optional<String>> imageSourceUsername() {
        return Codegen.optional(this.imageSourceUsername);
    }
    /**
     * This is the filepath of the raw image file
     * that will be uploaded to Glance. Conflicts with `image_source_url` and
     * `web_download`.
     * 
     */
    @Export(name="localFilePath", type=String.class, parameters={})
    private Output</* @Nullable */ String> localFilePath;

    /**
     * @return This is the filepath of the raw image file
     * that will be uploaded to Glance. Conflicts with `image_source_url` and
     * `web_download`.
     * 
     */
    public Output<Optional<String>> localFilePath() {
        return Codegen.optional(this.localFilePath);
    }
    /**
     * The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
     * 
     */
    @Export(name="metadata", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> metadata;

    /**
     * @return The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
     * 
     */
    public Output<Map<String,Object>> metadata() {
        return this.metadata;
    }
    /**
     * Amount of disk space (in GB) required to boot image.
     * Defaults to 0.
     * 
     */
    @Export(name="minDiskGb", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> minDiskGb;

    /**
     * @return Amount of disk space (in GB) required to boot image.
     * Defaults to 0.
     * 
     */
    public Output<Optional<Integer>> minDiskGb() {
        return Codegen.optional(this.minDiskGb);
    }
    /**
     * Amount of ram (in MB) required to boot image.
     * Defauts to 0.
     * 
     */
    @Export(name="minRamMb", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> minRamMb;

    /**
     * @return Amount of ram (in MB) required to boot image.
     * Defauts to 0.
     * 
     */
    public Output<Optional<Integer>> minRamMb() {
        return Codegen.optional(this.minRamMb);
    }
    /**
     * The name of the image.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the image.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The id of the openstack user who owns the image.
     * 
     */
    @Export(name="owner", type=String.class, parameters={})
    private Output<String> owner;

    /**
     * @return The id of the openstack user who owns the image.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * A map of key/value pairs to set freeform
     * information about an image. See the &#34;Notes&#34; section for further
     * information about properties.
     * 
     */
    @Export(name="properties", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> properties;

    /**
     * @return A map of key/value pairs to set freeform
     * information about an image. See the &#34;Notes&#34; section for further
     * information about properties.
     * 
     */
    public Output<Map<String,Object>> properties() {
        return this.properties;
    }
    /**
     * If true, image will not be deletable.
     * Defaults to false.
     * 
     */
    @Export(name="protected", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> protected_;

    /**
     * @return If true, image will not be deletable.
     * Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> protected_() {
        return Codegen.optional(this.protected_);
    }
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new Image.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new Image.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The path to the JSON-schema that represent
     * the image or image
     * 
     */
    @Export(name="schema", type=String.class, parameters={})
    private Output<String> schema;

    /**
     * @return The path to the JSON-schema that represent
     * the image or image
     * 
     */
    public Output<String> schema() {
        return this.schema;
    }
    /**
     * The size in bytes of the data associated with the image.
     * 
     */
    @Export(name="sizeBytes", type=Integer.class, parameters={})
    private Output<Integer> sizeBytes;

    /**
     * @return The size in bytes of the data associated with the image.
     * 
     */
    public Output<Integer> sizeBytes() {
        return this.sizeBytes;
    }
    /**
     * The status of the image. It can be &#34;queued&#34;, &#34;active&#34;
     * or &#34;saving&#34;.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the image. It can be &#34;queued&#34;, &#34;active&#34;
     * or &#34;saving&#34;.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The tags of the image. It must be a list of strings.
     * At this time, it is not possible to delete all tags of an image.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The tags of the image. It must be a list of strings.
     * At this time, it is not possible to delete all tags of an image.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * (**Deprecated** - use `updated_at` instead)
     * 
     * @deprecated
     * Use updated_at instead
     * 
     */
    @Deprecated /* Use updated_at instead */
    @Export(name="updateAt", type=String.class, parameters={})
    private Output<String> updateAt;

    /**
     * @return (**Deprecated** - use `updated_at` instead)
     * 
     */
    public Output<String> updateAt() {
        return this.updateAt;
    }
    /**
     * The date the image was last updated.
     * 
     */
    @Export(name="updatedAt", type=String.class, parameters={})
    private Output<String> updatedAt;

    /**
     * @return The date the image was last updated.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * If false, the checksum will not be verified
     * once the image is finished uploading. Conflicts with `web_download`.
     * Defaults to true when not using `web_download`.
     * 
     */
    @Export(name="verifyChecksum", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> verifyChecksum;

    /**
     * @return If false, the checksum will not be verified
     * once the image is finished uploading. Conflicts with `web_download`.
     * Defaults to true when not using `web_download`.
     * 
     */
    public Output<Optional<Boolean>> verifyChecksum() {
        return Codegen.optional(this.verifyChecksum);
    }
    /**
     * The visibility of the image. Must be one of
     * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. The ability to set the
     * visibility depends upon the configuration of the OpenStack cloud.
     * 
     */
    @Export(name="visibility", type=String.class, parameters={})
    private Output</* @Nullable */ String> visibility;

    /**
     * @return The visibility of the image. Must be one of
     * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. The ability to set the
     * visibility depends upon the configuration of the OpenStack cloud.
     * 
     */
    public Output<Optional<String>> visibility() {
        return Codegen.optional(this.visibility);
    }
    /**
     * If true, the &#34;web-download&#34; import method will
     * be used to let Openstack download the image directly from the remote source.
     * Conflicts with `local_file_path`. Defaults to false.
     * 
     */
    @Export(name="webDownload", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> webDownload;

    /**
     * @return If true, the &#34;web-download&#34; import method will
     * be used to let Openstack download the image directly from the remote source.
     * Conflicts with `local_file_path`. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> webDownload() {
        return Codegen.optional(this.webDownload);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Image(String name) {
        this(name, ImageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Image(String name, ImageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Image(String name, ImageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:images/image:Image", name, args == null ? ImageArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Image(String name, Output<String> id, @Nullable ImageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:images/image:Image", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "imageSourcePassword"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Image get(String name, Output<String> id, @Nullable ImageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Image(name, id, state, options);
    }
}
