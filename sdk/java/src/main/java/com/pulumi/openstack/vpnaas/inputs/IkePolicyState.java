// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.vpnaas.inputs.IkePolicyLifetimeArgs;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IkePolicyState extends com.pulumi.resources.ResourceArgs {

    public static final IkePolicyState Empty = new IkePolicyState();

    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     * 
     */
    @Import(name="authAlgorithm")
    private @Nullable Output<String> authAlgorithm;

    /**
     * @return The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     * 
     */
    public Optional<Output<String>> authAlgorithm() {
        return Optional.ofNullable(this.authAlgorithm);
    }

    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     * 
     */
    @Import(name="encryptionAlgorithm")
    private @Nullable Output<String> encryptionAlgorithm;

    /**
     * @return The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     * 
     */
    public Optional<Output<String>> encryptionAlgorithm() {
        return Optional.ofNullable(this.encryptionAlgorithm);
    }

    /**
     * The IKE mode. A valid value is v1 or v2. Default is v1.
     * Changing this updates the existing policy.
     * 
     */
    @Import(name="ikeVersion")
    private @Nullable Output<String> ikeVersion;

    /**
     * @return The IKE mode. A valid value is v1 or v2. Default is v1.
     * Changing this updates the existing policy.
     * 
     */
    public Optional<Output<String>> ikeVersion() {
        return Optional.ofNullable(this.ikeVersion);
    }

    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * 
     */
    @Import(name="lifetimes")
    private @Nullable Output<List<IkePolicyLifetimeArgs>> lifetimes;

    /**
     * @return The lifetime of the security association. Consists of Unit and Value.
     * 
     */
    public Optional<Output<List<IkePolicyLifetimeArgs>>> lifetimes() {
        return Optional.ofNullable(this.lifetimes);
    }

    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the policy. Changing this updates the name of
     * the existing policy.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
     * Changing this updates the existing policy.
     * 
     */
    @Import(name="pfs")
    private @Nullable Output<String> pfs;

    /**
     * @return The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
     * Changing this updates the existing policy.
     * 
     */
    public Optional<Output<String>> pfs() {
        return Optional.ofNullable(this.pfs);
    }

    /**
     * The IKE mode. A valid value is main, which is the default.
     * Changing this updates the existing policy.
     * 
     */
    @Import(name="phase1NegotiationMode")
    private @Nullable Output<String> phase1NegotiationMode;

    /**
     * @return The IKE mode. A valid value is main, which is the default.
     * Changing this updates the existing policy.
     * 
     */
    public Optional<Output<String>> phase1NegotiationMode() {
        return Optional.ofNullable(this.phase1NegotiationMode);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The owner of the policy. Required if admin wants to
     * create a service for another policy. Changing this creates a new policy.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the policy. Required if admin wants to
     * create a service for another policy. Changing this creates a new policy.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Map of additional options.
     * 
     */
    @Import(name="valueSpecs")
    private @Nullable Output<Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional options.
     * 
     */
    public Optional<Output<Map<String,Object>>> valueSpecs() {
        return Optional.ofNullable(this.valueSpecs);
    }

    private IkePolicyState() {}

    private IkePolicyState(IkePolicyState $) {
        this.authAlgorithm = $.authAlgorithm;
        this.description = $.description;
        this.encryptionAlgorithm = $.encryptionAlgorithm;
        this.ikeVersion = $.ikeVersion;
        this.lifetimes = $.lifetimes;
        this.name = $.name;
        this.pfs = $.pfs;
        this.phase1NegotiationMode = $.phase1NegotiationMode;
        this.region = $.region;
        this.tenantId = $.tenantId;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IkePolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IkePolicyState $;

        public Builder() {
            $ = new IkePolicyState();
        }

        public Builder(IkePolicyState defaults) {
            $ = new IkePolicyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authAlgorithm The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
         * Default is sha1. Changing this updates the algorithm of the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder authAlgorithm(@Nullable Output<String> authAlgorithm) {
            $.authAlgorithm = authAlgorithm;
            return this;
        }

        /**
         * @param authAlgorithm The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
         * Default is sha1. Changing this updates the algorithm of the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder authAlgorithm(String authAlgorithm) {
            return authAlgorithm(Output.of(authAlgorithm));
        }

        /**
         * @param description The human-readable description for the policy.
         * Changing this updates the description of the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The human-readable description for the policy.
         * Changing this updates the description of the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param encryptionAlgorithm The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
         * The default value is aes-128. Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder encryptionAlgorithm(@Nullable Output<String> encryptionAlgorithm) {
            $.encryptionAlgorithm = encryptionAlgorithm;
            return this;
        }

        /**
         * @param encryptionAlgorithm The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
         * The default value is aes-128. Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder encryptionAlgorithm(String encryptionAlgorithm) {
            return encryptionAlgorithm(Output.of(encryptionAlgorithm));
        }

        /**
         * @param ikeVersion The IKE mode. A valid value is v1 or v2. Default is v1.
         * Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder ikeVersion(@Nullable Output<String> ikeVersion) {
            $.ikeVersion = ikeVersion;
            return this;
        }

        /**
         * @param ikeVersion The IKE mode. A valid value is v1 or v2. Default is v1.
         * Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder ikeVersion(String ikeVersion) {
            return ikeVersion(Output.of(ikeVersion));
        }

        /**
         * @param lifetimes The lifetime of the security association. Consists of Unit and Value.
         * 
         * @return builder
         * 
         */
        public Builder lifetimes(@Nullable Output<List<IkePolicyLifetimeArgs>> lifetimes) {
            $.lifetimes = lifetimes;
            return this;
        }

        /**
         * @param lifetimes The lifetime of the security association. Consists of Unit and Value.
         * 
         * @return builder
         * 
         */
        public Builder lifetimes(List<IkePolicyLifetimeArgs> lifetimes) {
            return lifetimes(Output.of(lifetimes));
        }

        /**
         * @param lifetimes The lifetime of the security association. Consists of Unit and Value.
         * 
         * @return builder
         * 
         */
        public Builder lifetimes(IkePolicyLifetimeArgs... lifetimes) {
            return lifetimes(List.of(lifetimes));
        }

        /**
         * @param name The name of the policy. Changing this updates the name of
         * the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the policy. Changing this updates the name of
         * the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param pfs The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
         * Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder pfs(@Nullable Output<String> pfs) {
            $.pfs = pfs;
            return this;
        }

        /**
         * @param pfs The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
         * Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder pfs(String pfs) {
            return pfs(Output.of(pfs));
        }

        /**
         * @param phase1NegotiationMode The IKE mode. A valid value is main, which is the default.
         * Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder phase1NegotiationMode(@Nullable Output<String> phase1NegotiationMode) {
            $.phase1NegotiationMode = phase1NegotiationMode;
            return this;
        }

        /**
         * @param phase1NegotiationMode The IKE mode. A valid value is main, which is the default.
         * Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder phase1NegotiationMode(String phase1NegotiationMode) {
            return phase1NegotiationMode(Output.of(phase1NegotiationMode));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a VPN service. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * service.
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
         * A Networking client is needed to create a VPN service. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * service.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tenantId The owner of the policy. Required if admin wants to
         * create a service for another policy. Changing this creates a new policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the policy. Required if admin wants to
         * create a service for another policy. Changing this creates a new policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param valueSpecs Map of additional options.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(@Nullable Output<Map<String,Object>> valueSpecs) {
            $.valueSpecs = valueSpecs;
            return this;
        }

        /**
         * @param valueSpecs Map of additional options.
         * 
         * @return builder
         * 
         */
        public Builder valueSpecs(Map<String,Object> valueSpecs) {
            return valueSpecs(Output.of(valueSpecs));
        }

        public IkePolicyState build() {
            return $;
        }
    }

}
