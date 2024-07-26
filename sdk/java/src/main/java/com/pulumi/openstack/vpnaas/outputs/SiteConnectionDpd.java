// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SiteConnectionDpd {
    /**
     * @return The dead peer detection (DPD) action.
     * A valid value is clear, hold, restart, disabled, or restart-by-peer.
     * Default value is hold.
     * 
     */
    private @Nullable String action;
    /**
     * @return The dead peer detection (DPD) interval, in seconds.
     * A valid value is a positive integer.
     * Default is 30.
     * 
     */
    private @Nullable Integer interval;
    /**
     * @return The dead peer detection (DPD) timeout in seconds.
     * A valid value is a positive integer that is greater than the DPD interval value.
     * Default is 120.
     * 
     */
    private @Nullable Integer timeout;

    private SiteConnectionDpd() {}
    /**
     * @return The dead peer detection (DPD) action.
     * A valid value is clear, hold, restart, disabled, or restart-by-peer.
     * Default value is hold.
     * 
     */
    public Optional<String> action() {
        return Optional.ofNullable(this.action);
    }
    /**
     * @return The dead peer detection (DPD) interval, in seconds.
     * A valid value is a positive integer.
     * Default is 30.
     * 
     */
    public Optional<Integer> interval() {
        return Optional.ofNullable(this.interval);
    }
    /**
     * @return The dead peer detection (DPD) timeout in seconds.
     * A valid value is a positive integer that is greater than the DPD interval value.
     * Default is 120.
     * 
     */
    public Optional<Integer> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SiteConnectionDpd defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String action;
        private @Nullable Integer interval;
        private @Nullable Integer timeout;
        public Builder() {}
        public Builder(SiteConnectionDpd defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.action = defaults.action;
    	      this.interval = defaults.interval;
    	      this.timeout = defaults.timeout;
        }

        @CustomType.Setter
        public Builder action(@Nullable String action) {

            this.action = action;
            return this;
        }
        @CustomType.Setter
        public Builder interval(@Nullable Integer interval) {

            this.interval = interval;
            return this;
        }
        @CustomType.Setter
        public Builder timeout(@Nullable Integer timeout) {

            this.timeout = timeout;
            return this;
        }
        public SiteConnectionDpd build() {
            final var _resultValue = new SiteConnectionDpd();
            _resultValue.action = action;
            _resultValue.interval = interval;
            _resultValue.timeout = timeout;
            return _resultValue;
        }
    }
}