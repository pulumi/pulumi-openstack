// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecretPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecretPlainArgs Empty = new GetSecretPlainArgs();

    /**
     * Select the Secret with an ACL that contains the user.
     * Project scope is ignored. Defaults to `false`.
     * 
     */
    @Import(name="aclOnly")
    private @Nullable Boolean aclOnly;

    /**
     * @return Select the Secret with an ACL that contains the user.
     * Project scope is ignored. Defaults to `false`.
     * 
     */
    public Optional<Boolean> aclOnly() {
        return Optional.ofNullable(this.aclOnly);
    }

    /**
     * The Secret algorithm.
     * 
     */
    @Import(name="algorithm")
    private @Nullable String algorithm;

    /**
     * @return The Secret algorithm.
     * 
     */
    public Optional<String> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * The Secret bit length.
     * 
     */
    @Import(name="bitLength")
    private @Nullable Integer bitLength;

    /**
     * @return The Secret bit length.
     * 
     */
    public Optional<Integer> bitLength() {
        return Optional.ofNullable(this.bitLength);
    }

    /**
     * Date filter to select the Secret with
     * created matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    @Import(name="createdAtFilter")
    private @Nullable String createdAtFilter;

    /**
     * @return Date filter to select the Secret with
     * created matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    public Optional<String> createdAtFilter() {
        return Optional.ofNullable(this.createdAtFilter);
    }

    /**
     * Date filter to select the Secret with
     * expiration matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    @Import(name="expirationFilter")
    private @Nullable String expirationFilter;

    /**
     * @return Date filter to select the Secret with
     * expiration matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    public Optional<String> expirationFilter() {
        return Optional.ofNullable(this.expirationFilter);
    }

    /**
     * The Secret mode.
     * 
     */
    @Import(name="mode")
    private @Nullable String mode;

    /**
     * @return The Secret mode.
     * 
     */
    public Optional<String> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The Secret name.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The Secret name.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a secret. If omitted, the `region`
     * argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to fetch a secret. If omitted, the `region`
     * argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The Secret type. For more information see
     * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
     * 
     */
    @Import(name="secretType")
    private @Nullable String secretType;

    /**
     * @return The Secret type. For more information see
     * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
     * 
     */
    public Optional<String> secretType() {
        return Optional.ofNullable(this.secretType);
    }

    /**
     * Date filter to select the Secret with
     * updated matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    @Import(name="updatedAtFilter")
    private @Nullable String updatedAtFilter;

    /**
     * @return Date filter to select the Secret with
     * updated matching the specified criteria. See Date Filters below for more
     * detail.
     * 
     */
    public Optional<String> updatedAtFilter() {
        return Optional.ofNullable(this.updatedAtFilter);
    }

    private GetSecretPlainArgs() {}

    private GetSecretPlainArgs(GetSecretPlainArgs $) {
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
    public static Builder builder(GetSecretPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecretPlainArgs $;

        public Builder() {
            $ = new GetSecretPlainArgs();
        }

        public Builder(GetSecretPlainArgs defaults) {
            $ = new GetSecretPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclOnly Select the Secret with an ACL that contains the user.
         * Project scope is ignored. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder aclOnly(@Nullable Boolean aclOnly) {
            $.aclOnly = aclOnly;
            return this;
        }

        /**
         * @param algorithm The Secret algorithm.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable String algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param bitLength The Secret bit length.
         * 
         * @return builder
         * 
         */
        public Builder bitLength(@Nullable Integer bitLength) {
            $.bitLength = bitLength;
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
        public Builder createdAtFilter(@Nullable String createdAtFilter) {
            $.createdAtFilter = createdAtFilter;
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
        public Builder expirationFilter(@Nullable String expirationFilter) {
            $.expirationFilter = expirationFilter;
            return this;
        }

        /**
         * @param mode The Secret mode.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable String mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param name The Secret name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
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
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        /**
         * @param secretType The Secret type. For more information see
         * [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
         * 
         * @return builder
         * 
         */
        public Builder secretType(@Nullable String secretType) {
            $.secretType = secretType;
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
        public Builder updatedAtFilter(@Nullable String updatedAtFilter) {
            $.updatedAtFilter = updatedAtFilter;
            return this;
        }

        public GetSecretPlainArgs build() {
            return $;
        }
    }

}
