// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAuthScopeRole {
    /**
     * @return The ID of the role.
     * 
     */
    private String roleId;
    /**
     * @return The name of the role.
     * 
     */
    private String roleName;

    private GetAuthScopeRole() {}
    /**
     * @return The ID of the role.
     * 
     */
    public String roleId() {
        return this.roleId;
    }
    /**
     * @return The name of the role.
     * 
     */
    public String roleName() {
        return this.roleName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAuthScopeRole defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String roleId;
        private String roleName;
        public Builder() {}
        public Builder(GetAuthScopeRole defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.roleId = defaults.roleId;
    	      this.roleName = defaults.roleName;
        }

        @CustomType.Setter
        public Builder roleId(String roleId) {
            this.roleId = Objects.requireNonNull(roleId);
            return this;
        }
        @CustomType.Setter
        public Builder roleName(String roleName) {
            this.roleName = Objects.requireNonNull(roleName);
            return this;
        }
        public GetAuthScopeRole build() {
            final var _resultValue = new GetAuthScopeRole();
            _resultValue.roleId = roleId;
            _resultValue.roleName = roleName;
            return _resultValue;
        }
    }
}
