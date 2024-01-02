// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.networking.outputs.GetRouterExternalFixedIp;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRouterResult {
    private @Nullable Boolean adminStateUp;
    /**
     * @return The set of string tags applied on the router.
     * 
     */
    private List<String> allTags;
    /**
     * @return The availability zone that is used to make router resources highly available.
     * 
     */
    private List<String> availabilityZoneHints;
    private @Nullable String description;
    private @Nullable Boolean distributed;
    /**
     * @return The value that points out if the Source NAT is enabled on the router.
     * 
     */
    private Boolean enableSnat;
    /**
     * @return The external fixed IPs of the router.
     * 
     */
    private List<GetRouterExternalFixedIp> externalFixedIps;
    /**
     * @return The network UUID of an external gateway for the router.
     * 
     */
    private String externalNetworkId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String name;
    private @Nullable String region;
    private @Nullable String routerId;
    private @Nullable String status;
    private @Nullable List<String> tags;
    private @Nullable String tenantId;

    private GetRouterResult() {}
    public Optional<Boolean> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }
    /**
     * @return The set of string tags applied on the router.
     * 
     */
    public List<String> allTags() {
        return this.allTags;
    }
    /**
     * @return The availability zone that is used to make router resources highly available.
     * 
     */
    public List<String> availabilityZoneHints() {
        return this.availabilityZoneHints;
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public Optional<Boolean> distributed() {
        return Optional.ofNullable(this.distributed);
    }
    /**
     * @return The value that points out if the Source NAT is enabled on the router.
     * 
     */
    public Boolean enableSnat() {
        return this.enableSnat;
    }
    /**
     * @return The external fixed IPs of the router.
     * 
     */
    public List<GetRouterExternalFixedIp> externalFixedIps() {
        return this.externalFixedIps;
    }
    /**
     * @return The network UUID of an external gateway for the router.
     * 
     */
    public String externalNetworkId() {
        return this.externalNetworkId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    public Optional<String> routerId() {
        return Optional.ofNullable(this.routerId);
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

    public static Builder builder(GetRouterResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean adminStateUp;
        private List<String> allTags;
        private List<String> availabilityZoneHints;
        private @Nullable String description;
        private @Nullable Boolean distributed;
        private Boolean enableSnat;
        private List<GetRouterExternalFixedIp> externalFixedIps;
        private String externalNetworkId;
        private String id;
        private @Nullable String name;
        private @Nullable String region;
        private @Nullable String routerId;
        private @Nullable String status;
        private @Nullable List<String> tags;
        private @Nullable String tenantId;
        public Builder() {}
        public Builder(GetRouterResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminStateUp = defaults.adminStateUp;
    	      this.allTags = defaults.allTags;
    	      this.availabilityZoneHints = defaults.availabilityZoneHints;
    	      this.description = defaults.description;
    	      this.distributed = defaults.distributed;
    	      this.enableSnat = defaults.enableSnat;
    	      this.externalFixedIps = defaults.externalFixedIps;
    	      this.externalNetworkId = defaults.externalNetworkId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.region = defaults.region;
    	      this.routerId = defaults.routerId;
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
        public Builder allTags(List<String> allTags) {
            if (allTags == null) {
              throw new MissingRequiredPropertyException("GetRouterResult", "allTags");
            }
            this.allTags = allTags;
            return this;
        }
        public Builder allTags(String... allTags) {
            return allTags(List.of(allTags));
        }
        @CustomType.Setter
        public Builder availabilityZoneHints(List<String> availabilityZoneHints) {
            if (availabilityZoneHints == null) {
              throw new MissingRequiredPropertyException("GetRouterResult", "availabilityZoneHints");
            }
            this.availabilityZoneHints = availabilityZoneHints;
            return this;
        }
        public Builder availabilityZoneHints(String... availabilityZoneHints) {
            return availabilityZoneHints(List.of(availabilityZoneHints));
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder distributed(@Nullable Boolean distributed) {

            this.distributed = distributed;
            return this;
        }
        @CustomType.Setter
        public Builder enableSnat(Boolean enableSnat) {
            if (enableSnat == null) {
              throw new MissingRequiredPropertyException("GetRouterResult", "enableSnat");
            }
            this.enableSnat = enableSnat;
            return this;
        }
        @CustomType.Setter
        public Builder externalFixedIps(List<GetRouterExternalFixedIp> externalFixedIps) {
            if (externalFixedIps == null) {
              throw new MissingRequiredPropertyException("GetRouterResult", "externalFixedIps");
            }
            this.externalFixedIps = externalFixedIps;
            return this;
        }
        public Builder externalFixedIps(GetRouterExternalFixedIp... externalFixedIps) {
            return externalFixedIps(List.of(externalFixedIps));
        }
        @CustomType.Setter
        public Builder externalNetworkId(String externalNetworkId) {
            if (externalNetworkId == null) {
              throw new MissingRequiredPropertyException("GetRouterResult", "externalNetworkId");
            }
            this.externalNetworkId = externalNetworkId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRouterResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {

            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder routerId(@Nullable String routerId) {

            this.routerId = routerId;
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
        public GetRouterResult build() {
            final var _resultValue = new GetRouterResult();
            _resultValue.adminStateUp = adminStateUp;
            _resultValue.allTags = allTags;
            _resultValue.availabilityZoneHints = availabilityZoneHints;
            _resultValue.description = description;
            _resultValue.distributed = distributed;
            _resultValue.enableSnat = enableSnat;
            _resultValue.externalFixedIps = externalFixedIps;
            _resultValue.externalNetworkId = externalNetworkId;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.region = region;
            _resultValue.routerId = routerId;
            _resultValue.status = status;
            _resultValue.tags = tags;
            _resultValue.tenantId = tenantId;
            return _resultValue;
        }
    }
}
