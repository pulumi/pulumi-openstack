// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.InstanceArgs;
import com.pulumi.openstack.compute.inputs.InstanceState;
import com.pulumi.openstack.compute.outputs.InstanceBlockDevice;
import com.pulumi.openstack.compute.outputs.InstanceNetwork;
import com.pulumi.openstack.compute.outputs.InstancePersonality;
import com.pulumi.openstack.compute.outputs.InstanceSchedulerHint;
import com.pulumi.openstack.compute.outputs.InstanceVendorOptions;
import com.pulumi.openstack.compute.outputs.InstanceVolume;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="openstack:compute/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * The first detected Fixed IPv4 address.
     * 
     */
    @Export(name="accessIpV4", refs={String.class}, tree="[0]")
    private Output<String> accessIpV4;

    /**
     * @return The first detected Fixed IPv4 address.
     * 
     */
    public Output<String> accessIpV4() {
        return this.accessIpV4;
    }
    /**
     * The first detected Fixed IPv6 address.
     * 
     */
    @Export(name="accessIpV6", refs={String.class}, tree="[0]")
    private Output<String> accessIpV6;

    /**
     * @return The first detected Fixed IPv6 address.
     * 
     */
    public Output<String> accessIpV6() {
        return this.accessIpV6;
    }
    /**
     * The administrative password to assign to the server.
     * Changing this changes the root password on the existing server.
     * 
     */
    @Export(name="adminPass", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> adminPass;

    /**
     * @return The administrative password to assign to the server.
     * Changing this changes the root password on the existing server.
     * 
     */
    public Output<Optional<String>> adminPass() {
        return Codegen.optional(this.adminPass);
    }
    @Export(name="allMetadata", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> allMetadata;

    public Output<Map<String,Object>> allMetadata() {
        return this.allMetadata;
    }
    /**
     * The collection of tags assigned on the instance, which have
     * been explicitly and implicitly added.
     * 
     */
    @Export(name="allTags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allTags;

    /**
     * @return The collection of tags assigned on the instance, which have
     * been explicitly and implicitly added.
     * 
     */
    public Output<List<String>> allTags() {
        return this.allTags;
    }
    /**
     * The availability zone in which to create
     * the server. Conflicts with `availability_zone_hints`. Changing this creates
     * a new server.
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return The availability zone in which to create
     * the server. Conflicts with `availability_zone_hints`. Changing this creates
     * a new server.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * The availability zone in which to
     * create the server. This argument is preferred to `availability_zone`, when
     * scheduling the server on a
     * [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
     * host or node. Conflicts with `availability_zone`. Changing this creates a
     * new server.
     * 
     */
    @Export(name="availabilityZoneHints", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> availabilityZoneHints;

    /**
     * @return The availability zone in which to
     * create the server. This argument is preferred to `availability_zone`, when
     * scheduling the server on a
     * [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
     * host or node. Conflicts with `availability_zone`. Changing this creates a
     * new server.
     * 
     */
    public Output<Optional<String>> availabilityZoneHints() {
        return Codegen.optional(this.availabilityZoneHints);
    }
    /**
     * Configuration of block devices. The block_device
     * structure is documented below. Changing this creates a new server.
     * You can specify multiple block devices which will create an instance with
     * multiple disks. This configuration is very flexible, so please see the
     * following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
     * for more information.
     * 
     */
    @Export(name="blockDevices", refs={List.class,InstanceBlockDevice.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceBlockDevice>> blockDevices;

    /**
     * @return Configuration of block devices. The block_device
     * structure is documented below. Changing this creates a new server.
     * You can specify multiple block devices which will create an instance with
     * multiple disks. This configuration is very flexible, so please see the
     * following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
     * for more information.
     * 
     */
    public Output<Optional<List<InstanceBlockDevice>>> blockDevices() {
        return Codegen.optional(this.blockDevices);
    }
    /**
     * Whether to use the config_drive feature to
     * configure the instance. Changing this creates a new server.
     * 
     */
    @Export(name="configDrive", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> configDrive;

    /**
     * @return Whether to use the config_drive feature to
     * configure the instance. Changing this creates a new server.
     * 
     */
    public Output<Optional<Boolean>> configDrive() {
        return Codegen.optional(this.configDrive);
    }
    /**
     * The creation time of the instance.
     * 
     */
    @Export(name="created", refs={String.class}, tree="[0]")
    private Output<String> created;

    /**
     * @return The creation time of the instance.
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * The flavor ID of
     * the desired flavor for the server. Changing this resizes the existing server.
     * 
     */
    @Export(name="flavorId", refs={String.class}, tree="[0]")
    private Output<String> flavorId;

    /**
     * @return The flavor ID of
     * the desired flavor for the server. Changing this resizes the existing server.
     * 
     */
    public Output<String> flavorId() {
        return this.flavorId;
    }
    /**
     * The name of the
     * desired flavor for the server. Changing this resizes the existing server.
     * 
     */
    @Export(name="flavorName", refs={String.class}, tree="[0]")
    private Output<String> flavorName;

    /**
     * @return The name of the
     * desired flavor for the server. Changing this resizes the existing server.
     * 
     */
    public Output<String> flavorName() {
        return this.flavorName;
    }
    /**
     * @deprecated
     * Use the openstack_compute_floatingip_associate_v2 resource instead
     * 
     */
    @Deprecated /* Use the openstack_compute_floatingip_associate_v2 resource instead */
    @Export(name="floatingIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> floatingIp;

    public Output<Optional<String>> floatingIp() {
        return Codegen.optional(this.floatingIp);
    }
    /**
     * Whether to force the OpenStack instance to be
     * forcefully deleted. This is useful for environments that have reclaim / soft
     * deletion enabled.
     * 
     */
    @Export(name="forceDelete", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDelete;

    /**
     * @return Whether to force the OpenStack instance to be
     * forcefully deleted. This is useful for environments that have reclaim / soft
     * deletion enabled.
     * 
     */
    public Output<Optional<Boolean>> forceDelete() {
        return Codegen.optional(this.forceDelete);
    }
    /**
     * (Optional; Required if `image_name` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The image ID of
     * the desired image for the server. Changing this rebuilds the existing
     * server.
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output<String> imageId;

    /**
     * @return (Optional; Required if `image_name` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The image ID of
     * the desired image for the server. Changing this rebuilds the existing
     * server.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * (Optional; Required if `image_id` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The name of the
     * desired image for the server. Changing this rebuilds the existing server.
     * 
     */
    @Export(name="imageName", refs={String.class}, tree="[0]")
    private Output<String> imageName;

    /**
     * @return (Optional; Required if `image_id` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The name of the
     * desired image for the server. Changing this rebuilds the existing server.
     * 
     */
    public Output<String> imageName() {
        return this.imageName;
    }
    /**
     * The name of a key pair to put on the server. The key
     * pair must already be created and associated with the tenant&#39;s account.
     * Changing this creates a new server.
     * 
     */
    @Export(name="keyPair", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keyPair;

    /**
     * @return The name of a key pair to put on the server. The key
     * pair must already be created and associated with the tenant&#39;s account.
     * Changing this creates a new server.
     * 
     */
    public Output<Optional<String>> keyPair() {
        return Codegen.optional(this.keyPair);
    }
    /**
     * Metadata key/value pairs to make available from
     * within the instance. Changing this updates the existing server metadata.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> metadata;

    /**
     * @return Metadata key/value pairs to make available from
     * within the instance. Changing this updates the existing server metadata.
     * 
     */
    public Output<Optional<Map<String,Object>>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * A unique name for the resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Special string for `network` option to create
     * the server. `network_mode` can be `&#34;auto&#34;` or `&#34;none&#34;`.
     * Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
     * 
     */
    @Export(name="networkMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> networkMode;

    /**
     * @return Special string for `network` option to create
     * the server. `network_mode` can be `&#34;auto&#34;` or `&#34;none&#34;`.
     * Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
     * 
     */
    public Output<Optional<String>> networkMode() {
        return Codegen.optional(this.networkMode);
    }
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new server.
     * 
     */
    @Export(name="networks", refs={List.class,InstanceNetwork.class}, tree="[0,1]")
    private Output<List<InstanceNetwork>> networks;

    /**
     * @return An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new server.
     * 
     */
    public Output<List<InstanceNetwork>> networks() {
        return this.networks;
    }
    /**
     * Customize the personality of an instance by
     * defining one or more files and their contents. The personality structure
     * is described below. Changing this rebuilds the existing server.
     * 
     */
    @Export(name="personalities", refs={List.class,InstancePersonality.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstancePersonality>> personalities;

    /**
     * @return Customize the personality of an instance by
     * defining one or more files and their contents. The personality structure
     * is described below. Changing this rebuilds the existing server.
     * 
     */
    public Output<Optional<List<InstancePersonality>>> personalities() {
        return Codegen.optional(this.personalities);
    }
    /**
     * Provide the VM state. Only &#39;active&#39;, &#39;shutoff&#39;
     * and &#39;shelved_offloaded&#39; are supported values.
     * *Note*: If the initial power_state is the shutoff
     * the VM will be stopped immediately after build and the provisioners like
     * remote-exec or files are not supported.
     * 
     */
    @Export(name="powerState", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> powerState;

    /**
     * @return Provide the VM state. Only &#39;active&#39;, &#39;shutoff&#39;
     * and &#39;shelved_offloaded&#39; are supported values.
     * *Note*: If the initial power_state is the shutoff
     * the VM will be stopped immediately after build and the provisioners like
     * remote-exec or files are not supported.
     * 
     */
    public Output<Optional<String>> powerState() {
        return Codegen.optional(this.powerState);
    }
    /**
     * The region in which to create the server instance. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new server.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the server instance. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new server.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Provide the Nova scheduler with hints on how
     * the instance should be launched. The available hints are described below.
     * 
     */
    @Export(name="schedulerHints", refs={List.class,InstanceSchedulerHint.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceSchedulerHint>> schedulerHints;

    /**
     * @return Provide the Nova scheduler with hints on how
     * the instance should be launched. The available hints are described below.
     * 
     */
    public Output<Optional<List<InstanceSchedulerHint>>> schedulerHints() {
        return Codegen.optional(this.schedulerHints);
    }
    /**
     * An array of one or more security group names
     * to associate with the server. Changing this results in adding/removing
     * security groups from the existing server. *Note*: When attaching the
     * instance to networks using Ports, place the security groups on the Port
     * and not the instance. *Note*: Names should be used and not ids, as ids
     * trigger unnecessary updates.
     * 
     */
    @Export(name="securityGroups", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroups;

    /**
     * @return An array of one or more security group names
     * to associate with the server. Changing this results in adding/removing
     * security groups from the existing server. *Note*: When attaching the
     * instance to networks using Ports, place the security groups on the Port
     * and not the instance. *Note*: Names should be used and not ids, as ids
     * trigger unnecessary updates.
     * 
     */
    public Output<List<String>> securityGroups() {
        return this.securityGroups;
    }
    /**
     * Whether to try stop instance gracefully
     * before destroying it, thus giving chance for guest OS daemons to stop correctly.
     * If instance doesn&#39;t stop within timeout, it will be destroyed anyway.
     * 
     */
    @Export(name="stopBeforeDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stopBeforeDestroy;

    /**
     * @return Whether to try stop instance gracefully
     * before destroying it, thus giving chance for guest OS daemons to stop correctly.
     * If instance doesn&#39;t stop within timeout, it will be destroyed anyway.
     * 
     */
    public Output<Optional<Boolean>> stopBeforeDestroy() {
        return Codegen.optional(this.stopBeforeDestroy);
    }
    /**
     * A set of string tags for the instance. Changing this
     * updates the existing instance tags.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A set of string tags for the instance. Changing this
     * updates the existing instance tags.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The time when the instance was last updated.
     * 
     */
    @Export(name="updated", refs={String.class}, tree="[0]")
    private Output<String> updated;

    /**
     * @return The time when the instance was last updated.
     * 
     */
    public Output<String> updated() {
        return this.updated;
    }
    /**
     * The user data to provide when launching the instance.
     * Changing this creates a new server.
     * 
     */
    @Export(name="userData", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userData;

    /**
     * @return The user data to provide when launching the instance.
     * Changing this creates a new server.
     * 
     */
    public Output<Optional<String>> userData() {
        return Codegen.optional(this.userData);
    }
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     * 
     */
    @Export(name="vendorOptions", refs={InstanceVendorOptions.class}, tree="[0]")
    private Output</* @Nullable */ InstanceVendorOptions> vendorOptions;

    /**
     * @return Map of additional vendor-specific options.
     * Supported options are described below.
     * 
     */
    public Output<Optional<InstanceVendorOptions>> vendorOptions() {
        return Codegen.optional(this.vendorOptions);
    }
    /**
     * @deprecated
     * Use block_device or openstack_compute_volume_attach_v2 instead
     * 
     */
    @Deprecated /* Use block_device or openstack_compute_volume_attach_v2 instead */
    @Export(name="volumes", refs={List.class,InstanceVolume.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceVolume>> volumes;

    public Output<Optional<List<InstanceVolume>>> volumes() {
        return Codegen.optional(this.volumes);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, @Nullable InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, @Nullable InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/instance:Instance", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "adminPass"
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
