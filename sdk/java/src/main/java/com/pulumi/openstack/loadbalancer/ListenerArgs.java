// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ListenerArgs extends com.pulumi.resources.ResourceArgs {

    public static final ListenerArgs Empty = new ListenerArgs();

    /**
     * The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * A list of CIDR blocks that are permitted to connect to this listener, denying
     * all other source addresses. If not present, defaults to allow all.
     * 
     */
    @Import(name="allowedCidrs")
    private @Nullable Output<List<String>> allowedCidrs;

    /**
     * @return A list of CIDR blocks that are permitted to connect to this listener, denying
     * all other source addresses. If not present, defaults to allow all.
     * 
     */
    public Optional<Output<List<String>>> allowedCidrs() {
        return Optional.ofNullable(this.allowedCidrs);
    }

    /**
     * The maximum number of connections allowed
     * for the Listener.
     * 
     */
    @Import(name="connectionLimit")
    private @Nullable Output<Integer> connectionLimit;

    /**
     * @return The maximum number of connections allowed
     * for the Listener.
     * 
     */
    public Optional<Output<Integer>> connectionLimit() {
        return Optional.ofNullable(this.connectionLimit);
    }

    /**
     * The ID of the default pool with which the
     * Listener is associated.
     * 
     */
    @Import(name="defaultPoolId")
    private @Nullable Output<String> defaultPoolId;

    /**
     * @return The ID of the default pool with which the
     * Listener is associated.
     * 
     */
    public Optional<Output<String>> defaultPoolId() {
        return Optional.ofNullable(this.defaultPoolId);
    }

    /**
     * A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     * 
     */
    @Import(name="defaultTlsContainerRef")
    private @Nullable Output<String> defaultTlsContainerRef;

    /**
     * @return A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     * 
     */
    public Optional<Output<String>> defaultTlsContainerRef() {
        return Optional.ofNullable(this.defaultTlsContainerRef);
    }

    /**
     * Human-readable description for the Listener.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description for the Listener.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The list of key value pairs representing headers to insert
     * into the request before it is sent to the backend members. Changing this updates the headers of the
     * existing listener.
     * 
     */
    @Import(name="insertHeaders")
    private @Nullable Output<Map<String,Object>> insertHeaders;

    /**
     * @return The list of key value pairs representing headers to insert
     * into the request before it is sent to the backend members. Changing this updates the headers of the
     * existing listener.
     * 
     */
    public Optional<Output<Map<String,Object>>> insertHeaders() {
        return Optional.ofNullable(this.insertHeaders);
    }

    /**
     * The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     * 
     */
    @Import(name="loadbalancerId", required=true)
    private Output<String> loadbalancerId;

    /**
     * @return The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     * 
     */
    public Output<String> loadbalancerId() {
        return this.loadbalancerId;
    }

    /**
     * Human-readable name for the Listener. Does not have
     * to be unique.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Human-readable name for the Listener. Does not have
     * to be unique.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The protocol - can either be TCP, HTTP, HTTPS,
     * TERMINATED_HTTPS, UDP (supported only in Octavia), SCTP (supported only
     * in **Octavia minor version &gt;= 2.23**) or PROMETHEUS (supported only in
     * **Octavia minor version &gt;=2.25**). Changing this creates a new Listener.
     * 
     */
    @Import(name="protocol", required=true)
    private Output<String> protocol;

    /**
     * @return The protocol - can either be TCP, HTTP, HTTPS,
     * TERMINATED_HTTPS, UDP (supported only in Octavia), SCTP (supported only
     * in **Octavia minor version &gt;= 2.23**) or PROMETHEUS (supported only in
     * **Octavia minor version &gt;=2.25**). Changing this creates a new Listener.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }

    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new Listener.
     * 
     */
    @Import(name="protocolPort", required=true)
    private Output<Integer> protocolPort;

    /**
     * @return The port on which to listen for client traffic.
     * Changing this creates a new Listener.
     * 
     */
    public Output<Integer> protocolPort() {
        return this.protocolPort;
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     * 
     */
    @Import(name="sniContainerRefs")
    private @Nullable Output<List<String>> sniContainerRefs;

    /**
     * @return A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     * 
     */
    public Optional<Output<List<String>>> sniContainerRefs() {
        return Optional.ofNullable(this.sniContainerRefs);
    }

    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * The client inactivity timeout in milliseconds.
     * 
     */
    @Import(name="timeoutClientData")
    private @Nullable Output<Integer> timeoutClientData;

    /**
     * @return The client inactivity timeout in milliseconds.
     * 
     */
    public Optional<Output<Integer>> timeoutClientData() {
        return Optional.ofNullable(this.timeoutClientData);
    }

    /**
     * The member connection timeout in milliseconds.
     * 
     */
    @Import(name="timeoutMemberConnect")
    private @Nullable Output<Integer> timeoutMemberConnect;

    /**
     * @return The member connection timeout in milliseconds.
     * 
     */
    public Optional<Output<Integer>> timeoutMemberConnect() {
        return Optional.ofNullable(this.timeoutMemberConnect);
    }

    /**
     * The member inactivity timeout in milliseconds.
     * 
     */
    @Import(name="timeoutMemberData")
    private @Nullable Output<Integer> timeoutMemberData;

    /**
     * @return The member inactivity timeout in milliseconds.
     * 
     */
    public Optional<Output<Integer>> timeoutMemberData() {
        return Optional.ofNullable(this.timeoutMemberData);
    }

    /**
     * The time in milliseconds, to wait for additional
     * TCP packets for content inspection.
     * 
     */
    @Import(name="timeoutTcpInspect")
    private @Nullable Output<Integer> timeoutTcpInspect;

    /**
     * @return The time in milliseconds, to wait for additional
     * TCP packets for content inspection.
     * 
     */
    public Optional<Output<Integer>> timeoutTcpInspect() {
        return Optional.ofNullable(this.timeoutTcpInspect);
    }

    private ListenerArgs() {}

    private ListenerArgs(ListenerArgs $) {
        this.adminStateUp = $.adminStateUp;
        this.allowedCidrs = $.allowedCidrs;
        this.connectionLimit = $.connectionLimit;
        this.defaultPoolId = $.defaultPoolId;
        this.defaultTlsContainerRef = $.defaultTlsContainerRef;
        this.description = $.description;
        this.insertHeaders = $.insertHeaders;
        this.loadbalancerId = $.loadbalancerId;
        this.name = $.name;
        this.protocol = $.protocol;
        this.protocolPort = $.protocolPort;
        this.region = $.region;
        this.sniContainerRefs = $.sniContainerRefs;
        this.tenantId = $.tenantId;
        this.timeoutClientData = $.timeoutClientData;
        this.timeoutMemberConnect = $.timeoutMemberConnect;
        this.timeoutMemberData = $.timeoutMemberData;
        this.timeoutTcpInspect = $.timeoutTcpInspect;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerArgs $;

        public Builder() {
            $ = new ListenerArgs();
        }

        public Builder(ListenerArgs defaults) {
            $ = new ListenerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp The administrative state of the Listener.
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
         * @param adminStateUp The administrative state of the Listener.
         * A valid value is true (UP) or false (DOWN).
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param allowedCidrs A list of CIDR blocks that are permitted to connect to this listener, denying
         * all other source addresses. If not present, defaults to allow all.
         * 
         * @return builder
         * 
         */
        public Builder allowedCidrs(@Nullable Output<List<String>> allowedCidrs) {
            $.allowedCidrs = allowedCidrs;
            return this;
        }

        /**
         * @param allowedCidrs A list of CIDR blocks that are permitted to connect to this listener, denying
         * all other source addresses. If not present, defaults to allow all.
         * 
         * @return builder
         * 
         */
        public Builder allowedCidrs(List<String> allowedCidrs) {
            return allowedCidrs(Output.of(allowedCidrs));
        }

        /**
         * @param allowedCidrs A list of CIDR blocks that are permitted to connect to this listener, denying
         * all other source addresses. If not present, defaults to allow all.
         * 
         * @return builder
         * 
         */
        public Builder allowedCidrs(String... allowedCidrs) {
            return allowedCidrs(List.of(allowedCidrs));
        }

        /**
         * @param connectionLimit The maximum number of connections allowed
         * for the Listener.
         * 
         * @return builder
         * 
         */
        public Builder connectionLimit(@Nullable Output<Integer> connectionLimit) {
            $.connectionLimit = connectionLimit;
            return this;
        }

        /**
         * @param connectionLimit The maximum number of connections allowed
         * for the Listener.
         * 
         * @return builder
         * 
         */
        public Builder connectionLimit(Integer connectionLimit) {
            return connectionLimit(Output.of(connectionLimit));
        }

        /**
         * @param defaultPoolId The ID of the default pool with which the
         * Listener is associated.
         * 
         * @return builder
         * 
         */
        public Builder defaultPoolId(@Nullable Output<String> defaultPoolId) {
            $.defaultPoolId = defaultPoolId;
            return this;
        }

        /**
         * @param defaultPoolId The ID of the default pool with which the
         * Listener is associated.
         * 
         * @return builder
         * 
         */
        public Builder defaultPoolId(String defaultPoolId) {
            return defaultPoolId(Output.of(defaultPoolId));
        }

        /**
         * @param defaultTlsContainerRef A reference to a Barbican Secrets
         * container which stores TLS information. This is required if the protocol
         * is `TERMINATED_HTTPS`. See
         * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
         * for more information.
         * 
         * @return builder
         * 
         */
        public Builder defaultTlsContainerRef(@Nullable Output<String> defaultTlsContainerRef) {
            $.defaultTlsContainerRef = defaultTlsContainerRef;
            return this;
        }

        /**
         * @param defaultTlsContainerRef A reference to a Barbican Secrets
         * container which stores TLS information. This is required if the protocol
         * is `TERMINATED_HTTPS`. See
         * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
         * for more information.
         * 
         * @return builder
         * 
         */
        public Builder defaultTlsContainerRef(String defaultTlsContainerRef) {
            return defaultTlsContainerRef(Output.of(defaultTlsContainerRef));
        }

        /**
         * @param description Human-readable description for the Listener.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description for the Listener.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param insertHeaders The list of key value pairs representing headers to insert
         * into the request before it is sent to the backend members. Changing this updates the headers of the
         * existing listener.
         * 
         * @return builder
         * 
         */
        public Builder insertHeaders(@Nullable Output<Map<String,Object>> insertHeaders) {
            $.insertHeaders = insertHeaders;
            return this;
        }

        /**
         * @param insertHeaders The list of key value pairs representing headers to insert
         * into the request before it is sent to the backend members. Changing this updates the headers of the
         * existing listener.
         * 
         * @return builder
         * 
         */
        public Builder insertHeaders(Map<String,Object> insertHeaders) {
            return insertHeaders(Output.of(insertHeaders));
        }

        /**
         * @param loadbalancerId The load balancer on which to provision this
         * Listener. Changing this creates a new Listener.
         * 
         * @return builder
         * 
         */
        public Builder loadbalancerId(Output<String> loadbalancerId) {
            $.loadbalancerId = loadbalancerId;
            return this;
        }

        /**
         * @param loadbalancerId The load balancer on which to provision this
         * Listener. Changing this creates a new Listener.
         * 
         * @return builder
         * 
         */
        public Builder loadbalancerId(String loadbalancerId) {
            return loadbalancerId(Output.of(loadbalancerId));
        }

        /**
         * @param name Human-readable name for the Listener. Does not have
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
         * @param name Human-readable name for the Listener. Does not have
         * to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param protocol The protocol - can either be TCP, HTTP, HTTPS,
         * TERMINATED_HTTPS, UDP (supported only in Octavia), SCTP (supported only
         * in **Octavia minor version &gt;= 2.23**) or PROMETHEUS (supported only in
         * **Octavia minor version &gt;=2.25**). Changing this creates a new Listener.
         * 
         * @return builder
         * 
         */
        public Builder protocol(Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol The protocol - can either be TCP, HTTP, HTTPS,
         * TERMINATED_HTTPS, UDP (supported only in Octavia), SCTP (supported only
         * in **Octavia minor version &gt;= 2.23**) or PROMETHEUS (supported only in
         * **Octavia minor version &gt;=2.25**). Changing this creates a new Listener.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param protocolPort The port on which to listen for client traffic.
         * Changing this creates a new Listener.
         * 
         * @return builder
         * 
         */
        public Builder protocolPort(Output<Integer> protocolPort) {
            $.protocolPort = protocolPort;
            return this;
        }

        /**
         * @param protocolPort The port on which to listen for client traffic.
         * Changing this creates a new Listener.
         * 
         * @return builder
         * 
         */
        public Builder protocolPort(Integer protocolPort) {
            return protocolPort(Output.of(protocolPort));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create an . If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * Listener.
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
         * Listener.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param sniContainerRefs A list of references to Barbican Secrets
         * containers which store SNI information. See
         * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
         * for more information.
         * 
         * @return builder
         * 
         */
        public Builder sniContainerRefs(@Nullable Output<List<String>> sniContainerRefs) {
            $.sniContainerRefs = sniContainerRefs;
            return this;
        }

        /**
         * @param sniContainerRefs A list of references to Barbican Secrets
         * containers which store SNI information. See
         * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
         * for more information.
         * 
         * @return builder
         * 
         */
        public Builder sniContainerRefs(List<String> sniContainerRefs) {
            return sniContainerRefs(Output.of(sniContainerRefs));
        }

        /**
         * @param sniContainerRefs A list of references to Barbican Secrets
         * containers which store SNI information. See
         * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
         * for more information.
         * 
         * @return builder
         * 
         */
        public Builder sniContainerRefs(String... sniContainerRefs) {
            return sniContainerRefs(List.of(sniContainerRefs));
        }

        /**
         * @param tenantId Required for admins. The UUID of the tenant who owns
         * the Listener.  Only administrative users can specify a tenant UUID
         * other than their own. Changing this creates a new Listener.
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
         * the Listener.  Only administrative users can specify a tenant UUID
         * other than their own. Changing this creates a new Listener.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param timeoutClientData The client inactivity timeout in milliseconds.
         * 
         * @return builder
         * 
         */
        public Builder timeoutClientData(@Nullable Output<Integer> timeoutClientData) {
            $.timeoutClientData = timeoutClientData;
            return this;
        }

        /**
         * @param timeoutClientData The client inactivity timeout in milliseconds.
         * 
         * @return builder
         * 
         */
        public Builder timeoutClientData(Integer timeoutClientData) {
            return timeoutClientData(Output.of(timeoutClientData));
        }

        /**
         * @param timeoutMemberConnect The member connection timeout in milliseconds.
         * 
         * @return builder
         * 
         */
        public Builder timeoutMemberConnect(@Nullable Output<Integer> timeoutMemberConnect) {
            $.timeoutMemberConnect = timeoutMemberConnect;
            return this;
        }

        /**
         * @param timeoutMemberConnect The member connection timeout in milliseconds.
         * 
         * @return builder
         * 
         */
        public Builder timeoutMemberConnect(Integer timeoutMemberConnect) {
            return timeoutMemberConnect(Output.of(timeoutMemberConnect));
        }

        /**
         * @param timeoutMemberData The member inactivity timeout in milliseconds.
         * 
         * @return builder
         * 
         */
        public Builder timeoutMemberData(@Nullable Output<Integer> timeoutMemberData) {
            $.timeoutMemberData = timeoutMemberData;
            return this;
        }

        /**
         * @param timeoutMemberData The member inactivity timeout in milliseconds.
         * 
         * @return builder
         * 
         */
        public Builder timeoutMemberData(Integer timeoutMemberData) {
            return timeoutMemberData(Output.of(timeoutMemberData));
        }

        /**
         * @param timeoutTcpInspect The time in milliseconds, to wait for additional
         * TCP packets for content inspection.
         * 
         * @return builder
         * 
         */
        public Builder timeoutTcpInspect(@Nullable Output<Integer> timeoutTcpInspect) {
            $.timeoutTcpInspect = timeoutTcpInspect;
            return this;
        }

        /**
         * @param timeoutTcpInspect The time in milliseconds, to wait for additional
         * TCP packets for content inspection.
         * 
         * @return builder
         * 
         */
        public Builder timeoutTcpInspect(Integer timeoutTcpInspect) {
            return timeoutTcpInspect(Output.of(timeoutTcpInspect));
        }

        public ListenerArgs build() {
            $.loadbalancerId = Objects.requireNonNull($.loadbalancerId, "expected parameter 'loadbalancerId' to be non-null");
            $.protocol = Objects.requireNonNull($.protocol, "expected parameter 'protocol' to be non-null");
            $.protocolPort = Objects.requireNonNull($.protocolPort, "expected parameter 'protocolPort' to be non-null");
            return $;
        }
    }

}
