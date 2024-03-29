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


public final class QosDscpMarkingRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final QosDscpMarkingRuleArgs Empty = new QosDscpMarkingRuleArgs();

    /**
     * The value of DSCP mark. Changing this updates the DSCP mark value existing
     * QoS DSCP marking rule.
     * 
     */
    @Import(name="dscpMark", required=true)
    private Output<Integer> dscpMark;

    /**
     * @return The value of DSCP mark. Changing this updates the DSCP mark value existing
     * QoS DSCP marking rule.
     * 
     */
    public Output<Integer> dscpMark() {
        return this.dscpMark;
    }

    /**
     * The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
     * 
     */
    @Import(name="qosPolicyId", required=true)
    private Output<String> qosPolicyId;

    /**
     * @return The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
     * 
     */
    public Output<String> qosPolicyId() {
        return this.qosPolicyId;
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private QosDscpMarkingRuleArgs() {}

    private QosDscpMarkingRuleArgs(QosDscpMarkingRuleArgs $) {
        this.dscpMark = $.dscpMark;
        this.qosPolicyId = $.qosPolicyId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QosDscpMarkingRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QosDscpMarkingRuleArgs $;

        public Builder() {
            $ = new QosDscpMarkingRuleArgs();
        }

        public Builder(QosDscpMarkingRuleArgs defaults) {
            $ = new QosDscpMarkingRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dscpMark The value of DSCP mark. Changing this updates the DSCP mark value existing
         * QoS DSCP marking rule.
         * 
         * @return builder
         * 
         */
        public Builder dscpMark(Output<Integer> dscpMark) {
            $.dscpMark = dscpMark;
            return this;
        }

        /**
         * @param dscpMark The value of DSCP mark. Changing this updates the DSCP mark value existing
         * QoS DSCP marking rule.
         * 
         * @return builder
         * 
         */
        public Builder dscpMark(Integer dscpMark) {
            return dscpMark(Output.of(dscpMark));
        }

        /**
         * @param qosPolicyId The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(Output<String> qosPolicyId) {
            $.qosPolicyId = qosPolicyId;
            return this;
        }

        /**
         * @param qosPolicyId The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(String qosPolicyId) {
            return qosPolicyId(Output.of(qosPolicyId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
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
         * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public QosDscpMarkingRuleArgs build() {
            if ($.dscpMark == null) {
                throw new MissingRequiredPropertyException("QosDscpMarkingRuleArgs", "dscpMark");
            }
            if ($.qosPolicyId == null) {
                throw new MissingRequiredPropertyException("QosDscpMarkingRuleArgs", "qosPolicyId");
            }
            return $;
        }
    }

}
