// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerV1AclRead {
    /**
     * @return The date the container was created.
     * 
     */
    private @Nullable String createdAt;
    private @Nullable Boolean projectAccess;
    /**
     * @return The date the container was last updated.
     * 
     */
    private @Nullable String updatedAt;
    private @Nullable List<String> users;

    private ContainerV1AclRead() {}
    /**
     * @return The date the container was created.
     * 
     */
    public Optional<String> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }
    public Optional<Boolean> projectAccess() {
        return Optional.ofNullable(this.projectAccess);
    }
    /**
     * @return The date the container was last updated.
     * 
     */
    public Optional<String> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }
    public List<String> users() {
        return this.users == null ? List.of() : this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerV1AclRead defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String createdAt;
        private @Nullable Boolean projectAccess;
        private @Nullable String updatedAt;
        private @Nullable List<String> users;
        public Builder() {}
        public Builder(ContainerV1AclRead defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.projectAccess = defaults.projectAccess;
    	      this.updatedAt = defaults.updatedAt;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder createdAt(@Nullable String createdAt) {

            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder projectAccess(@Nullable Boolean projectAccess) {

            this.projectAccess = projectAccess;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(@Nullable String updatedAt) {

            this.updatedAt = updatedAt;
            return this;
        }
        @CustomType.Setter
        public Builder users(@Nullable List<String> users) {

            this.users = users;
            return this;
        }
        public Builder users(String... users) {
            return users(List.of(users));
        }
        public ContainerV1AclRead build() {
            final var _resultValue = new ContainerV1AclRead();
            _resultValue.createdAt = createdAt;
            _resultValue.projectAccess = projectAccess;
            _resultValue.updatedAt = updatedAt;
            _resultValue.users = users;
            return _resultValue;
        }
    }
}
