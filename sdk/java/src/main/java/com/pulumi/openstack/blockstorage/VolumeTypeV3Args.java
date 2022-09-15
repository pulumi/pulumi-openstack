// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeTypeV3Args extends com.pulumi.resources.ResourceArgs {

    public static final VolumeTypeV3Args Empty = new VolumeTypeV3Args();

    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing volume type.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description of the port. Changing
     * this updates the `description` of an existing volume type.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Key/Value pairs of metadata for the volume type.
     * 
     */
    @Import(name="extraSpecs")
    private @Nullable Output<Map<String,Object>> extraSpecs;

    /**
     * @return Key/Value pairs of metadata for the volume type.
     * 
     */
    public Optional<Output<Map<String,Object>>> extraSpecs() {
        return Optional.ofNullable(this.extraSpecs);
    }

    /**
     * Whether the volume type is public. Changing
     * this updates the `is_public` of an existing volume type.
     * 
     */
    @Import(name="isPublic")
    private @Nullable Output<Boolean> isPublic;

    /**
     * @return Whether the volume type is public. Changing
     * this updates the `is_public` of an existing volume type.
     * 
     */
    public Optional<Output<Boolean>> isPublic() {
        return Optional.ofNullable(this.isPublic);
    }

    /**
     * Name of the volume type.  Changing this
     * updates the `name` of an existing volume type.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the volume type.  Changing this
     * updates the `name` of an existing volume type.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private VolumeTypeV3Args() {}

    private VolumeTypeV3Args(VolumeTypeV3Args $) {
        this.description = $.description;
        this.extraSpecs = $.extraSpecs;
        this.isPublic = $.isPublic;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeTypeV3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeTypeV3Args $;

        public Builder() {
            $ = new VolumeTypeV3Args();
        }

        public Builder(VolumeTypeV3Args defaults) {
            $ = new VolumeTypeV3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Human-readable description of the port. Changing
         * this updates the `description` of an existing volume type.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description of the port. Changing
         * this updates the `description` of an existing volume type.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param extraSpecs Key/Value pairs of metadata for the volume type.
         * 
         * @return builder
         * 
         */
        public Builder extraSpecs(@Nullable Output<Map<String,Object>> extraSpecs) {
            $.extraSpecs = extraSpecs;
            return this;
        }

        /**
         * @param extraSpecs Key/Value pairs of metadata for the volume type.
         * 
         * @return builder
         * 
         */
        public Builder extraSpecs(Map<String,Object> extraSpecs) {
            return extraSpecs(Output.of(extraSpecs));
        }

        /**
         * @param isPublic Whether the volume type is public. Changing
         * this updates the `is_public` of an existing volume type.
         * 
         * @return builder
         * 
         */
        public Builder isPublic(@Nullable Output<Boolean> isPublic) {
            $.isPublic = isPublic;
            return this;
        }

        /**
         * @param isPublic Whether the volume type is public. Changing
         * this updates the `is_public` of an existing volume type.
         * 
         * @return builder
         * 
         */
        public Builder isPublic(Boolean isPublic) {
            return isPublic(Output.of(isPublic));
        }

        /**
         * @param name Name of the volume type.  Changing this
         * updates the `name` of an existing volume type.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the volume type.  Changing this
         * updates the `name` of an existing volume type.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to create the volume. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to create the volume. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public VolumeTypeV3Args build() {
            return $;
        }
    }

}