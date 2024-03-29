// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QosAssociationV3State extends com.pulumi.resources.ResourceArgs {

    public static final QosAssociationV3State Empty = new QosAssociationV3State();

    /**
     * ID of the qos to associate. Changing this creates
     * a new qos association.
     * 
     */
    @Import(name="qosId")
    private @Nullable Output<String> qosId;

    /**
     * @return ID of the qos to associate. Changing this creates
     * a new qos association.
     * 
     */
    public Optional<Output<String>> qosId() {
        return Optional.ofNullable(this.qosId);
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
    @Import(name="volumeTypeId")
    private @Nullable Output<String> volumeTypeId;

    /**
     * @return ID of the volume_type to associate.
     * Changing this creates a new qos association.
     * 
     */
    public Optional<Output<String>> volumeTypeId() {
        return Optional.ofNullable(this.volumeTypeId);
    }

    private QosAssociationV3State() {}

    private QosAssociationV3State(QosAssociationV3State $) {
        this.qosId = $.qosId;
        this.region = $.region;
        this.volumeTypeId = $.volumeTypeId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QosAssociationV3State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QosAssociationV3State $;

        public Builder() {
            $ = new QosAssociationV3State();
        }

        public Builder(QosAssociationV3State defaults) {
            $ = new QosAssociationV3State(Objects.requireNonNull(defaults));
        }

        /**
         * @param qosId ID of the qos to associate. Changing this creates
         * a new qos association.
         * 
         * @return builder
         * 
         */
        public Builder qosId(@Nullable Output<String> qosId) {
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
        public Builder volumeTypeId(@Nullable Output<String> volumeTypeId) {
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

        public QosAssociationV3State build() {
            return $;
        }
    }

}
