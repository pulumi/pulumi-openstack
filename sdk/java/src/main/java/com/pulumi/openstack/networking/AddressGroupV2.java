// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.AddressGroupV2Args;
import com.pulumi.openstack.networking.inputs.AddressGroupV2State;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 neutron address group resource within OpenStack.
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
 * import com.pulumi.openstack.networking.AddressGroupV2;
 * import com.pulumi.openstack.networking.AddressGroupV2Args;
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
 *         var group1 = new AddressGroupV2("group1", AddressGroupV2Args.builder()
 *             .name("group_1")
 *             .description("My neutron address group")
 *             .addresses(            
 *                 "192.168.0.1/32",
 *                 "2001:db8::1/128")
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
 * Address Groups can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:networking/addressGroupV2:AddressGroupV2 group_1 782fef9c-d03c-400a-9735-2f9af5681cb3
 * ```
 * 
 */
@ResourceType(type="openstack:networking/addressGroupV2:AddressGroupV2")
public class AddressGroupV2 extends com.pulumi.resources.CustomResource {
    /**
     * A list of CIDR blocks that define the addresses in
     * the address group. Each address must be a valid IPv4 or IPv6 CIDR block.
     * 
     */
    @Export(name="addresses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> addresses;

    /**
     * @return A list of CIDR blocks that define the addresses in
     * the address group. Each address must be a valid IPv4 or IPv6 CIDR block.
     * 
     */
    public Output<List<String>> addresses() {
        return this.addresses;
    }
    /**
     * A description of the address group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the address group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A name of the address group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name of the address group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The owner of the address group. Required if admin
     * wants to create a group for a specific project. Changing this creates a new
     * address group.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The owner of the address group. Required if admin
     * wants to create a group for a specific project. Changing this creates a new
     * address group.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V2 networking client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new address group.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new address group.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AddressGroupV2(java.lang.String name) {
        this(name, AddressGroupV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AddressGroupV2(java.lang.String name, AddressGroupV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AddressGroupV2(java.lang.String name, AddressGroupV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/addressGroupV2:AddressGroupV2", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AddressGroupV2(java.lang.String name, Output<java.lang.String> id, @Nullable AddressGroupV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/addressGroupV2:AddressGroupV2", name, state, makeResourceOptions(options, id), false);
    }

    private static AddressGroupV2Args makeArgs(AddressGroupV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AddressGroupV2Args.Empty : args;
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
    public static AddressGroupV2 get(java.lang.String name, Output<java.lang.String> id, @Nullable AddressGroupV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AddressGroupV2(name, id, state, options);
    }
}
