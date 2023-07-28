// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.QosMinimumBandwidthRuleArgs;
import com.pulumi.openstack.networking.inputs.QosMinimumBandwidthRuleState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron QoS minimum bandwidth rule resource within OpenStack.
 * 
 * ## Example Usage
 * ### Create a QoS Policy with some minimum bandwidth rule
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.QosPolicy;
 * import com.pulumi.openstack.networking.QosPolicyArgs;
 * import com.pulumi.openstack.networking.QosMinimumBandwidthRule;
 * import com.pulumi.openstack.networking.QosMinimumBandwidthRuleArgs;
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
 *         var qosPolicy1 = new QosPolicy(&#34;qosPolicy1&#34;, QosPolicyArgs.builder()        
 *             .description(&#34;min_kbps&#34;)
 *             .build());
 * 
 *         var minimumBandwidthRule1 = new QosMinimumBandwidthRule(&#34;minimumBandwidthRule1&#34;, QosMinimumBandwidthRuleArgs.builder()        
 *             .qosPolicyId(qosPolicy1.id())
 *             .minKbps(200)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * QoS minimum bandwidth rules can be imported using the `qos_policy_id/minimum_bandwidth_rule_id` format, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule minimum_bandwidth_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
 * ```
 * 
 */
@ResourceType(type="openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule")
public class QosMinimumBandwidthRule extends com.pulumi.resources.CustomResource {
    /**
     * The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
     * existing QoS minimum bandwidth rule.
     * 
     */
    @Export(name="direction", type=String.class, parameters={})
    private Output</* @Nullable */ String> direction;

    /**
     * @return The direction of traffic. Defaults to &#34;egress&#34;. Changing this updates the direction of the
     * existing QoS minimum bandwidth rule.
     * 
     */
    public Output<Optional<String>> direction() {
        return Codegen.optional(this.direction);
    }
    /**
     * The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     * 
     */
    @Export(name="minKbps", type=Integer.class, parameters={})
    private Output<Integer> minKbps;

    /**
     * @return The minimum kilobits per second. Changing this updates the min kbps value of the existing
     * QoS minimum bandwidth rule.
     * 
     */
    public Output<Integer> minKbps() {
        return this.minKbps;
    }
    /**
     * The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     * 
     */
    @Export(name="qosPolicyId", type=String.class, parameters={})
    private Output<String> qosPolicyId;

    /**
     * @return The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
     * 
     */
    public Output<String> qosPolicyId() {
        return this.qosPolicyId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QosMinimumBandwidthRule(String name) {
        this(name, QosMinimumBandwidthRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QosMinimumBandwidthRule(String name, QosMinimumBandwidthRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QosMinimumBandwidthRule(String name, QosMinimumBandwidthRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule", name, args == null ? QosMinimumBandwidthRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QosMinimumBandwidthRule(String name, Output<String> id, @Nullable QosMinimumBandwidthRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule", name, state, makeResourceOptions(options, id));
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
    public static QosMinimumBandwidthRule get(String name, Output<String> id, @Nullable QosMinimumBandwidthRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QosMinimumBandwidthRule(name, id, state, options);
    }
}
