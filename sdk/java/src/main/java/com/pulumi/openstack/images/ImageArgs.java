// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ImageArgs extends com.pulumi.resources.ResourceArgs {

    public static final ImageArgs Empty = new ImageArgs();

    /**
     * The container format. Must be one of &#34;bare&#34;,
     * &#34;ovf&#34;, &#34;aki&#34;, &#34;ari&#34;, &#34;ami&#34;, &#34;ova&#34;, &#34;docker&#34;, &#34;compressed&#34;.
     * 
     */
    @Import(name="containerFormat", required=true)
    private Output<String> containerFormat;

    /**
     * @return The container format. Must be one of &#34;bare&#34;,
     * &#34;ovf&#34;, &#34;aki&#34;, &#34;ari&#34;, &#34;ami&#34;, &#34;ova&#34;, &#34;docker&#34;, &#34;compressed&#34;.
     * 
     */
    public Output<String> containerFormat() {
        return this.containerFormat;
    }

    /**
     * If true, this provider will decompress downloaded
     * image before uploading it to OpenStack. Decompression algorithm is chosen by
     * checking &#34;Content-Type&#34; or `Content-Disposition` header to detect the
     * filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
     * Defaults to false. Changing this creates a new Image.
     * 
     */
    @Import(name="decompress")
    private @Nullable Output<Boolean> decompress;

    /**
     * @return If true, this provider will decompress downloaded
     * image before uploading it to OpenStack. Decompression algorithm is chosen by
     * checking &#34;Content-Type&#34; or `Content-Disposition` header to detect the
     * filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
     * Defaults to false. Changing this creates a new Image.
     * 
     */
    public Optional<Output<Boolean>> decompress() {
        return Optional.ofNullable(this.decompress);
    }

    /**
     * The disk format. Must be one of &#34;raw&#34;, &#34;vhd&#34;,
     * &#34;vhdx&#34;, &#34;vmdk&#34;, &#34;vdi&#34;, &#34;iso&#34;, &#34;ploop&#34;, &#34;qcow2&#34;, &#34;aki&#34;, &#34;ari&#34;, &#34;ami&#34;
     * 
     */
    @Import(name="diskFormat", required=true)
    private Output<String> diskFormat;

    /**
     * @return The disk format. Must be one of &#34;raw&#34;, &#34;vhd&#34;,
     * &#34;vhdx&#34;, &#34;vmdk&#34;, &#34;vdi&#34;, &#34;iso&#34;, &#34;ploop&#34;, &#34;qcow2&#34;, &#34;aki&#34;, &#34;ari&#34;, &#34;ami&#34;
     * 
     */
    public Output<String> diskFormat() {
        return this.diskFormat;
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
     * The password of basic auth to download
     * `image_source_url`.
     * 
     */
    @Import(name="imageSourcePassword")
    private @Nullable Output<String> imageSourcePassword;

    /**
     * @return The password of basic auth to download
     * `image_source_url`.
     * 
     */
    public Optional<Output<String>> imageSourcePassword() {
        return Optional.ofNullable(this.imageSourcePassword);
    }

    /**
     * This is the url of the raw image. If
     * `web_download` is not used, then the image will be downloaded in the
     * `image_cache_path` before being uploaded to Glance. Conflicts with
     * `local_file_path`.
     * 
     */
    @Import(name="imageSourceUrl")
    private @Nullable Output<String> imageSourceUrl;

    /**
     * @return This is the url of the raw image. If
     * `web_download` is not used, then the image will be downloaded in the
     * `image_cache_path` before being uploaded to Glance. Conflicts with
     * `local_file_path`.
     * 
     */
    public Optional<Output<String>> imageSourceUrl() {
        return Optional.ofNullable(this.imageSourceUrl);
    }

    /**
     * The username of basic auth to download
     * `image_source_url`.
     * 
     */
    @Import(name="imageSourceUsername")
    private @Nullable Output<String> imageSourceUsername;

    /**
     * @return The username of basic auth to download
     * `image_source_url`.
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
     * Amount of disk space (in GB) required to boot
     * image. Defaults to 0.
     * 
     */
    @Import(name="minDiskGb")
    private @Nullable Output<Integer> minDiskGb;

    /**
     * @return Amount of disk space (in GB) required to boot
     * image. Defaults to 0.
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
     * A map of key/value pairs to set freeform
     * information about an image. See the &#34;Notes&#34; section for further information
     * about properties.
     * 
     */
    @Import(name="properties")
    private @Nullable Output<Map<String,Object>> properties;

    /**
     * @return A map of key/value pairs to set freeform
     * information about an image. See the &#34;Notes&#34; section for further information
     * about properties.
     * 
     */
    public Optional<Output<Map<String,Object>>> properties() {
        return Optional.ofNullable(this.properties);
    }

    /**
     * If true, image will not be deletable. Defaults to
     * false.
     * 
     */
    @Import(name="protected")
    private @Nullable Output<Boolean> protected_;

    /**
     * @return If true, image will not be deletable. Defaults to
     * false.
     * 
     */
    public Optional<Output<Boolean>> protected_() {
        return Optional.ofNullable(this.protected_);
    }

    /**
     * The region in which to obtain the V2 Glance client. A
     * Glance client is needed to create an Image that can be used with a compute
     * instance. If omitted, the `region` argument of the provider is used. Changing
     * this creates a new Image.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Glance client. A
     * Glance client is needed to create an Image that can be used with a compute
     * instance. If omitted, the `region` argument of the provider is used. Changing
     * this creates a new Image.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The tags of the image. It must be a list of strings. At
     * this time, it is not possible to delete all tags of an image.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return The tags of the image. It must be a list of strings. At
     * this time, it is not possible to delete all tags of an image.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * If false, the checksum will not be verified
     * once the image is finished uploading. Conflicts with `web_download`. Defaults
     * to true when not using `web_download`.
     * 
     */
    @Import(name="verifyChecksum")
    private @Nullable Output<Boolean> verifyChecksum;

    /**
     * @return If false, the checksum will not be verified
     * once the image is finished uploading. Conflicts with `web_download`. Defaults
     * to true when not using `web_download`.
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
     * If true, the &#34;web-download&#34; import method will be
     * used to let Openstack download the image directly from the remote source.
     * Conflicts with `local_file_path`. Defaults to false.
     * 
     */
    @Import(name="webDownload")
    private @Nullable Output<Boolean> webDownload;

    /**
     * @return If true, the &#34;web-download&#34; import method will be
     * used to let Openstack download the image directly from the remote source.
     * Conflicts with `local_file_path`. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> webDownload() {
        return Optional.ofNullable(this.webDownload);
    }

    private ImageArgs() {}

    private ImageArgs(ImageArgs $) {
        this.containerFormat = $.containerFormat;
        this.decompress = $.decompress;
        this.diskFormat = $.diskFormat;
        this.hidden = $.hidden;
        this.imageCachePath = $.imageCachePath;
        this.imageId = $.imageId;
        this.imageSourcePassword = $.imageSourcePassword;
        this.imageSourceUrl = $.imageSourceUrl;
        this.imageSourceUsername = $.imageSourceUsername;
        this.localFilePath = $.localFilePath;
        this.minDiskGb = $.minDiskGb;
        this.minRamMb = $.minRamMb;
        this.name = $.name;
        this.properties = $.properties;
        this.protected_ = $.protected_;
        this.region = $.region;
        this.tags = $.tags;
        this.verifyChecksum = $.verifyChecksum;
        this.visibility = $.visibility;
        this.webDownload = $.webDownload;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ImageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ImageArgs $;

        public Builder() {
            $ = new ImageArgs();
        }

        public Builder(ImageArgs defaults) {
            $ = new ImageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param containerFormat The container format. Must be one of &#34;bare&#34;,
         * &#34;ovf&#34;, &#34;aki&#34;, &#34;ari&#34;, &#34;ami&#34;, &#34;ova&#34;, &#34;docker&#34;, &#34;compressed&#34;.
         * 
         * @return builder
         * 
         */
        public Builder containerFormat(Output<String> containerFormat) {
            $.containerFormat = containerFormat;
            return this;
        }

        /**
         * @param containerFormat The container format. Must be one of &#34;bare&#34;,
         * &#34;ovf&#34;, &#34;aki&#34;, &#34;ari&#34;, &#34;ami&#34;, &#34;ova&#34;, &#34;docker&#34;, &#34;compressed&#34;.
         * 
         * @return builder
         * 
         */
        public Builder containerFormat(String containerFormat) {
            return containerFormat(Output.of(containerFormat));
        }

        /**
         * @param decompress If true, this provider will decompress downloaded
         * image before uploading it to OpenStack. Decompression algorithm is chosen by
         * checking &#34;Content-Type&#34; or `Content-Disposition` header to detect the
         * filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
         * Defaults to false. Changing this creates a new Image.
         * 
         * @return builder
         * 
         */
        public Builder decompress(@Nullable Output<Boolean> decompress) {
            $.decompress = decompress;
            return this;
        }

        /**
         * @param decompress If true, this provider will decompress downloaded
         * image before uploading it to OpenStack. Decompression algorithm is chosen by
         * checking &#34;Content-Type&#34; or `Content-Disposition` header to detect the
         * filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
         * Defaults to false. Changing this creates a new Image.
         * 
         * @return builder
         * 
         */
        public Builder decompress(Boolean decompress) {
            return decompress(Output.of(decompress));
        }

        /**
         * @param diskFormat The disk format. Must be one of &#34;raw&#34;, &#34;vhd&#34;,
         * &#34;vhdx&#34;, &#34;vmdk&#34;, &#34;vdi&#34;, &#34;iso&#34;, &#34;ploop&#34;, &#34;qcow2&#34;, &#34;aki&#34;, &#34;ari&#34;, &#34;ami&#34;
         * 
         * @return builder
         * 
         */
        public Builder diskFormat(Output<String> diskFormat) {
            $.diskFormat = diskFormat;
            return this;
        }

        /**
         * @param diskFormat The disk format. Must be one of &#34;raw&#34;, &#34;vhd&#34;,
         * &#34;vhdx&#34;, &#34;vmdk&#34;, &#34;vdi&#34;, &#34;iso&#34;, &#34;ploop&#34;, &#34;qcow2&#34;, &#34;aki&#34;, &#34;ari&#34;, &#34;ami&#34;
         * 
         * @return builder
         * 
         */
        public Builder diskFormat(String diskFormat) {
            return diskFormat(Output.of(diskFormat));
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
         * @param imageSourcePassword The password of basic auth to download
         * `image_source_url`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourcePassword(@Nullable Output<String> imageSourcePassword) {
            $.imageSourcePassword = imageSourcePassword;
            return this;
        }

        /**
         * @param imageSourcePassword The password of basic auth to download
         * `image_source_url`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourcePassword(String imageSourcePassword) {
            return imageSourcePassword(Output.of(imageSourcePassword));
        }

        /**
         * @param imageSourceUrl This is the url of the raw image. If
         * `web_download` is not used, then the image will be downloaded in the
         * `image_cache_path` before being uploaded to Glance. Conflicts with
         * `local_file_path`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourceUrl(@Nullable Output<String> imageSourceUrl) {
            $.imageSourceUrl = imageSourceUrl;
            return this;
        }

        /**
         * @param imageSourceUrl This is the url of the raw image. If
         * `web_download` is not used, then the image will be downloaded in the
         * `image_cache_path` before being uploaded to Glance. Conflicts with
         * `local_file_path`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourceUrl(String imageSourceUrl) {
            return imageSourceUrl(Output.of(imageSourceUrl));
        }

        /**
         * @param imageSourceUsername The username of basic auth to download
         * `image_source_url`.
         * 
         * @return builder
         * 
         */
        public Builder imageSourceUsername(@Nullable Output<String> imageSourceUsername) {
            $.imageSourceUsername = imageSourceUsername;
            return this;
        }

        /**
         * @param imageSourceUsername The username of basic auth to download
         * `image_source_url`.
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
         * @param minDiskGb Amount of disk space (in GB) required to boot
         * image. Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder minDiskGb(@Nullable Output<Integer> minDiskGb) {
            $.minDiskGb = minDiskGb;
            return this;
        }

        /**
         * @param minDiskGb Amount of disk space (in GB) required to boot
         * image. Defaults to 0.
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
         * @param properties A map of key/value pairs to set freeform
         * information about an image. See the &#34;Notes&#34; section for further information
         * about properties.
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
         * information about an image. See the &#34;Notes&#34; section for further information
         * about properties.
         * 
         * @return builder
         * 
         */
        public Builder properties(Map<String,Object> properties) {
            return properties(Output.of(properties));
        }

        /**
         * @param protected_ If true, image will not be deletable. Defaults to
         * false.
         * 
         * @return builder
         * 
         */
        public Builder protected_(@Nullable Output<Boolean> protected_) {
            $.protected_ = protected_;
            return this;
        }

        /**
         * @param protected_ If true, image will not be deletable. Defaults to
         * false.
         * 
         * @return builder
         * 
         */
        public Builder protected_(Boolean protected_) {
            return protected_(Output.of(protected_));
        }

        /**
         * @param region The region in which to obtain the V2 Glance client. A
         * Glance client is needed to create an Image that can be used with a compute
         * instance. If omitted, the `region` argument of the provider is used. Changing
         * this creates a new Image.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Glance client. A
         * Glance client is needed to create an Image that can be used with a compute
         * instance. If omitted, the `region` argument of the provider is used. Changing
         * this creates a new Image.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tags The tags of the image. It must be a list of strings. At
         * this time, it is not possible to delete all tags of an image.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags The tags of the image. It must be a list of strings. At
         * this time, it is not possible to delete all tags of an image.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags The tags of the image. It must be a list of strings. At
         * this time, it is not possible to delete all tags of an image.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param verifyChecksum If false, the checksum will not be verified
         * once the image is finished uploading. Conflicts with `web_download`. Defaults
         * to true when not using `web_download`.
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
         * once the image is finished uploading. Conflicts with `web_download`. Defaults
         * to true when not using `web_download`.
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
         * @param webDownload If true, the &#34;web-download&#34; import method will be
         * used to let Openstack download the image directly from the remote source.
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
         * @param webDownload If true, the &#34;web-download&#34; import method will be
         * used to let Openstack download the image directly from the remote source.
         * Conflicts with `local_file_path`. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder webDownload(Boolean webDownload) {
            return webDownload(Output.of(webDownload));
        }

        public ImageArgs build() {
            if ($.containerFormat == null) {
                throw new MissingRequiredPropertyException("ImageArgs", "containerFormat");
            }
            if ($.diskFormat == null) {
                throw new MissingRequiredPropertyException("ImageArgs", "diskFormat");
            }
            return $;
        }
    }

}