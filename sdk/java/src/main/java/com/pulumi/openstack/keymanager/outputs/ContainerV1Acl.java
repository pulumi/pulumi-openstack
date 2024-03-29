// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.openstack.keymanager.outputs.ContainerV1AclRead;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerV1Acl {
    private @Nullable ContainerV1AclRead read;

    private ContainerV1Acl() {}
    public Optional<ContainerV1AclRead> read() {
        return Optional.ofNullable(this.read);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerV1Acl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable ContainerV1AclRead read;
        public Builder() {}
        public Builder(ContainerV1Acl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.read = defaults.read;
        }

        @CustomType.Setter
        public Builder read(@Nullable ContainerV1AclRead read) {

            this.read = read;
            return this;
        }
        public ContainerV1Acl build() {
            final var _resultValue = new ContainerV1Acl();
            _resultValue.read = read;
            return _resultValue;
        }
    }
}
