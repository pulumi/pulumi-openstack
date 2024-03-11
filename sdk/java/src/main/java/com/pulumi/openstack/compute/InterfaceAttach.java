// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.InterfaceAttachArgs;
import com.pulumi.openstack.compute.inputs.InterfaceAttachState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Attaches a Network Interface (a Port) to an Instance using the OpenStack
 * Compute (Nova) v2 API.
 * 
 * ## Example Usage
 * 
 * ### Basic Attachment
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.InterfaceAttach;
 * import com.pulumi.openstack.compute.InterfaceAttachArgs;
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
 *         var instance1 = new Instance(&#34;instance1&#34;, InstanceArgs.builder()        
 *             .securityGroups(&#34;default&#34;)
 *             .build());
 * 
 *         var ai1 = new InterfaceAttach(&#34;ai1&#34;, InterfaceAttachArgs.builder()        
 *             .instanceId(instance1.id())
 *             .networkId(openstack_networking_port_v2.network_1().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Attachment Specifying a Fixed IP
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.InterfaceAttach;
 * import com.pulumi.openstack.compute.InterfaceAttachArgs;
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
 *         var instance1 = new Instance(&#34;instance1&#34;, InstanceArgs.builder()        
 *             .securityGroups(&#34;default&#34;)
 *             .build());
 * 
 *         var ai1 = new InterfaceAttach(&#34;ai1&#34;, InterfaceAttachArgs.builder()        
 *             .instanceId(instance1.id())
 *             .networkId(openstack_networking_port_v2.network_1().id())
 *             .fixedIp(&#34;10.0.10.10&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Attachment Using an Existing Port
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
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
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.InterfaceAttach;
 * import com.pulumi.openstack.compute.InterfaceAttachArgs;
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
 *         var instance1 = new Instance(&#34;instance1&#34;, InstanceArgs.builder()        
 *             .securityGroups(&#34;default&#34;)
 *             .build());
 * 
 *         var ai1 = new InterfaceAttach(&#34;ai1&#34;, InterfaceAttachArgs.builder()        
 *             .instanceId(instance1.id())
 *             .portId(port1.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Interface Attachments can be imported using the Instance ID and Port ID
 * separated by a slash, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:compute/interfaceAttach:InterfaceAttach ai_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
 * ```
 * 
 */
@ResourceType(type="openstack:compute/interfaceAttach:InterfaceAttach")
public class InterfaceAttach extends com.pulumi.resources.CustomResource {
    /**
     * An IP address to assosciate with the port.
     * _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
     * 
     */
    @Export(name="fixedIp", refs={String.class}, tree="[0]")
    private Output<String> fixedIp;

    /**
     * @return An IP address to assosciate with the port.
     * _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
     * 
     */
    public Output<String> fixedIp() {
        return this.fixedIp;
    }
    /**
     * The ID of the Instance to attach the Port or Network to.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the Instance to attach the Port or Network to.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The ID of the Network to attach to an Instance. A port will be created automatically.
     * _NOTE_: This option and `port_id` are mutually exclusive.
     * 
     */
    @Export(name="networkId", refs={String.class}, tree="[0]")
    private Output<String> networkId;

    /**
     * @return The ID of the Network to attach to an Instance. A port will be created automatically.
     * _NOTE_: This option and `port_id` are mutually exclusive.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }
    /**
     * The ID of the Port to attach to an Instance.
     * _NOTE_: This option and `network_id` are mutually exclusive.
     * 
     */
    @Export(name="portId", refs={String.class}, tree="[0]")
    private Output<String> portId;

    /**
     * @return The ID of the Port to attach to an Instance.
     * _NOTE_: This option and `network_id` are mutually exclusive.
     * 
     */
    public Output<String> portId() {
        return this.portId;
    }
    /**
     * The region in which to create the interface attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new attachment.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the interface attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new attachment.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InterfaceAttach(String name) {
        this(name, InterfaceAttachArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InterfaceAttach(String name, InterfaceAttachArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InterfaceAttach(String name, InterfaceAttachArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/interfaceAttach:InterfaceAttach", name, args == null ? InterfaceAttachArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InterfaceAttach(String name, Output<String> id, @Nullable InterfaceAttachState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/interfaceAttach:InterfaceAttach", name, state, makeResourceOptions(options, id));
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
    public static InterfaceAttach get(String name, Output<String> id, @Nullable InterfaceAttachState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InterfaceAttach(name, id, state, options);
    }
}
