// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem.outputs;

import com.pulumi.core.annotations.CustomType;
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
    private final String cidr;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final Integer ipVersion;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String name;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String networkType;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String neutronNetId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String neutronSubnetId;
    /**
     * @return The owner of the Share Network.
     * 
     */
    private final String projectId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String region;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable String securityServiceId;
    /**
     * @return The list of security service IDs associated with
     * the share network.
     * 
     */
    private final List<String> securityServiceIds;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final Integer segmentationId;

    @CustomType.Constructor
    private GetShareNetworkResult(
        @CustomType.Parameter("cidr") String cidr,
        @CustomType.Parameter("description") String description,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("ipVersion") Integer ipVersion,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("networkType") String networkType,
        @CustomType.Parameter("neutronNetId") String neutronNetId,
        @CustomType.Parameter("neutronSubnetId") String neutronSubnetId,
        @CustomType.Parameter("projectId") String projectId,
        @CustomType.Parameter("region") String region,
        @CustomType.Parameter("securityServiceId") @Nullable String securityServiceId,
        @CustomType.Parameter("securityServiceIds") List<String> securityServiceIds,
        @CustomType.Parameter("segmentationId") Integer segmentationId) {
        this.cidr = cidr;
        this.description = description;
        this.id = id;
        this.ipVersion = ipVersion;
        this.name = name;
        this.networkType = networkType;
        this.neutronNetId = neutronNetId;
        this.neutronSubnetId = neutronSubnetId;
        this.projectId = projectId;
        this.region = region;
        this.securityServiceId = securityServiceId;
        this.securityServiceIds = securityServiceIds;
        this.segmentationId = segmentationId;
    }

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

        public Builder() {
    	      // Empty
        }

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

        public Builder cidr(String cidr) {
            this.cidr = Objects.requireNonNull(cidr);
            return this;
        }
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder ipVersion(Integer ipVersion) {
            this.ipVersion = Objects.requireNonNull(ipVersion);
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder networkType(String networkType) {
            this.networkType = Objects.requireNonNull(networkType);
            return this;
        }
        public Builder neutronNetId(String neutronNetId) {
            this.neutronNetId = Objects.requireNonNull(neutronNetId);
            return this;
        }
        public Builder neutronSubnetId(String neutronSubnetId) {
            this.neutronSubnetId = Objects.requireNonNull(neutronSubnetId);
            return this;
        }
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        public Builder securityServiceId(@Nullable String securityServiceId) {
            this.securityServiceId = securityServiceId;
            return this;
        }
        public Builder securityServiceIds(List<String> securityServiceIds) {
            this.securityServiceIds = Objects.requireNonNull(securityServiceIds);
            return this;
        }
        public Builder securityServiceIds(String... securityServiceIds) {
            return securityServiceIds(List.of(securityServiceIds));
        }
        public Builder segmentationId(Integer segmentationId) {
            this.segmentationId = Objects.requireNonNull(segmentationId);
            return this;
        }        public GetShareNetworkResult build() {
            return new GetShareNetworkResult(cidr, description, id, ipVersion, name, networkType, neutronNetId, neutronSubnetId, projectId, region, securityServiceId, securityServiceIds, segmentationId);
        }
    }
}