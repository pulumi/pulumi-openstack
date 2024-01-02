// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerV1SecretRef {
    /**
     * @return The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
     * 
     */
    private @Nullable String name;
    /**
     * @return The secret reference / where to find the secret, URL.
     * 
     */
    private String secretRef;

    private ContainerV1SecretRef() {}
    /**
     * @return The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The secret reference / where to find the secret, URL.
     * 
     */
    public String secretRef() {
        return this.secretRef;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerV1SecretRef defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private String secretRef;
        public Builder() {}
        public Builder(ContainerV1SecretRef defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.secretRef = defaults.secretRef;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder secretRef(String secretRef) {
            if (secretRef == null) {
              throw new MissingRequiredPropertyException("ContainerV1SecretRef", "secretRef");
            }
            this.secretRef = secretRef;
            return this;
        }
        public ContainerV1SecretRef build() {
            final var _resultValue = new ContainerV1SecretRef();
            _resultValue.name = name;
            _resultValue.secretRef = secretRef;
            return _resultValue;
        }
    }
}
