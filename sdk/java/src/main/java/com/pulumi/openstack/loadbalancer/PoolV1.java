// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.PoolV1Args;
import com.pulumi.openstack.loadbalancer.inputs.PoolV1State;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 load balancer pool resource within OpenStack.
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
 * import com.pulumi.openstack.loadbalancer.PoolV1;
 * import com.pulumi.openstack.loadbalancer.PoolV1Args;
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
 *         var pool1 = new PoolV1("pool1", PoolV1Args.builder()
 *             .name("tf_test_lb_pool")
 *             .protocol("HTTP")
 *             .subnetId("12345")
 *             .lbMethod("ROUND_ROBIN")
 *             .lbProvider("haproxy")
 *             .monitorIds("67890")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Complete Load Balancing Stack Example
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Notes
 * 
 * The `member` block is deprecated in favor of the `openstack.loadbalancer.MemberV1` resource.
 * 
 * ## Import
 * 
 * Load Balancer Pools can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/poolV1:PoolV1 pool_1 b255e6ba-02ad-43e6-8951-3428ca26b713
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/poolV1:PoolV1")
public class PoolV1 extends com.pulumi.resources.CustomResource {
    /**
     * The algorithm used to distribute load between the
     * members of the pool. The current specification supports &#39;ROUND_ROBIN&#39; and
     * &#39;LEAST_CONNECTIONS&#39; as valid values for this attribute.
     * 
     */
    @Export(name="lbMethod", refs={String.class}, tree="[0]")
    private Output<String> lbMethod;

    /**
     * @return The algorithm used to distribute load between the
     * members of the pool. The current specification supports &#39;ROUND_ROBIN&#39; and
     * &#39;LEAST_CONNECTIONS&#39; as valid values for this attribute.
     * 
     */
    public Output<String> lbMethod() {
        return this.lbMethod;
    }
    /**
     * The backend load balancing provider. For example:
     * `haproxy`, `F5`, etc.
     * 
     */
    @Export(name="lbProvider", refs={String.class}, tree="[0]")
    private Output<String> lbProvider;

    /**
     * @return The backend load balancing provider. For example:
     * `haproxy`, `F5`, etc.
     * 
     */
    public Output<String> lbProvider() {
        return this.lbProvider;
    }
    /**
     * An existing node to add to the pool. Changing this
     * updates the members of the pool. The member object structure is documented
     * below. Please note that the `member` block is deprecated in favor of the
     * `openstack.loadbalancer.MemberV1` resource.
     * 
     * @deprecated
     * Use openstack.loadbalancer.MemberV1 instead
     * 
     */
    @Deprecated /* Use openstack.loadbalancer.MemberV1 instead */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> members;

    /**
     * @return An existing node to add to the pool. Changing this
     * updates the members of the pool. The member object structure is documented
     * below. Please note that the `member` block is deprecated in favor of the
     * `openstack.loadbalancer.MemberV1` resource.
     * 
     */
    public Output<Optional<List<String>>> members() {
        return Codegen.optional(this.members);
    }
    /**
     * A list of IDs of monitors to associate with the
     * pool.
     * 
     */
    @Export(name="monitorIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> monitorIds;

    /**
     * @return A list of IDs of monitors to associate with the
     * pool.
     * 
     */
    public Output<Optional<List<String>>> monitorIds() {
        return Codegen.optional(this.monitorIds);
    }
    /**
     * The name of the pool. Changing this updates the name of
     * the existing pool.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the pool. Changing this updates the name of
     * the existing pool.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The protocol used by the pool members, you can use
     * either &#39;TCP, &#39;HTTP&#39;, or &#39;HTTPS&#39;. Changing this creates a new pool.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return The protocol used by the pool members, you can use
     * either &#39;TCP, &#39;HTTP&#39;, or &#39;HTTPS&#39;. Changing this creates a new pool.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB pool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB pool.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB pool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB pool.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The network on which the members of the pool will be
     * located. Only members that are on this network can be added to the pool.
     * Changing this creates a new pool.
     * 
     */
    @Export(name="subnetId", refs={String.class}, tree="[0]")
    private Output<String> subnetId;

    /**
     * @return The network on which the members of the pool will be
     * located. Only members that are on this network can be added to the pool.
     * Changing this creates a new pool.
     * 
     */
    public Output<String> subnetId() {
        return this.subnetId;
    }
    /**
     * The owner of the pool. Required if admin wants to
     * create a pool member for another tenant. Changing this creates a new pool.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the pool. Required if admin wants to
     * create a pool member for another tenant. Changing this creates a new pool.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PoolV1(String name) {
        this(name, PoolV1Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PoolV1(String name, PoolV1Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PoolV1(String name, PoolV1Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/poolV1:PoolV1", name, args == null ? PoolV1Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PoolV1(String name, Output<String> id, @Nullable PoolV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/poolV1:PoolV1", name, state, makeResourceOptions(options, id));
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
    public static PoolV1 get(String name, Output<String> id, @Nullable PoolV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PoolV1(name, id, state, options);
    }
}
