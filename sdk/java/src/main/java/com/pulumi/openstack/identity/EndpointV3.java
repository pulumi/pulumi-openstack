// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.identity.EndpointV3Args;
import com.pulumi.openstack.identity.inputs.EndpointV3State;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V3 Endpoint resource within OpenStack Keystone.
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
 * import com.pulumi.openstack.identity.ServiceV3;
 * import com.pulumi.openstack.identity.ServiceV3Args;
 * import com.pulumi.openstack.identity.EndpointV3;
 * import com.pulumi.openstack.identity.EndpointV3Args;
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
 *         var service1 = new ServiceV3("service1", ServiceV3Args.builder()
 *             .name("my-service")
 *             .type("my-service-type")
 *             .build());
 * 
 *         var endpoint1 = new EndpointV3("endpoint1", EndpointV3Args.builder()
 *             .name("my-endpoint")
 *             .serviceId(service1.id())
 *             .endpointRegion(service1.region())
 *             .url("http://my-endpoint")
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
 * Endpoints can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:identity/endpointV3:EndpointV3 endpoint_1 5392472b-106a-4845-90c6-7c8445f18770
 * ```
 * 
 */
@ResourceType(type="openstack:identity/endpointV3:EndpointV3")
public class EndpointV3 extends com.pulumi.resources.CustomResource {
    /**
     * The endpoint region. The `region` and
     * `endpoint_region` can be different.
     * 
     */
    @Export(name="endpointRegion", refs={String.class}, tree="[0]")
    private Output<String> endpointRegion;

    /**
     * @return The endpoint region. The `region` and
     * `endpoint_region` can be different.
     * 
     */
    public Output<String> endpointRegion() {
        return this.endpointRegion;
    }
    /**
     * The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     * 
     */
    @Export(name="interface", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> interface_;

    /**
     * @return The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     * 
     */
    public Output<Optional<String>> interface_() {
        return Codegen.optional(this.interface_);
    }
    /**
     * The endpoint name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The endpoint name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The endpoint service ID.
     * 
     */
    @Export(name="serviceId", refs={String.class}, tree="[0]")
    private Output<String> serviceId;

    /**
     * @return The endpoint service ID.
     * 
     */
    public Output<String> serviceId() {
        return this.serviceId;
    }
    /**
     * The service name of the endpoint.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The service name of the endpoint.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * The service type of the endpoint.
     * 
     */
    @Export(name="serviceType", refs={String.class}, tree="[0]")
    private Output<String> serviceType;

    /**
     * @return The service type of the endpoint.
     * 
     */
    public Output<String> serviceType() {
        return this.serviceType;
    }
    /**
     * The endpoint url.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The endpoint url.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EndpointV3(java.lang.String name) {
        this(name, EndpointV3Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EndpointV3(java.lang.String name, EndpointV3Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EndpointV3(java.lang.String name, EndpointV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/endpointV3:EndpointV3", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EndpointV3(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/endpointV3:EndpointV3", name, state, makeResourceOptions(options, id), false);
    }

    private static EndpointV3Args makeArgs(EndpointV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EndpointV3Args.Empty : args;
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
    public static EndpointV3 get(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EndpointV3(name, id, state, options);
    }
}
