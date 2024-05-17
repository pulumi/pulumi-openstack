// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.QosDscpMarkingRuleArgs;
import com.pulumi.openstack.networking.inputs.QosDscpMarkingRuleState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron QoS DSCP marking rule resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ### Create a QoS Policy with some DSCP marking rule
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
 * import com.pulumi.openstack.networking.QosDscpMarkingRule;
 * import com.pulumi.openstack.networking.QosDscpMarkingRuleArgs;
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
 *             .description("dscp_mark")
 *             .build());
 * 
 *         var dscpMarkingRule1 = new QosDscpMarkingRule("dscpMarkingRule1", QosDscpMarkingRuleArgs.builder()
 *             .qosPolicyId(qosPolicy1.id())
 *             .dscpMark(26)
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
 * QoS DSCP marking rules can be imported using the `qos_policy_id/dscp_marking_rule_id` format, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule dscp_marking_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
 * ```
 * 
 */
@ResourceType(type="openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule")
public class QosDscpMarkingRule extends com.pulumi.resources.CustomResource {
    /**
     * The value of DSCP mark. Changing this updates the DSCP mark value existing
     * QoS DSCP marking rule.
     * 
     */
    @Export(name="dscpMark", refs={Integer.class}, tree="[0]")
    private Output<Integer> dscpMark;

    /**
     * @return The value of DSCP mark. Changing this updates the DSCP mark value existing
     * QoS DSCP marking rule.
     * 
     */
    public Output<Integer> dscpMark() {
        return this.dscpMark;
    }
    /**
     * The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
     * 
     */
    @Export(name="qosPolicyId", refs={String.class}, tree="[0]")
    private Output<String> qosPolicyId;

    /**
     * @return The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
     * 
     */
    public Output<String> qosPolicyId() {
        return this.qosPolicyId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QosDscpMarkingRule(String name) {
        this(name, QosDscpMarkingRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QosDscpMarkingRule(String name, QosDscpMarkingRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QosDscpMarkingRule(String name, QosDscpMarkingRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule", name, args == null ? QosDscpMarkingRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QosDscpMarkingRule(String name, Output<String> id, @Nullable QosDscpMarkingRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule", name, state, makeResourceOptions(options, id));
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
    public static QosDscpMarkingRule get(String name, Output<String> id, @Nullable QosDscpMarkingRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QosDscpMarkingRule(name, id, state, options);
    }
}
