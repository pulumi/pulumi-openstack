// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.LoadBalancerArgs;
import com.pulumi.openstack.loadbalancer.inputs.LoadBalancerState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 loadbalancer resource within OpenStack.
 * 
 * &gt; **Note:** This resource has attributes that depend on octavia minor versions.
 * Please ensure your Openstack cloud supports the required minor version.
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
 * import com.pulumi.openstack.LbLoadbalancerV2;
 * import com.pulumi.openstack.LbLoadbalancerV2Args;
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
 *         var lb1 = new LbLoadbalancerV2("lb1", LbLoadbalancerV2Args.builder()
 *             .vipSubnetId("d9415786-5f1a-428b-b35f-2f1523e146d2")
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
 * Load Balancer can be imported using the Load Balancer ID, e.g.:
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/loadBalancer:LoadBalancer loadbalancer_1 19bcfdc7-c521-4a7e-9459-6750bd16df76
 * ```
 * 
 * @deprecated
 * openstack.loadbalancer/loadbalancer.LoadBalancer has been deprecated in favor of openstack.index/lbloadbalancerv2.LbLoadbalancerV2
 * 
 */
@Deprecated /* openstack.loadbalancer/loadbalancer.LoadBalancer has been deprecated in favor of openstack.index/lbloadbalancerv2.LbLoadbalancerV2 */
@ResourceType(type="openstack:loadbalancer/loadBalancer:LoadBalancer")
public class LoadBalancer extends com.pulumi.resources.CustomResource {
    /**
     * The administrative state of the Loadbalancer.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    @Export(name="adminStateUp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminStateUp;

    /**
     * @return The administrative state of the Loadbalancer.
     * A valid value is true (UP) or false (DOWN).
     * 
     */
    public Output<Optional<Boolean>> adminStateUp() {
        return Codegen.optional(this.adminStateUp);
    }
    /**
     * The availability zone of the Loadbalancer.
     * Changing this creates a new loadbalancer. Available only for Octavia
     * **minor version 2.14 or later**.
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> availabilityZone;

    /**
     * @return The availability zone of the Loadbalancer.
     * Changing this creates a new loadbalancer. Available only for Octavia
     * **minor version 2.14 or later**.
     * 
     */
    public Output<Optional<String>> availabilityZone() {
        return Codegen.optional(this.availabilityZone);
    }
    /**
     * Human-readable description for the Loadbalancer.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Human-readable description for the Loadbalancer.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The UUID of a flavor. Changing this creates a new
     * loadbalancer.
     * 
     */
    @Export(name="flavorId", refs={String.class}, tree="[0]")
    private Output<String> flavorId;

    /**
     * @return The UUID of a flavor. Changing this creates a new
     * loadbalancer.
     * 
     */
    public Output<String> flavorId() {
        return this.flavorId;
    }
    /**
     * The name of the provider. Changing this
     * creates a new loadbalancer.
     * 
     */
    @Export(name="loadbalancerProvider", refs={String.class}, tree="[0]")
    private Output<String> loadbalancerProvider;

    /**
     * @return The name of the provider. Changing this
     * creates a new loadbalancer.
     * 
     */
    public Output<String> loadbalancerProvider() {
        return this.loadbalancerProvider;
    }
    /**
     * Human-readable name for the Loadbalancer. Does not have
     * to be unique.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Human-readable name for the Loadbalancer. Does not have
     * to be unique.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * A list of security group IDs to apply to the
     * loadbalancer. The security groups must be specified by ID and not name (as
     * opposed to how they are configured with the Compute Instance).
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroupIds;

    /**
     * @return A list of security group IDs to apply to the
     * loadbalancer. The security groups must be specified by ID and not name (as
     * opposed to how they are configured with the Compute Instance).
     * 
     */
    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * A list of simple strings assigned to the loadbalancer.
     * Available only for Octavia **minor version 2.5 or later**.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of simple strings assigned to the loadbalancer.
     * Available only for Octavia **minor version 2.5 or later**.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Loadbalancer.  Only administrative users can specify a tenant UUID
     * other than their own.  Changing this creates a new loadbalancer.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return Required for admins. The UUID of the tenant who owns
     * the Loadbalancer.  Only administrative users can specify a tenant UUID
     * other than their own.  Changing this creates a new loadbalancer.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * The ip address of the load balancer.
     * Changing this creates a new loadbalancer.
     * 
     */
    @Export(name="vipAddress", refs={String.class}, tree="[0]")
    private Output<String> vipAddress;

    /**
     * @return The ip address of the load balancer.
     * Changing this creates a new loadbalancer.
     * 
     */
    public Output<String> vipAddress() {
        return this.vipAddress;
    }
    /**
     * The network on which to allocate the
     * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is available only for Octavia.
     * 
     */
    @Export(name="vipNetworkId", refs={String.class}, tree="[0]")
    private Output<String> vipNetworkId;

    /**
     * @return The network on which to allocate the
     * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is available only for Octavia.
     * 
     */
    public Output<String> vipNetworkId() {
        return this.vipNetworkId;
    }
    /**
     * The port UUID that the loadbalancer will use.
     * Changing this creates a new loadbalancer. It is available only for Octavia.
     * 
     */
    @Export(name="vipPortId", refs={String.class}, tree="[0]")
    private Output<String> vipPortId;

    /**
     * @return The port UUID that the loadbalancer will use.
     * Changing this creates a new loadbalancer. It is available only for Octavia.
     * 
     */
    public Output<String> vipPortId() {
        return this.vipPortId;
    }
    /**
     * The subnet on which to allocate the
     * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is required to Neutron LBaaS but optional for Octavia.
     * 
     */
    @Export(name="vipSubnetId", refs={String.class}, tree="[0]")
    private Output<String> vipSubnetId;

    /**
     * @return The subnet on which to allocate the
     * Loadbalancer&#39;s address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is required to Neutron LBaaS but optional for Octavia.
     * 
     */
    public Output<String> vipSubnetId() {
        return this.vipSubnetId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancer(String name) {
        this(name, LoadBalancerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancer(String name, @Nullable LoadBalancerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancer(String name, @Nullable LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/loadBalancer:LoadBalancer", name, args == null ? LoadBalancerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoadBalancer(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/loadBalancer:LoadBalancer", name, state, makeResourceOptions(options, id));
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
    public static LoadBalancer get(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancer(name, id, state, options);
    }
}
