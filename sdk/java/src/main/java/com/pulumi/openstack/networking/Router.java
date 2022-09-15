// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.RouterArgs;
import com.pulumi.openstack.networking.inputs.RouterState;
import com.pulumi.openstack.networking.outputs.RouterExternalFixedIp;
import com.pulumi.openstack.networking.outputs.RouterVendorOptions;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 router resource within OpenStack.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Router;
 * import com.pulumi.openstack.networking.RouterArgs;
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
 *         var router1 = new Router(&#34;router1&#34;, RouterArgs.builder()        
 *             .adminStateUp(true)
 *             .externalNetworkId(&#34;f67f0d72-0ddf-11e4-9d95-e1f29f417e2f&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Routers can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:networking/router:Router router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2
 * ```
 * 
 */
@ResourceType(type="openstack:networking/router:Router")
public class Router extends com.pulumi.resources.CustomResource {
    /**
     * Administrative up/down status for the router
     * (must be &#34;true&#34; or &#34;false&#34; if provided). Changing this updates the
     * `admin_state_up` of an existing router.
     * 
     */
    @Export(name="adminStateUp", type=Boolean.class, parameters={})
    private Output<Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the router
     * (must be &#34;true&#34; or &#34;false&#34; if provided). Changing this updates the
     * `admin_state_up` of an existing router.
     * 
     */
    public Output<Boolean> adminStateUp() {
        return this.adminStateUp;
    }
    /**
     * The collection of tags assigned on the router, which have been
     * explicitly and implicitly added.
     * 
     */
    @Export(name="allTags", type=List.class, parameters={String.class})
    private Output<List<String>> allTags;

    /**
     * @return The collection of tags assigned on the router, which have been
     * explicitly and implicitly added.
     * 
     */
    public Output<List<String>> allTags() {
        return this.allTags;
    }
    /**
     * An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new router.
     * 
     */
    @Export(name="availabilityZoneHints", type=List.class, parameters={String.class})
    private Output<List<String>> availabilityZoneHints;

    /**
     * @return An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new router.
     * 
     */
    public Output<List<String>> availabilityZoneHints() {
        return this.availabilityZoneHints;
    }
    /**
     * Human-readable description for the router.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-readable description for the router.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Indicates whether or not to create a
     * distributed router. The default policy setting in Neutron restricts
     * usage of this property to administrative users only.
     * 
     */
    @Export(name="distributed", type=Boolean.class, parameters={})
    private Output<Boolean> distributed;

    /**
     * @return Indicates whether or not to create a
     * distributed router. The default policy setting in Neutron restricts
     * usage of this property to administrative users only.
     * 
     */
    public Output<Boolean> distributed() {
        return this.distributed;
    }
    /**
     * Enable Source NAT for the router. Valid values are
     * &#34;true&#34; or &#34;false&#34;. An `external_network_id` has to be set in order to
     * set this property. Changing this updates the `enable_snat` of the router.
     * Setting this value **requires** an **ext-gw-mode** extension to be enabled
     * in OpenStack Neutron.
     * 
     */
    @Export(name="enableSnat", type=Boolean.class, parameters={})
    private Output<Boolean> enableSnat;

    /**
     * @return Enable Source NAT for the router. Valid values are
     * &#34;true&#34; or &#34;false&#34;. An `external_network_id` has to be set in order to
     * set this property. Changing this updates the `enable_snat` of the router.
     * Setting this value **requires** an **ext-gw-mode** extension to be enabled
     * in OpenStack Neutron.
     * 
     */
    public Output<Boolean> enableSnat() {
        return this.enableSnat;
    }
    /**
     * An external fixed IP for the router. This
     * can be repeated. The structure is described below. An `external_network_id`
     * has to be set in order to set this property. Changing this updates the
     * external fixed IPs of the router.
     * 
     */
    @Export(name="externalFixedIps", type=List.class, parameters={RouterExternalFixedIp.class})
    private Output<List<RouterExternalFixedIp>> externalFixedIps;

    /**
     * @return An external fixed IP for the router. This
     * can be repeated. The structure is described below. An `external_network_id`
     * has to be set in order to set this property. Changing this updates the
     * external fixed IPs of the router.
     * 
     */
    public Output<List<RouterExternalFixedIp>> externalFixedIps() {
        return this.externalFixedIps;
    }
    /**
     * The
     * network UUID of an external gateway for the router. A router with an
     * external gateway is required if any compute instances or load balancers
     * will be using floating IPs. Changing this updates the external gateway
     * of an existing router.
     * 
     * @deprecated
     * use external_network_id instead
     * 
     */
    @Deprecated /* use external_network_id instead */
    @Export(name="externalGateway", type=String.class, parameters={})
    private Output<String> externalGateway;

    /**
     * @return The
     * network UUID of an external gateway for the router. A router with an
     * external gateway is required if any compute instances or load balancers
     * will be using floating IPs. Changing this updates the external gateway
     * of an existing router.
     * 
     */
    public Output<String> externalGateway() {
        return this.externalGateway;
    }
    /**
     * The network UUID of an external gateway
     * for the router. A router with an external gateway is required if any
     * compute instances or load balancers will be using floating IPs. Changing
     * this updates the external gateway of the router.
     * 
     */
    @Export(name="externalNetworkId", type=String.class, parameters={})
    private Output<String> externalNetworkId;

    /**
     * @return The network UUID of an external gateway
     * for the router. A router with an external gateway is required if any
     * compute instances or load balancers will be using floating IPs. Changing
     * this updates the external gateway of the router.
     * 
     */
    public Output<String> externalNetworkId() {
        return this.externalNetworkId;
    }
    /**
     * A list of external subnet IDs to try over
     * each to obtain a fixed IP for the router. If a subnet ID in a list has
     * exhausted floating IP pool, the next subnet ID will be tried. This argument is
     * used only during the router creation and allows to set only one external fixed
     * IP. Conflicts with an `external_fixed_ip` argument.
     * 
     */
    @Export(name="externalSubnetIds", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> externalSubnetIds;

    /**
     * @return A list of external subnet IDs to try over
     * each to obtain a fixed IP for the router. If a subnet ID in a list has
     * exhausted floating IP pool, the next subnet ID will be tried. This argument is
     * used only during the router creation and allows to set only one external fixed
     * IP. Conflicts with an `external_fixed_ip` argument.
     * 
     */
    public Output<Optional<List<String>>> externalSubnetIds() {
        return Codegen.optional(this.externalSubnetIds);
    }
    /**
     * A unique name for the router. Changing this
     * updates the `name` of an existing router.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A unique name for the router. Changing this
     * updates the `name` of an existing router.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A set of string tags for the router.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A set of string tags for the router.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a router for another tenant. Changing this creates a new router.
     * 
     */
    @Export(name="tenantId", type=String.class, parameters={})
    private Output<String> tenantId;

    /**
     * @return The owner of the floating IP. Required if admin wants
     * to create a router for another tenant. Changing this creates a new router.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * Map of additional driver-specific options.
     * 
     */
    @Export(name="valueSpecs", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional driver-specific options.
     * 
     */
    public Output<Optional<Map<String,Object>>> valueSpecs() {
        return Codegen.optional(this.valueSpecs);
    }
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     * 
     */
    @Export(name="vendorOptions", type=RouterVendorOptions.class, parameters={})
    private Output</* @Nullable */ RouterVendorOptions> vendorOptions;

    /**
     * @return Map of additional vendor-specific options.
     * Supported options are described below.
     * 
     */
    public Output<Optional<RouterVendorOptions>> vendorOptions() {
        return Codegen.optional(this.vendorOptions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Router(String name) {
        this(name, RouterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Router(String name, @Nullable RouterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Router(String name, @Nullable RouterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/router:Router", name, args == null ? RouterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Router(String name, Output<String> id, @Nullable RouterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/router:Router", name, state, makeResourceOptions(options, id));
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
    public static Router get(String name, Output<String> id, @Nullable RouterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Router(name, id, state, options);
    }
}