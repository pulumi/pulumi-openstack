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
    private String description;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String domainId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable Boolean isDomain;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String name;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String parentId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String projectId;
    /**
     * @return The region the project is located in.
     * 
     */
    private String region;
    /**
     * @return See Argument Reference above.
     * 
     */
    private List<String> tags;

    private GetProjectResult() {}
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
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
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
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String domainId;
        private @Nullable Boolean enabled;
        private String id;
        private @Nullable Boolean isDomain;
        private @Nullable String name;
        private @Nullable String parentId;
        private @Nullable String projectId;
        private String region;
        private List<String> tags;
        public Builder() {}
        public Builder(GetProjectResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.domainId = defaults.domainId;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.isDomain = defaults.isDomain;
    	      this.name = defaults.name;
    	      this.parentId = defaults.parentId;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder domainId(String domainId) {
            this.domainId = Objects.requireNonNull(domainId);
            return this;
        }
        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder isDomain(@Nullable Boolean isDomain) {
            this.isDomain = isDomain;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder parentId(@Nullable String parentId) {
            this.parentId = parentId;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        public GetProjectResult build() {
            final var o = new GetProjectResult();
            o.description = description;
            o.domainId = domainId;
            o.enabled = enabled;
            o.id = id;
            o.isDomain = isDomain;
            o.name = name;
            o.parentId = parentId;
            o.projectId = projectId;
            o.region = region;
            o.tags = tags;
            return o;
        }
    }
}
