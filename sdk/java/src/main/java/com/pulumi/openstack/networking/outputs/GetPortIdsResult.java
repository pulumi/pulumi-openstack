// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPortIdsResult {
    private final @Nullable Boolean adminStateUp;
    private final @Nullable String description;
    private final @Nullable String deviceId;
    private final @Nullable String deviceOwner;
    private final @Nullable String dnsName;
    private final @Nullable String fixedIp;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final List<String> ids;
    private final @Nullable String macAddress;
    private final @Nullable String name;
    private final @Nullable String networkId;
    private final @Nullable String projectId;
    private final @Nullable String region;
    private final @Nullable List<String> securityGroupIds;
    private final @Nullable String sortDirection;
    private final @Nullable String sortKey;
    private final @Nullable String status;
    private final @Nullable List<String> tags;
    private final @Nullable String tenantId;

    @CustomType.Constructor
    private GetPortIdsResult(
        @CustomType.Parameter("adminStateUp") @Nullable Boolean adminStateUp,
        @CustomType.Parameter("description") @Nullable String description,
        @CustomType.Parameter("deviceId") @Nullable String deviceId,
        @CustomType.Parameter("deviceOwner") @Nullable String deviceOwner,
        @CustomType.Parameter("dnsName") @Nullable String dnsName,
        @CustomType.Parameter("fixedIp") @Nullable String fixedIp,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("ids") List<String> ids,
        @CustomType.Parameter("macAddress") @Nullable String macAddress,
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("networkId") @Nullable String networkId,
        @CustomType.Parameter("projectId") @Nullable String projectId,
        @CustomType.Parameter("region") @Nullable String region,
        @CustomType.Parameter("securityGroupIds") @Nullable List<String> securityGroupIds,
        @CustomType.Parameter("sortDirection") @Nullable String sortDirection,
        @CustomType.Parameter("sortKey") @Nullable String sortKey,
        @CustomType.Parameter("status") @Nullable String status,
        @CustomType.Parameter("tags") @Nullable List<String> tags,
        @CustomType.Parameter("tenantId") @Nullable String tenantId) {
        this.adminStateUp = adminStateUp;
        this.description = description;
        this.deviceId = deviceId;
        this.deviceOwner = deviceOwner;
        this.dnsName = dnsName;
        this.fixedIp = fixedIp;
        this.id = id;
        this.ids = ids;
        this.macAddress = macAddress;
        this.name = name;
        this.networkId = networkId;
        this.projectId = projectId;
        this.region = region;
        this.securityGroupIds = securityGroupIds;
        this.sortDirection = sortDirection;
        this.sortKey = sortKey;
        this.status = status;
        this.tags = tags;
        this.tenantId = tenantId;
    }

    public Optional<Boolean> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public Optional<String> deviceId() {
        return Optional.ofNullable(this.deviceId);
    }
    public Optional<String> deviceOwner() {
        return Optional.ofNullable(this.deviceOwner);
    }
    public Optional<String> dnsName() {
        return Optional.ofNullable(this.dnsName);
    }
    public Optional<String> fixedIp() {
        return Optional.ofNullable(this.fixedIp);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<String> ids() {
        return this.ids;
    }
    public Optional<String> macAddress() {
        return Optional.ofNullable(this.macAddress);
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> networkId() {
        return Optional.ofNullable(this.networkId);
    }
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    public List<String> securityGroupIds() {
        return this.securityGroupIds == null ? List.of() : this.securityGroupIds;
    }
    public Optional<String> sortDirection() {
        return Optional.ofNullable(this.sortDirection);
    }
    public Optional<String> sortKey() {
        return Optional.ofNullable(this.sortKey);
    }
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public List<String> tags() {
        return this.tags == null ? List.of() : this.tags;
    }
    public Optional<String> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPortIdsResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean adminStateUp;
        private @Nullable String description;
        private @Nullable String deviceId;
        private @Nullable String deviceOwner;
        private @Nullable String dnsName;
        private @Nullable String fixedIp;
        private String id;
        private List<String> ids;
        private @Nullable String macAddress;
        private @Nullable String name;
        private @Nullable String networkId;
        private @Nullable String projectId;
        private @Nullable String region;
        private @Nullable List<String> securityGroupIds;
        private @Nullable String sortDirection;
        private @Nullable String sortKey;
        private @Nullable String status;
        private @Nullable List<String> tags;
        private @Nullable String tenantId;

        public Builder() {
    	      // Empty
        }

        public Builder(GetPortIdsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminStateUp = defaults.adminStateUp;
    	      this.description = defaults.description;
    	      this.deviceId = defaults.deviceId;
    	      this.deviceOwner = defaults.deviceOwner;
    	      this.dnsName = defaults.dnsName;
    	      this.fixedIp = defaults.fixedIp;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.macAddress = defaults.macAddress;
    	      this.name = defaults.name;
    	      this.networkId = defaults.networkId;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.securityGroupIds = defaults.securityGroupIds;
    	      this.sortDirection = defaults.sortDirection;
    	      this.sortKey = defaults.sortKey;
    	      this.status = defaults.status;
    	      this.tags = defaults.tags;
    	      this.tenantId = defaults.tenantId;
        }

        public Builder adminStateUp(@Nullable Boolean adminStateUp) {
            this.adminStateUp = adminStateUp;
            return this;
        }
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        public Builder deviceId(@Nullable String deviceId) {
            this.deviceId = deviceId;
            return this;
        }
        public Builder deviceOwner(@Nullable String deviceOwner) {
            this.deviceOwner = deviceOwner;
            return this;
        }
        public Builder dnsName(@Nullable String dnsName) {
            this.dnsName = dnsName;
            return this;
        }
        public Builder fixedIp(@Nullable String fixedIp) {
            this.fixedIp = fixedIp;
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        public Builder macAddress(@Nullable String macAddress) {
            this.macAddress = macAddress;
            return this;
        }
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder networkId(@Nullable String networkId) {
            this.networkId = networkId;
            return this;
        }
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        public Builder securityGroupIds(@Nullable List<String> securityGroupIds) {
            this.securityGroupIds = securityGroupIds;
            return this;
        }
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }
        public Builder sortDirection(@Nullable String sortDirection) {
            this.sortDirection = sortDirection;
            return this;
        }
        public Builder sortKey(@Nullable String sortKey) {
            this.sortKey = sortKey;
            return this;
        }
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public Builder tags(@Nullable List<String> tags) {
            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        public Builder tenantId(@Nullable String tenantId) {
            this.tenantId = tenantId;
            return this;
        }        public GetPortIdsResult build() {
            return new GetPortIdsResult(adminStateUp, description, deviceId, deviceOwner, dnsName, fixedIp, id, ids, macAddress, name, networkId, projectId, region, securityGroupIds, sortDirection, sortKey, status, tags, tenantId);
        }
    }
}