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


public final class MonitorArgs extends com.pulumi.resources.ResourceArgs {

    public static final MonitorArgs Empty = new MonitorArgs();

    /**
     * The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * The time, in seconds, between sending probes to members.
     * 
     */
    @Import(name="delay", required=true)
    private Output<Integer> delay;

    /**
     * @return The time, in seconds, between sending probes to members.
     * 
     */
    public Output<Integer> delay() {
        return this.delay;
    }

    /**
     * The domain name to use in the HTTP host header
     * health monitor requests. Supported in Octavia API version 2.10 or later.
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return The domain name to use in the HTTP host header
     * health monitor requests. Supported in Octavia API version 2.10 or later.
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * &#34;200&#34;, a list like &#34;200, 202&#34; or a range like &#34;200-202&#34;. Default is &#34;200&#34;.
     * 
     */
    @Import(name="expectedCodes")
    private @Nullable Output<String> expectedCodes;

    /**
     * @return Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * &#34;200&#34;, a list like &#34;200, 202&#34; or a range like &#34;200-202&#34;. Default is &#34;200&#34;.
     * 
     */
    public Optional<Output<String>> expectedCodes() {
        return Optional.ofNullable(this.expectedCodes);
    }

    /**
     * Required for HTTP(S) types. The HTTP method that
     * the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
     * OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET.
     * 
     */
    @Import(name="httpMethod")
    private @Nullable Output<String> httpMethod;

    /**
     * @return Required for HTTP(S) types. The HTTP method that
     * the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
     * OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET.
     * 
     */
    public Optional<Output<String>> httpMethod() {
        return Optional.ofNullable(this.httpMethod);
    }

    /**
     * Required for HTTP(S) types. The HTTP version that
     * the health monitor uses for requests. One of `1.0` or 1.1`is supported
     * for HTTP(S) monitors. The default is`1.0`. Supported in Octavia API version
     * 2.10 or later.
     * 
     */
    @Import(name="httpVersion")
    private @Nullable Output<String> httpVersion;

    /**
     * @return Required for HTTP(S) types. The HTTP version that
     * the health monitor uses for requests. One of `1.0` or 1.1`is supported
     * for HTTP(S) monitors. The default is`1.0`. Supported in Octavia API version
     * 2.10 or later.
     * 
     */
    public Optional<Output<String>> httpVersion() {
        return Optional.ofNullable(this.httpVersion);
    }

    /**
     * Number of permissible ping failures before
     * changing the member&#39;s status to INACTIVE. Must be a number between 1
     * and 10.
     * 
     */
    @Import(name="maxRetries", required=true)
    private Output<Integer> maxRetries;

    /**
     * @return Number of permissible ping failures before
     * changing the member&#39;s status to INACTIVE. Must be a number between 1
     * and 10.
     * 
     */
    public Output<Integer> maxRetries() {
        return this.maxRetries;
    }

    /**
     * Number of permissible ping failures before
     * changing the member&#39;s status to ERROR. Must be a number between 1 and 10.
     * The default is 3. Changing this updates the max_retries_down of the
     * existing monitor.
     * 
     */
    @Import(name="maxRetriesDown")
    private @Nullable Output<Integer> maxRetriesDown;

    /**
     * @return Number of permissible ping failures before
     * changing the member&#39;s status to ERROR. Must be a number between 1 and 10.
     * The default is 3. Changing this updates the max_retries_down of the
     * existing monitor.
     * 
     */
    public Optional<Output<Integer>> maxRetriesDown() {
        return Optional.ofNullable(this.maxRetriesDown);
    }

    /**
     * The Name of the Monitor.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The Name of the Monitor.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The id of the pool that this monitor will be assigned to.
     * 
     */
    @Import(name="poolId", required=true)
    private Output<String> poolId;

