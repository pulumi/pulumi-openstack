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
    private @Nullable Boolean adminStateUp;
    private @Nullable String description;
    private @Nullable String deviceId;
    private @Nullable String deviceOwner;
    private @Nullable String dnsName;
    private @Nullable String fixedIp;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
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

    private GetPortIdsResult() {}
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
    @CustomType.Builder
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
        public Builder() {}
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

        @CustomType.Setter
        public Builder adminStateUp(@Nullable Boolean adminStateUp) {
            this.adminStateUp = adminStateUp;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder deviceId(@Nullable String deviceId) {
            this.deviceId = deviceId;
            return this;
        }
        @CustomType.Setter
        public Builder deviceOwner(@Nullable String deviceOwner) {
            this.deviceOwner = deviceOwner;
            return this;
        }
        @CustomType.Setter
        public Builder dnsName(@Nullable String dnsName) {
            this.dnsName = dnsName;
            return this;
        }
        @CustomType.Setter
        public Builder fixedIp(@Nullable String fixedIp) {
            this.fixedIp = fixedIp;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            this.ids = Objects.requireNonNull(ids);
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder macAddress(@Nullable String macAddress) {
            this.macAddress = macAddress;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder networkId(@Nullable String networkId) {
            this.networkId = networkId;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder securityGroupIds(@Nullable List<String> securityGroupIds) {
            this.securityGroupIds = securityGroupIds;
            return this;
        }
        public Builder securityGroupIds(String... securityGroupIds) {
            return securityGroupIds(List.of(securityGroupIds));
        }
        @CustomType.Setter
        public Builder sortDirection(@Nullable String sortDirection) {
            this.sortDirection = sortDirection;
            return this;
        }
        @CustomType.Setter
        public Builder sortKey(@Nullable String sortKey) {
            this.sortKey = sortKey;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
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
        public Builder tenantId(@Nullable String tenantId) {
            this.tenantId = tenantId;
            return this;
        }
        public GetPortIdsResult build() {
            final var o = new GetPortIdsResult();
            o.adminStateUp = adminStateUp;
            o.description = description;
            o.deviceId = deviceId;
            o.deviceOwner = deviceOwner;
            o.dnsName = dnsName;
            o.fixedIp = fixedIp;
            o.id = id;
            o.ids = ids;
            o.macAddress = macAddress;
            o.name = name;
            o.networkId = networkId;
            o.projectId = projectId;
            o.region = region;
            o.securityGroupIds = securityGroupIds;
            o.sortDirection = sortDirection;
            o.sortKey = sortKey;
            o.status = status;
            o.tags = tags;
            o.tenantId = tenantId;
            return o;
        }
    }
}
