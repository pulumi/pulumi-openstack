// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetPortBinding {
    /**
     * @return The ID of the host, which has the allocatee port.
     * 
     */
    private String hostId;
    /**
     * @return A JSON string containing the binding profile information.
     * 
     */
    private String profile;
    /**
     * @return A map of JSON strings containing additional details for this
     * specific binding.
     * 
     */
    private Map<String,Object> vifDetails;
    /**
     * @return The VNIC type of the port binding.
     * 
     */
    private String vifType;
    /**
     * @return VNIC type for the port.
     * 
     */
    private String vnicType;

    private GetPortBinding() {}
    /**
     * @return The ID of the host, which has the allocatee port.
     * 
     */
    public String hostId() {
        return this.hostId;
    }
    /**
     * @return A JSON string containing the binding profile information.
     * 
     */
    public String profile() {
        return this.profile;
    }
    /**
     * @return A map of JSON strings containing additional details for this
     * specific binding.
     * 
     */
    public Map<String,Object> vifDetails() {
        return this.vifDetails;
    }
    /**
     * @return The VNIC type of the port binding.
     * 
     */
    public String vifType() {
        return this.vifType;
    }
    /**
     * @return VNIC type for the port.
     * 
     */
    public String vnicType() {
        return this.vnicType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPortBinding defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String hostId;
        private String profile;
        private Map<String,Object> vifDetails;
        private String vifType;
        private String vnicType;
        public Builder() {}
        public Builder(GetPortBinding defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.hostId = defaults.hostId;
    	      this.profile = defaults.profile;
    	      this.vifDetails = defaults.vifDetails;
    	      this.vifType = defaults.vifType;
    	      this.vnicType = defaults.vnicType;
        }

        @CustomType.Setter
        public Builder hostId(String hostId) {
            this.hostId = Objects.requireNonNull(hostId);
            return this;
        }
        @CustomType.Setter
        public Builder profile(String profile) {
            this.profile = Objects.requireNonNull(profile);
            return this;
        }
        @CustomType.Setter
        public Builder vifDetails(Map<String,Object> vifDetails) {
            this.vifDetails = Objects.requireNonNull(vifDetails);
            return this;
        }
        @CustomType.Setter
        public Builder vifType(String vifType) {
            this.vifType = Objects.requireNonNull(vifType);
            return this;
        }
        @CustomType.Setter
        public Builder vnicType(String vnicType) {
            this.vnicType = Objects.requireNonNull(vnicType);
            return this;
        }
        public GetPortBinding build() {
            final var o = new GetPortBinding();
            o.hostId = hostId;
            o.profile = profile;
            o.vifDetails = vifDetails;
            o.vifType = vifType;
            o.vnicType = vnicType;
            return o;
        }
    }
}
