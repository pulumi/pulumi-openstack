// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.SecGroupRuleArgs;
import com.pulumi.openstack.networking.inputs.SecGroupRuleState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 neutron security group rule resource within OpenStack.
 * Unlike Nova security groups, neutron separates the group from the rules
 * and also allows an admin to target a specific tenant_id.
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
 * import com.pulumi.openstack.networking.SecGroup;
 * import com.pulumi.openstack.networking.SecGroupArgs;
 * import com.pulumi.openstack.networking.SecGroupRule;
 * import com.pulumi.openstack.networking.SecGroupRuleArgs;
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
 *         var secgroup1 = new SecGroup("secgroup1", SecGroupArgs.builder()
 *             .name("secgroup_1")
 *             .description("My neutron security group")
 *             .build());
 * 
 *         var secgroupRule1 = new SecGroupRule("secgroupRule1", SecGroupRuleArgs.builder()
 *             .direction("ingress")
 *             .ethertype("IPv4")
 *             .protocol("tcp")
 *             .portRangeMin(22)
 *             .portRangeMax(22)
 *             .remoteIpPrefix("0.0.0.0/0")
 *             .securityGroupId(secgroup1.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &gt; **Note:** To expose the full port-range 1:65535, use `0` for `port_range_min`
 * and `port_range_max`.
 * 
 * ## Import
 * 
 * Security Group Rules can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:networking/secGroupRule:SecGroupRule secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745
 * ```
 * 
 */
@ResourceType(type="openstack:networking/secGroupRule:SecGroupRule")
public class SecGroupRule extends com.pulumi.resources.CustomResource {
    /**
     * A description of the rule. Changing this creates a new security group rule.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the rule. Changing this creates a new security group rule.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The direction of the rule, valid values are __ingress__
     * or __egress__. Changing this creates a new security group rule.
     * 
     */
    @Export(name="direction", refs={String.class}, tree="[0]")
    private Output<String> direction;

    /**
     * @return The direction of the rule, valid values are __ingress__
     * or __egress__. Changing this creates a new security group rule.
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }
    /**
     * The layer 3 protocol type, valid values are __IPv4__
     * or __IPv6__. Changing this creates a new security group rule.
     * 
     */
    @Export(name="ethertype", refs={String.class}, tree="[0]")
    private Output<String> ethertype;

    /**
     * @return The layer 3 protocol type, valid values are __IPv4__
     * or __IPv6__. Changing this creates a new security group rule.
     * 
     */
    public Output<String> ethertype() {
        return this.ethertype;
    }
    /**
     * The higher part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     * 
     */
    @Export(name="portRangeMax", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> portRangeMax;

    /**
     * @return The higher part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     * 
     */
    public Output<Optional<Integer>> portRangeMax() {
        return Codegen.optional(this.portRangeMax);
    }
    /**
     * The lower part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     * 
     */
    @Export(name="portRangeMin", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> portRangeMin;

    /**
     * @return The lower part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     * 
     */
    public Output<Optional<Integer>> portRangeMin() {
        return Codegen.optional(this.portRangeMin);
    }
    /**
     * The layer 4 protocol type, valid values are
     * following. Changing this creates a new security group rule. This is required
     * if you want to specify a port range.
     * * empty string or omitted (any protocol)
     * * integer value between 0 and 255 (valid IP protocol number)
     * * __tcp__
     * * __udp__
     * * __icmp__
     * * __ah__
     * * __dccp__
     * * __egp__
     * * __esp__
     * * __gre__
     * * __igmp__
     * * __ipv6-encap__
     * * __ipv6-frag__
     * * __ipv6-icmp__
     * * __ipv6-nonxt__
     * * __ipv6-opts__
     * * __ipv6-route__
     * * __ospf__
     * * __pgm__
     * * __rsvp__
     * * __sctp__
     * * __udplite__
     * * __vrrp__
     * * __ipip__
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> protocol;

    /**
     * @return The layer 4 protocol type, valid values are
     * following. Changing this creates a new security group rule. This is required
     * if you want to specify a port range.
     * * empty string or omitted (any protocol)
     * * integer value between 0 and 255 (valid IP protocol number)
     * * __tcp__
     * * __udp__
     * * __icmp__
     * * __ah__
     * * __dccp__
     * * __egp__
     * * __esp__
     * * __gre__
     * * __igmp__
     * * __ipv6-encap__
     * * __ipv6-frag__
     * * __ipv6-icmp__
     * * __ipv6-nonxt__
     * * __ipv6-opts__
     * * __ipv6-route__
     * * __ospf__
     * * __pgm__
     * * __rsvp__
     * * __sctp__
     * * __udplite__
     * * __vrrp__
     * * __ipip__
     * 
     */
    public Output<Optional<String>> protocol() {
        return Codegen.optional(this.protocol);
    }
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group rule.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group rule.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The remote address group id, the value
     * needs to be an OpenStack ID of an address group in the same tenant. Changing
     * this creates a new security group rule. This argument is mutually exclusive
     * with `remote_ip_prefix` and `remote_group_id`.
     * 
     */
    @Export(name="remoteAddressGroupId", refs={String.class}, tree="[0]")
    private Output<String> remoteAddressGroupId;

    /**
     * @return The remote address group id, the value
     * needs to be an OpenStack ID of an address group in the same tenant. Changing
     * this creates a new security group rule. This argument is mutually exclusive
     * with `remote_ip_prefix` and `remote_group_id`.
     * 
     */
    public Output<String> remoteAddressGroupId() {
        return this.remoteAddressGroupId;
    }
    /**
     * The remote group id, the value needs to be an
     * Openstack ID of a security group in the same tenant. Changing this creates
     * a new security group rule.
     * 
     */
    @Export(name="remoteGroupId", refs={String.class}, tree="[0]")
    private Output<String> remoteGroupId;

    /**
     * @return The remote group id, the value needs to be an
     * Openstack ID of a security group in the same tenant. Changing this creates
     * a new security group rule.
     * 
     */
    public Output<String> remoteGroupId() {
        return this.remoteGroupId;
    }
    /**
     * The remote CIDR, the value needs to be a valid
     * CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
     * 
     */
    @Export(name="remoteIpPrefix", refs={String.class}, tree="[0]")
    private Output<String> remoteIpPrefix;

    /**
     * @return The remote CIDR, the value needs to be a valid
     * CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
     * 
     */
    public Output<String> remoteIpPrefix() {
        return this.remoteIpPrefix;
    }
    /**
     * The security group id the rule should belong
     * to, the value needs to be an Openstack ID of a security group in the same
     * tenant. Changing this creates a new security group rule.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output<String> securityGroupId;

    /**
     * @return The security group id the rule should belong
     * to, the value needs to be an Openstack ID of a security group in the same
     * tenant. Changing this creates a new security group rule.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group rule.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group rule.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecGroupRule(java.lang.String name) {
        this(name, SecGroupRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecGroupRule(java.lang.String name, SecGroupRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecGroupRule(java.lang.String name, SecGroupRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/secGroupRule:SecGroupRule", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecGroupRule(java.lang.String name, Output<java.lang.String> id, @Nullable SecGroupRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/secGroupRule:SecGroupRule", name, state, makeResourceOptions(options, id), false);
    }

    private static SecGroupRuleArgs makeArgs(SecGroupRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecGroupRuleArgs.Empty : args;
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
    public static SecGroupRule get(java.lang.String name, Output<java.lang.String> id, @Nullable SecGroupRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecGroupRule(name, id, state, options);
    }
}
