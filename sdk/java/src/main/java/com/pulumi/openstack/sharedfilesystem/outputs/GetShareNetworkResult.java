// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetShareNetworkResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String cidr;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private Integer ipVersion;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String name;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String networkType;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String neutronNetId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String neutronSubnetId;
    /**
     * @return The owner of the Share Network.
     * 
     */
    private String projectId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String region;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String securityServiceId;
    /**
     * @return The list of security service IDs associated with
     * the share network.
     * 
     */
    private List<String> securityServiceIds;
    /**
     * @return See Argument Reference above.
     * 
     */
    private Integer segmentationId;

    private GetShareNetworkResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String cidr() {
        return this.cidr;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String description() {
        return this.description;
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
    public Integer ipVersion() {
        return this.ipVersion;
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
    public String networkType() {
        return this.networkType;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String neutronNetId() {
        return this.neutronNetId;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String neutronSubnetId() {
        return this.neutronSubnetId;
    }
    /**
     * @return The owner of the Share Network.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> securityServiceId() {
        return Optional.ofNullable(this.securityServiceId);
    }
    /**
     * @return The list of security service IDs associated with
     * the share network.
     * 
     */
    public List<String> securityServiceIds() {
        return this.securityServiceIds;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Integer segmentationId() {
        return this.segmentationId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetShareNetworkResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cidr;
        private String description;
        private String id;
        private Integer ipVersion;
        private String name;
        private String networkType;
        private String neutronNetId;
        private String neutronSubnetId;
        private String projectId;
        private String region;
        private @Nullable String securityServiceId;
        private List<String> securityServiceIds;
        private Integer segmentationId;
        public Builder() {}
        public Builder(GetShareNetworkResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cidr = defaults.cidr;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.ipVersion = defaults.ipVersion;
    	      this.name = defaults.name;
    	      this.networkType = defaults.networkType;
    	      this.neutronNetId = defaults.neutronNetId;
    	      this.neutronSubnetId = defaults.neutronSubnetId;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.securityServiceId = defaults.securityServiceId;
    	      this.securityServiceIds = defaults.securityServiceIds;
    	      this.segmentationId = defaults.segmentationId;
        }

        @CustomType.Setter
        public Builder cidr(String cidr) {
            if (cidr == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "cidr");
            }
            this.cidr = cidr;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ipVersion(Integer ipVersion) {
            if (ipVersion == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "ipVersion");
            }
            this.ipVersion = ipVersion;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder networkType(String networkType) {
            if (networkType == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "networkType");
            }
            this.networkType = networkType;
            return this;
        }
        @CustomType.Setter
        public Builder neutronNetId(String neutronNetId) {
            if (neutronNetId == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "neutronNetId");
            }
            this.neutronNetId = neutronNetId;
            return this;
        }
        @CustomType.Setter
        public Builder neutronSubnetId(String neutronSubnetId) {
            if (neutronSubnetId == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "neutronSubnetId");
            }
            this.neutronSubnetId = neutronSubnetId;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder securityServiceId(@Nullable String securityServiceId) {

            this.securityServiceId = securityServiceId;
            return this;
        }
        @CustomType.Setter
        public Builder securityServiceIds(List<String> securityServiceIds) {
            if (securityServiceIds == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "securityServiceIds");
            }
            this.securityServiceIds = securityServiceIds;
            return this;
        }
        public Builder securityServiceIds(String... securityServiceIds) {
            return securityServiceIds(List.of(securityServiceIds));
        }
        @CustomType.Setter
        public Builder segmentationId(Integer segmentationId) {
            if (segmentationId == null) {
              throw new MissingRequiredPropertyException("GetShareNetworkResult", "segmentationId");
            }
            this.segmentationId = segmentationId;
            return this;
        }
        public GetShareNetworkResult build() {
            final var _resultValue = new GetShareNetworkResult();
            _resultValue.cidr = cidr;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.ipVersion = ipVersion;
            _resultValue.name = name;
            _resultValue.networkType = networkType;
            _resultValue.neutronNetId = neutronNetId;
            _resultValue.neutronSubnetId = neutronSubnetId;
            _resultValue.projectId = projectId;
            _resultValue.region = region;
            _resultValue.securityServiceId = securityServiceId;
            _resultValue.securityServiceIds = securityServiceIds;
            _resultValue.segmentationId = segmentationId;
            return _resultValue;
        }
    }
}
