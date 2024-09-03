// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.ServerGroupArgs;
import com.pulumi.openstack.compute.inputs.ServerGroupState;
import com.pulumi.openstack.compute.outputs.ServerGroupRules;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Server Group resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ### Compute service API version 2.63 or below:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.ServerGroup;
 * import com.pulumi.openstack.compute.ServerGroupArgs;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.inputs.InstanceSchedulerHintArgs;
 * import com.pulumi.openstack.compute.inputs.InstanceNetworkArgs;
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
 *         var test_sg = new ServerGroup("test-sg", ServerGroupArgs.builder()
 *             .name("my-sg")
 *             .policies("anti-affinity")
 *             .build());
 * 
 *         var test_instance = new Instance("test-instance", InstanceArgs.builder()
 *             .name("my-instance")
 *             .imageId("ad091b52-742f-469e-8f3c-fd81cadf0743")
 *             .flavorId("3")
 *             .schedulerHints(InstanceSchedulerHintArgs.builder()
 *                 .group(test_sg.id())
 *                 .build())
 *             .networks(InstanceNetworkArgs.builder()
 *                 .name("my_network")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Compute service API version 2.64 or above:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.ServerGroup;
 * import com.pulumi.openstack.compute.ServerGroupArgs;
 * import com.pulumi.openstack.compute.inputs.ServerGroupRulesArgs;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.inputs.InstanceSchedulerHintArgs;
 * import com.pulumi.openstack.compute.inputs.InstanceNetworkArgs;
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
 *         var test_sg = new ServerGroup("test-sg", ServerGroupArgs.builder()
 *             .name("my-sg")
 *             .policies("anti-affinity")
 *             .rules(ServerGroupRulesArgs.builder()
 *                 .maxServerPerHost(3)
 *                 .build())
 *             .build());
 * 
 *         var test_instance = new Instance("test-instance", InstanceArgs.builder()
 *             .name("my-instance")
 *             .imageId("ad091b52-742f-469e-8f3c-fd81cadf0743")
 *             .flavorId("3")
 *             .schedulerHints(InstanceSchedulerHintArgs.builder()
 *                 .group(test_sg.id())
 *                 .build())
 *             .networks(InstanceNetworkArgs.builder()
 *                 .name("my_network")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Policies
 * 
 * * `affinity` - All instances/servers launched in this group will be hosted on
 * the same compute node.
 * 
 * * `anti-affinity` - All instances/servers launched in this group will be
 * hosted on different compute nodes.
 * 
 * * `soft-affinity` - All instances/servers launched in this group will be hosted
 * on the same compute node if possible, but if not possible they
 * still will be scheduled instead of failure. To use this policy your
 * OpenStack environment should support Compute service API 2.15 or above.
 * 
 * * `soft-anti-affinity` - All instances/servers launched in this group will be
 * hosted on different compute nodes if possible, but if not possible they
 * still will be scheduled instead of failure. To use this policy your
 * OpenStack environment should support Compute service API 2.15 or above.
 * 
 * ## Import
 * 
 * Server Groups can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:compute/serverGroup:ServerGroup test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
 * ```
 * 
 */
@ResourceType(type="openstack:compute/serverGroup:ServerGroup")
public class ServerGroup extends com.pulumi.resources.CustomResource {
    /**
     * The instances that are part of this server group.
     * 
     */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> members;

    /**
     * @return The instances that are part of this server group.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * A unique name for the server group. Changing this creates
     * a new server group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the server group. Changing this creates
     * a new server group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A list of exactly one policy name to associate with
     * the server group. See the Policies section for more information. Changing this
     * creates a new server group.
     * 
     */
    @Export(name="policies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> policies;

    /**
     * @return A list of exactly one policy name to associate with
     * the server group. See the Policies section for more information. Changing this
     * creates a new server group.
     * 
     */
    public Output<Optional<List<String>>> policies() {
        return Codegen.optional(this.policies);
    }
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The rules which are applied to specified `policy`. Currently,
     * only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
     * 
     */
    @Export(name="rules", refs={ServerGroupRules.class}, tree="[0]")
    private Output<ServerGroupRules> rules;

    /**
     * @return The rules which are applied to specified `policy`. Currently,
     * only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
     * 
     */
    public Output<ServerGroupRules> rules() {
        return this.rules;
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
    public ServerGroup(java.lang.String name) {
        this(name, ServerGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServerGroup(java.lang.String name, @Nullable ServerGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServerGroup(java.lang.String name, @Nullable ServerGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/serverGroup:ServerGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServerGroup(java.lang.String name, Output<java.lang.String> id, @Nullable ServerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/serverGroup:ServerGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static ServerGroupArgs makeArgs(@Nullable ServerGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServerGroupArgs.Empty : args;
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
    public static ServerGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable ServerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServerGroup(name, id, state, options);
    }
}
