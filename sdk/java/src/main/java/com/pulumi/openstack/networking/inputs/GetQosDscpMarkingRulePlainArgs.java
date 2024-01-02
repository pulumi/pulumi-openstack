// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetQosDscpMarkingRulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetQosDscpMarkingRulePlainArgs Empty = new GetQosDscpMarkingRulePlainArgs();

    /**
     * The value of a DSCP mark.
     * 
     */
    @Import(name="dscpMark")
    private @Nullable Integer dscpMark;

    /**
     * @return The value of a DSCP mark.
     * 
     */
    public Optional<Integer> dscpMark() {
        return Optional.ofNullable(this.dscpMark);
    }

    /**
     * The QoS policy reference.
     * 
     */
    @Import(name="qosPolicyId", required=true)
    private String qosPolicyId;

    /**
     * @return The QoS policy reference.
     * 
     */
    public String qosPolicyId() {
        return this.qosPolicyId;
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetQosDscpMarkingRulePlainArgs() {}

    private GetQosDscpMarkingRulePlainArgs(GetQosDscpMarkingRulePlainArgs $) {
        this.dscpMark = $.dscpMark;
        this.qosPolicyId = $.qosPolicyId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetQosDscpMarkingRulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetQosDscpMarkingRulePlainArgs $;

        public Builder() {
            $ = new GetQosDscpMarkingRulePlainArgs();
        }

        public Builder(GetQosDscpMarkingRulePlainArgs defaults) {
            $ = new GetQosDscpMarkingRulePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dscpMark The value of a DSCP mark.
         * 
         * @return builder
         * 
         */
        public Builder dscpMark(@Nullable Integer dscpMark) {
            $.dscpMark = dscpMark;
            return this;
        }

        /**
         * @param qosPolicyId The QoS policy reference.
         * 
         * @return builder
         * 
         */
        public Builder qosPolicyId(String qosPolicyId) {
            $.qosPolicyId = qosPolicyId;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
         * `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetQosDscpMarkingRulePlainArgs build() {
            if ($.qosPolicyId == null) {
                throw new MissingRequiredPropertyException("GetQosDscpMarkingRulePlainArgs", "qosPolicyId");
            }
            return $;
        }
    }

}
