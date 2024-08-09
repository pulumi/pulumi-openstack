// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.dns;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.dns.TransferRequestArgs;
import com.pulumi.openstack.dns.inputs.TransferRequestState;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a DNS zone transfer request in the OpenStack DNS Service.
 * 
 * ## Example Usage
 * 
 * ### Automatically detect the correct network
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.dns.Zone;
 * import com.pulumi.openstack.dns.ZoneArgs;
 * import com.pulumi.openstack.dns.TransferRequest;
 * import com.pulumi.openstack.dns.TransferRequestArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var exampleZone = new Zone("exampleZone", ZoneArgs.builder()
 *             .name("example.com.")
 *             .email("jdoe}{@literal @}{@code example.com")
 *             .description("An example zone")
 *             .ttl(3000)
 *             .type("PRIMARY")
 *             .build());
 * 
 *         var request1 = new TransferRequest("request1", TransferRequestArgs.builder()
 *             .zoneId(exampleZone.id())
 *             .description("a transfer request")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource can be imported by specifying the transferRequest ID:
 * 
 * ```sh
 * $ pulumi import openstack:dns/transferRequest:TransferRequest request_1 request_id
 * ```
 * 
 */
@ResourceType(type="openstack:dns/transferRequest:TransferRequest")
public class TransferRequest extends com.pulumi.resources.CustomResource {
    /**
     * A description of the zone tranfer request.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the zone tranfer request.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     * 
     */
    @Export(name="disableStatusCheck", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableStatusCheck;

    /**
     * @return Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     * 
     */
    public Output<Optional<Boolean>> disableStatusCheck() {
        return Codegen.optional(this.disableStatusCheck);
    }
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    public Output<String> key() {
        return this.key;
    }
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The target Project ID to transfer to.
     * 
     */
    @Export(name="targetProjectId", refs={String.class}, tree="[0]")
    private Output<String> targetProjectId;

    /**
     * @return The target Project ID to transfer to.
     * 
     */
    public Output<String> targetProjectId() {
        return this.targetProjectId;
    }
    /**
     * Map of additional options. Changing this creates a
     * new transfer request.
     * 
     */
    @Export(name="valueSpecs", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional options. Changing this creates a
     * new transfer request.
     * 
     */
    public Output<Optional<Map<String,Object>>> valueSpecs() {
        return Codegen.optional(this.valueSpecs);
    }
    /**
     * The ID of the zone for which to create the transfer
     * request.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The ID of the zone for which to create the transfer
     * request.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TransferRequest(java.lang.String name) {
        this(name, TransferRequestArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TransferRequest(java.lang.String name, TransferRequestArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TransferRequest(java.lang.String name, TransferRequestArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:dns/transferRequest:TransferRequest", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TransferRequest(java.lang.String name, Output<java.lang.String> id, @Nullable TransferRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:dns/transferRequest:TransferRequest", name, state, makeResourceOptions(options, id), false);
    }

    private static TransferRequestArgs makeArgs(TransferRequestArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TransferRequestArgs.Empty : args;
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
    public static TransferRequest get(java.lang.String name, Output<java.lang.String> id, @Nullable TransferRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TransferRequest(name, id, state, options);
    }
}
