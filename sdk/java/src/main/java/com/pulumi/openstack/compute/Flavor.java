// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.FlavorArgs;
import com.pulumi.openstack.compute.inputs.FlavorState;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 flavor resource within OpenStack.
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
 * import com.pulumi.openstack.compute.Flavor;
 * import com.pulumi.openstack.compute.FlavorArgs;
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
 *         var test_flavor = new Flavor("test-flavor", FlavorArgs.builder()
 *             .name("my-flavor")
 *             .ram("8096")
 *             .vcpus("2")
 *             .disk("20")
 *             .extraSpecs(Map.ofEntries(
 *                 Map.entry("hw:cpu_policy", "CPU-POLICY"),
 *                 Map.entry("hw:cpu_thread_policy", "CPU-THREAD-POLICY")
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
 * Flavors can be imported using the `ID`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:compute/flavor:Flavor my-flavor 4142e64b-1b35-44a0-9b1e-5affc7af1106
 * ```
 * 
 */
@ResourceType(type="openstack:compute/flavor:Flavor")
public class Flavor extends com.pulumi.resources.CustomResource {
    /**
     * The description of the flavor. Changing this
     * updates the description of the flavor. Requires microversion &gt;= 2.55.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the flavor. Changing this
     * updates the description of the flavor. Requires microversion &gt;= 2.55.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The amount of disk space in GiB to use for the root
     * (/) partition. Changing this creates a new flavor.
     * 
     */
    @Export(name="disk", refs={Integer.class}, tree="[0]")
    private Output<Integer> disk;

    /**
     * @return The amount of disk space in GiB to use for the root
     * (/) partition. Changing this creates a new flavor.
     * 
     */
    public Output<Integer> disk() {
        return this.disk;
    }
    /**
     * The amount of ephemeral in GiB. If unspecified,
     * the default is 0. Changing this creates a new flavor.
     * 
     */
    @Export(name="ephemeral", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ephemeral;

    /**
     * @return The amount of ephemeral in GiB. If unspecified,
     * the default is 0. Changing this creates a new flavor.
     * 
     */
    public Output<Optional<Integer>> ephemeral() {
        return Codegen.optional(this.ephemeral);
    }
    /**
     * Key/Value pairs of metadata for the flavor.
     * 
     */
    @Export(name="extraSpecs", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> extraSpecs;

    /**
     * @return Key/Value pairs of metadata for the flavor.
     * 
     */
    public Output<Map<String,String>> extraSpecs() {
        return this.extraSpecs;
    }
    /**
     * Unique ID (integer or UUID) of flavor to create. Changing
     * this creates a new flavor.
     * 
     */
    @Export(name="flavorId", refs={String.class}, tree="[0]")
    private Output<String> flavorId;

    /**
     * @return Unique ID (integer or UUID) of flavor to create. Changing
     * this creates a new flavor.
     * 
     */
    public Output<String> flavorId() {
        return this.flavorId;
    }
    /**
     * Whether the flavor is public. Changing this creates
     * a new flavor.
     * 
     */
    @Export(name="isPublic", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isPublic;

    /**
     * @return Whether the flavor is public. Changing this creates
     * a new flavor.
     * 
     */
    public Output<Optional<Boolean>> isPublic() {
        return Codegen.optional(this.isPublic);
    }
    /**
     * A unique name for the flavor. Changing this creates a new
     * flavor.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the flavor. Changing this creates a new
     * flavor.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The amount of RAM to use, in megabytes. Changing this
     * creates a new flavor.
     * 
     */
    @Export(name="ram", refs={Integer.class}, tree="[0]")
    private Output<Integer> ram;

    /**
     * @return The amount of RAM to use, in megabytes. Changing this
     * creates a new flavor.
     * 
     */
    public Output<Integer> ram() {
        return this.ram;
    }
    /**
     * The region in which to obtain the V2 Compute client.
     * Flavors are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * Flavors are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * RX/TX bandwith factor. The default is 1. Changing
     * this creates a new flavor.
     * 
     */
    @Export(name="rxTxFactor", refs={Double.class}, tree="[0]")
    private Output</* @Nullable */ Double> rxTxFactor;

    /**
     * @return RX/TX bandwith factor. The default is 1. Changing
     * this creates a new flavor.
     * 
     */
    public Output<Optional<Double>> rxTxFactor() {
        return Codegen.optional(this.rxTxFactor);
    }
    /**
     * The amount of disk space in megabytes to use. If
     * unspecified, the default is 0. Changing this creates a new flavor.
     * 
     */
    @Export(name="swap", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> swap;

    /**
     * @return The amount of disk space in megabytes to use. If
     * unspecified, the default is 0. Changing this creates a new flavor.
     * 
     */
    public Output<Optional<Integer>> swap() {
        return Codegen.optional(this.swap);
    }
    /**
     * The number of virtual CPUs to use. Changing this creates
     * a new flavor.
     * 
     */
    @Export(name="vcpus", refs={Integer.class}, tree="[0]")
    private Output<Integer> vcpus;

    /**
     * @return The number of virtual CPUs to use. Changing this creates
     * a new flavor.
     * 
     */
    public Output<Integer> vcpus() {
        return this.vcpus;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Flavor(java.lang.String name) {
        this(name, FlavorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Flavor(java.lang.String name, FlavorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Flavor(java.lang.String name, FlavorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/flavor:Flavor", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Flavor(java.lang.String name, Output<java.lang.String> id, @Nullable FlavorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/flavor:Flavor", name, state, makeResourceOptions(options, id), false);
    }

    private static FlavorArgs makeArgs(FlavorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FlavorArgs.Empty : args;
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
    public static Flavor get(java.lang.String name, Output<java.lang.String> id, @Nullable FlavorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Flavor(name, id, state, options);
    }
}
