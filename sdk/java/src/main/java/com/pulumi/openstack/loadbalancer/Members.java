// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.MembersArgs;
import com.pulumi.openstack.loadbalancer.inputs.MembersState;
import com.pulumi.openstack.loadbalancer.outputs.MembersMember;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 members resource within OpenStack (batch members update).
 * 
 * &gt; **Note:** This resource has attributes that depend on octavia minor versions.
 * Please ensure your Openstack cloud supports the required minor version.
 * 
 * &gt; **Note:** This resource works only within Octavia API. For
 * legacy Neutron LBaaS v2 extension please use
 * openstack.loadbalancer.Member resource.
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
 * import com.pulumi.openstack.loadbalancer.Members;
 * import com.pulumi.openstack.loadbalancer.MembersArgs;
 * import com.pulumi.openstack.loadbalancer.inputs.MembersMemberArgs;
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
 *         var members1 = new Members("members1", MembersArgs.builder()
 *             .poolId("935685fb-a896-40f9-9ff4-ae531a3a00fe")
 *             .members(            
 *                 MembersMemberArgs.builder()
 *                     .address("192.168.199.23")
 *                     .protocolPort(8080)
 *                     .build(),
 *                 MembersMemberArgs.builder()
 *                     .address("192.168.199.24")
 *                     .protocolPort(8080)
 *                     .build())
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
 * Load Balancer Pool Members can be imported using the Pool ID, e.g.:
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/members:Members members_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/members:Members")
public class Members extends com.pulumi.resources.CustomResource {
    /**
     * A set of dictionaries containing member parameters. The
     * structure is described below.
     * 
     */
    @Export(name="members", refs={List.class,MembersMember.class}, tree="[0,1]")
    private Output</* @Nullable */ List<MembersMember>> members;

    /**
     * @return A set of dictionaries containing member parameters. The
     * structure is described below.
     * 
     */
    public Output<Optional<List<MembersMember>>> members() {
        return Codegen.optional(this.members);
    }
    /**
     * The id of the pool that members will be assigned to.
     * Changing this creates a new members resource.
     * 
     */
    @Export(name="poolId", refs={String.class}, tree="[0]")
    private Output<String> poolId;

    /**
     * @return The id of the pool that members will be assigned to.
     * Changing this creates a new members resource.
     * 
     */
    public Output<String> poolId() {
        return this.poolId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create pool members. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * members resource.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create pool members. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * members resource.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Members(String name) {
        this(name, MembersArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Members(String name, MembersArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Members(String name, MembersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/members:Members", name, args == null ? MembersArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Members(String name, Output<String> id, @Nullable MembersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/members:Members", name, state, makeResourceOptions(options, id));
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
    public static Members get(String name, Output<String> id, @Nullable MembersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Members(name, id, state, options);
    }
}
