// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MembersMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final MembersMemberArgs Empty = new MembersMemberArgs();

    /**
     * The IP address of the members to receive traffic from
     * the load balancer.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return The IP address of the members to receive traffic from
     * the load balancer.
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
     * A bool that indicates whether the member is
     * backup. **Requires octavia minor version 2.1 or later**.
     * 
     */
    @Import(name="backup")
    private @Nullable Output<Boolean> backup;

    /**
     * @return A bool that indicates whether the member is
     * backup. **Requires octavia minor version 2.1 or later**.
     * 
     */
    public Optional<Output<Boolean>> backup() {
        return Optional.ofNullable(this.backup);
    }

    /**
     * The unique ID for the members.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return The unique ID for the members.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * An alternate IP address used for health
     * monitoring a backend member.
     * 
     */
    @Import(name="monitorAddress")
    private @Nullable Output<String> monitorAddress;

    /**
     * @return An alternate IP address used for health
     * monitoring a backend member.
     * 
     */
    public Optional<Output<String>> monitorAddress() {
        return Optional.ofNullable(this.monitorAddress);
    }

    /**
     * An alternate protocol port used for health
     * monitoring a backend member.
     * 
     */
    @Import(name="monitorPort")
    private @Nullable Output<Integer> monitorPort;

    /**
     * @return An alternate protocol port used for health
     * monitoring a backend member.
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
     * The port on which to listen for client traffic.
     * 
     */
    @Import(name="protocolPort", required=true)
    private Output<Integer> protocolPort;

    /**
     * @return The port on which to listen for client traffic.
     * 
     */
    public Output<Integer> protocolPort() {
        return this.protocolPort;
    }

    /**
     * The subnet in which to access the member.
     * 
     */
    @Import(name="subnetId")
    private @Nullable Output<String> subnetId;

    /**
     * @return The subnet in which to access the member.
     * 
     */
    public Optional<Output<String>> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    /**
     * A positive integer value that indicates the relative
     * portion of traffic that this members should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     * 
     */
    @Import(name="weight")
    private @Nullable Output<Integer> weight;

    /**
     * @return A positive integer value that indicates the relative
     * portion of traffic that this members should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     * 
     */
    public Optional<Output<Integer>> weight() {
        return Optional.ofNullable(this.weight);
    }

    private MembersMemberArgs() {}

    private MembersMemberArgs(MembersMemberArgs $) {
        this.address = $.address;
        this.adminStateUp = $.adminStateUp;
        this.backup = $.backup;
        this.id = $.id;
        this.monitorAddress = $.monitorAddress;
        this.monitorPort = $.monitorPort;
        this.name = $.name;
        this.protocolPort = $.protocolPort;
        this.subnetId = $.subnetId;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MembersMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MembersMemberArgs $;

        public Builder() {
            $ = new MembersMemberArgs();
        }

        public Builder(MembersMemberArgs defaults) {
            $ = new MembersMemberArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address The IP address of the members to receive traffic from
         * the load balancer.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address The IP address of the members to receive traffic from
         * the load balancer.
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
         * @param backup A bool that indicates whether the member is
         * backup. **Requires octavia minor version 2.1 or later**.
         * 
         * @return builder
         * 
         */
        public Builder backup(@Nullable Output<Boolean> backup) {
            $.backup = backup;
            return this;
        }

        /**
         * @param backup A bool that indicates whether the member is
         * backup. **Requires octavia minor version 2.1 or later**.
         * 
         * @return builder
         * 
         */
        public Builder backup(Boolean backup) {
            return backup(Output.of(backup));
        }

        /**
         * @param id The unique ID for the members.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The unique ID for the members.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param monitorAddress An alternate IP address used for health
         * monitoring a backend member.
         * 
         * @return builder
         * 
         */
        public Builder monitorAddress(@Nullable Output<String> monitorAddress) {
            $.monitorAddress = monitorAddress;
            return this;
        }

        /**
         * @param monitorAddress An alternate IP address used for health
         * monitoring a backend member.
         * 
         * @return builder
         * 
         */
        public Builder monitorAddress(String monitorAddress) {
            return monitorAddress(Output.of(monitorAddress));
        }

        /**
         * @param monitorPort An alternate protocol port used for health
         * monitoring a backend member.
         * 
         * @return builder
         * 
         */
        public Builder monitorPort(@Nullable Output<Integer> monitorPort) {
            $.monitorPort = monitorPort;
            return this;
        }

        /**
         * @param monitorPort An alternate protocol port used for health
         * monitoring a backend member.
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
         * @param protocolPort The port on which to listen for client traffic.
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
         * 
         * @return builder
         * 
         */
        public Builder protocolPort(Integer protocolPort) {
            return protocolPort(Output.of(protocolPort));
        }

        /**
         * @param subnetId The subnet in which to access the member.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(@Nullable Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId The subnet in which to access the member.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        /**
         * @param weight A positive integer value that indicates the relative
         * portion of traffic that this members should receive from the pool. For
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
         * portion of traffic that this members should receive from the pool. For
         * example, a member with a weight of 10 receives five times as much traffic
         * as a member with a weight of 2. Defaults to 1.
         * 
         * @return builder
         * 
         */
        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public MembersMemberArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.protocolPort = Objects.requireNonNull($.protocolPort, "expected parameter 'protocolPort' to be non-null");
            return $;
        }
    }

}
