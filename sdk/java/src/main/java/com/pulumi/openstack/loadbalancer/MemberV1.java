// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.MemberV1Args;
import com.pulumi.openstack.loadbalancer.inputs.MemberV1State;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 load balancer member resource within OpenStack.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.loadbalancer.MemberV1;
 * import com.pulumi.openstack.loadbalancer.MemberV1Args;
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
 *         var member1 = new MemberV1(&#34;member1&#34;, MemberV1Args.builder()        
 *             .address(&#34;192.168.0.10&#34;)
 *             .poolId(&#34;d9415786-5f1a-428b-b35f-2f1523e146d2&#34;)
 *             .port(80)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Load Balancer Members can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:loadbalancer/memberV1:MemberV1 member_1 a7498676-4fe4-4243-a864-2eaaf18c73df
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/memberV1:MemberV1")
public class MemberV1 extends com.pulumi.resources.CustomResource {
    /**
     * The IP address of the member. Changing this creates a
     * new member.
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return The IP address of the member. Changing this creates a
     * new member.
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * The administrative state of the member.
     * Acceptable values are &#39;true&#39; and &#39;false&#39;. Changing this value updates the
     * state of the existing member.
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> adminStateUp;

    /**
     * @return The administrative state of the member.
     * Acceptable values are &#39;true&#39; and &#39;false&#39;. Changing this value updates the
     * state of the existing member.
     * 
     */
    public Output<Boolean> adminStateUp() {
        return this.adminStateUp;
    }
    /**
     * The ID of the LB pool. Changing this creates a new
     * member.
     * 
     */
    @Export(name="poolId", refs={String.class}, tree="[0]")
    private Output<String> poolId;

    /**
     * @return The ID of the LB pool. Changing this creates a new
     * member.
     * 
     */
    public Output<String> poolId() {
        return this.poolId;
    }
    /**
     * An integer representing the port on which the member is
     * hosted. Changing this creates a new member.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return An integer representing the port on which the member is
     * hosted. Changing this creates a new member.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The owner of the member. Required if admin wants to
     * create a member for another tenant. Changing this creates a new member.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tenantId;

    /**
     * @return The owner of the member. Required if admin wants to
     * create a member for another tenant. Changing this creates a new member.
     * 
     */
    public Output<Optional<String>> tenantId() {
        return Codegen.optional(this.tenantId);
    }
    @Export(name="weight", refs={Integer.class}, tree="[0]")
    private Output<Integer> weight;

    public Output<Integer> weight() {
        return this.weight;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MemberV1(String name) {
        this(name, MemberV1Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MemberV1(String name, MemberV1Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MemberV1(String name, MemberV1Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/memberV1:MemberV1", name, args == null ? MemberV1Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MemberV1(String name, Output<String> id, @Nullable MemberV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/memberV1:MemberV1", name, state, makeResourceOptions(options, id));
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
    public static MemberV1 get(String name, Output<String> id, @Nullable MemberV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MemberV1(name, id, state, options);
    }
}
