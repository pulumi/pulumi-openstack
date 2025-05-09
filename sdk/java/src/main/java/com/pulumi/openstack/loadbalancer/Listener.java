// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.ListenerArgs;
import com.pulumi.openstack.loadbalancer.inputs.ListenerState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 listener resource within OpenStack.
 * 
 * &gt; **Note:** This resource has attributes that depend on octavia minor versions.
 * Please ensure your Openstack cloud supports the required minor version.
 * 
 * ## Example Usage
 * 
 * ### Simple listener
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.loadbalancer.Listener;
 * import com.pulumi.openstack.loadbalancer.ListenerArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var listener1 = new Listener("listener1", ListenerArgs.builder()
 *             .protocol("HTTP")
 *             .protocolPort(8080)
 *             .loadbalancerId("d9415786-5f1a-428b-b35f-2f1523e146d2")
 *             .insertHeaders(Map.of("X-Forwarded-For", "true"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Listener with TLS and client certificate authentication
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.keymanager.SecretV1;
 * import com.pulumi.openstack.keymanager.SecretV1Args;
 * import com.pulumi.std.StdFunctions;
 * import com.pulumi.std.inputs.Filebase64Args;
 * import com.pulumi.std.inputs.FileArgs;
 * import com.pulumi.openstack.networking.NetworkingFunctions;
 * import com.pulumi.openstack.networking.inputs.GetSubnetArgs;
 * import com.pulumi.openstack.loadbalancer.LoadBalancer;
 * import com.pulumi.openstack.loadbalancer.LoadBalancerArgs;
 * import com.pulumi.openstack.loadbalancer.Listener;
 * import com.pulumi.openstack.loadbalancer.ListenerArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var certificate1 = new SecretV1("certificate1", SecretV1Args.builder()
 *             .name("certificate")
 *             .payload(StdFunctions.filebase64(Filebase64Args.builder()
 *                 .input("snakeoil.p12")
 *                 .build()).result())
 *             .payloadContentEncoding("base64")
 *             .payloadContentType("application/octet-stream")
 *             .build());
 * 
 *         var caCertificate1 = new SecretV1("caCertificate1", SecretV1Args.builder()
 *             .name("certificate")
 *             .payload(StdFunctions.file(FileArgs.builder()
 *                 .input("CA.pem")
 *                 .build()).result())
 *             .secretType("certificate")
 *             .payloadContentType("text/plain")
 *             .build());
 * 
 *         final var subnet1 = NetworkingFunctions.getSubnet(GetSubnetArgs.builder()
 *             .name("my-subnet")
 *             .build());
 * 
 *         var lb1 = new LoadBalancer("lb1", LoadBalancerArgs.builder()
 *             .name("loadbalancer")
 *             .vipSubnetId(subnet1.id())
 *             .build());
 * 
 *         var listener1 = new Listener("listener1", ListenerArgs.builder()
 *             .name("https")
 *             .protocol("TERMINATED_HTTPS")
 *             .protocolPort(443)
 *             .loadbalancerId(lb1.id())
 *             .defaultTlsContainerRef(certificate1)
 *             .clientAuthentication("OPTIONAL")
 *             .clientCaTlsContainerRef(caCertificate2.secretRef())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Load Balancer Listener can be imported using the Listener ID, e.g.:
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/listener:Listener listener_1 b67ce64e-8b26-405d-afeb-4a078901f15a
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/listener:Listener")
public class Listener extends com.pulumi.resources.CustomResource {
    /**
     * The administrative state of the Listener. A
     * valid value is true (UP) or false (DOWN).
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the Listener. A
     * valid value is true (UP) or false (DOWN).
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * A list of CIDR blocks that are permitted to
     * connect to this listener, denying all other source addresses. If not present,
     * defaults to allow all.
     * 
     */
    @Export(name="allowedCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedCidrs;

    /**
     * @return A list of CIDR blocks that are permitted to
     * connect to this listener, denying all other source addresses. If not present,
     * defaults to allow all.
     * 
     */
    public Output<Optional<List<String>>> allowedCidrs() {
        return Codegen.optional(this.allowedCidrs);
    }
    /**
     * A list of ALPN protocols. Available protocols:
     * `http/1.0`, `http/1.1`, `h2`. Supported only in **Octavia minor version &gt;=
     * 2.20**.
     * 
     */
    @Export(name="alpnProtocols", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> alpnProtocols;

    /**
     * @return A list of ALPN protocols. Available protocols:
     * `http/1.0`, `http/1.1`, `h2`. Supported only in **Octavia minor version &gt;=
     * 2.20**.
     * 
     */
    public Output<List<String>> alpnProtocols() {
        return this.alpnProtocols;
    }
    /**
     * The TLS client authentication mode.
     * Available options: `NONE`, `OPTIONAL` or `MANDATORY`. Requires
     * `TERMINATED_HTTPS` listener protocol and the `client_ca_tls_container_ref`.
     * Supported only in **Octavia minor version &gt;= 2.8**.
     * 
     */
    @Export(name="clientAuthentication", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientAuthentication;

    /**
     * @return The TLS client authentication mode.
     * Available options: `NONE`, `OPTIONAL` or `MANDATORY`. Requires
     * `TERMINATED_HTTPS` listener protocol and the `client_ca_tls_container_ref`.
     * Supported only in **Octavia minor version &gt;= 2.8**.
     * 
     */
    public Output<Optional<String>> clientAuthentication() {
        return Codegen.optional(this.clientAuthentication);
    }
    /**
     * The ref of the key manager service
     * secret containing a PEM format client CA certificate bundle for
     * `TERMINATED_HTTPS` listeners. Required if `client_authentication` is
     * `OPTIONAL` or `MANDATORY`. Supported only in **Octavia minor version &gt;=
     * 2.8**.
     * 
     */
    @Export(name="clientCaTlsContainerRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientCaTlsContainerRef;

    /**
     * @return The ref of the key manager service
     * secret containing a PEM format client CA certificate bundle for
     * `TERMINATED_HTTPS` listeners. Required if `client_authentication` is
     * `OPTIONAL` or `MANDATORY`. Supported only in **Octavia minor version &gt;=
     * 2.8**.
     * 
     */
    public Output<Optional<String>> clientCaTlsContainerRef() {
        return Codegen.optional(this.clientCaTlsContainerRef);
    }
    /**
     * The URI of the key manager service
     * secret containing a PEM format CA revocation list file for `TERMINATED_HTTPS`
     * listeners. Supported only in **Octavia minor version &gt;= 2.8**.
     * 
     */
    @Export(name="clientCrlContainerRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientCrlContainerRef;

    /**
     * @return The URI of the key manager service
     * secret containing a PEM format CA revocation list file for `TERMINATED_HTTPS`
     * listeners. Supported only in **Octavia minor version &gt;= 2.8**.
     * 
     */
    public Output<Optional<String>> clientCrlContainerRef() {
        return Codegen.optional(this.clientCrlContainerRef);
    }
    /**
     * The maximum number of connections allowed for
     * the Listener.
     * 
     */
    @Export(name="connectionLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> connectionLimit;

    /**
     * @return The maximum number of connections allowed for
     * the Listener.
     * 
     */
    public Output<Integer> connectionLimit() {
        return this.connectionLimit;
    }
    /**
     * The ID of the default pool with which the
     * Listener is associated.
     * 
     */
    @Export(name="defaultPoolId", refs={String.class}, tree="[0]")
    private Output<String> defaultPoolId;

    /**
     * @return The ID of the default pool with which the
     * Listener is associated.
     * 
     */
    public Output<String> defaultPoolId() {
        return this.defaultPoolId;
    }
    /**
     * A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol is
     * `TERMINATED_HTTPS`. See
     * [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
     * for more information.
     * 
     */
    @Export(name="defaultTlsContainerRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultTlsContainerRef;

    /**
     * @return A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol is
     * `TERMINATED_HTTPS`. See
     * [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
     * for more information.
     * 
     */
    public Output<Optional<String>> defaultTlsContainerRef() {
        return Codegen.optional(this.defaultTlsContainerRef);
    }
    /**
     * Human-readable description for the Listener.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-readable description for the Listener.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Defines whether the
     * **includeSubDomains** directive should be added to the
     * Strict-Transport-Security HTTP response header. This requires setting the
     * `hsts_max_age` option as well in order to become effective. Requires
     * `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
     * version &gt;= 2.27**.
     * 
     */
    @Export(name="hstsIncludeSubdomains", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hstsIncludeSubdomains;

    /**
     * @return Defines whether the
     * **includeSubDomains** directive should be added to the
     * Strict-Transport-Security HTTP response header. This requires setting the
     * `hsts_max_age` option as well in order to become effective. Requires
     * `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
     * version &gt;= 2.27**.
     * 
     */
    public Output<Optional<Boolean>> hstsIncludeSubdomains() {
        return Codegen.optional(this.hstsIncludeSubdomains);
    }
    /**
     * The value of the **max_age** directive for the
     * Strict-Transport-Security HTTP response header. Setting this enables HTTP
     * Strict Transport Security (HSTS) for the TLS-terminated listener. Requires
     * `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
     * version &gt;= 2.27**.
     * 
     */
    @Export(name="hstsMaxAge", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> hstsMaxAge;

    /**
     * @return The value of the **max_age** directive for the
     * Strict-Transport-Security HTTP response header. Setting this enables HTTP
     * Strict Transport Security (HSTS) for the TLS-terminated listener. Requires
     * `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia minor
     * version &gt;= 2.27**.
     * 
     */
    public Output<Optional<Integer>> hstsMaxAge() {
        return Codegen.optional(this.hstsMaxAge);
    }
    /**
     * Defines whether the **preload** directive should
     * be added to the Strict-Transport-Security HTTP response header. This requires
     * setting the `hsts_max_age` option as well in order to become effective.
     * Requires `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia
     * minor version &gt;= 2.27**.
     * 
     */
    @Export(name="hstsPreload", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hstsPreload;

    /**
     * @return Defines whether the **preload** directive should
     * be added to the Strict-Transport-Security HTTP response header. This requires
     * setting the `hsts_max_age` option as well in order to become effective.
     * Requires `TERMINATED_HTTPS` listener protocol. Supported only in **Octavia
     * minor version &gt;= 2.27**.
     * 
     */
    public Output<Optional<Boolean>> hstsPreload() {
        return Codegen.optional(this.hstsPreload);
    }
    /**
     * The list of key value pairs representing
     * headers to insert into the request before it is sent to the backend members.
     * Changing this updates the headers of the existing listener.
     * 
     */
    @Export(name="insertHeaders", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> insertHeaders;

    /**
     * @return The list of key value pairs representing
     * headers to insert into the request before it is sent to the backend members.
     * Changing this updates the headers of the existing listener.
     * 
     */
    public Output<Optional<Map<String,String>>> insertHeaders() {
        return Codegen.optional(this.insertHeaders);
    }
    /**
     * The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     * 
     */
    @Export(name="loadbalancerId", refs={String.class}, tree="[0]")
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
     * Human-readable name for the Listener. Does not have to be
     * unique.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Human-readable name for the Listener. Does not have to be
     * unique.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The protocol can be either `TCP`, `HTTP`, `HTTPS`,
     * `TERMINATED_HTTPS`, `UDP`, `SCTP` (supported only in **Octavia minor version
     * \&gt;= 2.23**), or `PROMETHEUS` (supported only in **Octavia minor version &gt;=
     * 2.25**). Changing this creates a new Listener.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return The protocol can be either `TCP`, `HTTP`, `HTTPS`,
     * `TERMINATED_HTTPS`, `UDP`, `SCTP` (supported only in **Octavia minor version
     * \&gt;= 2.23**), or `PROMETHEUS` (supported only in **Octavia minor version &gt;=
     * 2.25**). Changing this creates a new Listener.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The port on which to listen for client traffic.
     * * Changing this creates a new Listener.
     * 
     */
    @Export(name="protocolPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> protocolPort;

    /**
     * @return The port on which to listen for client traffic.
     * * Changing this creates a new Listener.
     * 
     */
    public Output<Integer> protocolPort() {
        return this.protocolPort;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a listener. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new Listener.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a listener. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new Listener.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
     * for more information.
     * 
     */
    @Export(name="sniContainerRefs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> sniContainerRefs;

    /**
     * @return A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://docs.openstack.org/octavia/latest/user/guides/basic-cookbook.html#deploy-a-tls-terminated-https-load-balancer)
     * for more information.
     * 
     */
    public Output<Optional<List<String>>> sniContainerRefs() {
        return Codegen.optional(this.sniContainerRefs);
    }
    /**
     * A list of simple strings assigned to the pool. Available
     * for Octavia **minor version 2.5 or later**.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of simple strings assigned to the pool. Available
     * for Octavia **minor version 2.5 or later**.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID other than
     * their own. Changing this creates a new Listener.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID other than
     * their own. Changing this creates a new Listener.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * The client inactivity timeout in
     * milliseconds.
     * 
     */
    @Export(name="timeoutClientData", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeoutClientData;

    /**
     * @return The client inactivity timeout in
     * milliseconds.
     * 
     */
    public Output<Integer> timeoutClientData() {
        return this.timeoutClientData;
    }
    /**
     * The member connection timeout in
     * milliseconds.
     * 
     */
    @Export(name="timeoutMemberConnect", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeoutMemberConnect;

    /**
     * @return The member connection timeout in
     * milliseconds.
     * 
     */
    public Output<Integer> timeoutMemberConnect() {
        return this.timeoutMemberConnect;
    }
    /**
     * The member inactivity timeout in
     * milliseconds.
     * 
     */
    @Export(name="timeoutMemberData", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeoutMemberData;

    /**
     * @return The member inactivity timeout in
     * milliseconds.
     * 
     */
    public Output<Integer> timeoutMemberData() {
        return this.timeoutMemberData;
    }
    /**
     * The time in milliseconds, to wait for
     * additional TCP packets for content inspection.
     * 
     */
    @Export(name="timeoutTcpInspect", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeoutTcpInspect;

    /**
     * @return The time in milliseconds, to wait for
     * additional TCP packets for content inspection.
     * 
     */
    public Output<Integer> timeoutTcpInspect() {
        return this.timeoutTcpInspect;
    }
    /**
     * List of ciphers in OpenSSL format
     * (colon-separated). See
     * https://www.openssl.org/docs/man1.1.1/man1/ciphers.html for more information.
     * Supported only in **Octavia minor version &gt;= 2.15**.
     * 
     */
    @Export(name="tlsCiphers", refs={String.class}, tree="[0]")
    private Output<String> tlsCiphers;

    /**
     * @return List of ciphers in OpenSSL format
     * (colon-separated). See
     * https://www.openssl.org/docs/man1.1.1/man1/ciphers.html for more information.
     * Supported only in **Octavia minor version &gt;= 2.15**.
     * 
     */
    public Output<String> tlsCiphers() {
        return this.tlsCiphers;
    }
    /**
     * A list of TLS protocol versions. Available
     * versions: `TLSv1`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Supported only in
     * **Octavia minor version &gt;= 2.17**.
     * 
     */
    @Export(name="tlsVersions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tlsVersions;

    /**
     * @return A list of TLS protocol versions. Available
     * versions: `TLSv1`, `TLSv1.1`, `TLSv1.2`, `TLSv1.3`. Supported only in
     * **Octavia minor version &gt;= 2.17**.
     * 
     */
    public Output<List<String>> tlsVersions() {
        return this.tlsVersions;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Listener(java.lang.String name) {
        this(name, ListenerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Listener(java.lang.String name, ListenerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Listener(java.lang.String name, ListenerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/listener:Listener", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Listener(java.lang.String name, Output<java.lang.String> id, @Nullable ListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/listener:Listener", name, state, makeResourceOptions(options, id), false);
    }

    private static ListenerArgs makeArgs(ListenerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ListenerArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Listener get(java.lang.String name, Output<java.lang.String> id, @Nullable ListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Listener(name, id, state, options);
    }
}
