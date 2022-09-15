// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetContainerSecretRef {
    /**
     * @return The Container name.
     * 
     */
    private final @Nullable String name;
    /**
     * @return The secret reference / where to find the secret, URL.
     * 
     */
    private final @Nullable String secretRef;

    @CustomType.Constructor
    private GetContainerSecretRef(
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("secretRef") @Nullable String secretRef) {
        this.name = name;
        this.secretRef = secretRef;
    }

    /**
     * @return The Container name.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The secret reference / where to find the secret, URL.
     * 
     */
    public Optional<String> secretRef() {
        return Optional.ofNullable(this.secretRef);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetContainerSecretRef defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String name;
        private @Nullable String secretRef;

        public Builder() {
    	      // Empty
        }

        public Builder(GetContainerSecretRef defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.secretRef = defaults.secretRef;
        }

        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder secretRef(@Nullable String secretRef) {
            this.secretRef = secretRef;
            return this;
        }        public GetContainerSecretRef build() {
            return new GetContainerSecretRef(name, secretRef);
        }
    }
}