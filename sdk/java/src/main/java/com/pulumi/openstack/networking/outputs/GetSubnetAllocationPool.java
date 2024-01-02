// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSubnetAllocationPool {
    private String end;
    private String start;

    private GetSubnetAllocationPool() {}
    public String end() {
        return this.end;
    }
    public String start() {
        return this.start;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSubnetAllocationPool defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String end;
        private String start;
        public Builder() {}
        public Builder(GetSubnetAllocationPool defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.end = defaults.end;
    	      this.start = defaults.start;
        }

        @CustomType.Setter
        public Builder end(String end) {
            if (end == null) {
              throw new MissingRequiredPropertyException("GetSubnetAllocationPool", "end");
            }
            this.end = end;
            return this;
        }
        @CustomType.Setter
        public Builder start(String start) {
            if (start == null) {
              throw new MissingRequiredPropertyException("GetSubnetAllocationPool", "start");
            }
            this.start = start;
            return this;
        }
        public GetSubnetAllocationPool build() {
            final var _resultValue = new GetSubnetAllocationPool();
            _resultValue.end = end;
            _resultValue.start = start;
            return _resultValue;
        }
    }
}
