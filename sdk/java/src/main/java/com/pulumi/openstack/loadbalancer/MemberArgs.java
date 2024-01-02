// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final MemberArgs Empty = new MemberArgs();

    /**
     * The IP address of the member to receive traffic from
     * the load balancer. Changing this creates a new member.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return The IP address of the member to receive traffic from
     * the load balancer. Changing this creates a new member.
     * 
     */
    public Output<String> address() {
        return this.address;
    }

    /**
     * The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * Boolean that indicates whether that member works as a backup or not. Available
     * only for Octavia &gt;= 2.1.
     * 
     */
    @Import(name="backup")
    private @Nullable Output<Boolean> backup;

    /**
     * @return Boolean that indicates whether that member works as a backup or not. Available
     * only for Octavia &gt;= 2.1.
     * 
     */
    public Optional<Output<Boolean>> backup() {
        return Optional.ofNullable(this.backup);
    }

    /**
     * An alternate IP address used for health monitoring a backend member.
     * Available only for Octavia
     * 
     */
    @Import(name="monitorAddress")
    private @Nullable Output<String> monitorAddress;

    /**
     * @return An alternate IP address used for health monitoring a backend member.
     * Available only for Octavia
     * 
     */
    public Optional<Output<String>> monitorAddress() {
        return Optional.ofNullable(this.monitorAddress);
    }

    /**
     * An alternate protocol port used for health monitoring a backend member.
     * Available only for Octavia
     * 
     */
    @Import(name="monitorPort")
    private @Nullable Output<Integer> monitorPort;

    /**
     * @return An alternate protocol port used for health monitoring a backend member.
     * Available only for Octavia
     * 
     */
    public Optional<Output<Integer>> monitorPort() {
        return Optional.ofNullable(this.monitorPort);
    }

    /**
     * Human-readable name for the member.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Human-readable name for the member.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The id of the pool that this member will be assigned
     * to. Changing this creates a new member.
     * 
     */
    @Import(name="poolId", required=true)
    private Output<String> poolId;

    /**
     * @return The id of the pool that this member will be assigned
     * to. Changing this creates a new member.
     * 
     */
    public Output<String> poolId() {
        return this.poolId;
    }

    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new member.
     * 
     */
    @Import(name="protocolPort", required=true)
    private Output<Integer> protocolPort;

    /**
     * @return The port on which to listen for client traffic.
     * Changing this creates a new member.
     * 
     */
    public Output<Integer> protocolPort() {
        return this.protocolPort;
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a member. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new member.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a member. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new member.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The subnet in which to access the member. Changing
     * this creates a new member.
     * 
     */
    @Import(name="subnetId")
    private @Nullable Output<String> subnetId;

    /**
     * @return The subnet in which to access the member. Changing
     * this creates a new member.
     * 
     */
    public Optional<Output<String>> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    /**
     * Required for admins. The UUID of the tenant who owns
     * the member.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new member.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the member.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new member.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * A positive integer value that indicates the relative
     * portion of traffic that this member should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    /**
     * @return A positive integer value that indicates the relative
     * portion of traffic that this member should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     * 
     */
    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private MemberArgs() {}

    private MemberArgs(MemberArgs $) {
        this.address = $.address;
        this.adminStateUp = $.adminStateUp;
        this.backup = $.backup;
        this.monitorAddress = $.monitorAddress;
        this.monitorPort = $.monitorPort;
        this.name = $.name;
        this.poolId = $.poolId;
        this.protocolPort = $.protocolPort;
        this.region = $.region;
        this.subnetId = $.subnetId;
        this.tenantId = $.tenantId;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MemberArgs $;

        public Builder() {
            $ = new MemberArgs();
        }

        public Builder(MemberArgs defaults) {
            $ = new MemberArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The IP address of the member to receive traffic from
         * the load balancer. Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address The IP address of the member to receive traffic from
         * the load balancer. Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param adminStateUp The administrative state of the member.
         * A valid value is true (UP) or false (DOWN). Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp The administrative state of the member.
         * A valid value is true (UP) or false (DOWN). Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param backup Boolean that indicates whether that member works as a backup or not. Available
         * only for Octavia &gt;= 2.1.
         * 
         * @return builder
         * 
         */
        public Builder backup(@Nullable Output<Boolean> backup) {
            $.backup = backup;
            return this;
        }

        /**
         * @param backup Boolean that indicates whether that member works as a backup or not. Available
         * only for Octavia &gt;= 2.1.
         * 
         * @return builder
         * 
         */
        public Builder backup(Boolean backup) {
            return backup(Output.of(backup));
        }

        /**
         * @param monitorAddress An alternate IP address used for health monitoring a backend member.
         * Available only for Octavia
         * 
         * @return builder
         * 
         */
        public Builder monitorAddress(@Nullable Output<String> monitorAddress) {
            $.monitorAddress = monitorAddress;
            return this;
        }

        /**
         * @param monitorAddress An alternate IP address used for health monitoring a backend member.
         * Available only for Octavia
         * 
         * @return builder
         * 
         */
        public Builder monitorAddress(String monitorAddress) {
            return monitorAddress(Output.of(monitorAddress));
        }

        /**
         * @param monitorPort An alternate protocol port used for health monitoring a backend member.
         * Available only for Octavia
         * 
         * @return builder
         * 
         */
        public Builder monitorPort(@Nullable Output<Integer> monitorPort) {
            $.monitorPort = monitorPort;
            return this;
        }

        /**
         * @param monitorPort An alternate protocol port used for health monitoring a backend member.
         * Available only for Octavia
         * 
         * @return builder
         * 
         */
        public Builder monitorPort(Integer monitorPort) {
            return monitorPort(Output.of(monitorPort));
        }

        /**
         * @param name Human-readable name for the member.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Human-readable name for the member.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param poolId The id of the pool that this member will be assigned
         * to. Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder poolId(Output<String> poolId) {
            $.poolId = poolId;
            return this;
        }

        /**
         * @param poolId The id of the pool that this member will be assigned
         * to. Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder poolId(String poolId) {
            return poolId(Output.of(poolId));
        }

        /**
         * @param protocolPort The port on which to listen for client traffic.
         * Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder protocolPort(Output<Integer> protocolPort) {
            $.protocolPort = protocolPort;
            return this;
        }

        /**
         * @param protocolPort The port on which to listen for client traffic.
         * Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder protocolPort(Integer protocolPort) {
            return protocolPort(Output.of(protocolPort));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a member. If omitted, the `region`
         * argument of the provider is used. Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a member. If omitted, the `region`
         * argument of the provider is used. Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param subnetId The subnet in which to access the member. Changing
         * this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(@Nullable Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId The subnet in which to access the member. Changing
         * this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        /**
         * @param tenantId Required for admins. The UUID of the tenant who owns
         * the member.  Only administrative users can specify a tenant UUID
         * other than their own. Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId Required for admins. The UUID of the tenant who owns
         * the member.  Only administrative users can specify a tenant UUID
         * other than their own. Changing this creates a new member.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param weight A positive integer value that indicates the relative
         * portion of traffic that this member should receive from the pool. For
         * example, a member with a weight of 10 receives five times as much traffic
         * as a member with a weight of 2. Defaults to 1.
         * 
         * @return builder
         * 
         */
        public Builder weight(@Nullable Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        /**
         * @param weight A positive integer value that indicates the relative
         * portion of traffic that this member should receive from the pool. For
         * example, a member with a weight of 10 receives five times as much traffic
         * as a member with a weight of 2. Defaults to 1.
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public MemberArgs build() {
            if ($.address == null) {
                throw new MissingRequiredPropertyException("MemberArgs", "address");
            }
            if ($.poolId == null) {
                throw new MissingRequiredPropertyException("MemberArgs", "poolId");
            }
            if ($.protocolPort == null) {
                throw new MissingRequiredPropertyException("MemberArgs", "protocolPort");
            }
            return $;
        }
    }

}
