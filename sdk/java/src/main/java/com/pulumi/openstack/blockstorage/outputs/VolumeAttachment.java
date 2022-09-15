// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class VolumeAttachment {
    private final @Nullable String device;
    private final @Nullable String id;
    private final @Nullable String instanceId;

    @CustomType.Constructor
    private VolumeAttachment(
        @CustomType.Parameter("device") @Nullable String device,
        @CustomType.Parameter("id") @Nullable String id,
        @CustomType.Parameter("instanceId") @Nullable String instanceId) {
        this.device = device;
        this.id = id;
        this.instanceId = instanceId;
    }

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    public Optional<String> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(VolumeAttachment defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String device;
        private @Nullable String id;
        private @Nullable String instanceId;

        public Builder() {
    	      // Empty
        }

        public Builder(VolumeAttachment defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.instanceId = defaults.instanceId;
        }

        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        public Builder instanceId(@Nullable String instanceId) {
            this.instanceId = instanceId;
            return this;
        }        public VolumeAttachment build() {
            return new VolumeAttachment(device, id, instanceId);
        }
    }
}