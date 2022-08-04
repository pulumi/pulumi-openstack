// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.objectstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TempUrlState extends com.pulumi.resources.ResourceArgs {

    public static final TempUrlState Empty = new TempUrlState();

    /**
     * The container name the object belongs to.
     * 
     */
    @Import(name="container")
    private @Nullable Output<String> container;

    /**
     * @return The container name the object belongs to.
     * 
     */
    public Optional<Output<String>> container() {
        return Optional.ofNullable(this.container);
    }

    /**
     * The method allowed when accessing this URL.
     * Valid values are `GET`, and `POST`. Default is `GET`.
     * 
     */
    @Import(name="method")
    private @Nullable Output<String> method;

    /**
     * @return The method allowed when accessing this URL.
     * Valid values are `GET`, and `POST`. Default is `GET`.
     * 
     */
    public Optional<Output<String>> method() {
        return Optional.ofNullable(this.method);
    }

    /**
     * The object name the tempurl is for.
     * 
     */
    @Import(name="object")
    private @Nullable Output<String> object;

    /**
     * @return The object name the tempurl is for.
     * 
     */
    public Optional<Output<String>> object() {
        return Optional.ofNullable(this.object);
    }

    /**
     * Whether to automatically regenerate the URL when
     * it has expired. If set to true, this will create a new resource with a new
     * ID and new URL. Defaults to false.
     * 
     */
    @Import(name="regenerate")
    private @Nullable Output<Boolean> regenerate;

    /**
     * @return Whether to automatically regenerate the URL when
     * it has expired. If set to true, this will create a new resource with a new
     * ID and new URL. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> regenerate() {
        return Optional.ofNullable(this.regenerate);
    }

    /**
     * The region the tempurl is located in.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region the tempurl is located in.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    @Import(name="split")
    private @Nullable Output<String> split;

    public Optional<Output<String>> split() {
        return Optional.ofNullable(this.split);
    }

    /**
     * The TTL, in seconds, for the URL. For how long it should
     * be valid.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return The TTL, in seconds, for the URL. For how long it should
     * be valid.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * The URL
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The URL
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private TempUrlState() {}

    private TempUrlState(TempUrlState $) {
        this.container = $.container;
        this.method = $.method;
        this.object = $.object;
        this.regenerate = $.regenerate;
        this.region = $.region;
        this.split = $.split;
        this.ttl = $.ttl;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TempUrlState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TempUrlState $;

        public Builder() {
            $ = new TempUrlState();
        }

        public Builder(TempUrlState defaults) {
            $ = new TempUrlState(Objects.requireNonNull(defaults));
        }

        /**
         * @param container The container name the object belongs to.
         * 
         * @return builder
         * 
         */
        public Builder container(@Nullable Output<String> container) {
            $.container = container;
            return this;
        }

        /**
         * @param container The container name the object belongs to.
         * 
         * @return builder
         * 
         */
        public Builder container(String container) {
            return container(Output.of(container));
        }

        /**
         * @param method The method allowed when accessing this URL.
         * Valid values are `GET`, and `POST`. Default is `GET`.
         * 
         * @return builder
         * 
         */
        public Builder method(@Nullable Output<String> method) {
            $.method = method;
            return this;
        }

        /**
         * @param method The method allowed when accessing this URL.
         * Valid values are `GET`, and `POST`. Default is `GET`.
         * 
         * @return builder
         * 
         */
        public Builder method(String method) {
            return method(Output.of(method));
        }

        /**
         * @param object The object name the tempurl is for.
         * 
         * @return builder
         * 
         */
        public Builder object(@Nullable Output<String> object) {
            $.object = object;
            return this;
        }

        /**
         * @param object The object name the tempurl is for.
         * 
         * @return builder
         * 
         */
        public Builder object(String object) {
            return object(Output.of(object));
        }

        /**
         * @param regenerate Whether to automatically regenerate the URL when
         * it has expired. If set to true, this will create a new resource with a new
         * ID and new URL. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder regenerate(@Nullable Output<Boolean> regenerate) {
            $.regenerate = regenerate;
            return this;
        }

        /**
         * @param regenerate Whether to automatically regenerate the URL when
         * it has expired. If set to true, this will create a new resource with a new
         * ID and new URL. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder regenerate(Boolean regenerate) {
            return regenerate(Output.of(regenerate));
        }

        /**
         * @param region The region the tempurl is located in.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region the tempurl is located in.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public Builder split(@Nullable Output<String> split) {
            $.split = split;
            return this;
        }

        public Builder split(String split) {
            return split(Output.of(split));
        }

        /**
         * @param ttl The TTL, in seconds, for the URL. For how long it should
         * be valid.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The TTL, in seconds, for the URL. For how long it should
         * be valid.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param url The URL
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public TempUrlState build() {
            return $;
        }
    }

}
