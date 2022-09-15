// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRoleResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String domainId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String name;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String region;

    @CustomType.Constructor
    private GetRoleResult(
        @CustomType.Parameter("domainId") String domainId,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("region") String region) {
        this.domainId = domainId;
        this.id = id;
        this.name = name;
        this.region = region;
    }

    /**
     * @return See Argument Reference above.
     * 
     */
    public String domainId() {
        return this.domainId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String region() {
        return this.region;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRoleResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String domainId;
        private String id;
        private String name;
        private String region;

        public Builder() {
    	      // Empty
        }

        public Builder(GetRoleResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domainId = defaults.domainId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.region = defaults.region;
        }

        public Builder domainId(String domainId) {
            this.domainId = Objects.requireNonNull(domainId);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }        public GetRoleResult build() {
            return new GetRoleResult(domainId, id, name, region);
        }
    }
}