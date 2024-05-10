// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.QosBandwidthLimitRuleArgs;
import com.pulumi.openstack.networking.inputs.QosBandwidthLimitRuleState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron QoS bandwidth limit rule resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ### Create a QoS Policy with some bandwidth limit rule
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.QosPolicy;
 * import com.pulumi.openstack.networking.QosPolicyArgs;
 * import com.pulumi.openstack.networking.QosBandwidthLimitRule;
 * import com.pulumi.openstack.networking.QosBandwidthLimitRuleArgs;
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
 *         var qosPolicy1 = new QosPolicy("qosPolicy1", QosPolicyArgs.builder()        
 *             .name("qos_policy_1")
 *             .description("bw_limit")
 *             .build());
 * 
 *         var bwLimitRule1 = new QosBandwidthLimitRule("bwLimitRule1", QosBandwidthLimitRuleArgs.builder()        
 *             .qosPolicyId(qosPolicy1.id())
 *             .maxKbps(3000)
 *             .maxBurstKbps(300)
 *             .direction("egress")
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
 * QoS bandwidth limit rules can be imported using the `qos_policy_id/bandwidth_limit_rule` format, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule bw_limit_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
 * ```
 * 
 */
@ResourceType(type="openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule")
public class QosBandwidthLimitRule extends com.pulumi.resources.CustomResource {
    /**
     * The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
     * existing QoS bandwidth limit rule.
     * 
     */
    @Export(name="direction", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> direction;

    /**
     * @return The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
     * existing QoS bandwidth limit rule.
     * 
     */
    public Output<Optional<String>> direction() {
        return Codegen.optional(this.direction);
    }
    /**
     * The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
     * maximum burst size in kilobits of the existing QoS bandwidth limit rule.
     * 
     */
    @Export(name="maxBurstKbps", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxBurstKbps;

    /**
     * @return The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the
     * maximum burst size in kilobits of the existing QoS bandwidth limit rule.
     * 
     */
    public Output<Optional<Integer>> maxBurstKbps() {
        return Codegen.optional(this.maxBurstKbps);
    }
    /**
     * The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
     * maximum kilobits per second of the existing QoS bandwidth limit rule.
     * 
     */
    @Export(name="maxKbps", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxKbps;

    /**
     * @return The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the
     * maximum kilobits per second of the existing QoS bandwidth limit rule.
     * 
     */
    public Output<Integer> maxKbps() {
        return this.maxKbps;
    }
    /**
     * The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
     * 
     */
    @Export(name="qosPolicyId", refs={String.class}, tree="[0]")
    private Output<String> qosPolicyId;

    /**
     * @return The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.
     * 
     */
    public Output<String> qosPolicyId() {
        return this.qosPolicyId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QosBandwidthLimitRule(String name) {
        this(name, QosBandwidthLimitRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QosBandwidthLimitRule(String name, QosBandwidthLimitRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QosBandwidthLimitRule(String name, QosBandwidthLimitRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule", name, args == null ? QosBandwidthLimitRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QosBandwidthLimitRule(String name, Output<String> id, @Nullable QosBandwidthLimitRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule", name, state, makeResourceOptions(options, id));
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
    public static QosBandwidthLimitRule get(String name, Output<String> id, @Nullable QosBandwidthLimitRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QosBandwidthLimitRule(name, id, state, options);
    }
}
