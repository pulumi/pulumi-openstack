// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.networking.outputs.GetSubnetAllocationPool;
import com.pulumi.openstack.networking.outputs.GetSubnetHostRoute;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSubnetResult {
    /**
     * @return A set of string tags applied on the subnet.
     * 
     */
    private List<String> allTags;
    /**
     * @return Allocation pools of the subnet.
     * 
     */
    private List<GetSubnetAllocationPool> allocationPools;
    private String cidr;
    private String description;
    private @Nullable Boolean dhcpEnabled;
    /**
     * @return DNS Nameservers of the subnet.
     * 
     */
    private List<String> dnsNameservers;
    /**
     * @return Whether the subnet has DHCP enabled or not.
     * 
     */
    private Boolean enableDhcp;
    private String gatewayIp;
    /**
     * @return Host Routes of the subnet.
     * 
     */
    private List<GetSubnetHostRoute> hostRoutes;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer ipVersion;
    private String ipv6AddressMode;
    private String ipv6RaMode;
    private String name;
    private String networkId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String region;
    /**
     * @return Service types of the subnet.
     * 
     */
    private List<String> serviceTypes;
    private String subnetId;
    private String subnetpoolId;
    private @Nullable List<String> tags;
    private String tenantId;

    private GetSubnetResult() {}
    /**
     * @return A set of string tags applied on the subnet.
     * 
     */
    public List<String> allTags() {
        return this.allTags;
    }
    /**
     * @return Allocation pools of the subnet.
     * 
     */
    public List<GetSubnetAllocationPool> allocationPools() {
        return this.allocationPools;
    }
    public String cidr() {
        return this.cidr;
    }
    public String description() {
        return this.description;
    }
    public Optional<Boolean> dhcpEnabled() {
        return Optional.ofNullable(this.dhcpEnabled);
    }
    /**
     * @return DNS Nameservers of the subnet.
     * 
     */
    public List<String> dnsNameservers() {
        return this.dnsNameservers;
    }
    /**
     * @return Whether the subnet has DHCP enabled or not.
     * 
     */
    public Boolean enableDhcp() {
        return this.enableDhcp;
    }
    public String gatewayIp() {
        return this.gatewayIp;
    }
    /**
     * @return Host Routes of the subnet.
     * 
     */
    public List<GetSubnetHostRoute> hostRoutes() {
        return this.hostRoutes;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Integer ipVersion() {
        return this.ipVersion;
    }
    public String ipv6AddressMode() {
        return this.ipv6AddressMode;
    }
    public String ipv6RaMode() {
        return this.ipv6RaMode;
    }
    public String name() {
        return this.name;
    }
    public String networkId() {
        return this.networkId;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return Service types of the subnet.
     * 
     */
    public List<String> serviceTypes() {
        return this.serviceTypes;
    }
    public String subnetId() {
        return this.subnetId;
    }
    public String subnetpoolId() {
        return this.subnetpoolId;
    }
    public List<String> tags() {
        return this.tags == null ? List.of() : this.tags;
    }
    public String tenantId() {
        return this.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSubnetResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> allTags;
        private List<GetSubnetAllocationPool> allocationPools;
        private String cidr;
        private String description;
        private @Nullable Boolean dhcpEnabled;
        private List<String> dnsNameservers;
        private Boolean enableDhcp;
        private String gatewayIp;
        private List<GetSubnetHostRoute> hostRoutes;
        private String id;
        private Integer ipVersion;
        private String ipv6AddressMode;
        private String ipv6RaMode;
        private String name;
        private String networkId;
        private String region;
        private List<String> serviceTypes;
        private String subnetId;
        private String subnetpoolId;
        private @Nullable List<String> tags;
        private String tenantId;
        public Builder() {}
        public Builder(GetSubnetResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allTags = defaults.allTags;
    	      this.allocationPools = defaults.allocationPools;
    	      this.cidr = defaults.cidr;
    	      this.description = defaults.description;
    	      this.dhcpEnabled = defaults.dhcpEnabled;
    	      this.dnsNameservers = defaults.dnsNameservers;
    	      this.enableDhcp = defaults.enableDhcp;
    	      this.gatewayIp = defaults.gatewayIp;
    	      this.hostRoutes = defaults.hostRoutes;
    	      this.id = defaults.id;
    	      this.ipVersion = defaults.ipVersion;
    	      this.ipv6AddressMode = defaults.ipv6AddressMode;
    	      this.ipv6RaMode = defaults.ipv6RaMode;
    	      this.name = defaults.name;
    	      this.networkId = defaults.networkId;
    	      this.region = defaults.region;
    	      this.serviceTypes = defaults.serviceTypes;
    	      this.subnetId = defaults.subnetId;
    	      this.subnetpoolId = defaults.subnetpoolId;
    	      this.tags = defaults.tags;
    	      this.tenantId = defaults.tenantId;
        }

        @CustomType.Setter
        public Builder allTags(List<String> allTags) {
            if (allTags == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "allTags");
            }
            this.allTags = allTags;
            return this;
        }
        public Builder allTags(String... allTags) {
            return allTags(List.of(allTags));
        }
        @CustomType.Setter
        public Builder allocationPools(List<GetSubnetAllocationPool> allocationPools) {
            if (allocationPools == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "allocationPools");
            }
            this.allocationPools = allocationPools;
            return this;
        }
        public Builder allocationPools(GetSubnetAllocationPool... allocationPools) {
            return allocationPools(List.of(allocationPools));
        }
        @CustomType.Setter
        public Builder cidr(String cidr) {
            if (cidr == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "cidr");
            }
            this.cidr = cidr;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder dhcpEnabled(@Nullable Boolean dhcpEnabled) {

            this.dhcpEnabled = dhcpEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder dnsNameservers(List<String> dnsNameservers) {
            if (dnsNameservers == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "dnsNameservers");
            }
            this.dnsNameservers = dnsNameservers;
            return this;
        }
        public Builder dnsNameservers(String... dnsNameservers) {
            return dnsNameservers(List.of(dnsNameservers));
        }
        @CustomType.Setter
        public Builder enableDhcp(Boolean enableDhcp) {
            if (enableDhcp == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "enableDhcp");
            }
            this.enableDhcp = enableDhcp;
            return this;
        }
        @CustomType.Setter
        public Builder gatewayIp(String gatewayIp) {
            if (gatewayIp == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "gatewayIp");
            }
            this.gatewayIp = gatewayIp;
            return this;
        }
        @CustomType.Setter
        public Builder hostRoutes(List<GetSubnetHostRoute> hostRoutes) {
            if (hostRoutes == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "hostRoutes");
            }
            this.hostRoutes = hostRoutes;
            return this;
        }
        public Builder hostRoutes(GetSubnetHostRoute... hostRoutes) {
            return hostRoutes(List.of(hostRoutes));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ipVersion(Integer ipVersion) {
            if (ipVersion == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "ipVersion");
            }
            this.ipVersion = ipVersion;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6AddressMode(String ipv6AddressMode) {
            if (ipv6AddressMode == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "ipv6AddressMode");
            }
            this.ipv6AddressMode = ipv6AddressMode;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6RaMode(String ipv6RaMode) {
            if (ipv6RaMode == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "ipv6RaMode");
            }
            this.ipv6RaMode = ipv6RaMode;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder networkId(String networkId) {
            if (networkId == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "networkId");
            }
            this.networkId = networkId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder serviceTypes(List<String> serviceTypes) {
            if (serviceTypes == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "serviceTypes");
            }
            this.serviceTypes = serviceTypes;
            return this;
        }
        public Builder serviceTypes(String... serviceTypes) {
            return serviceTypes(List.of(serviceTypes));
        }
        @CustomType.Setter
        public Builder subnetId(String subnetId) {
            if (subnetId == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "subnetId");
            }
            this.subnetId = subnetId;
            return this;
        }
        @CustomType.Setter
        public Builder subnetpoolId(String subnetpoolId) {
            if (subnetpoolId == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "subnetpoolId");
            }
            this.subnetpoolId = subnetpoolId;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable List<String> tags) {

            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder tenantId(String tenantId) {
            if (tenantId == null) {
              throw new MissingRequiredPropertyException("GetSubnetResult", "tenantId");
            }
            this.tenantId = tenantId;
            return this;
        }
        public GetSubnetResult build() {
            final var _resultValue = new GetSubnetResult();
            _resultValue.allTags = allTags;
            _resultValue.allocationPools = allocationPools;
            _resultValue.cidr = cidr;
            _resultValue.description = description;
            _resultValue.dhcpEnabled = dhcpEnabled;
            _resultValue.dnsNameservers = dnsNameservers;
            _resultValue.enableDhcp = enableDhcp;
            _resultValue.gatewayIp = gatewayIp;
            _resultValue.hostRoutes = hostRoutes;
            _resultValue.id = id;
            _resultValue.ipVersion = ipVersion;
            _resultValue.ipv6AddressMode = ipv6AddressMode;
            _resultValue.ipv6RaMode = ipv6RaMode;
            _resultValue.name = name;
            _resultValue.networkId = networkId;
            _resultValue.region = region;
            _resultValue.serviceTypes = serviceTypes;
            _resultValue.subnetId = subnetId;
            _resultValue.subnetpoolId = subnetpoolId;
            _resultValue.tags = tags;
            _resultValue.tenantId = tenantId;
            return _resultValue;
        }
    }
}
