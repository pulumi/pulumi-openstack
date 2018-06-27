import * as pulumi from "@pulumi/pulumi";
/**
 * This resource is experimental and may be removed in the future! Feedback
 * is requested if you find this resource useful or if you find any problems
 * with it.
 *
 * Creates a general purpose attachment connection to a Block
 * Storage volume using the OpenStack Block Storage (Cinder) v2 API.
 * Depending on your Block Storage service configuration, this
 * resource can assist in attaching a volume to a non-OpenStack resource
 * such as a bare-metal server or a remote virtual machine in a
 * different cloud provider.
 *
 * This does not actually attach a volume to an instance. Please use
 * the `openstack_compute_volume_attach_v2` resource for that.
 */
export declare class VolumeAttachV2 extends pulumi.CustomResource {
    /**
     * Get an existing VolumeAttachV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeAttachV2State): VolumeAttachV2;
    /**
     * Specify whether to attach the volume as Read-Only
     * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
     * If left unspecified, the Block Storage API will apply a default of `rw`.
     */
    readonly attachMode: pulumi.Output<string | undefined>;
    /**
     * This is a map of key/value pairs that contain the connection
     * information. You will want to pass this information to a provisioner
     * script to finalize the connection. See below for more information.
     */
    readonly data: pulumi.Output<{
        [key: string]: any;
    }>;
    /**
     * The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     */
    readonly device: pulumi.Output<string | undefined>;
    /**
     * The storage driver that the volume is based on.
     */
    readonly driverVolumeType: pulumi.Output<string>;
    /**
     * The host to attach the volume to.
     */
    readonly hostName: pulumi.Output<string>;
    /**
     * The iSCSI initiator string to make the connection.
     */
    readonly initiator: pulumi.Output<string | undefined>;
    /**
     * The IP address of the `host_name` above.
     */
    readonly ipAddress: pulumi.Output<string | undefined>;
    /**
     * A mount point base name for shared storage.
     */
    readonly mountPointBase: pulumi.Output<string>;
    /**
     * Whether to connect to this volume via multipath.
     */
    readonly multipath: pulumi.Output<boolean | undefined>;
    /**
     * The iSCSI initiator OS type.
     */
    readonly osType: pulumi.Output<string | undefined>;
    /**
     * The iSCSI initiator platform.
     */
    readonly platform: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V2 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    readonly volumeId: pulumi.Output<string>;
    /**
     * A wwnn name. Used for Fibre Channel connections.
     */
    readonly wwnn: pulumi.Output<string | undefined>;
    /**
     * An array of wwpn strings. Used for Fibre Channel
     * connections.
     */
    readonly wwpns: pulumi.Output<string[] | undefined>;
    /**
     * Create a VolumeAttachV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeAttachV2Args, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering VolumeAttachV2 resources.
 */
export interface VolumeAttachV2State {
    /**
     * Specify whether to attach the volume as Read-Only
     * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
     * If left unspecified, the Block Storage API will apply a default of `rw`.
     */
    readonly attachMode?: pulumi.Input<string>;
    /**
     * This is a map of key/value pairs that contain the connection
     * information. You will want to pass this information to a provisioner
     * script to finalize the connection. See below for more information.
     */
    readonly data?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     */
    readonly device?: pulumi.Input<string>;
    /**
     * The storage driver that the volume is based on.
     */
    readonly driverVolumeType?: pulumi.Input<string>;
    /**
     * The host to attach the volume to.
     */
    readonly hostName?: pulumi.Input<string>;
    /**
     * The iSCSI initiator string to make the connection.
     */
    readonly initiator?: pulumi.Input<string>;
    /**
     * The IP address of the `host_name` above.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * A mount point base name for shared storage.
     */
    readonly mountPointBase?: pulumi.Input<string>;
    /**
     * Whether to connect to this volume via multipath.
     */
    readonly multipath?: pulumi.Input<boolean>;
    /**
     * The iSCSI initiator OS type.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The iSCSI initiator platform.
     */
    readonly platform?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    readonly volumeId?: pulumi.Input<string>;
    /**
     * A wwnn name. Used for Fibre Channel connections.
     */
    readonly wwnn?: pulumi.Input<string>;
    /**
     * An array of wwpn strings. Used for Fibre Channel
     * connections.
     */
    readonly wwpns?: pulumi.Input<pulumi.Input<string>[]>;
}
/**
 * The set of arguments for constructing a VolumeAttachV2 resource.
 */
export interface VolumeAttachV2Args {
    /**
     * Specify whether to attach the volume as Read-Only
     * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
     * If left unspecified, the Block Storage API will apply a default of `rw`.
     */
    readonly attachMode?: pulumi.Input<string>;
    /**
     * The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     */
    readonly device?: pulumi.Input<string>;
    /**
     * The host to attach the volume to.
     */
    readonly hostName: pulumi.Input<string>;
    /**
     * The iSCSI initiator string to make the connection.
     */
    readonly initiator?: pulumi.Input<string>;
    /**
     * The IP address of the `host_name` above.
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * Whether to connect to this volume via multipath.
     */
    readonly multipath?: pulumi.Input<boolean>;
    /**
     * The iSCSI initiator OS type.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The iSCSI initiator platform.
     */
    readonly platform?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    readonly volumeId: pulumi.Input<string>;
    /**
     * A wwnn name. Used for Fibre Channel connections.
     */
    readonly wwnn?: pulumi.Input<string>;
    /**
     * An array of wwpn strings. Used for Fibre Channel
     * connections.
     */
    readonly wwpns?: pulumi.Input<pulumi.Input<string>[]>;
}
