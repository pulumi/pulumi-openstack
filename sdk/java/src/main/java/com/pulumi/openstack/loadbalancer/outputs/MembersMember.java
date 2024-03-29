// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class MembersMember {
    /**
     * @return The IP address of the members to receive traffic from
     * the load balancer.
     * 
     */
    private String address;
    /**
     * @return The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     * 
     */
    private @Nullable Boolean adminStateUp;
    /**
     * @return A bool that indicates whether the member is
     * backup. **Requires octavia minor version 2.1 or later**.
     * 
     */
    private @Nullable Boolean backup;
    /**
     * @return The unique ID for the members.
     * 
     */
    private @Nullable String id;
    /**
     * @return An alternate IP address used for health
     * monitoring a backend member.
     * 
     */
    private @Nullable String monitorAddress;
    /**
     * @return An alternate protocol port used for health
     * monitoring a backend member.
     * 
     */
    private @Nullable Integer monitorPort;
    /**
     * @return Human-readable name for the member.
     * 
     */
    private @Nullable String name;
    /**
     * @return The port on which to listen for client traffic.
     * 
     */
    private Integer protocolPort;
    /**
     * @return The subnet in which to access the member.
     * 
     */
    private @Nullable String subnetId;
    /**
     * @return A positive integer value that indicates the relative
     * portion of traffic that this members should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     * 
     */
    private @Nullable Integer weight;

    private MembersMember() {}
    /**
     * @return The IP address of the members to receive traffic from
     * the load balancer.
     * 
     */
    public String address() {
        return this.address;
    }
    /**
     * @return The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     * 
     */
    public Optional<Boolean> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }
    /**
     * @return A bool that indicates whether the member is
     * backup. **Requires octavia minor version 2.1 or later**.
     * 
     */
    public Optional<Boolean> backup() {
        return Optional.ofNullable(this.backup);
    }
    /**
     * @return The unique ID for the members.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return An alternate IP address used for health
     * monitoring a backend member.
     * 
     */
    public Optional<String> monitorAddress() {
        return Optional.ofNullable(this.monitorAddress);
    }
    /**
     * @return An alternate protocol port used for health
     * monitoring a backend member.
     * 
     */
    public Optional<Integer> monitorPort() {
        return Optional.ofNullable(this.monitorPort);
    }
    /**
     * @return Human-readable name for the member.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The port on which to listen for client traffic.
     * 
     */
    public Integer protocolPort() {
        return this.protocolPort;
    }
    /**
     * @return The subnet in which to access the member.
     * 
     */
    public Optional<String> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }
    /**
     * @return A positive integer value that indicates the relative
     * portion of traffic that this members should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     * 
     */
    public Optional<Integer> weight() {
        return Optional.ofNullable(this.weight);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MembersMember defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private @Nullable Boolean adminStateUp;
        private @Nullable Boolean backup;
        private @Nullable String id;
        private @Nullable String monitorAddress;
        private @Nullable Integer monitorPort;
        private @Nullable String name;
        private Integer protocolPort;
        private @Nullable String subnetId;
        private @Nullable Integer weight;
        public Builder() {}
        public Builder(MembersMember defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.adminStateUp = defaults.adminStateUp;
    	      this.backup = defaults.backup;
    	      this.id = defaults.id;
    	      this.monitorAddress = defaults.monitorAddress;
    	      this.monitorPort = defaults.monitorPort;
    	      this.name = defaults.name;
    	      this.protocolPort = defaults.protocolPort;
    	      this.subnetId = defaults.subnetId;
    	      this.weight = defaults.weight;
        }

        @CustomType.Setter
        public Builder address(String address) {
            if (address == null) {
              throw new MissingRequiredPropertyException("MembersMember", "address");
            }
            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder adminStateUp(@Nullable Boolean adminStateUp) {

            this.adminStateUp = adminStateUp;
            return this;
        }
        @CustomType.Setter
        public Builder backup(@Nullable Boolean backup) {

            this.backup = backup;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder monitorAddress(@Nullable String monitorAddress) {

            this.monitorAddress = monitorAddress;
            return this;
        }
        @CustomType.Setter
        public Builder monitorPort(@Nullable Integer monitorPort) {

            this.monitorPort = monitorPort;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder protocolPort(Integer protocolPort) {
            if (protocolPort == null) {
              throw new MissingRequiredPropertyException("MembersMember", "protocolPort");
            }
            this.protocolPort = protocolPort;
            return this;
        }
        @CustomType.Setter
        public Builder subnetId(@Nullable String subnetId) {

            this.subnetId = subnetId;
            return this;
        }
        @CustomType.Setter
        public Builder weight(@Nullable Integer weight) {

            this.weight = weight;
            return this;
        }
        public MembersMember build() {
            final var _resultValue = new MembersMember();
            _resultValue.address = address;
            _resultValue.adminStateUp = adminStateUp;
            _resultValue.backup = backup;
            _resultValue.id = id;
            _resultValue.monitorAddress = monitorAddress;
            _resultValue.monitorPort = monitorPort;
            _resultValue.name = name;
            _resultValue.protocolPort = protocolPort;
            _resultValue.subnetId = subnetId;
            _resultValue.weight = weight;
            return _resultValue;
        }
    }
}
