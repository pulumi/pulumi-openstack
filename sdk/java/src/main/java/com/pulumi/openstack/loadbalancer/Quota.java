// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.loadbalancer.QuotaArgs;
import com.pulumi.openstack.loadbalancer.inputs.QuotaState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a V2 load balancer quota resource within OpenStack.
 * 
 * &gt; **Note:** This usually requires admin privileges.
 * 
 * &gt; **Note:** This resource is only available for Octavia.
 * 
 * &gt; **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack
 *    API in case of delete call.
 * 
 * &gt; **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
 *    created with zero value.
 * 
 * &gt; **Note:** This resource has attributes that depend on octavia minor versions.
 * Please ensure your Openstack cloud supports the required minor version.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.identity.Project;
 * import com.pulumi.openstack.loadbalancer.Quota;
 * import com.pulumi.openstack.loadbalancer.QuotaArgs;
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
 *         var project1 = new Project(&#34;project1&#34;);
 * 
 *         var quota1 = new Quota(&#34;quota1&#34;, QuotaArgs.builder()        
 *             .projectId(project1.id())
 *             .loadbalancer(6)
 *             .listener(7)
 *             .member(8)
 *             .pool(9)
 *             .healthMonitor(10)
 *             .l7Policy(11)
 *             .l7Rule(12)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Quotas can be imported using the `project_id/region_name`, where region_name is the
 * one defined is the Openstack credentials that are in use. E.g.
 * 
 * ```sh
 * $ pulumi import openstack:loadbalancer/quota:Quota quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
 * ```
 * 
 */
@ResourceType(type="openstack:loadbalancer/quota:Quota")
public class Quota extends com.pulumi.resources.CustomResource {
    /**
     * Quota value for health_monitors. Changing
     * this updates the existing quota. Omitting it sets it to 0.
     * 
     */
    @Export(name="healthMonitor", refs={Integer.class}, tree="[0]")
    private Output<Integer> healthMonitor;

    /**
     * @return Quota value for health_monitors. Changing
     * this updates the existing quota. Omitting it sets it to 0.
     * 
     */
    public Output<Integer> healthMonitor() {
        return this.healthMonitor;
    }
    /**
     * Quota value for l7_policies. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     * 
     */
    @Export(name="l7Policy", refs={Integer.class}, tree="[0]")
    private Output<Integer> l7Policy;

    /**
     * @return Quota value for l7_policies. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     * 
     */
    public Output<Integer> l7Policy() {
        return this.l7Policy;
    }
    /**
     * Quota value for l7_rules. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     * 
     */
    @Export(name="l7Rule", refs={Integer.class}, tree="[0]")
    private Output<Integer> l7Rule;

    /**
     * @return Quota value for l7_rules. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     * 
     */
    public Output<Integer> l7Rule() {
        return this.l7Rule;
    }
    /**
     * Quota value for listeners. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     * 
     */
    @Export(name="listener", refs={Integer.class}, tree="[0]")
    private Output<Integer> listener;

    /**
     * @return Quota value for listeners. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     * 
     */
    public Output<Integer> listener() {
        return this.listener;
    }
    /**
     * Quota value for loadbalancers. Changing this
     * updates the existing quota. Omitting it sets it to 0.
     * 
     */
    @Export(name="loadbalancer", refs={Integer.class}, tree="[0]")
    private Output<Integer> loadbalancer;

    /**
     * @return Quota value for loadbalancers. Changing this
     * updates the existing quota. Omitting it sets it to 0.
     * 
     */
    public Output<Integer> loadbalancer() {
        return this.loadbalancer;
    }
    /**
     * Quota value for members. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     * 
     */
    @Export(name="member", refs={Integer.class}, tree="[0]")
    private Output<Integer> member;

    /**
     * @return Quota value for members. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     * 
     */
    public Output<Integer> member() {
        return this.member;
    }
    /**
     * Quota value for pools. Changing this updates the
     * the existing quota. Omitting it sets it to 0.
     * 
     */
    @Export(name="pool", refs={Integer.class}, tree="[0]")
    private Output<Integer> pool;

    /**
     * @return Quota value for pools. Changing this updates the
     * the existing quota. Omitting it sets it to 0.
     * 
     */
    public Output<Integer> pool() {
        return this.pool;
    }
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quota.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return ID of the project to manage quotas. Changing this
     * creates a new quota.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Region in which to manage quotas. Changing this
     * creates a new quota. If ommited, the region of the credentials is used.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Region in which to manage quotas. Changing this
     * creates a new quota. If ommited, the region of the credentials is used.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Quota(String name) {
        this(name, QuotaArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Quota(String name, QuotaArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Quota(String name, QuotaArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/quota:Quota", name, args == null ? QuotaArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Quota(String name, Output<String> id, @Nullable QuotaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:loadbalancer/quota:Quota", name, state, makeResourceOptions(options, id));
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
    public static Quota get(String name, Output<String> id, @Nullable QuotaState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Quota(name, id, state, options);
    }
}
