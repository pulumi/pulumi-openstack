// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.blockstorage.VolumeAttachArgs;
import com.pulumi.openstack.blockstorage.inputs.VolumeAttachState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; **Note:** This resource usually requires admin privileges.
 * 
 * &gt; **Note:** This resource does not actually attach a volume to an instance.
 * Please use the `openstack.compute.VolumeAttach` resource for that.
 * 
 * &gt; **Note:** All arguments including the `data` computed attribute will be
 * stored in the raw state as plain-text. Read more about sensitive data in
 * state.
 * 
 * Creates a general purpose attachment connection to a Block
 * Storage volume using the OpenStack Block Storage (Cinder) v3 API.
 * 
 * Depending on your Block Storage service configuration, this
 * resource can assist in attaching a volume to a non-OpenStack resource
 * such as a bare-metal server or a remote virtual machine in a
 * different cloud provider.
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
 * import com.pulumi.openstack.blockstorage.VolumeAttach;
 * import com.pulumi.openstack.blockstorage.VolumeAttachArgs;
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
 *             .name("volume_1")
 *             .size(1)
 *             .build());
 * 
 *         var va1 = new VolumeAttach("va1", VolumeAttachArgs.builder()
 *             .volumeId(volume1.id())
 *             .device("auto")
 *             .hostName("devstack")
 *             .ipAddress("192.168.255.10")
 *             .initiator("iqn.1993-08.org.debian:01:e9861fb1859")
 *             .osType("linux2")
 *             .platform("x86_64")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Volume Connection Data
 * 
 * Upon creation of this resource, a `data` exported attribute will be available.
 * This attribute is a set of key/value pairs that contains the information
 * required to complete the block storage connection.
 * 
 * As an example, creating an iSCSI-based volume will return the following:
 * 
 * This information can then be fed into a provisioner or a template shell script,
 * where the final result would look something like:
 * 
 * The contents of `data` will vary from each Block Storage service. You must have
 * a good understanding of how the service is configured and how to make the
 * appropriate final connection. However, if used correctly, this has the
 * flexibility to be able to attach OpenStack Block Storage volumes to
 * non-OpenStack resources.
 * 
 * ## Import
 * 
 * It is not possible to import this resource.
 * 
 */
@ResourceType(type="openstack:blockstorage/volumeAttach:VolumeAttach")
public class VolumeAttach extends com.pulumi.resources.CustomResource {
    /**
     * Specify whether to attach the volume as Read-Only
     * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
     * If left unspecified, the Block Storage API will apply a default of `rw`.
     * 
     */
    @Export(name="attachMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> attachMode;

    /**
     * @return Specify whether to attach the volume as Read-Only
     * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
     * If left unspecified, the Block Storage API will apply a default of `rw`.
     * 
     */
    public Output<Optional<String>> attachMode() {
        return Codegen.optional(this.attachMode);
    }
    /**
     * This is a map of key/value pairs that contain the connection
     * information. You will want to pass this information to a provisioner
     * script to finalize the connection. See below for more information.
     * 
     */
    @Export(name="data", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> data;

    /**
     * @return This is a map of key/value pairs that contain the connection
     * information. You will want to pass this information to a provisioner
     * script to finalize the connection. See below for more information.
     * 
     */
    public Output<Map<String,String>> data() {
        return this.data;
    }
    /**
     * The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     * 
     */
    @Export(name="device", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> device;

    /**
     * @return The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     * 
     */
    public Output<Optional<String>> device() {
        return Codegen.optional(this.device);
    }
    /**
     * The storage driver that the volume is based on.
     * 
     */
    @Export(name="driverVolumeType", refs={String.class}, tree="[0]")
    private Output<String> driverVolumeType;

    /**
     * @return The storage driver that the volume is based on.
     * 
     */
    public Output<String> driverVolumeType() {
        return this.driverVolumeType;
    }
    /**
     * The host to attach the volume to.
     * 
     */
    @Export(name="hostName", refs={String.class}, tree="[0]")
    private Output<String> hostName;

    /**
     * @return The host to attach the volume to.
     * 
     */
    public Output<String> hostName() {
        return this.hostName;
    }
    /**
     * The iSCSI initiator string to make the connection.
     * 
     */
    @Export(name="initiator", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> initiator;

    /**
     * @return The iSCSI initiator string to make the connection.
     * 
     */
    public Output<Optional<String>> initiator() {
        return Codegen.optional(this.initiator);
    }
    /**
     * The IP address of the `host_name` above.
     * 
     */
    @Export(name="ipAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipAddress;

    /**
     * @return The IP address of the `host_name` above.
     * 
     */
    public Output<Optional<String>> ipAddress() {
        return Codegen.optional(this.ipAddress);
    }
    /**
     * A mount point base name for shared storage.
     * 
     */
    @Export(name="mountPointBase", refs={String.class}, tree="[0]")
    private Output<String> mountPointBase;

    /**
     * @return A mount point base name for shared storage.
     * 
     */
    public Output<String> mountPointBase() {
        return this.mountPointBase;
    }
    /**
     * Whether to connect to this volume via multipath.
     * 
     */
    @Export(name="multipath", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> multipath;

    /**
     * @return Whether to connect to this volume via multipath.
     * 
     */
    public Output<Optional<Boolean>> multipath() {
        return Codegen.optional(this.multipath);
    }
    /**
     * The iSCSI initiator OS type.
     * 
     */
    @Export(name="osType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> osType;

    /**
     * @return The iSCSI initiator OS type.
     * 
     */
    public Output<Optional<String>> osType() {
        return Codegen.optional(this.osType);
    }
    /**
     * The iSCSI initiator platform.
     * 
     */
    @Export(name="platform", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> platform;

    /**
     * @return The iSCSI initiator platform.
     * 
     */
    public Output<Optional<String>> platform() {
        return Codegen.optional(this.platform);
    }
    /**
     * The region in which to obtain the V3 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V3 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     * 
     */
    public Output<String> region() {
        return this.region;
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
     * A wwnn name. Used for Fibre Channel connections.
     * 
     */
    @Export(name="wwnn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> wwnn;

    /**
     * @return A wwnn name. Used for Fibre Channel connections.
     * 
     */
    public Output<Optional<String>> wwnn() {
        return Codegen.optional(this.wwnn);
    }
    /**
     * An array of wwpn strings. Used for Fibre Channel
     * connections.
     * 
     */
    @Export(name="wwpns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> wwpns;

    /**
     * @return An array of wwpn strings. Used for Fibre Channel
     * connections.
     * 
     */
    public Output<Optional<List<String>>> wwpns() {
        return Codegen.optional(this.wwpns);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VolumeAttach(java.lang.String name) {
        this(name, VolumeAttachArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VolumeAttach(java.lang.String name, VolumeAttachArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VolumeAttach(java.lang.String name, VolumeAttachArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/volumeAttach:VolumeAttach", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private VolumeAttach(java.lang.String name, Output<java.lang.String> id, @Nullable VolumeAttachState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/volumeAttach:VolumeAttach", name, state, makeResourceOptions(options, id), false);
    }

    private static VolumeAttachArgs makeArgs(VolumeAttachArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? VolumeAttachArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "data"
            ))
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
    public static VolumeAttach get(java.lang.String name, Output<java.lang.String> id, @Nullable VolumeAttachState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VolumeAttach(name, id, state, options);
    }
}
