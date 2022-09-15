// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.PortForwardingV2Args;
import com.pulumi.openstack.networking.inputs.PortForwardingV2State;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 portforwarding resource within OpenStack.
 * 
 * ## Example Usage
 * ### Simple portforwarding
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.PortForwardingV2;
 * import com.pulumi.openstack.networking.PortForwardingV2Args;
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
 *         var pf1 = new PortForwardingV2(&#34;pf1&#34;, PortForwardingV2Args.builder()        
 *             .externalPort(7233)
 *             .floatingipId(&#34;7a52eb59-7d47-415d-a884-046666a6fbae&#34;)
 *             .internalPort(25)
 *             .internalPortId(&#34;b930d7f6-ceb7-40a0-8b81-a425dd994ccf&#34;)
 *             .protocol(&#34;tcp&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="openstack:networking/portForwardingV2:PortForwardingV2")
public class PortForwardingV2 extends com.pulumi.resources.CustomResource {
    /**
     * A text describing the port forwarding. Changing this
     * updates the `description` of an existing port forwarding.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A text describing the port forwarding. Changing this
     * updates the `description` of an existing port forwarding.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The TCP/UDP/other protocol port number of the port forwarding. Changing this
     * updates the `external_port` of an existing port forwarding.
     * 
     */
    @Export(name="externalPort", type=Integer.class, parameters={})
    private Output<Integer> externalPort;

    /**
     * @return The TCP/UDP/other protocol port number of the port forwarding. Changing this
     * updates the `external_port` of an existing port forwarding.
     * 
     */
    public Output<Integer> externalPort() {
        return this.externalPort;
    }
    /**
     * The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
     * 
     */
    @Export(name="floatingipId", type=String.class, parameters={})
    private Output<String> floatingipId;

    /**
     * @return The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
     * 
     */
    public Output<String> floatingipId() {
        return this.floatingipId;
    }
    /**
     * The fixed IPv4 address of the Neutron port associated with the port forwarding.
     * Changing this updates the `internal_ip_address` of an existing port forwarding.
     * 
     */
    @Export(name="internalIpAddress", type=String.class, parameters={})
    private Output<String> internalIpAddress;

    /**
     * @return The fixed IPv4 address of the Neutron port associated with the port forwarding.
     * Changing this updates the `internal_ip_address` of an existing port forwarding.
     * 
     */
    public Output<String> internalIpAddress() {
        return this.internalIpAddress;
    }
    /**
     * The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
     * port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
     * 
     */
    @Export(name="internalPort", type=Integer.class, parameters={})
    private Output<Integer> internalPort;

    /**
     * @return The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
     * port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
     * 
     */
    public Output<Integer> internalPort() {
        return this.internalPort;
    }
    /**
     * The ID of the Neutron port associated with the port forwarding. Changing
     * this updates the `internal_port_id` of an existing port forwarding.
     * 
     */
    @Export(name="internalPortId", type=String.class, parameters={})
    private Output<String> internalPortId;

    /**
     * @return The ID of the Neutron port associated with the port forwarding. Changing
     * this updates the `internal_port_id` of an existing port forwarding.
     * 
     */
    public Output<String> internalPortId() {
        return this.internalPortId;
    }
    /**
     * The IP protocol used in the port forwarding. Changing this updates the `protocol`
     * of an existing port forwarding.
     * 
     */
    @Export(name="protocol", type=String.class, parameters={})
    private Output<String> protocol;

    /**
     * @return The IP protocol used in the port forwarding. Changing this updates the `protocol`
     * of an existing port forwarding.
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port forwarding. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port forwarding.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port forwarding. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port forwarding.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PortForwardingV2(String name) {
        this(name, PortForwardingV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PortForwardingV2(String name, PortForwardingV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PortForwardingV2(String name, PortForwardingV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/portForwardingV2:PortForwardingV2", name, args == null ? PortForwardingV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PortForwardingV2(String name, Output<String> id, @Nullable PortForwardingV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/portForwardingV2:PortForwardingV2", name, state, makeResourceOptions(options, id));
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
    public static PortForwardingV2 get(String name, Output<String> id, @Nullable PortForwardingV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PortForwardingV2(name, id, state, options);
    }
}