    /**
     * @return The id of the pool that this monitor will be assigned to.
     * 
     */
    public Output<String> poolId() {
        return this.poolId;
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a monitor. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a monitor. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     * 
     */
    @Import(name="timeout", required=true)
    private Output<Integer> timeout;

    /**
     * @return Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     * 
     */
    public Output<Integer> timeout() {
        return this.timeout;
    }

    /**
     * The type of probe, which is PING, TCP, HTTP, HTTPS,
     * TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
     * verify the member state. Changing this creates a new monitor.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of probe, which is PING, TCP, HTTP, HTTPS,
     * TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
     * verify the member state. Changing this creates a new monitor.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Default is `/`.
     * 
     */
    @Import(name="urlPath")
    private @Nullable Output<String> urlPath;

    /**
     * @return Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Default is `/`.
     * 
     */
    public Optional<Output<String>> urlPath() {
        return Optional.ofNullable(this.urlPath);
    }

    private MonitorArgs() {}

    private MonitorArgs(MonitorArgs $) {
        this.adminStateUp = $.adminStateUp;
        this.delay = $.delay;
        this.domainName = $.domainName;
        this.expectedCodes = $.expectedCodes;
        this.httpMethod = $.httpMethod;
        this.httpVersion = $.httpVersion;
        this.maxRetries = $.maxRetries;
        this.maxRetriesDown = $.maxRetriesDown;
        this.name = $.name;
        this.poolId = $.poolId;
        this.region = $.region;
        this.tenantId = $.tenantId;
        this.timeout = $.timeout;
        this.type = $.type;
        this.urlPath = $.urlPath;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MonitorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MonitorArgs $;

        public Builder() {
            $ = new MonitorArgs();
        }

        public Builder(MonitorArgs defaults) {
            $ = new MonitorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp The administrative state of the monitor.
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
         * @param adminStateUp The administrative state of the monitor.
         * A valid value is true (UP) or false (DOWN).
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param delay The time, in seconds, between sending probes to members.
         * 
         * @return builder
         * 
         */
        public Builder delay(Output<Integer> delay) {
            $.delay = delay;
            return this;
        }

        /**
         * @param delay The time, in seconds, between sending probes to members.
         * 
         * @return builder
         * 
         */
        public Builder delay(Integer delay) {
            return delay(Output.of(delay));
        }

        /**
         * @param domainName The domain name to use in the HTTP host header
         * health monitor requests. Supported in Octavia API version 2.10 or later.
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName The domain name to use in the HTTP host header
         * health monitor requests. Supported in Octavia API version 2.10 or later.
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param expectedCodes Required for HTTP(S) types. Expected HTTP codes
         * for a passing HTTP(S) monitor. You can either specify a single status like
         * &#34;200&#34;, a list like &#34;200, 202&#34; or a range like &#34;200-202&#34;. Default is &#34;200&#34;.
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
         * &#34;200&#34;, a list like &#34;200, 202&#34; or a range like &#34;200-202&#34;. Default is &#34;200&#34;.
         * 
         * @return builder
         * 
         */
        public Builder expectedCodes(String expectedCodes) {
            return expectedCodes(Output.of(expectedCodes));
        }

        /**
         * @param httpMethod Required for HTTP(S) types. The HTTP method that
         * the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
         * OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET.
         * 
         * @return builder
         * 
         */
        public Builder httpMethod(@Nullable Output<String> httpMethod) {
            $.httpMethod = httpMethod;
            return this;
        }

        /**
         * @param httpMethod Required for HTTP(S) types. The HTTP method that
         * the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
         * OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET.
         * 
         * @return builder
         * 
         */
        public Builder httpMethod(String httpMethod) {
            return httpMethod(Output.of(httpMethod));
        }

        /**
         * @param httpVersion Required for HTTP(S) types. The HTTP version that
         * the health monitor uses for requests. One of `1.0` or 1.1`is supported
         * for HTTP(S) monitors. The default is`1.0`. Supported in Octavia API version
         * 2.10 or later.
         * 
         * @return builder
         * 
         */
        public Builder httpVersion(@Nullable Output<String> httpVersion) {
            $.httpVersion = httpVersion;
            return this;
        }

        /**
         * @param httpVersion Required for HTTP(S) types. The HTTP version that
         * the health monitor uses for requests. One of `1.0` or 1.1`is supported
         * for HTTP(S) monitors. The default is`1.0`. Supported in Octavia API version
         * 2.10 or later.
         * 
         * @return builder
         * 
         */
        public Builder httpVersion(String httpVersion) {
            return httpVersion(Output.of(httpVersion));
        }

        /**
         * @param maxRetries Number of permissible ping failures before
         * changing the member&#39;s status to INACTIVE. Must be a number between 1
         * and 10.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(Output<Integer> maxRetries) {
            $.maxRetries = maxRetries;
            return this;
        }

        /**
         * @param maxRetries Number of permissible ping failures before
         * changing the member&#39;s status to INACTIVE. Must be a number between 1
         * and 10.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(Integer maxRetries) {
            return maxRetries(Output.of(maxRetries));
        }

        /**
         * @param maxRetriesDown Number of permissible ping failures before
         * changing the member&#39;s status to ERROR. Must be a number between 1 and 10.
         * The default is 3. Changing this updates the max_retries_down of the
         * existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder maxRetriesDown(@Nullable Output<Integer> maxRetriesDown) {
            $.maxRetriesDown = maxRetriesDown;
            return this;
        }

        /**
         * @param maxRetriesDown Number of permissible ping failures before
         * changing the member&#39;s status to ERROR. Must be a number between 1 and 10.
         * The default is 3. Changing this updates the max_retries_down of the
         * existing monitor.
         * 
         * @return builder
         * 
         */
        public Builder maxRetriesDown(Integer maxRetriesDown) {
            return maxRetriesDown(Output.of(maxRetriesDown));
        }

        /**
         * @param name The Name of the Monitor.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The Name of the Monitor.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param poolId The id of the pool that this monitor will be assigned to.
         * 
         * @return builder
         * 
         */
        public Builder poolId(Output<String> poolId) {
            $.poolId = poolId;
            return this;
        }

        /**
         * @param poolId The id of the pool that this monitor will be assigned to.
         * 
         * @return builder
         * 
         */
        public Builder poolId(String poolId) {
            return poolId(Output.of(poolId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a monitor. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * monitor.
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
         * A Networking client is needed to create a monitor. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * monitor.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param tenantId Required for admins. The UUID of the tenant who owns
         * the monitor.  Only administrative users can specify a tenant UUID
         * other than their own. Changing this creates a new monitor.
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
         * the monitor.  Only administrative users can specify a tenant UUID
         * other than their own. Changing this creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param timeout Maximum number of seconds for a monitor to wait for a
         * ping reply before it times out. The value must be less than the delay
         * value.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Output<Integer> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Maximum number of seconds for a monitor to wait for a
         * ping reply before it times out. The value must be less than the delay
         * value.
         * 
         * @return builder
         * 
         */
        public Builder timeout(Integer timeout) {
            return timeout(Output.of(timeout));
        }

        /**
         * @param type The type of probe, which is PING, TCP, HTTP, HTTPS,
         * TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
         * verify the member state. Changing this creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of probe, which is PING, TCP, HTTP, HTTPS,
         * TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
         * verify the member state. Changing this creates a new monitor.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param urlPath Required for HTTP(S) types. URI path that will be
         * accessed if monitor type is HTTP or HTTPS. Default is `/`.
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
         * accessed if monitor type is HTTP or HTTPS. Default is `/`.
         * 
         * @return builder
         * 
         */
        public Builder urlPath(String urlPath) {
            return urlPath(Output.of(urlPath));
        }

        public MonitorArgs build() {
            if ($.delay == null) {
                throw new MissingRequiredPropertyException("MonitorArgs", "delay");
            }
            if ($.maxRetries == null) {
                throw new MissingRequiredPropertyException("MonitorArgs", "maxRetries");
            }
            if ($.poolId == null) {
                throw new MissingRequiredPropertyException("MonitorArgs", "poolId");
            }
            if ($.timeout == null) {
                throw new MissingRequiredPropertyException("MonitorArgs", "timeout");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("MonitorArgs", "type");
            }
            return $;
        }
    }

}
