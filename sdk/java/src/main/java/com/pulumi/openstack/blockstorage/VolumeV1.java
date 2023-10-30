// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.blockstorage.VolumeV1Args;
import com.pulumi.openstack.blockstorage.inputs.VolumeV1State;
import com.pulumi.openstack.blockstorage.outputs.VolumeV1Attachment;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 volume resource within OpenStack.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.blockstorage.VolumeV1;
 * import com.pulumi.openstack.blockstorage.VolumeV1Args;
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
 *         var volume1 = new VolumeV1(&#34;volume1&#34;, VolumeV1Args.builder()        
 *             .description(&#34;first test volume&#34;)
 *             .region(&#34;RegionOne&#34;)
 *             .size(3)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Volumes can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:blockstorage/volumeV1:VolumeV1 volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
 * ```
 * 
 */
@ResourceType(type="openstack:blockstorage/volumeV1:VolumeV1")
public class VolumeV1 extends com.pulumi.resources.CustomResource {
    /**
     * If a volume is attached to an instance, this attribute will
     * display the Attachment ID, Instance ID, and the Device as the Instance
     * sees it.
     * 
     */
    @Export(name="attachments", type=List.class, parameters={VolumeV1Attachment.class})
    private Output<List<VolumeV1Attachment>> attachments;

    /**
     * @return If a volume is attached to an instance, this attribute will
     * display the Attachment ID, Instance ID, and the Device as the Instance
     * sees it.
     * 
     */
    public Output<List<VolumeV1Attachment>> attachments() {
        return this.attachments;
    }
    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     * 
     */
    @Export(name="availabilityZone", type=String.class, parameters={})
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
     * A description of the volume. Changing this updates
     * the volume&#39;s description.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
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
     * The image ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    @Export(name="imageId", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageId;

    /**
     * @return The image ID from which to create the volume.
     * Changing this creates a new volume.
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
    @Export(name="metadata", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> metadata;

    /**
     * @return Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     * 
     */
    public Output<Map<String,Object>> metadata() {
        return this.metadata;
    }
    /**
     * A unique name for the volume. Changing this updates the
     * volume&#39;s name.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
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
    @Export(name="region", type=String.class, parameters={})
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
     * The size of the volume to create (in gigabytes). Changing
     * this creates a new volume.
     * 
     */
    @Export(name="size", type=Integer.class, parameters={})
    private Output<Integer> size;

    /**
     * @return The size of the volume to create (in gigabytes). Changing
     * this creates a new volume.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    @Export(name="snapshotId", type=String.class, parameters={})
    private Output</* @Nullable */ String> snapshotId;

    /**
     * @return The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    public Output<Optional<String>> snapshotId() {
        return Codegen.optional(this.snapshotId);
    }
    /**
     * The volume ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    @Export(name="sourceVolId", type=String.class, parameters={})
    private Output</* @Nullable */ String> sourceVolId;

    /**
     * @return The volume ID from which to create the volume.
     * Changing this creates a new volume.
     * 
     */
    public Output<Optional<String>> sourceVolId() {
        return Codegen.optional(this.sourceVolId);
    }
    /**
     * The type of volume to create.
     * Changing this creates a new volume.
     * 
     */
    @Export(name="volumeType", type=String.class, parameters={})
    private Output<String> volumeType;

    /**
     * @return The type of volume to create.
     * Changing this creates a new volume.
     * 
     */
    public Output<String> volumeType() {
        return this.volumeType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VolumeV1(String name) {
        this(name, VolumeV1Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VolumeV1(String name, VolumeV1Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VolumeV1(String name, VolumeV1Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/volumeV1:VolumeV1", name, args == null ? VolumeV1Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VolumeV1(String name, Output<String> id, @Nullable VolumeV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/volumeV1:VolumeV1", name, state, makeResourceOptions(options, id));
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
    public static VolumeV1 get(String name, Output<String> id, @Nullable VolumeV1State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VolumeV1(name, id, state, options);
    }
}
