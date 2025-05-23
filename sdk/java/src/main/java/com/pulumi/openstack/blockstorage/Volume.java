// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.blockstorage.VolumeArgs;
import com.pulumi.openstack.blockstorage.inputs.VolumeState;
import com.pulumi.openstack.blockstorage.outputs.VolumeAttachment;
import com.pulumi.openstack.blockstorage.outputs.VolumeSchedulerHint;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V3 volume resource within OpenStack.
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
 * import com.pulumi.openstack.blockstorage.Volume;
 * import com.pulumi.openstack.blockstorage.VolumeArgs;
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
 *         var volume1 = new Volume("volume1", VolumeArgs.builder()
 *             .region("RegionOne")
 *             .name("volume_1")
 *             .description("first test volume")
 *             .size(3)
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
 * Volumes can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:blockstorage/volume:Volume volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
 * ```
 * 
 */
@ResourceType(type="openstack:blockstorage/volume:Volume")
public class Volume extends com.pulumi.resources.CustomResource {
    /**
     * If a volume is attached to an instance, this attribute will
     * display the Attachment ID, Instance ID, and the Device as the Instance
     * sees it.
     * 
     */
    @Export(name="attachments", refs={List.class,VolumeAttachment.class}, tree="[0,1]")
    private Output<List<VolumeAttachment>> attachments;

    /**
     * @return If a volume is attached to an instance, this attribute will
     * display the Attachment ID, Instance ID, and the Device as the Instance
     * sees it.
     * 
     */
    public Output<List<VolumeAttachment>> attachments() {
        return this.attachments;
    }
    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return The availability zone for the volume.
     * Changing this creates a new volume.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * The backup ID from which to create the volume.
     * Conflicts with `snapshot_id`, `source_vol_id`, `image_id`. Changing this
     * creates a new volume. Requires microversion &gt;= 3.47.
     * 
     */
    @Export(name="backupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backupId;

    /**
     * @return The backup ID from which to create the volume.
     * Conflicts with `snapshot_id`, `source_vol_id`, `image_id`. Changing this
     * creates a new volume. Requires microversion &gt;= 3.47.
     * 
     */
    public Output<Optional<String>> backupId() {
        return Codegen.optional(this.backupId);
    }
    /**
     * The consistency group to place the volume
     * in.
     * 
     */
    @Export(name="consistencyGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> consistencyGroupId;

    /**
     * @return The consistency group to place the volume
     * in.
     * 
     */
    public Output<Optional<String>> consistencyGroupId() {
        return Codegen.optional(this.consistencyGroupId);
    }
    /**
     * A description of the volume. Changing this updates
     * the volume&#39;s description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the volume. Changing this updates
     * the volume&#39;s description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * When this option is set it allows extending
     * attached volumes. Note: updating size of an attached volume requires Cinder
     * support for version 3.42 and a compatible storage driver.
     * 
     */
    @Export(name="enableOnlineResize", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableOnlineResize;

    /**
     * @return When this option is set it allows extending
     * attached volumes. Note: updating size of an attached volume requires Cinder
     * support for version 3.42 and a compatible storage driver.
     * 
     */
    public Output<Optional<Boolean>> enableOnlineResize() {
        return Codegen.optional(this.enableOnlineResize);
    }
    /**
     * The image ID from which to create the volume.
     * Conflicts with `snapshot_id`, `source_vol_id`, `backup_id`. Changing this
     * creates a new volume.
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> imageId;

    /**
     * @return The image ID from which to create the volume.
     * Conflicts with `snapshot_id`, `source_vol_id`, `backup_id`. Changing this
     * creates a new volume.
     * 
     */
    public Output<Optional<String>> imageId() {
        return Codegen.optional(this.imageId);
    }
    /**
     * Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> metadata;

    /**
     * @return Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     * 
     */
    public Output<Map<String,String>> metadata() {
        return this.metadata;
    }
    /**
     * A unique name for the volume. Changing this updates the
     * volume&#39;s name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the volume. Changing this updates the
     * volume&#39;s name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Provide the Cinder scheduler with hints on where
     * to instantiate a volume in the OpenStack cloud. The available hints are described below.
     * 
     */
    @Export(name="schedulerHints", refs={List.class,VolumeSchedulerHint.class}, tree="[0,1]")
    private Output</* @Nullable */ List<VolumeSchedulerHint>> schedulerHints;

    /**
     * @return Provide the Cinder scheduler with hints on where
     * to instantiate a volume in the OpenStack cloud. The available hints are described below.
     * 
     */
    public Output<Optional<List<VolumeSchedulerHint>>> schedulerHints() {
        return Codegen.optional(this.schedulerHints);
    }
    /**
     * The size of the volume to create (in gigabytes).
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return The size of the volume to create (in gigabytes).
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * The snapshot ID from which to create the volume.
     * Conflicts with `source_vol_id`, `image_id`, `backup_id`. Changing this
     * creates a new volume.
     * 
     */
    @Export(name="snapshotId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snapshotId;

    /**
     * @return The snapshot ID from which to create the volume.
     * Conflicts with `source_vol_id`, `image_id`, `backup_id`. Changing this
     * creates a new volume.
     * 
     */
    public Output<Optional<String>> snapshotId() {
        return Codegen.optional(this.snapshotId);
    }
    /**
     * The volume ID to replicate with.
     * 
     */
    @Export(name="sourceReplica", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceReplica;

    /**
     * @return The volume ID to replicate with.
     * 
     */
    public Output<Optional<String>> sourceReplica() {
        return Codegen.optional(this.sourceReplica);
    }
    /**
     * The volume ID from which to create the volume.
     * Conflicts with `snapshot_id`, `image_id`, `backup_id`. Changing this
     * creates a new volume.
     * 
     */
    @Export(name="sourceVolId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceVolId;

    /**
     * @return The volume ID from which to create the volume.
     * Conflicts with `snapshot_id`, `image_id`, `backup_id`. Changing this
     * creates a new volume.
     * 
     */
    public Output<Optional<String>> sourceVolId() {
        return Codegen.optional(this.sourceVolId);
    }
    /**
     * Migration policy when changing `volume_type`.
     * `&#34;never&#34;` *(default)* prevents migration to another storage backend, while `&#34;on-demand&#34;`
     * allows migration if needed. Applicable only when updating `volume_type`.
     * 
     */
    @Export(name="volumeRetypePolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> volumeRetypePolicy;

    /**
     * @return Migration policy when changing `volume_type`.
     * `&#34;never&#34;` *(default)* prevents migration to another storage backend, while `&#34;on-demand&#34;`
     * allows migration if needed. Applicable only when updating `volume_type`.
     * 
     */
    public Output<Optional<String>> volumeRetypePolicy() {
        return Codegen.optional(this.volumeRetypePolicy);
    }
    /**
     * The type of volume to create or update.
     * Changing this will attempt an in-place retype operation; migration depends on `volume_retype_policy`.
     * 
     */
    @Export(name="volumeType", refs={String.class}, tree="[0]")
    private Output<String> volumeType;

    /**
     * @return The type of volume to create or update.
     * Changing this will attempt an in-place retype operation; migration depends on `volume_retype_policy`.
     * 
     */
    public Output<String> volumeType() {
        return this.volumeType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Volume(java.lang.String name) {
        this(name, VolumeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Volume(java.lang.String name, VolumeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Volume(java.lang.String name, VolumeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/volume:Volume", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Volume(java.lang.String name, Output<java.lang.String> id, @Nullable VolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/volume:Volume", name, state, makeResourceOptions(options, id), false);
    }

    private static VolumeArgs makeArgs(VolumeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VolumeArgs.Empty : args;
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
    public static Volume get(java.lang.String name, Output<java.lang.String> id, @Nullable VolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Volume(name, id, state, options);
    }
}
