// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class SubnetAllocationPool {
    /**
     * @return The ending address.
     * 
     */
    private String end;
    /**
     * @return The starting address.
     * 
     */
    private String start;

    private SubnetAllocationPool() {}
    /**
     * @return The ending address.
     * 
     */
    public String end() {
        return this.end;
    }
    /**
     * @return The starting address.
     * 
     */
    public String start() {
        return this.start;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SubnetAllocationPool defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String end;
        private String start;
        public Builder() {}
        public Builder(SubnetAllocationPool defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.end = defaults.end;
    	      this.start = defaults.start;
        }

        @CustomType.Setter
        public Builder end(String end) {
            this.end = Objects.requireNonNull(end);
            return this;
        }
        @CustomType.Setter
        public Builder start(String start) {
            this.start = Objects.requireNonNull(start);
            return this;
        }
        public SubnetAllocationPool build() {
            final var o = new SubnetAllocationPool();
            o.end = end;
            o.start = start;
            return o;
        }
    }
}
