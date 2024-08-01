// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.AddressScopeArgs;
import com.pulumi.openstack.networking.inputs.AddressScopeState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron addressscope resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ### Create an Address-scope
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.AddressScope;
 * import com.pulumi.openstack.networking.AddressScopeArgs;
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
 *         var addressscope1 = new AddressScope("addressscope1", AddressScopeArgs.builder()
 *             .name("addressscope_1")
 *             .ipVersion(6)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Create a Subnet Pool from an Address-scope
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.AddressScope;
 * import com.pulumi.openstack.networking.AddressScopeArgs;
 * import com.pulumi.openstack.networking.SubnetPool;
 * import com.pulumi.openstack.networking.SubnetPoolArgs;
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
 *         var addressscope1 = new AddressScope("addressscope1", AddressScopeArgs.builder()
 *             .name("addressscope_1")
 *             .ipVersion(6)
 *             .build());
 * 
 *         var subnetpool1 = new SubnetPool("subnetpool1", SubnetPoolArgs.builder()
 *             .name("subnetpool_1")
 *             .prefixes(            
 *                 "fdf7:b13d:dead:beef::/64",
 *                 "fd65:86cc:a334:39b7::/64")
 *             .addressScopeId(addressscope1.id())
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
 * Address-scopes can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:networking/addressScope:AddressScope addressscope_1 9cc35860-522a-4d35-974d-51d4b011801e
 * ```
 * 
 */
@ResourceType(type="openstack:networking/addressScope:AddressScope")
public class AddressScope extends com.pulumi.resources.CustomResource {
    /**
     * IP version, either 4 (default) or 6. Changing this
     * creates a new address-scope.
     * 
     */
    @Export(name="ipVersion", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ipVersion;

    /**
     * @return IP version, either 4 (default) or 6. Changing this
     * creates a new address-scope.
     * 
     */
    public Output<Optional<Integer>> ipVersion() {
        return Codegen.optional(this.ipVersion);
    }
    /**
     * The name of the address-scope. Changing this updates the
     * name of the existing address-scope.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the address-scope. Changing this updates the
     * name of the existing address-scope.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The owner of the address-scope. Required if admin
     * wants to create a address-scope for another project. Changing this creates a
     * new address-scope.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The owner of the address-scope. Required if admin
     * wants to create a address-scope for another project. Changing this creates a
     * new address-scope.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron address-scope. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * address-scope.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron address-scope. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * address-scope.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Indicates whether this address-scope is shared across
     * all projects. Changing this updates the shared status of the existing
     * address-scope.
     * 
     */
    @Export(name="shared", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> shared;

    /**
     * @return Indicates whether this address-scope is shared across
     * all projects. Changing this updates the shared status of the existing
     * address-scope.
     * 
     */
    public Output<Boolean> shared() {
        return this.shared;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AddressScope(String name) {
        this(name, AddressScopeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AddressScope(String name, @Nullable AddressScopeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AddressScope(String name, @Nullable AddressScopeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/addressScope:AddressScope", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private AddressScope(String name, Output<String> id, @Nullable AddressScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/addressScope:AddressScope", name, state, makeResourceOptions(options, id));
    }

    private static AddressScopeArgs makeArgs(@Nullable AddressScopeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AddressScopeArgs.Empty : args;
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
    public static AddressScope get(String name, Output<String> id, @Nullable AddressScopeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AddressScope(name, id, state, options);
    }
}
