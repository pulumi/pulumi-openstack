// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.dns;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TransferAcceptArgs extends com.pulumi.resources.ResourceArgs {

    public static final TransferAcceptArgs Empty = new TransferAcceptArgs();

    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack accept returned success.
     * 
     */
    @Import(name="disableStatusCheck")
    private @Nullable Output<Boolean> disableStatusCheck;

    /**
     * @return Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack accept returned success.
     * 
     */
    public Optional<Output<Boolean>> disableStatusCheck() {
        return Optional.ofNullable(this.disableStatusCheck);
    }

    /**
     * The transfer key.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return The transfer key.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Map of additional options. Changing this creates a
     * new transfer accept.
     * 
     */
    @Import(name="valueSpecs")
    private @Nullable Output<Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional options. Changing this creates a
     * new transfer accept.
     * 
     */
    public Optional<Output<Map<String,Object>>> valueSpecs() {
        return Optional.ofNullable(this.valueSpecs);
    }

    /**
     * The ID of the zone transfer request.
     * 
     */
    @Import(name="zoneTransferRequestId", required=true)
    private Output<String> zoneTransferRequestId;

    /**
     * @return The ID of the zone transfer request.
     * 
     */
    public Output<String> zoneTransferRequestId() {
        return this.zoneTransferRequestId;
    }

    private TransferAcceptArgs() {}

    private TransferAcceptArgs(TransferAcceptArgs $) {
        this.disableStatusCheck = $.disableStatusCheck;
        this.key = $.key;
        this.region = $.region;
        this.valueSpecs = $.valueSpecs;
        this.zoneTransferRequestId = $.zoneTransferRequestId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TransferAcceptArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TransferAcceptArgs $;

        public Builder() {
            $ = new TransferAcceptArgs();
        }

        public Builder(TransferAcceptArgs defaults) {
            $ = new TransferAcceptArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disableStatusCheck Disable wait for zone to reach ACTIVE
         * status. The check is enabled by default. If this argument is true, zone
         * will be considered as created/updated if OpenStack accept returned success.
         * 
         * @return builder
         * 
         */
        public Builder disableStatusCheck(@Nullable Output<Boolean> disableStatusCheck) {
            $.disableStatusCheck = disableStatusCheck;
            return this;
        }

        /**
         * @param disableStatusCheck Disable wait for zone to reach ACTIVE
         * status. The check is enabled by default. If this argument is true, zone
         * will be considered as created/updated if OpenStack accept returned success.
         * 
         * @return builder
         * 
         */
        public Builder disableStatusCheck(Boolean disableStatusCheck) {
            return disableStatusCheck(Output.of(disableStatusCheck));
        }

        /**
         * @param key The transfer key.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The transfer key.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * Keypairs are associated with accounts, but a Compute client is needed to
         * create one. If omitted, the `region` argument of the provider is used.
         * Changing this creates a new DNS zone.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * Keypairs are associated with accounts, but a Compute client is needed to
         * create one. If omitted, the `region` argument of the provider is used.
         * Changing this creates a new DNS zone.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param valueSpecs Map of additional options. Changing this creates a
         * new transfer accept.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(@Nullable Output<Map<String,Object>> valueSpecs) {
            $.valueSpecs = valueSpecs;
            return this;
        }

        /**
         * @param valueSpecs Map of additional options. Changing this creates a
         * new transfer accept.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(Map<String,Object> valueSpecs) {
            return valueSpecs(Output.of(valueSpecs));
        }

        /**
         * @param zoneTransferRequestId The ID of the zone transfer request.
         * 
         * @return builder
         * 
         */
        public Builder zoneTransferRequestId(Output<String> zoneTransferRequestId) {
            $.zoneTransferRequestId = zoneTransferRequestId;
            return this;
        }

        /**
         * @param zoneTransferRequestId The ID of the zone transfer request.
         * 
         * @return builder
         * 
         */
        public Builder zoneTransferRequestId(String zoneTransferRequestId) {
            return zoneTransferRequestId(Output.of(zoneTransferRequestId));
        }

        public TransferAcceptArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("TransferAcceptArgs", "key");
            }
            if ($.zoneTransferRequestId == null) {
                throw new MissingRequiredPropertyException("TransferAcceptArgs", "zoneTransferRequestId");
            }
            return $;
        }
    }

}
