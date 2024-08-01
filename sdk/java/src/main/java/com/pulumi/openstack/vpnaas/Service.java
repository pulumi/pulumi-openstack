// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.vpnaas.ServiceArgs;
import com.pulumi.openstack.vpnaas.inputs.ServiceState;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron VPN service resource within OpenStack.
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
 * import com.pulumi.openstack.vpnaas.Service;
 * import com.pulumi.openstack.vpnaas.ServiceArgs;
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
 *         var service1 = new Service("service1", ServiceArgs.builder()
 *             .name("my_service")
 *             .routerId("14a75700-fc03-4602-9294-26ee44f366b3")
 *             .adminStateUp("true")
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
 * Services can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:vpnaas/service:Service service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 * 
 */
@ResourceType(type="openstack:vpnaas/service:Service")
public class Service extends com.pulumi.resources.CustomResource {
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing service.
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing service.
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * The human-readable description for the service.
     * Changing this updates the description of the existing service.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The human-readable description for the service.
     * Changing this updates the description of the existing service.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The read-only external (public) IPv4 address that is used for the VPN service.
     * 
     */
    @Export(name="externalV4Ip", refs={String.class}, tree="[0]")
    private Output<String> externalV4Ip;

    /**
     * @return The read-only external (public) IPv4 address that is used for the VPN service.
     * 
     */
    public Output<String> externalV4Ip() {
        return this.externalV4Ip;
    }
    /**
     * The read-only external (public) IPv6 address that is used for the VPN service.
     * 
     */
    @Export(name="externalV6Ip", refs={String.class}, tree="[0]")
    private Output<String> externalV6Ip;

    /**
     * @return The read-only external (public) IPv6 address that is used for the VPN service.
     * 
     */
    public Output<String> externalV6Ip() {
        return this.externalV6Ip;
    }
    /**
     * The name of the service. Changing this updates the name of
     * the existing service.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the service. Changing this updates the name of
     * the existing service.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The ID of the router. Changing this creates a new service.
     * 
     */
    @Export(name="routerId", refs={String.class}, tree="[0]")
    private Output<String> routerId;

    /**
     * @return The ID of the router. Changing this creates a new service.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }
    /**
     * Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * SubnetID is the ID of the subnet. Default is null.
     * 
     */
    @Export(name="subnetId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subnetId;

    /**
     * @return SubnetID is the ID of the subnet. Default is null.
     * 
     */
    public Output<Optional<String>> subnetId() {
        return Codegen.optional(this.subnetId);
    }
    /**
     * The owner of the service. Required if admin wants to
     * create a service for another project. Changing this creates a new service.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the service. Required if admin wants to
     * create a service for another project. Changing this creates a new service.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * Map of additional options.
     * 
     */
    @Export(name="valueSpecs", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional options.
     * 
     */
    public Output<Optional<Map<String,Object>>> valueSpecs() {
        return Codegen.optional(this.valueSpecs);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Service(String name) {
        this(name, ServiceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Service(String name, ServiceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Service(String name, ServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:vpnaas/service:Service", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private Service(String name, Output<String> id, @Nullable ServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:vpnaas/service:Service", name, state, makeResourceOptions(options, id));
    }

    private static ServiceArgs makeArgs(ServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceArgs.Empty : args;
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
    public static Service get(String name, Output<String> id, @Nullable ServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Service(name, id, state, options);
    }
}
