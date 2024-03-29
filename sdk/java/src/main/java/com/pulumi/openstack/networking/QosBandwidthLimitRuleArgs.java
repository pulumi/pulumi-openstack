// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QosBandwidthLimitRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final QosBandwidthLimitRuleArgs Empty = new QosBandwidthLimitRuleArgs();

    /**
     * The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
     * existing QoS bandwidth limit rule.
     * 
     */
    @Import(name="direction")
    private @Nullable Output<String> direction;

    /**
     * @return The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
     * existing QoS bandwidth limit rule.
     * 
     */
    public Optional<Output<String>> direction() {
        return Optional.ofNullable(this.direction);
    }

    /**
     * The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
     * maximum burst size in kilobits of the existing QoS bandwidth limit rule.
     * 
     */
    @Import(name="maxBurstKbps")
    private @Nullable Output<Integer> maxBurstKbps;

    /**
     * @return The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
     * maximum burst size in kilobits of the existing QoS bandwidth limit rule.
     * 
     */
    public Optional<Output<Integer>> maxBurstKbps() {
        return Optional.ofNullable(this.maxBurstKbps);
    }

    /**
     * The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
     * maximum kilobits per second of the existing QoS bandwidth limit rule.
     * 
     */
    @Import(name="maxKbps", required=true)
    private Output<Integer> maxKbps;

    /**
     * @return The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
     * maximum kilobits per second of the existing QoS bandwidth limit rule.
     * 
     */
    public Output<Integer> maxKbps() {
        return this.maxKbps;
    }

    /**
     * The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
     * 
     */
    @Import(name="qosPolicyId", required=true)
    private Output<String> qosPolicyId;

    /**
     * @return The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
     * 
     */
    public Output<String> qosPolicyId() {
        return this.qosPolicyId;
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private QosBandwidthLimitRuleArgs() {}

    private QosBandwidthLimitRuleArgs(QosBandwidthLimitRuleArgs $) {
        this.direction = $.direction;
        this.maxBurstKbps = $.maxBurstKbps;
        this.maxKbps = $.maxKbps;
        this.qosPolicyId = $.qosPolicyId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QosBandwidthLimitRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QosBandwidthLimitRuleArgs $;

        public Builder() {
            $ = new QosBandwidthLimitRuleArgs();
        }

        public Builder(QosBandwidthLimitRuleArgs defaults) {
            $ = new QosBandwidthLimitRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param direction The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
         * existing QoS bandwidth limit rule.
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
         * existing QoS bandwidth limit rule.
         * 
         * @return builder
         * 
         */
        public Builder direction(String direction) {
            return direction(Output.of(direction));
        }

        /**
         * @param maxBurstKbps The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
         * maximum burst size in kilobits of the existing QoS bandwidth limit rule.
         * 
         * @return builder
         * 
         */
        public Builder maxBurstKbps(@Nullable Output<Integer> maxBurstKbps) {
            $.maxBurstKbps = maxBurstKbps;
            return this;
        }

        /**
         * @param maxBurstKbps The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
         * maximum burst size in kilobits of the existing QoS bandwidth limit rule.
         * 
         * @return builder
         * 
         */
        public Builder maxBurstKbps(Integer maxBurstKbps) {
            return maxBurstKbps(Output.of(maxBurstKbps));
        }

        /**
         * @param maxKbps The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
         * maximum kilobits per second of the existing QoS bandwidth limit rule.
         * 
         * @return builder
         * 
         */
        public Builder maxKbps(Output<Integer> maxKbps) {
            $.maxKbps = maxKbps;
            return this;
        }

        /**
         * @param maxKbps The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
         * maximum kilobits per second of the existing QoS bandwidth limit rule.
         * 
         * @return builder
         * 
         */
        public Builder maxKbps(Integer maxKbps) {
            return maxKbps(Output.of(maxKbps));
        }

        /**
         * @param qosPolicyId The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(Output<String> qosPolicyId) {
            $.qosPolicyId = qosPolicyId;
            return this;
        }

        /**
         * @param qosPolicyId The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(String qosPolicyId) {
            return qosPolicyId(Output.of(qosPolicyId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
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
         * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public QosBandwidthLimitRuleArgs build() {
            if ($.maxKbps == null) {
                throw new MissingRequiredPropertyException("QosBandwidthLimitRuleArgs", "maxKbps");
            }
            if ($.qosPolicyId == null) {
                throw new MissingRequiredPropertyException("QosBandwidthLimitRuleArgs", "qosPolicyId");
            }
            return $;
        }
    }

}
