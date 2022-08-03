// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.openstack.keymanager.outputs.SecretV1AclRead;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretV1Acl {
    private final @Nullable SecretV1AclRead read;

    @CustomType.Constructor
    private SecretV1Acl(@CustomType.Parameter("read") @Nullable SecretV1AclRead read) {
        this.read = read;
    }

    public Optional<SecretV1AclRead> read() {
        return Optional.ofNullable(this.read);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretV1Acl defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable SecretV1AclRead read;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretV1Acl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.read = defaults.read;
        }

        public Builder read(@Nullable SecretV1AclRead read) {
            this.read = read;
            return this;
        }        public SecretV1Acl build() {
            return new SecretV1Acl(read);
        }
    }
}
