// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAvailabilityZonesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return The names of the availability zones, ordered alphanumerically, that match the queried `state`
     * 
     */
    private final List<String> names;
    private final String region;
    private final @Nullable String state;

    @CustomType.Constructor
    private GetAvailabilityZonesResult(
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("names") List<String> names,
        @CustomType.Parameter("region") String region,
        @CustomType.Parameter("state") @Nullable String state) {
        this.id = id;
        this.names = names;
        this.region = region;
        this.state = state;
    }

    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The names of the availability zones, ordered alphanumerically, that match the queried `state`
     * 
     */
    public List<String> names() {
        return this.names;
    }
    public String region() {
        return this.region;
    }
    public Optional<String> state() {
        return Optional.ofNullable(this.state);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAvailabilityZonesResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String id;
        private List<String> names;
        private String region;
        private @Nullable String state;

        public Builder() {
    	      // Empty
        }

        public Builder(GetAvailabilityZonesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.names = defaults.names;
    	      this.region = defaults.region;
    	      this.state = defaults.state;
        }

        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder names(List<String> names) {
            this.names = Objects.requireNonNull(names);
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        public Builder state(@Nullable String state) {
            this.state = state;
            return this;
        }        public GetAvailabilityZonesResult build() {
            return new GetAvailabilityZonesResult(id, names, region, state);
        }
    }
}
