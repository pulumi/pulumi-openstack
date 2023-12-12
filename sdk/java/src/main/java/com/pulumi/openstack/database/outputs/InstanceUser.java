// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceUser {
    /**
     * @return A list of databases that user will have access to. If not specified,
     * user has access to all databases on th einstance. Changing this creates a new instance.
     * 
     */
    private @Nullable List<String> databases;
    /**
     * @return An ip address or % sign indicating what ip addresses can connect with
     * this user credentials. Changing this creates a new instance.
     * 
     */
    private @Nullable String host;
    /**
     * @return Username to be created on new instance. Changing this creates a
     * new instance.
     * 
     */
    private String name;
    /**
     * @return User&#39;s password. Changing this creates a
     * new instance.
     * 
     */
    private @Nullable String password;

    private InstanceUser() {}
    /**
     * @return A list of databases that user will have access to. If not specified,
     * user has access to all databases on th einstance. Changing this creates a new instance.
     * 
     */
    public List<String> databases() {
        return this.databases == null ? List.of() : this.databases;
    }
    /**
     * @return An ip address or % sign indicating what ip addresses can connect with
     * this user credentials. Changing this creates a new instance.
     * 
     */
    public Optional<String> host() {
        return Optional.ofNullable(this.host);
    }
    /**
     * @return Username to be created on new instance. Changing this creates a
     * new instance.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return User&#39;s password. Changing this creates a
     * new instance.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceUser defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> databases;
        private @Nullable String host;
        private String name;
        private @Nullable String password;
        public Builder() {}
        public Builder(InstanceUser defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.databases = defaults.databases;
    	      this.host = defaults.host;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
        }

        @CustomType.Setter
        public Builder databases(@Nullable List<String> databases) {
            this.databases = databases;
            return this;
        }
        public Builder databases(String... databases) {
            return databases(List.of(databases));
        }
        @CustomType.Setter
        public Builder host(@Nullable String host) {
            this.host = host;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        public InstanceUser build() {
            final var _resultValue = new InstanceUser();
            _resultValue.databases = databases;
            _resultValue.host = host;
            _resultValue.name = name;
            _resultValue.password = password;
            return _resultValue;
        }
    }
}
