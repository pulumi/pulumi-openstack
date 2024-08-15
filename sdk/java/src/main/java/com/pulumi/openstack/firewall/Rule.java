// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.firewall.RuleArgs;
import com.pulumi.openstack.firewall.inputs.RuleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a v1 firewall rule resource within OpenStack.
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
 * import com.pulumi.openstack.firewall.Rule;
 * import com.pulumi.openstack.firewall.RuleArgs;
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
 *         var rule1 = new Rule("rule1", RuleArgs.builder()
 *             .name("my_rule")
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
 * $ pulumi import openstack:firewall/rule:Rule rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
 * ```
 * 
 */
@ResourceType(type="openstack:firewall/rule:Rule")
public class Rule extends com.pulumi.resources.CustomResource {
    /**
     * Action to be taken ( must be &#34;allow&#34; or &#34;deny&#34;) when the
     * firewall rule matches. Changing this updates the `action` of an existing
     * firewall rule.
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output<String> action;

    /**
     * @return Action to be taken ( must be &#34;allow&#34; or &#34;deny&#34;) when the
     * firewall rule matches. Changing this updates the `action` of an existing
     * firewall rule.
     * 
     */
    public Output<String> action() {
        return this.action;
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
     * firewall rule.
     * 
     */
    @Export(name="destinationPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPort;

    /**
     * @return The destination port on which the firewall
     * rule operates. Changing this updates the `destination_port` of an existing
     * firewall rule.
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
     * IP version, either 4 (default) or 6. Changing this
     * updates the `ip_version` of an existing firewall rule.
     * 
     */
    @Export(name="ipVersion", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ipVersion;

    /**
     * @return IP version, either 4 (default) or 6. Changing this
     * updates the `ip_version` of an existing firewall rule.
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
     * The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule.
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return The protocol type on which the firewall rule operates.
     * Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
     * `protocol` of an existing firewall rule.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The region in which to obtain the v1 Compute client.
     * A Compute client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the v1 Compute client.
     * A Compute client is needed to create a firewall rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall rule.
     * 
     */
    public Output<String> region() {
        return this.region;
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
     * firewall rule.
     * 
     */
    @Export(name="sourcePort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePort;

    /**
     * @return The source port on which the firewall
     * rule operates. Changing this updates the `source_port` of an existing
     * firewall rule.
     * 
     */
    public Output<Optional<String>> sourcePort() {
        return Codegen.optional(this.sourcePort);
    }
    /**
     * The owner of the firewall rule. Required if admin
     * wants to create a firewall rule for another tenant. Changing this creates a
     * new firewall rule.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tenantId;

    /**
     * @return The owner of the firewall rule. Required if admin
     * wants to create a firewall rule for another tenant. Changing this creates a
     * new firewall rule.
     * 
     */
    public Output<Optional<String>> tenantId() {
        return Codegen.optional(this.tenantId);
    }
    /**
     * Map of additional options.
     * 
     */
    @Export(name="valueSpecs", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> valueSpecs;

    /**
     * @return Map of additional options.
     * 
     */
    public Output<Optional<Map<String,String>>> valueSpecs() {
        return Codegen.optional(this.valueSpecs);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Rule(java.lang.String name) {
        this(name, RuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Rule(java.lang.String name, RuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Rule(java.lang.String name, RuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:firewall/rule:Rule", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Rule(java.lang.String name, Output<java.lang.String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:firewall/rule:Rule", name, state, makeResourceOptions(options, id), false);
    }

    private static RuleArgs makeArgs(RuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RuleArgs.Empty : args;
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
    public static Rule get(java.lang.String name, Output<java.lang.String> id, @Nullable RuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Rule(name, id, state, options);
    }
}
