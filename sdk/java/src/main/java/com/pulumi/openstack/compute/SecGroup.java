// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.SecGroupArgs;
import com.pulumi.openstack.compute.inputs.SecGroupState;
import com.pulumi.openstack.compute.outputs.SecGroupRule;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Manages a V2 security group resource within OpenStack.
 * 
 * Please note that managing security groups through the OpenStack Compute API
 * has been deprecated. Unless you are using an older OpenStack environment, it is
 * recommended to use the `openstack.networking.SecGroup`
 * and `openstack.networking.SecGroupRule`
 * resources instead, which uses the OpenStack Networking API.
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
 * import com.pulumi.openstack.compute.SecGroup;
 * import com.pulumi.openstack.compute.SecGroupArgs;
 * import com.pulumi.openstack.compute.inputs.SecGroupRuleArgs;
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
 *         var secgroup1 = new SecGroup("secgroup1", SecGroupArgs.builder()
 *             .name("my_secgroup")
 *             .description("my security group")
 *             .rules(            
 *                 SecGroupRuleArgs.builder()
 *                     .fromPort(22)
 *                     .toPort(22)
 *                     .ipProtocol("tcp")
 *                     .cidr("0.0.0.0/0")
 *                     .build(),
 *                 SecGroupRuleArgs.builder()
 *                     .fromPort(80)
 *                     .toPort(80)
 *                     .ipProtocol("tcp")
 *                     .cidr("0.0.0.0/0")
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Notes
 * 
 * ### Referencing Security Groups
 * 
 * When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
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
 *         var test_server = new Instance("test-server", InstanceArgs.builder()
 *             .name("tf-test")
 *             .imageId("ad091b52-742f-469e-8f3c-fd81cadf0743")
 *             .flavorId("3")
 *             .keyPair("my_key_pair_name")
 *             .securityGroups(secgroup1.name())
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
 * Security Groups can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:compute/secGroup:SecGroup my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
 * ```
 * 
 */
@ResourceType(type="openstack:compute/secGroup:SecGroup")
public class SecGroup extends com.pulumi.resources.CustomResource {
    /**
     * A description for the security group. Changing this
     * updates the `description` of an existing security group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return A description for the security group. Changing this
     * updates the `description` of an existing security group.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * A unique name for the security group. Changing this
     * updates the `name` of an existing security group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the security group. Changing this
     * updates the `name` of an existing security group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a security group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a security group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A rule describing how the security group operates. The
     * rule object structure is documented below. Changing this updates the
     * security group rules. As shown in the example above, multiple rule blocks
     * may be used.
     * 
     */
    @Export(name="rules", refs={List.class,SecGroupRule.class}, tree="[0,1]")
    private Output<List<SecGroupRule>> rules;

    /**
     * @return A rule describing how the security group operates. The
     * rule object structure is documented below. Changing this updates the
     * security group rules. As shown in the example above, multiple rule blocks
     * may be used.
     * 
     */
    public Output<List<SecGroupRule>> rules() {
        return this.rules;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecGroup(String name) {
        this(name, SecGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecGroup(String name, SecGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecGroup(String name, SecGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/secGroup:SecGroup", name, args == null ? SecGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecGroup(String name, Output<String> id, @Nullable SecGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/secGroup:SecGroup", name, state, makeResourceOptions(options, id));
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
    public static SecGroup get(String name, Output<String> id, @Nullable SecGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecGroup(name, id, state, options);
    }
}
