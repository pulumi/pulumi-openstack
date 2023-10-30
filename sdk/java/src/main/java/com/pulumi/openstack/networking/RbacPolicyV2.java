// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.RbacPolicyV2Args;
import com.pulumi.openstack.networking.inputs.RbacPolicyV2State;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The RBAC policy resource contains functionality for working with Neutron RBAC
 * Policies. Role-Based Access Control (RBAC) policy framework enables both
 * operators and users to grant access to resources for specific projects.
 * 
 * Sharing an object with a specific project is accomplished by creating a
 * policy entry that permits the target project the `access_as_shared` action
 * on that object.
 * 
 * To make a network available as an external network for specific projects
 * rather than all projects, use the `access_as_external` action.
 * If a network is marked as external during creation, it now implicitly creates
 * a wildcard RBAC policy granting everyone access to preserve previous behavior
 * before this feature was added.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.RbacPolicyV2;
 * import com.pulumi.openstack.networking.RbacPolicyV2Args;
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
 *         var network1 = new Network(&#34;network1&#34;, NetworkArgs.builder()        
 *             .adminStateUp(&#34;true&#34;)
 *             .build());
 * 
 *         var rbacPolicy1 = new RbacPolicyV2(&#34;rbacPolicy1&#34;, RbacPolicyV2Args.builder()        
 *             .action(&#34;access_as_shared&#34;)
 *             .objectId(network1.id())
 *             .objectType(&#34;network&#34;)
 *             .targetTenant(&#34;20415a973c9e45d3917f078950644697&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RBAC policies can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:networking/rbacPolicyV2:RbacPolicyV2 rbac_policy_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
 * ```
 * 
 */
@ResourceType(type="openstack:networking/rbacPolicyV2:RbacPolicyV2")
public class RbacPolicyV2 extends com.pulumi.resources.CustomResource {
    /**
     * Action for the RBAC policy. Can either be
     * `access_as_external` or `access_as_shared`.
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output<String> action;

    /**
     * @return Action for the RBAC policy. Can either be
     * `access_as_external` or `access_as_shared`.
     * 
     */
    public Output<String> action() {
        return this.action;
    }
    /**
     * The ID of the `object_type` resource. An
     * `object_type` of `network` returns a network ID and an `object_type` of
     * `qos_policy` returns a QoS ID.
     * 
     */
    @Export(name="objectId", refs={String.class}, tree="[0]")
    private Output<String> objectId;

    /**
     * @return The ID of the `object_type` resource. An
     * `object_type` of `network` returns a network ID and an `object_type` of
     * `qos_policy` returns a QoS ID.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }
    /**
     * The type of the object that the RBAC policy
     * affects. Can be one of the following: `address_scope`, `address_group`,
     * `network`, `qos_policy`, `security_group` or `subnetpool`.
     * 
     */
    @Export(name="objectType", refs={String.class}, tree="[0]")
    private Output<String> objectType;

    /**
     * @return The type of the object that the RBAC policy
     * affects. Can be one of the following: `address_scope`, `address_group`,
     * `network`, `qos_policy`, `security_group` or `subnetpool`.
     * 
     */
    public Output<String> objectType() {
        return this.objectType;
    }
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The ID of the tenant to which the RBAC policy
     * will be enforced.
     * 
     */
    @Export(name="targetTenant", refs={String.class}, tree="[0]")
    private Output<String> targetTenant;

    /**
     * @return The ID of the tenant to which the RBAC policy
     * will be enforced.
     * 
     */
    public Output<String> targetTenant() {
        return this.targetTenant;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RbacPolicyV2(String name) {
        this(name, RbacPolicyV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RbacPolicyV2(String name, RbacPolicyV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RbacPolicyV2(String name, RbacPolicyV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/rbacPolicyV2:RbacPolicyV2", name, args == null ? RbacPolicyV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RbacPolicyV2(String name, Output<String> id, @Nullable RbacPolicyV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/rbacPolicyV2:RbacPolicyV2", name, state, makeResourceOptions(options, id));
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
    public static RbacPolicyV2 get(String name, Output<String> id, @Nullable RbacPolicyV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RbacPolicyV2(name, id, state, options);
    }
}
