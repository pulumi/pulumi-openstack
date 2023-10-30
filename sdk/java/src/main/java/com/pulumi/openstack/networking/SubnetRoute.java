// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.SubnetRouteArgs;
import com.pulumi.openstack.networking.inputs.SubnetRouteState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates a routing entry on a OpenStack V2 subnet.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Router;
 * import com.pulumi.openstack.networking.RouterArgs;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Subnet;
 * import com.pulumi.openstack.networking.SubnetArgs;
 * import com.pulumi.openstack.networking.SubnetRoute;
 * import com.pulumi.openstack.networking.SubnetRouteArgs;
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
 *         var router1 = new Router(&#34;router1&#34;, RouterArgs.builder()        
 *             .adminStateUp(&#34;true&#34;)
 *             .build());
 * 
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
 *         var subnetRoute1 = new SubnetRoute(&#34;subnetRoute1&#34;, SubnetRouteArgs.builder()        
 *             .subnetId(subnet1.id())
 *             .destinationCidr(&#34;10.0.1.0/24&#34;)
 *             .nextHop(&#34;192.168.199.254&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Routing entries can be imported using a combined ID using the following format`&lt;subnet_id&gt;-route-&lt;destination_cidr&gt;-&lt;next_hop&gt;`
 * 
 * ```sh
 *  $ pulumi import openstack:networking/subnetRoute:SubnetRoute subnet_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
 * ```
 * 
 */
@ResourceType(type="openstack:networking/subnetRoute:SubnetRoute")
public class SubnetRoute extends com.pulumi.resources.CustomResource {
    /**
     * CIDR block to match on the packet’s destination IP. Changing
     * this creates a new routing entry.
     * 
     */
    @Export(name="destinationCidr", type=String.class, parameters={})
    private Output<String> destinationCidr;

    /**
     * @return CIDR block to match on the packet’s destination IP. Changing
     * this creates a new routing entry.
     * 
     */
    public Output<String> destinationCidr() {
        return this.destinationCidr;
    }
    /**
     * IP address of the next hop gateway.  Changing
     * this creates a new routing entry.
     * 
     */
    @Export(name="nextHop", type=String.class, parameters={})
    private Output<String> nextHop;

    /**
     * @return IP address of the next hop gateway.  Changing
     * this creates a new routing entry.
     * 
     */
    public Output<String> nextHop() {
        return this.nextHop;
    }
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * ID of the subnet this routing entry belongs to. Changing
     * this creates a new routing entry.
     * 
     */
    @Export(name="subnetId", type=String.class, parameters={})
    private Output<String> subnetId;

    /**
     * @return ID of the subnet this routing entry belongs to. Changing
     * this creates a new routing entry.
     * 
     */
    public Output<String> subnetId() {
        return this.subnetId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SubnetRoute(String name) {
        this(name, SubnetRouteArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SubnetRoute(String name, SubnetRouteArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SubnetRoute(String name, SubnetRouteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/subnetRoute:SubnetRoute", name, args == null ? SubnetRouteArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SubnetRoute(String name, Output<String> id, @Nullable SubnetRouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/subnetRoute:SubnetRoute", name, state, makeResourceOptions(options, id));
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
    public static SubnetRoute get(String name, Output<String> id, @Nullable SubnetRouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SubnetRoute(name, id, state, options);
    }
}
