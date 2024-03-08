// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.VolumeAttachArgs;
import com.pulumi.openstack.compute.inputs.VolumeAttachState;
import com.pulumi.openstack.compute.outputs.VolumeAttachVendorOptions;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Attaches a Block Storage Volume to an Instance using the OpenStack
 * Compute (Nova) v2 API.
 * 
 * ## Example Usage
 * 
 * ### Basic attachment of a single volume to a single instance
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.blockstorage.Volume;
 * import com.pulumi.openstack.blockstorage.VolumeArgs;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.VolumeAttach;
 * import com.pulumi.openstack.compute.VolumeAttachArgs;
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
 *         var volume1 = new Volume(&#34;volume1&#34;, VolumeArgs.builder()        
 *             .size(1)
 *             .build());
 * 
 *         var instance1 = new Instance(&#34;instance1&#34;, InstanceArgs.builder()        
 *             .securityGroups(&#34;default&#34;)
 *             .build());
 * 
 *         var va1 = new VolumeAttach(&#34;va1&#34;, VolumeAttachArgs.builder()        
 *             .instanceId(instance1.id())
 *             .volumeId(volume1.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Using Multiattach-enabled volumes
 * 
 * Multiattach Volumes are dependent upon your OpenStack cloud and not all
 * clouds support multiattach.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.blockstorage.Volume;
 * import com.pulumi.openstack.blockstorage.VolumeArgs;
 * import com.pulumi.openstack.compute.Instance;
 * import com.pulumi.openstack.compute.InstanceArgs;
 * import com.pulumi.openstack.compute.VolumeAttach;
 * import com.pulumi.openstack.compute.VolumeAttachArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var volume1 = new Volume(&#34;volume1&#34;, VolumeArgs.builder()        
 *             .size(1)
 *             .multiattach(true)
 *             .build());
 * 
 *         var instance1 = new Instance(&#34;instance1&#34;, InstanceArgs.builder()        
 *             .securityGroups(&#34;default&#34;)
 *             .build());
 * 
 *         var instance2 = new Instance(&#34;instance2&#34;, InstanceArgs.builder()        
 *             .securityGroups(&#34;default&#34;)
 *             .build());
 * 
 *         var va1 = new VolumeAttach(&#34;va1&#34;, VolumeAttachArgs.builder()        
 *             .instanceId(instance1.id())
 *             .volumeId(volume1.id())
 *             .multiattach(true)
 *             .build());
 * 
 *         var va2 = new VolumeAttach(&#34;va2&#34;, VolumeAttachArgs.builder()        
 *             .instanceId(instance2.id())
 *             .volumeId(volume1.id())
 *             .multiattach(true)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(&#34;openstack_compute_volume_attach_v2.va_1&#34;)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * It is recommended to use `depends_on` for the attach resources
 * to enforce the volume attachments to happen one at a time.
 * 
 * ## Import
 * 
 * Volume Attachments can be imported using the Instance ID and Volume ID
 * separated by a slash, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:compute/volumeAttach:VolumeAttach va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
 * ```
 * 
 */
@ResourceType(type="openstack:compute/volumeAttach:VolumeAttach")
public class VolumeAttach extends com.pulumi.resources.CustomResource {
    @Export(name="device", refs={String.class}, tree="[0]")
    private Output<String> device;

    public Output<String> device() {
        return this.device;
    }
    /**
     * The ID of the Instance to attach the Volume to.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the Instance to attach the Volume to.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Enable attachment of multiattach-capable volumes.
     * 
     */
    @Export(name="multiattach", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> multiattach;

    /**
     * @return Enable attachment of multiattach-capable volumes.
     * 
     */
    public Output<Optional<Boolean>> multiattach() {
        return Codegen.optional(this.multiattach);
    }
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     * 
     */
    @Export(name="vendorOptions", refs={VolumeAttachVendorOptions.class}, tree="[0]")
    private Output</* @Nullable */ VolumeAttachVendorOptions> vendorOptions;

    /**
     * @return Map of additional vendor-specific options.
     * Supported options are described below.
     * 
     */
    public Output<Optional<VolumeAttachVendorOptions>> vendorOptions() {
        return Codegen.optional(this.vendorOptions);
    }
    /**
     * The ID of the Volume to attach to an Instance.
     * 
     */
    @Export(name="volumeId", refs={String.class}, tree="[0]")
    private Output<String> volumeId;

    /**
     * @return The ID of the Volume to attach to an Instance.
     * 
     */
    public Output<String> volumeId() {
        return this.volumeId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VolumeAttach(String name) {
        this(name, VolumeAttachArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VolumeAttach(String name, VolumeAttachArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VolumeAttach(String name, VolumeAttachArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/volumeAttach:VolumeAttach", name, args == null ? VolumeAttachArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VolumeAttach(String name, Output<String> id, @Nullable VolumeAttachState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/volumeAttach:VolumeAttach", name, state, makeResourceOptions(options, id));
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
    public static VolumeAttach get(String name, Output<String> id, @Nullable VolumeAttachState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VolumeAttach(name, id, state, options);
    }
}
