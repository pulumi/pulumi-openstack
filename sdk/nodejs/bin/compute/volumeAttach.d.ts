import * as pulumi from "@pulumi/pulumi";
/**
 * Attaches a Block Storage Volume to an Instance using the OpenStack
 * Compute (Nova) v2 API.
 */
export declare class VolumeAttach extends pulumi.CustomResource {
    /**
     * Get an existing VolumeAttach resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeAttachState): VolumeAttach;
    /**
     * The device of the volume attachment (ex: `/dev/vdc`).
     * _NOTE_: Being able to specify a device is dependent upon the hypervisor in
     * use. There is a chance that the device specified in Terraform will not be
     * the same device the hypervisor chose. If this happens, Terraform will wish
     * to update the device upon subsequent applying which will cause the volume
     * to be detached and reattached indefinitely. Please use with caution.
     */
    readonly device: pulumi.Output<string>;
    /**
     * The ID of the Instance to attach the Volume to.
     */
    readonly instanceId: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    readonly volumeId: pulumi.Output<string>;
    /**
     * Create a VolumeAttach resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeAttachArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering VolumeAttach resources.
 */
export interface VolumeAttachState {
    /**
     * The device of the volume attachment (ex: `/dev/vdc`).
     * _NOTE_: Being able to specify a device is dependent upon the hypervisor in
     * use. There is a chance that the device specified in Terraform will not be
     * the same device the hypervisor chose. If this happens, Terraform will wish
     * to update the device upon subsequent applying which will cause the volume
     * to be detached and reattached indefinitely. Please use with caution.
     */
    readonly device?: pulumi.Input<string>;
    /**
     * The ID of the Instance to attach the Volume to.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    readonly volumeId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a VolumeAttach resource.
 */
export interface VolumeAttachArgs {
    /**
     * The device of the volume attachment (ex: `/dev/vdc`).
     * _NOTE_: Being able to specify a device is dependent upon the hypervisor in
     * use. There is a chance that the device specified in Terraform will not be
     * the same device the hypervisor chose. If this happens, Terraform will wish
     * to update the device upon subsequent applying which will cause the volume
     * to be detached and reattached indefinitely. Please use with caution.
     */
    readonly device?: pulumi.Input<string>;
    /**
     * The ID of the Instance to attach the Volume to.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a volume attachment. If omitted, the
     * `region` argument of the provider is used. Changing this creates a
     * new volume attachment.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The ID of the Volume to attach to an Instance.
     */
    readonly volumeId: pulumi.Input<string>;
}
