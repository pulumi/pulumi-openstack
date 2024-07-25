// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.BgpvpnNetworkAssociateV2Args;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.inputs.BgpvpnNetworkAssociateV2State;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a V2 BGP VPN network association resource within OpenStack.
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
 * import com.pulumi.openstack.BgpvpnNetworkAssociateV2;
 * import com.pulumi.openstack.BgpvpnNetworkAssociateV2Args;
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
 *         var association1 = new BgpvpnNetworkAssociateV2("association1", BgpvpnNetworkAssociateV2Args.builder()
 *             .bgpvpnId("e7189337-5684-46ee-bcb1-44f1a57066c9")
 *             .networkId("de83d56c-4d2f-44f7-ac24-af393252204f")
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
 * BGP VPN network associations can be imported using the BGP VPN ID and BGP VPN
 * 
 * network association ID separated by a slash, e.g.:
 * 
 * hcl
 * 
 * ```sh
 * $ pulumi import openstack:index/bgpvpnNetworkAssociateV2:BgpvpnNetworkAssociateV2 association_1 2145aaa9-edaa-44fb-9815-e47a96677a72/67bb952a-f9d1-4fc8-ae84-082253a879d4
 * ```
 * 
 */
@ResourceType(type="openstack:index/bgpvpnNetworkAssociateV2:BgpvpnNetworkAssociateV2")
public class BgpvpnNetworkAssociateV2 extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the BGP VPN to which the network will be
     * associated. Changing this creates a new BGP VPN network association
     * 
     */
    @Export(name="bgpvpnId", refs={String.class}, tree="[0]")
    private Output<String> bgpvpnId;

    /**
     * @return The ID of the BGP VPN to which the network will be
     * associated. Changing this creates a new BGP VPN network association
     * 
     */
    public Output<String> bgpvpnId() {
        return this.bgpvpnId;
    }
    /**
     * The ID of the network to be associated with the BGP
     * VPN. Changing this creates a new BGP VPN network association.
     * 
     */
    @Export(name="networkId", refs={String.class}, tree="[0]")
    private Output<String> networkId;

    /**
     * @return The ID of the network to be associated with the BGP
     * VPN. Changing this creates a new BGP VPN network association.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }
    /**
     * The ID of the project that owns the BGP VPN network
     * association. Only administrative and users with `advsvc` role can specify a
     * project ID other than their own. Changing this creates a new BGP VPN network
     * association.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project that owns the BGP VPN network
     * association. Only administrative and users with `advsvc` role can specify a
     * project ID other than their own. Changing this creates a new BGP VPN network
     * association.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a BGP VPN network association. If
     * omitted, the `region` argument of the provider is used. Changing this creates
     * a new BGP VPN network association.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a BGP VPN network association. If
     * omitted, the `region` argument of the provider is used. Changing this creates
     * a new BGP VPN network association.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BgpvpnNetworkAssociateV2(String name) {
        this(name, BgpvpnNetworkAssociateV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BgpvpnNetworkAssociateV2(String name, BgpvpnNetworkAssociateV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BgpvpnNetworkAssociateV2(String name, BgpvpnNetworkAssociateV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:index/bgpvpnNetworkAssociateV2:BgpvpnNetworkAssociateV2", name, args == null ? BgpvpnNetworkAssociateV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BgpvpnNetworkAssociateV2(String name, Output<String> id, @Nullable BgpvpnNetworkAssociateV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:index/bgpvpnNetworkAssociateV2:BgpvpnNetworkAssociateV2", name, state, makeResourceOptions(options, id));
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
    public static BgpvpnNetworkAssociateV2 get(String name, Output<String> id, @Nullable BgpvpnNetworkAssociateV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BgpvpnNetworkAssociateV2(name, id, state, options);
    }
}
