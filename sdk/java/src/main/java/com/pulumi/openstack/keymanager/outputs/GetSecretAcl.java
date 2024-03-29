// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.keymanager.outputs.GetSecretAclRead;
import java.util.Objects;

@CustomType
public final class GetSecretAcl {
    private GetSecretAclRead read;

    private GetSecretAcl() {}
    public GetSecretAclRead read() {
        return this.read;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretAcl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private GetSecretAclRead read;
        public Builder() {}
        public Builder(GetSecretAcl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.read = defaults.read;
        }

        @CustomType.Setter
        public Builder read(GetSecretAclRead read) {
            if (read == null) {
              throw new MissingRequiredPropertyException("GetSecretAcl", "read");
            }
            this.read = read;
            return this;
        }
        public GetSecretAcl build() {
            final var _resultValue = new GetSecretAcl();
            _resultValue.read = read;
            return _resultValue;
        }
    }
}
