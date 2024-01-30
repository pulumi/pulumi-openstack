// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.PortArgs;
import com.pulumi.openstack.networking.inputs.PortState;
import com.pulumi.openstack.networking.outputs.PortAllowedAddressPair;
import com.pulumi.openstack.networking.outputs.PortBinding;
import com.pulumi.openstack.networking.outputs.PortExtraDhcpOption;
import com.pulumi.openstack.networking.outputs.PortFixedIp;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 port resource within OpenStack.
 * 
 * &gt; **Note:** Ports do not get an IP if the network they are attached
 * to does not have a subnet. If you create the subnet resource in the
 * same run as the port, make sure to use `fixed_ip.subnet_id` or
 * `depends_on` to enforce the subnet resource creation before the port
 * creation is triggered. More info here
 * 
 * ## Example Usage
 * ### Simple port
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Port;
 * import com.pulumi.openstack.networking.PortArgs;
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
 *         var network1 = new Network(&#34;network1&#34;, NetworkArgs.builder()        
 *             .adminStateUp(&#34;true&#34;)
 *             .build());
 * 
 *         var port1 = new Port(&#34;port1&#34;, PortArgs.builder()        
 *             .networkId(network1.id())
 *             .adminStateUp(&#34;true&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Port defining fixed_ip.subnet_id
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Subnet;
 * import com.pulumi.openstack.networking.SubnetArgs;
 * import com.pulumi.openstack.networking.Port;
 * import com.pulumi.openstack.networking.PortArgs;
 * import com.pulumi.openstack.networking.inputs.PortFixedIpArgs;
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
 *         var network1 = new Network(&#34;network1&#34;, NetworkArgs.builder()        
 *             .adminStateUp(&#34;true&#34;)
 *             .build());
 * 
 *         var subnet1 = new Subnet(&#34;subnet1&#34;, SubnetArgs.builder()        
 *             .networkId(network1.id())
 *             .cidr(&#34;192.168.199.0/24&#34;)
 *             .build());
 * 
 *         var port1 = new Port(&#34;port1&#34;, PortArgs.builder()        
 *             .networkId(network1.id())
 *             .adminStateUp(&#34;true&#34;)
 *             .fixedIps(PortFixedIpArgs.builder()
 *                 .subnetId(subnet1.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Port with physical binding information
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Port;
 * import com.pulumi.openstack.networking.PortArgs;
 * import com.pulumi.openstack.networking.inputs.PortBindingArgs;
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
 *         var network1 = new Network(&#34;network1&#34;, NetworkArgs.builder()        
 *             .adminStateUp(&#34;true&#34;)
 *             .build());
 * 
 *         var port1 = new Port(&#34;port1&#34;, PortArgs.builder()        
 *             .networkId(network1.id())
 *             .deviceId(&#34;cdf70fcf-c161-4f24-9c70-96b3f5a54b71&#34;)
 *             .deviceOwner(&#34;baremetal:none&#34;)
 *             .adminStateUp(&#34;true&#34;)
 *             .binding(PortBindingArgs.builder()
 *                 .hostId(&#34;b080b9cf-46e0-4ce8-ad47-0fd4accc872b&#34;)
 *                 .vnicType(&#34;baremetal&#34;)
 *                 .profile(&#34;&#34;&#34;
 * {
 *   &#34;local_link_information&#34;: [
 *     {
 *       &#34;switch_info&#34;: &#34;info1&#34;,
 *       &#34;port_id&#34;: &#34;Ethernet3/4&#34;,
 *       &#34;switch_id&#34;: &#34;12:34:56:78:9A:BC&#34;
 *     },
 *     {
 *       &#34;switch_info&#34;: &#34;info2&#34;,
 *       &#34;port_id&#34;: &#34;Ethernet3/4&#34;,
 *       &#34;switch_id&#34;: &#34;12:34:56:78:9A:BD&#34;
 *     }
 *   ],
 *   &#34;vlan_type&#34;: &#34;allowed&#34;
 * }
 *                 &#34;&#34;&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Notes
 * 
 * ### Ports and Instances
 * 
 * There are some notes to consider when connecting Instances to networks using
 * Ports. Please see the `openstack.compute.Instance` documentation for further
 * documentation.
 * 
 * ## Import
 * 
 * Ports can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:networking/port:Port port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
 * ```
 * 
 */
@ResourceType(type="openstack:networking/port:Port")
public class Port extends com.pulumi.resources.CustomResource {
    /**
     * Administrative up/down status for the port
     * (must be `true` or `false` if provided). Changing this updates the
     * `admin_state_up` of an existing port.
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the port
     * (must be `true` or `false` if provided). Changing this updates the
     * `admin_state_up` of an existing port.
     * 
     */
    public Output<Boolean> adminStateUp() {
        return this.adminStateUp;
    }
    /**
     * The collection of Fixed IP addresses on the port in the
     * order returned by the Network v2 API.
     * 
     */
    @Export(name="allFixedIps", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allFixedIps;

    /**
     * @return The collection of Fixed IP addresses on the port in the
     * order returned by the Network v2 API.
     * 
     */
    public Output<List<String>> allFixedIps() {
        return this.allFixedIps;
    }
    /**
     * The collection of Security Group IDs on the port
     * which have been explicitly and implicitly added.
     * 
     */
    @Export(name="allSecurityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allSecurityGroupIds;

    /**
     * @return The collection of Security Group IDs on the port
     * which have been explicitly and implicitly added.
     * 
     */
    public Output<List<String>> allSecurityGroupIds() {
        return this.allSecurityGroupIds;
    }
    /**
     * The collection of tags assigned on the port, which have been
     * explicitly and implicitly added.
     * 
     */
    @Export(name="allTags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allTags;

    /**
     * @return The collection of tags assigned on the port, which have been
     * explicitly and implicitly added.
     * 
     */
    public Output<List<String>> allTags() {
        return this.allTags;
    }
    /**
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     * 
     */
    @Export(name="allowedAddressPairs", refs={List.class,PortAllowedAddressPair.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PortAllowedAddressPair>> allowedAddressPairs;

    /**
     * @return An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     * 
     */
    public Output<Optional<List<PortAllowedAddressPair>>> allowedAddressPairs() {
        return Codegen.optional(this.allowedAddressPairs);
    }
    /**
     * The port binding allows to specify binding information
     * for the port. The structure is described below.
     * 
     */
    @Export(name="binding", refs={PortBinding.class}, tree="[0]")
    private Output<PortBinding> binding;

    /**
     * @return The port binding allows to specify binding information
     * for the port. The structure is described below.
     * 
     */
    public Output<PortBinding> binding() {
        return this.binding;
    }
    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing port.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-readable description of the port. Changing
     * this updates the `description` of an existing port.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     * 
     */
    @Export(name="deviceId", refs={String.class}, tree="[0]")
    private Output<String> deviceId;

    /**
     * @return The ID of the device attached to the port. Changing this
     * creates a new port.
     * 
     */
    public Output<String> deviceId() {
        return this.deviceId;
    }
    /**
     * The device owner of the port. Changing this creates
     * a new port.
     * 
     */
    @Export(name="deviceOwner", refs={String.class}, tree="[0]")
    private Output<String> deviceOwner;

    /**
     * @return The device owner of the port. Changing this creates
     * a new port.
     * 
     */
    public Output<String> deviceOwner() {
        return this.deviceOwner;
    }
    /**
     * The list of maps representing port DNS assignments.
     * 
     */
    @Export(name="dnsAssignments", refs={List.class,Map.class,String.class,Object.class}, tree="[0,[1,2,3]]")
    private Output<List<Map<String,Object>>> dnsAssignments;

    /**
     * @return The list of maps representing port DNS assignments.
     * 
     */
    public Output<List<Map<String,Object>>> dnsAssignments() {
        return this.dnsAssignments;
    }
    /**
     * The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     * 
     */
    @Export(name="dnsName", refs={String.class}, tree="[0]")
    private Output<String> dnsName;

    /**
     * @return The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     * 
     */
    public Output<String> dnsName() {
        return this.dnsName;
    }
    /**
     * An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     * 
     */
    @Export(name="extraDhcpOptions", refs={List.class,PortExtraDhcpOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PortExtraDhcpOption>> extraDhcpOptions;

    /**
     * @return An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     * 
     */
    public Output<Optional<List<PortExtraDhcpOption>>> extraDhcpOptions() {
        return Codegen.optional(this.extraDhcpOptions);
    }
    /**
     * An array of desired IPs for
     * this port. The structure is described below.
     * 
     */
    @Export(name="fixedIps", refs={List.class,PortFixedIp.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PortFixedIp>> fixedIps;

    /**
     * @return An array of desired IPs for
     * this port. The structure is described below.
     * 
     */
    public Output<Optional<List<PortFixedIp>>> fixedIps() {
        return Codegen.optional(this.fixedIps);
    }
    /**
     * Specify a specific MAC address for the port. Changing
     * this creates a new port.
     * 
     */
    @Export(name="macAddress", refs={String.class}, tree="[0]")
    private Output<String> macAddress;

    /**
     * @return Specify a specific MAC address for the port. Changing
     * this creates a new port.
     * 
     */
    public Output<String> macAddress() {
        return this.macAddress;
    }
    /**
     * A unique name for the port. Changing this
     * updates the `name` of an existing port.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the port. Changing this
     * updates the `name` of an existing port.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the network to attach the port to. Changing
     * this creates a new port.
     * 
     */
    @Export(name="networkId", refs={String.class}, tree="[0]")
    private Output<String> networkId;

    /**
     * @return The ID of the network to attach the port to. Changing
     * this creates a new port.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }
    /**
     * Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     * 
     */
    @Export(name="noFixedIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noFixedIp;

    /**
     * @return Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     * 
     */
    public Output<Optional<Boolean>> noFixedIp() {
        return Codegen.optional(this.noFixedIp);
    }
    /**
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the port will yield to the default
     * behavior of the Networking service, which is to usually apply the &#34;default&#34;
     * security group.
     * 
     */
    @Export(name="noSecurityGroups", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noSecurityGroups;

    /**
     * @return If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the port will yield to the default
     * behavior of the Networking service, which is to usually apply the &#34;default&#34;
     * security group.
     * 
     */
    public Output<Optional<Boolean>> noSecurityGroups() {
        return Codegen.optional(this.noSecurityGroups);
    }
    /**
     * Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of `true`. Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     * 
     */
    @Export(name="portSecurityEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> portSecurityEnabled;

    /**
     * @return Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of `true`. Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     * 
     */
    public Output<Boolean> portSecurityEnabled() {
        return this.portSecurityEnabled;
    }
    /**
     * Reference to the associated QoS policy.
     * 
     */
    @Export(name="qosPolicyId", refs={String.class}, tree="[0]")
    private Output<String> qosPolicyId;

    /**
     * @return Reference to the associated QoS policy.
     * 
     */
    public Output<String> qosPolicyId() {
        return this.qosPolicyId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> securityGroupIds;

    /**
     * @return A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     * 
     */
    public Output<Optional<List<String>>> securityGroupIds() {
        return Codegen.optional(this.securityGroupIds);
    }
    /**
     * A set of string tags for the port.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A set of string tags for the port.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The owner of the port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
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
    public Port(String name) {
        this(name, PortArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Port(String name, PortArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Port(String name, PortArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/port:Port", name, args == null ? PortArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Port(String name, Output<String> id, @Nullable PortState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/port:Port", name, state, makeResourceOptions(options, id));
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
    public static Port get(String name, Output<String> id, @Nullable PortState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Port(name, id, state, options);
    }
}
