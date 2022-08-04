// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecGroupRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecGroupRuleArgs Empty = new SecGroupRuleArgs();

    /**
     * Required if `from_group_id` or `self` is empty. The IP range
     * that will be the source of network traffic to the security group. Use 0.0.0.0/0
     * to allow all IP addresses. Changing this creates a new security group rule. Cannot
     * be combined with `from_group_id` or `self`.
     * 
     */
    @Import(name="cidr")
    private @Nullable Output<String> cidr;

    /**
     * @return Required if `from_group_id` or `self` is empty. The IP range
     * that will be the source of network traffic to the security group. Use 0.0.0.0/0
     * to allow all IP addresses. Changing this creates a new security group rule. Cannot
     * be combined with `from_group_id` or `self`.
     * 
     */
    public Optional<Output<String>> cidr() {
        return Optional.ofNullable(this.cidr);
    }

    /**
     * Required if `cidr` or `self` is empty. The ID of a
     * group from which to forward traffic to the parent group. Changing this creates a
     * new security group rule. Cannot be combined with `cidr` or `self`.
     * 
     */
    @Import(name="fromGroupId")
    private @Nullable Output<String> fromGroupId;

    /**
     * @return Required if `cidr` or `self` is empty. The ID of a
     * group from which to forward traffic to the parent group. Changing this creates a
     * new security group rule. Cannot be combined with `cidr` or `self`.
     * 
     */
    public Optional<Output<String>> fromGroupId() {
        return Optional.ofNullable(this.fromGroupId);
    }

    /**
     * An integer representing the lower bound of the port
     * range to open. Changing this creates a new security group rule.
     * 
     */
    @Import(name="fromPort", required=true)
    private Output<Integer> fromPort;

    /**
     * @return An integer representing the lower bound of the port
     * range to open. Changing this creates a new security group rule.
     * 
     */
    public Output<Integer> fromPort() {
        return this.fromPort;
    }

    @Import(name="id")
    private @Nullable Output<String> id;

    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * The protocol type that will be allowed. Changing
     * this creates a new security group rule.
     * 
     */
    @Import(name="ipProtocol", required=true)
    private Output<String> ipProtocol;

    /**
     * @return The protocol type that will be allowed. Changing
     * this creates a new security group rule.
     * 
     */
    public Output<String> ipProtocol() {
        return this.ipProtocol;
    }

    /**
     * Required if `cidr` and `from_group_id` is empty. If true,
     * the security group itself will be added as a source to this ingress rule. Cannot
     * be combined with `cidr` or `from_group_id`.
     * 
     */
    @Import(name="self")
    private @Nullable Output<Boolean> self;

    /**
     * @return Required if `cidr` and `from_group_id` is empty. If true,
     * the security group itself will be added as a source to this ingress rule. Cannot
     * be combined with `cidr` or `from_group_id`.
     * 
     */
    public Optional<Output<Boolean>> self() {
        return Optional.ofNullable(this.self);
    }

    /**
     * An integer representing the upper bound of the port
     * range to open. Changing this creates a new security group rule.
     * 
     */
    @Import(name="toPort", required=true)
    private Output<Integer> toPort;

    /**
     * @return An integer representing the upper bound of the port
     * range to open. Changing this creates a new security group rule.
     * 
     */
    public Output<Integer> toPort() {
        return this.toPort;
    }

    private SecGroupRuleArgs() {}

    private SecGroupRuleArgs(SecGroupRuleArgs $) {
        this.cidr = $.cidr;
        this.fromGroupId = $.fromGroupId;
        this.fromPort = $.fromPort;
        this.id = $.id;
        this.ipProtocol = $.ipProtocol;
        this.self = $.self;
        this.toPort = $.toPort;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecGroupRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecGroupRuleArgs $;

        public Builder() {
            $ = new SecGroupRuleArgs();
        }

        public Builder(SecGroupRuleArgs defaults) {
            $ = new SecGroupRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidr Required if `from_group_id` or `self` is empty. The IP range
         * that will be the source of network traffic to the security group. Use 0.0.0.0/0
         * to allow all IP addresses. Changing this creates a new security group rule. Cannot
         * be combined with `from_group_id` or `self`.
         * 
         * @return builder
         * 
         */
        public Builder cidr(@Nullable Output<String> cidr) {
            $.cidr = cidr;
            return this;
        }

        /**
         * @param cidr Required if `from_group_id` or `self` is empty. The IP range
         * that will be the source of network traffic to the security group. Use 0.0.0.0/0
         * to allow all IP addresses. Changing this creates a new security group rule. Cannot
         * be combined with `from_group_id` or `self`.
         * 
         * @return builder
         * 
         */
        public Builder cidr(String cidr) {
            return cidr(Output.of(cidr));
        }

        /**
         * @param fromGroupId Required if `cidr` or `self` is empty. The ID of a
         * group from which to forward traffic to the parent group. Changing this creates a
         * new security group rule. Cannot be combined with `cidr` or `self`.
         * 
         * @return builder
         * 
         */
        public Builder fromGroupId(@Nullable Output<String> fromGroupId) {
            $.fromGroupId = fromGroupId;
            return this;
        }

        /**
         * @param fromGroupId Required if `cidr` or `self` is empty. The ID of a
         * group from which to forward traffic to the parent group. Changing this creates a
         * new security group rule. Cannot be combined with `cidr` or `self`.
         * 
         * @return builder
         * 
         */
        public Builder fromGroupId(String fromGroupId) {
            return fromGroupId(Output.of(fromGroupId));
        }

        /**
         * @param fromPort An integer representing the lower bound of the port
         * range to open. Changing this creates a new security group rule.
         * 
         * @return builder
         * 
         */
        public Builder fromPort(Output<Integer> fromPort) {
            $.fromPort = fromPort;
            return this;
        }

        /**
         * @param fromPort An integer representing the lower bound of the port
         * range to open. Changing this creates a new security group rule.
         * 
         * @return builder
         * 
         */
        public Builder fromPort(Integer fromPort) {
            return fromPort(Output.of(fromPort));
        }

        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param ipProtocol The protocol type that will be allowed. Changing
         * this creates a new security group rule.
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(Output<String> ipProtocol) {
            $.ipProtocol = ipProtocol;
            return this;
        }

        /**
         * @param ipProtocol The protocol type that will be allowed. Changing
         * this creates a new security group rule.
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(String ipProtocol) {
            return ipProtocol(Output.of(ipProtocol));
        }

        /**
         * @param self Required if `cidr` and `from_group_id` is empty. If true,
         * the security group itself will be added as a source to this ingress rule. Cannot
         * be combined with `cidr` or `from_group_id`.
         * 
         * @return builder
         * 
         */
        public Builder self(@Nullable Output<Boolean> self) {
            $.self = self;
            return this;
        }

        /**
         * @param self Required if `cidr` and `from_group_id` is empty. If true,
         * the security group itself will be added as a source to this ingress rule. Cannot
         * be combined with `cidr` or `from_group_id`.
         * 
         * @return builder
         * 
         */
        public Builder self(Boolean self) {
            return self(Output.of(self));
        }

        /**
         * @param toPort An integer representing the upper bound of the port
         * range to open. Changing this creates a new security group rule.
         * 
         * @return builder
         * 
         */
        public Builder toPort(Output<Integer> toPort) {
            $.toPort = toPort;
            return this;
        }

        /**
         * @param toPort An integer representing the upper bound of the port
         * range to open. Changing this creates a new security group rule.
         * 
         * @return builder
         * 
         */
        public Builder toPort(Integer toPort) {
            return toPort(Output.of(toPort));
        }

        public SecGroupRuleArgs build() {
            $.fromPort = Objects.requireNonNull($.fromPort, "expected parameter 'fromPort' to be non-null");
            $.ipProtocol = Objects.requireNonNull($.ipProtocol, "expected parameter 'ipProtocol' to be non-null");
            $.toPort = Objects.requireNonNull($.toPort, "expected parameter 'toPort' to be non-null");
            return $;
        }
    }

}
