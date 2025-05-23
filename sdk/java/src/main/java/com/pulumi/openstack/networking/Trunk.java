// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.TrunkArgs;
import com.pulumi.openstack.networking.inputs.TrunkState;
import com.pulumi.openstack.networking.outputs.TrunkSubPort;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a networking V2 trunk resource within OpenStack.
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
 * import com.pulumi.openstack.networking.Port;
 * import com.pulumi.openstack.networking.PortArgs;
 * import com.pulumi.openstack.networking.Trunk;
 * import com.pulumi.openstack.networking.TrunkArgs;
 * import com.pulumi.openstack.networking.inputs.TrunkSubPortArgs;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.inputs.InstanceNetworkArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *             .name("network_1")
 *             .adminStateUp(true)
 *             .build());
 * 
 *         var subnet1 = new Subnet("subnet1", SubnetArgs.builder()
 *             .name("subnet_1")
 *             .networkId(network1.id())
 *             .cidr("192.168.1.0/24")
 *             .ipVersion(4)
 *             .enableDhcp(true)
 *             .noGateway(true)
 *             .build());
 * 
 *         var parentPort1 = new Port("parentPort1", PortArgs.builder()
 *             .name("parent_port_1")
 *             .networkId(network1.id())
 *             .adminStateUp(true)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(subnet1)
 *                 .build());
 * 
 *         var subport1 = new Port("subport1", PortArgs.builder()
 *             .name("subport_1")
 *             .networkId(network1.id())
 *             .adminStateUp(true)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(subnet1)
 *                 .build());
 * 
 *         var trunk1 = new Trunk("trunk1", TrunkArgs.builder()
 *             .name("trunk_1")
 *             .adminStateUp(true)
 *             .portId(parentPort1.id())
 *             .subPorts(TrunkSubPortArgs.builder()
 *                 .portId(subport1.id())
 *                 .segmentationId(1)
 *                 .segmentationType("vlan")
 *                 .build())
 *             .build());
 * 
 *         var instance1 = new Instance("instance1", InstanceArgs.builder()
 *             .name("instance_1")
 *             .securityGroups("default")
 *             .networks(InstanceNetworkArgs.builder()
 *                 .port(trunk1.portId())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="openstack:networking/trunk:Trunk")
public class Trunk extends com.pulumi.resources.CustomResource {
    /**
     * Administrative up/down status for the trunk
     * (must be &#34;true&#34; or &#34;false&#34; if provided). Changing this updates the
     * `admin_state_up` of an existing trunk.
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the trunk
     * (must be &#34;true&#34; or &#34;false&#34; if provided). Changing this updates the
     * `admin_state_up` of an existing trunk.
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * The collection of tags assigned on the trunk, which have been
     * explicitly and implicitly added.
     * 
     */
    @Export(name="allTags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allTags;

    /**
     * @return The collection of tags assigned on the trunk, which have been
     * explicitly and implicitly added.
     * 
     */
    public Output<List<String>> allTags() {
        return this.allTags;
    }
    /**
     * Human-readable description of the trunk. Changing this
     * updates the name of the existing trunk.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-readable description of the trunk. Changing this
     * updates the name of the existing trunk.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A unique name for the trunk. Changing this
     * updates the `name` of an existing trunk.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the trunk. Changing this
     * updates the `name` of an existing trunk.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the port to be used as the parent port of the
     * trunk. This is the port that should be used as the compute instance network
     * port. Changing this creates a new trunk.
     * 
     */
    @Export(name="portId", refs={String.class}, tree="[0]")
    private Output<String> portId;

    /**
     * @return The ID of the port to be used as the parent port of the
     * trunk. This is the port that should be used as the compute instance network
     * port. Changing this creates a new trunk.
     * 
     */
    public Output<String> portId() {
        return this.portId;
    }
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a trunk. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * trunk.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to create a trunk. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * trunk.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The set of ports that will be made subports of the trunk.
     * The structure of each subport is described below.
     * 
     */
    @Export(name="subPorts", refs={List.class,TrunkSubPort.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TrunkSubPort>> subPorts;

    /**
     * @return The set of ports that will be made subports of the trunk.
     * The structure of each subport is described below.
     * 
     */
    public Output<Optional<List<TrunkSubPort>>> subPorts() {
        return Codegen.optional(this.subPorts);
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
     * The owner of the Trunk. Required if admin wants
     * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the Trunk. Required if admin wants
     * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Trunk(java.lang.String name) {
        this(name, TrunkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Trunk(java.lang.String name, TrunkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Trunk(java.lang.String name, TrunkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/trunk:Trunk", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Trunk(java.lang.String name, Output<java.lang.String> id, @Nullable TrunkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/trunk:Trunk", name, state, makeResourceOptions(options, id), false);
    }

    private static TrunkArgs makeArgs(TrunkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TrunkArgs.Empty : args;
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
    public static Trunk get(java.lang.String name, Output<java.lang.String> id, @Nullable TrunkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Trunk(name, id, state, options);
    }
}
