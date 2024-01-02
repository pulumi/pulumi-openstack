// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetFwGroupV2Result {
    /**
     * @return See Argument Reference above.
     * 
     */
    private Boolean adminStateUp;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String description;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String egressFirewallPolicyId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String groupId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String ingressFirewallPolicyId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private @Nullable String name;
    /**
     * @return Ports associated with the firewall group.
     * 
     */
    private List<String> ports;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String projectId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String region;
    /**
     * @return See Argument Reference above.
     * 
     */
    private Boolean shared;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String status;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String tenantId;

    private GetFwGroupV2Result() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public Boolean adminStateUp() {
        return this.adminStateUp;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> egressFirewallPolicyId() {
        return Optional.ofNullable(this.egressFirewallPolicyId);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> groupId() {
        return Optional.ofNullable(this.groupId);
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
    public Optional<String> ingressFirewallPolicyId() {
        return Optional.ofNullable(this.ingressFirewallPolicyId);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Ports associated with the firewall group.
     * 
     */
    public List<String> ports() {
        return this.ports;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
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
     * @return See Argument Reference above.
     * 
     */
    public Boolean shared() {
        return this.shared;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String tenantId() {
        return this.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFwGroupV2Result defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean adminStateUp;
        private @Nullable String description;
        private @Nullable String egressFirewallPolicyId;
        private @Nullable String groupId;
        private String id;
        private @Nullable String ingressFirewallPolicyId;
        private @Nullable String name;
        private List<String> ports;
        private String projectId;
        private String region;
        private Boolean shared;
        private String status;
        private String tenantId;
        public Builder() {}
        public Builder(GetFwGroupV2Result defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminStateUp = defaults.adminStateUp;
    	      this.description = defaults.description;
    	      this.egressFirewallPolicyId = defaults.egressFirewallPolicyId;
    	      this.groupId = defaults.groupId;
    	      this.id = defaults.id;
    	      this.ingressFirewallPolicyId = defaults.ingressFirewallPolicyId;
    	      this.name = defaults.name;
    	      this.ports = defaults.ports;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.shared = defaults.shared;
    	      this.status = defaults.status;
    	      this.tenantId = defaults.tenantId;
        }

        @CustomType.Setter
        public Builder adminStateUp(Boolean adminStateUp) {
            if (adminStateUp == null) {
              throw new MissingRequiredPropertyException("GetFwGroupV2Result", "adminStateUp");
            }
            this.adminStateUp = adminStateUp;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder egressFirewallPolicyId(@Nullable String egressFirewallPolicyId) {

            this.egressFirewallPolicyId = egressFirewallPolicyId;
            return this;
        }
        @CustomType.Setter
        public Builder groupId(@Nullable String groupId) {

            this.groupId = groupId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetFwGroupV2Result", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ingressFirewallPolicyId(@Nullable String ingressFirewallPolicyId) {

            this.ingressFirewallPolicyId = ingressFirewallPolicyId;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder ports(List<String> ports) {
            if (ports == null) {
              throw new MissingRequiredPropertyException("GetFwGroupV2Result", "ports");
            }
            this.ports = ports;
            return this;
        }
        public Builder ports(String... ports) {
            return ports(List.of(ports));
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetFwGroupV2Result", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetFwGroupV2Result", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder shared(Boolean shared) {
            if (shared == null) {
              throw new MissingRequiredPropertyException("GetFwGroupV2Result", "shared");
            }
            this.shared = shared;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetFwGroupV2Result", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tenantId(String tenantId) {
            if (tenantId == null) {
              throw new MissingRequiredPropertyException("GetFwGroupV2Result", "tenantId");
            }
            this.tenantId = tenantId;
            return this;
        }
        public GetFwGroupV2Result build() {
            final var _resultValue = new GetFwGroupV2Result();
            _resultValue.adminStateUp = adminStateUp;
            _resultValue.description = description;
            _resultValue.egressFirewallPolicyId = egressFirewallPolicyId;
            _resultValue.groupId = groupId;
            _resultValue.id = id;
            _resultValue.ingressFirewallPolicyId = ingressFirewallPolicyId;
            _resultValue.name = name;
            _resultValue.ports = ports;
            _resultValue.projectId = projectId;
            _resultValue.region = region;
            _resultValue.shared = shared;
            _resultValue.status = status;
            _resultValue.tenantId = tenantId;
            return _resultValue;
        }
    }
}
