// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.vpnaas.EndpointGroupArgs;
import com.pulumi.openstack.vpnaas.inputs.EndpointGroupState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron Endpoint Group resource within OpenStack.
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
 * import com.pulumi.openstack.vpnaas.EndpointGroup;
 * import com.pulumi.openstack.vpnaas.EndpointGroupArgs;
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
 *         var group1 = new EndpointGroup("group1", EndpointGroupArgs.builder()
 *             .name("Group 1")
 *             .type("cidr")
 *             .endpoints(            
 *                 "10.2.0.0/24",
 *                 "10.3.0.0/24")
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
 * Groups can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:vpnaas/endpointGroup:EndpointGroup group_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 * 
 */
@ResourceType(type="openstack:vpnaas/endpointGroup:EndpointGroup")
public class EndpointGroup extends com.pulumi.resources.CustomResource {
    /**
     * The human-readable description for the group.
     * Changing this updates the description of the existing group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The human-readable description for the group.
     * Changing this updates the description of the existing group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * List of endpoints of the same type, for the endpoint group. The values will depend on the type.
     * Changing this creates a new group.
     * 
     */
    @Export(name="endpoints", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> endpoints;

    /**
     * @return List of endpoints of the same type, for the endpoint group. The values will depend on the type.
     * Changing this creates a new group.
     * 
     */
    public Output<Optional<List<String>>> endpoints() {
        return Codegen.optional(this.endpoints);
    }
    /**
     * The name of the group. Changing this updates the name of
     * the existing group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the group. Changing this updates the name of
     * the existing group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an endpoint group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * group.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an endpoint group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * group.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The owner of the group. Required if admin wants to
     * create an endpoint group for another project. Changing this creates a new group.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the group. Required if admin wants to
     * create an endpoint group for another project. Changing this creates a new group.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
     * Changing this creates a new group.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
     * Changing this creates a new group.
     * 
     */
    public Output<String> type() {
        return this.type;
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
    public EndpointGroup(java.lang.String name) {
        this(name, EndpointGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EndpointGroup(java.lang.String name, @Nullable EndpointGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EndpointGroup(java.lang.String name, @Nullable EndpointGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:vpnaas/endpointGroup:EndpointGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EndpointGroup(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:vpnaas/endpointGroup:EndpointGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static EndpointGroupArgs makeArgs(@Nullable EndpointGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EndpointGroupArgs.Empty : args;
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
    public static EndpointGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EndpointGroup(name, id, state, options);
    }
}
