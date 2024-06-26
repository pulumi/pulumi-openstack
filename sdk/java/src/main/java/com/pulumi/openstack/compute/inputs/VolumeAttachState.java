// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.compute.inputs.VolumeAttachVendorOptionsArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeAttachState extends com.pulumi.resources.ResourceArgs {

    public static final VolumeAttachState Empty = new VolumeAttachState();

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * The ID of the Instance to attach the Volume to.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The ID of the Instance to attach the Volume to.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Enable attachment of multiattach-capable volumes.
     * 
     */
    @Import(name="multiattach")
    private @Nullable Output<Boolean> multiattach;

    /**
     * @return Enable attachment of multiattach-capable volumes.
     * 
     */
    public Optional<Output<Boolean>> multiattach() {
        return Optional.ofNullable(this.multiattach);
    }

    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Add a device role tag that is applied to the volume when
     * attaching it to the VM. Changing this creates a new volume attachment with
     * the new tag. Requires microversion &gt;= 2.49.
     * 
     */
    @Import(name="tag")
    private @Nullable Output<String> tag;

    /**
     * @return Add a device role tag that is applied to the volume when
     * attaching it to the VM. Changing this creates a new volume attachment with
     * the new tag. Requires microversion &gt;= 2.49.
     * 
     */
    public Optional<Output<String>> tag() {
        return Optional.ofNullable(this.tag);
    }

    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     * 
     */
    @Import(name="vendorOptions")
    private @Nullable Output<VolumeAttachVendorOptionsArgs> vendorOptions;

    /**
     * @return Map of additional vendor-specific options.
     * Supported options are described below.
     * 
     */
    public Optional<Output<VolumeAttachVendorOptionsArgs>> vendorOptions() {
        return Optional.ofNullable(this.vendorOptions);
    }

    /**
     * The ID of the Volume to attach to an Instance.
     * 
     */
    @Import(name="volumeId")
    private @Nullable Output<String> volumeId;

    /**
     * @return The ID of the Volume to attach to an Instance.
     * 
     */
    public Optional<Output<String>> volumeId() {
        return Optional.ofNullable(this.volumeId);
    }

    private VolumeAttachState() {}

    private VolumeAttachState(VolumeAttachState $) {
        this.device = $.device;
        this.instanceId = $.instanceId;
        this.multiattach = $.multiattach;
        this.region = $.region;
        this.tag = $.tag;
        this.vendorOptions = $.vendorOptions;
        this.volumeId = $.volumeId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeAttachState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeAttachState $;

        public Builder() {
            $ = new VolumeAttachState();
        }

        public Builder(VolumeAttachState defaults) {
            $ = new VolumeAttachState(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        /**
         * @param instanceId The ID of the Instance to attach the Volume to.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The ID of the Instance to attach the Volume to.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param multiattach Enable attachment of multiattach-capable volumes.
         * 
         * @return builder
         * 
         */
        public Builder multiattach(@Nullable Output<Boolean> multiattach) {
            $.multiattach = multiattach;
            return this;
        }

        /**
         * @param multiattach Enable attachment of multiattach-capable volumes.
         * 
         * @return builder
         * 
         */
        public Builder multiattach(Boolean multiattach) {
            return multiattach(Output.of(multiattach));
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * A Compute client is needed to create a volume attachment. If omitted, the
         * `region` argument of the provider is used. Changing this creates a
         * new volume attachment.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * A Compute client is needed to create a volume attachment. If omitted, the
         * `region` argument of the provider is used. Changing this creates a
         * new volume attachment.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tag Add a device role tag that is applied to the volume when
         * attaching it to the VM. Changing this creates a new volume attachment with
         * the new tag. Requires microversion &gt;= 2.49.
         * 
         * @return builder
         * 
         */
        public Builder tag(@Nullable Output<String> tag) {
            $.tag = tag;
            return this;
        }

        /**
         * @param tag Add a device role tag that is applied to the volume when
         * attaching it to the VM. Changing this creates a new volume attachment with
         * the new tag. Requires microversion &gt;= 2.49.
         * 
         * @return builder
         * 
         */
        public Builder tag(String tag) {
            return tag(Output.of(tag));
        }

        /**
         * @param vendorOptions Map of additional vendor-specific options.
         * Supported options are described below.
         * 
         * @return builder
         * 
         */
        public Builder vendorOptions(@Nullable Output<VolumeAttachVendorOptionsArgs> vendorOptions) {
            $.vendorOptions = vendorOptions;
            return this;
        }

        /**
         * @param vendorOptions Map of additional vendor-specific options.
         * Supported options are described below.
         * 
         * @return builder
         * 
         */
        public Builder vendorOptions(VolumeAttachVendorOptionsArgs vendorOptions) {
            return vendorOptions(Output.of(vendorOptions));
        }

        /**
         * @param volumeId The ID of the Volume to attach to an Instance.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(@Nullable Output<String> volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        /**
         * @param volumeId The ID of the Volume to attach to an Instance.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(String volumeId) {
            return volumeId(Output.of(volumeId));
        }

        public VolumeAttachState build() {
            return $;
        }
    }

}
