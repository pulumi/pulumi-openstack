// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PoolPersistence {
    /**
     * @return The name of the cookie if persistence mode is set
     * appropriately. Required if `type = APP_COOKIE`.
     * 
     */
    private @Nullable String cookieName;
    /**
     * @return The type of persistence mode. The current specification
     * supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
     * 
     */
    private String type;

    private PoolPersistence() {}
    /**
     * @return The name of the cookie if persistence mode is set
     * appropriately. Required if `type = APP_COOKIE`.
     * 
     */
    public Optional<String> cookieName() {
        return Optional.ofNullable(this.cookieName);
    }
    /**
     * @return The type of persistence mode. The current specification
     * supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PoolPersistence defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cookieName;
        private String type;
        public Builder() {}
        public Builder(PoolPersistence defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cookieName = defaults.cookieName;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder cookieName(@Nullable String cookieName) {
            this.cookieName = cookieName;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public PoolPersistence build() {
            final var o = new PoolPersistence();
            o.cookieName = cookieName;
            o.type = type;
            return o;
        }
    }
}
