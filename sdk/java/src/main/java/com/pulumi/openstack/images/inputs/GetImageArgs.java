// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetImageArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetImageArgs Empty = new GetImageArgs();

    /**
     * The container format of the image.
     * 
     */
    @Import(name="containerFormat")
    private @Nullable Output<String> containerFormat;

    /**
     * @return The container format of the image.
     * 
     */
    public Optional<Output<String>> containerFormat() {
        return Optional.ofNullable(this.containerFormat);
    }

    /**
     * The disk format of the image.
     * 
     */
    @Import(name="diskFormat")
    private @Nullable Output<String> diskFormat;

    /**
     * @return The disk format of the image.
     * 
     */
    public Optional<Output<String>> diskFormat() {
        return Optional.ofNullable(this.diskFormat);
    }

    /**
     * Whether or not the image is hidden from public list.
     * 
     */
    @Import(name="hidden")
    private @Nullable Output<Boolean> hidden;

    /**
     * @return Whether or not the image is hidden from public list.
     * 
     */
    public Optional<Output<Boolean>> hidden() {
        return Optional.ofNullable(this.hidden);
    }

    /**
     * The status of the image. Must be one of
     * &#34;accepted&#34;, &#34;pending&#34;, &#34;rejected&#34;, or &#34;all&#34;.
     * 
     */
    @Import(name="memberStatus")
    private @Nullable Output<String> memberStatus;

    /**
     * @return The status of the image. Must be one of
     * &#34;accepted&#34;, &#34;pending&#34;, &#34;rejected&#34;, or &#34;all&#34;.
     * 
     */
    public Optional<Output<String>> memberStatus() {
        return Optional.ofNullable(this.memberStatus);
    }

    /**
     * If more than one result is returned, use the most
     * recent image.
     * 
     */
    @Import(name="mostRecent")
    private @Nullable Output<Boolean> mostRecent;

    /**
     * @return If more than one result is returned, use the most
     * recent image.
     * 
     */
    public Optional<Output<Boolean>> mostRecent() {
        return Optional.ofNullable(this.mostRecent);
    }

    /**
     * The name of the image. Cannot be used simultaneously with
     * `name_regex`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the image. Cannot be used simultaneously with
     * `name_regex`.
     * 
     */
    public Optional<Output<String>> name() {
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
    private @Nullable Output<String> nameRegex;

    /**
     * @return The regular expressian of the name of the image.
     * Cannot be used simultaneously with `name`. Unlike filtering by `name` the
     * `name_regex` filtering does by client on the result of OpenStack search
     * query.
     * 
     */
    public Optional<Output<String>> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }

    /**
     * The owner (UUID) of the image.
     * 
     */
    @Import(name="owner")
    private @Nullable Output<String> owner;

    /**
     * @return The owner (UUID) of the image.
     * 
     */
    public Optional<Output<String>> owner() {
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
    private @Nullable Output<Map<String,String>> properties;

    /**
     * @return a map of key/value pairs to match an image with.
     * All specified properties must be matched. Unlike other options filtering by
     * `properties` does by client on the result of OpenStack search query.
     * Filtering is applied if server responce contains at least 2 images. In case
     * there is only one image the `properties` ignores.
     * 
     */
    public Optional<Output<Map<String,String>>> properties() {
        return Optional.ofNullable(this.properties);
    }

    /**
     * The region in which to obtain the V2 Glance client. A
     * Glance client is needed to create an Image that can be used with a compute
     * instance. If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Glance client. A
     * Glance client is needed to create an Image that can be used with a compute
     * instance. If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The maximum size (in bytes) of the image to return.
     * 
     */
    @Import(name="sizeMax")
    private @Nullable Output<Integer> sizeMax;

    /**
     * @return The maximum size (in bytes) of the image to return.
     * 
     */
    public Optional<Output<Integer>> sizeMax() {
        return Optional.ofNullable(this.sizeMax);
    }

    /**
     * The minimum size (in bytes) of the image to return.
     * 
     */
    @Import(name="sizeMin")
    private @Nullable Output<Integer> sizeMin;

    /**
     * @return The minimum size (in bytes) of the image to return.
     * 
     */
    public Optional<Output<Integer>> sizeMin() {
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
    private @Nullable Output<String> sort;

    /**
     * @return Sorts the response by one or more attribute and sort
     * direction combinations. You can also set multiple sort keys and directions.
     * Default direction is `desc`. Use the comma (,) character to separate multiple
     * values. For example expression `sort = &#34;name:asc,status&#34;` sorts ascending by
     * name and descending by status.
     * 
     */
    public Optional<Output<String>> sort() {
        return Optional.ofNullable(this.sort);
    }

    /**
     * Search for images with a specific tag.
     * 
     */
    @Import(name="tag")
    private @Nullable Output<String> tag;

    /**
     * @return Search for images with a specific tag.
     * 
     */
    public Optional<Output<String>> tag() {
        return Optional.ofNullable(this.tag);
    }

    /**
     * A list of tags required to be set on the image (all
     * specified tags must be in the images tag list for it to be matched).
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of tags required to be set on the image (all
     * specified tags must be in the images tag list for it to be matched).
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The visibility of the image. Must be one of
     * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. Defaults to &#34;private&#34;.
     * 
     */
    @Import(name="visibility")
    private @Nullable Output<String> visibility;

    /**
     * @return The visibility of the image. Must be one of
     * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. Defaults to &#34;private&#34;.
     * 
     */
    public Optional<Output<String>> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    private GetImageArgs() {}

    private GetImageArgs(GetImageArgs $) {
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
    public static Builder builder(GetImageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetImageArgs $;

        public Builder() {
            $ = new GetImageArgs();
        }

        public Builder(GetImageArgs defaults) {
            $ = new GetImageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param containerFormat The container format of the image.
         * 
         * @return builder
         * 
         */
        public Builder containerFormat(@Nullable Output<String> containerFormat) {
            $.containerFormat = containerFormat;
            return this;
        }

        /**
         * @param containerFormat The container format of the image.
         * 
         * @return builder
         * 
         */
        public Builder containerFormat(String containerFormat) {
            return containerFormat(Output.of(containerFormat));
        }

        /**
         * @param diskFormat The disk format of the image.
         * 
         * @return builder
         * 
         */
        public Builder diskFormat(@Nullable Output<String> diskFormat) {
            $.diskFormat = diskFormat;
            return this;
        }

        /**
         * @param diskFormat The disk format of the image.
         * 
         * @return builder
         * 
         */
        public Builder diskFormat(String diskFormat) {
            return diskFormat(Output.of(diskFormat));
        }

        /**
         * @param hidden Whether or not the image is hidden from public list.
         * 
         * @return builder
         * 
         */
        public Builder hidden(@Nullable Output<Boolean> hidden) {
            $.hidden = hidden;
            return this;
        }

        /**
         * @param hidden Whether or not the image is hidden from public list.
         * 
         * @return builder
         * 
         */
        public Builder hidden(Boolean hidden) {
            return hidden(Output.of(hidden));
        }

        /**
         * @param memberStatus The status of the image. Must be one of
         * &#34;accepted&#34;, &#34;pending&#34;, &#34;rejected&#34;, or &#34;all&#34;.
         * 
         * @return builder
         * 
         */
        public Builder memberStatus(@Nullable Output<String> memberStatus) {
            $.memberStatus = memberStatus;
            return this;
        }

        /**
         * @param memberStatus The status of the image. Must be one of
         * &#34;accepted&#34;, &#34;pending&#34;, &#34;rejected&#34;, or &#34;all&#34;.
         * 
         * @return builder
         * 
         */
        public Builder memberStatus(String memberStatus) {
            return memberStatus(Output.of(memberStatus));
        }

        /**
         * @param mostRecent If more than one result is returned, use the most
         * recent image.
         * 
         * @return builder
         * 
         */
        public Builder mostRecent(@Nullable Output<Boolean> mostRecent) {
            $.mostRecent = mostRecent;
            return this;
        }

        /**
         * @param mostRecent If more than one result is returned, use the most
         * recent image.
         * 
         * @return builder
         * 
         */
        public Builder mostRecent(Boolean mostRecent) {
            return mostRecent(Output.of(mostRecent));
        }

        /**
         * @param name The name of the image. Cannot be used simultaneously with
         * `name_regex`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the image. Cannot be used simultaneously with
         * `name_regex`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
        public Builder nameRegex(@Nullable Output<String> nameRegex) {
            $.nameRegex = nameRegex;
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
        public Builder nameRegex(String nameRegex) {
            return nameRegex(Output.of(nameRegex));
        }

        /**
         * @param owner The owner (UUID) of the image.
         * 
         * @return builder
         * 
         */
        public Builder owner(@Nullable Output<String> owner) {
            $.owner = owner;
            return this;
        }

        /**
         * @param owner The owner (UUID) of the image.
         * 
         * @return builder
         * 
         */
        public Builder owner(String owner) {
            return owner(Output.of(owner));
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
        public Builder properties(@Nullable Output<Map<String,String>> properties) {
            $.properties = properties;
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
        public Builder properties(Map<String,String> properties) {
            return properties(Output.of(properties));
        }

        /**
         * @param region The region in which to obtain the V2 Glance client. A
         * Glance client is needed to create an Image that can be used with a compute
         * instance. If omitted, the `region` argument of the provider is used.
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
         * instance. If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param sizeMax The maximum size (in bytes) of the image to return.
         * 
         * @return builder
         * 
         */
        public Builder sizeMax(@Nullable Output<Integer> sizeMax) {
            $.sizeMax = sizeMax;
            return this;
        }

        /**
         * @param sizeMax The maximum size (in bytes) of the image to return.
         * 
         * @return builder
         * 
         */
        public Builder sizeMax(Integer sizeMax) {
            return sizeMax(Output.of(sizeMax));
        }

        /**
         * @param sizeMin The minimum size (in bytes) of the image to return.
         * 
         * @return builder
         * 
         */
        public Builder sizeMin(@Nullable Output<Integer> sizeMin) {
            $.sizeMin = sizeMin;
            return this;
        }

        /**
         * @param sizeMin The minimum size (in bytes) of the image to return.
         * 
         * @return builder
         * 
         */
        public Builder sizeMin(Integer sizeMin) {
            return sizeMin(Output.of(sizeMin));
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
        public Builder sort(@Nullable Output<String> sort) {
            $.sort = sort;
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
        public Builder sort(String sort) {
            return sort(Output.of(sort));
        }

        /**
         * @param tag Search for images with a specific tag.
         * 
         * @return builder
         * 
         */
        public Builder tag(@Nullable Output<String> tag) {
            $.tag = tag;
            return this;
        }

        /**
         * @param tag Search for images with a specific tag.
         * 
         * @return builder
         * 
         */
        public Builder tag(String tag) {
            return tag(Output.of(tag));
        }

        /**
         * @param tags A list of tags required to be set on the image (all
         * specified tags must be in the images tag list for it to be matched).
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
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
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
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
        public Builder visibility(@Nullable Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        /**
         * @param visibility The visibility of the image. Must be one of
         * &#34;public&#34;, &#34;private&#34;, &#34;community&#34;, or &#34;shared&#34;. Defaults to &#34;private&#34;.
         * 
         * @return builder
         * 
         */
        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        public GetImageArgs build() {
            return $;
        }
    }

}
