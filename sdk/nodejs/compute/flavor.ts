// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 flavor resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const test_flavor = new openstack.compute.Flavor("test-flavor", {
 *     disk: 20,
 *     extraSpecs: {
 *         "hw:cpu_policy": "CPU-POLICY",
 *         "hw:cpu_thread_policy": "CPU-THREAD-POLICY",
 *     },
 *     ram: 8096,
 *     vcpus: 2,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_flavor_v2.html.markdown.
 */
export class Flavor extends pulumi.CustomResource {
    /**
     * Get an existing Flavor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlavorState, opts?: pulumi.CustomResourceOptions): Flavor {
        return new Flavor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:compute/flavor:Flavor';

    /**
     * Returns true if the given object is an instance of Flavor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Flavor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Flavor.__pulumiType;
    }

    /**
     * The amount of disk space in gigabytes to use for the root
     * (/) partition. Changing this creates a new flavor.
     */
    public readonly disk!: pulumi.Output<number>;
    public readonly ephemeral!: pulumi.Output<number | undefined>;
    /**
     * Key/Value pairs of metadata for the flavor.
     */
    public readonly extraSpecs!: pulumi.Output<{[key: string]: any}>;
    /**
     * Whether the flavor is public. Changing this creates
     * a new flavor.
     */
    public readonly isPublic!: pulumi.Output<boolean | undefined>;
    /**
     * A unique name for the flavor. Changing this creates a new
     * flavor.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The amount of RAM to use, in megabytes. Changing this
     * creates a new flavor.
     */
    public readonly ram!: pulumi.Output<number>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Flavors are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * RX/TX bandwith factor. The default is 1. Changing
     * this creates a new flavor.
     */
    public readonly rxTxFactor!: pulumi.Output<number | undefined>;
    /**
     * The amount of disk space in megabytes to use. If
     * unspecified, the default is 0. Changing this creates a new flavor.
     */
    public readonly swap!: pulumi.Output<number | undefined>;
    /**
     * The number of virtual CPUs to use. Changing this creates
     * a new flavor.
     */
    public readonly vcpus!: pulumi.Output<number>;

    /**
     * Create a Flavor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlavorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlavorArgs | FlavorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FlavorState | undefined;
            inputs["disk"] = state ? state.disk : undefined;
            inputs["ephemeral"] = state ? state.ephemeral : undefined;
            inputs["extraSpecs"] = state ? state.extraSpecs : undefined;
            inputs["isPublic"] = state ? state.isPublic : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ram"] = state ? state.ram : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["rxTxFactor"] = state ? state.rxTxFactor : undefined;
            inputs["swap"] = state ? state.swap : undefined;
            inputs["vcpus"] = state ? state.vcpus : undefined;
        } else {
            const args = argsOrState as FlavorArgs | undefined;
            if (!args || args.disk === undefined) {
                throw new Error("Missing required property 'disk'");
            }
            if (!args || args.ram === undefined) {
                throw new Error("Missing required property 'ram'");
            }
            if (!args || args.vcpus === undefined) {
                throw new Error("Missing required property 'vcpus'");
            }
            inputs["disk"] = args ? args.disk : undefined;
            inputs["ephemeral"] = args ? args.ephemeral : undefined;
            inputs["extraSpecs"] = args ? args.extraSpecs : undefined;
            inputs["isPublic"] = args ? args.isPublic : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ram"] = args ? args.ram : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["rxTxFactor"] = args ? args.rxTxFactor : undefined;
            inputs["swap"] = args ? args.swap : undefined;
            inputs["vcpus"] = args ? args.vcpus : undefined;
        }
        super(Flavor.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Flavor resources.
 */
export interface FlavorState {
    /**
     * The amount of disk space in gigabytes to use for the root
     * (/) partition. Changing this creates a new flavor.
     */
    readonly disk?: pulumi.Input<number>;
    readonly ephemeral?: pulumi.Input<number>;
    /**
     * Key/Value pairs of metadata for the flavor.
     */
    readonly extraSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether the flavor is public. Changing this creates
     * a new flavor.
     */
    readonly isPublic?: pulumi.Input<boolean>;
    /**
     * A unique name for the flavor. Changing this creates a new
     * flavor.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The amount of RAM to use, in megabytes. Changing this
     * creates a new flavor.
     */
    readonly ram?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Flavors are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * RX/TX bandwith factor. The default is 1. Changing
     * this creates a new flavor.
     */
    readonly rxTxFactor?: pulumi.Input<number>;
    /**
     * The amount of disk space in megabytes to use. If
     * unspecified, the default is 0. Changing this creates a new flavor.
     */
    readonly swap?: pulumi.Input<number>;
    /**
     * The number of virtual CPUs to use. Changing this creates
     * a new flavor.
     */
    readonly vcpus?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Flavor resource.
 */
export interface FlavorArgs {
    /**
     * The amount of disk space in gigabytes to use for the root
     * (/) partition. Changing this creates a new flavor.
     */
    readonly disk: pulumi.Input<number>;
    readonly ephemeral?: pulumi.Input<number>;
    /**
     * Key/Value pairs of metadata for the flavor.
     */
    readonly extraSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether the flavor is public. Changing this creates
     * a new flavor.
     */
    readonly isPublic?: pulumi.Input<boolean>;
    /**
     * A unique name for the flavor. Changing this creates a new
     * flavor.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The amount of RAM to use, in megabytes. Changing this
     * creates a new flavor.
     */
    readonly ram: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Flavors are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * RX/TX bandwith factor. The default is 1. Changing
     * this creates a new flavor.
     */
    readonly rxTxFactor?: pulumi.Input<number>;
    /**
     * The amount of disk space in megabytes to use. If
     * unspecified, the default is 0. Changing this creates a new flavor.
     */
    readonly swap?: pulumi.Input<number>;
    /**
     * The number of virtual CPUs to use. Changing this creates
     * a new flavor.
     */
    readonly vcpus: pulumi.Input<number>;
}
