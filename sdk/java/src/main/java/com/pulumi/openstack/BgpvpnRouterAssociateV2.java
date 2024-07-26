// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.BgpvpnRouterAssociateV2Args;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.inputs.BgpvpnRouterAssociateV2State;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a V2 BGP VPN router association resource within OpenStack.
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
 * import com.pulumi.openstack.BgpvpnRouterAssociateV2;
 * import com.pulumi.openstack.BgpvpnRouterAssociateV2Args;
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
 *         var association1 = new BgpvpnRouterAssociateV2("association1", BgpvpnRouterAssociateV2Args.builder()
 *             .bgpvpnId("d57d39e1-dc63-44fd-8cbd-a4e1488100c5")
 *             .routerId("423fa80f-e0d7-4d02-a9a5-8b8c05812bf6")
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
 * BGP VPN router associations can be imported using the BGP VPN ID and BGP VPN
 * 
 * router association ID separated by a slash, e.g.:
 * 
 * hcl
 * 
 * ```sh
 * $ pulumi import openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2 association_1 e26d509e-fc2d-4fb5-8562-619911a9a6bc/3cc9df2d-80db-4536-8ba6-295d1d0f723f
 * ```
 * 
 */
@ResourceType(type="openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2")
public class BgpvpnRouterAssociateV2 extends com.pulumi.resources.CustomResource {
    /**
     * A boolean flag indicating whether extra
     * routes should be advertised. Defaults to true.
     * 
     */
    @Export(name="advertiseExtraRoutes", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> advertiseExtraRoutes;

    /**
     * @return A boolean flag indicating whether extra
     * routes should be advertised. Defaults to true.
     * 
     */
    public Output<Boolean> advertiseExtraRoutes() {
        return this.advertiseExtraRoutes;
    }
    /**
     * The ID of the BGP VPN to which the router will be
     * associated. Changing this creates a new BGP VPN router association.
     * 
     */
    @Export(name="bgpvpnId", refs={String.class}, tree="[0]")
    private Output<String> bgpvpnId;

    /**
     * @return The ID of the BGP VPN to which the router will be
     * associated. Changing this creates a new BGP VPN router association.
     * 
     */
    public Output<String> bgpvpnId() {
        return this.bgpvpnId;
    }
    /**
     * The ID of the project that owns the BGP VPN router
     * association. Only administrative and users with `advsvc` role can specify a
     * project ID other than their own. Changing this creates a new BGP VPN router
     * association.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project that owns the BGP VPN router
     * association. Only administrative and users with `advsvc` role can specify a
     * project ID other than their own. Changing this creates a new BGP VPN router
     * association.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a BGP VPN router association. If
     * omitted, the `region` argument of the provider is used. Changing this creates
     * a new BGP VPN router association.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a BGP VPN router association. If
     * omitted, the `region` argument of the provider is used. Changing this creates
     * a new BGP VPN router association.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The ID of the router to be associated with the BGP
     * VPN. Changing this creates a new BGP VPN router association.
     * 
     */
    @Export(name="routerId", refs={String.class}, tree="[0]")
    private Output<String> routerId;

    /**
     * @return The ID of the router to be associated with the BGP
     * VPN. Changing this creates a new BGP VPN router association.
     * 
     */
    public Output<String> routerId() {
        return this.routerId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BgpvpnRouterAssociateV2(String name) {
        this(name, BgpvpnRouterAssociateV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BgpvpnRouterAssociateV2(String name, BgpvpnRouterAssociateV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BgpvpnRouterAssociateV2(String name, BgpvpnRouterAssociateV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2", name, args == null ? BgpvpnRouterAssociateV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BgpvpnRouterAssociateV2(String name, Output<String> id, @Nullable BgpvpnRouterAssociateV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2", name, state, makeResourceOptions(options, id));
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
    public static BgpvpnRouterAssociateV2 get(String name, Output<String> id, @Nullable BgpvpnRouterAssociateV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BgpvpnRouterAssociateV2(name, id, state, options);
    }
}