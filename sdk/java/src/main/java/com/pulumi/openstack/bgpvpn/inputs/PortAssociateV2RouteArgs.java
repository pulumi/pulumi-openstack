// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.bgpvpn.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PortAssociateV2RouteArgs extends com.pulumi.resources.ResourceArgs {

    public static final PortAssociateV2RouteArgs Empty = new PortAssociateV2RouteArgs();

    /**
     * The ID of the BGP VPN to be advertised. Required
     * if `type` is `bgpvpn`. Conflicts with `prefix`.
     * 
     */
    @Import(name="bgpvpnId")
    private @Nullable Output<String> bgpvpnId;

    /**
     * @return The ID of the BGP VPN to be advertised. Required
     * if `type` is `bgpvpn`. Conflicts with `prefix`.
     * 
     */
    public Optional<Output<String>> bgpvpnId() {
        return Optional.ofNullable(this.bgpvpnId);
    }

    /**
     * The BGP LOCAL\_PREF value of the routes that will
     * be advertised.
     * 
     */
    @Import(name="localPref")
    private @Nullable Output<Integer> localPref;

    /**
     * @return The BGP LOCAL\_PREF value of the routes that will
     * be advertised.
     * 
     */
    public Optional<Output<Integer>> localPref() {
        return Optional.ofNullable(this.localPref);
    }

    /**
     * The CIDR prefix (v4 or v6) to be advertised. Required
     * if `type` is `prefix`. Conflicts with `bgpvpn_id`.
     * 
     */
    @Import(name="prefix")
    private @Nullable Output<String> prefix;

    /**
     * @return The CIDR prefix (v4 or v6) to be advertised. Required
     * if `type` is `prefix`. Conflicts with `bgpvpn_id`.
     * 
     */
    public Optional<Output<String>> prefix() {
        return Optional.ofNullable(this.prefix);
    }

    /**
     * Can be `prefix` or `bgpvpn`. For the `prefix` type, the
     * CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
     * `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Can be `prefix` or `bgpvpn`. For the `prefix` type, the
     * CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
     * `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private PortAssociateV2RouteArgs() {}

    private PortAssociateV2RouteArgs(PortAssociateV2RouteArgs $) {
        this.bgpvpnId = $.bgpvpnId;
        this.localPref = $.localPref;
        this.prefix = $.prefix;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PortAssociateV2RouteArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PortAssociateV2RouteArgs $;

        public Builder() {
            $ = new PortAssociateV2RouteArgs();
        }

        public Builder(PortAssociateV2RouteArgs defaults) {
            $ = new PortAssociateV2RouteArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bgpvpnId The ID of the BGP VPN to be advertised. Required
         * if `type` is `bgpvpn`. Conflicts with `prefix`.
         * 
         * @return builder
         * 
         */
        public Builder bgpvpnId(@Nullable Output<String> bgpvpnId) {
            $.bgpvpnId = bgpvpnId;
            return this;
        }

        /**
         * @param bgpvpnId The ID of the BGP VPN to be advertised. Required
         * if `type` is `bgpvpn`. Conflicts with `prefix`.
         * 
         * @return builder
         * 
         */
        public Builder bgpvpnId(String bgpvpnId) {
            return bgpvpnId(Output.of(bgpvpnId));
        }

        /**
         * @param localPref The BGP LOCAL\_PREF value of the routes that will
         * be advertised.
         * 
         * @return builder
         * 
         */
        public Builder localPref(@Nullable Output<Integer> localPref) {
            $.localPref = localPref;
            return this;
        }

        /**
         * @param localPref The BGP LOCAL\_PREF value of the routes that will
         * be advertised.
         * 
         * @return builder
         * 
         */
        public Builder localPref(Integer localPref) {
            return localPref(Output.of(localPref));
        }

        /**
         * @param prefix The CIDR prefix (v4 or v6) to be advertised. Required
         * if `type` is `prefix`. Conflicts with `bgpvpn_id`.
         * 
         * @return builder
         * 
         */
        public Builder prefix(@Nullable Output<String> prefix) {
            $.prefix = prefix;
            return this;
        }

        /**
         * @param prefix The CIDR prefix (v4 or v6) to be advertised. Required
         * if `type` is `prefix`. Conflicts with `bgpvpn_id`.
         * 
         * @return builder
         * 
         */
        public Builder prefix(String prefix) {
            return prefix(Output.of(prefix));
        }

        /**
         * @param type Can be `prefix` or `bgpvpn`. For the `prefix` type, the
         * CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
         * `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Can be `prefix` or `bgpvpn`. For the `prefix` type, the
         * CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
         * `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public PortAssociateV2RouteArgs build() {
            if ($.type == null) {
                throw new MissingRequiredPropertyException("PortAssociateV2RouteArgs", "type");
            }
            return $;
        }
    }

}
