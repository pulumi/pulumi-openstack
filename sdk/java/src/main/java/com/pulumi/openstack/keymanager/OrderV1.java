// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.keymanager.OrderV1Args;
import com.pulumi.openstack.keymanager.inputs.OrderV1State;
import com.pulumi.openstack.keymanager.outputs.OrderV1Meta;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a V1 Barbican order resource within OpenStack.
 * 
 * ## Example Usage
 * ### Symmetric key order
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.keymanager.OrderV1;
 * import com.pulumi.openstack.keymanager.OrderV1Args;
 * import com.pulumi.openstack.keymanager.inputs.OrderV1MetaArgs;
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
 *         var order1 = new OrderV1(&#34;order1&#34;, OrderV1Args.builder()        
 *             .meta(OrderV1MetaArgs.builder()
 *                 .algorithm(&#34;aes&#34;)
 *                 .bitLength(256)
 *                 .mode(&#34;cbc&#34;)
 *                 .name(&#34;mysecret&#34;)
 *                 .build())
 *             .type(&#34;key&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Asymmetric key pair order
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.keymanager.OrderV1;
 * import com.pulumi.openstack.keymanager.OrderV1Args;
 * import com.pulumi.openstack.keymanager.inputs.OrderV1MetaArgs;
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
 *         var order1 = new OrderV1(&#34;order1&#34;, OrderV1Args.builder()        
 *             .meta(OrderV1MetaArgs.builder()
 *                 .algorithm(&#34;rsa&#34;)
 *                 .bitLength(4096)
 *                 .name(&#34;mysecret&#34;)
 *                 .build())
 *             .type(&#34;asymmetric&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Orders can be imported using the order id (the last part of the order reference), e.g.:
 * 
 * ```sh
 *  $ pulumi import openstack:keymanager/orderV1:OrderV1 order_1 0c6cd26a-c012-4d7b-8034-057c0f1c2953
 * ```
 * 
 */
@ResourceType(type="openstack:keymanager/orderV1:OrderV1")
public class OrderV1 extends com.pulumi.resources.CustomResource {
    /**
     * The container reference / where to find the container.
     * 
     */
    @Export(name="containerRef", type=String.class, parameters={})
    private Output<String> containerRef;

    /**
     * @return The container reference / where to find the container.
     * 
     */
    public Output<String> containerRef() {
        return this.containerRef;
    }
    /**
     * The date the order was created.
     * 
     */
    @Export(name="created", type=String.class, parameters={})
    private Output<String> created;

    /**
     * @return The date the order was created.
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * The creator of the order.
     * 
     */
    @Export(name="creatorId", type=String.class, parameters={})
    private Output<String> creatorId;

    /**
     * @return The creator of the order.
     * 
     */
    public Output<String> creatorId() {
        return this.creatorId;
    }
    /**
     * Dictionary containing the order metadata used to generate the order. The structure is described below.
     * 
     */
    @Export(name="meta", type=OrderV1Meta.class, parameters={})
    private Output<OrderV1Meta> meta;

    /**
     * @return Dictionary containing the order metadata used to generate the order. The structure is described below.
     * 
     */
    public Output<OrderV1Meta> meta() {
        return this.meta;
    }
    /**
     * The order reference / where to find the order.
     * 
     */
    @Export(name="orderRef", type=String.class, parameters={})
    private Output<String> orderRef;

    /**
     * @return The order reference / where to find the order.
     * 
     */
    public Output<String> orderRef() {
        return this.orderRef;
    }
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a order. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 order.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a order. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 order.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The secret reference / where to find the secret.
     * 
     */
    @Export(name="secretRef", type=String.class, parameters={})
    private Output<String> secretRef;

    /**
     * @return The secret reference / where to find the secret.
     * 
     */
    public Output<String> secretRef() {
        return this.secretRef;
    }
    /**
     * The status of the order.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The status of the order.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The sub status of the order.
     * 
     */
    @Export(name="subStatus", type=String.class, parameters={})
    private Output<String> subStatus;

    /**
     * @return The sub status of the order.
     * 
     */
    public Output<String> subStatus() {
        return this.subStatus;
    }
    /**
     * The sub status message of the order.
     * 
     */
    @Export(name="subStatusMessage", type=String.class, parameters={})
    private Output<String> subStatusMessage;

    /**
     * @return The sub status message of the order.
     * 
     */
    public Output<String> subStatusMessage() {
        return this.subStatusMessage;
    }
    /**
     * The type of key to be generated. Must be one of `asymmetric`, `key`.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The type of key to be generated. Must be one of `asymmetric`, `key`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The date the order was last updated.
     * 
     */
    @Export(name="updated", type=String.class, parameters={})
    private Output<String> updated;

    /**
     * @return The date the order was last updated.
     * 
     */
    public Output<String> updated() {
        return this.updated;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OrderV1(String name) {
        this(name, OrderV1Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrderV1(String name, OrderV1Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrderV1(String name, OrderV1Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:keymanager/orderV1:OrderV1", name, args == null ? OrderV1Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OrderV1(String name, Output<String> id, @Nullable OrderV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:keymanager/orderV1:OrderV1", name, state, makeResourceOptions(options, id));
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
    public static OrderV1 get(String name, Output<String> id, @Nullable OrderV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrderV1(name, id, state, options);
    }
}
