// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.PoolArgs;
import com.pulumi.openstack.loadbalancer.inputs.PoolState;
import com.pulumi.openstack.loadbalancer.outputs.PoolPersistence;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 pool resource within OpenStack.
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
 * import com.pulumi.openstack.loadbalancer.Pool;
 * import com.pulumi.openstack.loadbalancer.PoolArgs;
 * import com.pulumi.openstack.loadbalancer.inputs.PoolPersistenceArgs;
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
 *         var pool1 = new Pool("pool1", PoolArgs.builder()        
 *             .protocol("HTTP")
 *             .lbMethod("ROUND_ROBIN")
 *             .listenerId("d9415786-5f1a-428b-b35f-2f1523e146d2")
 *             .persistence(PoolPersistenceArgs.builder()
 *                 .type("APP_COOKIE")
 *                 .cookieName("testCookie")
 *                 .build())
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
 * Load Balancer Pool can be imported using the Pool ID, e.g.:
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/pool:Pool pool_1 60ad9ee4-249a-4d60-a45b-aa60e046c513
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/pool:Pool")
public class Pool extends com.pulumi.resources.CustomResource {
    /**
     * The administrative state of the pool.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the pool.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * Human-readable description for the pool.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-readable description for the pool.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The load balancing algorithm to
     * distribute traffic to the pool&#39;s members. Must be one of
     * ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT.
     * 
     */
    @Export(name="lbMethod", refs={String.class}, tree="[0]")
    private Output<String> lbMethod;

    /**
     * @return The load balancing algorithm to
     * distribute traffic to the pool&#39;s members. Must be one of
     * ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT.
     * 
     */
    public Output<String> lbMethod() {
        return this.lbMethod;
    }
    /**
     * The Listener on which the members of the pool
     * will be associated with. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     * 
     */
    @Export(name="listenerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> listenerId;

    /**
     * @return The Listener on which the members of the pool
     * will be associated with. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     * 
     */
    public Output<Optional<String>> listenerId() {
        return Codegen.optional(this.listenerId);
    }
    /**
     * The load balancer on which to provision this
     * pool. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     * 
     */
    @Export(name="loadbalancerId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loadbalancerId;

    /**
     * @return The load balancer on which to provision this
     * pool. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     * 
     */
    public Output<Optional<String>> loadbalancerId() {
        return Codegen.optional(this.loadbalancerId);
    }
    /**
     * Human-readable name for the pool.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Human-readable name for the pool.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Omit this field to prevent session persistence.  Indicates
     * whether connections in the same session will be processed by the same Pool
     * member or not. Changing this creates a new pool.
     * 
     */
    @Export(name="persistence", refs={PoolPersistence.class}, tree="[0]")
    private Output<PoolPersistence> persistence;

    /**
     * @return Omit this field to prevent session persistence.  Indicates
     * whether connections in the same session will be processed by the same Pool
     * member or not. Changing this creates a new pool.
     * 
     */
    public Output<PoolPersistence> persistence() {
        return this.persistence;
    }
    /**
     * The protocol - can either be TCP, HTTP, HTTPS, PROXY,
     * UDP, PROXYV2 (**Octavia minor version &gt;= 2.22**) or SCTP
     * (**Octavia minor version &gt;= 2.23**). Changing this creates a new pool.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return The protocol - can either be TCP, HTTP, HTTPS, PROXY,
     * UDP, PROXYV2 (**Octavia minor version &gt;= 2.22**) or SCTP
     * (**Octavia minor version &gt;= 2.23**). Changing this creates a new pool.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * pool.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * pool.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Required for admins. The UUID of the tenant who owns
     * the pool.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new pool.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the pool.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new pool.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Pool(String name) {
        this(name, PoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Pool(String name, PoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Pool(String name, PoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/pool:Pool", name, args == null ? PoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Pool(String name, Output<String> id, @Nullable PoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/pool:Pool", name, state, makeResourceOptions(options, id));
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
    public static Pool get(String name, Output<String> id, @Nullable PoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Pool(name, id, state, options);
    }
}
