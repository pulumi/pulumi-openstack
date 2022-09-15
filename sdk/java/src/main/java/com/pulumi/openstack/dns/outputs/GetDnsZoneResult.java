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
    private final @Nullable String allProjects;
    /**
     * @return Attributes of the DNS Service scheduler.
     * 
     */
    private final Map<String,Object> attributes;
    /**
     * @return The time the zone was created.
     * 
     */
    private final String createdAt;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable String description;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable String email;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return An array of master DNS servers. When `type` is  `SECONDARY`.
     * 
     */
    private final List<String> masters;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable String name;
    /**
     * @return The ID of the pool hosting the zone.
     * 
     */
    private final String poolId;
    /**
     * @return The project ID that owns the zone.
     * 
     */
    private final String projectId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String region;
    /**
     * @return The serial number of the zone.
     * 
     */
    private final Integer serial;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable String status;
    /**
     * @return The time the zone was transferred.
     * 
     */
    private final String transferredAt;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable Integer ttl;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable String type;
    /**
     * @return The time the zone was last updated.
     * 
     */
    private final String updatedAt;
    /**
     * @return The version of the zone.
     * 
     */
    private final Integer version;

    @CustomType.Constructor
    private GetDnsZoneResult(
        @CustomType.Parameter("allProjects") @Nullable String allProjects,
        @CustomType.Parameter("attributes") Map<String,Object> attributes,
        @CustomType.Parameter("createdAt") String createdAt,
        @CustomType.Parameter("description") @Nullable String description,
        @CustomType.Parameter("email") @Nullable String email,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("masters") List<String> masters,
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("poolId") String poolId,
        @CustomType.Parameter("projectId") String projectId,
        @CustomType.Parameter("region") String region,
        @CustomType.Parameter("serial") Integer serial,
        @CustomType.Parameter("status") @Nullable String status,
        @CustomType.Parameter("transferredAt") String transferredAt,
        @CustomType.Parameter("ttl") @Nullable Integer ttl,
        @CustomType.Parameter("type") @Nullable String type,
        @CustomType.Parameter("updatedAt") String updatedAt,
        @CustomType.Parameter("version") Integer version) {
        this.allProjects = allProjects;
        this.attributes = attributes;
        this.createdAt = createdAt;
        this.description = description;
        this.email = email;
        this.id = id;
        this.masters = masters;
        this.name = name;
        this.poolId = poolId;
        this.projectId = projectId;
        this.region = region;
        this.serial = serial;
        this.status = status;
        this.transferredAt = transferredAt;
        this.ttl = ttl;
        this.type = type;
        this.updatedAt = updatedAt;
        this.version = version;
    }

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

        public Builder() {
    	      // Empty
        }

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

        public Builder allProjects(@Nullable String allProjects) {
            this.allProjects = allProjects;
            return this;
        }
        public Builder attributes(Map<String,Object> attributes) {
            this.attributes = Objects.requireNonNull(attributes);
            return this;
        }
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        public Builder email(@Nullable String email) {
            this.email = email;
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder masters(List<String> masters) {
            this.masters = Objects.requireNonNull(masters);
            return this;
        }
        public Builder masters(String... masters) {
            return masters(List.of(masters));
        }
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder poolId(String poolId) {
            this.poolId = Objects.requireNonNull(poolId);
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
        public Builder serial(Integer serial) {
            this.serial = Objects.requireNonNull(serial);
            return this;
        }
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public Builder transferredAt(String transferredAt) {
            this.transferredAt = Objects.requireNonNull(transferredAt);
            return this;
        }
        public Builder ttl(@Nullable Integer ttl) {
            this.ttl = ttl;
            return this;
        }
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        public Builder version(Integer version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }        public GetDnsZoneResult build() {
            return new GetDnsZoneResult(allProjects, attributes, createdAt, description, email, id, masters, name, poolId, projectId, region, serial, status, transferredAt, ttl, type, updatedAt, version);
        }
    }
}