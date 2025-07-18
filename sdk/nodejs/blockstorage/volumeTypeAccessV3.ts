// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 block storage volume type access resource within OpenStack.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const project1 = new openstack.identity.Project("project_1", {name: "project_1"});
 * const volumeType1 = new openstack.blockstorage.VolumeTypeV3("volume_type_1", {
 *     name: "volume_type_1",
 *     isPublic: false,
 * });
 * const volumeTypeAccess = new openstack.blockstorage.VolumeTypeAccessV3("volume_type_access", {
 *     projectId: project1.id,
 *     volumeTypeId: volumeType1.id,
 * });
 * ```
 *
 * ## Import
 *
 * Volume types access can be imported using the `volume_type_id/project_id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3 volume_type_access 941793f0-0a34-4bc4-b72e-a6326ae58283/ed498e81f0cc448bae0ad4f8f21bf67f
 * ```
 */
export class VolumeTypeAccessV3 extends pulumi.CustomResource {
    /**
     * Get an existing VolumeTypeAccessV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeTypeAccessV3State, opts?: pulumi.CustomResourceOptions): VolumeTypeAccessV3 {
        return new VolumeTypeAccessV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3';

    /**
     * Returns true if the given object is an instance of VolumeTypeAccessV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeTypeAccessV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeTypeAccessV3.__pulumiType;
    }

    /**
     * ID of the project to give access to. Changing this
     * creates a new resource.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * ID of the volume type to give access to. Changing
     * this creates a new resource.
     */
    public readonly volumeTypeId!: pulumi.Output<string>;

    /**
     * Create a VolumeTypeAccessV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeTypeAccessV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeTypeAccessV3Args | VolumeTypeAccessV3State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeTypeAccessV3State | undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["volumeTypeId"] = state ? state.volumeTypeId : undefined;
        } else {
            const args = argsOrState as VolumeTypeAccessV3Args | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.volumeTypeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeTypeId'");
            }
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["volumeTypeId"] = args ? args.volumeTypeId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VolumeTypeAccessV3.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VolumeTypeAccessV3 resources.
 */
export interface VolumeTypeAccessV3State {
    /**
     * ID of the project to give access to. Changing this
     * creates a new resource.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    region?: pulumi.Input<string>;
    /**
     * ID of the volume type to give access to. Changing
     * this creates a new resource.
     */
    volumeTypeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VolumeTypeAccessV3 resource.
 */
export interface VolumeTypeAccessV3Args {
    /**
     * ID of the project to give access to. Changing this
     * creates a new resource.
     */
    projectId: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    region?: pulumi.Input<string>;
    /**
     * ID of the volume type to give access to. Changing
     * this creates a new resource.
     */
    volumeTypeId: pulumi.Input<string>;
}
