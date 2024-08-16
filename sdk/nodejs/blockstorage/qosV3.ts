// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 block storage Quality-Of-Servirce (qos) resource within OpenStack.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const qos = new openstack.blockstorage.QosV3("qos", {
 *     name: "foo",
 *     consumer: "back-end",
 *     specs: {
 *         read_iops_sec: "40000",
 *         write_iops_sec: "40000",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Qos can be imported using the `qos_id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:blockstorage/qosV3:QosV3 qos 941793f0-0a34-4bc4-b72e-a6326ae58283
 * ```
 */
export class QosV3 extends pulumi.CustomResource {
    /**
     * Get an existing QosV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QosV3State, opts?: pulumi.CustomResourceOptions): QosV3 {
        return new QosV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:blockstorage/qosV3:QosV3';

    /**
     * Returns true if the given object is an instance of QosV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QosV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QosV3.__pulumiType;
    }

    /**
     * The consumer of qos. Can be one of `front-end`,
     * `back-end` or `both`. Changing this updates the `consumer` of an
     * existing qos.
     */
    public readonly consumer!: pulumi.Output<string | undefined>;
    /**
     * Name of the qos.  Changing this creates a new qos.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the qos. If omitted,
     * the `region` argument of the provider is used. Changing this creates
     * a new qos.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Key/Value pairs of specs for the qos.
     */
    public readonly specs!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a QosV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: QosV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QosV3Args | QosV3State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QosV3State | undefined;
            resourceInputs["consumer"] = state ? state.consumer : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["specs"] = state ? state.specs : undefined;
        } else {
            const args = argsOrState as QosV3Args | undefined;
            resourceInputs["consumer"] = args ? args.consumer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["specs"] = args ? args.specs : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(QosV3.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QosV3 resources.
 */
export interface QosV3State {
    /**
     * The consumer of qos. Can be one of `front-end`,
     * `back-end` or `both`. Changing this updates the `consumer` of an
     * existing qos.
     */
    consumer?: pulumi.Input<string>;
    /**
     * Name of the qos.  Changing this creates a new qos.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the qos. If omitted,
     * the `region` argument of the provider is used. Changing this creates
     * a new qos.
     */
    region?: pulumi.Input<string>;
    /**
     * Key/Value pairs of specs for the qos.
     */
    specs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a QosV3 resource.
 */
export interface QosV3Args {
    /**
     * The consumer of qos. Can be one of `front-end`,
     * `back-end` or `both`. Changing this updates the `consumer` of an
     * existing qos.
     */
    consumer?: pulumi.Input<string>;
    /**
     * Name of the qos.  Changing this creates a new qos.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the qos. If omitted,
     * the `region` argument of the provider is used. Changing this creates
     * a new qos.
     */
    region?: pulumi.Input<string>;
    /**
     * Key/Value pairs of specs for the qos.
     */
    specs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
