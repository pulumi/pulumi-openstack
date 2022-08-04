// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OrderV1Meta {
    /**
     * @return Algorithm to use for key generation.
     * 
     */
    private final String algorithm;
    /**
     * @return - Bit lenght of key to be generated.
     * 
     */
    private final Integer bitLength;
    /**
     * @return This is a UTC timestamp in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ. If set, the secret will not be available after this time.
     * 
     */
    private final @Nullable String expiration;
    /**
     * @return The mode to use for key generation.
     * 
     */
    private final @Nullable String mode;
    /**
     * @return The name of the secret set by the user.
     * 
     */
    private final @Nullable String name;
    /**
     * @return The media type for the content of the secrets payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
     * 
     */
    private final @Nullable String payloadContentType;

    @CustomType.Constructor
    private OrderV1Meta(
        @CustomType.Parameter("algorithm") String algorithm,
        @CustomType.Parameter("bitLength") Integer bitLength,
        @CustomType.Parameter("expiration") @Nullable String expiration,
        @CustomType.Parameter("mode") @Nullable String mode,
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("payloadContentType") @Nullable String payloadContentType) {
        this.algorithm = algorithm;
        this.bitLength = bitLength;
        this.expiration = expiration;
        this.mode = mode;
        this.name = name;
        this.payloadContentType = payloadContentType;
    }

    /**
     * @return Algorithm to use for key generation.
     * 
     */
    public String algorithm() {
        return this.algorithm;
    }
    /**
     * @return - Bit lenght of key to be generated.
     * 
     */
    public Integer bitLength() {
        return this.bitLength;
    }
    /**
     * @return This is a UTC timestamp in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ. If set, the secret will not be available after this time.
     * 
     */
    public Optional<String> expiration() {
        return Optional.ofNullable(this.expiration);
    }
    /**
     * @return The mode to use for key generation.
     * 
     */
    public Optional<String> mode() {
        return Optional.ofNullable(this.mode);
    }
    /**
     * @return The name of the secret set by the user.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The media type for the content of the secrets payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
     * 
     */
    public Optional<String> payloadContentType() {
        return Optional.ofNullable(this.payloadContentType);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrderV1Meta defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String algorithm;
        private Integer bitLength;
        private @Nullable String expiration;
        private @Nullable String mode;
        private @Nullable String name;
        private @Nullable String payloadContentType;

        public Builder() {
    	      // Empty
        }

        public Builder(OrderV1Meta defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithm = defaults.algorithm;
    	      this.bitLength = defaults.bitLength;
    	      this.expiration = defaults.expiration;
    	      this.mode = defaults.mode;
    	      this.name = defaults.name;
    	      this.payloadContentType = defaults.payloadContentType;
        }

        public Builder algorithm(String algorithm) {
            this.algorithm = Objects.requireNonNull(algorithm);
            return this;
        }
        public Builder bitLength(Integer bitLength) {
            this.bitLength = Objects.requireNonNull(bitLength);
            return this;
        }
        public Builder expiration(@Nullable String expiration) {
            this.expiration = expiration;
            return this;
        }
        public Builder mode(@Nullable String mode) {
            this.mode = mode;
            return this;
        }
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder payloadContentType(@Nullable String payloadContentType) {
            this.payloadContentType = payloadContentType;
            return this;
        }        public OrderV1Meta build() {
            return new OrderV1Meta(algorithm, bitLength, expiration, mode, name, payloadContentType);
        }
    }
}
