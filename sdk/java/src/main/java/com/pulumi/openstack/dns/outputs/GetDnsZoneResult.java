// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.dns.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDnsZoneResult {
    private @Nullable String allProjects;
    /**
     * @return Attributes of the DNS Service scheduler.
     * 
     */
    private Map<String,Object> attributes;
    /**
     * @return The time the zone was created.
     * 
     */
    private String createdAt;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String description;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String email;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return An array of master DNS servers. When `type` is  `SECONDARY`.
     * 
     */
    private List<String> masters;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String name;
    /**
     * @return The ID of the pool hosting the zone.
     * 
     */
    private String poolId;
    /**
     * @return The project ID that owns the zone.
     * 
     */
    private String projectId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String region;
    /**
     * @return The serial number of the zone.
     * 
     */
    private Integer serial;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String status;
    /**
     * @return The time the zone was transferred.
     * 
     */
    private String transferredAt;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable Integer ttl;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String type;
    /**
     * @return The time the zone was last updated.
     * 
     */
    private String updatedAt;
    /**
     * @return The version of the zone.
     * 
     */
    private Integer version;

    private GetDnsZoneResult() {}
    public Optional<String> allProjects() {
        return Optional.ofNullable(this.allProjects);
    }
    /**
     * @return Attributes of the DNS Service scheduler.
     * 
     */
    public Map<String,Object> attributes() {
        return this.attributes;
    }
    /**
     * @return The time the zone was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> email() {
        return Optional.ofNullable(this.email);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return An array of master DNS servers. When `type` is  `SECONDARY`.
     * 
     */
    public List<String> masters() {
        return this.masters;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The ID of the pool hosting the zone.
     * 
     */
    public String poolId() {
        return this.poolId;
    }
    /**
     * @return The project ID that owns the zone.
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
     * @return The serial number of the zone.
     * 
     */
    public Integer serial() {
        return this.serial;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    /**
     * @return The time the zone was transferred.
     * 
     */
    public String transferredAt() {
        return this.transferredAt;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<Integer> ttl() {
        return Optional.ofNullable(this.ttl);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    /**
     * @return The time the zone was last updated.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return The version of the zone.
     * 
     */
    public Integer version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDnsZoneResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String allProjects;
        private Map<String,Object> attributes;
        private String createdAt;
        private @Nullable String description;
        private @Nullable String email;
        private String id;
        private List<String> masters;
        private @Nullable String name;
        private String poolId;
        private String projectId;
        private String region;
        private Integer serial;
        private @Nullable String status;
        private String transferredAt;
        private @Nullable Integer ttl;
        private @Nullable String type;
        private String updatedAt;
        private Integer version;
        public Builder() {}
        public Builder(GetDnsZoneResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allProjects = defaults.allProjects;
    	      this.attributes = defaults.attributes;
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.email = defaults.email;
    	      this.id = defaults.id;
    	      this.masters = defaults.masters;
    	      this.name = defaults.name;
    	      this.poolId = defaults.poolId;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.serial = defaults.serial;
    	      this.status = defaults.status;
    	      this.transferredAt = defaults.transferredAt;
    	      this.ttl = defaults.ttl;
    	      this.type = defaults.type;
    	      this.updatedAt = defaults.updatedAt;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder allProjects(@Nullable String allProjects) {
            this.allProjects = allProjects;
            return this;
        }
        @CustomType.Setter
        public Builder attributes(Map<String,Object> attributes) {
            this.attributes = Objects.requireNonNull(attributes);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder email(@Nullable String email) {
            this.email = email;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder masters(List<String> masters) {
            this.masters = Objects.requireNonNull(masters);
            return this;
        }
        public Builder masters(String... masters) {
            return masters(List.of(masters));
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder poolId(String poolId) {
            this.poolId = Objects.requireNonNull(poolId);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder serial(Integer serial) {
            this.serial = Objects.requireNonNull(serial);
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder transferredAt(String transferredAt) {
            this.transferredAt = Objects.requireNonNull(transferredAt);
            return this;
        }
        @CustomType.Setter
        public Builder ttl(@Nullable Integer ttl) {
            this.ttl = ttl;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder version(Integer version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetDnsZoneResult build() {
            final var o = new GetDnsZoneResult();
            o.allProjects = allProjects;
            o.attributes = attributes;
            o.createdAt = createdAt;
            o.description = description;
            o.email = email;
            o.id = id;
            o.masters = masters;
            o.name = name;
            o.poolId = poolId;
            o.projectId = projectId;
            o.region = region;
            o.serial = serial;
            o.status = status;
            o.transferredAt = transferredAt;
            o.ttl = ttl;
            o.type = type;
            o.updatedAt = updatedAt;
            o.version = version;
            return o;
        }
    }
}
