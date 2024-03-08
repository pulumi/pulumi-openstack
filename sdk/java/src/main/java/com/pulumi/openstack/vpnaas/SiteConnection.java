// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.vpnaas.SiteConnectionArgs;
import com.pulumi.openstack.vpnaas.inputs.SiteConnectionState;
import com.pulumi.openstack.vpnaas.outputs.SiteConnectionDpd;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron IPSec site connection resource within OpenStack.
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
 * import com.pulumi.openstack.vpnaas.SiteConnection;
 * import com.pulumi.openstack.vpnaas.SiteConnectionArgs;
 * import com.pulumi.openstack.vpnaas.inputs.SiteConnectionDpdArgs;
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
 *         var conn1 = new SiteConnection(&#34;conn1&#34;, SiteConnectionArgs.builder()        
 *             .ikepolicyId(openstack_vpnaas_ike_policy_v2.policy_2().id())
 *             .ipsecpolicyId(openstack_vpnaas_ipsec_policy_v2.policy_1().id())
 *             .vpnserviceId(openstack_vpnaas_service_v2.service_1().id())
 *             .psk(&#34;secret&#34;)
 *             .peerAddress(&#34;192.168.10.1&#34;)
 *             .localEpGroupId(openstack_vpnaas_endpoint_group_v2.group_2().id())
 *             .peerEpGroupId(openstack_vpnaas_endpoint_group_v2.group_1().id())
 *             .dpds(SiteConnectionDpdArgs.builder()
 *                 .action(&#34;restart&#34;)
 *                 .timeout(42)
 *                 .interval(21)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Site Connections can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:vpnaas/siteConnection:SiteConnection conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 * 
 */
@ResourceType(type="openstack:vpnaas/siteConnection:SiteConnection")
public class SiteConnection extends com.pulumi.resources.CustomResource {
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing connection.
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing connection.
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * The human-readable description for the connection.
     * Changing this updates the description of the existing connection.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The human-readable description for the connection.
     * Changing this updates the description of the existing connection.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A dictionary with dead peer detection (DPD) protocol controls.
     * 
     */
    @Export(name="dpds", refs={List.class,SiteConnectionDpd.class}, tree="[0,1]")
    private Output<List<SiteConnectionDpd>> dpds;

    /**
     * @return A dictionary with dead peer detection (DPD) protocol controls.
     * 
     */
    public Output<List<SiteConnectionDpd>> dpds() {
        return this.dpds;
    }
    /**
     * The ID of the IKE policy. Changing this creates a new connection.
     * 
     */
    @Export(name="ikepolicyId", refs={String.class}, tree="[0]")
    private Output<String> ikepolicyId;

    /**
     * @return The ID of the IKE policy. Changing this creates a new connection.
     * 
     */
    public Output<String> ikepolicyId() {
        return this.ikepolicyId;
    }
    /**
     * A valid value is response-only or bi-directional. Default is bi-directional.
     * 
     */
    @Export(name="initiator", refs={String.class}, tree="[0]")
    private Output<String> initiator;

    /**
     * @return A valid value is response-only or bi-directional. Default is bi-directional.
     * 
     */
    public Output<String> initiator() {
        return this.initiator;
    }
    /**
     * The ID of the IPsec policy. Changing this creates a new connection.
     * 
     */
    @Export(name="ipsecpolicyId", refs={String.class}, tree="[0]")
    private Output<String> ipsecpolicyId;

    /**
     * @return The ID of the IPsec policy. Changing this creates a new connection.
     * 
     */
    public Output<String> ipsecpolicyId() {
        return this.ipsecpolicyId;
    }
    /**
     * The ID for the endpoint group that contains private subnets for the local side of the connection.
     * You must specify this parameter with the peer_ep_group_id parameter unless
     * in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
     * Changing this updates the existing connection.
     * 
     */
    @Export(name="localEpGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> localEpGroupId;

    /**
     * @return The ID for the endpoint group that contains private subnets for the local side of the connection.
     * You must specify this parameter with the peer_ep_group_id parameter unless
     * in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
     * Changing this updates the existing connection.
     * 
     */
    public Output<Optional<String>> localEpGroupId() {
        return Codegen.optional(this.localEpGroupId);
    }
    /**
     * An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
     * Most often, local ID would be domain name, email address, etc.
     * If this is not configured then the external IP address will be used as the ID.
     * 
     */
    @Export(name="localId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> localId;

    /**
     * @return An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
     * Most often, local ID would be domain name, email address, etc.
     * If this is not configured then the external IP address will be used as the ID.
     * 
     */
    public Output<Optional<String>> localId() {
        return Codegen.optional(this.localId);
    }
    /**
     * The maximum transmission unit (MTU) value to address fragmentation.
     * Minimum value is 68 for IPv4, and 1280 for IPv6.
     * 
     */
    @Export(name="mtu", refs={Integer.class}, tree="[0]")
    private Output<Integer> mtu;

    /**
     * @return The maximum transmission unit (MTU) value to address fragmentation.
     * Minimum value is 68 for IPv4, and 1280 for IPv6.
     * 
     */
    public Output<Integer> mtu() {
        return this.mtu;
    }
    /**
     * The name of the connection. Changing this updates the name of
     * the existing connection.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the connection. Changing this updates the name of
     * the existing connection.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The peer gateway public IPv4 or IPv6 address or FQDN.
     * 
     */
    @Export(name="peerAddress", refs={String.class}, tree="[0]")
    private Output<String> peerAddress;

    /**
     * @return The peer gateway public IPv4 or IPv6 address or FQDN.
     * 
     */
    public Output<String> peerAddress() {
        return this.peerAddress;
    }
    /**
     * Unique list of valid peer private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; .
     * 
     */
    @Export(name="peerCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> peerCidrs;

    /**
     * @return Unique list of valid peer private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; .
     * 
     */
    public Output<Optional<List<String>>> peerCidrs() {
        return Codegen.optional(this.peerCidrs);
    }
    /**
     * The ID for the endpoint group that contains private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; for the peer side of the connection.
     * You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
     * where peer_cidrs is provided with a subnet_id for the VPN service.
     * 
     */
    @Export(name="peerEpGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> peerEpGroupId;

    /**
     * @return The ID for the endpoint group that contains private CIDRs in the form &lt; net_address &gt; / &lt; prefix &gt; for the peer side of the connection.
     * You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
     * where peer_cidrs is provided with a subnet_id for the VPN service.
     * 
     */
    public Output<Optional<String>> peerEpGroupId() {
        return Codegen.optional(this.peerEpGroupId);
    }
    /**
     * The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
     * Typically, this value matches the peer_address value.
     * Changing this updates the existing policy.
     * 
     */
    @Export(name="peerId", refs={String.class}, tree="[0]")
    private Output<String> peerId;

    /**
     * @return The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
     * Typically, this value matches the peer_address value.
     * Changing this updates the existing policy.
     * 
     */
    public Output<String> peerId() {
        return this.peerId;
    }
    /**
     * The pre-shared key. A valid value is any string.
     * 
     */
    @Export(name="psk", refs={String.class}, tree="[0]")
    private Output<String> psk;

    /**
     * @return The pre-shared key. A valid value is any string.
     * 
     */
    public Output<String> psk() {
        return this.psk;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec site connection. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * site connection.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec site connection. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * site connection.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The owner of the connection. Required if admin wants to
     * create a connection for another project. Changing this creates a new connection.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the connection. Required if admin wants to
     * create a connection for another project. Changing this creates a new connection.
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
     * The ID of the VPN service. Changing this creates a new connection.
     * 
     */
    @Export(name="vpnserviceId", refs={String.class}, tree="[0]")
    private Output<String> vpnserviceId;

    /**
     * @return The ID of the VPN service. Changing this creates a new connection.
     * 
     */
    public Output<String> vpnserviceId() {
        return this.vpnserviceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SiteConnection(String name) {
        this(name, SiteConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SiteConnection(String name, SiteConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SiteConnection(String name, SiteConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:vpnaas/siteConnection:SiteConnection", name, args == null ? SiteConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SiteConnection(String name, Output<String> id, @Nullable SiteConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:vpnaas/siteConnection:SiteConnection", name, state, makeResourceOptions(options, id));
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
    public static SiteConnection get(String name, Output<String> id, @Nullable SiteConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SiteConnection(name, id, state, options);
    }
}
