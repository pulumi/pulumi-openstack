// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.identity.outputs.GetAuthScopeServiceCatalogEndpoint;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAuthScopeServiceCatalog {
    /**
     * @return A list of endpoints for the service.
     * 
     */
    private List<GetAuthScopeServiceCatalogEndpoint> endpoints;
    /**
     * @return The ID of the endpoint.
     * 
     */
    private String id;
    /**
     * @return The name of the scope. This is an arbitrary name which is
     * only used as a unique identifier so an actual token isn&#39;t used as the ID.
     * 
     */
    private String name;
    /**
     * @return The type of the service.
     * 
     */
    private String type;

    private GetAuthScopeServiceCatalog() {}
    /**
     * @return A list of endpoints for the service.
     * 
     */
    public List<GetAuthScopeServiceCatalogEndpoint> endpoints() {
        return this.endpoints;
    }
    /**
     * @return The ID of the endpoint.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the scope. This is an arbitrary name which is
     * only used as a unique identifier so an actual token isn&#39;t used as the ID.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The type of the service.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAuthScopeServiceCatalog defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetAuthScopeServiceCatalogEndpoint> endpoints;
        private String id;
        private String name;
        private String type;
        public Builder() {}
        public Builder(GetAuthScopeServiceCatalog defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endpoints = defaults.endpoints;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder endpoints(List<GetAuthScopeServiceCatalogEndpoint> endpoints) {
            if (endpoints == null) {
              throw new MissingRequiredPropertyException("GetAuthScopeServiceCatalog", "endpoints");
            }
            this.endpoints = endpoints;
            return this;
        }
        public Builder endpoints(GetAuthScopeServiceCatalogEndpoint... endpoints) {
            return endpoints(List.of(endpoints));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetAuthScopeServiceCatalog", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetAuthScopeServiceCatalog", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetAuthScopeServiceCatalog", "type");
            }
            this.type = type;
            return this;
        }
        public GetAuthScopeServiceCatalog build() {
            final var _resultValue = new GetAuthScopeServiceCatalog();
            _resultValue.endpoints = endpoints;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
