// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.firewall.RuleV2Args;
import com.pulumi.openstack.firewall.inputs.RuleV2State;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a v2 firewall rule resource within OpenStack.
 * 
 * &gt; **Note:** Firewall v2 has no support for OVN currently.
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
 * import com.pulumi.openstack.firewall.RuleV2;
 * import com.pulumi.openstack.firewall.RuleV2Args;
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
 *         var rule2 = new RuleV2("rule2", RuleV2Args.builder()
 *             .name("firewall_rule")
 *             .description("drop TELNET traffic")
 *             .action("deny")
 *             .protocol("tcp")
 *             .destinationPort("23")
 *             .enabled("true")
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
 * Firewall Rules can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:firewall/ruleV2:RuleV2 rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
 * ```
 * 
 */
@ResourceType(type="openstack:firewall/ruleV2:RuleV2")
public class RuleV2 extends com.pulumi.resources.CustomResource {
    /**
     * Action to be taken (must be &#34;allow&#34;, &#34;deny&#34; or &#34;reject&#34;)
     * when the firewall rule matches. Changing this updates the `action` of an
     * existing firewall rule. Default is `deny`.
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> action;

    /**
     * @return Action to be taken (must be &#34;allow&#34;, &#34;deny&#34; or &#34;reject&#34;)
     * when the firewall rule matches. Changing this updates the `action` of an
     * existing firewall rule. Default is `deny`.
     * 
     */
    public Output<Optional<String>> action() {
        return Codegen.optional(this.action);
    }
    /**
     * A description for the firewall rule. Changing this
     * updates the `description` of an existing firewall rule.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description for the firewall rule. Changing this
     * updates the `description` of an existing firewall rule.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The destination IP address on which the
     * firewall rule operates. Changing this updates the `destination_ip_address`
     * of an existing firewall rule.
     * 
     */
    @Export(name="destinationIpAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationIpAddress;

    /**
     * @return The destination IP address on which the
     * firewall rule operates. Changing this updates the `destination_ip_address`
     * of an existing firewall rule.
     * 
     */
    public Output<Optional<String>> destinationIpAddress() {
        return Codegen.optional(this.destinationIpAddress);
    }
    /**
     * The destination port on which the firewall
     * rule operates. Changing this updates the `destination_port` of an existing
     * firewall rule. Require not `any` or empty protocol.
     * 
     */
    @Export(name="destinationPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPort;

    /**
     * @return The destination port on which the firewall
     * rule operates. Changing this updates the `destination_port` of an existing
     * firewall rule. Require not `any` or empty protocol.
     * 
     */
    public Output<Optional<String>> destinationPort() {
        return Codegen.optional(this.destinationPort);
    }
    /**
     * Enabled status for the firewall rule (must be &#34;true&#34;
     * or &#34;false&#34; if provided - defaults to &#34;true&#34;). Changing this updates the
     * `enabled` status of an existing firewall rule.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Enabled status for the firewall rule (must be &#34;true&#34;
     * or &#34;false&#34; if provided - defaults to &#34;true&#34;). Changing this updates the
     * `enabled` status of an existing firewall rule.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * IP version, either 4 or 6. Changing this
     * updates the `ip_version` of an existing firewall rule. Default is `4`.
     * 
     */
    @Export(name="ipVersion", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ipVersion;

    /**
     * @return IP version, either 4 or 6. Changing this
     * updates the `ip_version` of an existing firewall rule. Default is `4`.
     * 
     */
    public Output<Optional<Integer>> ipVersion() {
        return Codegen.optional(this.ipVersion);
    }
    /**
     * A unique name for the firewall rule. Changing this
     * updates the `name` of an existing firewall rule.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the firewall rule. Changing this
     * updates the `name` of an existing firewall rule.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * This argument conflicts and is interchangeable
     * with `tenant_id`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another project. Changing this creates a new
     * firewall rule.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return This argument conflicts and is interchangeable
     * with `tenant_id`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another project. Changing this creates a new
     * firewall rule.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * (Optional; Required if `source_port` or `destination_port` is not
     * empty) The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule. Default is `any`.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> protocol;

    /**
     * @return (Optional; Required if `source_port` or `destination_port` is not
     * empty) The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule. Default is `any`.
     * 
     */
    public Output<Optional<String>> protocol() {
        return Codegen.optional(this.protocol);
    }
    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Sharing status of the firewall rule (must be &#34;true&#34;
     * or &#34;false&#34; if provided). If this is &#34;true&#34; the policy is visible to, and
     * can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall policy. On
     * 
     */
    @Export(name="shared", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> shared;

    /**
     * @return Sharing status of the firewall rule (must be &#34;true&#34;
     * or &#34;false&#34; if provided). If this is &#34;true&#34; the policy is visible to, and
     * can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall policy. On
     * 
     */
    public Output<Optional<Boolean>> shared() {
        return Codegen.optional(this.shared);
    }
    /**
     * The source IP address on which the firewall
     * rule operates. Changing this updates the `source_ip_address` of an existing
     * firewall rule.
     * 
     */
    @Export(name="sourceIpAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceIpAddress;

    /**
     * @return The source IP address on which the firewall
     * rule operates. Changing this updates the `source_ip_address` of an existing
     * firewall rule.
     * 
     */
    public Output<Optional<String>> sourceIpAddress() {
        return Codegen.optional(this.sourceIpAddress);
    }
    /**
     * The source port on which the firewall
     * rule operates. Changing this updates the `source_port` of an existing
     * firewall rule. Require not `any` or empty protocol.
     * 
     */
    @Export(name="sourcePort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePort;

    /**
     * @return The source port on which the firewall
     * rule operates. Changing this updates the `source_port` of an existing
     * firewall rule. Require not `any` or empty protocol.
     * 
     */
    public Output<Optional<String>> sourcePort() {
        return Codegen.optional(this.sourcePort);
    }
    /**
     * This argument conflicts and is interchangeable
     * with `project_id`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another tenant. Changing this creates a new
     * firewall rule.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return This argument conflicts and is interchangeable
     * with `project_id`. The owner of the firewall rule. Required if admin wants
     * to create a firewall rule for another tenant. Changing this creates a new
     * firewall rule.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RuleV2(String name) {
        this(name, RuleV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RuleV2(String name, @Nullable RuleV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RuleV2(String name, @Nullable RuleV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:firewall/ruleV2:RuleV2", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private RuleV2(String name, Output<String> id, @Nullable RuleV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:firewall/ruleV2:RuleV2", name, state, makeResourceOptions(options, id));
    }

    private static RuleV2Args makeArgs(@Nullable RuleV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RuleV2Args.Empty : args;
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
    public static RuleV2 get(String name, Output<String> id, @Nullable RuleV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RuleV2(name, id, state, options);
    }
}
