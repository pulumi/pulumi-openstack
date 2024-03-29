// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceVolumeArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceVolumeArgs Empty = new InstanceVolumeArgs();

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="id")
    private @Nullable Output<String> id;

    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    @Import(name="volumeId", required=true)
    private Output<String> volumeId;

    public Output<String> volumeId() {
        return this.volumeId;
    }

    private InstanceVolumeArgs() {}

    private InstanceVolumeArgs(InstanceVolumeArgs $) {
        this.device = $.device;
        this.id = $.id;
        this.volumeId = $.volumeId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceVolumeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceVolumeArgs $;

        public Builder() {
            $ = new InstanceVolumeArgs();
        }

        public Builder(InstanceVolumeArgs defaults) {
            $ = new InstanceVolumeArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        public Builder id(String id) {
            return id(Output.of(id));
        }

        public Builder volumeId(Output<String> volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        public Builder volumeId(String volumeId) {
            return volumeId(Output.of(volumeId));
        }

        public InstanceVolumeArgs build() {
            if ($.volumeId == null) {
                throw new MissingRequiredPropertyException("InstanceVolumeArgs", "volumeId");
            }
            return $;
        }
    }

}
