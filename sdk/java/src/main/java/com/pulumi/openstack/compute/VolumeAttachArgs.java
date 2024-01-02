// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.compute.inputs.VolumeAttachVendorOptionsArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeAttachArgs extends com.pulumi.resources.ResourceArgs {

    public static final VolumeAttachArgs Empty = new VolumeAttachArgs();

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * The ID of the Instance to attach the Volume to.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return The ID of the Instance to attach the Volume to.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
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
    @Import(name="volumeId", required=true)
    private Output<String> volumeId;

    /**
     * @return The ID of the Volume to attach to an Instance.
     * 
     */
    public Output<String> volumeId() {
        return this.volumeId;
    }

    private VolumeAttachArgs() {}

    private VolumeAttachArgs(VolumeAttachArgs $) {
        this.device = $.device;
        this.instanceId = $.instanceId;
        this.multiattach = $.multiattach;
        this.region = $.region;
        this.vendorOptions = $.vendorOptions;
        this.volumeId = $.volumeId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeAttachArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeAttachArgs $;

        public Builder() {
            $ = new VolumeAttachArgs();
        }

        public Builder(VolumeAttachArgs defaults) {
            $ = new VolumeAttachArgs(Objects.requireNonNull(defaults));
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
        public Builder instanceId(Output<String> instanceId) {
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
        public Builder volumeId(Output<String> volumeId) {
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

        public VolumeAttachArgs build() {
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("VolumeAttachArgs", "instanceId");
            }
            if ($.volumeId == null) {
                throw new MissingRequiredPropertyException("VolumeAttachArgs", "volumeId");
            }
            return $;
        }
    }

}
