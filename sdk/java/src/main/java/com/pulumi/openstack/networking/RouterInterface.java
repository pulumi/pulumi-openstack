// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.RouterInterfaceArgs;
import com.pulumi.openstack.networking.inputs.RouterInterfaceState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 router interface resource within OpenStack.
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
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Subnet;
 * import com.pulumi.openstack.networking.SubnetArgs;
 * import com.pulumi.openstack.networking.Router;
 * import com.pulumi.openstack.networking.RouterArgs;
 * import com.pulumi.openstack.networking.RouterInterface;
 * import com.pulumi.openstack.networking.RouterInterfaceArgs;
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
 *             .ipVersion(4)
 *             .build());
 * 
 *         var router1 = new Router(&#34;router1&#34;, RouterArgs.builder()        
 *             .externalNetworkId(&#34;f67f0d72-0ddf-11e4-9d95-e1f29f417e2f&#34;)
 *             .build());
 * 
 *         var routerInterface1 = new RouterInterface(&#34;routerInterface1&#34;, RouterInterfaceArgs.builder()        
 *             .routerId(router1.id())
 *             .subnetId(subnet1.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Router Interfaces can be imported using the port `id`, e.g.
 * 
 * $ openstack port list --router &lt;router name or id&gt;
 * 
 * ```sh
 * $ pulumi import openstack:networking/routerInterface:RouterInterface int_1 port_id
 * ```
 * 
 */
@ResourceType(type="openstack:networking/routerInterface:RouterInterface")
public class RouterInterface extends com.pulumi.resources.CustomResource {
    /**
     * A boolean indicating whether the routes from the
     * corresponding router ID should be deleted so that the router interface can
     * be destroyed without any errors. The default value is `false`.
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return A boolean indicating whether the routes from the
     * corresponding router ID should be deleted so that the router interface can
     * be destroyed without any errors. The default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * ID of the port this interface connects to. Changing
     * this creates a new router interface.
     * 
     */
    @Export(name="portId", refs={String.class}, tree="[0]")
    private Output<String> portId;

    /**
     * @return ID of the port this interface connects to. Changing
     * this creates a new router interface.
     * 
     */
    public Output<String> portId() {
        return this.portId;
    }
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router interface.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router interface.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * ID of the router this interface belongs to. Changing
     * this creates a new router interface.
     * 
     */
    @Export(name="routerId", refs={String.class}, tree="[0]")
    private Output<String> routerId;

    /**
     * @return ID of the router this interface belongs to. Changing
     * this creates a new router interface.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }
    /**
     * ID of the subnet this interface connects to. Changing
     * this creates a new router interface.
     * 
     */
    @Export(name="subnetId", refs={String.class}, tree="[0]")
    private Output<String> subnetId;

    /**
     * @return ID of the subnet this interface connects to. Changing
     * this creates a new router interface.
     * 
     */
    public Output<String> subnetId() {
        return this.subnetId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RouterInterface(String name) {
        this(name, RouterInterfaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RouterInterface(String name, RouterInterfaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RouterInterface(String name, RouterInterfaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/routerInterface:RouterInterface", name, args == null ? RouterInterfaceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RouterInterface(String name, Output<String> id, @Nullable RouterInterfaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/routerInterface:RouterInterface", name, state, makeResourceOptions(options, id));
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
    public static RouterInterface get(String name, Output<String> id, @Nullable RouterInterfaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RouterInterface(name, id, state, options);
    }
}
