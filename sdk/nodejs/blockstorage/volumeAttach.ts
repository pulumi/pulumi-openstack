// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > **Note:** This resource usually requires admin privileges.
 *
 * > **Note:** This resource does not actually attach a volume to an instance.
 * Please use the `openstack.compute.VolumeAttach` resource for that.
 *
 * > **Note:** All arguments including the `data` computed attribute will be
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
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const volume1 = new openstack.blockstorage.Volume("volume_1", {
 *     name: "volume_1",
 *     size: 1,
 * });
 * const va1 = new openstack.blockstorage.VolumeAttach("va_1", {
 *     volumeId: volume1.id,
 *     device: "auto",
 *     hostName: "devstack",
 *     ipAddress: "192.168.255.10",
 *     initiator: "iqn.1993-08.org.debian:01:e9861fb1859",
 *     osType: "linux2",
 *     platform: "x86_64",
 * });
 * ```
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
    public /*out*/ readonly data!: pulumi.Output<{[key: string]: string}>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeAttachState | undefined;
            resourceInputs["attachMode"] = state ? state.attachMode : undefined;
            resourceInputs["data"] = state ? state.data : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["driverVolumeType"] = state ? state.driverVolumeType : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["initiator"] = state ? state.initiator : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["mountPointBase"] = state ? state.mountPointBase : undefined;
            resourceInputs["multipath"] = state ? state.multipath : undefined;
            resourceInputs["osType"] = state ? state.osType : undefined;
            resourceInputs["platform"] = state ? state.platform : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["volumeId"] = state ? state.volumeId : undefined;
            resourceInputs["wwnn"] = state ? state.wwnn : undefined;
            resourceInputs["wwpns"] = state ? state.wwpns : undefined;
        } else {
            const args = argsOrState as VolumeAttachArgs | undefined;
            if ((!args || args.hostName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostName'");
            }
            if ((!args || args.volumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeId'");
            }
            resourceInputs["attachMode"] = args ? args.attachMode : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["initiator"] = args ? args.initiator : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["multipath"] = args ? args.multipath : undefined;
            resourceInputs["osType"] = args ? args.osType : undefined;
            resourceInputs["platform"] = args ? args.platform : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
            resourceInputs["wwnn"] = args ? args.wwnn : undefined;
            resourceInputs["wwpns"] = args ? args.wwpns : undefined;
            resourceInputs["data"] = undefined /*out*/;
            resourceInputs["driverVolumeType"] = undefined /*out*/;
            resourceInputs["mountPointBase"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["data"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(VolumeAttach.__pulumiType, name, resourceInputs, opts);
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
    attachMode?: pulumi.Input<string>;
    /**
     * This is a map of key/value pairs that contain the connection
     * information. You will want to pass this information to a provisioner
     * script to finalize the connection. See below for more information.
     */
    data?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     */
    device?: pulumi.Input<string>;
    /**
     * The storage driver that the volume is based on.
     */
    driverVolumeType?: pulumi.Input<string>;
    /**
     * The host to attach the volume to.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The iSCSI initiator string to make the connection.
     */
    initiator?: pulumi.Input<string>;
    /**
     * The IP address of the `hostName` above.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * A mount point base name for shared storage.
     */
    mountPointBase?: pulumi.Input<string>;
    /**
     * Whether to connect to this volume via multipath.
     */
    multipath?: pulumi.Input<boolean>;
    /**
     * The iSCSI initiator OS type.
     */
    osType?: pulumi.Input<string>;
    /**
     * The iSCSI initiator platform.
     */
    platform?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    volumeId?: pulumi.Input<string>;
    /**
     * A wwnn name. Used for Fibre Channel connections.
     */
    wwnn?: pulumi.Input<string>;
    /**
     * An array of wwpn strings. Used for Fibre Channel
     * connections.
     */
    wwpns?: pulumi.Input<pulumi.Input<string>[]>;
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
    attachMode?: pulumi.Input<string>;
    /**
     * The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     */
    device?: pulumi.Input<string>;
    /**
     * The host to attach the volume to.
     */
    hostName: pulumi.Input<string>;
    /**
     * The iSCSI initiator string to make the connection.
     */
    initiator?: pulumi.Input<string>;
    /**
     * The IP address of the `hostName` above.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Whether to connect to this volume via multipath.
     */
    multipath?: pulumi.Input<boolean>;
    /**
     * The iSCSI initiator OS type.
     */
    osType?: pulumi.Input<string>;
    /**
     * The iSCSI initiator platform.
     */
    platform?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    volumeId: pulumi.Input<string>;
    /**
     * A wwnn name. Used for Fibre Channel connections.
     */
    wwnn?: pulumi.Input<string>;
    /**
     * An array of wwpn strings. Used for Fibre Channel
     * connections.
     */
    wwpns?: pulumi.Input<pulumi.Input<string>[]>;
}
