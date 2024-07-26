// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSnapshotV3Args extends com.pulumi.resources.InvokeArgs {

    public static final GetSnapshotV3Args Empty = new GetSnapshotV3Args();

    /**
     * Pick the most recently created snapshot if there
     * are multiple results.
     * 
     */
    @Import(name="mostRecent")
    private @Nullable Output<Boolean> mostRecent;

    /**
     * @return Pick the most recently created snapshot if there
     * are multiple results.
     * 
     */
    public Optional<Output<Boolean>> mostRecent() {
        return Optional.ofNullable(this.mostRecent);
    }

    /**
     * The name of the snapshot.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the snapshot.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V3 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V3 Block Storage
     * client. If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The status of the snapshot.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The status of the snapshot.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The ID of the snapshot&#39;s volume.
     * 
     */
    @Import(name="volumeId")
    private @Nullable Output<String> volumeId;

    /**
     * @return The ID of the snapshot&#39;s volume.
     * 
     */
    public Optional<Output<String>> volumeId() {
        return Optional.ofNullable(this.volumeId);
    }

    private GetSnapshotV3Args() {}

    private GetSnapshotV3Args(GetSnapshotV3Args $) {
        this.mostRecent = $.mostRecent;
        this.name = $.name;
        this.region = $.region;
        this.status = $.status;
        this.volumeId = $.volumeId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSnapshotV3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSnapshotV3Args $;

        public Builder() {
            $ = new GetSnapshotV3Args();
        }

        public Builder(GetSnapshotV3Args defaults) {
            $ = new GetSnapshotV3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param mostRecent Pick the most recently created snapshot if there
         * are multiple results.
         * 
         * @return builder
         * 
         */
        public Builder mostRecent(@Nullable Output<Boolean> mostRecent) {
            $.mostRecent = mostRecent;
            return this;
        }

        /**
         * @param mostRecent Pick the most recently created snapshot if there
         * are multiple results.
         * 
         * @return builder
         * 
         */
        public Builder mostRecent(Boolean mostRecent) {
            return mostRecent(Output.of(mostRecent));
        }

        /**
         * @param name The name of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V3 Block Storage
         * client. If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V3 Block Storage
         * client. If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param status The status of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The status of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param volumeId The ID of the snapshot&#39;s volume.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(@Nullable Output<String> volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        /**
         * @param volumeId The ID of the snapshot&#39;s volume.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(String volumeId) {
            return volumeId(Output.of(volumeId));
        }

        public GetSnapshotV3Args build() {
            return $;
        }
    }

}