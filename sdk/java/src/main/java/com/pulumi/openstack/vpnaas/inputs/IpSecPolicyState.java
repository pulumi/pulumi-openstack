// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.vpnaas.inputs.IpSecPolicyLifetimeArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpSecPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final IpSecPolicyState Empty = new IpSecPolicyState();

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
     * The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     * 
     */
    @Import(name="encapsulationMode")
    private @Nullable Output<String> encapsulationMode;

    /**
     * @return The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     * 
     */
    public Optional<Output<String>> encapsulationMode() {
        return Optional.ofNullable(this.encapsulationMode);
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
     * The lifetime of the security association. Consists of Unit and Value.
     * 
     */
    @Import(name="lifetimes")
    private @Nullable Output<List<IpSecPolicyLifetimeArgs>> lifetimes;

    /**
     * @return The lifetime of the security association. Consists of Unit and Value.
     * 
     */
    public Optional<Output<List<IpSecPolicyLifetimeArgs>>> lifetimes() {
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
     * The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
     * is group5. Changing this updates the existing policy.
     * 
     */
    @Import(name="pfs")
    private @Nullable Output<String> pfs;

    /**
     * @return The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
     * is group5. Changing this updates the existing policy.
     * 
     */
    public Optional<Output<String>> pfs() {
        return Optional.ofNullable(this.pfs);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * The transform protocol. Valid values are esp, ah and ah-esp.
     * Changing this updates the existing policy. Default is ESP.
     * 
     */
    @Import(name="transformProtocol")
    private @Nullable Output<String> transformProtocol;

    /**
     * @return The transform protocol. Valid values are esp, ah and ah-esp.
     * Changing this updates the existing policy. Default is ESP.
     * 
     */
    public Optional<Output<String>> transformProtocol() {
        return Optional.ofNullable(this.transformProtocol);
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

    private IpSecPolicyState() {}

    private IpSecPolicyState(IpSecPolicyState $) {
        this.authAlgorithm = $.authAlgorithm;
        this.description = $.description;
        this.encapsulationMode = $.encapsulationMode;
        this.encryptionAlgorithm = $.encryptionAlgorithm;
        this.lifetimes = $.lifetimes;
        this.name = $.name;
        this.pfs = $.pfs;
        this.region = $.region;
        this.tenantId = $.tenantId;
        this.transformProtocol = $.transformProtocol;
        this.valueSpecs = $.valueSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpSecPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpSecPolicyState $;

        public Builder() {
            $ = new IpSecPolicyState();
        }

        public Builder(IpSecPolicyState defaults) {
            $ = new IpSecPolicyState(Objects.requireNonNull(defaults));
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
         * @param encapsulationMode The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
         * Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder encapsulationMode(@Nullable Output<String> encapsulationMode) {
            $.encapsulationMode = encapsulationMode;
            return this;
        }

        /**
         * @param encapsulationMode The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
         * Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder encapsulationMode(String encapsulationMode) {
            return encapsulationMode(Output.of(encapsulationMode));
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
         * @param lifetimes The lifetime of the security association. Consists of Unit and Value.
         * 
         * @return builder
         * 
         */
        public Builder lifetimes(@Nullable Output<List<IpSecPolicyLifetimeArgs>> lifetimes) {
            $.lifetimes = lifetimes;
            return this;
        }

        /**
         * @param lifetimes The lifetime of the security association. Consists of Unit and Value.
         * 
         * @return builder
         * 
         */
        public Builder lifetimes(List<IpSecPolicyLifetimeArgs> lifetimes) {
            return lifetimes(Output.of(lifetimes));
        }

        /**
         * @param lifetimes The lifetime of the security association. Consists of Unit and Value.
         * 
         * @return builder
         * 
         */
        public Builder lifetimes(IpSecPolicyLifetimeArgs... lifetimes) {
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
         * @param pfs The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
         * is group5. Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder pfs(@Nullable Output<String> pfs) {
            $.pfs = pfs;
            return this;
        }

        /**
         * @param pfs The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
         * is group5. Changing this updates the existing policy.
         * 
         * @return builder
         * 
         */
        public Builder pfs(String pfs) {
            return pfs(Output.of(pfs));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create an IPSec policy. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * policy.
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
         * A Networking client is needed to create an IPSec policy. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * policy.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tenantId The owner of the policy. Required if admin wants to
         * create a policy for another project. Changing this creates a new policy.
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
         * create a policy for another project. Changing this creates a new policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param transformProtocol The transform protocol. Valid values are esp, ah and ah-esp.
         * Changing this updates the existing policy. Default is ESP.
         * 
         * @return builder
         * 
         */
        public Builder transformProtocol(@Nullable Output<String> transformProtocol) {
            $.transformProtocol = transformProtocol;
            return this;
        }

        /**
         * @param transformProtocol The transform protocol. Valid values are esp, ah and ah-esp.
         * Changing this updates the existing policy. Default is ESP.
         * 
         * @return builder
         * 
         */
        public Builder transformProtocol(String transformProtocol) {
            return transformProtocol(Output.of(transformProtocol));
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

        public IpSecPolicyState build() {
            return $;
        }
    }

}
