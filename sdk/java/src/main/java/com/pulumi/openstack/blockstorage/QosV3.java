// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.blockstorage.QosV3Args;
import com.pulumi.openstack.blockstorage.inputs.QosV3State;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V3 block storage Quality-Of-Servirce (qos) resource within OpenStack.
 * 
 * &gt; **Note:** This usually requires admin privileges.
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
 * import com.pulumi.openstack.blockstorage.QosV3;
 * import com.pulumi.openstack.blockstorage.QosV3Args;
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
 *         var qos = new QosV3("qos", QosV3Args.builder()
 *             .name("foo")
 *             .consumer("back-end")
 *             .specs(Map.ofEntries(
 *                 Map.entry("read_iops_sec", "40000"),
 *                 Map.entry("write_iops_sec", "40000")
 *             ))
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
 * Qos can be imported using the `qos_id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:blockstorage/qosV3:QosV3 qos 941793f0-0a34-4bc4-b72e-a6326ae58283
 * ```
 * 
 */
@ResourceType(type="openstack:blockstorage/qosV3:QosV3")
public class QosV3 extends com.pulumi.resources.CustomResource {
    /**
     * The consumer of qos. Can be one of `front-end`,
     * `back-end` or `both`. Changing this updates the `consumer` of an
     * existing qos.
     * 
     */
    @Export(name="consumer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> consumer;

    /**
     * @return The consumer of qos. Can be one of `front-end`,
     * `back-end` or `both`. Changing this updates the `consumer` of an
     * existing qos.
     * 
     */
    public Output<Optional<String>> consumer() {
        return Codegen.optional(this.consumer);
    }
    /**
     * Name of the qos.  Changing this creates a new qos.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the qos.  Changing this creates a new qos.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to create the qos. If omitted,
     * the `region` argument of the provider is used. Changing this creates
     * a new qos.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the qos. If omitted,
     * the `region` argument of the provider is used. Changing this creates
     * a new qos.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Key/Value pairs of specs for the qos.
     * 
     */
    @Export(name="specs", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> specs;

    /**
     * @return Key/Value pairs of specs for the qos.
     * 
     */
    public Output<Optional<Map<String,String>>> specs() {
        return Codegen.optional(this.specs);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QosV3(java.lang.String name) {
        this(name, QosV3Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QosV3(java.lang.String name, @Nullable QosV3Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QosV3(java.lang.String name, @Nullable QosV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/qosV3:QosV3", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private QosV3(java.lang.String name, Output<java.lang.String> id, @Nullable QosV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/qosV3:QosV3", name, state, makeResourceOptions(options, id), false);
    }

    private static QosV3Args makeArgs(@Nullable QosV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? QosV3Args.Empty : args;
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
    public static QosV3 get(java.lang.String name, Output<java.lang.String> id, @Nullable QosV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QosV3(name, id, state, options);
    }
}
