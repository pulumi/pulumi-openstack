// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class L7PolicyV2Args extends com.pulumi.resources.ResourceArgs {

    public static final L7PolicyV2Args Empty = new L7PolicyV2Args();

    /**
     * The L7 Policy action - can either be REDIRECT\_TO\_POOL,
     * REDIRECT\_TO\_URL or REJECT.
     * 
     */
    @Import(name="action", required=true)
    private Output<String> action;

    /**
     * @return The L7 Policy action - can either be REDIRECT\_TO\_POOL,
     * REDIRECT\_TO\_URL or REJECT.
     * 
     */
    public Output<String> action() {
        return this.action;
    }

    /**
     * The administrative state of the L7 Policy.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the L7 Policy.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * Human-readable description for the L7 Policy.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description for the L7 Policy.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The Listener on which the L7 Policy will be associated with.
     * Changing this creates a new L7 Policy.
     * 
     */
    @Import(name="listenerId", required=true)
    private Output<String> listenerId;

    /**
     * @return The Listener on which the L7 Policy will be associated with.
     * Changing this creates a new L7 Policy.
     * 
     */
    public Output<String> listenerId() {
        return this.listenerId;
    }

    /**
     * Human-readable name for the L7 Policy. Does not have
     * to be unique.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Human-readable name for the L7 Policy. Does not have
     * to be unique.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The position of this policy on the listener. Positions start at 1.
     * 
     */
    @Import(name="position")
    private @Nullable Output<Integer> position;

    /**
     * @return The position of this policy on the listener. Positions start at 1.
     * 
     */
    public Optional<Output<Integer>> position() {
        return Optional.ofNullable(this.position);
    }

    /**
     * Integer. Requests matching this policy will be\
     * redirected to the specified URL or Prefix URL with the HTTP response code.
     * Valid if action is REDIRECT\_TO\_URL or REDIRECT\_PREFIX. Valid options are:
     * 301, 302, 303, 307, or 308. Default is 302. New in octavia version 2.9
     * 
     */
    @Import(name="redirectHttpCode")
    private @Nullable Output<Integer> redirectHttpCode;

    /**
     * @return Integer. Requests matching this policy will be\
     * redirected to the specified URL or Prefix URL with the HTTP response code.
     * Valid if action is REDIRECT\_TO\_URL or REDIRECT\_PREFIX. Valid options are:
     * 301, 302, 303, 307, or 308. Default is 302. New in octavia version 2.9
     * 
     */
    public Optional<Output<Integer>> redirectHttpCode() {
        return Optional.ofNullable(this.redirectHttpCode);
    }

    /**
     * Requests matching this policy will be redirected to the
     * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
     * 
     */
    @Import(name="redirectPoolId")
    private @Nullable Output<String> redirectPoolId;

    /**
     * @return Requests matching this policy will be redirected to the
     * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
     * 
     */
    public Optional<Output<String>> redirectPoolId() {
        return Optional.ofNullable(this.redirectPoolId);
    }

    /**
     * Requests matching this policy will be redirected to
     * this Prefix URL. Only valid if action is REDIRECT\_PREFIX.
     * 
     */
    @Import(name="redirectPrefix")
    private @Nullable Output<String> redirectPrefix;

    /**
     * @return Requests matching this policy will be redirected to
     * this Prefix URL. Only valid if action is REDIRECT\_PREFIX.
     * 
     */
    public Optional<Output<String>> redirectPrefix() {
        return Optional.ofNullable(this.redirectPrefix);
    }

    /**
     * Requests matching this policy will be redirected to this URL.
     * Only valid if action is REDIRECT\_TO\_URL.
     * 
     */
    @Import(name="redirectUrl")
    private @Nullable Output<String> redirectUrl;

    /**
     * @return Requests matching this policy will be redirected to this URL.
     * Only valid if action is REDIRECT\_TO\_URL.
     * 
     */
    public Optional<Output<String>> redirectUrl() {
        return Optional.ofNullable(this.redirectUrl);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Policy.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Policy.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Policy.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Policy.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the L7 Policy.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Policy.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private L7PolicyV2Args() {}

    private L7PolicyV2Args(L7PolicyV2Args $) {
        this.action = $.action;
        this.adminStateUp = $.adminStateUp;
        this.description = $.description;
        this.listenerId = $.listenerId;
        this.name = $.name;
        this.position = $.position;
        this.redirectHttpCode = $.redirectHttpCode;
        this.redirectPoolId = $.redirectPoolId;
        this.redirectPrefix = $.redirectPrefix;
        this.redirectUrl = $.redirectUrl;
        this.region = $.region;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(L7PolicyV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private L7PolicyV2Args $;

        public Builder() {
            $ = new L7PolicyV2Args();
        }

        public Builder(L7PolicyV2Args defaults) {
            $ = new L7PolicyV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The L7 Policy action - can either be REDIRECT\_TO\_POOL,
         * REDIRECT\_TO\_URL or REJECT.
         * 
         * @return builder
         * 
         */
        public Builder action(Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action The L7 Policy action - can either be REDIRECT\_TO\_POOL,
         * REDIRECT\_TO\_URL or REJECT.
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param adminStateUp The administrative state of the L7 Policy.
         * A valid value is true (UP) or false (DOWN).
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp The administrative state of the L7 Policy.
         * A valid value is true (UP) or false (DOWN).
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param description Human-readable description for the L7 Policy.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description for the L7 Policy.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param listenerId The Listener on which the L7 Policy will be associated with.
         * Changing this creates a new L7 Policy.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(Output<String> listenerId) {
            $.listenerId = listenerId;
            return this;
        }

        /**
         * @param listenerId The Listener on which the L7 Policy will be associated with.
         * Changing this creates a new L7 Policy.
         * 
         * @return builder
         * 
         */
        public Builder listenerId(String listenerId) {
            return listenerId(Output.of(listenerId));
        }

        /**
         * @param name Human-readable name for the L7 Policy. Does not have
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
         * @param name Human-readable name for the L7 Policy. Does not have
         * to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param position The position of this policy on the listener. Positions start at 1.
         * 
         * @return builder
         * 
         */
        public Builder position(@Nullable Output<Integer> position) {
            $.position = position;
            return this;
        }

        /**
         * @param position The position of this policy on the listener. Positions start at 1.
         * 
         * @return builder
         * 
         */
        public Builder position(Integer position) {
            return position(Output.of(position));
        }

        /**
         * @param redirectHttpCode Integer. Requests matching this policy will be\
         * redirected to the specified URL or Prefix URL with the HTTP response code.
         * Valid if action is REDIRECT\_TO\_URL or REDIRECT\_PREFIX. Valid options are:
         * 301, 302, 303, 307, or 308. Default is 302. New in octavia version 2.9
         * 
         * @return builder
         * 
         */
        public Builder redirectHttpCode(@Nullable Output<Integer> redirectHttpCode) {
            $.redirectHttpCode = redirectHttpCode;
            return this;
        }

        /**
         * @param redirectHttpCode Integer. Requests matching this policy will be\
         * redirected to the specified URL or Prefix URL with the HTTP response code.
         * Valid if action is REDIRECT\_TO\_URL or REDIRECT\_PREFIX. Valid options are:
         * 301, 302, 303, 307, or 308. Default is 302. New in octavia version 2.9
         * 
         * @return builder
         * 
         */
        public Builder redirectHttpCode(Integer redirectHttpCode) {
            return redirectHttpCode(Output.of(redirectHttpCode));
        }

        /**
         * @param redirectPoolId Requests matching this policy will be redirected to the
         * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
         * 
         * @return builder
         * 
         */
        public Builder redirectPoolId(@Nullable Output<String> redirectPoolId) {
            $.redirectPoolId = redirectPoolId;
            return this;
        }

        /**
         * @param redirectPoolId Requests matching this policy will be redirected to the
         * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
         * 
         * @return builder
         * 
         */
        public Builder redirectPoolId(String redirectPoolId) {
            return redirectPoolId(Output.of(redirectPoolId));
        }

        /**
         * @param redirectPrefix Requests matching this policy will be redirected to
         * this Prefix URL. Only valid if action is REDIRECT\_PREFIX.
         * 
         * @return builder
         * 
         */
        public Builder redirectPrefix(@Nullable Output<String> redirectPrefix) {
            $.redirectPrefix = redirectPrefix;
            return this;
        }

        /**
         * @param redirectPrefix Requests matching this policy will be redirected to
         * this Prefix URL. Only valid if action is REDIRECT\_PREFIX.
         * 
         * @return builder
         * 
         */
        public Builder redirectPrefix(String redirectPrefix) {
            return redirectPrefix(Output.of(redirectPrefix));
        }

        /**
         * @param redirectUrl Requests matching this policy will be redirected to this URL.
         * Only valid if action is REDIRECT\_TO\_URL.
         * 
         * @return builder
         * 
         */
        public Builder redirectUrl(@Nullable Output<String> redirectUrl) {
            $.redirectUrl = redirectUrl;
            return this;
        }

        /**
         * @param redirectUrl Requests matching this policy will be redirected to this URL.
         * Only valid if action is REDIRECT\_TO\_URL.
         * 
         * @return builder
         * 
         */
        public Builder redirectUrl(String redirectUrl) {
            return redirectUrl(Output.of(redirectUrl));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create an . If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * L7 Policy.
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
         * A Networking client is needed to create an . If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * L7 Policy.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tenantId Required for admins. The UUID of the tenant who owns
         * the L7 Policy.  Only administrative users can specify a tenant UUID
         * other than their own. Changing this creates a new L7 Policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId Required for admins. The UUID of the tenant who owns
         * the L7 Policy.  Only administrative users can specify a tenant UUID
         * other than their own. Changing this creates a new L7 Policy.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public L7PolicyV2Args build() {
            if ($.action == null) {
                throw new MissingRequiredPropertyException("L7PolicyV2Args", "action");
            }
            if ($.listenerId == null) {
                throw new MissingRequiredPropertyException("L7PolicyV2Args", "listenerId");
            }
            return $;
        }
    }

}