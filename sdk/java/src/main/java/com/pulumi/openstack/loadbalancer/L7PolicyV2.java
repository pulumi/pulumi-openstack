// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.L7PolicyV2Args;
import com.pulumi.openstack.loadbalancer.inputs.L7PolicyV2State;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Load Balancer L7 Policy resource within OpenStack.
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
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Subnet;
 * import com.pulumi.openstack.networking.SubnetArgs;
 * import com.pulumi.openstack.LbLoadbalancerV2;
 * import com.pulumi.openstack.LbLoadbalancerV2Args;
 * import com.pulumi.openstack.loadbalancer.Listener;
 * import com.pulumi.openstack.loadbalancer.ListenerArgs;
 * import com.pulumi.openstack.loadbalancer.Pool;
 * import com.pulumi.openstack.loadbalancer.PoolArgs;
 * import com.pulumi.openstack.loadbalancer.L7PolicyV2;
 * import com.pulumi.openstack.loadbalancer.L7PolicyV2Args;
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
 *         var network1 = new Network("network1", NetworkArgs.builder()
 *             .name("network_1")
 *             .adminStateUp(true)
 *             .build());
 * 
 *         var subnet1 = new Subnet("subnet1", SubnetArgs.builder()
 *             .name("subnet_1")
 *             .cidr("192.168.199.0/24")
 *             .ipVersion(4)
 *             .networkId(network1.id())
 *             .build());
 * 
 *         var loadbalancer1 = new LbLoadbalancerV2("loadbalancer1", LbLoadbalancerV2Args.builder()
 *             .name("loadbalancer_1")
 *             .vipSubnetId(subnet1.id())
 *             .build());
 * 
 *         var listener1 = new Listener("listener1", ListenerArgs.builder()
 *             .name("listener_1")
 *             .protocol("HTTP")
 *             .protocolPort(8080)
 *             .loadbalancerId(loadbalancer1.id())
 *             .build());
 * 
 *         var pool1 = new Pool("pool1", PoolArgs.builder()
 *             .name("pool_1")
 *             .protocol("HTTP")
 *             .lbMethod("ROUND_ROBIN")
 *             .loadbalancerId(loadbalancer1.id())
 *             .build());
 * 
 *         var l7policy1 = new L7PolicyV2("l7policy1", L7PolicyV2Args.builder()
 *             .name("test")
 *             .action("REDIRECT_TO_POOL")
 *             .description("test l7 policy")
 *             .position(1)
 *             .listenerId(listener1.id())
 *             .redirectPoolId(pool1.id())
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
 * Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.:
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/l7PolicyV2:L7PolicyV2 l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/l7PolicyV2:L7PolicyV2")
public class L7PolicyV2 extends com.pulumi.resources.CustomResource {
    /**
     * The L7 Policy action - can either be REDIRECT\_TO\_POOL,
     * REDIRECT\_TO\_URL or REJECT.
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
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
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the L7 Policy.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * Human-readable description for the L7 Policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-readable description for the L7 Policy.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Listener on which the L7 Policy will be associated with.
     * Changing this creates a new L7 Policy.
     * 
     */
    @Export(name="listenerId", refs={String.class}, tree="[0]")
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
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Human-readable name for the L7 Policy. Does not have
     * to be unique.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The position of this policy on the listener. Positions start at 1.
     * 
     */
    @Export(name="position", refs={Integer.class}, tree="[0]")
    private Output<Integer> position;

    /**
     * @return The position of this policy on the listener. Positions start at 1.
     * 
     */
    public Output<Integer> position() {
        return this.position;
    }
    /**
     * Integer. Requests matching this policy will be\
     * redirected to the specified URL or Prefix URL with the HTTP response code.
     * Valid if action is REDIRECT\_TO\_URL or REDIRECT\_PREFIX. Valid options are:
     * 301, 302, 303, 307, or 308. Default is 302. New in octavia version 2.9
     * 
     */
    @Export(name="redirectHttpCode", refs={Integer.class}, tree="[0]")
    private Output<Integer> redirectHttpCode;

    /**
     * @return Integer. Requests matching this policy will be\
     * redirected to the specified URL or Prefix URL with the HTTP response code.
     * Valid if action is REDIRECT\_TO\_URL or REDIRECT\_PREFIX. Valid options are:
     * 301, 302, 303, 307, or 308. Default is 302. New in octavia version 2.9
     * 
     */
    public Output<Integer> redirectHttpCode() {
        return this.redirectHttpCode;
    }
    /**
     * Requests matching this policy will be redirected to the
     * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
     * 
     */
    @Export(name="redirectPoolId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> redirectPoolId;

    /**
     * @return Requests matching this policy will be redirected to the
     * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
     * 
     */
    public Output<Optional<String>> redirectPoolId() {
        return Codegen.optional(this.redirectPoolId);
    }
    /**
     * Requests matching this policy will be redirected to
     * this Prefix URL. Only valid if action is REDIRECT\_PREFIX.
     * 
     */
    @Export(name="redirectPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> redirectPrefix;

    /**
     * @return Requests matching this policy will be redirected to
     * this Prefix URL. Only valid if action is REDIRECT\_PREFIX.
     * 
     */
    public Output<Optional<String>> redirectPrefix() {
        return Codegen.optional(this.redirectPrefix);
    }
    /**
     * Requests matching this policy will be redirected to this URL.
     * Only valid if action is REDIRECT\_TO\_URL.
     * 
     */
    @Export(name="redirectUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> redirectUrl;

    /**
     * @return Requests matching this policy will be redirected to this URL.
     * Only valid if action is REDIRECT\_TO\_URL.
     * 
     */
    public Output<Optional<String>> redirectUrl() {
        return Codegen.optional(this.redirectUrl);
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an L7 policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Policy.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an L7 policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Policy.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Policy.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Policy.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the L7 Policy.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Policy.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public L7PolicyV2(java.lang.String name) {
        this(name, L7PolicyV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public L7PolicyV2(java.lang.String name, L7PolicyV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public L7PolicyV2(java.lang.String name, L7PolicyV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/l7PolicyV2:L7PolicyV2", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private L7PolicyV2(java.lang.String name, Output<java.lang.String> id, @Nullable L7PolicyV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/l7PolicyV2:L7PolicyV2", name, state, makeResourceOptions(options, id), false);
    }

    private static L7PolicyV2Args makeArgs(L7PolicyV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? L7PolicyV2Args.Empty : args;
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
    public static L7PolicyV2 get(java.lang.String name, Output<java.lang.String> id, @Nullable L7PolicyV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new L7PolicyV2(name, id, state, options);
    }
}
