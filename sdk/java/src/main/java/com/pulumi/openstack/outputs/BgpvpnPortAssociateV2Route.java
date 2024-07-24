// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BgpvpnPortAssociateV2Route {
    /**
     * @return The ID of the BGP VPN to be advertised. Required
     * if `type` is `bgpvpn`. Conflicts with `prefix`.
     * 
     */
    private @Nullable String bgpvpnId;
    /**
     * @return The BGP LOCAL\_PREF value of the routes that will
     * be advertised.
     * 
     */
    private @Nullable Integer localPref;
    /**
     * @return The CIDR prefix (v4 or v6) to be advertised. Required
     * if `type` is `prefix`. Conflicts with `bgpvpn_id`.
     * 
     */
    private @Nullable String prefix;
    /**
     * @return Can be `prefix` or `bgpvpn`. For the `prefix` type, the
     * CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
     * `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
     * 
     */
    private String type;

    private BgpvpnPortAssociateV2Route() {}
    /**
     * @return The ID of the BGP VPN to be advertised. Required
     * if `type` is `bgpvpn`. Conflicts with `prefix`.
     * 
     */
    public Optional<String> bgpvpnId() {
        return Optional.ofNullable(this.bgpvpnId);
    }
    /**
     * @return The BGP LOCAL\_PREF value of the routes that will
     * be advertised.
     * 
     */
    public Optional<Integer> localPref() {
        return Optional.ofNullable(this.localPref);
    }
    /**
     * @return The CIDR prefix (v4 or v6) to be advertised. Required
     * if `type` is `prefix`. Conflicts with `bgpvpn_id`.
     * 
     */
    public Optional<String> prefix() {
        return Optional.ofNullable(this.prefix);
    }
    /**
     * @return Can be `prefix` or `bgpvpn`. For the `prefix` type, the
     * CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
     * `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BgpvpnPortAssociateV2Route defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String bgpvpnId;
        private @Nullable Integer localPref;
        private @Nullable String prefix;
        private String type;
        public Builder() {}
        public Builder(BgpvpnPortAssociateV2Route defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bgpvpnId = defaults.bgpvpnId;
    	      this.localPref = defaults.localPref;
    	      this.prefix = defaults.prefix;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder bgpvpnId(@Nullable String bgpvpnId) {

            this.bgpvpnId = bgpvpnId;
            return this;
        }
        @CustomType.Setter
        public Builder localPref(@Nullable Integer localPref) {

            this.localPref = localPref;
            return this;
        }
        @CustomType.Setter
        public Builder prefix(@Nullable String prefix) {

            this.prefix = prefix;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("BgpvpnPortAssociateV2Route", "type");
            }
            this.type = type;
            return this;
        }
        public BgpvpnPortAssociateV2Route build() {
            final var _resultValue = new BgpvpnPortAssociateV2Route();
            _resultValue.bgpvpnId = bgpvpnId;
            _resultValue.localPref = localPref;
            _resultValue.prefix = prefix;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
