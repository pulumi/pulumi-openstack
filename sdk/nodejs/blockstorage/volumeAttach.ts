// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource is experimental and may be removed in the future! Feedback
 * is requested if you find this resource useful or if you find any problems
 * with it.
 *
 * Creates a general purpose attachment connection to a Block
 * Storage volume using the OpenStack Block Storage (Cinder) v3 API.
 * Depending on your Block Storage service configuration, this
 * resource can assist in attaching a volume to a non-OpenStack resource
 * such as a bare-metal server or a remote virtual machine in a
 * different cloud provider.
 *
 * This does not actually attach a volume to an instance. Please use
 * the `openstack.compute.VolumeAttach` resource for that.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const volume1 = new openstack.blockstorage.Volume("volume_1", {
 *     size: 1,
 * });
 * const va1 = new openstack.blockstorage.VolumeAttach("va_1", {
 *     device: "auto",
 *     hostName: "devstack",
 *     initiator: "iqn.1993-08.org.debian:01:e9861fb1859",
 *     ipAddress: "192.168.255.10",
 *     osType: "linux2",
 *     platform: "x86_64",
 *     volumeId: volume1.id,
 * });
 * ```
 */
export class VolumeAttach extends pulumi.CustomResource {
    /**
     * Get an existing VolumeAttach resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeAttachState, opts?: pulumi.CustomResourceOptions): VolumeAttach {
        return new VolumeAttach(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:blockstorage/volumeAttach:VolumeAttach';

    /**
     * Returns true if the given object is an instance of VolumeAttach.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeAttach {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeAttach.__pulumiType;
    }

    /**
     * Specify whether to attach the volume as Read-Only
     * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
     * If left unspecified, the Block Storage API will apply a default of `rw`.
     */
    public readonly attachMode!: pulumi.Output<string | undefined>;
    /**
     * This is a map of key/value pairs that contain the connection
     * information. You will want to pass this information to a provisioner
     * script to finalize the connection. See below for more information.
     */
    public /*out*/ readonly data!: pulumi.Output<{[key: string]: any}>;
    /**
     * The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * The storage driver that the volume is based on.
     */
    public /*out*/ readonly driverVolumeType!: pulumi.Output<string>;
    /**
     * The host to attach the volume to.
     */
    public readonly hostName!: pulumi.Output<string>;
    /**
     * The iSCSI initiator string to make the connection.
     */
    public readonly initiator!: pulumi.Output<string | undefined>;
    /**
     * The IP address of the `hostName` above.
     */
    public readonly ipAddress!: pulumi.Output<string | undefined>;
    /**
     * A mount point base name for shared storage.
     */
    public /*out*/ readonly mountPointBase!: pulumi.Output<string>;
    /**
     * Whether to connect to this volume via multipath.
     */
    public readonly multipath!: pulumi.Output<boolean | undefined>;
    /**
     * The iSCSI initiator OS type.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * The iSCSI initiator platform.
     */
    public readonly platform!: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V3 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    public readonly volumeId!: pulumi.Output<string>;
    /**
     * A wwnn name. Used for Fibre Channel connections.
     */
    public readonly wwnn!: pulumi.Output<string | undefined>;
    /**
     * An array of wwpn strings. Used for Fibre Channel
     * connections.
     */
    public readonly wwpns!: pulumi.Output<string[] | undefined>;

    /**
     * Create a VolumeAttach resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeAttachArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeAttachArgs | VolumeAttachState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VolumeAttachState | undefined;
            inputs["attachMode"] = state ? state.attachMode : undefined;
            inputs["data"] = state ? state.data : undefined;
            inputs["device"] = state ? state.device : undefined;
            inputs["driverVolumeType"] = state ? state.driverVolumeType : undefined;
            inputs["hostName"] = state ? state.hostName : undefined;
            inputs["initiator"] = state ? state.initiator : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["mountPointBase"] = state ? state.mountPointBase : undefined;
            inputs["multipath"] = state ? state.multipath : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["platform"] = state ? state.platform : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["volumeId"] = state ? state.volumeId : undefined;
            inputs["wwnn"] = state ? state.wwnn : undefined;
            inputs["wwpns"] = state ? state.wwpns : undefined;
        } else {
            const args = argsOrState as VolumeAttachArgs | undefined;
            if (!args || args.hostName === undefined) {
                throw new Error("Missing required property 'hostName'");
            }
            if (!args || args.volumeId === undefined) {
                throw new Error("Missing required property 'volumeId'");
            }
            inputs["attachMode"] = args ? args.attachMode : undefined;
            inputs["device"] = args ? args.device : undefined;
            inputs["hostName"] = args ? args.hostName : undefined;
            inputs["initiator"] = args ? args.initiator : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["multipath"] = args ? args.multipath : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["platform"] = args ? args.platform : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["volumeId"] = args ? args.volumeId : undefined;
            inputs["wwnn"] = args ? args.wwnn : undefined;
            inputs["wwpns"] = args ? args.wwpns : undefined;
            inputs["data"] = undefined /*out*/;
            inputs["driverVolumeType"] = undefined /*out*/;
            inputs["mountPointBase"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VolumeAttach.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VolumeAttach resources.
 */
export interface VolumeAttachState {
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
    readonly data?: pulumi.Input<{[key: string]: any}>;
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
     * The IP address of the `hostName` above.
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
     * The region in which to obtain the V3 Block Storage
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
 * The set of arguments for constructing a VolumeAttach resource.
 */
export interface VolumeAttachArgs {
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
     * The IP address of the `hostName` above.
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
     * The region in which to obtain the V3 Block Storage
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
