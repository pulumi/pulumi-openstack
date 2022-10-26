// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAuthScopeServiceCatalogEndpoint {
    /**
     * @return The ID of the endpoint.
     * 
     */
    private String id;
    /**
     * @return The interface of the endpoint.
     * 
     */
    private String interface_;
    /**
     * @return The region in which to obtain the V3 Identity client.
     * A Identity client is needed to retrieve tokens IDs. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    private String region;
    /**
     * @return The region ID of the endpoint.
     * 
     */
    private String regionId;
    /**
     * @return The URL of the endpoint.
     * 
     */
    private String url;

    private GetAuthScopeServiceCatalogEndpoint() {}
    /**
     * @return The ID of the endpoint.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The interface of the endpoint.
     * 
     */
    public String interface_() {
        return this.interface_;
    }
    /**
     * @return The region in which to obtain the V3 Identity client.
     * A Identity client is needed to retrieve tokens IDs. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The region ID of the endpoint.
     * 
     */
    public String regionId() {
        return this.regionId;
    }
    /**
     * @return The URL of the endpoint.
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAuthScopeServiceCatalogEndpoint defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String interface_;
        private String region;
        private String regionId;
        private String url;
        public Builder() {}
        public Builder(GetAuthScopeServiceCatalogEndpoint defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.interface_ = defaults.interface_;
    	      this.region = defaults.region;
    	      this.regionId = defaults.regionId;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter("interface")
        public Builder interface_(String interface_) {
            this.interface_ = Objects.requireNonNull(interface_);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder regionId(String regionId) {
            this.regionId = Objects.requireNonNull(regionId);
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public GetAuthScopeServiceCatalogEndpoint build() {
            final var o = new GetAuthScopeServiceCatalogEndpoint();
            o.id = id;
            o.interface_ = interface_;
            o.region = region;
            o.regionId = regionId;
            o.url = url;
            return o;
        }
    }
}
