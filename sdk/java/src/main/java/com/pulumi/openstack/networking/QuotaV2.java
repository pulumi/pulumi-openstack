// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.QuotaV2Args;
import com.pulumi.openstack.networking.inputs.QuotaV2State;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a V2 networking quota resource within OpenStack.
 * 
 * &gt; **Note:** This usually requires admin privileges.
 * 
 * &gt; **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
 *     in case of delete call.
 * 
 * &gt; **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
 *     created with zero value.
 * 
 * ## Import
 * 
 * Quotas can be imported using the `project_id/region_name`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:networking/quotaV2:QuotaV2 quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
 * ```
 * 
 */
@ResourceType(type="openstack:networking/quotaV2:QuotaV2")
public class QuotaV2 extends com.pulumi.resources.CustomResource {
    /**
     * Quota value for floating IPs. Changing this updates the
     * existing quota.
     * 
     */
    @Export(name="floatingip", refs={Integer.class}, tree="[0]")
    private Output<Integer> floatingip;

    /**
     * @return Quota value for floating IPs. Changing this updates the
     * existing quota.
     * 
     */
    public Output<Integer> floatingip() {
        return this.floatingip;
    }
    /**
     * Quota value for networks. Changing this updates the
     * existing quota.
     * 
     */
    @Export(name="network", refs={Integer.class}, tree="[0]")
    private Output<Integer> network;

    /**
     * @return Quota value for networks. Changing this updates the
     * existing quota.
     * 
     */
    public Output<Integer> network() {
        return this.network;
    }
    /**
     * Quota value for ports. Changing this updates the
     * existing quota.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return Quota value for ports. Changing this updates the
     * existing quota.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * ID of the project to manage quota. Changing this
     * creates new quota.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return ID of the project to manage quota. Changing this
     * creates new quota.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Quota value for RBAC policies.
     * Changing this updates the existing quota.
     * 
     */
    @Export(name="rbacPolicy", refs={Integer.class}, tree="[0]")
    private Output<Integer> rbacPolicy;

    /**
     * @return Quota value for RBAC policies.
     * Changing this updates the existing quota.
     * 
     */
    public Output<Integer> rbacPolicy() {
        return this.rbacPolicy;
    }
    /**
     * The region in which to create the quota. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates new quota.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the quota. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates new quota.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Quota value for routers. Changing this updates the
     * existing quota.
     * 
     */
    @Export(name="router", refs={Integer.class}, tree="[0]")
    private Output<Integer> router;

    /**
     * @return Quota value for routers. Changing this updates the
     * existing quota.
     * 
     */
    public Output<Integer> router() {
        return this.router;
    }
    /**
     * Quota value for security groups. Changing
     * this updates the existing quota.
     * 
     */
    @Export(name="securityGroup", refs={Integer.class}, tree="[0]")
    private Output<Integer> securityGroup;

    /**
     * @return Quota value for security groups. Changing
     * this updates the existing quota.
     * 
     */
    public Output<Integer> securityGroup() {
        return this.securityGroup;
    }
    /**
     * Quota value for security group rules.
     * Changing this updates the existing quota.
     * 
     */
    @Export(name="securityGroupRule", refs={Integer.class}, tree="[0]")
    private Output<Integer> securityGroupRule;

    /**
     * @return Quota value for security group rules.
     * Changing this updates the existing quota.
     * 
     */
    public Output<Integer> securityGroupRule() {
        return this.securityGroupRule;
    }
    /**
     * Quota value for subnets. Changing
     * this updates the existing quota.
     * 
     */
    @Export(name="subnet", refs={Integer.class}, tree="[0]")
    private Output<Integer> subnet;

    /**
     * @return Quota value for subnets. Changing
     * this updates the existing quota.
     * 
     */
    public Output<Integer> subnet() {
        return this.subnet;
    }
    /**
     * Quota value for subnetpools.
     * Changing this updates the existing quota.
     * 
     */
    @Export(name="subnetpool", refs={Integer.class}, tree="[0]")
    private Output<Integer> subnetpool;

    /**
     * @return Quota value for subnetpools.
     * Changing this updates the existing quota.
     * 
     */
    public Output<Integer> subnetpool() {
        return this.subnetpool;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QuotaV2(String name) {
        this(name, QuotaV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QuotaV2(String name, QuotaV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QuotaV2(String name, QuotaV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/quotaV2:QuotaV2", name, args == null ? QuotaV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QuotaV2(String name, Output<String> id, @Nullable QuotaV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/quotaV2:QuotaV2", name, state, makeResourceOptions(options, id));
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
    public static QuotaV2 get(String name, Output<String> id, @Nullable QuotaV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QuotaV2(name, id, state, options);
    }
}
