// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.MonitorV1Args;
import com.pulumi.openstack.loadbalancer.inputs.MonitorV1State;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 load balancer monitor resource within OpenStack.
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
 * import com.pulumi.openstack.loadbalancer.MonitorV1;
 * import com.pulumi.openstack.loadbalancer.MonitorV1Args;
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
 *         var monitor1 = new MonitorV1("monitor1", MonitorV1Args.builder()
 *             .type("PING")
 *             .delay(30)
 *             .timeout(5)
 *             .maxRetries(3)
 *             .adminStateUp("true")
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
 * Load Balancer Members can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/monitorV1:MonitorV1 monitor_1 119d7530-72e9-449a-aa97-124a5ef1992c
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/monitorV1:MonitorV1")
public class MonitorV1 extends com.pulumi.resources.CustomResource {
    /**
     * The administrative state of the monitor.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
     * state of the existing monitor.
     * 
     */
    @Export(name="adminStateUp", refs={String.class}, tree="[0]")
    private Output<String> adminStateUp;

    /**
     * @return The administrative state of the monitor.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value updates the
     * state of the existing monitor.
     * 
     */
    public Output<String> adminStateUp() {
        return this.adminStateUp;
    }
    /**
     * The time, in seconds, between sending probes to members.
     * Changing this creates a new monitor.
     * 
     */
    @Export(name="delay", refs={Integer.class}, tree="[0]")
    private Output<Integer> delay;

    /**
     * @return The time, in seconds, between sending probes to members.
     * Changing this creates a new monitor.
     * 
     */
    public Output<Integer> delay() {
        return this.delay;
    }
    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * &#34;200&#34;, or a range like &#34;200-202&#34;. Changing this updates the expected_codes
     * of the existing monitor.
     * 
     */
    @Export(name="expectedCodes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expectedCodes;

    /**
     * @return Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * &#34;200&#34;, or a range like &#34;200-202&#34;. Changing this updates the expected_codes
     * of the existing monitor.
     * 
     */
    public Output<Optional<String>> expectedCodes() {
        return Codegen.optional(this.expectedCodes);
    }
    /**
     * Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it defaults
     * to &#34;GET&#34;. Changing this updates the http_method of the existing monitor.
     * 
     */
    @Export(name="httpMethod", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> httpMethod;

    /**
     * @return Required for HTTP(S) types. The HTTP method used
     * for requests by the monitor. If this attribute is not specified, it defaults
     * to &#34;GET&#34;. Changing this updates the http_method of the existing monitor.
     * 
     */
    public Output<Optional<String>> httpMethod() {
        return Codegen.optional(this.httpMethod);
    }
    /**
     * Number of permissible ping failures before changing
     * the member&#39;s status to INACTIVE. Must be a number between 1 and 10. Changing
     * this updates the max_retries of the existing monitor.
     * 
     */
    @Export(name="maxRetries", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxRetries;

    /**
     * @return Number of permissible ping failures before changing
     * the member&#39;s status to INACTIVE. Must be a number between 1 and 10. Changing
     * this updates the max_retries of the existing monitor.
     * 
     */
    public Output<Integer> maxRetries() {
        return this.maxRetries;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB monitor. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB monitor.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB monitor. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB monitor.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The owner of the monitor. Required if admin wants to
     * create a monitor for another tenant. Changing this creates a new monitor.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the monitor. Required if admin wants to
     * create a monitor for another tenant. Changing this creates a new monitor.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay value.
     * Changing this updates the timeout of the existing monitor.
     * 
     */
    @Export(name="timeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeout;

    /**
     * @return Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay value.
     * Changing this updates the timeout of the existing monitor.
     * 
     */
    public Output<Integer> timeout() {
        return this.timeout;
    }
    /**
     * The type of probe, which is PING, TCP, HTTP, or HTTPS,
     * that is sent by the monitor to verify the member state. Changing this
     * creates a new monitor.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of probe, which is PING, TCP, HTTP, or HTTPS,
     * that is sent by the monitor to verify the member state. Changing this
     * creates a new monitor.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Changing this updates the
     * url_path of the existing monitor.
     * 
     */
    @Export(name="urlPath", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> urlPath;

    /**
     * @return Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Changing this updates the
     * url_path of the existing monitor.
     * 
     */
    public Output<Optional<String>> urlPath() {
        return Codegen.optional(this.urlPath);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MonitorV1(String name) {
        this(name, MonitorV1Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MonitorV1(String name, MonitorV1Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MonitorV1(String name, MonitorV1Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/monitorV1:MonitorV1", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private MonitorV1(String name, Output<String> id, @Nullable MonitorV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/monitorV1:MonitorV1", name, state, makeResourceOptions(options, id));
    }

    private static MonitorV1Args makeArgs(MonitorV1Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MonitorV1Args.Empty : args;
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
    public static MonitorV1 get(String name, Output<String> id, @Nullable MonitorV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MonitorV1(name, id, state, options);
    }
}
