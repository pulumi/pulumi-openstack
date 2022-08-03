// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecretArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecretArgs Empty = new GetSecretArgs();

    /**
     * Select the Secret with an ACL that contains the user.
     * Project scope is ignored. Defaults to `false`.
     * 
     */
    @Import(name="aclOnly")
    private @Nullable Output<Boolean> aclOnly;

    /**
     * @return Select the Secret with an ACL that contains the user.
     * Project scope is ignored. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> aclOnly() {
        return Optional.ofNullable(this.aclOnly);
    }

    /**
     * The Secret algorithm.
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return The Secret algorithm.
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * The Secret bit length.
     * 
     */
    @Import(name="bitLength")
    private @Nullable Output<Integer> bitLength;

    /**
     * @return The Secret bit length.
     * 
     */
    public Optional<Output<Integer>> bitLength() {
        return Optional.ofNullable(this.bitLength);
    }

    /**
     * Date filter to select the Secret with
     * created matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    @Import(name="createdAtFilter")
    private @Nullable Output<String> createdAtFilter;

    /**
     * @return Date filter to select the Secret with
     * created matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    public Optional<Output<String>> createdAtFilter() {
        return Optional.ofNullable(this.createdAtFilter);
    }

    /**
     * Date filter to select the Secret with
     * expiration matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    @Import(name="expirationFilter")
    private @Nullable Output<String> expirationFilter;

    /**
     * @return Date filter to select the Secret with
     * expiration matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    public Optional<Output<String>> expirationFilter() {
        return Optional.ofNullable(this.expirationFilter);
    }

    /**
     * The Secret mode.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return The Secret mode.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The Secret name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The Secret name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a secret. If omitted, the `region`
     * argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a secret. If omitted, the `region`
     * argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The Secret type. For more information see
     * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
     * 
     */
    @Import(name="secretType")
    private @Nullable Output<String> secretType;

    /**
     * @return The Secret type. For more information see
     * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
     * 
     */
    public Optional<Output<String>> secretType() {
        return Optional.ofNullable(this.secretType);
    }

    /**
     * Date filter to select the Secret with
     * updated matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    @Import(name="updatedAtFilter")
    private @Nullable Output<String> updatedAtFilter;

    /**
     * @return Date filter to select the Secret with
     * updated matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    public Optional<Output<String>> updatedAtFilter() {
        return Optional.ofNullable(this.updatedAtFilter);
    }

    private GetSecretArgs() {}

    private GetSecretArgs(GetSecretArgs $) {
        this.aclOnly = $.aclOnly;
        this.algorithm = $.algorithm;
        this.bitLength = $.bitLength;
        this.createdAtFilter = $.createdAtFilter;
        this.expirationFilter = $.expirationFilter;
        this.mode = $.mode;
        this.name = $.name;
        this.region = $.region;
        this.secretType = $.secretType;
        this.updatedAtFilter = $.updatedAtFilter;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecretArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecretArgs $;

        public Builder() {
            $ = new GetSecretArgs();
        }

        public Builder(GetSecretArgs defaults) {
            $ = new GetSecretArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclOnly Select the Secret with an ACL that contains the user.
         * Project scope is ignored. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder aclOnly(@Nullable Output<Boolean> aclOnly) {
            $.aclOnly = aclOnly;
            return this;
        }

        /**
         * @param aclOnly Select the Secret with an ACL that contains the user.
         * Project scope is ignored. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder aclOnly(Boolean aclOnly) {
            return aclOnly(Output.of(aclOnly));
        }

        /**
         * @param algorithm The Secret algorithm.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm The Secret algorithm.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param bitLength The Secret bit length.
         * 
         * @return builder
         * 
         */
        public Builder bitLength(@Nullable Output<Integer> bitLength) {
            $.bitLength = bitLength;
            return this;
        }

        /**
         * @param bitLength The Secret bit length.
         * 
         * @return builder
         * 
         */
        public Builder bitLength(Integer bitLength) {
            return bitLength(Output.of(bitLength));
        }

        /**
         * @param createdAtFilter Date filter to select the Secret with
         * created matching the specified criteria. See Date Filters below for more
         * detail.
         * 
         * @return builder
         * 
         */
        public Builder createdAtFilter(@Nullable Output<String> createdAtFilter) {
            $.createdAtFilter = createdAtFilter;
            return this;
        }

        /**
         * @param createdAtFilter Date filter to select the Secret with
         * created matching the specified criteria. See Date Filters below for more
         * detail.
         * 
         * @return builder
         * 
         */
        public Builder createdAtFilter(String createdAtFilter) {
            return createdAtFilter(Output.of(createdAtFilter));
        }

        /**
         * @param expirationFilter Date filter to select the Secret with
         * expiration matching the specified criteria. See Date Filters below for more
         * detail.
         * 
         * @return builder
         * 
         */
        public Builder expirationFilter(@Nullable Output<String> expirationFilter) {
            $.expirationFilter = expirationFilter;
            return this;
        }

        /**
         * @param expirationFilter Date filter to select the Secret with
         * expiration matching the specified criteria. See Date Filters below for more
         * detail.
         * 
         * @return builder
         * 
         */
        public Builder expirationFilter(String expirationFilter) {
            return expirationFilter(Output.of(expirationFilter));
        }

        /**
         * @param mode The Secret mode.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The Secret mode.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name The Secret name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The Secret name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V1 KeyManager client.
         * A KeyManager client is needed to fetch a secret. If omitted, the `region`
         * argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V1 KeyManager client.
         * A KeyManager client is needed to fetch a secret. If omitted, the `region`
         * argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param secretType The Secret type. For more information see
         * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
         * 
         * @return builder
         * 
         */
        public Builder secretType(@Nullable Output<String> secretType) {
            $.secretType = secretType;
            return this;
        }

        /**
         * @param secretType The Secret type. For more information see
         * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
         * 
         * @return builder
         * 
         */
        public Builder secretType(String secretType) {
            return secretType(Output.of(secretType));
        }

        /**
         * @param updatedAtFilter Date filter to select the Secret with
         * updated matching the specified criteria. See Date Filters below for more
         * detail.
         * 
         * @return builder
         * 
         */
        public Builder updatedAtFilter(@Nullable Output<String> updatedAtFilter) {
            $.updatedAtFilter = updatedAtFilter;
            return this;
        }

        /**
         * @param updatedAtFilter Date filter to select the Secret with
         * updated matching the specified criteria. See Date Filters below for more
         * detail.
         * 
         * @return builder
         * 
         */
        public Builder updatedAtFilter(String updatedAtFilter) {
            return updatedAtFilter(Output.of(updatedAtFilter));
        }

        public GetSecretArgs build() {
            return $;
        }
    }

}
