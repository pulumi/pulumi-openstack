import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V3 volume resource within OpenStack.
 */
export declare class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState): Volume;
    /**
     * If a volume is attached to an instance, this attribute will
     * display the Attachment ID, Instance ID, and the Device as the Instance
     * sees it.
     */
    readonly attachments: pulumi.Output<{
        device: string;
        id: string;
        instanceId: string;
    }[]>;
    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     */
    readonly availabilityZone: pulumi.Output<string>;
    /**
     * The consistency group to place the volume
     * in.
     */
    readonly consistencyGroupId: pulumi.Output<string | undefined>;
    /**
     * A description of the volume. Changing this updates
     * the volume's description.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * When this option is set it allows extending
     * attached volumes. Note: updating size of an attached volume requires Cinder
     * support for version 3.42 and a compatible storage driver.
     */
    readonly enableOnlineResize: pulumi.Output<boolean | undefined>;
    /**
     * The image ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly imageId: pulumi.Output<string | undefined>;
    /**
     * Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     */
    readonly metadata: pulumi.Output<{
        [key: string]: any;
    }>;
    /**
     * A unique name for the volume. Changing this updates the
     * volume's name.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The size of the volume to create (in gigabytes).
     */
    readonly size: pulumi.Output<number>;
    /**
     * The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly snapshotId: pulumi.Output<string | undefined>;
    /**
     * The volume ID to replicate with.
     */
    readonly sourceReplica: pulumi.Output<string | undefined>;
    /**
     * The volume ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly sourceVolId: pulumi.Output<string | undefined>;
    /**
     * The type of volume to create.
     * Changing this creates a new volume.
     */
    readonly volumeType: pulumi.Output<string>;
    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.ResourceOptions);
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
    readonly attachments?: pulumi.Input<{
        device?: pulumi.Input<string>;
        id?: pulumi.Input<string>;
        instanceId?: pulumi.Input<string>;
    }[]>;
    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The consistency group to place the volume
     * in.
     */
    readonly consistencyGroupId?: pulumi.Input<string>;
    /**
     * A description of the volume. Changing this updates
     * the volume's description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * When this option is set it allows extending
     * attached volumes. Note: updating size of an attached volume requires Cinder
     * support for version 3.42 and a compatible storage driver.
     */
    readonly enableOnlineResize?: pulumi.Input<boolean>;
    /**
     * The image ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly imageId?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     */
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A unique name for the volume. Changing this updates the
     * volume's name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The size of the volume to create (in gigabytes).
     */
    readonly size?: pulumi.Input<number>;
    /**
     * The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * The volume ID to replicate with.
     */
    readonly sourceReplica?: pulumi.Input<string>;
    /**
     * The volume ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly sourceVolId?: pulumi.Input<string>;
    /**
     * The type of volume to create.
     * Changing this creates a new volume.
     */
    readonly volumeType?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * The availability zone for the volume.
     * Changing this creates a new volume.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The consistency group to place the volume
     * in.
     */
    readonly consistencyGroupId?: pulumi.Input<string>;
    /**
     * A description of the volume. Changing this updates
     * the volume's description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * When this option is set it allows extending
     * attached volumes. Note: updating size of an attached volume requires Cinder
     * support for version 3.42 and a compatible storage driver.
     */
    readonly enableOnlineResize?: pulumi.Input<boolean>;
    /**
     * The image ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly imageId?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to associate with the volume.
     * Changing this updates the existing volume metadata.
     */
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A unique name for the volume. Changing this updates the
     * volume's name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The size of the volume to create (in gigabytes).
     */
    readonly size: pulumi.Input<number>;
    /**
     * The snapshot ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * The volume ID to replicate with.
     */
    readonly sourceReplica?: pulumi.Input<string>;
    /**
     * The volume ID from which to create the volume.
     * Changing this creates a new volume.
     */
    readonly sourceVolId?: pulumi.Input<string>;
    /**
     * The type of volume to create.
     * Changing this creates a new volume.
     */
    readonly volumeType?: pulumi.Input<string>;
}
