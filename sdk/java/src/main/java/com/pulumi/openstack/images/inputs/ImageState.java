// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ImageState extends com.pulumi.resources.ResourceArgs {

    public static final ImageState Empty = new ImageState();

    /**
     * The checksum of the data associated with the image.
     * 
     */
    @Import(name="checksum")
    private @Nullable Output<String> checksum;

    /**
     * @return The checksum of the data associated with the image.
     * 
     */
    public Optional<Output<String>> checksum() {
        return Optional.ofNullable(this.checksum);
    }

    /**
     * The container format. Must be one of
     * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;bare&#34;, &#34;ovf&#34;.
     * 
     */
    @Import(name="containerFormat")
    private @Nullable Output<String> containerFormat;

    /**
     * @return The container format. Must be one of
     * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;bare&#34;, &#34;ovf&#34;.
     * 
     */
    public Optional<Output<String>> containerFormat() {
        return Optional.ofNullable(this.containerFormat);
    }

    /**
     * The date the image was created.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The date the image was created.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * The disk format. Must be one of
     * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;vhd&#34;, &#34;vmdk&#34;, &#34;raw&#34;, &#34;qcow2&#34;, &#34;vdi&#34;, &#34;iso&#34;.
     * 
     */
    @Import(name="diskFormat")
    private @Nullable Output<String> diskFormat;

    /**
     * @return The disk format. Must be one of
     * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;vhd&#34;, &#34;vmdk&#34;, &#34;raw&#34;, &#34;qcow2&#34;, &#34;vdi&#34;, &#34;iso&#34;.
     * 
     */
    public Optional<Output<String>> diskFormat() {
        return Optional.ofNullable(this.diskFormat);
    }

    /**
     * the trailing path after the glance
     * endpoint that represent the location of the image
     * or the path to retrieve it.
     * 
     */
    @Import(name="file")
    private @Nullable Output<String> file;

    /**
     * @return the trailing path after the glance
     * endpoint that represent the location of the image
     * or the path to retrieve it.
     * 
     */
    public Optional<Output<String>> file() {
        return Optional.ofNullable(this.file);
    }

    /**
     * If true, image will be hidden from public list.
     * Defaults to false.
     * 
     */
    @Import(name="hidden")
    private @Nullable Output<Boolean> hidden;

    /**
     * @return If true, image will be hidden from public list.
     * Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> hidden() {
        return Optional.ofNullable(this.hidden);
    }

    @Import(name="imageCachePath")
    private @Nullable Output<String> imageCachePath;

    public Optional<Output<String>> imageCachePath() {
        return Optional.ofNullable(this.imageCachePath);
    }

    /**
     * Unique ID (valid UUID) of image to create. Changing
     * this creates a new image.
     * 
     */
    @Import(name="imageId")
    private @Nullable Output<String> imageId;

    /**
     * @return Unique ID (valid UUID) of image to create. Changing
     * this creates a new image.
     * 
     */
    public Optional<Output<String>> imageId() {
        return Optional.ofNullable(this.imageId);
    }

    /**
     * The password of basic auth to download `image_source_url`.
     * 
     */
    @Import(name="imageSourcePassword")
    private @Nullable Output<String> imageSourcePassword;

    /**
     * @return The password of basic auth to download `image_source_url`.
     * 
     */
    public Optional<Output<String>> imageSourcePassword() {
        return Optional.ofNullable(this.imageSourcePassword);
    }

    /**
     * This is the url of the raw image. If `web_download`
     * is not used, then the image will be downloaded in the `image_cache_path` before
     * being uploaded to Glance.
     * Conflicts with `local_file_path`.
     * 
     */
    @Import(name="imageSourceUrl")
    private @Nullable Output<String> imageSourceUrl;

    /**
     * @return This is the url of the raw image. If `web_download`
     * is not used, then the image will be downloaded in the `image_cache_path` before
     * being uploaded to Glance.
     * Conflicts with `local_file_path`.
     * 
     */
    public Optional<Output<String>> imageSourceUrl() {
        return Optional.ofNullable(this.imageSourceUrl);
    }

    /**
     * The username of basic auth to download `image_source_url`.
     * 
     */
    @Import(name="imageSourceUsername")
    private @Nullable Output<String> imageSourceUsername;

    /**
     * @return The username of basic auth to download `image_source_url`.
     * 
     */
    public Optional<Output<String>> imageSourceUsername() {
        return Optional.ofNullable(this.imageSourceUsername);
    }

    /**
     * This is the filepath of the raw image file
     * that will be uploaded to Glance. Conflicts with `image_source_url` and
     * `web_download`.
     * 
     */
    @Import(name="localFilePath")
    private @Nullable Output<String> localFilePath;

    /**
     * @return This is the filepath of the raw image file
     * that will be uploaded to Glance. Conflicts with `image_source_url` and
     * `web_download`.
     * 
     */
    public Optional<Output<String>> localFilePath() {
        return Optional.ofNullable(this.localFilePath);
    }

    /**
     * The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,Object>> metadata;

    /**
     * @return The metadata associated with the image.
     * Image metadata allow for meaningfully define the image properties
     * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
     * 
     */
    public Optional<Output<Map<String,Object>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * Amount of disk space (in GB) required to boot image.
     * Defaults to 0.
     * 
     */
    @Import(name="minDiskGb")
    private @Nullable Output<Integer> minDiskGb;

    /**
     * @return Amount of disk space (in GB) required to boot image.
     * Defaults to 0.
     * 
     */
    public Optional<Output<Integer>> minDiskGb() {
        return Optional.ofNullable(this.minDiskGb);
    }

    /**
     * Amount of ram (in MB) required to boot image.
     * Defauts to 0.
     * 
     */
    @Import(name="minRamMb")
    private @Nullable Output<Integer> minRamMb;

    /**
     * @return Amount of ram (in MB) required to boot image.
     * Defauts to 0.
     * 
     */
    public Optional<Output<Integer>> minRamMb() {
        return Optional.ofNullable(this.minRamMb);
    }

    /**
     * The name of the image.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the image.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The id of the openstack user who owns the image.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return The id of the openstack user who owns the image.
     * 
     */
    public Optional<Output<String>> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * A map of key/value pairs to set freeform
     * information about an image. See the &#34;Notes&#34; section for further
     * information about properties.
     * 
     */
    @Import(name="properties")
    private @Nullable Output<Map<String,Object>> properties;

    /**
     * @return A map of key/value pairs to set freeform
     * information about an image. See the &#34;Notes&#34; section for further
     * information about properties.
     * 
     */
    public Optional<Output<Map<String,Object>>> properties() {
        return Optional.ofNullable(this.properties);
    }

    /**
     * If true, image will not be deletable.
     * Defaults to false.
     * 
     */
    @Import(name="protected")
    private @Nullable Output<Boolean> protected_;

    /**
     * @return If true, image will not be deletable.
     * Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> protected_() {
        return Optional.ofNullable(this.protected_);
    }

    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new Image.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Glance client.
     * A Glance client is needed to create an Image that can be used with
     * a compute instance. If omitted, the `region` argument of the provider
     * is used. Changing this creates a new Image.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The path to the JSON-schema that represent
     * the image or image
     * 
     */
    @Import(name="schema")
    private @Nullable Output<String> schema;

    /**
     * @return The path to the JSON-schema that represent
     * the image or image
     * 
     */
    public Optional<Output<String>> schema() {
        return Optional.ofNullable(this.schema);
    }

    /**
     * The size in bytes of the data associated with the image.
     * 
     */
    @Import(name="sizeBytes")
    private @Nullable Output<Integer> sizeBytes;

    /**
     * @return The size in bytes of the data associated with the image.
     * 
     */
    public Optional<Output<Integer>> sizeBytes() {
        return Optional.ofNullable(this.sizeBytes);
    }

    /**
     * The status of the image. It can be &#34;queued&#34;, &#34;active&#34;
     * or &#34;saving&#34;.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the image. It can be &#34;queued&#34;, &#34;active&#34;
     * or &#34;saving&#34;.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The tags of the image. It must be a list of strings.
     * At this time, it is not possible to delete all tags of an image.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags of the image. It must be a list of strings.
     * At this time, it is not possible to delete all tags of an image.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * (**Deprecated** - use `updated_at` instead)
     * 
     * @deprecated
     * Use updated_at instead
     * 
     */
    @Deprecated /* Use updated_at instead */
    @Import(name="updateAt")
    private @Nullable Output<String> updateAt;

    /**
     * @return (**Deprecated** - use `updated_at` instead)
     * 
     * @deprecated
     * Use updated_at instead
     * 
     */
    @Deprecated /* Use updated_at instead */
    public Optional<Output<String>> updateAt() {
        return Optional.ofNullable(this.updateAt);
    }

    /**
     * The date the image was last updated.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return The date the image was last updated.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    /**
     * If false, the checksum will not be verified
     * once the image is finished uploading. Conflicts with `web_download`.
     * Defaults to true when not using `web_download`.
     * 
     */
    @Import(name="verifyChecksum")
    private @Nullable Output<Boolean> verifyChecksum;

    /**
     * @return If false, the checksum will not be verified
     * once the image is finished uploading. Conflicts with `web_download`.
     * Defaults to true when not using `web_download`.
     * 
     */
    public Optional<Output<Boolean>> verifyChecksum() {
        return Optional.ofNullable(this.verifyChecksum);
    }

    /**
     * The visibility of the image. Must be one of
     * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. The ability to set the
     * visibility depends upon the configuration of the OpenStack cloud.
     * 
     */
    @Import(name="visibility")
    private @Nullable Output<String> visibility;

    /**
     * @return The visibility of the image. Must be one of
     * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. The ability to set the
     * visibility depends upon the configuration of the OpenStack cloud.
     * 
     */
    public Optional<Output<String>> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    /**
     * If true, the &#34;web-download&#34; import method will
     * be used to let Openstack download the image directly from the remote source.
     * Conflicts with `local_file_path`. Defaults to false.
     * 
     */
    @Import(name="webDownload")
    private @Nullable Output<Boolean> webDownload;

    /**
     * @return If true, the &#34;web-download&#34; import method will
     * be used to let Openstack download the image directly from the remote source.
     * Conflicts with `local_file_path`. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> webDownload() {
        return Optional.ofNullable(this.webDownload);
    }

    private ImageState() {}

    private ImageState(ImageState $) {
        this.checksum = $.checksum;
        this.containerFormat = $.containerFormat;
        this.createdAt = $.createdAt;
        this.diskFormat = $.diskFormat;
        this.file = $.file;
        this.hidden = $.hidden;
        this.imageCachePath = $.imageCachePath;
        this.imageId = $.imageId;
        this.imageSourcePassword = $.imageSourcePassword;
        this.imageSourceUrl = $.imageSourceUrl;
        this.imageSourceUsername = $.imageSourceUsername;
        this.localFilePath = $.localFilePath;
        this.metadata = $.metadata;
        this.minDiskGb = $.minDiskGb;
        this.minRamMb = $.minRamMb;
        this.name = $.name;
        this.owner = $.owner;
        this.properties = $.properties;
        this.protected_ = $.protected_;
        this.region = $.region;
        this.schema = $.schema;
        this.sizeBytes = $.sizeBytes;
        this.status = $.status;
        this.tags = $.tags;
        this.updateAt = $.updateAt;
        this.updatedAt = $.updatedAt;
        this.verifyChecksum = $.verifyChecksum;
        this.visibility = $.visibility;
        this.webDownload = $.webDownload;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ImageState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ImageState $;

        public Builder() {
            $ = new ImageState();
        }

        public Builder(ImageState defaults) {
            $ = new ImageState(Objects.requireNonNull(defaults));
        }

        /**
         * @param checksum The checksum of the data associated with the image.
         * 
         * @return builder
         * 
         */
        public Builder checksum(@Nullable Output<String> checksum) {
            $.checksum = checksum;
            return this;
        }

        /**
         * @param checksum The checksum of the data associated with the image.
         * 
         * @return builder
         * 
         */
        public Builder checksum(String checksum) {
            return checksum(Output.of(checksum));
        }

        /**
         * @param containerFormat The container format. Must be one of
         * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;bare&#34;, &#34;ovf&#34;.
         * 
         * @return builder
         * 
         */
        public Builder containerFormat(@Nullable Output<String> containerFormat) {
            $.containerFormat = containerFormat;
            return this;
        }

        /**
         * @param containerFormat The container format. Must be one of
         * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;bare&#34;, &#34;ovf&#34;.
         * 
         * @return builder
         * 
         */
        public Builder containerFormat(String containerFormat) {
            return containerFormat(Output.of(containerFormat));
        }

        /**
         * @param createdAt The date the image was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The date the image was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param diskFormat The disk format. Must be one of
         * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;vhd&#34;, &#34;vmdk&#34;, &#34;raw&#34;, &#34;qcow2&#34;, &#34;vdi&#34;, &#34;iso&#34;.
         * 
         * @return builder
         * 
         */
        public Builder diskFormat(@Nullable Output<String> diskFormat) {
            $.diskFormat = diskFormat;
            return this;
        }

        /**
         * @param diskFormat The disk format. Must be one of
         * &#34;ami&#34;, &#34;ari&#34;, &#34;aki&#34;, &#34;vhd&#34;, &#34;vmdk&#34;, &#34;raw&#34;, &#34;qcow2&#34;, &#34;vdi&#34;, &#34;iso&#34;.
         * 
         * @return builder
         * 
         */
        public Builder diskFormat(String diskFormat) {
            return diskFormat(Output.of(diskFormat));
        }

        /**
         * @param file the trailing path after the glance
         * endpoint that represent the location of the image
         * or the path to retrieve it.
         * 
         * @return builder
         * 
         */
        public Builder file(@Nullable Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file the trailing path after the glance
         * endpoint that represent the location of the image
         * or the path to retrieve it.
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        /**
         * @param hidden If true, image will be hidden from public list.
         * Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder hidden(@Nullable Output<Boolean> hidden) {
            $.hidden = hidden;
            return this;
        }

        /**
         * @param hidden If true, image will be hidden from public list.
         * Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder hidden(Boolean hidden) {
            return hidden(Output.of(hidden));
        }

        public Builder imageCachePath(@Nullable Output<String> imageCachePath) {
            $.imageCachePath = imageCachePath;
            return this;
        }

        public Builder imageCachePath(String imageCachePath) {
            return imageCachePath(Output.of(imageCachePath));
        }

        /**
         * @param imageId Unique ID (valid UUID) of image to create. Changing
         * this creates a new image.
         * 
         * @return builder
         * 
         */
        public Builder imageId(@Nullable Output<String> imageId) {
            $.imageId = imageId;
            return this;
        }

        /**
         * @param imageId Unique ID (valid UUID) of image to create. Changing
         * this creates a new image.
         * 
         * @return builder
         * 
         */
        public Builder imageId(String imageId) {
            return imageId(Output.of(imageId));
        }

        /**
         * @param imageSourcePassword The password of basic auth to download `image_source_url`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourcePassword(@Nullable Output<String> imageSourcePassword) {
            $.imageSourcePassword = imageSourcePassword;
            return this;
        }

        /**
         * @param imageSourcePassword The password of basic auth to download `image_source_url`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourcePassword(String imageSourcePassword) {
            return imageSourcePassword(Output.of(imageSourcePassword));
        }

        /**
         * @param imageSourceUrl This is the url of the raw image. If `web_download`
         * is not used, then the image will be downloaded in the `image_cache_path` before
         * being uploaded to Glance.
         * Conflicts with `local_file_path`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourceUrl(@Nullable Output<String> imageSourceUrl) {
            $.imageSourceUrl = imageSourceUrl;
            return this;
        }

        /**
         * @param imageSourceUrl This is the url of the raw image. If `web_download`
         * is not used, then the image will be downloaded in the `image_cache_path` before
         * being uploaded to Glance.
         * Conflicts with `local_file_path`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourceUrl(String imageSourceUrl) {
            return imageSourceUrl(Output.of(imageSourceUrl));
        }

        /**
         * @param imageSourceUsername The username of basic auth to download `image_source_url`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourceUsername(@Nullable Output<String> imageSourceUsername) {
            $.imageSourceUsername = imageSourceUsername;
            return this;
        }

        /**
         * @param imageSourceUsername The username of basic auth to download `image_source_url`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourceUsername(String imageSourceUsername) {
            return imageSourceUsername(Output.of(imageSourceUsername));
        }

        /**
         * @param localFilePath This is the filepath of the raw image file
         * that will be uploaded to Glance. Conflicts with `image_source_url` and
         * `web_download`.
         * 
         * @return builder
         * 
         */
        public Builder localFilePath(@Nullable Output<String> localFilePath) {
            $.localFilePath = localFilePath;
            return this;
        }

        /**
         * @param localFilePath This is the filepath of the raw image file
         * that will be uploaded to Glance. Conflicts with `image_source_url` and
         * `web_download`.
         * 
         * @return builder
         * 
         */
        public Builder localFilePath(String localFilePath) {
            return localFilePath(Output.of(localFilePath));
        }

        /**
         * @param metadata The metadata associated with the image.
         * Image metadata allow for meaningfully define the image properties
         * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,Object>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata The metadata associated with the image.
         * Image metadata allow for meaningfully define the image properties
         * and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,Object> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param minDiskGb Amount of disk space (in GB) required to boot image.
         * Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder minDiskGb(@Nullable Output<Integer> minDiskGb) {
            $.minDiskGb = minDiskGb;
            return this;
        }

        /**
         * @param minDiskGb Amount of disk space (in GB) required to boot image.
         * Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder minDiskGb(Integer minDiskGb) {
            return minDiskGb(Output.of(minDiskGb));
        }

        /**
         * @param minRamMb Amount of ram (in MB) required to boot image.
         * Defauts to 0.
         * 
         * @return builder
         * 
         */
        public Builder minRamMb(@Nullable Output<Integer> minRamMb) {
            $.minRamMb = minRamMb;
            return this;
        }

        /**
         * @param minRamMb Amount of ram (in MB) required to boot image.
         * Defauts to 0.
         * 
         * @return builder
         * 
         */
        public Builder minRamMb(Integer minRamMb) {
            return minRamMb(Output.of(minRamMb));
        }

        /**
         * @param name The name of the image.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the image.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param owner The id of the openstack user who owns the image.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The id of the openstack user who owns the image.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
        }

        /**
         * @param properties A map of key/value pairs to set freeform
         * information about an image. See the &#34;Notes&#34; section for further
         * information about properties.
         * 
         * @return builder
         * 
         */
        public Builder properties(@Nullable Output<Map<String,Object>> properties) {
            $.properties = properties;
            return this;
        }

        /**
         * @param properties A map of key/value pairs to set freeform
         * information about an image. See the &#34;Notes&#34; section for further
         * information about properties.
         * 
         * @return builder
         * 
         */
        public Builder properties(Map<String,Object> properties) {
            return properties(Output.of(properties));
        }

        /**
         * @param protected_ If true, image will not be deletable.
         * Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder protected_(@Nullable Output<Boolean> protected_) {
            $.protected_ = protected_;
            return this;
        }

        /**
         * @param protected_ If true, image will not be deletable.
         * Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder protected_(Boolean protected_) {
            return protected_(Output.of(protected_));
        }

        /**
         * @param region The region in which to obtain the V2 Glance client.
         * A Glance client is needed to create an Image that can be used with
         * a compute instance. If omitted, the `region` argument of the provider
         * is used. Changing this creates a new Image.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Glance client.
         * A Glance client is needed to create an Image that can be used with
         * a compute instance. If omitted, the `region` argument of the provider
         * is used. Changing this creates a new Image.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param schema The path to the JSON-schema that represent
         * the image or image
         * 
         * @return builder
         * 
         */
        public Builder schema(@Nullable Output<String> schema) {
            $.schema = schema;
            return this;
        }

        /**
         * @param schema The path to the JSON-schema that represent
         * the image or image
         * 
         * @return builder
         * 
         */
        public Builder schema(String schema) {
            return schema(Output.of(schema));
        }

        /**
         * @param sizeBytes The size in bytes of the data associated with the image.
         * 
         * @return builder
         * 
         */
        public Builder sizeBytes(@Nullable Output<Integer> sizeBytes) {
            $.sizeBytes = sizeBytes;
            return this;
        }

        /**
         * @param sizeBytes The size in bytes of the data associated with the image.
         * 
         * @return builder
         * 
         */
        public Builder sizeBytes(Integer sizeBytes) {
            return sizeBytes(Output.of(sizeBytes));
        }

        /**
         * @param status The status of the image. It can be &#34;queued&#34;, &#34;active&#34;
         * or &#34;saving&#34;.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the image. It can be &#34;queued&#34;, &#34;active&#34;
         * or &#34;saving&#34;.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags The tags of the image. It must be a list of strings.
         * At this time, it is not possible to delete all tags of an image.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of the image. It must be a list of strings.
         * At this time, it is not possible to delete all tags of an image.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags of the image. It must be a list of strings.
         * At this time, it is not possible to delete all tags of an image.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param updateAt (**Deprecated** - use `updated_at` instead)
         * 
         * @return builder
         * 
         * @deprecated
         * Use updated_at instead
         * 
         */
        @Deprecated /* Use updated_at instead */
        public Builder updateAt(@Nullable Output<String> updateAt) {
            $.updateAt = updateAt;
            return this;
        }

        /**
         * @param updateAt (**Deprecated** - use `updated_at` instead)
         * 
         * @return builder
         * 
         * @deprecated
         * Use updated_at instead
         * 
         */
        @Deprecated /* Use updated_at instead */
        public Builder updateAt(String updateAt) {
            return updateAt(Output.of(updateAt));
        }

        /**
         * @param updatedAt The date the image was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt The date the image was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        /**
         * @param verifyChecksum If false, the checksum will not be verified
         * once the image is finished uploading. Conflicts with `web_download`.
         * Defaults to true when not using `web_download`.
         * 
         * @return builder
         * 
         */
        public Builder verifyChecksum(@Nullable Output<Boolean> verifyChecksum) {
            $.verifyChecksum = verifyChecksum;
            return this;
        }

        /**
         * @param verifyChecksum If false, the checksum will not be verified
         * once the image is finished uploading. Conflicts with `web_download`.
         * Defaults to true when not using `web_download`.
         * 
         * @return builder
         * 
         */
        public Builder verifyChecksum(Boolean verifyChecksum) {
            return verifyChecksum(Output.of(verifyChecksum));
        }

        /**
         * @param visibility The visibility of the image. Must be one of
         * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. The ability to set the
         * visibility depends upon the configuration of the OpenStack cloud.
         * 
         * @return builder
         * 
         */
        public Builder visibility(@Nullable Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        /**
         * @param visibility The visibility of the image. Must be one of
         * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. The ability to set the
         * visibility depends upon the configuration of the OpenStack cloud.
         * 
         * @return builder
         * 
         */
        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        /**
         * @param webDownload If true, the &#34;web-download&#34; import method will
         * be used to let Openstack download the image directly from the remote source.
         * Conflicts with `local_file_path`. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder webDownload(@Nullable Output<Boolean> webDownload) {
            $.webDownload = webDownload;
            return this;
        }

        /**
         * @param webDownload If true, the &#34;web-download&#34; import method will
         * be used to let Openstack download the image directly from the remote source.
         * Conflicts with `local_file_path`. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder webDownload(Boolean webDownload) {
            return webDownload(Output.of(webDownload));
        }

        public ImageState build() {
            return $;
        }
    }

}
