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
import java.lang.Object;
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
     * The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * A list of CIDR blocks that are permitted to connect to this listener, denying
     * all other source addresses. If not present, defaults to allow all.
     * 
     */
    @Export(name="allowedCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedCidrs;

    /**
     * @return A list of CIDR blocks that are permitted to connect to this listener, denying
     * all other source addresses. If not present, defaults to allow all.
     * 
     */
    public Output<Optional<List<String>>> allowedCidrs() {
        return Codegen.optional(this.allowedCidrs);
    }
    /**
     * The maximum number of connections allowed
     * for the Listener.
     * 
     */
    @Export(name="connectionLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> connectionLimit;

    /**
     * @return The maximum number of connections allowed
     * for the Listener.
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
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     * 
     */
    @Export(name="defaultTlsContainerRef", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultTlsContainerRef;

    /**
     * @return A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
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
     * The list of key value pairs representing headers to insert
     * into the request before it is sent to the backend members. Changing this updates the headers of the
     * existing listener.
     * 
     */
    @Export(name="insertHeaders", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> insertHeaders;

    /**
     * @return The list of key value pairs representing headers to insert
     * into the request before it is sent to the backend members. Changing this updates the headers of the
     * existing listener.
     * 
     */
    public Output<Optional<Map<String,Object>>> insertHeaders() {
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
     * Human-readable name for the Listener. Does not have
     * to be unique.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Human-readable name for the Listener. Does not have
     * to be unique.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The protocol - can either be TCP, HTTP, HTTPS,
     * TERMINATED_HTTPS, UDP (supported only in Octavia), SCTP (supported only
     * in **Octavia minor version &gt;= 2.23**) or PROMETHEUS (supported only in
     * **Octavia minor version &gt;=2.25**). Changing this creates a new Listener.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
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
    @Export(name="protocolPort", refs={Integer.class}, tree="[0]")
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
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     * 
     */
    @Export(name="sniContainerRefs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> sniContainerRefs;

    /**
     * @return A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     * 
     */
    public Output<Optional<List<String>>> sniContainerRefs() {
        return Codegen.optional(this.sniContainerRefs);
    }
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * The client inactivity timeout in milliseconds.
     * 
     */
    @Export(name="timeoutClientData", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeoutClientData;

    /**
     * @return The client inactivity timeout in milliseconds.
     * 
     */
    public Output<Integer> timeoutClientData() {
        return this.timeoutClientData;
    }
    /**
     * The member connection timeout in milliseconds.
     * 
     */
    @Export(name="timeoutMemberConnect", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeoutMemberConnect;

    /**
     * @return The member connection timeout in milliseconds.
     * 
     */
    public Output<Integer> timeoutMemberConnect() {
        return this.timeoutMemberConnect;
    }
    /**
     * The member inactivity timeout in milliseconds.
     * 
     */
    @Export(name="timeoutMemberData", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeoutMemberData;

    /**
     * @return The member inactivity timeout in milliseconds.
     * 
     */
    public Output<Integer> timeoutMemberData() {
        return this.timeoutMemberData;
    }
    /**
     * The time in milliseconds, to wait for additional
     * TCP packets for content inspection.
     * 
     */
    @Export(name="timeoutTcpInspect", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeoutTcpInspect;

    /**
     * @return The time in milliseconds, to wait for additional
     * TCP packets for content inspection.
     * 
     */
    public Output<Integer> timeoutTcpInspect() {
        return this.timeoutTcpInspect;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Listener(String name) {
        this(name, ListenerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Listener(String name, ListenerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Listener(String name, ListenerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/listener:Listener", name, args == null ? ListenerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Listener(String name, Output<String> id, @Nullable ListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/listener:Listener", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Listener get(String name, Output<String> id, @Nullable ListenerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Listener(name, id, state, options);
    }
}
