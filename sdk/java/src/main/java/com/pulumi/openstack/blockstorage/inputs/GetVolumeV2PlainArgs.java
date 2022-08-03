// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVolumeV2PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVolumeV2PlainArgs Empty = new GetVolumeV2PlainArgs();

    /**
     * Indicates if the volume is bootable.
     * 
     */
    @Import(name="bootable")
    private @Nullable String bootable;

    /**
     * @return Indicates if the volume is bootable.
     * 
     */
    public Optional<String> bootable() {
        return Optional.ofNullable(this.bootable);
    }

    /**
     * Metadata key/value pairs associated with the volume.
     * 
     */
    @Import(name="metadata")
    private @Nullable Map<String,Object> metadata;

    /**
     * @return Metadata key/value pairs associated with the volume.
     * 
     */
    public Optional<Map<String,Object>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * The name of the volume.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the volume.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V2 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V2 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The status of the volume.
     * 
     */
    @Import(name="status")
    private @Nullable String status;

    /**
     * @return The status of the volume.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The type of the volume.
     * 
     */
    @Import(name="volumeType")
    private @Nullable String volumeType;

    /**
     * @return The type of the volume.
     * 
     */
    public Optional<String> volumeType() {
        return Optional.ofNullable(this.volumeType);
    }

    private GetVolumeV2PlainArgs() {}

    private GetVolumeV2PlainArgs(GetVolumeV2PlainArgs $) {
        this.bootable = $.bootable;
        this.metadata = $.metadata;
        this.name = $.name;
        this.region = $.region;
        this.status = $.status;
        this.volumeType = $.volumeType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVolumeV2PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVolumeV2PlainArgs $;

        public Builder() {
            $ = new GetVolumeV2PlainArgs();
        }

        public Builder(GetVolumeV2PlainArgs defaults) {
            $ = new GetVolumeV2PlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bootable Indicates if the volume is bootable.
         * 
         * @return builder
         * 
         */
        public Builder bootable(@Nullable String bootable) {
            $.bootable = bootable;
            return this;
        }

        /**
         * @param metadata Metadata key/value pairs associated with the volume.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Map<String,Object> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param name The name of the volume.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Block Storage
         * client. If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        /**
         * @param status The status of the volume.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable String status) {
            $.status = status;
            return this;
        }

        /**
         * @param volumeType The type of the volume.
         * 
         * @return builder
         * 
         */
        public Builder volumeType(@Nullable String volumeType) {
            $.volumeType = volumeType;
            return this;
        }

        public GetVolumeV2PlainArgs build() {
            return $;
        }
    }

}
