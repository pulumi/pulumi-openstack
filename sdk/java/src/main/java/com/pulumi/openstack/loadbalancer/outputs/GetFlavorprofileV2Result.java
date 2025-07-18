// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetFlavorprofileV2Result {
    /**
     * @return Extra data of the flavorprofile depending on the provider.
     * 
     */
    private String flavorData;
    private String flavorprofileId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the flavorprofile.
     * 
     */
    private String name;
    /**
     * @return The name of the provider that the flavorprofile uses.
     * 
     */
    private String providerName;
    private String region;

    private GetFlavorprofileV2Result() {}
    /**
     * @return Extra data of the flavorprofile depending on the provider.
     * 
     */
    public String flavorData() {
        return this.flavorData;
    }
    public String flavorprofileId() {
        return this.flavorprofileId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the flavorprofile.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The name of the provider that the flavorprofile uses.
     * 
     */
    public String providerName() {
        return this.providerName;
    }
    public String region() {
        return this.region;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFlavorprofileV2Result defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String flavorData;
        private String flavorprofileId;
        private String id;
        private String name;
        private String providerName;
        private String region;
        public Builder() {}
        public Builder(GetFlavorprofileV2Result defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.flavorData = defaults.flavorData;
    	      this.flavorprofileId = defaults.flavorprofileId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.providerName = defaults.providerName;
    	      this.region = defaults.region;
        }

        @CustomType.Setter
        public Builder flavorData(String flavorData) {
            if (flavorData == null) {
              throw new MissingRequiredPropertyException("GetFlavorprofileV2Result", "flavorData");
            }
            this.flavorData = flavorData;
            return this;
        }
        @CustomType.Setter
        public Builder flavorprofileId(String flavorprofileId) {
            if (flavorprofileId == null) {
              throw new MissingRequiredPropertyException("GetFlavorprofileV2Result", "flavorprofileId");
            }
            this.flavorprofileId = flavorprofileId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetFlavorprofileV2Result", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetFlavorprofileV2Result", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder providerName(String providerName) {
            if (providerName == null) {
              throw new MissingRequiredPropertyException("GetFlavorprofileV2Result", "providerName");
            }
            this.providerName = providerName;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetFlavorprofileV2Result", "region");
            }
            this.region = region;
            return this;
        }
        public GetFlavorprofileV2Result build() {
            final var _resultValue = new GetFlavorprofileV2Result();
            _resultValue.flavorData = flavorData;
            _resultValue.flavorprofileId = flavorprofileId;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.providerName = providerName;
            _resultValue.region = region;
            return _resultValue;
        }
    }
}
