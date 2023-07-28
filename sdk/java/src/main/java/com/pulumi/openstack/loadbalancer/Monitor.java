// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.MonitorArgs;
import com.pulumi.openstack.loadbalancer.inputs.MonitorState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 monitor resource within OpenStack.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.loadbalancer.Monitor;
 * import com.pulumi.openstack.loadbalancer.MonitorArgs;
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
 *         var monitor1 = new Monitor(&#34;monitor1&#34;, MonitorArgs.builder()        
 *             .poolId(openstack_lb_pool_v2.pool_1().id())
 *             .type(&#34;PING&#34;)
 *             .delay(20)
 *             .timeout(10)
 *             .maxRetries(5)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Load Balancer Pool Monitor can be imported using the Monitor ID, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2
 * ```
 * 
 *  In case of using OpenContrail, the import may not work properly. If you face an issue, try to import the monitor providing its parent pool ID
 * 
 * ```sh
 *  $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2/708bc224-0f8c-4981-ac82-97095fe051b6
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/monitor:Monitor")
public class Monitor extends com.pulumi.resources.CustomResource {
    /**
     * The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Export(name="adminStateUp", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * The time, in seconds, between sending probes to members.
     * 
     */
    @Export(name="delay", type=Integer.class, parameters={})
    private Output<Integer> delay;

    /**
     * @return The time, in seconds, between sending probes to members.
     * 
     */
    public Output<Integer> delay() {
        return this.delay;
    }
    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * &#34;200&#34;, or a range like &#34;200-202&#34;.
     * 
     */
    @Export(name="expectedCodes", type=String.class, parameters={})
    private Output<String> expectedCodes;

    /**
     * @return Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * &#34;200&#34;, or a range like &#34;200-202&#34;.
     * 
     */
    public Output<String> expectedCodes() {
        return this.expectedCodes;
    }
    /**
     * Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it
     * defaults to &#34;GET&#34;.
     * 
     */
    @Export(name="httpMethod", type=String.class, parameters={})
    private Output<String> httpMethod;

    /**
     * @return Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it
     * defaults to &#34;GET&#34;.
     * 
     */
    public Output<String> httpMethod() {
        return this.httpMethod;
    }
    /**
     * Number of permissible ping failures before
     * changing the member&#39;s status to INACTIVE. Must be a number between 1
     * and 10.
     * 
     */
    @Export(name="maxRetries", type=Integer.class, parameters={})
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
     * Number of permissible ping failures befor changing the member&#39;s
     * status to ERROR. Must be a number between 1 and 10 (supported only in Octavia).
     * Changing this updates the max_retries_down of the existing monitor.
     * 
     */
    @Export(name="maxRetriesDown", type=Integer.class, parameters={})
    private Output<Integer> maxRetriesDown;

    /**
     * @return Number of permissible ping failures befor changing the member&#39;s
     * status to ERROR. Must be a number between 1 and 10 (supported only in Octavia).
     * Changing this updates the max_retries_down of the existing monitor.
     * 
     */
    public Output<Integer> maxRetriesDown() {
        return this.maxRetriesDown;
    }
    /**
     * The Name of the Monitor.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The Name of the Monitor.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The id of the pool that this monitor will be assigned to.
     * 
     */
    @Export(name="poolId", type=String.class, parameters={})
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
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     * 
     */
    @Export(name="tenantId", type=String.class, parameters={})
    private Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     * 
     */
    @Export(name="timeout", type=Integer.class, parameters={})
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
     * TLS-HELLO or UDP-CONNECT (supported only in Octavia), that is sent by the load
     * balancer to verify the member state. Changing this creates a new monitor.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The type of probe, which is PING, TCP, HTTP, HTTPS,
     * TLS-HELLO or UDP-CONNECT (supported only in Octavia), that is sent by the load
     * balancer to verify the member state. Changing this creates a new monitor.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS.
     * 
     */
    @Export(name="urlPath", type=String.class, parameters={})
    private Output<String> urlPath;

    /**
     * @return Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS.
     * 
     */
    public Output<String> urlPath() {
        return this.urlPath;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Monitor(String name) {
        this(name, MonitorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Monitor(String name, MonitorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Monitor(String name, MonitorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/monitor:Monitor", name, args == null ? MonitorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Monitor(String name, Output<String> id, @Nullable MonitorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/monitor:Monitor", name, state, makeResourceOptions(options, id));
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
    public static Monitor get(String name, Output<String> id, @Nullable MonitorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Monitor(name, id, state, options);
    }
}
