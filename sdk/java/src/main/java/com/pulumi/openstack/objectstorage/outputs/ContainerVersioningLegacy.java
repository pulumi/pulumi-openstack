// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.objectstorage.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ContainerVersioningLegacy {
    /**
     * @return Container in which versions will be stored.
     * 
     */
    private String location;
    /**
     * @return Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
     * 
     */
    private String type;

    private ContainerVersioningLegacy() {}
    /**
     * @return Container in which versions will be stored.
     * 
     */
    public String location() {
        return this.location;
    }
    /**
     * @return Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerVersioningLegacy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String location;
        private String type;
        public Builder() {}
        public Builder(ContainerVersioningLegacy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.location = defaults.location;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder location(String location) {
            if (location == null) {
              throw new MissingRequiredPropertyException("ContainerVersioningLegacy", "location");
            }
            this.location = location;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("ContainerVersioningLegacy", "type");
            }
            this.type = type;
            return this;
        }
        public ContainerVersioningLegacy build() {
            final var _resultValue = new ContainerVersioningLegacy();
            _resultValue.location = location;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
