// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetFlavorResult {
    private @Nullable String description;
    private @Nullable Integer disk;
    /**
     * @return Key/Value pairs of metadata for the flavor.
     * 
     */
    private Map<String,String> extraSpecs;
    private @Nullable String flavorId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Boolean isPublic;
    private @Nullable Integer minDisk;
    private @Nullable Integer minRam;
    private @Nullable String name;
    private @Nullable Integer ram;
    private String region;
    private @Nullable Double rxTxFactor;
    private @Nullable Integer swap;
    private @Nullable Integer vcpus;

    private GetFlavorResult() {}
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public Optional<Integer> disk() {
        return Optional.ofNullable(this.disk);
    }
    /**
     * @return Key/Value pairs of metadata for the flavor.
     * 
     */
    public Map<String,String> extraSpecs() {
        return this.extraSpecs;
    }
    public Optional<String> flavorId() {
        return Optional.ofNullable(this.flavorId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Boolean> isPublic() {
        return Optional.ofNullable(this.isPublic);
    }
    public Optional<Integer> minDisk() {
        return Optional.ofNullable(this.minDisk);
    }
    public Optional<Integer> minRam() {
        return Optional.ofNullable(this.minRam);
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<Integer> ram() {
        return Optional.ofNullable(this.ram);
    }
    public String region() {
        return this.region;
    }
    public Optional<Double> rxTxFactor() {
        return Optional.ofNullable(this.rxTxFactor);
    }
    public Optional<Integer> swap() {
        return Optional.ofNullable(this.swap);
    }
    public Optional<Integer> vcpus() {
        return Optional.ofNullable(this.vcpus);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFlavorResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable Integer disk;
        private Map<String,String> extraSpecs;
        private @Nullable String flavorId;
        private String id;
        private @Nullable Boolean isPublic;
        private @Nullable Integer minDisk;
        private @Nullable Integer minRam;
        private @Nullable String name;
        private @Nullable Integer ram;
        private String region;
        private @Nullable Double rxTxFactor;
        private @Nullable Integer swap;
        private @Nullable Integer vcpus;
        public Builder() {}
        public Builder(GetFlavorResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.disk = defaults.disk;
    	      this.extraSpecs = defaults.extraSpecs;
    	      this.flavorId = defaults.flavorId;
    	      this.id = defaults.id;
    	      this.isPublic = defaults.isPublic;
    	      this.minDisk = defaults.minDisk;
    	      this.minRam = defaults.minRam;
    	      this.name = defaults.name;
    	      this.ram = defaults.ram;
    	      this.region = defaults.region;
    	      this.rxTxFactor = defaults.rxTxFactor;
    	      this.swap = defaults.swap;
    	      this.vcpus = defaults.vcpus;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder disk(@Nullable Integer disk) {

            this.disk = disk;
            return this;
        }
        @CustomType.Setter
        public Builder extraSpecs(Map<String,String> extraSpecs) {
            if (extraSpecs == null) {
              throw new MissingRequiredPropertyException("GetFlavorResult", "extraSpecs");
            }
            this.extraSpecs = extraSpecs;
            return this;
        }
        @CustomType.Setter
        public Builder flavorId(@Nullable String flavorId) {

            this.flavorId = flavorId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetFlavorResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isPublic(@Nullable Boolean isPublic) {

            this.isPublic = isPublic;
            return this;
        }
        @CustomType.Setter
        public Builder minDisk(@Nullable Integer minDisk) {

            this.minDisk = minDisk;
            return this;
        }
        @CustomType.Setter
        public Builder minRam(@Nullable Integer minRam) {

            this.minRam = minRam;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder ram(@Nullable Integer ram) {

            this.ram = ram;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetFlavorResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder rxTxFactor(@Nullable Double rxTxFactor) {

            this.rxTxFactor = rxTxFactor;
            return this;
        }
        @CustomType.Setter
        public Builder swap(@Nullable Integer swap) {

            this.swap = swap;
            return this;
        }
        @CustomType.Setter
        public Builder vcpus(@Nullable Integer vcpus) {

            this.vcpus = vcpus;
            return this;
        }
        public GetFlavorResult build() {
            final var _resultValue = new GetFlavorResult();
            _resultValue.description = description;
            _resultValue.disk = disk;
            _resultValue.extraSpecs = extraSpecs;
            _resultValue.flavorId = flavorId;
            _resultValue.id = id;
            _resultValue.isPublic = isPublic;
            _resultValue.minDisk = minDisk;
            _resultValue.minRam = minRam;
            _resultValue.name = name;
            _resultValue.ram = ram;
            _resultValue.region = region;
            _resultValue.rxTxFactor = rxTxFactor;
            _resultValue.swap = swap;
            _resultValue.vcpus = vcpus;
            return _resultValue;
        }
    }
}
