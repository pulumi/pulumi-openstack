// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QosMinimumBandwidthRuleState extends com.pulumi.resources.ResourceArgs {

    public static final QosMinimumBandwidthRuleState Empty = new QosMinimumBandwidthRuleState();

    /**
     * The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
     * existing QoS minimum bandwidth rule.
     * 
     */
    @Import(name="direction")
    private @Nullable Output<String> direction;

    /**
     * @return The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
     * existing QoS minimum bandwidth rule.
     * 
     */
    public Optional<Output<String>> direction() {
        return Optional.ofNullable(this.direction);
    }

    /**
     * The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     * 
     */
    @Import(name="minKbps")
    private @Nullable Output<Integer> minKbps;

    /**
     * @return The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     * 
     */
    public Optional<Output<Integer>> minKbps() {
        return Optional.ofNullable(this.minKbps);
    }

    /**
     * The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     * 
     */
    @Import(name="qosPolicyId")
    private @Nullable Output<String> qosPolicyId;

    /**
     * @return The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     * 
     */
    public Optional<Output<String>> qosPolicyId() {
        return Optional.ofNullable(this.qosPolicyId);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private QosMinimumBandwidthRuleState() {}

    private QosMinimumBandwidthRuleState(QosMinimumBandwidthRuleState $) {
        this.direction = $.direction;
        this.minKbps = $.minKbps;
        this.qosPolicyId = $.qosPolicyId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QosMinimumBandwidthRuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QosMinimumBandwidthRuleState $;

        public Builder() {
            $ = new QosMinimumBandwidthRuleState();
        }

        public Builder(QosMinimumBandwidthRuleState defaults) {
            $ = new QosMinimumBandwidthRuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param direction The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
         * existing QoS minimum bandwidth rule.
         * 
         * @return builder
         * 
         */
        public Builder direction(@Nullable Output<String> direction) {
            $.direction = direction;
            return this;
        }

        /**
         * @param direction The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
         * existing QoS minimum bandwidth rule.
         * 
         * @return builder
         * 
         */
        public Builder direction(String direction) {
            return direction(Output.of(direction));
        }

        /**
         * @param minKbps The minimum kilobits per second. Changing this updates the min kbps value of the existing
         * QoS minimum bandwidth rule.
         * 
         * @return builder
         * 
         */
        public Builder minKbps(@Nullable Output<Integer> minKbps) {
            $.minKbps = minKbps;
            return this;
        }

        /**
         * @param minKbps The minimum kilobits per second. Changing this updates the min kbps value of the existing
         * QoS minimum bandwidth rule.
         * 
         * @return builder
         * 
         */
        public Builder minKbps(Integer minKbps) {
            return minKbps(Output.of(minKbps));
        }

        /**
         * @param qosPolicyId The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(@Nullable Output<String> qosPolicyId) {
            $.qosPolicyId = qosPolicyId;
            return this;
        }

        /**
         * @param qosPolicyId The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(String qosPolicyId) {
            return qosPolicyId(Output.of(qosPolicyId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
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
         * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public QosMinimumBandwidthRuleState build() {
            return $;
        }
    }

}
