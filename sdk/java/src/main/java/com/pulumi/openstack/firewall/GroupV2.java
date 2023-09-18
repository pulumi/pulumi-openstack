// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.firewall.GroupV2Args;
import com.pulumi.openstack.firewall.inputs.GroupV2State;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a v2 firewall group resource within OpenStack.
 * 
 * &gt; **Note:** Firewall v2 has no support for OVN currently.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.firewall.RuleV2;
 * import com.pulumi.openstack.firewall.RuleV2Args;
 * import com.pulumi.openstack.firewall.PolicyV2;
 * import com.pulumi.openstack.firewall.PolicyV2Args;
 * import com.pulumi.openstack.firewall.GroupV2;
 * import com.pulumi.openstack.firewall.GroupV2Args;
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
 *         var rule1 = new RuleV2(&#34;rule1&#34;, RuleV2Args.builder()        
 *             .description(&#34;drop TELNET traffic&#34;)
 *             .action(&#34;deny&#34;)
 *             .protocol(&#34;tcp&#34;)
 *             .destinationPort(&#34;23&#34;)
 *             .enabled(&#34;true&#34;)
 *             .build());
 * 
 *         var rule2 = new RuleV2(&#34;rule2&#34;, RuleV2Args.builder()        
 *             .description(&#34;drop NTP traffic&#34;)
 *             .action(&#34;deny&#34;)
 *             .protocol(&#34;udp&#34;)
 *             .destinationPort(&#34;123&#34;)
 *             .enabled(&#34;false&#34;)
 *             .build());
 * 
 *         var policy1 = new PolicyV2(&#34;policy1&#34;, PolicyV2Args.builder()        
 *             .rules(rule1.id())
 *             .build());
 * 
 *         var policy2 = new PolicyV2(&#34;policy2&#34;, PolicyV2Args.builder()        
 *             .rules(rule2.id())
 *             .build());
 * 
 *         var group1 = new GroupV2(&#34;group1&#34;, GroupV2Args.builder()        
 *             .ingressFirewallPolicyId(policy1.id())
 *             .egressFirewallPolicyId(policy2.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Firewall groups can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:firewall/groupV2:GroupV2 group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97
 * ```
 * 
 */
@ResourceType(type="openstack:firewall/groupV2:GroupV2")
public class GroupV2 extends com.pulumi.resources.CustomResource {
    /**
     * Administrative up/down status for the firewall
     * group (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
     * Changing this updates the `admin_state_up` of an existing firewall group.
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the firewall
     * group (must be &#34;true&#34; or &#34;false&#34; if provided - defaults to &#34;true&#34;).
     * Changing this updates the `admin_state_up` of an existing firewall group.
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * A description for the firewall group. Changing this
     * updates the `description` of an existing firewall group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description for the firewall group. Changing this
     * updates the `description` of an existing firewall group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The egress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `egress_firewall_policy_id` of an existing firewall group.
     * 
     */
    @Export(name="egressFirewallPolicyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> egressFirewallPolicyId;

    /**
     * @return The egress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `egress_firewall_policy_id` of an existing firewall group.
     * 
     */
    public Output<Optional<String>> egressFirewallPolicyId() {
        return Codegen.optional(this.egressFirewallPolicyId);
    }
    /**
     * The ingress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `ingress_firewall_policy_id` of an existing firewall group.
     * 
     */
    @Export(name="ingressFirewallPolicyId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ingressFirewallPolicyId;

    /**
     * @return The ingress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `ingress_firewall_policy_id` of an existing firewall group.
     * 
     */
    public Output<Optional<String>> ingressFirewallPolicyId() {
        return Codegen.optional(this.ingressFirewallPolicyId);
    }
    /**
     * A name for the firewall group. Changing this
     * updates the `name` of an existing firewall.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name for the firewall group. Changing this
     * updates the `name` of an existing firewall.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Port(s) to associate this firewall group
     * with. Must be a list of strings. Changing this updates the associated ports
     * of an existing firewall group.
     * 
     */
    @Export(name="ports", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ports;

    /**
     * @return Port(s) to associate this firewall group
     * with. Must be a list of strings. Changing this updates the associated ports
     * of an existing firewall group.
     * 
     */
    public Output<Optional<List<String>>> ports() {
        return Codegen.optional(this.ports);
    }
    /**
     * This argument conflict and interchangeable with
     * `tenant_id`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another project. Changing this creates a new
     * firewall group.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return This argument conflict and interchangeable with
     * `tenant_id`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another project. Changing this creates a new
     * firewall group.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall group.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall group.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Sharing status of the firewall group (must be &#34;true&#34;
     * or &#34;false&#34; if provided). If this is &#34;true&#34; the firewall group is visible to,
     * and can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall group. Only administrative users
     * can specify if the firewall group should be shared.
     * 
     */
    @Export(name="shared", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> shared;

    /**
     * @return Sharing status of the firewall group (must be &#34;true&#34;
     * or &#34;false&#34; if provided). If this is &#34;true&#34; the firewall group is visible to,
     * and can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall group. Only administrative users
     * can specify if the firewall group should be shared.
     * 
     */
    public Output<Optional<Boolean>> shared() {
        return Codegen.optional(this.shared);
    }
    /**
     * The status of the firewall group.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the firewall group.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * This argument conflict and interchangeable with
     * `project_id`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another tenant. Changing this creates a new
     * firewall group.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return This argument conflict and interchangeable with
     * `project_id`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another tenant. Changing this creates a new
     * firewall group.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupV2(String name) {
        this(name, GroupV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupV2(String name, @Nullable GroupV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupV2(String name, @Nullable GroupV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:firewall/groupV2:GroupV2", name, args == null ? GroupV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupV2(String name, Output<String> id, @Nullable GroupV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:firewall/groupV2:GroupV2", name, state, makeResourceOptions(options, id));
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
    public static GroupV2 get(String name, Output<String> id, @Nullable GroupV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupV2(name, id, state, options);
    }
}
