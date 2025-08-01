// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSubnetIdsV2Result {
    private @Nullable String cidr;
    private @Nullable String description;
    private @Nullable Boolean dhcpEnabled;
    private @Nullable Boolean dnsPublishFixedIp;
    private @Nullable String gatewayIp;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable Integer ipVersion;
    private @Nullable String ipv6AddressMode;
    private String ipv6RaMode;
    private @Nullable String name;
    private @Nullable String nameRegex;
    private @Nullable String networkId;
    private String region;
    private @Nullable String segmentId;
    private @Nullable String sortDirection;
    private @Nullable String sortKey;
    private @Nullable String subnetpoolId;
    private @Nullable List<String> tags;
    private @Nullable String tenantId;

    private GetSubnetIdsV2Result() {}
    public Optional<String> cidr() {
        return Optional.ofNullable(this.cidr);
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public Optional<Boolean> dhcpEnabled() {
        return Optional.ofNullable(this.dhcpEnabled);
    }
    public Optional<Boolean> dnsPublishFixedIp() {
        return Optional.ofNullable(this.dnsPublishFixedIp);
    }
    public Optional<String> gatewayIp() {
        return Optional.ofNullable(this.gatewayIp);
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
    public Optional<Integer> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }
    public Optional<String> ipv6AddressMode() {
        return Optional.ofNullable(this.ipv6AddressMode);
    }
    public String ipv6RaMode() {
        return this.ipv6RaMode;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public Optional<String> networkId() {
        return Optional.ofNullable(this.networkId);
    }
    public String region() {
        return this.region;
    }
    public Optional<String> segmentId() {
        return Optional.ofNullable(this.segmentId);
    }
    public Optional<String> sortDirection() {
        return Optional.ofNullable(this.sortDirection);
    }
    public Optional<String> sortKey() {
        return Optional.ofNullable(this.sortKey);
    }
    public Optional<String> subnetpoolId() {
        return Optional.ofNullable(this.subnetpoolId);
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

    public static Builder builder(GetSubnetIdsV2Result defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cidr;
        private @Nullable String description;
        private @Nullable Boolean dhcpEnabled;
        private @Nullable Boolean dnsPublishFixedIp;
        private @Nullable String gatewayIp;
        private String id;
        private List<String> ids;
        private @Nullable Integer ipVersion;
        private @Nullable String ipv6AddressMode;
        private String ipv6RaMode;
        private @Nullable String name;
        private @Nullable String nameRegex;
        private @Nullable String networkId;
        private String region;
        private @Nullable String segmentId;
        private @Nullable String sortDirection;
        private @Nullable String sortKey;
        private @Nullable String subnetpoolId;
        private @Nullable List<String> tags;
        private @Nullable String tenantId;
        public Builder() {}
        public Builder(GetSubnetIdsV2Result defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cidr = defaults.cidr;
    	      this.description = defaults.description;
    	      this.dhcpEnabled = defaults.dhcpEnabled;
    	      this.dnsPublishFixedIp = defaults.dnsPublishFixedIp;
    	      this.gatewayIp = defaults.gatewayIp;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.ipVersion = defaults.ipVersion;
    	      this.ipv6AddressMode = defaults.ipv6AddressMode;
    	      this.ipv6RaMode = defaults.ipv6RaMode;
    	      this.name = defaults.name;
    	      this.nameRegex = defaults.nameRegex;
    	      this.networkId = defaults.networkId;
    	      this.region = defaults.region;
    	      this.segmentId = defaults.segmentId;
    	      this.sortDirection = defaults.sortDirection;
    	      this.sortKey = defaults.sortKey;
    	      this.subnetpoolId = defaults.subnetpoolId;
    	      this.tags = defaults.tags;
    	      this.tenantId = defaults.tenantId;
        }

        @CustomType.Setter
        public Builder cidr(@Nullable String cidr) {

            this.cidr = cidr;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder dhcpEnabled(@Nullable Boolean dhcpEnabled) {

            this.dhcpEnabled = dhcpEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder dnsPublishFixedIp(@Nullable Boolean dnsPublishFixedIp) {

            this.dnsPublishFixedIp = dnsPublishFixedIp;
            return this;
        }
        @CustomType.Setter
        public Builder gatewayIp(@Nullable String gatewayIp) {

            this.gatewayIp = gatewayIp;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSubnetIdsV2Result", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            if (ids == null) {
              throw new MissingRequiredPropertyException("GetSubnetIdsV2Result", "ids");
            }
            this.ids = ids;
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder ipVersion(@Nullable Integer ipVersion) {

            this.ipVersion = ipVersion;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6AddressMode(@Nullable String ipv6AddressMode) {

            this.ipv6AddressMode = ipv6AddressMode;
            return this;
        }
        @CustomType.Setter
        public Builder ipv6RaMode(String ipv6RaMode) {
            if (ipv6RaMode == null) {
              throw new MissingRequiredPropertyException("GetSubnetIdsV2Result", "ipv6RaMode");
            }
            this.ipv6RaMode = ipv6RaMode;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nameRegex(@Nullable String nameRegex) {

            this.nameRegex = nameRegex;
            return this;
        }
        @CustomType.Setter
        public Builder networkId(@Nullable String networkId) {

            this.networkId = networkId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetSubnetIdsV2Result", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder segmentId(@Nullable String segmentId) {

            this.segmentId = segmentId;
            return this;
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
        public Builder subnetpoolId(@Nullable String subnetpoolId) {

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
        public Builder tenantId(@Nullable String tenantId) {

            this.tenantId = tenantId;
            return this;
        }
        public GetSubnetIdsV2Result build() {
            final var _resultValue = new GetSubnetIdsV2Result();
            _resultValue.cidr = cidr;
            _resultValue.description = description;
            _resultValue.dhcpEnabled = dhcpEnabled;
            _resultValue.dnsPublishFixedIp = dnsPublishFixedIp;
            _resultValue.gatewayIp = gatewayIp;
            _resultValue.id = id;
            _resultValue.ids = ids;
            _resultValue.ipVersion = ipVersion;
            _resultValue.ipv6AddressMode = ipv6AddressMode;
            _resultValue.ipv6RaMode = ipv6RaMode;
            _resultValue.name = name;
            _resultValue.nameRegex = nameRegex;
            _resultValue.networkId = networkId;
            _resultValue.region = region;
            _resultValue.segmentId = segmentId;
            _resultValue.sortDirection = sortDirection;
            _resultValue.sortKey = sortKey;
            _resultValue.subnetpoolId = subnetpoolId;
            _resultValue.tags = tags;
            _resultValue.tenantId = tenantId;
            return _resultValue;
        }
    }
}
