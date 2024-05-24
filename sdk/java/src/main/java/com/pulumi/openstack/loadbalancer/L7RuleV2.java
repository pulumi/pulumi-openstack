// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.L7RuleV2Args;
import com.pulumi.openstack.loadbalancer.inputs.L7RuleV2State;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 L7 Rule resource within OpenStack.
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
 * import com.pulumi.openstack.loadbalancer.L7RuleV2;
 * import com.pulumi.openstack.loadbalancer.L7RuleV2Args;
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
 *             .adminStateUp("true")
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
 *             .action("REDIRECT_TO_URL")
 *             .description("test description")
 *             .position(1)
 *             .listenerId(listener1.id())
 *             .redirectUrl("http://www.example.com")
 *             .build());
 * 
 *         var l7rule1 = new L7RuleV2("l7rule1", L7RuleV2Args.builder()
 *             .l7policyId(l7policy1.id())
 *             .type("PATH")
 *             .compareType("EQUAL_TO")
 *             .value("/api")
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
 * Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID
 * separated by a slash, e.g.:
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/l7RuleV2:L7RuleV2 l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/l7RuleV2:L7RuleV2")
public class L7RuleV2 extends com.pulumi.resources.CustomResource {
    /**
     * The administrative state of the L7 Rule.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the L7 Rule.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * The comparison type for the L7 rule - can either be
     * CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
     * 
     */
    @Export(name="compareType", refs={String.class}, tree="[0]")
    private Output<String> compareType;

    /**
     * @return The comparison type for the L7 rule - can either be
     * CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
     * 
     */
    public Output<String> compareType() {
        return this.compareType;
    }
    /**
     * When true the logic of the rule is inverted. For example, with invert
     * true, equal to would become not equal to. Default is false.
     * 
     */
    @Export(name="invert", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> invert;

    /**
     * @return When true the logic of the rule is inverted. For example, with invert
     * true, equal to would become not equal to. Default is false.
     * 
     */
    public Output<Optional<Boolean>> invert() {
        return Codegen.optional(this.invert);
    }
    /**
     * The key to use for the comparison. For example, the name of the cookie to
     * evaluate. Valid when `type` is set to COOKIE or HEADER.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> key;

    /**
     * @return The key to use for the comparison. For example, the name of the cookie to
     * evaluate. Valid when `type` is set to COOKIE or HEADER.
     * 
     */
    public Output<Optional<String>> key() {
        return Codegen.optional(this.key);
    }
    /**
     * The ID of the L7 Policy to query. Changing this creates a new
     * L7 Rule.
     * 
     */
    @Export(name="l7policyId", refs={String.class}, tree="[0]")
    private Output<String> l7policyId;

    /**
     * @return The ID of the L7 Policy to query. Changing this creates a new
     * L7 Rule.
     * 
     */
    public Output<String> l7policyId() {
        return this.l7policyId;
    }
    /**
     * The ID of the Listener owning this resource.
     * 
     */
    @Export(name="listenerId", refs={String.class}, tree="[0]")
    private Output<String> listenerId;

    /**
     * @return The ID of the Listener owning this resource.
     * 
     */
    public Output<String> listenerId() {
        return this.listenerId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Rule.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Rule.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Rule.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Rule.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the L7 Rule.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Rule.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
     * HOST\_NAME or PATH.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
     * HOST\_NAME or PATH.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The value to use for the comparison. For example, the file type to
     * compare.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The value to use for the comparison. For example, the file type to
     * compare.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public L7RuleV2(String name) {
        this(name, L7RuleV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public L7RuleV2(String name, L7RuleV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public L7RuleV2(String name, L7RuleV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/l7RuleV2:L7RuleV2", name, args == null ? L7RuleV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private L7RuleV2(String name, Output<String> id, @Nullable L7RuleV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/l7RuleV2:L7RuleV2", name, state, makeResourceOptions(options, id));
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
    public static L7RuleV2 get(String name, Output<String> id, @Nullable L7RuleV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new L7RuleV2(name, id, state, options);
    }
}
