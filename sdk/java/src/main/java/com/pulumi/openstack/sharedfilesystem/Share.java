// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.sharedfilesystem.ShareArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.ShareState;
import com.pulumi.openstack.sharedfilesystem.outputs.ShareExportLocation;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Use this resource to configure a share.
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
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Subnet;
 * import com.pulumi.openstack.networking.SubnetArgs;
 * import com.pulumi.openstack.sharedfilesystem.ShareNetwork;
 * import com.pulumi.openstack.sharedfilesystem.ShareNetworkArgs;
 * import com.pulumi.openstack.sharedfilesystem.Share;
 * import com.pulumi.openstack.sharedfilesystem.ShareArgs;
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
 *         var network1 = new Network("network1", NetworkArgs.builder()
 *             .name("network_1")
 *             .adminStateUp("true")
 *             .build());
 * 
 *         var subnet1 = new Subnet("subnet1", SubnetArgs.builder()
 *             .name("subnet_1")
 *             .cidr("192.168.199.0/24")
 *             .ipVersion(4)
 *             .networkId(network1.id())
 *             .build());
 * 
 *         var sharenetwork1 = new ShareNetwork("sharenetwork1", ShareNetworkArgs.builder()
 *             .name("test_sharenetwork")
 *             .description("test share network with security services")
 *             .neutronNetId(network1.id())
 *             .neutronSubnetId(subnet1.id())
 *             .build());
 * 
 *         var share1 = new Share("share1", ShareArgs.builder()
 *             .name("nfs_share")
 *             .description("test share description")
 *             .shareProto("NFS")
 *             .size(1)
 *             .shareNetworkId(sharenetwork1.id())
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
 * This resource can be imported by specifying the ID of the share:
 * 
 * ```sh
 * $ pulumi import openstack:sharedfilesystem/share:Share share_1 id
 * ```
 * 
 */
@ResourceType(type="openstack:sharedfilesystem/share:Share")
public class Share extends com.pulumi.resources.CustomResource {
    /**
     * The map of metadata, assigned on the share, which has been
     * explicitly and implicitly added.
     * 
     */
    @Export(name="allMetadata", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> allMetadata;

    /**
     * @return The map of metadata, assigned on the share, which has been
     * explicitly and implicitly added.
     * 
     */
    public Output<Map<String,Object>> allMetadata() {
        return this.allMetadata;
    }
    /**
     * The share availability zone. Changing this creates a
     * new share.
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return The share availability zone. Changing this creates a
     * new share.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * The human-readable description for the share.
     * Changing this updates the description of the existing share.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The human-readable description for the share.
     * Changing this updates the description of the existing share.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * A list of export locations. For example, when a share server
     * has more than one network interface, it can have multiple export locations.
     * 
     */
    @Export(name="exportLocations", refs={List.class,ShareExportLocation.class}, tree="[0,1]")
    private Output<List<ShareExportLocation>> exportLocations;

    /**
     * @return A list of export locations. For example, when a share server
     * has more than one network interface, it can have multiple export locations.
     * 
     */
    public Output<List<ShareExportLocation>> exportLocations() {
        return this.exportLocations;
    }
    /**
     * Indicates whether a share has replicas or not.
     * 
     */
    @Export(name="hasReplicas", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> hasReplicas;

    /**
     * @return Indicates whether a share has replicas or not.
     * 
     */
    public Output<Boolean> hasReplicas() {
        return this.hasReplicas;
    }
    /**
     * The share host name.
     * 
     */
    @Export(name="host", refs={String.class}, tree="[0]")
    private Output<String> host;

    /**
     * @return The share host name.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * The level of visibility for the share. Set to true to make
     * share public. Set to false to make it private. Default value is false. Changing this
     * updates the existing share.
     * 
     */
    @Export(name="isPublic", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isPublic;

    /**
     * @return The level of visibility for the share. Set to true to make
     * share public. Set to false to make it private. Default value is false. Changing this
     * updates the existing share.
     * 
     */
    public Output<Optional<Boolean>> isPublic() {
        return Codegen.optional(this.isPublic);
    }
    /**
     * One or more metadata key and value pairs as a dictionary of
     * strings.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> metadata;

    /**
     * @return One or more metadata key and value pairs as a dictionary of
     * strings.
     * 
     */
    public Output<Optional<Map<String,Object>>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * The name of the share. Changing this updates the name
     * of the existing share.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the share. Changing this updates the name
     * of the existing share.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The owner of the Share.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The owner of the Share.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share. Changing this
     * creates a new share.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share. Changing this
     * creates a new share.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The share replication type.
     * 
     */
    @Export(name="replicationType", refs={String.class}, tree="[0]")
    private Output<String> replicationType;

    /**
     * @return The share replication type.
     * 
     */
    public Output<String> replicationType() {
        return this.replicationType;
    }
    /**
     * The UUID of a share network where the share server exists
     * or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
     * the share_network_id value from the snapshot is used. Changing this creates a new share.
     * 
     */
    @Export(name="shareNetworkId", refs={String.class}, tree="[0]")
    private Output<String> shareNetworkId;

    /**
     * @return The UUID of a share network where the share server exists
     * or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
     * the share_network_id value from the snapshot is used. Changing this creates a new share.
     * 
     */
    public Output<String> shareNetworkId() {
        return this.shareNetworkId;
    }
    /**
     * The share protocol - can either be NFS, CIFS,
     * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
     * 
     */
    @Export(name="shareProto", refs={String.class}, tree="[0]")
    private Output<String> shareProto;

    /**
     * @return The share protocol - can either be NFS, CIFS,
     * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
     * 
     */
    public Output<String> shareProto() {
        return this.shareProto;
    }
    /**
     * The UUID of the share server.
     * 
     */
    @Export(name="shareServerId", refs={String.class}, tree="[0]")
    private Output<String> shareServerId;

    /**
     * @return The UUID of the share server.
     * 
     */
    public Output<String> shareServerId() {
        return this.shareServerId;
    }
    /**
     * The share type name. If you omit this parameter, the default
     * share type is used.
     * 
     */
    @Export(name="shareType", refs={String.class}, tree="[0]")
    private Output<String> shareType;

    /**
     * @return The share type name. If you omit this parameter, the default
     * share type is used.
     * 
     */
    public Output<String> shareType() {
        return this.shareType;
    }
    /**
     * The share size, in GBs. The requested share size cannot be greater
     * than the allowed GB quota. Changing this resizes the existing share.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return The share size, in GBs. The requested share size cannot be greater
     * than the allowed GB quota. Changing this resizes the existing share.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * The UUID of the share&#39;s base snapshot. Changing this creates
     * a new share.
     * 
     */
    @Export(name="snapshotId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snapshotId;

    /**
     * @return The UUID of the share&#39;s base snapshot. Changing this creates
     * a new share.
     * 
     */
    public Output<Optional<String>> snapshotId() {
        return Codegen.optional(this.snapshotId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Share(String name) {
        this(name, ShareArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Share(String name, ShareArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Share(String name, ShareArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:sharedfilesystem/share:Share", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private Share(String name, Output<String> id, @Nullable ShareState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:sharedfilesystem/share:Share", name, state, makeResourceOptions(options, id));
    }

    private static ShareArgs makeArgs(ShareArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ShareArgs.Empty : args;
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
    public static Share get(String name, Output<String> id, @Nullable ShareState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Share(name, id, state, options);
    }
}
