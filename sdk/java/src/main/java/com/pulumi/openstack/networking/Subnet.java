// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.SubnetArgs;
import com.pulumi.openstack.networking.inputs.SubnetState;
import com.pulumi.openstack.networking.outputs.SubnetAllocationPool;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron subnet resource within OpenStack.
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
 *             .name("tf_test_network")
 *             .adminStateUp("true")
 *             .build());
 * 
 *         var subnet1 = new Subnet("subnet1", SubnetArgs.builder()
 *             .networkId(network1.id())
 *             .cidr("192.168.199.0/24")
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
 * Subnets can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:networking/subnet:Subnet subnet_1 da4faf16-5546-41e4-8330-4d0002b74048
 * ```
 * 
 */
@ResourceType(type="openstack:networking/subnet:Subnet")
public class Subnet extends com.pulumi.resources.CustomResource {
    /**
     * The collection of ags assigned on the subnet, which have been
     * explicitly and implicitly added.
     * 
     */
    @Export(name="allTags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allTags;

    /**
     * @return The collection of ags assigned on the subnet, which have been
     * explicitly and implicitly added.
     * 
     */
    public Output<List<String>> allTags() {
        return this.allTags;
    }
    /**
     * A block declaring the start and end range of
     * the IP addresses available for use with DHCP in this subnet. Multiple
     * `allocation_pool` blocks can be declared, providing the subnet with more
     * than one range of IP addresses to use with DHCP. However, each IP range
     * must be from the same CIDR that the subnet is part of.
     * The `allocation_pool` block is documented below.
     * 
     */
    @Export(name="allocationPools", refs={List.class,SubnetAllocationPool.class}, tree="[0,1]")
    private Output<List<SubnetAllocationPool>> allocationPools;

    /**
     * @return A block declaring the start and end range of
     * the IP addresses available for use with DHCP in this subnet. Multiple
     * `allocation_pool` blocks can be declared, providing the subnet with more
     * than one range of IP addresses to use with DHCP. However, each IP range
     * must be from the same CIDR that the subnet is part of.
     * The `allocation_pool` block is documented below.
     * 
     */
    public Output<List<SubnetAllocationPool>> allocationPools() {
        return this.allocationPools;
    }
    /**
     * CIDR representing IP range for this subnet, based on IP
     * version. You can omit this option if you are creating a subnet from a
     * subnet pool.
     * 
     */
    @Export(name="cidr", refs={String.class}, tree="[0]")
    private Output<String> cidr;

    /**
     * @return CIDR representing IP range for this subnet, based on IP
     * version. You can omit this option if you are creating a subnet from a
     * subnet pool.
     * 
     */
    public Output<String> cidr() {
        return this.cidr;
    }
    /**
     * Human-readable description of the subnet. Changing this
     * updates the name of the existing subnet.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-readable description of the subnet. Changing this
     * updates the name of the existing subnet.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * An array of DNS name server names used by hosts
     * in this subnet. Changing this updates the DNS name servers for the existing
     * subnet.
     * 
     */
    @Export(name="dnsNameservers", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dnsNameservers;

    /**
     * @return An array of DNS name server names used by hosts
     * in this subnet. Changing this updates the DNS name servers for the existing
     * subnet.
     * 
     */
    public Output<Optional<List<String>>> dnsNameservers() {
        return Codegen.optional(this.dnsNameservers);
    }
    /**
     * Whether to publish DNS records for IPs
     * from this subnet. Defaults is false.
     * 
     */
    @Export(name="dnsPublishFixedIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dnsPublishFixedIp;

    /**
     * @return Whether to publish DNS records for IPs
     * from this subnet. Defaults is false.
     * 
     */
    public Output<Optional<Boolean>> dnsPublishFixedIp() {
        return Codegen.optional(this.dnsPublishFixedIp);
    }
    /**
     * The administrative state of the network.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value enables or
     * disables the DHCP capabilities of the existing subnet. Defaults to true.
     * 
     */
    @Export(name="enableDhcp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableDhcp;

    /**
     * @return The administrative state of the network.
     * Acceptable values are &#34;true&#34; and &#34;false&#34;. Changing this value enables or
     * disables the DHCP capabilities of the existing subnet. Defaults to true.
     * 
     */
    public Output<Optional<Boolean>> enableDhcp() {
        return Codegen.optional(this.enableDhcp);
    }
    /**
     * Default gateway used by devices in this subnet.
     * Leaving this blank and not setting `no_gateway` will cause a default
     * gateway of `.1` to be used. Changing this updates the gateway IP of the
     * existing subnet.
     * 
     */
    @Export(name="gatewayIp", refs={String.class}, tree="[0]")
    private Output<String> gatewayIp;

    /**
     * @return Default gateway used by devices in this subnet.
     * Leaving this blank and not setting `no_gateway` will cause a default
     * gateway of `.1` to be used. Changing this updates the gateway IP of the
     * existing subnet.
     * 
     */
    public Output<String> gatewayIp() {
        return this.gatewayIp;
    }
    /**
     * IP version, either 4 (default) or 6. Changing this creates a
     * new subnet.
     * 
     */
    @Export(name="ipVersion", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ipVersion;

    /**
     * @return IP version, either 4 (default) or 6. Changing this creates a
     * new subnet.
     * 
     */
    public Output<Optional<Integer>> ipVersion() {
        return Codegen.optional(this.ipVersion);
    }
    /**
     * The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     * 
     */
    @Export(name="ipv6AddressMode", refs={String.class}, tree="[0]")
    private Output<String> ipv6AddressMode;

    /**
     * @return The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     * 
     */
    public Output<String> ipv6AddressMode() {
        return this.ipv6AddressMode;
    }
    /**
     * The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     * 
     */
    @Export(name="ipv6RaMode", refs={String.class}, tree="[0]")
    private Output<String> ipv6RaMode;

    /**
     * @return The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     * 
     */
    public Output<String> ipv6RaMode() {
        return this.ipv6RaMode;
    }
    /**
     * The name of the subnet. Changing this updates the name of
     * the existing subnet.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the subnet. Changing this updates the name of
     * the existing subnet.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The UUID of the parent network. Changing this
     * creates a new subnet.
     * 
     */
    @Export(name="networkId", refs={String.class}, tree="[0]")
    private Output<String> networkId;

    /**
     * @return The UUID of the parent network. Changing this
     * creates a new subnet.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }
    /**
     * Do not set a gateway IP on this subnet. Changing
     * this removes or adds a default gateway IP of the existing subnet.
     * 
     */
    @Export(name="noGateway", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noGateway;

    /**
     * @return Do not set a gateway IP on this subnet. Changing
     * this removes or adds a default gateway IP of the existing subnet.
     * 
     */
    public Output<Optional<Boolean>> noGateway() {
        return Codegen.optional(this.noGateway);
    }
    /**
     * The prefix length to use when creating a subnet
     * from a subnet pool. The default subnet pool prefix length that was defined
     * when creating the subnet pool will be used if not provided. Changing this
     * creates a new subnet.
     * 
     */
    @Export(name="prefixLength", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> prefixLength;

    /**
     * @return The prefix length to use when creating a subnet
     * from a subnet pool. The default subnet pool prefix length that was defined
     * when creating the subnet pool will be used if not provided. Changing this
     * creates a new subnet.
     * 
     */
    public Output<Optional<Integer>> prefixLength() {
        return Codegen.optional(this.prefixLength);
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnet.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnet.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * An array of service types used by the subnet.
     * Changing this updates the service types for the existing subnet.
     * 
     */
    @Export(name="serviceTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> serviceTypes;

    /**
     * @return An array of service types used by the subnet.
     * Changing this updates the service types for the existing subnet.
     * 
     */
    public Output<List<String>> serviceTypes() {
        return this.serviceTypes;
    }
    /**
     * The ID of the subnetpool associated with the subnet.
     * 
     */
    @Export(name="subnetpoolId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subnetpoolId;

    /**
     * @return The ID of the subnetpool associated with the subnet.
     * 
     */
    public Output<Optional<String>> subnetpoolId() {
        return Codegen.optional(this.subnetpoolId);
    }
    /**
     * A set of string tags for the subnet.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A set of string tags for the subnet.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The owner of the subnet. Required if admin wants to
     * create a subnet for another tenant. Changing this creates a new subnet.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the subnet. Required if admin wants to
     * create a subnet for another tenant. Changing this creates a new subnet.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
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
    public Subnet(java.lang.String name) {
        this(name, SubnetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Subnet(java.lang.String name, SubnetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Subnet(java.lang.String name, SubnetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/subnet:Subnet", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Subnet(java.lang.String name, Output<java.lang.String> id, @Nullable SubnetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/subnet:Subnet", name, state, makeResourceOptions(options, id), false);
    }

    private static SubnetArgs makeArgs(SubnetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SubnetArgs.Empty : args;
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
    public static Subnet get(java.lang.String name, Output<java.lang.String> id, @Nullable SubnetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Subnet(name, id, state, options);
    }
}
