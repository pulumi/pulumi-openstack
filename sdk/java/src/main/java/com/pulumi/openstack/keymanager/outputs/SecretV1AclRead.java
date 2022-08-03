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
public final class SecretV1AclRead {
    /**
     * @return The date the secret ACL was created.
     * 
     */
    private final @Nullable String createdAt;
    /**
     * @return Whether the secret is accessible project wide.
     * Defaults to `true`.
     * 
     */
    private final @Nullable Boolean projectAccess;
    /**
     * @return The date the secret ACL was last updated.
     * 
     */
    private final @Nullable String updatedAt;
    /**
     * @return The list of user IDs, which are allowed to access the
     * secret, when `project_access` is set to `false`.
     * 
     */
    private final @Nullable List<String> users;

    @CustomType.Constructor
    private SecretV1AclRead(
        @CustomType.Parameter("createdAt") @Nullable String createdAt,
        @CustomType.Parameter("projectAccess") @Nullable Boolean projectAccess,
        @CustomType.Parameter("updatedAt") @Nullable String updatedAt,
        @CustomType.Parameter("users") @Nullable List<String> users) {
        this.createdAt = createdAt;
        this.projectAccess = projectAccess;
        this.updatedAt = updatedAt;
        this.users = users;
    }

    /**
     * @return The date the secret ACL was created.
     * 
     */
    public Optional<String> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }
    /**
     * @return Whether the secret is accessible project wide.
     * Defaults to `true`.
     * 
     */
    public Optional<Boolean> projectAccess() {
        return Optional.ofNullable(this.projectAccess);
    }
    /**
     * @return The date the secret ACL was last updated.
     * 
     */
    public Optional<String> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }
    /**
     * @return The list of user IDs, which are allowed to access the
     * secret, when `project_access` is set to `false`.
     * 
     */
    public List<String> users() {
        return this.users == null ? List.of() : this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretV1AclRead defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String createdAt;
        private @Nullable Boolean projectAccess;
        private @Nullable String updatedAt;
        private @Nullable List<String> users;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretV1AclRead defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.projectAccess = defaults.projectAccess;
    	      this.updatedAt = defaults.updatedAt;
    	      this.users = defaults.users;
        }

        public Builder createdAt(@Nullable String createdAt) {
            this.createdAt = createdAt;
            return this;
        }
        public Builder projectAccess(@Nullable Boolean projectAccess) {
            this.projectAccess = projectAccess;
            return this;
        }
        public Builder updatedAt(@Nullable String updatedAt) {
            this.updatedAt = updatedAt;
            return this;
        }
        public Builder users(@Nullable List<String> users) {
            this.users = users;
            return this;
        }
        public Builder users(String... users) {
            return users(List.of(users));
        }        public SecretV1AclRead build() {
            return new SecretV1AclRead(createdAt, projectAccess, updatedAt, users);
        }
    }
}
