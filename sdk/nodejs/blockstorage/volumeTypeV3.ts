// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 block storage volume type resource within OpenStack.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const volumeType1 = new openstack.blockstorage.VolumeTypeV3("volume_type_1", {
 *     description: "Volume type 1",
 *     extraSpecs: {
 *         capabilities: "gpu",
 *         volume_backend_name: "ssd",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Volume types can be imported using the `volume_type_id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:blockstorage/volumeTypeV3:VolumeTypeV3 volume_type_1 941793f0-0a34-4bc4-b72e-a6326ae58283
 * ```
 */
export class VolumeTypeV3 extends pulumi.CustomResource {
    /**
     * Get an existing VolumeTypeV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeTypeV3State, opts?: pulumi.CustomResourceOptions): VolumeTypeV3 {
        return new VolumeTypeV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:blockstorage/volumeTypeV3:VolumeTypeV3';

    /**
     * Returns true if the given object is an instance of VolumeTypeV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeTypeV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeTypeV3.__pulumiType;
    }

    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing volume type.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Key/Value pairs of metadata for the volume type.
     */
    public readonly extraSpecs!: pulumi.Output<{[key: string]: any}>;
    /**
     * Whether the volume type is public. Changing
     * this updates the `isPublic` of an existing volume type.
     */
    public readonly isPublic!: pulumi.Output<boolean>;
    /**
     * Name of the volume type.  Changing this
     * updates the `name` of an existing volume type.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a VolumeTypeV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VolumeTypeV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeTypeV3Args | VolumeTypeV3State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeTypeV3State | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["extraSpecs"] = state ? state.extraSpecs : undefined;
            inputs["isPublic"] = state ? state.isPublic : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as VolumeTypeV3Args | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["extraSpecs"] = args ? args.extraSpecs : undefined;
            inputs["isPublic"] = args ? args.isPublic : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VolumeTypeV3.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VolumeTypeV3 resources.
 */
export interface VolumeTypeV3State {
    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing volume type.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Key/Value pairs of metadata for the volume type.
     */
    readonly extraSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether the volume type is public. Changing
     * this updates the `isPublic` of an existing volume type.
     */
    readonly isPublic?: pulumi.Input<boolean>;
    /**
     * Name of the volume type.  Changing this
     * updates the `name` of an existing volume type.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VolumeTypeV3 resource.
 */
export interface VolumeTypeV3Args {
    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing volume type.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Key/Value pairs of metadata for the volume type.
     */
    readonly extraSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether the volume type is public. Changing
     * this updates the `isPublic` of an existing volume type.
     */
    readonly isPublic?: pulumi.Input<boolean>;
    /**
     * Name of the volume type.  Changing this
     * updates the `name` of an existing volume type.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    readonly region?: pulumi.Input<string>;
}