// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.blockstorage.QuoteSetV2Args;
import com.pulumi.openstack.blockstorage.inputs.QuoteSetV2State;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 block storage quotaset resource within OpenStack.
 * 
 * &gt; **Note:** This usually requires admin privileges.
 * 
 * &gt; **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
 *     in case of delete call.
 * 
 * &gt; **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
 *     created with zero value. This excludes volume type quota.
 * 
 * ## Import
 * 
 * Quotasets can be imported using the `project_id/region`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:blockstorage/quoteSetV2:QuoteSetV2 quotaset_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
 * ```
 * 
 */
@ResourceType(type="openstack:blockstorage/quoteSetV2:QuoteSetV2")
public class QuoteSetV2 extends com.pulumi.resources.CustomResource {
    /**
     * Quota value for backup gigabytes. Changing
     * this updates the existing quotaset.
     * 
     */
    @Export(name="backupGigabytes", refs={Integer.class}, tree="[0]")
    private Output<Integer> backupGigabytes;

    /**
     * @return Quota value for backup gigabytes. Changing
     * this updates the existing quotaset.
     * 
     */
    public Output<Integer> backupGigabytes() {
        return this.backupGigabytes;
    }
    /**
     * Quota value for backups. Changing this updates the
     * existing quotaset.
     * 
     */
    @Export(name="backups", refs={Integer.class}, tree="[0]")
    private Output<Integer> backups;

    /**
     * @return Quota value for backups. Changing this updates the
     * existing quotaset.
     * 
     */
    public Output<Integer> backups() {
        return this.backups;
    }
    /**
     * Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     * 
     */
    @Export(name="gigabytes", refs={Integer.class}, tree="[0]")
    private Output<Integer> gigabytes;

    /**
     * @return Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     * 
     */
    public Output<Integer> gigabytes() {
        return this.gigabytes;
    }
    /**
     * Quota value for groups. Changing this updates the
     * existing quotaset.
     * 
     */
    @Export(name="groups", refs={Integer.class}, tree="[0]")
    private Output<Integer> groups;

    /**
     * @return Quota value for groups. Changing this updates the
     * existing quotaset.
     * 
     */
    public Output<Integer> groups() {
        return this.groups;
    }
    /**
     * Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     * 
     */
    @Export(name="perVolumeGigabytes", refs={Integer.class}, tree="[0]")
    private Output<Integer> perVolumeGigabytes;

    /**
     * @return Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     * 
     */
    public Output<Integer> perVolumeGigabytes() {
        return this.perVolumeGigabytes;
    }
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Quota value for snapshots. Changing this updates the
     * existing quotaset.
     * 
     */
    @Export(name="snapshots", refs={Integer.class}, tree="[0]")
    private Output<Integer> snapshots;

    /**
     * @return Quota value for snapshots. Changing this updates the
     * existing quotaset.
     * 
     */
    public Output<Integer> snapshots() {
        return this.snapshots;
    }
    /**
     * Key/Value pairs for setting quota for
     * volumes types. Possible keys are `snapshots_&lt;volume_type_name&gt;`,
     * `volumes_&lt;volume_type_name&gt;` and `gigabytes_&lt;volume_type_name&gt;`.
     * 
     */
    @Export(name="volumeTypeQuota", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> volumeTypeQuota;

    /**
     * @return Key/Value pairs for setting quota for
     * volumes types. Possible keys are `snapshots_&lt;volume_type_name&gt;`,
     * `volumes_&lt;volume_type_name&gt;` and `gigabytes_&lt;volume_type_name&gt;`.
     * 
     */
    public Output<Optional<Map<String,Object>>> volumeTypeQuota() {
        return Codegen.optional(this.volumeTypeQuota);
    }
    /**
     * Quota value for volumes. Changing this updates the
     * existing quotaset.
     * 
     */
    @Export(name="volumes", refs={Integer.class}, tree="[0]")
    private Output<Integer> volumes;

    /**
     * @return Quota value for volumes. Changing this updates the
     * existing quotaset.
     * 
     */
    public Output<Integer> volumes() {
        return this.volumes;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QuoteSetV2(java.lang.String name) {
        this(name, QuoteSetV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QuoteSetV2(java.lang.String name, QuoteSetV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QuoteSetV2(java.lang.String name, QuoteSetV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/quoteSetV2:QuoteSetV2", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private QuoteSetV2(java.lang.String name, Output<java.lang.String> id, @Nullable QuoteSetV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/quoteSetV2:QuoteSetV2", name, state, makeResourceOptions(options, id), false);
    }

    private static QuoteSetV2Args makeArgs(QuoteSetV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? QuoteSetV2Args.Empty : args;
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
    public static QuoteSetV2 get(java.lang.String name, Output<java.lang.String> id, @Nullable QuoteSetV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QuoteSetV2(name, id, state, options);
    }
}
