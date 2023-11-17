// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.keymanager.inputs.SecretV1AclArgs;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretV1Args extends com.pulumi.resources.ResourceArgs {

    public static final SecretV1Args Empty = new SecretV1Args();

    /**
     * Allows to control an access to a secret. Currently only the
     * `read` operation is supported. If not specified, the secret is accessible
     * project wide.
     * 
     */
    @Import(name="acl")
    private @Nullable Output<SecretV1AclArgs> acl;

    /**
     * @return Allows to control an access to a secret. Currently only the
     * `read` operation is supported. If not specified, the secret is accessible
     * project wide.
     * 
     */
    public Optional<Output<SecretV1AclArgs>> acl() {
        return Optional.ofNullable(this.acl);
    }

    /**
     * Metadata provided by a user or system for informational purposes.
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return Metadata provided by a user or system for informational purposes.
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * Metadata provided by a user or system for informational purposes.
     * 
     */
    @Import(name="bitLength")
    private @Nullable Output<Integer> bitLength;

    /**
     * @return Metadata provided by a user or system for informational purposes.
     * 
     */
    public Optional<Output<Integer>> bitLength() {
        return Optional.ofNullable(this.bitLength);
    }

    /**
     * The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
     * 
     */
    @Import(name="expiration")
    private @Nullable Output<String> expiration;

    /**
     * @return The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
     * 
     */
    public Optional<Output<String>> expiration() {
        return Optional.ofNullable(this.expiration);
    }

    /**
     * Additional Metadata for the secret.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,Object>> metadata;

    /**
     * @return Additional Metadata for the secret.
     * 
     */
    public Optional<Output<Map<String,Object>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * Metadata provided by a user or system for informational purposes.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return Metadata provided by a user or system for informational purposes.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * Human-readable name for the Secret. Does not have
     * to be unique.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Human-readable name for the Secret. Does not have
     * to be unique.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The secret&#39;s data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
     * 
     */
    @Import(name="payload")
    private @Nullable Output<String> payload;

    /**
     * @return The secret&#39;s data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
     * 
     */
    public Optional<Output<String>> payload() {
        return Optional.ofNullable(this.payload);
    }

    /**
     * The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
     * 
     */
    @Import(name="payloadContentEncoding")
    private @Nullable Output<String> payloadContentEncoding;

    /**
     * @return The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
     * 
     */
    public Optional<Output<String>> payloadContentEncoding() {
        return Optional.ofNullable(this.payloadContentEncoding);
    }

    /**
     * The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
     * 
     */
    @Import(name="payloadContentType")
    private @Nullable Output<String> payloadContentType;

    /**
     * @return The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
     * 
     */
    public Optional<Output<String>> payloadContentType() {
        return Optional.ofNullable(this.payloadContentType);
    }

    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a secret. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 secret.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a secret. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 secret.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
     * 
     */
    @Import(name="secretType")
    private @Nullable Output<String> secretType;

    /**
     * @return Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
     * 
     */
    public Optional<Output<String>> secretType() {
        return Optional.ofNullable(this.secretType);
    }

    private SecretV1Args() {}

    private SecretV1Args(SecretV1Args $) {
        this.acl = $.acl;
        this.algorithm = $.algorithm;
        this.bitLength = $.bitLength;
        this.expiration = $.expiration;
        this.metadata = $.metadata;
        this.mode = $.mode;
        this.name = $.name;
        this.payload = $.payload;
        this.payloadContentEncoding = $.payloadContentEncoding;
        this.payloadContentType = $.payloadContentType;
        this.region = $.region;
        this.secretType = $.secretType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretV1Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretV1Args $;

        public Builder() {
            $ = new SecretV1Args();
        }

        public Builder(SecretV1Args defaults) {
            $ = new SecretV1Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl Allows to control an access to a secret. Currently only the
         * `read` operation is supported. If not specified, the secret is accessible
         * project wide.
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<SecretV1AclArgs> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl Allows to control an access to a secret. Currently only the
         * `read` operation is supported. If not specified, the secret is accessible
         * project wide.
         * 
         * @return builder
         * 
         */
        public Builder acl(SecretV1AclArgs acl) {
            return acl(Output.of(acl));
        }

        /**
         * @param algorithm Metadata provided by a user or system for informational purposes.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Metadata provided by a user or system for informational purposes.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param bitLength Metadata provided by a user or system for informational purposes.
         * 
         * @return builder
         * 
         */
        public Builder bitLength(@Nullable Output<Integer> bitLength) {
            $.bitLength = bitLength;
            return this;
        }

        /**
         * @param bitLength Metadata provided by a user or system for informational purposes.
         * 
         * @return builder
         * 
         */
        public Builder bitLength(Integer bitLength) {
            return bitLength(Output.of(bitLength));
        }

        /**
         * @param expiration The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
         * 
         * @return builder
         * 
         */
        public Builder expiration(@Nullable Output<String> expiration) {
            $.expiration = expiration;
            return this;
        }

        /**
         * @param expiration The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
         * 
         * @return builder
         * 
         */
        public Builder expiration(String expiration) {
            return expiration(Output.of(expiration));
        }

        /**
         * @param metadata Additional Metadata for the secret.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,Object>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Additional Metadata for the secret.
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,Object> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param mode Metadata provided by a user or system for informational purposes.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode Metadata provided by a user or system for informational purposes.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name Human-readable name for the Secret. Does not have
         * to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Human-readable name for the Secret. Does not have
         * to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param payload The secret&#39;s data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
         * 
         * @return builder
         * 
         */
        public Builder payload(@Nullable Output<String> payload) {
            $.payload = payload;
            return this;
        }

        /**
         * @param payload The secret&#39;s data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
         * 
         * @return builder
         * 
         */
        public Builder payload(String payload) {
            return payload(Output.of(payload));
        }

        /**
         * @param payloadContentEncoding The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
         * 
         * @return builder
         * 
         */
        public Builder payloadContentEncoding(@Nullable Output<String> payloadContentEncoding) {
            $.payloadContentEncoding = payloadContentEncoding;
            return this;
        }

        /**
         * @param payloadContentEncoding The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
         * 
         * @return builder
         * 
         */
        public Builder payloadContentEncoding(String payloadContentEncoding) {
            return payloadContentEncoding(Output.of(payloadContentEncoding));
        }

        /**
         * @param payloadContentType The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
         * 
         * @return builder
         * 
         */
        public Builder payloadContentType(@Nullable Output<String> payloadContentType) {
            $.payloadContentType = payloadContentType;
            return this;
        }

        /**
         * @param payloadContentType The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
         * 
         * @return builder
         * 
         */
        public Builder payloadContentType(String payloadContentType) {
            return payloadContentType(Output.of(payloadContentType));
        }

        /**
         * @param region The region in which to obtain the V1 KeyManager client.
         * A KeyManager client is needed to create a secret. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * V1 secret.
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
         * A KeyManager client is needed to create a secret. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * V1 secret.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param secretType Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
         * 
         * @return builder
         * 
         */
        public Builder secretType(@Nullable Output<String> secretType) {
            $.secretType = secretType;
            return this;
        }

        /**
         * @param secretType Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
         * 
         * @return builder
         * 
         */
        public Builder secretType(String secretType) {
            return secretType(Output.of(secretType));
        }

        public SecretV1Args build() {
            return $;
        }
    }

}
