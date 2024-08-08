// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.MemberArgs;
import com.pulumi.openstack.loadbalancer.inputs.MemberState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 member resource within OpenStack.
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
 * import com.pulumi.openstack.loadbalancer.Member;
 * import com.pulumi.openstack.loadbalancer.MemberArgs;
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
 *         var member1 = new Member("member1", MemberArgs.builder()
 *             .poolId("935685fb-a896-40f9-9ff4-ae531a3a00fe")
 *             .address("192.168.199.23")
 *             .protocolPort(8080)
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
 * Load Balancer Pool Member can be imported using the Pool ID and Member ID
 * separated by a slash, e.g.:
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/member:Member member_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5/9563b79c-8460-47da-8a95-2711b746510f
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/member:Member")
public class Member extends com.pulumi.resources.CustomResource {
    /**
     * The IP address of the member to receive traffic from
     * the load balancer. Changing this creates a new member.
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return The IP address of the member to receive traffic from
     * the load balancer. Changing this creates a new member.
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * Boolean that indicates whether that member works as a backup or not. Available
     * only for Octavia &gt;= 2.1.
     * 
     */
    @Export(name="backup", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> backup;

    /**
     * @return Boolean that indicates whether that member works as a backup or not. Available
     * only for Octavia &gt;= 2.1.
     * 
     */
    public Output<Optional<Boolean>> backup() {
        return Codegen.optional(this.backup);
    }
    /**
     * An alternate IP address used for health monitoring a backend member.
     * Available only for Octavia
     * 
     */
    @Export(name="monitorAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> monitorAddress;

    /**
     * @return An alternate IP address used for health monitoring a backend member.
     * Available only for Octavia
     * 
     */
    public Output<Optional<String>> monitorAddress() {
        return Codegen.optional(this.monitorAddress);
    }
    /**
     * An alternate protocol port used for health monitoring a backend member.
     * Available only for Octavia
     * 
     */
    @Export(name="monitorPort", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> monitorPort;

    /**
     * @return An alternate protocol port used for health monitoring a backend member.
     * Available only for Octavia
     * 
     */
    public Output<Optional<Integer>> monitorPort() {
        return Codegen.optional(this.monitorPort);
    }
    /**
     * Human-readable name for the member.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Human-readable name for the member.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The id of the pool that this member will be assigned
     * to. Changing this creates a new member.
     * 
     */
    @Export(name="poolId", refs={String.class}, tree="[0]")
    private Output<String> poolId;

    /**
     * @return The id of the pool that this member will be assigned
     * to. Changing this creates a new member.
     * 
     */
    public Output<String> poolId() {
        return this.poolId;
    }
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new member.
     * 
     */
    @Export(name="protocolPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> protocolPort;

    /**
     * @return The port on which to listen for client traffic.
     * Changing this creates a new member.
     * 
     */
    public Output<Integer> protocolPort() {
        return this.protocolPort;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a member. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new member.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a member. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new member.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The subnet in which to access the member. Changing
     * this creates a new member.
     * 
     */
    @Export(name="subnetId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subnetId;

    /**
     * @return The subnet in which to access the member. Changing
     * this creates a new member.
     * 
     */
    public Output<Optional<String>> subnetId() {
        return Codegen.optional(this.subnetId);
    }
    /**
     * A list of simple strings assigned to the member.
     * Available only for Octavia &gt;= 2.5.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of simple strings assigned to the member.
     * Available only for Octavia &gt;= 2.5.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Required for admins. The UUID of the tenant who owns
     * the member.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new member.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the member.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new member.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * A positive integer value that indicates the relative
     * portion of traffic that this member should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     * 
     */
    @Export(name="weight", refs={Integer.class}, tree="[0]")
    private Output<Integer> weight;

    /**
     * @return A positive integer value that indicates the relative
     * portion of traffic that this member should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     * 
     */
    public Output<Integer> weight() {
        return this.weight;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Member(java.lang.String name) {
        this(name, MemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Member(java.lang.String name, MemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Member(java.lang.String name, MemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/member:Member", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Member(java.lang.String name, Output<java.lang.String> id, @Nullable MemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/member:Member", name, state, makeResourceOptions(options, id), false);
    }

    private static MemberArgs makeArgs(MemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MemberArgs.Empty : args;
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
    public static Member get(java.lang.String name, Output<java.lang.String> id, @Nullable MemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Member(name, id, state, options);
    }
}
