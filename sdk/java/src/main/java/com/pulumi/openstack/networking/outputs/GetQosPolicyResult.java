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
import javax.annotation.Nullable;

@CustomType
public final class GetQosPolicyResult {
    /**
     * @return The set of string tags applied on the QoS policy.
     * 
     */
    private List<String> allTags;
    /**
     * @return The time at which QoS policy was created.
     * 
     */
    private String createdAt;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private Boolean isDefault;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String name;
    private String projectId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String region;
    /**
     * @return The revision number of the QoS policy.
     * 
     */
    private Integer revisionNumber;
    /**
     * @return See Argument Reference above.
     * 
     */
    private Boolean shared;
    private @Nullable List<String> tags;
    /**
     * @return The time at which QoS policy was created.
     * 
     */
    private String updatedAt;

    private GetQosPolicyResult() {}
    /**
     * @return The set of string tags applied on the QoS policy.
     * 
     */
    public List<String> allTags() {
        return this.allTags;
    }
    /**
     * @return The time at which QoS policy was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
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
    public Boolean isDefault() {
        return this.isDefault;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String name() {
        return this.name;
    }
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
     * @return The revision number of the QoS policy.
     * 
     */
    public Integer revisionNumber() {
        return this.revisionNumber;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Boolean shared() {
        return this.shared;
    }
    public List<String> tags() {
        return this.tags == null ? List.of() : this.tags;
    }
    /**
     * @return The time at which QoS policy was created.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetQosPolicyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> allTags;
        private String createdAt;
        private String description;
        private String id;
        private Boolean isDefault;
        private String name;
        private String projectId;
        private String region;
        private Integer revisionNumber;
        private Boolean shared;
        private @Nullable List<String> tags;
        private String updatedAt;
        public Builder() {}
        public Builder(GetQosPolicyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allTags = defaults.allTags;
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.isDefault = defaults.isDefault;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.revisionNumber = defaults.revisionNumber;
    	      this.shared = defaults.shared;
    	      this.tags = defaults.tags;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder allTags(List<String> allTags) {
            if (allTags == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "allTags");
            }
            this.allTags = allTags;
            return this;
        }
        public Builder allTags(String... allTags) {
            return allTags(List.of(allTags));
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isDefault(Boolean isDefault) {
            if (isDefault == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "isDefault");
            }
            this.isDefault = isDefault;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder revisionNumber(Integer revisionNumber) {
            if (revisionNumber == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "revisionNumber");
            }
            this.revisionNumber = revisionNumber;
            return this;
        }
        @CustomType.Setter
        public Builder shared(Boolean shared) {
            if (shared == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "shared");
            }
            this.shared = shared;
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
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetQosPolicyResult", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        public GetQosPolicyResult build() {
            final var _resultValue = new GetQosPolicyResult();
            _resultValue.allTags = allTags;
            _resultValue.createdAt = createdAt;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.isDefault = isDefault;
            _resultValue.name = name;
            _resultValue.projectId = projectId;
            _resultValue.region = region;
            _resultValue.revisionNumber = revisionNumber;
            _resultValue.shared = shared;
            _resultValue.tags = tags;
            _resultValue.updatedAt = updatedAt;
            return _resultValue;
        }
    }
}
