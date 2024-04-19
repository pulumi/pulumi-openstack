// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.firewall.FirewallArgs;
import com.pulumi.openstack.firewall.inputs.FirewallState;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a v1 firewall resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.firewall.Rule;
 * import com.pulumi.openstack.firewall.RuleArgs;
 * import com.pulumi.openstack.firewall.Policy;
 * import com.pulumi.openstack.firewall.PolicyArgs;
 * import com.pulumi.openstack.firewall.Firewall;
 * import com.pulumi.openstack.firewall.FirewallArgs;
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
 *         var rule1 = new Rule(&#34;rule1&#34;, RuleArgs.builder()        
 *             .name(&#34;my-rule-1&#34;)
 *             .description(&#34;drop TELNET traffic&#34;)
 *             .action(&#34;deny&#34;)
 *             .protocol(&#34;tcp&#34;)
 *             .destinationPort(&#34;23&#34;)
 *             .enabled(&#34;true&#34;)
 *             .build());
 * 
 *         var rule2 = new Rule(&#34;rule2&#34;, RuleArgs.builder()        
 *             .name(&#34;my-rule-2&#34;)
 *             .description(&#34;drop NTP traffic&#34;)
 *             .action(&#34;deny&#34;)
 *             .protocol(&#34;udp&#34;)
 *             .destinationPort(&#34;123&#34;)
 *             .enabled(&#34;false&#34;)
 *             .build());
 * 
 *         var policy1 = new Policy(&#34;policy1&#34;, PolicyArgs.builder()        
 *             .name(&#34;my-policy&#34;)
 *             .rules(            
 *                 rule1.id(),
 *                 rule2.id())
 *             .build());
 * 
 *         var firewall1 = new Firewall(&#34;firewall1&#34;, FirewallArgs.builder()        
 *             .name(&#34;my-firewall&#34;)
 *             .policyId(policy1.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Firewalls can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:firewall/firewall:Firewall firewall_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97
 * ```
 * 
 */
@ResourceType(type="openstack:firewall/firewall:Firewall")
public class Firewall extends com.pulumi.resources.CustomResource {
    /**
     * Administrative up/down status for the firewall
     * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
     * Changing this updates the `admin_state_up` of an existing firewall.
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the firewall
     * (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
     * Changing this updates the `admin_state_up` of an existing firewall.
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     * 
     */
    @Export(name="associatedRouters", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> associatedRouters;

    /**
     * @return Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     * 
     */
    public Output<List<String>> associatedRouters() {
        return this.associatedRouters;
    }
    /**
     * A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Should this firewall not be associated with any routers
     * (must be &#34;true&#34; or &#34;false&#34; if provide - defaults to &#34;false&#34;).
     * Conflicts with `associated_routers`.
     * 
     */
    @Export(name="noRouters", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noRouters;

    /**
     * @return Should this firewall not be associated with any routers
     * (must be &#34;true&#34; or &#34;false&#34; if provide - defaults to &#34;false&#34;).
     * Conflicts with `associated_routers`.
     * 
     */
    public Output<Optional<Boolean>> noRouters() {
        return Codegen.optional(this.noRouters);
    }
    /**
     * The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     * 
     */
    @Export(name="policyId", refs={String.class}, tree="[0]")
    private Output<String> policyId;

    /**
     * @return The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     * 
     */
    public Output<String> policyId() {
        return this.policyId;
    }
    /**
     * The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * Map of additional options.
     * 
     */
    @Export(name="valueSpecs", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional options.
     * 
     */
    public Output<Optional<Map<String,Object>>> valueSpecs() {
        return Codegen.optional(this.valueSpecs);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Firewall(String name) {
        this(name, FirewallArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Firewall(String name, FirewallArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Firewall(String name, FirewallArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:firewall/firewall:Firewall", name, args == null ? FirewallArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Firewall(String name, Output<String> id, @Nullable FirewallState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:firewall/firewall:Firewall", name, state, makeResourceOptions(options, id));
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
    public static Firewall get(String name, Output<String> id, @Nullable FirewallState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Firewall(name, id, state, options);
    }
}
