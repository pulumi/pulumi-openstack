// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MonitorV1State extends com.pulumi.resources.ResourceArgs {

    public static final MonitorV1State Empty = new MonitorV1State();

    /**
     * The administrative state of the monitor.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
     * state of the existing monitor.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<String> adminStateUp;

    /**
     * @return The administrative state of the monitor.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
     * state of the existing monitor.
     * 
     */
    public Optional<Output<String>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * The time, in seconds, between sending probes to members.
     * Changing this creates a new monitor.
     * 
     */
    @Import(name="delay")
    private @Nullable Output<Integer> delay;

    /**
     * @return The time, in seconds, between sending probes to members.
     * Changing this creates a new monitor.
     * 
     */
    public Optional<Output<Integer>> delay() {
        return Optional.ofNullable(this.delay);
    }

    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * &#34;200&#34;, or a range like &#34;200-202&#34;. Changing this updates the expected_codes
     * of the existing monitor.
     * 
     */
    @Import(name="expectedCodes")
    private @Nullable Output<String> expectedCodes;

    /**
     * @return Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * &#34;200&#34;, or a range like &#34;200-202&#34;. Changing this updates the expected_codes
     * of the existing monitor.
     * 
     */
    public Optional<Output<String>> expectedCodes() {
        return Optional.ofNullable(this.expectedCodes);
    }

    /**
     * Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it defaults
     * to &#34;GET&#34;. Changing this updates the http_method of the existing monitor.
     * 
     */
    @Import(name="httpMethod")
    private @Nullable Output<String> httpMethod;

    /**
     * @return Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it defaults
     * to &#34;GET&#34;. Changing this updates the http_method of the existing monitor.
     * 
     */
    public Optional<Output<String>> httpMethod() {
        return Optional.ofNullable(this.httpMethod);
    }

    /**
     * Number of permissible ping failures before changing
     * the member&#39;s status to INACTIVE. Must be a number between 1 and 10. Changing
     * this updates the max_retries of the existing monitor.
     * 
     */
    @Import(name="maxRetries")
    private @Nullable Output<Integer> maxRetries;

    /**
     * @return Number of permissible ping failures before changing
     * the member&#39;s status to INACTIVE. Must be a number between 1 and 10. Changing
     * this updates the max_retries of the existing monitor.
     * 
     */
    public Optional<Output<Integer>> maxRetries() {
        return Optional.ofNullable(this.maxRetries);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB monitor. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB monitor.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB monitor. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB monitor.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The owner of the monitor. Required if admin wants to
     * create a monitor for another tenant. Changing this creates a new monitor.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the monitor. Required if admin wants to
     * create a monitor for another tenant. Changing this creates a new monitor.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay value.
     * Changing this updates the timeout of the existing monitor.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<Integer> timeout;

    /**
     * @return Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay value.
     * Changing this updates the timeout of the existing monitor.
     * 
     */
    public Optional<Output<Integer>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    /**
     * The type of probe, which is PING, TCP, HTTP, or HTTPS,
     * that is sent by the monitor to verify the member state. Changing this
     * creates a new monitor.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of probe, which is PING, TCP, HTTP, or HTTPS,
     * that is sent by the monitor to verify the member state. Changing this
     * creates a new monitor.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Changing this updates the
     * url_path of the existing monitor.
     * 
     */
    @Import(name="urlPath")
    private @Nullable Output<String> urlPath;

    /**
     * @return Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Changing this updates the
     * url_path of the existing monitor.
     * 
     */
    public Optional<Output<String>> urlPath() {
        return Optional.ofNullable(this.urlPath);
    }

    private MonitorV1State() {}

    private MonitorV1State(MonitorV1State $) {
        this.adminStateUp = $.adminStateUp;
        this.delay = $.delay;
        this.expectedCodes = $.expectedCodes;
        this.httpMethod = $.httpMethod;
        this.maxRetries = $.maxRetries;
        this.region = $.region;
        this.tenantId = $.tenantId;
        this.timeout = $.timeout;
        this.type = $.type;
        this.urlPath = $.urlPath;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MonitorV1State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MonitorV1State $;

        public Builder() {
            $ = new MonitorV1State();
        }

        public Builder(MonitorV1State defaults) {
            $ = new MonitorV1State(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp The administrative state of the monitor.
         * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
         * state of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<String> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp The administrative state of the monitor.
         * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
         * state of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(String adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param delay The time, in seconds, between sending probes to members.
         * Changing this creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder delay(@Nullable Output<Integer> delay) {
            $.delay = delay;
            return this;
        }

        /**
         * @param delay The time, in seconds, between sending probes to members.
         * Changing this creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder delay(Integer delay) {
            return delay(Output.of(delay));
        }

        /**
         * @param expectedCodes Required for HTTP(S) types. Expected HTTP codes
         * for a passing HTTP(S) monitor. You can either specify a single status like
         * &#34;200&#34;, or a range like &#34;200-202&#34;. Changing this updates the expected_codes
         * of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder expectedCodes(@Nullable Output<String> expectedCodes) {
            $.expectedCodes = expectedCodes;
            return this;
        }

        /**
         * @param expectedCodes Required for HTTP(S) types. Expected HTTP codes
         * for a passing HTTP(S) monitor. You can either specify a single status like
         * &#34;200&#34;, or a range like &#34;200-202&#34;. Changing this updates the expected_codes
         * of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder expectedCodes(String expectedCodes) {
            return expectedCodes(Output.of(expectedCodes));
        }

        /**
         * @param httpMethod Required for HTTP(S) types. The HTTP method used
         * for requests by the monitor. If this attribute is not specified, it defaults
         * to &#34;GET&#34;. Changing this updates the http_method of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder httpMethod(@Nullable Output<String> httpMethod) {
            $.httpMethod = httpMethod;
            return this;
        }

        /**
         * @param httpMethod Required for HTTP(S) types. The HTTP method used
         * for requests by the monitor. If this attribute is not specified, it defaults
         * to &#34;GET&#34;. Changing this updates the http_method of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder httpMethod(String httpMethod) {
            return httpMethod(Output.of(httpMethod));
        }

        /**
         * @param maxRetries Number of permissible ping failures before changing
         * the member&#39;s status to INACTIVE. Must be a number between 1 and 10. Changing
         * this updates the max_retries of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(@Nullable Output<Integer> maxRetries) {
            $.maxRetries = maxRetries;
            return this;
        }

        /**
         * @param maxRetries Number of permissible ping failures before changing
         * the member&#39;s status to INACTIVE. Must be a number between 1 and 10. Changing
         * this updates the max_retries of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(Integer maxRetries) {
            return maxRetries(Output.of(maxRetries));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create an LB monitor. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * LB monitor.
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
         * A Networking client is needed to create an LB monitor. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * LB monitor.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tenantId The owner of the monitor. Required if admin wants to
         * create a monitor for another tenant. Changing this creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the monitor. Required if admin wants to
         * create a monitor for another tenant. Changing this creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param timeout Maximum number of seconds for a monitor to wait for a
         * ping reply before it times out. The value must be less than the delay value.
         * Changing this updates the timeout of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Maximum number of seconds for a monitor to wait for a
         * ping reply before it times out. The value must be less than the delay value.
         * Changing this updates the timeout of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        /**
         * @param type The type of probe, which is PING, TCP, HTTP, or HTTPS,
         * that is sent by the monitor to verify the member state. Changing this
         * creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of probe, which is PING, TCP, HTTP, or HTTPS,
         * that is sent by the monitor to verify the member state. Changing this
         * creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param urlPath Required for HTTP(S) types. URI path that will be
         * accessed if monitor type is HTTP or HTTPS. Changing this updates the
         * url_path of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder urlPath(@Nullable Output<String> urlPath) {
            $.urlPath = urlPath;
            return this;
        }

        /**
         * @param urlPath Required for HTTP(S) types. URI path that will be
         * accessed if monitor type is HTTP or HTTPS. Changing this updates the
         * url_path of the existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder urlPath(String urlPath) {
            return urlPath(Output.of(urlPath));
        }

        public MonitorV1State build() {
            return $;
        }
    }

}
