// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetProjectResult {
    /**
     * @return The description of the project.
     * 
     */
    private final String description;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final String domainId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable Boolean enabled;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable Boolean isDomain;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable String name;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final @Nullable String parentId;
    /**
     * @return The region the project is located in.
     * 
     */
    private final String region;
    /**
     * @return See Argument Reference above.
     * 
     */
    private final List<String> tags;

    @CustomType.Constructor
    private GetProjectResult(
        @CustomType.Parameter("description") String description,
        @CustomType.Parameter("domainId") String domainId,
        @CustomType.Parameter("enabled") @Nullable Boolean enabled,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("isDomain") @Nullable Boolean isDomain,
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("parentId") @Nullable String parentId,
        @CustomType.Parameter("region") String region,
        @CustomType.Parameter("tags") List<String> tags) {
        this.description = description;
        this.domainId = domainId;
        this.enabled = enabled;
        this.id = id;
        this.isDomain = isDomain;
        this.name = name;
        this.parentId = parentId;
        this.region = region;
        this.tags = tags;
    }

    /**
     * @return The description of the project.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String domainId() {
        return this.domainId;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
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
    public Optional<Boolean> isDomain() {
        return Optional.ofNullable(this.isDomain);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> parentId() {
        return Optional.ofNullable(this.parentId);
    }
    /**
     * @return The region the project is located in.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String description;
        private String domainId;
        private @Nullable Boolean enabled;
        private String id;
        private @Nullable Boolean isDomain;
        private @Nullable String name;
        private @Nullable String parentId;
        private String region;
        private List<String> tags;

        public Builder() {
    	      // Empty
        }

        public Builder(GetProjectResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.domainId = defaults.domainId;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.isDomain = defaults.isDomain;
    	      this.name = defaults.name;
    	      this.parentId = defaults.parentId;
    	      this.region = defaults.region;
    	      this.tags = defaults.tags;
        }

        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        public Builder domainId(String domainId) {
            this.domainId = Objects.requireNonNull(domainId);
            return this;
        }
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder isDomain(@Nullable Boolean isDomain) {
            this.isDomain = isDomain;
            return this;
        }
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder parentId(@Nullable String parentId) {
            this.parentId = parentId;
            return this;
        }
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }        public GetProjectResult build() {
            return new GetProjectResult(description, domainId, enabled, id, isDomain, name, parentId, region, tags);
        }
    }
}
