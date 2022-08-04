// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeAttachVendorOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final VolumeAttachVendorOptionsArgs Empty = new VolumeAttachVendorOptionsArgs();

    /**
     * Boolean to control whether
     * to ignore volume status confirmation of the attached volume. This can be helpful
     * to work with some OpenStack clouds which don&#39;t have the Block Storage V3 API available.
     * 
     */
    @Import(name="ignoreVolumeConfirmation")
    private @Nullable Output<Boolean> ignoreVolumeConfirmation;

    /**
     * @return Boolean to control whether
     * to ignore volume status confirmation of the attached volume. This can be helpful
     * to work with some OpenStack clouds which don&#39;t have the Block Storage V3 API available.
     * 
     */
    public Optional<Output<Boolean>> ignoreVolumeConfirmation() {
        return Optional.ofNullable(this.ignoreVolumeConfirmation);
    }

    private VolumeAttachVendorOptionsArgs() {}

    private VolumeAttachVendorOptionsArgs(VolumeAttachVendorOptionsArgs $) {
        this.ignoreVolumeConfirmation = $.ignoreVolumeConfirmation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeAttachVendorOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeAttachVendorOptionsArgs $;

        public Builder() {
            $ = new VolumeAttachVendorOptionsArgs();
        }

        public Builder(VolumeAttachVendorOptionsArgs defaults) {
            $ = new VolumeAttachVendorOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ignoreVolumeConfirmation Boolean to control whether
         * to ignore volume status confirmation of the attached volume. This can be helpful
         * to work with some OpenStack clouds which don&#39;t have the Block Storage V3 API available.
         * 
         * @return builder
         * 
         */
        public Builder ignoreVolumeConfirmation(@Nullable Output<Boolean> ignoreVolumeConfirmation) {
            $.ignoreVolumeConfirmation = ignoreVolumeConfirmation;
            return this;
        }

        /**
         * @param ignoreVolumeConfirmation Boolean to control whether
         * to ignore volume status confirmation of the attached volume. This can be helpful
         * to work with some OpenStack clouds which don&#39;t have the Block Storage V3 API available.
         * 
         * @return builder
         * 
         */
        public Builder ignoreVolumeConfirmation(Boolean ignoreVolumeConfirmation) {
            return ignoreVolumeConfirmation(Output.of(ignoreVolumeConfirmation));
        }

        public VolumeAttachVendorOptionsArgs build() {
            return $;
        }
    }

}
