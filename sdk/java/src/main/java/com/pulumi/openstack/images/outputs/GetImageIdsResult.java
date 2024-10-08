// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetImageIdsResult {
    private @Nullable String containerFormat;
    private @Nullable String diskFormat;
    private @Nullable Boolean hidden;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<String> ids;
    private @Nullable String memberStatus;
    private @Nullable String name;
    private @Nullable String nameRegex;
    private @Nullable String owner;
    private @Nullable Map<String,String> properties;
    private String region;
    private @Nullable Integer sizeMax;
    private @Nullable Integer sizeMin;
    private @Nullable String sort;
    private @Nullable String tag;
    private @Nullable List<String> tags;
    private @Nullable String visibility;

    private GetImageIdsResult() {}
    public Optional<String> containerFormat() {
        return Optional.ofNullable(this.containerFormat);
    }
    public Optional<String> diskFormat() {
        return Optional.ofNullable(this.diskFormat);
    }
    public Optional<Boolean> hidden() {
        return Optional.ofNullable(this.hidden);
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
    public Optional<String> memberStatus() {
        return Optional.ofNullable(this.memberStatus);
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> nameRegex() {
        return Optional.ofNullable(this.nameRegex);
    }
    public Optional<String> owner() {
        return Optional.ofNullable(this.owner);
    }
    public Map<String,String> properties() {
        return this.properties == null ? Map.of() : this.properties;
    }
    public String region() {
        return this.region;
    }
    public Optional<Integer> sizeMax() {
        return Optional.ofNullable(this.sizeMax);
    }
    public Optional<Integer> sizeMin() {
        return Optional.ofNullable(this.sizeMin);
    }
    public Optional<String> sort() {
        return Optional.ofNullable(this.sort);
    }
    public Optional<String> tag() {
        return Optional.ofNullable(this.tag);
    }
    public List<String> tags() {
        return this.tags == null ? List.of() : this.tags;
    }
    public Optional<String> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetImageIdsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String containerFormat;
        private @Nullable String diskFormat;
        private @Nullable Boolean hidden;
        private String id;
        private List<String> ids;
        private @Nullable String memberStatus;
        private @Nullable String name;
        private @Nullable String nameRegex;
        private @Nullable String owner;
        private @Nullable Map<String,String> properties;
        private String region;
        private @Nullable Integer sizeMax;
        private @Nullable Integer sizeMin;
        private @Nullable String sort;
        private @Nullable String tag;
        private @Nullable List<String> tags;
        private @Nullable String visibility;
        public Builder() {}
        public Builder(GetImageIdsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.containerFormat = defaults.containerFormat;
    	      this.diskFormat = defaults.diskFormat;
    	      this.hidden = defaults.hidden;
    	      this.id = defaults.id;
    	      this.ids = defaults.ids;
    	      this.memberStatus = defaults.memberStatus;
    	      this.name = defaults.name;
    	      this.nameRegex = defaults.nameRegex;
    	      this.owner = defaults.owner;
    	      this.properties = defaults.properties;
    	      this.region = defaults.region;
    	      this.sizeMax = defaults.sizeMax;
    	      this.sizeMin = defaults.sizeMin;
    	      this.sort = defaults.sort;
    	      this.tag = defaults.tag;
    	      this.tags = defaults.tags;
    	      this.visibility = defaults.visibility;
        }

        @CustomType.Setter
        public Builder containerFormat(@Nullable String containerFormat) {

            this.containerFormat = containerFormat;
            return this;
        }
        @CustomType.Setter
        public Builder diskFormat(@Nullable String diskFormat) {

            this.diskFormat = diskFormat;
            return this;
        }
        @CustomType.Setter
        public Builder hidden(@Nullable Boolean hidden) {

            this.hidden = hidden;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetImageIdsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ids(List<String> ids) {
            if (ids == null) {
              throw new MissingRequiredPropertyException("GetImageIdsResult", "ids");
            }
            this.ids = ids;
            return this;
        }
        public Builder ids(String... ids) {
            return ids(List.of(ids));
        }
        @CustomType.Setter
        public Builder memberStatus(@Nullable String memberStatus) {

            this.memberStatus = memberStatus;
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
        public Builder owner(@Nullable String owner) {

            this.owner = owner;
            return this;
        }
        @CustomType.Setter
        public Builder properties(@Nullable Map<String,String> properties) {

            this.properties = properties;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetImageIdsResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder sizeMax(@Nullable Integer sizeMax) {

            this.sizeMax = sizeMax;
            return this;
        }
        @CustomType.Setter
        public Builder sizeMin(@Nullable Integer sizeMin) {

            this.sizeMin = sizeMin;
            return this;
        }
        @CustomType.Setter
        public Builder sort(@Nullable String sort) {

            this.sort = sort;
            return this;
        }
        @CustomType.Setter
        public Builder tag(@Nullable String tag) {

            this.tag = tag;
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
        public Builder visibility(@Nullable String visibility) {

            this.visibility = visibility;
            return this;
        }
        public GetImageIdsResult build() {
            final var _resultValue = new GetImageIdsResult();
            _resultValue.containerFormat = containerFormat;
            _resultValue.diskFormat = diskFormat;
            _resultValue.hidden = hidden;
            _resultValue.id = id;
            _resultValue.ids = ids;
            _resultValue.memberStatus = memberStatus;
            _resultValue.name = name;
            _resultValue.nameRegex = nameRegex;
            _resultValue.owner = owner;
            _resultValue.properties = properties;
            _resultValue.region = region;
            _resultValue.sizeMax = sizeMax;
            _resultValue.sizeMin = sizeMin;
            _resultValue.sort = sort;
            _resultValue.tag = tag;
            _resultValue.tags = tags;
            _resultValue.visibility = visibility;
            return _resultValue;
        }
    }
}
