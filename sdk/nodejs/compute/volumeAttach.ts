// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Attaches a Block Storage Volume to an Instance using the OpenStack
 * Compute (Nova) v2 API.
 *
 * ## Example Usage
 * ### Basic attachment of a single volume to a single instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const volume1 = new openstack.blockstorage.VolumeV2("volume1", {size: 1});
 * const instance1 = new openstack.compute.Instance("instance1", {securityGroups: ["default"]});
 * const va1 = new openstack.compute.VolumeAttach("va1", {
 *     instanceId: instance1.id,
 *     volumeId: volume1.id,
 * });
 * ```
 * ### Using Multiattach-enabled volumes
 *
 * Multiattach Volumes are dependent upon your OpenStack cloud and not all
 * clouds support multiattach.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const volume1 = new openstack.blockstorage.Volume("volume1", {
 *     multiattach: true,
 *     size: 1,
 * });
 * const instance1 = new openstack.compute.Instance("instance1", {securityGroups: ["default"]});
 * const instance2 = new openstack.compute.Instance("instance2", {securityGroups: ["default"]});
 * const va1 = new openstack.compute.VolumeAttach("va1", {
 *     instanceId: instance1.id,
 *     multiattach: true,
 *     volumeId: openstack_blockstorage_volume_v2.volume_1.id,
 * });
 * const va2 = new openstack.compute.VolumeAttach("va2", {
 *     instanceId: instance2.id,
 *     multiattach: true,
 *     volumeId: openstack_blockstorage_volume_v2.volume_1.id,
 * }, {
 *     dependsOn: ["openstack_compute_volume_attach_v2.va_1"],
 * });
 * ```
 *
 * It is recommended to use `dependsOn` for the attach resources
 * to enforce the volume attachments to happen one at a time.
 *
 * ## Import
 *
 * Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:compute/volumeAttach:VolumeAttach va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
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
    public static readonly __pulumiType = 'openstack:compute/volumeAttach:VolumeAttach';

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
     * See Argument Reference above. _NOTE_: The correctness of this
     * information is dependent upon the hypervisor in use. In some cases, this
     * should not be used as an authoritative piece of information.
     */
    public readonly device!: pulumi.Output<string>;
    /**
     * The ID of the Instance to attach the Volume to.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Enable attachment of multiattach-capable volumes.
     */
    public readonly multiattach!: pulumi.Output<boolean | undefined>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    public readonly vendorOptions!: pulumi.Output<outputs.compute.VolumeAttachVendorOptions | undefined>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    public readonly volumeId!: pulumi.Output<string>;

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
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["multiattach"] = state ? state.multiattach : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["vendorOptions"] = state ? state.vendorOptions : undefined;
            resourceInputs["volumeId"] = state ? state.volumeId : undefined;
        } else {
            const args = argsOrState as VolumeAttachArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.volumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeId'");
            }
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["multiattach"] = args ? args.multiattach : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["vendorOptions"] = args ? args.vendorOptions : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VolumeAttach.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VolumeAttach resources.
 */
export interface VolumeAttachState {
    /**
     * See Argument Reference above. _NOTE_: The correctness of this
     * information is dependent upon the hypervisor in use. In some cases, this
     * should not be used as an authoritative piece of information.
     */
    device?: pulumi.Input<string>;
    /**
     * The ID of the Instance to attach the Volume to.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Enable attachment of multiattach-capable volumes.
     */
    multiattach?: pulumi.Input<boolean>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     */
    region?: pulumi.Input<string>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    vendorOptions?: pulumi.Input<inputs.compute.VolumeAttachVendorOptions>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    volumeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VolumeAttach resource.
 */
export interface VolumeAttachArgs {
    /**
     * See Argument Reference above. _NOTE_: The correctness of this
     * information is dependent upon the hypervisor in use. In some cases, this
     * should not be used as an authoritative piece of information.
     */
    device?: pulumi.Input<string>;
    /**
     * The ID of the Instance to attach the Volume to.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Enable attachment of multiattach-capable volumes.
     */
    multiattach?: pulumi.Input<boolean>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     */
    region?: pulumi.Input<string>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    vendorOptions?: pulumi.Input<inputs.compute.VolumeAttachVendorOptions>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    volumeId: pulumi.Input<string>;
}
