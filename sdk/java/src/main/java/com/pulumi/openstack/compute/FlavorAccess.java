// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.FlavorAccessArgs;
import com.pulumi.openstack.compute.inputs.FlavorAccessState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a project access for flavor V2 resource within OpenStack.
 * 
 * &gt; **Note:** You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 * 
 * ***
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
 * import com.pulumi.openstack.identity.Project;
 * import com.pulumi.openstack.identity.ProjectArgs;
 * import com.pulumi.openstack.compute.Flavor;
 * import com.pulumi.openstack.compute.FlavorArgs;
 * import com.pulumi.openstack.compute.FlavorAccess;
 * import com.pulumi.openstack.compute.FlavorAccessArgs;
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
 *         var project1 = new Project("project1", ProjectArgs.builder()
 *             .name("my-project")
 *             .build());
 * 
 *         var flavor1 = new Flavor("flavor1", FlavorArgs.builder()
 *             .name("my-flavor")
 *             .ram(8096)
 *             .vcpus(2)
 *             .disk(20)
 *             .isPublic(false)
 *             .build());
 * 
 *         var access1 = new FlavorAccess("access1", FlavorAccessArgs.builder()
 *             .tenantId(project1.id())
 *             .flavorId(flavor1.id())
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
 * This resource can be imported by specifying all two arguments, separated
 * by a forward slash:
 * 
 * ```sh
 * $ pulumi import openstack:compute/flavorAccess:FlavorAccess access_1 flavor_id/tenant_id
 * ```
 * 
 */
@ResourceType(type="openstack:compute/flavorAccess:FlavorAccess")
public class FlavorAccess extends com.pulumi.resources.CustomResource {
    /**
     * The UUID of flavor to use. Changing this creates a new flavor access.
     * 
     */
    @Export(name="flavorId", refs={String.class}, tree="[0]")
    private Output<String> flavorId;

    /**
     * @return The UUID of flavor to use. Changing this creates a new flavor access.
     * 
     */
    public Output<String> flavorId() {
        return this.flavorId;
    }
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor access.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor access.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The UUID of tenant which is allowed to use the flavor.
     * Changing this creates a new flavor access.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The UUID of tenant which is allowed to use the flavor.
     * Changing this creates a new flavor access.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FlavorAccess(java.lang.String name) {
        this(name, FlavorAccessArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FlavorAccess(java.lang.String name, FlavorAccessArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FlavorAccess(java.lang.String name, FlavorAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/flavorAccess:FlavorAccess", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private FlavorAccess(java.lang.String name, Output<java.lang.String> id, @Nullable FlavorAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/flavorAccess:FlavorAccess", name, state, makeResourceOptions(options, id), false);
    }

    private static FlavorAccessArgs makeArgs(FlavorAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FlavorAccessArgs.Empty : args;
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
    public static FlavorAccess get(java.lang.String name, Output<java.lang.String> id, @Nullable FlavorAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FlavorAccess(name, id, state, options);
    }
}
