// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KeypairArgs extends com.pulumi.resources.ResourceArgs {

    public static final KeypairArgs Empty = new KeypairArgs();

    /**
     * A unique name for the keypair. Changing this creates a new
     * keypair.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name for the keypair. Changing this creates a new
     * keypair.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A pregenerated OpenSSH-formatted public key.
     * Changing this creates a new keypair. If a public key is not specified, then
     * a public/private key pair will be automatically generated. If a pair is
     * created, then destroying this resource means you will lose access to that
     * keypair forever.
     * 
     */
    @Import(name="publicKey")
    private @Nullable Output<String> publicKey;

    /**
     * @return A pregenerated OpenSSH-formatted public key.
     * Changing this creates a new keypair. If a public key is not specified, then
     * a public/private key pair will be automatically generated. If a pair is
     * created, then destroying this resource means you will lose access to that
     * keypair forever.
     * 
     */
    public Optional<Output<String>> publicKey() {
        return Optional.ofNullable(this.publicKey);
    }

    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new keypair.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new keypair.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * This allows administrative users to operate key-pairs
     * of specified user ID. For this feature your need to have openstack microversion
     * 2.10 (Liberty) or later.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<String> userId;

    /**
     * @return This allows administrative users to operate key-pairs
     * of specified user ID. For this feature your need to have openstack microversion
     * 2.10 (Liberty) or later.
     * 
     */
    public Optional<Output<String>> userId() {
        return Optional.ofNullable(this.userId);
    }

    /**
     * Map of additional options.
     * 
     */
    @Import(name="valueSpecs")
    private @Nullable Output<Map<String,String>> valueSpecs;

    /**
     * @return Map of additional options.
     * 
     */
    public Optional<Output<Map<String,String>>> valueSpecs() {
        return Optional.ofNullable(this.valueSpecs);
    }

    private KeypairArgs() {}

    private KeypairArgs(KeypairArgs $) {
        this.name = $.name;
        this.publicKey = $.publicKey;
        this.region = $.region;
        this.userId = $.userId;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KeypairArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KeypairArgs $;

        public Builder() {
            $ = new KeypairArgs();
        }

        public Builder(KeypairArgs defaults) {
            $ = new KeypairArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name A unique name for the keypair. Changing this creates a new
         * keypair.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name for the keypair. Changing this creates a new
         * keypair.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param publicKey A pregenerated OpenSSH-formatted public key.
         * Changing this creates a new keypair. If a public key is not specified, then
         * a public/private key pair will be automatically generated. If a pair is
         * created, then destroying this resource means you will lose access to that
         * keypair forever.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(@Nullable Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey A pregenerated OpenSSH-formatted public key.
         * Changing this creates a new keypair. If a public key is not specified, then
         * a public/private key pair will be automatically generated. If a pair is
         * created, then destroying this resource means you will lose access to that
         * keypair forever.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * Keypairs are associated with accounts, but a Compute client is needed to
         * create one. If omitted, the `region` argument of the provider is used.
         * Changing this creates a new keypair.
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
         * Changing this creates a new keypair.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param userId This allows administrative users to operate key-pairs
         * of specified user ID. For this feature your need to have openstack microversion
         * 2.10 (Liberty) or later.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId This allows administrative users to operate key-pairs
         * of specified user ID. For this feature your need to have openstack microversion
         * 2.10 (Liberty) or later.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        /**
         * @param valueSpecs Map of additional options.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(@Nullable Output<Map<String,String>> valueSpecs) {
            $.valueSpecs = valueSpecs;
            return this;
        }

        /**
         * @param valueSpecs Map of additional options.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(Map<String,String> valueSpecs) {
            return valueSpecs(Output.of(valueSpecs));
        }

        public KeypairArgs build() {
            return $;
        }
    }

}
