// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationCredentialAccessRule {
    /**
     * @return The ID of the existing access rule. The access rule ID of
     * another application credential can be provided.
     * 
     */
    private final @Nullable String id;
    /**
     * @return The request method that the application credential is
     * permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
     * `HEAD`, `PATCH`, `PUT` and `DELETE`.
     * 
     */
    private final String method;
    /**
     * @return The API path that the application credential is permitted
     * to access. May use named wildcards such as **{tag}** or the unnamed wildcard
     * **\*** to match against any string in the path up to a **{@literal /}**, or the recursive
     * wildcard **\*\*** to include **{@literal /}** in the matched path.
     * 
     */
    private final String path;
    /**
     * @return The service type identifier for the service that the
     * application credential is granted to access. Must be a service type that is
     * listed in the service catalog and not a code name for a service. E.g.
     * **identity**, **compute**, **volumev3**, **image**, **network**,
     * **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
     * 
     */
    private final String service;

    @CustomType.Constructor
    private ApplicationCredentialAccessRule(
        @CustomType.Parameter("id") @Nullable String id,
        @CustomType.Parameter("method") String method,
        @CustomType.Parameter("path") String path,
        @CustomType.Parameter("service") String service) {
        this.id = id;
        this.method = method;
        this.path = path;
        this.service = service;
    }

    /**
     * @return The ID of the existing access rule. The access rule ID of
     * another application credential can be provided.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The request method that the application credential is
     * permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
     * `HEAD`, `PATCH`, `PUT` and `DELETE`.
     * 
     */
    public String method() {
        return this.method;
    }
    /**
     * @return The API path that the application credential is permitted
     * to access. May use named wildcards such as **{tag}** or the unnamed wildcard
     * **\*** to match against any string in the path up to a **{@literal /}**, or the recursive
     * wildcard **\*\*** to include **{@literal /}** in the matched path.
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return The service type identifier for the service that the
     * application credential is granted to access. Must be a service type that is
     * listed in the service catalog and not a code name for a service. E.g.
     * **identity**, **compute**, **volumev3**, **image**, **network**,
     * **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
     * 
     */
    public String service() {
        return this.service;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationCredentialAccessRule defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String id;
        private String method;
        private String path;
        private String service;

        public Builder() {
    	      // Empty
        }

        public Builder(ApplicationCredentialAccessRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.method = defaults.method;
    	      this.path = defaults.path;
    	      this.service = defaults.service;
        }

        public Builder id(@Nullable String id) {
            this.id = id;
            return this;
        }
        public Builder method(String method) {
            this.method = Objects.requireNonNull(method);
            return this;
        }
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public Builder service(String service) {
            this.service = Objects.requireNonNull(service);
            return this;
        }        public ApplicationCredentialAccessRule build() {
            return new ApplicationCredentialAccessRule(id, method, path, service);
        }
    }
}