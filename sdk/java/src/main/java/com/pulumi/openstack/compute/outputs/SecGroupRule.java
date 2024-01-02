// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecGroupRule {
    /**
     * @return Required if `from_group_id` or `self` is empty. The IP range
     * that will be the source of network traffic to the security group. Use 0.0.0.0/0
     * to allow all IP addresses. Changing this creates a new security group rule. Cannot
     * be combined with `from_group_id` or `self`.
     * 
     */
    private @Nullable String cidr;
    /**
     * @return Required if `cidr` or `self` is empty. The ID of a
     * group from which to forward traffic to the parent group. Changing this creates a
     * new security group rule. Cannot be combined with `cidr` or `self`.
     * 
     */
    private @Nullable String fromGroupId;
    /**
     * @return An integer representing the lower bound of the port
     * range to open. Changing this creates a new security group rule.
     * 
     */
    private Integer fromPort;
    private @Nullable String id;
    /**
     * @return The protocol type that will be allowed. Changing
     * this creates a new security group rule.
     * 
     */
    private String ipProtocol;
    /**
     * @return Required if `cidr` and `from_group_id` is empty. If true,
     * the security group itself will be added as a source to this ingress rule. Cannot
     * be combined with `cidr` or `from_group_id`.
     * 
     */
    private @Nullable Boolean self;
    /**
     * @return An integer representing the upper bound of the port
     * range to open. Changing this creates a new security group rule.
     * 
     */
    private Integer toPort;

    private SecGroupRule() {}
    /**
     * @return Required if `from_group_id` or `self` is empty. The IP range
     * that will be the source of network traffic to the security group. Use 0.0.0.0/0
     * to allow all IP addresses. Changing this creates a new security group rule. Cannot
     * be combined with `from_group_id` or `self`.
     * 
     */
    public Optional<String> cidr() {
        return Optional.ofNullable(this.cidr);
    }
    /**
     * @return Required if `cidr` or `self` is empty. The ID of a
     * group from which to forward traffic to the parent group. Changing this creates a
     * new security group rule. Cannot be combined with `cidr` or `self`.
     * 
     */
    public Optional<String> fromGroupId() {
        return Optional.ofNullable(this.fromGroupId);
    }
    /**
     * @return An integer representing the lower bound of the port
     * range to open. Changing this creates a new security group rule.
     * 
     */
    public Integer fromPort() {
        return this.fromPort;
    }
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The protocol type that will be allowed. Changing
     * this creates a new security group rule.
     * 
     */
    public String ipProtocol() {
        return this.ipProtocol;
    }
    /**
     * @return Required if `cidr` and `from_group_id` is empty. If true,
     * the security group itself will be added as a source to this ingress rule. Cannot
     * be combined with `cidr` or `from_group_id`.
     * 
     */
    public Optional<Boolean> self() {
        return Optional.ofNullable(this.self);
    }
    /**
     * @return An integer representing the upper bound of the port
     * range to open. Changing this creates a new security group rule.
     * 
     */
    public Integer toPort() {
        return this.toPort;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecGroupRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cidr;
        private @Nullable String fromGroupId;
        private Integer fromPort;
        private @Nullable String id;
        private String ipProtocol;
        private @Nullable Boolean self;
        private Integer toPort;
        public Builder() {}
        public Builder(SecGroupRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cidr = defaults.cidr;
    	      this.fromGroupId = defaults.fromGroupId;
    	      this.fromPort = defaults.fromPort;
    	      this.id = defaults.id;
    	      this.ipProtocol = defaults.ipProtocol;
    	      this.self = defaults.self;
    	      this.toPort = defaults.toPort;
        }

        @CustomType.Setter
        public Builder cidr(@Nullable String cidr) {

            this.cidr = cidr;
            return this;
        }
        @CustomType.Setter
        public Builder fromGroupId(@Nullable String fromGroupId) {

            this.fromGroupId = fromGroupId;
            return this;
        }
        @CustomType.Setter
        public Builder fromPort(Integer fromPort) {
            if (fromPort == null) {
              throw new MissingRequiredPropertyException("SecGroupRule", "fromPort");
            }
            this.fromPort = fromPort;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ipProtocol(String ipProtocol) {
            if (ipProtocol == null) {
              throw new MissingRequiredPropertyException("SecGroupRule", "ipProtocol");
            }
            this.ipProtocol = ipProtocol;
            return this;
        }
        @CustomType.Setter
        public Builder self(@Nullable Boolean self) {

            this.self = self;
            return this;
        }
        @CustomType.Setter
        public Builder toPort(Integer toPort) {
            if (toPort == null) {
              throw new MissingRequiredPropertyException("SecGroupRule", "toPort");
            }
            this.toPort = toPort;
            return this;
        }
        public SecGroupRule build() {
            final var _resultValue = new SecGroupRule();
            _resultValue.cidr = cidr;
            _resultValue.fromGroupId = fromGroupId;
            _resultValue.fromPort = fromPort;
            _resultValue.id = id;
            _resultValue.ipProtocol = ipProtocol;
            _resultValue.self = self;
            _resultValue.toPort = toPort;
            return _resultValue;
        }
    }
}
