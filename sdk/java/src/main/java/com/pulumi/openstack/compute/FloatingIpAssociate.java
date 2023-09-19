// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.FloatingIpAssociateArgs;
import com.pulumi.openstack.compute.inputs.FloatingIpAssociateState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Associate a floating IP to an instance.
 * 
 * ## Example Usage
 * ### Automatically detect the correct network
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.networking.FloatingIp;
 * import com.pulumi.openstack.networking.FloatingIpArgs;
 * import com.pulumi.openstack.compute.FloatingIpAssociate;
 * import com.pulumi.openstack.compute.FloatingIpAssociateArgs;
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
 *         var instance1 = new Instance(&#34;instance1&#34;, InstanceArgs.builder()        
 *             .imageId(&#34;ad091b52-742f-469e-8f3c-fd81cadf0743&#34;)
 *             .flavorId(3)
 *             .keyPair(&#34;my_key_pair_name&#34;)
 *             .securityGroups(&#34;default&#34;)
 *             .build());
 * 
 *         var fip1FloatingIp = new FloatingIp(&#34;fip1FloatingIp&#34;, FloatingIpArgs.builder()        
 *             .pool(&#34;my_pool&#34;)
 *             .build());
 * 
 *         var fip1FloatingIpAssociate = new FloatingIpAssociate(&#34;fip1FloatingIpAssociate&#34;, FloatingIpAssociateArgs.builder()        
 *             .floatingIp(fip1FloatingIp.address())
 *             .instanceId(instance1.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Explicitly set the network to attach to
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.inputs.InstanceNetworkArgs;
 * import com.pulumi.openstack.networking.FloatingIp;
 * import com.pulumi.openstack.networking.FloatingIpArgs;
 * import com.pulumi.openstack.compute.FloatingIpAssociate;
 * import com.pulumi.openstack.compute.FloatingIpAssociateArgs;
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
 *         var instance1 = new Instance(&#34;instance1&#34;, InstanceArgs.builder()        
 *             .imageId(&#34;ad091b52-742f-469e-8f3c-fd81cadf0743&#34;)
 *             .flavorId(3)
 *             .keyPair(&#34;my_key_pair_name&#34;)
 *             .securityGroups(&#34;default&#34;)
 *             .networks(            
 *                 InstanceNetworkArgs.builder()
 *                     .name(&#34;my_network&#34;)
 *                     .build(),
 *                 InstanceNetworkArgs.builder()
 *                     .name(&#34;default&#34;)
 *                     .build())
 *             .build());
 * 
 *         var fip1FloatingIp = new FloatingIp(&#34;fip1FloatingIp&#34;, FloatingIpArgs.builder()        
 *             .pool(&#34;my_pool&#34;)
 *             .build());
 * 
 *         var fip1FloatingIpAssociate = new FloatingIpAssociate(&#34;fip1FloatingIpAssociate&#34;, FloatingIpAssociateArgs.builder()        
 *             .floatingIp(fip1FloatingIp.address())
 *             .instanceId(instance1.id())
 *             .fixedIp(instance1.networks().applyValue(networks -&gt; networks[1].fixedIpV4()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported by specifying all three arguments, separated by a forward slash:
 * 
 * ```sh
 *  $ pulumi import openstack:compute/floatingIpAssociate:FloatingIpAssociate fip_1 floating_ip/instance_id/fixed_ip
 * ```
 * 
 */
@ResourceType(type="openstack:compute/floatingIpAssociate:FloatingIpAssociate")
public class FloatingIpAssociate extends com.pulumi.resources.CustomResource {
    /**
     * The specific IP address to direct traffic to.
     * 
     */
    @Export(name="fixedIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> fixedIp;

    /**
     * @return The specific IP address to direct traffic to.
     * 
     */
    public Output<Optional<String>> fixedIp() {
        return Codegen.optional(this.fixedIp);
    }
    /**
     * The floating IP to associate.
     * 
     */
    @Export(name="floatingIp", refs={String.class}, tree="[0]")
    private Output<String> floatingIp;

    /**
     * @return The floating IP to associate.
     * 
     */
    public Output<String> floatingIp() {
        return this.floatingIp;
    }
    /**
     * The instance to associte the floating IP with.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The instance to associte the floating IP with.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new floatingip_associate.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new floatingip_associate.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    @Export(name="waitUntilAssociated", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> waitUntilAssociated;

    public Output<Optional<Boolean>> waitUntilAssociated() {
        return Codegen.optional(this.waitUntilAssociated);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FloatingIpAssociate(String name) {
        this(name, FloatingIpAssociateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FloatingIpAssociate(String name, FloatingIpAssociateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FloatingIpAssociate(String name, FloatingIpAssociateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/floatingIpAssociate:FloatingIpAssociate", name, args == null ? FloatingIpAssociateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FloatingIpAssociate(String name, Output<String> id, @Nullable FloatingIpAssociateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/floatingIpAssociate:FloatingIpAssociate", name, state, makeResourceOptions(options, id));
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
    public static FloatingIpAssociate get(String name, Output<String> id, @Nullable FloatingIpAssociateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FloatingIpAssociate(name, id, state, options);
    }
}
