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
public final class GetContainerAclRead {
    /**
     * @return The date the container ACL was created.
     * 
     */
    private String createdAt;
    /**
     * @return Whether the container is accessible project wide.
     * 
     */
    private @Nullable Boolean projectAccess;
    /**
     * @return The date the container ACL was last updated.
     * 
     */
    private String updatedAt;
    /**
     * @return The list of user IDs, which are allowed to access the container,
     * when `project_access` is set to `false`.
     * 
     */
    private @Nullable List<String> users;

    private GetContainerAclRead() {}
    /**
     * @return The date the container ACL was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Whether the container is accessible project wide.
     * 
     */
    public Optional<Boolean> projectAccess() {
        return Optional.ofNullable(this.projectAccess);
    }
    /**
     * @return The date the container ACL was last updated.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return The list of user IDs, which are allowed to access the container,
     * when `project_access` is set to `false`.
     * 
     */
    public List<String> users() {
        return this.users == null ? List.of() : this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetContainerAclRead defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private @Nullable Boolean projectAccess;
        private String updatedAt;
        private @Nullable List<String> users;
        public Builder() {}
        public Builder(GetContainerAclRead defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.projectAccess = defaults.projectAccess;
    	      this.updatedAt = defaults.updatedAt;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder projectAccess(@Nullable Boolean projectAccess) {
            this.projectAccess = projectAccess;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
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
        public GetContainerAclRead build() {
            final var o = new GetContainerAclRead();
            o.createdAt = createdAt;
            o.projectAccess = projectAccess;
            o.updatedAt = updatedAt;
            o.users = users;
            return o;
        }
    }
}
