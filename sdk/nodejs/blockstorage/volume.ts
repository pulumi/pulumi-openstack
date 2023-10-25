// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V3 volume resource within OpenStack.
 *
 * ## Import
 *
 * Volumes can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:blockstorage/volume:Volume volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:blockstorage/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * If a volume is attached to an instance, this attribute will
     * display the Attachment ID, Instance ID, and the Device as the Instance
     * sees it.
     */
    public /*out*/ readonly attachments!: pulumi.Output<outputs.blockstorage.VolumeAttachment[]>;
    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The consistency group to place the volume
     * in.
     */
    public readonly consistencyGroupId!: pulumi.Output<string | undefined>;
    /**
     * A description of the volume. Changing this updates
     * the volume's description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When this option is set it allows extending
     * attached volumes. Note: updating size of an attached volume requires Cinder
     * support for version 3.42 and a compatible storage driver.
     */
    public readonly enableOnlineResize!: pulumi.Output<boolean | undefined>;
    /**
     * The image ID from which to create the volume.
     * Changing this creates a new volume.
     */
    public readonly imageId!: pulumi.Output<string | undefined>;
    /**
     * Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * Allow the volume to be attached to more than one Compute instance.
     */
    public readonly multiattach!: pulumi.Output<boolean | undefined>;
    /**
     * A unique name for the volume. Changing this updates the
     * volume's name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Provide the Cinder scheduler with hints on where
     * to instantiate a volume in the OpenStack cloud. The available hints are described below.
     */
    public readonly schedulerHints!: pulumi.Output<outputs.blockstorage.VolumeSchedulerHint[] | undefined>;
    /**
     * The size of the volume to create (in gigabytes).
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * The volume ID to replicate with.
     */
    public readonly sourceReplica!: pulumi.Output<string | undefined>;
    /**
     * The volume ID from which to create the volume.
     * Changing this creates a new volume.
     */
    public readonly sourceVolId!: pulumi.Output<string | undefined>;
    /**
     * The type of volume to create.
     * Changing this creates a new volume.
     */
    public readonly volumeType!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            resourceInputs["attachments"] = state ? state.attachments : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["consistencyGroupId"] = state ? state.consistencyGroupId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableOnlineResize"] = state ? state.enableOnlineResize : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["multiattach"] = state ? state.multiattach : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["schedulerHints"] = state ? state.schedulerHints : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
            resourceInputs["sourceReplica"] = state ? state.sourceReplica : undefined;
            resourceInputs["sourceVolId"] = state ? state.sourceVolId : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["consistencyGroupId"] = args ? args.consistencyGroupId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableOnlineResize"] = args ? args.enableOnlineResize : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["multiattach"] = args ? args.multiattach : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["schedulerHints"] = args ? args.schedulerHints : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["sourceReplica"] = args ? args.sourceReplica : undefined;
            resourceInputs["sourceVolId"] = args ? args.sourceVolId : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["attachments"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * If a volume is attached to an instance, this attribute will
     * display the Attachment ID, Instance ID, and the Device as the Instance
     * sees it.
     */
    attachments?: pulumi.Input<pulumi.Input<inputs.blockstorage.VolumeAttachment>[]>;
    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The consistency group to place the volume
     * in.
     */
    consistencyGroupId?: pulumi.Input<string>;
    /**
     * A description of the volume. Changing this updates
     * the volume's description.
     */
    description?: pulumi.Input<string>;
    /**
     * When this option is set it allows extending
     * attached volumes. Note: updating size of an attached volume requires Cinder
     * support for version 3.42 and a compatible storage driver.
     */
    enableOnlineResize?: pulumi.Input<boolean>;
    /**
     * The image ID from which to create the volume.
     * Changing this creates a new volume.
     */
    imageId?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Allow the volume to be attached to more than one Compute instance.
     */
    multiattach?: pulumi.Input<boolean>;
    /**
     * A unique name for the volume. Changing this updates the
     * volume's name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     */
    region?: pulumi.Input<string>;
    /**
     * Provide the Cinder scheduler with hints on where
     * to instantiate a volume in the OpenStack cloud. The available hints are described below.
     */
    schedulerHints?: pulumi.Input<pulumi.Input<inputs.blockstorage.VolumeSchedulerHint>[]>;
    /**
     * The size of the volume to create (in gigabytes).
     */
    size?: pulumi.Input<number>;
    /**
     * The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The volume ID to replicate with.
     */
    sourceReplica?: pulumi.Input<string>;
    /**
     * The volume ID from which to create the volume.
     * Changing this creates a new volume.
     */
    sourceVolId?: pulumi.Input<string>;
    /**
     * The type of volume to create.
     * Changing this creates a new volume.
     */
    volumeType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The consistency group to place the volume
     * in.
     */
    consistencyGroupId?: pulumi.Input<string>;
    /**
     * A description of the volume. Changing this updates
     * the volume's description.
     */
    description?: pulumi.Input<string>;
    /**
     * When this option is set it allows extending
     * attached volumes. Note: updating size of an attached volume requires Cinder
     * support for version 3.42 and a compatible storage driver.
     */
    enableOnlineResize?: pulumi.Input<boolean>;
    /**
     * The image ID from which to create the volume.
     * Changing this creates a new volume.
     */
    imageId?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Allow the volume to be attached to more than one Compute instance.
     */
    multiattach?: pulumi.Input<boolean>;
    /**
     * A unique name for the volume. Changing this updates the
     * volume's name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     */
    region?: pulumi.Input<string>;
    /**
     * Provide the Cinder scheduler with hints on where
     * to instantiate a volume in the OpenStack cloud. The available hints are described below.
     */
    schedulerHints?: pulumi.Input<pulumi.Input<inputs.blockstorage.VolumeSchedulerHint>[]>;
    /**
     * The size of the volume to create (in gigabytes).
     */
    size: pulumi.Input<number>;
    /**
     * The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * The volume ID to replicate with.
     */
    sourceReplica?: pulumi.Input<string>;
    /**
     * The volume ID from which to create the volume.
     * Changing this creates a new volume.
     */
    sourceVolId?: pulumi.Input<string>;
    /**
     * The type of volume to create.
     * Changing this creates a new volume.
     */
    volumeType?: pulumi.Input<string>;
}
