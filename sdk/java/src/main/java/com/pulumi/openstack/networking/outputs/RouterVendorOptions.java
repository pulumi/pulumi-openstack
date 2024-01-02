// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RouterVendorOptions {
    /**
     * @return Boolean to control whether
     * the Router gateway is assigned during creation or updated after creation.
     * 
     */
    private @Nullable Boolean setRouterGatewayAfterCreate;

    private RouterVendorOptions() {}
    /**
     * @return Boolean to control whether
     * the Router gateway is assigned during creation or updated after creation.
     * 
     */
    public Optional<Boolean> setRouterGatewayAfterCreate() {
        return Optional.ofNullable(this.setRouterGatewayAfterCreate);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RouterVendorOptions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean setRouterGatewayAfterCreate;
        public Builder() {}
        public Builder(RouterVendorOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.setRouterGatewayAfterCreate = defaults.setRouterGatewayAfterCreate;
        }

        @CustomType.Setter
        public Builder setRouterGatewayAfterCreate(@Nullable Boolean setRouterGatewayAfterCreate) {

            this.setRouterGatewayAfterCreate = setRouterGatewayAfterCreate;
            return this;
        }
        public RouterVendorOptions build() {
            final var _resultValue = new RouterVendorOptions();
            _resultValue.setRouterGatewayAfterCreate = setRouterGatewayAfterCreate;
            return _resultValue;
        }
    }
}
