// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.blockstorage.inputs.VolumeV2SchedulerHintArgs;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeV2Args extends com.pulumi.resources.ResourceArgs {

    public static final VolumeV2Args Empty = new VolumeV2Args();

    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return The availability zone for the volume.
     * Changing this creates a new volume.
     * 
     */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * The consistency group to place the volume
     * in.
     * 
     */
    @Import(name="consistencyGroupId")
    private @Nullable Output<String> consistencyGroupId;

    /**
     * @return The consistency group to place the volume
     * in.
     * 
     */
    public Optional<Output<String>> consistencyGroupId() {
        return Optional.ofNullable(this.consistencyGroupId);
    }

    /**
     * A description of the volume. Changing this updates
     * the volume&#39;s description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the volume. Changing this updates
     * the volume&#39;s description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The image ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    @Import(name="imageId")
    private @Nullable Output<String> imageId;

    /**
     * @return The image ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    public Optional<Output<String>> imageId() {
        return Optional.ofNullable(this.imageId);
    }

    /**
     * Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,Object>> metadata;

    /**
     * @return Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     * 
     */
    public Optional<Output<Map<String,Object>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * A unique name for the volume. Changing this updates the
     * volume&#39;s name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name for the volume. Changing this updates the
     * volume&#39;s name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Provide the Cinder scheduler with hints on where
     * to instantiate a volume in the OpenStack cloud. The available hints are described below.
     * 
     */
    @Import(name="schedulerHints")
    private @Nullable Output<List<VolumeV2SchedulerHintArgs>> schedulerHints;

    /**
     * @return Provide the Cinder scheduler with hints on where
     * to instantiate a volume in the OpenStack cloud. The available hints are described below.
     * 
     */
    public Optional<Output<List<VolumeV2SchedulerHintArgs>>> schedulerHints() {
        return Optional.ofNullable(this.schedulerHints);
    }

    /**
     * The size of the volume to create (in gigabytes). Changing
     * this creates a new volume.
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return The size of the volume to create (in gigabytes). Changing
     * this creates a new volume.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    /**
     * The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    /**
     * The volume ID to replicate with.
     * 
     */
    @Import(name="sourceReplica")
    private @Nullable Output<String> sourceReplica;

    /**
     * @return The volume ID to replicate with.
     * 
     */
    public Optional<Output<String>> sourceReplica() {
        return Optional.ofNullable(this.sourceReplica);
    }

    /**
     * The volume ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    @Import(name="sourceVolId")
    private @Nullable Output<String> sourceVolId;

    /**
     * @return The volume ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    public Optional<Output<String>> sourceVolId() {
        return Optional.ofNullable(this.sourceVolId);
    }

    /**
     * The type of volume to create.
     * Changing this creates a new volume.
     * 
     */
    @Import(name="volumeType")
    private @Nullable Output<String> volumeType;

    /**
     * @return The type of volume to create.
     * Changing this creates a new volume.
     * 
     */
    public Optional<Output<String>> volumeType() {
        return Optional.ofNullable(this.volumeType);
    }

    private VolumeV2Args() {}

    private VolumeV2Args(VolumeV2Args $) {
        this.availabilityZone = $.availabilityZone;
        this.consistencyGroupId = $.consistencyGroupId;
        this.description = $.description;
        this.imageId = $.imageId;
        this.metadata = $.metadata;
        this.name = $.name;
        this.region = $.region;
        this.schedulerHints = $.schedulerHints;
        this.size = $.size;
        this.snapshotId = $.snapshotId;
        this.sourceReplica = $.sourceReplica;
        this.sourceVolId = $.sourceVolId;
        this.volumeType = $.volumeType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeV2Args $;

        public Builder() {
            $ = new VolumeV2Args();
        }

        public Builder(VolumeV2Args defaults) {
            $ = new VolumeV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param availabilityZone The availability zone for the volume.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The availability zone for the volume.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param consistencyGroupId The consistency group to place the volume
         * in.
         * 
         * @return builder
         * 
         */
        public Builder consistencyGroupId(@Nullable Output<String> consistencyGroupId) {
            $.consistencyGroupId = consistencyGroupId;
            return this;
        }

        /**
         * @param consistencyGroupId The consistency group to place the volume
         * in.
         * 
         * @return builder
         * 
         */
        public Builder consistencyGroupId(String consistencyGroupId) {
            return consistencyGroupId(Output.of(consistencyGroupId));
        }

        /**
         * @param description A description of the volume. Changing this updates
         * the volume&#39;s description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the volume. Changing this updates
         * the volume&#39;s description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param imageId The image ID from which to create the volume.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder imageId(@Nullable Output<String> imageId) {
            $.imageId = imageId;
            return this;
        }

        /**
         * @param imageId The image ID from which to create the volume.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder imageId(String imageId) {
            return imageId(Output.of(imageId));
        }

        /**
         * @param metadata Metadata key/value pairs to associate with the volume.
         * Changing this updates the existing volume metadata.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,Object>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Metadata key/value pairs to associate with the volume.
         * Changing this updates the existing volume metadata.
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,Object> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param name A unique name for the volume. Changing this updates the
         * volume&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name for the volume. Changing this updates the
         * volume&#39;s name.
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
         * creates a new volume.
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
         * creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param schedulerHints Provide the Cinder scheduler with hints on where
         * to instantiate a volume in the OpenStack cloud. The available hints are described below.
         * 
         * @return builder
         * 
         */
        public Builder schedulerHints(@Nullable Output<List<VolumeV2SchedulerHintArgs>> schedulerHints) {
            $.schedulerHints = schedulerHints;
            return this;
        }

        /**
         * @param schedulerHints Provide the Cinder scheduler with hints on where
         * to instantiate a volume in the OpenStack cloud. The available hints are described below.
         * 
         * @return builder
         * 
         */
        public Builder schedulerHints(List<VolumeV2SchedulerHintArgs> schedulerHints) {
            return schedulerHints(Output.of(schedulerHints));
        }

        /**
         * @param schedulerHints Provide the Cinder scheduler with hints on where
         * to instantiate a volume in the OpenStack cloud. The available hints are described below.
         * 
         * @return builder
         * 
         */
        public Builder schedulerHints(VolumeV2SchedulerHintArgs... schedulerHints) {
            return schedulerHints(List.of(schedulerHints));
        }

        /**
         * @param size The size of the volume to create (in gigabytes). Changing
         * this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The size of the volume to create (in gigabytes). Changing
         * this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param snapshotId The snapshot ID from which to create the volume.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId The snapshot ID from which to create the volume.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
        }

        /**
         * @param sourceReplica The volume ID to replicate with.
         * 
         * @return builder
         * 
         */
        public Builder sourceReplica(@Nullable Output<String> sourceReplica) {
            $.sourceReplica = sourceReplica;
            return this;
        }

        /**
         * @param sourceReplica The volume ID to replicate with.
         * 
         * @return builder
         * 
         */
        public Builder sourceReplica(String sourceReplica) {
            return sourceReplica(Output.of(sourceReplica));
        }

        /**
         * @param sourceVolId The volume ID from which to create the volume.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder sourceVolId(@Nullable Output<String> sourceVolId) {
            $.sourceVolId = sourceVolId;
            return this;
        }

        /**
         * @param sourceVolId The volume ID from which to create the volume.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder sourceVolId(String sourceVolId) {
            return sourceVolId(Output.of(sourceVolId));
        }

        /**
         * @param volumeType The type of volume to create.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder volumeType(@Nullable Output<String> volumeType) {
            $.volumeType = volumeType;
            return this;
        }

        /**
         * @param volumeType The type of volume to create.
         * Changing this creates a new volume.
         * 
         * @return builder
         * 
         */
        public Builder volumeType(String volumeType) {
            return volumeType(Output.of(volumeType));
        }

        public VolumeV2Args build() {
            $.size = Objects.requireNonNull($.size, "expected parameter 'size' to be non-null");
            return $;
        }
    }

}