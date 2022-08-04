// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QosAssociationV3Args extends com.pulumi.resources.ResourceArgs {

    public static final QosAssociationV3Args Empty = new QosAssociationV3Args();

    /**
     * ID of the qos to associate. Changing this creates
     * a new qos association.
     * 
     */
    @Import(name="qosId", required=true)
    private Output<String> qosId;

    /**
     * @return ID of the qos to associate. Changing this creates
     * a new qos association.
     * 
     */
    public Output<String> qosId() {
        return this.qosId;
    }

    /**
     * The region in which to create the qos association.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new qos association.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to create the qos association.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new qos association.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * ID of the volume_type to associate.
     * Changing this creates a new qos association.
     * 
     */
    @Import(name="volumeTypeId", required=true)
    private Output<String> volumeTypeId;

    /**
     * @return ID of the volume_type to associate.
     * Changing this creates a new qos association.
     * 
     */
    public Output<String> volumeTypeId() {
        return this.volumeTypeId;
    }

    private QosAssociationV3Args() {}

    private QosAssociationV3Args(QosAssociationV3Args $) {
        this.qosId = $.qosId;
        this.region = $.region;
        this.volumeTypeId = $.volumeTypeId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QosAssociationV3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QosAssociationV3Args $;

        public Builder() {
            $ = new QosAssociationV3Args();
        }

        public Builder(QosAssociationV3Args defaults) {
            $ = new QosAssociationV3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param qosId ID of the qos to associate. Changing this creates
         * a new qos association.
         * 
         * @return builder
         * 
         */
        public Builder qosId(Output<String> qosId) {
            $.qosId = qosId;
            return this;
        }

        /**
         * @param qosId ID of the qos to associate. Changing this creates
         * a new qos association.
         * 
         * @return builder
         * 
         */
        public Builder qosId(String qosId) {
            return qosId(Output.of(qosId));
        }

        /**
         * @param region The region in which to create the qos association.
         * If omitted, the `region` argument of the provider is used. Changing
         * this creates a new qos association.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to create the qos association.
         * If omitted, the `region` argument of the provider is used. Changing
         * this creates a new qos association.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param volumeTypeId ID of the volume_type to associate.
         * Changing this creates a new qos association.
         * 
         * @return builder
         * 
         */
        public Builder volumeTypeId(Output<String> volumeTypeId) {
            $.volumeTypeId = volumeTypeId;
            return this;
        }

        /**
         * @param volumeTypeId ID of the volume_type to associate.
         * Changing this creates a new qos association.
         * 
         * @return builder
         * 
         */
        public Builder volumeTypeId(String volumeTypeId) {
            return volumeTypeId(Output.of(volumeTypeId));
        }

        public QosAssociationV3Args build() {
            $.qosId = Objects.requireNonNull($.qosId, "expected parameter 'qosId' to be non-null");
            $.volumeTypeId = Objects.requireNonNull($.volumeTypeId, "expected parameter 'volumeTypeId' to be non-null");
            return $;
        }
    }

}
