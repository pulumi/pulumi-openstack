// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images.inputs;

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


public final class GetImagePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetImagePlainArgs Empty = new GetImagePlainArgs();

    /**
     * The container format of the image.
     * 
     */
    @Import(name="containerFormat")
    private @Nullable String containerFormat;

    /**
     * @return The container format of the image.
     * 
     */
    public Optional<String> containerFormat() {
        return Optional.ofNullable(this.containerFormat);
    }

    /**
     * The disk format of the image.
     * 
     */
    @Import(name="diskFormat")
    private @Nullable String diskFormat;

    /**
     * @return The disk format of the image.
     * 
     */
    public Optional<String> diskFormat() {
        return Optional.ofNullable(this.diskFormat);
    }

    /**
     * Whether or not the image is hidden from public list.
     * 
     */
    @Import(name="hidden")
    private @Nullable Boolean hidden;

    /**
     * @return Whether or not the image is hidden from public list.
     * 
     */
    public Optional<Boolean> hidden() {
        return Optional.ofNullable(this.hidden);
    }

    /**
     * The status of the image. Must be one of
     * &#34;accepted&#34;, &#34;pending&#34;, &#34;rejected&#34;, or &#34;all&#34;.
     * 
     */
    @Import(name="memberStatus")
    private @Nullable String memberStatus;

    /**
     * @return The status of the image. Must be one of
     * &#34;accepted&#34;, &#34;pending&#34;, &#34;rejected&#34;, or &#34;all&#34;.
     * 
     */
    public Optional<String> memberStatus() {
        return Optional.ofNullable(this.memberStatus);
    }

    /**
     * If more than one result is returned, use the most
     * recent image.
     * 
     */
    @Import(name="mostRecent")
    private @Nullable Boolean mostRecent;

    /**
     * @return If more than one result is returned, use the most
     * recent image.
     * 
     */
    public Optional<Boolean> mostRecent() {
        return Optional.ofNullable(this.mostRecent);
    }

    /**
     * The name of the image. Cannot be used simultaneously with
     * `name_regex`.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the image. Cannot be used simultaneously with
     * `name_regex`.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The regular expressian of the name of the image.
     * Cannot be used simultaneously with `name`. Unlike filtering by `name` the
     * `name_regex` filtering does by client on the result of OpenStack search
     * query.
     * 
     */
    @Import(name="nameRegex")
    private @Nullable String nameRegex;

    /**
     * @return The regular expressian of the name of the image.
     * Cannot be used simultaneously with `name`. Unlike filtering by `name` the
     * `name_regex` filtering does by client on the result of OpenStack search
     * query.
     * 
     */
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The owner (UUID) of the image.
     * 
     */
    @Import(name="owner")
    private @Nullable String owner;

    /**
     * @return The owner (UUID) of the image.
     * 
     */
    public Optional<String> owner() {
        return Optional.ofNullable(this.owner);
    }

    /**
     * a map of key/value pairs to match an image with.
     * All specified properties must be matched. Unlike other options filtering by
     * `properties` does by client on the result of OpenStack search query.
     * Filtering is applied if server responce contains at least 2 images. In case
     * there is only one image the `properties` ignores.
     * 
     */
    @Import(name="properties")
    private @Nullable Map<String,Object> properties;

    /**
     * @return a map of key/value pairs to match an image with.
     * All specified properties must be matched. Unlike other options filtering by
     * `properties` does by client on the result of OpenStack search query.
     * Filtering is applied if server responce contains at least 2 images. In case
     * there is only one image the `properties` ignores.
     * 
     */
    public Optional<Map<String,Object>> properties() {
        return Optional.ofNullable(this.properties);
    }

    /**
     * The region in which to obtain the V2 Glance client. A
     * Glance client is needed to create an Image that can be used with a compute
     * instance. If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V2 Glance client. A
     * Glance client is needed to create an Image that can be used with a compute
     * instance. If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The maximum size (in bytes) of the image to return.
     * 
     */
    @Import(name="sizeMax")
    private @Nullable Integer sizeMax;

    /**
     * @return The maximum size (in bytes) of the image to return.
     * 
     */
    public Optional<Integer> sizeMax() {
        return Optional.ofNullable(this.sizeMax);
    }

    /**
     * The minimum size (in bytes) of the image to return.
     * 
     */
    @Import(name="sizeMin")
    private @Nullable Integer sizeMin;

    /**
     * @return The minimum size (in bytes) of the image to return.
     * 
     */
    public Optional<Integer> sizeMin() {
        return Optional.ofNullable(this.sizeMin);
    }

    /**
     * Sorts the response by one or more attribute and sort
     * direction combinations. You can also set multiple sort keys and directions.
     * Default direction is `desc`. Use the comma (,) character to separate multiple
     * values. For example expression `sort = &#34;name:asc,status&#34;` sorts ascending by
     * name and descending by status.
     * 
     */
    @Import(name="sort")
    private @Nullable String sort;

    /**
     * @return Sorts the response by one or more attribute and sort
     * direction combinations. You can also set multiple sort keys and directions.
     * Default direction is `desc`. Use the comma (,) character to separate multiple
     * values. For example expression `sort = &#34;name:asc,status&#34;` sorts ascending by
     * name and descending by status.
     * 
     */
    public Optional<String> sort() {
        return Optional.ofNullable(this.sort);
    }

    /**
     * Search for images with a specific tag.
     * 
     */
    @Import(name="tag")
    private @Nullable String tag;

    /**
     * @return Search for images with a specific tag.
     * 
     */
    public Optional<String> tag() {
        return Optional.ofNullable(this.tag);
    }

    /**
     * A list of tags required to be set on the image (all
     * specified tags must be in the images tag list for it to be matched).
     * 
     */
    @Import(name="tags")
    private @Nullable List<String> tags;

    /**
     * @return A list of tags required to be set on the image (all
     * specified tags must be in the images tag list for it to be matched).
     * 
     */
    public Optional<List<String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The visibility of the image. Must be one of
     * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. Defaults to &#34;private&#34;.
     * 
     */
    @Import(name="visibility")
    private @Nullable String visibility;

    /**
     * @return The visibility of the image. Must be one of
     * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. Defaults to &#34;private&#34;.
     * 
     */
    public Optional<String> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    private GetImagePlainArgs() {}

    private GetImagePlainArgs(GetImagePlainArgs $) {
        this.containerFormat = $.containerFormat;
        this.diskFormat = $.diskFormat;
        this.hidden = $.hidden;
        this.memberStatus = $.memberStatus;
        this.mostRecent = $.mostRecent;
        this.name = $.name;
        this.nameRegex = $.nameRegex;
        this.owner = $.owner;
        this.properties = $.properties;
        this.region = $.region;
        this.sizeMax = $.sizeMax;
        this.sizeMin = $.sizeMin;
        this.sort = $.sort;
        this.tag = $.tag;
        this.tags = $.tags;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetImagePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetImagePlainArgs $;

        public Builder() {
            $ = new GetImagePlainArgs();
        }

        public Builder(GetImagePlainArgs defaults) {
            $ = new GetImagePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param containerFormat The container format of the image.
         * 
         * @return builder
         * 
         */
        public Builder containerFormat(@Nullable String containerFormat) {
            $.containerFormat = containerFormat;
            return this;
        }

        /**
         * @param diskFormat The disk format of the image.
         * 
         * @return builder
         * 
         */
        public Builder diskFormat(@Nullable String diskFormat) {
            $.diskFormat = diskFormat;
            return this;
        }

        /**
         * @param hidden Whether or not the image is hidden from public list.
         * 
         * @return builder
         * 
         */
        public Builder hidden(@Nullable Boolean hidden) {
            $.hidden = hidden;
            return this;
        }

        /**
         * @param memberStatus The status of the image. Must be one of
         * &#34;accepted&#34;, &#34;pending&#34;, &#34;rejected&#34;, or &#34;all&#34;.
         * 
         * @return builder
         * 
         */
        public Builder memberStatus(@Nullable String memberStatus) {
            $.memberStatus = memberStatus;
            return this;
        }

        /**
         * @param mostRecent If more than one result is returned, use the most
         * recent image.
         * 
         * @return builder
         * 
         */
        public Builder mostRecent(@Nullable Boolean mostRecent) {
            $.mostRecent = mostRecent;
            return this;
        }

        /**
         * @param name The name of the image. Cannot be used simultaneously with
         * `name_regex`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param nameRegex The regular expressian of the name of the image.
         * Cannot be used simultaneously with `name`. Unlike filtering by `name` the
         * `name_regex` filtering does by client on the result of OpenStack search
         * query.
         * 
         * @return builder
         * 
         */
        public Builder nameRegex(@Nullable String nameRegex) {
            $.nameRegex = nameRegex;
            return this;
        }

        /**
         * @param owner The owner (UUID) of the image.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable String owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param properties a map of key/value pairs to match an image with.
         * All specified properties must be matched. Unlike other options filtering by
         * `properties` does by client on the result of OpenStack search query.
         * Filtering is applied if server responce contains at least 2 images. In case
         * there is only one image the `properties` ignores.
         * 
         * @return builder
         * 
         */
        public Builder properties(@Nullable Map<String,Object> properties) {
            $.properties = properties;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Glance client. A
         * Glance client is needed to create an Image that can be used with a compute
         * instance. If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        /**
         * @param sizeMax The maximum size (in bytes) of the image to return.
         * 
         * @return builder
         * 
         */
        public Builder sizeMax(@Nullable Integer sizeMax) {
            $.sizeMax = sizeMax;
            return this;
        }

        /**
         * @param sizeMin The minimum size (in bytes) of the image to return.
         * 
         * @return builder
         * 
         */
        public Builder sizeMin(@Nullable Integer sizeMin) {
            $.sizeMin = sizeMin;
            return this;
        }

        /**
         * @param sort Sorts the response by one or more attribute and sort
         * direction combinations. You can also set multiple sort keys and directions.
         * Default direction is `desc`. Use the comma (,) character to separate multiple
         * values. For example expression `sort = &#34;name:asc,status&#34;` sorts ascending by
         * name and descending by status.
         * 
         * @return builder
         * 
         */
        public Builder sort(@Nullable String sort) {
            $.sort = sort;
            return this;
        }

        /**
         * @param tag Search for images with a specific tag.
         * 
         * @return builder
         * 
         */
        public Builder tag(@Nullable String tag) {
            $.tag = tag;
            return this;
        }

        /**
         * @param tags A list of tags required to be set on the image (all
         * specified tags must be in the images tag list for it to be matched).
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable List<String> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of tags required to be set on the image (all
         * specified tags must be in the images tag list for it to be matched).
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param visibility The visibility of the image. Must be one of
         * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. Defaults to &#34;private&#34;.
         * 
         * @return builder
         * 
         */
        public Builder visibility(@Nullable String visibility) {
            $.visibility = visibility;
            return this;
        }

        public GetImagePlainArgs build() {
            return $;
        }
    }

}
